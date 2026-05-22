package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPortfoliosRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPortfoliosRequestContent{}

// ListPortfoliosRequestContent struct for ListPortfoliosRequestContent
type ListPortfoliosRequestContent struct {
	PortfolioIdFilter *ObjectIdFilter    `json:"portfolioIdFilter,omitempty"`
	StateFilter       *EntityStateFilter `json:"stateFilter,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// whether to get a list of targetingClauses with extended data fields (creationDate, lastUpdateDate, servingStatus).
	IncludeExtendedDataFields *bool       `json:"includeExtendedDataFields,omitempty"`
	NameFilter                *NameFilter `json:"nameFilter,omitempty"`
}

// NewListPortfoliosRequestContent instantiates a new ListPortfoliosRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPortfoliosRequestContent() *ListPortfoliosRequestContent {
	this := ListPortfoliosRequestContent{}
	return &this
}

// NewListPortfoliosRequestContentWithDefaults instantiates a new ListPortfoliosRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPortfoliosRequestContentWithDefaults() *ListPortfoliosRequestContent {
	this := ListPortfoliosRequestContent{}
	return &this
}

// GetPortfolioIdFilter returns the PortfolioIdFilter field value if set, zero value otherwise.
func (o *ListPortfoliosRequestContent) GetPortfolioIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.PortfolioIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.PortfolioIdFilter
}

// GetPortfolioIdFilterOk returns a tuple with the PortfolioIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosRequestContent) GetPortfolioIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.PortfolioIdFilter) {
		return nil, false
	}
	return o.PortfolioIdFilter, true
}

// HasPortfolioIdFilter returns a boolean if a field has been set.
func (o *ListPortfoliosRequestContent) HasPortfolioIdFilter() bool {
	if o != nil && !IsNil(o.PortfolioIdFilter) {
		return true
	}

	return false
}

// SetPortfolioIdFilter gets a reference to the given ObjectIdFilter and assigns it to the PortfolioIdFilter field.
func (o *ListPortfoliosRequestContent) SetPortfolioIdFilter(v ObjectIdFilter) {
	o.PortfolioIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *ListPortfoliosRequestContent) GetStateFilter() EntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret EntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosRequestContent) GetStateFilterOk() (*EntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *ListPortfoliosRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given EntityStateFilter and assigns it to the StateFilter field.
func (o *ListPortfoliosRequestContent) SetStateFilter(v EntityStateFilter) {
	o.StateFilter = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListPortfoliosRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListPortfoliosRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListPortfoliosRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *ListPortfoliosRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *ListPortfoliosRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *ListPortfoliosRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *ListPortfoliosRequestContent) GetNameFilter() NameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret NameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortfoliosRequestContent) GetNameFilterOk() (*NameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *ListPortfoliosRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given NameFilter and assigns it to the NameFilter field.
func (o *ListPortfoliosRequestContent) SetNameFilter(v NameFilter) {
	o.NameFilter = &v
}

func (o ListPortfoliosRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortfolioIdFilter) {
		toSerialize["portfolioIdFilter"] = o.PortfolioIdFilter
	}
	if !IsNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	if !IsNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	return toSerialize, nil
}

type NullableListPortfoliosRequestContent struct {
	value *ListPortfoliosRequestContent
	isSet bool
}

func (v NullableListPortfoliosRequestContent) Get() *ListPortfoliosRequestContent {
	return v.value
}

func (v *NullableListPortfoliosRequestContent) Set(val *ListPortfoliosRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListPortfoliosRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListPortfoliosRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPortfoliosRequestContent(val *ListPortfoliosRequestContent) *NullableListPortfoliosRequestContent {
	return &NullableListPortfoliosRequestContent{value: val, isSet: true}
}

func (v NullableListPortfoliosRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPortfoliosRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
