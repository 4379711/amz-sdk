package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the UnlinkCarrierAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnlinkCarrierAccountResponse{}

// UnlinkCarrierAccountResponse The Response  for the UnlinkCarrierAccountResponse operation.
type UnlinkCarrierAccountResponse struct {
	// Is Carrier unlinked from Merchant
	IsUnlinked *bool `json:"isUnlinked,omitempty"`
}

// NewUnlinkCarrierAccountResponse instantiates a new UnlinkCarrierAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlinkCarrierAccountResponse() *UnlinkCarrierAccountResponse {
	this := UnlinkCarrierAccountResponse{}
	return &this
}

// NewUnlinkCarrierAccountResponseWithDefaults instantiates a new UnlinkCarrierAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlinkCarrierAccountResponseWithDefaults() *UnlinkCarrierAccountResponse {
	this := UnlinkCarrierAccountResponse{}
	return &this
}

// GetIsUnlinked returns the IsUnlinked field value if set, zero value otherwise.
func (o *UnlinkCarrierAccountResponse) GetIsUnlinked() bool {
	if o == nil || IsNil(o.IsUnlinked) {
		var ret bool
		return ret
	}
	return *o.IsUnlinked
}

// GetIsUnlinkedOk returns a tuple with the IsUnlinked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkCarrierAccountResponse) GetIsUnlinkedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnlinked) {
		return nil, false
	}
	return o.IsUnlinked, true
}

// HasIsUnlinked returns a boolean if a field has been set.
func (o *UnlinkCarrierAccountResponse) HasIsUnlinked() bool {
	if o != nil && !IsNil(o.IsUnlinked) {
		return true
	}

	return false
}

// SetIsUnlinked gets a reference to the given bool and assigns it to the IsUnlinked field.
func (o *UnlinkCarrierAccountResponse) SetIsUnlinked(v bool) {
	o.IsUnlinked = &v
}

func (o UnlinkCarrierAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsUnlinked) {
		toSerialize["isUnlinked"] = o.IsUnlinked
	}
	return toSerialize, nil
}

type NullableUnlinkCarrierAccountResponse struct {
	value *UnlinkCarrierAccountResponse
	isSet bool
}

func (v NullableUnlinkCarrierAccountResponse) Get() *UnlinkCarrierAccountResponse {
	return v.value
}

func (v *NullableUnlinkCarrierAccountResponse) Set(val *UnlinkCarrierAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlinkCarrierAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlinkCarrierAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlinkCarrierAccountResponse(val *UnlinkCarrierAccountResponse) *NullableUnlinkCarrierAccountResponse {
	return &NullableUnlinkCarrierAccountResponse{value: val, isSet: true}
}

func (v NullableUnlinkCarrierAccountResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnlinkCarrierAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
