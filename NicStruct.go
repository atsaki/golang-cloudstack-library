package cloudstack

type Nic struct {
	ResourceBase
	// the broadcast uri of the nic
	BroadcastUri NullString `json:"broadcasturi"`
	// device id for the network when plugged into the virtual machine
	DeviceId ID `json:"deviceid"`
	// the gateway of the nic
	Gateway NullString `json:"gateway"`
	// the ID of the nic
	Id ID `json:"id"`
	// the IPv6 address of network
	Ip6Address NullString `json:"ip6address"`
	// the cidr of IPv6 network
	Ip6Cidr NullString `json:"ip6cidr"`
	// the gateway of IPv6 network
	Ip6Gateway NullString `json:"ip6gateway"`
	// the ip address of the nic
	IpAddress NullString `json:"ipaddress"`
	// true if nic is default, false otherwise
	IsDefault NullBool `json:"isdefault"`
	// the isolation uri of the nic
	IsolationUri NullString `json:"isolationuri"`
	// true if nic is default, false otherwise
	MacAddress NullString `json:"macaddress"`
	// the netmask of the nic
	Netmask NullString `json:"netmask"`
	// the ID of the corresponding network
	NetworkId ID `json:"networkid"`
	// the name of the corresponding network
	NetworkName NullString `json:"networkname"`
	// the Secondary ipv4 addr of nic
	SecondaryIp []NullString `json:"secondaryip"`
	// the traffic type of the nic
	TrafficType NullString `json:"traffictype"`
	// the type of the nic
	Type NullString `json:"type"`
	// Id of the vm to which the nic belongs
	VirtualMachineId ID `json:"virtualmachineid"`
}
