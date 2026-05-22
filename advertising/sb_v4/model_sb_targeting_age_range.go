package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingAgeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingAgeRange{}

// SBTargetingAgeRange struct for SBTargetingAgeRange
type SBTargetingAgeRange struct {
	// Id of Age Range. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Age Range Refinement IDs.
	AgeRangeRefinementId string `json:"ageRangeRefinementId"`
	// Name of Age Range.
	Name *string `json:"name,omitempty"`
	// Translated name of Age Range based off locale sent in request.
	TranslatedName *string `json:"translatedName,omitempty"`
}

type _SBTargetingAgeRange SBTargetingAgeRange

// NewSBTargetingAgeRange instantiates a new SBTargetingAgeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingAgeRange(ageRangeRefinementId string) *SBTargetingAgeRange {
	this := SBTargetingAgeRange{}
	this.AgeRangeRefinementId = ageRangeRefinementId
	return &this
}

// NewSBTargetingAgeRangeWithDefaults instantiates a new SBTargetingAgeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingAgeRangeWithDefaults() *SBTargetingAgeRange {
	this := SBTargetingAgeRange{}
	return &this
}

// GetAgeRangeRefinementId returns the AgeRangeRefinementId field value
func (o *SBTargetingAgeRange) GetAgeRangeRefinementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgeRangeRefinementId
}

// GetAgeRangeRefinementIdOk returns a tuple with the AgeRangeRefinementId field value
// and a boolean to check if the value has been set.
func (o *SBTargetingAgeRange) GetAgeRangeRefinementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeRangeRefinementId, true
}

// SetAgeRangeRefinementId sets field value
func (o *SBTargetingAgeRange) SetAgeRangeRefinementId(v string) {
	o.AgeRangeRefinementId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SBTargetingAgeRange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingAgeRange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SBTargetingAgeRange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SBTargetingAgeRange) SetName(v string) {
	o.Name = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *SBTargetingAgeRange) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingAgeRange) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *SBTargetingAgeRange) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *SBTargetingAgeRange) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

func (o SBTargetingAgeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ageRangeRefinementId"] = o.AgeRangeRefinementId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TranslatedName) {
		toSerialize["translatedName"] = o.TranslatedName
	}
	return toSerialize, nil
}

type NullableSBTargetingAgeRange struct {
	value *SBTargetingAgeRange
	isSet bool
}

func (v NullableSBTargetingAgeRange) Get() *SBTargetingAgeRange {
	return v.value
}

func (v *NullableSBTargetingAgeRange) Set(val *SBTargetingAgeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingAgeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingAgeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingAgeRange(val *SBTargetingAgeRange) *NullableSBTargetingAgeRange {
	return &NullableSBTargetingAgeRange{value: val, isSet: true}
}

func (v NullableSBTargetingAgeRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingAgeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
