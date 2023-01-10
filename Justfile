@_default:
    just --list --unsorted

preview: run
    #!/bin/bash
    xdot <(dot /tmp/actual.dot)

run:
    go run src/main.go > /tmp/actual.dot

@test: run
    [ $(diff test/expected.dot /tmp/actual.dot | wc -l) -eq 0 ] && echo '\e[1;32mOK\e[0m' || echo "\e[1;31mERROR: test/expected.dot incompatible with /tmp/actual.dot\e[0m"
