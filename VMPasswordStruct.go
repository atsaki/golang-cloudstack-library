package cloudstack

type VMPassword struct {
	ResourceBase
	// The base64 encoded encrypted password of the VM
	EncryptedPassword NullString `json:"encryptedpassword"`
}
