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

run pr token:
    #!/bin/bash
    sources=$(find src/ -name *.go -not \( -name *_test.go \))
    go run $sources {{ pr }} {{ token }} > {{ default_output }}

@utest:
    go test -v src/*.go

#   @test: run
#       [ $(diff test/expected.dot {{ default_output }} | wc -l) -eq 0 ] && echo '\e[1;32mOK\e[0m' || echo "\e[1;31mERROR: test/expected.dot incompatible with /tmp/actual.dot\e[0m"

#   show_svg: run
#       dot -Tsvg /tmp/actual.dot -o /tmp/test.svg
#       firefox /tmp/test.svg
