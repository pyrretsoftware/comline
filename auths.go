package comline

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"strings"
)

type HardcodedKey struct {
	Key string
}

func (k HardcodedKey) GetKey() string {
	return k.Key
}

type Extension struct {
	Name string
	Description string
	URL         string
	Image       image.Image
}

func (e Extension) GetKey() string {
	sections := []string{e.Name, e.Description, e.URL}
	if e.Image != nil {
		b := bytes.Buffer{}	
		err := png.Encode(&b, e.Image)
		if err == nil {
			img := base64.StdEncoding.EncodeToString(b.Bytes())
			sections = append(sections, img)
		}
	}
	
	return "ext:" + strings.Join(sections, ";")
}