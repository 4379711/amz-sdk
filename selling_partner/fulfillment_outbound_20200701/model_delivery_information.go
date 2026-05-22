package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the DeliveryInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryInformation{}

// DeliveryInformation The delivery information for the package. This information is available after the package is delivered.
type DeliveryInformation struct {
	// A list of delivery documents for a package.
	DeliveryDocumentList []DeliveryDocument `json:"deliveryDocumentList,omitempty"`
	DropOffLocation      *DropOffLocation   `json:"dropOffLocation,omitempty"`
}

// NewDeliveryInformation instantiates a new DeliveryInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryInformation() *DeliveryInformation {
	this := DeliveryInformation{}
	return &this
}

// NewDeliveryInformationWithDefaults instantiates a new DeliveryInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryInformationWithDefaults() *DeliveryInformation {
	this := DeliveryInformation{}
	return &this
}

// GetDeliveryDocumentList returns the DeliveryDocumentList field value if set, zero value otherwise.
func (o *DeliveryInformation) GetDeliveryDocumentList() []DeliveryDocument {
	if o == nil || IsNil(o.DeliveryDocumentList) {
		var ret []DeliveryDocument
		return ret
	}
	return o.DeliveryDocumentList
}

// GetDeliveryDocumentListOk returns a tuple with the DeliveryDocumentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryInformation) GetDeliveryDocumentListOk() ([]DeliveryDocument, bool) {
	if o == nil || IsNil(o.DeliveryDocumentList) {
		return nil, false
	}
	return o.DeliveryDocumentList, true
}

// HasDeliveryDocumentList returns a boolean if a field has been set.
func (o *DeliveryInformation) HasDeliveryDocumentList() bool {
	if o != nil && !IsNil(o.DeliveryDocumentList) {
		return true
	}

	return false
}

// SetDeliveryDocumentList gets a reference to the given []DeliveryDocument and assigns it to the DeliveryDocumentList field.
func (o *DeliveryInformation) SetDeliveryDocumentList(v []DeliveryDocument) {
	o.DeliveryDocumentList = v
}

// GetDropOffLocation returns the DropOffLocation field value if set, zero value otherwise.
func (o *DeliveryInformation) GetDropOffLocation() DropOffLocation {
	if o == nil || IsNil(o.DropOffLocation) {
		var ret DropOffLocation
		return ret
	}
	return *o.DropOffLocation
}

// GetDropOffLocationOk returns a tuple with the DropOffLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryInformation) GetDropOffLocationOk() (*DropOffLocation, bool) {
	if o == nil || IsNil(o.DropOffLocation) {
		return nil, false
	}
	return o.DropOffLocation, true
}

// HasDropOffLocation returns a boolean if a field has been set.
func (o *DeliveryInformation) HasDropOffLocation() bool {
	if o != nil && !IsNil(o.DropOffLocation) {
		return true
	}

	return false
}

// SetDropOffLocation gets a reference to the given DropOffLocation and assigns it to the DropOffLocation field.
func (o *DeliveryInformation) SetDropOffLocation(v DropOffLocation) {
	o.DropOffLocation = &v
}

func (o DeliveryInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryDocumentList) {
		toSerialize["deliveryDocumentList"] = o.DeliveryDocumentList
	}
	if !IsNil(o.DropOffLocation) {
		toSerialize["dropOffLocation"] = o.DropOffLocation
	}
	return toSerialize, nil
}

type NullableDeliveryInformation struct {
	value *DeliveryInformation
	isSet bool
}

func (v NullableDeliveryInformation) Get() *DeliveryInformation {
	return v.value
}

func (v *NullableDeliveryInformation) Set(val *DeliveryInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryInformation(val *DeliveryInformation) *NullableDeliveryInformation {
	return &NullableDeliveryInformation{value: val, isSet: true}
}

func (v NullableDeliveryInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
