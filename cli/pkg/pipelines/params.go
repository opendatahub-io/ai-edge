/*
Copyright 2024. Open Data Hub Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pipelines

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Params struct {
	Name  string      `yaml:"name"`
	Value interface{} `yaml:"value"`
}

type RunParams struct {
	Params []Params `yaml:"params"`
}

func (p *RunParams) GetParamValue(name string) interface{} {
	for _, param := range p.Params {
		if param.Name == name {
			return param.Value
		}
	}
	return nil
}

// ToSimpleMap converts the RunParams struct to a simple map of string to interface{}
func (p *RunParams) ToSimpleMap() map[string]interface{} {
	params := make(map[string]interface{})
	for _, param := range p.Params {
		params[param.Name] = param.Value
	}
	return params
}

func ReadParams(paramsFile string) (*RunParams, error) {
	// Read YAML file
	data, err := os.ReadFile(paramsFile)
	if err != nil {
		log.Fatal("error reading file: ", err)
		return nil, err
	}

	// Unmarshal YAML to struct
	var runParams RunParams
	err = yaml.Unmarshal([]byte(data), &runParams)
	if err != nil {
		log.Fatal("error unmarshalling yaml: ", err)
		return nil, err
	}

	return &runParams, nil
}
