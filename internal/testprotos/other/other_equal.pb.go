// Code generated by protoc-gen-equal-go. DO NOT EDIT.
// source: internal/testprotos/other/other.proto

package other

func (x *OtherMessage) Equal(y *OtherMessage) bool {
	if x == y {
		return true
	}
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	if x.I != y.I {
		return false
	}
	return true
}
