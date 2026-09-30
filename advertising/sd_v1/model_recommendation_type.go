package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// RecommendationType Signifies a type of recommendation
type RecommendationType string

// List of RecommendationType
const (
	RECOMMENDATIONTYPE_PRODUCT RecommendationType = "PRODUCT"
)

// All allowed values of RecommendationType enum
var AllowedRecommendationTypeEnumValues = []RecommendationType{
	"PRODUCT",
}

func (v *RecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = RecommendationType(value)
	return nil
}

// NewRecommendationTypeFromValue returns a pointer to a valid RecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecommendationTypeFromValue(v string) (*RecommendationType, error) {
	ev := RecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecommendationType: valid values are %v", v, AllowedRecommendationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecommendationType) IsValid() bool {
	for _, existing := range AllowedRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationType value
func (v RecommendationType) Ptr() *RecommendationType {
	return &v
}

type NullableRecommendationType struct {
	value *RecommendationType
	isSet bool
}

func (v NullableRecommendationType) Get() *RecommendationType {
	return v.value
}

func (v *NullableRecommendationType) Set(val *RecommendationType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationType(val *RecommendationType) *NullableRecommendationType {
	return &NullableRecommendationType{value: val, isSet: true}
}

func (v NullableRecommendationType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecommendationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
