package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsPlacement You can enable controls to adjust your bid based on the placement location. Specify a location where you want to use bid controls. The percentage value set is the percentage of the original bid for which you want to have your bid adjustment increased. For example, a 50% Top of Search adjustment on a $1.00 bid would increase the bid to $1.50 for the opportunity to win a Top of Search placement. A further 100% Amazon Business adjustment would increase the bid to $3.00 for the Amazon Business Top of Search placement and to $2.00 for all other Amazon Business placements. The Amazon Business Bid Adjustment and Reporting for Sponsored Products will be coming soon to Bulksheets. | Predicate |  Placement | |-----------|------------| | `PLACEMENT_TOP` | Top of search (first page) | | `PLACEMENT_PRODUCT_PAGE` | Product pages | | `PLACEMENT_REST_OF_SEARCH` | Rest of the search | | `SITE_AMAZON_BUSINESS` | Site Amazon Business |
type SponsoredProductsPlacement string

// List of SponsoredProductsPlacement
const (
	SPONSOREDPRODUCTSPLACEMENT_PLACEMENT_TOP            SponsoredProductsPlacement = "PLACEMENT_TOP"
	SPONSOREDPRODUCTSPLACEMENT_PLACEMENT_PRODUCT_PAGE   SponsoredProductsPlacement = "PLACEMENT_PRODUCT_PAGE"
	SPONSOREDPRODUCTSPLACEMENT_PLACEMENT_REST_OF_SEARCH SponsoredProductsPlacement = "PLACEMENT_REST_OF_SEARCH"
	SPONSOREDPRODUCTSPLACEMENT_SITE_AMAZON_BUSINESS     SponsoredProductsPlacement = "SITE_AMAZON_BUSINESS"
)

// All allowed values of SponsoredProductsPlacement enum
var AllowedSponsoredProductsPlacementEnumValues = []SponsoredProductsPlacement{
	"PLACEMENT_TOP",
	"PLACEMENT_PRODUCT_PAGE",
	"PLACEMENT_REST_OF_SEARCH",
	"SITE_AMAZON_BUSINESS",
}

func (v *SponsoredProductsPlacement) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsPlacement(value)
	for _, existing := range AllowedSponsoredProductsPlacementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsPlacement", value)
}

// NewSponsoredProductsPlacementFromValue returns a pointer to a valid SponsoredProductsPlacement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsPlacementFromValue(v string) (*SponsoredProductsPlacement, error) {
	ev := SponsoredProductsPlacement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsPlacement: valid values are %v", v, AllowedSponsoredProductsPlacementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsPlacement) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsPlacementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsPlacement value
func (v SponsoredProductsPlacement) Ptr() *SponsoredProductsPlacement {
	return &v
}

type NullableSponsoredProductsPlacement struct {
	value *SponsoredProductsPlacement
	isSet bool
}

func (v NullableSponsoredProductsPlacement) Get() *SponsoredProductsPlacement {
	return v.value
}

func (v *NullableSponsoredProductsPlacement) Set(val *SponsoredProductsPlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsPlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsPlacement(val *SponsoredProductsPlacement) *NullableSponsoredProductsPlacement {
	return &NullableSponsoredProductsPlacement{value: val, isSet: true}
}

func (v NullableSponsoredProductsPlacement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
