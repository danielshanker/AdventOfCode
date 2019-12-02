import re
import sys
from collections import deque

inp = 4151

###########example###########
#inp = 42
#############################

grid = [[0 for x in range(0,301)]for y in range(0,301)]
for x in range(1,301):
        for y in range(1,301):
                rackID = x + 10
                powerLevel = rackID * y
                powerLevel += inp
                powerLevel *= rackID
                
                hund = int(powerLevel/100) % 10
                hund -= 5
                grid[x][y] = hund
                

biggest = 0
answer = 0
for x in range(1,298):
        for y in range(1,298):
              square = grid[x][y]
              square += grid[x+1][y]
              square += grid[x+2][y]
              square += grid[x][y+1]
              square += grid[x+1][y+1]
              square += grid[x+2][y+1]
              square += grid[x][y+2]
              square += grid[x+1][y+2]
              square += grid[x+2][y+2]

              if square > biggest:
                      biggest = square
                      answer = str(x)+","+str(y)


print "s1: " + str(answer)

def calcSize(grid, x, y, size):
        square = 0
        for i in range(0, size):
                for j in range(0, size):
                        square += grid[x+i][y+j]
        return square


biggest = 0
answer = 0
count = 0
for x in range(301,1,-1):
        for y in range(301,1,-1):
                sr = 301 - x
                sr2 = 301 - y
                
                for size in range(0,min(sr, sr2)):
                        if x + size > 301 or y + size > 301:
                                continue
                        square = calcSize(grid, x, y, size)
              
                        if square > biggest:
                                biggest = square
                                answer = str(x)+","+str(y)+","+str(size+1)
                        if count % 10000 == 0:
                                print biggest
                                print answer
                                print str(x)+","+str(y)+","+str(size+1)
                                print "-----"
                        count +=1


print "s2: " + str(answer)
