package attachecase

import (
	"testing"
)

func TestIotas(t *testing.T) {
	for _, test := range []struct {
		name      string
		iotaConst int
		expect    int
	}{
		{"File signifier", fsFile, 1000},
		{"Directory signifier", fsDir, 1001},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.iotaConst != test.expect {
				t.Errorf("Received %v, expected %v", test.iotaConst, test.expect)
			}
		})
	}
}

func TestFilesystem(t *testing.T) {
	for _, test := range []struct {
		name        string
		json        string
		expect      string
		expectError bool
	}{
		{},
	} {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
