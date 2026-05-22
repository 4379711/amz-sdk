package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateScheduledPackagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScheduledPackagesRequest{}

// CreateScheduledPackagesRequest The request body for the POST /easyShip/2022-03-23/packages/bulk API.
type CreateScheduledPackagesRequest struct {
	// A string of up to 255 characters.
	MarketplaceId string `json:"marketplaceId"`
	// An array allowing users to specify orders to be scheduled.
	OrderScheduleDetailsList []OrderScheduleDetails `json:"orderScheduleDetailsList"`
	LabelFormat              LabelFormat            `json:"labelFormat"`
}

type _CreateScheduledPackagesRequest CreateScheduledPackagesRequest

// NewCreateScheduledPackagesRequest instantiates a new CreateScheduledPackagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScheduledPackagesRequest(marketplaceId string, orderScheduleDetailsList []OrderScheduleDetails, labelFormat LabelFormat) *CreateScheduledPackagesRequest {
	this := CreateScheduledPackagesRequest{}
	this.MarketplaceId = marketplaceId
	this.OrderScheduleDetailsList = orderScheduleDetailsList
	this.LabelFormat = labelFormat
	return &this
}

// NewCreateScheduledPackagesRequestWithDefaults instantiates a new CreateScheduledPackagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScheduledPackagesRequestWithDefaults() *CreateScheduledPackagesRequest {
	this := CreateScheduledPackagesRequest{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *CreateScheduledPackagesRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackagesRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *CreateScheduledPackagesRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetOrderScheduleDetailsList returns the OrderScheduleDetailsList field value
func (o *CreateScheduledPackagesRequest) GetOrderScheduleDetailsList() []OrderScheduleDetails {
	if o == nil {
		var ret []OrderScheduleDetails
		return ret
	}

	return o.OrderScheduleDetailsList
}

// GetOrderScheduleDetailsListOk returns a tuple with the OrderScheduleDetailsList field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackagesRequest) GetOrderScheduleDetailsListOk() ([]OrderScheduleDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderScheduleDetailsList, true
}

// SetOrderScheduleDetailsList sets field value
func (o *CreateScheduledPackagesRequest) SetOrderScheduleDetailsList(v []OrderScheduleDetails) {
	o.OrderScheduleDetailsList = v
}

// GetLabelFormat returns the LabelFormat field value
func (o *CreateScheduledPackagesRequest) GetLabelFormat() LabelFormat {
	if o == nil {
		var ret LabelFormat
		return ret
	}

	return o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledPackagesRequest) GetLabelFormatOk() (*LabelFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelFormat, true
}

// SetLabelFormat sets field value
func (o *CreateScheduledPackagesRequest) SetLabelFormat(v LabelFormat) {
	o.LabelFormat = v
}

func (o CreateScheduledPackagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["orderScheduleDetailsList"] = o.OrderScheduleDetailsList
	toSerialize["labelFormat"] = o.LabelFormat
	return toSerialize, nil
}

type NullableCreateScheduledPackagesRequest struct {
	value *CreateScheduledPackagesRequest
	isSet bool
}

func (v NullableCreateScheduledPackagesRequest) Get() *CreateScheduledPackagesRequest {
	return v.value
}

func (v *NullableCreateScheduledPackagesRequest) Set(val *CreateScheduledPackagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScheduledPackagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScheduledPackagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScheduledPackagesRequest(val *CreateScheduledPackagesRequest) *NullableCreateScheduledPackagesRequest {
	return &NullableCreateScheduledPackagesRequest{value: val, isSet: true}
}

func (v NullableCreateScheduledPackagesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateScheduledPackagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
