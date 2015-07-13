package test

type Test struct {
	Title    string
	Resource map[string]interface{}
	TestCase func(controller *TestController) *TestScene
}
