// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto3test

import (
	"math"
	"testing"
	"time"

	"github.com/melias122/protoc-gen-go-equal/internal/testprotos/other"
	test3pb "github.com/melias122/protoc-gen-go-equal/internal/testprotos/test3" // "google.golang.org/protobuf/internal/testprotos/test3"
	testpb "github.com/melias122/protoc-gen-go-equal/internal/testprotos/test3"  // "google.golang.org/protobuf/internal/testprotos/test"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func mustAny(t *testing.T, msg proto.Message) *anypb.Any {
	anymsg, err := anypb.New(msg)
	if err != nil {
		t.Fatal(err)
	}
	return anymsg
}

func TestEqual(t *testing.T) {
	identicalPtrPb := &testpb.TestAllTypes{MapStringString: map[string]string{"a": "b", "c": "d"}}

	tests := []struct {
		x, y *testpb.TestAllTypes
		eq   bool
	}{
		{
			x:  nil,
			y:  nil,
			eq: true,
		}, {
			x:  (*testpb.TestAllTypes)(nil),
			y:  nil,
			eq: true, // proto.Equal had eq: false,
		}, {
			x:  (*testpb.TestAllTypes)(nil),
			y:  (*testpb.TestAllTypes)(nil),
			eq: true,
		}, {
			x:  new(testpb.TestAllTypes),
			y:  (*testpb.TestAllTypes)(nil),
			eq: false,
		}, {
			x:  new(testpb.TestAllTypes),
			y:  new(testpb.TestAllTypes),
			eq: true,
		},

		// Identical input pointers
		{
			x:  identicalPtrPb,
			y:  identicalPtrPb,
			eq: true,
		},

		// Singulars.
		// NestedMessage  singular_nested_message  = 98;
		// ForeignMessage singular_foreign_message = 99;
		// ImportMessage  singular_import_message  = 100;
		// NestedEnum     singular_nested_enum     = 101;
		// ForeignEnum    singular_foreign_enum    = 102;
		// ImportEnum     singular_import_enum     = 103;
		// Scalars.
		{
			x: &testpb.TestAllTypes{SingularInt32: 1},
			y: &testpb.TestAllTypes{SingularInt32: 2},
		}, {
			x: &testpb.TestAllTypes{SingularInt64: 1},
			y: &testpb.TestAllTypes{SingularInt64: 2},
		}, {
			x: &testpb.TestAllTypes{SingularUint32: 1},
			y: &testpb.TestAllTypes{SingularUint32: 2},
		}, {
			x: &testpb.TestAllTypes{SingularUint64: 1},
			y: &testpb.TestAllTypes{SingularUint64: 2},
		}, {
			x: &testpb.TestAllTypes{SingularSint32: 1},
			y: &testpb.TestAllTypes{SingularSint32: 2},
		}, {
			x: &testpb.TestAllTypes{SingularSint64: 1},
			y: &testpb.TestAllTypes{SingularSint64: 2},
		}, {
			x: &testpb.TestAllTypes{SingularFixed32: 1},
			y: &testpb.TestAllTypes{SingularFixed32: 2},
		}, {
			x: &testpb.TestAllTypes{SingularFixed64: 1},
			y: &testpb.TestAllTypes{SingularFixed64: 2},
		}, {
			x: &testpb.TestAllTypes{SingularSfixed32: 1},
			y: &testpb.TestAllTypes{SingularSfixed32: 2},
		}, {
			x: &testpb.TestAllTypes{SingularSfixed64: 1},
			y: &testpb.TestAllTypes{SingularSfixed64: 2},
		}, {
			x:  &testpb.TestAllTypes{SingularFloat: 0},
			y:  &testpb.TestAllTypes{SingularFloat: 0},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularDouble: 0},
			y:  &testpb.TestAllTypes{SingularDouble: 0},
			eq: true,
		}, {
			x: &testpb.TestAllTypes{SingularFloat: 1},
			y: &testpb.TestAllTypes{SingularFloat: 2},
		}, {
			x: &testpb.TestAllTypes{SingularDouble: 1},
			y: &testpb.TestAllTypes{SingularDouble: 2},
		}, {
			x: &testpb.TestAllTypes{SingularFloat: 2},
			y: &testpb.TestAllTypes{SingularFloat: 1},
		}, {
			x: &testpb.TestAllTypes{SingularDouble: 2},
			y: &testpb.TestAllTypes{SingularDouble: 1},
		}, {
			x: &testpb.TestAllTypes{SingularFloat: 0},
			y: &testpb.TestAllTypes{SingularFloat: float32(math.NaN())},
		}, {
			x: &testpb.TestAllTypes{SingularDouble: 0},
			y: &testpb.TestAllTypes{SingularDouble: math.NaN()},
		}, {
			x:  &testpb.TestAllTypes{SingularDouble: math.NaN()},
			y:  &testpb.TestAllTypes{SingularDouble: math.NaN()},
			eq: true,
		}, {
			x: &testpb.TestAllTypes{SingularBool: true},
			y: &testpb.TestAllTypes{SingularBool: false},
		}, {
			x: &testpb.TestAllTypes{SingularString: "a"},
			y: &testpb.TestAllTypes{SingularString: "b"},
		}, {
			x: &testpb.TestAllTypes{SingularBytes: []byte("a")},
			y: &testpb.TestAllTypes{SingularBytes: []byte("b")},
		}, {
			x: &testpb.TestAllTypes{SingularNestedEnum: testpb.TestAllTypes_FOO},
			y: &testpb.TestAllTypes{SingularNestedEnum: testpb.TestAllTypes_BAR},
		}, {
			x:  &testpb.TestAllTypes{SingularInt32: 2},
			y:  &testpb.TestAllTypes{SingularInt32: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularInt64: 2},
			y:  &testpb.TestAllTypes{SingularInt64: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularUint32: 2},
			y:  &testpb.TestAllTypes{SingularUint32: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularUint64: 2},
			y:  &testpb.TestAllTypes{SingularUint64: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularSint32: 2},
			y:  &testpb.TestAllTypes{SingularSint32: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularSint64: 2},
			y:  &testpb.TestAllTypes{SingularSint64: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularFixed32: 2},
			y:  &testpb.TestAllTypes{SingularFixed32: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularFixed64: 2},
			y:  &testpb.TestAllTypes{SingularFixed64: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularSfixed32: 2},
			y:  &testpb.TestAllTypes{SingularSfixed32: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularSfixed64: 2},
			y:  &testpb.TestAllTypes{SingularSfixed64: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularFloat: 2},
			y:  &testpb.TestAllTypes{SingularFloat: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularDouble: 2},
			y:  &testpb.TestAllTypes{SingularDouble: 2},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularFloat: float32(math.NaN())},
			y:  &testpb.TestAllTypes{SingularFloat: float32(math.NaN())},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularDouble: math.NaN()},
			y:  &testpb.TestAllTypes{SingularDouble: math.NaN()},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularBool: true},
			y:  &testpb.TestAllTypes{SingularBool: true},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularBool: false},
			y:  &testpb.TestAllTypes{SingularBool: false},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularString: "abc"},
			y:  &testpb.TestAllTypes{SingularString: "abc"},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularBytes: []byte("abc")},
			y:  &testpb.TestAllTypes{SingularBytes: []byte("abc")},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{SingularNestedEnum: testpb.TestAllTypes_FOO},
			y:  &testpb.TestAllTypes{SingularNestedEnum: testpb.TestAllTypes_FOO},
			eq: true,
		},

		// Scalars.
		{
			x: &testpb.TestAllTypes{OptionalInt32: proto.Int32(1)},
			y: &testpb.TestAllTypes{OptionalInt32: proto.Int32(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalInt64: proto.Int64(1)},
			y: &testpb.TestAllTypes{OptionalInt64: proto.Int64(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalUint32: proto.Uint32(1)},
			y: &testpb.TestAllTypes{OptionalUint32: proto.Uint32(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalUint64: proto.Uint64(1)},
			y: &testpb.TestAllTypes{OptionalUint64: proto.Uint64(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalSint32: proto.Int32(1)},
			y: &testpb.TestAllTypes{OptionalSint32: proto.Int32(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalSint64: proto.Int64(1)},
			y: &testpb.TestAllTypes{OptionalSint64: proto.Int64(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalFixed32: proto.Uint32(1)},
			y: &testpb.TestAllTypes{OptionalFixed32: proto.Uint32(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalFixed64: proto.Uint64(1)},
			y: &testpb.TestAllTypes{OptionalFixed64: proto.Uint64(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalSfixed32: proto.Int32(1)},
			y: &testpb.TestAllTypes{OptionalSfixed32: proto.Int32(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalSfixed64: proto.Int64(1)},
			y: &testpb.TestAllTypes{OptionalSfixed64: proto.Int64(2)},
		}, {
			x:  &testpb.TestAllTypes{OptionalFloat: proto.Float32(0)},
			y:  &testpb.TestAllTypes{OptionalFloat: proto.Float32(0)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(0)},
			y:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(0)},
			eq: true,
		}, {
			x: &testpb.TestAllTypes{OptionalFloat: proto.Float32(1)},
			y: &testpb.TestAllTypes{OptionalFloat: proto.Float32(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalDouble: proto.Float64(1)},
			y: &testpb.TestAllTypes{OptionalDouble: proto.Float64(2)},
		}, {
			x: &testpb.TestAllTypes{OptionalFloat: proto.Float32(2)},
			y: &testpb.TestAllTypes{OptionalFloat: proto.Float32(1)},
		}, {
			x: &testpb.TestAllTypes{OptionalDouble: proto.Float64(2)},
			y: &testpb.TestAllTypes{OptionalDouble: proto.Float64(1)},
		}, {
			x: &testpb.TestAllTypes{OptionalFloat: proto.Float32(0)},
			y: &testpb.TestAllTypes{OptionalFloat: proto.Float32(float32(math.NaN()))},
		}, {
			x: &testpb.TestAllTypes{OptionalDouble: proto.Float64(0)},
			y: &testpb.TestAllTypes{OptionalDouble: proto.Float64(float64(math.NaN()))},
		}, {
			x:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(float64(math.NaN()))},
			y:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(float64(math.NaN()))},
			eq: true,
		}, {
			x: &testpb.TestAllTypes{OptionalBool: proto.Bool(true)},
			y: &testpb.TestAllTypes{OptionalBool: proto.Bool(false)},
		}, {
			x: &testpb.TestAllTypes{OptionalString: proto.String("a")},
			y: &testpb.TestAllTypes{OptionalString: proto.String("b")},
		}, {
			x: &testpb.TestAllTypes{OptionalBytes: []byte("a")},
			y: &testpb.TestAllTypes{OptionalBytes: []byte("b")},
		}, {
			x: &testpb.TestAllTypes{OptionalNestedEnum: testpb.TestAllTypes_FOO.Enum()},
			y: &testpb.TestAllTypes{OptionalNestedEnum: testpb.TestAllTypes_BAR.Enum()},
		}, {
			x:  &testpb.TestAllTypes{OptionalInt32: proto.Int32(2)},
			y:  &testpb.TestAllTypes{OptionalInt32: proto.Int32(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalInt64: proto.Int64(2)},
			y:  &testpb.TestAllTypes{OptionalInt64: proto.Int64(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalUint32: proto.Uint32(2)},
			y:  &testpb.TestAllTypes{OptionalUint32: proto.Uint32(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalUint64: proto.Uint64(2)},
			y:  &testpb.TestAllTypes{OptionalUint64: proto.Uint64(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalSint32: proto.Int32(2)},
			y:  &testpb.TestAllTypes{OptionalSint32: proto.Int32(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalSint64: proto.Int64(2)},
			y:  &testpb.TestAllTypes{OptionalSint64: proto.Int64(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalFixed32: proto.Uint32(2)},
			y:  &testpb.TestAllTypes{OptionalFixed32: proto.Uint32(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalFixed64: proto.Uint64(2)},
			y:  &testpb.TestAllTypes{OptionalFixed64: proto.Uint64(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalSfixed32: proto.Int32(2)},
			y:  &testpb.TestAllTypes{OptionalSfixed32: proto.Int32(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalSfixed64: proto.Int64(2)},
			y:  &testpb.TestAllTypes{OptionalSfixed64: proto.Int64(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalFloat: proto.Float32(2)},
			y:  &testpb.TestAllTypes{OptionalFloat: proto.Float32(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(2)},
			y:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(2)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalFloat: proto.Float32(float32(math.NaN()))},
			y:  &testpb.TestAllTypes{OptionalFloat: proto.Float32(float32(math.NaN()))},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(float64(math.NaN()))},
			y:  &testpb.TestAllTypes{OptionalDouble: proto.Float64(float64(math.NaN()))},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalBool: proto.Bool(true)},
			y:  &testpb.TestAllTypes{OptionalBool: proto.Bool(true)},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalString: proto.String("abc")},
			y:  &testpb.TestAllTypes{OptionalString: proto.String("abc")},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalBytes: []byte("abc")},
			y:  &testpb.TestAllTypes{OptionalBytes: []byte("abc")},
			eq: true,
		}, {
			x:  &testpb.TestAllTypes{OptionalNestedEnum: testpb.TestAllTypes_FOO.Enum()},
			y:  &testpb.TestAllTypes{OptionalNestedEnum: testpb.TestAllTypes_FOO.Enum()},
			eq: true,
		},

		// Proto3 presence.
		{
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalInt32: proto.Int32(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalInt64: proto.Int64(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalUint32: proto.Uint32(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalUint64: proto.Uint64(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalSint32: proto.Int32(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalSint64: proto.Int64(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalFixed32: proto.Uint32(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalFixed64: proto.Uint64(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalSfixed32: proto.Int32(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalSfixed64: proto.Int64(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalFloat: proto.Float32(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalDouble: proto.Float64(0)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalBool: proto.Bool(false)},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalString: proto.String("")},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalBytes: []byte{}},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalNestedEnum: test3pb.TestAllTypes_FOO.Enum()},
		},

		// Messages.
		{
			x: &testpb.TestAllTypes{OptionalNestedMessage: &testpb.TestAllTypes_NestedMessage{
				A: proto.Int32(1),
			}},
			y: &testpb.TestAllTypes{OptionalNestedMessage: &testpb.TestAllTypes_NestedMessage{
				A: proto.Int32(2),
			}},
		}, {
			x: &testpb.TestAllTypes{},
			y: &testpb.TestAllTypes{OptionalNestedMessage: &testpb.TestAllTypes_NestedMessage{}},
		}, {
			x: &test3pb.TestAllTypes{},
			y: &test3pb.TestAllTypes{OptionalNestedMessage: &test3pb.TestAllTypes_NestedMessage{}},
		},

		// Lists.
		{
			x: &testpb.TestAllTypes{RepeatedInt32: []int32{1}},
			y: &testpb.TestAllTypes{RepeatedInt32: []int32{1, 2}},
		}, {
			x: &testpb.TestAllTypes{RepeatedInt32: []int32{1, 2}},
			y: &testpb.TestAllTypes{RepeatedInt32: []int32{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedInt64: []int64{1, 2}},
			y: &testpb.TestAllTypes{RepeatedInt64: []int64{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedUint32: []uint32{1, 2}},
			y: &testpb.TestAllTypes{RepeatedUint32: []uint32{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedUint64: []uint64{1, 2}},
			y: &testpb.TestAllTypes{RepeatedUint64: []uint64{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedSint32: []int32{1, 2}},
			y: &testpb.TestAllTypes{RepeatedSint32: []int32{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedSint64: []int64{1, 2}},
			y: &testpb.TestAllTypes{RepeatedSint64: []int64{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedFixed32: []uint32{1, 2}},
			y: &testpb.TestAllTypes{RepeatedFixed32: []uint32{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedFixed64: []uint64{1, 2}},
			y: &testpb.TestAllTypes{RepeatedFixed64: []uint64{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedSfixed32: []int32{1, 2}},
			y: &testpb.TestAllTypes{RepeatedSfixed32: []int32{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedSfixed64: []int64{1, 2}},
			y: &testpb.TestAllTypes{RepeatedSfixed64: []int64{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedFloat: []float32{1, 2}},
			y: &testpb.TestAllTypes{RepeatedFloat: []float32{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedDouble: []float64{1, 2}},
			y: &testpb.TestAllTypes{RepeatedDouble: []float64{1, 3}},
		}, {
			x: &testpb.TestAllTypes{RepeatedBool: []bool{true, false}},
			y: &testpb.TestAllTypes{RepeatedBool: []bool{true, true}},
		}, {
			x: &testpb.TestAllTypes{RepeatedString: []string{"a", "b"}},
			y: &testpb.TestAllTypes{RepeatedString: []string{"a", "c"}},
		}, {
			x: &testpb.TestAllTypes{RepeatedBytes: [][]byte{[]byte("a"), []byte("b")}},
			y: &testpb.TestAllTypes{RepeatedBytes: [][]byte{[]byte("a"), []byte("c")}},
		}, {
			x: &testpb.TestAllTypes{RepeatedNestedEnum: []testpb.TestAllTypes_NestedEnum{testpb.TestAllTypes_FOO}},
			y: &testpb.TestAllTypes{RepeatedNestedEnum: []testpb.TestAllTypes_NestedEnum{testpb.TestAllTypes_BAR}},
		}, {
			x: &testpb.TestAllTypes{RepeatedNestedMessage: []*testpb.TestAllTypes_NestedMessage{
				{A: proto.Int32(1)},
				{A: proto.Int32(2)},
			}},
			y: &testpb.TestAllTypes{RepeatedNestedMessage: []*testpb.TestAllTypes_NestedMessage{
				{A: proto.Int32(1)},
				{A: proto.Int32(3)},
			}},
		},

		// Maps: various configurations.
		{
			x: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2}},
			y: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{3: 4}},
		}, {
			x: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2}},
			y: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2, 3: 4}},
		}, {
			x: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2}},
		},

		// Maps: various types.
		{
			x: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapInt32Int32: map[int32]int32{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapInt64Int64: map[int64]int64{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapInt64Int64: map[int64]int64{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapUint32Uint32: map[uint32]uint32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapUint32Uint32: map[uint32]uint32{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapUint64Uint64: map[uint64]uint64{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapUint64Uint64: map[uint64]uint64{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapSint32Sint32: map[int32]int32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapSint32Sint32: map[int32]int32{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapSint64Sint64: map[int64]int64{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapSint64Sint64: map[int64]int64{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapFixed32Fixed32: map[uint32]uint32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapFixed32Fixed32: map[uint32]uint32{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapFixed64Fixed64: map[uint64]uint64{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapFixed64Fixed64: map[uint64]uint64{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapSfixed32Sfixed32: map[int32]int32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapSfixed32Sfixed32: map[int32]int32{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapSfixed64Sfixed64: map[int64]int64{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapSfixed64Sfixed64: map[int64]int64{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapInt32Float: map[int32]float32{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapInt32Float: map[int32]float32{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapInt32Double: map[int32]float64{1: 2, 3: 4}},
			y: &testpb.TestAllTypes{MapInt32Double: map[int32]float64{1: 2, 3: 5}},
		}, {
			x: &testpb.TestAllTypes{MapBoolBool: map[bool]bool{true: false, false: true}},
			y: &testpb.TestAllTypes{MapBoolBool: map[bool]bool{true: false, false: false}},
		}, {
			x: &testpb.TestAllTypes{MapStringString: map[string]string{"a": "b", "c": "d"}},
			y: &testpb.TestAllTypes{MapStringString: map[string]string{"a": "b", "c": "e"}},
		}, {
			x: &testpb.TestAllTypes{MapStringBytes: map[string][]byte{"a": []byte("b"), "c": []byte("d")}},
			y: &testpb.TestAllTypes{MapStringBytes: map[string][]byte{"a": []byte("b"), "c": []byte("e")}},
		}, {
			x: &testpb.TestAllTypes{MapStringNestedMessage: map[string]*testpb.TestAllTypes_NestedMessage{
				"a": {A: proto.Int32(1)},
				"b": {A: proto.Int32(2)},
			}},
			y: &testpb.TestAllTypes{MapStringNestedMessage: map[string]*testpb.TestAllTypes_NestedMessage{
				"a": {A: proto.Int32(1)},
				"b": {A: proto.Int32(3)},
			}},
		}, {
			x: &testpb.TestAllTypes{MapStringNestedEnum: map[string]testpb.TestAllTypes_NestedEnum{
				"a": testpb.TestAllTypes_FOO,
				"b": testpb.TestAllTypes_BAR,
			}},
			y: &testpb.TestAllTypes{MapStringNestedEnum: map[string]testpb.TestAllTypes_NestedEnum{
				"a": testpb.TestAllTypes_FOO,
				"b": testpb.TestAllTypes_BAZ,
			}},
		},

		// Oneof
		{
			x: &test3pb.TestAllTypes{
				OneofField: &test3pb.TestAllTypes_OneofUint32{
					OneofUint32: 1,
				},
			},
		},
		{
			x: &test3pb.TestAllTypes{
				OneofField: &test3pb.TestAllTypes_OneofWrappersStringValue{
					OneofWrappersStringValue: wrapperspb.String("s"),
				},
			},
		},

		// Known Types
		{
			x: &test3pb.TestAllTypes{
				Any: mustAny(t, &testpb.TestAllTypes{OptionalInt32: proto.Int32(1)}),
			},
			y: &test3pb.TestAllTypes{
				Any: mustAny(t, &testpb.TestAllTypes{OptionalInt32: proto.Int32(2)}),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				Any: mustAny(t, &testpb.TestAllTypes{OptionalInt32: proto.Int32(1)}),
			},
			y: &test3pb.TestAllTypes{
				Any: mustAny(t, &testpb.TestAllTypes{OptionalInt32: proto.Int32(1)}),
			},
			eq: true,
		},
		{
			x: &test3pb.TestAllTypes{
				Duration: durationpb.New(time.Second),
			},
			y: &test3pb.TestAllTypes{
				Duration: durationpb.New(time.Hour),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				Duration: durationpb.New(time.Hour),
			},
			y: &test3pb.TestAllTypes{
				Duration: durationpb.New(time.Hour),
			},
			eq: true,
		},
		{
			x: &test3pb.TestAllTypes{
				Empty: &emptypb.Empty{},
			},
			y: &test3pb.TestAllTypes{
				Empty: nil,
			},
		},
		{
			x: &test3pb.TestAllTypes{
				Empty: nil,
			},
			y: &test3pb.TestAllTypes{
				Empty: nil,
			},
			eq: true,
		},
		{
			x: &test3pb.TestAllTypes{
				Empty: &emptypb.Empty{},
			},
			y: &test3pb.TestAllTypes{
				Empty: &emptypb.Empty{},
			},
			eq: true,
		},
		{
			x: &test3pb.TestAllTypes{
				Timestamp: timestamppb.Now(),
			},
			y: &test3pb.TestAllTypes{
				Timestamp: timestamppb.New(time.Now().Add(time.Hour)),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				Timestamp: timestamppb.New(time.Time{}),
			},
			y: &test3pb.TestAllTypes{
				Timestamp: timestamppb.New(time.Time{}),
			},
			eq: true,
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersBoolValue: wrapperspb.Bool(true),
			},
			y: &test3pb.TestAllTypes{
				WrappersBoolValue: wrapperspb.Bool(false),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersBytesValue: wrapperspb.Bytes([]byte("hello")),
			},
			y: &test3pb.TestAllTypes{
				WrappersBytesValue: wrapperspb.Bytes([]byte("world")),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersDoubleValue: wrapperspb.Double(.5),
			},
			y: &test3pb.TestAllTypes{
				WrappersDoubleValue: wrapperspb.Double(.55),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersFloatValue: wrapperspb.Float(.5),
			},
			y: &test3pb.TestAllTypes{
				WrappersFloatValue: wrapperspb.Float(.55),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersInt32Value: wrapperspb.Int32(5),
			},
			y: &test3pb.TestAllTypes{
				WrappersInt32Value: wrapperspb.Int32(7),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersInt64Value: wrapperspb.Int64(5),
			},
			y: &test3pb.TestAllTypes{
				WrappersInt64Value: wrapperspb.Int64(7),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersStringValue: wrapperspb.String("s1"),
			},
			y: &test3pb.TestAllTypes{
				WrappersStringValue: wrapperspb.String("s2"),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersUint32Value: wrapperspb.UInt32(5),
			},
			y: &test3pb.TestAllTypes{
				WrappersUint32Value: wrapperspb.UInt32(7),
			},
		},
		{
			x: &test3pb.TestAllTypes{
				WrappersUint64Value: wrapperspb.UInt64(5),
			},
			y: &test3pb.TestAllTypes{
				WrappersUint64Value: wrapperspb.UInt64(7),
			},
		},

		{
			x: &test3pb.TestAllTypes{
				OtherMessage: &other.OtherMessage{I: 1},
			},
			y: &test3pb.TestAllTypes{
				OtherMessage: &other.OtherMessage{I: 2},
			},
		},
		{
			x: &test3pb.TestAllTypes{
				OtherMessage: &other.OtherMessage{I: 2},
			},
			y: &test3pb.TestAllTypes{
				OtherMessage: &other.OtherMessage{I: 1},
			},
		},
		{
			x: &test3pb.TestAllTypes{
				OtherMessage: &other.OtherMessage{I: 1},
			},
			y: &test3pb.TestAllTypes{
				OtherMessage: &other.OtherMessage{I: 1},
			},
			eq: true,
		},
	}

	for _, tt := range tests {
		if !tt.eq && !tt.x.Equal(tt.x) {
			t.Errorf("Equal(x, x) = false, want true\n==== x ====\n%v", prototext.Format(tt.x))
		}
		if !tt.eq && !tt.y.Equal(tt.y) {
			t.Errorf("Equal(y, y) = false, want true\n==== y ====\n%v", prototext.Format(tt.y))
		}
		if eq := tt.x.Equal(tt.y); eq != tt.eq {
			t.Errorf("Equal(x, y) = %v, want %v\n==== x ====\n%v==== y ====\n%v", eq, tt.eq, prototext.Format(tt.x), prototext.Format(tt.y))
		}
	}
}

func BenchmarkProtoEqualWithSmallEmpty(b *testing.B) {
	x := &testpb.ForeignMessage{}
	y := &testpb.ForeignMessage{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Equal(x, y)
	}
}

func BenchmarkEqualWithSmallEmpty(b *testing.B) {
	x := &testpb.ForeignMessage{}
	y := &testpb.ForeignMessage{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(y)
	}
}

func BenchmarkProtoEqualWithIdenticalPtrEmpty(b *testing.B) {
	x := &testpb.ForeignMessage{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Equal(x, x)
	}
}

func BenchmarkEqualWithIdenticalPtrEmpty(b *testing.B) {
	x := &testpb.ForeignMessage{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(x)
	}
}

func BenchmarkProtoEqualWithLargeEmpty(b *testing.B) {
	x := &testpb.TestAllTypes{}
	y := &testpb.TestAllTypes{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Equal(x, y)
	}
}

func BenchmarkEqualWithLargeEmpty(b *testing.B) {
	x := &testpb.TestAllTypes{}
	y := &testpb.TestAllTypes{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(y)
	}
}

func makeNested(depth int) *testpb.TestAllTypes {
	if depth <= 0 {
		return nil
	}
	return &testpb.TestAllTypes{
		OptionalNestedMessage: &testpb.TestAllTypes_NestedMessage{
			Corecursive: makeNested(depth - 1),
		},
	}
}

func BenchmarkProtoEqualWithDeeplyNestedEqual(b *testing.B) {
	x := makeNested(20)
	y := makeNested(20)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Equal(x, y)
	}
}

func BenchmarkEqualWithDeeplyNestedEqual(b *testing.B) {
	x := makeNested(20)
	y := makeNested(20)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(y)
	}
}

func BenchmarkProtoEqualWithDeeplyNestedDifferent(b *testing.B) {
	x := makeNested(20)
	y := makeNested(21)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Equal(x, y)
	}
}

func BenchmarkEqualWithDeeplyNestedDifferent(b *testing.B) {
	x := makeNested(20)
	y := makeNested(21)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(y)
	}
}

func BenchmarkProtoEqualWithDeeplyNestedIdenticalPtr(b *testing.B) {
	x := makeNested(20)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Equal(x, x)
	}
}

func BenchmarkEqualWithDeeplyNestedIdenticalPtr(b *testing.B) {
	x := makeNested(20)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Equal(x)
	}
}
