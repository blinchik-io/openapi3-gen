package schema

import (
	"reflect"
)

type SchemaObject struct {
	Type       string                `yaml:"type,omitempty"`
	Properties map[string]Properties `yaml:"properties,omitempty"`
}

func (s *Schema) AddSchemaObject(obj any) {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return
	}

	// Build a schema object for the entire struct
	sObj := SchemaObject{
		Type:       "object",
		Properties: map[string]Properties{},
	}

	for i := range t.NumField() {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		prop := fieldToSchema(field)
		if prop == nil {
			continue
		}

		propName := jsonName(field)
		sObj.Properties[propName] = *prop

	}

	s.Components.Schemas[t.String()] = sObj
}

type Properties struct {
	Type   string      `yaml:"type,omitempty"`
	Format string      `yaml:"format,omitempty"`
	Items  ItemsObject `yaml:"items,omitempty"`
}

type ItemsObject struct {
	Type string `yaml:"type,omitempty"`
	Ref  string `yaml:"$ref,omitempty"`
}
