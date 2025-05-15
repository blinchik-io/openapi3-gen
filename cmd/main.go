package main

import (
	"fmt"
	"log"
	"time"

	"github.com/blinchik-io/openapi3-generator/internal/schema"
	"github.com/goccy/go-yaml"
)

// // createCrimeRequest
// type createCrimeRequest struct {
// 	Location       string                 `json:"location" openapi:"type=string;example=Central Park, New York"`
// 	Date           *time.Time             `json:"date" openapi:"type=string;format=date-time;example=2023-01-15T14:30:00Z"`
// 	Perpetrator    []string               `json:"perpetrator" openapi:"type=array;items=string;example=[\"AAA\",\"BBB\",\"XXX\"]"`
// 	Victim         []string               `json:"victim" openapi:"type=array;items=string;example=[\"TTT\",\"WWW\",\"QQQ\"]"`
// 	Category       string                 `json:"category" openapi:"type=string;example=killing"`
// 	ReferenceNotes *string                `json:"reference_notes" openapi:"type=string;example=Police report #12345"`
// 	OtherNotes     *string                `json:"other_notes" openapi:"type=string;example=Witness accounts vary"`
// 	Victims        *[]createVictimRequest `json:"victims" openapi:"type=array;items=main.createVictimRequest;ref=true;example=[{\"full_name\":\"John Doe\"}]"`
// }
//
// // createVictimRequest
// type createVictimRequest struct {
// 	FullName        string  `json:"full_name" openapi:"type=string;example=firstname lastname"`
// 	Age             *int    `json:"age" openapi:"type=integer;example=23"`
// 	Religion        string  `json:"religion" openapi:"type=string;example=A"`
// 	Ethnicity       string  `json:"ethnicity" openapi:"type=string;example=A"`
// 	Gender          string  `json:"gender" openapi:"type=string;example=M"`
// 	ReferenceNotes  *string `json:"reference_notes" openapi:"type=string;example=sometext"`
// 	AdditionalNotes *string `json:"additional_notes" openapi:"type=string;example=sometext"`
// }

type registerUserRequest struct {
	Name     string `json:"name" openapi:"type=string"`
	Email    string `json:"email" openapi:"type=string"`
	Password string `json:"password" openapi:"type=string;format=password"`
}

type User struct {
	ID        int64     `json:"id" openapi:"type=string"`
	CreatedAt time.Time `json:"created_at" openapi:"type=string;format=date-time"`
	Name      string    `json:"name" openapi:"type=string"`
	Email     string    `json:"email" openapi:"type=string"`
	Activated bool      `json:"activated" openapi:"type=boolean"`
	Version   int       `json:"-"`
}

func main() {

	s := schema.NewSchema()
	s.OpenAPI = "3.0.3"
	s.Info = schema.Info{
		Title:   "My API",
		Version: "v1.0.0",
	}

	s.Servers = []schema.Server{
		{
			Url:         "{scheme}://localhost:4000/v1",
			Description: "Test 3en server.",
			Variables: map[string]schema.Variable{
				"scheme": {
					Enum:        []string{"http", "https"},
					Default:     "http",
					Description: "The supported protocols",
				},
			},
		},
	}

	s.AddSchemaObject(registerUserRequest{})
	s.AddSchemaObject(User{})

	userResponseObj := schema.NewResponsesObject()
	userResponseObj.Response.Description = "test"
	s.AddResponsesObject("202", User{}, *userResponseObj)

	userRequestObj := schema.NewRequestObject()
	userRequestObj.Description = "test"
	userRequestObj.Content = schema.NewContent("*/*")
	s.AddRequestObject(registerUserRequest{}, *userRequestObj)

	// s.AddSchemaObject(createCrimeRequest{})
	// s.AddSchemaObject(createVictimRequest{})

	// rObj := schema.NewRequestObject()
	// rObj.Description = "Create a crime"
	// rObj.Content = schema.NewContent("application/json")
	// s.AddRequestObject(createVictimRequest{}, *rObj)

	// pathItem := schema.NewPathItemObject()
	// pathItem.Summary = "List Crimes"
	// pathItem.Description = "Get a paginated list of crimes with filtering and sorting capabilities"
	//
	// pathItem.Get = &schema.Operation{
	// 	Tags: []string{"crimes"},
	// }
	//
	// s.AddPathItemObject("/crimes", *pathItem)

	yamlData, err := yaml.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(yamlData))
}
