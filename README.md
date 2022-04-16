# Pixels

Pixels is a game about pixel simulation made using Go and  
the Go Binding of Raylib 4.0;  

![image](https://user-images.githubusercontent.com/67066397/163684623-6ec27745-795c-45df-b6e0-13a977666308.png)
![image](https://user-images.githubusercontent.com/67066397/163684641-25f7dcdc-1527-4589-9d87-adacf40b513f.png)
![image](https://user-images.githubusercontent.com/67066397/163684666-5726fd25-e7b5-449a-adff-fea8a2ee8086.png)


# How the game works

You have a canvas. You drop one pixel using the right mouse button,  
or multiple holding the left mouse button. You can change colors using  
the scroll wheel, and clear the screen using the delete key.  
Each color hunts nearest enemies that have other colors. If a pixel is not being  
fed (eaten some other pixels) in 10 seconds it will die.

To get all the keybinds listed, Press X.  

In the top right corner, there is listed, in order,  
FPS (capped at 244), the numbers of pixels (all colors)  
and a rectangle colored in the current pixel color.  

# READ THIS

Even though the game is multi-threaded and somewhat optimized,  
the game still needs a decent CPU to run without issues.   
Currently tested on an i3 that is completely dogshit  
and there is few performance issues.  
Minimum Cores: 2  
Minimum Threads: 4  
Must have a 64-bit operating system (currently supporting only Windows 10 and 11)

# For developers

To build this, just install Go (1.18 or later)  
go into the game directory, open a terminal,
and type in .\run.bat to run or .\build.bat
to build the game in \build\ directory.

# Many thanks

Raylib: https://www.raylib.com/
Raylib discord: https://discord.com/invite/raylib
Raylib reddit: https://www.reddit.com/r/raylib/
