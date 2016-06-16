package util

import (
	"bytes"
	"encoding/base64"

	"github.com/tdewolff/minify"
	minifyJS "github.com/tdewolff/minify/js"
)

// WeaponizeJavascript weaponizes a javascript script by making it minifying
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

	return "eval(atob('" + buf.String() + "'))", nil
}
