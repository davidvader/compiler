// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package template

import "github.com/go-vela/types/yaml"

// Engine represents the interface for Vela integrating
// with the different supported template engines.
type Engine interface {
	// Render defines a function that combines
	// the template with the step.
	Render(template string, step *yaml.Step) (yaml.StepSlice, error)
}
