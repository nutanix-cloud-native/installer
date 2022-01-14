package nutanix

// Platform stores any global configuration used for Nutanix platforms.
type Platform struct {
	// PrismCentral is the domain name or IP address of the Prism Central.
	PrismCentral string `json:"prismCentral"`

	// Insecure disables certificate checking when connecting to Prism Central.
	Insecure bool `json:"insecure"`

	// Port is the port to use to connect to the Prism Central.
	Port string `json:"port"`

	// Username is the name of the user to use to connect to the Prism Central.
	Username string `json:"username"`

	// Password is the password for the user to use to connect to the Prism Central.
	Password string `json:"password"`

	// PrismElement is the name of the Prism Element cluster to use in the Prism Central.
	PrismElementUUID string `json:"prismElementUuid"`

	// DefaultStorageContainer is the default datastore to use for provisioning volumes.
	// DefaultStorageContainer string `json:"defaultStorageContainer"`

	// ClusterOSImage overrides the url provided in rhcos.json to download the RHCOS Image
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

	// Network specifies the name of the network to be used by the cluster.
	SubnetUUID string `json:"subnetUuid,omitempty"`
}

