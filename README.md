# GCE Metadata Exporter

This is an exporter for [Prometheus](https://prometheus.io/) that collects metadata from Google Compute Engine.

# Installation

## Docker

```
$ docker run atombender/gce-metadata-exporter:latest
```

## From source

This requires [dep](https://github.com/golang/dep), GNU Make and Go >= 1.7.

```shell
$ dep ensure -vendor-only
$ make build
```

# Usage

```shell
$ gce-metadata-exporter
```

# License

MIT license. See `LICENSE` file.
