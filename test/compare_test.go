package compare

import (
	"testing"

	cp "local.packages/compare"
)

func Test_01(t *testing.T) {
	cp.Compare(t, string("kaeru"), string("kaeru"))
}

func Test_02(t *testing.T) {
	cp.Compare(t, string("kaeru"), string("🐸"))
}

func Test_03(t *testing.T) {
	cp.CompareNotEqual(t, string("kaeru"), string("kaeru"))
}

func Test_04(t *testing.T) {
	cp.CompareNotEqual(t, string("kaeru"), string("🐸"))
}
