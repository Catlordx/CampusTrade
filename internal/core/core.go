package core

type option struct {
	disablePProf      bool
	disableSwagger    bool
	disablePrometheus bool
	enableCors        bool
}
type Option func(*option)

func WithDisablePProf() Option {
	return func(o *option) {
		o.disablePProf = true
	}
}
func WithDisableSwagger() Option {
	return func(o *option) {
		o.disableSwagger = true
	}
}
func WithDisablePrometheus() Option {
	return func(o *option) {
		o.disablePrometheus = true
	}
}
func WithEnableCors() Option {
	return func(o *option) {
		o.enableCors = true
	}
}
