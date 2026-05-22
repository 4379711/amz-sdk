package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdatePortfoliosRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePortfoliosRequestContent{}

// UpdatePortfoliosRequestContent struct for UpdatePortfoliosRequestContent
type UpdatePortfoliosRequestContent struct {
	// An array of portfolio with updated values.
	Portfolios []UpdatePortfolio `json:"portfolios"`
}

type _UpdatePortfoliosRequestContent UpdatePortfoliosRequestContent

// NewUpdatePortfoliosRequestContent instantiates a new UpdatePortfoliosRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePortfoliosRequestContent(portfolios []UpdatePortfolio) *UpdatePortfoliosRequestContent {
	this := UpdatePortfoliosRequestContent{}
	this.Portfolios = portfolios
	return &this
}

// NewUpdatePortfoliosRequestContentWithDefaults instantiates a new UpdatePortfoliosRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePortfoliosRequestContentWithDefaults() *UpdatePortfoliosRequestContent {
	this := UpdatePortfoliosRequestContent{}
	return &this
}

// GetPortfolios returns the Portfolios field value
func (o *UpdatePortfoliosRequestContent) GetPortfolios() []UpdatePortfolio {
	if o == nil {
		var ret []UpdatePortfolio
		return ret
	}

	return o.Portfolios
}

// GetPortfoliosOk returns a tuple with the Portfolios field value
// and a boolean to check if the value has been set.
func (o *UpdatePortfoliosRequestContent) GetPortfoliosOk() ([]UpdatePortfolio, bool) {
	if o == nil {
		return nil, false
	}
	return o.Portfolios, true
}

// SetPortfolios sets field value
func (o *UpdatePortfoliosRequestContent) SetPortfolios(v []UpdatePortfolio) {
	o.Portfolios = v
}

func (o UpdatePortfoliosRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["portfolios"] = o.Portfolios
	return toSerialize, nil
}

type NullableUpdatePortfoliosRequestContent struct {
	value *UpdatePortfoliosRequestContent
	isSet bool
}

func (v NullableUpdatePortfoliosRequestContent) Get() *UpdatePortfoliosRequestContent {
	return v.value
}

func (v *NullableUpdatePortfoliosRequestContent) Set(val *UpdatePortfoliosRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePortfoliosRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePortfoliosRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePortfoliosRequestContent(val *UpdatePortfoliosRequestContent) *NullableUpdatePortfoliosRequestContent {
	return &NullableUpdatePortfoliosRequestContent{value: val, isSet: true}
}

func (v NullableUpdatePortfoliosRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdatePortfoliosRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
