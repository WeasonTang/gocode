package cal

import (
	"testing"
)

func TestSub(t *testing.T) {
	res := getSub(7, 4)
	if res != 3 {
		t.Fatalf("getSub函数测试失败期望3 实际%v\n", res)
	}
	t.Logf("getSub函数测试成功\n")

}