package which_registry

import (
	"errors"
	"regexp"
	"strings"
)

// https://cloud.google.com/container-registry/docs/overview#registries
func isGoogleContainerRegistry(d string) (bool, error) {
	match, err := regexp.MatchString(".*gcr.io", d)
	if err != nil {
		return false, err
	}

	if match {
		location := strings.Replace(d, "gcr.io", "", 1)
		if location == "" || location == "us." || location == "eu." || location == "asia." {
			return true, nil
		}
		return false, errors.New("Invalid location: " + location)
	}
	return match, nil
}
