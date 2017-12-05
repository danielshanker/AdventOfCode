file = open("input.txt", "r")
index = 0
step = 0

maze = []

for line in file:
    maze.append(int(line))

while True:
    jumpVal = maze[index]
    maze[index] += 1
    index = index + jumpVal
    step += 1
    if index >= len(maze) or index < 0:
        break

print "s1: " + str(step)

file = open("input.txt", "r")
index = 0
step = 0

maze = []

for line in file:
    maze.append(int(line))

while True:
    jumpVal = maze[index]
    if jumpVal < 3:
        maze[index] += 1
    else:
        maze[index] -= 1        
    index = index + jumpVal
    step += 1
    if index >= len(maze) or index < 0:
        break

print "s2: " + str(step)

