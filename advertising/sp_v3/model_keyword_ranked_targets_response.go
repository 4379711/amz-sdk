package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordRankedTargetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordRankedTargetsResponse{}

// KeywordRankedTargetsResponse struct for KeywordRankedTargetsResponse
type KeywordRankedTargetsResponse struct {
	// A list of ranked keyword targets
	Targets []KeywordTargetResponse `json:"targets,omitempty"`
}

// NewKeywordRankedTargetsResponse instantiates a new KeywordRankedTargetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordRankedTargetsResponse() *KeywordRankedTargetsResponse {
	this := KeywordRankedTargetsResponse{}
	return &this
}

// NewKeywordRankedTargetsResponseWithDefaults instantiates a new KeywordRankedTargetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordRankedTargetsResponseWithDefaults() *KeywordRankedTargetsResponse {
	this := KeywordRankedTargetsResponse{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *KeywordRankedTargetsResponse) GetTargets() []KeywordTargetResponse {
	if o == nil || IsNil(o.Targets) {
		var ret []KeywordTargetResponse
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordRankedTargetsResponse) GetTargetsOk() ([]KeywordTargetResponse, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *KeywordRankedTargetsResponse) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []KeywordTargetResponse and assigns it to the Targets field.
func (o *KeywordRankedTargetsResponse) SetTargets(v []KeywordTargetResponse) {
	o.Targets = v
}

func (o KeywordRankedTargetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableKeywordRankedTargetsResponse struct {
	value *KeywordRankedTargetsResponse
	isSet bool
}

func (v NullableKeywordRankedTargetsResponse) Get() *KeywordRankedTargetsResponse {
	return v.value
}

func (v *NullableKeywordRankedTargetsResponse) Set(val *KeywordRankedTargetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordRankedTargetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordRankedTargetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordRankedTargetsResponse(val *KeywordRankedTargetsResponse) *NullableKeywordRankedTargetsResponse {
	return &NullableKeywordRankedTargetsResponse{value: val, isSet: true}
}

func (v NullableKeywordRankedTargetsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordRankedTargetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
