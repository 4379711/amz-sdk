package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the CompetitiveSummaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompetitiveSummaryResponse{}

// CompetitiveSummaryResponse The response for the individual `competitiveSummary` request in the batch operation.
type CompetitiveSummaryResponse struct {
	Status HttpStatusLine                 `json:"status"`
	Body   CompetitiveSummaryResponseBody `json:"body"`
}

type _CompetitiveSummaryResponse CompetitiveSummaryResponse

// NewCompetitiveSummaryResponse instantiates a new CompetitiveSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompetitiveSummaryResponse(status HttpStatusLine, body CompetitiveSummaryResponseBody) *CompetitiveSummaryResponse {
	this := CompetitiveSummaryResponse{}
	this.Status = status
	this.Body = body
	return &this
}

// NewCompetitiveSummaryResponseWithDefaults instantiates a new CompetitiveSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompetitiveSummaryResponseWithDefaults() *CompetitiveSummaryResponse {
	this := CompetitiveSummaryResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *CompetitiveSummaryResponse) GetStatus() HttpStatusLine {
	if o == nil {
		var ret HttpStatusLine
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponse) GetStatusOk() (*HttpStatusLine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CompetitiveSummaryResponse) SetStatus(v HttpStatusLine) {
	o.Status = v
}

// GetBody returns the Body field value
func (o *CompetitiveSummaryResponse) GetBody() CompetitiveSummaryResponseBody {
	if o == nil {
		var ret CompetitiveSummaryResponseBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponse) GetBodyOk() (*CompetitiveSummaryResponseBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *CompetitiveSummaryResponse) SetBody(v CompetitiveSummaryResponseBody) {
	o.Body = v
}

func (o CompetitiveSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

type NullableCompetitiveSummaryResponse struct {
	value *CompetitiveSummaryResponse
	isSet bool
}

func (v NullableCompetitiveSummaryResponse) Get() *CompetitiveSummaryResponse {
	return v.value
}

func (v *NullableCompetitiveSummaryResponse) Set(val *CompetitiveSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompetitiveSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompetitiveSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompetitiveSummaryResponse(val *CompetitiveSummaryResponse) *NullableCompetitiveSummaryResponse {
	return &NullableCompetitiveSummaryResponse{value: val, isSet: true}
}

func (v NullableCompetitiveSummaryResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompetitiveSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
