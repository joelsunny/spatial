import pygame
import pygame.freetype
import random
from pprint import pprint

'''
How to force a breadth first drawing without an explicit graph to traverse.
recursion is naturally dfs, inorder traversal.
need level order.

use map instead of single list

'''

# function to translate between linear order and 2d cordinates
def d2xy(d, size):
    pass

points = {}
def hilbert(x=256, y=256, size=512, l=0):
    if (size == 4):
        return
    else:
        if l not in points.keys():
            points[l] = []
        points[l] += [(x - size//2, y - size//2), (x + size//2, y - size//2), (x + size//2, y + size//2), (x - size//2, y + size//2)]
        hilbert(x - size//2, y - size//2, size//2, l+1)
        hilbert(x + size//2, y - size//2, size//2, l+1)
        hilbert(x + size//2, y + size//2, size//2, l+1)
        hilbert(x - size//2, y + size//2, size//2, l+1)


    
def main():
    global width, height, space, writer, g, clock, fall_time
    width  = 512
    height = 512
    n = 6  

    win = pygame.display.set_mode((512, 512))
    pygame.init()

    done = False
    win.fill((0,0,0))
    clock = pygame.time.Clock()
    print(points.keys())
    fi = points[6]
    for i,p in enumerate(fi):
        if i < len(fi)-1:
            pygame.draw.line(win,(255,124,123), p, fi[i+1], 1)
    pygame.display.update()
            
    
    while not done:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                pygame.quit()
                exit(0)

if __name__ == '__main__':
    hilbert()
    main()