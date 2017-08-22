package ir.ms.pb;

import android.util.Log;

public class RPC_ResponseBase {
{{range .Services}}
	public static class {{.Name}}_Base implements RPC_INTERFACES.{{.Name}}  {
	  {{- $SName := .Name -}}
	  {{ range .Methods}}
		public void {{.MethodName}}( {{.OutTypeName}} pbOut){
			Log.d("RPC:Response", "{{$SName}}_Base.{{$SName}}");
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