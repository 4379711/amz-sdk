package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAdditionalSellerInputsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdditionalSellerInputsResult{}

// GetAdditionalSellerInputsResult The payload for the `getAdditionalSellerInputs` operation.
type GetAdditionalSellerInputsResult struct {
	// A list of additional inputs.
	ShipmentLevelFields []AdditionalInputs `json:"ShipmentLevelFields,omitempty"`
	// A list of item level fields.
	ItemLevelFieldsList []ItemLevelFields `json:"ItemLevelFieldsList,omitempty"`
}

// NewGetAdditionalSellerInputsResult instantiates a new GetAdditionalSellerInputsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdditionalSellerInputsResult() *GetAdditionalSellerInputsResult {
	this := GetAdditionalSellerInputsResult{}
	return &this
}

// NewGetAdditionalSellerInputsResultWithDefaults instantiates a new GetAdditionalSellerInputsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdditionalSellerInputsResultWithDefaults() *GetAdditionalSellerInputsResult {
	this := GetAdditionalSellerInputsResult{}
	return &this
}

// GetShipmentLevelFields returns the ShipmentLevelFields field value if set, zero value otherwise.
func (o *GetAdditionalSellerInputsResult) GetShipmentLevelFields() []AdditionalInputs {
	if o == nil || IsNil(o.ShipmentLevelFields) {
		var ret []AdditionalInputs
		return ret
	}
	return o.ShipmentLevelFields
}

// GetShipmentLevelFieldsOk returns a tuple with the ShipmentLevelFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsResult) GetShipmentLevelFieldsOk() ([]AdditionalInputs, bool) {
	if o == nil || IsNil(o.ShipmentLevelFields) {
		return nil, false
	}
	return o.ShipmentLevelFields, true
}

// HasShipmentLevelFields returns a boolean if a field has been set.
func (o *GetAdditionalSellerInputsResult) HasShipmentLevelFields() bool {
	if o != nil && !IsNil(o.ShipmentLevelFields) {
		return true
	}

	return false
}

// SetShipmentLevelFields gets a reference to the given []AdditionalInputs and assigns it to the ShipmentLevelFields field.
func (o *GetAdditionalSellerInputsResult) SetShipmentLevelFields(v []AdditionalInputs) {
	o.ShipmentLevelFields = v
}

// GetItemLevelFieldsList returns the ItemLevelFieldsList field value if set, zero value otherwise.
func (o *GetAdditionalSellerInputsResult) GetItemLevelFieldsList() []ItemLevelFields {
	if o == nil || IsNil(o.ItemLevelFieldsList) {
		var ret []ItemLevelFields
		return ret
	}
	return o.ItemLevelFieldsList
}

// GetItemLevelFieldsListOk returns a tuple with the ItemLevelFieldsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdditionalSellerInputsResult) GetItemLevelFieldsListOk() ([]ItemLevelFields, bool) {
	if o == nil || IsNil(o.ItemLevelFieldsList) {
		return nil, false
	}
	return o.ItemLevelFieldsList, true
}

// HasItemLevelFieldsList returns a boolean if a field has been set.
func (o *GetAdditionalSellerInputsResult) HasItemLevelFieldsList() bool {
	if o != nil && !IsNil(o.ItemLevelFieldsList) {
		return true
	}

	return false
}

// SetItemLevelFieldsList gets a reference to the given []ItemLevelFields and assigns it to the ItemLevelFieldsList field.
func (o *GetAdditionalSellerInputsResult) SetItemLevelFieldsList(v []ItemLevelFields) {
	o.ItemLevelFieldsList = v
}

func (o GetAdditionalSellerInputsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentLevelFields) {
		toSerialize["ShipmentLevelFields"] = o.ShipmentLevelFields
	}
	if !IsNil(o.ItemLevelFieldsList) {
		toSerialize["ItemLevelFieldsList"] = o.ItemLevelFieldsList
	}
	return toSerialize, nil
}

type NullableGetAdditionalSellerInputsResult struct {
	value *GetAdditionalSellerInputsResult
	isSet bool
}

func (v NullableGetAdditionalSellerInputsResult) Get() *GetAdditionalSellerInputsResult {
	return v.value
}

func (v *NullableGetAdditionalSellerInputsResult) Set(val *GetAdditionalSellerInputsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdditionalSellerInputsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdditionalSellerInputsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdditionalSellerInputsResult(val *GetAdditionalSellerInputsResult) *NullableGetAdditionalSellerInputsResult {
	return &NullableGetAdditionalSellerInputsResult{value: val, isSet: true}
}

func (v NullableGetAdditionalSellerInputsResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAdditionalSellerInputsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
