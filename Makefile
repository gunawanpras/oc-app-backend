.PHONY: run docker-build docker-run docker-build-run

run:
	@cp .env.development .env
	@go build .
	@go run .

docker-build:
	@docker build . --file Dockerfile --build-arg ENVIRONMENT=development --no-cache --tag oc-app-backend

docker-run:
	@docker run --name=oc-app-backend -p 3000:3000 -d oc-app-backend:latest
	@docker ps

docker-build-run: docker-build docker-run