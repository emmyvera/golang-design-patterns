package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Opitions     *VideoOptions
	Encoder      Processor
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}

	fmt.Println("vd.NewVideoNew(): Video Created:", id, input)

	return Video{
		ID:           id,
		InputFile:    input,
		OutputDir:    output,
		EncodingType: encType,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Opitions:     ops,
	}
}

func (v *Video) encode() {
	var fileName string

	switch v.EncodingType {
	case "mp4":
		//encode the video
		fmt.Print("v.encode(): About to encode to mp4", v.ID)
		name, err := v.encodeToMP4()
		if err != nil {
			// send info to the notifyChane
			v.sendToNotifyChan(false, "", fmt.Sprintf("Encode failed for %d: %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.mp4", name)

	default:
		fmt.Println("v.encode(): Error tying to encode to mp4", v.ID)
		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d: invalid encoding type", v.ID))
		return

	}
	fmt.Println("v.encode(): Sending sucess message to for video id", v.ID, "to nitifyChan")
	v.sendToNotifyChan(true, fileName, fmt.Sprintf("video id: %d processed and saved as %s", v.ID, fmt.Sprintf("%s/%s", v.OutputDir, fileName)))
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := ""
	fmt.Println("v.encodeToMP4(): About to try to encode to mp4", v.ID)
	if !v.Opitions.RenameOutput {
		// Get the baseFile name
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		// Generate random name
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)

	if err != nil {
		return "", err
	}
	fmt.Println("v.encodeToMP4(): Sucessfully encoded Video id", v.ID)
	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	fmt.Println("v.sendToNotifyChan(): Sending message to notifyChan for video id", v.ID)
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Message:    message,
		OutputFile: fileName,
	}
}

func New(jobQuene chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	fmt.Println("New(): Creating worker pool")
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	// TODO: impl processor logic
	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		WorkerPool: workerPool,
		jobQueue:   jobQuene,
		maxWorkers: maxWorkers,
		Processor:  p,
	}
}
