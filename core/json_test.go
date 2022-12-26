package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONDataFilter(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		fields  []string
		want    []byte
		wantErr error
	}{
		{
			name: "basic success",
			jsonStr: `
				{
					"field1":"value1",
					"field2":{
						"nestedField1":"nestedValue1",
						"nestedField2":"nestedValue2"
					},
					"field3": [
						{
							"nestedField1": "nestedValue1",
							"nestedField2": "nestedValue2"
						}
					],
					"field4": {
						"nestedField1": {
							"nestedField1": "nestedValue1",
							"nestedField2": "nestedValue2"
						}
					}
				}
			`,
			fields: []string{"field1", "field2.nestedField1", "field3.nestedField1", "field4.nestedField1.nestedField1"},
			want:   []byte(`{"field2":{"nestedField2":"nestedValue2"},"field3":[{"nestedField2":"nestedValue2"}],"field4":{"nestedField1":{"nestedField2":"nestedValue2"}}}`),
		},
		{
			name: "success with array base",
			jsonStr: `
				[
					{
						"field1": "value1",
						"field2": "value2"
					}
				]
			`,
			fields: []string{"field1"},
			want:   []byte(`[{"field2":"value2"}]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &JSONDataFilter{}
			got, err := f.FilterUsingStringFields([]byte(tt.jsonStr), tt.fields)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
