// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package codegen

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/thriftrw/compile"
)

// TypeConverter can generate a function body that converts two thriftrw
// FieldGroups from one to another. It's assumed that the converted code
// operates on two variables, "in" and "out" and that both are a go struct.
type TypeConverter struct {
	Lines  []string
	Helper PackageNameResolver
}

// PackageNameResolver interface allows for resolving what the
// package name for a thrift file is. This depends on where the
// thrift-based structs are generated.
type PackageNameResolver interface {
	TypePackageName(thriftFile string) (string, error)
}

// append helper will add a line to TypeConverter
func (c *TypeConverter) append(parts ...string) {
	line := strings.Join(parts, "")
	c.Lines = append(c.Lines, line)
}

// appendf helper will add a formatted line to TypeConverter
func (c *TypeConverter) appendf(format string, parts ...interface{}) {
	line := fmt.Sprintf(format, parts...)
	c.Lines = append(c.Lines, line)
}

func (c *TypeConverter) getGoTypeName(
	valueType compile.TypeSpec,
) (string, error) {
	switch s := valueType.(type) {
	case *compile.BoolSpec:
		return "bool", nil
	case *compile.I8Spec:
		return "int8", nil
	case *compile.I16Spec:
		return "int16", nil
	case *compile.I32Spec:
		return "int32", nil
	case *compile.I64Spec:
		return "int64", nil
	case *compile.DoubleSpec:
		return "float64", nil
	case *compile.StringSpec:
		return "string", nil
	case *compile.BinarySpec:
		return "[]byte", nil
	case *compile.MapSpec:
		/* coverage ignore next line */
		panic("Not Implemented")
	case *compile.SetSpec:
		/* coverage ignore next line */
		panic("Not Implemented")
	case *compile.ListSpec:
		/* coverage ignore next line */
		panic("Not Implemented")
	case *compile.EnumSpec, *compile.StructSpec, *compile.TypedefSpec:
		return c.getIdentifierName(s)
	default:
		/* coverage ignore next line */
		panic(fmt.Sprintf("Unknown type (%T) %v", valueType, valueType))
	}
}

func (c *TypeConverter) getIdentifierName(
	fieldType compile.TypeSpec,
) (string, error) {
	pkgName, err := c.Helper.TypePackageName(fieldType.ThriftFile())
	if err != nil {
		return "", errors.Wrapf(
			err,
			"could not lookup fieldType when building converter for %s :",
			fieldType.ThriftName(),
		)
	}
	typeName := pkgName + "." + fieldType.ThriftName()
	return typeName, nil
}

func (c *TypeConverter) genConverterForStruct(
	toFieldName string,
	toFieldType *compile.StructSpec,
	fromFieldType compile.TypeSpec,
	fromIdentifier string,
	keyPrefix string,
	indent string,
) error {
	toIdentifier := "out." + keyPrefix

	typeName, err := c.getIdentifierName(toFieldType)
	if err != nil {
		return err
	}
	subToFields := toFieldType.Fields

	fromFieldStruct, ok := fromFieldType.(*compile.StructSpec)
	if !ok {
		return errors.Errorf(
			"could not convert struct fields, "+
				"incompatible type for %s :",
			toFieldName,
		)
	}

	c.append(indent, "if ", fromIdentifier, " != nil {")

	c.append(indent, "	", toIdentifier, " = &", typeName, "{}")

	subFromFields := fromFieldStruct.Fields
	err = c.genStructConverter(
		keyPrefix+".",
		indent+"	",
		subFromFields,
		subToFields,
	)
	if err != nil {
		return err
	}

	c.append(indent, "} else {")

	c.append(indent, "	", toIdentifier, " = nil")

	c.append(indent, "}")

	return nil
}

func (c *TypeConverter) genConverterForPrimitive(
	toField *compile.FieldSpec,
	toIdentifier string,
	fromIdentifier string,
	overiddenIdentifier string,
) error {
	typeName, err := c.getGoTypeName(toField.Type)
	if err != nil {
		return err
	}
	if overiddenIdentifier != "" {

	}
	if toField.Required {
		c.append(toIdentifier, " = ", typeName, "(", fromIdentifier, ")")
	} else {
		c.append(toIdentifier, " = (*", typeName, ")(", fromIdentifier, ")")
	}
	return nil
}

func (c *TypeConverter) genConverterForList(
	toFieldType *compile.ListSpec,
	toField *compile.FieldSpec,
	fromField *compile.FieldSpec,
	toIdentifier string,
	fromIdentifier string,
	keyPrefix string,
	indent string,
) error {
	typeName, err := c.getGoTypeName(toFieldType.ValueSpec)
	if err != nil {
		return err
	}

	valueStruct, isStruct := toFieldType.ValueSpec.(*compile.StructSpec)
	if isStruct {
		c.appendf(
			"%s = make([]*%s, len(%s))",
			toIdentifier, typeName, fromIdentifier,
		)
	} else {
		c.appendf(
			"%s = make([]%s, len(%s))",
			toIdentifier, typeName, fromIdentifier,
		)
	}

	c.append("for index, value := range ", fromIdentifier, " {")

	if isStruct {
		fromFieldType, ok := fromField.Type.(*compile.ListSpec)
		if !ok {
			return errors.Errorf(
				"Could not convert field (%s): type is not list",
				fromField.Name,
			)
		}

		err = c.genConverterForStruct(
			toField.Name,
			valueStruct,
			fromFieldType.ValueSpec,
			"value",
			keyPrefix+strings.Title(toField.Name)+"[index]",
			"	"+indent,
		)
		if err != nil {
			return err
		}
	} else {
		c.append("	", toIdentifier, "[index] = ", typeName, "(value)")
	}

	c.append("}")
	return nil
}

func (c *TypeConverter) genConverterForMap(
	toFieldType *compile.MapSpec,
	toField *compile.FieldSpec,
	fromField *compile.FieldSpec,
	toIdentifier string,
	fromIdentifier string,
	keyPrefix string,
	indent string,
) error {
	typeName, err := c.getGoTypeName(toFieldType.ValueSpec)
	if err != nil {
		return err
	}

	_, isStringKey := toFieldType.KeySpec.(*compile.StringSpec)
	if !isStringKey {
		return errors.Errorf(
			"could not convert key (%s), map is not string-keyed.",
			toField.Name,
		)
	}

	valueStruct, isStruct := toFieldType.ValueSpec.(*compile.StructSpec)
	if isStruct {
		c.appendf(
			"%s = make(map[string]*%s, len(%s))",
			toIdentifier, typeName, fromIdentifier,
		)
	} else {
		c.appendf(
			"%s = make(map[string]%s, len(%s))",
			toIdentifier, typeName, fromIdentifier,
		)
	}

	c.append("for key, value := range ", fromIdentifier, " {")

	if isStruct {
		fromFieldType, ok := fromField.Type.(*compile.MapSpec)
		if !ok {
			return errors.Errorf(
				"Could not convert field (%s): type is not map",
				fromField.Name,
			)
		}

		err = c.genConverterForStruct(
			toField.Name,
			valueStruct,
			fromFieldType.ValueSpec,
			"value",
			keyPrefix+strings.Title(toField.Name)+"[key]",
			"	"+indent,
		)
		if err != nil {
			return err
		}
	} else {
		c.append("	", toIdentifier, "[key] = ", typeName, "(value)")
	}

	c.append("}")
	return nil
}

func (c *TypeConverter) genStructConverter(
	keyPrefix string,
	indent string,
	fromFields []*compile.FieldSpec,
	toFields []*compile.FieldSpec,
	fieldMap map[*compile.FieldSpec]FieldMapperEntry,
) error {
	for i := 0; i < len(toFields); i++ {
		toField := toFields[i]

		// Check for same named field
		var fromField *compile.FieldSpec
		for j := 0; j < len(fromFields); j++ {
			if fromFields[j].Name == toField.Name {
				fromField = fromFields[j]
				break
			}
		}

		// Check for mapped field
		var overiddenField *compile.FieldSpec
		for k, v := range fieldMap {
			if v.dest.Name == toField.Name {
				if fromField == nil {
					fromField = v.dest
				} else {
					if v.override {
						// check for required/optional setting
						if !v.dest.Required {
							overiddenField = fromField
						}
						fromField = v.dest
					} else {
						// check for required/optional setting
						if !fromField.Required {
							overiddenField = v.dest
						}
					}
				}
			}
		}

		if fromField == nil {
			return errors.Errorf(
				"cannot map by name for the field %s",
				toField.Name,
			)
		}

		toIdentifier := indent + "out." + keyPrefix + strings.Title(toField.Name)
		fromIdentifier := "in." + keyPrefix + strings.Title(fromField.Name)
		overiddenIdentifier := ""
		if overiddenField != nil {
			overiddenIdentifier = "in." + keyPrefix + strings.Title(overiddenField.Name)
		}

		// Override thrift type names to avoid naming collisions between endpoint
		// and client types.
		switch toFieldType := toField.Type.(type) {
		case
			*compile.BoolSpec,
			*compile.I8Spec,
			*compile.I16Spec,
			*compile.I32Spec,
			*compile.EnumSpec,
			*compile.I64Spec,
			*compile.DoubleSpec,
			*compile.StringSpec:

			err := c.genConverterForPrimitive(
				toField, toIdentifier, fromIdentifier, overiddenIdentifier,
			)
			if err != nil {
				return err
			}
		case *compile.BinarySpec:
			c.append(toIdentifier, " = []byte(", fromIdentifier, ")")
		case *compile.TypedefSpec:
			typeName, err := c.getIdentifierName(toField.Type)
			if err != nil {
				return err
			}

			// TODO: typedef for struct is invalid here ...
			if toField.Required {
				c.append(
					toIdentifier, " = ", typeName, "(", fromIdentifier, ")",
				)
			} else {
				c.append(
					toIdentifier, " = (*", typeName, ")(", fromIdentifier, ")",
				)
			}

		case *compile.StructSpec:
			err := c.genConverterForStruct(
				toField.Name,
				toFieldType,
				fromField.Type,
				fromIdentifier,
				keyPrefix+strings.Title(toField.Name),
				indent,
			)
			if err != nil {
				return err
			}
		case *compile.ListSpec:
			err := c.genConverterForList(
				toFieldType,
				toField,
				fromField,
				toIdentifier,
				fromIdentifier,
				keyPrefix,
				indent,
			)
			if err != nil {
				return err
			}
		case *compile.MapSpec:
			err := c.genConverterForMap(
				toFieldType,
				toField,
				fromField,
				toIdentifier,
				fromIdentifier,
				keyPrefix,
				indent,
			)
			if err != nil {
				return err
			}
		default:
			// fmt.Printf("Unknown type %s for field %s \n",
			// 	toField.Type.TypeCode().String(), toField.Name,
			// )

			// pkgName, err := h.TypePackageName(toField.Type.ThriftFile())
			// if err != nil {
			// 	return nil, err
			// }
			// typeName := pkgName + "." + toField.Type.ThriftName()
			// line := toIdentifier + "(*" + typeName + ")" + postfix
			// c.Lines = append(c.Lines, line)
		}
	}

	return nil
}

// GenStructConverter will add lines to the TypeConverter for mapping
// from one go struct to another based on two thriftrw.FieldGroups
func (c *TypeConverter) GenStructConverter(
	fromFields []*compile.FieldSpec,
	toFields []*compile.FieldSpec,
	fieldMap map[*compile.FieldSpec]FieldMapperEntry,
) error {
	err := c.genStructConverter("", "", fromFields, toFields, fieldMap)
	if err != nil {
		return err
	}

	return nil
}

// FieldMapperEntry defines a destination field and optional arguments
// converting and overriding fields.
type FieldMapperEntry struct {
	dest          *compile.FieldSpec
	override      bool
	typeConverter string // TODO: implement
	transform     string // TODO: implement
}

func assignWithOverride(
	toIdentifier string,
	typeName string,
	fromIdentifier string,
	overridenIdentifier string,
) []string {

	line1 := toIdentifier + " = " + typeName + "(" + overridenIdentifier + ")"
	line2 := "if " + fromIdentifier + " != nil {"
	line3 := toIdentifier + " = " + typeName + "(" + fromIdentifier + ")"
	line4 := "}"

	return []string{line1, line2, line3, line4}
}
