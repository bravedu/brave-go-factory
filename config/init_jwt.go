package config

type jwt struct {
	JwtSecret  string `yaml:"jwt_secret"`
	JwtUserKey string `yaml:"jwt_user_key"`
}
