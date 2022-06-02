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
		{
			image:    "us-east1-docker.pkg.dev/my-project/my-repo/test-image",
			registry: GOOGLE_ARTIFACT_REGISTRY,
		},
		{
			image:    "us-docker.pkg.dev/my-project/my-repo/test-image",
			registry: GOOGLE_ARTIFACT_REGISTRY,
		},
		{
			image:    "gcr.io/example.com/my-project/image-name",
			registry: GOOGLE_CONTAINER_REGISTRY,
		},
		{
			image:    "us.gcr.io/builds/product1/dev/product1-app:beta-2.0",
			registry: GOOGLE_CONTAINER_REGISTRY,
		},
		{
			image:    "docker.elastic.co/elasticsearch/elasticsearch",
			registry: UNKNOWN,
		},
		{
			image:    "quay.io/bitnami/nginx",
			registry: QUAY_IO,
		},
		{
			image: "ghcr.io/parca-dev/parca:v0.11.1",
			registry: GITHUB_CONTAINER_REGISTRY,
		},
	}

	for _, test := range tests {
		r, err := Which(test.image)
		if err != nil {
			t.Fatalf("failed test: %v\n", err)
		}
		if r != test.registry {
			t.Fatalf("failed test: %v %v\n", test.image, r)
		}
	}
}

func TestParseRepo(t *testing.T) {
	r := "docker.io/library/nginx"
	if parseRepo(r) != "docker.io" {
		t.Fatalf("failed test: %v\n", r)
	}
}
