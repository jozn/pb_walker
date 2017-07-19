
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

   RPC_MessageReq
   RPC_MessageReqOffline
   Greeter
   RpcMsgs
   RPC_UserOffline
   RPC_User
   RPC_Auth
}

/////////////// Interfaces ////////////////

type RPC_MessageReq interface {
    GetLastChnagesForRoom(PB_ReqLastChangesForTheRoom, ) (*PB_ResponseLastChangesForTheRoom ,error)
}

type RPC_MessageReqOffline interface {
    SetLastSeen(PB_RequestSetLastSeenMessages, ) (*PB_ResponseSetLastSeenMessages ,error)
}

type Greeter interface {
    SayHello(PB_Message, ) (*PB_UserWithMe ,error)
}

type RpcMsgs interface {
    UploadNewMsg(PB_Message, ) (*PB_ResRpcAddMsg ,error)
}

type RPC_UserOffline interface {
    BlockUser(PB_UserParam_BlockUser, ) (*PB_UserOfflineResponse_BlockUser ,error)
    UnBlockUser(PB_UserParam_UnBlockUser, ) (*PB_UserOfflineResponse_UnBlockUser ,error)
    UpdateAbout(PB_UserParam_UpdateAbout, ) (*PB_UserOfflineResponse_UpdateAbout ,error)
    UpdateUserName(PB_UserParam_UpdateUserName, ) (*PB_UserOfflineResponse_UpdateUserName ,error)
    ChangePrivacy(PB_UserParam_ChangePrivacy, ) (*PB_UserResponseOffline_ChangePrivacy ,error)
    ChangeAvatar(PB_UserParam_ChangeAvatar, ) (*PB_UserOfflineResponse_ChangeAvatar ,error)
}

type RPC_User interface {
    CheckUserName(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    GetBlockedList(PB_UserParam_BlockedList, ) (*PB_UserResponse_BlockedList ,error)
}

type RPC_Auth interface {
    CheckPhone(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    SendCode(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    SendCodeToSms(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    SendCodeToTelgram(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    SingUp(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    SingIn(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
    LogOut(PB_UserParam_CheckUserName, ) (*PB_UserResponse_CheckUserName ,error)
}



////////////// map of rpc methods to all
func HandleRpcs(cmd PB_CommandToClient, params RPC_UserParam, rpcHandler RFC_AllHandlersInteract) {

	splits := strings.Split(cmd.Command, ".")

	if len(splits) != 2 {
		return
	}

	switch splits[0] {

    case "RPC_MessageReq": //each pb_service
            rpc,ok := rpcHandler.(RPC_MessageReq)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : RPC_MessageReq"))
                return
            }

            switch splits[1]  {
                case "GetLastChnagesForRoom": //each pb_service_method
                    load := &PB_ReqLastChangesForTheRoom{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetLastChnagesForRoom(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    case "RPC_MessageReqOffline": //each pb_service
            rpc,ok := rpcHandler.(RPC_MessageReqOffline)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : RPC_MessageReqOffline"))
                return
            }

            switch splits[1]  {
                case "SetLastSeen": //each pb_service_method
                    load := &PB_RequestSetLastSeenMessages{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SetLastSeen(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    case "Greeter": //each pb_service
            rpc,ok := rpcHandler.(Greeter)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : Greeter"))
                return
            }

            switch splits[1]  {
                case "SayHello": //each pb_service_method
                    load := &PB_Message{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SayHello(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    case "RpcMsgs": //each pb_service
            rpc,ok := rpcHandler.(RpcMsgs)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : RpcMsgs"))
                return
            }

            switch splits[1]  {
                case "UploadNewMsg": //each pb_service_method
                    load := &PB_Message{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UploadNewMsg(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    case "RPC_UserOffline": //each pb_service
            rpc,ok := rpcHandler.(RPC_UserOffline)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : RPC_UserOffline"))
                return
            }

            switch splits[1]  {
                case "BlockUser": //each pb_service_method
                    load := &PB_UserParam_BlockUser{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.BlockUser(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "UnBlockUser": //each pb_service_method
                    load := &PB_UserParam_UnBlockUser{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UnBlockUser(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "UpdateAbout": //each pb_service_method
                    load := &PB_UserParam_UpdateAbout{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UpdateAbout(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "UpdateUserName": //each pb_service_method
                    load := &PB_UserParam_UpdateUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.UpdateUserName(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "ChangePrivacy": //each pb_service_method
                    load := &PB_UserParam_ChangePrivacy{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.ChangePrivacy(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "ChangeAvatar": //each pb_service_method
                    load := &PB_UserParam_ChangeAvatar{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.ChangeAvatar(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    case "RPC_User": //each pb_service
            rpc,ok := rpcHandler.(RPC_User)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : RPC_User"))
                return
            }

            switch splits[1]  {
                case "CheckUserName": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.CheckUserName(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "GetBlockedList": //each pb_service_method
                    load := &PB_UserParam_BlockedList{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.GetBlockedList(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
    case "RPC_Auth": //each pb_service
            rpc,ok := rpcHandler.(RPC_Auth)
            if !ok {
                RPC_ResponseHandler.HandelError(errors.New("rpcHandler could not be cast to : RPC_Auth"))
                return
            }

            switch splits[1]  {
                case "CheckPhone": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.CheckPhone(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "SendCode": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SendCode(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "SendCodeToSms": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SendCodeToSms(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "SendCodeToTelgram": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SendCodeToTelgram(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "SingUp": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SingUp(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "SingIn": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.SingIn(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }
                case "LogOut": //each pb_service_method
                    load := &PB_UserParam_CheckUserName{}
                    err := proto.Unmarshal(cmd.Data, load)
                    if err == nil {
                        res, err := rpc.LogOut(load)
                        RPC_ResponseHandler.HandleOfflineResult(res,err)
                    }else{
                     RPC_ResponseHandler.HandelError(err)
                    }
                }}
}

