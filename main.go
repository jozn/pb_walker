package main

import (
	"encoding/json"
	"fmt"
	"github.com/dsymonds/gotoc/parser"

	"bytes"
	"github.com/dsymonds/gotoc/ast"
	"io/ioutil"
	"log"
	"ms/sun/helper"
	"os"
	"text/template"
)

////////// Service /////////
type ServiceView struct {
	Name    string
	Methods []MethodView
	Comment string
}

type MethodView struct {
	MethodName  string
	InTypeName  string
	OutTypeName string
}

////////// Messages /////////

type MessageView struct {
	Name    string
	Fields  []FieldView
	Comment string
}

type FieldView struct {
	FieldName string
	TypeName  string
	Repeated  bool
	TagNumber int
	GoType    string
	javaType  string
}

////////// Enums /////////

type EnumView struct {
	Name    string
	Fields  []EnumFieldView
	Comment string
}

type EnumFieldView struct {
	FieldName string
	TagNumber int
	PosNumber int
}

//////////////////////////

const DIR_OUTPUT = `C:\Go\_gopath\src\ms\sun\models\x\pb__gen_ant.go` //"./play/gen_sample_out.go"
const TEMPLATES_DIR  = "./templates/"

func main() {
	const DIR_PROTOS = `C:\Go\_gopath\src\ms\sun\models\protos`
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	protos := make([]string, len(files))
	for i, f := range files {
		protos[i] = f.Name()
	}

	//ast, err := parser.ParseFiles([]string{"1.proto"}, []string{"./play/"})
	ast, err := parser.ParseFiles(protos, []string{DIR_PROTOS})

	fmt.Println(ast, err)
	fmt.Println(err)
	for _, f := range ast.Files {
		//helper.PertyPrint(f)

		_ = f
	}

	helper.PertyPrint(ExtractAllServicesViews(ast))
	helper.PertyPrint(ExtractAllMessagesViews(ast))
	helper.PertyPrint(ExtractAllEnumsViews(ast))

	templSerivces(ExtractAllServicesViews(ast))

}

func ExtractAllServicesViews(pbFilesSet *ast.FileSet) []ServiceView {
	serviceViews := make([]ServiceView, 0)

	for _, pbFile := range pbFilesSet.Files {
		for _, ser := range pbFile.Services {
			serView := ServiceView{
				Name:    ser.Name,
				Comment: findComment(ser.Position, pbFile),
			}

			for _, m := range ser.Methods {
				mv := MethodView{
					MethodName:  m.Name,
					InTypeName:  m.InTypeName,
					OutTypeName: m.OutTypeName,
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
					FieldName: f.Name,
					TypeName:  f.TypeName,
					Repeated:  f.Repeated,
					TagNumber: f.Tag,
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

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}

func templSerivces(services []ServiceView) {
	tpl := template.New("go_interface")
    tplGoInterface,err := ioutil.ReadFile(TEMPLATES_DIR+ "rpc.go")
    noErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	noErr(err)

	s := struct {
		Services []ServiceView
	}{
		Services: services,
	}

	bts := bytes.NewBufferString("")
	err = tpl.Execute(bts, s)
	noErr(err)

	fmt.Println(bts.String())
	ioutil.WriteFile(DIR_OUTPUT, []byte(bts.String()), os.ModeType)
}

func noErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
