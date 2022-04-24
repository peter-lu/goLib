package ArrayLib

import (
	"fmt"
	"testing"
)

func TestInSlice(t *testing.T) {
	fmt.Println(ArrayObj.InSlice(1, []interface{}{1, "2", 3, "4", "5"}))
	fmt.Println(ArrayObj.InSlice(2, []interface{}{1, "2", 3, "4", "5"}))
}
