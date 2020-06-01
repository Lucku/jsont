package jsont

const (
	_defIndexLeaves = true
	_defValidate    = true
)

// Option is a configuration option of the JSONTransformer and forms part of the functional options pattern
// described, e.g., in the Uber Go Style Guide: https://github.com/uber-go/guide/blob/master/style.md#functional-options
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

func withIndexLeaves(i bool) Option {
	return indexLeavesOption(i)
}

func withValidate(v bool) Option {
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
