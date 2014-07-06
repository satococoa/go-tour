package main

import (
  "code.google.com/p/go-tour/pic"
  "image"
  "image/color"
)

type Image struct{
  Width, Height int
}

func (img Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, img.Width, img.Height)
}

func makeSlice(width, height int) [][]uint8 {
  slice := make([][]uint8, width)
  for j := range slice {
    slice[j] = make([]uint8, height)
    for k := range slice[j] {
      slice[j][k] = uint8(k ^ j)
    }
  }
  return slice
}

func (img Image) At(x, y int) color.Color {
  slice := makeSlice(img.Width, img.Height)
  return color.RGBA{slice[y][x], slice[y][x], 255, 255}
}

func main() {
  m := Image{100, 100}
  pic.ShowImage(&m)
}
