
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

   gg

}

/////////////// Interfaces ////////////////

type gg interface {
    GetUser(ChannelCreate, ) (*ChannelCreateResponse ,error)
}



////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToClient, params RPC_UserParam, rpcHandler RFC_AllHandlersInteract) {

	splits := strings.Split(cmd.Command, ".")

	if len(splits) != 2 {
		return
	}

	switch splits[0] {

    case "gg": //each pb_service
            rpc,ok := rpcHandler.(gg)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : gg"))
                return
            }

            switch splits[1]  {
                case "GetUser": //each pb_service_method
                    load := &ChannelCreate{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetUser(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    }
}

