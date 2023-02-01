package accounts

type Accounts struct {
	Accounts []Account
}

type Account struct {
	Id          string `yaml:"id"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Role        string `yaml:"role"`
}
