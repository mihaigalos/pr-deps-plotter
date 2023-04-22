#!/bin/sh

echo 'SVG<<EOF' >> $GITHUB_OUTPUT
go run $sources $1 $2 | dot -Tsvg >> $GITHUB_OUTPUT
echo 'EOF' >> $GITHUB_OUTPUT
cat $(echo $GITHUB_OUTPUT)

