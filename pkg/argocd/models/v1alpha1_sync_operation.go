// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1SyncOperation SyncOperation contains sync operation details.
//
// swagger:model v1alpha1SyncOperation
type V1alpha1SyncOperation struct {

	// DryRun will perform a `kubectl apply --dry-run` without actually performing the sync
	DryRun bool `json:"dryRun,omitempty"`

	// Manifests is an optional field that overrides sync source with a local directory for development
	Manifests []string `json:"manifests"`

	// Prune deletes resources that are no longer tracked in git
	Prune bool `json:"prune,omitempty"`

	// Resources describes which resources to sync
	Resources []*V1alpha1SyncOperationResource `json:"resources"`

	// Revision is the revision in which to sync the application to.
	// If omitted, will use the revision specified in app spec.
	Revision string `json:"revision,omitempty"`

	// source
	Source *V1alpha1ApplicationSource `json:"source,omitempty"`

	// SyncOptions provide per-sync sync-options, e.g. Validate=false
	SyncOptions []string `json:"syncOptions"`

	// sync strategy
	SyncStrategy *V1alpha1SyncStrategy `json:"syncStrategy,omitempty"`
}

// Validate validates this v1alpha1 sync operation
func (m *V1alpha1SyncOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1SyncOperation) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1SyncOperation) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1SyncOperation) validateSyncStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncStrategy) { // not required
		return nil
	}

	if m.SyncStrategy != nil {
		if err := m.SyncStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncStrategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1SyncOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1SyncOperation) UnmarshalBinary(b []byte) error {
	var res V1alpha1SyncOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}