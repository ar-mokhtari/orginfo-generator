package generator

type Runner interface {
	Name() string
	Init(args []string) error
	Run() error
}
