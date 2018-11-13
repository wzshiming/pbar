


help: 
	@echo "\`make styles_bind\` Style config bind to Go-source."
	@echo "\`make json_fmt\`    Style config format."

all: json_fmt styles_bind

styles_bind: tool_bindata
	go-bindata -pkg styles -o styles.go *.json

json_fmt: tool_jsonfmt
	jsonfmt -w *.json

tool_bindata:
	go get github.com/wzshiming/go-bindata/go-bindata

tool_jsonfmt:
	go get gopkg.in/ffmt.v1/cmd/jsonfmt