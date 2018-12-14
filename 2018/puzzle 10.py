import re
import sys
import time

file = open("input10.txt", "r")
#file = open("example.txt", "r")

lights = []

for line in file:
    matchObj = re.match("position=<(.*)> velocity=<(.*)>", line)
    lightInfo = {}
    if matchObj:
        pos = matchObj.group(1)
        vel = matchObj.group(2)
        x, y = pos.split(",")
        lightInfo['x'] = int(y)
        lightInfo['y'] = int(x)
        a, b = vel.split(",")
        lightInfo['xv'] = int(b)
        lightInfo['yv'] = int(a)
        lights.append(lightInfo)
count = 0
while 1:
    count +=1
    minx = 10000000
    miny = 10000000
    maxx = -10000000
    maxy = -10000000
    for light in lights:
        light['x'] += light['xv'] 
        light['y'] += light['yv'] 
        if light['x'] < minx:
            minx = light['x']
        if light['x'] > maxx:
            maxx = light['x']
        if light['y'] < miny:
            miny = light['y']
        if light['y'] > maxy:
            maxy = light['y']
    if maxx - minx > 70:
        continue
    for a in range(minx, maxx+1):
        for b in range(miny, maxy+1):
            lightFound = 0
            for light in lights:
                if light['x'] == a and light['y'] == b:
                    lightFound = 1
            if lightFound == 1:
                sys.stdout.write('#')
            else:
                sys.stdout.write('.')
        print ''
    print ''
    print count
    time.sleep(1)



