# Application make file

# Compile application in build/ folder
# usage : $ make build
build:
	go build -o build/

# Run the compiled build
# Usage : $ make run_build
run_prod:
	./build/demyst-assignment

# Run development build
# Usage : $ make run_dev
run_dev:
	go run main.go

# Run tests
# Usage : $ make test
test:
	go test ./...

# Run tests verbose
# Usage : $ make test_verbose
test_verbose:
	go test ./... -v

# Docker build image
# Usage : $ make docker_build
docker_build:
	docker build -t valindo_assignment -f ./Dockerfile .

# Docker build test image
# usage : $ make docker_build_test
docker_build_test:
	docker build -t valindo_assignment_test -f Dockerfile.testing .

# Docker run image
# usage : $ make docker_run
docker_run:
	docker run valindo_assignment

# Docker run test image
# usage : $ make docker_run_test
docker_run_test:
	docker run valindo_assignment_test

# Clean removes build artifacts from the folder
# Generally used during development
# Usage : $ make clean
clean:
	rm -rf build/