package which_registry

import "regexp"

// https://cloud.google.com/container-registry/docs/overview#registries
func isGoogleContainerRegistry(d string) (bool, error) {
	// TODO: location check
	match, err := regexp.MatchString(".*gcr.io", d)
	if err != nil {
		return false, err
	}
	return match, nil
}
