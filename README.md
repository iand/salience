# salience

Go package to detect interesting portions of images (salient region detection).

## Usage

One method "Crop" is provided which crops the supplied image to the width and height provided, ensuring the most interesting portion of the image is contained within the cropped area.

Crop works by moving a sliding window over the input image and selecting the one with the highest entropy.

Inspired by some of the ideas submitted here: http://codebrawl.com/contests/content-aware-image-cropping-with-chunkypng

Run the sample command line app like this:

go run /path/to/bin/cropper.go  /path/to/input.img /path/to/output.img 200 120

## Installation

Simply run

	go get github.com/iand/salience

Documentation is at [http://godoc.org/github.com/iand/salience](http://godoc.org/github.com/iand/salience)

## Authors

* [Ian Davis](http://github.com/iand) - <http://iandavis.com/>


## Credits


## Contributing

* Do submit your changes as a pull request
* Do your best to adhere to the existing coding conventions and idioms.
* Do run `go fmt` on the code before committing 
* Do feel free to add yourself to the [`CREDITS`](CREDITS) file and the
  corresponding list in the the `README.md`. Alphabetical order applies.
* Don't touch the `AUTHORS` file. If your contributions are significant
  enough, be assured we will eventually add you in there.
* Do note that in order for us to merge any non-trivial changes (as a rule
  of thumb, additions larger than about 15 lines of code), we need an
  explicit [public domain dedication][PDD] on record from you. Please include
  a the statement found in the `WAIVER` file with your pull request

## License

This is free and unencumbered public domain software. For more
information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
