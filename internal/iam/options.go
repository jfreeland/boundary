package iam

// GetOpts - iterate the inbound Options and return a struct
func GetOpts(opt ...Option) options {
	opts := getDefaultOptions()
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Option - how Options are passed as arguments
type Option func(*options)

// options = how options are represented
type options struct {
	withPublicId    string
	withName        string
	withScope       *Scope
	withDescription string
	withGroupGrants bool
}

func getDefaultOptions() options {
	return options{
		withPublicId:    "",
		withScope:       nil,
		withDescription: "",
		withGroupGrants: false,
		withName:        "",
	}
}

// WitPublicId provides an optional public id
func WitPublicId(id string) Option {
	return func(o *options) {
		o.withPublicId = id
	}
}

// WithDescription provides an optional description
func WithDescription(desc string) Option {
	return func(o *options) {
		o.withDescription = desc
	}
}

// WithScope provides an optional scope
func WithScope(s *Scope) Option {
	return func(o *options) {
		o.withScope = s
	}
}

// WithName provides an option to search by a friendly name
func WithName(name string) Option {
	return func(o *options) {
		o.withName = name
	}
}

// WithGroupGrants provides an option to include group grants
func WithGroupGrants(include bool) Option {
	return func(o *options) {
		o.withGroupGrants = include
	}
}
