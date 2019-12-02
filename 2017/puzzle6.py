file = open("input.txt", "r")
index = 0
step = 0
count = 0

for line in file:
    numList = line.split("\t")
numList = [int(x) for x in numList]
seen = {}
target = ''

while (1):
    index = 0
    highest = max(numList)
    string = ' '.join(str(x) for x in numList)
    if string == target:
        break
    if string in seen and target == '':
        target = string
    if target != '':
        count += 1
    step += 1
    seen[string] = 1
    for value in range(0, len(numList)):
        if numList[index] == highest:
            numList[index] = 0
            break
        index+=1
    index +=1
    while (1):
        if highest <= 0:
            break
        if index >= len(numList):
            index = 0
        numList[index] += 1
        index += 1
        highest -= 1
print ("sol 1: " + str(step))
print ("sol 2: " + str(count))
