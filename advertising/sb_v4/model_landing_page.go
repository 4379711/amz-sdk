package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the LandingPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LandingPage{}

// LandingPage struct for LandingPage
type LandingPage struct {
	Asins    []string         `json:"asins,omitempty"`
	PageType *LandingPageType `json:"pageType,omitempty"`
	// URL of an existing simple landing page or Store page. Vendors may also specify the URL of a custom landing page. If a custom URL is specified, the landing page must include the ASINs of at least three products that are advertised as part of the campaign. Do not include this property in the request if the asins property is also included, these properties are mutually exclusive. Note that brandVideo ads only support Store page as landing page and does not allow asins property.
	Url *string `json:"url,omitempty"`
}

// NewLandingPage instantiates a new LandingPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandingPage() *LandingPage {
	this := LandingPage{}
	return &this
}

// NewLandingPageWithDefaults instantiates a new LandingPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandingPageWithDefaults() *LandingPage {
	this := LandingPage{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *LandingPage) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *LandingPage) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *LandingPage) SetAsins(v []string) {
	o.Asins = v
}

// GetPageType returns the PageType field value if set, zero value otherwise.
func (o *LandingPage) GetPageType() LandingPageType {
	if o == nil || IsNil(o.PageType) {
		var ret LandingPageType
		return ret
	}
	return *o.PageType
}

// GetPageTypeOk returns a tuple with the PageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetPageTypeOk() (*LandingPageType, bool) {
	if o == nil || IsNil(o.PageType) {
		return nil, false
	}
	return o.PageType, true
}

// HasPageType returns a boolean if a field has been set.
func (o *LandingPage) HasPageType() bool {
	if o != nil && !IsNil(o.PageType) {
		return true
	}

	return false
}

// SetPageType gets a reference to the given LandingPageType and assigns it to the PageType field.
func (o *LandingPage) SetPageType(v LandingPageType) {
	o.PageType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LandingPage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LandingPage) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LandingPage) SetUrl(v string) {
	o.Url = &v
}

func (o LandingPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.PageType) {
		toSerialize["pageType"] = o.PageType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableLandingPage struct {
	value *LandingPage
	isSet bool
}

func (v NullableLandingPage) Get() *LandingPage {
	return v.value
}

func (v *NullableLandingPage) Set(val *LandingPage) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPage) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPage(val *LandingPage) *NullableLandingPage {
	return &NullableLandingPage{value: val, isSet: true}
}

func (v NullableLandingPage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
