import re

file = open("input.txt", "r")
towerParts = {}

for line in file:
    numList = line.split("\t")
    towerStep = line.split(" -> ")
    base = line.split(" (")[0]
    if len(towerStep) < 2:
        continue
    towerParts[base] = towerStep[1]

while (1):
    bottom = 1
    for i in towerParts:
        if re.search(base, towerParts[i], flags=0):
            base = i
            bottom = 0
            break
    if bottom == 1:
        break
print "sol 1: " + base
firstBase = base

file = open("input.txt", "r")
towerParts = {}
weights = {}

for line in file:
    numList = line.split("\t")
    matchObj =  re.match("(.*) \((\d+)\)", line, flags=0)
    if matchObj:
        base = matchObj.group(1)
        weights[base] = matchObj.group(2)
    towerParts[base] = ['top']
    matchObj =  re.match(".* -> (.*)", line, flags=0)
    if matchObj:
        towerParts[base] = matchObj.group(1).split(", ")

def findTop(baseLevel):
    weight = 0
    if baseLevel == 'top':
        return 0
    output = []
    for i in range (0, len(towerParts[baseLevel])):
        output.append(findTop(towerParts[baseLevel][i]))
    for i in range (0, len(output)):
        weight += output[i]
        if output[0] != output[i]:
            print "Inbalance " + baseLevel + " " + str(output[i] - output[0]) + str(output)
            for j in range (0, len(output)):
                print weights[towerParts[baseLevel][j]]
            return int(weights[baseLevel]) + weight
    return int(weights[baseLevel]) + weight

findTop(firstBase)


