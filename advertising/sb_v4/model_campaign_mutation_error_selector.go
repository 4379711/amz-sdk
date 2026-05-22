package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignMutationErrorSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMutationErrorSelector{}

// CampaignMutationErrorSelector struct for CampaignMutationErrorSelector
type CampaignMutationErrorSelector struct {
	DateError    *DateError    `json:"dateError,omitempty"`
	BiddingError *BiddingError `json:"biddingError,omitempty"`
	BudgetError  *BudgetError  `json:"budgetError,omitempty"`
	BillingError *BillingError `json:"billingError,omitempty"`
	RangeError   *RangeError   `json:"rangeError,omitempty"`
	OtherError   *OtherError   `json:"otherError,omitempty"`
}

// NewCampaignMutationErrorSelector instantiates a new CampaignMutationErrorSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMutationErrorSelector() *CampaignMutationErrorSelector {
	this := CampaignMutationErrorSelector{}
	return &this
}

// NewCampaignMutationErrorSelectorWithDefaults instantiates a new CampaignMutationErrorSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMutationErrorSelectorWithDefaults() *CampaignMutationErrorSelector {
	this := CampaignMutationErrorSelector{}
	return &this
}

// GetDateError returns the DateError field value if set, zero value otherwise.
func (o *CampaignMutationErrorSelector) GetDateError() DateError {
	if o == nil || IsNil(o.DateError) {
		var ret DateError
		return ret
	}
	return *o.DateError
}

// GetDateErrorOk returns a tuple with the DateError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationErrorSelector) GetDateErrorOk() (*DateError, bool) {
	if o == nil || IsNil(o.DateError) {
		return nil, false
	}
	return o.DateError, true
}

// HasDateError returns a boolean if a field has been set.
func (o *CampaignMutationErrorSelector) HasDateError() bool {
	if o != nil && !IsNil(o.DateError) {
		return true
	}

	return false
}

// SetDateError gets a reference to the given DateError and assigns it to the DateError field.
func (o *CampaignMutationErrorSelector) SetDateError(v DateError) {
	o.DateError = &v
}

// GetBiddingError returns the BiddingError field value if set, zero value otherwise.
func (o *CampaignMutationErrorSelector) GetBiddingError() BiddingError {
	if o == nil || IsNil(o.BiddingError) {
		var ret BiddingError
		return ret
	}
	return *o.BiddingError
}

// GetBiddingErrorOk returns a tuple with the BiddingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationErrorSelector) GetBiddingErrorOk() (*BiddingError, bool) {
	if o == nil || IsNil(o.BiddingError) {
		return nil, false
	}
	return o.BiddingError, true
}

// HasBiddingError returns a boolean if a field has been set.
func (o *CampaignMutationErrorSelector) HasBiddingError() bool {
	if o != nil && !IsNil(o.BiddingError) {
		return true
	}

	return false
}

// SetBiddingError gets a reference to the given BiddingError and assigns it to the BiddingError field.
func (o *CampaignMutationErrorSelector) SetBiddingError(v BiddingError) {
	o.BiddingError = &v
}

// GetBudgetError returns the BudgetError field value if set, zero value otherwise.
func (o *CampaignMutationErrorSelector) GetBudgetError() BudgetError {
	if o == nil || IsNil(o.BudgetError) {
		var ret BudgetError
		return ret
	}
	return *o.BudgetError
}

// GetBudgetErrorOk returns a tuple with the BudgetError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationErrorSelector) GetBudgetErrorOk() (*BudgetError, bool) {
	if o == nil || IsNil(o.BudgetError) {
		return nil, false
	}
	return o.BudgetError, true
}

// HasBudgetError returns a boolean if a field has been set.
func (o *CampaignMutationErrorSelector) HasBudgetError() bool {
	if o != nil && !IsNil(o.BudgetError) {
		return true
	}

	return false
}

// SetBudgetError gets a reference to the given BudgetError and assigns it to the BudgetError field.
func (o *CampaignMutationErrorSelector) SetBudgetError(v BudgetError) {
	o.BudgetError = &v
}

// GetBillingError returns the BillingError field value if set, zero value otherwise.
func (o *CampaignMutationErrorSelector) GetBillingError() BillingError {
	if o == nil || IsNil(o.BillingError) {
		var ret BillingError
		return ret
	}
	return *o.BillingError
}

// GetBillingErrorOk returns a tuple with the BillingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationErrorSelector) GetBillingErrorOk() (*BillingError, bool) {
	if o == nil || IsNil(o.BillingError) {
		return nil, false
	}
	return o.BillingError, true
}

// HasBillingError returns a boolean if a field has been set.
func (o *CampaignMutationErrorSelector) HasBillingError() bool {
	if o != nil && !IsNil(o.BillingError) {
		return true
	}

	return false
}

// SetBillingError gets a reference to the given BillingError and assigns it to the BillingError field.
func (o *CampaignMutationErrorSelector) SetBillingError(v BillingError) {
	o.BillingError = &v
}

// GetRangeError returns the RangeError field value if set, zero value otherwise.
func (o *CampaignMutationErrorSelector) GetRangeError() RangeError {
	if o == nil || IsNil(o.RangeError) {
		var ret RangeError
		return ret
	}
	return *o.RangeError
}

// GetRangeErrorOk returns a tuple with the RangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationErrorSelector) GetRangeErrorOk() (*RangeError, bool) {
	if o == nil || IsNil(o.RangeError) {
		return nil, false
	}
	return o.RangeError, true
}

// HasRangeError returns a boolean if a field has been set.
func (o *CampaignMutationErrorSelector) HasRangeError() bool {
	if o != nil && !IsNil(o.RangeError) {
		return true
	}

	return false
}

// SetRangeError gets a reference to the given RangeError and assigns it to the RangeError field.
func (o *CampaignMutationErrorSelector) SetRangeError(v RangeError) {
	o.RangeError = &v
}

// GetOtherError returns the OtherError field value if set, zero value otherwise.
func (o *CampaignMutationErrorSelector) GetOtherError() OtherError {
	if o == nil || IsNil(o.OtherError) {
		var ret OtherError
		return ret
	}
	return *o.OtherError
}

// GetOtherErrorOk returns a tuple with the OtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutationErrorSelector) GetOtherErrorOk() (*OtherError, bool) {
	if o == nil || IsNil(o.OtherError) {
		return nil, false
	}
	return o.OtherError, true
}

// HasOtherError returns a boolean if a field has been set.
func (o *CampaignMutationErrorSelector) HasOtherError() bool {
	if o != nil && !IsNil(o.OtherError) {
		return true
	}

	return false
}

// SetOtherError gets a reference to the given OtherError and assigns it to the OtherError field.
func (o *CampaignMutationErrorSelector) SetOtherError(v OtherError) {
	o.OtherError = &v
}

func (o CampaignMutationErrorSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateError) {
		toSerialize["dateError"] = o.DateError
	}
	if !IsNil(o.BiddingError) {
		toSerialize["biddingError"] = o.BiddingError
	}
	if !IsNil(o.BudgetError) {
		toSerialize["budgetError"] = o.BudgetError
	}
	if !IsNil(o.BillingError) {
		toSerialize["billingError"] = o.BillingError
	}
	if !IsNil(o.RangeError) {
		toSerialize["rangeError"] = o.RangeError
	}
	if !IsNil(o.OtherError) {
		toSerialize["otherError"] = o.OtherError
	}
	return toSerialize, nil
}

type NullableCampaignMutationErrorSelector struct {
	value *CampaignMutationErrorSelector
	isSet bool
}

func (v NullableCampaignMutationErrorSelector) Get() *CampaignMutationErrorSelector {
	return v.value
}

func (v *NullableCampaignMutationErrorSelector) Set(val *CampaignMutationErrorSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMutationErrorSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMutationErrorSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMutationErrorSelector(val *CampaignMutationErrorSelector) *NullableCampaignMutationErrorSelector {
	return &NullableCampaignMutationErrorSelector{value: val, isSet: true}
}

func (v NullableCampaignMutationErrorSelector) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignMutationErrorSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
