file = open("input.txt", "r")
inst = []
connection = [0]

for line in file:
    inst.append(line.rstrip())

for i in range(0, 100):
    for j in range (0, len(inst)):
        node = inst[j].split (" <-> ")
        end = node[1].split(',')
        if int(node[0]) in connection:
            for a in end:
                if int(a) not in connection:
                    connection.append(int(a))

print "sol 1: " + str(len(connection))

count = 0
while (1):
    rList = []
    count += 1
    node = inst[0].split (" <-> ")
    end = node[1].split(',')
    connection = [int(node[0])]
    for i in range(0, 100):
        for j in range (0, len(inst)):
            node = inst[j].split (" <-> ")
            end = node[1].split(',')
            remove = 0
            if int(node[0]) in connection:
                if inst[j] not in rList:
                    rList.append(inst[j])
                for a in end:
                    if int(a) not in connection:
                        connection.append(int(a))
    for a in rList:
        inst.remove(a)
    if len(inst) == 0:
        break
print count
