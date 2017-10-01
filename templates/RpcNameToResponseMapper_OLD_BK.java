package ir.ms.pb;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import ir.ms.pb.PB_SyncParam_GetGeneralUpdates;

/**
 * Created by Hamid on 9/30/2017.
 */

public class RpcNameToResponseMapper {
	private static Map<String, Class< ? extends com.google.protobuf.GeneratedMessageLite>> mp = new ConcurrentHashMap<>();

	public static Map<String, Class< ? extends com.google.protobuf.GeneratedMessageLite>> getMap(){
		if(mp.size() < 1){
			fill();
		}
		return mp;
	}

	private static synchronized void fill(){
	 	{{range .Services}}
	 	// Service {{.Name}}
			{{$ser := . }}
			{{- range .Methods}}
		mp.put("{{$ser.Name}}.{{.MethodName}}", {{.OutTypeName}}.class);
			{{- end}}

	   	{{end}}

		//mp.put("Rpc_Msg.GetFull", PB_SyncParam_GetGeneralUpdates.class.getName());
	}
}
