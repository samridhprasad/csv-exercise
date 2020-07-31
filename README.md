# SCOIR Technical Interview for Back-End Engineers
This repo contains an exercise intended for Back-End Engineers.

## Instructions
1. Clone this repo and checkout the branch with your name.
1. Using technology of your choice, complete [the assignment](./Assignment.md).
1. Update this README with
    * a `How-To` section containing any instructions needed to execute your program.
    * an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Before the deadline, submit a pull request with your solution against your branch (please do not submit a PR against master).

## Assumptions
- The validation condition for matching internal_id was that it must be an 8 digit int where *every digit* is positive integer.
- No 3rd party dependencies/libraries for the solutions
- Pass input-directory, output-directory, error-directory via flags
### Build
```
 go build

```
### Run:
```
 ./csv-exercise -input-directory=/path/to/your/inputdir/ -output-directory=/path/to/your/outputdir/ -error-directory=/path/to/your/errordir/

```
## Future Improvements
 - If I had more time, I would have learned more about go routines and replace the `select{}` in main() meant to block exit so the input directory is watched continuously 
 - Would definitely prefer actual unit tests than hotswapping csv files to test different cases
 - I would also look into using well-tested file system event watchers to make this more production-ready 
 - Added [additional notes](Hindsight.md) for discussion with the team

## Expectations
1. Please take no more than 8 hours to work on this exercise. Complete as much as possible and then submit your solution.
1. This exercise is meant to showcase how you work. With consideration to the time limit, do your best to treat it like a production system.
