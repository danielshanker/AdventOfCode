import re
import sys

nodes = {}
x = 0
y = 0
file = open("input.txt", "r")
for line in file:
    nodes[x] = {}
    line = line.rstrip()
    for i in range(0,len(line)):
        nodes[x][y] = line[i]
        y += 1
    x += 1
    y = 0

x = (len(nodes)-1)/2
y = (len(nodes[0])-1)/2
direction = 0
infected = 0

for i in range(0,10000):
    if x not in nodes:
        nodes[x] = {}
        nodes[x][y] = '.'
    if y not in nodes[x]:
        nodes[x][y] = '.'

    if nodes[x][y] == '#':
        direction -= 1
        if direction == -1:
            direction = 3
        nodes[x][y] = '.'
    elif nodes[x][y] == '.':
        direction += 1
        if direction == 4:
            direction = 0
        nodes[x][y] = '#'
        infected += 1

    if direction == 0:
        x -= 1
    if direction == 1:
        y -= 1
    if direction == 2:
        x += 1
    if direction == 3:
        y += 1

print "sol 1: " + str(infected)

nodes = {}
x = 0
y = 0
file = open("input.txt", "r")
for line in file:
    nodes[x] = {}
    line = line.rstrip()
    for i in range(0,len(line)):
        nodes[x][y] = line[i]
        y += 1
    x += 1
    y = 0

x = (len(nodes)-1)/2
y = (len(nodes[0])-1)/2
direction = 0
infected = 0

for i in range(0,10000000):
    if x not in nodes:
        nodes[x] = {}
        nodes[x][y] = '.'
    if y not in nodes[x]:
        nodes[x][y] = '.'

    if nodes[x][y] == '#':
        direction -= 1
        if direction == -1:
            direction = 3
        nodes[x][y] = 'F'
    elif nodes[x][y] == '.':
        direction += 1
        if direction == 4:
            direction = 0
        nodes[x][y] = 'W'
    elif nodes[x][y] == 'F':
        direction += 2
        direction %= 4
        nodes[x][y] = '.'
    elif nodes[x][y] == 'W':
        nodes[x][y] = '#'
        infected += 1

    if direction == 0:
        x -= 1
    if direction == 1:
        y -= 1
    if direction == 2:
        x += 1
    if direction == 3:
        y += 1

print "sol 2: " + str(infected)
