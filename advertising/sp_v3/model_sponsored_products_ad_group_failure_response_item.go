package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsAdGroupFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupFailureResponseItem{}

// SponsoredProductsAdGroupFailureResponseItem struct for SponsoredProductsAdGroupFailureResponseItem
type SponsoredProductsAdGroupFailureResponseItem struct {
	// the index of the adGroup in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsAdGroupMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsAdGroupFailureResponseItem SponsoredProductsAdGroupFailureResponseItem

// NewSponsoredProductsAdGroupFailureResponseItem instantiates a new SponsoredProductsAdGroupFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupFailureResponseItem(index int32) *SponsoredProductsAdGroupFailureResponseItem {
	this := SponsoredProductsAdGroupFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsAdGroupFailureResponseItemWithDefaults instantiates a new SponsoredProductsAdGroupFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupFailureResponseItemWithDefaults() *SponsoredProductsAdGroupFailureResponseItem {
	this := SponsoredProductsAdGroupFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsAdGroupFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsAdGroupFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupFailureResponseItem) GetErrors() []SponsoredProductsAdGroupMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsAdGroupMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupFailureResponseItem) GetErrorsOk() ([]SponsoredProductsAdGroupMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsAdGroupMutationError and assigns it to the Errors field.
func (o *SponsoredProductsAdGroupFailureResponseItem) SetErrors(v []SponsoredProductsAdGroupMutationError) {
	o.Errors = v
}

func (o SponsoredProductsAdGroupFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAdGroupFailureResponseItem struct {
	value *SponsoredProductsAdGroupFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsAdGroupFailureResponseItem) Get() *SponsoredProductsAdGroupFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupFailureResponseItem) Set(val *SponsoredProductsAdGroupFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupFailureResponseItem(val *SponsoredProductsAdGroupFailureResponseItem) *NullableSponsoredProductsAdGroupFailureResponseItem {
	return &NullableSponsoredProductsAdGroupFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
