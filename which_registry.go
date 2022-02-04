package which_registry

import (
	"errors"
	"regexp"
	"strings"

	regions "github.com/jsonmaur/aws-regions/v2"
)

func parseRepo(r string) string {
	return strings.Split(r, "/")[0]
}

func validateRegion(region string) error {
	_, err := regions.LookupByCode(region)
	if err != nil {
		return errors.New("Invalid Region: " + region)
	}
	return nil
}

// https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_concepts
func isECRPrivate(d string) (bool, error) {
	s := strings.Split(d, ".")
	if len(s) <= 3 {
		return false, nil
	}

	err := validateRegion(s[3])
	if err != nil {
		return false, err
	}

	if s[1] == "dkr" && s[2] == "ecr" && s[4] == "amazonaws" && s[5] == "com" {
		return true, nil
	}
	return false, nil
}

// https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling#tag
func isGoogleArtifactRegistry(d string) (bool, error) {
	// TODO: location check
	match, err := regexp.MatchString(".*-docker.pkg.dev", d)
	if err != nil {
		return false, err
	}
	return match, nil
}

// https://cloud.google.com/container-registry/docs/overview#registries
func isGoogleContainerRegistry(d string) (bool, error) {
	// TODO: location check
	match, err := regexp.MatchString(".*gcr.io", d)
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
