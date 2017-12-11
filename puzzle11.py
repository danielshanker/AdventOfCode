file = open("input.txt", "r")
index = 0
step = 0
maze = {}

for line in file:
    line = line.rstrip()
    instr = line.split(',')

x = 0
y = 0
z = 0
maxD = 0

for dir in instr:
    if dir == 'n':
        y += 1
        z -= 1
    if dir == 's':
        y -= 1
        z += 1
    if dir == 'nw':
        y += 1
        x -= 1
    if dir == 'se':
        y -= 1
        x += 1
    if dir == 'ne':
        x += 1
        z -= 1
    if dir == 'sw':
        x -= 1
        z += 1
    distance = (abs(x) + abs(y) + abs (z))/2
    if distance > maxD:
        maxD = distance

distance = (abs(x) + abs(y) + abs (z))/2

print "x - " + str(x) + " y - " + str(y) + " z - " + str(z)
print distance
print maxD
