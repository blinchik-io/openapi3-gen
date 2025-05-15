package schema

type HeaderObject struct {
	Ref         string `yaml:"$ref,omitempty" json:"$ref,omitempty"`
	Description string `yaml:"description,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
	Deprecated  bool   `yaml:"deprecated,omitempty"`
}

type MediaTypeObject struct {
	Schema   SchemaObject              `yaml:"schema,omitempty"`
	Example  any                       `yaml:"example,omitempty"`
	Encoding map[string]EncodingObject `yaml:"encoding,omitempty"`
}
