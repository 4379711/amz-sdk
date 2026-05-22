package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the RejectedOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectedOrder{}

// RejectedOrder A order which we couldn't schedule on your behalf. It contains its id, and information on the error.
type RejectedOrder struct {
	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId string `json:"amazonOrderId"`
	Error         *Error `json:"error,omitempty"`
}

type _RejectedOrder RejectedOrder

// NewRejectedOrder instantiates a new RejectedOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectedOrder(amazonOrderId string) *RejectedOrder {
	this := RejectedOrder{}
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewRejectedOrderWithDefaults instantiates a new RejectedOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectedOrderWithDefaults() *RejectedOrder {
	this := RejectedOrder{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *RejectedOrder) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *RejectedOrder) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *RejectedOrder) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RejectedOrder) GetError() Error {
	if o == nil || IsNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectedOrder) GetErrorOk() (*Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RejectedOrder) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *RejectedOrder) SetError(v Error) {
	o.Error = &v
}

func (o RejectedOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amazonOrderId"] = o.AmazonOrderId
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableRejectedOrder struct {
	value *RejectedOrder
	isSet bool
}

func (v NullableRejectedOrder) Get() *RejectedOrder {
	return v.value
}

func (v *NullableRejectedOrder) Set(val *RejectedOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectedOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectedOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectedOrder(val *RejectedOrder) *NullableRejectedOrder {
	return &NullableRejectedOrder{value: val, isSet: true}
}

func (v NullableRejectedOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRejectedOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
