package cloudstack


// AddAccountToProject represents the paramter of AddAccountToProject
type AddAccountToProjectParameter struct {
	// name of the account to be added to the project
	Account NullString
	// email to which invitation to the project is going to be sent
	Email NullString
	// id of the project to add the account to
	ProjectId ID
}

func NewAddAccountToProjectParameter(projectid string) (p *AddAccountToProjectParameter) {
	p = new(AddAccountToProjectParameter)
	p.ProjectId.Set(projectid)
	return p
}

// Adds acoount to a project
func (c *Client) AddAccountToProject(p *AddAccountToProjectParameter) (*Result, error) {
	obj, err := c.Request("addAccountToProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// EnableAccount represents the paramter of EnableAccount
type EnableAccountParameter struct {
	// Enables specified account.
	Account NullString
	// Enables specified account in this domain.
	DomainId ID
	// Account id
	Id ID
}

func NewEnableAccountParameter() (p *EnableAccountParameter) {
	p = new(EnableAccountParameter)
	return p
}

// Enables an account
func (c *Client) EnableAccount(p *EnableAccountParameter) (*Account, error) {
	obj, err := c.Request("enableAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Account), err
}


// DeleteAccountFromProject represents the paramter of DeleteAccountFromProject
type DeleteAccountFromProjectParameter struct {
	// name of the account to be removed from the project
	Account NullString
	// id of the project to remove the account from
	ProjectId ID
}

func NewDeleteAccountFromProjectParameter(account string, projectid string) (p *DeleteAccountFromProjectParameter) {
	p = new(DeleteAccountFromProjectParameter)
	p.Account.Set(account)
	p.ProjectId.Set(projectid)
	return p
}

// Deletes account from the project
func (c *Client) DeleteAccountFromProject(p *DeleteAccountFromProjectParameter) (*Result, error) {
	obj, err := c.Request("deleteAccountFromProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// UpdateAccount represents the paramter of UpdateAccount
type UpdateAccountParameter struct {
	// the current account name
	Account NullString
	// details for account used to store specific parameters
	AccountDetails map[string]string
	// the ID of the domain where the account exists
	DomainId ID
	// Account id
	Id ID
	// Network domain for the account's networks; empty string will update
	// domainName with NULL value
	NetworkDomain NullString
	// new name for the account
	NewName NullString
}

func NewUpdateAccountParameter(newname string) (p *UpdateAccountParameter) {
	p = new(UpdateAccountParameter)
	p.NewName.Set(newname)
	return p
}

// Updates account information for the authenticated user
func (c *Client) UpdateAccount(p *UpdateAccountParameter) (*Account, error) {
	obj, err := c.Request("updateAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Account), err
}


// LockAccount represents the paramter of LockAccount
type LockAccountParameter struct {
	// Locks the specified account.
	Account NullString
	// Locks the specified account on this domain.
	DomainId ID
}

func NewLockAccountParameter(account string, domainid string) (p *LockAccountParameter) {
	p = new(LockAccountParameter)
	p.Account.Set(account)
	p.DomainId.Set(domainid)
	return p
}

// Locks an account
func (c *Client) LockAccount(p *LockAccountParameter) (*Account, error) {
	obj, err := c.Request("lockAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Account), err
}


// CreateAccount represents the paramter of CreateAccount
type CreateAccountParameter struct {
	// Creates the user under the specified account. If no account is specified, the
	// username will be used as the account name.
	Account NullString
	// details for account used to store specific parameters
	AccountDetails map[string]string
	// Account UUID, required for adding account from external provisioning system
	AccountId ID
	// Type of the account.  Specify 0 for user, 1 for root admin, and 2 for domain
	// admin
	AccountType NullNumber
	// Creates the user under the specified domain.
	DomainId ID
	// email
	Email NullString
	// firstname
	FirstName NullString
	// lastname
	LastName NullString
	// Network domain for the account's networks
	NetworkDomain NullString
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

func NewCreateAccountParameter(accounttype int, email string, firstname string, lastname string, password string, username string) (p *CreateAccountParameter) {
	p = new(CreateAccountParameter)
	p.AccountType.Set(accounttype)
	p.Email.Set(email)
	p.FirstName.Set(firstname)
	p.LastName.Set(lastname)
	p.Password.Set(password)
	p.UserName.Set(username)
	return p
}

// Creates an account
func (c *Client) CreateAccount(p *CreateAccountParameter) (*Account, error) {
	obj, err := c.Request("createAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Account), err
}


// DeleteAccount represents the paramter of DeleteAccount
type DeleteAccountParameter struct {
	// Account id
	Id ID
}

func NewDeleteAccountParameter(id string) (p *DeleteAccountParameter) {
	p = new(DeleteAccountParameter)
	p.Id.Set(id)
	return p
}

// Deletes a account, and all users associated with this account
func (c *Client) DeleteAccount(p *DeleteAccountParameter) (*Result, error) {
	obj, err := c.Request("deleteAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// DisableAccount represents the paramter of DisableAccount
type DisableAccountParameter struct {
	// Disables specified account.
	Account NullString
	// Disables specified account in this domain.
	DomainId ID
	// Account id
	Id ID
	// If true, only lock the account; else disable the account
	Lock NullBool
}

func NewDisableAccountParameter(lock bool) (p *DisableAccountParameter) {
	p = new(DisableAccountParameter)
	p.Lock.Set(lock)
	return p
}

// Disables an account
func (c *Client) DisableAccount(p *DisableAccountParameter) (*Account, error) {
	obj, err := c.Request("disableAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Account), err
}


// ListAccounts represents the paramter of ListAccounts
type ListAccountsParameter struct {
	// list accounts by account type. Valid account types are 1 (admin), 2
	// (domain-admin), and 0 (user).
	AccountType NullNumber
	// list only resources belonging to the domain specified
	DomainId ID
	// list account by account ID
	Id ID
	// list accounts by cleanuprequred attribute (values are true or false)
	IsCleanupRequired NullBool
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// list account by account name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// list accounts by state. Valid states are enabled, disabled, and locked.
	State NullString
}

func NewListAccountsParameter() (p *ListAccountsParameter) {
	p = new(ListAccountsParameter)
	return p
}

// Lists accounts and provides detailed account information for listed accounts
func (c *Client) ListAccounts(p *ListAccountsParameter) ([]*Account, error) {
	obj, err := c.Request("listAccounts", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Account), err
}


// MarkDefaultZoneForAccount represents the paramter of MarkDefaultZoneForAccount
type MarkDefaultZoneForAccountParameter struct {
	// Name of the account that is to be marked.
	Account NullString
	// Marks the account that belongs to the specified domain.
	DomainId ID
	// The Zone ID with which the account is to be marked.
	ZoneId ID
}

func NewMarkDefaultZoneForAccountParameter(account string, domainid string, zoneid string) (p *MarkDefaultZoneForAccountParameter) {
	p = new(MarkDefaultZoneForAccountParameter)
	p.Account.Set(account)
	p.DomainId.Set(domainid)
	p.ZoneId.Set(zoneid)
	return p
}

// Marks a default zone for this account
func (c *Client) MarkDefaultZoneForAccount(p *MarkDefaultZoneForAccountParameter) (*Account, error) {
	obj, err := c.Request("markDefaultZoneForAccount", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Account), err
}

