package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	img "github.com/brg-liuwei/gowkhtmltox/gowkhtmltoimage"
)

func main() {
	img.Init(false)

	conv := img.NewConvertor()
	conv.SetProperty("web.loadImages", "true")
	conv.SetProperty("fmt", "jpeg")
	// conv.SetProperty("in", "https://www.baidu.com/")
	conv.SetHtml("<html><body><h1>Hello World</h1></body></html>")

	if err := conv.Ready(); err != nil {
		panic("ready error")
	}

	now := time.Now()
	if err := conv.Run(); err != nil {
		panic("run error")
	}
	fmt.Fprintln(os.Stderr, "use time: ", time.Since(now))

	binary.Write(os.Stdout, binary.LittleEndian, conv.GetImage())

	img.DeInit()
}
