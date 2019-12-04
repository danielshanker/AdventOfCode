import re
import math

def compute(codeList):
    curPos = 0
    while (curPos <= len(codeList)):
        value = 0
        nextPos = curPos
        a = codeList[curPos + 1]
        b = codeList[curPos + 2]
        a = codeList[a]
        b = codeList[b]
        if codeList[curPos] == 1:
            value = a + b
            nextPos = curPos + 4
        elif codeList[curPos] == 2:
            value = a * b
            nextPos = curPos + 4
        elif codeList[curPos] == 99:
            nextPos = curPos + 1
            break
        else:
            print "An error has occurred"
        codeList[codeList[curPos+3]] = value
        curPos = nextPos
    return codeList[0]




codes = []
file = open("input2.txt", "r")
for line in file:
        line = line.rstrip()
        codes = line.split(",")
codes = [ int(x) for x in codes ]
originalCode = [ int(x) for x in codes ]

curPos = 0
codes[1] = 12
codes[2] = 2

print "s1: " + str(compute(codes))

target = 19690720
x = 0
y = 0

for i in range(100):
    for j in range(100):
        curCode = [ int(x) for x in originalCode ]
        curCode[1] = i
        curCode[2] = j
        curVal = compute(curCode)
        if curVal == target:
            x = i
            y = j
            print "s2: " + str((100 * x) + y)
            exit()

