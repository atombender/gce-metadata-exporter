FROM alpine:3.5 AS build
ENV GOPATH=/go PATH=$PATH:/go/bin
ARG DEP_VERSION=0.3.2
RUN \
     mkdir -p /go/bin \
  && apk update \
  && apk add go curl git dumb-init su-exec libc-dev \
  && curl -fL# -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 \
  && chmod +x /usr/local/bin/dep \
  && mkdir -p /go/src/github.com/atombender/gce-metadata-exporter
WORKDIR /go/src/github.com/atombender/gce-metadata-exporter
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only -v
COPY . ./
RUN go build -o /gce-metadata-exporter github.com/atombender/gce-metadata-exporter

FROM alpine:3.5
RUN apk update && apk add dumb-init su-exec
COPY --from=build /gce-metadata-exporter /usr/local/bin/
ENTRYPOINT ["dumb-init", "su-exec", "nobody", "gce-metadata-exporter"]
