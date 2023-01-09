@_default:
    just --list --unsorted

preview:
    #!/bin/bash
    xdot <(dot example.dot)
