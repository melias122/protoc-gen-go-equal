package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	mathPackage  = protogen.GoImportPath("math")
	protoPackage = protogen.GoImportPath("google.golang.org/protobuf/proto")
)

func genEqual(g *protogen.GeneratedFile, messages []*protogen.Message, proto3 bool) {
	for _, m := range messages {

		// Generate equal for nested messages
		if len(m.Messages) > 0 {
			genEqual(g, m.Messages, proto3)
		}

		// Do not generate extra message for map comparison
		if m.Desc.IsMapEntry() {
			continue
		}

		g.P()
		g.P(`func (x *`, m.GoIdent, `) Equal(y *`, m.GoIdent, `) bool {`)

		// Avoid comparison if both inputs are identical pointers
		g.P(`if x == y {`)
		g.P(`return true`)
		g.P(`}`)

		// Handle nil cases:
		// - messages are equal when both are nil
		// - skip comparison when one of the messages is nil
		g.P(`if x == nil || y == nil {`)
		g.P(`return x == nil && y == nil`)
		g.P(`}`)

		// Generate fields comparison
		for _, f := range m.Fields {

			fieldName := f.GoName

			switch {
			case f.Desc.IsList():
				g.P(`if len(x.` + fieldName + `) != len(y.` + fieldName + `) {`)
				g.P(`return false`)
				g.P(`}`)
				g.P(`for i := 0; i < len(x.` + fieldName + `); i++ {`)

				genEqualField(g, f, fieldName+`[i]`, proto3, true)

				g.P(`}`)

			case f.Desc.IsMap():
				g.P(`if len(x.` + fieldName + `) != len(y.` + fieldName + `) {`)
				g.P(`return false`)
				g.P(`}`)
				g.P(`for k := range x.` + fieldName + ` {`)
				g.P(`_, ok := y.` + fieldName + `[k]`)
				g.P(`if !ok {`)
				g.P(`return false`)
				g.P(`}`)

				genEqualField(g, f.Message.Fields[1], fieldName+`[k]`, proto3, true)

				g.P(`}`)

			default:
				genEqualField(g, f, fieldName, proto3, false)
			}
		}

		g.P(`return true`)
		g.P(`}`)
	}
}

func genEqualField(g *protogen.GeneratedFile, f *protogen.Field, fieldName string, proto3 bool, repeated bool) {
	// Some of these lines are stolen from vtprotobuf equal
	oneof := f.Oneof != nil && !f.Oneof.Desc.IsSynthetic() || (f.Desc.ContainingOneof() != nil && !proto3)
	nullable := (f.Message != nil || (f.Oneof != nil && f.Oneof.Desc.IsSynthetic()) || (!proto3 && !oneof)) && !repeated

	x, y := "x."+fieldName, "y."+fieldName
	if oneof {
		x, y = "x.Get"+fieldName+"()", "y.Get"+fieldName+"()"
	}

	switch f.Desc.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.StringKind:
		if nullable {
			g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {`)
		} else {
			g.P(`if `, x+` != `+y, ` {`)
		}
		g.P(`return false`)
		g.P(`}`)

	case protoreflect.FloatKind, protoreflect.DoubleKind:
		// TODO: math.IsNaN
		if nullable {
			g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || (`, mathPackage.Ident("IsNaN"), `(float64(*p)) && !`, mathPackage.Ident("IsNaN"), `(float64(*q)) || !`, mathPackage.Ident("IsNaN"), `(float64(*p)) && `, mathPackage.Ident("IsNaN"), `(float64(*q))) || (!`, mathPackage.Ident("IsNaN"), `(float64(*p)) && !`, mathPackage.Ident("IsNaN"), `(float64(*q)) && *p != *q))) {`)
		} else {
			g.P(`if (`, mathPackage.Ident("IsNaN"), `(float64(`, x, `)) && !`, mathPackage.Ident("IsNaN"), `(float64(`, y, `)) || !`, mathPackage.Ident("IsNaN"), `(float64(`, x, `)) && `, mathPackage.Ident("IsNaN"), `(float64(`, y, `))) || (!`, mathPackage.Ident("IsNaN"), `(float64(`, x, `)) && !`, mathPackage.Ident("IsNaN"), `(float64(`, y, `)) && `, x, ` != `, y, `) {`)
		}
		g.P(`return false`)
		g.P(`}`)

	case protoreflect.BytesKind:
		if nullable {
			g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || string(p) != string(q))) {`)
		} else {
			g.P(`if string(` + x + `) != string(` + y + `) {`)
		}
		g.P(`return false`)
		g.P(`}`)

	case protoreflect.MessageKind, protoreflect.GroupKind:
		switch f.Message.Location.SourceFile {
		case "google/protobuf/any.proto":
			g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || p.TypeUrl != q.TypeUrl || string(p.Value) != string(q.Value))) {`)
			g.P(`return false`)
			g.P(`}`)

		case "google/protobuf/duration.proto", "google/protobuf/timestamp.proto":
			g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || p.Seconds != q.Seconds || p.Nanos != q.Nanos)) {`)
			g.P(`return false`)
			g.P(`}`)

		case "google/protobuf/empty.proto":
			g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && q == nil) {`)
			g.P(`return false`)
			g.P(`}`)

		case "google/protobuf/wrappers.proto":
			wrappersBytes := f.Message.Fields[0].Desc.Kind() == protoreflect.BytesKind
			if wrappersBytes {
				g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || string(p.Value) != string(q.Value))) {`)
			} else {
				g.P(`if p, q := `, x, `, `, y, `; (p == nil && q != nil) || (p != nil && (q == nil || p.Value != q.Value)) {`)
			}
			g.P(`return false`)
			g.P(`}`)

		default:
			isLocalMessage := f.Message != nil && f.Message.Desc != nil && f.Message.Desc.ParentFile() != nil && isLocalPackage[string(f.Message.Desc.ParentFile().Package())]
			if isLocalMessage {
				g.P(`if !`, x, `.Equal(`, y, `) {`)
				g.P(`	return false`)
				g.P(`}`)
			} else {
				g.P(`if equal, ok := interface{}(`, x, `).(interface { Equal(*`, g.QualifiedGoIdent(f.Message.GoIdent), `) bool }); !ok || !equal.Equal(`, y, `) {`)
				g.P(`	return false`)
				g.P(`} else if !`, protoPackage.Ident("Equal"), `(`, x, `, `, y, `) {`)
				g.P(`	return false`)
				g.P(`}`)
			}
		}

	// Fallback to proto.Equal
	default:
		g.P(`if !`, protoPackage.Ident("Equal"), `(`+x+`, `+y+`) {`)
		g.P(`return false`)
		g.P(`}`)
	}
}
