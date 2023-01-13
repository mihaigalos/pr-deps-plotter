default_output := "/tmp/actual.dot"

tool := "pr-deps-plotter"
docker_image_version := "0.0.1"
docker_user_repo := "mihaigalos"
docker_image_dockerhub := docker_user_repo + "/" + tool + ":" + docker_image_version
docker_image_dockerhub_latest := docker_user_repo + "/" + tool + ":latest"

@_default:
    just --list --unsorted

preview: run
    #!/bin/bash
    xdot <(dot {{ default_output }})

build:
    sudo docker build \
        --network=host \
        --tag {{ docker_image_dockerhub }} \
        --tag {{ docker_image_dockerhub_latest }} \
        .

run:
    go run src/*.go > {{ default_output }}

@test: run
    [ $(diff test/expected.dot {{ default_output }} | wc -l) -eq 0 ] && echo '\e[1;32mOK\e[0m' || echo "\e[1;31mERROR: test/expected.dot incompatible with /tmp/actual.dot\e[0m"

show_svg: run
    dot -Tsvg /tmp/actual.dot -o /tmp/test.svg
    firefox /tmp/test.svg
