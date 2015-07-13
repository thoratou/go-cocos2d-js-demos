package test

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/thoratou/go-cocos2d-js/cc"
)

type TestScene struct {
	Scene       cc.Scene
	RunThisTest func()
	controller  *TestController
}

func NewTestScene(controller *TestController) *TestScene {
	scene := cc.NewScene()

	testScene := &TestScene{
		Scene:      scene,
		controller: controller,
	}

	label := cc.NewLabelTTF("Main Menu", "Arial", 20)
	menuItem := cc.NewMenuItemLabel(label, testScene.onMainMenuCallback, scene)
	menuItem.SetPosition(cc.NewPoint(cc.WinSize().Width()-50, 25))

	menu := cc.NewMenu(menuItem)
	menu.SetPosition(cc.NewPoint(0, 0))

	if js.Global.Get("sideIndexBar") != js.Undefined {
		scene.AddChildWithOrder(menu, 1)
	}

	return testScene
}

func (t *TestScene) onMainMenuCallback() {
	scene := cc.NewScene()
	layer := t.controller.Layer
	scene.AddChild(layer)
	//TODO transition progress
	cc.Director().RunScene(scene)
}
