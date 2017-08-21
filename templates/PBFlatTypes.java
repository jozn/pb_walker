package ir.ms.pb

public class PBFlatTypes {
{{range .Messages}}
	public class {{.Name}} {
	  {{- range .Fields}}
	   public {{.JavaType}} {{.FieldName}};
	  {{- end }}
	}
{{end}}
	
}
/*
{{range .Services}}
RPC_INTERFACES.{{.Name}} {{.Name}}_Handeler = null;
{{- end}}
	
*/