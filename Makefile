fNameSrc = main
fNameOut = main

test:
	go test ./src/...

t: test

build: test
	go build -o bin/${fNameSrc}.exe src/${fNameOut}.go

b: build

run:
	go run src/${fNameSrc}.go

r: run

execute:
	./bin/${fNameOut}

e: execute