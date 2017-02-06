package echorus

import (
	"fmt"
	"testing"

	"github.com/labstack/gommon/log"
)

func TestPrint(t *testing.T) {
	l := NewLogger()
	l.Debugj(log.JSON{"test": "gogo"})
	l.Debug("test")
	l.Debugf("aa %s", "gogo")

	a := log.JSON{"a": "b"}
	b := log.JSON{"c": "d"}
	c := l.MergeJSON(a, b)
	fmt.Println(c)
}
