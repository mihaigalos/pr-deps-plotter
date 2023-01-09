@_default:
    just --list --unsorted

preview:
    #!/bin/bash
    xdot <(dot test/expected.dot)

@test:
    go run src/main.go > /tmp/actual.dot
    [ $(diff test/expected.dot /tmp/actual.dot | wc -l) -eq 0 ] && echo '\e[1;32mOK\e[0m' || echo "\e[1;31mERROR: test/expected.dot incompatible with /tmp/actual.dot\e[0m"
