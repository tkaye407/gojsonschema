package gojsonschema

import "github.com/xeipuuv/gojsonreference"

type SubSchemaAccessor struct {
	parent *subSchema
	schema *subSchema
}

func NewSubSchemaAccessor(in *subSchema) *SubSchemaAccessor {
	if in != nil && in.refSchema != nil {
		return &SubSchemaAccessor{schema: in.refSchema, parent: in}
	}

	return &SubSchemaAccessor{schema: in}
}

func (s *SubSchemaAccessor) Title() *string {
	return s.schema.title
}

func (s *SubSchemaAccessor) UniqueItems() bool {
	return s.schema.uniqueItems
}

func (s *SubSchemaAccessor) AdditionalProperties() (bool, *SubSchemaAccessor) {
	if s.schema.additionalProperties == nil {
		return false, nil
	}

	switch typedValue := s.schema.additionalProperties.(type) {
	case bool:
		return typedValue, nil
	case *subSchema:
		return false, NewSubSchemaAccessor(typedValue)
	}

	return false, nil
}

func (s *SubSchemaAccessor) Description() *string {
	return s.schema.description
}

func (s *SubSchemaAccessor) Name() string {
	if s.parent != nil {
		return s.parent.property
	}

	return s.schema.property
}

func (s *SubSchemaAccessor) Ref() *gojsonreference.JsonReference {
	return s.schema.ref
}

func (s *SubSchemaAccessor) RefSchema() *SubSchemaAccessor {
	return NewSubSchemaAccessor(s.schema.refSchema)
}

func (s *SubSchemaAccessor) Schema() *SubSchemaAccessor {
	if s.schema == nil {
		return nil
	}
	return NewSubSchemaAccessor(s.schema)
}

func (s *SubSchemaAccessor) Parent() *SubSchemaAccessor {
	return NewSubSchemaAccessor(s.schema.parent)
}

func (s *SubSchemaAccessor) Items() []*SubSchemaAccessor {
	children := make([]*SubSchemaAccessor, 0, len(s.schema.itemsChildren))
	for _, c := range s.schema.itemsChildren {
		children = append(children, NewSubSchemaAccessor(c))
	}
	return children
}

func (s *SubSchemaAccessor) Properties() []*SubSchemaAccessor {
	children := make([]*SubSchemaAccessor, 0, len(s.schema.propertiesChildren))
	for _, c := range s.schema.propertiesChildren {
		children = append(children, NewSubSchemaAccessor(c))
	}
	return children
}

func (s *SubSchemaAccessor) Enum() []string {
	return s.schema.enum
}

func (s *SubSchemaAccessor) Required() []string {
	return s.schema.required
}

func (s *SubSchemaAccessor) Type() *string {
	if s.schema.types != nil && s.schema.types.IsTyped() {
		schemaType := s.schema.types.String()
		return &schemaType
	}

	if s.schema.bsonTypes != nil && s.schema.bsonTypes.IsTyped() {
		schemaType := s.schema.bsonTypes.String()
		return &schemaType
	}

	return nil
}

func (s *SubSchemaAccessor) Default() interface{} {
	return s.schema.defaultValue
}
