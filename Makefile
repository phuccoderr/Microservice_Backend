up:
	docker-compose up -d
down:
	docker-compose down
up_build:
	docker-compose up --build
build_discovery:
	chdir \discovery-server && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0