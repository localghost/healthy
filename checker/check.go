package checker

type Check interface {
	Configure(options map[string]interface{}) error

	Run() error
}

type CheckRegistry struct {
	providers map[string]func () Check
}

func (r *CheckRegistry) Add(name string, provider func() Check) {
	r.providers[name] = provider
}

func (r *CheckRegistry) CreateAndConfigure(name string, options map[string]interface{}) Check{
	provider := r.providers[name]()
	provider.Configure(options)
	return provider
}

var registry = CheckRegistry{
	providers: make(map[string]func () Check),
}

func init() {
	registry.Add("http", NewHttpCheck)
}

