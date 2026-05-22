package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdFailureResponseItem{}

// AdFailureResponseItem struct for AdFailureResponseItem
type AdFailureResponseItem struct {
	// the index of the ad in the array from the request body.
	Index float32 `json:"index"`
	// A list of validation errors.
	Errors []AdMutationError `json:"errors,omitempty"`
}

type _AdFailureResponseItem AdFailureResponseItem

// NewAdFailureResponseItem instantiates a new AdFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdFailureResponseItem(index float32) *AdFailureResponseItem {
	this := AdFailureResponseItem{}
	this.Index = index
	return &this
}

// NewAdFailureResponseItemWithDefaults instantiates a new AdFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdFailureResponseItemWithDefaults() *AdFailureResponseItem {
	this := AdFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *AdFailureResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AdFailureResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AdFailureResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AdFailureResponseItem) GetErrors() []AdMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []AdMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdFailureResponseItem) GetErrorsOk() ([]AdMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AdFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AdMutationError and assigns it to the Errors field.
func (o *AdFailureResponseItem) SetErrors(v []AdMutationError) {
	o.Errors = v
}

func (o AdFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableAdFailureResponseItem struct {
	value *AdFailureResponseItem
	isSet bool
}

func (v NullableAdFailureResponseItem) Get() *AdFailureResponseItem {
	return v.value
}

func (v *NullableAdFailureResponseItem) Set(val *AdFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAdFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAdFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdFailureResponseItem(val *AdFailureResponseItem) *NullableAdFailureResponseItem {
	return &NullableAdFailureResponseItem{value: val, isSet: true}
}

func (v NullableAdFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
