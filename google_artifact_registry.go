package which_registry

import "regexp"

// https://cloud.google.com/artifact-registry/docs/docker/pushing-and-pulling#tag
func isGoogleArtifactRegistry(d string) (bool, error) {
	// TODO: location check
	match, err := regexp.MatchString(".*-docker.pkg.dev", d)
	if err != nil {
		return false, err
	}
	return match, nil
}
