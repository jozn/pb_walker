package ant

////////// Service /////////
type ServiceView struct {
    Name    string
    Methods []MethodView
    Comment string
}

type MethodView struct {
    MethodName  string
    InTypeName  string
    OutTypeName string
}

////////// Messages /////////

type MessageView struct {
    Name    string
    Fields  []FieldView
    Comment string
}

type FieldView struct {
    FieldName string
    TypeName  string
    Repeated  bool
    TagNumber int
    GoType    string
    javaType  string
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
