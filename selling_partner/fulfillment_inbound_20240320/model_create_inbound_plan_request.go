package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateInboundPlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInboundPlanRequest{}

// CreateInboundPlanRequest The `createInboundPlan` request.
type CreateInboundPlanRequest struct {
	// Marketplaces where the items need to be shipped to. Currently only one marketplace can be selected in this request.
	DestinationMarketplaces []string `json:"destinationMarketplaces"`
	// Items included in this plan.
	Items []ItemInput `json:"items"`
	// Name for the Inbound Plan. If one isn't provided, a default name will be provided.
	Name          *string      `json:"name,omitempty"`
	SourceAddress AddressInput `json:"sourceAddress"`
}

type _CreateInboundPlanRequest CreateInboundPlanRequest

// NewCreateInboundPlanRequest instantiates a new CreateInboundPlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInboundPlanRequest(destinationMarketplaces []string, items []ItemInput, sourceAddress AddressInput) *CreateInboundPlanRequest {
	this := CreateInboundPlanRequest{}
	this.DestinationMarketplaces = destinationMarketplaces
	this.Items = items
	this.SourceAddress = sourceAddress
	return &this
}

// NewCreateInboundPlanRequestWithDefaults instantiates a new CreateInboundPlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInboundPlanRequestWithDefaults() *CreateInboundPlanRequest {
	this := CreateInboundPlanRequest{}
	return &this
}

// GetDestinationMarketplaces returns the DestinationMarketplaces field value
func (o *CreateInboundPlanRequest) GetDestinationMarketplaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DestinationMarketplaces
}

// GetDestinationMarketplacesOk returns a tuple with the DestinationMarketplaces field value
// and a boolean to check if the value has been set.
func (o *CreateInboundPlanRequest) GetDestinationMarketplacesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationMarketplaces, true
}

// SetDestinationMarketplaces sets field value
func (o *CreateInboundPlanRequest) SetDestinationMarketplaces(v []string) {
	o.DestinationMarketplaces = v
}

// GetItems returns the Items field value
func (o *CreateInboundPlanRequest) GetItems() []ItemInput {
	if o == nil {
		var ret []ItemInput
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreateInboundPlanRequest) GetItemsOk() ([]ItemInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CreateInboundPlanRequest) SetItems(v []ItemInput) {
	o.Items = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateInboundPlanRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundPlanRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateInboundPlanRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateInboundPlanRequest) SetName(v string) {
	o.Name = &v
}

// GetSourceAddress returns the SourceAddress field value
func (o *CreateInboundPlanRequest) GetSourceAddress() AddressInput {
	if o == nil {
		var ret AddressInput
		return ret
	}

	return o.SourceAddress
}

// GetSourceAddressOk returns a tuple with the SourceAddress field value
// and a boolean to check if the value has been set.
func (o *CreateInboundPlanRequest) GetSourceAddressOk() (*AddressInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAddress, true
}

// SetSourceAddress sets field value
func (o *CreateInboundPlanRequest) SetSourceAddress(v AddressInput) {
	o.SourceAddress = v
}

func (o CreateInboundPlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinationMarketplaces"] = o.DestinationMarketplaces
	toSerialize["items"] = o.Items
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["sourceAddress"] = o.SourceAddress
	return toSerialize, nil
}

type NullableCreateInboundPlanRequest struct {
	value *CreateInboundPlanRequest
	isSet bool
}

func (v NullableCreateInboundPlanRequest) Get() *CreateInboundPlanRequest {
	return v.value
}

func (v *NullableCreateInboundPlanRequest) Set(val *CreateInboundPlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInboundPlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInboundPlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInboundPlanRequest(val *CreateInboundPlanRequest) *NullableCreateInboundPlanRequest {
	return &NullableCreateInboundPlanRequest{value: val, isSet: true}
}

func (v NullableCreateInboundPlanRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateInboundPlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
