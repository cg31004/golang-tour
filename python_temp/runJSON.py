import json

with open("tmp.json", "r", encoding="utf-8") as f:
    jf = json.loads(f.read())

for k,v in jf.items():
    print(k)
    if k != "scheme":
        print(v)
        print("============================================")