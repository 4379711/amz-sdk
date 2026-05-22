package sellers

import (
	"github.com/bytedance/sonic"
)

// checks if the Business type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Business{}

// Business Information about the seller's business. Certain fields may be omitted depending on the seller's `businessType`.
type Business struct {
	// The registered business name.
	Name                      string  `json:"name"`
	RegisteredBusinessAddress Address `json:"registeredBusinessAddress"`
	// The seller's company registration number, if applicable. This field will be absent for individual sellers and sole proprietorships.
	CompanyRegistrationNumber *string `json:"companyRegistrationNumber,omitempty"`
	// The seller's company tax identification number, if applicable. This field will be present for certain business types only, such as sole proprietorships.
	CompanyTaxIdentificationNumber *string `json:"companyTaxIdentificationNumber,omitempty"`
	// The non-Latin script version of the registered business name, if applicable.
	NonLatinName *string `json:"nonLatinName,omitempty"`
}

type _Business Business

// NewBusiness instantiates a new Business object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusiness(name string, registeredBusinessAddress Address) *Business {
	this := Business{}
	this.Name = name
	this.RegisteredBusinessAddress = registeredBusinessAddress
	return &this
}

// NewBusinessWithDefaults instantiates a new Business object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessWithDefaults() *Business {
	this := Business{}
	return &this
}

// GetName returns the Name field value
func (o *Business) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Business) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Business) SetName(v string) {
	o.Name = v
}

// GetRegisteredBusinessAddress returns the RegisteredBusinessAddress field value
func (o *Business) GetRegisteredBusinessAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.RegisteredBusinessAddress
}

// GetRegisteredBusinessAddressOk returns a tuple with the RegisteredBusinessAddress field value
// and a boolean to check if the value has been set.
func (o *Business) GetRegisteredBusinessAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegisteredBusinessAddress, true
}

// SetRegisteredBusinessAddress sets field value
func (o *Business) SetRegisteredBusinessAddress(v Address) {
	o.RegisteredBusinessAddress = v
}

// GetCompanyRegistrationNumber returns the CompanyRegistrationNumber field value if set, zero value otherwise.
func (o *Business) GetCompanyRegistrationNumber() string {
	if o == nil || IsNil(o.CompanyRegistrationNumber) {
		var ret string
		return ret
	}
	return *o.CompanyRegistrationNumber
}

// GetCompanyRegistrationNumberOk returns a tuple with the CompanyRegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetCompanyRegistrationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyRegistrationNumber) {
		return nil, false
	}
	return o.CompanyRegistrationNumber, true
}

// HasCompanyRegistrationNumber returns a boolean if a field has been set.
func (o *Business) HasCompanyRegistrationNumber() bool {
	if o != nil && !IsNil(o.CompanyRegistrationNumber) {
		return true
	}

	return false
}

// SetCompanyRegistrationNumber gets a reference to the given string and assigns it to the CompanyRegistrationNumber field.
func (o *Business) SetCompanyRegistrationNumber(v string) {
	o.CompanyRegistrationNumber = &v
}

// GetCompanyTaxIdentificationNumber returns the CompanyTaxIdentificationNumber field value if set, zero value otherwise.
func (o *Business) GetCompanyTaxIdentificationNumber() string {
	if o == nil || IsNil(o.CompanyTaxIdentificationNumber) {
		var ret string
		return ret
	}
	return *o.CompanyTaxIdentificationNumber
}

// GetCompanyTaxIdentificationNumberOk returns a tuple with the CompanyTaxIdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetCompanyTaxIdentificationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyTaxIdentificationNumber) {
		return nil, false
	}
	return o.CompanyTaxIdentificationNumber, true
}

// HasCompanyTaxIdentificationNumber returns a boolean if a field has been set.
func (o *Business) HasCompanyTaxIdentificationNumber() bool {
	if o != nil && !IsNil(o.CompanyTaxIdentificationNumber) {
		return true
	}

	return false
}

// SetCompanyTaxIdentificationNumber gets a reference to the given string and assigns it to the CompanyTaxIdentificationNumber field.
func (o *Business) SetCompanyTaxIdentificationNumber(v string) {
	o.CompanyTaxIdentificationNumber = &v
}

// GetNonLatinName returns the NonLatinName field value if set, zero value otherwise.
func (o *Business) GetNonLatinName() string {
	if o == nil || IsNil(o.NonLatinName) {
		var ret string
		return ret
	}
	return *o.NonLatinName
}

// GetNonLatinNameOk returns a tuple with the NonLatinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetNonLatinNameOk() (*string, bool) {
	if o == nil || IsNil(o.NonLatinName) {
		return nil, false
	}
	return o.NonLatinName, true
}

// HasNonLatinName returns a boolean if a field has been set.
func (o *Business) HasNonLatinName() bool {
	if o != nil && !IsNil(o.NonLatinName) {
		return true
	}

	return false
}

// SetNonLatinName gets a reference to the given string and assigns it to the NonLatinName field.
func (o *Business) SetNonLatinName(v string) {
	o.NonLatinName = &v
}

func (o Business) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["registeredBusinessAddress"] = o.RegisteredBusinessAddress
	if !IsNil(o.CompanyRegistrationNumber) {
		toSerialize["companyRegistrationNumber"] = o.CompanyRegistrationNumber
	}
	if !IsNil(o.CompanyTaxIdentificationNumber) {
		toSerialize["companyTaxIdentificationNumber"] = o.CompanyTaxIdentificationNumber
	}
	if !IsNil(o.NonLatinName) {
		toSerialize["nonLatinName"] = o.NonLatinName
	}
	return toSerialize, nil
}

type NullableBusiness struct {
	value *Business
	isSet bool
}

func (v NullableBusiness) Get() *Business {
	return v.value
}

func (v *NullableBusiness) Set(val *Business) {
	v.value = val
	v.isSet = true
}

func (v NullableBusiness) IsSet() bool {
	return v.isSet
}

func (v *NullableBusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusiness(val *Business) *NullableBusiness {
	return &NullableBusiness{value: val, isSet: true}
}

func (v NullableBusiness) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
