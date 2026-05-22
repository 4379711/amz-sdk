package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateCollectionFormRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateCollectionFormRequest{}

// GenerateCollectionFormRequest The request schema Call to generate the collection form.
type GenerateCollectionFormRequest struct {
	// Object to pass additional information about the MCI Integrator shipperType: List of ClientReferenceDetail
	ClientReferenceDetails []ClientReferenceDetail `json:"clientReferenceDetails,omitempty"`
	// The carrier identifier for the offering, provided by the carrier.
	CarrierId       string   `json:"carrierId"`
	ShipFromAddress *Address `json:"shipFromAddress,omitempty"`
}

type _GenerateCollectionFormRequest GenerateCollectionFormRequest

// NewGenerateCollectionFormRequest instantiates a new GenerateCollectionFormRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateCollectionFormRequest(carrierId string) *GenerateCollectionFormRequest {
	this := GenerateCollectionFormRequest{}
	this.CarrierId = carrierId
	return &this
}

// NewGenerateCollectionFormRequestWithDefaults instantiates a new GenerateCollectionFormRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateCollectionFormRequestWithDefaults() *GenerateCollectionFormRequest {
	this := GenerateCollectionFormRequest{}
	return &this
}

// GetClientReferenceDetails returns the ClientReferenceDetails field value if set, zero value otherwise.
func (o *GenerateCollectionFormRequest) GetClientReferenceDetails() []ClientReferenceDetail {
	if o == nil || IsNil(o.ClientReferenceDetails) {
		var ret []ClientReferenceDetail
		return ret
	}
	return o.ClientReferenceDetails
}

// GetClientReferenceDetailsOk returns a tuple with the ClientReferenceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateCollectionFormRequest) GetClientReferenceDetailsOk() ([]ClientReferenceDetail, bool) {
	if o == nil || IsNil(o.ClientReferenceDetails) {
		return nil, false
	}
	return o.ClientReferenceDetails, true
}

// HasClientReferenceDetails returns a boolean if a field has been set.
func (o *GenerateCollectionFormRequest) HasClientReferenceDetails() bool {
	if o != nil && !IsNil(o.ClientReferenceDetails) {
		return true
	}

	return false
}

// SetClientReferenceDetails gets a reference to the given []ClientReferenceDetail and assigns it to the ClientReferenceDetails field.
func (o *GenerateCollectionFormRequest) SetClientReferenceDetails(v []ClientReferenceDetail) {
	o.ClientReferenceDetails = v
}

// GetCarrierId returns the CarrierId field value
func (o *GenerateCollectionFormRequest) GetCarrierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value
// and a boolean to check if the value has been set.
func (o *GenerateCollectionFormRequest) GetCarrierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierId, true
}

// SetCarrierId sets field value
func (o *GenerateCollectionFormRequest) SetCarrierId(v string) {
	o.CarrierId = v
}

// GetShipFromAddress returns the ShipFromAddress field value if set, zero value otherwise.
func (o *GenerateCollectionFormRequest) GetShipFromAddress() Address {
	if o == nil || IsNil(o.ShipFromAddress) {
		var ret Address
		return ret
	}
	return *o.ShipFromAddress
}

// GetShipFromAddressOk returns a tuple with the ShipFromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateCollectionFormRequest) GetShipFromAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.ShipFromAddress) {
		return nil, false
	}
	return o.ShipFromAddress, true
}

// HasShipFromAddress returns a boolean if a field has been set.
func (o *GenerateCollectionFormRequest) HasShipFromAddress() bool {
	if o != nil && !IsNil(o.ShipFromAddress) {
		return true
	}

	return false
}

// SetShipFromAddress gets a reference to the given Address and assigns it to the ShipFromAddress field.
func (o *GenerateCollectionFormRequest) SetShipFromAddress(v Address) {
	o.ShipFromAddress = &v
}

func (o GenerateCollectionFormRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientReferenceDetails) {
		toSerialize["clientReferenceDetails"] = o.ClientReferenceDetails
	}
	toSerialize["carrierId"] = o.CarrierId
	if !IsNil(o.ShipFromAddress) {
		toSerialize["shipFromAddress"] = o.ShipFromAddress
	}
	return toSerialize, nil
}

type NullableGenerateCollectionFormRequest struct {
	value *GenerateCollectionFormRequest
	isSet bool
}

func (v NullableGenerateCollectionFormRequest) Get() *GenerateCollectionFormRequest {
	return v.value
}

func (v *NullableGenerateCollectionFormRequest) Set(val *GenerateCollectionFormRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateCollectionFormRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateCollectionFormRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateCollectionFormRequest(val *GenerateCollectionFormRequest) *NullableGenerateCollectionFormRequest {
	return &NullableGenerateCollectionFormRequest{value: val, isSet: true}
}

func (v NullableGenerateCollectionFormRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateCollectionFormRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
