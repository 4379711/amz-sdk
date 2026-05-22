package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignMutationFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMutationFailureResponseItem{}

// CampaignMutationFailureResponseItem struct for CampaignMutationFailureResponseItem
type CampaignMutationFailureResponseItem struct {
	// the index of the campaign in the array from the request body.
	Index float32 `json:"index"`
	// A list of validation errors.
	Errors []CampaignMutationError `json:"errors,omitempty"`
}

type _CampaignMutationFailureResponseItem CampaignMutationFailureResponseItem

// NewCampaignMutationFailureResponseItem instantiates a new CampaignMutationFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMutationFailureResponseItem(index float32) *CampaignMutationFailureResponseItem {
	this := CampaignMutationFailureResponseItem{}
	this.Index = index
	return &this
}

// NewCampaignMutationFailureResponseItemWithDefaults instantiates a new CampaignMutationFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMutationFailureResponseItemWithDefaults() *CampaignMutationFailureResponseItem {
	this := CampaignMutationFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *CampaignMutationFailureResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *CampaignMutationFailureResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *CampaignMutationFailureResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CampaignMutationFailureResponseItem) GetErrors() []CampaignMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []CampaignMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationFailureResponseItem) GetErrorsOk() ([]CampaignMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CampaignMutationFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []CampaignMutationError and assigns it to the Errors field.
func (o *CampaignMutationFailureResponseItem) SetErrors(v []CampaignMutationError) {
	o.Errors = v
}

func (o CampaignMutationFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCampaignMutationFailureResponseItem struct {
	value *CampaignMutationFailureResponseItem
	isSet bool
}

func (v NullableCampaignMutationFailureResponseItem) Get() *CampaignMutationFailureResponseItem {
	return v.value
}

func (v *NullableCampaignMutationFailureResponseItem) Set(val *CampaignMutationFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMutationFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMutationFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMutationFailureResponseItem(val *CampaignMutationFailureResponseItem) *NullableCampaignMutationFailureResponseItem {
	return &NullableCampaignMutationFailureResponseItem{value: val, isSet: true}
}

func (v NullableCampaignMutationFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignMutationFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
