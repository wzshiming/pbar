package main

import (
	"math/rand"
	"time"

	"github.com/wzshiming/pbar"
)

func main() {

	renderer := pbar.NewRenderer()
	info1 := renderer.Add(pbar.NewNormalStyle("Hello world!"))
	info1.SetTotal(100)
	info2 := renderer.Add(pbar.NewBreakStyle("Hello world! hello all! biu biu biu!"))
	info2.SetTotal(100)
	info3 := renderer.Add(pbar.NewFillStyle("Hello world!"))
	info3.SetTotal(100)
	info4 := renderer.Add(pbar.NewFillRedStyle("Hello world!"))
	info4.SetTotal(100)
	info5 := renderer.Add(pbar.NewAddedStyle("Hello world!"))
	info5.SetTotal(100)
	info6 := renderer.Add(pbar.NewTTYShowStyle("Hello world! hello all! biu biu biu!"))
	info6.SetTotal(100)
	renderer.Start()
	go func() {
		for {
			time.Sleep(time.Second / 5)
			info1.AddCurrent(uint64(rand.Int63n(2)))
			info2.AddCurrent(uint64(rand.Int63n(2)))
			info3.AddCurrent(uint64(rand.Int63n(2)))
			info4.AddCurrent(uint64(rand.Int63n(2)))
			info5.AddCurrent(uint64(rand.Int63n(2)))
			info6.AddCurrent(uint64(rand.Int63n(2)))
		}
	}()
	renderer.Wait()
}
