# GCE Metadata Exporter

This is an exporter for [Prometheus](https://prometheus.io/) that collects metadata from Google Compute Engine.

# Installation

## Kubernetes

You can use the [Helm](https://helm.sh) chart under `chart`:

```shell
$ helm install ./helm/gce-metadata-exporter
```

## Docker

```
$ docker run atombender/gce-metadata-exporter:0.2.0
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
