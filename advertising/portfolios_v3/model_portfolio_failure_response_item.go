package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioFailureResponseItem{}

// PortfolioFailureResponseItem struct for PortfolioFailureResponseItem
type PortfolioFailureResponseItem struct {
	// the index of the portfolio in the array from the request body
	Index float32 `json:"index"`
	// a list of validation errors
	Errors []PortfolioMutationError `json:"errors,omitempty"`
}

type _PortfolioFailureResponseItem PortfolioFailureResponseItem

// NewPortfolioFailureResponseItem instantiates a new PortfolioFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioFailureResponseItem(index float32) *PortfolioFailureResponseItem {
	this := PortfolioFailureResponseItem{}
	this.Index = index
	return &this
}

// NewPortfolioFailureResponseItemWithDefaults instantiates a new PortfolioFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioFailureResponseItemWithDefaults() *PortfolioFailureResponseItem {
	this := PortfolioFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *PortfolioFailureResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *PortfolioFailureResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *PortfolioFailureResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PortfolioFailureResponseItem) GetErrors() []PortfolioMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []PortfolioMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioFailureResponseItem) GetErrorsOk() ([]PortfolioMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PortfolioFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []PortfolioMutationError and assigns it to the Errors field.
func (o *PortfolioFailureResponseItem) SetErrors(v []PortfolioMutationError) {
	o.Errors = v
}

func (o PortfolioFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePortfolioFailureResponseItem struct {
	value *PortfolioFailureResponseItem
	isSet bool
}

func (v NullablePortfolioFailureResponseItem) Get() *PortfolioFailureResponseItem {
	return v.value
}

func (v *NullablePortfolioFailureResponseItem) Set(val *PortfolioFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioFailureResponseItem(val *PortfolioFailureResponseItem) *NullablePortfolioFailureResponseItem {
	return &NullablePortfolioFailureResponseItem{value: val, isSet: true}
}

func (v NullablePortfolioFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
