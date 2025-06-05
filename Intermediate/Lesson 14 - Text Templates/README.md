Text Templates in GO

- Useful to generate structured texts like HTML, JSON, SQL Qurey

Actions
a. Variable Insertion: {{.fieldName}}
b. Pipelines: {{functionName .FieldName}}
c. Control Statements: {{if .Condition}} ... {{else}} ... {{end}}
d. Iteration: {{range .Slice}} ... {{end}}

Advanced Functions
a. Nested Templates: {{templete "name" .}}
b. Functions
c. Custom Delimiters
d. Error Handaling: template.Must()

Use Cases
a. HTML Creation
b. Email Templates
c. Code Generation
d. Document Generation

Best Practices
a. Seperation of Concer
b. Efficiency
c. Security

// Part of two packages text template package and html template package
