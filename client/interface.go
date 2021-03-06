package client

// APIClient is the interface implemented by the Alauda API client.
type APIClient interface {
	APIServer() string
	Account() string
	Token() string
	Initialize(string, string, string)
	AuthClient
	AppClient
	SpaceClient
	ClusterClient
	LoadBalancerClient
	RegistryClient
	ImageClient
	NodeClient
	ProjectClient
	NamespaceClient
}

// AuthClient is the API client for authentication related APIs.
type AuthClient interface {
	Login(*LoginData) (*LoginResult, error)
}

// AppClient is the API client for application related APIs.
type AppClient interface {
	ListApps(string, string, *ListAppsParams) (*ListAppsResult, error)
	InspectApp(string, string, string) (*App, error)
	GetAppYaml(string, string, string) (string, error)
	StartApp(string, string, string) error
	StopApp(string, string, string) error
	RemoveApp(string, string, string) error
	RunApp(string, string, string, string) error
}

// SpaceClient is the API client for space related APIs.
type SpaceClient interface {
	ListSpaces(*ListSpacesParams) (*ListSpacesResult, error)
	InspectSpace(string, *InspectSpaceParams) (*Space, error)
}

// ClusterClient is the API client for cluster related APIs.
type ClusterClient interface {
	ListClusters() (*ListClustersResult, error)
	InspectCluster(string) (*Cluster, error)
}

// LoadBalancerClient is the API client for LB related APIs.
type LoadBalancerClient interface {
	ListLoadBalancers(*ListLoadBalancersParams) (*ListLoadBalancersResult, error)
	InspectLoadBalancer(string) (*LoadBalancer, error)
	UpdateLoadBalancer(string, *UpdateLoadBalancerData) error
}

// RegistryClient is the API client for the registry related APIs.
type RegistryClient interface {
	ListRegistries() (*ListRegistriesResult, error)
	ListRegistryProjects(string) (*ListRegistryProjectsResult, error)
}

// ImageClient is the API client for the image related APIs.
type ImageClient interface {
	ListImages(string, string) (*ListImagesResult, error)
	ListImageTags(string, string, string) (*ListImageTagsResult, error)
}

// NodeClient is the API client for node related APIs.
type NodeClient interface {
	ListNodes(string) (*ListNodesResult, error)
	InspectNode(string, string) (*Node, error)
	CordonNode(string, string) error
	UncordonNode(string, string) error
	DrainNode(string, string) error
	SetNodeLabels(string, string, *SetNodeLabelsData) error
}

// ProjectClient is the API client for project related APIs.
type ProjectClient interface {
	ListProjects() (*ListProjectsResult, error)
	InspectProject(string) (*Project, error)
}

// NamespaceClient is the API client for namespace related APIs.
type NamespaceClient interface {
	ListNamespaces(string) (*ListNamespacesResult, error)
	InspectNamespace(string, string) (*Namespace, error)
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ APIClient = &Client{}
