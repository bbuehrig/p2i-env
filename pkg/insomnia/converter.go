package insomnia

import (
	"encoding/json"
)

type postmanEnvironment struct {
	ID                   string         `json:"id"`
	Name                 string         `json:"name"`
	Values               []postmanValue `json:"values"`
	PostmanVariableScope string         `json:"_postman_variable_scope"`
	PostmanExportedAt    string         `json:"_postman_exported_at"`
	PostmanExportedUsing string         `json:"_postman_exported_using"`
}

type postmanValue struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

func ConvertPostmanEnvironment(body []byte) (string, error) {
	postmanEnv := &postmanEnvironment{}
	err := json.Unmarshal(body, postmanEnv)
	if err != nil {
		return "", err
	}

	parameterMap := make(map[string]string)
	for _, value := range postmanEnv.Values {
		parameterMap[value.Key] = value.Value
	}

	result, _ := json.Marshal(parameterMap)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
