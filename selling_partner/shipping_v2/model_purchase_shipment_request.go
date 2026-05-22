package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseShipmentRequest{}

// PurchaseShipmentRequest The request schema for the purchaseShipment operation.
type PurchaseShipmentRequest struct {
	// A unique token generated to identify a getRates operation.
	RequestToken string `json:"requestToken"`
	// An identifier for the rate (shipment offering) provided by a shipping service provider.
	RateId                         string                         `json:"rateId"`
	RequestedDocumentSpecification RequestedDocumentSpecification `json:"requestedDocumentSpecification"`
	// The value-added services to be added to a shipping service purchase.
	RequestedValueAddedServices []RequestedValueAddedService `json:"requestedValueAddedServices,omitempty"`
	// The additional inputs required to purchase a shipping offering, in JSON format. The JSON provided here must adhere to the JSON schema that is returned in the response to the getAdditionalInputs operation.  Additional inputs are only required when indicated by the requiresAdditionalInputs property in the response to the getRates operation.
	AdditionalInputs map[string]map[string]interface{} `json:"additionalInputs,omitempty"`
}

type _PurchaseShipmentRequest PurchaseShipmentRequest

// NewPurchaseShipmentRequest instantiates a new PurchaseShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseShipmentRequest(requestToken string, rateId string, requestedDocumentSpecification RequestedDocumentSpecification) *PurchaseShipmentRequest {
	this := PurchaseShipmentRequest{}
	this.RequestToken = requestToken
	this.RateId = rateId
	this.RequestedDocumentSpecification = requestedDocumentSpecification
	return &this
}

// NewPurchaseShipmentRequestWithDefaults instantiates a new PurchaseShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseShipmentRequestWithDefaults() *PurchaseShipmentRequest {
	this := PurchaseShipmentRequest{}
	return &this
}

// GetRequestToken returns the RequestToken field value
func (o *PurchaseShipmentRequest) GetRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestToken, true
}

// SetRequestToken sets field value
func (o *PurchaseShipmentRequest) SetRequestToken(v string) {
	o.RequestToken = v
}

// GetRateId returns the RateId field value
func (o *PurchaseShipmentRequest) GetRateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RateId
}

// GetRateIdOk returns a tuple with the RateId field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetRateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateId, true
}

// SetRateId sets field value
func (o *PurchaseShipmentRequest) SetRateId(v string) {
	o.RateId = v
}

// GetRequestedDocumentSpecification returns the RequestedDocumentSpecification field value
func (o *PurchaseShipmentRequest) GetRequestedDocumentSpecification() RequestedDocumentSpecification {
	if o == nil {
		var ret RequestedDocumentSpecification
		return ret
	}

	return o.RequestedDocumentSpecification
}

// GetRequestedDocumentSpecificationOk returns a tuple with the RequestedDocumentSpecification field value
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetRequestedDocumentSpecificationOk() (*RequestedDocumentSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedDocumentSpecification, true
}

// SetRequestedDocumentSpecification sets field value
func (o *PurchaseShipmentRequest) SetRequestedDocumentSpecification(v RequestedDocumentSpecification) {
	o.RequestedDocumentSpecification = v
}

// GetRequestedValueAddedServices returns the RequestedValueAddedServices field value if set, zero value otherwise.
func (o *PurchaseShipmentRequest) GetRequestedValueAddedServices() []RequestedValueAddedService {
	if o == nil || IsNil(o.RequestedValueAddedServices) {
		var ret []RequestedValueAddedService
		return ret
	}
	return o.RequestedValueAddedServices
}

// GetRequestedValueAddedServicesOk returns a tuple with the RequestedValueAddedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetRequestedValueAddedServicesOk() ([]RequestedValueAddedService, bool) {
	if o == nil || IsNil(o.RequestedValueAddedServices) {
		return nil, false
	}
	return o.RequestedValueAddedServices, true
}

// HasRequestedValueAddedServices returns a boolean if a field has been set.
func (o *PurchaseShipmentRequest) HasRequestedValueAddedServices() bool {
	if o != nil && !IsNil(o.RequestedValueAddedServices) {
		return true
	}

	return false
}

// SetRequestedValueAddedServices gets a reference to the given []RequestedValueAddedService and assigns it to the RequestedValueAddedServices field.
func (o *PurchaseShipmentRequest) SetRequestedValueAddedServices(v []RequestedValueAddedService) {
	o.RequestedValueAddedServices = v
}

// GetAdditionalInputs returns the AdditionalInputs field value if set, zero value otherwise.
func (o *PurchaseShipmentRequest) GetAdditionalInputs() map[string]map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInputs) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.AdditionalInputs
}

// GetAdditionalInputsOk returns a tuple with the AdditionalInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseShipmentRequest) GetAdditionalInputsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInputs) {
		return map[string]map[string]interface{}{}, false
	}
	return o.AdditionalInputs, true
}

// HasAdditionalInputs returns a boolean if a field has been set.
func (o *PurchaseShipmentRequest) HasAdditionalInputs() bool {
	if o != nil && !IsNil(o.AdditionalInputs) {
		return true
	}

	return false
}

// SetAdditionalInputs gets a reference to the given map[string]map[string]interface{} and assigns it to the AdditionalInputs field.
func (o *PurchaseShipmentRequest) SetAdditionalInputs(v map[string]map[string]interface{}) {
	o.AdditionalInputs = v
}

func (o PurchaseShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestToken"] = o.RequestToken
	toSerialize["rateId"] = o.RateId
	toSerialize["requestedDocumentSpecification"] = o.RequestedDocumentSpecification
	if !IsNil(o.RequestedValueAddedServices) {
		toSerialize["requestedValueAddedServices"] = o.RequestedValueAddedServices
	}
	if !IsNil(o.AdditionalInputs) {
		toSerialize["additionalInputs"] = o.AdditionalInputs
	}
	return toSerialize, nil
}

type NullablePurchaseShipmentRequest struct {
	value *PurchaseShipmentRequest
	isSet bool
}

func (v NullablePurchaseShipmentRequest) Get() *PurchaseShipmentRequest {
	return v.value
}

func (v *NullablePurchaseShipmentRequest) Set(val *PurchaseShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseShipmentRequest(val *PurchaseShipmentRequest) *NullablePurchaseShipmentRequest {
	return &NullablePurchaseShipmentRequest{value: val, isSet: true}
}

func (v NullablePurchaseShipmentRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
