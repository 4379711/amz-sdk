package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceLevelProductIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceLevelProductIdentifier{}

// SponsoredProductsMarketplaceLevelProductIdentifier struct for SponsoredProductsMarketplaceLevelProductIdentifier
type SponsoredProductsMarketplaceLevelProductIdentifier struct {
	Identifier  string                        `json:"identifier"`
	Marketplace *SponsoredProductsMarketplace `json:"marketplace,omitempty"`
}

type _SponsoredProductsMarketplaceLevelProductIdentifier SponsoredProductsMarketplaceLevelProductIdentifier

// NewSponsoredProductsMarketplaceLevelProductIdentifier instantiates a new SponsoredProductsMarketplaceLevelProductIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceLevelProductIdentifier(identifier string) *SponsoredProductsMarketplaceLevelProductIdentifier {
	this := SponsoredProductsMarketplaceLevelProductIdentifier{}
	this.Identifier = identifier
	return &this
}

// NewSponsoredProductsMarketplaceLevelProductIdentifierWithDefaults instantiates a new SponsoredProductsMarketplaceLevelProductIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceLevelProductIdentifierWithDefaults() *SponsoredProductsMarketplaceLevelProductIdentifier {
	this := SponsoredProductsMarketplaceLevelProductIdentifier{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) SetIdentifier(v string) {
	o.Identifier = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceLevelProductIdentifier) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

func (o SponsoredProductsMarketplaceLevelProductIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	return toSerialize, nil
}

type NullableSponsoredProductsMarketplaceLevelProductIdentifier struct {
	value *SponsoredProductsMarketplaceLevelProductIdentifier
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceLevelProductIdentifier) Get() *SponsoredProductsMarketplaceLevelProductIdentifier {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceLevelProductIdentifier) Set(val *SponsoredProductsMarketplaceLevelProductIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceLevelProductIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceLevelProductIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceLevelProductIdentifier(val *SponsoredProductsMarketplaceLevelProductIdentifier) *NullableSponsoredProductsMarketplaceLevelProductIdentifier {
	return &NullableSponsoredProductsMarketplaceLevelProductIdentifier{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceLevelProductIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceLevelProductIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
