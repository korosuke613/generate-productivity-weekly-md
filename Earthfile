# Earthfile

FROM golang:1.16
WORKDIR /work
ENV GO111MODULE=on

src:
    COPY ./cmd ./cmd
    COPY ./lib ./lib
    COPY ./main.go ./
    COPY go.mod go.sum ./

deps:
    FROM +src

    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

lint:
    FROM +src
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.38.0
    RUN golangci-lint run


build:
    FROM +deps

    RUN go build -o build/tempura ./main.go
    SAVE ARTIFACT build/tempura /tempura AS LOCAL build/tempura

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

docker:
    FROM gcr.io/distroless/base

    COPY +build/tempura /tempura
    ENTRYPOINT ["/tempura"]
    SAVE IMAGE tempura:latest
