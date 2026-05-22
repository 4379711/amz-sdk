package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdatePortfoliosResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePortfoliosResponseContent{}

// UpdatePortfoliosResponseContent struct for UpdatePortfoliosResponseContent
type UpdatePortfoliosResponseContent struct {
	Portfolios BulkPortfolioOperationResponse `json:"portfolios"`
}

type _UpdatePortfoliosResponseContent UpdatePortfoliosResponseContent

// NewUpdatePortfoliosResponseContent instantiates a new UpdatePortfoliosResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePortfoliosResponseContent(portfolios BulkPortfolioOperationResponse) *UpdatePortfoliosResponseContent {
	this := UpdatePortfoliosResponseContent{}
	this.Portfolios = portfolios
	return &this
}

// NewUpdatePortfoliosResponseContentWithDefaults instantiates a new UpdatePortfoliosResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePortfoliosResponseContentWithDefaults() *UpdatePortfoliosResponseContent {
	this := UpdatePortfoliosResponseContent{}
	return &this
}

// GetPortfolios returns the Portfolios field value
func (o *UpdatePortfoliosResponseContent) GetPortfolios() BulkPortfolioOperationResponse {
	if o == nil {
		var ret BulkPortfolioOperationResponse
		return ret
	}

	return o.Portfolios
}

// GetPortfoliosOk returns a tuple with the Portfolios field value
// and a boolean to check if the value has been set.
func (o *UpdatePortfoliosResponseContent) GetPortfoliosOk() (*BulkPortfolioOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Portfolios, true
}

// SetPortfolios sets field value
func (o *UpdatePortfoliosResponseContent) SetPortfolios(v BulkPortfolioOperationResponse) {
	o.Portfolios = v
}

func (o UpdatePortfoliosResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["portfolios"] = o.Portfolios
	return toSerialize, nil
}

type NullableUpdatePortfoliosResponseContent struct {
	value *UpdatePortfoliosResponseContent
	isSet bool
}

func (v NullableUpdatePortfoliosResponseContent) Get() *UpdatePortfoliosResponseContent {
	return v.value
}

func (v *NullableUpdatePortfoliosResponseContent) Set(val *UpdatePortfoliosResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePortfoliosResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePortfoliosResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePortfoliosResponseContent(val *UpdatePortfoliosResponseContent) *NullableUpdatePortfoliosResponseContent {
	return &NullableUpdatePortfoliosResponseContent{value: val, isSet: true}
}

func (v NullableUpdatePortfoliosResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdatePortfoliosResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
