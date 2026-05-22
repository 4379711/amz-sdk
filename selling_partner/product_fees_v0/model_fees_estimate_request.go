package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the FeesEstimateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesEstimateRequest{}

// FeesEstimateRequest A product, marketplace, and proposed price used to request estimated fees.
type FeesEstimateRequest struct {
	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`
	// When true, the offer is fulfilled by Amazon.
	IsAmazonFulfilled   *bool               `json:"IsAmazonFulfilled,omitempty"`
	PriceToEstimateFees PriceToEstimateFees `json:"PriceToEstimateFees"`
	// A unique identifier provided by the caller to track this request.
	Identifier                 string                      `json:"Identifier"`
	OptionalFulfillmentProgram *OptionalFulfillmentProgram `json:"OptionalFulfillmentProgram,omitempty"`
}

type _FeesEstimateRequest FeesEstimateRequest

// NewFeesEstimateRequest instantiates a new FeesEstimateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesEstimateRequest(marketplaceId string, priceToEstimateFees PriceToEstimateFees, identifier string) *FeesEstimateRequest {
	this := FeesEstimateRequest{}
	this.MarketplaceId = marketplaceId
	this.PriceToEstimateFees = priceToEstimateFees
	this.Identifier = identifier
	return &this
}

// NewFeesEstimateRequestWithDefaults instantiates a new FeesEstimateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesEstimateRequestWithDefaults() *FeesEstimateRequest {
	this := FeesEstimateRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *FeesEstimateRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *FeesEstimateRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetIsAmazonFulfilled returns the IsAmazonFulfilled field value if set, zero value otherwise.
func (o *FeesEstimateRequest) GetIsAmazonFulfilled() bool {
	if o == nil || IsNil(o.IsAmazonFulfilled) {
		var ret bool
		return ret
	}
	return *o.IsAmazonFulfilled
}

// GetIsAmazonFulfilledOk returns a tuple with the IsAmazonFulfilled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesEstimateRequest) GetIsAmazonFulfilledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmazonFulfilled) {
		return nil, false
	}
	return o.IsAmazonFulfilled, true
}

// HasIsAmazonFulfilled returns a boolean if a field has been set.
func (o *FeesEstimateRequest) HasIsAmazonFulfilled() bool {
	if o != nil && !IsNil(o.IsAmazonFulfilled) {
		return true
	}

	return false
}

// SetIsAmazonFulfilled gets a reference to the given bool and assigns it to the IsAmazonFulfilled field.
func (o *FeesEstimateRequest) SetIsAmazonFulfilled(v bool) {
	o.IsAmazonFulfilled = &v
}

// GetPriceToEstimateFees returns the PriceToEstimateFees field value
func (o *FeesEstimateRequest) GetPriceToEstimateFees() PriceToEstimateFees {
	if o == nil {
		var ret PriceToEstimateFees
		return ret
	}

	return o.PriceToEstimateFees
}

// GetPriceToEstimateFeesOk returns a tuple with the PriceToEstimateFees field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateRequest) GetPriceToEstimateFeesOk() (*PriceToEstimateFees, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceToEstimateFees, true
}

// SetPriceToEstimateFees sets field value
func (o *FeesEstimateRequest) SetPriceToEstimateFees(v PriceToEstimateFees) {
	o.PriceToEstimateFees = v
}

// GetIdentifier returns the Identifier field value
func (o *FeesEstimateRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *FeesEstimateRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *FeesEstimateRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetOptionalFulfillmentProgram returns the OptionalFulfillmentProgram field value if set, zero value otherwise.
func (o *FeesEstimateRequest) GetOptionalFulfillmentProgram() OptionalFulfillmentProgram {
	if o == nil || IsNil(o.OptionalFulfillmentProgram) {
		var ret OptionalFulfillmentProgram
		return ret
	}
	return *o.OptionalFulfillmentProgram
}

// GetOptionalFulfillmentProgramOk returns a tuple with the OptionalFulfillmentProgram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesEstimateRequest) GetOptionalFulfillmentProgramOk() (*OptionalFulfillmentProgram, bool) {
	if o == nil || IsNil(o.OptionalFulfillmentProgram) {
		return nil, false
	}
	return o.OptionalFulfillmentProgram, true
}

// HasOptionalFulfillmentProgram returns a boolean if a field has been set.
func (o *FeesEstimateRequest) HasOptionalFulfillmentProgram() bool {
	if o != nil && !IsNil(o.OptionalFulfillmentProgram) {
		return true
	}

	return false
}

// SetOptionalFulfillmentProgram gets a reference to the given OptionalFulfillmentProgram and assigns it to the OptionalFulfillmentProgram field.
func (o *FeesEstimateRequest) SetOptionalFulfillmentProgram(v OptionalFulfillmentProgram) {
	o.OptionalFulfillmentProgram = &v
}

func (o FeesEstimateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceId"] = o.MarketplaceId
	if !IsNil(o.IsAmazonFulfilled) {
		toSerialize["IsAmazonFulfilled"] = o.IsAmazonFulfilled
	}
	toSerialize["PriceToEstimateFees"] = o.PriceToEstimateFees
	toSerialize["Identifier"] = o.Identifier
	if !IsNil(o.OptionalFulfillmentProgram) {
		toSerialize["OptionalFulfillmentProgram"] = o.OptionalFulfillmentProgram
	}
	return toSerialize, nil
}

type NullableFeesEstimateRequest struct {
	value *FeesEstimateRequest
	isSet bool
}

func (v NullableFeesEstimateRequest) Get() *FeesEstimateRequest {
	return v.value
}

func (v *NullableFeesEstimateRequest) Set(val *FeesEstimateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesEstimateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesEstimateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesEstimateRequest(val *FeesEstimateRequest) *NullableFeesEstimateRequest {
	return &NullableFeesEstimateRequest{value: val, isSet: true}
}

func (v NullableFeesEstimateRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeesEstimateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
