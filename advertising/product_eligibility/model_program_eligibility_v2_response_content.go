package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProgramEligibilityV2ResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramEligibilityV2ResponseContent{}

// ProgramEligibilityV2ResponseContent An object of program eligibility responses for an advertiser.
type ProgramEligibilityV2ResponseContent struct {
	EligibilityStatusLists []MarketplaceEntitiesEligibilityStatusList `json:"eligibilityStatusLists,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewProgramEligibilityV2ResponseContent instantiates a new ProgramEligibilityV2ResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramEligibilityV2ResponseContent() *ProgramEligibilityV2ResponseContent {
	this := ProgramEligibilityV2ResponseContent{}
	return &this
}

// NewProgramEligibilityV2ResponseContentWithDefaults instantiates a new ProgramEligibilityV2ResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramEligibilityV2ResponseContentWithDefaults() *ProgramEligibilityV2ResponseContent {
	this := ProgramEligibilityV2ResponseContent{}
	return &this
}

// GetEligibilityStatusLists returns the EligibilityStatusLists field value if set, zero value otherwise.
func (o *ProgramEligibilityV2ResponseContent) GetEligibilityStatusLists() []MarketplaceEntitiesEligibilityStatusList {
	if o == nil || IsNil(o.EligibilityStatusLists) {
		var ret []MarketplaceEntitiesEligibilityStatusList
		return ret
	}
	return o.EligibilityStatusLists
}

// GetEligibilityStatusListsOk returns a tuple with the EligibilityStatusLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramEligibilityV2ResponseContent) GetEligibilityStatusListsOk() ([]MarketplaceEntitiesEligibilityStatusList, bool) {
	if o == nil || IsNil(o.EligibilityStatusLists) {
		return nil, false
	}
	return o.EligibilityStatusLists, true
}

// HasEligibilityStatusLists returns a boolean if a field has been set.
func (o *ProgramEligibilityV2ResponseContent) HasEligibilityStatusLists() bool {
	if o != nil && !IsNil(o.EligibilityStatusLists) {
		return true
	}

	return false
}

// SetEligibilityStatusLists gets a reference to the given []MarketplaceEntitiesEligibilityStatusList and assigns it to the EligibilityStatusLists field.
func (o *ProgramEligibilityV2ResponseContent) SetEligibilityStatusLists(v []MarketplaceEntitiesEligibilityStatusList) {
	o.EligibilityStatusLists = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ProgramEligibilityV2ResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramEligibilityV2ResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ProgramEligibilityV2ResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ProgramEligibilityV2ResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ProgramEligibilityV2ResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EligibilityStatusLists) {
		toSerialize["eligibilityStatusLists"] = o.EligibilityStatusLists
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableProgramEligibilityV2ResponseContent struct {
	value *ProgramEligibilityV2ResponseContent
	isSet bool
}

func (v NullableProgramEligibilityV2ResponseContent) Get() *ProgramEligibilityV2ResponseContent {
	return v.value
}

func (v *NullableProgramEligibilityV2ResponseContent) Set(val *ProgramEligibilityV2ResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramEligibilityV2ResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramEligibilityV2ResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramEligibilityV2ResponseContent(val *ProgramEligibilityV2ResponseContent) *NullableProgramEligibilityV2ResponseContent {
	return &NullableProgramEligibilityV2ResponseContent{value: val, isSet: true}
}

func (v NullableProgramEligibilityV2ResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProgramEligibilityV2ResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
