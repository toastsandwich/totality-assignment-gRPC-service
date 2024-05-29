# 🌐 Totality Assignment - gRPC Service

## 📜 Project Description

This project is a gRPC service for managing various functionalities as part of the Totality assignment. It is built using Go and includes various modules and configurations to support its operations.

## ✨ Features

- 🚀 gRPC server implementation
- ⚙️ Configuration management
- 💾 Database integration (simulation in the form of a variable)
- 📦 Modular code structure for scalability
- 📄 Logging functionality

## 📦 Installation

To install and set up the project, follow these steps:

1. **Clone the repository:**
    ```sh
    git clone https://github.com/toastsandwich/totality-assignment-gRPC-service/
    ```

2. **Navigate to the project directory:**
    ```sh
    cd totality-assignment-gRPC-service
    ```

3. **Install the dependencies:**
    ```sh
    go mod tidy
    ```

## 🚀 Usage

To start the gRPC server, run the following command:

```sh
make run
```
## 🐳 Docker

To start the gRPC server using Docker, follow these steps:

1. **Pull the image:**
    ```sh
    docker pull toastsandwich/totality-grpc-service:latest
    ```

2. **Get Docker image ID:**
    ```sh
    docker images
    ```

3. **Get Docker container ID:**
    ```sh
    docker ps
    ```

4. **Run the Docker image:**
    ```sh
    docker run -p <localport>:7000 <docker image id>
    ```

5. **Test services using Postman**
    ```sh
    grpc://localhost:<localport> -> service
    ```
---
