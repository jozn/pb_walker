package ant

import (
	"text/template"
)

var fns = template.FuncMap{
	"IsPBPrimateTypes":    IsPBPrimateTypes,
	"tPBTypeToGoFlatType": tPBTypeToGoFlatType,
	"tFlatTypeToGoPBType": tFlatTypeToGoPBType,
}

func IsPBPrimateTypes(pbType string) bool {
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

func tPBTypeToGoFlatType(field, pbType, fieldPerifx string) string {
	r := ""
	flatSr := pbTypesToGoFlatTypes(pbType)
	if pbType == flatSr {
		r = fieldPerifx + "." + field
	} else {
		r = flatSr + "(" + fieldPerifx + "." + field + ")"
	}

	return r
}

func tFlatTypeToGoPBType(field, pbType, fieldPerifx string) string {
	r := ""
	flatSr := pbTypesToGoFlatTypes(pbType)
	if pbType == flatSr {
		r = fieldPerifx + "." + field
	} else {
		r = pbType + "(" + fieldPerifx + "." + field + ")"
	}

	return r
}
