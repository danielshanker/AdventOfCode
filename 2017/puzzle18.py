import re
import time

file = open("input.txt", "r")
inst = []
for line in file:
    inst.append(line.rstrip())

sound = {}
regs = {}

i = 0
while (i < len(inst)):
    jump = 0
    matchObj = re.match("(\w+) (\w+)\s?(-*\w*)", inst[i], flags=0)
    if matchObj:
        instruction = matchObj.group(1)
        a = matchObj.group(2)
        b = matchObj.group(3)
    else:
        print "an error has occurred " + inst[i]
    if a not in regs and a != '1':
        regs[a] = 0
    if instruction == 'snd':
        sound[a] = int(regs[a])
    elif instruction == 'set':
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] = int(b)
        else:
            regs[a] = int(regs[b])
    elif instruction == 'add':
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] += int(b)
        else:
            regs[a] += int(regs[b])
    elif instruction == 'mul':
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] *= int(b)
        else:
            regs[a] *= int(regs[b])
    elif instruction == 'mod':
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] %= int(b)
        else:
            regs[a] %= int(regs[b])
    elif instruction == 'rcv' and int(regs[a]) != 0:
        if a in sound:
            print "sol 1: " + str(sound[a])
            break
    elif instruction == 'jgz' and int(regs[a]) > 0:
        m = re.match("-?\d+", b, flags=0)
        if m:
            i += int(b)
        else:
            i += int(regs[b])
        jump = 1
    if jump == 0:
        i += 1


file = open("input.txt", "r")
inst = []
for line in file:
    inst.append(line.rstrip())

count = 0
regsA = {}
regsB = {}
regsA['p'] = 0
regsB['p'] = 1
queueA = []
queueB = []

i = 0
j = 0
waitA = 0
waitB = 0

while (i < len(inst) or j < len(inst)):
    if waitA == 1 and waitB == 1:
        break
    if waitA == 0 and i < len(inst):
        jumpA = 0
        matchObj = re.match("(\w+) (\w+)\s?(-*\w*)", inst[i], flags=0)
        if matchObj:
            instruction = matchObj.group(1)
            a = matchObj.group(2)
            b = matchObj.group(3)
        else:
            print "an error has occurred " + inst[i]
        if a not in regsA and a != '1':
            regsA[a] = 0
        if instruction == 'snd':
            m = re.match("-?\d+", a, flags=0)
            if m:
                k = int(a)
            else:
                k = int(regsA[a])
            waitB = 0
            queueB.append(int(k))
        elif instruction == 'set':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsA[a] = int(b)
            else:
                regsA[a] = int(regsA[b])
        elif instruction == 'add':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsA[a] += int(b)
            else:
                regsA[a] += int(regsA[b])
        elif instruction == 'mul':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsA[a] *= int(b)
            else:
                regsA[a] *= int(regsA[b])
        elif instruction == 'mod':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsA[a] %= int(b)
            else:
                regsA[a] %= int(regsA[b])
        elif instruction == 'rcv':
            if len(queueA) == 0:
                waitA = 1
                jumpA = 1
            else:
                regsA[a] = queueA[0]
                queueA.pop(0)
        elif instruction == 'jgz':
            m = re.match("-?\d+", a, flags=0)
            if m:
                k = a
            else:
                k = regsA[a]
            if k > 0:
                m = re.match("-?\d+", b, flags=0)
                if m:
                    i += int(b)
                else:
                    i += int(regsA[b])
                jumpA = 1
        if jumpA == 0:
            i += 1
    
    if waitB == 0 and j < len(inst):
        jumpB = 0
        waitB = 0
        matchObj = re.match("(\w+) (\w+)\s?(-*\w*)", inst[j], flags=0)
        if matchObj:
            instruction = matchObj.group(1)
            a = matchObj.group(2)
            b = matchObj.group(3)
        else:
            print "an error has occurred " + inst[j]
        if a not in regsB and a != '1':
            regsB[a] = 0
        if instruction == 'snd':
            m = re.match("-?\d+", a, flags=0)
            if m:
                k = int(a)
            else:
                k = int(regsB[a])
            waitA = 0
            count += 1
            queueA.append(int(k))
        elif instruction == 'set':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsB[a] = int(b)
            else:
                regsB[a] = int(regsB[b])
        elif instruction == 'add':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsB[a] += int(b)
            else:
                regsB[a] += int(regsB[b])
        elif instruction == 'mul':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsB[a] *= int(b)
            else:
                regsB[a] *= int(regsB[b])
        elif instruction == 'mod':
            m = re.match("-?\d+", b, flags=0)
            if m:
                regsB[a] %= int(b)
            else:
                regsB[a] %= int(regsB[b])
        elif instruction == 'rcv':
            if len(queueB) == 0:
                waitB = 1
                jumpB = 1
            else:
                regsB[a] = queueB[0]
                queueB.pop(0)
        elif instruction == 'jgz':
            m = re.match("-?\d+", a, flags=0)
            if m:
                k = a
            else:
                k = regsB[a]
            if k > 0:
                m = re.match("-?\d+", b, flags=0)
                if m:
                    j += int(b)
                else:
                    j += int(regsB[b])
                jumpB = 1
        if jumpB == 0:
            j += 1
    if j >= len(inst):
        waitB = 1
    if i >= len(inst):
        waitA = 1
    
print "sol 2: " + str(count)
