package schema

import "reflect"

type RequestBodyObject struct {
	Description string                     `yaml:"description"`
	Content     map[string]MediaTypeObject `yaml:"content"`
	Required    bool                       `yaml:"required,omitempty"`
}

type EncodingObject struct {
	ContentType string
	Headers     map[string]HeaderObject
}

func (s *Schema) AddRequestObject(obj any, rObj RequestBodyObject) {

	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return
	}

	s.Components.RequestBodies[t.String()] = rObj
}

func NewRequestObject() *RequestBodyObject {
	return &RequestBodyObject{}
}

func NewContent(mediaType string) map[string]MediaTypeObject {
	return map[string]MediaTypeObject{
		mediaType: {
			Schema: SchemaObject{},
		},
	}
}
