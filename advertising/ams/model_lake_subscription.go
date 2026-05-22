package ams

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the LakeSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LakeSubscription{}

// LakeSubscription struct for LakeSubscription
type LakeSubscription struct {
	// AWS ARN of the principal associated with the subscription. Supported principal types: - AWS Account Principal
	Principal string `json:"principal"`
	// ISO8601 Timestamp
	CreatedDate time.Time `json:"createdDate"`
	// Additional details associated with the subscription
	Notes *string `json:"notes,omitempty"`
	// Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets.
	DataSetId string `json:"dataSetId"`
	// ISO8601 Timestamp
	UpdatedDate time.Time `json:"updatedDate"`
	// Unique subscription identifier
	SubscriptionId string                   `json:"subscriptionId"`
	Status         SubscriptionEntityStatus `json:"status"`
}

type _LakeSubscription LakeSubscription

// NewLakeSubscription instantiates a new LakeSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLakeSubscription(principal string, createdDate time.Time, dataSetId string, updatedDate time.Time, subscriptionId string, status SubscriptionEntityStatus) *LakeSubscription {
	this := LakeSubscription{}
	this.Principal = principal
	this.CreatedDate = createdDate
	this.DataSetId = dataSetId
	this.UpdatedDate = updatedDate
	this.SubscriptionId = subscriptionId
	this.Status = status
	return &this
}

// NewLakeSubscriptionWithDefaults instantiates a new LakeSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLakeSubscriptionWithDefaults() *LakeSubscription {
	this := LakeSubscription{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *LakeSubscription) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *LakeSubscription) SetPrincipal(v string) {
	o.Principal = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *LakeSubscription) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *LakeSubscription) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *LakeSubscription) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *LakeSubscription) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *LakeSubscription) SetNotes(v string) {
	o.Notes = &v
}

// GetDataSetId returns the DataSetId field value
func (o *LakeSubscription) GetDataSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetDataSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *LakeSubscription) SetDataSetId(v string) {
	o.DataSetId = v
}

// GetUpdatedDate returns the UpdatedDate field value
func (o *LakeSubscription) GetUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedDate, true
}

// SetUpdatedDate sets field value
func (o *LakeSubscription) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *LakeSubscription) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *LakeSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetStatus returns the Status field value
func (o *LakeSubscription) GetStatus() SubscriptionEntityStatus {
	if o == nil {
		var ret SubscriptionEntityStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LakeSubscription) GetStatusOk() (*SubscriptionEntityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LakeSubscription) SetStatus(v SubscriptionEntityStatus) {
	o.Status = v
}

func (o LakeSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["createdDate"] = o.CreatedDate
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["dataSetId"] = o.DataSetId
	toSerialize["updatedDate"] = o.UpdatedDate
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableLakeSubscription struct {
	value *LakeSubscription
	isSet bool
}

func (v NullableLakeSubscription) Get() *LakeSubscription {
	return v.value
}

func (v *NullableLakeSubscription) Set(val *LakeSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableLakeSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableLakeSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLakeSubscription(val *LakeSubscription) *NullableLakeSubscription {
	return &NullableLakeSubscription{value: val, isSet: true}
}

func (v NullableLakeSubscription) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLakeSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
