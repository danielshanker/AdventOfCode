import re
import time

file = open("input16.txt", "r")
#file = open("example.txt", "r")

def fillOp():
    fill = {}
    for i in ('addr', 'addi', 'mulr', 'muli', 'banr', 'bani', 'borr', 'bori', 'setr', 'seti', 'gtir', 'gtri', 'gtrr', 'eqir', 'eqri', 'eqrr'):
        fill[i] = 1
    return fill

instr = [1]
i = 0
instr[0] = {}
for line in file:
    if line[0] == "%":
        break
    if line[0] == 'B':
        a, b = line.split('[')
        before, ignore = b.split(']')
        instr[i]['before'] =  before
    elif line[0] == 'A':
        a, b = line.split('[')
        before, ignore = b.split(']')
        instr[i]['after'] = before
    elif line[0] == '\n':
        i += 1
        instr.append({})
        continue
    else:
        instr[i]['inst'] = line

answer = 0
for curInst in instr:
    if not any(curInst):
        continue
    beforeRegister = []
    afterRegister = []
    beforeRegister = curInst['before'].split(',')
    afterRegister = curInst['after'].split(',')
    opCode, a, b, c = curInst['inst'].split(' ')
    opp = []
    #addr
    if int(beforeRegister[int(a)]) + int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('addr')
    #addi
    if int(beforeRegister[int(a)]) + int(b) == int(afterRegister[int(c)]):
        opp.append('addi')
    #mulr
    if int(beforeRegister[int(a)]) * int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('mulr')
    #muli
    if int(beforeRegister[int(a)]) * int(b) == int(afterRegister[int(c)]):
        opp.append('muli')
    #banr
    if int(beforeRegister[int(a)]) & int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('banr')
    #bani
    if int(beforeRegister[int(a)]) & int(b) == int(afterRegister[int(c)]):
        opp.append('bani')
    #borr
    if int(beforeRegister[int(a)]) | int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('borr')
    #bori
    if int(beforeRegister[int(a)]) | int(b) == int(afterRegister[int(c)]):
        opp.append('bori')
    #setr
    if int(beforeRegister[int(a)]) == int(afterRegister[int(c)]):
        opp.append('setr')
    #seti
    if int(a) == int(afterRegister[int(c)]):
        opp.append('bori')
    # gtir
    if int(a) > int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('gtir')
    if int(a) <= int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('gtir')
    # gtri
    if int(beforeRegister[int(a)]) > int(b) and int(afterRegister[int(c)]) == 1:
        opp.append('gtir')
    if int(beforeRegister[int(a)]) <= int(b) and int(afterRegister[int(c)]) == 0:
        opp.append('gtir')
    # gtrr
    if int(beforeRegister[int(a)]) > int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('gtrr')
    if int(beforeRegister[int(a)]) <= int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('gtrr')
    # gtir
    if int(a) == int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('gtir')
    if int(a) != int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('gtir')
    # gtri
    if int(beforeRegister[int(a)]) == int(b) and int(afterRegister[int(c)]) == 1:
        opp.append('gtir')
    if int(beforeRegister[int(a)]) != int(b) and int(afterRegister[int(c)]) == 0:
        opp.append('gtir')
    # gtrr
    if int(beforeRegister[int(a)]) == int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('gtrr')
    if int(beforeRegister[int(a)]) != int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('gtrr')
    if len(opp) >= 3:
        answer += 1

print "s1: " + str(answer)

answer = 0
opCodes = {}
for curInst in instr:
    if not any(curInst):
        continue
    beforeRegister = []
    afterRegister = []
    beforeRegister = curInst['before'].split(',')
    afterRegister = curInst['after'].split(',')
    opCode, a, b, c = curInst['inst'].split(' ')
    if opCode not in opCodes:
        opCodes[opCode] = fillOp()
    opp = []
    #addr
    if int(beforeRegister[int(a)]) + int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('addr')
    #addi
    if int(beforeRegister[int(a)]) + int(b) == int(afterRegister[int(c)]):
        opp.append('addi')
    #mulr
    if int(beforeRegister[int(a)]) * int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('mulr')
    #muli
    if int(beforeRegister[int(a)]) * int(b) == int(afterRegister[int(c)]):
        opp.append('muli')
    #banr
    if int(beforeRegister[int(a)]) & int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('banr')
    #bani
    if int(beforeRegister[int(a)]) & int(b) == int(afterRegister[int(c)]):
        opp.append('bani')
    #borr
    if int(beforeRegister[int(a)]) | int(beforeRegister[int(b)]) == int(afterRegister[int(c)]):
        opp.append('borr')
    #bori
    if int(beforeRegister[int(a)]) | int(b) == int(afterRegister[int(c)]):
        opp.append('bori')
    #setr
    if int(beforeRegister[int(a)]) == int(afterRegister[int(c)]):
        opp.append('setr')
    #seti
    if int(a) == int(afterRegister[int(c)]):
        opp.append('seti')
    # gtir
    if int(a) > int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('gtir')
    if int(a) <= int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('gtir')
    # gtri
    if int(beforeRegister[int(a)]) > int(b) and int(afterRegister[int(c)]) == 1:
        opp.append('gtri')
    if int(beforeRegister[int(a)]) <= int(b) and int(afterRegister[int(c)]) == 0:
        opp.append('gtri')
    # gtrr
    if int(beforeRegister[int(a)]) > int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('gtrr')
    if int(beforeRegister[int(a)]) <= int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('gtrr')
    # eqir
    if int(a) == int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('eqir')
    if int(a) != int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('eqir')
    # eqri
    if int(beforeRegister[int(a)]) == int(b) and int(afterRegister[int(c)]) == 1:
        opp.append('eqri')
    if int(beforeRegister[int(a)]) != int(b) and int(afterRegister[int(c)]) == 0:
        opp.append('eqri')
    # eqrr
    if int(beforeRegister[int(a)]) == int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 1:
        opp.append('eqrr')
    if int(beforeRegister[int(a)]) != int(beforeRegister[int(b)]) and int(afterRegister[int(c)]) == 0:
        opp.append('eqrr')

    for x in ('addr', 'addi', 'mulr', 'muli', 'banr', 'bani', 'borr', 'bori', 'setr', 'seti', 'gtir', 'gtri', 'gtrr', 'eqir', 'eqri', 'eqrr'):
        if x not in opp:
            opCodes[opCode][x] = 0
            

finalOp = [0 for x in range(16)]
count = 0
while count < 15:
    for i in range(16):
        found = []
        for j in opCodes[str(i)]:
            if opCodes[str(i)][j] == 1 and j not in finalOp:
                found.append(j)
        if len(found) == 1:
            finalOp[i] = found[0]
            count +=1




file = open("input16.txt", "r")
flag = 0
test = []
for line in file:
    if line[0] == "%":
        flag = 1
        continue
    if flag == 0:
        continue
    test.append(line)

register = [0 for x in range(4)]

for curInst in test:
    opCode, a, b, c = curInst.split(' ')
    opCode = finalOp[int(opCode)]
    #print opCode +  " " a + " " + b + " " + c + " " 
    #print register
    a = int(a)
    b = int(b)
    c = int(c)
    #addr
    if opCode == 'addr':
        register[c] = register[a] + register[b]
    #addi
    if opCode == 'addi':
        register[c] = register[a] + b
    #mulr
    if opCode == 'mulr':
        register[c] = register[a] * register[b]
    #muli
    if opCode == 'muli':
        register[c] = register[a] * b
    #banr
    if opCode == 'banr':
        register[c] = register[a] & register[b]
    #bani
    if opCode == 'bani':
        register[c] = register[a] & b
    #bori
    if opCode == 'bori':
        register[c] = register[a] | b
    #borr
    if opCode == 'borr':
        register[c] = register[a] | register[b]
    #setr
    if opCode == 'setr':
        register[c] = register[a]
    #seti
    if opCode == 'seti':
        register[c] = a
    # gtir
    if opCode == 'gtir':
        if a > register[b]:
            register[c] = 1
        else:
            register[c] = 0
    # gtri
    if opCode == 'gtri':
        if register[a] > b:
            register[c] = 1
        else:
            register[c] = 0
    # gtrr
    if opCode == 'gtrr':
        if register[a] > register[b]:
            register[c] = 1
        else:
            register[c] = 0
    # eqir
    if opCode == 'eqir':
        if a == register[b]:
            register[c] = 1
        else:
            register[c] = 0
    # eqri
    if opCode == 'eqri':
        if register[a] == b:
            register[c] = 1
        else:
            register[c] = 0
    # eqrr
    if opCode == 'eqrr':
        if register[a] == register[b]:
            register[c] = 1
        else:
            register[c] = 0
#    print register
print "s2: " + str(register[0])
