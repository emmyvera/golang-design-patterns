# Golang Design Patterns

This repository contains my implementations of common design patterns in Golang, exploring how each pattern can be used effectively in real-world applications. Through these exercises, I deepened my understanding of how to structure, optimize, and scale code using patterns that enhance maintainability and flexibility. I also made this for my personal code snippet :-)

## Overview

This repo covers several key design patterns, each implemented in Go with clear examples and explanations:

- **Factory Pattern**: Encapsulates object creation, providing an interface for creating objects without specifying the exact class of object.
- **Abstract Factory Pattern**: Extends the Factory pattern by creating factories that generate families of related objects.
- **Repository Pattern**: Provides a data access layer that abstracts interactions with data sources, promoting separation of concerns.
- **Singleton Pattern**: Ensures a class has only one instance and provides a global point of access to that instance.
- **Builder Pattern**: Separates the construction of a complex object from its representation, allowing for more flexible creation.
- **Adapter Pattern**: Allows incompatible interfaces to work together by acting as a bridge between them.
- **Worker Pool Pattern**: Manages a pool of goroutines to handle concurrent tasks, controlling resource usage effectively.

Although the **Repository Pattern** and **Singleton Pattern** are missing here because I did used them in a real life project  

## Prerequisites

To try out these patterns, you'll need:

- [Golang](https://golang.org/dl/) version 1.16 or higher

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/golang-design-patterns.git
   cd golang-design-patterns

## Key Takeaways
Learning these patterns provided insights into:

Structuring code to enhance reusability and readability
Handling complex object creation in a simplified manner
Managing concurrency efficiently with Go's goroutines and channels

Acknowledgments
Thanks to Trevor Sawler (Ph.D.) creator of Streamline your development by learning how common design patterns are implement in Go on Udemy for providing a solid foundation in Golang design patterns and best practices.