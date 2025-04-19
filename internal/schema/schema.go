package schema

type Schema struct {
	OpenAPI    string      `yaml:"openapi"`
	Info       Info        `yaml:"info"`
	Components *Components `yaml:"components,omitempty"`
}

func NewSchema() *Schema {
	return &Schema{
		Components: &Components{
			Schemas:       make(map[string]SchemaObject),
			RequestBodies: make(map[string]RequestBodyObject),
		},
	}
}

type Info struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

type Components struct {
	Schemas       map[string]SchemaObject      `yaml:"schemas,omitempty"`
	RequestBodies map[string]RequestBodyObject `yaml:"requestBodies,omitempty"`
}
