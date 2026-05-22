package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalKeywordServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalKeywordServingStatus{}

// SponsoredProductsGlobalKeywordServingStatus struct for SponsoredProductsGlobalKeywordServingStatus
type SponsoredProductsGlobalKeywordServingStatus struct {
	// Serving status details of Keyword
	StatusReasons []SponsoredProductsGlobalKeywordServingStatusReason `json:"statusReasons,omitempty"`
	// The marketplace serving statuses.
	MarketplaceKeywordServingStatus []SponsoredProductsMarketplaceLevelKeywordServingStatus `json:"marketplaceKeywordServingStatus,omitempty"`
	KeywordServingStatus            *SponsoredProductsGlobalEntityServingStatus             `json:"keywordServingStatus,omitempty"`
}

// NewSponsoredProductsGlobalKeywordServingStatus instantiates a new SponsoredProductsGlobalKeywordServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalKeywordServingStatus() *SponsoredProductsGlobalKeywordServingStatus {
	this := SponsoredProductsGlobalKeywordServingStatus{}
	return &this
}

// NewSponsoredProductsGlobalKeywordServingStatusWithDefaults instantiates a new SponsoredProductsGlobalKeywordServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalKeywordServingStatusWithDefaults() *SponsoredProductsGlobalKeywordServingStatus {
	this := SponsoredProductsGlobalKeywordServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordServingStatus) GetStatusReasons() []SponsoredProductsGlobalKeywordServingStatusReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalKeywordServingStatusReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalKeywordServingStatusReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalKeywordServingStatusReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsGlobalKeywordServingStatus) SetStatusReasons(v []SponsoredProductsGlobalKeywordServingStatusReason) {
	o.StatusReasons = v
}

// GetMarketplaceKeywordServingStatus returns the MarketplaceKeywordServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordServingStatus) GetMarketplaceKeywordServingStatus() []SponsoredProductsMarketplaceLevelKeywordServingStatus {
	if o == nil || IsNil(o.MarketplaceKeywordServingStatus) {
		var ret []SponsoredProductsMarketplaceLevelKeywordServingStatus
		return ret
	}
	return o.MarketplaceKeywordServingStatus
}

// GetMarketplaceKeywordServingStatusOk returns a tuple with the MarketplaceKeywordServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordServingStatus) GetMarketplaceKeywordServingStatusOk() ([]SponsoredProductsMarketplaceLevelKeywordServingStatus, bool) {
	if o == nil || IsNil(o.MarketplaceKeywordServingStatus) {
		return nil, false
	}
	return o.MarketplaceKeywordServingStatus, true
}

// HasMarketplaceKeywordServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordServingStatus) HasMarketplaceKeywordServingStatus() bool {
	if o != nil && !IsNil(o.MarketplaceKeywordServingStatus) {
		return true
	}

	return false
}

// SetMarketplaceKeywordServingStatus gets a reference to the given []SponsoredProductsMarketplaceLevelKeywordServingStatus and assigns it to the MarketplaceKeywordServingStatus field.
func (o *SponsoredProductsGlobalKeywordServingStatus) SetMarketplaceKeywordServingStatus(v []SponsoredProductsMarketplaceLevelKeywordServingStatus) {
	o.MarketplaceKeywordServingStatus = v
}

// GetKeywordServingStatus returns the KeywordServingStatus field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordServingStatus) GetKeywordServingStatus() SponsoredProductsGlobalEntityServingStatus {
	if o == nil || IsNil(o.KeywordServingStatus) {
		var ret SponsoredProductsGlobalEntityServingStatus
		return ret
	}
	return *o.KeywordServingStatus
}

// GetKeywordServingStatusOk returns a tuple with the KeywordServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordServingStatus) GetKeywordServingStatusOk() (*SponsoredProductsGlobalEntityServingStatus, bool) {
	if o == nil || IsNil(o.KeywordServingStatus) {
		return nil, false
	}
	return o.KeywordServingStatus, true
}

// HasKeywordServingStatus returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordServingStatus) HasKeywordServingStatus() bool {
	if o != nil && !IsNil(o.KeywordServingStatus) {
		return true
	}

	return false
}

// SetKeywordServingStatus gets a reference to the given SponsoredProductsGlobalEntityServingStatus and assigns it to the KeywordServingStatus field.
func (o *SponsoredProductsGlobalKeywordServingStatus) SetKeywordServingStatus(v SponsoredProductsGlobalEntityServingStatus) {
	o.KeywordServingStatus = &v
}

func (o SponsoredProductsGlobalKeywordServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.MarketplaceKeywordServingStatus) {
		toSerialize["marketplaceKeywordServingStatus"] = o.MarketplaceKeywordServingStatus
	}
	if !IsNil(o.KeywordServingStatus) {
		toSerialize["keywordServingStatus"] = o.KeywordServingStatus
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalKeywordServingStatus struct {
	value *SponsoredProductsGlobalKeywordServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordServingStatus) Get() *SponsoredProductsGlobalKeywordServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordServingStatus) Set(val *SponsoredProductsGlobalKeywordServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordServingStatus(val *SponsoredProductsGlobalKeywordServingStatus) *NullableSponsoredProductsGlobalKeywordServingStatus {
	return &NullableSponsoredProductsGlobalKeywordServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
