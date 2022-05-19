#!/usr/bin/env bash

image_name="asppj/go-build-tools:v0.1.0"

docker buildx build -t ${image_name} --platform=linux/arm64,linux/amd64 . --push
