package config

type wxweb struct {
	AccessTokenUrl string `yaml:"access_token_url"`
	UserinfoUrl    string `yaml:"userinfo_url"`
}

type wxh5 struct {
	AccessTokenUrl       string `yaml:"access_token_url"`
	UserinfoUrl          string `yaml:"userinfo_url"`
	PublicAccessTokenUrl string `yaml:"public_access_token_url"`
	TicketUrl            string `yaml:"ticket_url"`
}
