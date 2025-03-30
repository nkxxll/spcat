package content

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
)

func GetContent(adress string) []byte {
	if fi, err := os.Stat(adress); err == nil {
		return getFileContent(adress, fi)
	}
	if urlAdr, err := url.Parse(adress); err == nil {
		return getWebContent(urlAdr)
	}
	return nil
}

func getWebContent(url *url.URL) []byte {
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading web content!", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		fmt.Fprintln(os.Stderr, "Content Type is empty!")
		os.Exit(1)
	}
	if contentType == "image/jpeg" {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading response body!")
			os.Exit(1)
		}
		return getJPEGContent(body)
	}
	if contentType != "image/png" {
		fmt.Fprintln(os.Stderr, "Content Type PNG is the only content type that is accepted at the moment!")
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading response body!")
		os.Exit(1)
	}
	return body
}

func getFileContent(adress string, fi os.FileInfo) []byte {
	if fi.IsDir() {
		fmt.Fprintln(os.Stderr, "Interactive is not jet implemented!")
		os.Exit(1)
	}
	content, err := os.ReadFile(adress)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file!")
		os.Exit(1)
	}
	return content
}

// getJPEGContent aka convert jpg bytes to png bytes and return them todo
// better solution you can also encode rgba but I can't get it to work with the
// limited documentation for the kitty graphics protocol
func getJPEGContent(body []byte) []byte {
	img, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding image")
		os.Exit(1)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding image")
		os.Exit(1)
	}
	return buf.Bytes()
}
