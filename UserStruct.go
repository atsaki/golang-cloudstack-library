package cloudstack

type User struct {
	ResourceBase
	// the account name of the user
	Account NullString `json:"account"`
	// the account ID of the user
	AccountId ID `json:"accountid"`
	// the account type of the user
	AccountType NullNumber `json:"accounttype"`
	// the api key of the user
	ApiKey NullString `json:"apikey"`
	// the date and time the user account was created
	Created NullString `json:"created"`
	// the domain name of the user
	Domain NullString `json:"domain"`
	// the domain ID of the user
	DomainId ID `json:"domainid"`
	// the user email address
	Email NullString `json:"email"`
	// the user firstname
	FirstName NullString `json:"firstname"`
	// the user ID
	Id ID `json:"id"`
	// the boolean value representing if the updating target is in caller's child
	// domain
	IsCallerChildDomain NullBool `json:"iscallerchilddomain"`
	// true if user is default, false otherwise
	IsDefault NullBool `json:"isdefault"`
	// the user lastname
	LastName NullString `json:"lastname"`
	// the secret key of the user
	SecretKey NullString `json:"secretkey"`
	// the user state
	State NullString `json:"state"`
	// the timezone user was created in
	TimeZone NullString `json:"timezone"`
	// the user name
	UserName NullString `json:"username"`
}

type LdapUser struct {
	ResourceBase
	// The user's domain
	Domain NullString `json:"domain"`
	// The user's email
	Email NullString `json:"email"`
	// The user's firstname
	FirstName NullString `json:"firstname"`
	// The user's lastname
	LastName NullString `json:"lastname"`
	// The user's principle
	Principal NullString `json:"principal"`
	// The user's username
	UserName NullString `json:"username"`
}

type UserKeys struct {
	ResourceBase
	// the api key of the registered user
	ApiKey NullString `json:"apikey"`
	// the secret key of the registered user
	SecretKey NullString `json:"secretkey"`
}
