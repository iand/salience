# salience

Go package to detect interesting portions of images (salient region detection).

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iand/salience)
[![Check Status](https://github.com/iand/salience/actions/workflows/check.yml/badge.svg)](https://github.com/iand/salience/actions/workflows/check.yml)
[![Test Status](https://github.com/iand/salience/actions/workflows/test.yml/badge.svg)](https://github.com/iand/salience/actions/workflows/test.yml)

## Usage

One method "Crop" is provided which crops the supplied image to the width and height provided, ensuring the most interesting portion of the image is contained within the cropped area.

Crop works by moving a sliding window over the input image and selecting the one with the highest entropy.

Inspired by some of the ideas submitted here: http://codebrawl.com/contests/content-aware-image-cropping-with-chunkypng

Run the sample command line app like this:

go run /path/to/bin/cropper.go  /path/to/input.img /path/to/output.img 200 120

## Installation

Simply run

	go get github.com/iand/salience

Documentation is at [https://pkg.go.dev/github.com/iand/salience](https://pkg.go.dev/github.com/iand/salience)

## Authors

* [Ian Davis](http://github.com/iand) - <http://iandavis.com/>


## Contributors


## Contributing

* Do submit your changes as a pull request
* Do your best to adhere to the existing coding conventions and idioms.
* Do run `go fmt` on the code before committing 
* Do feel free to add yourself to the [`CREDITS`](CREDITS) file and the
  corresponding Contributors list in the [`README.md`](README.md). 
  Alphabetical order applies.
* Don't touch the [`AUTHORS`](AUTHORS) file. An existing author will add you if 
  your contributions are significant enough.
* Do note that in order for any non-trivial changes to be merged (as a rule
  of thumb, additions larger than about 15 lines of code), an explicit
  Public Domain Dedication needs to be on record from you. Please include
  a copy of the statement found in the [`WAIVER`](WAIVER) file with your pull request

## License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.
