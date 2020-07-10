all: pull rrefresh

.PHONY: build
build:
	docker build -t tbcom-1 .

.PHONY: clean
clean:
	docker rm tbc; docker rmi tbcom-1

.PHONY: drun
drun:
	docker run -t -it -d --name tbc -p 8080:8080 tbcom-1

.PHONY: pull
pull:
	git pull origin master

.PHONY: run
run:
	go run main.go

.PHONY: rrefresh
rrefresh: pull stop clean build drun

.PHONY: stop
stop:
	docker stop tbc
