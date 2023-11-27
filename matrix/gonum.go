package matrix

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type matrix struct{ *mat.Dense }

func init() {
	// blas64.Use(netlib.Implementation{})
}

func withDense(dense *mat.Dense) Matrix { return &matrix{Dense: dense} }
func WithShape(r, c int) Matrix         { return NewMat(Shape{R: r, C: c}) }
func NewMat(s Shape) Matrix             { return withDense(mat.NewDense(s.R, s.C, make([]float64, s.R*s.C))) }
func NewRand(shape Shape) Matrix {
	nums := []float64{}
	for i := 0; i < shape.R*shape.C; i++ {
		nums = append(nums, random(0, 1))
	}
	return &matrix{Dense: mat.NewDense(shape.R, shape.C, nums)}
}
func NewRandN(shape Shape, options ...RandSeedOption) Matrix {
	option := ApplyRandSeedOption(options...)
	nums := []float64{}
	for i := 0; i < shape.R*shape.C; i++ {
		seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
		var random *rand.Rand
		if option.seed == nil {
			random = rand.New(rand.NewSource(seed.Int64()))
		} else {
			random = rand.New(rand.NewSource(int64(*option.seed)))
		}
		nums = append(nums, random.NormFloat64())
	}
	return &matrix{Dense: mat.NewDense(shape.R, shape.C, nums)}
}
func NewEmpty() Matrix {
	var dense mat.Dense
	return withDense(&dense)
}
func NewFull(shape Shape, v float64) Matrix {
	m := NewMat(shape)
	m.Full(v)
	return m
}
func New1D(dence ...float64) Matrix { return withDense(mat.NewDense(1, len(dence), dence)) }
func New2D(dence [][]float64) Matrix {
	tmpDence := []float64{}
	for _, denceArr := range dence {
		tmpDence = append(tmpDence, denceArr...)
	}
	return &matrix{Dense: mat.NewDense(len(dence), len(dence[0]), tmpDence)}
}
func NewRandUniform(min, max float64, shape Shape) Matrix {
	uniform := distuv.Uniform{Max: max, Min: min}
	return NewMat(shape).CopyApply(func(f float64) float64 {
		return uniform.Rand()
	})
}
func Eye(row, col int) Matrix {
	result := NewMat(Shape{R: row, C: col})
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i-j == 0 {
				result.Set(i, j, 1)
			}
		}
	}

	return result
}
func Linspace(start, stop float64, num int) Matrix {
	result := make([]float64, num)
	step := (stop - start) / float64(num-1)

	for i := range result {
		result[i] = start + float64(i)*step
	}

	result[len(result)-1] = stop

	return New1D(result...)
}
func (m *matrix) Set(i, j int, v float64)     { m.Dense.Set(i, j, v) }
func (m *matrix) SetRow(idx int, v []float64) { m.Dense.SetRow(idx, v) }
func (m *matrix) MatMul(v Matrix)             { m.Dense.Mul(m, v) }
func (m *matrix) Log()                        { m.Apply(func(v float64) float64 { return math.Log(v) }) }
func (m *matrix) Scale(f float64)             { m.Dense.Scale(f, m) }
func (m *matrix) Full(f float64)              { m.Apply(func(v float64) float64 { return f }) }
func (m *matrix) Reshape(s Shape)             { m.Dense = mat.NewDense(s.R, s.C, m.RawMatrix().Data) }
func (m *matrix) Shape() (result Shape)       { result.R, result.C = m.Dims(); return }
func (m *matrix) Flatten() []float64          { return m.RawMatrix().Data }
func (m *matrix) Dims() (row, col int)        { return m.Dense.Dims() }
func (m *matrix) Rows() int                   { return m.Dense.RawMatrix().Rows }
func (m *matrix) Cols() int                   { return m.Dense.RawMatrix().Cols }
func (m *matrix) IsEmpty() bool               { return m.Dense.IsEmpty() }

func (m *matrix) Copy() Matrix                  { return withDense(mat.DenseCopyOf(m.Dense)) }
func (m *matrix) CopyT() Matrix                 { return withDense(mat.DenseCopyOf(m.Dense.T())) }
func (m *matrix) CopyAdd(v Matrix) Matrix       { return add(m, v) }
func (m *matrix) CopySub(v Matrix) Matrix       { return sub(m, v) }
func (m *matrix) CopyMul(v Matrix) Matrix       { return mul(m, v) }
func (m *matrix) CopyDiv(v Matrix) Matrix       { return div(m, v) }
func (m *matrix) CopyAddFloat(v float64) Matrix { return addf(m, v) }
func (m *matrix) CopySubFloat(v float64) Matrix { return subf(m, v) }
func (m *matrix) CopyMulFloat(v float64) Matrix { return mulf(m, v) }
func (m *matrix) CopyDivFloat(v float64) Matrix { return divf(m, v) }
func (m *matrix) CopyMatMul(v Matrix) Matrix    { return matMul(m, v) }
func (m *matrix) CopyScale(f float64) Matrix    { return scale(f, m) }
func (m *matrix) CopyLog() Matrix               { return m.CopyApply(func(v float64) float64 { return math.Log(v) }) }
func (m *matrix) CopyClip(min, max float64) Matrix {
	return m.CopyApply(func(v float64) float64 {
		if v > max {
			return max
		}
		if v < min {
			return min
		}
		return v
	})
}
func (m *matrix) CopyFull(f float64) Matrix {
	result := m.Copy()
	result.Full(f)
	return result
}
func (m *matrix) CopyApply(fn func(v float64) float64) Matrix {
	return m.CopyApplyWithIndex(func(i, j int, v float64) float64 {
		return fn(v)
	})
}
func (m *matrix) CopyApplyWithIndex(fn func(i, j int, v float64) float64) Matrix {
	result := m.Copy()
	result.ApplyWithIndex(fn)
	return result
}
func (m *matrix) CopyReshape(s Shape) Matrix {
	return &matrix{Dense: mat.NewDense(s.R, s.C, m.RawMatrix().Data)}
}
func (m *matrix) Apply(fn func(v float64) float64) {
	m.ApplyWithIndex(func(i, j int, v float64) float64 { return fn(v) })
}
func (m *matrix) ApplyWithIndex(fn func(i, j int, v float64) float64) {
	var dence mat.Dense
	dence.Apply(fn, m.Dense)
	m.Dense = &dence
}
func (m *matrix) Len() int {
	if m.Shape().R == 1 {
		return m.Shape().C
	}
	return m.Shape().R
}
func (m *matrix) Swap(i1, j1, i2, j2 int) {
	a := m.Dense.At(i1, j1)
	b := m.Dense.At(i2, j2)
	m.Dense.Set(i1, j1, b)
	m.Dense.Set(i2, j2, a)
}
func (m *matrix) Shuffle() {
	shape := m.Shape()
	m.Each(func(i, j int, v float64) {
		m.Swap(i, j, rand.Intn(shape.R), rand.Intn(shape.C))
	})
}
func (m *matrix) Cat(indexs interface{}) Matrix {
	switch idx := indexs.(type) {
	case float64:
		if m.Rows() == 1 {
			return New1D(m.At(0, int(idx)))
		}
		return New1D(m.RawRowView(int(idx))...)
	case int:
		if m.Rows() == 1 {
			return New1D(m.At(0, idx))
		}
		return New1D(m.RawRowView(idx)...)
	case []int:
		if m.Shape().R == 1 {
			result := make([]float64, 0, len(idx))
			for _, idx := range idx {
				result = append(result, m.At(0, idx))
			}
			return New1D(result...)
		}
		result := [][]float64{}
		for _, idx := range idx {
			result = append(result, m.RawRowView(idx))
		}
		return New2D(result)
	case []float64:
		if m.Shape().R == 1 {
			result := make([]float64, 0, len(idx))
			for _, idx := range idx {
				result = append(result, m.At(0, int(idx)))
			}
			return New1D(result...)
		}
		result := [][]float64{}
		for _, idx := range idx {
			result = append(result, m.RawRowView(int(idx)))
		}
		return New2D(result)
	}

	panic("invlid arg index")
}
func (m *matrix) ArgMax(options ...Option) Matrix {
	option := mergeOption(options...)
	if option.axis == nil {
		tmp := m.Flatten()
		maxValue := tmp[0]
		result := 0
		for i, v := range tmp {
			if maxValue < v {
				maxValue = v
				result = i
			}
		}
		return New1D(float64(result))
	}

	if *option.axis == 0 {
		result := []float64{}
		for j := 0; j < m.Cols(); j++ {
			maxValue := m.At(0, j)
			tmp := 0
			for i := 1; i < m.Rows(); i++ {
				if maxValue < m.At(i, j) {
					maxValue = m.At(i, j)
					tmp = i
				}
			}
			result = append(result, float64(tmp))
		}
		return New1D(result...)
	}

	result := []float64{}
	for i := 0; i < m.Rows(); i++ {
		maxValue := m.At(i, 0)
		tmp := 0
		for j := 1; j < m.Cols(); j++ {
			if maxValue < m.At(i, j) {
				maxValue = m.At(i, j)
				tmp = j
			}
		}
		result = append(result, float64(tmp))
	}
	return New1D(result...)
}

func (m *matrix) Mean(options ...Option) Matrix {
	option := mergeOption(options...)
	if option.axis == nil {
		tmp := m.Flatten()
		return New1D(m.Sum().At(0, 0) / float64(len(tmp)))
	}

	panic("no implments")
}
func (m *matrix) Sum(options ...Option) Matrix {
	opt := mergeOption(options...)
	if opt.axis != nil {
		shape := m.Shape()
		switch *opt.axis {
		case 0:
			tmp := make([]float64, shape.C)
			m.Each(func(i, j int, _ float64) { tmp[j] += m.At(i, j) })
			return New1D(tmp...)
		case 1:
			tmp := make([]float64, shape.R)
			m.Each(func(i, j int, _ float64) { tmp[i] += m.At(i, j) })
			return &matrix{Dense: mat.NewDense(shape.R, 1, tmp)}
		default:
			panic("missing axis options")
		}
	} else {
		return New1D(mat.Sum(m))
	}
}

func (m *matrix) Max(options ...Option) Matrix {
	opt := mergeOption(options...)
	if opt.axis == nil {

		return New1D(slices.Max(m.Flatten()))
	}

	result := []float64{}
	if *opt.axis == 0 {
		for i := 0; i < m.Dense.RawMatrix().Cols; i++ {
			tmp := []float64{}
			for j := 0; j < m.Dense.RawMatrix().Rows; j++ {
				tmp = append(tmp, m.At(j, i))
			}
			result = append(result, slices.Max(tmp))
		}
		return New1D(result...)
	}

	if *opt.axis == 1 {
		for i := 0; i < m.Dense.RawMatrix().Rows; i++ {
			result = append(result, slices.Max(m.Dense.RawRowView(i)))
		}
		return New1D(result...).CopyT()
	}

	return nil
}
func (m *matrix) Allclose(a Matrix, rtol, atol *float64) bool {
	if rtol == nil {
		rtol = new(float64)
		*rtol = 1e-5
	}
	if atol == nil {
		atol = new(float64)
		*atol = 1e-8
	}

	for i := 0; i < m.Dense.RawMatrix().Rows; i++ {
		for j := 0; j < m.Dense.RawMatrix().Cols; j++ {
			aij := m.At(i, j)
			bij := a.At(i, j)
			if math.Abs(aij-bij) > (*atol + (*rtol)*math.Abs(bij)) {
				return false
			}
		}
	}
	return true
}
func (m *matrix) Broadcast(s Shape) Matrix {
	result := NewMat(s)

	if r, c := m.Dims(); r == 1 && c == 1 {
		result.ApplyWithIndex(func(i, j int, v float64) float64 { return m.At(0, 0) })
	} else if s.R == r {
		result.ApplyWithIndex(func(i, j int, v float64) float64 { return m.At(i, 0) })
	} else if s.C == c {
		result.ApplyWithIndex(func(i, j int, v float64) float64 { return m.At(0, j) })
	} else {
		panic("missing shape")
	}

	return result
}
func (m *matrix) Each(fn func(i, j int, v float64)) {
	for i := 0; i < m.RawMatrix().Rows; i++ {
		for j := 0; j < m.RawMatrix().Cols; j++ {
			fn(i, j, m.At(i, j))
		}
	}
}
func (m *matrix) EachR(fn func(i int, v []float64)) {
	for i := 0; i < m.RawMatrix().Rows; i++ {
		tmp := []float64{}
		for j := 0; j < m.RawMatrix().Cols; j++ {
			tmp = append(tmp, m.At(i, j))
		}
		fn(i, tmp)
	}
}
func (m *matrix) EachC(fn func(i int, v []float64)) {
	for i := 0; i < m.RawMatrix().Cols; i++ {
		tmp := []float64{}
		for j := 0; j < m.RawMatrix().Rows; j++ {
			tmp = append(tmp, m.At(j, i))
		}
		fn(i, tmp)
	}
}
func (m *matrix) OnOff(fn func(i, j int, v float64) bool) Matrix {
	return m.CopyApplyWithIndex(func(i, j int, v float64) float64 {
		if fn(i, j, v) {
			return 1
		}
		return 0
	})
}
func (m *matrix) Clip(min, max float64) Matrix {
	m.Apply(func(v float64) float64 {
		if v > max {
			return max
		}
		if v < min {
			return min
		}
		return v
	})
	return m
}
func (m *matrix) Array() [][]float64 {
	result := [][]float64{}
	for i := 0; i < m.Rows(); i++ {
		result = append(result, m.Dense.RawRowView(i))
	}

	return result
}

func (m *matrix) Search(x interface{}, y interface{}) []float64 {
	result := []float64{}
	intx, xok := x.([]int)
	inty, yok := y.([]int)
	if !xok {
		if floatx, ok := x.([]float64); ok {
			for _, fx := range floatx {
				intx = append(intx, int(fx))
			}
		} else {
			panic("missing x type")
		}
	}
	if !yok {
		if floaty, ok := y.([]float64); ok {
			for _, fy := range floaty {
				inty = append(inty, int(fy))
			}
		} else {
			panic("missing x type")
		}
	}
	for i := range intx {
		result = append(result, m.At(intx[i], inty[i]))
	}

	return result
}

func (m *matrix) String() string {
	return fmt.Sprintf(" %v", mat.Formatted(m, mat.Prefix(" "), mat.Squeeze()))
}
