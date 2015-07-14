package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/thoratou/go-cocos2d-js/cc"
)

//resources

var (
	resouces = map[string]interface{}{
		"HelloWorld_png":    "res/HelloWorld.png",
		"CloseNormal_png":   "res/CloseNormal.png",
		"CloseSelected_png": "res/CloseSelected.png",
	}
)

// HelloworldScene

type helloworldLayer struct {
	cc.Layer
	sprite cc.Sprite
}

func NewHelloWorldLayer() cc.Layer {
	layer := &helloworldLayer{Layer: cc.NewLayer(), sprite: nil}

	size := cc.WinSize()

	closeNormal := resouces["CloseNormal_png"].(string)
	closeSelected := resouces["CloseSelected_png"].(string)
	closeCallback := func(_ *js.Object) {
		cc.Log("Menu is clicked!")
	}

	closeItem := cc.NewMenuItemImageAllArgs(
		&closeNormal,
		&closeSelected,
		nil,
		&closeCallback,
		layer.Layer)

	closeItem.Attr(map[string]interface{}{
		"x":       size.Width() - 20,
		"y":       20,
		"anchorX": 0.5,
		"anchorY": 0.5,
	})

	menu := cc.NewMenu(closeItem)
	//menu.SetPosition(cc.NewPoint(0, 0)) works too
	menu.SetPositionX(0)
	menu.SetPositionY(0)
	layer.AddChildWithOrder(menu, 1)

	//helloLabel := js.Global.Get("cc").Get("LabelTTF").New("Hello World", "Arial", 38)
	helloLabel := cc.NewLabelTTF("Hello World", "Arial", 38)

	// position the label on the center of the screen
	helloLabel.SetPosition(cc.NewPoint(size.Width()/2, 0))

	// add the label as a child to this layer
	layer.AddChildWithOrder(helloLabel, 5)

	// add "HelloWorld" splash screen"
	layer.sprite = cc.NewSprite(resouces["HelloWorld_png"].(string))
	layer.sprite.Attr(map[string]interface{}{
		"x":        size.Width() / 2,
		"y":        size.Height() / 2,
		"scale":    0.5,
		"rotation": 180,
	})
	layer.AddChildWithOrder(layer.sprite, 0)

	layer.sprite.RunAction(
		cc.NewSequence(
			cc.NewRotateTo(2, 0, 0),
			cc.NewScaleTo(2, 1, 1),
		),
	)

	helloLabel.RunAction(
		cc.NewSpawn(
			cc.NewMoveBy(2.5, cc.NewPoint(0, size.Height()-40)),
			cc.NewTintTo(2.5, 255, 125, 0),
		),
	)

	return layer.Layer
}

// HelloworldScene

func NewHelloWorldScene() cc.Scene {
	scene := cc.NewScene()
	scene.SetOnEnter(func() {
		scene.OnEnterSuper()
		layer := NewHelloWorldLayer()
		scene.AddChild(layer)
	})
	return scene
}

// main

func main() {
	cc.Game.SetOnStart(func() {
		cc.Log("starts from main.go")
		cc.View().AdjustViewPort(true)
		cc.View().SetDesignResolutionSize(800, 450, cc.SHOW_ALL)
		cc.View().ResizeWithBrowserSize(true)

		cc.LoaderScene().Preload(resouces, func() {
			cc.Director().RunScene(NewHelloWorldScene())
		})
	})

	cc.Game.Run()
}
