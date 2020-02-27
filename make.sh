#!/bin/sh
set -eu
go build -o codegen ./_codegen
x() {
	GOOS=$1
	GOARCH=$2
	export GOOS GOARCH
	file=generated_${GOOS}_${GOARCH}.go
	echo "Generating $file (${GOOS}/${GOARCH})"
	./codegen >$file
	gofmt -s -w $file
	echo "Testing $file (${GOOS}/${GOARCH})"
	go vet $file
	go build .
}
set -x
rm -rf generated_*.go
set +x
for f in $(go tool dist list); do
	set -f
	IFS='/'
	set -- $f
	goos=$1
	goarch=$2
	set +f
	unset IFS
	case $goos in
	js | plan9 | windows)
		echo "Skipping ${goos}/${goarch}: Unsupported"
		;;
	android)
		# https://golang.org/pkg/go/build says "Using GOOS=android matches build tags and files as for GOOS=linux in addition to android tags and files."
		echo "Skipping ${goos}/${goarch}: Covered by linux/${goarch}"
		;;
	illumos)
		# https://golang.org/pkg/go/build says "Using GOOS=illumos matches build tags and files as for GOOS=solaris in addition to illumos tags and files."
		echo "Skipping ${goos}/${goarch}: Covered by solaris/${goarch}"
		;;
	*)
		x $goos $goarch
		;;
	esac
done
rm -f codegen
