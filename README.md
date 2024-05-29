# Totality Assignment - gRPC Service

## Project Description

This project is a gRPC service for managing various functionalities as part of the Totality assignment. It is built using Go and includes various modules and configurations to support its operations.

## Features

- gRPC server implementation
- Configuration management
- Database integration (here just a simulation in form of a variable)
- Modular code structure for scalability
- Logging functionality

## Installation

To install and set up the project, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/toastsandwich/totality-assignment-gRPC-service/
    ```

2. Navigate to the project directory:
    ```sh
    cd totality-assignment-gRPC-servive
    ```

3. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Usage

To start the gRPC server, run the following command:

```sh
make run
```

## Docker
To start gRPC server on docker

1. Pull the image:
    ```docker
    docker pull toastsandwich/totality-grpc-service:latest
    ```

2. Get docker image id
    ```docker
    docker images
    ```

3. Get docker container id
    ```docker
    docker ps
    ```

4. Run the docker image
    ```docker
    docker run -p<localport>:<7000> <docker image id>
    ```
5. you can now use postman to test services
