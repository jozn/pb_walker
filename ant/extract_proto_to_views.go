package ant

import (
	"github.com/dsymonds/gotoc/ast"
	"strings"
)

//deps
func ExtractAllServicesViews(pbFilesSet *ast.FileSet) []ServiceView {
	serviceViews := make([]ServiceView, 0)

	for _, pbFile := range pbFilesSet.Files {
		for _, ser := range pbFile.Services {
			serView := ServiceView{
				Name:    ser.Name,
				Comment: findComment(ser.Position, pbFile),
				Hash:    Hash32(ser.Name),
			}

			for _, m := range ser.Methods {
				mv := MethodView{
					MethodName:     m.Name,
					InTypeName:     m.InTypeName,
					GoInTypeName:   strings.Replace(m.InTypeName, ".", "_", -1),
					OutTypeName:    m.OutTypeName,
					GoOutTypeName:  strings.Replace(m.OutTypeName, ".", "_", -1),
					Hash:           Hash32(m.Name),
					FullMethodName: serView.Name + "." + m.Name,
				}
				serView.Methods = append(serView.Methods, mv)
			}
			serviceViews = append(serviceViews, serView)
		}
	}

	return serviceViews
}

func ExtractAllMessagesViews(pbFilesSet *ast.FileSet) []MessageView {
	messageViews := make([]MessageView, 0)

	for _, pbFile := range pbFilesSet.Files {
		for _, msg := range pbFile.Messages {
			msgView := MessageView{
				Name:    msg.Name,
				Comment: findComment(msg.Position, pbFile),
			}

			for _, f := range msg.Fields {
				mv := FieldView{
					FieldName:  f.Name,
					TypeName:   f.TypeName,
					Repeated:   f.Repeated,
					TagNumber:  f.Tag,
					GoType:     pbTypesToGoType(f.TypeName),
					GoFlatType: pbTypesToGoFlatTypes(f.TypeName),
					JavaType:   pbTypesToJavaType(f.TypeName),
				}
				msgView.Fields = append(msgView.Fields, mv)
			}
			messageViews = append(messageViews, msgView)
		}
	}

	return messageViews
}

func ExtractAllEnumsViews(pbFilesSet *ast.FileSet) []EnumView {
	enumViews := make([]EnumView, 0)

	for _, pbFile := range pbFilesSet.Files {
		for _, enum := range pbFile.Enums {
			enumView := EnumView{
				Name:    enum.Name,
				Comment: findComment(enum.Position, pbFile),
			}

			for _, value := range enum.Values {
				fieldView := EnumFieldView{
					FieldName: value.Name,
					TagNumber: int(value.Number),
					PosNumber: int(value.Number),
				}
				enumView.Fields = append(enumView.Fields, fieldView)
			}
			enumViews = append(enumViews, enumView)
		}
	}

	return enumViews
}

func findComment(pos ast.Position, pbFile *ast.File) string {
	for _, com := range pbFile.Comments {
		if com.End.Line == pos.Line-1 && len(com.Text) > 0 {
			return com.Text[len(com.Text)-1]
		}
	}
	return ""
}
