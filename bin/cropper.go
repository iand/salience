/*
This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
*/

package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"strconv"

	"github.com/iand/salience"
)

// A simple command line program for finding the most interesting section of an image
func main() {
	if len(os.Args) < 5 {
		println("Please supply input image filename, output filename and output image width and height as arguments")
		os.Exit(1)
	}
	finName := os.Args[1]
	foutName := os.Args[2]
	widthStr := os.Args[3]
	heightStr := os.Args[4]

	width, err := strconv.ParseInt(widthStr, 10, 0)
	if err != nil {
		fmt.Printf("Error parsing image width argument: %s\n", err.Error())
		os.Exit(1)
	}

	height, err := strconv.ParseInt(heightStr, 10, 0)
	if err != nil {
		fmt.Printf("Error parsing image height argument: %s\n", err.Error())
		os.Exit(1)
	}

	fin, err := os.OpenFile(finName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Error reading input image: %s\n", err.Error())
		os.Exit(1)
	}

	img, _, err := image.Decode(fin)
	if err != nil {
		fmt.Printf("Error decoding input image: %s\n", err.Error())
		os.Exit(1)
	}

	fout, err := os.OpenFile(foutName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error writing output image: %s\n", err.Error())
		os.Exit(1)
	}

	imgOut := salience.Crop(img, int(width), int(height))

	if err = png.Encode(fout, imgOut); err != nil {
		fmt.Printf("Error encoding output image: %s\n", err.Error())
		os.Exit(1)
	}
}
