GOSOURCES := $(wildcard */*.go)

all: build

build: $(GOSOURCES) go.mod
	GOOS=linux go build -o build/cnt-manager.linux
	GOOS=darwin go build -o build/cnt-manager.osx

test-osx: build
	./build/cnt-manager.osx

install: build
	cp -v ./build/cnt-manager.linux /usr/local/sbin/cnt-manager