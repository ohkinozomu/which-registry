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

	isecrpm, err := isECRPrivate(domain)
	if err != nil {
		return r, err
	}
	if isecrpm {
		return ECR_PRIVATE, nil
	}

	isgar, err := isGoogleArtifactRegistry(domain)
	if err != nil {
		return r, err
	}
	if isgar {
		return GOOGLE_ARTIFACT_REGISTRY, nil
	}

	isgcr, err := isGoogleContainerRegistry(domain)
	if err != nil {
		return r, err
	}
	if isgcr {
		return GOOGLE_CONTAINER_REGISTRY, nil
	}

	return r, nil
}
