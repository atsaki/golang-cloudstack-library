package cloudstack

// ImportLdapUsers represents the paramter of ImportLdapUsers
type ImportLdapUsersParameter struct {
	// Creates the user under the specified account. If no account is specified, the
	// username will be used as the account name.
	Account NullString
	// details for account used to store specific parameters
	AccountDetails map[string]string
	// Type of the account.  Specify 0 for user, 1 for root admin, and 2 for domain
	// admin
	AccountType NullNumber
	// Specifies the domain to which the ldap users are to be imported. If no domain
	// is specified, a domain will created using group parameter. If the group is
	// also not specified, a domain name based on the OU information will be
	// created. If no OU hierarchy exists, will be defaulted to ROOT domain
	DomainId ID
	// Specifies the group name from which the ldap users are to be imported. If no
	// group is specified, all the users will be imported.
	Group NullString
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
	// Specifies a timezone for this command. For more information on the timezone
	// parameter, see Time Zone Format.
	TimeZone NullString
}

func NewImportLdapUsersParameter(accounttype int) (p *ImportLdapUsersParameter) {
	p = new(ImportLdapUsersParameter)
	p.AccountType.Set(accounttype)
	return p
}

// Import LDAP users
func (c *Client) ImportLdapUsers(p *ImportLdapUsersParameter) (*LdapUser, error) {
	obj, err := c.Request("importLdapUsers", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*LdapUser), err
}

// ListLdapUsers represents the paramter of ListLdapUsers
type ListLdapUsersParameter struct {
	// List by keyword
	Keyword NullString
	// Determines whether all ldap users are returned or just non-cloudstack users
	ListType NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListLdapUsersParameter() (p *ListLdapUsersParameter) {
	p = new(ListLdapUsersParameter)
	return p
}

// Lists all LDAP Users
func (c *Client) ListLdapUsers(p *ListLdapUsersParameter) ([]*LdapUser, error) {
	obj, err := c.Request("listLdapUsers", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*LdapUser), err
}

// EnableUser represents the paramter of EnableUser
type EnableUserParameter struct {
	// Enables user by user ID.
	Id ID
}

func NewEnableUserParameter(id string) (p *EnableUserParameter) {
	p = new(EnableUserParameter)
	p.Id.Set(id)
	return p
}

// Enables a user account
func (c *Client) EnableUser(p *EnableUserParameter) (*User, error) {
	obj, err := c.Request("enableUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*User), err
}

// CreateUser represents the paramter of CreateUser
type CreateUserParameter struct {
	// Creates the user under the specified account. If no account is specified, the
	// username will be used as the account name.
	Account NullString
	// Creates the user under the specified domain. Has to be accompanied with the
	// account parameter
	DomainId ID
	// email
	Email NullString
	// firstname
	FirstName NullString
	// lastname
	LastName NullString
	// Clear text password (Default hashed to SHA256SALT). If you wish to use any
	// other hashing algorithm, you would need to write a custom authentication
	// adapter See Docs section.
	Password NullString
	// Specifies a timezone for this command. For more information on the timezone
	// parameter, see Time Zone Format.
	TimeZone NullString
	// User UUID, required for adding account from external provisioning system
	UserId ID
	// Unique username.
	UserName NullString
}

func NewCreateUserParameter(account string, email string, firstname string, lastname string, password string, username string) (p *CreateUserParameter) {
	p = new(CreateUserParameter)
	p.Account.Set(account)
	p.Email.Set(email)
	p.FirstName.Set(firstname)
	p.LastName.Set(lastname)
	p.Password.Set(password)
	p.UserName.Set(username)
	return p
}

// Creates a user for an account that already exists
func (c *Client) CreateUser(p *CreateUserParameter) (*User, error) {
	obj, err := c.Request("createUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*User), err
}

// DeleteUser represents the paramter of DeleteUser
type DeleteUserParameter struct {
	// id of the user to be deleted
	Id ID
}

func NewDeleteUserParameter(id string) (p *DeleteUserParameter) {
	p = new(DeleteUserParameter)
	p.Id.Set(id)
	return p
}

// Deletes a user for an account
func (c *Client) DeleteUser(p *DeleteUserParameter) (*Result, error) {
	obj, err := c.Request("deleteUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// GetUser represents the paramter of GetUser
type GetUserParameter struct {
	// API key of the user
	UserApiKey NullString
}

func NewGetUserParameter(userapikey string) (p *GetUserParameter) {
	p = new(GetUserParameter)
	p.UserApiKey.Set(userapikey)
	return p
}

// Find user account by API key
func (c *Client) GetUser(p *GetUserParameter) (*User, error) {
	obj, err := c.Request("getUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*User), err
}

// UpdateUser represents the paramter of UpdateUser
type UpdateUserParameter struct {
	// email
	Email NullString
	// first name
	FirstName NullString
	// User uuid
	Id ID
	// last name
	LastName NullString
	// Clear text password (default hashed to SHA256SALT). If you wish to use any
	// other hasing algorithm, you would need to write a custom authentication
	// adapter
	Password NullString
	// Specifies a timezone for this command. For more information on the timezone
	// parameter, see Time Zone Format.
	TimeZone NullString
	// The API key for the user. Must be specified with userSecretKey
	UserApiKey NullString
	// Unique username
	UserName NullString
	// The secret key for the user. Must be specified with userSecretKey
	UserSecretKey NullString
}

func NewUpdateUserParameter(id string) (p *UpdateUserParameter) {
	p = new(UpdateUserParameter)
	p.Id.Set(id)
	return p
}

// Updates a user account
func (c *Client) UpdateUser(p *UpdateUserParameter) (*User, error) {
	obj, err := c.Request("updateUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*User), err
}

// DisableUser represents the paramter of DisableUser
type DisableUserParameter struct {
	// Disables user by user ID.
	Id ID
}

func NewDisableUserParameter(id string) (p *DisableUserParameter) {
	p = new(DisableUserParameter)
	p.Id.Set(id)
	return p
}

// Disables a user account
func (c *Client) DisableUser(p *DisableUserParameter) (*User, error) {
	obj, err := c.Request("disableUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*User), err
}

// RegisterUserKeys represents the paramter of RegisterUserKeys
type RegisterUserKeysParameter struct {
	// User id
	Id ID
}

func NewRegisterUserKeysParameter(id string) (p *RegisterUserKeysParameter) {
	p = new(RegisterUserKeysParameter)
	p.Id.Set(id)
	return p
}

// This command allows a user to register for the developer API, returning a
// secret key and an API key. This request is made through the integration API
// port, so it is a privileged command and must be made on behalf of a user. It
// is up to the implementer just how the username and password are entered, and
// then how that translates to an integration API request. Both secret key and
// API key should be returned to the user
func (c *Client) RegisterUserKeys(p *RegisterUserKeysParameter) (*UserKeys, error) {
	obj, err := c.Request("registerUserKeys", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*UserKeys), err
}

// LockUser represents the paramter of LockUser
type LockUserParameter struct {
	// Locks user by user ID.
	Id ID
}

func NewLockUserParameter(id string) (p *LockUserParameter) {
	p = new(LockUserParameter)
	p.Id.Set(id)
	return p
}

// Locks a user account
func (c *Client) LockUser(p *LockUserParameter) (*User, error) {
	obj, err := c.Request("lockUser", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*User), err
}

// ListUsers represents the paramter of ListUsers
type ListUsersParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// List users by account type. Valid types include admin, domain-admin,
	// read-only-admin, or user.
	AccountType NullNumber
	// list only resources belonging to the domain specified
	DomainId ID
	// List user by ID.
	Id ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll  NullBool
	Page     NullNumber
	PageSize NullNumber
	// List users by state of the user account.
	State NullString
	// List user by the username
	UserName NullString
}

func NewListUsersParameter() (p *ListUsersParameter) {
	p = new(ListUsersParameter)
	return p
}

// Lists user accounts
func (c *Client) ListUsers(p *ListUsersParameter) ([]*User, error) {
	obj, err := c.Request("listUsers", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*User), err
}
