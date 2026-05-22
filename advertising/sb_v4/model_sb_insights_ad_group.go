package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsAdGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsAdGroup{}

// SBInsightsAdGroup The ad group settings.
type SBInsightsAdGroup struct {
	Keywords []SBInsightsKeyword `json:"keywords,omitempty"`
	AdFormat SBInsightsAdFormat  `json:"adFormat"`
}

type _SBInsightsAdGroup SBInsightsAdGroup

// NewSBInsightsAdGroup instantiates a new SBInsightsAdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsAdGroup(adFormat SBInsightsAdFormat) *SBInsightsAdGroup {
	this := SBInsightsAdGroup{}
	this.AdFormat = adFormat
	return &this
}

// NewSBInsightsAdGroupWithDefaults instantiates a new SBInsightsAdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsAdGroupWithDefaults() *SBInsightsAdGroup {
	this := SBInsightsAdGroup{}
	return &this
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SBInsightsAdGroup) GetKeywords() []SBInsightsKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SBInsightsKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsAdGroup) GetKeywordsOk() ([]SBInsightsKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SBInsightsAdGroup) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SBInsightsKeyword and assigns it to the Keywords field.
func (o *SBInsightsAdGroup) SetKeywords(v []SBInsightsKeyword) {
	o.Keywords = v
}

// GetAdFormat returns the AdFormat field value
func (o *SBInsightsAdGroup) GetAdFormat() SBInsightsAdFormat {
	if o == nil {
		var ret SBInsightsAdFormat
		return ret
	}

	return o.AdFormat
}

// GetAdFormatOk returns a tuple with the AdFormat field value
// and a boolean to check if the value has been set.
func (o *SBInsightsAdGroup) GetAdFormatOk() (*SBInsightsAdFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdFormat, true
}

// SetAdFormat sets field value
func (o *SBInsightsAdGroup) SetAdFormat(v SBInsightsAdFormat) {
	o.AdFormat = v
}

func (o SBInsightsAdGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	toSerialize["adFormat"] = o.AdFormat
	return toSerialize, nil
}

type NullableSBInsightsAdGroup struct {
	value *SBInsightsAdGroup
	isSet bool
}

func (v NullableSBInsightsAdGroup) Get() *SBInsightsAdGroup {
	return v.value
}

func (v *NullableSBInsightsAdGroup) Set(val *SBInsightsAdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsAdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsAdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsAdGroup(val *SBInsightsAdGroup) *NullableSBInsightsAdGroup {
	return &NullableSBInsightsAdGroup{value: val, isSet: true}
}

func (v NullableSBInsightsAdGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsAdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
