FROM golang:alpine3.17 as base

COPY . /src
RUN cd /src/src \
    &&  go build -o pr-deps-plotter

FROM alpine:3.17 as tool
ARG USER

RUN apk update && \
    apk add \
        graphviz

COPY --from=base /src/src/pr-deps-plotter /usr/local/bin

RUN adduser -D ${USER}

ENTRYPOINT [ "/bin/sh", "-c", "pr-deps-plotter $1 $2 | dot -Tsvg" ]

