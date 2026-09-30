package sd_v1

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

type targetingJSONKind uint8

const (
	targetingJSONNone targetingJSONKind = iota
	targetingJSONContent
	targetingJSONScalar
	targetingJSONLegacy
	targetingJSONNested
)

// Each nested predicate carries its own immutable snapshot so that unknown
// fields stay with the predicate when callers move or copy array elements.
type targetingJSONState struct {
	raw      []byte
	fields   map[string]json.RawMessage
	baseline map[string]json.RawMessage
	kind     targetingJSONKind
}

func targetingJSONObject(data []byte) (map[string]json.RawMessage, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

func targetingKindFromJSON(data []byte, allowLegacy bool) (targetingJSONKind, error) {
	fields, err := targetingJSONObject(data)
	if err != nil {
		return targetingJSONNone, err
	}
	if fields == nil {
		return targetingJSONNone, nil
	}
	value := bytes.TrimSpace(fields["value"])
	if len(value) > 0 && value[0] == '[' {
		return targetingJSONNested, nil
	}
	if _, exists := fields["eventType"]; allowLegacy && exists {
		return targetingJSONLegacy, nil
	}
	var predicateType string
	if value, exists := fields["type"]; exists {
		if err := sonic.Unmarshal(value, &predicateType); err != nil {
			return targetingJSONNone, err
		}
	}
	if predicateType == "contentCategorySameAs" {
		return targetingJSONContent, nil
	}
	return targetingJSONScalar, nil
}

func encodeTargetingJSON(typed any) ([]byte, error) {
	// ToMap keeps a non-nil empty optional slice distinguishable from nil.
	// That distinction lets callers clear an originally empty nested value.
	if mapped, ok := typed.(MappedNullable); ok {
		fields, err := mapped.ToMap()
		if err != nil {
			return nil, err
		}
		return sonic.Marshal(fields)
	}
	return sonic.Marshal(typed)
}

func newTargetingJSONState(data []byte, typed any, kind targetingJSONKind) (*targetingJSONState, error) {
	fields, err := targetingJSONObject(data)
	if err != nil {
		return nil, err
	}
	encoded, err := encodeTargetingJSON(typed)
	if err != nil {
		return nil, err
	}
	baseline, err := targetingJSONObject(encoded)
	if err != nil {
		return nil, err
	}
	return &targetingJSONState{
		raw:      bytes.Clone(data),
		fields:   fields,
		baseline: baseline,
		kind:     kind,
	}, nil
}

func sameTargetingJSON(a, b json.RawMessage) bool {
	var compactA, compactB bytes.Buffer
	if err := json.Compact(&compactA, a); err != nil {
		return len(a) == 0 && len(b) == 0
	}
	if err := json.Compact(&compactB, b); err != nil {
		return false
	}
	return bytes.Equal(compactA.Bytes(), compactB.Bytes())
}

func (state *targetingJSONState) marshal(typed any, kind targetingJSONKind) ([]byte, error) {
	encoded, err := encodeTargetingJSON(typed)
	if err != nil || state == nil {
		return encoded, err
	}
	current, err := targetingJSONObject(encoded)
	if err != nil {
		return nil, err
	}
	merged := make(map[string]json.RawMessage, len(state.fields)+len(current))
	for key, value := range state.fields {
		merged[key] = value
	}
	changed := state.kind != kind
	if changed {
		// A branch replacement owns all schema fields, including fields omitted
		// by the new branch; unrelated extension fields remain intact.
		delete(merged, "type")
		delete(merged, "value")
		delete(merged, "eventType")
		for key, value := range current {
			merged[key] = value
		}
	} else {
		for key, old := range state.baseline {
			value, exists := current[key]
			if !exists {
				delete(merged, key)
				changed = true
			} else if !sameTargetingJSON(old, value) {
				merged[key] = value
				changed = true
			}
		}
		for key, value := range current {
			if _, exists := state.baseline[key]; !exists {
				merged[key] = value
				changed = true
			}
		}
	}
	if !changed {
		return bytes.Clone(state.raw), nil
	}
	return json.Marshal(merged)
}

type targetingJSONVariant struct {
	kind  targetingJSONKind
	value any
}

func marshalTargetingUnion(state *targetingJSONState, variants ...targetingJSONVariant) ([]byte, error) {
	var selected *targetingJSONVariant
	for i := range variants {
		if IsNil(variants[i].value) {
			continue
		}
		if selected != nil {
			return nil, fmt.Errorf("targeting expression must contain only one variant")
		}
		selected = &variants[i]
	}
	if selected == nil {
		return []byte("null"), nil
	}
	return state.marshal(selected.value, selected.kind)
}

type targetingPredicateBaseJSON TargetingPredicateBase

func (dst *TargetingPredicateBase) UnmarshalJSON(data []byte) error {
	var decoded targetingPredicateBaseJSON
	if err := sonic.Unmarshal(data, &decoded); err != nil {
		return err
	}
	state, err := newTargetingJSONState(data, decoded, targetingJSONScalar)
	if err != nil {
		return err
	}
	decoded.jsonState = state
	*dst = TargetingPredicateBase(decoded)
	return nil
}

func (src TargetingPredicateBase) MarshalJSON() ([]byte, error) {
	return src.jsonState.marshal(targetingPredicateBaseJSON(src), targetingJSONScalar)
}

func (dst *SDTargetingPredicateBaseV31) UnmarshalJSON(data []byte) error {
	var decoded _SDTargetingPredicateBaseV31
	if err := sonic.Unmarshal(data, &decoded); err != nil {
		return err
	}
	state, err := newTargetingJSONState(data, decoded, targetingJSONScalar)
	if err != nil {
		return err
	}
	decoded.jsonState = state
	*dst = SDTargetingPredicateBaseV31(decoded)
	return nil
}

func (src SDTargetingPredicateBaseV31) MarshalJSON() ([]byte, error) {
	return src.jsonState.marshal(_SDTargetingPredicateBaseV31(src), targetingJSONScalar)
}

// UnmarshalJSON selects the structural variant before retaining extension fields.
func (dst *TargetingExpressionInner) UnmarshalJSON(data []byte) error {
	kind, err := targetingKindFromJSON(data, true)
	if err != nil {
		return err
	}
	var decoded TargetingExpressionInner
	switch kind {
	case targetingJSONContent:
		err = sonic.Unmarshal(data, &decoded.ContentTargetingPredicate)
	case targetingJSONScalar:
		err = sonic.Unmarshal(data, &decoded.TargetingPredicate)
	case targetingJSONLegacy:
		err = sonic.Unmarshal(data, &decoded.TargetingPredicateLegacy)
	case targetingJSONNested:
		err = sonic.Unmarshal(data, &decoded.TargetingPredicateNested)
	}
	if err != nil {
		return err
	}
	decoded.jsonState, err = newTargetingJSONState(data, decoded.GetActualInstance(), kind)
	if err != nil {
		return err
	}
	*dst = decoded
	return nil
}

// MarshalJSON preserves extension fields while reflecting edits to the selected variant.
func (src TargetingExpressionInner) MarshalJSON() ([]byte, error) {
	return marshalTargetingUnion(src.jsonState,
		targetingJSONVariant{targetingJSONContent, src.ContentTargetingPredicate},
		targetingJSONVariant{targetingJSONScalar, src.TargetingPredicate},
		targetingJSONVariant{targetingJSONLegacy, src.TargetingPredicateLegacy},
		targetingJSONVariant{targetingJSONNested, src.TargetingPredicateNested},
	)
}

// UnmarshalJSON selects the structural variant before retaining extension fields.
func (dst *SDTargetExpressionV32) UnmarshalJSON(data []byte) error {
	kind, err := targetingKindFromJSON(data, false)
	if err != nil {
		return err
	}
	var decoded SDTargetExpressionV32
	switch kind {
	case targetingJSONContent:
		err = sonic.Unmarshal(data, &decoded.SDContentTargetingPredicateV31)
	case targetingJSONNested:
		err = sonic.Unmarshal(data, &decoded.SDTargetingPredicateNestedV31)
	case targetingJSONScalar:
		err = sonic.Unmarshal(data, &decoded.SDTargetingPredicateV31)
	}
	if err != nil {
		return err
	}
	decoded.jsonState, err = newTargetingJSONState(data, decoded.GetActualInstance(), kind)
	if err != nil {
		return err
	}
	*dst = decoded
	return nil
}

// MarshalJSON preserves extension fields while reflecting edits to the selected variant.
func (src SDTargetExpressionV32) MarshalJSON() ([]byte, error) {
	return marshalTargetingUnion(src.jsonState,
		targetingJSONVariant{targetingJSONContent, src.SDContentTargetingPredicateV31},
		targetingJSONVariant{targetingJSONNested, src.SDTargetingPredicateNestedV31},
		targetingJSONVariant{targetingJSONScalar, src.SDTargetingPredicateV31},
	)
}
