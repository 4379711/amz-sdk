package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateNegativeKeyword{}

// SponsoredProductsUpdateNegativeKeyword struct for SponsoredProductsUpdateNegativeKeyword
type SponsoredProductsUpdateNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string                                      `json:"keywordId"`
	State     *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
}

type _SponsoredProductsUpdateNegativeKeyword SponsoredProductsUpdateNegativeKeyword

// NewSponsoredProductsUpdateNegativeKeyword instantiates a new SponsoredProductsUpdateNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateNegativeKeyword(keywordId string) *SponsoredProductsUpdateNegativeKeyword {
	this := SponsoredProductsUpdateNegativeKeyword{}
	this.KeywordId = keywordId
	return &this
}

// NewSponsoredProductsUpdateNegativeKeywordWithDefaults instantiates a new SponsoredProductsUpdateNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateNegativeKeywordWithDefaults() *SponsoredProductsUpdateNegativeKeyword {
	this := SponsoredProductsUpdateNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsUpdateNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsUpdateNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateNegativeKeyword) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

func (o SponsoredProductsUpdateNegativeKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateNegativeKeyword struct {
	value *SponsoredProductsUpdateNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsUpdateNegativeKeyword) Get() *SponsoredProductsUpdateNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsUpdateNegativeKeyword) Set(val *SponsoredProductsUpdateNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateNegativeKeyword(val *SponsoredProductsUpdateNegativeKeyword) *NullableSponsoredProductsUpdateNegativeKeyword {
	return &NullableSponsoredProductsUpdateNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
