// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StatusResponse Health and status information of daemon
// swagger:model StatusResponse

type StatusResponse struct {

	// Status of Cilium daemon
	Cilium *Status `json:"cilium,omitempty"`

	// Status of cluster
	Cluster *ClusterStatus `json:"cluster,omitempty"`

	// Status of local container runtime
	ContainerRuntime *Status `json:"container-runtime,omitempty"`

	// Status of all endpoint controllers
	Controllers ControllerStatuses `json:"controllers"`

	// Status of IP address management
	IPAM *IPAMStatus `json:"ipam,omitempty"`

	// Status of Kubernetes integration
	Kubernetes *K8sStatus `json:"kubernetes,omitempty"`

	// Status of key/value datastore
	Kvstore *Status `json:"kvstore,omitempty"`

	// Status of the node monitor
	NodeMonitor *MonitorStatus `json:"nodeMonitor,omitempty"`

	// Status of proxy
	Proxy *ProxyStatus `json:"proxy,omitempty"`

	// List of stale information in the status
	Stale map[string]strfmt.DateTime `json:"stale,omitempty"`
}

/* polymorph StatusResponse cilium false */

/* polymorph StatusResponse cluster false */

/* polymorph StatusResponse container-runtime false */

/* polymorph StatusResponse controllers false */

/* polymorph StatusResponse ipam false */

/* polymorph StatusResponse kubernetes false */

/* polymorph StatusResponse kvstore false */

/* polymorph StatusResponse nodeMonitor false */

/* polymorph StatusResponse proxy false */

/* polymorph StatusResponse stale false */

// Validate validates this status response
func (m *StatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCilium(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerRuntime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAM(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKvstore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodeMonitor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusResponse) validateCilium(formats strfmt.Registry) error {

	if swag.IsZero(m.Cilium) { // not required
		return nil
	}

	if m.Cilium != nil {

		if err := m.Cilium.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cilium")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {

		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateContainerRuntime(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerRuntime) { // not required
		return nil
	}

	if m.ContainerRuntime != nil {

		if err := m.ContainerRuntime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container-runtime")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIPAM(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAM) { // not required
		return nil
	}

	if m.IPAM != nil {

		if err := m.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipam")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKubernetes(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if m.Kubernetes != nil {

		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKvstore(formats strfmt.Registry) error {

	if swag.IsZero(m.Kvstore) { // not required
		return nil
	}

	if m.Kvstore != nil {

		if err := m.Kvstore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstore")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateNodeMonitor(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeMonitor) { // not required
		return nil
	}

	if m.NodeMonitor != nil {

		if err := m.NodeMonitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeMonitor")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {

		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusResponse) UnmarshalBinary(b []byte) error {
	var res StatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
