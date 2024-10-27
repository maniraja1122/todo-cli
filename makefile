all: build run install clean

build: main.go
	go build main.go
	mkdir -p bin
	cp -f main bin/todo
	mv -f main bin/td
	echo "Built Successful, Make sure to move the bin/todo and bin/td file to your device's bin folder"

run:
	./bin/todo

install:
	sudo mv -f bin/todo /usr/local/bin
	sudo mv -f bin/td /usr/local/bin

clean:
	rm -rf bin
	rm -f main
