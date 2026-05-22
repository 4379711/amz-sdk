package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Brand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Brand{}

// Brand struct for Brand
type Brand struct {
	// Name of brand. This field is OPTIONAL if the Brand object is being used as an input.
	Name *string `json:"name,omitempty"`
	// Id of brand. This field is REQUIRED if the Brand object is being used as an input. Use the GetRefinementsForCategory to retrieve Brand Node IDs.
	Id *string `json:"id,omitempty"`
}

// NewBrand instantiates a new Brand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrand() *Brand {
	this := Brand{}
	return &this
}

// NewBrandWithDefaults instantiates a new Brand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandWithDefaults() *Brand {
	this := Brand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Brand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Brand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Brand) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Brand) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Brand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Brand) SetId(v string) {
	o.Id = &v
}

func (o Brand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableBrand struct {
	value *Brand
	isSet bool
}

func (v NullableBrand) Get() *Brand {
	return v.value
}

func (v *NullableBrand) Set(val *Brand) {
	v.value = val
	v.isSet = true
}

func (v NullableBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrand(val *Brand) *NullableBrand {
	return &NullableBrand{value: val, isSet: true}
}

func (v NullableBrand) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
