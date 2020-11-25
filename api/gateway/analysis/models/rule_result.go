// Code generated by go-swagger; DO NOT EDIT.

package models

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuleResult rule result
//
// swagger:model ruleResult
type RuleResult struct {

	// alert context error
	AlertContextError string `json:"alertContextError,omitempty"`

	// alert context output
	AlertContextOutput string `json:"alertContextOutput,omitempty"`

	// dedup error
	DedupError string `json:"dedupError,omitempty"`

	// dedup output
	DedupOutput string `json:"dedupOutput,omitempty"`

	// True if the result includes an error.
	Errored bool `json:"errored,omitempty"`

	// An error produced before running any of the rule functions, like import or syntax error.
	GenericError string `json:"genericError,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// True if the test returned the expected result. Always false if the result includes an error.
	Passed bool `json:"passed,omitempty"`

	// rule error
	RuleError string `json:"ruleError,omitempty"`

	// rule Id
	RuleID string `json:"ruleId,omitempty"`

	// rule output
	RuleOutput bool `json:"ruleOutput,omitempty"`

	// test name
	TestName string `json:"testName,omitempty"`

	// title error
	TitleError string `json:"titleError,omitempty"`

	// title output
	TitleOutput string `json:"titleOutput,omitempty"`
}

// Validate validates this rule result
func (m *RuleResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleResult) UnmarshalBinary(b []byte) error {
	var res RuleResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
