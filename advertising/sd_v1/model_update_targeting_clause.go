package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTargetingClause{}

// UpdateTargetingClause struct for UpdateTargetingClause
type UpdateTargetingClause struct {
	State *string `json:"state,omitempty"`
	// The bid will override the adGroup bid if specified. This field is not used for negative targeting clauses. The bid must be less than the maximum allowable bid for the campaign's marketplace; for a list of maximum allowable bids, find the [\"Bid constraints by marketplace\" table in our documentation overview](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace). You cannot manually set a bid when the targeting clause's adGroup has an enabled optimization rule.
	Bid      NullableFloat32 `json:"bid,omitempty"`
	TargetId int64           `json:"targetId"`
}

type _UpdateTargetingClause UpdateTargetingClause

// NewUpdateTargetingClause instantiates a new UpdateTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTargetingClause(targetId int64) *UpdateTargetingClause {
	this := UpdateTargetingClause{}
	this.TargetId = targetId
	return &this
}

// NewUpdateTargetingClauseWithDefaults instantiates a new UpdateTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTargetingClauseWithDefaults() *UpdateTargetingClause {
	this := UpdateTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateTargetingClause) SetState(v string) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTargetingClause) GetBid() float32 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float32
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTargetingClause) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *UpdateTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat32 and assigns it to the Bid field.
func (o *UpdateTargetingClause) SetBid(v float32) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *UpdateTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *UpdateTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

// GetTargetId returns the TargetId field value
func (o *UpdateTargetingClause) GetTargetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *UpdateTargetingClause) GetTargetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *UpdateTargetingClause) SetTargetId(v int64) {
	o.TargetId = v
}

func (o UpdateTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	toSerialize["targetId"] = o.TargetId
	return toSerialize, nil
}

type NullableUpdateTargetingClause struct {
	value *UpdateTargetingClause
	isSet bool
}

func (v NullableUpdateTargetingClause) Get() *UpdateTargetingClause {
	return v.value
}

func (v *NullableUpdateTargetingClause) Set(val *UpdateTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTargetingClause(val *UpdateTargetingClause) *NullableUpdateTargetingClause {
	return &NullableUpdateTargetingClause{value: val, isSet: true}
}

func (v NullableUpdateTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
