file = open("input.txt", "r")
wall = {}

maxPos = 0
for line in file:
    line = line.rstrip()
    (pos, depth) = line.split(": ")
    wall[int(pos)] = int(depth)
    if pos > maxPos:
        maxPos = int(pos)

hit = []
for i in range (0, maxPos+1):
    if i in wall:
        wallDepth = wall[i]
        wallPos = i % (2*(wallDepth - 1))
        if wallPos == 0:
            hit.append(i)
severity = 0
for i in hit:
    severity += i * wall[i]
print "sol 1: " + str(severity)

j = 9
while (1):
    j += 1
    beenHit = 0
    for i in range (0, maxPos+1):
        if i in wall:
            wallDepth = wall[i]
            wallPos = (i + j) % (2*(wallDepth - 1))
            if wallPos == 0:
                beenHit = 1
                break
    if beenHit == 0:
        print "sol 2: " + str(j)
        break
