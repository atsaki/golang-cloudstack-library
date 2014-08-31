package cloudstack

type Serviceoffering struct {
	Cpunumber                 NullInt64  `json:"cpunumber"`
	Cpuspeed                  NullInt64  `json:"cpuspeed"`
	Created                   NullString `json:"created"`
	Defaultuse                NullBool   `json:"defaultuse"`
	Deploymentplanner         NullString `json:"deploymentplanner"`
	DiskBytesReadRate         NullInt64  `json:"diskbytesreadrate"`
	DiskBytesWriteRate        NullInt64  `json:"diskbyteswriterate"`
	DiskIopsReadRate          NullInt64  `json:"diskiopsreadrate"`
	DiskIopsWriteRate         NullInt64  `json:"diskiopswriterate"`
	Displaytext               NullString `json:"displaytext"`
	Domain                    NullString `json:"domain"`
	Domainid                  ID         `json:"domainid"`
	Hosttags                  NullString `json:"hosttags"`
	Hypervisorsnapshotreserve NullInt64  `json:"hypervisorsnapshotreserve"`
	Id                        ID         `json:"id"`
	Iscustomized              NullBool   `json:"iscustomized"`
	Iscustomizediops          NullBool   `json:"iscustomizediops"`
	Issystem                  NullBool   `json:"issystem"`
	Isvolatile                NullBool   `json:"isvolatile"`
	Limitcpuuse               NullBool   `json:"limitcpuuse"`
	Maxiops                   NullInt64  `json:"maxiops"`
	Memory                    NullInt64  `json:"memory"`
	Miniops                   NullInt64  `json:"miniops"`
	Name                      NullString `json:"name"`
	Networkrate               NullInt64  `json:"networkrate"`
	Offerha                   NullBool   `json:"offerha"`
	Serviceofferingdetails    NullString `json:"serviceofferingdetails"`
	Storagetype               NullString `json:"storagetype"`
	Systemvmtype              NullString `json:"systemvmtype"`
	Tags                      NullString `json:"tags"`
}
