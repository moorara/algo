package radixsort

import "testing"

func TestQuick3WayString(t *testing.T) {
	tests := []struct {
		a []string
		w int
	}{
		{
			a: []string{"Milad", "Mona", "Milad", "Mona"},
		},
		{
			a: []string{
				"Docker", "Kubernetes", "Prometheus", "Docker",
				"Terraform", "Vault", "Consul",
				"Linkerd", "Istio",
				"Kafka", "NATS",
				"CockroachDB", "ArangoDB",
				"Go", "JavaScript", "TypeScript", "Go",
				"gRPC", "GraphQL", "gRPC",
				"React", "Redux", "Vue", "React", "Redux",
			},
		},
	}

	for _, tc := range tests {
		Quick3WayString(tc.a)

		if !isSorted[string](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}
