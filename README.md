# MELI Item detail 

- [More details about the project](docs/readme/architecture.md)
- [Used prompts](docs/readme/prompts.md)
- [Spanish version](docs/readme/esp/spanish.md)

## Description

This project is a simple API to get item details from Mercado Libre. It uses Go and Redis for caching.

## Building

To build the project, you need to have Go installed on your machine. Then, you can run the following commands:
```sh
go mod tidy
go build -o meli-api
```

## Configuration

The project uses a `config.yaml` file for configuration. You can modify this file to change the API port and other settings. The config file is located at `configs/config.yaml`.

## Running

To run the project, you need to have Redis running on your machine:

```sh
docker compose -f deployment/docker-compose-services.yml up -d
```

Be sure that Redis is running before starting the API. And you have to redefine the environment variable `REDIS_PASSWORD` in the file `deployment/docker-compose-services.yml` with your password. Adding the device path `/var/data/redis` to the volume driver_opts in the docker-compose-services.yml file allows the data to persist across container restarts.

And then start the API with the following command:

```sh
./meli-api -c configs/config.yaml
```

## Database

The project uses a JSON file as a database. The path to the database file is specified in the `config.yaml` file under the `database` section.

## Metrics

The project uses Prometheus for metrics collection. The metrics server runs on port 9090 by default. You can change this port in the `config.yaml` file under the `metrics` section. 

Additionally there are custom metrics:
- http\_request\_duration\_seconds
- http\_requests\_total

They can be found bt visiting `/metrics` endpoint of the API server.

## Testing

The project includes unit tests for the API endpoints. You can run the tests using the following command:
```sh
go test ./...
```

## Deployment

There are docker compose files provided in `deployment`. You can use them to deploy the application locally or on a production environment.

`data-storage` is a volume that is used to store data. It is mounted at `/data/storage` inside the container.Outside it takes a folder with JSON files as a database. 

To run the application using docker compose, navigate to the `deployment` directory and run:

```sh
docker compose -f docker-compose.yml up --build
```

## CI/CD

The project uses GitHub Actions for continuous integration and delivery. The workflow is triggered on every push. It lints and tests the application. If there is a release with a tag, it builds docker containers and pushes them to github packages. File to found: [.github/workflows/ci.yml](.github/workflows/ci.yml)

## Swagger

The project includes Swagger documentation for API endpoints. You can access it by visiting `/swagger/index.html` endpoint of the API server.