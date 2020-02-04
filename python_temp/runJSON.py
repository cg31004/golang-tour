import json

with open("tmp.json", "r", encoding="utf-8") as f:
    jf = json.loads(f.read())

print(len(jf["items"]))
