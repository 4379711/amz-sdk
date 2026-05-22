package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferEligibilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferEligibilityRequest{}

// OfferEligibilityRequest A offer advertising eligibility request object
type OfferEligibilityRequest struct {
	// A List of offerIds
	Offers []Offers `json:"offers"`
	// This field is used to specify eligibility checks for a given adProduct. Eligibility checks may vary across adProduct.
	AdProduct string `json:"adProduct"`
	// Id of a retailer. You can retrieve this from the /adsAccounts API.
	RetailerId string `json:"retailerId"`
	// Set to the locale string in the table below to specify the language in which the response is returned
	Locale string `json:"locale"`
}

type _OfferEligibilityRequest OfferEligibilityRequest

// NewOfferEligibilityRequest instantiates a new OfferEligibilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferEligibilityRequest(offers []Offers, adProduct string, retailerId string, locale string) *OfferEligibilityRequest {
	this := OfferEligibilityRequest{}
	this.Offers = offers
	this.AdProduct = adProduct
	this.RetailerId = retailerId
	this.Locale = locale
	return &this
}

// NewOfferEligibilityRequestWithDefaults instantiates a new OfferEligibilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferEligibilityRequestWithDefaults() *OfferEligibilityRequest {
	this := OfferEligibilityRequest{}
	var adProduct string = "SPONSORED_PRODUCTS"
	this.AdProduct = adProduct
	return &this
}

// GetOffers returns the Offers field value
func (o *OfferEligibilityRequest) GetOffers() []Offers {
	if o == nil {
		var ret []Offers
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *OfferEligibilityRequest) GetOffersOk() ([]Offers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *OfferEligibilityRequest) SetOffers(v []Offers) {
	o.Offers = v
}

// GetAdProduct returns the AdProduct field value
func (o *OfferEligibilityRequest) GetAdProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdProduct
}

// GetAdProductOk returns a tuple with the AdProduct field value
// and a boolean to check if the value has been set.
func (o *OfferEligibilityRequest) GetAdProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdProduct, true
}

// SetAdProduct sets field value
func (o *OfferEligibilityRequest) SetAdProduct(v string) {
	o.AdProduct = v
}

// GetRetailerId returns the RetailerId field value
func (o *OfferEligibilityRequest) GetRetailerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetailerId
}

// GetRetailerIdOk returns a tuple with the RetailerId field value
// and a boolean to check if the value has been set.
func (o *OfferEligibilityRequest) GetRetailerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetailerId, true
}

// SetRetailerId sets field value
func (o *OfferEligibilityRequest) SetRetailerId(v string) {
	o.RetailerId = v
}

// GetLocale returns the Locale field value
func (o *OfferEligibilityRequest) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *OfferEligibilityRequest) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *OfferEligibilityRequest) SetLocale(v string) {
	o.Locale = v
}

func (o OfferEligibilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offers"] = o.Offers
	toSerialize["adProduct"] = o.AdProduct
	toSerialize["retailerId"] = o.RetailerId
	toSerialize["locale"] = o.Locale
	return toSerialize, nil
}

type NullableOfferEligibilityRequest struct {
	value *OfferEligibilityRequest
	isSet bool
}

func (v NullableOfferEligibilityRequest) Get() *OfferEligibilityRequest {
	return v.value
}

func (v *NullableOfferEligibilityRequest) Set(val *OfferEligibilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferEligibilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferEligibilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferEligibilityRequest(val *OfferEligibilityRequest) *NullableOfferEligibilityRequest {
	return &NullableOfferEligibilityRequest{value: val, isSet: true}
}

func (v NullableOfferEligibilityRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferEligibilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
