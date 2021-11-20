# Folder Loctions
SRC_DIR 	:= ./src
BLD_DIR 	:= ./bin

compile:
	go build -o $(BLD_DIR)/main $(SRC_DIR)/main.go

release:
	$(BLD_DIR)/main

run:
	go run $(SRC_DIR)/main.go

# Clean The Build Folder To Allow For A Complete Rebuild
clean:
	rm -f $(BLD_DIR)/*
	go mod tidy

keys:
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

test:
	go build -o $(BLD_DIR)/main $(SRC_DIR)/main.go
	$(BLD_DIR)/main

configtest:
	go build -o $(BLD_DIR)/main $(SRC_DIR)/main.go
	$(BLD_DIR)/main -c

help:
	go build -o $(BLD_DIR)/main $(SRC_DIR)/main.go
	$(BLD_DIR)/main -h

setup:
	mkdir bin