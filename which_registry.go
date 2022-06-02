package which_registry

import (
	"strings"
)

func parseRepo(r string) string {
	return strings.Split(r, "/")[0]
}

func Which(image string) (Registry, error) {
	var r Registry
	repo, _, _, err := ParseImageName(image)
	if err != nil {
		return r, err
	}

	domain := parseRepo(repo)

	if isDockerHub(domain) {
		return DOCKER_HUB, nil
	} else if isECRPublic(domain) {
		return ECR_PUBLIC, nil
	} else if isQuayIO(domain) {
		return QUAY_IO, nil
	}

	isECRp, err := isECRPrivate(domain)
	if err != nil {
		return r, err
	}
	if isECRp {
		return ECR_PRIVATE, nil
	}

	isGAR, err := isGoogleArtifactRegistry(domain)
	if err != nil {
		return r, err
	}
	if isGAR {
		return GOOGLE_ARTIFACT_REGISTRY, nil
	}

	isGCR, err := isGoogleContainerRegistry(domain)
	if err != nil {
		return r, err
	}
	if isGCR {
		return GOOGLE_CONTAINER_REGISTRY, nil
	}

	if isGitHubContainerRegistry(domain) {
		return GITHUB_CONTAINER_REGISTRY, nil
	}

	return r, nil
}
