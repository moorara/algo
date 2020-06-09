package compare

import (
	"testing"
	"time"
)

func TestUint(t *testing.T) {
	tests := []struct {
		name           string
		a, b           uint
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Uint(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		name           string
		a, b           uint8
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Uint8(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		name           string
		a, b           uint16
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Uint16(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name           string
		a, b           uint32
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Uint32(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name           string
		a, b           uint64
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Uint64(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name           string
		a, b           int
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Int(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	tests := []struct {
		name           string
		a, b           int8
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Int8(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	tests := []struct {
		name           string
		a, b           int16
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Int16(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		name           string
		a, b           int32
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Int32(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		name           string
		a, b           int64
		expectedResult int
	}{
		{"Less", 1, 2, -1},
		{"Equal", 2, 2, 0},
		{"Greater", 3, 2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Int64(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		name           string
		a, b           float32
		expectedResult int
	}{
		{"Less", 1.0, 2.0, -1},
		{"Equal", 2.0, 2.0, 0},
		{"Greater", 3.0, 2.0, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Float32(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name           string
		a, b           float64
		expectedResult int
	}{
		{"Less", 1.0, 2.0, -1},
		{"Equal", 2.0, 2.0, 0},
		{"Greater", 3.0, 2.0, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Float64(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name           string
		a, b           string
		expectedResult int
	}{
		{"Less", "A", "B", -1},
		{"Equal", "B", "B", 0},
		{"Greater", "C", "B", 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := String(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestTime(t *testing.T) {
	t1, _ := time.Parse(time.Kitchen, "10:00AM")
	t2, _ := time.Parse(time.Kitchen, "12:00PM")
	t3, _ := time.Parse(time.Kitchen, "20:00PM")

	tests := []struct {
		name           string
		a, b           time.Time
		expectedResult int
	}{
		{"Less", t1, t2, -1},
		{"Equal", t2, t2, 0},
		{"Greater", t3, t2, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := Time(tc.a, tc.b); result != tc.expectedResult {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}
