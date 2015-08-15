package ActionsTests

import (
	"github.com/thoratou/go-cocos2d-js-demos/samples/test"
	"github.com/thoratou/go-cocos2d-js/cc"
)

type actionMove struct {
	TestScene *test.TestScene
	Grossini  cc.Sprite
	Tamara    cc.Sprite
	Kathia    cc.Sprite
}

func (a *actionMove) centerSprites() {
	winSize := cc.Director().GetWinSize()

	a.Grossini.SetPositionX(winSize.Width() / 2)
	a.Grossini.SetPositionY(winSize.Height() / 2)
	a.Tamara.SetPositionX(winSize.Width() / 4)
	a.Tamara.SetPositionY(winSize.Height() / 2)
	a.Kathia.SetPositionX(3 * winSize.Width() / 4)
	a.Kathia.SetPositionY(winSize.Height() / 2)
}

func NewScene(controller *test.TestController) *test.TestScene {
	ts := test.NewTestScene(controller)
	am := &actionMove{
		TestScene: ts,
		Grossini:  cc.NewSprite("res/Images/grossini.png"),
		Tamara:    cc.NewSprite("res/Images/grossinis_sister1.png"),
		Kathia:    cc.NewSprite("res/Images/grossinis_sister2.png"),
	}

	am.Grossini.SetScale(0.5)
	am.Tamara.SetScale(0.5)
	am.Kathia.SetScale(0.5)

	ts.Scene.AddChildWithTag(am.Grossini, 1)
	ts.Scene.AddChildWithTag(am.Tamara, 2)
	ts.Scene.AddChildWithTag(am.Kathia, 3)

	winSize := cc.Director().GetWinSize()
	am.Grossini.SetPositionX(winSize.Width() / 2)
	am.Grossini.SetPositionY(winSize.Height() / 3)
	am.Tamara.SetPositionX(winSize.Width() / 2)
	am.Tamara.SetPositionY(2 * winSize.Height() / 3)
	am.Kathia.SetPositionX(winSize.Width() / 2)
	am.Kathia.SetPositionY(winSize.Height() / 2)

	ts.Scene.SetOnEnter(func() {
		ts.Scene.OnEnterSuper()

		cc.Log("Move test started")
		am.centerSprites()
		winSize := cc.Director().GetWinSize()

		actionTo := cc.NewMoveTo(2, cc.NewPoint(winSize.Width()-40, winSize.Height()-40))
		actionBy := cc.NewMoveBy(1, cc.NewPoint(80, 80))
		actionByBack := actionBy.Reverse()

		am.Tamara.RunAction(actionTo)
		am.Grossini.RunAction(cc.NewSequence(actionBy, actionByBack))
		am.Kathia.RunAction(cc.NewMoveTo(1, cc.NewPoint(40, 40)))
	})

	ts.RunThisTest = func() {
		cc.Director().RunScene(ts.Scene)
	}
	return ts
}
