import re

inputVal = "31,2,85,1,80,109,35,63,98,255,0,13,105,254,128,33"
length = 256

iList = []
inputs = []
for i in range (0, length):
    iList.append(i)
for i in range (0, len(inputVal)):
    a = ord(inputVal[i])
    inputs.append(a)
inputs.append(17)
inputs.append(31)
inputs.append(73)
inputs.append(47)
inputs.append(23)
print (inputs)

def reverseList(a, start, end):
    b = []
    j = 1
    for i in range (start, end):
        w = end - j
        w = w % length
        b.append(a[w])
        j += 1
    j = 0
    for i in range (start, end):
        w = end - j
        w = w % length
        iw = i % length
        a[iw] = b[j]
        j += 1
    return a

skipSize = 0
pos = 0

for q in range (0, 64):
    for k in range (0, len(inputs)):
        val = int(inputs[k])
        iList = reverseList(iList, pos, pos+val)
        oPos = pos
        newPos = pos + val + skipSize
        pos = newPos % length
        skipSize += 1

string = ''
for i in range (0, 16):
    start = 16*i
    variable = iList[start] ^ iList[start + 1] ^ iList[start + 2] ^ iList[start + 3] ^ iList[start + 4] ^ iList[start + 5] ^ iList[start + 6] ^ iList[start + 7] ^ iList[start + 8] ^ iList[start + 9] ^ iList[start + 10] ^ iList[start + 11] ^ iList[start + 12] ^ iList[start + 13] ^ iList[start + 14] ^ iList[start + 15]
    xor = hex(variable)
    if len(xor) < 4:
        xor = str(0)+str(xor[2])
    else:
        xor = str(xor[2]) + str(xor[3])
    string += xor

print (string)

