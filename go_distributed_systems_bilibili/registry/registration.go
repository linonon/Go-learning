package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName // A service required B service
	ServiceUpdateURL string        // service's URL which need to be updated
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEnrty struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEnrty
	Removed []patchEnrty
}
