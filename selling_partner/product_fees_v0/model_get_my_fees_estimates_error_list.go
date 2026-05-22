package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMyFeesEstimatesErrorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMyFeesEstimatesErrorList{}

// GetMyFeesEstimatesErrorList A list of error responses returned when a request is unsuccessful.
type GetMyFeesEstimatesErrorList struct {
	Errors []Error `json:"errors"`
}

type _GetMyFeesEstimatesErrorList GetMyFeesEstimatesErrorList

// NewGetMyFeesEstimatesErrorList instantiates a new GetMyFeesEstimatesErrorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMyFeesEstimatesErrorList(errors []Error) *GetMyFeesEstimatesErrorList {
	this := GetMyFeesEstimatesErrorList{}
	this.Errors = errors
	return &this
}

// NewGetMyFeesEstimatesErrorListWithDefaults instantiates a new GetMyFeesEstimatesErrorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMyFeesEstimatesErrorListWithDefaults() *GetMyFeesEstimatesErrorList {
	this := GetMyFeesEstimatesErrorList{}
	return &this
}

// GetErrors returns the Errors field value
func (o *GetMyFeesEstimatesErrorList) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *GetMyFeesEstimatesErrorList) GetErrorsOk() ([]Error, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *GetMyFeesEstimatesErrorList) SetErrors(v []Error) {
	o.Errors = v
}

func (o GetMyFeesEstimatesErrorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableGetMyFeesEstimatesErrorList struct {
	value *GetMyFeesEstimatesErrorList
	isSet bool
}

func (v NullableGetMyFeesEstimatesErrorList) Get() *GetMyFeesEstimatesErrorList {
	return v.value
}

func (v *NullableGetMyFeesEstimatesErrorList) Set(val *GetMyFeesEstimatesErrorList) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMyFeesEstimatesErrorList) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMyFeesEstimatesErrorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMyFeesEstimatesErrorList(val *GetMyFeesEstimatesErrorList) *NullableGetMyFeesEstimatesErrorList {
	return &NullableGetMyFeesEstimatesErrorList{value: val, isSet: true}
}

func (v NullableGetMyFeesEstimatesErrorList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMyFeesEstimatesErrorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
