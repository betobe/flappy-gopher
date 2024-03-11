# Flappy gopher

A game created with the Go language + SDL library

Game made by following [JustForFunc](https://www.youtube.com/@JustForFunc) tutorial to get an overview of the language capabilities

## Requirements
#### OSX:
```
brew install sdl2
brew install sdl2_image
brew install sdl2_ttf
brew install sdl2_mixer
```

## How to run it
1. Install dependencies with ``` make dependencies ```
2. Build the application with ``` make build ```
3. Run it! ``` ./dist/main ```

## Troubleshooting
At an early stage of develpment you may find that the game window doesn't show up but no error is triggered. You can find a hack to temporary avoid that problem in the first commit (main.go#L44) of this repository.
