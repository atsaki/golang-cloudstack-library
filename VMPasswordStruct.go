package cloudstack

type VMPassword struct {
	// The base64 encoded encrypted password of the VM
	EncryptedPassword NullString `json:"encryptedpassword"`
	// CloudStack API Client
	Client *Client
}
