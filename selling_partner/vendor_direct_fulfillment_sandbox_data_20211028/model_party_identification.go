package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the PartyIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartyIdentification{}

// PartyIdentification The identification object for the party information. For example, warehouse code or vendor code. Please refer to specific party for more details.
type PartyIdentification struct {
	// Assigned identification for the party. For example, warehouse code or vendor code. Please refer to specific party for more details.
	PartyId string `json:"partyId"`
}

type _PartyIdentification PartyIdentification

// NewPartyIdentification instantiates a new PartyIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyIdentification(partyId string) *PartyIdentification {
	this := PartyIdentification{}
	this.PartyId = partyId
	return &this
}

// NewPartyIdentificationWithDefaults instantiates a new PartyIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyIdentificationWithDefaults() *PartyIdentification {
	this := PartyIdentification{}
	return &this
}

// GetPartyId returns the PartyId field value
func (o *PartyIdentification) GetPartyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartyId, true
}

// SetPartyId sets field value
func (o *PartyIdentification) SetPartyId(v string) {
	o.PartyId = v
}

func (o PartyIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partyId"] = o.PartyId
	return toSerialize, nil
}

type NullablePartyIdentification struct {
	value *PartyIdentification
	isSet bool
}

func (v NullablePartyIdentification) Get() *PartyIdentification {
	return v.value
}

func (v *NullablePartyIdentification) Set(val *PartyIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyIdentification(val *PartyIdentification) *NullablePartyIdentification {
	return &NullablePartyIdentification{value: val, isSet: true}
}

func (v NullablePartyIdentification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePartyIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
