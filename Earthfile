# Earthfile

FROM golang:1.16
WORKDIR /work
ENV GO111MODULE=on

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

    COPY ./cmd ./cmd
    COPY ./lib ./lib
    COPY ./main.go ./

lint:
    FROM +deps
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.38.0
    RUN golangci-lint run


binary:
    FROM +deps

    COPY .git ./.git
    RUN export VERSION=$(git tag | tail -1) \
     && go build -o build/tempura \
        -ldflags "-s -w -X \"github.com/korosuke613/tempura/cmd.version=$VERSION\" " \
        ./main.go
    RUN ./build/tempura -v
    SAVE ARTIFACT build/tempura /tempura AS LOCAL build/tempura


docker:
    FROM gcr.io/distroless/base

    COPY +binary/tempura /tempura
    ENTRYPOINT ["/tempura"]
    SAVE IMAGE tempura:latest

build:
    BUILD +binary
    BUILD +docker

test:
    FROM +deps

    COPY examples ./examples
    RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./...
    SAVE ARTIFACT coverage.txt AS LOCAL build/coverage.txt

goreleaser-setup:
    FROM +deps
    RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
    RUN mv ./bin/goreleaser /usr/local/bin/goreleaser
    COPY . /work
    RUN git reset --hard HEAD


release-dryrun:
    FROM +goreleaser-setup
    RUN goreleaser --snapshot --skip-publish --rm-dist
    SAVE ARTIFACT dist AS LOCAL build/dist

release:
    FROM +goreleaser-setup
    RUN --push --secret GITHUB_TOKEN=+secrets/GITHUB_TOKEN goreleaser release


all:
    BUILD +build
    BUILD +test
    BUILD +lint
    BUILD +docker
    BUILD +release-dryrun
