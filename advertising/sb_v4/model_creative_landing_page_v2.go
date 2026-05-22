package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeLandingPageV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeLandingPageV2{}

// CreativeLandingPageV2 Landing page V2, where type is String with allowed values listed, and url or asins of that type.
type CreativeLandingPageV2 struct {
	// The list of asins on the landingPage If type is PRODUCT_LIST. A minimum of 3 asins are required. For the 'PRODUCT_LIST' type, the asins property is mandatory, and the url should not be included.
	Asins []string `json:"asins,omitempty"`
	// Supported types are PRODUCT_LIST, STORE, DETAIL_PAGE, CUSTOM_URL. More could be added in future.
	Type *string `json:"type,omitempty"`
	// The url of the landingPage. When including the 'asins' property in the request, do not include this property, as they are mutually exclusive. For the PRODUCT_LIST type, the asins property is mandatory, and the url should not be included.
	Url *string `json:"url,omitempty"`
}

// NewCreativeLandingPageV2 instantiates a new CreativeLandingPageV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeLandingPageV2() *CreativeLandingPageV2 {
	this := CreativeLandingPageV2{}
	return &this
}

// NewCreativeLandingPageV2WithDefaults instantiates a new CreativeLandingPageV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeLandingPageV2WithDefaults() *CreativeLandingPageV2 {
	this := CreativeLandingPageV2{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreativeLandingPageV2) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeLandingPageV2) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreativeLandingPageV2) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreativeLandingPageV2) SetAsins(v []string) {
	o.Asins = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreativeLandingPageV2) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeLandingPageV2) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreativeLandingPageV2) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreativeLandingPageV2) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreativeLandingPageV2) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeLandingPageV2) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreativeLandingPageV2) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreativeLandingPageV2) SetUrl(v string) {
	o.Url = &v
}

func (o CreativeLandingPageV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCreativeLandingPageV2 struct {
	value *CreativeLandingPageV2
	isSet bool
}

func (v NullableCreativeLandingPageV2) Get() *CreativeLandingPageV2 {
	return v.value
}

func (v *NullableCreativeLandingPageV2) Set(val *CreativeLandingPageV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeLandingPageV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeLandingPageV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeLandingPageV2(val *CreativeLandingPageV2) *NullableCreativeLandingPageV2 {
	return &NullableCreativeLandingPageV2{value: val, isSet: true}
}

func (v NullableCreativeLandingPageV2) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeLandingPageV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
