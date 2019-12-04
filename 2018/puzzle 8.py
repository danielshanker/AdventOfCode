import re
import time

file = open("input8.txt", "r")
file = open("example.txt", "r")

node = []
for line in file:
    node = line.split(' ')
node[len(node)-1] = node[len(node)-1].rstrip()

    

print "s1: " + str('')
