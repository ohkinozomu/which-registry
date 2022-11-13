package which_registry

import (
	"regexp"
)

// https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling#tag
func isGoogleArtifactRegistry(d string) (bool, error) {
	match, err := regexp.MatchString(".*-docker.pkg.dev", d)
	if err != nil {
		return false, err
	}

	if match {
		return true, nil
	}
	return false, nil
}
