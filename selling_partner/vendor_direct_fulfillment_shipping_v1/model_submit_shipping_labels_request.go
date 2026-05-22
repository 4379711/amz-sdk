package vendor_direct_fulfillment_shipping_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitShippingLabelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitShippingLabelsRequest{}

// SubmitShippingLabelsRequest The request schema for the submitShippingLabelRequest operation.
type SubmitShippingLabelsRequest struct {
	// An array of shipping label requests to be processed.
	ShippingLabelRequests []ShippingLabelRequest `json:"shippingLabelRequests,omitempty"`
}

// NewSubmitShippingLabelsRequest instantiates a new SubmitShippingLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitShippingLabelsRequest() *SubmitShippingLabelsRequest {
	this := SubmitShippingLabelsRequest{}
	return &this
}

// NewSubmitShippingLabelsRequestWithDefaults instantiates a new SubmitShippingLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitShippingLabelsRequestWithDefaults() *SubmitShippingLabelsRequest {
	this := SubmitShippingLabelsRequest{}
	return &this
}

// GetShippingLabelRequests returns the ShippingLabelRequests field value if set, zero value otherwise.
func (o *SubmitShippingLabelsRequest) GetShippingLabelRequests() []ShippingLabelRequest {
	if o == nil || IsNil(o.ShippingLabelRequests) {
		var ret []ShippingLabelRequest
		return ret
	}
	return o.ShippingLabelRequests
}

// GetShippingLabelRequestsOk returns a tuple with the ShippingLabelRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitShippingLabelsRequest) GetShippingLabelRequestsOk() ([]ShippingLabelRequest, bool) {
	if o == nil || IsNil(o.ShippingLabelRequests) {
		return nil, false
	}
	return o.ShippingLabelRequests, true
}

// HasShippingLabelRequests returns a boolean if a field has been set.
func (o *SubmitShippingLabelsRequest) HasShippingLabelRequests() bool {
	if o != nil && !IsNil(o.ShippingLabelRequests) {
		return true
	}

	return false
}

// SetShippingLabelRequests gets a reference to the given []ShippingLabelRequest and assigns it to the ShippingLabelRequests field.
func (o *SubmitShippingLabelsRequest) SetShippingLabelRequests(v []ShippingLabelRequest) {
	o.ShippingLabelRequests = v
}

func (o SubmitShippingLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShippingLabelRequests) {
		toSerialize["shippingLabelRequests"] = o.ShippingLabelRequests
	}
	return toSerialize, nil
}

type NullableSubmitShippingLabelsRequest struct {
	value *SubmitShippingLabelsRequest
	isSet bool
}

func (v NullableSubmitShippingLabelsRequest) Get() *SubmitShippingLabelsRequest {
	return v.value
}

func (v *NullableSubmitShippingLabelsRequest) Set(val *SubmitShippingLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitShippingLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitShippingLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitShippingLabelsRequest(val *SubmitShippingLabelsRequest) *NullableSubmitShippingLabelsRequest {
	return &NullableSubmitShippingLabelsRequest{value: val, isSet: true}
}

func (v NullableSubmitShippingLabelsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitShippingLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
