APP_BINARY=cronApi.exe

build:
	@echo Building App binary...
	set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ./bin/api/${APP_BINARY} ./cmd/api 
	@echo Done!
build_docker:
	@echo Building Docker App...
	docker build -t ilingu/cron-api .
	@echo Done!
start: build_docker
	@echo Starting application...
	docker run -it --rm -p 3001:3001 ilingu/cron-api
	@echo Done!
push: build_docker
	@echo Pushing Application...
	docker push ilingu/cron-api
	@echo Done!
