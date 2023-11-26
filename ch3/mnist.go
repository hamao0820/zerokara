package main

import (
	"compress/gzip"
	"encoding/binary"
	"errors"
	"image"
	"image/color"
	"io"
	"os"
	"path"
)

var (
	// ErrFormat indicates that the file has not been recognised.
	ErrFormat = errors.New("mnist: invalid format")

	// ErrSize indicates that the labels and images count mismatch.
	ErrSize = errors.New("mnist: size mismatch")
)

// MINST image dimension in pixels.
const (
	Width  = 28
	Height = 28
)

// MNIST database file names.
const (
	TrainingImageFileName = "data/train-images-idx3-ubyte.gz"
	TrainingLabelFileName = "data/train-labels-idx1-ubyte.gz"
	TestImageFileName     = "data/t10k-images-idx3-ubyte.gz"
	TestLabelFileName     = "data/t10k-labels-idx1-ubyte.gz"
)

// Image represents a MNIST image. It is a array a bytes representing the color.
// 0 is black (the background) and 255 is white (the digit color).
type Image [Width * Height]byte

// Label is the digit label from 0 to 9.
type Label int8

// Set represents the data set with the images paired with the labels.
type Set struct {
	Images []*Image
	Labels []Label
}

type imageFileHeader struct {
	Magic     int32
	NumImages int32
	Height    int32
	Width     int32
}

type labelFileHeader struct {
	Magic     int32
	NumLabels int32
}

// Magic keys are used to check file formats.
const (
	imageMagic = 0x00000803
	labelMagic = 0x00000801
)

// readImage reads a image from the file and returns it.
func readImage(r io.Reader) (*Image, error) {
	img := &Image{}
	err := binary.Read(r, binary.BigEndian, img)
	return img, err
}

// LoadImageFile opens the image file, parses it, and returns the data in order.
func LoadImageFile(name string) ([]*Image, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	header := imageFileHeader{}

	err = binary.Read(reader, binary.BigEndian, &header)
	if err != nil {
		return nil, err
	}

	if header.Magic != imageMagic ||
		header.Width != Width ||
		header.Height != Height {
		return nil, ErrFormat
	}

	images := make([]*Image, header.NumImages)
	for i := int32(0); i < header.NumImages; i++ {
		images[i], err = readImage(reader)
		if err != nil {
			return nil, err
		}
	}

	return images, nil
}

// LoadLabelFile opens the label file, parses it, and returns the labels in
// order.
func LoadLabelFile(name string) ([]Label, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	header := labelFileHeader{}

	err = binary.Read(reader, binary.BigEndian, &header)
	if err != nil {
		return nil, err
	}

	if header.Magic != labelMagic {
		return nil, err
	}

	labels := make([]Label, header.NumLabels)
	for i := int32(0); i < header.NumLabels; i++ {
		err = binary.Read(reader, binary.BigEndian, &labels[i])
		if err != nil {
			return nil, err
		}
	}

	return labels, nil
}

// ColorModel implements the image.Image interface.
func (img *Image) ColorModel() color.Model {
	return color.GrayModel
}

// Bounds implements the image.Image interface.
func (img *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{Width, Height},
	}
}

// At implements the image.Image interface.
func (img *Image) At(x, y int) color.Color {
	return color.Gray{Y: img[y*Width+x]}
}

// Set modifies the pixel at (x,y).
func (img *Image) Set(x, y int, v byte) {
	img[y*Width+x] = v
}

// LoadSet loads the images and labels, check if the counts match and returns
// a set.
func LoadSet(imageName, labelName string) (*Set, error) {
	images, err := LoadImageFile(imageName)
	if err != nil {
		return nil, err
	}

	labels, err := LoadLabelFile(labelName)
	if err != nil {
		return nil, err
	}

	if len(images) != len(labels) {
		return nil, ErrSize
	}

	set := &Set{
		Images: images,
		Labels: labels,
	}

	return set, nil
}

// Count returns the number of images and labels in the set.
func (s *Set) Count() int {
	return len(s.Labels)
}

// Get returns the i-th image and its label.
func (s *Set) Get(i int) (*Image, Label) {
	return s.Images[i], s.Labels[i]
}

// Load loads the whole MINST database and returns the training set and the test
// set.
func Load(dir string) (training, test *Set, err error) {
	training, err = LoadSet(path.Join(dir, TrainingImageFileName),
		path.Join(dir, TrainingLabelFileName))
	if err != nil {
		return nil, nil, err
	}

	test, err = LoadSet(path.Join(dir, TestImageFileName),
		path.Join(dir, TestLabelFileName))
	if err != nil {
		return nil, nil, err
	}

	return
}
