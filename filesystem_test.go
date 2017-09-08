package attachecase

import (
	"testing"
)

func FilesystemTest(t *testing.T) {
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
