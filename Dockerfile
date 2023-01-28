FROM golang:alpine3.17 as base


COPY . /src
RUN cd /src/src \
    &&  go build -o pr-deps-plotter

FROM alpine:3.17 as tool
ARG ARCH
ARG JUST_VERSION
ARG USER

COPY --from=base /src/src/pr-deps-plotter /usr/local/bin
RUN cd $(mktemp -d) && \
    wget https://github.com/casey/just/releases/download/${JUST_VERSION}/just-${JUST_VERSION}-${ARCH}-unknown-linux-musl.tar.gz && \
    tar -xzvf *.tar.gz && \
    mv just /usr/local/bin

RUN rm -rf /tmp/tmp.*
RUN adduser -D ${USER}
COPY Justfile /

ENTRYPOINT [ "pr-deps-plotter" ]

