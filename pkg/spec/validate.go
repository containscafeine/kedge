/*
Copyright 2017 The Kedge Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spec

import (
	"fmt"

	"github.com/pkg/errors"
)

func validateVolumeClaims(vcs []VolumeClaim) error {
	// find the duplicate volume claim names, if found any then error out
	vcmap := make(map[string]interface{})
	for _, vc := range vcs {
		if _, ok := vcmap[vc.Name]; !ok {
			// value here does not matter
			vcmap[vc.Name] = nil
		} else {
			return fmt.Errorf("duplicate entry of volume claim %q", vc.Name)
		}
	}
	return nil
}

func validateKubernetesVersion(kVersion string) error {
	switch kVersion {
	case latestKubernetesVersion:
		return nil
	default:
		return fmt.Errorf("unsupported Kubernetes version: %v", kVersion)
	}
}

func ValidateApp(app *App) error {

	// validate kubernetesVersion
	if err := validateKubernetesVersion(app.KubernetesVersion); err != nil {
		return errors.Wrap(err, "error validating kubernetesVersion")
	}

	// validate volumeclaims
	if err := validateVolumeClaims(app.VolumeClaims); err != nil {
		return errors.Wrap(err, "error validating volume claims")
	}

	return nil
}
