import re

file = open("input3.txt", "r")
fabric = [[0 for x in range(1000)] for y in range(1000)]
overlap = 0
for line in file:
        garbage, useful = line.split("@ ")
        loc, size = useful.split(": ")
        xv, yv = loc.split(",")
        xs, ys = size.split("x")
        for x in range(int(xv), int(xv)+int(xs)):
                for y in range(int(yv), int(yv)+int(ys)):
                        fabric[x][y] += 1
                        if fabric[x][y] == 2:
                                overlap += 1


print "s1: " + str(overlap)
file = open("input3.txt", "r")
lineNo = "0"
for line in file:
        overlap = 0
        garbage, useful = line.split(" @ ")
        loc, size = useful.split(": ")
        xv, yv = loc.split(",")
        xs, ys = size.split("x")
        for x in range(int(xv), int(xv)+int(xs)):
                for y in range(int(yv), int(yv)+int(ys)):
                        if fabric[x][y] >= 2:
                                overlap = 1
                                next
                if overlap == 1:
                        next
        if overlap == 0:
                lineNo += garbage


print "s2: " + str(lineNo)
