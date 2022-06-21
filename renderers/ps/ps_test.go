package ps

import (
	"bytes"
	"testing"

	"github.com/eihigh/canvas"
)

func TestPS(t *testing.T) {
	w := &bytes.Buffer{}
	ps := New(w, 100, 80, nil)
	ps.setColor(canvas.Red)
	//test.String(t, string(w.Bytes()), "")
}
