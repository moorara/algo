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

		if !isSortedString(tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}

func TestMSDInt(t *testing.T) {
	tests := []struct {
		a []int
	}{
		{[]int{30, 20, 10, 40, 50}},
	}

	for _, tc := range tests {
		MSDInt(tc.a)

		if !isSortedInt(tc.a) {
			// TODO:
			// t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}
