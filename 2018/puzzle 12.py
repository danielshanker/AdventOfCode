import re
import sys
from collections import deque

answer = 0
file = open("input12.txt", "r")
#file = open ("example.txt", "r")
padding = 5000
pots = ['.' for x in range(padding)]
rules = []
for line in file:
        matchObj = re.match("initial state: (.*)", line)
        if matchObj:
                potList = matchObj.group(1)
                for i in potList:
                        pots.append(i.rstrip())
                continue
        start, end = line.split(' => ')
        rule = {}
        rule['ll'] = start[0]
        rule['l'] = start[1]
        rule['c'] = start[2]
        rule['r'] = start[3]
        rule['rr'] = start[4]
        rule['n'] = end.rstrip()
        rules.append(rule)
for i in range(padding):
        pots.append('.')

newPots = list(pots)


lastActivePots = ''.join(pots)
pattern = 0
finalGen = 0
for gen in range(500):
        for i in range(2,len(pots)-2):
                found = 0
                for rule in rules:
                        if rule['c'] == pots[i] and rule['l'] == pots[i-1] and rule['ll'] == pots[i-2] and rule['r'] == pots[i+1] and rule['rr'] == pots[i+2]:                                
                                newPots[i] = rule['n']
                                found = 1
                if found == 0:
                        newPots[i] = '.'
        pots = list(newPots)
        
        answer = 0
        a = -1 * padding
        for i in pots:
                if i == '#':
                        answer += a
                a += 1
        if (gen == 20):
                print "s1: " + str(answer)

        active = 0
        activePots = ''
        matchObj = re.match("^\.+(#.*#)\.+$", ''.join(pots))
        if matchObj:
                activePots = matchObj.group(1)
        if lastActivePots == activePots:
                pattern = 1
        lastActivePots = activePots
        if pattern == 1:
                x = 0
                for i in pots:
                        if i == '#':
                                pattern = x
                                finalGen = gen
                                break
                        x+=1
                break

        #print str(gen)+ ' ' + activePots

bigNum = 50000000000
answer = 0
first = bigNum + pattern - finalGen - padding - 1
a = first
for i in lastActivePots:
        if i == '#':
                answer += a
        a += 1
print "s2: " + str(answer)

