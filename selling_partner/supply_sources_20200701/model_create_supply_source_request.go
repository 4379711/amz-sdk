package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSupplySourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSupplySourceRequest{}

// CreateSupplySourceRequest A request to create a supply source.
type CreateSupplySourceRequest struct {
	// The seller-provided unique supply source code.
	SupplySourceCode string `json:"supplySourceCode"`
	// The custom alias for this supply source
	Alias   string  `json:"alias"`
	Address Address `json:"address"`
}

type _CreateSupplySourceRequest CreateSupplySourceRequest

// NewCreateSupplySourceRequest instantiates a new CreateSupplySourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSupplySourceRequest(supplySourceCode string, alias string, address Address) *CreateSupplySourceRequest {
	this := CreateSupplySourceRequest{}
	this.SupplySourceCode = supplySourceCode
	this.Alias = alias
	this.Address = address
	return &this
}

// NewCreateSupplySourceRequestWithDefaults instantiates a new CreateSupplySourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSupplySourceRequestWithDefaults() *CreateSupplySourceRequest {
	this := CreateSupplySourceRequest{}
	return &this
}

// GetSupplySourceCode returns the SupplySourceCode field value
func (o *CreateSupplySourceRequest) GetSupplySourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplySourceCode
}

// GetSupplySourceCodeOk returns a tuple with the SupplySourceCode field value
// and a boolean to check if the value has been set.
func (o *CreateSupplySourceRequest) GetSupplySourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplySourceCode, true
}

// SetSupplySourceCode sets field value
func (o *CreateSupplySourceRequest) SetSupplySourceCode(v string) {
	o.SupplySourceCode = v
}

// GetAlias returns the Alias field value
func (o *CreateSupplySourceRequest) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *CreateSupplySourceRequest) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *CreateSupplySourceRequest) SetAlias(v string) {
	o.Alias = v
}

// GetAddress returns the Address field value
func (o *CreateSupplySourceRequest) GetAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateSupplySourceRequest) GetAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateSupplySourceRequest) SetAddress(v Address) {
	o.Address = v
}

func (o CreateSupplySourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supplySourceCode"] = o.SupplySourceCode
	toSerialize["alias"] = o.Alias
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableCreateSupplySourceRequest struct {
	value *CreateSupplySourceRequest
	isSet bool
}

func (v NullableCreateSupplySourceRequest) Get() *CreateSupplySourceRequest {
	return v.value
}

func (v *NullableCreateSupplySourceRequest) Set(val *CreateSupplySourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSupplySourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSupplySourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSupplySourceRequest(val *CreateSupplySourceRequest) *NullableCreateSupplySourceRequest {
	return &NullableCreateSupplySourceRequest{value: val, isSet: true}
}

func (v NullableCreateSupplySourceRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSupplySourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
