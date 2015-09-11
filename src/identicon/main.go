package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {

	// enter the email that you would like to become an identicon
	var email string

	fmt.Println("Please enter your email than press ENTER!")
	_, error := fmt.Scanf("%s", &email)
	if error != nil {
		return
	}

	md5Hasher := md5.New()
	md5Hasher.Write([]byte(email))
	hashedEmail := hex.EncodeToString(md5Hasher.Sum(nil))

	// we will create a 5x5 block that will hold our identicon
	// each block will be 16x16 pixels wide

	blockSide := 16

	identicon := image.NewRGBA(image.Rect(0, 0, blockSide*5,
		blockSide*5))

	bounds := identicon.Bounds()

	colorOne := GetColor([]byte(hashedEmail[:1])[0])
	colorTwo := GetColor([]byte(hashedEmail[1:2])[0])

	if colorOne == colorTwo {
		colorTwo = color.RGBA{255, 255, 255, 0}
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			identicon.Set(x, y, colorTwo)
		}
	}

	for i := 2; i < 32; i += 2 {
		// we will add consecutive bytes together which will equal
		// a max of 32. if 0-7, or 16-23 then box will be shaded
		// any other value it will not
		addedBytes := HexAdd([]byte(hashedEmail[i : i+1])[0],
			[]byte(hashedEmail[i+1 : i+2])[0])

		if addedBytes >= 8 && addedBytes <= 15 || addedBytes >= 8 && addedBytes <= 15 {
			continue
		}

		column := ((i - 2) % 6) / 2
		row := (i - 2) / 6

		// this tells us that we should start the painting of the box at
		// x == 32 pixels
		xStart := blockSide * column
		xEnd := xStart + blockSide

		yStart := blockSide * row
		yEnd := yStart + blockSide

		for y := yStart; y < yEnd; y++ {
			for x := xStart; x < xEnd; x++ {
				identicon.Set(x, y, colorOne)
			}
		}

		if column == 0 || column == 1 {
			column += 4 - column*2

			xStart = blockSide * column
			xEnd = xStart + blockSide

			yStart = blockSide * row
			yEnd = yStart + blockSide

			for y := yStart; y < yEnd; y++ {
				for x := xStart; x < xEnd; x++ {
					identicon.Set(x, y, colorOne)
				}
			}
		}

	}

	newImg, _ := os.Create("identicon.jpeg")
	defer newImg.Close()

	jpeg.Encode(newImg, identicon, &jpeg.Options{jpeg.DefaultQuality})
}
