run:
	go run main.go

rrefresh:
	docker stop tbc && docker rm tbc && docker rmi tbcom-1 . && docker build -t tbcom-1 . && docker run -it -d --name tbc -p 8080:8080 tbcom-1
