package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vlad-belogrudov/gopl/pkg/mandelbrot"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error
		if err = r.ParseForm(); err != nil {
			log.Println("cannot parse params: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var xmin, xmax, ymin, ymax float64 = -2, 2, -2, 2
		width, height := 1920, 1080
		for k, v := range r.Form {
			switch k {
			case "width":
				width, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse width: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "height":
				height, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse height ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "xrange":
				xrange, err := strconv.ParseFloat(v[0], 10)
				if err != nil {
					log.Println("cannot parse xrange: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				xmin = -xrange / 2
				xmax = xrange / 2
			case "yrange":
				yrange, err := strconv.ParseFloat(v[0], 10)
				if err != nil {
					log.Println("cannot parse yrange: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				ymin = -yrange / 2
				ymax = yrange / 2
			}
		}
		mandelbrot.Mandelbrot(w, xmin, xmax, ymin, ymax, width, height)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
