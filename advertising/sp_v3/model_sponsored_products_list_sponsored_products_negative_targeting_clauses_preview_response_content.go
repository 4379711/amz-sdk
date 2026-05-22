package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent{}

// SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent struct for SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent
type SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken                *string                                    `json:"nextToken,omitempty"`
	NegativeTargetingClauses []SponsoredProductsNegativeTargetingClause `json:"negativeTargetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent instantiates a new SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent {
	this := SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent {
	this := SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) GetNegativeTargetingClauses() []SponsoredProductsNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		var ret []SponsoredProductsNegativeTargetingClause
		return ret
	}
	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// HasNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) HasNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.NegativeTargetingClauses) {
		return true
	}

	return false
}

// SetNegativeTargetingClauses gets a reference to the given []SponsoredProductsNegativeTargetingClause and assigns it to the NegativeTargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) SetNegativeTargetingClauses(v []SponsoredProductsNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.NegativeTargetingClauses) {
		toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent struct {
	value *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) Get() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) Set(val *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent(val *SponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesPreviewResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
