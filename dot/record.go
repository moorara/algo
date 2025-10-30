package dot

import (
	"bytes"
	"strings"
)

var replacer = strings.NewReplacer(
	`|`, `\|`,
	`{`, `\{`,
	`}`, `\}`,
	`<`, `\<`,
	`>`, `\>`,
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
		fields = append(fields, f.DOT())
	}

	return strings.Join(fields, " | ")
}

// Field represents a field.
type Field interface {
	DOT() string
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

// DOT generates a DOT representation of the simpleField object.
func (f *simpleField) DOT() string {
	var b bytes.Buffer

	if f.Name != "" {
		b.WriteRune('<')
		b.WriteString(f.Name)
		b.WriteRune('>')
	}

	if f.Label != "" {
		label := f.Label
		label = replacer.Replace(label)

		b.WriteString(label)
	}

	return b.String()
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

// DOT generates a DOT representation of the complexField object.
func (f *complexField) DOT() string {
	var b bytes.Buffer

	b.WriteString("{ ")
	b.WriteString(f.Record.Label())
	b.WriteString(" }")

	return b.String()
}
