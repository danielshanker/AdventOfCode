import re

file = open("input.txt", "r")
index = 0
step = 0
maze = {}
inst = []

for line in file:
    matchObj =  re.match("(.*) (.*) (.*) if (.*) (.*) (.*)", line, flags=0)
    if matchObj:
        if matchObj.group(1) not in maze:
            maze[matchObj.group(1)] = 0
    inst.append(line)

biggest2 = 0

for i in range (0, len(inst)):
    doIt = 0
    matchObj =  re.match("(.*) (.*) (.*) if (.*) (.*) (.*)", inst[i], flags=0)
    if matchObj:
        a = matchObj.group(1)
        b = matchObj.group(2)
        c = int(matchObj.group(3))
        d = matchObj.group(4)
        e = matchObj.group(5)
        f = int(matchObj.group(6))
    if e == '==':
        if maze[d] == f:
            doIt = 1
    elif e == '!=':
        if maze[d] != f:
            doIt = 1
    elif e == '>':
        if maze[d] > f:
            doIt = 1
    elif e == '>=':
        if maze[d] >= f:
            doIt = 1
    elif e == '<':
        if maze[d] < f:
            doIt = 1
    elif e == '<=':
        if maze[d] <= f:
            doIt = 1
    if doIt == 1:
        if b == 'inc':
            maze[a] += c
        else:
            maze[a] -= c
    if maze[a] > biggest2:
        biggest2 = maze[a]

biggest = maze[a]
for key in maze:
    if maze[key] > biggest:
        biggest = maze[key]

print ("sol 1: " + str(biggest))
print ("sol 2: " + str(biggest2))
