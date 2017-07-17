package main

import (
	"encoding/json"
	"fmt"
	"github.com/dsymonds/gotoc/parser"

	"bytes"
	"github.com/dsymonds/gotoc/ast"
	"log"
	"ms/sun/helper"
	"text/template"
    "io/ioutil"
    "os"
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

func main() {
	ast, err := parser.ParseFiles([]string{"1.proto"}, []string{"./play/"})

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

func templSerivcesEach_DEP(services []ServiceView) {
	tpl := template.New("go_interface")
	tpl, err := tpl.Parse(tplGoInterface)
	noErr(err)

	bts := bytes.NewBufferString("")
	//tpl.Execute(bts,nil)

	for _, s := range services {
		err = tpl.Execute(bts, s)
		noErr(err)

	}
	fmt.Println(bts.String())

}

func templSerivces(services []ServiceView) {
	tpl := template.New("go_interface")
	tpl, err := tpl.Parse(tplGoInterface)
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
    ioutil.WriteFile("./play/gen_sample_out.go",[]byte(bts.String()), os.ModeType)
}

func noErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

const tplGoInterface = `
package x

import (
    "strings"
    "github.com/golang/protobuf/proto"
    "errors"
)


type RPC_UserParam interface {
	GetUserId() int
	IsUser() bool
}

type RPC_ResponseHandlerInterface interface {
	HandleOfflineResult(interface{}, error) int
	IsUserOnlineResult(interface{}, error) bool
	HandelError(error)
}

var RPC_ResponseHandler RPC_ResponseHandlerInterface

//note: rpc methods cant have equal name they must be different even in different rpc services
type RFC_AllHandlersInteract interface {
{{range .Services}}
   {{.Name}}
{{end}}
}

/////////////// Interfaces ////////////////
{{range .Services}}
type {{.Name}} interface {
{{- range .Methods}}
    {{.MethodName}}({{.InTypeName}}, ) (*{{.OutTypeName}} ,error)
{{- end}}
}
{{end}}


////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToClient, params RPC_UserParam, rpcHandler RFC_AllHandlersInteract) {

	splits := strings.Split(cmd.Command, ".")

	if len(splits) != 2 {
		return
	}

	switch splits[0] {
{{range .Services}}
    case "{{.Name}}": //each pb_service
            rpc,ok := rpcHandler.({{.Name}})
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : {{.Name}}"))
                return
            }

            switch splits[1]  {
            {{- range .Methods}}
                case "{{.MethodName}}": //each pb_service_method
                    load := &{{.InTypeName}}{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.{{.MethodName}}(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
            {{- end -}}
            }
    }
{{- end}}
}

`

type sss interface {
	SomeMethod(string) (int, error)
}
