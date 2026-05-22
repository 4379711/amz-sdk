package replenishment20221107

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Metric The metric name and description.
type Metric string

// List of Metric
const (
	METRIC_SHIPPED_SUBSCRIPTION_UNITS                 Metric = "SHIPPED_SUBSCRIPTION_UNITS"
	METRIC_TOTAL_SUBSCRIPTIONS_REVENUE                Metric = "TOTAL_SUBSCRIPTIONS_REVENUE"
	METRIC_ACTIVE_SUBSCRIPTIONS                       Metric = "ACTIVE_SUBSCRIPTIONS"
	METRIC_NOT_DELIVERED_DUE_TO_OOS                   Metric = "NOT_DELIVERED_DUE_TO_OOS"
	METRIC_SUBSCRIBER_NON_SUBSCRIBER_AVERAGE_REVENUE  Metric = "SUBSCRIBER_NON_SUBSCRIBER_AVERAGE_REVENUE"
	METRIC_LOST_REVENUE_DUE_TO_OOS                    Metric = "LOST_REVENUE_DUE_TO_OOS"
	METRIC_SUBSCRIBER_NON_SUBSCRIBER_AVERAGE_REORDERS Metric = "SUBSCRIBER_NON_SUBSCRIBER_AVERAGE_REORDERS"
	METRIC_COUPONS_REVENUE_PENETRATION                Metric = "COUPONS_REVENUE_PENETRATION"
	METRIC_REVENUE_BY_DELIVERIES                      Metric = "REVENUE_BY_DELIVERIES"
	METRIC_SUBSCRIBER_RETENTION                       Metric = "SUBSCRIBER_RETENTION"
	METRIC_REVENUE_PENETRATION_BY_SELLER_FUNDING      Metric = "REVENUE_PENETRATION_BY_SELLER_FUNDING"
	METRIC_SHARE_OF_COUPON_SUBSCRIPTIONS              Metric = "SHARE_OF_COUPON_SUBSCRIPTIONS"
)

// All allowed values of Metric enum
var AllowedMetricEnumValues = []Metric{
	METRIC_SHIPPED_SUBSCRIPTION_UNITS,
	METRIC_TOTAL_SUBSCRIPTIONS_REVENUE,
	METRIC_ACTIVE_SUBSCRIPTIONS,
	METRIC_NOT_DELIVERED_DUE_TO_OOS,
	METRIC_SUBSCRIBER_NON_SUBSCRIBER_AVERAGE_REVENUE,
	METRIC_LOST_REVENUE_DUE_TO_OOS,
	METRIC_SUBSCRIBER_NON_SUBSCRIBER_AVERAGE_REORDERS,
	METRIC_COUPONS_REVENUE_PENETRATION,
	METRIC_REVENUE_BY_DELIVERIES,
	METRIC_SUBSCRIBER_RETENTION,
	METRIC_REVENUE_PENETRATION_BY_SELLER_FUNDING,
	METRIC_SHARE_OF_COUPON_SUBSCRIPTIONS,
}

func (v *Metric) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Metric(value)
	for _, existing := range AllowedMetricEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Metric", value)
}

// NewMetricFromValue returns a pointer to a valid Metric
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricFromValue(v string) (*Metric, error) {
	ev := Metric(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Metric: valid values are %v", v, AllowedMetricEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Metric) IsValid() bool {
	for _, existing := range AllowedMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Metric value
func (v Metric) Ptr() *Metric {
	return &v
}

type NullableMetric struct {
	value *Metric
	isSet bool
}

func (v NullableMetric) Get() *Metric {
	return v.value
}

func (v *NullableMetric) Set(val *Metric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetric(val *Metric) *NullableMetric {
	return &NullableMetric{value: val, isSet: true}
}

func (v NullableMetric) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
