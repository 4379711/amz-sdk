package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AgeRangeLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRangeLoP{}

// AgeRangeLoP struct for AgeRangeLoP
type AgeRangeLoP struct {
	// Name of Age Range.
	Name *string `json:"name,omitempty"`
	// Id of Age Range. Use the POST /sp/targets/category/{categoryId}/refinements endpoint to retrieve Age Range Node IDs.
	Id *string `json:"id,omitempty"`
	// Translated name of Age Range based off locale sent in request.
	TranslatedName *string `json:"translatedName,omitempty"`
}

// NewAgeRangeLoP instantiates a new AgeRangeLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRangeLoP() *AgeRangeLoP {
	this := AgeRangeLoP{}
	return &this
}

// NewAgeRangeLoPWithDefaults instantiates a new AgeRangeLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRangeLoPWithDefaults() *AgeRangeLoP {
	this := AgeRangeLoP{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgeRangeLoP) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRangeLoP) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgeRangeLoP) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgeRangeLoP) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgeRangeLoP) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRangeLoP) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgeRangeLoP) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgeRangeLoP) SetId(v string) {
	o.Id = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *AgeRangeLoP) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRangeLoP) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *AgeRangeLoP) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *AgeRangeLoP) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

func (o AgeRangeLoP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TranslatedName) {
		toSerialize["translatedName"] = o.TranslatedName
	}
	return toSerialize, nil
}

type NullableAgeRangeLoP struct {
	value *AgeRangeLoP
	isSet bool
}

func (v NullableAgeRangeLoP) Get() *AgeRangeLoP {
	return v.value
}

func (v *NullableAgeRangeLoP) Set(val *AgeRangeLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRangeLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRangeLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRangeLoP(val *AgeRangeLoP) *NullableAgeRangeLoP {
	return &NullableAgeRangeLoP{value: val, isSet: true}
}

func (v NullableAgeRangeLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAgeRangeLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
