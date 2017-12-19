import re

file = open("input.txt", "r")
maze = []
for line in file:
    maze.append(line.rstrip('\n'))

x = 0
y = 0
direction = 'd'
for i in range (0, len(maze[0])):
    if maze[0][i] != ' ':
        x = i

letters = ''
steps = 0
while (1):
    steps += 1
    if direction == 'd':
        y += 1
        if maze[y][x] == '|':
            continue
        elif maze[y][x] == '+':
            if maze[y][x+1] != ' ' and maze[y][x+1] != '|':
                direction = 'r'
            else:
                direction = 'l'
        elif maze[y][x] == '-':
            continue
        elif maze[y][x] == ' ':
            break
        else:
            letters += maze[y][x]

    elif direction == 'u':
        y -= 1
        if maze[y][x] == '|':
            continue
        elif maze[y][x] == '+':
            if maze[y][x+1] != ' ' and maze[y][x+1] != '|':
                direction = 'r'
            else:
                direction = 'l'
        elif maze[y][x] == '-':
            continue
        elif maze[y][x] == ' ':
            break
        else:
            letters += maze[y][x]



    elif direction == 'r':
        x += 1
        if maze[y][x] == '-':
            continue
        elif maze[y][x] == '+':
            if maze[y+1][x] != ' ' and maze[y+1][x] != '-':
                direction = 'd'
            else:
                direction = 'u'
        elif maze[y][x] == '|':
            continue
        elif maze[y][x] == ' ':
            break
        else:
            letters += maze[y][x]


    elif direction == 'l':
        x -= 1
        if maze[y][x] == '-':
            continue
        elif maze[y][x] == '+':
            if maze[y+1][x] != ' ' and maze[y+1][x] != '-':
                direction = 'd'
            else:
                direction = 'u'
        elif maze[y][x] == '|':
            continue
        elif maze[y][x] == ' ':
            break
        else:
            letters += maze[y][x]
print "sol 1: " + letters
print "sol 2: " + str(steps)
