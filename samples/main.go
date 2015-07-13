package main

import (
	"github.com/thoratou/go-cocos2d-js-demos/samples/ActionsTests"
	"github.com/thoratou/go-cocos2d-js-demos/samples/test"
	"github.com/thoratou/go-cocos2d-js/cc"
)

//resources

var (
	resources = map[string]interface{}{
		"HelloWorld_png":    "res/HelloWorld.png",
		"CloseNormal_png":   "res/CloseNormal.png",
		"CloseSelected_png": "res/CloseSelected.png",
	}
	testNames = []test.Test{
		test.Test{
			Title:    "Actions Test",
			Resource: map[string]interface{}{},
			TestCase: func(controller *test.TestController) *test.TestScene {
				return ActionsTests.NewScene(controller)
			},
		},
	}
)

// main

func main() {
	cc.Game.SetOnStart(func() {
		cc.Log("starts from main.go", "begin tests")
		cc.View().AdjustViewPort(true)
		cc.View().SetDesignResolutionSize(800, 450, cc.SHOW_ALL)
		cc.View().ResizeWithBrowserSize(true)

		cc.LoaderScene().Preload(resources, func() {
			scene := cc.NewScene()
			scene.AddChild(test.NewTestController(testNames, resources).Layer)
			cc.Director().RunScene(scene)
		})
	})

	cc.Game.Run()
}
