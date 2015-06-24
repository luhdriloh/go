package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	// using the os library we must first open the file up
	image_name := "yeah.jpg"
	fimage1, _ := os.Open(image_name)
	defer fimage1.Close()

	// fimage1 is of interface io.Reader since File (its type) implements the Read
	// method
	image1, _, _ := image.Decode(fimage1)

	m := image.NewRGBA(image.Rect(0, 0, 2880, 1800))
	draw.Draw(m, m.Bounds(), image1, image.Point{0, 0}, draw.Src)

	m_bounds := m.Bounds()

	x_min := m_bounds.Min.X
	x_max := m_bounds.Max.X

	y_min := m_bounds.Min.Y
	y_max := m_bounds.Max.Y

	for y := y_min; y < y_max; y++ {
		for x := x_min; x < x_max/2; x++ {
			color_on_right := m.At(x, y)
			color_on_left := m.At(x_max-x-1, y)

			m.Set(x_max-x-1, y, color_on_right)
			m.Set(x, y, color_on_left)
		}
	}

	new_image_name := "reverse" + image_name
	toimg, _ := os.Create(new_image_name)
	defer toimg.Close()

	jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})
}
