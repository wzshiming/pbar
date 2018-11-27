package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wzshiming/pbar"
)

var (
	w = flag.Bool("w", false, "Write the changes to the file")
)

func init() {
	flag.Usage = func() {
		w := os.Stdout
		fmt.Fprintf(w, "pbarstylefmt:\n")
		fmt.Fprintf(w, "Usage:\n")
		fmt.Fprintf(w, "    %s [Options] file1 [filen ...]\n", os.Args[0])
		fmt.Fprintf(w, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
	for _, file := range args {
		err := format(file, *w)
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}
	}
}

func format(file string, w bool) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	i := &pbar.Marks{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return err
	}

	ret, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return err
	}
	if !w {
		fmt.Print(string(ret))
		return nil
	}

	err = ioutil.WriteFile(file, ret, 0666)
	if err != nil {
		return err
	}
	return nil
}
