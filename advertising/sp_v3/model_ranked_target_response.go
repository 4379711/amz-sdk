package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RankedTargetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankedTargetResponse{}

// RankedTargetResponse struct for RankedTargetResponse
type RankedTargetResponse struct {
	// A list of ranked keyword targets
	KeywordTargetList []RecKeywordTarget `json:"keywordTargetList,omitempty"`
}

// NewRankedTargetResponse instantiates a new RankedTargetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankedTargetResponse() *RankedTargetResponse {
	this := RankedTargetResponse{}
	return &this
}

// NewRankedTargetResponseWithDefaults instantiates a new RankedTargetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankedTargetResponseWithDefaults() *RankedTargetResponse {
	this := RankedTargetResponse{}
	return &this
}

// GetKeywordTargetList returns the KeywordTargetList field value if set, zero value otherwise.
func (o *RankedTargetResponse) GetKeywordTargetList() []RecKeywordTarget {
	if o == nil || IsNil(o.KeywordTargetList) {
		var ret []RecKeywordTarget
		return ret
	}
	return o.KeywordTargetList
}

// GetKeywordTargetListOk returns a tuple with the KeywordTargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedTargetResponse) GetKeywordTargetListOk() ([]RecKeywordTarget, bool) {
	if o == nil || IsNil(o.KeywordTargetList) {
		return nil, false
	}
	return o.KeywordTargetList, true
}

// HasKeywordTargetList returns a boolean if a field has been set.
func (o *RankedTargetResponse) HasKeywordTargetList() bool {
	if o != nil && !IsNil(o.KeywordTargetList) {
		return true
	}

	return false
}

// SetKeywordTargetList gets a reference to the given []RecKeywordTarget and assigns it to the KeywordTargetList field.
func (o *RankedTargetResponse) SetKeywordTargetList(v []RecKeywordTarget) {
	o.KeywordTargetList = v
}

func (o RankedTargetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordTargetList) {
		toSerialize["keywordTargetList"] = o.KeywordTargetList
	}
	return toSerialize, nil
}

type NullableRankedTargetResponse struct {
	value *RankedTargetResponse
	isSet bool
}

func (v NullableRankedTargetResponse) Get() *RankedTargetResponse {
	return v.value
}

func (v *NullableRankedTargetResponse) Set(val *RankedTargetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRankedTargetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRankedTargetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankedTargetResponse(val *RankedTargetResponse) *NullableRankedTargetResponse {
	return &NullableRankedTargetResponse{value: val, isSet: true}
}

func (v NullableRankedTargetResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRankedTargetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
