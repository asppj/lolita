
FROM docker.io/golang:1.13

WORKDIR /app

RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go get -u -v \
#     github.com/mdempsky/gocode \
#     github.com/uudashr/gopkgs/cmd/gopkgs \
#     github.com/ramya-rao-a/go-outline \
#     github.com/acroca/go-symbols \
#     golang.org/x/tools/cmd/guru \
#     golang.org/x/tools/cmd/gorename \
#     github.com/rogpeppe/godef \
#     github.com/zmb3/gogetdoc \
#     github.com/sqs/goreturns \
#     golang.org/x/tools/cmd/goimports \
#     golang.org/x/lint/golint \
#     github.com/alecthomas/gometalinter \
#     github.com/golangci/golangci-lint/cmd/golangci-lint \
#     github.com/mgechev/revive \
#     github.com/derekparker/delve/cmd/dlv 2>&1
#
# # gocode-gomod
#RUN go get -x -d github.com/stamblerre/gocode \
#     && go build -o gocode-gomod github.com/stamblerre/gocode \
#     && mv gocode-gomod $GOPATH/bin/

# Install git, process tools, lsb-release (common in install instructions for CLIs)
#RUN apt-get update && apt-get -y install git procps lsb-release

# Clean up
#RUN apt-get autoremove -y \
#    && apt-get clean -y \
#    && rm -rf /var/lib/apt/lists/*

# Set the default shell to bash instead of sh
ENV SHELL /bin/bash