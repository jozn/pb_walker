package ir.ms.pb;

import android.util.Log;

import java.util.HashMap;
import java.util.Map;

public class RPC_HANDLERS {
{{range .Services}}
public interface {{.Name}} {
  {{- range .Methods}}
    void {{.MethodName}}( {{.OutTypeName}} pb, boolean handled);
  {{- end }}
}
{{- end}}

{{range .Services}}
  public static class {{.Name}}_Empty implements {{.Name}}{
  {{$SName := .Name }}
  {{- range .Methods}}
  	@Override
    public void {{.MethodName}}( {{.OutTypeName}} pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC '{{$SName}}.{{.MethodName}}' ");
    }

  {{- end }}
  }
{{- end}}

	/////////////////////////////////// Maper of Handling methods /////////////////////////////////
	public static interface HandleRowRpcResponse {
		void handle(Object pb,boolean handled);
	}


	{{range .Services}}
	public static RPC_HANDLERS.{{.Name}} {{.Name}}_Default_Handler = new RPC_HANDLERS.{{.Name}}_Empty();
	{{- end}}


	public static Map<String,HandleRowRpcResponse > maper = new HashMap<>();

	public static void initMap(){
		{{range .Services}}
			{{$SName := .Name }}
              {{- range .Methods}}
              	maper.put("{{$SName}}.{{.MethodName}}", (pb, handled)->{
                	if(pb instanceof {{.OutTypeName}}){
                		{{$SName}}_Default_Handler.{{.MethodName}}(({{.OutTypeName}}) pb, handled);
                	}else{
                		Log.d("RPC", " can not convert response object to {{.OutTypeName}} in rpc: .{{.MethodName}} ");
                	}
                });
              {{end}}
        {{- end}}
	}
	
}
/*
{{range .Services}}
RPC_HANDLERS.{{.Name}} {{.Name}}_Default_Handler = new RPC_HANDLERS.{{.Name}} {{.Name}}_Empty();
{{- end}}
	
*/