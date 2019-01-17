package lengthconv

import "fmt"

type Meter float64
type CentiMeter float64

func (m Meter) String() string       { return fmt.Sprintf("%gM", m) }
func (cm CentiMeter) String() string { return fmt.Sprintf("%gCM", cm) }

func CMToM(cm CentiMeter) Meter { return Meter(cm / 100) }
func MToCM(cm Meter) CentiMeter { return CentiMeter(cm * 100) }
