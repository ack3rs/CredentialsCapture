FROM golang:alpine as credentialscapture 
RUN apk update && apk add --no-cache git inotify-tools jq

COPY . /src
RUN cd /src && GIT_COMMIT=$(git rev-list -1 HEAD) && VERSION=$(git describe --tags --abbrev=0) && go build -o /go/credentialscapture -ldflags "-X main.GitCommit=$GIT_COMMIT" -ldflags "-X main.VERSION=$VERSION"

COPY Docker/docker-entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

