#!/bin/bash

VERSION=0.3.0
docker build -t atombender/gce-metadata-exporter:${VERSION} .
docker tag atombender/gce-metadata-exporter:${VERSION} atombender/gce-metadata-exporter:latest
docker push atombender/gce-metadata-exporter:${VERSION}
docker push atombender/gce-metadata-exporter:latest
