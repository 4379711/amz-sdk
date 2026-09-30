#!/usr/bin/env python3
"""Connect regenerated models to the SDK's maintained JSON codecs."""

from pathlib import Path
import re
import subprocess


ROOT = Path(__file__).resolve().parents[1]
MODELS = (
    ("sp_v3", "model_optimization_rules_api_swagger_rule_criteria.go",
     "OptimizationRulesAPIRuleCriteria", ("UnmarshalJSON",), False),
    ("sd_v1", "model_targeting_expression_inner.go",
     "TargetingExpressionInner", ("UnmarshalJSON", "MarshalJSON"), True),
    ("sd_v1", "model_sd_target_expression_v32.go",
     "SDTargetExpressionV32", ("UnmarshalJSON", "MarshalJSON"), True),
    ("sd_v1", "model_targeting_predicate_base.go",
     "TargetingPredicateBase", ("UnmarshalJSON", "MarshalJSON"), True),
    ("sd_v1", "model_sd_targeting_predicate_base_v31.go",
     "SDTargetingPredicateBaseV31", ("UnmarshalJSON", "MarshalJSON"), True),
)


def update_model(source, model, methods, needs_state):
    pattern = (
        r"^(?://[^\n]*\n)*func \(\w+ \*?" + re.escape(model)
        + r"\) (?:" + "|".join(methods) + r")\([^\n]*\)[^\n]*\{\n.*?^\}\n"
    )
    source = re.sub(pattern, "", source, flags=re.M | re.S)
    if needs_state:
        pattern = r"(type " + re.escape(model) + r" struct \{\n)(.*?)(^\})"
        match = re.search(pattern, source, re.M | re.S)
        if match is None:
            raise ValueError(f"Missing generated struct {model}")
        fields = match.group(2)
        # The maintained codec owns the raw bytes and their typed snapshot.
        fields = re.sub(r"(?:\t//[^\n]*\n)*\traw\s+\[\]byte[^\n]*\n", "", fields)
        if not re.search(r"\bjsonState\s+\*targetingJSONState\b", fields):
            fields += "\tjsonState *targetingJSONState\n"
        source = source[:match.start(2)] + fields + source[match.end(2):]
    for package, path in (("fmt", "fmt"), ("bytes", "bytes"),
                          ("json", "encoding/json"), ("validator", "gopkg.in/validator.v2")):
        without_import = re.sub(
            r'^\s*"' + re.escape(path) + r'"\s*\n', "", source, flags=re.M
        )
        if not re.search(r"\b" + package + r"\.", without_import):
            source = without_import
    return source


def main():
    for codec in ("sp_v3/rule_criteria_json.go", "sd_v1/targeting_expression_json.go"):
        if not (ROOT / "advertising" / codec).is_file():
            raise FileNotFoundError(f"Keep the maintained codec advertising/{codec}")
    changed = []
    for package, name, model, methods, needs_state in MODELS:
        path = ROOT / "advertising" / package / name
        original = path.read_text()
        source = update_model(original, model, methods, needs_state)
        formatted = subprocess.run(
            ["gofmt"], input=source, text=True, capture_output=True, check=True
        ).stdout
        if formatted != original:
            path.write_text(formatted)
            changed.append(path)
    print(f"Updated {len(changed)} JSON model files")


if __name__ == "__main__":
    main()
