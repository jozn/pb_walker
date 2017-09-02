package ant

////////// Service /////////
type ServiceView struct {
	Name    string
	Methods []MethodView
	Comment string
	Hash    int32
	Options []OptionsView
}

type MethodView struct {
	MethodName  string
	InTypeName  string
	OutTypeName string
	Hash        int32
	Options     []OptionsView
}

////////// Messages /////////

type MessageView struct {
	Name    string
	Fields  []FieldView
	Comment string
	Options []OptionsView
}

type FieldView struct {
	FieldName  string
	TypeName   string
	Repeated   bool
	TagNumber  int
	GoType     string
	GoFlatType string
	JavaType   string
	Options    []OptionsView
}

////////// Enums /////////

type EnumView struct {
	Name    string
	Fields  []EnumFieldView
	Comment string
	Options []OptionsView
}

type EnumFieldView struct {
	FieldName string
	TagNumber int
	PosNumber int
	Options   []OptionsView
}

/////////// Tag /////////
type OptionsView struct {
	OptionName  string
	OptionValue string
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
