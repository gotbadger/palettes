package palettes

import (
    // "fmt"
    "image/color"
    "math"
)

func Rainbow(c int) []color.RGBA {
  p := make([]color.RGBA, c)
  for i := 0; i < c; i++ {
    p[c-i-1] = color.RGBA{
        sinToDec(c, i, 0 * math.Pi * 2/3),
        sinToDec(c, i, 2 * math.Pi * 2/3),
        sinToDec(c, i, 1 * math.Pi * 2/3),
        255,
    }
  }
  return p
}

func sinToDec(c int, i int, phase float64) uint8 {
  s := math.Sin(math.Pi / float64(c) * 2 * float64(i) + phase)
  return uint8(math.Floor(s * 127) + 128)
}
