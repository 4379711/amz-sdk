package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreatePortfoliosRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePortfoliosRequestContent{}

// CreatePortfoliosRequestContent struct for CreatePortfoliosRequestContent
type CreatePortfoliosRequestContent struct {
	// An array of portfolio to create.
	Portfolios []CreatePortfolio `json:"portfolios"`
}

type _CreatePortfoliosRequestContent CreatePortfoliosRequestContent

// NewCreatePortfoliosRequestContent instantiates a new CreatePortfoliosRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePortfoliosRequestContent(portfolios []CreatePortfolio) *CreatePortfoliosRequestContent {
	this := CreatePortfoliosRequestContent{}
	this.Portfolios = portfolios
	return &this
}

// NewCreatePortfoliosRequestContentWithDefaults instantiates a new CreatePortfoliosRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePortfoliosRequestContentWithDefaults() *CreatePortfoliosRequestContent {
	this := CreatePortfoliosRequestContent{}
	return &this
}

// GetPortfolios returns the Portfolios field value
func (o *CreatePortfoliosRequestContent) GetPortfolios() []CreatePortfolio {
	if o == nil {
		var ret []CreatePortfolio
		return ret
	}

	return o.Portfolios
}

// GetPortfoliosOk returns a tuple with the Portfolios field value
// and a boolean to check if the value has been set.
func (o *CreatePortfoliosRequestContent) GetPortfoliosOk() ([]CreatePortfolio, bool) {
	if o == nil {
		return nil, false
	}
	return o.Portfolios, true
}

// SetPortfolios sets field value
func (o *CreatePortfoliosRequestContent) SetPortfolios(v []CreatePortfolio) {
	o.Portfolios = v
}

func (o CreatePortfoliosRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["portfolios"] = o.Portfolios
	return toSerialize, nil
}

type NullableCreatePortfoliosRequestContent struct {
	value *CreatePortfoliosRequestContent
	isSet bool
}

func (v NullableCreatePortfoliosRequestContent) Get() *CreatePortfoliosRequestContent {
	return v.value
}

func (v *NullableCreatePortfoliosRequestContent) Set(val *CreatePortfoliosRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePortfoliosRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePortfoliosRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePortfoliosRequestContent(val *CreatePortfoliosRequestContent) *NullableCreatePortfoliosRequestContent {
	return &NullableCreatePortfoliosRequestContent{value: val, isSet: true}
}

func (v NullableCreatePortfoliosRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreatePortfoliosRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
