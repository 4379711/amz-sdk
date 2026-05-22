package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AgeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRange{}

// AgeRange struct for AgeRange
type AgeRange struct {
	// Name of Age Range. This field is OPTIONAL if the Age Range object is being used as an input.
	Name *string `json:"name,omitempty"`
	// Id of Age Range. This field is REQUIRED if the Age Range object is being used as an input. Use the GetRefinementsForCategory to retrieve Age Range Node IDs.
	Id *string `json:"id,omitempty"`
}

// NewAgeRange instantiates a new AgeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRange() *AgeRange {
	this := AgeRange{}
	return &this
}

// NewAgeRangeWithDefaults instantiates a new AgeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRangeWithDefaults() *AgeRange {
	this := AgeRange{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgeRange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgeRange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgeRange) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgeRange) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRange) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgeRange) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgeRange) SetId(v string) {
	o.Id = &v
}

func (o AgeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAgeRange struct {
	value *AgeRange
	isSet bool
}

func (v NullableAgeRange) Get() *AgeRange {
	return v.value
}

func (v *NullableAgeRange) Set(val *AgeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRange(val *AgeRange) *NullableAgeRange {
	return &NullableAgeRange{value: val, isSet: true}
}

func (v NullableAgeRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAgeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
