package ant

////////// Service /////////
type ServiceView struct {
	Name    string
	Methods []MethodView
	Comment string
	Hash    int32
}

type MethodView struct {
	MethodName  string
	InTypeName  string
	OutTypeName string
	Hash        int32
}

////////// Messages /////////

type MessageView struct {
	Name    string
	Fields  []FieldView
	Comment string
}

type FieldView struct {
	FieldName  string
	TypeName   string
	Repeated   bool
	TagNumber  int
	GoType     string
	GoFlatType string
	javaType   string
}

////////// Enums /////////

type EnumView struct {
	Name    string
	Fields  []EnumFieldView
	Comment string
}

type EnumFieldView struct {
	FieldName string
	TagNumber int
	PosNumber int
}

/////////////////////////////////////////
///////////// Extractor /////////////////

type GenOut struct {
	Services []ServiceView
	Messages []MessageView
	Enums    []EnumView

	OutGoEnumsStr string
	OutGoRPCsStr  string
	OutJavaStr    string
}
