# Server Management Application

This repository contains the front-end and API code for a simple Server Management Application.

## Usage

⚠️ For the purpose of the exercise I've committed `.env` files

To run the application locally, follow the steps below.

### Prerequisites

- [Docker](https://www.docker.com/get-started/) must be installed to run the API.
- [node](https://nodejs.org/fr/download) version 18 or higher must be installed to run frontend.
- [pnpm](https://pnpm.io/) must be installed to run the front-end.

## Front-End

To launch the front-end, use the following command:

```shell
make run-front
```

This command will perform the following steps:

- Navigate to the server-manage-front directory.
- Install the dependencies using `pnpm install`.
- Run the development server with `pnpm run dev`.

  The frontend will be available at http://localhost:5173/

## Back-End (API)

To launch the api, use the following command:

```shell
make run-api
```

This command will perform the following steps:

- Navigate to the server-manage-api directory.
- Start the API using `docker-compose up`.

  The API will be available at http://localhost:9000.

#### Stopping the API

To stop the API, press Ctrl + C in the terminal where the API is running.
You can also run `make stop-api` that will:

- navigate to server-manage-api directory
- run `docker-compose down`

### Running Unit Tests

To execute tests for the API, follow these steps:

Run the following command:

```shell
make test-api

```

This command will perform the following steps:

- Navigate to the server-manage-api directory.
- Run the unit tests by executing make test.

### API Documentation

Once the API is running, you can access the API documentation at http://localhost:9000/swagger/index.html.

### API Endpoints

Local Endpoint: `http://localhost:9000`

The following endpoints are available in the app:

### Health Check

- Method: GET
- Path: `/health`
- Description: Returns the status of the api.

**Example Response:**

```json
{
  "status": "up"
}
```

### List Servers

- Method: GET
- Path: `/servers`
- Description: Retrieves a list of all servers.

**Example Response:**

```json
{
  "servers": [
    {
      "id": "a5cb11fd-03cd-410d-94df-06fb328f6573",
      "name": "Atlas",
      "type": "small",
      "status": "stopped"
    },
    {
      "id": "7bf55aeb-d0c1-4f29-a6e2-fe0a8854cb34",
      "name": "Cyprus",
      "type": "medium",
      "status": "running"
    }
  ],
  "total": 2
}
```

### Get Server

- Method: GET
- Path: `/servers/:id`
- Description: Retrieves information about a server by its ID.

**Example Response:**

```json
{
  "id": "7bf55aeb-d0c1-4f29-a6e2-fe0a8854cb34",
  "name": "Cyprus",
  "type": "medium",
  "status": "running"
}
```

### Create a Server

- Method: POST
- Path: `/servers`
- Description: Creates a new server with the provided details.

**Example Request Body:**

```json
{
  "Name": "Orla",
  "Type": "small",
  "Status": "running"
}
```

A server type must be either small, medium or large.
A server status must be either starting, running, stopping, stopped.

**Example Response**

```json
{
  "message": "Server Orla was created with id 7951302e-6690-46f8-aeb1-9488e11a483f",
  "status": 200
}
```

### Delete Server

- Method: DELETE
- Path: `/servers/:id`
- Description: Deletes a server by its ID.

**Example Response:**

```json
{
  "message": "Server with id a5cb11fd-03cd-410d-94df-06fb328f6573 was deleted",
  "status": 200
}


