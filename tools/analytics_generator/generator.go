package main

import (
	"fmt"
	"go/format"
	"strings"
)

func generate(schemaStr string, packageName string) (string, string, error) {
	schema, err := Json[Schema](schemaStr)
	if err != nil {
		return "", "", err
	}

	if strings.HasSuffix(schema.Self.Name, "_ctx") {
		return generateContext(schema, packageName)
	}
	return generateEvent(schema, packageName)
}

func generateContext(schema Schema, packageName string) (string, string, error) {
	props := make([]string, len(schema.Properties))
	for propName, propSchema := range schema.Properties {
		props = append(props, fmt.Sprintf("%s %s `json:\"%s\"`", formatPascalCase(propName), typesMap[propSchema.Type], propName))
	}
	code, err := excecuteTemplate(contextCodeTemplate, contextCodeProps{
		PackageName: packageName,
		Name:        formatPascalCase(schema.Self.Name),
		Props:       strings.Join(props, "\n"),
		SchemaId: fmt.Sprintf(
			"%s/%s/%s/%s",
			schema.Self.Namespace,
			schema.Self.Category,
			schema.Self.Name,
			schema.Self.Version,
		),
	})
	if err != nil {
		return "", "", err
	}

	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		return "", "", err
	}

	return string(formattedCode), fmt.Sprintf("%s.go", schema.Self.Name), nil
}

func generateEvent(schema Schema, packageName string) (string, string, error) {
	props := make([]string, len(schema.Properties))
	_type := ""
	_action := ""
	for propName, propSchema := range schema.Properties {
		if propName == "type" {
			_type = propSchema.Const
			continue
		}
		if propName == "action" {
			_action = propSchema.Const
			continue
		}
		props = append(props, fmt.Sprintf("%s %s `json:\"%s\"`", formatPascalCase(propName), typesMap[propSchema.Type], propName))
	}
	code, err := excecuteTemplate(eventCodeTemplate, eventCodeProps{
		PackageName: packageName,
		Name:        formatPascalCase(schema.Self.Name),
		Props:       strings.Join(props, "\n"),
		SchemaId:    schema.Id,
		Type:        _type,
		Action:      _action,
	})
	if err != nil {
		return "", "", err
	}

	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		return "", "", err
	}

	return string(formattedCode), fmt.Sprintf("%s_event.go", schema.Self.Name), nil
}
