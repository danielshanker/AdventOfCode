import re
import sys

answer = 0
file = open("input6.txt", "r")

coordinates = []
for line in file:
        coordinates.append(line)
#coordinates = ("1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9")
example = 500

grid = [[0 for x in range(example)] for y in range(example)]
regionSize = 0

for i in range (0, example):
        for j in range (0,example):
                smallest = 100000
                a = 0
                regionFlag = 0
                distTotal = 0
                for line in coordinates:
                        x, y = line.split (', ')
                        dist = abs(i - int(x)) + abs(j - int(y))
                        if dist == 0:
                                grid[j][i] = a
                                smallest = dist
                        elif dist == smallest:
                                grid[j][i] = '.'
                        elif dist < smallest:
                                smallest = dist
                                grid[j][i] = a
                        distTotal += dist
                        if distTotal >= 10000:
                                regionFlag = 1
                        a += 1
                if regionFlag == 0:
                        regionSize += 1

numbers = [0 for x in range(len(coordinates))]
infinite = []
for i in range (0, example):
        for j in range (0, example):
                value = grid[j][i]
                if i == 0 or j == 0 or i == example - 1 or j == example - 1:
                        if value not in infinite:
                                infinite.append(value)
                        continue
                if value == '.':
                        continue
                numbers[value] += 1
print infinite

biggest = 0
val = 0
for i in range(len(coordinates)):
        if i in infinite:
                continue
        if numbers[i] > biggest:
                biggest = numbers[i]
                val = i
print "s1: " + str(biggest)
print "s2: " + str(regionSize)
