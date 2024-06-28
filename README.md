# Chain of Responsibility Pattern in Go

## Overview
This repository demonstrates the application of the Chain of Responsibility design pattern in Go. The project highlights how to create a chain of handler objects to process requests, showcasing flexibility and best practices in design patterns and unit testing.

## Pattern Description
The Chain of Responsibility pattern is used to pass a request along a chain of handlers. Each handler decides either to process the request or to pass it to the next handler in the chain. This pattern is useful for scenarios such as event handling systems, middleware, and logging.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and usage of handlers.
- **pkg/chain/**: Houses the Chain of Responsibility implementation.
  - **handler.go**: Defines the `Handler` interface.
  - **base_handler.go**: Provides a base implementation of the `Handler` interface.
  - **concrete_handler_a.go**: Implements `ConcreteHandlerA` to handle specific requests.
  - **concrete_handler_b.go**: Implements `ConcreteHandlerB` to handle specific requests.
  - **base_handler_test.go**: Unit tests for `BaseHandler`.
  - **concrete_handler_a_test.go**: Unit tests for `ConcreteHandlerA`.
  - **concrete_handler_b_test.go**: Unit tests for `ConcreteHandlerB`.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Chain-Of-Responsibility_Pattern.git
cd Go_Chain-Of-Responsibility_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp