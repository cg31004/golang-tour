import json

with open("042-loadjson/gamelist.json", "r", encoding="utf-8") as f :
    jf = json.loads(f.read())
# for k in jf:
#     print(k)
data = jf['Items']
channelcode = list()

for d in data :
    channelcode.append(d["channelcode"])

print(channelcode)
scc = set(channelcode)
print(scc)