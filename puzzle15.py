
inputA = 679
inputB = 771
aList = []
bList = []


prevA = inputA;
prevB = inputB;
count = 0
for j in range(0,40000000):
    a = prevA * 16807
    b = prevB * 48271

    a = a % 2147483647
    b = b % 2147483647

    prevA = a
    prevB = b

    binaryA = str(bin(a)[2:].zfill(32))
    binaryB = str(bin(b)[2:].zfill(32))
    if a % 4 == 0:
        aList.append(a)
    if b % 8 == 0:
        bList.append(b)
    noMatch = 0
    for i in range(16,32):
        if binaryA[i] != binaryB[i]:
            noMatch = 1
            break
    if noMatch == 0:
        count +=1
print "sol 1: " + str(count)

count = 0
for i in range(0, 5000000):
    if i >= len(bList):
        break
    binaryA = str(bin(aList[i])[2:].zfill(32))
    binaryB = str(bin(bList[i])[2:].zfill(32))
    noMatch = 0
    for j in range(16,32):
        if binaryA[j] != binaryB[j]:
            noMatch = 1
            break
    if noMatch == 0:
        count +=1
print "sol 2: " + str( count)
