package ArrayLib

type ArrayInt32 struct {
}
type ArrayInt32INTF interface {
	InSlice(val int32, slice []int32) bool
}

var ArrayInt32Obj ArrayInt32INTF = (*ArrayInt32)(nil)

//checkout val in slice
func (ArrayInt32Obj *ArrayInt32) InSlice(val int32, slice []int32) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
