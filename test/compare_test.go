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
