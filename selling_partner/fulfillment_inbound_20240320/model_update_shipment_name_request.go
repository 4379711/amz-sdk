package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateShipmentNameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentNameRequest{}

// UpdateShipmentNameRequest The `updateShipmentName` request.
type UpdateShipmentNameRequest struct {
	// A human-readable name to update the shipment name to.
	Name string `json:"name"`
}

type _UpdateShipmentNameRequest UpdateShipmentNameRequest

// NewUpdateShipmentNameRequest instantiates a new UpdateShipmentNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentNameRequest(name string) *UpdateShipmentNameRequest {
	this := UpdateShipmentNameRequest{}
	this.Name = name
	return &this
}

// NewUpdateShipmentNameRequestWithDefaults instantiates a new UpdateShipmentNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentNameRequestWithDefaults() *UpdateShipmentNameRequest {
	this := UpdateShipmentNameRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateShipmentNameRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentNameRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateShipmentNameRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateShipmentNameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateShipmentNameRequest struct {
	value *UpdateShipmentNameRequest
	isSet bool
}

func (v NullableUpdateShipmentNameRequest) Get() *UpdateShipmentNameRequest {
	return v.value
}

func (v *NullableUpdateShipmentNameRequest) Set(val *UpdateShipmentNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentNameRequest(val *UpdateShipmentNameRequest) *NullableUpdateShipmentNameRequest {
	return &NullableUpdateShipmentNameRequest{value: val, isSet: true}
}

func (v NullableUpdateShipmentNameRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateShipmentNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
