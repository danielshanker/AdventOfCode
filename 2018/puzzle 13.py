import re
import sys

answer = 0
file = open("input13.txt", "r")
file = open ("example.txt", "r")

track = [[0 for x in range(6)] for y in range(15)]
cart = {}
cartNum = 0

y = 0
for line in file:
    x = 0
    for i in range(15):
        trackPiece = {}
        trackPiece['step'] = 0
        if i >= len(line):
            trackPiece['orig'] = ' ' 
            trackPiece['current'] = ' ' 
        else:
            trackPiece['current'] = line[i]
            if line[i] == '>' or line[i] == '<':
                trackPiece['orig'] = '-'
                cart[cartNum] = str(x) + ',' + str(y)
            elif line[i] == '^' or line[i] == 'v':
                trackPiece['orig'] = '-'
            else:
                trackPiece['orig'] = line[i]
        track[x][y] = trackPiece 
        x += 1
    y+=1


for y in range (6):
    for x in range(15):
        sys.stdout.write(str(track[x][y]['orig']))
    print ''

for y in range (6):
    for x in range(15):
        sys.stdout.write(str(track[x][y]['current']))
    print ''

print "VVVVVVVVVVVVVVVVVVVVVVVV"

