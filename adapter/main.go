package main

// 适配器是一种结构型设计模式， 它能使不兼容的对象能够相互合作。
// 适配器可担任两个对象间的封装器， 它会接收对于一个对象的调用， 并将其转换为另一个对象可识别的格式和接口。

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows{}
	// windowsMachineAdapter实现了insertIntoLightningPort，在其中调用windows的USB
	// windows没有实现insertIntoLightningPort，实现的是insertIntoUSBPort
	// 有点多此一举的感觉，因为代码较简单，没有发挥出来吧
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
