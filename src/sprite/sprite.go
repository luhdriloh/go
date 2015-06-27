package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

type Sprite struct {
	spriteHeight int
	spriteWidth  int
	sprites      []image.Image
}

func NewSprite(height, width int, spriteSheet image.Image) (*Sprite, error) {
	spriteRect := spriteSheet.Bounds()

	sheetWidth := spriteRect.Dx()
	sheetHeight := spriteRect.Dy()

	spriteImage := image.NewRGBA(image.Rect(spriteRect.Min.X, spriteRect.Min.Y, sheetWidth, sheetHeight))
	draw.Draw(spriteImage, spriteImage.Bounds(), spriteSheet, image.Point{0, 0}, draw.Src)

	// if sprite height and width do not
	if sheetWidth%width != 0 || sheetHeight%height != 0 {
		return nil, errors.New("Your sprite widths/heights not compatible with your image size")
	}

	// the following tells us how many rows and columns of sprites there are
	xImages := sheetWidth / width
	yImages := sheetHeight / height

	// lets create our slice to return
	spriteSlice := make([]image.Image, 0, xImages*yImages)

	for y := 0; y < yImages; y++ {
		for x := 0; x < xImages; x++ {
			littleSprite := spriteImage.SubImage(image.Rect(width*x, height*y, x*width+width, y*height+height))
			spriteSlice = append(spriteSlice, littleSprite)
		}
	}

	return &Sprite{height, width, spriteSlice}, nil
}

func main() {
	spriteFile, _ := os.Open("red.png")
	defer spriteFile.Close()

	spriteSheet, _ := png.Decode(spriteFile)

	spriteStruct, _ := NewSprite(32, 32, spriteSheet)

	for i, image := range spriteStruct.sprites {
		img, _ := os.Create(fmt.Sprintf("%dblue.png", i))
		defer img.Close()
		png.Encode(img, image)
	}
}
