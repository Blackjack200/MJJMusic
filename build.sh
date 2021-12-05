#!/bin/bash

function build() {
		  echo "[+] Build GOOS=${1} ARCH=${2}"
		  	    NAME="mjjmusic_${1}_${2}"
			    	      GOOS=${1} GOARCH=${2} go build -ldflags="-w -s" -o "$NAME" main.go
				      	        mv "$NAME" bin/"$NAME"
}

rm -rdf bin 2>/dev/null >/dev/null
mkdir bin 2>/dev/null >/dev/null

build linux amd64 &
build darwin amd64 &
wait