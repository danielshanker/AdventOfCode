import re

file = open("input.txt", "r")
posX = []
velX = []
accX = []
posY = []
velY = []
accY = []
posZ = []
velZ = []
accZ = []
part = []
q = 0
for line in file:
    line = line.rstrip()
    matchObj = re.match("p=<(-?\w+),(-?\w+),(-?\w+)>, v=<(-?\w+),(-?\w+),(-?\w+)>, a=<(-?\w+),(-?\w+),(-?\w+)>", line, flags=0)
    if matchObj:
        posX.append(int(matchObj.group(1)))
        posY.append(int(matchObj.group(2)))
        posZ.append(int(matchObj.group(3)))
        velX.append(int(matchObj.group(4)))
        velY.append(int(matchObj.group(5)))
        velZ.append(int(matchObj.group(6)))
        accX.append(int(matchObj.group(7)))
        accY.append(int(matchObj.group(8)))
        accZ.append(int(matchObj.group(9)))
        part.append(q)
        q += 1


particle = 0
smallest = int(abs(accX[0])) + int(abs(accY[0])) + int(abs(accZ[0]))
for i in range(0, len(accX)):
    acceleration = int(abs(accX[i])) + int(abs(accY[i])) + int(abs(accZ[i]))
    if acceleration < smallest:
        smallest = acceleration
        particle = i
print "sol 1: " + str(particle)

file = open("input.txt", "r")
posX = []
velX = []
accX = []
posY = []
velY = []
accY = []
posZ = []
velZ = []
accZ = []
for line in file:
    line = line.rstrip()
    matchObj = re.match("p=<(-?\w+),(-?\w+),(-?\w+)>, v=<(-?\w+),(-?\w+),(-?\w+)>, a=<(-?\w+),(-?\w+),(-?\w+)>", line, flags=0)
    if matchObj:
        posX.append(int(matchObj.group(1)))
        posY.append(int(matchObj.group(2)))
        posZ.append(int(matchObj.group(3)))
        velX.append(int(matchObj.group(4)))
        velY.append(int(matchObj.group(5)))
        velZ.append(int(matchObj.group(6)))
        accX.append(int(matchObj.group(7)))
        accY.append(int(matchObj.group(8)))
        accZ.append(int(matchObj.group(9)))


#for j in range(0,1000):
for j in range(0,500):
    collided = []
    for i in range(0,len(accX)):
        velX[i] += accX[i]
        velY[i] += accY[i]
        velZ[i] += accZ[i]
    
        posX[i] += velX[i]
        posY[i] += velY[i]
        posZ[i] += velZ[i]

    for i in range(0,len(accX)):
        col = 0
        for k in range (0, i):
            if posX[i] == posX[k] and posY[i] == posY[k] and posZ[i] == posZ[k]:
                col = 1
                if k not in collided:
                    collided.append(k)
        if col == 1:
            collided.append(i)
    for i in sorted(collided, reverse=True):
        del accX[i]
        del accY[i]
        del accZ[i]
        del velX[i]
        del velY[i]
        del velZ[i]
        del posX[i]
        del posY[i]
        del posZ[i]
        del part[i]
print "sol 2: " + str(len(posX))
            
