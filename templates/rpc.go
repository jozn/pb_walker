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
type RPC_AllHandlersInteract interface {
{{range .Services}}
   {{.Name}}
{{- end}}
}

/////////////// Interfaces ////////////////
{{range .Services}}
type {{.Name}} interface {
{{- range .Methods}}
    {{.MethodName}}({{.InTypeName}} ) ({{.OutTypeName}} ,error)
{{- end}}
}
{{end}}


////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToClient, params RPC_UserParam, rpcHandler RPC_AllHandlersInteract) {

    splits := strings.Split(cmd.Command, ".")

    if len(splits) != 2 {
        return
    }

    switch splits[0] {
{{range .Services}}
    case "{{.Name}}":
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
                        res, err := rpc.{{.MethodName}}(*load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
            {{- end}}
            }
{{- end}}
}
}