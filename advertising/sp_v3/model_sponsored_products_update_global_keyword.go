package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalKeyword{}

// SponsoredProductsUpdateGlobalKeyword struct for SponsoredProductsUpdateGlobalKeyword
type SponsoredProductsUpdateGlobalKeyword struct {
	// The identifier of the keyword.
	KeywordId string `json:"keywordId"`
	// Name for the Keyword
	Name        *string                                           `json:"name,omitempty"`
	State       *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	Bid         *SponsoredProductsGlobalBid                       `json:"bid,omitempty"`
	KeywordText *SponsoredProductsGlobalKeywordText               `json:"keywordText,omitempty"`
}

type _SponsoredProductsUpdateGlobalKeyword SponsoredProductsUpdateGlobalKeyword

// NewSponsoredProductsUpdateGlobalKeyword instantiates a new SponsoredProductsUpdateGlobalKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalKeyword(keywordId string) *SponsoredProductsUpdateGlobalKeyword {
	this := SponsoredProductsUpdateGlobalKeyword{}
	this.KeywordId = keywordId
	return &this
}

// NewSponsoredProductsUpdateGlobalKeywordWithDefaults instantiates a new SponsoredProductsUpdateGlobalKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalKeywordWithDefaults() *SponsoredProductsUpdateGlobalKeyword {
	this := SponsoredProductsUpdateGlobalKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsUpdateGlobalKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsUpdateGlobalKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalKeyword) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalKeyword) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalKeyword) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalKeyword) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalKeyword) GetBid() SponsoredProductsGlobalBid {
	if o == nil || IsNil(o.Bid) {
		var ret SponsoredProductsGlobalBid
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) GetBidOk() (*SponsoredProductsGlobalBid, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given SponsoredProductsGlobalBid and assigns it to the Bid field.
func (o *SponsoredProductsUpdateGlobalKeyword) SetBid(v SponsoredProductsGlobalBid) {
	o.Bid = &v
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalKeyword) GetKeywordText() SponsoredProductsGlobalKeywordText {
	if o == nil || IsNil(o.KeywordText) {
		var ret SponsoredProductsGlobalKeywordText
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) GetKeywordTextOk() (*SponsoredProductsGlobalKeywordText, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalKeyword) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given SponsoredProductsGlobalKeywordText and assigns it to the KeywordText field.
func (o *SponsoredProductsUpdateGlobalKeyword) SetKeywordText(v SponsoredProductsGlobalKeywordText) {
	o.KeywordText = &v
}

func (o SponsoredProductsUpdateGlobalKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.KeywordText) {
		toSerialize["keywordText"] = o.KeywordText
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateGlobalKeyword struct {
	value *SponsoredProductsUpdateGlobalKeyword
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalKeyword) Get() *SponsoredProductsUpdateGlobalKeyword {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalKeyword) Set(val *SponsoredProductsUpdateGlobalKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalKeyword(val *SponsoredProductsUpdateGlobalKeyword) *NullableSponsoredProductsUpdateGlobalKeyword {
	return &NullableSponsoredProductsUpdateGlobalKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
