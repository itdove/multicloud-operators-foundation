package aws

// MachinePoolPlatform stores the configuration for a machine pool
// installed on AWS.
type MachinePoolPlatform struct {
	// Zones is list of availability zones that can be used.
	Zones []string `json:"zones,omitempty"`

	// Subnets is the list of subnets to which to attach the machines.
	// There must be exactly one subnet for each availability zone used.
	Subnets []string `json:"subnets,omitempty"`

	// InstanceType defines the ec2 instance type.
	// eg. m4-large
	InstanceType string `json:"type"`

	// EC2RootVolume defines the storage for ec2 instance.
	EC2RootVolume `json:"rootVolume"`

	// SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.
	// +optional
	SpotMarketOptions *SpotMarketOptions `json:"spotMarketOptions,omitempty"`
}

// SpotMarketOptions defines the options available to a user when configuring
// Machines to run on Spot instances.
// Most users should provide an empty struct.
type SpotMarketOptions struct {
	// The maximum price the user is willing to pay for their instances
	// Default: On-Demand price
	// +optional
	MaxPrice *string `json:"maxPrice,omitempty"`
}

// EC2RootVolume defines the storage for an ec2 instance.
type EC2RootVolume struct {
	// IOPS defines the iops for the storage.
	IOPS int `json:"iops"`
	// Size defines the size of the storage.
	Size int `json:"size"`
	// Type defines the type of the storage.
	Type string `json:"type"`
}
