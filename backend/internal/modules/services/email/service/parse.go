package service

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
)

type TemplateName string

const (
	Verification TemplateName = "verification"
)

var templateByType = map[TemplateName]string{
	Verification: verificationTemplate,
}

//go:embed templates/verification.tmpl
var verificationTemplate string

func parseTemplate(templateName TemplateName, data any) (string, error) {
	tmplString, ok := templateByType[templateName]
	if !ok {
		return "", fmt.Errorf("неизвестный шаблон: %s", templateName)
	}

	tmpl, err := template.New(string(templateName)).Parse(tmplString)
	if err != nil {
		return "", err
	}

	var parsedBuffer bytes.Buffer
	if err := tmpl.Execute(&parsedBuffer, data); err != nil {
		return "", err
	}

	return parsedBuffer.String(), nil
}
