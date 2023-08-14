BINARY_NAME=steganographer
OUTPUT=bin
 
build:
	@echo "Building..."
	go build -o ${OUTPUT}/${BINARY_NAME} main.go
 
run: build
	@echo "Running..."
	./${OUTPUT}/${BINARY_NAME}
 
clean:
	@echo "Cleaning..."
	go clean
	rm -rfd ./${OUTPUT}


