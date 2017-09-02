// server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var width, height = 600, 320
var cells, xyrange = 100, 30.0
var xyscale, zscale = float64(width) / float64(2) /xyrange, float64(height) * 0.4
var angle = math.Pi / 6

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	r.ParseForm()
	for k, v := range r.Form {
		val := v[0]
		switch {
			case k == "height":
				height, _ = strconv.Atoi(val)
			case k == "width":
				width, _ = strconv.Atoi(val)
		}

	}
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			color := "#FF0000"
			if az < 0 {
				color = "#0000FF"
			}
			fmt.Fprintf(w, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	sin_a, cos_a := math.Sin(angle), math.Cos(angle)

	// compute surface height
	z := f(x, y)

	// Project (x,y,z) isometrically onto a 2-D SVG canvas (sx, sy).
	sx := float64(width)/2.0 + (x-y)*cos_a*xyscale
	sy := float64(height)/2.0 + (x+y)*sin_a*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
