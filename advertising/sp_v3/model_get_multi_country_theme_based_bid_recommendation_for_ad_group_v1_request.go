package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request - struct for GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
type GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request struct {
	MultiCountryAdGroupThemeBasedBidRecommendationRequest *MultiCountryAdGroupThemeBasedBidRecommendationRequest
	MultiCountryAsinsThemeBasedBidRecommendationRequest   *MultiCountryAsinsThemeBasedBidRecommendationRequest
}

// MultiCountryAdGroupThemeBasedBidRecommendationRequestAsGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request is a convenience function that returns MultiCountryAdGroupThemeBasedBidRecommendationRequest wrapped in GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
func MultiCountryAdGroupThemeBasedBidRecommendationRequestAsGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request(v *MultiCountryAdGroupThemeBasedBidRecommendationRequest) GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	return GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request{
		MultiCountryAdGroupThemeBasedBidRecommendationRequest: v,
	}
}

// MultiCountryAsinsThemeBasedBidRecommendationRequestAsGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request is a convenience function that returns MultiCountryAsinsThemeBasedBidRecommendationRequest wrapped in GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
func MultiCountryAsinsThemeBasedBidRecommendationRequestAsGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request(v *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	return GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request{
		MultiCountryAsinsThemeBasedBidRecommendationRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BIDS_FOR_EXISTING_AD_GROUP'
	if jsonDict["recommendationType"] == "BIDS_FOR_EXISTING_AD_GROUP" {
		// try to unmarshal JSON data into MultiCountryAdGroupThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.MultiCountryAdGroupThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.MultiCountryAdGroupThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.MultiCountryAdGroupThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request as MultiCountryAdGroupThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BIDS_FOR_NEW_AD_GROUP'
	if jsonDict["recommendationType"] == "BIDS_FOR_NEW_AD_GROUP" {
		// try to unmarshal JSON data into MultiCountryAsinsThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.MultiCountryAsinsThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.MultiCountryAsinsThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.MultiCountryAsinsThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request as MultiCountryAsinsThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MultiCountryAdGroupThemeBasedBidRecommendationRequest'
	if jsonDict["recommendationType"] == "MultiCountryAdGroupThemeBasedBidRecommendationRequest" {
		// try to unmarshal JSON data into MultiCountryAdGroupThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.MultiCountryAdGroupThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.MultiCountryAdGroupThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.MultiCountryAdGroupThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request as MultiCountryAdGroupThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MultiCountryAsinsThemeBasedBidRecommendationRequest'
	if jsonDict["recommendationType"] == "MultiCountryAsinsThemeBasedBidRecommendationRequest" {
		// try to unmarshal JSON data into MultiCountryAsinsThemeBasedBidRecommendationRequest
		err = sonic.Unmarshal(data, &dst.MultiCountryAsinsThemeBasedBidRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.MultiCountryAsinsThemeBasedBidRecommendationRequest, return on the first match
		} else {
			dst.MultiCountryAsinsThemeBasedBidRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request as MultiCountryAsinsThemeBasedBidRecommendationRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) MarshalJSON() ([]byte, error) {
	if src.MultiCountryAdGroupThemeBasedBidRecommendationRequest != nil {
		return sonic.Marshal(&src.MultiCountryAdGroupThemeBasedBidRecommendationRequest)
	}

	if src.MultiCountryAsinsThemeBasedBidRecommendationRequest != nil {
		return sonic.Marshal(&src.MultiCountryAsinsThemeBasedBidRecommendationRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MultiCountryAdGroupThemeBasedBidRecommendationRequest != nil {
		return obj.MultiCountryAdGroupThemeBasedBidRecommendationRequest
	}

	if obj.MultiCountryAsinsThemeBasedBidRecommendationRequest != nil {
		return obj.MultiCountryAsinsThemeBasedBidRecommendationRequest
	}

	// all schemas are nil
	return nil
}

type NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request struct {
	value *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
	isSet bool
}

func (v NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) Get() *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	return v.value
}

func (v *NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) Set(val *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request(val *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) *NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	return &NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request{value: val, isSet: true}
}

func (v NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
