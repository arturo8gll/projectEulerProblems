verificar=[3,5]
count=0
for i in range(1,1000):
    for j in verificar:
        if (i%j)==0:
            print(f"{j}:{i}")
            count=count+i
            break
print(count)
