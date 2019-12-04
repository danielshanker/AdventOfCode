import re

wire1 = []
wire2 = []
file = open("input3.txt", "r")
i = 0
for line in file:
    line = line.rstrip()
    if i == 0:
        wire1 = line.split(",")
        i = i + 1
    if i == 1:
        wire2 = line.split(",")

#wire1 = "R8,U5,L5,D3".split(",")
#wire2 = "U7,R6,D4,L4".split(",")

grid = {}
grid[0] = {}
grid[0][0] = 1
curx = 0
cury = 0
for inst in wire1:
    direction = inst[0:1]
    move = inst[1:]
    if direction == 'R':
        for i in range (1,int(move)+1):
            curx = curx+1
            if curx not in grid:
                grid[curx] = {}
            grid[curx][cury] = 1
    if direction == 'L':
        for i in range (1,int(move)+1):
            curx = curx-1
            if curx not in grid:
                grid[curx] = {}
            grid[curx][cury] = 1
    if direction == 'U':
        for i in range (1,int(move)+1):
            cury = cury+1
            grid[curx][cury] = 1
    if direction == 'D':
        for i in range (1,int(move)+1):
            cury = cury-1
            grid[curx][cury] = 1

curx = 0
cury = 0
for inst in wire2:
    direction = inst[0:1]
    move = inst[1:]
    if direction == 'R':
        for i in range (1,int(move)+1):
            curx = curx+1
            if curx in grid and cury in grid[curx]:
                grid[curx][cury] = 2
    if direction == 'L':
        for i in range (1,int(move)+1):
            curx = curx-1
            if curx in grid and cury in grid[curx]:
                grid[curx][cury] = 2
    if direction == 'U':
        for i in range (1,int(move)+1):
            cury = cury+1
            if curx in grid and cury in grid[curx]:
                grid[curx][cury] = 2
    if direction == 'D':
        for i in range (1,int(move)+1):
            cury = cury-1
            if curx in grid and cury in grid[curx]:
                grid[curx][cury] = 2


minVal = 10000000
for i in grid:
    for j in grid[i]:
        if grid[i][j] == 2:
            if abs(i) + abs(j) < minVal:
                minVal = abs(i) + abs(j)

print "s1 = " + str(minVal)


grid = {}
grid[0] = {}
grid[0][0] = {}
curx = 0
cury = 0
step = 0
for inst in wire1:
    direction = inst[0:1]
    move = inst[1:]
    if direction == 'R':
        for i in range (1,int(move)+1):
            step = step + 1
            curx = curx+1
            if curx not in grid:
                grid[curx] = {}
            if cury not in grid[curx]:
                grid[curx][cury] = {}
                grid[curx][cury]["steps"] = step
    if direction == 'L':
        for i in range (1,int(move)+1):
            step = step + 1
            curx = curx-1
            if curx not in grid:
                grid[curx] = {}
            if cury not in grid[curx]:
                grid[curx][cury] = {}
                grid[curx][cury]["steps"] = step
    if direction == 'U':
        for i in range (1,int(move)+1):
            step = step + 1
            cury = cury+1
            if cury not in grid[curx]:
                grid[curx][cury] = {}
                grid[curx][cury]["steps"] = step
    if direction == 'D':
        for i in range (1,int(move)+1):
            step = step + 1
            cury = cury-1
            if cury not in grid[curx]:
                grid[curx][cury] = {}
                grid[curx][cury]["steps"] = step

curx = 0
cury = 0
step = 0
for inst in wire2:
    direction = inst[0:1]
    move = inst[1:]
    if direction == 'R':
        for i in range (1,int(move)+1):
            step = step + 1
            curx = curx+1
            if curx in grid and cury in grid[curx] and "steps" in grid[curx][cury] and "x" not in grid[curx][cury]:
                grid[curx][cury]["x"] = 2
                grid[curx][cury]["steps"] = grid[curx][cury]["steps"] + step 
    if direction == 'L':
        for i in range (1,int(move)+1):
            step = step + 1
            curx = curx-1
            if curx in grid and cury in grid[curx] and "steps" in grid[curx][cury] and "x" not in grid[curx][cury]:
                grid[curx][cury]["x"] = 2
                grid[curx][cury]["steps"] = grid[curx][cury]["steps"] + step 
    if direction == 'U':
        for i in range (1,int(move)+1):
            step = step + 1
            cury = cury+1
            if curx in grid and cury in grid[curx] and "steps" in grid[curx][cury] and "x" not in grid[curx][cury]:
                grid[curx][cury]["x"] = 2
                grid[curx][cury]["steps"] = grid[curx][cury]["steps"] + step 
    if direction == 'D':
        for i in range (1,int(move)+1):
            step = step + 1
            cury = cury-1
            if curx in grid and cury in grid[curx] and "steps" in grid[curx][cury] and "x" not in grid[curx][cury]:
                grid[curx][cury]["x"] = 2
                grid[curx][cury]["steps"] = grid[curx][cury]["steps"] + step 

minVal = 10000000
for i in grid:
    for j in grid[i]:
        if j in grid[i] and "x" in grid[i][j]:
            if grid[i][j]["steps"] < minVal:
                minVal = grid[i][j]["steps"]

print "s2 = " + str(minVal)

