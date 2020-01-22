from _switch import switch

v = "ten"
for case in switch(v):
    if case('one'):
        print(1)
        break
    if case('two'):
        print(2)
        break
    if case('three'):
        print(3)
        break
    if case('eleven'):
        print(11)
        break
    if case('ten'):
        print(10)
        break
    if case():
        print("nothing here")
        