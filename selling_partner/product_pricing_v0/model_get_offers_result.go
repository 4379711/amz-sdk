package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetOffersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOffersResult{}

// GetOffersResult The payload for the getListingOffers and getItemOffers operations.
type GetOffersResult struct {
	// A marketplace identifier.
	MarketplaceID string `json:"MarketplaceID"`
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`
	// The stock keeping unit (SKU) of the item.
	SKU           *string       `json:"SKU,omitempty"`
	ItemCondition ConditionType `json:"ItemCondition"`
	// The status of the operation.
	Status     string         `json:"status"`
	Identifier ItemIdentifier `json:"Identifier"`
	Summary    Summary        `json:"Summary"`
	// A list of offer details. The list is the same length as the TotalOfferCount in the Summary or 20, whichever is less.
	Offers []OfferDetail `json:"Offers"`
}

type _GetOffersResult GetOffersResult

// NewGetOffersResult instantiates a new GetOffersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOffersResult(marketplaceID string, itemCondition ConditionType, status string, identifier ItemIdentifier, summary Summary, offers []OfferDetail) *GetOffersResult {
	this := GetOffersResult{}
	this.MarketplaceID = marketplaceID
	this.ItemCondition = itemCondition
	this.Status = status
	this.Identifier = identifier
	this.Summary = summary
	this.Offers = offers
	return &this
}

// NewGetOffersResultWithDefaults instantiates a new GetOffersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOffersResultWithDefaults() *GetOffersResult {
	this := GetOffersResult{}
	return &this
}

// GetMarketplaceID returns the MarketplaceID field value
func (o *GetOffersResult) GetMarketplaceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceID
}

// GetMarketplaceIDOk returns a tuple with the MarketplaceID field value
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetMarketplaceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceID, true
}

// SetMarketplaceID sets field value
func (o *GetOffersResult) SetMarketplaceID(v string) {
	o.MarketplaceID = v
}

// GetASIN returns the ASIN field value if set, zero value otherwise.
func (o *GetOffersResult) GetASIN() string {
	if o == nil || IsNil(o.ASIN) {
		var ret string
		return ret
	}
	return *o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetASINOk() (*string, bool) {
	if o == nil || IsNil(o.ASIN) {
		return nil, false
	}
	return o.ASIN, true
}

// HasASIN returns a boolean if a field has been set.
func (o *GetOffersResult) HasASIN() bool {
	if o != nil && !IsNil(o.ASIN) {
		return true
	}

	return false
}

// SetASIN gets a reference to the given string and assigns it to the ASIN field.
func (o *GetOffersResult) SetASIN(v string) {
	o.ASIN = &v
}

// GetSKU returns the SKU field value if set, zero value otherwise.
func (o *GetOffersResult) GetSKU() string {
	if o == nil || IsNil(o.SKU) {
		var ret string
		return ret
	}
	return *o.SKU
}

// GetSKUOk returns a tuple with the SKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetSKUOk() (*string, bool) {
	if o == nil || IsNil(o.SKU) {
		return nil, false
	}
	return o.SKU, true
}

// HasSKU returns a boolean if a field has been set.
func (o *GetOffersResult) HasSKU() bool {
	if o != nil && !IsNil(o.SKU) {
		return true
	}

	return false
}

// SetSKU gets a reference to the given string and assigns it to the SKU field.
func (o *GetOffersResult) SetSKU(v string) {
	o.SKU = &v
}

// GetItemCondition returns the ItemCondition field value
func (o *GetOffersResult) GetItemCondition() ConditionType {
	if o == nil {
		var ret ConditionType
		return ret
	}

	return o.ItemCondition
}

// GetItemConditionOk returns a tuple with the ItemCondition field value
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetItemConditionOk() (*ConditionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCondition, true
}

// SetItemCondition sets field value
func (o *GetOffersResult) SetItemCondition(v ConditionType) {
	o.ItemCondition = v
}

// GetStatus returns the Status field value
func (o *GetOffersResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetOffersResult) SetStatus(v string) {
	o.Status = v
}

// GetIdentifier returns the Identifier field value
func (o *GetOffersResult) GetIdentifier() ItemIdentifier {
	if o == nil {
		var ret ItemIdentifier
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetIdentifierOk() (*ItemIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GetOffersResult) SetIdentifier(v ItemIdentifier) {
	o.Identifier = v
}

// GetSummary returns the Summary field value
func (o *GetOffersResult) GetSummary() Summary {
	if o == nil {
		var ret Summary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetSummaryOk() (*Summary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *GetOffersResult) SetSummary(v Summary) {
	o.Summary = v
}

// GetOffers returns the Offers field value
func (o *GetOffersResult) GetOffers() []OfferDetail {
	if o == nil {
		var ret []OfferDetail
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *GetOffersResult) GetOffersOk() ([]OfferDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *GetOffersResult) SetOffers(v []OfferDetail) {
	o.Offers = v
}

func (o GetOffersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceID"] = o.MarketplaceID
	if !IsNil(o.ASIN) {
		toSerialize["ASIN"] = o.ASIN
	}
	if !IsNil(o.SKU) {
		toSerialize["SKU"] = o.SKU
	}
	toSerialize["ItemCondition"] = o.ItemCondition
	toSerialize["status"] = o.Status
	toSerialize["Identifier"] = o.Identifier
	toSerialize["Summary"] = o.Summary
	toSerialize["Offers"] = o.Offers
	return toSerialize, nil
}

type NullableGetOffersResult struct {
	value *GetOffersResult
	isSet bool
}

func (v NullableGetOffersResult) Get() *GetOffersResult {
	return v.value
}

func (v *NullableGetOffersResult) Set(val *GetOffersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOffersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOffersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOffersResult(val *GetOffersResult) *NullableGetOffersResult {
	return &NullableGetOffersResult{value: val, isSet: true}
}

func (v NullableGetOffersResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetOffersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
