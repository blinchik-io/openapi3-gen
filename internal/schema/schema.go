package schema

type Schema struct {
	OpenAPI    string                    `yaml:"openapi"`
	Info       Info                      `yaml:"info"`
	Servers    []Server                  `yaml:"servers"`
	Components *Components               `yaml:"components,omitempty"`
	Paths      map[string]PathItemObject `yaml:"paths,omitempty"`
}

func NewSchema() *Schema {
	return &Schema{
		Components: &Components{
			Schemas:       make(map[string]SchemaObject),
			RequestBodies: make(map[string]RequestBodyObject),
			Responses:     make(map[string]ResponsesObject),
		},
		Paths: make(map[string]PathItemObject),
	}
}

type Info struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

type Server struct {
	Url         string              `yaml:"url"`
	Description string              `yaml:"description,omitempty"`
	Variables   map[string]Variable `yaml:"variables,omitempty"`
}

type Variable struct {
	Enum        []string `yaml:"enum,omitempty"`
	Default     string   `yaml:"default"`
	Description string   `yaml:"description,omitempty"`
}

type Components struct {
	Schemas       map[string]SchemaObject      `yaml:"schemas,omitempty"`
	RequestBodies map[string]RequestBodyObject `yaml:"requestBodies,omitempty"`
	Responses     map[string]ResponsesObject   `yaml:"responses,omitempty"`
}
