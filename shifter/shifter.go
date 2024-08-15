package shifter

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func shiftImageToRight(src image.Image, shift int) *image.RGBA {
	// Get the bounds of the source image
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Create a new blank RGBA image
	newImage := image.NewRGBA(bounds)

	// Draw the original image into the new image but shifted to the right
	draw.Draw(newImage, image.Rect(shift, 0, width, height), src, image.Point{0, 0}, draw.Src)

	return newImage
}

func CreateFrame(src image.Image, shift int) error {
	// Shift the image to the right
	shiftedImage := shiftImageToRight(src, shift)

	// Create the output file
	outputFile, err := os.Create(fmt.Sprintf("output/frame_%d.jpg", shift/10))
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Save the image
	err = jpeg.Encode(outputFile, shiftedImage, nil)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}

func GetImageFromFile(path string) (image.Image, error) {
	// Open the source image
	file, err := os.Open("image.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	// Decode the image
	src, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	return src, nil
}
