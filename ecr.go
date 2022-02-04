package which_registry

import (
	"errors"
	"strings"

	regions "github.com/jsonmaur/aws-regions/v2"
)

func isECRPublic(domain string) bool {
	return domain == "public.ecr.aws"
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

	if s[1] == "dkr" && s[2] == "ecr" && s[4] == "amazonaws" && s[5] == "com" {
		err := validateRegion(s[3])
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
