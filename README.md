## About
The goal of this project is to build a simple CLI tool that computes some basic flight planning info, such as:

- Weight & balance data
- Runway and climb performance metrics
- Enroute/cruise performance for long distance flying

In addition, I will rely on `aviationweather.gov/data/api` to fetch any relevant atmospheric/airport data. At the end of the day, I'd like this to be a fully functional supplement to a flight planning workflow.


## Roadmap

### Minimum viable product:
- [ ] Basic W&B Calculation (hardcoded loadsheet)
    - [ ] Using supplied default values
    - [ ] Interactive custom loading
- [ ] Basic Enroute Performance
    - [ ] ETE, ETD, ETA
    - [ ] Landing weight & CG
    - [ ] TOC/TOD planning
    - [ ] Fuel planning
- [ ] Basic Runway Performance
    - [ ] Hardcoded density table
    - [ ] Wind calculator

### Expanded Features
Interactive addition of...
- [ ] New aircraft model info:
    - [ ] Basic limitations + EW/CG
    - [ ] Custom W&B loadsheet*
- [ ] Model-specific performance tables:
    - [ ] Runway data
    - [ ] T/F/D to Climb
    - [ ] Cruise

Additional features:
- [ ] Basic IFR performance (dep gradient/app slope)
- Standalone:
    - [ ] Wind computer
    - [ ] '3:1' descent planner (target angle or rate)
    - [ ] Conversion Tools