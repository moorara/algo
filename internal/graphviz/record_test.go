package graphviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRecord(t *testing.T) {
	tests := []struct {
		name          string
		fields        []Field
		expectedLabel string
	}{
		{
			name: "OK",
			fields: []Field{
				&simpleField{Name: "left", Label: "Left"},
				&complexField{
					Record: Record{
						Fields: []Field{
							&simpleField{Label: "a"},
							&complexField{
								Record: Record{
									Fields: []Field{
										&simpleField{Label: "b"},
										&simpleField{Name: "middle", Label: "c"},
										&simpleField{Label: "d"},
									},
								},
							},
							&simpleField{Label: "e"},
						},
					},
				},
				&simpleField{Name: "right", Label: "Right"},
			},
			expectedLabel: "<left>Left | { a | { b | <middle>c | d } | e } | <right>Right",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRecord(tc.fields...)
			assert.Equal(t, tc.expectedLabel, r.Label())
		})
	}
}

func TestRecord_AddField(t *testing.T) {
	tests := []struct {
		name          string
		fields        []Field
		expectedLabel string
	}{
		{
			name: "OK",
			fields: []Field{
				&simpleField{Name: "left", Label: "Left"},
				&complexField{
					Record: Record{
						Fields: []Field{
							&simpleField{Label: "a"},
							&complexField{
								Record: Record{
									Fields: []Field{
										&simpleField{Label: "b"},
										&simpleField{Name: "middle", Label: "c"},
										&simpleField{Label: "d"},
									},
								},
							},
							&simpleField{Label: "e"},
						},
					},
				},
				&simpleField{Name: "right", Label: "Right"},
			},
			expectedLabel: "<left>Left | { a | { b | <middle>c | d } | e } | <right>Right",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRecord()
			for _, f := range tc.fields {
				r.AddField(f)
			}

			assert.Equal(t, tc.expectedLabel, r.Label())
		})
	}
}

func TestNewSimpleField(t *testing.T) {
	tests := []struct {
		name            string
		fieldName       string
		label           string
		expectedDotCode string
	}{
		{
			name:            "Left",
			fieldName:       "l",
			label:           "left",
			expectedDotCode: "<l>left",
		},
		{
			name:            "Right",
			fieldName:       "r",
			label:           "right",
			expectedDotCode: "<r>right",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := NewSimpleField(tc.fieldName, tc.label)
			assert.Equal(t, tc.expectedDotCode, f.dotCode())
		})
	}
}

func TestNewComplexField(t *testing.T) {
	tests := []struct {
		name            string
		record          Record
		expectedDotCode string
	}{
		{
			name: "OK",
			record: Record{
				Fields: []Field{
					&simpleField{Label: "a"},
					&complexField{
						Record: Record{
							Fields: []Field{
								&simpleField{Label: "b"},
								&simpleField{Name: "id", Label: "c"},
								&simpleField{Label: "d"},
							},
						},
					},
					&simpleField{Label: "e"},
				},
			},
			expectedDotCode: "{ a | { b | <id>c | d } | e }",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := NewComplexField(tc.record)
			assert.Equal(t, tc.expectedDotCode, f.dotCode())
		})
	}
}
