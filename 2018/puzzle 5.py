import re
import time

file = open("input5.txt", "r")

for line in file:
        polymer = line

#polymer = "dabAcCaCBAcCcaDA"
newPolymer = polymer
originalPolymer = polymer
index = 0
count = 0

##while 1:
##        break
##        lastChar = ""
##        out = 1
##        count +=1
##        index = 0
##        for char in polymer:
##                if char.lower() == lastChar.lower():
##                        if (char.islower() and lastChar.isupper()) or (char.isupper() and lastChar.islower()):
##                                newPolymer = polymer[:index-1] + polymer[index+1:]
##                                polymer = newPolymer
##                                out = 0
##                                break
##                lastChar = char
##                index +=1
##        if out == 1:
##                break
##
##print "s1: " + str(len(polymer))


polymer = originalPolymer
index = 0
count = 0
smallest = len(polymer)
from string import ascii_lowercase
for i in (ascii_lowercase):
        polymer = originalPolymer
        polymer = polymer.replace(i, '')
        polymer = polymer.replace(i.upper(), '')
        newPolymer = polymer
        while 1:
                lastChar = ""
                out = 1
                count +=1
                index = 0
                for char in polymer:
                        if char.lower() == lastChar.lower():
                                if (char.islower() and lastChar.isupper()) or (char.isupper() and lastChar.islower()):
                                        newPolymer = polymer[:index-1] + polymer[index+1:]
                                        polymer = newPolymer
                                        out = 0
                                        break
                        lastChar = char
                        index +=1
                if count % 500 == 0:
                        print count
                if out == 1:
                        break
        print i
        if smallest > len(polymer):
                smallest = len(polymer)
        

print "s2: " + str(smallest)
