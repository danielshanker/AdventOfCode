file = open("input.txt", "r")
total = 0

for line in file:
        number = line.split("\t")
        s = number[0]
        b = number[0]
        for i in number:
                if int(i) < int(s):
                        s = i
                if int(i) > int(b):
                        b = i
        checksum = int(b) - int(s)
        total += checksum
print "s1: " + str(total)

total = 0
file = open("input.txt", "r")

for line in file:
        found = 0
        number = line.split("\t")
        for i in number:
                for j in number:                        
                        if i != j and int(i) % int(j) == 0:
                                total += int(i)/int(j)
                                found = 1
                                break
                if found == 1:
                        break
print "s2: " + str(total)
