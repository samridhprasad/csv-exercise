# Team Discussion

On revisiting my code after a week, I found nuances that I had previously missed. 
- I wrote unit tests to confirm line-by-line validation was its own testable entity `validator_test.go`
- I wanted to improve the visibility of the PR by using Github Actions to run `go test -v` prior to merging
- I ran `go fmt` to fix any formatting mishaps 
- I found a bug where validation errors with names > 15 characters were not caught
//TODO: Bring PR on scoir's origin up to date with forked `samridhprasad/csv-exercise/{sprasad}` 

# Checklist for equivalence partitions:
- Test each validation constraint with blank, invalid & valid inputs

## ID Constraint 
-[x] passes on valid 8-digit non-zero input
-[x] raises appropriate validation error on invalid record
-[x] handles blank gracefully

## Name Constraints
### First Name Check
-[x] passes on valid 1-15 char string 
-[x] raises appropriate validation error on invalid record
-[x] handles blank gracefully

### Middle Name Check
-[x] passes on valid 0-15 char string 
-[x] raises appropriate validation error on invalid record
-[x] handles blank gracefully

### Last Name Check
-[x] passes on valid 1-15 char string 
-[x] raises appropriate validation error on invalid record
-[x] handles blank gracefully

## Phone Number Constraint
-[x] passes on valid ###-###-#### input
-[x] raises appropriate validation error on invalid record
-[x] handles blank gracefully

## Header Alignment Check
-[x] passes on valid h=5
-[x] raises appropriate validation error on invalid headers
-[x] handles blank gracefully 
