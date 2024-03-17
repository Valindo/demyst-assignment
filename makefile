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

# Clean removes build artifacts from the folder
# Generally used during development
# Usage : $ make clearn
clean:
	rm -rf build/