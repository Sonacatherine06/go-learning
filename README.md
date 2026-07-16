# Go Learning

This repository contains beginner-friendly Go (Golang) programs created while learning the Go programming language.

## Programs

- Hello World
- Calculator
- Factorial
- Fibonacci Series
- Prime Number Checker
- Reverse String
- Palindrome Checker
- Array Sum
- Student Struct

## Technologies Used

- Go (Golang)
- Docker

## How to Run Using Go

```bash
go run helloworld.go
```

Replace `helloworld.go` with the name of any Go program you want to run.

## How to Run Using Docker

Build the Docker image:

```bash
docker build --build-arg FILE=helloworld.go -t hello-app .
```

Run the Docker container:

```bash
docker run --rm hello-app
```

To run another program, replace `helloworld.go` with any of the following:

- calculator.go
- factorial.go
- fibonacci.go
- palindrome.go
- primenumber_check.go
- reverse_string.go
- student.go
- sum_of_array.go
