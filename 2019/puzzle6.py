import re

def findChildren(forest, count):
    global totalCount 
    print forest["id"] + str(count)
    totalCount = totalCount + count
    if 'children' not in forest:
        return count
    for i in forest["children"]:
        findChildren(i,count+1)
    return count

def findRoots(forest, roots):
    newRoots = [ x for x in roots ]
    if forest["id"] == "SAN":
        global santaRoots
        santaRoots = newRoots
    if forest["id"] == "YOU":
        global youRoots
        youRoots = newRoots
    if 'children' not in forest:
        return
    newRoots.append(forest["id"])
    for i in forest["children"]:
        findRoots(i, newRoots)
    return


file = open("input6.txt", "r")
i = 0
orbits = ["COM)COM"]
for line in file:
    orbits.append(line.rstrip())

#orbits = ["COM)COM", "COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L", "K)YOU", "I)SAN)"]
nodes = {}
nodes["COM"] = { "id" : "COM" }

for i in orbits:
    values = i.split(")")
    nodes[values[1]] = { "id" : values[1] }

forest = []

for i in orbits:
    values = i.split(")")
    a = values[0]
    b = values[1]
    node = nodes[b]

    if b == a:
        forest.append(node)
    else:
        parent = nodes[a]
        if not 'children' in parent:
            parent['children'] = []
        children = parent['children']
        children.append(node)

totalCount = 0
#count = findChildren(forest[0], 0)
print "s1 = " + str(totalCount)


totalCount = 0
santaRoots = []
youRoots = []

findRoots(forest[0], [])
print santaRoots
print youRoots

steps = 0

for i in youRoots:
    if i not in santaRoots:
        steps = steps + 1

for i in santaRoots:
    if i not in youRoots:
        steps = steps + 1

print "s2 = " + str(steps)
