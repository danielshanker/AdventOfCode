import re

file = open("input.txt", "r")

for line in file:
    inputVal = line
total = 0
groupVal = 0
ignore = 0
inGarb = 0
garbCount = 0
for i in range (0, len(inputVal)):
    if ignore == 1:
        ignore = 0
        continue
    if inputVal[i] == '!':
        ignore = 1
        continue
    if inputVal[i] == '>':
        inGarb = 0
    if inGarb == 1:
        garbCount += 1
        continue
    if inputVal[i] == '<':
        inGarb = 1
    if inputVal[i] == '{':
        groupVal += 1
    if inputVal[i] == '}':
        total += groupVal
        groupVal -= 1
print ("sol 1: " + str(total))
print ("sol 2: " + str(garbCount))
