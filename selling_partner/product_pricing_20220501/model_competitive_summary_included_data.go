package product_pricing_20220501

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CompetitiveSummaryIncludedData The supported data types in the `getCompetitiveSummary` API.
type CompetitiveSummaryIncludedData string

// List of CompetitiveSummaryIncludedData
const (
	COMPETITIVESUMMARYINCLUDEDDATA_FEATURED_BUYING_OPTIONS CompetitiveSummaryIncludedData = "featuredBuyingOptions"
	COMPETITIVESUMMARYINCLUDEDDATA_REFERENCE_PRICES        CompetitiveSummaryIncludedData = "referencePrices"
	COMPETITIVESUMMARYINCLUDEDDATA_LOWEST_PRICED_OFFERS    CompetitiveSummaryIncludedData = "lowestPricedOffers"
)

// All allowed values of CompetitiveSummaryIncludedData enum
var AllowedCompetitiveSummaryIncludedDataEnumValues = []CompetitiveSummaryIncludedData{
	COMPETITIVESUMMARYINCLUDEDDATA_FEATURED_BUYING_OPTIONS,
	COMPETITIVESUMMARYINCLUDEDDATA_REFERENCE_PRICES,
	COMPETITIVESUMMARYINCLUDEDDATA_LOWEST_PRICED_OFFERS,
}

func (v *CompetitiveSummaryIncludedData) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompetitiveSummaryIncludedData(value)
	for _, existing := range AllowedCompetitiveSummaryIncludedDataEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompetitiveSummaryIncludedData", value)
}

// NewCompetitiveSummaryIncludedDataFromValue returns a pointer to a valid CompetitiveSummaryIncludedData
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompetitiveSummaryIncludedDataFromValue(v string) (*CompetitiveSummaryIncludedData, error) {
	ev := CompetitiveSummaryIncludedData(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompetitiveSummaryIncludedData: valid values are %v", v, AllowedCompetitiveSummaryIncludedDataEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompetitiveSummaryIncludedData) IsValid() bool {
	for _, existing := range AllowedCompetitiveSummaryIncludedDataEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompetitiveSummaryIncludedData value
func (v CompetitiveSummaryIncludedData) Ptr() *CompetitiveSummaryIncludedData {
	return &v
}

type NullableCompetitiveSummaryIncludedData struct {
	value *CompetitiveSummaryIncludedData
	isSet bool
}

func (v NullableCompetitiveSummaryIncludedData) Get() *CompetitiveSummaryIncludedData {
	return v.value
}

func (v *NullableCompetitiveSummaryIncludedData) Set(val *CompetitiveSummaryIncludedData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompetitiveSummaryIncludedData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompetitiveSummaryIncludedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompetitiveSummaryIncludedData(val *CompetitiveSummaryIncludedData) *NullableCompetitiveSummaryIncludedData {
	return &NullableCompetitiveSummaryIncludedData{value: val, isSet: true}
}

func (v NullableCompetitiveSummaryIncludedData) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompetitiveSummaryIncludedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
