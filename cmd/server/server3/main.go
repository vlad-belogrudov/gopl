package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vlad-belogrudov/gopl/pkg/color"
	"github.com/vlad-belogrudov/gopl/pkg/surface"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error
		if err = r.ParseForm(); err != nil {
			log.Println("cannot parse params: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params := surface.DefaultSurfaceParams()
		for k, v := range r.Form {
			switch k {
			case "width":
				params.Width, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse width: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "height":
				params.Height, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse height ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "xyrange":
				params.XYRange, err = strconv.ParseFloat(v[0], 10)
				if err != nil {
					log.Println("cannot parse xyrange: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "cells":
				params.Cells, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse cells: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "topcolor":
				c, err := strconv.ParseUint(v[0], 0, 24)
				if err != nil {
					log.Println("cannot parse topcolor: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				params.TopColor = color.Color{
					Red:   color.ColorByte((c & (0xFF << 16)) >> 16),
					Green: color.ColorByte((c & (0xFF << 8)) >> 8),
					Blue:  color.ColorByte(c & 0xFF),
				}
			case "bottomcolor":
				c, err := strconv.ParseUint(v[0], 0, 24)
				if err != nil {
					log.Println("cannot parse bottomcolor: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				params.BottomColor = color.Color{
					Red:   color.ColorByte((c & (0xFF << 16)) >> 16),
					Green: color.ColorByte((c & (0xFF << 8)) >> 8),
					Blue:  color.ColorByte(c & 0xFF)}
			}
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Surface(w, params)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
