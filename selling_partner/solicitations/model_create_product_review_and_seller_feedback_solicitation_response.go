package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateProductReviewAndSellerFeedbackSolicitationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductReviewAndSellerFeedbackSolicitationResponse{}

// CreateProductReviewAndSellerFeedbackSolicitationResponse The response schema for the createProductReviewAndSellerFeedbackSolicitation operation.
type CreateProductReviewAndSellerFeedbackSolicitationResponse struct {
	// A list of error responses returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

// NewCreateProductReviewAndSellerFeedbackSolicitationResponse instantiates a new CreateProductReviewAndSellerFeedbackSolicitationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductReviewAndSellerFeedbackSolicitationResponse() *CreateProductReviewAndSellerFeedbackSolicitationResponse {
	this := CreateProductReviewAndSellerFeedbackSolicitationResponse{}
	return &this
}

// NewCreateProductReviewAndSellerFeedbackSolicitationResponseWithDefaults instantiates a new CreateProductReviewAndSellerFeedbackSolicitationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductReviewAndSellerFeedbackSolicitationResponseWithDefaults() *CreateProductReviewAndSellerFeedbackSolicitationResponse {
	this := CreateProductReviewAndSellerFeedbackSolicitationResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateProductReviewAndSellerFeedbackSolicitationResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductReviewAndSellerFeedbackSolicitationResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateProductReviewAndSellerFeedbackSolicitationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CreateProductReviewAndSellerFeedbackSolicitationResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o CreateProductReviewAndSellerFeedbackSolicitationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateProductReviewAndSellerFeedbackSolicitationResponse struct {
	value *CreateProductReviewAndSellerFeedbackSolicitationResponse
	isSet bool
}

func (v NullableCreateProductReviewAndSellerFeedbackSolicitationResponse) Get() *CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return v.value
}

func (v *NullableCreateProductReviewAndSellerFeedbackSolicitationResponse) Set(val *CreateProductReviewAndSellerFeedbackSolicitationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductReviewAndSellerFeedbackSolicitationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductReviewAndSellerFeedbackSolicitationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductReviewAndSellerFeedbackSolicitationResponse(val *CreateProductReviewAndSellerFeedbackSolicitationResponse) *NullableCreateProductReviewAndSellerFeedbackSolicitationResponse {
	return &NullableCreateProductReviewAndSellerFeedbackSolicitationResponse{value: val, isSet: true}
}

func (v NullableCreateProductReviewAndSellerFeedbackSolicitationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateProductReviewAndSellerFeedbackSolicitationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
