package cloudstack

type LoadBalancerRuleInstance struct {
	ResourceBase
	// IP addresses of the vm set of lb rule
	LbVmIpAddresses []NullString `json:"lbvmipaddresses"`
	// the user vm set for lb rule
	LoadBalancerRuleInstance NullString `json:"loadbalancerruleinstance"`
}
