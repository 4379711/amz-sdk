package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyRequestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyRequestResult{}

// BrandSafetyRequestResult struct for BrandSafetyRequestResult
type BrandSafetyRequestResult struct {
	Status *BrandSafetyDenyListDomainUpdateResultStatus `json:"status,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
	// The identifier of the Brand Safety Deny List Domain.
	DomainId *int64 `json:"domainId,omitempty"`
	// The website or app identifier.
	Name *string `json:"name,omitempty"`
}

// NewBrandSafetyRequestResult instantiates a new BrandSafetyRequestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyRequestResult() *BrandSafetyRequestResult {
	this := BrandSafetyRequestResult{}
	var status BrandSafetyDenyListDomainUpdateResultStatus = BRANDSAFETYDENYLISTDOMAINUPDATERESULTSTATUS_SUCCESS
	this.Status = &status
	return &this
}

// NewBrandSafetyRequestResultWithDefaults instantiates a new BrandSafetyRequestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyRequestResultWithDefaults() *BrandSafetyRequestResult {
	this := BrandSafetyRequestResult{}
	var status BrandSafetyDenyListDomainUpdateResultStatus = BRANDSAFETYDENYLISTDOMAINUPDATERESULTSTATUS_SUCCESS
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BrandSafetyRequestResult) GetStatus() BrandSafetyDenyListDomainUpdateResultStatus {
	if o == nil || IsNil(o.Status) {
		var ret BrandSafetyDenyListDomainUpdateResultStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestResult) GetStatusOk() (*BrandSafetyDenyListDomainUpdateResultStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BrandSafetyRequestResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BrandSafetyDenyListDomainUpdateResultStatus and assigns it to the Status field.
func (o *BrandSafetyRequestResult) SetStatus(v BrandSafetyDenyListDomainUpdateResultStatus) {
	o.Status = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BrandSafetyRequestResult) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestResult) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BrandSafetyRequestResult) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BrandSafetyRequestResult) SetDetails(v string) {
	o.Details = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *BrandSafetyRequestResult) GetDomainId() int64 {
	if o == nil || IsNil(o.DomainId) {
		var ret int64
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestResult) GetDomainIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *BrandSafetyRequestResult) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given int64 and assigns it to the DomainId field.
func (o *BrandSafetyRequestResult) SetDomainId(v int64) {
	o.DomainId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrandSafetyRequestResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrandSafetyRequestResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrandSafetyRequestResult) SetName(v string) {
	o.Name = &v
}

func (o BrandSafetyRequestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.DomainId) {
		toSerialize["domainId"] = o.DomainId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableBrandSafetyRequestResult struct {
	value *BrandSafetyRequestResult
	isSet bool
}

func (v NullableBrandSafetyRequestResult) Get() *BrandSafetyRequestResult {
	return v.value
}

func (v *NullableBrandSafetyRequestResult) Set(val *BrandSafetyRequestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyRequestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyRequestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyRequestResult(val *BrandSafetyRequestResult) *NullableBrandSafetyRequestResult {
	return &NullableBrandSafetyRequestResult{value: val, isSet: true}
}

func (v NullableBrandSafetyRequestResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyRequestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
