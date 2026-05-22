package finances_v0

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the DebtRecoveryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebtRecoveryItem{}

// DebtRecoveryItem An item of a debt payment or debt adjustment.
type DebtRecoveryItem struct {
	RecoveryAmount *Currency `json:"RecoveryAmount,omitempty"`
	OriginalAmount *Currency `json:"OriginalAmount,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	GroupBeginDate *time.Time `json:"GroupBeginDate,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	GroupEndDate *time.Time `json:"GroupEndDate,omitempty"`
}

// NewDebtRecoveryItem instantiates a new DebtRecoveryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebtRecoveryItem() *DebtRecoveryItem {
	this := DebtRecoveryItem{}
	return &this
}

// NewDebtRecoveryItemWithDefaults instantiates a new DebtRecoveryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebtRecoveryItemWithDefaults() *DebtRecoveryItem {
	this := DebtRecoveryItem{}
	return &this
}

// GetRecoveryAmount returns the RecoveryAmount field value if set, zero value otherwise.
func (o *DebtRecoveryItem) GetRecoveryAmount() Currency {
	if o == nil || IsNil(o.RecoveryAmount) {
		var ret Currency
		return ret
	}
	return *o.RecoveryAmount
}

// GetRecoveryAmountOk returns a tuple with the RecoveryAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebtRecoveryItem) GetRecoveryAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.RecoveryAmount) {
		return nil, false
	}
	return o.RecoveryAmount, true
}

// HasRecoveryAmount returns a boolean if a field has been set.
func (o *DebtRecoveryItem) HasRecoveryAmount() bool {
	if o != nil && !IsNil(o.RecoveryAmount) {
		return true
	}

	return false
}

// SetRecoveryAmount gets a reference to the given Currency and assigns it to the RecoveryAmount field.
func (o *DebtRecoveryItem) SetRecoveryAmount(v Currency) {
	o.RecoveryAmount = &v
}

// GetOriginalAmount returns the OriginalAmount field value if set, zero value otherwise.
func (o *DebtRecoveryItem) GetOriginalAmount() Currency {
	if o == nil || IsNil(o.OriginalAmount) {
		var ret Currency
		return ret
	}
	return *o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebtRecoveryItem) GetOriginalAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.OriginalAmount) {
		return nil, false
	}
	return o.OriginalAmount, true
}

// HasOriginalAmount returns a boolean if a field has been set.
func (o *DebtRecoveryItem) HasOriginalAmount() bool {
	if o != nil && !IsNil(o.OriginalAmount) {
		return true
	}

	return false
}

// SetOriginalAmount gets a reference to the given Currency and assigns it to the OriginalAmount field.
func (o *DebtRecoveryItem) SetOriginalAmount(v Currency) {
	o.OriginalAmount = &v
}

// GetGroupBeginDate returns the GroupBeginDate field value if set, zero value otherwise.
func (o *DebtRecoveryItem) GetGroupBeginDate() time.Time {
	if o == nil || IsNil(o.GroupBeginDate) {
		var ret time.Time
		return ret
	}
	return *o.GroupBeginDate
}

// GetGroupBeginDateOk returns a tuple with the GroupBeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebtRecoveryItem) GetGroupBeginDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GroupBeginDate) {
		return nil, false
	}
	return o.GroupBeginDate, true
}

// HasGroupBeginDate returns a boolean if a field has been set.
func (o *DebtRecoveryItem) HasGroupBeginDate() bool {
	if o != nil && !IsNil(o.GroupBeginDate) {
		return true
	}

	return false
}

// SetGroupBeginDate gets a reference to the given time.Time and assigns it to the GroupBeginDate field.
func (o *DebtRecoveryItem) SetGroupBeginDate(v time.Time) {
	o.GroupBeginDate = &v
}

// GetGroupEndDate returns the GroupEndDate field value if set, zero value otherwise.
func (o *DebtRecoveryItem) GetGroupEndDate() time.Time {
	if o == nil || IsNil(o.GroupEndDate) {
		var ret time.Time
		return ret
	}
	return *o.GroupEndDate
}

// GetGroupEndDateOk returns a tuple with the GroupEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebtRecoveryItem) GetGroupEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GroupEndDate) {
		return nil, false
	}
	return o.GroupEndDate, true
}

// HasGroupEndDate returns a boolean if a field has been set.
func (o *DebtRecoveryItem) HasGroupEndDate() bool {
	if o != nil && !IsNil(o.GroupEndDate) {
		return true
	}

	return false
}

// SetGroupEndDate gets a reference to the given time.Time and assigns it to the GroupEndDate field.
func (o *DebtRecoveryItem) SetGroupEndDate(v time.Time) {
	o.GroupEndDate = &v
}

func (o DebtRecoveryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecoveryAmount) {
		toSerialize["RecoveryAmount"] = o.RecoveryAmount
	}
	if !IsNil(o.OriginalAmount) {
		toSerialize["OriginalAmount"] = o.OriginalAmount
	}
	if !IsNil(o.GroupBeginDate) {
		toSerialize["GroupBeginDate"] = o.GroupBeginDate
	}
	if !IsNil(o.GroupEndDate) {
		toSerialize["GroupEndDate"] = o.GroupEndDate
	}
	return toSerialize, nil
}

type NullableDebtRecoveryItem struct {
	value *DebtRecoveryItem
	isSet bool
}

func (v NullableDebtRecoveryItem) Get() *DebtRecoveryItem {
	return v.value
}

func (v *NullableDebtRecoveryItem) Set(val *DebtRecoveryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDebtRecoveryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDebtRecoveryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebtRecoveryItem(val *DebtRecoveryItem) *NullableDebtRecoveryItem {
	return &NullableDebtRecoveryItem{value: val, isSet: true}
}

func (v NullableDebtRecoveryItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDebtRecoveryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
