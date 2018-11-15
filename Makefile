



.PHONY: test
test: pbar
	cd examples && ./example1.sh

pbar: style_fmt styles_bind sh_fmt
	go get github.com/wzshiming/pbar/cmd/pbar

styles_bind: tool_bindata
	cd styles && go-bindata -pkg styles -o styles.go *.json

style_fmt: tool_stylefmt
	cd styles && stylefmt -w *.json

sh_fmt: tool_shfmt
	cd examples && shfmt -w *.sh

tool_bindata:
	go get github.com/wzshiming/go-bindata/go-bindata

tool_stylefmt:
	go get github.com/wzshiming/pbar/cmd/stylefmt

tool_shfmt:
	go get mvdan.cc/sh/cmd/shfmt
