FROM golang:1.15

VOLUME /workspace

RUN apt update && \
    apt install -y \
        vim \
        rsyslog \
        strace \
        graphviz \
        netcat \
        dnsutils \
        net-tools

RUN go get -v \
        golang.org/x/tools/gopls \
        golang.org/x/tools/cmd/godoc \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint

WORKDIR /workspace
    
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["/bin/bash"]
