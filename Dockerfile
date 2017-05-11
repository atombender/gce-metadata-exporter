FROM alpine:3.5
RUN apk update && apk add dumb-init su-exec
COPY build/gce-metadata-exporter /usr/bin/
ENTRYPOINT ["dumb-init", "su-exec", "nobody", "gce-metadata-exporter"]
