package product_eligibility

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Resource The advertising resource of which you wish to check feature access for. Example: Sponsored Display Campaign.
type Resource string

// List of Resource
const (
	RESOURCE_SD_CAMPAIGNS                    Resource = "sd:campaigns"
	RESOURCE_ST_CAMPAIGNS                    Resource = "st:campaigns"
	RESOURCE_DSP_CAMPAIGNS                   Resource = "dsp:campaigns"
	RESOURCE_UNIFIED_REPORT_CENTER           Resource = "unifiedReportCenter"
	RESOURCE_SD_CAMPAIGNS_ASINLESS           Resource = "sd:campaigns:asinless"
	RESOURCE_ST_CAMPAIGNS_ASINLESS           Resource = "st:campaigns:asinless"
	RESOURCE_MULTI_TOUCH_ATTRIBUTION_METRICS Resource = "multiTouchAttributionMetrics"
)

// All allowed values of Resource enum
var AllowedResourceEnumValues = []Resource{
	"sd:campaigns",
	"st:campaigns",
	"dsp:campaigns",
	"unifiedReportCenter",
	"sd:campaigns:asinless",
	"st:campaigns:asinless",
	"multiTouchAttributionMetrics",
}

func (v *Resource) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Resource(value)
	for _, existing := range AllowedResourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Resource", value)
}

// NewResourceFromValue returns a pointer to a valid Resource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceFromValue(v string) (*Resource, error) {
	ev := Resource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Resource: valid values are %v", v, AllowedResourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Resource) IsValid() bool {
	for _, existing := range AllowedResourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Resource value
func (v Resource) Ptr() *Resource {
	return &v
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
