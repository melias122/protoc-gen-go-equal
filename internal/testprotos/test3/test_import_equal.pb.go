// Code generated by protoc-gen-equal-go. DO NOT EDIT.
// source: internal/testprotos/test3/test_import.proto

package test3

func (x *ImportMessage) Equal(y *ImportMessage) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	if x == y {
		return true
	}
	return true
}