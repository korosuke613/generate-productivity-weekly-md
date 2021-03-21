package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"text/template"
)

type Input map[string]interface{}

type Tempura struct {
	Template         string
	TemplateFilePath string
	Input            Input
}

func (t *Tempura) SetInputFromString(jsonStr string) error {
	if err := json.Unmarshal([]byte(jsonStr), &t.Input); err != nil {
		return err
	}
	return nil
}

func (t *Tempura) SetInputFromJSON(jsonPath string) error {
	raw, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	if err := t.SetInputFromString(string(raw)); err != nil {
		return err
	}

	return nil
}

func (t *Tempura) getTemplate() (*template.Template, error) {
	if t.Template != "" {
		temp, err := template.New("").Parse(t.Template)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}

	if t.TemplateFilePath != "" {
		temp, err := template.ParseFiles(t.TemplateFilePath)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}

	return nil, errors.New("you should set template")
}

func (t *Tempura) Fill(output *string) error {
	temp, err := t.getTemplate()
	if err != nil {
		return err
	}

	var b bytes.Buffer
	if err := temp.Execute(&b, t.Input); err != nil {
		return err
	}

	*output = b.String()

	return nil
}
