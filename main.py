from dataclasses import dataclass

@dataclass
class CruiseData:
    galsPerHr: float
    groundspeed: float
    distance: float

@dataclass
class FlightData:
    hrsETE: float
    gals: float


def getFlightData(input: CruiseData) -> FlightData:
    hrs = input.distance / input.groundspeed
    fuelBurn = hrs * input.galsPerHr
    return FlightData(hrsETE = hrs, gals = fuelBurn)

def promptCruiseData() -> FlightData:
    gph = float(input("Enter fuel burn rate (gal/hr): "))
    gs = float(input("Enter planned groundspeed: "))
    dist = float(input("Enter leg length (nautical miles): "))
    flightData = getFlightData(CruiseData(gph, gs, dist))

    return flightData


def main():
    cruise = promptCruiseData()
    print(f"Estimated time enroute: {cruise.hrsETE:.2f}hrs")
    print(f"Fuel consumed enroute: {cruise.gals:.1f}gals")


if __name__ == "__main__":
    main()

