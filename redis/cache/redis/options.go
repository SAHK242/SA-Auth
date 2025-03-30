package rediscache

var (
	defaultOptions = &options{
		addr:         "localhost:6379",
		minIdleConns: 3,
	}
)

const (
	defaultAddr         = "localhost:6379"
	defaultMinIdleConns = 3
)

type (
	options struct {
		addr         string
		minIdleConns int
	}

	Option func(*options)
)

func WithAddr(addr string) Option {
	return func(o *options) {
		if addr == "" {
			o.addr = defaultAddr
		} else {
			o.addr = addr
		}
	}
}

func WithMinIdleConns(minIdleConns int) Option {
	return func(o *options) {
		if minIdleConns == 0 {
			o.minIdleConns = defaultMinIdleConns
		} else {
			o.minIdleConns = minIdleConns
		}
	}
}

func eval(opts []Option) *options {
	o := &options{
		addr:         defaultOptions.addr,
		minIdleConns: defaultOptions.minIdleConns,
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}
