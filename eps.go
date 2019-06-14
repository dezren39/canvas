package canvas

import (
	"fmt"
	"image/color"
	"io"
)

var psEllipseDef = `/ellipse {
/endangle exch def
/startangle exch def
/yrad exch def
/xrad exch def
/y exch def
/x exch def
/savematrix matrix currentmatrix def
x y translate
xrad yrad scale
0 0 1 startangle endangle arc
savematrix setmatrix
} def /ellipsen {
/endangle exch def
/startangle exch def
/yrad exch def
/xrad exch def
/y exch def
/x exch def
/savematrix matrix currentmatrix def
x y translate
xrad yrad scale
0 0 1 startangle endangle arcn
savematrix setmatrix
} def`

type EPSWriter struct {
	io.Writer
	color color.RGBA
}

func NewEPSWriter(writer io.Writer, width, height float64) *EPSWriter {
	w := &EPSWriter{
		Writer: writer,
		color:  Black,
	}

	fmt.Fprintf(w, "%%!PS-Adobe-3.0 EPSF-3.0\n%%%%BoundingBox: 0 0 %.5f %.5f\n", width, height)
	fmt.Fprintf(w, psEllipseDef)

	// TODO: generate preview

	return w
}

func (w *EPSWriter) SetColor(color color.RGBA) {
	if color != w.color {
		// TODO: transparency
		fmt.Fprintf(w, " %f %f %f setrgbcolor", float64(color.R)/255.0, float64(color.G)/255.0, float64(color.B)/255.0)
		w.color = color
	}
}