package symboltable

import (
	"fmt"
	"testing"
)

func TestHashOpts_verify(t *testing.T) {
	tests := []struct {
		opts HashOpts
	}{
		{HashOpts{InitialCap: 8}},
		{HashOpts{InitialCap: 32}},
		{HashOpts{InitialCap: 256}},
		{HashOpts{InitialCap: 1024}},
	}

	for i, tc := range tests {
		name := fmt.Sprintf("%d", i+1)
		t.Run(name, func(t *testing.T) {
			tc.opts.verify()
		})
	}
}
