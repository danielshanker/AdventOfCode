import re
import random

file = open("input.txt", "r")
inputVals = []
for line in file:
    inputVals.append(line.rstrip())
    
global s
global l
global strong
global long
strong = 0
long = 0
s = []
l = []

def buildBridge(curVal, bridge, newBridge):
    longest = 0
    global strong
    global long
    for i in bridge:
        a,b = i.split('/')
        send = ''
        if curVal == a:
            send = b
        if curVal == b:
            send = a
        if send != '':
            bridgeB = []
            for j in bridge:
                if j != i:
                    bridgeB.append(j)
            if a == curVal:
                send = b
            else:
                send = a
            newBridgeB = []
            for j in newBridge:
                newBridgeB.append(j)
            newBridgeB.append(i)
            buildBridge(send, bridgeB, newBridgeB)
    length = 0
    for i in newBridge:
        a,b = i.split('/')
        length += int(a) + int(b)
    s.append(length)
    l.append(len(newBridge))
    if long < len(newBridge):
        long = len(newBridge)
    elif long == len(newBridge):
        if strong < length:
            strong = length
            long = len(newBridge)


buildBridge('0', inputVals, [])
print ("sol 1 : " + str(max(s)))
print ("sol 2 : " + str(strong))
