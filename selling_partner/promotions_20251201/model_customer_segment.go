package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomerSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSegment{}

// CustomerSegment The customer segment. When `segmentType` is `BRAND`, `segmentId` contains the segment / audience ID and `segmentDetails.brandSegmentDetails` must be provided. When `segmentType` is `PROGRAM`, `segmentId` contains the program name (for example, `PRIME_EXCLUSIVE`, `MOM`, `STUDENT`).
type CustomerSegment struct {
	SegmentType CustomerSegmentType `json:"segmentType"`
	// The segment identifier. For `BRAND`: the segment / audience ID. For `PROGRAM`: the program name (for example, `PRIME_EXCLUSIVE`, `MOM`, `STUDENT` etc).
	SegmentId      string                  `json:"segmentId"`
	SegmentDetails *CustomerSegmentDetails `json:"segmentDetails,omitempty"`
}

// NewCustomerSegment instantiates a new CustomerSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSegment(segmentType CustomerSegmentType, segmentId string) *CustomerSegment {
	this := CustomerSegment{}
	this.SegmentType = segmentType
	this.SegmentId = segmentId
	return &this
}

// NewCustomerSegmentWithDefaults instantiates a new CustomerSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSegmentWithDefaults() *CustomerSegment {
	this := CustomerSegment{}
	return &this
}

// GetSegmentType returns the SegmentType field value
func (o *CustomerSegment) GetSegmentType() CustomerSegmentType {
	if o == nil {
		var ret CustomerSegmentType
		return ret
	}

	return o.SegmentType
}

// GetSegmentTypeOk returns a tuple with the SegmentType field value
// and a boolean to check if the value has been set.
func (o *CustomerSegment) GetSegmentTypeOk() (*CustomerSegmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentType, true
}

// SetSegmentType sets field value
func (o *CustomerSegment) SetSegmentType(v CustomerSegmentType) {
	o.SegmentType = v
}

// GetSegmentId returns the SegmentId field value
func (o *CustomerSegment) GetSegmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value
// and a boolean to check if the value has been set.
func (o *CustomerSegment) GetSegmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentId, true
}

// SetSegmentId sets field value
func (o *CustomerSegment) SetSegmentId(v string) {
	o.SegmentId = v
}

// GetSegmentDetails returns the SegmentDetails field value if set, zero value otherwise.
func (o *CustomerSegment) GetSegmentDetails() CustomerSegmentDetails {
	if o == nil || IsNil(o.SegmentDetails) {
		var ret CustomerSegmentDetails
		return ret
	}
	return *o.SegmentDetails
}

// GetSegmentDetailsOk returns a tuple with the SegmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSegment) GetSegmentDetailsOk() (*CustomerSegmentDetails, bool) {
	if o == nil || IsNil(o.SegmentDetails) {
		return nil, false
	}
	return o.SegmentDetails, true
}

// HasSegmentDetails returns a boolean if a field has been set.
func (o *CustomerSegment) HasSegmentDetails() bool {
	if o != nil && !IsNil(o.SegmentDetails) {
		return true
	}

	return false
}

// SetSegmentDetails gets a reference to the given CustomerSegmentDetails and assigns it to the SegmentDetails field.
func (o *CustomerSegment) SetSegmentDetails(v CustomerSegmentDetails) {
	o.SegmentDetails = &v
}

func (o CustomerSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["segmentType"] = o.SegmentType
	toSerialize["segmentId"] = o.SegmentId
	if !IsNil(o.SegmentDetails) {
		toSerialize["segmentDetails"] = o.SegmentDetails
	}
	return toSerialize, nil
}

type NullableCustomerSegment struct {
	value *CustomerSegment
	isSet bool
}

func (v NullableCustomerSegment) Get() *CustomerSegment {
	return v.value
}

func (v *NullableCustomerSegment) Set(val *CustomerSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSegment(val *CustomerSegment) *NullableCustomerSegment {
	return &NullableCustomerSegment{value: val, isSet: true}
}

func (v NullableCustomerSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomerSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
