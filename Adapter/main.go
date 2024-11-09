package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id" xml:"id"`
	Name     string `json:"name" xml:"name"`
	Company  string `json:"company" xml:"company"`
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
}

type DataInterface interface {
	GetData() (*User, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*User, error) {
	return rs.Remote.GetData()
}

type JSONBackend struct {
}

func (jb *JSONBackend) GetData() (*User, error) {
	resp, err := http.Get("https://json-placeholder.mock.beeceptor.com/users/1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type XMLBackend struct {
}

func (xb *XMLBackend) GetData() (*User, error) {
	xmlFile := `
	<?xml version="1.0" encoding="UTF-8" ?>
	<root>
		<id>1</id>
		<name>Emily Johnson</name>
		<company>1</company>
		<username>1</username>
		<email>1</email>
	</root>
	`

	var user User
	_ = xml.Unmarshal([]byte(xmlFile), &user)

	return &user, nil
}

func main() {
	user := getRemoteData()
	fmt.Println("User without adapter: \t", user.ID, user.Name)

	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	userFromJSON, _ := jsonAdapter.CallRemoteService()
	fmt.Println("User with JSON adapter: \t", userFromJSON.ID, userFromJSON.Name)

	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	userFromXML, _ := xmlAdapter.CallRemoteService()
	fmt.Println("User with XML adapter: \t", userFromXML.ID, userFromXML.Name)

}

 func getRemoteData() *User {
	resp, err := http.Get("https://json-placeholder.mock.beeceptor.com/users/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	return &user

}
