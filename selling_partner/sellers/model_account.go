package sellers

import (
	"github.com/bytedance/sonic"
)

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account The response schema for the `getAccount` operation.
type Account struct {
	// List of marketplace participations.
	MarketplaceParticipationList []MarketplaceParticipation `json:"marketplaceParticipationList"`
	// The type of business registered for the seller account.
	BusinessType string `json:"businessType"`
	// The selling plan details.
	SellingPlan    string          `json:"sellingPlan"`
	Business       *Business       `json:"business,omitempty"`
	PrimaryContact *PrimaryContact `json:"primaryContact,omitempty"`
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(marketplaceParticipationList []MarketplaceParticipation, businessType string, sellingPlan string) *Account {
	this := Account{}
	this.MarketplaceParticipationList = marketplaceParticipationList
	this.BusinessType = businessType
	this.SellingPlan = sellingPlan
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetMarketplaceParticipationList returns the MarketplaceParticipationList field value
func (o *Account) GetMarketplaceParticipationList() []MarketplaceParticipation {
	if o == nil {
		var ret []MarketplaceParticipation
		return ret
	}

	return o.MarketplaceParticipationList
}

// GetMarketplaceParticipationListOk returns a tuple with the MarketplaceParticipationList field value
// and a boolean to check if the value has been set.
func (o *Account) GetMarketplaceParticipationListOk() ([]MarketplaceParticipation, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceParticipationList, true
}

// SetMarketplaceParticipationList sets field value
func (o *Account) SetMarketplaceParticipationList(v []MarketplaceParticipation) {
	o.MarketplaceParticipationList = v
}

// GetBusinessType returns the BusinessType field value
func (o *Account) GetBusinessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessType
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value
// and a boolean to check if the value has been set.
func (o *Account) GetBusinessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessType, true
}

// SetBusinessType sets field value
func (o *Account) SetBusinessType(v string) {
	o.BusinessType = v
}

// GetSellingPlan returns the SellingPlan field value
func (o *Account) GetSellingPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellingPlan
}

// GetSellingPlanOk returns a tuple with the SellingPlan field value
// and a boolean to check if the value has been set.
func (o *Account) GetSellingPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingPlan, true
}

// SetSellingPlan sets field value
func (o *Account) SetSellingPlan(v string) {
	o.SellingPlan = v
}

// GetBusiness returns the Business field value if set, zero value otherwise.
func (o *Account) GetBusiness() Business {
	if o == nil || IsNil(o.Business) {
		var ret Business
		return ret
	}
	return *o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBusinessOk() (*Business, bool) {
	if o == nil || IsNil(o.Business) {
		return nil, false
	}
	return o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *Account) HasBusiness() bool {
	if o != nil && !IsNil(o.Business) {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given Business and assigns it to the Business field.
func (o *Account) SetBusiness(v Business) {
	o.Business = &v
}

// GetPrimaryContact returns the PrimaryContact field value if set, zero value otherwise.
func (o *Account) GetPrimaryContact() PrimaryContact {
	if o == nil || IsNil(o.PrimaryContact) {
		var ret PrimaryContact
		return ret
	}
	return *o.PrimaryContact
}

// GetPrimaryContactOk returns a tuple with the PrimaryContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetPrimaryContactOk() (*PrimaryContact, bool) {
	if o == nil || IsNil(o.PrimaryContact) {
		return nil, false
	}
	return o.PrimaryContact, true
}

// HasPrimaryContact returns a boolean if a field has been set.
func (o *Account) HasPrimaryContact() bool {
	if o != nil && !IsNil(o.PrimaryContact) {
		return true
	}

	return false
}

// SetPrimaryContact gets a reference to the given PrimaryContact and assigns it to the PrimaryContact field.
func (o *Account) SetPrimaryContact(v PrimaryContact) {
	o.PrimaryContact = &v
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceParticipationList"] = o.MarketplaceParticipationList
	toSerialize["businessType"] = o.BusinessType
	toSerialize["sellingPlan"] = o.SellingPlan
	if !IsNil(o.Business) {
		toSerialize["business"] = o.Business
	}
	if !IsNil(o.PrimaryContact) {
		toSerialize["primaryContact"] = o.PrimaryContact
	}
	return toSerialize, nil
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
