# Go Sonic Go

Author Pablo Picouto Garcia

Is this repo useful? Please ⭑Star this repository and share the love.

![My image](img/title.png)

A Sonic engine implemented using [pixel](https://github.com/faiface/pixel) for animations and [beep](https://github.com/faiface/beep) for audio in Go.

https://user-images.githubusercontent.com/2054461/211116034-de51001f-e8f9-4282-bec2-b1da2f1584d8.mov

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





