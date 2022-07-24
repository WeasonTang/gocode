package cal
import (
	"testing"
)


func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper(10)测试失败期望55 实际%v\n", res)
	}
	t.Logf("test addUpper(10) succes\n")
}