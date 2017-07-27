package ant

import (
	"strings"
	"text/template"
)

var fns = template.FuncMap{
	"tIsPBPrimateTypes":    tIsPBPrimateTypes,
	"tPBTypeToGoFlatType":  tPBTypeToGoFlatType,
	"tFlatTypeToGoPBType":  tFlatTypeToGoPBType,
	"tFlatTypeToGoPBType2": tFlatTypeToGoPBType2,
}

func tIsPBPrimateTypes(pbType string) bool {
	r := false
	switch pbType {
	case "int64", "sint64", "int32",
		"sint32", "uint32", "uint64", "fixed32",
		"fixed64", "sfixed32", "sfixed64":
		r = true
	case "double":
		r = true
	case "float":
		r = true
	case "bool":
		r = true
	case "string":
		r = true
	case "bytes":
		r = true
	}
	return r
}

func tPBTypeToGoFlatType(field FieldView, fieldPerifx string) string {
	r := ""
	flatSr := pbTypesToGoFlatTypes(field.TypeName)
	if field.TypeName == flatSr {
		r = fieldPerifx + "." + field.FieldName
	} else {
		if field.Repeated {
			m := "helper.Slice" + strings.Title(field.TypeName) + "To" + strings.Title(flatSr)
			r = m + "(" + fieldPerifx + "." + field.FieldName + ")"
		} else {
			r = flatSr + "(" + fieldPerifx + "." + field.FieldName + ")"
		}
	}

	return r
}

func tFlatTypeToGoPBType2(field FieldView, fieldPerifx string) string {
	r := ""

	flatSr := pbTypesToGoFlatTypes(field.TypeName)
	goSr := pbTypesToGoType(field.TypeName)
	if goSr == flatSr {
		r = fieldPerifx + "." + field.FieldName
	} else {
		if field.Repeated {
			m := "helper.Slice" + strings.Title(flatSr) + "To" + strings.Title(goSr)
			r = m + "(" + fieldPerifx + "." + field.FieldName + ")"
		} else {
			r = goSr + "(" + fieldPerifx + "." + field.FieldName + ")"
		}
	}

	return r
}

//////////////////////// Deprecated //////////////////

func tFlatTypeToGoPBType(field, pbType, fieldPerifx string) string {
	r := ""
	flatSr := pbTypesToGoFlatTypes(pbType)
	goSr := pbTypesToGoType(pbType)
	if goSr == flatSr {
		r = fieldPerifx + "." + field
	} else {
		r = goSr + "(" + fieldPerifx + "." + field + ")"
	}

	return r
}
