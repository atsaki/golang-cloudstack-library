package cloudstack

type Sshkeypair struct {
	Fingerprint NullString `json:"fingerprint"`
	Name        NullString `json:"name"`
}
