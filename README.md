# GCE Metadata Exporter

This is an exporter for [Prometheus](https://prometheus.io/) that collects metadata from Google Compute Engine.

# Installation

## Docker

```
$ docker run atombender/gce-metadata-exporter:latest
```

## From source

This requires [Glide](https://glide.sh/) and Go >= 1.7.

```shell
$ mkdir -p $GOPATH/src/github.com/atombender
$ cd $GOPATH/src/github.com/atombender
$ git clone https://github.com/atombender/gce-metadata-exporter
$ cd gce-metadata-exporter
$ glide install --strip-vendor
$ go build -o gce-metadata-exporter *.go
```

# Usage

```shell
$ gce-metadata-exporter
```

# License

MIT license. See `LICENSE` file.
