

.PHONY: all
all: pbar crun

pbar: pbar_style_fmt styles_bind sh_fmt
	go install github.com/wzshiming/pbar/cmd/pbar

styles_bind: go-bindata
	cd styles && go-bindata --pkg styles -o styles.go *.json

pbar_style_fmt: pbarstylefmt
	cd styles && pbarstylefmt -w *.json

sh_fmt: shfmt
	cd examples && shfmt -w *.sh

go-bindata:
	go get github.com/wzshiming/go-bindata/cmd/go-bindata

pbarstylefmt:
	go install github.com/wzshiming/pbar/cmd/pbarstylefmt

shfmt:
	go get mvdan.cc/sh/cmd/shfmt

crun:
	go get github.com/wzshiming/crun/cmd/crun
