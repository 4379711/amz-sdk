package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetCampaignAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetCampaignAttributes{}

// SponsoredProductsTargetCampaignAttributes struct for SponsoredProductsTargetCampaignAttributes
type SponsoredProductsTargetCampaignAttributes struct {
	// The identifier of the target marketplace.
	MarketplaceId string `json:"marketplaceId"`
	// The end date for the campaign in formats according to https://tools.ietf.org/html/rfc3339#section-5.6.
	EndDate *string `json:"endDate,omitempty"`
	// The name to be appended to the campaign. If new name already exists, a number will be appended i.e. if \"Campaign Name Copy\" exist, we will name it \"Campaign Name Copy 1\"
	NameSuffix string `json:"nameSuffix"`
	// The start date of the campaign in formats according to https://tools.ietf.org/html/rfc3339#section-5.6.
	StartDate *string `json:"startDate,omitempty"`
	// The advertiser id per the targeted marketplace. Advertiser id per marketplace can fetched through /v2/profiles API.
	AdvertiserId string `json:"advertiserId"`
	// The budget for the campaign.
	Budget *float64                                   `json:"budget,omitempty"`
	Status SponsoredProductsCreateOrUpdateEntityState `json:"status"`
}

type _SponsoredProductsTargetCampaignAttributes SponsoredProductsTargetCampaignAttributes

// NewSponsoredProductsTargetCampaignAttributes instantiates a new SponsoredProductsTargetCampaignAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetCampaignAttributes(marketplaceId string, nameSuffix string, advertiserId string, status SponsoredProductsCreateOrUpdateEntityState) *SponsoredProductsTargetCampaignAttributes {
	this := SponsoredProductsTargetCampaignAttributes{}
	this.MarketplaceId = marketplaceId
	this.NameSuffix = nameSuffix
	this.AdvertiserId = advertiserId
	this.Status = status
	return &this
}

// NewSponsoredProductsTargetCampaignAttributesWithDefaults instantiates a new SponsoredProductsTargetCampaignAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetCampaignAttributesWithDefaults() *SponsoredProductsTargetCampaignAttributes {
	this := SponsoredProductsTargetCampaignAttributes{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *SponsoredProductsTargetCampaignAttributes) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *SponsoredProductsTargetCampaignAttributes) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SponsoredProductsTargetCampaignAttributes) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SponsoredProductsTargetCampaignAttributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SponsoredProductsTargetCampaignAttributes) SetEndDate(v string) {
	o.EndDate = &v
}

// GetNameSuffix returns the NameSuffix field value
func (o *SponsoredProductsTargetCampaignAttributes) GetNameSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameSuffix
}

// GetNameSuffixOk returns a tuple with the NameSuffix field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetNameSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameSuffix, true
}

// SetNameSuffix sets field value
func (o *SponsoredProductsTargetCampaignAttributes) SetNameSuffix(v string) {
	o.NameSuffix = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SponsoredProductsTargetCampaignAttributes) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SponsoredProductsTargetCampaignAttributes) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SponsoredProductsTargetCampaignAttributes) SetStartDate(v string) {
	o.StartDate = &v
}

// GetAdvertiserId returns the AdvertiserId field value
func (o *SponsoredProductsTargetCampaignAttributes) GetAdvertiserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvertiserId
}

// GetAdvertiserIdOk returns a tuple with the AdvertiserId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetAdvertiserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvertiserId, true
}

// SetAdvertiserId sets field value
func (o *SponsoredProductsTargetCampaignAttributes) SetAdvertiserId(v string) {
	o.AdvertiserId = v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *SponsoredProductsTargetCampaignAttributes) GetBudget() float64 {
	if o == nil || IsNil(o.Budget) {
		var ret float64
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *SponsoredProductsTargetCampaignAttributes) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float64 and assigns it to the Budget field.
func (o *SponsoredProductsTargetCampaignAttributes) SetBudget(v float64) {
	o.Budget = &v
}

// GetStatus returns the Status field value
func (o *SponsoredProductsTargetCampaignAttributes) GetStatus() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetCampaignAttributes) GetStatusOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SponsoredProductsTargetCampaignAttributes) SetStatus(v SponsoredProductsCreateOrUpdateEntityState) {
	o.Status = v
}

func (o SponsoredProductsTargetCampaignAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["nameSuffix"] = o.NameSuffix
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	toSerialize["advertiserId"] = o.AdvertiserId
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableSponsoredProductsTargetCampaignAttributes struct {
	value *SponsoredProductsTargetCampaignAttributes
	isSet bool
}

func (v NullableSponsoredProductsTargetCampaignAttributes) Get() *SponsoredProductsTargetCampaignAttributes {
	return v.value
}

func (v *NullableSponsoredProductsTargetCampaignAttributes) Set(val *SponsoredProductsTargetCampaignAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetCampaignAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetCampaignAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetCampaignAttributes(val *SponsoredProductsTargetCampaignAttributes) *NullableSponsoredProductsTargetCampaignAttributes {
	return &NullableSponsoredProductsTargetCampaignAttributes{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetCampaignAttributes) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetCampaignAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
