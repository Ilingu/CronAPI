APP_BINARY=cronApi.exe

build:
	@echo Building App binary...
	set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ./${APP_BINARY} ./cmd/api 
	@echo Done!