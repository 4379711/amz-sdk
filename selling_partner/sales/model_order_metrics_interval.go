package sales

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderMetricsInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderMetricsInterval{}

// OrderMetricsInterval Contains order metrics.
type OrderMetricsInterval struct {
	// The interval of time based on requested granularity (ex. Hour, Day, etc.) If this is the first or the last interval from the list, it might contain incomplete data if the requested interval doesn't align with the requested granularity (ex. request interval 2018-09-01T02:00:00Z--2018-09-04T19:00:00Z and granularity day will result in Sept 1st UTC day and Sept 4th UTC days having partial data).
	Interval string `json:"interval"`
	// The number of units in orders based on the specified filters.
	UnitCount int32 `json:"unitCount"`
	// The number of order items based on the specified filters.
	OrderItemCount int32 `json:"orderItemCount"`
	// The number of orders based on the specified filters.
	OrderCount       int32 `json:"orderCount"`
	AverageUnitPrice Money `json:"averageUnitPrice"`
	TotalSales       Money `json:"totalSales"`
}

type _OrderMetricsInterval OrderMetricsInterval

// NewOrderMetricsInterval instantiates a new OrderMetricsInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderMetricsInterval(interval string, unitCount int32, orderItemCount int32, orderCount int32, averageUnitPrice Money, totalSales Money) *OrderMetricsInterval {
	this := OrderMetricsInterval{}
	this.Interval = interval
	this.UnitCount = unitCount
	this.OrderItemCount = orderItemCount
	this.OrderCount = orderCount
	this.AverageUnitPrice = averageUnitPrice
	this.TotalSales = totalSales
	return &this
}

// NewOrderMetricsIntervalWithDefaults instantiates a new OrderMetricsInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderMetricsIntervalWithDefaults() *OrderMetricsInterval {
	this := OrderMetricsInterval{}
	return &this
}

// GetInterval returns the Interval field value
func (o *OrderMetricsInterval) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *OrderMetricsInterval) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *OrderMetricsInterval) SetInterval(v string) {
	o.Interval = v
}

// GetUnitCount returns the UnitCount field value
func (o *OrderMetricsInterval) GetUnitCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitCount
}

// GetUnitCountOk returns a tuple with the UnitCount field value
// and a boolean to check if the value has been set.
func (o *OrderMetricsInterval) GetUnitCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitCount, true
}

// SetUnitCount sets field value
func (o *OrderMetricsInterval) SetUnitCount(v int32) {
	o.UnitCount = v
}

// GetOrderItemCount returns the OrderItemCount field value
func (o *OrderMetricsInterval) GetOrderItemCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderItemCount
}

// GetOrderItemCountOk returns a tuple with the OrderItemCount field value
// and a boolean to check if the value has been set.
func (o *OrderMetricsInterval) GetOrderItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemCount, true
}

// SetOrderItemCount sets field value
func (o *OrderMetricsInterval) SetOrderItemCount(v int32) {
	o.OrderItemCount = v
}

// GetOrderCount returns the OrderCount field value
func (o *OrderMetricsInterval) GetOrderCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderCount
}

// GetOrderCountOk returns a tuple with the OrderCount field value
// and a boolean to check if the value has been set.
func (o *OrderMetricsInterval) GetOrderCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderCount, true
}

// SetOrderCount sets field value
func (o *OrderMetricsInterval) SetOrderCount(v int32) {
	o.OrderCount = v
}

// GetAverageUnitPrice returns the AverageUnitPrice field value
func (o *OrderMetricsInterval) GetAverageUnitPrice() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.AverageUnitPrice
}

// GetAverageUnitPriceOk returns a tuple with the AverageUnitPrice field value
// and a boolean to check if the value has been set.
func (o *OrderMetricsInterval) GetAverageUnitPriceOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageUnitPrice, true
}

// SetAverageUnitPrice sets field value
func (o *OrderMetricsInterval) SetAverageUnitPrice(v Money) {
	o.AverageUnitPrice = v
}

// GetTotalSales returns the TotalSales field value
func (o *OrderMetricsInterval) GetTotalSales() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.TotalSales
}

// GetTotalSalesOk returns a tuple with the TotalSales field value
// and a boolean to check if the value has been set.
func (o *OrderMetricsInterval) GetTotalSalesOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSales, true
}

// SetTotalSales sets field value
func (o *OrderMetricsInterval) SetTotalSales(v Money) {
	o.TotalSales = v
}

func (o OrderMetricsInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	toSerialize["unitCount"] = o.UnitCount
	toSerialize["orderItemCount"] = o.OrderItemCount
	toSerialize["orderCount"] = o.OrderCount
	toSerialize["averageUnitPrice"] = o.AverageUnitPrice
	toSerialize["totalSales"] = o.TotalSales
	return toSerialize, nil
}

type NullableOrderMetricsInterval struct {
	value *OrderMetricsInterval
	isSet bool
}

func (v NullableOrderMetricsInterval) Get() *OrderMetricsInterval {
	return v.value
}

func (v *NullableOrderMetricsInterval) Set(val *OrderMetricsInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderMetricsInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderMetricsInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderMetricsInterval(val *OrderMetricsInterval) *NullableOrderMetricsInterval {
	return &NullableOrderMetricsInterval{value: val, isSet: true}
}

func (v NullableOrderMetricsInterval) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderMetricsInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
