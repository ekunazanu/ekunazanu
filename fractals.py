#!/usr/bin/env python3

import numpy as np
import matplotlib.pyplot as plt

# Julia Set
xDomain, yDomain = np.linspace(-2, 2, 2000), np.linspace(-2, 2, 2000)
max_iterations = 1000
colormap = 'gray_r'
c = complex(-2, 0)
# Mandelbrot Set
# xDomain, yDomain = np.linspace(-0.914531120987654315192, -0.330394552098765441863, 3000), np.linspace(0.315550324938271605116, 0.765528888888888888892, 2000)
# max_iterations = 100
# colormap = 'pink_r'
# z = complex(0, 0)

bound = 2
power = 2

iterationArray = []
for y in yDomain:
    row = []
    for x in xDomain:
        z = complex(x,y)     # julia
        # c = complex(x, y)    # mandelbrot
        # z = complex(0, 0)    # mandelbrot
        for iterationNumber in range(max_iterations):
            if(abs(z) >= bound):
                row.append(iterationNumber)
                break
            else: z = z**power + c
        else:
            row.append(max_iterations)
    iterationArray.append(row)

ax = plt.axes()
ax.set_aspect('equal')
plt.axis('off')
plt.pcolormesh(xDomain, yDomain, iterationArray, cmap=colormap)
plt.savefig('fractal.png', bbox_inches="tight", dpi=400)
