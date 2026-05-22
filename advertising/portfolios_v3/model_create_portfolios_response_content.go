package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CreatePortfoliosResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePortfoliosResponseContent{}

// CreatePortfoliosResponseContent struct for CreatePortfoliosResponseContent
type CreatePortfoliosResponseContent struct {
	Portfolios BulkPortfolioOperationResponse `json:"portfolios"`
}

type _CreatePortfoliosResponseContent CreatePortfoliosResponseContent

// NewCreatePortfoliosResponseContent instantiates a new CreatePortfoliosResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePortfoliosResponseContent(portfolios BulkPortfolioOperationResponse) *CreatePortfoliosResponseContent {
	this := CreatePortfoliosResponseContent{}
	this.Portfolios = portfolios
	return &this
}

// NewCreatePortfoliosResponseContentWithDefaults instantiates a new CreatePortfoliosResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePortfoliosResponseContentWithDefaults() *CreatePortfoliosResponseContent {
	this := CreatePortfoliosResponseContent{}
	return &this
}

// GetPortfolios returns the Portfolios field value
func (o *CreatePortfoliosResponseContent) GetPortfolios() BulkPortfolioOperationResponse {
	if o == nil {
		var ret BulkPortfolioOperationResponse
		return ret
	}

	return o.Portfolios
}

// GetPortfoliosOk returns a tuple with the Portfolios field value
// and a boolean to check if the value has been set.
func (o *CreatePortfoliosResponseContent) GetPortfoliosOk() (*BulkPortfolioOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Portfolios, true
}

// SetPortfolios sets field value
func (o *CreatePortfoliosResponseContent) SetPortfolios(v BulkPortfolioOperationResponse) {
	o.Portfolios = v
}

func (o CreatePortfoliosResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["portfolios"] = o.Portfolios
	return toSerialize, nil
}

type NullableCreatePortfoliosResponseContent struct {
	value *CreatePortfoliosResponseContent
	isSet bool
}

func (v NullableCreatePortfoliosResponseContent) Get() *CreatePortfoliosResponseContent {
	return v.value
}

func (v *NullableCreatePortfoliosResponseContent) Set(val *CreatePortfoliosResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePortfoliosResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePortfoliosResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePortfoliosResponseContent(val *CreatePortfoliosResponseContent) *NullableCreatePortfoliosResponseContent {
	return &NullableCreatePortfoliosResponseContent{value: val, isSet: true}
}

func (v NullableCreatePortfoliosResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreatePortfoliosResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
