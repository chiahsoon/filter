package core

import (
	"encoding/json"
	"strings"
)

type JSONDataFilter struct{}

// FilterUsingStringFields filters the fields in the input JSON string based on the provided list of field names.
// The nested order of the fields are delimited by dots - ".".
// It returns a new JSON string with the specified fields removed.
func (f *JSONDataFilter) FilterUsingStringFields(data []byte, fields []string) ([]byte, error) {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	// Recursively remove the specified fields from the map
	f.removeFields(jsonData, fields)

	// Marshal the map back into a JSON string
	filteredJSON, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return filteredJSON, nil
}

// removeFields is a recursive function that removes the specified fields from the input.
// It also removes the fields from arrays/lists if they are present.
func (f *JSONDataFilter) removeFields(data interface{}, fields []string) {
	switch v := data.(type) {
	case map[string]interface{}:
		// If the data is a map, iterate through it and remove the specified fields
		for key := range v {
			for _, field := range fields {
				hasNestedFieldToRemove := strings.HasPrefix(field, key+".")
				isFieldToRemove := strings.EqualFold(key, field)
				if hasNestedFieldToRemove {
					// If the field has a prefix matching the current key, it means it is a nested field
					// Recursively remove the field from the nested map
					f.removeFields(v[key], []string{strings.TrimPrefix(field, key+".")})
					break
				} else if isFieldToRemove {
					// If the field matches the current key, it means it is a top-level field
					// Remove it from the map
					delete(v, key)
					break
				}
			}

			// If the value is a nested map, recursively remove the fields from it
			if nestedData, ok := v[key].(map[string]interface{}); ok {
				f.removeFields(nestedData, fields)
			}
		}
	case []interface{}:
		// If the data is a slice, iterate through it and remove the specified fields from the elements
		for i := range v {
			f.removeFields(v[i], fields)
		}
	}
}
