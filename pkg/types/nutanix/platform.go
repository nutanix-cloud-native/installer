package nutanix

// Platform stores any global configuration used for Nutanix platforms.
type Platform struct {
	// PrismCentral is the endpoint (address and port) and credentials to
	// connect to the Prism Central.
	PrismCentral PrismCentral `json:"prismCentral"`

	// PrismElements holds a list of uuid and the endpoint (address and port) of
	// the Prism Elements (clusters). Currently we only support one
	// Prism Element (cluster) for an OpenShift cluster, where all the Nutanix resources (VMs, subnets, etc.)
	// used in the OpenShift cluster are located. In the future, we may support the Nutanix resources (VMs, etc.)
	// used in the OpenShift cluster can come from multiple Prism Elements (clusters) of the Prism Cental.
	PrismElements []PrismElement `json:"prismElements"`

	// ClusterOSImage overrides the url provided in rhcos.json to download the RHCOS Image
	//
	// +optional
	ClusterOSImage string `json:"clusterOSImage,omitempty"`

	// APIVIP is the virtual IP address for the api endpoint
	//
	// +kubebuilder:validation:format=ip
	// +optional
	APIVIP string `json:"apiVIP,omitempty"`

	// IngressVIP is the virtual IP address for ingress
	//
	// +kubebuilder:validation:format=ip
	// +optional
	IngressVIP string `json:"ingressVIP,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on Nutanix for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	// SubnetUUIDs identifies the network subnets to be used by the cluster.
	// Currently we only support one subnet for an OpenShift cluster.
	SubnetUUIDs []string `json:"subnetUUID"`
}

// PrismCentral holds the endpoint and credentials data used to connect to the Prism Central
type PrismCentral struct {
	// Endpoint holds the address and port of the Prism Central
	Endpoint PrismEndpoint `json:"endpoint"`

	// Username is the name of the user to connect to the Prism Central
	Username string `json:"username"`

	// Password is the password for the user to connect to the Prism Central
	Password string `json:"password"`
}

// PrismElement holds the uuid, endpoint of the Prism Element (cluster)
type PrismElement struct {
	// UUID is the UUID of the Prism Element (cluster)
	UUID string `json:"uuid"`

	// Endpoint holds the address and port of the Prism Element
	Endpoint PrismEndpoint `json:"endpoint"`
}

// PrismEndpoint holds the endpoint address and port to access the Nutanix Prism Central or Element (cluster)
type PrismEndpoint struct {
	// address is the endpoint address (DNS name or IP address) of the Nutanix Prism Central or Element (cluster)
	Address string `json:"address"`

	// port is the port number to access the Nutanix Prism Central or Element (cluster)
	Port int32 `json:"port"`
}
