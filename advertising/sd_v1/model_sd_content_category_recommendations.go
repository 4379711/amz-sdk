package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDContentCategoryRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDContentCategoryRecommendations{}

// SDContentCategoryRecommendations A recommended content category to target ads on
type SDContentCategoryRecommendations struct {
	// The content category value
	ContentCategory *string `json:"contentCategory,omitempty"`
	// The content category name
	Name *string `json:"name,omitempty"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation
	Rank *int32 `json:"rank,omitempty"`
}

// NewSDContentCategoryRecommendations instantiates a new SDContentCategoryRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDContentCategoryRecommendations() *SDContentCategoryRecommendations {
	this := SDContentCategoryRecommendations{}
	return &this
}

// NewSDContentCategoryRecommendationsWithDefaults instantiates a new SDContentCategoryRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDContentCategoryRecommendationsWithDefaults() *SDContentCategoryRecommendations {
	this := SDContentCategoryRecommendations{}
	return &this
}

// GetContentCategory returns the ContentCategory field value if set, zero value otherwise.
func (o *SDContentCategoryRecommendations) GetContentCategory() string {
	if o == nil || IsNil(o.ContentCategory) {
		var ret string
		return ret
	}
	return *o.ContentCategory
}

// GetContentCategoryOk returns a tuple with the ContentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDContentCategoryRecommendations) GetContentCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ContentCategory) {
		return nil, false
	}
	return o.ContentCategory, true
}

// HasContentCategory returns a boolean if a field has been set.
func (o *SDContentCategoryRecommendations) HasContentCategory() bool {
	if o != nil && !IsNil(o.ContentCategory) {
		return true
	}

	return false
}

// SetContentCategory gets a reference to the given string and assigns it to the ContentCategory field.
func (o *SDContentCategoryRecommendations) SetContentCategory(v string) {
	o.ContentCategory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDContentCategoryRecommendations) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDContentCategoryRecommendations) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDContentCategoryRecommendations) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDContentCategoryRecommendations) SetName(v string) {
	o.Name = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SDContentCategoryRecommendations) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDContentCategoryRecommendations) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SDContentCategoryRecommendations) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *SDContentCategoryRecommendations) SetRank(v int32) {
	o.Rank = &v
}

func (o SDContentCategoryRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentCategory) {
		toSerialize["contentCategory"] = o.ContentCategory
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableSDContentCategoryRecommendations struct {
	value *SDContentCategoryRecommendations
	isSet bool
}

func (v NullableSDContentCategoryRecommendations) Get() *SDContentCategoryRecommendations {
	return v.value
}

func (v *NullableSDContentCategoryRecommendations) Set(val *SDContentCategoryRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableSDContentCategoryRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableSDContentCategoryRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDContentCategoryRecommendations(val *SDContentCategoryRecommendations) *NullableSDContentCategoryRecommendations {
	return &NullableSDContentCategoryRecommendations{value: val, isSet: true}
}

func (v NullableSDContentCategoryRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDContentCategoryRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
