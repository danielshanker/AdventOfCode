import re
import time

file = open("input.txt", "r")
inst = []
for line in file:
    inst.append(line.rstrip())

regs = {}
for i in ('a','b','c','d','e','f','g','h'):
    regs[i] = 0
count = 0
i = 0
while (i < len(inst)):
    jump = 0
    matchObj = re.match("(\w+) (\w+)\s?(-*\w*)", inst[i], flags=0)
    if matchObj:
        instruction = matchObj.group(1)
        a = matchObj.group(2)
        b = matchObj.group(3)
    else:
        print ("an error has occurred " + inst[i])
    if a not in regs and a != '1':
        regs[a] = 0
    if instruction == 'sub':
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] -= int(b)
        else:
            regs[a] -= int(regs[b])
    elif instruction == 'set':
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] = int(b)
        else:
            regs[a] = int(regs[b])
    elif instruction == 'mul':
        count += 1
        m = re.match("-?\d+", b, flags=0)
        if m:
            regs[a] *= int(b)
        else:
            regs[a] *= int(regs[b])
    elif instruction == 'jnz':
        m = re.match("-?\d+", a, flags=0)
        if m:
            val = int(a)
        else:
            val = int(regs[a])
        if val != 0:
            m = re.match("-?\d+", b, flags=0)
            if m:
                i += int(b)
            else:
                i += int(regs[b])
            jump = 1
    if jump == 0:
        i += 1
print ("sol 1: " + str(count))



a = 1
b = 99*100 + 100000
c = b + 17000
d = 0
e = 0
f = 0
g = 0
h = 0
x = 0
while b <= c:
    for d in range(2, b):
        if b%d == 0:
            h += 1
            break
    b += 17
    x += 1
print ("sol 2: " + str(h))





