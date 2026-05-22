package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the CompetitiveSummaryBatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompetitiveSummaryBatchResponse{}

// CompetitiveSummaryBatchResponse The response schema for the `competitiveSummaryBatch` operation.
type CompetitiveSummaryBatchResponse struct {
	// The response list for the `competitiveSummaryBatch` operation.
	Responses []CompetitiveSummaryResponse `json:"responses"`
}

type _CompetitiveSummaryBatchResponse CompetitiveSummaryBatchResponse

// NewCompetitiveSummaryBatchResponse instantiates a new CompetitiveSummaryBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompetitiveSummaryBatchResponse(responses []CompetitiveSummaryResponse) *CompetitiveSummaryBatchResponse {
	this := CompetitiveSummaryBatchResponse{}
	this.Responses = responses
	return &this
}

// NewCompetitiveSummaryBatchResponseWithDefaults instantiates a new CompetitiveSummaryBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompetitiveSummaryBatchResponseWithDefaults() *CompetitiveSummaryBatchResponse {
	this := CompetitiveSummaryBatchResponse{}
	return &this
}

// GetResponses returns the Responses field value
func (o *CompetitiveSummaryBatchResponse) GetResponses() []CompetitiveSummaryResponse {
	if o == nil {
		var ret []CompetitiveSummaryResponse
		return ret
	}

	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryBatchResponse) GetResponsesOk() ([]CompetitiveSummaryResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Responses, true
}

// SetResponses sets field value
func (o *CompetitiveSummaryBatchResponse) SetResponses(v []CompetitiveSummaryResponse) {
	o.Responses = v
}

func (o CompetitiveSummaryBatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["responses"] = o.Responses
	return toSerialize, nil
}

type NullableCompetitiveSummaryBatchResponse struct {
	value *CompetitiveSummaryBatchResponse
	isSet bool
}

func (v NullableCompetitiveSummaryBatchResponse) Get() *CompetitiveSummaryBatchResponse {
	return v.value
}

func (v *NullableCompetitiveSummaryBatchResponse) Set(val *CompetitiveSummaryBatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompetitiveSummaryBatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompetitiveSummaryBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompetitiveSummaryBatchResponse(val *CompetitiveSummaryBatchResponse) *NullableCompetitiveSummaryBatchResponse {
	return &NullableCompetitiveSummaryBatchResponse{value: val, isSet: true}
}

func (v NullableCompetitiveSummaryBatchResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompetitiveSummaryBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
