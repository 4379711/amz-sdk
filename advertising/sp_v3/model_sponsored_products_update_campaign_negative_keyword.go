package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateCampaignNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateCampaignNegativeKeyword{}

// SponsoredProductsUpdateCampaignNegativeKeyword struct for SponsoredProductsUpdateCampaignNegativeKeyword
type SponsoredProductsUpdateCampaignNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string                                      `json:"keywordId"`
	State     *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
}

type _SponsoredProductsUpdateCampaignNegativeKeyword SponsoredProductsUpdateCampaignNegativeKeyword

// NewSponsoredProductsUpdateCampaignNegativeKeyword instantiates a new SponsoredProductsUpdateCampaignNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateCampaignNegativeKeyword(keywordId string) *SponsoredProductsUpdateCampaignNegativeKeyword {
	this := SponsoredProductsUpdateCampaignNegativeKeyword{}
	this.KeywordId = keywordId
	return &this
}

// NewSponsoredProductsUpdateCampaignNegativeKeywordWithDefaults instantiates a new SponsoredProductsUpdateCampaignNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateCampaignNegativeKeywordWithDefaults() *SponsoredProductsUpdateCampaignNegativeKeyword {
	this := SponsoredProductsUpdateCampaignNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateCampaignNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

func (o SponsoredProductsUpdateCampaignNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateCampaignNegativeKeyword struct {
	value *SponsoredProductsUpdateCampaignNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsUpdateCampaignNegativeKeyword) Get() *SponsoredProductsUpdateCampaignNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsUpdateCampaignNegativeKeyword) Set(val *SponsoredProductsUpdateCampaignNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateCampaignNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateCampaignNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateCampaignNegativeKeyword(val *SponsoredProductsUpdateCampaignNegativeKeyword) *NullableSponsoredProductsUpdateCampaignNegativeKeyword {
	return &NullableSponsoredProductsUpdateCampaignNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateCampaignNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateCampaignNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
