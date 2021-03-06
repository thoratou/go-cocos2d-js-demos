package test

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/thoratou/go-cocos2d-js/cc"
)

const (
	lineSpace = 40
)

//test controller
type TestController struct {
	Layer                   cc.LayerGradient
	itemMenu                cc.Menu
	beginPos                int
	isMouseDown             bool
	autoTestEnabled         bool
	autoTestCurrentTestName string
	curPos                  cc.Point
	yOffset                 int
	testNames               []Test
	resources               map[string]interface{}
}

func NewTestController(testNames []Test, resources map[string]interface{}) *TestController {
	testController := &TestController{
		Layer: cc.NewLayerGradient(
			cc.NewColor(0, 0, 0, 255),
			cc.NewColor(0x46, 0x82, 0xB4, 255),
		),
		itemMenu:                nil,
		beginPos:                0,
		isMouseDown:             false,
		autoTestEnabled:         false,
		autoTestCurrentTestName: "N/A",
		curPos:                  cc.NewPoint(0, 0),
		yOffset:                 0,
		testNames:               testNames,
		resources:               resources,
	}

	//OnEnter
	testController.Layer.SetOnEnter(func() {
		testController.Layer.OnEnterSuper()
		testController.itemMenu.SetPositionX(testController.yOffset)
	})

	//globals
	director := cc.Director()
	winSize := director.GetWinSize()

	//add close menu
	closeItem := cc.NewMenuItemImage(
		resources["CloseNormal_png"].(string),
		resources["CloseSelected_png"].(string),
		func() {
			location := js.Global.Get("location")
			if location != nil {
				location.Call("replace", "http://www.google.com")
			} else {
				cc.Log("No location found, cannot close")
			}
		},
		testController.Layer)
	closeItem.SetPositionX(winSize.Width() - 30)
	closeItem.SetPositionY(winSize.Height() - 30)

	subItem1 := cc.NewMenuItemFontWithString("Automated Test: Off")
	subItem1.SetFontSize(18)
	subItem2 := cc.NewMenuItemFontWithString("Automated Test: On")
	subItem2.SetFontSize(18)

	toggleAutoTestItem := cc.NewMenuItemToggle(subItem1, subItem2)
	toggleAutoTestItem.SetCallback(testController.OnToggleAutoTest, testController.Layer)
	toggleAutoTestItem.SetPositionX(winSize.Width() - toggleAutoTestItem.GetWidth()/2 - 10)
	toggleAutoTestItem.SetPositionY(20)
	if testController.autoTestEnabled {
		toggleAutoTestItem.SetSelectedIndex(1)
	}

	menu := cc.NewMenu(closeItem, toggleAutoTestItem)
	menu.SetPositionX(0)
	menu.SetPositionY(0)

	// add menu items for tests
	testController.itemMenu = cc.NewMenu()

	//Add items for tests
	for i, testCase := range testNames {
		label := cc.NewLabelTTF(testCase.Title, "Arial", 24)
		menuItem := cc.NewMenuItemLabel(label, testController.OnMenuCallback, testController.Layer)
		testController.itemMenu.AddChildWithOrder(menuItem, i+10000)
		menuItem.SetPositionX(winSize.Width() / 2)
		menuItem.SetPositionY(winSize.Height() - (i+1)*lineSpace)
		menuItem.SetEnabled(true)
	}

	testController.itemMenu.SetWidth(winSize.Width())
	testController.itemMenu.SetHeight(1 * lineSpace)
	testController.itemMenu.SetPositionX(testController.curPos.X())
	testController.itemMenu.SetPositionY(testController.curPos.Y())
	testController.Layer.AddChild(testController.itemMenu)
	testController.Layer.AddChildWithOrder(menu, 1)

	return testController
}

func (t *TestController) OnMenuCallback(sender cc.Node) {
	cc.Log("called OnMenuCallback")
	t.yOffset = t.itemMenu.GetPositionY()
	idx := sender.GetLocalZOrder() - 10000

	t.autoTestCurrentTestName = t.testNames[idx].Title
	cc.Log("Load scene:", t.autoTestCurrentTestName)

	testCase := t.testNames[idx]
	res := testCase.Resource
	cc.LoaderScene().Preload(res, func() {
		scene := testCase.TestCase(t)
		if scene != nil {
			cc.Log("Run scene:", t.autoTestCurrentTestName)
			scene.RunThisTest()
		}
	})
}

func (t *TestController) OnToggleAutoTest(_ cc.Node) {
	cc.Log("called OnToggleAutoTest")
	t.autoTestEnabled = !t.autoTestEnabled
}

func (t *TestController) Clone() *TestController {
	return NewTestController(t.testNames, t.resources)
}
