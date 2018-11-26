import re

length = 256
output = []
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

for l in range (0,128):
    inputVal = "oundnydw-" + str(l)
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
    output.append(string)

oneCount = 0
grid = [[0 for x in range(128)] for y in range(128)]
k = 0
for i in output:
    binString = ''
    for j in range(0, len(i)):
        binary = bin(int(i[j], 16))[2:].zfill(4)
        binString += str(binary)
    oneCount += int(binString.count('1'))
    for j in range(0, len(binString)):
        grid[k][j] = int(binString[j])
    k += 1
print ("Sol 1: " + str(oneCount))

def addGroup(grid, x, y, group):
    if grid[x][y] != 1:
        return (grid, 0)
    
    grid[x][y] = group
    if x != 0 and grid[x-1][y] != 0:
        (grid, _) = addGroup(grid, x-1, y, group)
    if x != 127 and grid[x+1][y] != 0:
        (grid, _) = addGroup(grid, x+1, y, group)
    if y != 0 and grid[x][y-1] != 0:
        (grid, _) = addGroup(grid, x, y-1, group)
    if y != 127 and grid[x][y+1] != 0:
        (grid, _) = addGroup(grid, x, y+1, group)
    return (grid, 1)

group = 2
for i in range(0,128):
    for j in range(0,128):
        (grid, groupAdd) = addGroup(grid, i, j, group)
        group += groupAdd

print ("sol 2: " + str(group - 2))
