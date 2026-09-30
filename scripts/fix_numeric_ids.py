#!/usr/bin/env python3
"""Preserve numeric Amazon entity IDs as int64 after generating the clients."""

from pathlib import Path
import re
import subprocess


ROOT = Path(__file__).resolve().parents[1]
MODEL_FIELDS = {
    "model_campaign_response_ex.go": ("CampaignId",),
    "model_ad_group_response_ex.go": ("AdGroupId", "CampaignId"),
    "model_product_ad_response_ex.go": ("AdId", "AdGroupId", "CampaignId"),
    "model_product_ad_response.go": ("AdId",),
    "model_targeting_clause_ex.go": ("TargetId", "AdGroupId", "CampaignId"),
    "model_negative_targeting_clause_ex.go": ("TargetId", "AdGroupId"),
    "model_creative.go": ("CreativeId",),
    "model_creative_update.go": ("CreativeId",),
    "model_creative_response.go": ("CreativeId",),
    "model_create_creative.go": ("AdGroupId",),
    "model_creative_moderation.go": ("CreativeId",),
}


def update_model(source, fields):
    for field in fields:
        source = re.sub(
            rf"(\b{field}\s+\*?)float32(?=\s+`json:)", r"\1int64", source
        )
        argument = field[0].lower() + field[1:]
        source = re.sub(rf"\b{argument} float32\b", f"{argument} int64", source)
        source = re.sub(
            rf"func \(o \*\w+\) (?:Get{field}(?:Ok)?|Set{field})\([^\n]*\)[^\n]*\{{.*?\n\}}",
            lambda match: match.group().replace("float32", "int64"),
            source,
            flags=re.S,
        )
        source = re.sub(
            rf"(// Set{field} gets a reference to the given )float32",
            r"\1int64",
            source,
        )
    return source


def main():
    updates = {}
    for package in ("sp_v3", "sb_v4"):
        path = ROOT / "advertising" / package / "api_budget_rules.go"
        updates[path] = re.sub(
            r"(\bcampaignId\s+)float32\b", r"\1int64", path.read_text()
        )
    for name, fields in MODEL_FIELDS.items():
        path = ROOT / "advertising/sd_v1" / name
        updates[path] = update_model(path.read_text(), fields)

    changed = []
    for path, source in updates.items():
        if source != path.read_text():
            path.write_text(source)
            changed.append(str(path))
    if changed:
        subprocess.run(["gofmt", "-w", *changed], check=True)
    print(f"Updated {len(changed)} numeric ID files")


if __name__ == "__main__":
    main()
