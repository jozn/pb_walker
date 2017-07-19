package ir.ms.pb

public class RPC_INTERFACES {
{{range .Services}}
public interface {{.Name}} {
  {{- range .Methods}}
    {{.MethodName}}( {{.OutTypeName}} );
  {{- end -}}
}
{{- end}}
	
}
/*
{{range .Services}}
RPC_INTERFACES.{{.Name}} {{.Name}}_Handeler = null;
{{- end}}
	
*/