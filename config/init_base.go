package config

type baseCnf struct {
	Env         string `yaml:"dev"`
	ProjectName string `yaml:"project_name"`
	ProjectPort string `yaml:"project_port"`
}
