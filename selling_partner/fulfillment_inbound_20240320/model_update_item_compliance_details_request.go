package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateItemComplianceDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateItemComplianceDetailsRequest{}

// UpdateItemComplianceDetailsRequest The `updateItemComplianceDetails` request.
type UpdateItemComplianceDetailsRequest struct {
	// The merchant SKU, a merchant-supplied identifier for a specific SKU.
	Msku       string     `json:"msku"`
	TaxDetails TaxDetails `json:"taxDetails"`
}

type _UpdateItemComplianceDetailsRequest UpdateItemComplianceDetailsRequest

// NewUpdateItemComplianceDetailsRequest instantiates a new UpdateItemComplianceDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateItemComplianceDetailsRequest(msku string, taxDetails TaxDetails) *UpdateItemComplianceDetailsRequest {
	this := UpdateItemComplianceDetailsRequest{}
	this.Msku = msku
	this.TaxDetails = taxDetails
	return &this
}

// NewUpdateItemComplianceDetailsRequestWithDefaults instantiates a new UpdateItemComplianceDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateItemComplianceDetailsRequestWithDefaults() *UpdateItemComplianceDetailsRequest {
	this := UpdateItemComplianceDetailsRequest{}
	return &this
}

// GetMsku returns the Msku field value
func (o *UpdateItemComplianceDetailsRequest) GetMsku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msku
}

// GetMskuOk returns a tuple with the Msku field value
// and a boolean to check if the value has been set.
func (o *UpdateItemComplianceDetailsRequest) GetMskuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msku, true
}

// SetMsku sets field value
func (o *UpdateItemComplianceDetailsRequest) SetMsku(v string) {
	o.Msku = v
}

// GetTaxDetails returns the TaxDetails field value
func (o *UpdateItemComplianceDetailsRequest) GetTaxDetails() TaxDetails {
	if o == nil {
		var ret TaxDetails
		return ret
	}

	return o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value
// and a boolean to check if the value has been set.
func (o *UpdateItemComplianceDetailsRequest) GetTaxDetailsOk() (*TaxDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxDetails, true
}

// SetTaxDetails sets field value
func (o *UpdateItemComplianceDetailsRequest) SetTaxDetails(v TaxDetails) {
	o.TaxDetails = v
}

func (o UpdateItemComplianceDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msku"] = o.Msku
	toSerialize["taxDetails"] = o.TaxDetails
	return toSerialize, nil
}

type NullableUpdateItemComplianceDetailsRequest struct {
	value *UpdateItemComplianceDetailsRequest
	isSet bool
}

func (v NullableUpdateItemComplianceDetailsRequest) Get() *UpdateItemComplianceDetailsRequest {
	return v.value
}

func (v *NullableUpdateItemComplianceDetailsRequest) Set(val *UpdateItemComplianceDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateItemComplianceDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateItemComplianceDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateItemComplianceDetailsRequest(val *UpdateItemComplianceDetailsRequest) *NullableUpdateItemComplianceDetailsRequest {
	return &NullableUpdateItemComplianceDetailsRequest{value: val, isSet: true}
}

func (v NullableUpdateItemComplianceDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateItemComplianceDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
