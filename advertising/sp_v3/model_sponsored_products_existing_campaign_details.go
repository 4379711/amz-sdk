package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsExistingCampaignDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsExistingCampaignDetails{}

// SponsoredProductsExistingCampaignDetails The request object for creating a new target promotion group with existing campaigns. Please note that the adGroupIds provided need to contain the same Ad ASINs/SKUs combination as the Auto-Targeting adGroup for the target promotion group.
type SponsoredProductsExistingCampaignDetails struct {
	// AdGroupIds of existing manual campaigns to be used as part of the Target Promotion Group for     promoting product targets.
	ProductCampaignAdGroupIds []string `json:"productCampaignAdGroupIds,omitempty"`
	// AdGroupIds of existing manual campaigns to be used as part of the Target Promotion Group for     promoting keyword targets.
	KeywordCampaignAdGroupIds []string `json:"keywordCampaignAdGroupIds,omitempty"`
}

// NewSponsoredProductsExistingCampaignDetails instantiates a new SponsoredProductsExistingCampaignDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsExistingCampaignDetails() *SponsoredProductsExistingCampaignDetails {
	this := SponsoredProductsExistingCampaignDetails{}
	return &this
}

// NewSponsoredProductsExistingCampaignDetailsWithDefaults instantiates a new SponsoredProductsExistingCampaignDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsExistingCampaignDetailsWithDefaults() *SponsoredProductsExistingCampaignDetails {
	this := SponsoredProductsExistingCampaignDetails{}
	return &this
}

// GetProductCampaignAdGroupIds returns the ProductCampaignAdGroupIds field value if set, zero value otherwise.
func (o *SponsoredProductsExistingCampaignDetails) GetProductCampaignAdGroupIds() []string {
	if o == nil || IsNil(o.ProductCampaignAdGroupIds) {
		var ret []string
		return ret
	}
	return o.ProductCampaignAdGroupIds
}

// GetProductCampaignAdGroupIdsOk returns a tuple with the ProductCampaignAdGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsExistingCampaignDetails) GetProductCampaignAdGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductCampaignAdGroupIds) {
		return nil, false
	}
	return o.ProductCampaignAdGroupIds, true
}

// HasProductCampaignAdGroupIds returns a boolean if a field has been set.
func (o *SponsoredProductsExistingCampaignDetails) HasProductCampaignAdGroupIds() bool {
	if o != nil && !IsNil(o.ProductCampaignAdGroupIds) {
		return true
	}

	return false
}

// SetProductCampaignAdGroupIds gets a reference to the given []string and assigns it to the ProductCampaignAdGroupIds field.
func (o *SponsoredProductsExistingCampaignDetails) SetProductCampaignAdGroupIds(v []string) {
	o.ProductCampaignAdGroupIds = v
}

// GetKeywordCampaignAdGroupIds returns the KeywordCampaignAdGroupIds field value if set, zero value otherwise.
func (o *SponsoredProductsExistingCampaignDetails) GetKeywordCampaignAdGroupIds() []string {
	if o == nil || IsNil(o.KeywordCampaignAdGroupIds) {
		var ret []string
		return ret
	}
	return o.KeywordCampaignAdGroupIds
}

// GetKeywordCampaignAdGroupIdsOk returns a tuple with the KeywordCampaignAdGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsExistingCampaignDetails) GetKeywordCampaignAdGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.KeywordCampaignAdGroupIds) {
		return nil, false
	}
	return o.KeywordCampaignAdGroupIds, true
}

// HasKeywordCampaignAdGroupIds returns a boolean if a field has been set.
func (o *SponsoredProductsExistingCampaignDetails) HasKeywordCampaignAdGroupIds() bool {
	if o != nil && !IsNil(o.KeywordCampaignAdGroupIds) {
		return true
	}

	return false
}

// SetKeywordCampaignAdGroupIds gets a reference to the given []string and assigns it to the KeywordCampaignAdGroupIds field.
func (o *SponsoredProductsExistingCampaignDetails) SetKeywordCampaignAdGroupIds(v []string) {
	o.KeywordCampaignAdGroupIds = v
}

func (o SponsoredProductsExistingCampaignDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductCampaignAdGroupIds) {
		toSerialize["productCampaignAdGroupIds"] = o.ProductCampaignAdGroupIds
	}
	if !IsNil(o.KeywordCampaignAdGroupIds) {
		toSerialize["keywordCampaignAdGroupIds"] = o.KeywordCampaignAdGroupIds
	}
	return toSerialize, nil
}

type NullableSponsoredProductsExistingCampaignDetails struct {
	value *SponsoredProductsExistingCampaignDetails
	isSet bool
}

func (v NullableSponsoredProductsExistingCampaignDetails) Get() *SponsoredProductsExistingCampaignDetails {
	return v.value
}

func (v *NullableSponsoredProductsExistingCampaignDetails) Set(val *SponsoredProductsExistingCampaignDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsExistingCampaignDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsExistingCampaignDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsExistingCampaignDetails(val *SponsoredProductsExistingCampaignDetails) *NullableSponsoredProductsExistingCampaignDetails {
	return &NullableSponsoredProductsExistingCampaignDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsExistingCampaignDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsExistingCampaignDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
