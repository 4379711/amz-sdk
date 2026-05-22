package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the LinkObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkObject{}

// LinkObject A Link object.
type LinkObject struct {
	// A URI for this object.
	Href string `json:"href"`
	// An identifier for this object.
	Name *string `json:"name,omitempty"`
}

type _LinkObject LinkObject

// NewLinkObject instantiates a new LinkObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkObject(href string) *LinkObject {
	this := LinkObject{}
	this.Href = href
	return &this
}

// NewLinkObjectWithDefaults instantiates a new LinkObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkObjectWithDefaults() *LinkObject {
	this := LinkObject{}
	return &this
}

// GetHref returns the Href field value
func (o *LinkObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LinkObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LinkObject) SetHref(v string) {
	o.Href = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinkObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinkObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinkObject) SetName(v string) {
	o.Name = &v
}

func (o LinkObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLinkObject struct {
	value *LinkObject
	isSet bool
}

func (v NullableLinkObject) Get() *LinkObject {
	return v.value
}

func (v *NullableLinkObject) Set(val *LinkObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkObject(val *LinkObject) *NullableLinkObject {
	return &NullableLinkObject{value: val, isSet: true}
}

func (v NullableLinkObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLinkObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
