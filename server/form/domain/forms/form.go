package forms

import (
	"errors"
	"github.com/ashkan-maleki/go-form-builder/server/pkg/domain"
)

type CreateForm struct {
}

func (f CreateForm) Name() string {
	return "CreateForm"
}

type Form struct {
	Groups []FieldGroup
	Name   string
	ID     string
}

func (f *Form) Process(command domain.Command) ([]domain.Event, error) {
	var events []domain.Event
	switch cmd := command.(type) {
	case *CreateForm:
		return f.CreateForm(cmd)
	default:
		return events, errors.New("no such command")
	}
}

func (f *Form) CreateForm(cmd *CreateForm) ([]domain.Event, error) {
	var events []domain.Event
	return events, nil
}

type FieldGroup struct {
	Fields []Field
	Name   string
}

type FieldType struct {
	Name string
}

type FieldValue struct {
	Name string
}

type Field struct {
	Name  string
	Type  FieldType
	Value []FieldValue
}
