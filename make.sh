#!/bin/sh
set -eux
go build -o codegen ./_codegen
for goos in linux darwin dragonfly freebsd netbsd openbsd solaris; do
	GOOS=$goos ./codegen >generated_${goos}.go
	gofmt -s -w generated_${goos}.go
	GOOS=$goos go vet generated_${goos}.go
	GOOS=$goos go build .
done
rm -f codegen
