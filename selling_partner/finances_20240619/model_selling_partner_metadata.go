package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the SellingPartnerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellingPartnerMetadata{}

// SellingPartnerMetadata Metadata describing the seller.
type SellingPartnerMetadata struct {
	// Unique seller identifier.
	SellingPartnerId *string `json:"sellingPartnerId,omitempty"`
	// Account type of transaction.
	AccountType *string `json:"accountType,omitempty"`
	// Marketplace identifier of transaction.
	MarketplaceId *string `json:"marketplaceId,omitempty"`
}

// NewSellingPartnerMetadata instantiates a new SellingPartnerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellingPartnerMetadata() *SellingPartnerMetadata {
	this := SellingPartnerMetadata{}
	return &this
}

// NewSellingPartnerMetadataWithDefaults instantiates a new SellingPartnerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellingPartnerMetadataWithDefaults() *SellingPartnerMetadata {
	this := SellingPartnerMetadata{}
	return &this
}

// GetSellingPartnerId returns the SellingPartnerId field value if set, zero value otherwise.
func (o *SellingPartnerMetadata) GetSellingPartnerId() string {
	if o == nil || IsNil(o.SellingPartnerId) {
		var ret string
		return ret
	}
	return *o.SellingPartnerId
}

// GetSellingPartnerIdOk returns a tuple with the SellingPartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingPartnerMetadata) GetSellingPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellingPartnerId) {
		return nil, false
	}
	return o.SellingPartnerId, true
}

// HasSellingPartnerId returns a boolean if a field has been set.
func (o *SellingPartnerMetadata) HasSellingPartnerId() bool {
	if o != nil && !IsNil(o.SellingPartnerId) {
		return true
	}

	return false
}

// SetSellingPartnerId gets a reference to the given string and assigns it to the SellingPartnerId field.
func (o *SellingPartnerMetadata) SetSellingPartnerId(v string) {
	o.SellingPartnerId = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *SellingPartnerMetadata) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingPartnerMetadata) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *SellingPartnerMetadata) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *SellingPartnerMetadata) SetAccountType(v string) {
	o.AccountType = &v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *SellingPartnerMetadata) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingPartnerMetadata) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *SellingPartnerMetadata) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *SellingPartnerMetadata) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

func (o SellingPartnerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SellingPartnerId) {
		toSerialize["sellingPartnerId"] = o.SellingPartnerId
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	return toSerialize, nil
}

type NullableSellingPartnerMetadata struct {
	value *SellingPartnerMetadata
	isSet bool
}

func (v NullableSellingPartnerMetadata) Get() *SellingPartnerMetadata {
	return v.value
}

func (v *NullableSellingPartnerMetadata) Set(val *SellingPartnerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSellingPartnerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSellingPartnerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellingPartnerMetadata(val *SellingPartnerMetadata) *NullableSellingPartnerMetadata {
	return &NullableSellingPartnerMetadata{value: val, isSet: true}
}

func (v NullableSellingPartnerMetadata) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSellingPartnerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
