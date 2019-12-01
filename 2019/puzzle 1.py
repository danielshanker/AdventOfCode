import re
import math

masses = []
file = open("input1.txt", "r")
totalMass = 0
for line in file:
        masses.append(int(line))
        totalMass = totalMass + (math.floor(int(line)/3) - 2)
print "s1: " + str(totalMass)

totalFuelMass = 0
for mass in masses:
        fuelMass = math.floor(int(mass)/3) - 2
        totalFuelMass = totalFuelMass + fuelMass
        while fuelMass > 0:
                fuelMass = math.floor(int(fuelMass)/3) - 2
                if fuelMass > 0:
                        totalFuelMass = totalFuelMass + fuelMass
print "s2: " + str(totalFuelMass)
