package main

import (
	"fmt"
)

type (
	feet   int
	pounds float64
	inches float64
	knots  float64
)

type Aircraft struct {
	Reg         string    `json:"registration"`
	Model       string    `json:"model"`
	Limitations LimSheet  `json:"limits"`
	LoadData    LoadSheet `json:"load_data"`
}

type LoadSheet struct {
	MaxGW      pounds   `json:"max_gw"`
	EmptyLbs   pounds   `json:"empty_weight"`
	EmptyCG    inches   `json:"empty_cg"`
	CGEnvelope CGLimits `json:"cg_envelope"`
}

type CGLimits struct {
	FwdLimits []CGLimit `json:"fwd_cg_limits"`
	AftLimits []CGLimit `json:"aft_cg_limits"`
}

type CGLimit struct {
	Weight pounds `json:"weight"`
	Arm    inches `json:"arm"`
}

type LimSheet struct {
	MaxGW pounds `json:"max_gw"`
	Vne   knots  `json:"vne"`
	Vs0   knots  `json:"vs0"`
	Ceil  feet   `json:"ceiling"`
	Seats int    `json:"seats"`
}

/*
type Perf struct {
	distTO    feet
	distLDG   feet
	climbRate feet
	cruise    struct {
		TAS       knots
		burn      gallons
		endurance time.Duration
	}
}

type Leg struct {
	points   [2]string
	length   nauticalMi
	burn     gallons
	duration time.Duration
}

type FlightPlan struct {
	legs []Leg
}

func fmtUnit(value any) string {
	switch value.(type) {
	case feet:
		return fmt.Sprintf("%.0fft", value)
	case pounds:
		return fmt.Sprintf("%.1flbs", value)
	case inches:
		return fmt.Sprintf("%.1fin", value)
	case gallons:
		return fmt.Sprintf("%.1fgal", value)
	case knots:
		return fmt.Sprintf("%.0fkts", value)
	case nauticalMi:
		return fmt.Sprintf("%.1fnm", value)
	}
	return fmt.Sprintf("%.1f", value)
}
*/

func main() {
	fmt.Println("Flight Computer")

}
