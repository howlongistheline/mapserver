package main

import (
  "image/color"

  "github.com/flopp/go-staticmaps"
  "github.com/fogleman/gg"
  "github.com/golang/geo/s2"
)

func main() {
  ctx := sm.NewContext()
  ctx.SetSize(400, 400)
  ctx.SetZoom(18)
  ctx.AddMarker(sm.NewMarker(s2.LatLngFromDegrees(-41.29839171894141, 174.7704205880379), color.RGBA{0xff, 0, 0, 0xff}, 16.0))

  img, err := ctx.Render()
  if err != nil {
    panic(err)
  }

  if err := gg.SavePNG("my-map.png", img); err != nil {
    panic(err)
  }
}
