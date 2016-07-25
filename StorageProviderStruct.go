package cloudstack

type StorageProvider struct {
	ResourceBase
	// the name of the storage provider
	Name NullString `json:"name"`
	// the type of the storage provider: primary or image provider
	Type NullString `json:"type"`
}
