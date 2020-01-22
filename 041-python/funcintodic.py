def a(x):
    return x

def b(x):
    return x*2

c = {"a":a, "b":b}

f = input("func use:               ")
v = int(input("input value:            "))

print(c[f](v))