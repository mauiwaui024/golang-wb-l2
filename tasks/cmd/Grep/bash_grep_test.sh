#!/bin/bash

# Test file paths
FILE1="file1.txt"
FILE2="file2.txt"
FILE3="file3.txt"

# Test patterns
PATTERN="hello"

# Flags to test
FLAGS=(
    "-B 2"       # Print 2 lines before matching
    "-A 2"       # Print 2 lines after matching
    "-C 2"       # Print 2 lines of output context
    "-c"         # Print only a count of matched lines
    "-i"         # Ignore case distinctions
    "-v"         # Selected lines are those not matching any of the specified patterns
    "-F"         # Interpret pattern as a fixed string, not a regular expression
    "-n"         
)

# Variables to count successful and failed tests
successful_tests=0
failed_tests=0

# Test each flag with both your grep and standard grep
for FLAG in "${FLAGS[@]}"; do
    echo "Testing flag: $FLAG"

    echo "Testing with ./my_grep:"
    result_my_grep1=$(./my_grep $FLAG "$PATTERN" $FILE1)
    result_my_grep2=$(./my_grep $FLAG "$PATTERN" $FILE2)
    result_my_grep3=$(./my_grep $FLAG "$PATTERN" $FILE3)
    result_grep1=$(grep $FLAG "$PATTERN" $FILE1)
    result_grep2=$(grep $FLAG "$PATTERN" $FILE2)
    result_grep3=$(grep $FLAG "$PATTERN" $FILE3)
    
    if [[ $result_my_grep1 != $result_grep1 || 
          $result_my_grep2 != $result_grep2 || 
          $result_my_grep3 != $result_grep3 ]]; then
        echo "Failed test for $FLAG"
        echo "Expected:"
        echo "$result_grep1"
        echo "$result_grep2"
        echo "$result_grep3"
        echo "Got:"
        echo "$result_my_grep1"
        echo "$result_my_grep2"
        echo "$result_my_grep3"
        ((failed_tests++))
    else
        ((successful_tests++))
    fi

    echo "------------------------------------------"
done

echo "Successful tests: $successful_tests"
echo "Failed tests: $failed_tests"
