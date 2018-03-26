all:
	go run src/*.go
build:
	$$(cd src && go build -o ../bin/bekit)
