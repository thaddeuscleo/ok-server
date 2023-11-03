# Ok-Server

## Description
Ok-Server is a simple Go application designed to respond with "ok" when accessed at the root route. Its primary purpose is to test TCP connections and server functionality.

## Purpose
The main purpose of the Ok-Server project is to serve as a basic testing tool for TCP connections. By running this server, you can verify that your server environment is operational and can handle incoming HTTP requests.

## Getting Started
Follow these instructions to get the Ok-Server up and running on your system.

### Prerequisites
- [Go](https://golang.org/doc/install) should be installed on your system.

### Installation
1. Clone this repository to your local machine:

    ```sh
    git clone https://github.com/your-username/ok-server.git
    cd ok-server
    ```

2. Build the project:

    ```sh
    make run
    ```

### Usage
1. Run the Ok-Server with a specified port number:

    ```sh
    ./ok-server 8080
    ```

    Replace `8080` with your desired port number.

2. Once the server is running, open a web browser or use a tool like `curl` to make a request to the server:

    ```sh
    curl http://localhost:8080
    ```

    You should receive a response with "ok."

3. The server logs will display "Request Accepted" each time a request is made.