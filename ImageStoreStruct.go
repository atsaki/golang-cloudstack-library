package cloudstack

type ImageStore struct {
	ResourceBase
	// the details of the image store
	Details []NullString `json:"details"`
	// the ID of the image store
	Id ID `json:"id"`
	// the name of the image store
	Name NullString `json:"name"`
	// the protocol of the image store
	Protocol NullString `json:"protocol"`
	// the provider name of the image store
	ProviderName NullString `json:"providername"`
	// the scope of the image store
	Scope NullString `json:"scope"`
	// the url of the image store
	Url NullString `json:"url"`
	// the Zone ID of the image store
	ZoneId ID `json:"zoneid"`
	// the Zone name of the image store
	ZoneName NullString `json:"zonename"`
}
