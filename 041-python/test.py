import sys
import json

reload(sys)
sys.setdefaultencoding('uft-8')

def test(argv):
    print(argv)
    res = {
        "price" : 0.01,
        "name": "test"
    }
    final = json.dump(res)
    return final