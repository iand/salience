salience
========

Go package to detect interesting portions of images (salient region detection).

One method "Crop" is provided which crops the supplied image to the width and height provided, ensuring the most interesting portion of the image is contained within the cropped area.

Crop works by moving a sliding window over the input image and selecting the one with the highest entropy.

Inspired by some of the ideas submitted here: http://codebrawl.com/contests/content-aware-image-cropping-with-chunkypng

Run the sample command line app like this:

go run /path/to/bin/cropper.go  /path/to/input.img /path/to/output.img 200 120

INSTALLATION
============

Simply run

	go get github.com/iand/salience

Documentation is at [http://go.pkgdoc.org/github.com/iand/salience](http://go.pkgdoc.org/github.com/iand/salience)

LICENSE
=======
This code and associated documentation is in the public domain.

To the extent possible under law, Ian Davis has waived all copyright
and related or neighboring rights to this file. This work is published 
from the United Kingdom. 

TIP
===
If you like this code and want to show your appreciation, I accept bitcoin tips at 1NMjYDmQq9X2m8oSSieGh6J6tmJY11K47X