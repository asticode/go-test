package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Flags
var (
	count = flag.Int("c", 10, "the number of times the action is repeated")

	content = `package $$$

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astisub"
	"github.com/asticode/go-astitodo"
	"github.com/asticode/go-astits"
)

func Test() {
	_ = astilectron.Event{}
	_ = astisub.Subtitles{}
	_ = astitodo.TODOs{}
	_ = astits.Demuxer{}
}
`
)

func main() {
	flag.Parse()
	astilog.FlagInit()
	var err error
	for idx := 0; idx < *count; idx++ {
		var n = fmt.Sprintf("test%d", idx+1)
		var d = "./" + n
		var f = d + "/" + n + ".go"
		astilog.Debugf("Removing %s", d)
		if err = os.RemoveAll(d); err != nil {
			astilog.Fatal(errors.Wrapf(err, "removeAll of %s failed", d))
		}
		astilog.Debugf("Mkdiring %s", d)
		if err = os.MkdirAll(d, 0755); err != nil {
			astilog.Fatal(errors.Wrapf(err, "mkdirAll of %s failed", d))
		}
		var w io.WriteCloser
		astilog.Debugf("Creating %s", f)
		if w, err = os.Create(f); err != nil {
			astilog.Fatal(errors.Wrapf(err, "creating %s failed", f))
		}
		defer w.Close()
		astilog.Debugf("Writing to %s", f)
		if _, err = w.Write([]byte(strings.Replace(content, "$$$", n, -1))); err != nil {
			astilog.Fatal(errors.Wrapf(err, "writing to %s failed", f))
		}
	}
}
