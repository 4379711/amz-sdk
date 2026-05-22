package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateInboundPlanNameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInboundPlanNameRequest{}

// UpdateInboundPlanNameRequest The `updateInboundPlanName` request.
type UpdateInboundPlanNameRequest struct {
	// A human-readable name to update the inbound plan name to.
	Name string `json:"name"`
}

type _UpdateInboundPlanNameRequest UpdateInboundPlanNameRequest

// NewUpdateInboundPlanNameRequest instantiates a new UpdateInboundPlanNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInboundPlanNameRequest(name string) *UpdateInboundPlanNameRequest {
	this := UpdateInboundPlanNameRequest{}
	this.Name = name
	return &this
}

// NewUpdateInboundPlanNameRequestWithDefaults instantiates a new UpdateInboundPlanNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInboundPlanNameRequestWithDefaults() *UpdateInboundPlanNameRequest {
	this := UpdateInboundPlanNameRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateInboundPlanNameRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateInboundPlanNameRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateInboundPlanNameRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateInboundPlanNameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateInboundPlanNameRequest struct {
	value *UpdateInboundPlanNameRequest
	isSet bool
}

func (v NullableUpdateInboundPlanNameRequest) Get() *UpdateInboundPlanNameRequest {
	return v.value
}

func (v *NullableUpdateInboundPlanNameRequest) Set(val *UpdateInboundPlanNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInboundPlanNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInboundPlanNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInboundPlanNameRequest(val *UpdateInboundPlanNameRequest) *NullableUpdateInboundPlanNameRequest {
	return &NullableUpdateInboundPlanNameRequest{value: val, isSet: true}
}

func (v NullableUpdateInboundPlanNameRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateInboundPlanNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
