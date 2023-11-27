package main

import (
	"image/color"

	"github.com/hamao0820/zerokara/matrix"
	"github.com/petar/GoMNIST"
)

const (
	TRAIN_URL      = "http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz"
	LABEL_URL      = "http://yann.lecun.com/exdb/mnist/train-labels-idx1-ubyte.gz"
	TEST_TRAIN_URL = "http://yann.lecun.com/exdb/mnist/t10k-images-idx3-ubyte.gz"
	TEST_LABEL_URL = "http://yann.lecun.com/exdb/mnist/t10k-labels-idx1-ubyte.gz"

	LOCAL_DATA_PATH        = "data"
	LOCAL_TRAIN_PATH       = LOCAL_DATA_PATH + "/train-images-idx3-ubyte.gz"
	LOCAL_LABELS_PATH      = LOCAL_DATA_PATH + "/train-labels-idx1-ubyte.gz"
	LOCAL_TEST_TRAIN_PATH  = LOCAL_DATA_PATH + "/t10k-images-idx3-ubyte.gz"
	LOCAL_TEST_LABELS_PATH = LOCAL_DATA_PATH + "/t10k-labels-idx1-ubyte.gz"
)

type (
	datasetOption struct {
		Train           *bool
		Transform       func(matrix.Matrix) matrix.Matrix
		TargetTransform func(matrix.Matrix) matrix.Matrix
	}
	DatasetOption func(*datasetOption)
	Dataset       interface {
		Get(idx interface{}) (matrix.Matrix, matrix.Matrix)
		Len() int
		IsTrain() bool
	}
	dataset struct {
		data            matrix.Matrix
		label           matrix.Matrix
		transform       func(matrix.Matrix) matrix.Matrix
		targetTransform func(matrix.Matrix) matrix.Matrix
		prepare         func()
		isTrain         bool
	}

	mnist struct{ Dataset }
)

func ApplyDataSetOpt(options ...DatasetOption) datasetOption {
	option := datasetOption{}
	for _, opt := range options {
		opt(&option)
	}
	return option
}

func NewMnist(options ...DatasetOption) Dataset {
	opt := ApplyDataSetOpt(options...)

	trainSet, testSet, err := GoMNIST.Load(LOCAL_DATA_PATH)
	if err != nil {
		panic(err)
	}

	if opt.Train != nil && *opt.Train {
		trainData, trainLabels := loadMatrix(trainSet)
		return &mnist{Dataset: ExtendsDataset(trainData, trainLabels, func() {}, options...)}
	}

	testData, testLabels := loadMatrix(testSet)
	return &mnist{Dataset: ExtendsDataset(testData, testLabels, func() {}, options...)}

}

func loadMatrix(s *GoMNIST.Set) (data, labels matrix.Matrix) {
	mat := matrix.NewMat(matrix.Shape{R: len(s.Images), C: s.NRow * s.NRow})
	lMat := matrix.NewMat(matrix.Shape{R: len(s.Labels), C: 1})

	for i := 0; i < len(s.Images); i++ {
		i := i
		img, label := s.Get(i)
		b := img.Bounds()
		// p := plot.New()
		// p.Add(plotter.NewImage(img, 0, 0, 200, 200))
		// err := p.Save(5*vg.Centimeter, 5*vg.Centimeter, "mnist/"+cnv.MustStr(label)+"_"+cnv.MustStr(i)+".png")
		// if err != nil {
		// 	log.Debug("error saving image plot: %v\n", err)
		// }
		row := make([]float64, 0, s.NRow*s.NRow)
		for x := 0; x < b.Max.Y; x++ {
			for y := 0; y < b.Max.X; y++ {
				row = append(row, float64(img.At(y, x).(color.Gray).Y))
			}
		}

		mat.SetRow(i, row)
		lMat.SetRow(i, []float64{float64(label)})
	}

	return mat, lMat
}

func TransformData(fn func(matrix.Matrix) matrix.Matrix) DatasetOption {
	return func(so *datasetOption) {
		so.Transform = fn
	}
}
func TransformLabel(fn func(matrix.Matrix) matrix.Matrix) DatasetOption {
	return func(so *datasetOption) {
		so.TargetTransform = fn
	}
}
func Train(train bool) DatasetOption {
	return func(so *datasetOption) {
		so.Train = &train
	}
}

func ExtendsDataset(data, label matrix.Matrix, prepare func(), options ...DatasetOption) Dataset {
	option := ApplyDataSetOpt(options...)
	if option.TargetTransform == nil {
		option.TargetTransform = func(x matrix.Matrix) matrix.Matrix { return x }
	}
	if option.Transform == nil {
		option.Transform = func(x matrix.Matrix) matrix.Matrix { return x }
	}

	instance := &dataset{
		prepare:         prepare,
		data:            data,
		label:           label,
		transform:       option.Transform,
		targetTransform: option.TargetTransform,
	}
	if option.Train != nil {
		instance.isTrain = *option.Train
	}
	if instance.prepare != nil {
		instance.prepare()
	}

	return instance
}

func (d *dataset) Get(idx interface{}) (matrix.Matrix, matrix.Matrix) {
	if d.label == nil {
		return d.transform(d.data.Cat(idx)), nil
	}
	return d.transform(d.data.Cat(idx)), d.targetTransform(d.label.Cat(idx))
}

func (d *dataset) Len() int {
	return d.data.Len()
}
func (d *dataset) IsTrain() bool {
	return d.isTrain
}
