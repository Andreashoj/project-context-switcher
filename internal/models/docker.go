package models

type Service struct {
	Image       string            `yaml:"image"`
	Ports       []string          `yaml:"ports"`
	Environment map[string]string `yaml:"environment"`
}

type DockerCompose struct {
	Services map[string]Service `yaml:"services"`
}
