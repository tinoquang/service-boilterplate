APP:=service-boilerplate

build:
	go build -v -o ${APP} .

run: build
	./${APP}

go-mod:
	@go mod tidy -v && go mod vendor
