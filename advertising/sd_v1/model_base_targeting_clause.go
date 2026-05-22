package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BaseTargetingClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTargetingClause{}

// BaseTargetingClause struct for BaseTargetingClause
type BaseTargetingClause struct {
	State *string `json:"state,omitempty"`
	// The bid will override the adGroup bid if specified. This field is not used for negative targeting clauses. The bid must be less than the maximum allowable bid for the campaign's marketplace; for a list of maximum allowable bids, find the [\"Bid constraints by marketplace\" table in our documentation overview](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace). You cannot manually set a bid when the targeting clause's adGroup has an enabled optimization rule.
	Bid NullableFloat32 `json:"bid,omitempty"`
}

// NewBaseTargetingClause instantiates a new BaseTargetingClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTargetingClause() *BaseTargetingClause {
	this := BaseTargetingClause{}
	return &this
}

// NewBaseTargetingClauseWithDefaults instantiates a new BaseTargetingClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTargetingClauseWithDefaults() *BaseTargetingClause {
	this := BaseTargetingClause{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseTargetingClause) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTargetingClause) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseTargetingClause) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BaseTargetingClause) SetState(v string) {
	o.State = &v
}

// GetBid returns the Bid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseTargetingClause) GetBid() float32 {
	if o == nil || IsNil(o.Bid.Get()) {
		var ret float32
		return ret
	}
	return *o.Bid.Get()
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseTargetingClause) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bid.Get(), o.Bid.IsSet()
}

// HasBid returns a boolean if a field has been set.
func (o *BaseTargetingClause) HasBid() bool {
	if o != nil && o.Bid.IsSet() {
		return true
	}

	return false
}

// SetBid gets a reference to the given NullableFloat32 and assigns it to the Bid field.
func (o *BaseTargetingClause) SetBid(v float32) {
	o.Bid.Set(&v)
}

// SetBidNil sets the value for Bid to be an explicit nil
func (o *BaseTargetingClause) SetBidNil() {
	o.Bid.Set(nil)
}

// UnsetBid ensures that no value is present for Bid, not even an explicit nil
func (o *BaseTargetingClause) UnsetBid() {
	o.Bid.Unset()
}

func (o BaseTargetingClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Bid.IsSet() {
		toSerialize["bid"] = o.Bid.Get()
	}
	return toSerialize, nil
}

type NullableBaseTargetingClause struct {
	value *BaseTargetingClause
	isSet bool
}

func (v NullableBaseTargetingClause) Get() *BaseTargetingClause {
	return v.value
}

func (v *NullableBaseTargetingClause) Set(val *BaseTargetingClause) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTargetingClause) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTargetingClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTargetingClause(val *BaseTargetingClause) *NullableBaseTargetingClause {
	return &NullableBaseTargetingClause{value: val, isSet: true}
}

func (v NullableBaseTargetingClause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBaseTargetingClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
