import re

file = open("input4.txt", "r")

sched = []
for line in file:
        sched.append(line)
sched.sort()

example = "[1518-11-01 00:00] Guard #10 begins shift\n[1518-11-01 00:05] falls asleep\n[1518-11-01 00:25] wakes up\n[1518-11-01 00:30] falls asleep\n[1518-11-01 00:55] wakes up\n[1518-11-01 23:58] Guard #99 begins shift\n[1518-11-02 00:40] falls asleep\n[1518-11-02 00:50] wakes up\n[1518-11-03 00:05] Guard #10 begins shift\n[1518-11-03 00:24] falls asleep\n[1518-11-03 00:29] wakes up\n[1518-11-04 00:02] Guard #99 begins shift\n[1518-11-04 00:36] falls asleep\n[1518-11-04 00:46] wakes up\n[1518-11-05 00:03] Guard #99 begins shift\n[1518-11-05 00:45] falls asleep\n[1518-11-05 00:55] wakes up"
#sched = example.split("\n")

curGuard = -1
curStartTime = -1
sleepTime = {}
sleepMins = {}
for line in sched:
        time, action = line.split("] ")
        garb, minute = time.split(":")
        matchObj = re.match("Guard #(.*) begins", action)
        if matchObj:
                curGuard = int(matchObj.group(1))

        matchObj = re.match("falls asleep", action)
        if matchObj:
                curStartTime = int(minute)
                
        matchObj = re.match("wakes up", action)
        if matchObj:
                if curGuard in sleepTime:
                        sleepTime[curGuard] += (int(minute) - curStartTime)
                else:
                        sleepTime[curGuard] = (int(minute) - curStartTime)

                for i in range(int(curStartTime), int(minute)):
                        key = str(curGuard) + ":" + str(i)
                        if key in sleepMins:                                
                                sleepMins[key] += 1
                        else:
                                sleepMins[key] = 1

mostMins = 0
guard = -1
for i in sleepTime:
        if sleepTime[i] > mostMins:
                guard = i
                mostMins = sleepTime[i]

maxMin = 0
curMaxMin = 0
for i in sleepMins:
        g, m = i.split(":")
        if int(g) != guard:
                continue
        if sleepMins[i] > maxMin:
                curMaxMin = m
                maxMin = sleepMins[i]

answer = int(curMaxMin) * int(guard)
print "s1: " + str(answer)

import re

file = open("input4.txt", "r")

sched = []
for line in file:
        sched.append(line)
sched.sort()

example = "[1518-11-01 00:00] Guard #10 begins shift\n[1518-11-01 00:05] falls asleep\n[1518-11-01 00:25] wakes up\n[1518-11-01 00:30] falls asleep\n[1518-11-01 00:55] wakes up\n[1518-11-01 23:58] Guard #99 begins shift\n[1518-11-02 00:40] falls asleep\n[1518-11-02 00:50] wakes up\n[1518-11-03 00:05] Guard #10 begins shift\n[1518-11-03 00:24] falls asleep\n[1518-11-03 00:29] wakes up\n[1518-11-04 00:02] Guard #99 begins shift\n[1518-11-04 00:36] falls asleep\n[1518-11-04 00:46] wakes up\n[1518-11-05 00:03] Guard #99 begins shift\n[1518-11-05 00:45] falls asleep\n[1518-11-05 00:55] wakes up"
#sched = example.split("\n")

curGuard = -1
curStartTime = -1
sleepTime = {}
sleepMins = {}
for line in sched:
        time, action = line.split("] ")
        garb, minute = time.split(":")
        matchObj = re.match("Guard #(.*) begins", action)
        if matchObj:
                curGuard = int(matchObj.group(1))

        matchObj = re.match("falls asleep", action)
        if matchObj:
                curStartTime = int(minute)
                
        matchObj = re.match("wakes up", action)
        if matchObj:
                if curGuard in sleepTime:
                        sleepTime[curGuard] += (int(minute) - curStartTime)
                else:
                        sleepTime[curGuard] = (int(minute) - curStartTime)

                for i in range(int(curStartTime), int(minute)):
                        key = str(curGuard) + ":" + str(i)
                        if key in sleepMins:                                
                                sleepMins[key] += 1
                        else:
                                sleepMins[key] = 1

maxMin = 0
curMaxMin = 0
maxGuard = 0
for i in sleepMins:
        g, m = i.split(":")

        if sleepMins[i] > maxMin:
                curMaxMin = m
                maxMin = sleepMins[i]
                maxGuard = g

answer = int(curMaxMin) * int(maxGuard)
print "s2: " + str(answer)
