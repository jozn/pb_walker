package ant

import (
	"fmt"
	"github.com/emicklei/proto"
	"ms/sun/helper"
	"os"
)

func ExtractAllServicesViewsV2(protos []*proto.Proto) []ServiceView {
	serviceViews := make([]ServiceView, 0)

	for _, pto := range protos {
		for _, entry := range pto.Elements {
			if ser, ok := entry.(*proto.Service); ok {
				serView := ServiceView{
					Name:    ser.Name,
					Comment: findCommentV2(ser.Comment),
					Hash:    Hash32(ser.Name),
				}

				for _, m2 := range ser.Elements {
					if m, ok := m2.(*proto.RPC); ok {
						mv := MethodView{
							MethodName:  m.Name,
							InTypeName:  m.RequestType,
							OutTypeName: m.ReturnsType,
							Hash:        Hash32(m.Name),
						}
						serView.Methods = append(serView.Methods, mv)
					}
				}
				serviceViews = append(serviceViews, serView)
			}
		}
	}

	return serviceViews
}

func ExtractAllMessagesViewsV2(protos []*proto.Proto) []MessageView {
	messageViews := make([]MessageView, 0)

	for _, pto := range protos {
		for _, ele := range pto.Elements {
			if msg, ok := ele.(*proto.Message); ok {
				msgView := MessageView{
					Name:    msg.Name,
					Comment: findCommentV2(msg.Comment),
				}

				for _, f2 := range msg.Elements {
					if f, ok := f2.(*proto.NormalField); ok {
						mv := FieldView{
							FieldName:  f.Name,
							TypeName:   f.Type,
							Repeated:   f.Repeated,
							TagNumber:  f.Sequence,
							GoType:     pbTypesToGoType(f.Type),
							GoFlatType: pbTypesToGoFlatTypes(f.Type),
							JavaType:   pbTypesToJavaType(f.Type),
						}
						msgView.Fields = append(msgView.Fields, mv)
					}
				}
				messageViews = append(messageViews, msgView)
			}
		}
	}

	return messageViews
}

func ExtractAllEnumsViewsV2(protos []*proto.Proto) []EnumView {
	enumViews := make([]EnumView, 0)

	for _, pto := range protos {
		for _, ele := range pto.Elements {
			if enum, ok := ele.(*proto.Enum); ok {
				enumView := EnumView{
					Name:    enum.Name,
					Comment: findCommentV2(enum.Comment),
				}

				pos := 0
				for _, ele := range enum.Elements {
					if value, ok := ele.(*proto.EnumField); ok {
						fieldView := EnumFieldView{
							FieldName: value.Name,
							TagNumber: int(value.Integer),
							PosNumber: int(pos),
						}
						pos++
						enumView.Fields = append(enumView.Fields, fieldView)
					}
				}
				enumViews = append(enumViews, enumView)
			}
		}
	}

	return enumViews
}

func findCommentV2(com *proto.Comment) string {

	return ""
}

func xxx() {

	reader, _ := os.Open(`C:\Go\_gopath\src\ms\ants\play\1.proto`)
	defer reader.Close()
	parser := proto.NewParser(reader)
	def, err := parser.Parse()
	helper.NoErr(err)
	messageViews := make([]MessageView, 0)
	for _, v := range def.Elements {
		if msg, ok := v.(*proto.Message); ok {
			fmt.Println(msg)
			msgView := MessageView{
				Name:    msg.Name,
				Comment: "ccc",
			}

			for _, f2 := range msg.Elements {
				if f, ok := f2.(*proto.NormalField); ok {
					mv := FieldView{
						FieldName:  f.Name,
						TypeName:   f.Type,
						Repeated:   f.Repeated,
						TagNumber:  f.Sequence,
						GoType:     pbTypesToGoType(f.Type),
						GoFlatType: pbTypesToGoFlatTypes(f.Type),
						JavaType:   pbTypesToJavaType(f.Type),
					}
					msgView.Fields = append(msgView.Fields, mv)
				}

			}
			messageViews = append(messageViews, msgView)
			//if msg, ok := v.(*proto.Message); ok {
			//	fmt.Println(msg)
			//}
		}
		//fmt.Println(v)
	}
	helper.PertyPrint(messageViews)
	helper.PertyPrint(def)

	/*def.Elements[0].Accept(&vv{})
	formatter := proto.NewFormatter(os.Stdout, " ")

	formatter.Format(def)*/

}

type vv struct {
}

func (m *vv) VisitMessage(m1 *proto.Message) {
	print("VisitMessage me")
}

func (m *vv) VisitService(v *proto.Service) {
	print("VisitService me")
}

func (m *vv) VisitSyntax(s *proto.Syntax) {
	print("VisitSyntax me")
	print(s.Comment)
	print(s.InlineComment)
	print(s.Value)

	s.Accept(m)

}

func (m *vv) VisitPackage(p *proto.Package) {
	print("VisitPackage me")
}

func (m *vv) VisitOption(o *proto.Option) {
	print("VisitOption me")
}

func (m *vv) VisitImport(i *proto.Import) {
	print("VisitImport me")
}

func (m *vv) VisitNormalField(i *proto.NormalField) {
	print("VisitNormalField me")
}

func (m *vv) VisitEnumField(i *proto.EnumField) {
	print("implement me")
}

func (m *vv) VisitEnum(e *proto.Enum) {
	print("VisitEnum me")
}

func (m *vv) VisitComment(e *proto.Comment) {
	print("VisitComment me")
}

func (m *vv) VisitOneof(o *proto.Oneof) {
	print("VisitOneof me")
}

func (m *vv) VisitOneofField(o *proto.OneOfField) {
	print("VisitOneofField me")
}

func (m *vv) VisitReserved(r *proto.Reserved) {
	print("VisitReserved me")
}

func (m *vv) VisitRPC(r *proto.RPC) {
	print("VisitRPC me")
}

func (m *vv) VisitMapField(f *proto.MapField) {
	print("VisitMapField me")
}

func (m *vv) VisitGroup(g *proto.Group) {
	print("VisitGroup me")
}

func (m *vv) VisitExtensions(e *proto.Extensions) {
	print("VisitExtensions me")
}
