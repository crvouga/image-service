.PHONY: docker-up docker-down docker

# Docker image and container names
IMAGE_NAME := image-service
CONTAINER_NAME := image-service
PORT := 8080
CONTAINER_PORT := 80

# Build and run the Docker container
docker-up:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .
	@echo "Removing existing container if it exists..."
	-docker stop $(CONTAINER_NAME) 2>/dev/null || true
	-docker rm $(CONTAINER_NAME) 2>/dev/null || true
	@echo "Starting container..."
	docker run -d --name $(CONTAINER_NAME) -p $(PORT):$(CONTAINER_PORT) $(IMAGE_NAME)
	@echo "Container is running on http://localhost:$(PORT)"

# Stop and remove the Docker container
docker-down:
	@echo "Stopping container..."
	-docker stop $(CONTAINER_NAME) 2>/dev/null || true
	@echo "Removing container..."
	-docker rm $(CONTAINER_NAME) 2>/dev/null || true
	@echo "Container stopped and removed"

# Run docker-down then docker-up
docker: docker-down docker-up
	@echo "Docker container restarted"
