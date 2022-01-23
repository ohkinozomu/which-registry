package which_registry

import (
	"testing"
)

func TestWhich(t *testing.T) {
	type test struct {
		image    string
		registry Registry
	}

	tests := []test{
		{
			image:    "nginx:alpine",
			registry: DOCKER_HUB,
		},
		{
			image:    "public.ecr.aws/nginx/nginx:1.21",
			registry: ECR_PUBLIC,
		},
		{
			image:    "111111111111.dkr.ecr.us-east-1.amazonaws.com/nginx:1.21",
			registry: ECR_PRIVATE,
		},
	}

	for _, test := range tests {
		r, err := Which(test.image)
		if err != nil {
			t.Fatalf("failed test: %v\n", err)
		}
		if r != test.registry {
			t.Fatalf("failed test: %v\n", r)
		}
	}
}

func TestParseRepo(t *testing.T) {
	r := "docker.io/library/nginx"
	if parseRepo(r) != "docker.io" {
		t.Fatalf("failed test: %v\n", r)
	}
}
