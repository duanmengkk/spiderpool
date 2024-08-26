// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CoordinatorConfig Coordinator config
//
// swagger:model CoordinatorConfig
type CoordinatorConfig struct {

	// detect gateway
	DetectGateway bool `json:"detectGateway,omitempty"`

	// detect IP conflict
	DetectIPConflict bool `json:"detectIPConflict,omitempty"`

	// hijack c ID r
	HijackCIDR []string `json:"hijackCIDR"`

	// host rule table
	HostRuleTable int64 `json:"hostRuleTable,omitempty"`

	// mode
	// Required: true
	Mode *string `json:"mode"`

	// overlay pod c ID r
	// Required: true
	OverlayPodCIDR []string `json:"overlayPodCIDR"`

	// pod default route n i c
	PodDefaultRouteNIC string `json:"podDefaultRouteNIC,omitempty"`

	// pod m a c prefix
	PodMACPrefix string `json:"podMACPrefix,omitempty"`

	// pod r p filter
	PodRPFilter int64 `json:"podRPFilter,omitempty"`

	// service c ID r
	// Required: true
	ServiceCIDR []string `json:"serviceCIDR"`

	// tune pod routes
	// Required: true
	TunePodRoutes *bool `json:"tunePodRoutes"`

	// tx queue len
	TxQueueLen int64 `json:"txQueueLen,omitempty"`
}

// Validate validates this coordinator config
func (m *CoordinatorConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverlayPodCIDR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCIDR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunePodRoutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CoordinatorConfig) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorConfig) validateOverlayPodCIDR(formats strfmt.Registry) error {

	if err := validate.Required("overlayPodCIDR", "body", m.OverlayPodCIDR); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorConfig) validateServiceCIDR(formats strfmt.Registry) error {

	if err := validate.Required("serviceCIDR", "body", m.ServiceCIDR); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorConfig) validateTunePodRoutes(formats strfmt.Registry) error {

	if err := validate.Required("tunePodRoutes", "body", m.TunePodRoutes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this coordinator config based on context it is used
func (m *CoordinatorConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoordinatorConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoordinatorConfig) UnmarshalBinary(b []byte) error {
	var res CoordinatorConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
