package main

import (
"fmt"
"image"
"log"
"os"

// Package image/jpeg is not used explicitly in the code below,
// but is imported for its initialization side-effect, which allows
// image.Decode to understand JPEG formatted images. Uncomment these
// two lines to also understand GIF and PNG images:
// _ "image/gif"
// _ "image/png"
_ "image/jpeg"
)

func main() {
	curDir, _ := os.Getwd()

	path := curDir+"/dsp/lab1/images/"
// Open the file.
file, err := os.Open(path+"1.jpg")
if err != nil {
log.Fatal(err)
}
defer file.Close()

// Decode the image.
m, _, err := image.Decode(file)
if err != nil {
log.Fatal(err)
}
bounds := m.Bounds()

// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
//
// An image's bounds do not necessarily start at (0, 0), so the two loops start
// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
// likely to result in better memory access patterns than X first and Y second.
var histogram [16][4]int
for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
for x := bounds.Min.X; x < bounds.Max.X; x++ {
r, g, b, a := m.At(x, y).RGBA()
// A color's RGBA method returns values in the range [0, 65535].
// Shifting by 12 reduces this to the range [0, 15].
histogram[r>>12][0]++
histogram[g>>12][1]++
histogram[b>>12][2]++
histogram[a>>12][3]++
}
}

// Print the results.
fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
for i, x := range histogram {
fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
}
}