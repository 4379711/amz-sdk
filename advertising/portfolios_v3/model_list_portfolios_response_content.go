package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPortfoliosResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPortfoliosResponseContent{}

// ListPortfoliosResponseContent struct for ListPortfoliosResponseContent
type ListPortfoliosResponseContent struct {
	// The total number of entities
	TotalResults *float32 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken  *string     `json:"nextToken,omitempty"`
	Portfolios []Portfolio `json:"portfolios,omitempty"`
}

// NewListPortfoliosResponseContent instantiates a new ListPortfoliosResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPortfoliosResponseContent() *ListPortfoliosResponseContent {
	this := ListPortfoliosResponseContent{}
	return &this
}

// NewListPortfoliosResponseContentWithDefaults instantiates a new ListPortfoliosResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPortfoliosResponseContentWithDefaults() *ListPortfoliosResponseContent {
	this := ListPortfoliosResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ListPortfoliosResponseContent) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosResponseContent) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ListPortfoliosResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *ListPortfoliosResponseContent) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListPortfoliosResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListPortfoliosResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListPortfoliosResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetPortfolios returns the Portfolios field value if set, zero value otherwise.
func (o *ListPortfoliosResponseContent) GetPortfolios() []Portfolio {
	if o == nil || IsNil(o.Portfolios) {
		var ret []Portfolio
		return ret
	}
	return o.Portfolios
}

// GetPortfoliosOk returns a tuple with the Portfolios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosResponseContent) GetPortfoliosOk() ([]Portfolio, bool) {
	if o == nil || IsNil(o.Portfolios) {
		return nil, false
	}
	return o.Portfolios, true
}

// HasPortfolios returns a boolean if a field has been set.
func (o *ListPortfoliosResponseContent) HasPortfolios() bool {
	if o != nil && !IsNil(o.Portfolios) {
		return true
	}

	return false
}

// SetPortfolios gets a reference to the given []Portfolio and assigns it to the Portfolios field.
func (o *ListPortfoliosResponseContent) SetPortfolios(v []Portfolio) {
	o.Portfolios = v
}

func (o ListPortfoliosResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Portfolios) {
		toSerialize["portfolios"] = o.Portfolios
	}
	return toSerialize, nil
}

type NullableListPortfoliosResponseContent struct {
	value *ListPortfoliosResponseContent
	isSet bool
}

func (v NullableListPortfoliosResponseContent) Get() *ListPortfoliosResponseContent {
	return v.value
}

func (v *NullableListPortfoliosResponseContent) Set(val *ListPortfoliosResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListPortfoliosResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListPortfoliosResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPortfoliosResponseContent(val *ListPortfoliosResponseContent) *NullableListPortfoliosResponseContent {
	return &NullableListPortfoliosResponseContent{value: val, isSet: true}
}

func (v NullableListPortfoliosResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPortfoliosResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
