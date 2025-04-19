package schema

import "reflect"

type RequestBodyObject struct {
	Description string `yaml:"description,omitempty"`
	Content     MediaTypeObject
	Required    bool `yaml:"required,omitempty"`
}

type MediaTypeObject struct {
	Schema   SchemaObject              `yaml:"schema,omitempty"`
	Example  any                       `yaml:"example,omitempty"`
	Encoding map[string]EncodingObject `yaml:"encoding,omitempty"`
}

type EncodingObject struct {
	ContentType string
	Headers     map[string]HeaderObject
}

type HeaderObject struct {
	Description string `yaml:"description,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
	Deprecated  bool   `yaml:"deprecated,omitempty"`
}

func (s *Schema) AddRequestObject(obj any) {

	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return
	}

	rObj := RequestBodyObject{}
	s.Components.RequestBodies[t.String()] = rObj
}
