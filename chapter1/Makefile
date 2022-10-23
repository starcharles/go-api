BINARY_NAME=main
 
start:
	air

build:
	go mod tidy
	go build -o ${BINARY_NAME} main.go
 
clean:
	go clean
	rm ${BINARY_NAME}