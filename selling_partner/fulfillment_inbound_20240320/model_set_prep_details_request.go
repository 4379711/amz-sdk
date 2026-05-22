package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SetPrepDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPrepDetailsRequest{}

// SetPrepDetailsRequest The `setPrepDetails` request.
type SetPrepDetailsRequest struct {
	// The marketplace ID. For a list of possible values, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// A list of MSKUs and related prep details.
	MskuPrepDetails []MskuPrepDetailInput `json:"mskuPrepDetails"`
}

type _SetPrepDetailsRequest SetPrepDetailsRequest

// NewSetPrepDetailsRequest instantiates a new SetPrepDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPrepDetailsRequest(marketplaceId string, mskuPrepDetails []MskuPrepDetailInput) *SetPrepDetailsRequest {
	this := SetPrepDetailsRequest{}
	this.MarketplaceId = marketplaceId
	this.MskuPrepDetails = mskuPrepDetails
	return &this
}

// NewSetPrepDetailsRequestWithDefaults instantiates a new SetPrepDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPrepDetailsRequestWithDefaults() *SetPrepDetailsRequest {
	this := SetPrepDetailsRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *SetPrepDetailsRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *SetPrepDetailsRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *SetPrepDetailsRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetMskuPrepDetails returns the MskuPrepDetails field value
func (o *SetPrepDetailsRequest) GetMskuPrepDetails() []MskuPrepDetailInput {
	if o == nil {
		var ret []MskuPrepDetailInput
		return ret
	}

	return o.MskuPrepDetails
}

// GetMskuPrepDetailsOk returns a tuple with the MskuPrepDetails field value
// and a boolean to check if the value has been set.
func (o *SetPrepDetailsRequest) GetMskuPrepDetailsOk() ([]MskuPrepDetailInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.MskuPrepDetails, true
}

// SetMskuPrepDetails sets field value
func (o *SetPrepDetailsRequest) SetMskuPrepDetails(v []MskuPrepDetailInput) {
	o.MskuPrepDetails = v
}

func (o SetPrepDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["mskuPrepDetails"] = o.MskuPrepDetails
	return toSerialize, nil
}

type NullableSetPrepDetailsRequest struct {
	value *SetPrepDetailsRequest
	isSet bool
}

func (v NullableSetPrepDetailsRequest) Get() *SetPrepDetailsRequest {
	return v.value
}

func (v *NullableSetPrepDetailsRequest) Set(val *SetPrepDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPrepDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPrepDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPrepDetailsRequest(val *SetPrepDetailsRequest) *NullableSetPrepDetailsRequest {
	return &NullableSetPrepDetailsRequest{value: val, isSet: true}
}

func (v NullableSetPrepDetailsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSetPrepDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
