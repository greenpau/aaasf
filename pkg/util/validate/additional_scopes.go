package validate

import (
	"github.com/greenpau/go-authcrunch/pkg/errors"
	"regexp"
)

// AdditionalScopes verifies if the provided additional_scopes argument is valid
func AdditionalScopes(additionalScopes string) error {
	compile, _ := regexp.Compile("[^\\w%20$]+")
	match := compile.MatchString(additionalScopes)
	if match == false {
		return errors.ErrInvalidAdditionalScopes
	}
	return nil
}
