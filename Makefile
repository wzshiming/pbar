

.PHONY: all
all: pbar crun

pbar: style_fmt styles_bind sh_fmt
	go get github.com/wzshiming/pbar/cmd/pbar

styles_bind: go-bindata
	cd styles && go-bindata -pkg styles -o styles.go *.json

style_fmt: stylefmt
	cd styles && stylefmt -w *.json

sh_fmt: shfmt
	cd examples && shfmt -w *.sh

go-bindata:
	go get github.com/wzshiming/go-bindata/go-bindata

stylefmt:
	go get github.com/wzshiming/pbar/cmd/stylefmt

shfmt:
	go get mvdan.cc/sh/cmd/shfmt

crun:
	go get github.com/wzshiming/crun/cmd/crun
