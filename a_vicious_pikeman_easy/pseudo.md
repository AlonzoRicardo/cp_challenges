# Lore
You are given a number of problems which were designed for a contest in which the pikeman participated. For each problem, you are given the estimated time in minutes for solving the problem. Calculate the maximum number of problems a pikeman can solve in the contest, and the minimum penalty he can get, under the assumptions that these estimations are correct. You may assume that the pikemen are very efficient: submissions are always correct, and after submitting a problem they start solving the next problem immediately.

# Input
Input starts with two integers on a single line `1≤𝑁≤10ˆ4` and `1≤𝑇≤10ˆ9`, the number of problems in the ancient contest and the total length of the contest in minutes. Then follows a line with four integers `1≤𝐴,𝐵,𝐶,𝑡0≤10ˆ6`, where `𝑡0 (𝑡0 ≤ 𝐶)` specifies the time in minutes required for solving the first problem, and the rest of the times `𝑡1,…,𝑡𝑁−1` are given by: 

`𝑡𝑖=((𝐴𝑡𝑖−1+𝐵)mod 𝐶)+1, 𝑖 ∈ [1,𝑁−1]`

# Output
Output should consist of two integers: the maximum number of problems a pikeman can solve within the time limit, and the total penalty he will get for solving them. As the penalty might be huge, print it modulo `1000000007`. Print them on the same line, separated by a single space.

# Sample Input
```            
1 3
2 2 2 1
```

# Sample Output
```
1 1
```

# Pseudocode
```
<!-- Given Variables -->
TotalNumberOfProblems = 1≤TotalNumberOfProblems≤10ˆ4
TotalContestLength = 1≤TotalContestLength≤10ˆ9
A = 1≤𝐴≤10ˆ6
B = 1≤𝐵≤10ˆ6
C = 1≤𝐶≤10ˆ6
FirstProblemResolveTime = (FirstProblemResolveTime ≤ 𝐶) and 1≤FirstProblemResolveTime≤10ˆ6

<!-- Created Variables -->
problemSolvingTimes = [FirstProblemResolveTime]
sortedProblemSolvingTimes = []
problemsSolved = 0
totalTimeSolvingProblems = 0

for n in TotalNumberOfProblems
    time = ((𝐴*n−1+𝐵)mod 𝐶)+1
    push time to problemSolvingTimes

sortedProblemSolvingTimes = sort(ProblemSolvingTimes)

while totalTimeSolvingProblems <= TotalContestLength
    if sortedProblemSolvingTimes is empty
        break loop;
    else
        problemTime = pop(sortedProblemSolvingTimes, 1)
        totalTimeSolvingProblems += problemTime
        problemsSolved++

print "problemsSolved totalTimeSolvingProblems%1000000007"
```
