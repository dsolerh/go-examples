package main

import (
	"strings"
	"text/template"
)

func buildTemplate(templateName, templateString string) (*template.Template, error) {
	return template.
		New(templateName).
		Parse(templateString)
}

func excecuteTemplate(t *template.Template, data any) (string, error) {
	var b strings.Builder
	err := t.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

const contextCodeTemplateStr = `
package {{.PackageName}}

import "bitbucket.org/whatwapp/wakama/analytics"

const {{.Name}}SchemaId = "{{.SchemaId}}"
const {{.Name}}Schema = "depot:" + {{.Name}}SchemaId

type {{.Name}}Props struct {
	{{.Props}}
}

type {{.Name}} struct {
	analytics.Schema
	{{.Name}}Props
}

func (ctx *{{.Name}}) GetId() string {
	return {{.Name}}SchemaId
}

func New{{.Name}}(props {{.Name}}Props) *{{.Name}} {
	return &{{.Name}}{
		Schema:              analytics.Schema{Schema: {{.Name}}Schema},
		{{.Name}}Props: props,
	}
}`

type contextCodeProps struct {
	PackageName string
	Name        string
	Props       string
	SchemaId    string
}

var contextCodeTemplate = must(buildTemplate("contextCode", contextCodeTemplateStr))

const eventCodeTemplateStr = `
package {{.PackageName}}

import "bitbucket.org/whatwapp/wakama/analytics"

const {{.Name}}EventSchema = "{{.SchemaId}}"

type {{.Name}}EventProps struct {
	{{.Props}}
}

type {{.Name}}Event struct {
	analytics.Schema
	analytics.AnalyticsEvent
	{{.Name}}EventProps
}

func New{{.Name}}Event(props {{.Name}}EventProps) *{{.Name}}Event {
	return &{{.Name}}Event{
		Schema: analytics.Schema{Schema: {{.Name}}EventSchema},
		AnalyticsEvent: analytics.AnalyticsEvent{
			Type:   "{{.Type}}",
 			Action: "{{.Action}}",
		},
		{{.Name}}EventProps: props,
	}
}`

type eventCodeProps struct {
	PackageName string
	Name        string
	Props       string
	SchemaId    string
	Type        string
	Action      string
}

var eventCodeTemplate = must(buildTemplate("eventCode", eventCodeTemplateStr))
