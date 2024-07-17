## Quick Start

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### Building the Application

1. Build the Docker image:

    ```sh
    make build
    ```

3. Run application container:

    ```sh
    make run
    ```
## API Endpoints

- **GET /**: Welcome message.
- **GET /health**: Health check endpoint.