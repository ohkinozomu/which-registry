package which_registry

import (
	"errors"
	"regexp"
	"strings"

	regions "github.com/ohkinozomu/gcp-regions"
)

// https://cloud.google.com/artifact-registry/docs/repo-locations#location-mr
func isMultiRegion(l string) bool {
	return l == "asia" || l == "europe" || l == "us"
}

// https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling#tag
func isGoogleArtifactRegistry(d string) (bool, error) {
	match, err := regexp.MatchString(".*-docker.pkg.dev", d)
	if err != nil {
		return false, err
	}

	if match {
		location := strings.Replace(d, "-docker.pkg.dev", "", 1)
		if !regions.IsValid(location) && !isMultiRegion(location) {
			return false, errors.New("Invalid location: " + location)
		}
		return true, nil
	}
	return false, nil
}
