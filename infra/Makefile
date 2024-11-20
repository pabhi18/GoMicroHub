FRONT_END = frontApp
BROKER_BINARY = brokerApp
Auth_Binary = authApp

.PHONY: up up_build down build_broker build_auth build_frontend run

## up : starts all containers in the background without forcing build
up: 
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started"

## up_build : stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_frontend
	@echo "Stopping docker-compose images (if running...)"
	docker-compose down
	@echo "Building (when required ) and starting docker images..."
	docker-compose up --build -d 
	@echo "Docker images built and started"

## stop docker compose 
down: 
	@echo "Stopping docker-compose images"
	docker-compose down
	@echo "Done!"

## build_frontend: builds the frontend binary as a Linux executable
build_frontend:
	@echo "Building frontend binary as a macOS executable..."
	cd ../front-end && go build -o ${FRONT_END} ./cmd/web
	@echo "Build Done!"


## build_broker: builds the broker binary as a Linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the broker binary as a Linux executable
build_auth:
	@echo "Building auth binary..." 
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${Auth_Binary} ./cmd/api
	@echo "Done!"

## run: runs the frontend binary
run:
	@echo "Running frontend binary..."
	cd ../front-end && ./${FRONT_END}
	@echo "Running Done!"
