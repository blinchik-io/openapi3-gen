package schema

type PathItemObject struct {
	Ref         string `yaml:"$ref,omitempty" json:"$ref,omitempty"`
	Summary     string `yaml:"summary,omitempty" json:"summary,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	Get     *Operation `yaml:"get,omitempty" json:"get,omitempty"`
	Put     *Operation `yaml:"put,omitempty" json:"put,omitempty"`
	Post    *Operation `yaml:"post,omitempty" json:"post,omitempty"`
	Delete  *Operation `yaml:"delete,omitempty" json:"delete,omitempty"`
	Options *Operation `yaml:"options,omitempty" json:"options,omitempty"`
	Head    *Operation `yaml:"head,omitempty" json:"head,omitempty"`
	Patch   *Operation `yaml:"patch,omitempty" json:"patch,omitempty"`
	Trace   *Operation `yaml:"trace,omitempty" json:"trace,omitempty"`

	Servers []Server `yaml:"servers,omitempty" json:"servers,omitempty"`
	// Parameters []ParameterOrRef `yaml:"parameters,omitempty" json:"parameters,omitempty"`
}

type Operation struct {
	Tags        []string `yaml:"tags,omitempty" json:"tags,omitempty"`
	Summary     string   `yaml:"summary,omitempty" json:"summary,omitempty"`
	Description string   `yaml:"description,omitempty" json:"description,omitempty"`
	// ExternalDocs *ExternalDocumentation   `yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
	OperationID string `yaml:"operationId,omitempty" json:"operationId,omitempty"`
	// Parameters   []ParameterOrRef         `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	// RequestBody  *RequestBodyOrRef        `yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
	// Responses    Responses                `yaml:"responses" json:"responses"` // required field
	// Callbacks    map[string]CallbackOrRef `yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
	Deprecated bool `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	// Security     []SecurityRequirement    `yaml:"security,omitempty" json:"security,omitempty"`
	Servers []Server `yaml:"servers,omitempty" json:"servers,omitempty"`
}

func (s *Schema) AddPathItemObject(path string, pathItem PathItemObject) {
	s.Paths[path] = pathItem
}

func NewPathItemObject() *PathItemObject {
	return &PathItemObject{}
}
