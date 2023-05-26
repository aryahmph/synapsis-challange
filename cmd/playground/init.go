package main

import (
	"fmt"
	xenditCommon "synapsis-challange/common/xendit"
)

func main() {
	xenditManager := xenditCommon.NewXenditManager("xnd_development_ye78i7MafcV7CF1UPhdFS6m4NjfzXfqtYjlQ9KDiJ3fNMZXG3AFdYfStgtbpZ")
	payment, err := xenditManager.NewVirtualAccountPayment("va-example", 50_000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(payment)
}
