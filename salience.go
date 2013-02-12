/*
  PUBLIC DOMAIN STATEMENT
  To the extent possible under law, Ian Davis has waived all copyright
  and related or neighboring rights to this Source Code file.
  This work is published from the United Kingdom.
*/

// Inspired by https://gist.github.com/zaeleus/a54cd41137b678935c91

// Crops an image to its most interesting area
package salience

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type Section struct {
	x, y int
	e    float64
}

// Crops an image to its most interesting area with the specified extents
func Crop(img image.Image, cropWidth, cropHeight int) image.Image {
	r := img.Bounds()
	imageWidth := r.Max.X - r.Min.X
	imageHeight := r.Max.Y - r.Min.Y

	x, y := 0, 0
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
	cropped := image.NewRGBA(r)
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			cropped.Set(x, y, img.At(x, y))
		}
	}
	return cropped
}

func greyscale(img image.Image) *image.Gray {
	grey := image.NewGray(img.Bounds())

	if rgba, ok := img.(*image.RGBA); ok {
		// Optimised path
		for p := 0; p < len(grey.Pix); p++ {
			i := p * 4
			r := uint32(rgba.Pix[i])
			g := uint32(rgba.Pix[i+1])
			b := uint32(rgba.Pix[i+2])
			grey.Pix[p] = uint8((r*299 + g*587 + b*114) / 1000)

		}
	} else {
		// Generic path
		for p := 0; p < len(grey.Pix); p++ {
			r, g, b, _ := img.At(p%grey.Stride, p/grey.Stride).RGBA()
			grey.Pix[p] = uint8((r*299 + g*587 + b*114) / 1000)
		}
	}
	return grey
}

func histogram(img image.Image) map[uint8]int {
	h := make(map[uint8]int, 0)
	r := img.Bounds()
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			greyval := uint8((r*299 + g*587 + b*114) / 1000)
			h[greyval] += 1
		}
	}
	return h
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
				freq[diff+(arraySize-1)/2] += 1
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

	fmt.Printf("Entropy of (%d, %d) (%d, %d) is %0.2f\n", r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, e)

	return e

}

func greyvalue(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int((r*299 + g*587 + b*114) / 1000)
}
