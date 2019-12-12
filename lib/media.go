package lib

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mattn/go-sixel"
)

// Media is display media lib
type Media struct{}

// ShowImage is display image
func (m *Media) ShowImage(url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("cannot get image: %v: %v", err, res.Status)
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("cannot decode image %v", err)
	}

	img, _, err := image.Decode(bytes.NewReader(buf))

	if err != nil {
		return fmt.Errorf("cannot decode image %v", err)
	}

	enc := sixel.NewEncoder(os.Stdout)
	return enc.Encode(img)
}
