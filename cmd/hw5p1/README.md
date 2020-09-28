# Frequency Analysis

## How to Run
Have any version of `Go` installed, run `go run main.go` from inside this directory. The program will take input from stdin.

Alternatively, I included binaries for Windows/Mac/Linux, although only Linux is tested.


## Complexity Analysis
"N" is the number of characters in the input.

Reading from stdin takes "N" time.
Converting input to lowercase takes "N" time.
Writing lowercase ascii characters to new string takes "N" time (and <="N" space).
Filling initial frequency map takes constant (26 operations) time.
Filling frequency map with total occurrences of each character of input takes "N" time.
Outputting the difference in frequencies of each character takes constant time (another 26 operations).

Total time = O(4N) + O(1) = O(n)
