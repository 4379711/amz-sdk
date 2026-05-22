package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsResponse{}

// CreativeRecommendationsResponse struct for CreativeRecommendationsResponse
type CreativeRecommendationsResponse struct {
	// Ordered list of Primary Headline recommendation groups.
	PrimaryHeadlines [][]TextRecommendation `json:"primaryHeadlines,omitempty"`
	// Ordered list of Secondary Headline recommendation groups.
	SecondaryHeadlines [][]TextRecommendation `json:"secondaryHeadlines,omitempty"`
}

// NewCreativeRecommendationsResponse instantiates a new CreativeRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsResponse() *CreativeRecommendationsResponse {
	this := CreativeRecommendationsResponse{}
	return &this
}

// NewCreativeRecommendationsResponseWithDefaults instantiates a new CreativeRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsResponseWithDefaults() *CreativeRecommendationsResponse {
	this := CreativeRecommendationsResponse{}
	return &this
}

// GetPrimaryHeadlines returns the PrimaryHeadlines field value if set, zero value otherwise.
func (o *CreativeRecommendationsResponse) GetPrimaryHeadlines() [][]TextRecommendation {
	if o == nil || IsNil(o.PrimaryHeadlines) {
		var ret [][]TextRecommendation
		return ret
	}
	return o.PrimaryHeadlines
}

// GetPrimaryHeadlinesOk returns a tuple with the PrimaryHeadlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsResponse) GetPrimaryHeadlinesOk() ([][]TextRecommendation, bool) {
	if o == nil || IsNil(o.PrimaryHeadlines) {
		return nil, false
	}
	return o.PrimaryHeadlines, true
}

// HasPrimaryHeadlines returns a boolean if a field has been set.
func (o *CreativeRecommendationsResponse) HasPrimaryHeadlines() bool {
	if o != nil && !IsNil(o.PrimaryHeadlines) {
		return true
	}

	return false
}

// SetPrimaryHeadlines gets a reference to the given [][]TextRecommendation and assigns it to the PrimaryHeadlines field.
func (o *CreativeRecommendationsResponse) SetPrimaryHeadlines(v [][]TextRecommendation) {
	o.PrimaryHeadlines = v
}

// GetSecondaryHeadlines returns the SecondaryHeadlines field value if set, zero value otherwise.
func (o *CreativeRecommendationsResponse) GetSecondaryHeadlines() [][]TextRecommendation {
	if o == nil || IsNil(o.SecondaryHeadlines) {
		var ret [][]TextRecommendation
		return ret
	}
	return o.SecondaryHeadlines
}

// GetSecondaryHeadlinesOk returns a tuple with the SecondaryHeadlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsResponse) GetSecondaryHeadlinesOk() ([][]TextRecommendation, bool) {
	if o == nil || IsNil(o.SecondaryHeadlines) {
		return nil, false
	}
	return o.SecondaryHeadlines, true
}

// HasSecondaryHeadlines returns a boolean if a field has been set.
func (o *CreativeRecommendationsResponse) HasSecondaryHeadlines() bool {
	if o != nil && !IsNil(o.SecondaryHeadlines) {
		return true
	}

	return false
}

// SetSecondaryHeadlines gets a reference to the given [][]TextRecommendation and assigns it to the SecondaryHeadlines field.
func (o *CreativeRecommendationsResponse) SetSecondaryHeadlines(v [][]TextRecommendation) {
	o.SecondaryHeadlines = v
}

func (o CreativeRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryHeadlines) {
		toSerialize["primaryHeadlines"] = o.PrimaryHeadlines
	}
	if !IsNil(o.SecondaryHeadlines) {
		toSerialize["secondaryHeadlines"] = o.SecondaryHeadlines
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationsResponse struct {
	value *CreativeRecommendationsResponse
	isSet bool
}

func (v NullableCreativeRecommendationsResponse) Get() *CreativeRecommendationsResponse {
	return v.value
}

func (v *NullableCreativeRecommendationsResponse) Set(val *CreativeRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsResponse(val *CreativeRecommendationsResponse) *NullableCreativeRecommendationsResponse {
	return &NullableCreativeRecommendationsResponse{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
