package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	bytesPackage = protogen.GoImportPath("bytes")
	mathPackage  = protogen.GoImportPath("math")

	protoPackage = protogen.GoImportPath("google.golang.org/protobuf/proto")
)

func genEqual(g *protogen.GeneratedFile, messages []*protogen.Message) {
	for _, m := range messages {

		// Generate equal for nested messages
		if len(m.Messages) > 0 {
			genEqual(g, m.Messages)
		}

		// Do not generate extra message for map comparison
		if m.Desc.IsMapEntry() {
			continue
		}

		g.P()
		g.P(`func (x *`, m.GoIdent, `) Equal(y *`, m.GoIdent, `) bool {`)

		// Handle nil cases:
		// - messages are equal when both are nil
		// - skip comparison when one of the messages is nil
		g.P(`if x == nil || y == nil {`)
		g.P(`return x == nil && y == nil`)
		g.P(`}`)

		// Avoid comparison if both inputs are identical pointers
		g.P(`if x == y {`)
		g.P(`return true`)
		g.P(`}`)

		// Message has no fields
		if len(m.Fields) == 0 {
			g.P(`return true`)
		}

		var (
			ret = "return "
			and = " &&"
		)

		// Generate fields comparison
		for i, f := range m.Fields {

			// For last comparison do not write &&
			if i == len(m.Fields)-1 {
				and = ""
			}

			fieldName := f.GoName

			optional := f.Desc.HasOptionalKeyword()
			if optional {
				g.P(ret, `func() bool {`)
				g.P(`if x.` + fieldName + ` == nil || y.` + fieldName + ` == nil {`)
				g.P(`return x.` + fieldName + ` == nil && y.` + fieldName + ` == nil`)
				g.P(`}`)
				g.P(`ox, oy := x.Get` + fieldName + `(), y.Get` + fieldName + `()`)
				g.P(`return true &&`)

				and = ""
				ret = ""
			}

			switch {
			case f.Desc.IsList():
				g.P(ret, `func() bool {`)
				g.P(`if len(x.` + fieldName + `) != len(y.` + fieldName + `) {`)
				g.P(`return false`)
				g.P(`}`)
				g.P(`for i := 0; i < len(x.` + fieldName + `); i++ {`)
				g.P(`equal :=`)

				genEqualField(g, f, fieldName+`[i]`, "", "", optional)

				g.P(`if !equal {`)
				g.P(`return false`)
				g.P(`}`)
				g.P(`}`)
				g.P(`return true`)
				g.P(`}()`, and)

			case f.Desc.IsMap():
				g.P(ret, `func() bool {`)
				g.P(`if len(x.` + fieldName + `) != len(y.` + fieldName + `) {`)
				g.P(`return false`)
				g.P(`}`)
				g.P(`for k := range x.` + fieldName + ` {`)
				g.P(`_, hasKey := y.` + fieldName + `[k]`)
				g.P(`equal := hasKey &&`)

				genEqualField(g, f.Message.Fields[1], fieldName+`[k]`, "", "", optional)

				g.P(`if !equal {`)
				g.P(`return false`)
				g.P(`}`)
				g.P(`}`)
				g.P(`return true`)
				g.P(`}()`, and)

			default:
				genEqualField(g, f, fieldName, ret, and, optional)
			}

			if optional {
				if i == len(m.Fields)-1 {
					and = ""
				} else {
					and = " &&"
				}

				g.P(`}()`, and)
			}

			ret = ""
		}

		g.P(`}`)
	}
}

func genEqualField(g *protogen.GeneratedFile, f *protogen.Field, fieldName, ret, and string, optional bool) {
	x, y := "x."+fieldName, "y."+fieldName
	if optional {
		x, y = "ox", "oy"
	} else if f.Oneof != nil {
		x = `x.Get` + fieldName + `()`
		y = `y.Get` + fieldName + `()`
	}

	switch f.Desc.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.StringKind:
		g.P(ret, x+` == `+y, and)

	case protoreflect.FloatKind, protoreflect.DoubleKind:
		// This is what proto.Equal is doing with floats
		g.P(ret, `func() bool {`)
		g.P(`if `, mathPackage.Ident("IsNaN"), `(float64(`+x+`)) || `, mathPackage.Ident("IsNaN"), `(float64(`+y+`)) {`)
		g.P(`return `, mathPackage.Ident("IsNaN"), `(float64(`+x+`)) && `, mathPackage.Ident("IsNaN"), `(float64(`+y+`))`)
		g.P(`}`)
		g.P(`return `, x+` == `+y)
		g.P(`}()`, and)

	case protoreflect.BytesKind:
		g.P(ret, bytesPackage.Ident("Equal"), `(`+x+`, `+y+`)`, and)

	case protoreflect.MessageKind, protoreflect.GroupKind:
		switch f.Message.Location.SourceFile {
		case "google/protobuf/any.proto":
			g.P(ret, `func() bool {`)
			g.P(`if ` + x + ` == nil || ` + y + ` == nil {`)
			g.P(`return ` + x + ` == nil && ` + y + ` == nil`)
			g.P(`}`)
			g.P(`return `+x+`.TypeUrl == `+y+`.TypeUrl && `, bytesPackage.Ident("Equal"), `(`+x+`.Value, `+y+`.Value)`)
			g.P(`}()`, and)

		case "google/protobuf/api.proto",
			"google/protobuf/field_mask.proto",
			"google/protobuf/source_context.proto",
			"google/protobuf/struct.proto",
			"google/protobuf/type.proto":
			g.P(ret, protoPackage.Ident("Equal"), `(`+x+`, `+y+`)`, and)

		case "google/protobuf/duration.proto", "google/protobuf/timestamp.proto":
			g.P(ret, `func() bool {`)
			g.P(`if ` + x + ` == nil || ` + y + ` == nil {`)
			g.P(`return ` + x + ` == nil && ` + y + ` == nil`)
			g.P(`}`)
			g.P(`return ` + x + `.Seconds == ` + y + `.Seconds && ` + x + `.Nanos == ` + y + `.Nanos`)
			g.P(`}()`, and)

		case "google/protobuf/empty.proto":
			g.P(ret, `(`+x+` == nil && `+y+` == nil || `+x+` != nil && `+y+` != nil)`, and)

		case "google/protobuf/wrappers.proto":
			g.P(ret, `func() bool {`)
			g.P(`if ` + x + ` == nil || ` + y + ` == nil {`)
			g.P(`return ` + x + ` == nil && ` + y + ` == nil`)
			g.P(`}`)
			genEqualField(g, f.Message.Fields[0], "Get"+fieldName+"().Value", "return ", "", false)
			g.P(`}()`, and)

		default:
			g.P(ret, x+`.Equal(`+y+`)`, and)
		}

	default:
		g.P(ret, protoPackage.Ident("Equal"), `(`+x+`, `+y+`)`, and)
	}
}
