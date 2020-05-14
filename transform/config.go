package transform

const (
	_defIndexLeaves = true
	_defValidate    = true
)

type Option interface {
	apply(opts *options)
}

type options struct {
	indexLeaves bool
	validate    bool
}

type (
	indexLeavesOption bool
	validateOption    bool
)

func (i indexLeavesOption) apply(opts *options) {
	opts.indexLeaves = true
}

func (v validateOption) apply(opts *options) {
	opts.validate = true
}

func WithIndexLeaves(i bool) Option {
	return indexLeavesOption(i)
}

func WithValidate(v bool) Option {
	return validateOption(v)
}

func newOptions(opts ...Option) options {

	opt := options{
		indexLeaves: _defIndexLeaves,
		validate:    _defValidate,
	}

	for _, o := range opts {
		o.apply(&opt)
	}

	return opt
}
