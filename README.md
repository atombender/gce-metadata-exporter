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
$ docker run atombender/gce-metadata-exporter:0.3.0
```

# Usage

```shell
$ gce-metadata-exporter
```

# License

MIT license. See `LICENSE` file.
