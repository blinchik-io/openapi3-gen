package schema

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type ParsedTags struct {
	Typ     string
	Format  string
	Example string
	Items   string
	Ref     bool
}

func parseJSONTagName(tag string) string {
	if tag == "" {
		return ""
	}
	parts := strings.Split(tag, ",")
	return parts[0]
}

// parseOpenapiTag is a simple parser for your "type=...,format=...,example=..." tags
func parseOpenapiTag(tag string) ParsedTags {
	var tagValues ParsedTags

	// Example tag: "type=string;format=date-time;example=2023-04-08"
	segments := strings.Split(tag, ";")
	for _, seg := range segments {
		kv := strings.SplitN(seg, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		val := strings.TrimSpace(kv[1])

		switch key {
		case "type":
			tagValues.Typ = val
		case "format":
			tagValues.Format = val
		case "items":
			tagValues.Items = val
		case "ref":
			ref, err := strconv.ParseBool(val)
			if err != nil {
				log.Fatal(err)
			}
			tagValues.Ref = ref
		case "example":
			tagValues.Example = val
		}
	}
	return tagValues
}

func jsonName(f reflect.StructField) string {
	tag := f.Tag.Get("json")
	if name := parseJSONTagName(tag); name != "" {
		return name
	}
	return f.Name
}

func fieldToSchema(f reflect.StructField) *Properties {
	openapi := f.Tag.Get("openapi")
	if openapi == "" {
		return nil
	}

	var prop Properties

	tag := parseOpenapiTag(openapi)

	prop.Type = tag.Typ

	if len(tag.Format) != 0 {
		prop.Format = tag.Format
	}

	if tag.Ref {
		prop.Items.Ref = fmt.Sprintf("#/components/schemas/%s", tag.Items)
	} else {
		prop.Items.Type = tag.Items
	}

	return &prop

}
