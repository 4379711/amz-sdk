package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmTransportationOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmTransportationOptionsRequest{}

// ConfirmTransportationOptionsRequest The `confirmTransportationOptions` request.
type ConfirmTransportationOptionsRequest struct {
	// Information needed to confirm one of the available transportation options.
	TransportationSelections []TransportationSelection `json:"transportationSelections"`
}

type _ConfirmTransportationOptionsRequest ConfirmTransportationOptionsRequest

// NewConfirmTransportationOptionsRequest instantiates a new ConfirmTransportationOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmTransportationOptionsRequest(transportationSelections []TransportationSelection) *ConfirmTransportationOptionsRequest {
	this := ConfirmTransportationOptionsRequest{}
	this.TransportationSelections = transportationSelections
	return &this
}

// NewConfirmTransportationOptionsRequestWithDefaults instantiates a new ConfirmTransportationOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmTransportationOptionsRequestWithDefaults() *ConfirmTransportationOptionsRequest {
	this := ConfirmTransportationOptionsRequest{}
	return &this
}

// GetTransportationSelections returns the TransportationSelections field value
func (o *ConfirmTransportationOptionsRequest) GetTransportationSelections() []TransportationSelection {
	if o == nil {
		var ret []TransportationSelection
		return ret
	}

	return o.TransportationSelections
}

// GetTransportationSelectionsOk returns a tuple with the TransportationSelections field value
// and a boolean to check if the value has been set.
func (o *ConfirmTransportationOptionsRequest) GetTransportationSelectionsOk() ([]TransportationSelection, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransportationSelections, true
}

// SetTransportationSelections sets field value
func (o *ConfirmTransportationOptionsRequest) SetTransportationSelections(v []TransportationSelection) {
	o.TransportationSelections = v
}

func (o ConfirmTransportationOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transportationSelections"] = o.TransportationSelections
	return toSerialize, nil
}

type NullableConfirmTransportationOptionsRequest struct {
	value *ConfirmTransportationOptionsRequest
	isSet bool
}

func (v NullableConfirmTransportationOptionsRequest) Get() *ConfirmTransportationOptionsRequest {
	return v.value
}

func (v *NullableConfirmTransportationOptionsRequest) Set(val *ConfirmTransportationOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmTransportationOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmTransportationOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmTransportationOptionsRequest(val *ConfirmTransportationOptionsRequest) *NullableConfirmTransportationOptionsRequest {
	return &NullableConfirmTransportationOptionsRequest{value: val, isSet: true}
}

func (v NullableConfirmTransportationOptionsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmTransportationOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
