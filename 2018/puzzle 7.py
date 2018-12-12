import re
import string

di=dict(zip(string.letters,[ord(c)%32 for c in string.letters]))

answer = 0
file = open("input7.txt", "r")
#file = open ("example.txt", "r")
steps = []

for line in file:
        steps.append(line)
        print line

nodes = []
requires = {}

for a in steps:
        matchObj = re.match("Step (.) must be finished before step (.) can begin.", a)
        if matchObj:
                first = matchObj.group(1)
                second = matchObj.group(2)
                if first not in nodes:
                        nodes.append(first)
                if second not in nodes:
                        nodes.append(second)
                if second not in requires:
                        requires[second] = {}
                requires[second][first] = 0

nodes.sort()
completed = []
ready = {}
for i in range(26):
        for a in nodes:
                flag = 0
                if a in completed:
                        continue
                if a not in requires or a in ready:
                        flag = 1
                        completed.append(a)
                        for b in requires:
                                if b in requires and a in requires[b] and requires[b][a] == 0:
                                        requires[b][a] = 1
                for b in requires:
                        doneFlag = 0
                        for c in requires[b]:
                                if requires[b][c] == 0:
                                        doneFlag = 1
                                        break
                        if doneFlag == 0:
                                ready[b] = 1
                if flag == 1:
                        break
                        

print "s1: " + str(''.join(completed))

########################################
answer = 0

nodes = []
requires = {}

for a in steps:
        matchObj = re.match("Step (.) must be finished before step (.) can begin.", a)
        if matchObj:
                first = matchObj.group(1)
                second = matchObj.group(2)
                if first not in nodes:
                        nodes.append(first)
                if second not in nodes:
                        nodes.append(second)
                if second not in requires:
                        requires[second] = {}
                requires[second][first] = 0

nodes.sort()
completed = []
ready = {}
inProgress = {}
free = 5
time = {}
totalTime = 0
for i in nodes:
        time[i] = di[i] + 60

while(1):
        #print free
        #print ready
        totalTime += 1
        for a in nodes:
                #print a
                if a in completed:
                        continue
                if a not in requires or a in ready:
                        if a not in inProgress:
                                if free > 0:
                                        flag = 1
                                        inProgress[a] = 1
                                        free -= 1
                                else:
                                        break
        for x in inProgress:
                time[x] -= 1
                if time[x] == 0:
                        print free
                        print completed
                        free += 1
                        completed.append(x)
                        print free
                        print completed
                        for b in requires:
                                if b in requires and x in requires[b] and requires[b][x] == 0:
                                        requires[b][x] = 1
        for x in completed:
                if x in inProgress:
                       del inProgress[x]

        for b in requires:
                doneFlag = 0
                for c in requires[b]:
                        if requires[b][c] == 0:
                                doneFlag = 1
                                break
                if doneFlag == 0:
                        ready[b] = 1
        if len(completed) == len(nodes):
                break

print "s2: " + str(totalTime)


