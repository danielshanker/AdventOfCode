import re

recipes = 894501
scores = [3, 7]

pos1 = 0
pos2 = 1
answer = ''

for i in range (0, recipes+10):
    r1 = scores[pos1]
    r2 = scores[pos2]
    newR = r1 + r2
    if newR < 10:
        scores.append(newR)
    else:
        a = str(newR)
        scores.append(int(a[0]))
        scores.append(int(a[1]))
    pos1 = (pos1 + 1 + scores[pos1]) % len(scores)
    pos2 = (pos2 + 1 + scores[pos2]) % len(scores)

for i in range(recipes, recipes + 10):
    answer += str(scores[i])
print "s1: " + str(answer)

scores = [3, 7]

pos1 = 0
pos2 = 1
answer = ''

while (1):
    r1 = scores[pos1]
    r2 = scores[pos2]
    newR = r1 + r2
    if newR < 10:
        scores.append(newR)
    else:
        a = str(newR)
        scores.append(int(a[0]))
        scores.append(int(a[1]))
    pos1 = (pos1 + 1 + scores[pos1]) % len(scores)
    pos2 = (pos2 + 1 + scores[pos2]) % len(scores)

    if len(scores) > len(str(recipes)):
        a = ''
        for j in range(7, 1, -1):
            a += str(scores[len(scores)-j])
        if a == str(recipes):
            break
print "s2: " + str(len(scores) - len(str(recipes)))
