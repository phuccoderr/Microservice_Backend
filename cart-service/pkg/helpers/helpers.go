package helpers

import "encoding/json"

func StructToMap(data interface{}) (map[string]interface{}, error) {
	// Chuyển đổi struct thành JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Chuyển đổi JSON thành map
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
