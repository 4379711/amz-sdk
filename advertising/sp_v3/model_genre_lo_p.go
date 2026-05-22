package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GenreLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenreLoP{}

// GenreLoP struct for GenreLoP
type GenreLoP struct {
	// Name of Genre.
	Name *string `json:"name,omitempty"`
	// Id of Genre. Use the POST /sp/targets/category/{categoryId}/refinements endpoint to retrieve Genre Node IDs.
	Id *string `json:"id,omitempty"`
	// Translated name of the Genre based off locale send in the query parameter.
	TranslatedName *string `json:"translatedName,omitempty"`
}

// NewGenreLoP instantiates a new GenreLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenreLoP() *GenreLoP {
	this := GenreLoP{}
	return &this
}

// NewGenreLoPWithDefaults instantiates a new GenreLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenreLoPWithDefaults() *GenreLoP {
	this := GenreLoP{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GenreLoP) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreLoP) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GenreLoP) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GenreLoP) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GenreLoP) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreLoP) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GenreLoP) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GenreLoP) SetId(v string) {
	o.Id = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *GenreLoP) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreLoP) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *GenreLoP) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *GenreLoP) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

func (o GenreLoP) ToMap() (map[string]interface{}, error) {
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

type NullableGenreLoP struct {
	value *GenreLoP
	isSet bool
}

func (v NullableGenreLoP) Get() *GenreLoP {
	return v.value
}

func (v *NullableGenreLoP) Set(val *GenreLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableGenreLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableGenreLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenreLoP(val *GenreLoP) *NullableGenreLoP {
	return &NullableGenreLoP{value: val, isSet: true}
}

func (v NullableGenreLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenreLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
