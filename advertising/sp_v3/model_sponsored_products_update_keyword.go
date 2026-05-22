package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateKeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateKeyword{}

// SponsoredProductsUpdateKeyword struct for SponsoredProductsUpdateKeyword
type SponsoredProductsUpdateKeyword struct {
	// The identifier of the keyword.
	KeywordId string                                      `json:"keywordId"`
	State     *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
	// Bid associated with this keyword. Applicable to biddable match types only. For more information about bid constraints by marketplace, see [bid limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace).
	Bid NullableFloat64 `json:"bid,omitempty"`
}

type _SponsoredProductsUpdateKeyword SponsoredProductsUpdateKeyword

// NewSponsoredProductsUpdateKeyword instantiates a new SponsoredProductsUpdateKeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateKeyword(keywordId string) *SponsoredProductsUpdateKeyword {
	this := SponsoredProductsUpdateKeyword{}
	this.KeywordId = keywordId
	return &this
}

// NewSponsoredProductsUpdateKeywordWithDefaults instantiates a new SponsoredProductsUpdateKeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateKeywordWithDefaults() *SponsoredProductsUpdateKeyword {
	this := SponsoredProductsUpdateKeyword{}
	return &this
}

// GetKeywordId returns the KeywordId field value
func (o *SponsoredProductsUpdateKeyword) GetKeywordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateKeyword) GetKeywordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordId, true
}

// SetKeywordId sets field value
func (o *SponsoredProductsUpdateKeyword) SetKeywordId(v string) {
	o.KeywordId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateKeyword) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateKeyword) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateKeyword) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateKeyword) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SponsoredProductsUpdateKeyword) GetBid() float64 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float64
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SponsoredProductsUpdateKeyword) GetBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateKeyword) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat64 and assigns it to the Bid field.
func (o *SponsoredProductsUpdateKeyword) SetBid(v float64) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *SponsoredProductsUpdateKeyword) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *SponsoredProductsUpdateKeyword) UnsetBid() {
	o.Bid.Unset()
}

func (o SponsoredProductsUpdateKeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordId"] = o.KeywordId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateKeyword struct {
	value *SponsoredProductsUpdateKeyword
	isSet bool
}

func (v NullableSponsoredProductsUpdateKeyword) Get() *SponsoredProductsUpdateKeyword {
	return v.value
}

func (v *NullableSponsoredProductsUpdateKeyword) Set(val *SponsoredProductsUpdateKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateKeyword(val *SponsoredProductsUpdateKeyword) *NullableSponsoredProductsUpdateKeyword {
	return &NullableSponsoredProductsUpdateKeyword{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateKeyword) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
