package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendationNewCampaignsException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendationNewCampaignsException{}

// BudgetRecommendationNewCampaignsException struct for BudgetRecommendationNewCampaignsException
type BudgetRecommendationNewCampaignsException struct {
	Message *string `json:"message,omitempty"`
}

// NewBudgetRecommendationNewCampaignsException instantiates a new BudgetRecommendationNewCampaignsException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendationNewCampaignsException() *BudgetRecommendationNewCampaignsException {
	this := BudgetRecommendationNewCampaignsException{}
	return &this
}

// NewBudgetRecommendationNewCampaignsExceptionWithDefaults instantiates a new BudgetRecommendationNewCampaignsException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationNewCampaignsExceptionWithDefaults() *BudgetRecommendationNewCampaignsException {
	this := BudgetRecommendationNewCampaignsException{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BudgetRecommendationNewCampaignsException) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationNewCampaignsException) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BudgetRecommendationNewCampaignsException) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BudgetRecommendationNewCampaignsException) SetMessage(v string) {
	o.Message = &v
}

func (o BudgetRecommendationNewCampaignsException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableBudgetRecommendationNewCampaignsException struct {
	value *BudgetRecommendationNewCampaignsException
	isSet bool
}

func (v NullableBudgetRecommendationNewCampaignsException) Get() *BudgetRecommendationNewCampaignsException {
	return v.value
}

func (v *NullableBudgetRecommendationNewCampaignsException) Set(val *BudgetRecommendationNewCampaignsException) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendationNewCampaignsException) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendationNewCampaignsException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendationNewCampaignsException(val *BudgetRecommendationNewCampaignsException) *NullableBudgetRecommendationNewCampaignsException {
	return &NullableBudgetRecommendationNewCampaignsException{value: val, isSet: true}
}

func (v NullableBudgetRecommendationNewCampaignsException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendationNewCampaignsException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
