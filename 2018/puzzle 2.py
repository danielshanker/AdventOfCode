import re

file = open("input2.txt", "r")
boxID = 0
two = 0
three = 0
for line in file:
        twoFlag = 0
        threeFlag = 0
        map = {}
        for l in line:
                if l in map:
                        map[l] += 1
                else:
                        map[l] = 1
        for key in map:
                if map[key] == 2 and twoFlag == 0:
                        two +=1
                        twoFlag = 1
                if map[key] == 3 and threeFlag == 0:
                        three +=1
                        threeFlag = 1

boxID = three * two
print "s1: " + str(boxID)

boxID = 0

out = 0
inp = []
value1 = ""
value2 = ""
file = open("input2.txt", "r")
for line in file:
        inp.append(line)

for ID in inp:
        for secondID in inp:
                i = 0
                misses = 0
                for char in ID:
                        if ID == secondID:
                                next
                        if ID[i] != secondID[i]:
                                misses += 1
                        if misses >= 2:
                                break
                        i += 1
                if misses == 1:
                        value1 = ID
                        value2 = secondID
                        value = ID + " " + secondID
                        out = 1
                        break
                if out == 1:
                        break

i = 0
finalAnswer = ""
for char in value1:
        if value1[i] == value2[i]:
                finalAnswer += value1[i]
        i += 1
        
print "s2: " + str(finalAnswer)
