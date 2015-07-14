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
	testScene := &TestScene{
		Scene:      cc.NewScene(),
		controller: controller,
	}

	label := cc.NewLabelTTF("Main Menu", "Arial", 20)
	menuItem := cc.NewMenuItemLabelAllArgs(label, testScene.OnMainMenuCallback, testScene.Scene)
	menuItem.SetPosition(cc.NewPoint(cc.WinSize().Width()-50, 25))

	menu := cc.NewMenu(menuItem)
	menu.SetPosition(cc.NewPoint(0, 0))

	testScene.Scene.AddChildWithOrder(menu, 1)

	return testScene
}

func (t *TestScene) OnMainMenuCallback(_ *js.Object) {
	cc.Log("begin OnMainMenuCallback")
	scene := cc.NewScene()
	layer := t.controller.Clone().Layer
	scene.AddChild(layer)
	//TODO transition progress
	cc.Director().RunScene(scene)
}
