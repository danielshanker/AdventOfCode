import re

twoRules = []
threeRules = []
file = open("input.txt", "r")
for line in file:
    line = line.rstrip()
    if len(line) == 20:
        twoRules.append(line)
    else:
        threeRules.append(line)
art = []
art.append('.#.')
art.append('..#')
art.append('###')

def rotateSection(section):
    rotSection = []
    for i in range(0,len(section)):
        rotSection.append('')
        for j in range(0,len(section)):
            rotSection[i] += str(section[len(section)-1-j][i])
    return rotSection


def checkRule(section, rule):
    inp, out = rule.split(' => ')
    a = inp.split('/')
    match = 0
    ra = []
    #standard
    for i in range(0,len(section)):
        if section[i] == a[i]:
            match += 1
        else: 
            break
    if match == len(section):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    match = 0
    # verticle flip
    for i in range(0,len(section)):
        if section[len(section)-1-i] == a[i]:
            match += 1
        else: 
            break
    if match == len(section):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    match = 0
    # horizontal flip
    for i in range(0,len(section)):
        if section[i][::-1] == a[i]:
            match += 1
        else: 
            break
    if match == len(section):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    match = 0
    # rotate 90 
    rotSection = rotateSection(section)
    for i in range(0,len(rotSection)):
        if rotSection[i] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra
    match = 0
    # verticle flip
    for i in range(0,len(rotSection)):
        if rotSection[len(rotSection)-1-i] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    match = 0
    # horizontal flip
    for i in range(0,len(rotSection)):
        if rotSection[i][::-1] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    # rotate 180 
    match = 0
    rotSection = rotateSection(rotSection)
    for i in range(0,len(rotSection)):
        if rotSection[i] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra
    match = 0
    # verticle flip
    for i in range(0,len(rotSection)):
        if rotSection[len(rotSection)-1-i] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    match = 0
    # horizontal flip
    for i in range(0,len(rotSection)):
        if rotSection[i][::-1] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra
    # rotate 270 
    match = 0
    rotSection = rotateSection(rotSection)
    for i in range(0,len(rotSection)):
        if rotSection[i] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra
    # verticle flip
    match = 0
    for i in range(0,len(rotSection)):
        if rotSection[len(rotSection)-1-i] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra

    match = 0
    # horizontal flip
    for i in range(0,len(rotSection)):
        if rotSection[i][::-1] == a[i]:
            match += 1
        else: 
            break
    if match == len(rotSection):
        ret = out.split('/')
        for i in ret:
            ra.append(i)
        return ra
    return section 

    
    

def byTwo(art, rules):
    i = 0
    newSec = [['' for x in range(len(art)/2)] for y in range(len(art)/2)] 
    while i < len(art):
        j = 0
        while j < len(art):
            for rule in rules:
                section = [art[i][j:j+2], art[i+1][j:j+2]]
                a = checkRule(section, rule)
                if len(a) == 3:
                    newSec[i/2][j/2] = a
                    break
            j += 2
        i += 2

    newArt = [''] * len(newSec)*len(newSec[0][0])
    count = 0
    x = 0
    for i in range(0, len(newSec)):
        for j in range(0,len(newSec[0][0])):
            for k in range(0, len(newSec[0])):
                newArt[count] += newSec[i][k][j]
                x += 1
                if x % len(newSec) == 0:
                    count += 1

    return newArt


def byThree(art, rules):
    i = 0
    newSec = [['' for x in range(len(art)/3)] for y in range(len(art)/3)] 
    while i < len(art):
        j = 0
        while j < len(art):
            for rule in rules:
                section = [art[i][j:j+3], art[i+1][j:j+3], art[i+2][j:j+3]]
                a = checkRule(section, rule)
                if len(a) == 4:
                    newSec[i/3][j/3] = a
                    break
            j += 3
        i += 3

    newArt = [''] * len(newSec)*len(newSec[0][0])
    count = 0
    x = 0
    for i in range(0, len(newSec)):
        for j in range(0,len(newSec[0][0])):
            for k in range(0, len(newSec[0])):
                newArt[count] += newSec[i][k][j]
                x += 1
                if x % len(newSec) == 0:
                    count += 1

    return newArt


for i in range(0,18):
    if len(art) % 2 == 0:
        art = byTwo(art, twoRules)
    elif len(art) % 3 == 0:
        art = byThree(art, threeRules)
    if i == 4:
        count = 0
        for i in art:
            count += i.count('#')
        print "sol 1: " + str(count)

count = 0
for i in art:
    count += i.count('#')
print "sol 2: " + str(count)
