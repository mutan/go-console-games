go-image := golang:1.16.0-buster
go-dir := /go/src
docker := docker container run --rm -it
workdir := -w $(go-dir)
volume := -v ${CURDIR}/rooms:$(go-dir)

run_rooms:
	$(docker) $(workdir) $(volume) $(go-image) go run main.go

build_rooms:
	$(docker) $(workdir) $(volume) $(go-image) go build -v -o rooms.app
