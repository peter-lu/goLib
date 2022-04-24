package ArrayLib

type ArrayString struct {
}
type ArrayStringINTF interface {
	InSlice(val string, slice []string) bool
}

var ArrayStringObj ArrayStringINTF = (*ArrayString)(nil)

//checkout val in slice
func (ArrayStringObj *ArrayString) InSlice(val string, slice []string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
