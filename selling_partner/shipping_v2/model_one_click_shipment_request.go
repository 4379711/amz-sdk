package shipping_v2

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the OneClickShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneClickShipmentRequest{}

// OneClickShipmentRequest The request schema for the OneClickShipment operation. When the channelType is not Amazon, shipTo is required and when channelType is Amazon shipTo is ignored.
type OneClickShipmentRequest struct {
	ShipTo   *Address `json:"shipTo,omitempty"`
	ShipFrom Address  `json:"shipFrom"`
	ReturnTo *Address `json:"returnTo,omitempty"`
	// The ship date and time (the requested pickup). This defaults to the current date and time.
	ShipDate   *time.Time  `json:"shipDate,omitempty"`
	GoodsOwner *GoodsOwner `json:"goodsOwner,omitempty"`
	// A list of packages to be shipped through a shipping service offering.
	Packages []Package `json:"packages"`
	// The value-added services to be added to a shipping service purchase.
	ValueAddedServicesDetails []OneClickShipmentValueAddedService `json:"valueAddedServicesDetails,omitempty"`
	// A list of tax detail information.
	TaxDetails                    []TaxDetail                    `json:"taxDetails,omitempty"`
	ChannelDetails                ChannelDetails                 `json:"channelDetails"`
	LabelSpecifications           RequestedDocumentSpecification `json:"labelSpecifications"`
	ServiceSelection              ServiceSelection               `json:"serviceSelection"`
	ShipperInstruction            *ShipperInstruction            `json:"shipperInstruction,omitempty"`
	DestinationAccessPointDetails *AccessPointDetails            `json:"destinationAccessPointDetails,omitempty"`
}

type _OneClickShipmentRequest OneClickShipmentRequest

// NewOneClickShipmentRequest instantiates a new OneClickShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneClickShipmentRequest(shipFrom Address, packages []Package, channelDetails ChannelDetails, labelSpecifications RequestedDocumentSpecification, serviceSelection ServiceSelection) *OneClickShipmentRequest {
	this := OneClickShipmentRequest{}
	this.ShipFrom = shipFrom
	this.Packages = packages
	this.ChannelDetails = channelDetails
	this.LabelSpecifications = labelSpecifications
	this.ServiceSelection = serviceSelection
	return &this
}

// NewOneClickShipmentRequestWithDefaults instantiates a new OneClickShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneClickShipmentRequestWithDefaults() *OneClickShipmentRequest {
	this := OneClickShipmentRequest{}
	return &this
}

// GetShipTo returns the ShipTo field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetShipTo() Address {
	if o == nil || IsNil(o.ShipTo) {
		var ret Address
		return ret
	}
	return *o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetShipToOk() (*Address, bool) {
	if o == nil || IsNil(o.ShipTo) {
		return nil, false
	}
	return o.ShipTo, true
}

// HasShipTo returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasShipTo() bool {
	if o != nil && !IsNil(o.ShipTo) {
		return true
	}

	return false
}

// SetShipTo gets a reference to the given Address and assigns it to the ShipTo field.
func (o *OneClickShipmentRequest) SetShipTo(v Address) {
	o.ShipTo = &v
}

// GetShipFrom returns the ShipFrom field value
func (o *OneClickShipmentRequest) GetShipFrom() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipFrom
}

// GetShipFromOk returns a tuple with the ShipFrom field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetShipFromOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFrom, true
}

// SetShipFrom sets field value
func (o *OneClickShipmentRequest) SetShipFrom(v Address) {
	o.ShipFrom = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetReturnTo() Address {
	if o == nil || IsNil(o.ReturnTo) {
		var ret Address
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetReturnToOk() (*Address, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given Address and assigns it to the ReturnTo field.
func (o *OneClickShipmentRequest) SetReturnTo(v Address) {
	o.ReturnTo = &v
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetShipDate() time.Time {
	if o == nil || IsNil(o.ShipDate) {
		var ret time.Time
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetShipDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShipDate) {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasShipDate() bool {
	if o != nil && !IsNil(o.ShipDate) {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given time.Time and assigns it to the ShipDate field.
func (o *OneClickShipmentRequest) SetShipDate(v time.Time) {
	o.ShipDate = &v
}

// GetGoodsOwner returns the GoodsOwner field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetGoodsOwner() GoodsOwner {
	if o == nil || IsNil(o.GoodsOwner) {
		var ret GoodsOwner
		return ret
	}
	return *o.GoodsOwner
}

// GetGoodsOwnerOk returns a tuple with the GoodsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetGoodsOwnerOk() (*GoodsOwner, bool) {
	if o == nil || IsNil(o.GoodsOwner) {
		return nil, false
	}
	return o.GoodsOwner, true
}

// HasGoodsOwner returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasGoodsOwner() bool {
	if o != nil && !IsNil(o.GoodsOwner) {
		return true
	}

	return false
}

// SetGoodsOwner gets a reference to the given GoodsOwner and assigns it to the GoodsOwner field.
func (o *OneClickShipmentRequest) SetGoodsOwner(v GoodsOwner) {
	o.GoodsOwner = &v
}

// GetPackages returns the Packages field value
func (o *OneClickShipmentRequest) GetPackages() []Package {
	if o == nil {
		var ret []Package
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetPackagesOk() ([]Package, bool) {
	if o == nil {
		return nil, false
	}
	return o.Packages, true
}

// SetPackages sets field value
func (o *OneClickShipmentRequest) SetPackages(v []Package) {
	o.Packages = v
}

// GetValueAddedServicesDetails returns the ValueAddedServicesDetails field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetValueAddedServicesDetails() []OneClickShipmentValueAddedService {
	if o == nil || IsNil(o.ValueAddedServicesDetails) {
		var ret []OneClickShipmentValueAddedService
		return ret
	}
	return o.ValueAddedServicesDetails
}

// GetValueAddedServicesDetailsOk returns a tuple with the ValueAddedServicesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetValueAddedServicesDetailsOk() ([]OneClickShipmentValueAddedService, bool) {
	if o == nil || IsNil(o.ValueAddedServicesDetails) {
		return nil, false
	}
	return o.ValueAddedServicesDetails, true
}

// HasValueAddedServicesDetails returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasValueAddedServicesDetails() bool {
	if o != nil && !IsNil(o.ValueAddedServicesDetails) {
		return true
	}

	return false
}

// SetValueAddedServicesDetails gets a reference to the given []OneClickShipmentValueAddedService and assigns it to the ValueAddedServicesDetails field.
func (o *OneClickShipmentRequest) SetValueAddedServicesDetails(v []OneClickShipmentValueAddedService) {
	o.ValueAddedServicesDetails = v
}

// GetTaxDetails returns the TaxDetails field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetTaxDetails() []TaxDetail {
	if o == nil || IsNil(o.TaxDetails) {
		var ret []TaxDetail
		return ret
	}
	return o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetTaxDetailsOk() ([]TaxDetail, bool) {
	if o == nil || IsNil(o.TaxDetails) {
		return nil, false
	}
	return o.TaxDetails, true
}

// HasTaxDetails returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasTaxDetails() bool {
	if o != nil && !IsNil(o.TaxDetails) {
		return true
	}

	return false
}

// SetTaxDetails gets a reference to the given []TaxDetail and assigns it to the TaxDetails field.
func (o *OneClickShipmentRequest) SetTaxDetails(v []TaxDetail) {
	o.TaxDetails = v
}

// GetChannelDetails returns the ChannelDetails field value
func (o *OneClickShipmentRequest) GetChannelDetails() ChannelDetails {
	if o == nil {
		var ret ChannelDetails
		return ret
	}

	return o.ChannelDetails
}

// GetChannelDetailsOk returns a tuple with the ChannelDetails field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetChannelDetailsOk() (*ChannelDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelDetails, true
}

// SetChannelDetails sets field value
func (o *OneClickShipmentRequest) SetChannelDetails(v ChannelDetails) {
	o.ChannelDetails = v
}

// GetLabelSpecifications returns the LabelSpecifications field value
func (o *OneClickShipmentRequest) GetLabelSpecifications() RequestedDocumentSpecification {
	if o == nil {
		var ret RequestedDocumentSpecification
		return ret
	}

	return o.LabelSpecifications
}

// GetLabelSpecificationsOk returns a tuple with the LabelSpecifications field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetLabelSpecificationsOk() (*RequestedDocumentSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSpecifications, true
}

// SetLabelSpecifications sets field value
func (o *OneClickShipmentRequest) SetLabelSpecifications(v RequestedDocumentSpecification) {
	o.LabelSpecifications = v
}

// GetServiceSelection returns the ServiceSelection field value
func (o *OneClickShipmentRequest) GetServiceSelection() ServiceSelection {
	if o == nil {
		var ret ServiceSelection
		return ret
	}

	return o.ServiceSelection
}

// GetServiceSelectionOk returns a tuple with the ServiceSelection field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetServiceSelectionOk() (*ServiceSelection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceSelection, true
}

// SetServiceSelection sets field value
func (o *OneClickShipmentRequest) SetServiceSelection(v ServiceSelection) {
	o.ServiceSelection = v
}

// GetShipperInstruction returns the ShipperInstruction field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetShipperInstruction() ShipperInstruction {
	if o == nil || IsNil(o.ShipperInstruction) {
		var ret ShipperInstruction
		return ret
	}
	return *o.ShipperInstruction
}

// GetShipperInstructionOk returns a tuple with the ShipperInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetShipperInstructionOk() (*ShipperInstruction, bool) {
	if o == nil || IsNil(o.ShipperInstruction) {
		return nil, false
	}
	return o.ShipperInstruction, true
}

// HasShipperInstruction returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasShipperInstruction() bool {
	if o != nil && !IsNil(o.ShipperInstruction) {
		return true
	}

	return false
}

// SetShipperInstruction gets a reference to the given ShipperInstruction and assigns it to the ShipperInstruction field.
func (o *OneClickShipmentRequest) SetShipperInstruction(v ShipperInstruction) {
	o.ShipperInstruction = &v
}

// GetDestinationAccessPointDetails returns the DestinationAccessPointDetails field value if set, zero value otherwise.
func (o *OneClickShipmentRequest) GetDestinationAccessPointDetails() AccessPointDetails {
	if o == nil || IsNil(o.DestinationAccessPointDetails) {
		var ret AccessPointDetails
		return ret
	}
	return *o.DestinationAccessPointDetails
}

// GetDestinationAccessPointDetailsOk returns a tuple with the DestinationAccessPointDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentRequest) GetDestinationAccessPointDetailsOk() (*AccessPointDetails, bool) {
	if o == nil || IsNil(o.DestinationAccessPointDetails) {
		return nil, false
	}
	return o.DestinationAccessPointDetails, true
}

// HasDestinationAccessPointDetails returns a boolean if a field has been set.
func (o *OneClickShipmentRequest) HasDestinationAccessPointDetails() bool {
	if o != nil && !IsNil(o.DestinationAccessPointDetails) {
		return true
	}

	return false
}

// SetDestinationAccessPointDetails gets a reference to the given AccessPointDetails and assigns it to the DestinationAccessPointDetails field.
func (o *OneClickShipmentRequest) SetDestinationAccessPointDetails(v AccessPointDetails) {
	o.DestinationAccessPointDetails = &v
}

func (o OneClickShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipTo) {
		toSerialize["shipTo"] = o.ShipTo
	}
	toSerialize["shipFrom"] = o.ShipFrom
	if !IsNil(o.ReturnTo) {
		toSerialize["returnTo"] = o.ReturnTo
	}
	if !IsNil(o.ShipDate) {
		toSerialize["shipDate"] = o.ShipDate
	}
	if !IsNil(o.GoodsOwner) {
		toSerialize["goodsOwner"] = o.GoodsOwner
	}
	toSerialize["packages"] = o.Packages
	if !IsNil(o.ValueAddedServicesDetails) {
		toSerialize["valueAddedServicesDetails"] = o.ValueAddedServicesDetails
	}
	if !IsNil(o.TaxDetails) {
		toSerialize["taxDetails"] = o.TaxDetails
	}
	toSerialize["channelDetails"] = o.ChannelDetails
	toSerialize["labelSpecifications"] = o.LabelSpecifications
	toSerialize["serviceSelection"] = o.ServiceSelection
	if !IsNil(o.ShipperInstruction) {
		toSerialize["shipperInstruction"] = o.ShipperInstruction
	}
	if !IsNil(o.DestinationAccessPointDetails) {
		toSerialize["destinationAccessPointDetails"] = o.DestinationAccessPointDetails
	}
	return toSerialize, nil
}

type NullableOneClickShipmentRequest struct {
	value *OneClickShipmentRequest
	isSet bool
}

func (v NullableOneClickShipmentRequest) Get() *OneClickShipmentRequest {
	return v.value
}

func (v *NullableOneClickShipmentRequest) Set(val *OneClickShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOneClickShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOneClickShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneClickShipmentRequest(val *OneClickShipmentRequest) *NullableOneClickShipmentRequest {
	return &NullableOneClickShipmentRequest{value: val, isSet: true}
}

func (v NullableOneClickShipmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOneClickShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
