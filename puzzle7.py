import re

file = open("input.txt", "r")
index = 0
step = 0
count = 0
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

file = open("input.txt", "r")
index = 0
step = 0
count = 0
towerParts = {}
weights = {}

for line in file:
    numList = line.split("\t")
    towerStep = line.split(" -> ")
    matchObj =  re.match("(.*) \((\d+)\)", towerStep[0], flags=0)
    if matchObj:
        weight = matchObj.group(2)
        base = matchObj.group(1)
    weights[base] = weight
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
print "sol 2: " + str(weights)
