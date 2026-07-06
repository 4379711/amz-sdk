package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemCancellation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemCancellation{}

// ItemCancellation The cancellation information of the order item.
type ItemCancellation struct {
	CancellationRequest *ItemCancellationRequest `json:"cancellationRequest,omitempty"`
}

// NewItemCancellation instantiates a new ItemCancellation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemCancellation() *ItemCancellation {
	this := ItemCancellation{}
	return &this
}

// NewItemCancellationWithDefaults instantiates a new ItemCancellation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemCancellationWithDefaults() *ItemCancellation {
	this := ItemCancellation{}
	return &this
}

// GetCancellationRequest returns the CancellationRequest field value if set, zero value otherwise.
func (o *ItemCancellation) GetCancellationRequest() ItemCancellationRequest {
	if o == nil || IsNil(o.CancellationRequest) {
		var ret ItemCancellationRequest
		return ret
	}
	return *o.CancellationRequest
}

// GetCancellationRequestOk returns a tuple with the CancellationRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCancellation) GetCancellationRequestOk() (*ItemCancellationRequest, bool) {
	if o == nil || IsNil(o.CancellationRequest) {
		return nil, false
	}
	return o.CancellationRequest, true
}

// HasCancellationRequest returns a boolean if a field has been set.
func (o *ItemCancellation) HasCancellationRequest() bool {
	if o != nil && !IsNil(o.CancellationRequest) {
		return true
	}

	return false
}

// SetCancellationRequest gets a reference to the given ItemCancellationRequest and assigns it to the CancellationRequest field.
func (o *ItemCancellation) SetCancellationRequest(v ItemCancellationRequest) {
	o.CancellationRequest = &v
}

func (o ItemCancellation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancellationRequest) {
		toSerialize["cancellationRequest"] = o.CancellationRequest
	}
	return toSerialize, nil
}

type NullableItemCancellation struct {
	value *ItemCancellation
	isSet bool
}

func (v NullableItemCancellation) Get() *ItemCancellation {
	return v.value
}

func (v *NullableItemCancellation) Set(val *ItemCancellation) {
	v.value = val
	v.isSet = true
}

func (v NullableItemCancellation) IsSet() bool {
	return v.isSet
}

func (v *NullableItemCancellation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemCancellation(val *ItemCancellation) *NullableItemCancellation {
	return &NullableItemCancellation{value: val, isSet: true}
}

func (v NullableItemCancellation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemCancellation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
