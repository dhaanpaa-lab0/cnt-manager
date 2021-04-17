package models

import "gopkg.in/yaml.v2"

type ContainerConfigFile struct {
	ContainerName string `yaml:"container_name"`
	ContainerLocation string `yaml:"container_location"`
	ContainerPort int `yaml:"container_port"`
}

func (receiver ContainerConfigFile) Save() string {
	b, err := yaml.Marshal(receiver)
	if err == nil {
		return ""
	} else {
		return string(b)
	}

}