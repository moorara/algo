package graphviz

import (
	"bytes"
	"strings"
)

// Record represents a record.
type Record struct {
	Fields []Field
}

// NewRecord creates a new record.
func NewRecord(fields ...Field) Record {
	return Record{
		Fields: fields,
	}
}

// AddField adds a new field to the record.
func (r *Record) AddField(f Field) {
	r.Fields = append(r.Fields, f)
}

// Label returns the label for a node with record or Mrecord shape.
func (r *Record) Label() string {
	fields := []string{}
	for _, f := range r.Fields {
		fields = append(fields, f.dotCode())
	}

	return strings.Join(fields, " | ")
}

// Field represents a field.
type Field interface {
	dotCode() string
}

// simpleField represents a simple field.
type simpleField struct {
	Name  string
	Label string
}

// NewSimpleField creates a new simple field.
func NewSimpleField(name, label string) Field {
	return &simpleField{
		Name:  name,
		Label: label,
	}
}

// DotCode generates the Graphviz dot language code.
func (f *simpleField) dotCode() string {
	buf := new(bytes.Buffer)

	if f.Name != "" {
		buf.WriteString("<" + f.Name + ">")
	}

	if f.Label != "" {
		buf.WriteString(f.Label)
	}

	return buf.String()
}

// complexField represents a complex field.
type complexField struct {
	Record Record
}

// NewComplexField creates a new complex field.
func NewComplexField(record Record) Field {
	return &complexField{
		Record: record,
	}
}

// DotCode generates the Graphviz dot language code.
func (f *complexField) dotCode() string {
	buf := new(bytes.Buffer)

	buf.WriteString("{ ")
	buf.WriteString(f.Record.Label())
	buf.WriteString(" }")

	return buf.String()
}
