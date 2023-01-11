default_output := "/tmp/actual.dot"

@_default:
    just --list --unsorted

preview: run
    #!/bin/bash
    xdot <(dot {{ default_output }})

run:
    go run src/*.go > {{ default_output }}

@test: run
    [ $(diff test/expected.dot {{ default_output }} | wc -l) -eq 0 ] && echo '\e[1;32mOK\e[0m' || echo "\e[1;31mERROR: test/expected.dot incompatible with /tmp/actual.dot\e[0m"

show_svg: run
    dot -Tsvg /tmp/actual.dot -o /tmp/test.svg
    firefox /tmp/test.svg
