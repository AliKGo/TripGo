build:
	go build -o ride-hail-system .

format:
	gofumpt -l -w .

up:
	docker-compose up --build -d

down:
	docker-compose down 

nuke:
	docker-compose down -v
