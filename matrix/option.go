package matrix

type (
	Option func(*option)
	option struct {
		axis     *int
		keepDims bool
	}
)

func mergeOption(options ...Option) (result option) {
	for _, opt := range options {
		opt(&result)
	}

	return
}

func Axis(axis int) Option {
	return func(so *option) {
		so.axis = &axis
	}
}

func KeepDims(keepDims bool) Option {
	return func(so *option) {
		so.keepDims = keepDims
	}
}
