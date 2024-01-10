package forms

type Form struct {
	Groups []FormGroup
	Name   string
	ID     string
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

type FormGroup struct {
	Fields []Field
	Name   string
}
