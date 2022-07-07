package radixsort

import "testing"

func TestMSDString(t *testing.T) {
	tests := []struct {
		a []string
	}{
		{
			a: []string{"Milad", "Mona", "Milad", "Mona"},
		},
		{
			a: []string{
				"Docker", "Kubernetes", "Prometheus",
				"Terraform", "Vault", "Consul",
				"Linkerd", "Istio",
				"Kafka", "NATS",
				"CockroachDB", "ArangoDB",
				"Go", "JavaScript", "TypeScript",
				"gRPC", "GraphQL",
				"React", "Redux", "Vue",
			},
		},
	}

	for _, tc := range tests {
		MSDString(tc.a)

		if !isSorted[string](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}

func TestMSDInt(t *testing.T) {
	tests := []struct {
		a []int
	}{
		{[]int{30, -20, 10, -40, 50}},
		{[]int{90, -85, 80, -75, 70, -65, 60, -55, 50, -45, 40, -35, 30, -25, 20, -15, 10}},
	}

	for _, tc := range tests {
		MSDInt(tc.a)

		if !isSorted[int](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}

func TestMSDUint(t *testing.T) {
	tests := []struct {
		a []uint
	}{
		{[]uint{30, 20, 10, 40, 50}},
		{[]uint{90, 85, 80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10}},
	}

	for _, tc := range tests {
		MSDUint(tc.a)

		if !isSorted[uint](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}
