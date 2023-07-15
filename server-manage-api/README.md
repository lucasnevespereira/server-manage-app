# server-manage-api

Simple REST API to manage servers

API provides the following endpoints:

- Create a server
- List servers
- Get a server by ID
- Delete a server by ID

A server has the following properties:

- Unique ID
- Name
- Type (small, medium, large)
- Status (starting, running, stopping, stopped)

## Usage

#### Prerequisites

Before running the API, make sure you have Docker and Docker Compose installed on your machine.

### Running the API

```bash
docker-compose up
```

or

```bash
make run
```

This command will build the necessary Docker images and start the API container.

Wait for the API to start. You will see log messages indicating the API is running.

Once the API is up and running, you can access it at http://localhost:9000

#### Stopping the API

To stop the API, press `Ctrl + C` in the terminal where the API is running. This will gracefully stop the API container.

You can also run `docker-compose down`


### API Documentation

Once the API is running, you can access the API documentation at http://localhost:9000/swagger/index.html.


### Running Unit Tests

To execute tests for the API, follow these steps:

Run the following command:

```shell
make test
```

## API Calls and Responses

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
```