package ir.ms.pb;

import com.mardomsara.social.pipe.*;
import android.util.Log;

public class RPC {
{{range .Services}}
public static class {{.Name}} {
	{{$SName := .Name }}

	{{- range .Methods}}
        public static interface {{.MethodName}}_ResultHandler{
    		void onResult({{.OutTypeName}} res);
        }
    {{- end -}}

  {{- range .Methods}}
    public static void {{.MethodName}}( {{.InTypeName}} param ,{{.MethodName}}_ResultHandler resultHandler, ErrorCallback errorCallback){
		SuccessCallback callback = null;
		if(resultHandler != null){
			callback = new SuccessCallback() {
				@Override
				public void handle(byte[] data) {
					Log.i("RPC ws", "handling rpc respnse for: {{.MethodName}} with respose class {{.OutTypeName}}");
					try {
						{{.OutTypeName}} d ={{.OutTypeName}}.parseFrom(data);
						resultHandler.onResult(d);
					}catch (com.google.protobuf.InvalidProtocolBufferException e){
						Log.d("RPC ws", "parsing protocol buffer is faild: {{.OutTypeName}}");
					}
				}
			};
		}

		Pipe.send("{{$SName}}.{{.MethodName}}", param, callback ,errorCallback);
    }
  {{- end -}}
}
{{- end}}
	
}
/*
{{range .Services}}
RPC_INTERFACES.{{.Name}} {{.Name}}_Handeler = null;
{{- end}}
	
*/