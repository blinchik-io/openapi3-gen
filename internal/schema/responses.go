package schema

import "reflect"

type ResponsesObject struct {
	Ref      string    `yaml:"$ref,omitempty" json:"$ref,omitempty"`
	Response *Response `yaml:",inline,omitempty" json:",inline,omitempty"`
}

type Response struct {
	Description string                     `yaml:"description" json:"description"`
	Headers     map[string]HeaderObject    `yaml:"headers,omitempty" json:"headers,omitempty"`
	Content     map[string]MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
}

func (s *Schema) AddResponsesObject(resp string, obj any, respObj ResponsesObject) {

	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return
	}

	s.Components.Responses[t.String()+"_"+resp] = respObj
}

func NewResponsesObject() *ResponsesObject {
	return &ResponsesObject{
		Response: &Response{},
	}
}
