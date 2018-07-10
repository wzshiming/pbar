package main

import (
	"math/rand"
	"time"

	"github.com/wzshiming/pbar"
)

func main() {

	renderer := pbar.NewRenderer()
	info1 := renderer.Add(pbar.NewNormalStyle("Hello world!"))
	info1.SetTotal(200)
	info2 := renderer.Add(pbar.NewBreakStyle("Hello world!"))
	info2.SetTotal(200)
	info3 := renderer.Add(pbar.NewFillStyle("Hello world!"))
	info3.SetTotal(200)
	info4 := renderer.Add(pbar.NewPercentStyle("Hello world!"))
	info4.SetTotal(200)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			info1.AddCurrent(uint64(rand.Int63n(2)))
			info2.AddCurrent(uint64(rand.Int63n(2)))
			info3.AddCurrent(uint64(rand.Int63n(2)))
			info4.AddCurrent(uint64(rand.Int63n(2)))
		}
	}()
	renderer.Wait()
}
