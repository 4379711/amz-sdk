package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ModerationResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModerationResultsResponse{}

// ModerationResultsResponse struct for ModerationResultsResponse
type ModerationResultsResponse struct {
	ModerationResults []ModerationResult `json:"moderationResults,omitempty"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewModerationResultsResponse instantiates a new ModerationResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerationResultsResponse() *ModerationResultsResponse {
	this := ModerationResultsResponse{}
	return &this
}

// NewModerationResultsResponseWithDefaults instantiates a new ModerationResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerationResultsResponseWithDefaults() *ModerationResultsResponse {
	this := ModerationResultsResponse{}
	return &this
}

// GetModerationResults returns the ModerationResults field value if set, zero value otherwise.
func (o *ModerationResultsResponse) GetModerationResults() []ModerationResult {
	if o == nil || IsNil(o.ModerationResults) {
		var ret []ModerationResult
		return ret
	}
	return o.ModerationResults
}

// GetModerationResultsOk returns a tuple with the ModerationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsResponse) GetModerationResultsOk() ([]ModerationResult, bool) {
	if o == nil || IsNil(o.ModerationResults) {
		return nil, false
	}
	return o.ModerationResults, true
}

// HasModerationResults returns a boolean if a field has been set.
func (o *ModerationResultsResponse) HasModerationResults() bool {
	if o != nil && !IsNil(o.ModerationResults) {
		return true
	}

	return false
}

// SetModerationResults gets a reference to the given []ModerationResult and assigns it to the ModerationResults field.
func (o *ModerationResultsResponse) SetModerationResults(v []ModerationResult) {
	o.ModerationResults = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ModerationResultsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ModerationResultsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ModerationResultsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ModerationResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModerationResults) {
		toSerialize["moderationResults"] = o.ModerationResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableModerationResultsResponse struct {
	value *ModerationResultsResponse
	isSet bool
}

func (v NullableModerationResultsResponse) Get() *ModerationResultsResponse {
	return v.value
}

func (v *NullableModerationResultsResponse) Set(val *ModerationResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationResultsResponse(val *ModerationResultsResponse) *NullableModerationResultsResponse {
	return &NullableModerationResultsResponse{value: val, isSet: true}
}

func (v NullableModerationResultsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
