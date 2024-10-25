all: build run install clean

build: main.go
	go build main.go
	mkdir -p bin
	mv -f main bin/todo
	echo "Built Successful, Make sure to move the bin/todo file to your device's bin folder"

run:
	./bin/todo

install:
	sudo mv -f bin/todo /usr/local/bin

clean:
	rm -rf bin
	rm -f main
