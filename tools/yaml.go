package tools

import "gopkg.in/yaml.v3"

func ToYaml(d interface{}) []byte {
	data, err := yaml.Marshal(d)

	if err != nil {
		panic(err)
	}

	return data
}
