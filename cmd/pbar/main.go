package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/wzshiming/pbar/run"
	"github.com/wzshiming/pbar/styles"
)

var (
	reg   = flag.String("r", `^(?P<Title>.+)\s+(?P<Current>\d+)/(?P<Total>\d+)$`, "Match the regularity of the text")
	title = flag.String("t", "Title", "Unique identification fields for multiple bars")
	style = flag.String("s", "normal", "Select the style of the progress bar")
)

func init() {
	flag.Usage = func() {
		w := os.Stderr
		fmt.Fprintf(w, "jsonfmt:\n")
		fmt.Fprintf(w, "Usage:\n")
		fmt.Fprintf(w, "    %s [Options] -- Match progress information for progress bar from input\n", os.Args[0])
		fmt.Fprintf(w, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	bar, err := styles.OpenBuiltinStyle(*style)
	if err != nil {
		conf, err0 := ioutil.ReadFile(*style)
		if err0 != nil {
			fmt.Fprint(os.Stderr, err)
			flag.Usage()
			return
		}
		bar, err = styles.NewConfig(conf)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			flag.Usage()
			return
		}
	}

	reg, err := regexp.Compile(*reg)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		flag.Usage()
		return
	}

	run.RunBar(os.Stdin, os.Stdout, bar, reg, *title)
}
