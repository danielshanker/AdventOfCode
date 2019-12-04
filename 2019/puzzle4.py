import re

inp = "246540-787419"
start = 246540
end = 787420
passwords = 0

for code in range (start, end):
    codeStr = str(code)
    double = 0
    for i in range(5):
        if int(codeStr[i]) > int(codeStr[i+1]):
            double = 0
            break
        if int(codeStr[i]) == int(codeStr[i+1]):
            double = 1
    if double == 1:
        passwords = passwords + 1


print "s1 = " + str(passwords)

passwords = 0
for code in range (start, end):
    codeStr = str(code)
    badFlag = 0
    doubleDigit = []
    doubleDigit = [ 0 for x in range(10) ]
    for i in range(5):
        if int(codeStr[i]) > int(codeStr[i+1]):
            badFlag = 1
            break
        if int(codeStr[i]) == int(codeStr[i+1]):
            doubleDigit[int(codeStr[i])] = doubleDigit[int(codeStr[i])] + 1

    for i in range(10):
        if doubleDigit[i] == 1 and badFlag == 0:
            passwords = passwords + 1
            break


print "s2 = " + str(passwords)
