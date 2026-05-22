package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the CompetitiveSummaryBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompetitiveSummaryBatchRequest{}

// CompetitiveSummaryBatchRequest The `competitiveSummary` batch request data.
type CompetitiveSummaryBatchRequest struct {
	// A batched list of `competitiveSummary` requests.
	Requests []CompetitiveSummaryRequest `json:"requests"`
}

type _CompetitiveSummaryBatchRequest CompetitiveSummaryBatchRequest

// NewCompetitiveSummaryBatchRequest instantiates a new CompetitiveSummaryBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompetitiveSummaryBatchRequest(requests []CompetitiveSummaryRequest) *CompetitiveSummaryBatchRequest {
	this := CompetitiveSummaryBatchRequest{}
	this.Requests = requests
	return &this
}

// NewCompetitiveSummaryBatchRequestWithDefaults instantiates a new CompetitiveSummaryBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompetitiveSummaryBatchRequestWithDefaults() *CompetitiveSummaryBatchRequest {
	this := CompetitiveSummaryBatchRequest{}
	return &this
}

// GetRequests returns the Requests field value
func (o *CompetitiveSummaryBatchRequest) GetRequests() []CompetitiveSummaryRequest {
	if o == nil {
		var ret []CompetitiveSummaryRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryBatchRequest) GetRequestsOk() ([]CompetitiveSummaryRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *CompetitiveSummaryBatchRequest) SetRequests(v []CompetitiveSummaryRequest) {
	o.Requests = v
}

func (o CompetitiveSummaryBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requests"] = o.Requests
	return toSerialize, nil
}

type NullableCompetitiveSummaryBatchRequest struct {
	value *CompetitiveSummaryBatchRequest
	isSet bool
}

func (v NullableCompetitiveSummaryBatchRequest) Get() *CompetitiveSummaryBatchRequest {
	return v.value
}

func (v *NullableCompetitiveSummaryBatchRequest) Set(val *CompetitiveSummaryBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCompetitiveSummaryBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCompetitiveSummaryBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompetitiveSummaryBatchRequest(val *CompetitiveSummaryBatchRequest) *NullableCompetitiveSummaryBatchRequest {
	return &NullableCompetitiveSummaryBatchRequest{value: val, isSet: true}
}

func (v NullableCompetitiveSummaryBatchRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompetitiveSummaryBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
