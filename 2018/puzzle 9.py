import re
import sys
from collections import deque

players = 486
lastMarble = 70833

###########example###########
#players = 10
#lastMarble = 1618
#############################


circle = deque([0])
score = [0 for x in range(players)]
currentMarble = 0
for i in range(1, lastMarble + 1):
        if i % 23 == 0:
                circle.rotate(7)
                curElf = i % players
                score[curElf] += i
                score[curElf] += circle.pop()
                circle.rotate(-1)
        else:
                circle.rotate(-1)
                circle.append(i)
        #print circle

highScore = 0
for i in score:
        if i > highScore:
                highScore = i

print "s1: " + str(highScore)

circle = deque([0])
score = [0 for x in range(players)]
currentMarble = 0
for i in range(1, (lastMarble*100) + 1):
        if i % 23 == 0:
                circle.rotate(7)
                curElf = i % players
                score[curElf] += i
                score[curElf] += circle.pop()
                circle.rotate(-1)
        else:
                circle.rotate(-1)
                circle.append(i)
        #print circle

highScore = 0
for i in score:
        if i > highScore:
                highScore = i

print "s1: " + str(highScore)
