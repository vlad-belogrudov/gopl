package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vlad-belogrudov/gopl/pkg/lissajous"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error
		if err = r.ParseForm(); err != nil {
			log.Println("cannot parse params: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params := lissajous.DefaultLissajousParams()
		for k, v := range r.Form {
			switch k {
			case "cycles":
				params.Cycles, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse cycles: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "res":
				params.Res, err = strconv.ParseFloat(v[0], 10)
				if err != nil {
					log.Println("cannot parse res: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "xsize":
				params.Xsize, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse xsize: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "ysize":
				params.Ysize, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse ysize: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "frames":
				params.Frames, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse frames: ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			case "delay":
				params.Delay, err = strconv.Atoi(v[0])
				if err != nil {
					log.Println("cannot parse delay ", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
		}
		lissajous.Lissajous(w, params)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
