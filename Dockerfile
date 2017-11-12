FROM ubuntu:14.04

RUN apt-get update && \
apt-get install -y software-properties-common python-software-properties

RUN add-apt-repository ppa:gophers/archive && \
apt update && \
apt-get install -y golang-1.9-go && \
mkdir -p /go/src/

ENV PATH="/usr/lib/go-1.9/bin:${PATH}"

COPY . /go/src
