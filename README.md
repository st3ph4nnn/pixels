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
Currently tested on an i3 that is completely dogshit and there were few performance issues.  
Minimum Cores: 2  
Minimum Threads: 4  
Must have a 64-bit operating system (currently supporting only Windows 10 and 11)

# Releases

The recommended version is the latest one released in the releases tab (sidebar)  
If you want to play the latest version available (beta, bugs may happen), check  
the build folder in the repository.

# For developers

## Windows

You will need:

- Go: https://go.dev/dl/  
- MinGW w64devkit: https://github.com/skeeto/w64devkit/releases/download/v1.12.0/w64devkit-1.12.0.zip (You need to add it to path)
 
After that, download the project and extract it somewhere  
Open a terminal in the build folder.  
There are two args you **must** specify.  
1. OS (LINUX, WINDOWS) = the os you want to build to (you can cross-compile)
2. RUN (RUN, BUILD) = either RUN or BUILD the game

Use those with the **mingw32-make.exe** like this:
**mingw32-make.exe OS=WINDOWS RUN=RUN**  

If you build, it will either show up in the windows folder.

## Linux

You will need:

- Go: golang-go in apt and dnf
- g++: g++ in apt and dnf

### Ubuntu-based distributions:

**sudo apt-get install g++ golang-go libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev**

### Redhat-based distributions:

**sudo dnf install g++ golang-go mesa-libGL-devel libXi-devel libXcursor-devel libXrandr-devel libXinerama-devel**

After that, download and extract the project somewhere  
Open a terminal in the build folder.  
There are two args you **must** specify.  
1. OS (LINUX, WINDOWS) = the os you want to build to (you can cross-compile)
2. RUN (RUN, BUILD) = either RUN or BUILD the game

Use those with the **make** like this:
**make OS=LINUX RUN=RUN**  

If you build, it will show up in the linux folder.

# Many thanks

Raylib: https://www.raylib.com/  
Raylib discord: https://discord.com/invite/raylib  
Raylib reddit: https://www.reddit.com/r/raylib/
