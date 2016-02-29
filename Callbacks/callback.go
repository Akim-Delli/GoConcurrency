package main

import (
	"fmt"
)

func main() {
	po := new(PurchaseOrder)
	po.value = 42.27

	ch := make(chan *PurchaseOrder)

	go savePO(po, ch)

	newPo := <-ch
	fmt.Printf("PO: %v", newPo)
}

type PurchaseOrder struct {
	Number int
	value  float64
}

func savePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 124

	callback <- po
}
