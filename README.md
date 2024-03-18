# Coding Assignment

## Contents
1. [Description](#Description)
2. [Setup and prerequisite](#Setup-and-prerequisite)
3. [Testing](#Testing)
4. [Docker commands](#Docker-commands)
5. [Notes](#Notes)

## Description
The following application fetch's todos from https://jsonplaceholder.typicode.com/todos, the requirement was to fetch the first 20 even todos and display them to indicate if they are completed or not

## Setup and prerequisite
Please make sure you have `>=go:1.21` setup locally if you want to run this without docker

Clone the repository first
```
$ git clone 
```

Install any dependencies
```
$ go mod download
```

Setup env variables
```
$ cp .env.example .env
```

Build application
```
$ make build
```

Run Application locally
```
$ make run_dev
``` 

## Testing
Run test normally
```
$ make test
```
Run test with verbose flag
```
$ make test_verbose
```

## Docker commands
Docker build application
```
$ make docker_build
```

Docker run application
```
$ make docker_run
```

Docker build test
```
$ make docker_build_test
```

Docker run tests
```
$ make docker_run_test
```

## Notes
- `[X]` and `[ ]` are used to indicate of the todo is completed or not respectively
- Displaying the Todos internally uses `FetchTodos` and `generateUrl` , this was done so that incase we decide to introduce an API that would retrive the Todos, it would not chnage the functionality much
- The above point opens room for extenstions when we need to mark the todo as completed or not

