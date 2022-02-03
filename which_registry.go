package which_registry

import (
	"regexp"
	"strings"
)

func parseRepo(r string) string {
	return strings.Split(r, "/")[0]
}

// https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_concepts
func isECRPrivate(d string) bool {
	s := strings.Split(d, ".")
	// TODO: region check?
	if s[1] == "dkr" && s[2] == "ecr" && s[4] == "amazonaws" && s[5] == "com" {
		return true
	}
	return false
}

// https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling#tag
func isArtifactRegistry(d string) (bool, error) {

	// TODO: location check
	match, err := regexp.MatchString(".*-docker.pkg.dev", d)
	if err != nil {
		return false, err
	}
	return match, nil
}

func Which(image string) (Registry, error) {
	var r Registry
	repo, _, _, err := ParseImageName(image)
	if err != nil {
		return r, err
	}

	domain := parseRepo(repo)

	if domain == "docker.io" {
		return DOCKER_HUB, nil
	} else if domain == "public.ecr.aws" {
		return ECR_PUBLIC, nil
	}

	if isECRPrivate(domain) {
		return ECR_PRIVATE, nil
	}

	isar, err := isArtifactRegistry(domain)
	if err != nil {
		return r, err
	}
	if isar {
		return ARTIFACT_REGISTRY, nil
	}

	return r, nil
}
