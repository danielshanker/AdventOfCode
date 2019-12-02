import re

file = open("input1.txt", "r")
frequency = 0
for line in file:
        frequency = frequency + int(line)
print "s1: " + str(frequency)

frequency = 0
map = {}
out = 0
inp = []
file = open("input1.txt", "r")
for line in file:
        inp.append(line)

while 1:        
        for line in inp:
                frequency = frequency + int(line)
                if frequency in map:
                        out = 1
                        break;
                map[frequency] = 1
        if out == 1:
                break

print "s2: " + str(frequency)
