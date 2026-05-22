package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProgramEligibilityV2RequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramEligibilityV2RequestContent{}

// ProgramEligibilityV2RequestContent A request to evaluate account level eligibility for Amazon ad programs (Sponsored Products, Sponsored Brands, Sponsored Display, Stores, DirectToConsumer, Amazon Attribution, etc).
type ProgramEligibilityV2RequestContent struct {
	// Max results for pagination
	MaxResults *float32 `json:"maxResults,omitempty"`
	// The pagination token that is required to go to the next page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewProgramEligibilityV2RequestContent instantiates a new ProgramEligibilityV2RequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramEligibilityV2RequestContent() *ProgramEligibilityV2RequestContent {
	this := ProgramEligibilityV2RequestContent{}
	return &this
}

// NewProgramEligibilityV2RequestContentWithDefaults instantiates a new ProgramEligibilityV2RequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramEligibilityV2RequestContentWithDefaults() *ProgramEligibilityV2RequestContent {
	this := ProgramEligibilityV2RequestContent{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ProgramEligibilityV2RequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramEligibilityV2RequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ProgramEligibilityV2RequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ProgramEligibilityV2RequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ProgramEligibilityV2RequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramEligibilityV2RequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ProgramEligibilityV2RequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ProgramEligibilityV2RequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ProgramEligibilityV2RequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableProgramEligibilityV2RequestContent struct {
	value *ProgramEligibilityV2RequestContent
	isSet bool
}

func (v NullableProgramEligibilityV2RequestContent) Get() *ProgramEligibilityV2RequestContent {
	return v.value
}

func (v *NullableProgramEligibilityV2RequestContent) Set(val *ProgramEligibilityV2RequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramEligibilityV2RequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramEligibilityV2RequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramEligibilityV2RequestContent(val *ProgramEligibilityV2RequestContent) *NullableProgramEligibilityV2RequestContent {
	return &NullableProgramEligibilityV2RequestContent{value: val, isSet: true}
}

func (v NullableProgramEligibilityV2RequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProgramEligibilityV2RequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
