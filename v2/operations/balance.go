package operations


// Example:
//
//	charge, update := &omise.Balance{}, &RetrieveBalance{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveBalance struct {
	Base
}

func (req *RetrieveBalance) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/balance",
		ContentType: "application/json",
	}
}

