import re

file = open("input.txt", "r")

for line in file:
    inputVal = line.rstrip()
progs = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p']
progsI = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p']


def spin(progs, num):
    newP = []
    for i in range(len(progs)-num, len(progs)):
        newP.append(progs[i])
    for i in range(0, len(progs) - num):
        newP.append(progs[i])
    return newP

def exchange(progs, a, b):
    temp = progs[a]
    progs[a] = progs[b]
    progs[b] = temp
    return progs

def partner(progs, a, b):
    i = progs.index(a)
    j = progs.index(b)
    temp = progs[i]
    progs[i] = progs[j]
    progs[j] = temp
    return progs

inst = inputVal.split(',')
count = 0

while (1):
    count += 1
    for i in inst:
        matchObj = re.match("(\w)(\w+)/(\w+)", i, flags=0)
        if matchObj:
            instruction = matchObj.group(1)
            a = matchObj.group(2)
            b = matchObj.group(3)
        else:
            matchObj = re.match("s(\w+)", i, flags=0)
            instruction = 's'
            a = matchObj.group(1)
        if instruction == 's':
            progs = spin(progs, int(a))
        elif instruction == 'x':
            progs = exchange(progs, int(a), int(b))
        elif instruction == 'p':
            progs = partner(progs, a, b)
        else:
            print "an error occurred "+ str(i)
    if count == 1:
        print ("sol 1: " + ''.join([str(x) for x in progs]))
    if progs == progsI:
        break

loopCount = 1000000000 % count
for j in range (0, loopCount):
    count += 1
    for i in inst:
        matchObj = re.match("(\w)(\w+)/(\w+)", i, flags=0)
        if matchObj:
            instruction = matchObj.group(1)
            a = matchObj.group(2)
            b = matchObj.group(3)
        else:
            matchObj = re.match("s(\w+)", i, flags=0)
            instruction = 's'
            a = matchObj.group(1)
        if instruction == 's':
            progs = spin(progs, int(a))
        elif instruction == 'x':
            progs = exchange(progs, int(a), int(b))
        elif instruction == 'p':
            progs = partner(progs, a, b)
        else:
            print "an error occurred "+ str(i)


print ("sol 2: " + ''.join([str(x) for x in progs]))
