.PHONY: create-build-dir run build clean

run: build/edulsp
	./build/edulsp

create-build-dir:
	mkdir -p build

build: create-build-dir
	go build -o build/edulsp src/main.go

clean:
	rm -rf log.txt build
