package transfers_20240601

import (
	"github.com/bytedance/sonic"
)

// checks if the InitiatePayoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiatePayoutRequest{}

// InitiatePayoutRequest The request schema for the `initiatePayout` operation.
type InitiatePayoutRequest struct {
	// The identifier of the Amazon marketplace. For the list of all marketplace IDs, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// The account type in the selected marketplace for which a payout must be initiated. For supported EU marketplaces, the only account type is `Standard Orders`.
	AccountType string `json:"accountType"`
}

type _InitiatePayoutRequest InitiatePayoutRequest

// NewInitiatePayoutRequest instantiates a new InitiatePayoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiatePayoutRequest(marketplaceId string, accountType string) *InitiatePayoutRequest {
	this := InitiatePayoutRequest{}
	this.MarketplaceId = marketplaceId
	this.AccountType = accountType
	return &this
}

// NewInitiatePayoutRequestWithDefaults instantiates a new InitiatePayoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiatePayoutRequestWithDefaults() *InitiatePayoutRequest {
	this := InitiatePayoutRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *InitiatePayoutRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *InitiatePayoutRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *InitiatePayoutRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetAccountType returns the AccountType field value
func (o *InitiatePayoutRequest) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *InitiatePayoutRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *InitiatePayoutRequest) SetAccountType(v string) {
	o.AccountType = v
}

func (o InitiatePayoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["accountType"] = o.AccountType
	return toSerialize, nil
}

type NullableInitiatePayoutRequest struct {
	value *InitiatePayoutRequest
	isSet bool
}

func (v NullableInitiatePayoutRequest) Get() *InitiatePayoutRequest {
	return v.value
}

func (v *NullableInitiatePayoutRequest) Set(val *InitiatePayoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiatePayoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiatePayoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiatePayoutRequest(val *InitiatePayoutRequest) *NullableInitiatePayoutRequest {
	return &NullableInitiatePayoutRequest{value: val, isSet: true}
}

func (v NullableInitiatePayoutRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInitiatePayoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
