salience
========

Go package to detect interesting portions of images.

One method "Crop" is provided which crops the supplied image to the width and height provided, ensuring the most interesting portion of the image is contained within the cropped area.

Crop works by moving a sliding window over the input image and selecting the one with the highest entropy.

Inspired by some of the ideas submitted here: http://codebrawl.com/contests/content-aware-image-cropping-with-chunkypng

Run the sample command line app like this:

go run /path/to/bin/cropper.go  /path/to/input.img /path/to/output.img 200