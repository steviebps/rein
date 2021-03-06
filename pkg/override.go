package rein

import (
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/mod/semver"
)

// Override is a Toggle value to be consumed by and restricted to a semantic version range
type Override struct {
	MinimumVersion string      `json:"minimumVersion"`
	MaximumVersion string      `json:"maximumVersion"`
	Value          interface{} `json:"value"`
}

// UnmarshalJSON Custom UnmarshalJSON method for validating Override
func (o *Override) UnmarshalJSON(b []byte) error {

	var alias overrideAlias

	err := json.Unmarshal(b, &alias)
	if err != nil {
		return err
	}

	*o = alias.toOverride()

	if o.Value == nil {
		return errors.New("Override value cannot be empty/nil")
	}

	if isValidMin := semver.IsValid(o.MinimumVersion); !isValidMin {
		return fmt.Errorf("%q is not a valid semantic version", o.MinimumVersion)
	}

	if isValidMax := semver.IsValid(o.MaximumVersion); !isValidMax {
		return fmt.Errorf("%q is not a valid semantic version", o.MaximumVersion)
	}

	// if minimum version is greater than maximum version
	if semver.Compare(o.MinimumVersion, o.MaximumVersion) == 1 {
		return fmt.Errorf("An override with the minimum version of %v is greater than its maximum version (%v)", o.MinimumVersion, o.MaximumVersion)
	}

	return nil
}

type overrideAlias Override

func (o overrideAlias) toOverride() Override {
	return Override{
		o.MinimumVersion,
		o.MaximumVersion,
		o.Value,
	}
}
