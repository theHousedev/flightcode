from dataclasses import dataclass

@dataclass
class CruiseData:
    galsPerHr: float
    gndspeed: float
    distance: float

@dataclass
class FlightData:
    hrsETE: float
    gals: float

@dataclass
class DescentData:
    offsetTOD: float    # miles before target to begin descent
    rateFPM: int        # feet/min descent rate

# Cruise flight calculations
def promptCruiseData() -> FlightData:
    gph = float(input("Enter fuel burn rate (gal/hr): "))
    gs = float(input("Enter planned groundspeed: "))
    dist = float(input("Enter leg length (nautical miles): "))
    
    hrs = dist / gs
    fuelBurn = hrs * gph
    return FlightData(hrsETE = hrs, gals = fuelBurn)

# Descent planning calculations
def promptDescentPlan() -> DescentData:
    altCurrent = float(input("Current altitude: "))
    altTarget = float(input("Target altitude: "))
    angle = float(input("Target descent angle (default 3.0°): "))

    altToDescend = altCurrent - altTarget



def main():
    # present offer for type of calculation:
    #   - cruise (fuel burn/ETE)
    #   - performance (TO/LDG dist, climb rate)
    #   - '3:1 rule' descent planning
    #   - gross weight speed recalculation
    #   - weight shift formula?

    cruise = promptCruiseData()
    print(f"Estimated time enroute: {cruise.hrsETE:.2f}hrs")
    print(f"Fuel consumed enroute: {cruise.gals:.1f}gals")


if __name__ == "__main__":
    main()

