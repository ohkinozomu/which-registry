package which_registry

import (
	"strings"
)

func isECRPublic(domain string) bool {
	return domain == "public.ecr.aws"
}

// https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_concepts
func isECRPrivate(d string) (bool, error) {
	s := strings.Split(d, ".")
	if len(s) <= 3 {
		return false, nil
	}

	if s[1] == "dkr" && s[2] == "ecr" && s[4] == "amazonaws" && s[5] == "com" {
		return true, nil
	}
	return false, nil
}
