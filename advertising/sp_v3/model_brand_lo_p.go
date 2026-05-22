package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandLoP{}

// BrandLoP struct for BrandLoP
type BrandLoP struct {
	// Name of brand.
	Name *string `json:"name,omitempty"`
	// Id of brand.
	Id *string `json:"id,omitempty"`
}

// NewBrandLoP instantiates a new BrandLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandLoP() *BrandLoP {
	this := BrandLoP{}
	return &this
}

// NewBrandLoPWithDefaults instantiates a new BrandLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandLoPWithDefaults() *BrandLoP {
	this := BrandLoP{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrandLoP) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLoP) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrandLoP) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrandLoP) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrandLoP) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLoP) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrandLoP) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrandLoP) SetId(v string) {
	o.Id = &v
}

func (o BrandLoP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableBrandLoP struct {
	value *BrandLoP
	isSet bool
}

func (v NullableBrandLoP) Get() *BrandLoP {
	return v.value
}

func (v *NullableBrandLoP) Set(val *BrandLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandLoP(val *BrandLoP) *NullableBrandLoP {
	return &NullableBrandLoP{value: val, isSet: true}
}

func (v NullableBrandLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
