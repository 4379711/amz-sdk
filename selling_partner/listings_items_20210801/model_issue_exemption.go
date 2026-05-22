package listings_items_20210801

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the IssueExemption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueExemption{}

// IssueExemption Conveying the status of the listed enforcement actions and, if applicable, provides information about the exemption's expiry date.
type IssueExemption struct {
	// This field indicates the current exemption status for the listed enforcement actions. It can take values such as `EXEMPT`, signifying permanent exemption, `EXEMPT_UNTIL_EXPIRY_DATE` indicating temporary exemption until a specified date, or `NOT_EXEMPT` signifying no exemptions, and enforcement actions were already applied.
	Status string `json:"status"`
	// This field represents the timestamp, following the ISO 8601 format, which specifies the date when temporary exemptions, if applicable, will expire, and Amazon will begin enforcing the listed actions.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
}

type _IssueExemption IssueExemption

// NewIssueExemption instantiates a new IssueExemption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueExemption(status string) *IssueExemption {
	this := IssueExemption{}
	this.Status = status
	return &this
}

// NewIssueExemptionWithDefaults instantiates a new IssueExemption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueExemptionWithDefaults() *IssueExemption {
	this := IssueExemption{}
	return &this
}

// GetStatus returns the Status field value
func (o *IssueExemption) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IssueExemption) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IssueExemption) SetStatus(v string) {
	o.Status = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *IssueExemption) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueExemption) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *IssueExemption) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *IssueExemption) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

func (o IssueExemption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	return toSerialize, nil
}

type NullableIssueExemption struct {
	value *IssueExemption
	isSet bool
}

func (v NullableIssueExemption) Get() *IssueExemption {
	return v.value
}

func (v *NullableIssueExemption) Set(val *IssueExemption) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueExemption) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueExemption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueExemption(val *IssueExemption) *NullableIssueExemption {
	return &NullableIssueExemption{value: val, isSet: true}
}

func (v NullableIssueExemption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIssueExemption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
