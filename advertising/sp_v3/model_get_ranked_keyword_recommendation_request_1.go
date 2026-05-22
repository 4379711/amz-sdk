package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// GetRankedKeywordRecommendationRequest1 - struct for GetRankedKeywordRecommendationRequest1
type GetRankedKeywordRecommendationRequest1 struct {
	RankedKeywordTargetsForAdGroupRequest *RankedKeywordTargetsForAdGroupRequest
	RankedKeywordTargetsForAsinsRequest   *RankedKeywordTargetsForAsinsRequest
}

// RankedKeywordTargetsForAdGroupRequestAsGetRankedKeywordRecommendationRequest1 is a convenience function that returns RankedKeywordTargetsForAdGroupRequest wrapped in GetRankedKeywordRecommendationRequest1
func RankedKeywordTargetsForAdGroupRequestAsGetRankedKeywordRecommendationRequest1(v *RankedKeywordTargetsForAdGroupRequest) GetRankedKeywordRecommendationRequest1 {
	return GetRankedKeywordRecommendationRequest1{
		RankedKeywordTargetsForAdGroupRequest: v,
	}
}

// RankedKeywordTargetsForAsinsRequestAsGetRankedKeywordRecommendationRequest1 is a convenience function that returns RankedKeywordTargetsForAsinsRequest wrapped in GetRankedKeywordRecommendationRequest1
func RankedKeywordTargetsForAsinsRequestAsGetRankedKeywordRecommendationRequest1(v *RankedKeywordTargetsForAsinsRequest) GetRankedKeywordRecommendationRequest1 {
	return GetRankedKeywordRecommendationRequest1{
		RankedKeywordTargetsForAsinsRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRankedKeywordRecommendationRequest1) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'KEYWORDS_FOR_ADGROUP'
	if jsonDict["recommendationType"] == "KEYWORDS_FOR_ADGROUP" {
		// try to unmarshal JSON data into RankedKeywordTargetsForAdGroupRequest
		err = sonic.Unmarshal(data, &dst.RankedKeywordTargetsForAdGroupRequest)
		if err == nil {
			return nil // data stored in dst.RankedKeywordTargetsForAdGroupRequest, return on the first match
		} else {
			dst.RankedKeywordTargetsForAdGroupRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest1 as RankedKeywordTargetsForAdGroupRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'KEYWORDS_FOR_ASINS'
	if jsonDict["recommendationType"] == "KEYWORDS_FOR_ASINS" {
		// try to unmarshal JSON data into RankedKeywordTargetsForAsinsRequest
		err = sonic.Unmarshal(data, &dst.RankedKeywordTargetsForAsinsRequest)
		if err == nil {
			return nil // data stored in dst.RankedKeywordTargetsForAsinsRequest, return on the first match
		} else {
			dst.RankedKeywordTargetsForAsinsRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest1 as RankedKeywordTargetsForAsinsRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RankedKeywordTargetsForAdGroupRequest'
	if jsonDict["recommendationType"] == "RankedKeywordTargetsForAdGroupRequest" {
		// try to unmarshal JSON data into RankedKeywordTargetsForAdGroupRequest
		err = sonic.Unmarshal(data, &dst.RankedKeywordTargetsForAdGroupRequest)
		if err == nil {
			return nil // data stored in dst.RankedKeywordTargetsForAdGroupRequest, return on the first match
		} else {
			dst.RankedKeywordTargetsForAdGroupRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest1 as RankedKeywordTargetsForAdGroupRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RankedKeywordTargetsForAsinsRequest'
	if jsonDict["recommendationType"] == "RankedKeywordTargetsForAsinsRequest" {
		// try to unmarshal JSON data into RankedKeywordTargetsForAsinsRequest
		err = sonic.Unmarshal(data, &dst.RankedKeywordTargetsForAsinsRequest)
		if err == nil {
			return nil // data stored in dst.RankedKeywordTargetsForAsinsRequest, return on the first match
		} else {
			dst.RankedKeywordTargetsForAsinsRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest1 as RankedKeywordTargetsForAsinsRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRankedKeywordRecommendationRequest1) MarshalJSON() ([]byte, error) {
	if src.RankedKeywordTargetsForAdGroupRequest != nil {
		return sonic.Marshal(&src.RankedKeywordTargetsForAdGroupRequest)
	}

	if src.RankedKeywordTargetsForAsinsRequest != nil {
		return sonic.Marshal(&src.RankedKeywordTargetsForAsinsRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRankedKeywordRecommendationRequest1) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RankedKeywordTargetsForAdGroupRequest != nil {
		return obj.RankedKeywordTargetsForAdGroupRequest
	}

	if obj.RankedKeywordTargetsForAsinsRequest != nil {
		return obj.RankedKeywordTargetsForAsinsRequest
	}

	// all schemas are nil
	return nil
}

type NullableGetRankedKeywordRecommendationRequest1 struct {
	value *GetRankedKeywordRecommendationRequest1
	isSet bool
}

func (v NullableGetRankedKeywordRecommendationRequest1) Get() *GetRankedKeywordRecommendationRequest1 {
	return v.value
}

func (v *NullableGetRankedKeywordRecommendationRequest1) Set(val *GetRankedKeywordRecommendationRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRankedKeywordRecommendationRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRankedKeywordRecommendationRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRankedKeywordRecommendationRequest1(val *GetRankedKeywordRecommendationRequest1) *NullableGetRankedKeywordRecommendationRequest1 {
	return &NullableGetRankedKeywordRecommendationRequest1{value: val, isSet: true}
}

func (v NullableGetRankedKeywordRecommendationRequest1) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRankedKeywordRecommendationRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
