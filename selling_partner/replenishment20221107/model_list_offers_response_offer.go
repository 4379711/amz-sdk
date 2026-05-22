package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOffersResponseOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOffersResponseOffer{}

// ListOffersResponseOffer An object which contains details about an offer.
type ListOffersResponseOffer struct {
	// The SKU. This property is only supported for sellers and not for vendors.
	Sku *string `json:"sku,omitempty"`
	// The Amazon Standard Identification Number (ASIN).
	Asin *string `json:"asin,omitempty"`
	// The marketplace identifier. The supported marketplaces for both sellers and vendors are US, CA, ES, UK, FR, IT, IN, DE and JP. The supported marketplaces for vendors only are BR, AU, MX, AE and NL. Refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids) to find the identifier for the marketplace.
	MarketplaceId             *string                    `json:"marketplaceId,omitempty"`
	Eligibility               *EligibilityStatus         `json:"eligibility,omitempty"`
	OfferProgramConfiguration *OfferProgramConfiguration `json:"offerProgramConfiguration,omitempty"`
	ProgramType               *ProgramType               `json:"programType,omitempty"`
	// A list of vendor codes associated with the offer.
	VendorCodes []string `json:"vendorCodes,omitempty"`
}

// NewListOffersResponseOffer instantiates a new ListOffersResponseOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOffersResponseOffer() *ListOffersResponseOffer {
	this := ListOffersResponseOffer{}
	return &this
}

// NewListOffersResponseOfferWithDefaults instantiates a new ListOffersResponseOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOffersResponseOfferWithDefaults() *ListOffersResponseOffer {
	this := ListOffersResponseOffer{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ListOffersResponseOffer) SetSku(v string) {
	o.Sku = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ListOffersResponseOffer) SetAsin(v string) {
	o.Asin = &v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *ListOffersResponseOffer) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetEligibility returns the Eligibility field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetEligibility() EligibilityStatus {
	if o == nil || IsNil(o.Eligibility) {
		var ret EligibilityStatus
		return ret
	}
	return *o.Eligibility
}

// GetEligibilityOk returns a tuple with the Eligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetEligibilityOk() (*EligibilityStatus, bool) {
	if o == nil || IsNil(o.Eligibility) {
		return nil, false
	}
	return o.Eligibility, true
}

// HasEligibility returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasEligibility() bool {
	if o != nil && !IsNil(o.Eligibility) {
		return true
	}

	return false
}

// SetEligibility gets a reference to the given EligibilityStatus and assigns it to the Eligibility field.
func (o *ListOffersResponseOffer) SetEligibility(v EligibilityStatus) {
	o.Eligibility = &v
}

// GetOfferProgramConfiguration returns the OfferProgramConfiguration field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetOfferProgramConfiguration() OfferProgramConfiguration {
	if o == nil || IsNil(o.OfferProgramConfiguration) {
		var ret OfferProgramConfiguration
		return ret
	}
	return *o.OfferProgramConfiguration
}

// GetOfferProgramConfigurationOk returns a tuple with the OfferProgramConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetOfferProgramConfigurationOk() (*OfferProgramConfiguration, bool) {
	if o == nil || IsNil(o.OfferProgramConfiguration) {
		return nil, false
	}
	return o.OfferProgramConfiguration, true
}

// HasOfferProgramConfiguration returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasOfferProgramConfiguration() bool {
	if o != nil && !IsNil(o.OfferProgramConfiguration) {
		return true
	}

	return false
}

// SetOfferProgramConfiguration gets a reference to the given OfferProgramConfiguration and assigns it to the OfferProgramConfiguration field.
func (o *ListOffersResponseOffer) SetOfferProgramConfiguration(v OfferProgramConfiguration) {
	o.OfferProgramConfiguration = &v
}

// GetProgramType returns the ProgramType field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetProgramType() ProgramType {
	if o == nil || IsNil(o.ProgramType) {
		var ret ProgramType
		return ret
	}
	return *o.ProgramType
}

// GetProgramTypeOk returns a tuple with the ProgramType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetProgramTypeOk() (*ProgramType, bool) {
	if o == nil || IsNil(o.ProgramType) {
		return nil, false
	}
	return o.ProgramType, true
}

// HasProgramType returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasProgramType() bool {
	if o != nil && !IsNil(o.ProgramType) {
		return true
	}

	return false
}

// SetProgramType gets a reference to the given ProgramType and assigns it to the ProgramType field.
func (o *ListOffersResponseOffer) SetProgramType(v ProgramType) {
	o.ProgramType = &v
}

// GetVendorCodes returns the VendorCodes field value if set, zero value otherwise.
func (o *ListOffersResponseOffer) GetVendorCodes() []string {
	if o == nil || IsNil(o.VendorCodes) {
		var ret []string
		return ret
	}
	return o.VendorCodes
}

// GetVendorCodesOk returns a tuple with the VendorCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponseOffer) GetVendorCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.VendorCodes) {
		return nil, false
	}
	return o.VendorCodes, true
}

// HasVendorCodes returns a boolean if a field has been set.
func (o *ListOffersResponseOffer) HasVendorCodes() bool {
	if o != nil && !IsNil(o.VendorCodes) {
		return true
	}

	return false
}

// SetVendorCodes gets a reference to the given []string and assigns it to the VendorCodes field.
func (o *ListOffersResponseOffer) SetVendorCodes(v []string) {
	o.VendorCodes = v
}

func (o ListOffersResponseOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.Eligibility) {
		toSerialize["eligibility"] = o.Eligibility
	}
	if !IsNil(o.OfferProgramConfiguration) {
		toSerialize["offerProgramConfiguration"] = o.OfferProgramConfiguration
	}
	if !IsNil(o.ProgramType) {
		toSerialize["programType"] = o.ProgramType
	}
	if !IsNil(o.VendorCodes) {
		toSerialize["vendorCodes"] = o.VendorCodes
	}
	return toSerialize, nil
}

type NullableListOffersResponseOffer struct {
	value *ListOffersResponseOffer
	isSet bool
}

func (v NullableListOffersResponseOffer) Get() *ListOffersResponseOffer {
	return v.value
}

func (v *NullableListOffersResponseOffer) Set(val *ListOffersResponseOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableListOffersResponseOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableListOffersResponseOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOffersResponseOffer(val *ListOffersResponseOffer) *NullableListOffersResponseOffer {
	return &NullableListOffersResponseOffer{value: val, isSet: true}
}

func (v NullableListOffersResponseOffer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOffersResponseOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
