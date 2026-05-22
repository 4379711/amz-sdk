package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDCategoryRecommendationV33 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDCategoryRecommendationV33{}

// SDCategoryRecommendationV33 A recommended category to target ads on
type SDCategoryRecommendationV33 struct {
	// The category identifier
	Category *int32 `json:"category,omitempty"`
	// The category name
	Name *string `json:"name,omitempty"`
	// The translated category name by requested locale, field will not be provided if locale is not provided or campaign localization service is down.
	TranslatedName *string `json:"translatedName,omitempty"`
	// The path of the category within the category catalogue.
	Path []string `json:"path,omitempty"`
	// The translated path of the category within the category catalogue by requested locale, field will not be provided if locale is not provided or campaign localization is down.
	TranslatedPath           []string                                             `json:"translatedPath,omitempty"`
	TargetableAsinCountRange *SDCategoryRecommendationV33TargetableAsinCountRange `json:"targetableAsinCountRange,omitempty"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation.
	Rank *int32 `json:"rank,omitempty"`
}

// NewSDCategoryRecommendationV33 instantiates a new SDCategoryRecommendationV33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDCategoryRecommendationV33() *SDCategoryRecommendationV33 {
	this := SDCategoryRecommendationV33{}
	return &this
}

// NewSDCategoryRecommendationV33WithDefaults instantiates a new SDCategoryRecommendationV33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDCategoryRecommendationV33WithDefaults() *SDCategoryRecommendationV33 {
	this := SDCategoryRecommendationV33{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetCategory() int32 {
	if o == nil || IsNil(o.Category) {
		var ret int32
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetCategoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given int32 and assigns it to the Category field.
func (o *SDCategoryRecommendationV33) SetCategory(v int32) {
	o.Category = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDCategoryRecommendationV33) SetName(v string) {
	o.Name = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *SDCategoryRecommendationV33) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetPath() []string {
	if o == nil || IsNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetPathOk() ([]string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *SDCategoryRecommendationV33) SetPath(v []string) {
	o.Path = v
}

// GetTranslatedPath returns the TranslatedPath field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetTranslatedPath() []string {
	if o == nil || IsNil(o.TranslatedPath) {
		var ret []string
		return ret
	}
	return o.TranslatedPath
}

// GetTranslatedPathOk returns a tuple with the TranslatedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetTranslatedPathOk() ([]string, bool) {
	if o == nil || IsNil(o.TranslatedPath) {
		return nil, false
	}
	return o.TranslatedPath, true
}

// HasTranslatedPath returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasTranslatedPath() bool {
	if o != nil && !IsNil(o.TranslatedPath) {
		return true
	}

	return false
}

// SetTranslatedPath gets a reference to the given []string and assigns it to the TranslatedPath field.
func (o *SDCategoryRecommendationV33) SetTranslatedPath(v []string) {
	o.TranslatedPath = v
}

// GetTargetableAsinCountRange returns the TargetableAsinCountRange field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetTargetableAsinCountRange() SDCategoryRecommendationV33TargetableAsinCountRange {
	if o == nil || IsNil(o.TargetableAsinCountRange) {
		var ret SDCategoryRecommendationV33TargetableAsinCountRange
		return ret
	}
	return *o.TargetableAsinCountRange
}

// GetTargetableAsinCountRangeOk returns a tuple with the TargetableAsinCountRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetTargetableAsinCountRangeOk() (*SDCategoryRecommendationV33TargetableAsinCountRange, bool) {
	if o == nil || IsNil(o.TargetableAsinCountRange) {
		return nil, false
	}
	return o.TargetableAsinCountRange, true
}

// HasTargetableAsinCountRange returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasTargetableAsinCountRange() bool {
	if o != nil && !IsNil(o.TargetableAsinCountRange) {
		return true
	}

	return false
}

// SetTargetableAsinCountRange gets a reference to the given SDCategoryRecommendationV33TargetableAsinCountRange and assigns it to the TargetableAsinCountRange field.
func (o *SDCategoryRecommendationV33) SetTargetableAsinCountRange(v SDCategoryRecommendationV33TargetableAsinCountRange) {
	o.TargetableAsinCountRange = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SDCategoryRecommendationV33) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDCategoryRecommendationV33) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SDCategoryRecommendationV33) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *SDCategoryRecommendationV33) SetRank(v int32) {
	o.Rank = &v
}

func (o SDCategoryRecommendationV33) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TranslatedName) {
		toSerialize["translatedName"] = o.TranslatedName
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.TranslatedPath) {
		toSerialize["translatedPath"] = o.TranslatedPath
	}
	if !IsNil(o.TargetableAsinCountRange) {
		toSerialize["targetableAsinCountRange"] = o.TargetableAsinCountRange
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableSDCategoryRecommendationV33 struct {
	value *SDCategoryRecommendationV33
	isSet bool
}

func (v NullableSDCategoryRecommendationV33) Get() *SDCategoryRecommendationV33 {
	return v.value
}

func (v *NullableSDCategoryRecommendationV33) Set(val *SDCategoryRecommendationV33) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCategoryRecommendationV33) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCategoryRecommendationV33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCategoryRecommendationV33(val *SDCategoryRecommendationV33) *NullableSDCategoryRecommendationV33 {
	return &NullableSDCategoryRecommendationV33{value: val, isSet: true}
}

func (v NullableSDCategoryRecommendationV33) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCategoryRecommendationV33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
