# Server Management Application

This repository contains the front-end and API code for the Server Management Application.

## Getting Started

To run the application locally, follow the steps below.

### Prerequisites

- [Docker](https://www.docker.com/get-started/) must be installed to run the API.
- [pnpm](https://pnpm.io/) must be installed to run the front-end.

### Front-End

To launch the front-end, use the following command:

```shell
make run-front
```

This command will perform the following steps:

- Navigate to the server-manage-front directory.
- Install the dependencies using `pnpm install`.
- Run the development server with `pnpm run dev`.

### Back-End (API)

To launch the front-end, use the following command:

```shell
make run-api
```

This command will perform the following steps:

- Navigate to the server-manage-api directory.
- Start the API using `docker-compose up`.

  The API will be available at http://localhost:9000.


