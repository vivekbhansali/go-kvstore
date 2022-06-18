.PHONY: build clean run tidy

NAME=app

build:
	go build -o ${NAME} main.go

clean:
	rm -f ${NAME}

run: build
	./app

tidy:
	go mod tidy