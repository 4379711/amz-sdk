package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalCampaignNegativeKeyword{}

// SponsoredProductsUpdateGlobalCampaignNegativeKeyword struct for SponsoredProductsUpdateGlobalCampaignNegativeKeyword
type SponsoredProductsUpdateGlobalCampaignNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// Name for the keyword
	Name        *string                                           `json:"name,omitempty"`
	State       *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	KeywordText *SponsoredProductsGlobalNegativeKeywordText       `json:"keywordText,omitempty"`
}

type _SponsoredProductsUpdateGlobalCampaignNegativeKeyword SponsoredProductsUpdateGlobalCampaignNegativeKeyword

// NewSponsoredProductsUpdateGlobalCampaignNegativeKeyword instantiates a new SponsoredProductsUpdateGlobalCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalCampaignNegativeKeyword(keywordId string) *SponsoredProductsUpdateGlobalCampaignNegativeKeyword {
	this := SponsoredProductsUpdateGlobalCampaignNegativeKeyword{}
	this.KeywordId = keywordId
	return &this
}

// NewSponsoredProductsUpdateGlobalCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsUpdateGlobalCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalCampaignNegativeKeywordWithDefaults() *SponsoredProductsUpdateGlobalCampaignNegativeKeyword {
	this := SponsoredProductsUpdateGlobalCampaignNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetKeywordText() SponsoredProductsGlobalNegativeKeywordText {
	if o == nil || IsNil(o.KeywordText) {
		var ret SponsoredProductsGlobalNegativeKeywordText
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalNegativeKeywordText, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given SponsoredProductsGlobalNegativeKeywordText and assigns it to the KeywordText field.
func (o *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) SetKeywordText(v SponsoredProductsGlobalNegativeKeywordText) {
	o.KeywordText = &v
}

func (o SponsoredProductsUpdateGlobalCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.KeywordText) {
		toSerialize["keywordText"] = o.KeywordText
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword struct {
	value *SponsoredProductsUpdateGlobalCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword) Get() *SponsoredProductsUpdateGlobalCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword) Set(val *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword(val *SponsoredProductsUpdateGlobalCampaignNegativeKeyword) *NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword {
	return &NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
