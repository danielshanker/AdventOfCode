import re
import math

def compute(codeList, inputVal):
    curPos = 0
    output = 0
    while (curPos <= len(codeList)):
        value = 0
        nextPos = curPos
        opCode = str(codeList[curPos])
        regC = opCode[-5:-4]
        regB = opCode[-4:-3]
        regA = opCode[-3:-2]
        inst = int(opCode[-2:])

        if inst == 99:
            break

        if inst == 2 or inst == 1 or inst == 7 or inst == 8:
            if regA == "1":
                a = codeList[curPos + 1]
            else:
                a = codeList[curPos + 1]
                a = codeList[a]
            if regB == "1":
                b = codeList[curPos + 2]
            else:
                b = codeList[curPos + 2]
                b = codeList[b]
            if regC == "1":
                c = codeList[curPos + 3]
            else:
                c = codeList[curPos + 3]
        elif inst == 3 or inst == 4:
            if regA == "1":
                a = codeList[curPos + 1]
            else:
                a = codeList[curPos + 1]
                a = codeList[a]
        elif inst == 5 or inst == 6:
            if regA == "1":
                a = codeList[curPos + 1]
            else:
                a = codeList[curPos + 1]
                a = codeList[a]
            if regB == "1":
                b = codeList[curPos + 2]
            else:
                b = codeList[curPos + 2]
                b = codeList[b]
        else:
            print "error inst - " + str(inst)
            exit()
            

        if inst == 01:
            value = a + b
            codeList[c] = value
            nextPos = curPos + 4
        elif inst == 02:
            value = a * b
            codeList[c] = value
            nextPos = curPos + 4
        elif inst == 03:
            codeList[codeList[curPos + 1]] = inputVal
            nextPos = curPos + 2
        elif inst == 04:
            if output != 0:
                print "output was not 0"
                exit()
            if regA == "1":
                output = codeList[curPos + 1]
            else:
                output = codeList[codeList[curPos + 1]]
            nextPos = curPos + 2
        elif inst == 05:
            if a != 0:
                nextPos = b
            else:
                nextPos = curPos + 3
        elif inst == 06:
            if a == 0:
                nextPos = b
            else:
                nextPos = curPos + 3
        elif inst == 07:
            if a < b:
                codeList[c] = 1
            else:
                codeList[c] = 0
            nextPos = curPos + 4
        elif inst == 8:
            if a == b:
                codeList[c] = 1
            else:
                codeList[c] = 0
            nextPos = curPos + 4
        elif inst == 99:
            break
        else:
            print codeList[curPos:curPos+4]
            print "An error has occurred curPos=" + str(curPos) + " instruction = " + str(inst)
            print codeList
            exit()
        curPos = nextPos
    return output




codes = []
file = open("input5.txt", "r")
for line in file:
        line = line.rstrip()
        codes = line.split(",")
codes = [ int(x) for x in codes ]
originalCode = [ int(x) for x in codes ]

print "s1: " + str(compute(codes, 1))
codes = originalCode
print "s2: " + str(compute(codes, 5))
