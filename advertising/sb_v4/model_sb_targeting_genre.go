package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingGenre type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingGenre{}

// SBTargetingGenre struct for SBTargetingGenre
type SBTargetingGenre struct {
	// Id of Genre. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Genre Refinement IDs.
	GenreRefinementId string `json:"genreRefinementId"`
	// Name of Genre.
	Name *string `json:"name,omitempty"`
	// Translated name of Genre based off locale sent in request.
	TranslatedName *string `json:"translatedName,omitempty"`
}

type _SBTargetingGenre SBTargetingGenre

// NewSBTargetingGenre instantiates a new SBTargetingGenre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingGenre(genreRefinementId string) *SBTargetingGenre {
	this := SBTargetingGenre{}
	this.GenreRefinementId = genreRefinementId
	return &this
}

// NewSBTargetingGenreWithDefaults instantiates a new SBTargetingGenre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingGenreWithDefaults() *SBTargetingGenre {
	this := SBTargetingGenre{}
	return &this
}

// GetGenreRefinementId returns the GenreRefinementId field value
func (o *SBTargetingGenre) GetGenreRefinementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GenreRefinementId
}

// GetGenreRefinementIdOk returns a tuple with the GenreRefinementId field value
// and a boolean to check if the value has been set.
func (o *SBTargetingGenre) GetGenreRefinementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenreRefinementId, true
}

// SetGenreRefinementId sets field value
func (o *SBTargetingGenre) SetGenreRefinementId(v string) {
	o.GenreRefinementId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SBTargetingGenre) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGenre) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SBTargetingGenre) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SBTargetingGenre) SetName(v string) {
	o.Name = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *SBTargetingGenre) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGenre) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *SBTargetingGenre) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *SBTargetingGenre) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

func (o SBTargetingGenre) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["genreRefinementId"] = o.GenreRefinementId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TranslatedName) {
		toSerialize["translatedName"] = o.TranslatedName
	}
	return toSerialize, nil
}

type NullableSBTargetingGenre struct {
	value *SBTargetingGenre
	isSet bool
}

func (v NullableSBTargetingGenre) Get() *SBTargetingGenre {
	return v.value
}

func (v *NullableSBTargetingGenre) Set(val *SBTargetingGenre) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingGenre) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingGenre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingGenre(val *SBTargetingGenre) *NullableSBTargetingGenre {
	return &NullableSBTargetingGenre{value: val, isSet: true}
}

func (v NullableSBTargetingGenre) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingGenre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
