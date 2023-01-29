default_output := "/tmp/actual.dot"

arch := `lscpu | grep Architecture | cut -d ' ' -f 21`
tool := "pr-deps-plotter"
docker_image_version := "0.0.1"
docker_user_repo := "mihaigalos"
docker_image_dockerhub := docker_user_repo + "/" + tool + ":" + docker_image_version
docker_image_dockerhub_latest := docker_user_repo + "/" + tool + ":latest"
just_version := "1.12.0"
user := "user"

@_default:
    just --list --unsorted

#preview: run
#    #!/bin/bash
#    xdot <(dot {{ default_output }})

dockerize:
    sudo docker build \
        --build-arg=ARCH={{ arch }} \
        --build-arg=JUST_VERSION={{ just_version }} \
        --build-arg=USER={{ user }} \
        --network=host \
        --tag {{ docker_image_dockerhub }} \
        --tag {{ docker_image_dockerhub_latest }} \
        .
run_docker pr token:
    touch {{ default_output }}
    sudo docker run \
    --rm -it \
    --net=host \
    -v {{ default_output }}:{{ default_output }} \
    mihaigalos/pr-deps-plotter:0.0.1 \
        {{ pr }} {{ token }} > {{ default_output }}

run pr token:
    #!/bin/bash
    sources=$(find src/ -name *.go -not \( -name *_test.go \))
    go run $sources {{ pr }} {{ token }} > {{ default_output }}

@utest:
    go test -v src/*.go

#   show_svg: run
#       dot -Tsvg /tmp/actual.dot -o /tmp/test.svg
#       firefox /tmp/test.svg
