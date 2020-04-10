package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"strconv"
	_ "strconv"
	"strings"

	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"

	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createImage(lat float64, long float64) string {
	ctx := sm.NewContext()
	ctx.SetSize(400, 400)
	ctx.SetZoom(18)
	ctx.AddMarker(sm.NewMarker(s2.LatLngFromDegrees(lat, long), color.RGBA{0xff, 0, 0, 0xff}, 16.0))

	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}
	name := fmt.Sprintf("%v%f%v%f%v", "images/", lat, ",", long, ".png")
	if err := gg.SavePNG(name, img); err != nil {
		panic(err)
	}
	return name
}

func handler(w http.ResponseWriter, r *http.Request) {
	request := r.URL.String()
	if strings.Contains(request, "K3340") { //put some API security here if the service gets abused.
		request = strings.TrimLeft(request, "/")
		parsed := strings.Split(request, ",")
		lat, _ := strconv.ParseFloat(parsed[0], 64)
		long, _ := strconv.ParseFloat(parsed[1], 64)
		path := createImage(lat, long)
		w.Header().Set("Content-Type", "image/png")
		image, _ := ioutil.ReadFile(path)
		w.Write(image)
	} else {
		fmt.Fprintf(w, "Error: no API key found in your request!")
	}

}
