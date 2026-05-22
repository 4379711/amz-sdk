package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the DirectPurchaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectPurchaseRequest{}

// DirectPurchaseRequest The request schema for the directPurchaseShipment operation. When the channel type is Amazon, the shipTo address is not required and will be ignored.
type DirectPurchaseRequest struct {
	ShipTo   *Address `json:"shipTo,omitempty"`
	ShipFrom *Address `json:"shipFrom,omitempty"`
	ReturnTo *Address `json:"returnTo,omitempty"`
	// A list of packages to be shipped through a shipping service offering.
	Packages            []Package                       `json:"packages,omitempty"`
	ChannelDetails      ChannelDetails                  `json:"channelDetails"`
	LabelSpecifications *RequestedDocumentSpecification `json:"labelSpecifications,omitempty"`
}

type _DirectPurchaseRequest DirectPurchaseRequest

// NewDirectPurchaseRequest instantiates a new DirectPurchaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectPurchaseRequest(channelDetails ChannelDetails) *DirectPurchaseRequest {
	this := DirectPurchaseRequest{}
	this.ChannelDetails = channelDetails
	return &this
}

// NewDirectPurchaseRequestWithDefaults instantiates a new DirectPurchaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectPurchaseRequestWithDefaults() *DirectPurchaseRequest {
	this := DirectPurchaseRequest{}
	return &this
}

// GetShipTo returns the ShipTo field value if set, zero value otherwise.
func (o *DirectPurchaseRequest) GetShipTo() Address {
	if o == nil || IsNil(o.ShipTo) {
		var ret Address
		return ret
	}
	return *o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseRequest) GetShipToOk() (*Address, bool) {
	if o == nil || IsNil(o.ShipTo) {
		return nil, false
	}
	return o.ShipTo, true
}

// HasShipTo returns a boolean if a field has been set.
func (o *DirectPurchaseRequest) HasShipTo() bool {
	if o != nil && !IsNil(o.ShipTo) {
		return true
	}

	return false
}

// SetShipTo gets a reference to the given Address and assigns it to the ShipTo field.
func (o *DirectPurchaseRequest) SetShipTo(v Address) {
	o.ShipTo = &v
}

// GetShipFrom returns the ShipFrom field value if set, zero value otherwise.
func (o *DirectPurchaseRequest) GetShipFrom() Address {
	if o == nil || IsNil(o.ShipFrom) {
		var ret Address
		return ret
	}
	return *o.ShipFrom
}

// GetShipFromOk returns a tuple with the ShipFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseRequest) GetShipFromOk() (*Address, bool) {
	if o == nil || IsNil(o.ShipFrom) {
		return nil, false
	}
	return o.ShipFrom, true
}

// HasShipFrom returns a boolean if a field has been set.
func (o *DirectPurchaseRequest) HasShipFrom() bool {
	if o != nil && !IsNil(o.ShipFrom) {
		return true
	}

	return false
}

// SetShipFrom gets a reference to the given Address and assigns it to the ShipFrom field.
func (o *DirectPurchaseRequest) SetShipFrom(v Address) {
	o.ShipFrom = &v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *DirectPurchaseRequest) GetReturnTo() Address {
	if o == nil || IsNil(o.ReturnTo) {
		var ret Address
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseRequest) GetReturnToOk() (*Address, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *DirectPurchaseRequest) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given Address and assigns it to the ReturnTo field.
func (o *DirectPurchaseRequest) SetReturnTo(v Address) {
	o.ReturnTo = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *DirectPurchaseRequest) GetPackages() []Package {
	if o == nil || IsNil(o.Packages) {
		var ret []Package
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseRequest) GetPackagesOk() ([]Package, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *DirectPurchaseRequest) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []Package and assigns it to the Packages field.
func (o *DirectPurchaseRequest) SetPackages(v []Package) {
	o.Packages = v
}

// GetChannelDetails returns the ChannelDetails field value
func (o *DirectPurchaseRequest) GetChannelDetails() ChannelDetails {
	if o == nil {
		var ret ChannelDetails
		return ret
	}

	return o.ChannelDetails
}

// GetChannelDetailsOk returns a tuple with the ChannelDetails field value
// and a boolean to check if the value has been set.
func (o *DirectPurchaseRequest) GetChannelDetailsOk() (*ChannelDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelDetails, true
}

// SetChannelDetails sets field value
func (o *DirectPurchaseRequest) SetChannelDetails(v ChannelDetails) {
	o.ChannelDetails = v
}

// GetLabelSpecifications returns the LabelSpecifications field value if set, zero value otherwise.
func (o *DirectPurchaseRequest) GetLabelSpecifications() RequestedDocumentSpecification {
	if o == nil || IsNil(o.LabelSpecifications) {
		var ret RequestedDocumentSpecification
		return ret
	}
	return *o.LabelSpecifications
}

// GetLabelSpecificationsOk returns a tuple with the LabelSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPurchaseRequest) GetLabelSpecificationsOk() (*RequestedDocumentSpecification, bool) {
	if o == nil || IsNil(o.LabelSpecifications) {
		return nil, false
	}
	return o.LabelSpecifications, true
}

// HasLabelSpecifications returns a boolean if a field has been set.
func (o *DirectPurchaseRequest) HasLabelSpecifications() bool {
	if o != nil && !IsNil(o.LabelSpecifications) {
		return true
	}

	return false
}

// SetLabelSpecifications gets a reference to the given RequestedDocumentSpecification and assigns it to the LabelSpecifications field.
func (o *DirectPurchaseRequest) SetLabelSpecifications(v RequestedDocumentSpecification) {
	o.LabelSpecifications = &v
}

func (o DirectPurchaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipTo) {
		toSerialize["shipTo"] = o.ShipTo
	}
	if !IsNil(o.ShipFrom) {
		toSerialize["shipFrom"] = o.ShipFrom
	}
	if !IsNil(o.ReturnTo) {
		toSerialize["returnTo"] = o.ReturnTo
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	toSerialize["channelDetails"] = o.ChannelDetails
	if !IsNil(o.LabelSpecifications) {
		toSerialize["labelSpecifications"] = o.LabelSpecifications
	}
	return toSerialize, nil
}

type NullableDirectPurchaseRequest struct {
	value *DirectPurchaseRequest
	isSet bool
}

func (v NullableDirectPurchaseRequest) Get() *DirectPurchaseRequest {
	return v.value
}

func (v *NullableDirectPurchaseRequest) Set(val *DirectPurchaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectPurchaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectPurchaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectPurchaseRequest(val *DirectPurchaseRequest) *NullableDirectPurchaseRequest {
	return &NullableDirectPurchaseRequest{value: val, isSet: true}
}

func (v NullableDirectPurchaseRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDirectPurchaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
