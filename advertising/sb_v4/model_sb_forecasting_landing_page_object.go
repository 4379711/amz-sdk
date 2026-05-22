package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingLandingPageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingLandingPageObject{}

// SBForecastingLandingPageObject struct for SBForecastingLandingPageObject
type SBForecastingLandingPageObject struct {
	// Landing page information.
	LandingPageUrl *string `json:"landingPageUrl,omitempty"`
}

// NewSBForecastingLandingPageObject instantiates a new SBForecastingLandingPageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingLandingPageObject() *SBForecastingLandingPageObject {
	this := SBForecastingLandingPageObject{}
	return &this
}

// NewSBForecastingLandingPageObjectWithDefaults instantiates a new SBForecastingLandingPageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingLandingPageObjectWithDefaults() *SBForecastingLandingPageObject {
	this := SBForecastingLandingPageObject{}
	return &this
}

// GetLandingPageUrl returns the LandingPageUrl field value if set, zero value otherwise.
func (o *SBForecastingLandingPageObject) GetLandingPageUrl() string {
	if o == nil || IsNil(o.LandingPageUrl) {
		var ret string
		return ret
	}
	return *o.LandingPageUrl
}

// GetLandingPageUrlOk returns a tuple with the LandingPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingLandingPageObject) GetLandingPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPageUrl) {
		return nil, false
	}
	return o.LandingPageUrl, true
}

// HasLandingPageUrl returns a boolean if a field has been set.
func (o *SBForecastingLandingPageObject) HasLandingPageUrl() bool {
	if o != nil && !IsNil(o.LandingPageUrl) {
		return true
	}

	return false
}

// SetLandingPageUrl gets a reference to the given string and assigns it to the LandingPageUrl field.
func (o *SBForecastingLandingPageObject) SetLandingPageUrl(v string) {
	o.LandingPageUrl = &v
}

func (o SBForecastingLandingPageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LandingPageUrl) {
		toSerialize["landingPageUrl"] = o.LandingPageUrl
	}
	return toSerialize, nil
}

type NullableSBForecastingLandingPageObject struct {
	value *SBForecastingLandingPageObject
	isSet bool
}

func (v NullableSBForecastingLandingPageObject) Get() *SBForecastingLandingPageObject {
	return v.value
}

func (v *NullableSBForecastingLandingPageObject) Set(val *SBForecastingLandingPageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingLandingPageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingLandingPageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingLandingPageObject(val *SBForecastingLandingPageObject) *NullableSBForecastingLandingPageObject {
	return &NullableSBForecastingLandingPageObject{value: val, isSet: true}
}

func (v NullableSBForecastingLandingPageObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingLandingPageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
