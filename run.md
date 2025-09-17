# TLDR

1. Run redis:

```sh
docker compose -f deployment/docker-compose-services.yml up -d
```

2. Build and run the API:

```sh
go mod tidy
go build -o meli-api
```

3. Start the API with the following command:

```sh
./meli-api -c configs/config.yaml
```

3.1. Or run the API using Docker Compose:

```sh
docker compose -f docker-compose.yml up --build
```