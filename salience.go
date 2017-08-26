/*
  This is free and unencumbered software released into the public domain. For more
  information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
*/

// Crops an image to its most interesting area
package salience

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type Section struct {
	x, y int
	e    float64
}

// Crop crops an image to its most interesting area with the specified extents
func Crop(img image.Image, cropWidth, cropHeight int) image.Image {
	r := img.Bounds()
	imageWidth := r.Max.X - r.Min.X
	imageHeight := r.Max.Y - r.Min.Y

	if cropWidth > imageWidth {
		cropWidth = imageWidth
	}

	if cropHeight > imageHeight {
		cropHeight = imageHeight
	}

	var x, y int
	sliceStep := imageWidth / 8
	if imageHeight/8 < sliceStep {
		sliceStep = imageHeight / 8
	}
	bestSection := Section{0, 0, 0.0}

	for x = 0; x < imageWidth-cropWidth; x += sliceStep {
		for y = 0; y < imageHeight-cropHeight; y += sliceStep {
			e := entropy(img, image.Rect(x, y, x+cropWidth, y+cropHeight))

			if e > bestSection.e {
				bestSection.e = e
				bestSection.x = x
				bestSection.y = y
			}
		}
	}
	return crop(img, image.Rect(bestSection.x, bestSection.y, bestSection.x+cropWidth, bestSection.y+cropHeight))
}

func crop(img image.Image, r image.Rectangle) image.Image {
	cropped := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(cropped, cropped.Bounds(), img, r.Min, draw.Src)
	return cropped
}

// Calculate the entropy of a portion of an image
// From http://www.astro.cornell.edu/research/projects/compression/entropy.html
func entropy(img image.Image, r image.Rectangle) float64 {
	arraySize := 256*2 - 1
	freq := make([]float64, arraySize)

	for x := r.Min.X; x < r.Max.X-1; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			diff := greyvalue(img.At(x, y)) - greyvalue(img.At(x+1, y))
			if -(arraySize+1)/2 < diff && diff < (arraySize+1)/2 {
				freq[diff+(arraySize-1)/2]++
			}
		}
	}

	n := 0.0
	for _, v := range freq {
		n += v
	}

	e := 0.0
	for i := 0; i < len(freq); i++ {
		freq[i] = freq[i] / n
		if freq[i] != 0.0 {
			e -= freq[i] * math.Log2(freq[i])
		}
	}

	return e

}

func greyvalue(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int((r*299 + g*587 + b*114) / 1000)
}
