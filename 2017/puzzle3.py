input = 312051
square = 1
steps = 1
x = 0
y = 0

while square <= input:
    if square == input:
        break
    # right
    for i in range (0, steps):
        square += 1
        x += 1
        if square == input:
            break
    if square == input:
        break

    # up
    if square == input:
        break
    for i in range (0, steps):
        square += 1
        y += 1
        if square == input:
            break
    steps += 1
    # left
    if square == input:
        break
    for i in range (0, steps):
        square += 1
        x -= 1
        if square == input:
            break

    # down
    if square == input:
        break
    for i in range (0, steps):
        square += 1
        y -= 1
        if square == input:
            break
    steps += 1
print ("solution 1: " + str(abs(x) + abs(y)))

def checkAdjacent(grid, x, y):
    squareVal = 0
    squareVal += grid[x+1][y]
    squareVal += grid[x-1][y]
    squareVal += grid[x+1][y+1]
    squareVal += grid[x+1][y-1]
    squareVal += grid[x-1][y-1]
    squareVal += grid[x-1][y+1]
    squareVal += grid[x][y+1]
    squareVal += grid[x][y-1]
    return squareVal

input = 312051
square = 1
steps = 1
gridSize = 7
x = gridSize
y = gridSize
grid = [[0 for j in range(gridSize*2)] for k in range(gridSize*2)]

grid[x][y] = 1

while square <= input:
    if square > input:
        break
    # right
    for i in range (0, steps):
        x += 1
        square = checkAdjacent(grid, x, y)
        grid[x][y] = square
        if square > input:
            break
    if square == input:
        break

    # up
    if square > input:
        break
    for i in range (0, steps):
        y += 1
        square = checkAdjacent(grid, x, y)
        grid[x][y] = square
        if square > input:
            break
    steps += 1
    # left
    if square > input:
        break
    for i in range (0, steps):
        x -= 1
        square = checkAdjacent(grid, x, y)
        grid[x][y] = square
        if square > input:
            break

    # down
    if square > input:
        break
    for i in range (0, steps):
        y -= 1
        square = checkAdjacent(grid, x, y)
        grid[x][y] = square
        if square > input:
            break
    steps += 1
    
print ("solution 2: " + str(square))
