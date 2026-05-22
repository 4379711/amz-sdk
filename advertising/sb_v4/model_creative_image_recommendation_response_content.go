package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeImageRecommendationResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeImageRecommendationResponseContent{}

// CreativeImageRecommendationResponseContent struct for CreativeImageRecommendationResponseContent
type CreativeImageRecommendationResponseContent struct {
	// The total number of results returned by an operation.
	TotalResults *float32 `json:"totalResults,omitempty"`
	// Recommendations are sorted on relevancy score, i.e. more relevant image has lesser array index value
	Recommendations []CreativeImageRecommendationEntry `json:"recommendations,omitempty"`
}

// NewCreativeImageRecommendationResponseContent instantiates a new CreativeImageRecommendationResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeImageRecommendationResponseContent() *CreativeImageRecommendationResponseContent {
	this := CreativeImageRecommendationResponseContent{}
	return &this
}

// NewCreativeImageRecommendationResponseContentWithDefaults instantiates a new CreativeImageRecommendationResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeImageRecommendationResponseContentWithDefaults() *CreativeImageRecommendationResponseContent {
	this := CreativeImageRecommendationResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *CreativeImageRecommendationResponseContent) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationResponseContent) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *CreativeImageRecommendationResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *CreativeImageRecommendationResponseContent) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *CreativeImageRecommendationResponseContent) GetRecommendations() []CreativeImageRecommendationEntry {
	if o == nil || IsNil(o.Recommendations) {
		var ret []CreativeImageRecommendationEntry
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeImageRecommendationResponseContent) GetRecommendationsOk() ([]CreativeImageRecommendationEntry, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *CreativeImageRecommendationResponseContent) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []CreativeImageRecommendationEntry and assigns it to the Recommendations field.
func (o *CreativeImageRecommendationResponseContent) SetRecommendations(v []CreativeImageRecommendationEntry) {
	o.Recommendations = v
}

func (o CreativeImageRecommendationResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableCreativeImageRecommendationResponseContent struct {
	value *CreativeImageRecommendationResponseContent
	isSet bool
}

func (v NullableCreativeImageRecommendationResponseContent) Get() *CreativeImageRecommendationResponseContent {
	return v.value
}

func (v *NullableCreativeImageRecommendationResponseContent) Set(val *CreativeImageRecommendationResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeImageRecommendationResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeImageRecommendationResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeImageRecommendationResponseContent(val *CreativeImageRecommendationResponseContent) *NullableCreativeImageRecommendationResponseContent {
	return &NullableCreativeImageRecommendationResponseContent{value: val, isSet: true}
}

func (v NullableCreativeImageRecommendationResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeImageRecommendationResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
