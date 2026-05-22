package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SPGetBudgetRulesRecommendationRequest - struct for SPGetBudgetRulesRecommendationRequest
type SPGetBudgetRulesRecommendationRequest struct {
	SPBudgetRulesRecommendationEventRequest *SPBudgetRulesRecommendationEventRequest
}

// SPBudgetRulesRecommendationEventRequestAsSPGetBudgetRulesRecommendationRequest is a convenience function that returns SPBudgetRulesRecommendationEventRequest wrapped in SPGetBudgetRulesRecommendationRequest
func SPBudgetRulesRecommendationEventRequestAsSPGetBudgetRulesRecommendationRequest(v *SPBudgetRulesRecommendationEventRequest) SPGetBudgetRulesRecommendationRequest {
	return SPGetBudgetRulesRecommendationRequest{
		SPBudgetRulesRecommendationEventRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SPGetBudgetRulesRecommendationRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVENTS_FOR_EXISTING_CAMPAIGN'
	if jsonDict["recommendationType"] == "EVENTS_FOR_EXISTING_CAMPAIGN" {
		// try to unmarshal JSON data into SPBudgetRulesRecommendationEventRequest
		err = sonic.Unmarshal(data, &dst.SPBudgetRulesRecommendationEventRequest)
		if err == nil {
			return nil // data stored in dst.SPBudgetRulesRecommendationEventRequest, return on the first match
		} else {
			dst.SPBudgetRulesRecommendationEventRequest = nil
			return fmt.Errorf("failed to unmarshal SPGetBudgetRulesRecommendationRequest as SPBudgetRulesRecommendationEventRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SPBudgetRulesRecommendationEventRequest'
	if jsonDict["recommendationType"] == "SPBudgetRulesRecommendationEventRequest" {
		// try to unmarshal JSON data into SPBudgetRulesRecommendationEventRequest
		err = sonic.Unmarshal(data, &dst.SPBudgetRulesRecommendationEventRequest)
		if err == nil {
			return nil // data stored in dst.SPBudgetRulesRecommendationEventRequest, return on the first match
		} else {
			dst.SPBudgetRulesRecommendationEventRequest = nil
			return fmt.Errorf("failed to unmarshal SPGetBudgetRulesRecommendationRequest as SPBudgetRulesRecommendationEventRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SPGetBudgetRulesRecommendationRequest) MarshalJSON() ([]byte, error) {
	if src.SPBudgetRulesRecommendationEventRequest != nil {
		return sonic.Marshal(&src.SPBudgetRulesRecommendationEventRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SPGetBudgetRulesRecommendationRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SPBudgetRulesRecommendationEventRequest != nil {
		return obj.SPBudgetRulesRecommendationEventRequest
	}

	// all schemas are nil
	return nil
}

type NullableSPGetBudgetRulesRecommendationRequest struct {
	value *SPGetBudgetRulesRecommendationRequest
	isSet bool
}

func (v NullableSPGetBudgetRulesRecommendationRequest) Get() *SPGetBudgetRulesRecommendationRequest {
	return v.value
}

func (v *NullableSPGetBudgetRulesRecommendationRequest) Set(val *SPGetBudgetRulesRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSPGetBudgetRulesRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSPGetBudgetRulesRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPGetBudgetRulesRecommendationRequest(val *SPGetBudgetRulesRecommendationRequest) *NullableSPGetBudgetRulesRecommendationRequest {
	return &NullableSPGetBudgetRulesRecommendationRequest{value: val, isSet: true}
}

func (v NullableSPGetBudgetRulesRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPGetBudgetRulesRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
