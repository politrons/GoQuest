# Go Sonic Go

![My image](img/title.png)

A Sonic engine implemented using [pixel](https://github.com/faiface/pixel) animation and [beep](https://github.com/faiface/beep) for audio in Go.

[SonicEngine](main/SonicEngine.go) contains all the logic of animations of 

```
Run, Jump, Wait, Loop up, Look down, Spin.
```
![My image](img/stop.png)
![My image](img/wait.png)
![My image](img/jump.png)

[SonicGame](main/SonicGame.go) is the ```main``` class that create the Window, load all the [sprites](sprites/) and use 
```pixel``` lib to get all Key events to invoke the specific animation of the engine.

All the sprites have been used thanks to this author.

![My image](img/sprites.png)


