package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// GetGlobalRankedKeywordRecommendationRequest - struct for GetGlobalRankedKeywordRecommendationRequest
type GetGlobalRankedKeywordRecommendationRequest struct {
	GlobalRankedKeywordTargetsForAdGroupRequest *GlobalRankedKeywordTargetsForAdGroupRequest
	GlobalRankedKeywordTargetsForAsinsRequest   *GlobalRankedKeywordTargetsForAsinsRequest
}

// GlobalRankedKeywordTargetsForAdGroupRequestAsGetGlobalRankedKeywordRecommendationRequest is a convenience function that returns GlobalRankedKeywordTargetsForAdGroupRequest wrapped in GetGlobalRankedKeywordRecommendationRequest
func GlobalRankedKeywordTargetsForAdGroupRequestAsGetGlobalRankedKeywordRecommendationRequest(v *GlobalRankedKeywordTargetsForAdGroupRequest) GetGlobalRankedKeywordRecommendationRequest {
	return GetGlobalRankedKeywordRecommendationRequest{
		GlobalRankedKeywordTargetsForAdGroupRequest: v,
	}
}

// GlobalRankedKeywordTargetsForAsinsRequestAsGetGlobalRankedKeywordRecommendationRequest is a convenience function that returns GlobalRankedKeywordTargetsForAsinsRequest wrapped in GetGlobalRankedKeywordRecommendationRequest
func GlobalRankedKeywordTargetsForAsinsRequestAsGetGlobalRankedKeywordRecommendationRequest(v *GlobalRankedKeywordTargetsForAsinsRequest) GetGlobalRankedKeywordRecommendationRequest {
	return GetGlobalRankedKeywordRecommendationRequest{
		GlobalRankedKeywordTargetsForAsinsRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGlobalRankedKeywordRecommendationRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'KEYWORDS_FOR_ADGROUP'
	if jsonDict["recommendationType"] == "KEYWORDS_FOR_ADGROUP" {
		// try to unmarshal JSON data into GlobalRankedKeywordTargetsForAdGroupRequest
		err = sonic.Unmarshal(data, &dst.GlobalRankedKeywordTargetsForAdGroupRequest)
		if err == nil {
			return nil // data stored in dst.GlobalRankedKeywordTargetsForAdGroupRequest, return on the first match
		} else {
			dst.GlobalRankedKeywordTargetsForAdGroupRequest = nil
			return fmt.Errorf("failed to unmarshal GetGlobalRankedKeywordRecommendationRequest as GlobalRankedKeywordTargetsForAdGroupRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'KEYWORDS_FOR_ASINS'
	if jsonDict["recommendationType"] == "KEYWORDS_FOR_ASINS" {
		// try to unmarshal JSON data into GlobalRankedKeywordTargetsForAsinsRequest
		err = sonic.Unmarshal(data, &dst.GlobalRankedKeywordTargetsForAsinsRequest)
		if err == nil {
			return nil // data stored in dst.GlobalRankedKeywordTargetsForAsinsRequest, return on the first match
		} else {
			dst.GlobalRankedKeywordTargetsForAsinsRequest = nil
			return fmt.Errorf("failed to unmarshal GetGlobalRankedKeywordRecommendationRequest as GlobalRankedKeywordTargetsForAsinsRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GlobalRankedKeywordTargetsForAdGroupRequest'
	if jsonDict["recommendationType"] == "GlobalRankedKeywordTargetsForAdGroupRequest" {
		// try to unmarshal JSON data into GlobalRankedKeywordTargetsForAdGroupRequest
		err = sonic.Unmarshal(data, &dst.GlobalRankedKeywordTargetsForAdGroupRequest)
		if err == nil {
			return nil // data stored in dst.GlobalRankedKeywordTargetsForAdGroupRequest, return on the first match
		} else {
			dst.GlobalRankedKeywordTargetsForAdGroupRequest = nil
			return fmt.Errorf("failed to unmarshal GetGlobalRankedKeywordRecommendationRequest as GlobalRankedKeywordTargetsForAdGroupRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GlobalRankedKeywordTargetsForAsinsRequest'
	if jsonDict["recommendationType"] == "GlobalRankedKeywordTargetsForAsinsRequest" {
		// try to unmarshal JSON data into GlobalRankedKeywordTargetsForAsinsRequest
		err = sonic.Unmarshal(data, &dst.GlobalRankedKeywordTargetsForAsinsRequest)
		if err == nil {
			return nil // data stored in dst.GlobalRankedKeywordTargetsForAsinsRequest, return on the first match
		} else {
			dst.GlobalRankedKeywordTargetsForAsinsRequest = nil
			return fmt.Errorf("failed to unmarshal GetGlobalRankedKeywordRecommendationRequest as GlobalRankedKeywordTargetsForAsinsRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGlobalRankedKeywordRecommendationRequest) MarshalJSON() ([]byte, error) {
	if src.GlobalRankedKeywordTargetsForAdGroupRequest != nil {
		return sonic.Marshal(&src.GlobalRankedKeywordTargetsForAdGroupRequest)
	}

	if src.GlobalRankedKeywordTargetsForAsinsRequest != nil {
		return sonic.Marshal(&src.GlobalRankedKeywordTargetsForAsinsRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGlobalRankedKeywordRecommendationRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GlobalRankedKeywordTargetsForAdGroupRequest != nil {
		return obj.GlobalRankedKeywordTargetsForAdGroupRequest
	}

	if obj.GlobalRankedKeywordTargetsForAsinsRequest != nil {
		return obj.GlobalRankedKeywordTargetsForAsinsRequest
	}

	// all schemas are nil
	return nil
}

type NullableGetGlobalRankedKeywordRecommendationRequest struct {
	value *GetGlobalRankedKeywordRecommendationRequest
	isSet bool
}

func (v NullableGetGlobalRankedKeywordRecommendationRequest) Get() *GetGlobalRankedKeywordRecommendationRequest {
	return v.value
}

func (v *NullableGetGlobalRankedKeywordRecommendationRequest) Set(val *GetGlobalRankedKeywordRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGlobalRankedKeywordRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGlobalRankedKeywordRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGlobalRankedKeywordRecommendationRequest(val *GetGlobalRankedKeywordRecommendationRequest) *NullableGetGlobalRankedKeywordRecommendationRequest {
	return &NullableGetGlobalRankedKeywordRecommendationRequest{value: val, isSet: true}
}

func (v NullableGetGlobalRankedKeywordRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetGlobalRankedKeywordRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
