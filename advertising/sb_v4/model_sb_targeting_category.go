package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingCategory{}

// SBTargetingCategory Category details.
type SBTargetingCategory struct {
	AsinCountRange *SBTargetingIntegerRange `json:"asinCountRange,omitempty"`
	// If the category is targetable or not.
	IsTargetable *bool `json:"isTargetable,omitempty"`
	// The category refinement id of the parent category. Missing parentCategoryRefinementId signifies this is a root category.
	ParentCategoryRefinementId *string                         `json:"parentCategoryRefinementId,omitempty"`
	EstimatedReach             *SBTargetingEstimatedReachRange `json:"estimatedReach,omitempty"`
	// Name of category.
	Name *string `json:"name,omitempty"`
	// Translated name of the category.
	TranslatedName *string `json:"translatedName,omitempty"`
	// The category refinement id. Please use /sb/targets/categories or /sb/recommendations/targets/category to retrieve category IDs.
	CategoryRefinementId *string `json:"categoryRefinementId,omitempty"`
}

// NewSBTargetingCategory instantiates a new SBTargetingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingCategory() *SBTargetingCategory {
	this := SBTargetingCategory{}
	return &this
}

// NewSBTargetingCategoryWithDefaults instantiates a new SBTargetingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingCategoryWithDefaults() *SBTargetingCategory {
	this := SBTargetingCategory{}
	return &this
}

// GetAsinCountRange returns the AsinCountRange field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetAsinCountRange() SBTargetingIntegerRange {
	if o == nil || IsNil(o.AsinCountRange) {
		var ret SBTargetingIntegerRange
		return ret
	}
	return *o.AsinCountRange
}

// GetAsinCountRangeOk returns a tuple with the AsinCountRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetAsinCountRangeOk() (*SBTargetingIntegerRange, bool) {
	if o == nil || IsNil(o.AsinCountRange) {
		return nil, false
	}
	return o.AsinCountRange, true
}

// HasAsinCountRange returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasAsinCountRange() bool {
	if o != nil && !IsNil(o.AsinCountRange) {
		return true
	}

	return false
}

// SetAsinCountRange gets a reference to the given SBTargetingIntegerRange and assigns it to the AsinCountRange field.
func (o *SBTargetingCategory) SetAsinCountRange(v SBTargetingIntegerRange) {
	o.AsinCountRange = &v
}

// GetIsTargetable returns the IsTargetable field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetIsTargetable() bool {
	if o == nil || IsNil(o.IsTargetable) {
		var ret bool
		return ret
	}
	return *o.IsTargetable
}

// GetIsTargetableOk returns a tuple with the IsTargetable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetIsTargetableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTargetable) {
		return nil, false
	}
	return o.IsTargetable, true
}

// HasIsTargetable returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasIsTargetable() bool {
	if o != nil && !IsNil(o.IsTargetable) {
		return true
	}

	return false
}

// SetIsTargetable gets a reference to the given bool and assigns it to the IsTargetable field.
func (o *SBTargetingCategory) SetIsTargetable(v bool) {
	o.IsTargetable = &v
}

// GetParentCategoryRefinementId returns the ParentCategoryRefinementId field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetParentCategoryRefinementId() string {
	if o == nil || IsNil(o.ParentCategoryRefinementId) {
		var ret string
		return ret
	}
	return *o.ParentCategoryRefinementId
}

// GetParentCategoryRefinementIdOk returns a tuple with the ParentCategoryRefinementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetParentCategoryRefinementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCategoryRefinementId) {
		return nil, false
	}
	return o.ParentCategoryRefinementId, true
}

// HasParentCategoryRefinementId returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasParentCategoryRefinementId() bool {
	if o != nil && !IsNil(o.ParentCategoryRefinementId) {
		return true
	}

	return false
}

// SetParentCategoryRefinementId gets a reference to the given string and assigns it to the ParentCategoryRefinementId field.
func (o *SBTargetingCategory) SetParentCategoryRefinementId(v string) {
	o.ParentCategoryRefinementId = &v
}

// GetEstimatedReach returns the EstimatedReach field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetEstimatedReach() SBTargetingEstimatedReachRange {
	if o == nil || IsNil(o.EstimatedReach) {
		var ret SBTargetingEstimatedReachRange
		return ret
	}
	return *o.EstimatedReach
}

// GetEstimatedReachOk returns a tuple with the EstimatedReach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetEstimatedReachOk() (*SBTargetingEstimatedReachRange, bool) {
	if o == nil || IsNil(o.EstimatedReach) {
		return nil, false
	}
	return o.EstimatedReach, true
}

// HasEstimatedReach returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasEstimatedReach() bool {
	if o != nil && !IsNil(o.EstimatedReach) {
		return true
	}

	return false
}

// SetEstimatedReach gets a reference to the given SBTargetingEstimatedReachRange and assigns it to the EstimatedReach field.
func (o *SBTargetingCategory) SetEstimatedReach(v SBTargetingEstimatedReachRange) {
	o.EstimatedReach = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SBTargetingCategory) SetName(v string) {
	o.Name = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *SBTargetingCategory) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

// GetCategoryRefinementId returns the CategoryRefinementId field value if set, zero value otherwise.
func (o *SBTargetingCategory) GetCategoryRefinementId() string {
	if o == nil || IsNil(o.CategoryRefinementId) {
		var ret string
		return ret
	}
	return *o.CategoryRefinementId
}

// GetCategoryRefinementIdOk returns a tuple with the CategoryRefinementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingCategory) GetCategoryRefinementIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryRefinementId) {
		return nil, false
	}
	return o.CategoryRefinementId, true
}

// HasCategoryRefinementId returns a boolean if a field has been set.
func (o *SBTargetingCategory) HasCategoryRefinementId() bool {
	if o != nil && !IsNil(o.CategoryRefinementId) {
		return true
	}

	return false
}

// SetCategoryRefinementId gets a reference to the given string and assigns it to the CategoryRefinementId field.
func (o *SBTargetingCategory) SetCategoryRefinementId(v string) {
	o.CategoryRefinementId = &v
}

func (o SBTargetingCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AsinCountRange) {
		toSerialize["asinCountRange"] = o.AsinCountRange
	}
	if !IsNil(o.IsTargetable) {
		toSerialize["isTargetable"] = o.IsTargetable
	}
	if !IsNil(o.ParentCategoryRefinementId) {
		toSerialize["parentCategoryRefinementId"] = o.ParentCategoryRefinementId
	}
	if !IsNil(o.EstimatedReach) {
		toSerialize["estimatedReach"] = o.EstimatedReach
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TranslatedName) {
		toSerialize["translatedName"] = o.TranslatedName
	}
	if !IsNil(o.CategoryRefinementId) {
		toSerialize["categoryRefinementId"] = o.CategoryRefinementId
	}
	return toSerialize, nil
}

type NullableSBTargetingCategory struct {
	value *SBTargetingCategory
	isSet bool
}

func (v NullableSBTargetingCategory) Get() *SBTargetingCategory {
	return v.value
}

func (v *NullableSBTargetingCategory) Set(val *SBTargetingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingCategory(val *SBTargetingCategory) *NullableSBTargetingCategory {
	return &NullableSBTargetingCategory{value: val, isSet: true}
}

func (v NullableSBTargetingCategory) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
