package main

type (
	feet   int
	pounds float64
	inches float64
	knots  float64
)

type ModelsFile struct {
	Models map[string]Model `json:"models"`
}

type Model struct {
	Name       string     `json:"name"`
	Limits     Limits     `json:"limits"`
	CGEnvelope CGEnvelope `json:"cg_envelope"`
}

type Limits struct {
	MaxGW pounds `json:"maxGW"`
	Vne   knots  `json:"vne"`
	Vs0   knots  `json:"vs0"`
	Ceil  feet   `json:"ceiling"`
	Seats int    `json:"seats"`
}

type CGEnvelope struct {
	Fwd []CGLimit `json:"fwd_limits"`
	Aft []CGLimit `json:"aft_limits"`
}

type CGLimit struct {
	Weight pounds `json:"weight"`
	Arm    inches `json:"arm"`
}

type AircraftFile struct {
	Aircraft []Aircraft `json:"aircraft"`
}

type Aircraft struct {
	Reg         string `json:"reg"`
	ModelID     string `json:"model_id"`
	EmptyWeight pounds `json:"empty_weight"`
	EmptyCG     inches `json:"empty_cg"`
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
*/
