ARG GO_VERSION=1.11
FROM golang:${GO_VERSION}-alpine AS build
RUN apk add --update curl git libc-dev gcc
RUN mkdir -p /mnt/build
WORKDIR /mnt/build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o ./gce-metadata-exporter ./...

FROM golang:${GO_VERSION}-alpine
RUN apk update && apk add dumb-init su-exec
COPY --from=build /mnt/build/gce-metadata-exporter /usr/local/bin/
ENTRYPOINT ["dumb-init", "su-exec", "nobody", "gce-metadata-exporter"]
