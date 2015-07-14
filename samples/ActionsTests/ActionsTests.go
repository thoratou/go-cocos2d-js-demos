package ActionsTests

import (
	"github.com/thoratou/go-cocos2d-js-demos/samples/test"
	"github.com/thoratou/go-cocos2d-js/cc"
)

func NewScene(controller *test.TestController) *test.TestScene {
	ts := test.NewTestScene(controller)
	ts.RunThisTest = func() {
		cc.Director().RunScene(ts.Scene)
	}
	return ts
}
