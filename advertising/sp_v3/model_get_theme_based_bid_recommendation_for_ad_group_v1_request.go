package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// GetThemeBasedBidRecommendationForAdGroupV1Request - struct for GetThemeBasedBidRecommendationForAdGroupV1Request
type GetThemeBasedBidRecommendationForAdGroupV1Request struct {
	AdGroupThemeBasedBidRecommendationRequestV4 *AdGroupThemeBasedBidRecommendationRequestV4
	AsinsThemeBasedBidRecommendationRequestV4   *AsinsThemeBasedBidRecommendationRequestV4
}

// AdGroupThemeBasedBidRecommendationRequestV4AsGetThemeBasedBidRecommendationForAdGroupV1Request is a convenience function that returns AdGroupThemeBasedBidRecommendationRequestV4 wrapped in GetThemeBasedBidRecommendationForAdGroupV1Request
func AdGroupThemeBasedBidRecommendationRequestV4AsGetThemeBasedBidRecommendationForAdGroupV1Request(v *AdGroupThemeBasedBidRecommendationRequestV4) GetThemeBasedBidRecommendationForAdGroupV1Request {
	return GetThemeBasedBidRecommendationForAdGroupV1Request{
		AdGroupThemeBasedBidRecommendationRequestV4: v,
	}
}

// AsinsThemeBasedBidRecommendationRequestV4AsGetThemeBasedBidRecommendationForAdGroupV1Request is a convenience function that returns AsinsThemeBasedBidRecommendationRequestV4 wrapped in GetThemeBasedBidRecommendationForAdGroupV1Request
func AsinsThemeBasedBidRecommendationRequestV4AsGetThemeBasedBidRecommendationForAdGroupV1Request(v *AsinsThemeBasedBidRecommendationRequestV4) GetThemeBasedBidRecommendationForAdGroupV1Request {
	return GetThemeBasedBidRecommendationForAdGroupV1Request{
		AsinsThemeBasedBidRecommendationRequestV4: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetThemeBasedBidRecommendationForAdGroupV1Request) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BIDS_FOR_EXISTING_AD_GROUP'
	if jsonDict["recommendationType"] == "BIDS_FOR_EXISTING_AD_GROUP" {
		// try to unmarshal JSON data into AdGroupThemeBasedBidRecommendationRequestV4
		err = sonic.Unmarshal(data, &dst.AdGroupThemeBasedBidRecommendationRequestV4)
		if err == nil {
			return nil // data stored in dst.AdGroupThemeBasedBidRecommendationRequestV4, return on the first match
		} else {
			dst.AdGroupThemeBasedBidRecommendationRequestV4 = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request as AdGroupThemeBasedBidRecommendationRequestV4: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BIDS_FOR_NEW_AD_GROUP'
	if jsonDict["recommendationType"] == "BIDS_FOR_NEW_AD_GROUP" {
		// try to unmarshal JSON data into AsinsThemeBasedBidRecommendationRequestV4
		err = sonic.Unmarshal(data, &dst.AsinsThemeBasedBidRecommendationRequestV4)
		if err == nil {
			return nil // data stored in dst.AsinsThemeBasedBidRecommendationRequestV4, return on the first match
		} else {
			dst.AsinsThemeBasedBidRecommendationRequestV4 = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request as AsinsThemeBasedBidRecommendationRequestV4: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AdGroupThemeBasedBidRecommendationRequestV4'
	if jsonDict["recommendationType"] == "AdGroupThemeBasedBidRecommendationRequestV4" {
		// try to unmarshal JSON data into AdGroupThemeBasedBidRecommendationRequestV4
		err = sonic.Unmarshal(data, &dst.AdGroupThemeBasedBidRecommendationRequestV4)
		if err == nil {
			return nil // data stored in dst.AdGroupThemeBasedBidRecommendationRequestV4, return on the first match
		} else {
			dst.AdGroupThemeBasedBidRecommendationRequestV4 = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request as AdGroupThemeBasedBidRecommendationRequestV4: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AsinsThemeBasedBidRecommendationRequestV4'
	if jsonDict["recommendationType"] == "AsinsThemeBasedBidRecommendationRequestV4" {
		// try to unmarshal JSON data into AsinsThemeBasedBidRecommendationRequestV4
		err = sonic.Unmarshal(data, &dst.AsinsThemeBasedBidRecommendationRequestV4)
		if err == nil {
			return nil // data stored in dst.AsinsThemeBasedBidRecommendationRequestV4, return on the first match
		} else {
			dst.AsinsThemeBasedBidRecommendationRequestV4 = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request as AsinsThemeBasedBidRecommendationRequestV4: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetThemeBasedBidRecommendationForAdGroupV1Request) MarshalJSON() ([]byte, error) {
	if src.AdGroupThemeBasedBidRecommendationRequestV4 != nil {
		return sonic.Marshal(&src.AdGroupThemeBasedBidRecommendationRequestV4)
	}

	if src.AsinsThemeBasedBidRecommendationRequestV4 != nil {
		return sonic.Marshal(&src.AsinsThemeBasedBidRecommendationRequestV4)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetThemeBasedBidRecommendationForAdGroupV1Request) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdGroupThemeBasedBidRecommendationRequestV4 != nil {
		return obj.AdGroupThemeBasedBidRecommendationRequestV4
	}

	if obj.AsinsThemeBasedBidRecommendationRequestV4 != nil {
		return obj.AsinsThemeBasedBidRecommendationRequestV4
	}

	// all schemas are nil
	return nil
}

type NullableGetThemeBasedBidRecommendationForAdGroupV1Request struct {
	value *GetThemeBasedBidRecommendationForAdGroupV1Request
	isSet bool
}

func (v NullableGetThemeBasedBidRecommendationForAdGroupV1Request) Get() *GetThemeBasedBidRecommendationForAdGroupV1Request {
	return v.value
}

func (v *NullableGetThemeBasedBidRecommendationForAdGroupV1Request) Set(val *GetThemeBasedBidRecommendationForAdGroupV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThemeBasedBidRecommendationForAdGroupV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThemeBasedBidRecommendationForAdGroupV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThemeBasedBidRecommendationForAdGroupV1Request(val *GetThemeBasedBidRecommendationForAdGroupV1Request) *NullableGetThemeBasedBidRecommendationForAdGroupV1Request {
	return &NullableGetThemeBasedBidRecommendationForAdGroupV1Request{value: val, isSet: true}
}

func (v NullableGetThemeBasedBidRecommendationForAdGroupV1Request) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetThemeBasedBidRecommendationForAdGroupV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
