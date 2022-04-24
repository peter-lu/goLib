package ArrayLib

type ArrayInt64 struct {
}
type ArrayInt64INTF interface {
	InSlice(val int64, slice []int64) bool
}

var ArrayInt64Obj ArrayInt64INTF = (*ArrayInt64)(nil)

//checkout val in slice
func (ArrayInt64Obj *ArrayInt64) InSlice(val int64, slice []int64) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
