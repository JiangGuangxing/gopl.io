package main

import (
	"fmt"
	"gopl.io/ch2/lengthconv"
)

func main() {
	//for _, arg := range os.Args[1:] {
	//	t, err := strconv.ParseFloat(arg, 64)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	//		os.Exit(1)
	//	}
	//	cm := lengthconv.CentiMeter(t)
	//	m := lengthconv.Meter(t)
	//	fmt.Printf("%s = %s, %s = %s\n",
	//		cm, lengthconv.CMToM(cm), m, lengthconv.MToCM(m))
	//}

	t := 12
	cm := lengthconv.CentiMeter(t)
	m := lengthconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		cm, lengthconv.CMToM(cm), m, lengthconv.MToCM(m))
}
