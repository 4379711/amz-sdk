package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Subpage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subpage{}

// Subpage struct for Subpage
type Subpage struct {
	PageTitle *string `json:"pageTitle,omitempty"`
	Asin      *string `json:"asin,omitempty"`
	Url       *string `json:"url,omitempty"`
}

// NewSubpage instantiates a new Subpage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubpage() *Subpage {
	this := Subpage{}
	return &this
}

// NewSubpageWithDefaults instantiates a new Subpage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubpageWithDefaults() *Subpage {
	this := Subpage{}
	return &this
}

// GetPageTitle returns the PageTitle field value if set, zero value otherwise.
func (o *Subpage) GetPageTitle() string {
	if o == nil || IsNil(o.PageTitle) {
		var ret string
		return ret
	}
	return *o.PageTitle
}

// GetPageTitleOk returns a tuple with the PageTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subpage) GetPageTitleOk() (*string, bool) {
	if o == nil || IsNil(o.PageTitle) {
		return nil, false
	}
	return o.PageTitle, true
}

// HasPageTitle returns a boolean if a field has been set.
func (o *Subpage) HasPageTitle() bool {
	if o != nil && !IsNil(o.PageTitle) {
		return true
	}

	return false
}

// SetPageTitle gets a reference to the given string and assigns it to the PageTitle field.
func (o *Subpage) SetPageTitle(v string) {
	o.PageTitle = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *Subpage) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subpage) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *Subpage) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *Subpage) SetAsin(v string) {
	o.Asin = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Subpage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subpage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Subpage) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Subpage) SetUrl(v string) {
	o.Url = &v
}

func (o Subpage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageTitle) {
		toSerialize["pageTitle"] = o.PageTitle
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSubpage struct {
	value *Subpage
	isSet bool
}

func (v NullableSubpage) Get() *Subpage {
	return v.value
}

func (v *NullableSubpage) Set(val *Subpage) {
	v.value = val
	v.isSet = true
}

func (v NullableSubpage) IsSet() bool {
	return v.isSet
}

func (v *NullableSubpage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubpage(val *Subpage) *NullableSubpage {
	return &NullableSubpage{value: val, isSet: true}
}

func (v NullableSubpage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubpage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
