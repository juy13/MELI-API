#!/bin/bash

# Generate Swagger documentation for the Go API
swag init --dir package/server/api -o docs/swagger --parseDependency --parseInternal
