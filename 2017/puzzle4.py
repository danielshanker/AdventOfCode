file = open("input.txt", "r").read().split('\n')
total = 0

for line in file:
    if line == "":
        continue 
    words = line.split(" ")
    words.sort()
    index = 0
    for a in words:
        if index == len(words) - 1:
            total += 1
        else:
            if words[index] == words[index + 1]:
                break
        index += 1
print "s1: " + str(total)

file = open("input.txt", "r").read().split('\n')
total = 0

for line in file:
    if line == "":
        continue 
    words = line.split(" ")
    index = 0
    wordsAlpha = []
    for b in words:
        wordsAlpha.append(''.join(sorted(b)))
        index += 1
    wordsAlpha.sort()
    index = 0
    for a in wordsAlpha:
        if index == len(wordsAlpha) - 1:
            total += 1
        else:
            if wordsAlpha[index] == wordsAlpha[index + 1]:
                break
        index += 1
print "s1: " + str(total)

