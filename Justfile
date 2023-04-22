tool := "pr-deps-plotter"
docker_image_version := "0.0.1"
docker_user_repo := "mihaigalos"
docker_image_dockerhub := docker_user_repo + "/" + tool + ":" + docker_image_version
docker_image_dockerhub_latest := docker_user_repo + "/" + tool + ":latest"

@_default:
    just --list --unsorted

dockerize:
    sudo docker build \
        --network=host \
        --tag {{ docker_image_dockerhub }} \
        --tag {{ docker_image_dockerhub_latest }} \
        .

run_docker pr token:
    sudo docker run \
    --rm -it \
    --net=host \
    mihaigalos/pr-deps-plotter:0.0.1 \
        -- \
        {{ pr }} {{ token }}

run pr token:
    #!/bin/bash
    sources=$(find src/ -name *.go -not \( -name *_test.go \))
    go run $sources {{ pr }} {{ token }}

@test:
    go test -v src/*.go

push:
    sudo docker push {{ docker_image_dockerhub }}
    sudo docker push {{ docker_image_dockerhub_latest }}
