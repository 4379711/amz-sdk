package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDAudienceRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDAudienceRecommendation{}

// SDAudienceRecommendation A recommended standard Amazon audience to target ads on
type SDAudienceRecommendation struct {
	// The audience identifier
	Audience *string `json:"audience,omitempty"`
	// The Amazon audience name
	Name *string `json:"name,omitempty"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation
	Rank *int32 `json:"rank,omitempty"`
}

// NewSDAudienceRecommendation instantiates a new SDAudienceRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDAudienceRecommendation() *SDAudienceRecommendation {
	this := SDAudienceRecommendation{}
	return &this
}

// NewSDAudienceRecommendationWithDefaults instantiates a new SDAudienceRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDAudienceRecommendationWithDefaults() *SDAudienceRecommendation {
	this := SDAudienceRecommendation{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *SDAudienceRecommendation) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAudienceRecommendation) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *SDAudienceRecommendation) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *SDAudienceRecommendation) SetAudience(v string) {
	o.Audience = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDAudienceRecommendation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAudienceRecommendation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDAudienceRecommendation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDAudienceRecommendation) SetName(v string) {
	o.Name = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SDAudienceRecommendation) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAudienceRecommendation) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SDAudienceRecommendation) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *SDAudienceRecommendation) SetRank(v int32) {
	o.Rank = &v
}

func (o SDAudienceRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableSDAudienceRecommendation struct {
	value *SDAudienceRecommendation
	isSet bool
}

func (v NullableSDAudienceRecommendation) Get() *SDAudienceRecommendation {
	return v.value
}

func (v *NullableSDAudienceRecommendation) Set(val *SDAudienceRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableSDAudienceRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableSDAudienceRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDAudienceRecommendation(val *SDAudienceRecommendation) *NullableSDAudienceRecommendation {
	return &NullableSDAudienceRecommendation{value: val, isSet: true}
}

func (v NullableSDAudienceRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDAudienceRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
