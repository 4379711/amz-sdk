package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeLandingPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeLandingPage{}

// CreativeLandingPage Landing page.
type CreativeLandingPage struct {
	// The list of asins on the landingPage If type is PRODUCT_LIST.
	Asins []string                 `json:"asins,omitempty"`
	Type  *CreativeLandingPageType `json:"type,omitempty"`
	// The url of the landingPage.
	Value *string `json:"value,omitempty"`
}

// NewCreativeLandingPage instantiates a new CreativeLandingPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeLandingPage() *CreativeLandingPage {
	this := CreativeLandingPage{}
	return &this
}

// NewCreativeLandingPageWithDefaults instantiates a new CreativeLandingPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeLandingPageWithDefaults() *CreativeLandingPage {
	this := CreativeLandingPage{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreativeLandingPage) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeLandingPage) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreativeLandingPage) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreativeLandingPage) SetAsins(v []string) {
	o.Asins = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreativeLandingPage) GetType() CreativeLandingPageType {
	if o == nil || IsNil(o.Type) {
		var ret CreativeLandingPageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeLandingPage) GetTypeOk() (*CreativeLandingPageType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreativeLandingPage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CreativeLandingPageType and assigns it to the Type field.
func (o *CreativeLandingPage) SetType(v CreativeLandingPageType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreativeLandingPage) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeLandingPage) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreativeLandingPage) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreativeLandingPage) SetValue(v string) {
	o.Value = &v
}

func (o CreativeLandingPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCreativeLandingPage struct {
	value *CreativeLandingPage
	isSet bool
}

func (v NullableCreativeLandingPage) Get() *CreativeLandingPage {
	return v.value
}

func (v *NullableCreativeLandingPage) Set(val *CreativeLandingPage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeLandingPage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeLandingPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeLandingPage(val *CreativeLandingPage) *NullableCreativeLandingPage {
	return &NullableCreativeLandingPage{value: val, isSet: true}
}

func (v NullableCreativeLandingPage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeLandingPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
