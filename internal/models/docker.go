package models

type Service struct {
	Image       string            `yaml:"image" json:"image"`
	Ports       []string          `yaml:"ports" json:"ports"`
	Environment map[string]string `yaml:"environment" json:"environment"`
}

type DockerCompose struct {
	Services map[string]Service `yaml:"services" json:"services"`
}
