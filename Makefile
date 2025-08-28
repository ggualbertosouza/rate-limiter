BINARY_NAME=api
OUT_DIR=bin/${BINARY_NAME}

build-image: clean
	echo "Starting app image build"
	go build -o ${OUT_DIR} src/cmd/main.go

run:
	@go run src/cmd/main.go

clean:
	echo "Limpando ambiente"
	rm -rf bin/

.PHONY=run-dev