APP:=service-boilerplate
RELEASE_TIME  := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
RELEASE_VER   := $(if $(CI_COMMIT_ID),$(CI_COMMIT_ID),$(shell git rev-parse --short HEAD))
LDFLAGS       := "-s -w -X \"main.ldBuildDate=${RELEASE_TIME}\" -X main.ldGitCommit=${RELEASE_VER}"

build:
	@go build -v -ldflags ${LDFLAGS} -o ${APP} main.go

run: build
	# print version
	@./${APP} -v

	# run binary
	@./${APP}

go-mod:
	@go mod tidy -v && go mod vendor
