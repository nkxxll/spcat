package main

import (
	_ "image/jpeg"
	"os"
	"spcat/pkg/renderer"
)

func main() {
	file := "../fixtures/test.jpg"
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		os.Exit(1)
	}

	stat, err := f.Stat()
	if err != nil {
		os.Exit(1)
	}

	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	if err != nil {
		os.Exit(1)
	}

	termImg := renderer.GenerateKittyPic(buf, renderer.KittyOptions)
	os.Stdout.Write(termImg)
}
