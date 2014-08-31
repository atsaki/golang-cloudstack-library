package cloudstack

type Nic struct {
	Broadcasturi     NullString   `json:"broadcasturi"`
	Deviceid         ID           `json:"deviceid"`
	Gateway          NullString   `json:"gateway"`
	Id               ID           `json:"id"`
	Ip6address       NullString   `json:"ip6address"`
	Ip6cidr          NullString   `json:"ip6cidr"`
	Ip6gateway       NullString   `json:"ip6gateway"`
	Ipaddress        NullString   `json:"ipaddress"`
	Isdefault        NullBool     `json:"isdefault"`
	Isolationuri     NullString   `json:"isolationuri"`
	Macaddress       NullString   `json:"macaddress"`
	Netmask          NullString   `json:"netmask"`
	Networkid        ID           `json:"networkid"`
	Networkname      NullString   `json:"networkname"`
	Secondaryip      []NullString `json:"secondaryip"`
	Traffictype      NullString   `json:"traffictype"`
	Type             NullString   `json:"type"`
	Virtualmachineid ID           `json:"virtualmachineid"`
}
