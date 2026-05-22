package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalNegativeKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalNegativeKeyword{}

// SponsoredProductsUpdateGlobalNegativeKeyword struct for SponsoredProductsUpdateGlobalNegativeKeyword
type SponsoredProductsUpdateGlobalNegativeKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// Name for the keyword
	Name        *string                                           `json:"name,omitempty"`
	State       *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	KeywordText *SponsoredProductsGlobalNegativeKeywordText       `json:"keywordText,omitempty"`
}

type _SponsoredProductsUpdateGlobalNegativeKeyword SponsoredProductsUpdateGlobalNegativeKeyword

// NewSponsoredProductsUpdateGlobalNegativeKeyword instantiates a new SponsoredProductsUpdateGlobalNegativeKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalNegativeKeyword(keywordId string) *SponsoredProductsUpdateGlobalNegativeKeyword {
	this := SponsoredProductsUpdateGlobalNegativeKeyword{}
	this.KeywordId = keywordId
	return &this
}

// NewSponsoredProductsUpdateGlobalNegativeKeywordWithDefaults instantiates a new SponsoredProductsUpdateGlobalNegativeKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalNegativeKeywordWithDefaults() *SponsoredProductsUpdateGlobalNegativeKeyword {
	this := SponsoredProductsUpdateGlobalNegativeKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetKeywordText() SponsoredProductsGlobalNegativeKeywordText {
	if o == nil || IsNil(o.KeywordText) {
		var ret SponsoredProductsGlobalNegativeKeywordText
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalNegativeKeywordText, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given SponsoredProductsGlobalNegativeKeywordText and assigns it to the KeywordText field.
func (o *SponsoredProductsUpdateGlobalNegativeKeyword) SetKeywordText(v SponsoredProductsGlobalNegativeKeywordText) {
	o.KeywordText = &v
}

func (o SponsoredProductsUpdateGlobalNegativeKeyword) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsUpdateGlobalNegativeKeyword struct {
	value *SponsoredProductsUpdateGlobalNegativeKeyword
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalNegativeKeyword) Get() *SponsoredProductsUpdateGlobalNegativeKeyword {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalNegativeKeyword) Set(val *SponsoredProductsUpdateGlobalNegativeKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalNegativeKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalNegativeKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalNegativeKeyword(val *SponsoredProductsUpdateGlobalNegativeKeyword) *NullableSponsoredProductsUpdateGlobalNegativeKeyword {
	return &NullableSponsoredProductsUpdateGlobalNegativeKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalNegativeKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalNegativeKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
