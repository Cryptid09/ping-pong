package main

import(
  "log"
  "github.com/hajimehoshi/ebiten/v2"
)

const (
  screenWidth =  640
  screenHeight = 480
  ballSpeed = 5
  paddleSpeed = 6
)

type Object struct {
  X, Y, W, H int
}

type Paddle struct{
  Object
}

type Ball struct {
  Object 
  dxdt int
  dydt int
}

type Game struct {
  paddle Paddle
  ball Ball
  highScore int
}

func main (){
  ebiten.SetWindowTitle("Pong in Ebitengine")  
  ebiten.SetWindowSize(screenWidth, screenHeight)

  g := &Game{}
 
  err := ebiten.RunGame(g)
  if err != nil{
    log.Fatal(err)
  }
}

func (g *Game) Layout(outsideWidth, outsideHeight int)(int,int){
  return screenWidth , screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Update() error{
  return nil
}

