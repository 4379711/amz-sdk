package profiles_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountInfo{}

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	// The identifier of the marketplace to which the account is associated.
	MarketplaceStringId *string `json:"marketplaceStringId,omitempty"`
	// Identifier for sellers and vendors. Note that this value is not unique and may be the same across marketplace.
	Id   *string      `json:"id,omitempty"`
	Type *AccountType `json:"type,omitempty"`
	// Account name.
	Name *string `json:"name,omitempty"`
	// The account subtype.
	SubType *string `json:"subType,omitempty"`
	// Only present for Vendors, this returns whether the Advertiser has set up a valid payment method or not.
	ValidPaymentMethod *bool `json:"validPaymentMethod,omitempty"`
}

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetMarketplaceStringId returns the MarketplaceStringId field value if set, zero value otherwise.
func (o *AccountInfo) GetMarketplaceStringId() string {
	if o == nil || IsNil(o.MarketplaceStringId) {
		var ret string
		return ret
	}
	return *o.MarketplaceStringId
}

// GetMarketplaceStringIdOk returns a tuple with the MarketplaceStringId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetMarketplaceStringIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceStringId) {
		return nil, false
	}
	return o.MarketplaceStringId, true
}

// HasMarketplaceStringId returns a boolean if a field has been set.
func (o *AccountInfo) HasMarketplaceStringId() bool {
	if o != nil && !IsNil(o.MarketplaceStringId) {
		return true
	}

	return false
}

// SetMarketplaceStringId gets a reference to the given string and assigns it to the MarketplaceStringId field.
func (o *AccountInfo) SetMarketplaceStringId(v string) {
	o.MarketplaceStringId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountInfo) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountInfo) GetType() AccountType {
	if o == nil || IsNil(o.Type) {
		var ret AccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AccountType and assigns it to the Type field.
func (o *AccountInfo) SetType(v AccountType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountInfo) SetName(v string) {
	o.Name = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *AccountInfo) GetSubType() string {
	if o == nil || IsNil(o.SubType) {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *AccountInfo) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *AccountInfo) SetSubType(v string) {
	o.SubType = &v
}

// GetValidPaymentMethod returns the ValidPaymentMethod field value if set, zero value otherwise.
func (o *AccountInfo) GetValidPaymentMethod() bool {
	if o == nil || IsNil(o.ValidPaymentMethod) {
		var ret bool
		return ret
	}
	return *o.ValidPaymentMethod
}

// GetValidPaymentMethodOk returns a tuple with the ValidPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetValidPaymentMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidPaymentMethod) {
		return nil, false
	}
	return o.ValidPaymentMethod, true
}

// HasValidPaymentMethod returns a boolean if a field has been set.
func (o *AccountInfo) HasValidPaymentMethod() bool {
	if o != nil && !IsNil(o.ValidPaymentMethod) {
		return true
	}

	return false
}

// SetValidPaymentMethod gets a reference to the given bool and assigns it to the ValidPaymentMethod field.
func (o *AccountInfo) SetValidPaymentMethod(v bool) {
	o.ValidPaymentMethod = &v
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceStringId) {
		toSerialize["marketplaceStringId"] = o.MarketplaceStringId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SubType) {
		toSerialize["subType"] = o.SubType
	}
	if !IsNil(o.ValidPaymentMethod) {
		toSerialize["validPaymentMethod"] = o.ValidPaymentMethod
	}
	return toSerialize, nil
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
