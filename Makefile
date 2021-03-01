go-image := golang:1.16.0-buster
go-dir := /go/src

run_rooms:
	docker container run --rm -it -w $(go-dir) -v ${CURDIR}/rooms:$(go-dir) --name rooms $(go-image) go run main.go

build_rooms:
	docker container run --rm -it -w $(go-dir) -v ${CURDIR}/rooms:$(go-dir) --name rooms $(go-image) go build -v -o rooms.app
