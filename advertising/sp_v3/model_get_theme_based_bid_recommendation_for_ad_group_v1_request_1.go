package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// GetThemeBasedBidRecommendationForAdGroupV1Request1 - struct for GetThemeBasedBidRecommendationForAdGroupV1Request1
type GetThemeBasedBidRecommendationForAdGroupV1Request1 struct {
	AdGroupThemeBasedBidRecommendationRequest *AdGroupThemeBasedBidRecommendationRequest
	AsinsThemeBasedBidRecommendationRequest   *AsinsThemeBasedBidRecommendationRequest
}

// AdGroupThemeBasedBidRecommendationRequestAsGetThemeBasedBidRecommendationForAdGroupV1Request1 is a convenience function that returns AdGroupThemeBasedBidRecommendationRequest wrapped in GetThemeBasedBidRecommendationForAdGroupV1Request1
func AdGroupThemeBasedBidRecommendationRequestAsGetThemeBasedBidRecommendationForAdGroupV1Request1(v *AdGroupThemeBasedBidRecommendationRequest) GetThemeBasedBidRecommendationForAdGroupV1Request1 {
	return GetThemeBasedBidRecommendationForAdGroupV1Request1{
		AdGroupThemeBasedBidRecommendationRequest: v,
	}
}

// AsinsThemeBasedBidRecommendationRequestAsGetThemeBasedBidRecommendationForAdGroupV1Request1 is a convenience function that returns AsinsThemeBasedBidRecommendationRequest wrapped in GetThemeBasedBidRecommendationForAdGroupV1Request1
func AsinsThemeBasedBidRecommendationRequestAsGetThemeBasedBidRecommendationForAdGroupV1Request1(v *AsinsThemeBasedBidRecommendationRequest) GetThemeBasedBidRecommendationForAdGroupV1Request1 {
	return GetThemeBasedBidRecommendationForAdGroupV1Request1{
		AsinsThemeBasedBidRecommendationRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetThemeBasedBidRecommendationForAdGroupV1Request1) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BIDS_FOR_EXISTING_AD_GROUP'
	if jsonDict["recommendationType"] == "BIDS_FOR_EXISTING_AD_GROUP" {
		// try to unmarshal JSON data into AdGroupThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AdGroupThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AdGroupThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.AdGroupThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request1 as AdGroupThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BIDS_FOR_NEW_AD_GROUP'
	if jsonDict["recommendationType"] == "BIDS_FOR_NEW_AD_GROUP" {
		// try to unmarshal JSON data into AsinsThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AsinsThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AsinsThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.AsinsThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request1 as AsinsThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AdGroupThemeBasedBidRecommendationRequest'
	if jsonDict["recommendationType"] == "AdGroupThemeBasedBidRecommendationRequest" {
		// try to unmarshal JSON data into AdGroupThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AdGroupThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AdGroupThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.AdGroupThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request1 as AdGroupThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AsinsThemeBasedBidRecommendationRequest'
	if jsonDict["recommendationType"] == "AsinsThemeBasedBidRecommendationRequest" {
		// try to unmarshal JSON data into AsinsThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AsinsThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AsinsThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.AsinsThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetThemeBasedBidRecommendationForAdGroupV1Request1 as AsinsThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetThemeBasedBidRecommendationForAdGroupV1Request1) MarshalJSON() ([]byte, error) {
	if src.AdGroupThemeBasedBidRecommendationRequest != nil {
		return sonic.Marshal(&src.AdGroupThemeBasedBidRecommendationRequest)
	}

	if src.AsinsThemeBasedBidRecommendationRequest != nil {
		return sonic.Marshal(&src.AsinsThemeBasedBidRecommendationRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetThemeBasedBidRecommendationForAdGroupV1Request1) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdGroupThemeBasedBidRecommendationRequest != nil {
		return obj.AdGroupThemeBasedBidRecommendationRequest
	}

	if obj.AsinsThemeBasedBidRecommendationRequest != nil {
		return obj.AsinsThemeBasedBidRecommendationRequest
	}

	// all schemas are nil
	return nil
}

type NullableGetThemeBasedBidRecommendationForAdGroupV1Request1 struct {
	value *GetThemeBasedBidRecommendationForAdGroupV1Request1
	isSet bool
}

func (v NullableGetThemeBasedBidRecommendationForAdGroupV1Request1) Get() *GetThemeBasedBidRecommendationForAdGroupV1Request1 {
	return v.value
}

func (v *NullableGetThemeBasedBidRecommendationForAdGroupV1Request1) Set(val *GetThemeBasedBidRecommendationForAdGroupV1Request1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThemeBasedBidRecommendationForAdGroupV1Request1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThemeBasedBidRecommendationForAdGroupV1Request1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThemeBasedBidRecommendationForAdGroupV1Request1(val *GetThemeBasedBidRecommendationForAdGroupV1Request1) *NullableGetThemeBasedBidRecommendationForAdGroupV1Request1 {
	return &NullableGetThemeBasedBidRecommendationForAdGroupV1Request1{value: val, isSet: true}
}

func (v NullableGetThemeBasedBidRecommendationForAdGroupV1Request1) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetThemeBasedBidRecommendationForAdGroupV1Request1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
