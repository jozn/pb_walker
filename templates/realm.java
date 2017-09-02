package com.mardomsara.social.models_realm.pb_realm;

import io.realm.RealmObject;
import io.realm.annotations.PrimaryKey;
import ir.ms.pb.*;


{{$pk := (index .Fields 0).FieldName }}
public class {{.RealmClass}} extends RealmObject {
   {{- range $i ,$el := .Fields }}
		{{if (eq $pk  $el.FieldName ) }} @PrimaryKey 
		{{ end -}}
	public {{$el.JavaType}} {{$el.FieldName }}; 			//{{$i}}  PB {{$el.TagNumber}}
   {{- end }}
	

	public static {{.RealmClass}} fromPB({{.Name}} pb){
		{{.RealmClass}} t = new {{.RealmClass}}();
		{{ range .Fields}}
	    t.{{.FieldName}} = pb.get{{.FieldName}}();
	    {{- end }}

	    return t;
	}

}
	/*
	folding

	PBFlatTypes.{{.Name}} t = new PBFlatTypes.{{.Name}}();
	{{- range .Fields}}
    t.{{.FieldName}} = pb.get{{.FieldName}}();
    {{- end }}
	*/

	/*
	PBFlatTypes.{{.Name}} t = new PBFlatTypes.{{.Name}}();
	{{- range .Fields}}
	t.{{.FieldName}} = ;
	{{- end }}
	*/

	/*
	{{.Name}} t = new {{.Name}}();
	{{- range .Fields}}
	t.{{.FieldName}} = m.get{{.FieldName}}() ;
	{{- end }}
	
*/
