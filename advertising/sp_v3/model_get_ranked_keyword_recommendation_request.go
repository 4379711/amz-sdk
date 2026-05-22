package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// GetRankedKeywordRecommendationRequest - struct for GetRankedKeywordRecommendationRequest
type GetRankedKeywordRecommendationRequest struct {
	AdGroupKeywordTargetRankRecommendationRequest *AdGroupKeywordTargetRankRecommendationRequest
	AsinsKeywordTargetRankRecommendationRequest   *AsinsKeywordTargetRankRecommendationRequest
}

// AdGroupKeywordTargetRankRecommendationRequestAsGetRankedKeywordRecommendationRequest is a convenience function that returns AdGroupKeywordTargetRankRecommendationRequest wrapped in GetRankedKeywordRecommendationRequest
func AdGroupKeywordTargetRankRecommendationRequestAsGetRankedKeywordRecommendationRequest(v *AdGroupKeywordTargetRankRecommendationRequest) GetRankedKeywordRecommendationRequest {
	return GetRankedKeywordRecommendationRequest{
		AdGroupKeywordTargetRankRecommendationRequest: v,
	}
}

// AsinsKeywordTargetRankRecommendationRequestAsGetRankedKeywordRecommendationRequest is a convenience function that returns AsinsKeywordTargetRankRecommendationRequest wrapped in GetRankedKeywordRecommendationRequest
func AsinsKeywordTargetRankRecommendationRequestAsGetRankedKeywordRecommendationRequest(v *AsinsKeywordTargetRankRecommendationRequest) GetRankedKeywordRecommendationRequest {
	return GetRankedKeywordRecommendationRequest{
		AsinsKeywordTargetRankRecommendationRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRankedKeywordRecommendationRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'KEYWORDS_FOR_ADGROUP'
	if jsonDict["recommendationType"] == "KEYWORDS_FOR_ADGROUP" {
		// try to unmarshal JSON data into AdGroupKeywordTargetRankRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AdGroupKeywordTargetRankRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AdGroupKeywordTargetRankRecommendationRequest, return on the first match
		} else {
			dst.AdGroupKeywordTargetRankRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest as AdGroupKeywordTargetRankRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'KEYWORDS_FOR_ASINS'
	if jsonDict["recommendationType"] == "KEYWORDS_FOR_ASINS" {
		// try to unmarshal JSON data into AsinsKeywordTargetRankRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AsinsKeywordTargetRankRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AsinsKeywordTargetRankRecommendationRequest, return on the first match
		} else {
			dst.AsinsKeywordTargetRankRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest as AsinsKeywordTargetRankRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AdGroupKeywordTargetRankRecommendationRequest'
	if jsonDict["recommendationType"] == "AdGroupKeywordTargetRankRecommendationRequest" {
		// try to unmarshal JSON data into AdGroupKeywordTargetRankRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AdGroupKeywordTargetRankRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AdGroupKeywordTargetRankRecommendationRequest, return on the first match
		} else {
			dst.AdGroupKeywordTargetRankRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest as AdGroupKeywordTargetRankRecommendationRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AsinsKeywordTargetRankRecommendationRequest'
	if jsonDict["recommendationType"] == "AsinsKeywordTargetRankRecommendationRequest" {
		// try to unmarshal JSON data into AsinsKeywordTargetRankRecommendationRequest
		err = sonic.Unmarshal(data, &dst.AsinsKeywordTargetRankRecommendationRequest)
		if err == nil {
			return nil // data stored in dst.AsinsKeywordTargetRankRecommendationRequest, return on the first match
		} else {
			dst.AsinsKeywordTargetRankRecommendationRequest = nil
			return fmt.Errorf("failed to unmarshal GetRankedKeywordRecommendationRequest as AsinsKeywordTargetRankRecommendationRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRankedKeywordRecommendationRequest) MarshalJSON() ([]byte, error) {
	if src.AdGroupKeywordTargetRankRecommendationRequest != nil {
		return sonic.Marshal(&src.AdGroupKeywordTargetRankRecommendationRequest)
	}

	if src.AsinsKeywordTargetRankRecommendationRequest != nil {
		return sonic.Marshal(&src.AsinsKeywordTargetRankRecommendationRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRankedKeywordRecommendationRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdGroupKeywordTargetRankRecommendationRequest != nil {
		return obj.AdGroupKeywordTargetRankRecommendationRequest
	}

	if obj.AsinsKeywordTargetRankRecommendationRequest != nil {
		return obj.AsinsKeywordTargetRankRecommendationRequest
	}

	// all schemas are nil
	return nil
}

type NullableGetRankedKeywordRecommendationRequest struct {
	value *GetRankedKeywordRecommendationRequest
	isSet bool
}

func (v NullableGetRankedKeywordRecommendationRequest) Get() *GetRankedKeywordRecommendationRequest {
	return v.value
}

func (v *NullableGetRankedKeywordRecommendationRequest) Set(val *GetRankedKeywordRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRankedKeywordRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRankedKeywordRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRankedKeywordRecommendationRequest(val *GetRankedKeywordRecommendationRequest) *NullableGetRankedKeywordRecommendationRequest {
	return &NullableGetRankedKeywordRecommendationRequest{value: val, isSet: true}
}

func (v NullableGetRankedKeywordRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRankedKeywordRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
