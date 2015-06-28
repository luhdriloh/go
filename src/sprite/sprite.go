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
	leftSprites  []image.Image
	rightSprites []image.Image
	upSprites    []image.Image
	downSprites  []image.Image
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

	// loop over our sprite sheet and cut each individual one out from top left to bottom right
	for y := 0; y < yImages; y++ {
		for x := 0; x < xImages; x++ {
			littleSprite := spriteImage.SubImage(image.Rect(width*x, height*y, x*width+width, y*height+height))
			spriteSlice = append(spriteSlice, littleSprite)
		}
	}

	leftImage := make([]image.Image, 0, xImages*yImages)
	rightImage := make([]image.Image, 0, xImages*yImages)
	upImage := make([]image.Image, 0, xImages*yImages)
	downImage := make([]image.Image, 0, xImages*yImages)

	return &Sprite{height, width, spriteSlice, leftImage, rightImage, upImage, downImage}, nil
}

// the following sets your left or right sprites movement images. start off with the stationary sprite
func setSprites(sprites []image.Image, imageSlice []image.Image, values []int) (bool, error, []image.Image) {

	for _, value := range values {
		if value >= len(sprites) {
			return false, errors.New("Index is to big"), imageSlice
		}

		imageSlice = append(imageSlice, sprites[value])
	}

	return true, nil, imageSlice
}

func (p *Sprite) SetLeftSprites(values ...int) (bool, error) {
	worked, errorValue, images := setSprites(p.sprites, p.leftSprites, values)
	p.leftSprites = images

	return worked, errorValue
}

func (p *Sprite) SetRightSprites(values ...int) (bool, error) {
	worked, errorValue, images := setSprites(p.sprites, p.rightSprites, values)
	p.rightSprites = images

	return worked, errorValue
}

func (p *Sprite) SetUpSprites(values ...int) (bool, error) {
	worked, errorValue, images := setSprites(p.sprites, p.upSprites, values)
	p.upSprites = images

	return worked, errorValue
}

func (p *Sprite) SetDownprites(values ...int) (bool, error) {
	worked, errorValue, images := setSprites(p.sprites, p.downSprites, values)
	p.downSprites = images

	return worked, errorValue
}

func main() {
	spriteFile, _ := os.Open("red.png")
	defer spriteFile.Close()

	spriteSheet, _ := png.Decode(spriteFile)

	spriteStruct, _ := NewSprite(32, 32, spriteSheet)

	spriteStruct.SetLeftSprites(3, 5, 4)

	for i, image := range spriteStruct.leftSprites {
		img, _ := os.Create(fmt.Sprintf("%dblue.png", i))
		defer img.Close()
		png.Encode(img, image)
	}
}
