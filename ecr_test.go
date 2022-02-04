package which_registry

import (
	"testing"
)

func TestValidateRegion(t *testing.T) {
	region := "us-east-1"
	err := validateRegion(region)
	if err != nil {
		t.Fatalf("failed test: %v\n", region)
	}

	region = "us-east-1000000"
	err = validateRegion(region)
	if err.Error() != "Invalid Region: "+region {
		t.Fatalf("failed test: %v\n", region)
	}
}
