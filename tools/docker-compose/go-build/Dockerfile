FROM golang:1.18.2-buster
WORKDIR /workspace

ENV GOPROXY=https://goproxy.io,https://goproxy.cn,direct

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    make \
    gcc g++ \
    pkg-config \
    openssh-client \
    bash 
