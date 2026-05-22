package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDCategoryRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDCategoryRecommendation{}

// SDCategoryRecommendation A recommended category to target ads on
type SDCategoryRecommendation struct {
	// The category identifier
	Category *int32 `json:"category,omitempty"`
	// The category name
	Name *string `json:"name,omitempty"`
	// The path of the category within the category catalogue.
	Path                     []string                                          `json:"path,omitempty"`
	TargetableAsinCountRange *SDCategoryRecommendationTargetableAsinCountRange `json:"targetableAsinCountRange,omitempty"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation
	Rank *int32 `json:"rank,omitempty"`
}

// NewSDCategoryRecommendation instantiates a new SDCategoryRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDCategoryRecommendation() *SDCategoryRecommendation {
	this := SDCategoryRecommendation{}
	return &this
}

// NewSDCategoryRecommendationWithDefaults instantiates a new SDCategoryRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDCategoryRecommendationWithDefaults() *SDCategoryRecommendation {
	this := SDCategoryRecommendation{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SDCategoryRecommendation) GetCategory() int32 {
	if o == nil || IsNil(o.Category) {
		var ret int32
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendation) GetCategoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SDCategoryRecommendation) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given int32 and assigns it to the Category field.
func (o *SDCategoryRecommendation) SetCategory(v int32) {
	o.Category = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDCategoryRecommendation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDCategoryRecommendation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDCategoryRecommendation) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SDCategoryRecommendation) GetPath() []string {
	if o == nil || IsNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendation) GetPathOk() ([]string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SDCategoryRecommendation) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *SDCategoryRecommendation) SetPath(v []string) {
	o.Path = v
}

// GetTargetableAsinCountRange returns the TargetableAsinCountRange field value if set, zero value otherwise.
func (o *SDCategoryRecommendation) GetTargetableAsinCountRange() SDCategoryRecommendationTargetableAsinCountRange {
	if o == nil || IsNil(o.TargetableAsinCountRange) {
		var ret SDCategoryRecommendationTargetableAsinCountRange
		return ret
	}
	return *o.TargetableAsinCountRange
}

// GetTargetableAsinCountRangeOk returns a tuple with the TargetableAsinCountRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendation) GetTargetableAsinCountRangeOk() (*SDCategoryRecommendationTargetableAsinCountRange, bool) {
	if o == nil || IsNil(o.TargetableAsinCountRange) {
		return nil, false
	}
	return o.TargetableAsinCountRange, true
}

// HasTargetableAsinCountRange returns a boolean if a field has been set.
func (o *SDCategoryRecommendation) HasTargetableAsinCountRange() bool {
	if o != nil && !IsNil(o.TargetableAsinCountRange) {
		return true
	}

	return false
}

// SetTargetableAsinCountRange gets a reference to the given SDCategoryRecommendationTargetableAsinCountRange and assigns it to the TargetableAsinCountRange field.
func (o *SDCategoryRecommendation) SetTargetableAsinCountRange(v SDCategoryRecommendationTargetableAsinCountRange) {
	o.TargetableAsinCountRange = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SDCategoryRecommendation) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendation) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SDCategoryRecommendation) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *SDCategoryRecommendation) SetRank(v int32) {
	o.Rank = &v
}

func (o SDCategoryRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.TargetableAsinCountRange) {
		toSerialize["targetableAsinCountRange"] = o.TargetableAsinCountRange
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableSDCategoryRecommendation struct {
	value *SDCategoryRecommendation
	isSet bool
}

func (v NullableSDCategoryRecommendation) Get() *SDCategoryRecommendation {
	return v.value
}

func (v *NullableSDCategoryRecommendation) Set(val *SDCategoryRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCategoryRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCategoryRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCategoryRecommendation(val *SDCategoryRecommendation) *NullableSDCategoryRecommendation {
	return &NullableSDCategoryRecommendation{value: val, isSet: true}
}

func (v NullableSDCategoryRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCategoryRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
