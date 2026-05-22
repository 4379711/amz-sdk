package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the BuyerTaxInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyerTaxInformation{}

// BuyerTaxInformation Contains the business invoice tax information. Available only in the TR marketplace.
type BuyerTaxInformation struct {
	// Business buyer's company legal name.
	BuyerLegalCompanyName *string `json:"BuyerLegalCompanyName,omitempty"`
	// Business buyer's address.
	BuyerBusinessAddress *string `json:"BuyerBusinessAddress,omitempty"`
	// Business buyer's tax registration ID.
	BuyerTaxRegistrationId *string `json:"BuyerTaxRegistrationId,omitempty"`
	// Business buyer's tax office.
	BuyerTaxOffice *string `json:"BuyerTaxOffice,omitempty"`
}

// NewBuyerTaxInformation instantiates a new BuyerTaxInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyerTaxInformation() *BuyerTaxInformation {
	this := BuyerTaxInformation{}
	return &this
}

// NewBuyerTaxInformationWithDefaults instantiates a new BuyerTaxInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerTaxInformationWithDefaults() *BuyerTaxInformation {
	this := BuyerTaxInformation{}
	return &this
}

// GetBuyerLegalCompanyName returns the BuyerLegalCompanyName field value if set, zero value otherwise.
func (o *BuyerTaxInformation) GetBuyerLegalCompanyName() string {
	if o == nil || IsNil(o.BuyerLegalCompanyName) {
		var ret string
		return ret
	}
	return *o.BuyerLegalCompanyName
}

// GetBuyerLegalCompanyNameOk returns a tuple with the BuyerLegalCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxInformation) GetBuyerLegalCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerLegalCompanyName) {
		return nil, false
	}
	return o.BuyerLegalCompanyName, true
}

// HasBuyerLegalCompanyName returns a boolean if a field has been set.
func (o *BuyerTaxInformation) HasBuyerLegalCompanyName() bool {
	if o != nil && !IsNil(o.BuyerLegalCompanyName) {
		return true
	}

	return false
}

// SetBuyerLegalCompanyName gets a reference to the given string and assigns it to the BuyerLegalCompanyName field.
func (o *BuyerTaxInformation) SetBuyerLegalCompanyName(v string) {
	o.BuyerLegalCompanyName = &v
}

// GetBuyerBusinessAddress returns the BuyerBusinessAddress field value if set, zero value otherwise.
func (o *BuyerTaxInformation) GetBuyerBusinessAddress() string {
	if o == nil || IsNil(o.BuyerBusinessAddress) {
		var ret string
		return ret
	}
	return *o.BuyerBusinessAddress
}

// GetBuyerBusinessAddressOk returns a tuple with the BuyerBusinessAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxInformation) GetBuyerBusinessAddressOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerBusinessAddress) {
		return nil, false
	}
	return o.BuyerBusinessAddress, true
}

// HasBuyerBusinessAddress returns a boolean if a field has been set.
func (o *BuyerTaxInformation) HasBuyerBusinessAddress() bool {
	if o != nil && !IsNil(o.BuyerBusinessAddress) {
		return true
	}

	return false
}

// SetBuyerBusinessAddress gets a reference to the given string and assigns it to the BuyerBusinessAddress field.
func (o *BuyerTaxInformation) SetBuyerBusinessAddress(v string) {
	o.BuyerBusinessAddress = &v
}

// GetBuyerTaxRegistrationId returns the BuyerTaxRegistrationId field value if set, zero value otherwise.
func (o *BuyerTaxInformation) GetBuyerTaxRegistrationId() string {
	if o == nil || IsNil(o.BuyerTaxRegistrationId) {
		var ret string
		return ret
	}
	return *o.BuyerTaxRegistrationId
}

// GetBuyerTaxRegistrationIdOk returns a tuple with the BuyerTaxRegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxInformation) GetBuyerTaxRegistrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerTaxRegistrationId) {
		return nil, false
	}
	return o.BuyerTaxRegistrationId, true
}

// HasBuyerTaxRegistrationId returns a boolean if a field has been set.
func (o *BuyerTaxInformation) HasBuyerTaxRegistrationId() bool {
	if o != nil && !IsNil(o.BuyerTaxRegistrationId) {
		return true
	}

	return false
}

// SetBuyerTaxRegistrationId gets a reference to the given string and assigns it to the BuyerTaxRegistrationId field.
func (o *BuyerTaxInformation) SetBuyerTaxRegistrationId(v string) {
	o.BuyerTaxRegistrationId = &v
}

// GetBuyerTaxOffice returns the BuyerTaxOffice field value if set, zero value otherwise.
func (o *BuyerTaxInformation) GetBuyerTaxOffice() string {
	if o == nil || IsNil(o.BuyerTaxOffice) {
		var ret string
		return ret
	}
	return *o.BuyerTaxOffice
}

// GetBuyerTaxOfficeOk returns a tuple with the BuyerTaxOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerTaxInformation) GetBuyerTaxOfficeOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerTaxOffice) {
		return nil, false
	}
	return o.BuyerTaxOffice, true
}

// HasBuyerTaxOffice returns a boolean if a field has been set.
func (o *BuyerTaxInformation) HasBuyerTaxOffice() bool {
	if o != nil && !IsNil(o.BuyerTaxOffice) {
		return true
	}

	return false
}

// SetBuyerTaxOffice gets a reference to the given string and assigns it to the BuyerTaxOffice field.
func (o *BuyerTaxInformation) SetBuyerTaxOffice(v string) {
	o.BuyerTaxOffice = &v
}

func (o BuyerTaxInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerLegalCompanyName) {
		toSerialize["BuyerLegalCompanyName"] = o.BuyerLegalCompanyName
	}
	if !IsNil(o.BuyerBusinessAddress) {
		toSerialize["BuyerBusinessAddress"] = o.BuyerBusinessAddress
	}
	if !IsNil(o.BuyerTaxRegistrationId) {
		toSerialize["BuyerTaxRegistrationId"] = o.BuyerTaxRegistrationId
	}
	if !IsNil(o.BuyerTaxOffice) {
		toSerialize["BuyerTaxOffice"] = o.BuyerTaxOffice
	}
	return toSerialize, nil
}

type NullableBuyerTaxInformation struct {
	value *BuyerTaxInformation
	isSet bool
}

func (v NullableBuyerTaxInformation) Get() *BuyerTaxInformation {
	return v.value
}

func (v *NullableBuyerTaxInformation) Set(val *BuyerTaxInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyerTaxInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyerTaxInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyerTaxInformation(val *BuyerTaxInformation) *NullableBuyerTaxInformation {
	return &NullableBuyerTaxInformation{value: val, isSet: true}
}

func (v NullableBuyerTaxInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBuyerTaxInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
