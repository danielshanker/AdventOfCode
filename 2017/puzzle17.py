inputVal = 359
sl = [0]

curPos = 0
for i in range(1,2018):
    curPos = ((curPos + inputVal) % len(sl)) + 1
    sl.insert(curPos, i)

a = (sl.index(2017) + 1) % 2018
print ("sol 1: " + str(sl[a]))


sl = [0]
zPos = 0
sol = 0
for i in range(1,50000001):
    curPos = ((curPos + inputVal) % i) + 1
    if curPos == zPos:
        zPos += 1
    if curPos == zPos + 1:
       sol = i 
        
print ("sol 2: " + str(sol))
