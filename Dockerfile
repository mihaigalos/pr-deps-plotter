FROM golang:alpine3.17 as base

COPY . /src

RUN cd /src/src \
    &&  go build -o pr-deps-plotter

FROM alpine:3.17 as tool

COPY --from=base /src/src/pr-deps-plotter /usr/local/bin

ENTRYPOINT [ "pr-deps-plotter" ]

