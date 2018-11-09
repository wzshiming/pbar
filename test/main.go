package main

import (
	"math/rand"
	"time"

	"github.com/wzshiming/pbar"
	"github.com/wzshiming/pbar/styles"
)

func main() {

	renderer := pbar.NewRenderer()
	pbar1, _ := styles.Normal.New(map[string]interface{}{
		"Content": "Hello world!",
	})
	info1 := renderer.Add(pbar1)
	info1.SetTotal(100)
	go func() {
		for {
			time.Sleep(time.Second / 5)
			info1.AddCurrent(uint64(rand.Int63n(2)))
		}
	}()
	renderer.Wait()
}
