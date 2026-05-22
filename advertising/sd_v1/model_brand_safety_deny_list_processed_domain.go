package sd_v1

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyDenyListProcessedDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyDenyListProcessedDomain{}

// BrandSafetyDenyListProcessedDomain struct for BrandSafetyDenyListProcessedDomain
type BrandSafetyDenyListProcessedDomain struct {
	// The identifier of the Brand Safety List domain.
	DomainId *int64 `json:"domainId,omitempty"`
	// The website or app identifier. This can be in the form of full domain (eg. 'example.com' or 'example.net'), or mobile app identifier (eg. 'com.example.app' for Android apps or '1234567890' for iOS apps)
	Name  *string                         `json:"name,omitempty"`
	Type  *BrandSafetyDenyListDomainType  `json:"type,omitempty"`
	State *BrandSafetyDenyListDomainState `json:"state,omitempty"`
	// The date time the domain was created at. Format YYYY-MM-ddT:HH:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The date time the domain was last modified. Format YYYY-MM-ddT:HH:mm:ssZ
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewBrandSafetyDenyListProcessedDomain instantiates a new BrandSafetyDenyListProcessedDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyDenyListProcessedDomain() *BrandSafetyDenyListProcessedDomain {
	this := BrandSafetyDenyListProcessedDomain{}
	var state BrandSafetyDenyListDomainState = BRANDSAFETYDENYLISTDOMAINSTATE_ENABLED
	this.State = &state
	return &this
}

// NewBrandSafetyDenyListProcessedDomainWithDefaults instantiates a new BrandSafetyDenyListProcessedDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyDenyListProcessedDomainWithDefaults() *BrandSafetyDenyListProcessedDomain {
	this := BrandSafetyDenyListProcessedDomain{}
	var state BrandSafetyDenyListDomainState = BRANDSAFETYDENYLISTDOMAINSTATE_ENABLED
	this.State = &state
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *BrandSafetyDenyListProcessedDomain) GetDomainId() int64 {
	if o == nil || IsNil(o.DomainId) {
		var ret int64
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListProcessedDomain) GetDomainIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *BrandSafetyDenyListProcessedDomain) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given int64 and assigns it to the DomainId field.
func (o *BrandSafetyDenyListProcessedDomain) SetDomainId(v int64) {
	o.DomainId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrandSafetyDenyListProcessedDomain) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListProcessedDomain) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrandSafetyDenyListProcessedDomain) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrandSafetyDenyListProcessedDomain) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BrandSafetyDenyListProcessedDomain) GetType() BrandSafetyDenyListDomainType {
	if o == nil || IsNil(o.Type) {
		var ret BrandSafetyDenyListDomainType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListProcessedDomain) GetTypeOk() (*BrandSafetyDenyListDomainType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BrandSafetyDenyListProcessedDomain) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BrandSafetyDenyListDomainType and assigns it to the Type field.
func (o *BrandSafetyDenyListProcessedDomain) SetType(v BrandSafetyDenyListDomainType) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BrandSafetyDenyListProcessedDomain) GetState() BrandSafetyDenyListDomainState {
	if o == nil || IsNil(o.State) {
		var ret BrandSafetyDenyListDomainState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListProcessedDomain) GetStateOk() (*BrandSafetyDenyListDomainState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BrandSafetyDenyListProcessedDomain) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given BrandSafetyDenyListDomainState and assigns it to the State field.
func (o *BrandSafetyDenyListProcessedDomain) SetState(v BrandSafetyDenyListDomainState) {
	o.State = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BrandSafetyDenyListProcessedDomain) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListProcessedDomain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BrandSafetyDenyListProcessedDomain) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BrandSafetyDenyListProcessedDomain) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *BrandSafetyDenyListProcessedDomain) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListProcessedDomain) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *BrandSafetyDenyListProcessedDomain) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *BrandSafetyDenyListProcessedDomain) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o BrandSafetyDenyListProcessedDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DomainId) {
		toSerialize["domainId"] = o.DomainId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableBrandSafetyDenyListProcessedDomain struct {
	value *BrandSafetyDenyListProcessedDomain
	isSet bool
}

func (v NullableBrandSafetyDenyListProcessedDomain) Get() *BrandSafetyDenyListProcessedDomain {
	return v.value
}

func (v *NullableBrandSafetyDenyListProcessedDomain) Set(val *BrandSafetyDenyListProcessedDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyDenyListProcessedDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyDenyListProcessedDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyDenyListProcessedDomain(val *BrandSafetyDenyListProcessedDomain) *NullableBrandSafetyDenyListProcessedDomain {
	return &NullableBrandSafetyDenyListProcessedDomain{value: val, isSet: true}
}

func (v NullableBrandSafetyDenyListProcessedDomain) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyDenyListProcessedDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
