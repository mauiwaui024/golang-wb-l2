#!/bin/bash

# Test file paths
FILE1="file1.txt"
FILE2="file2.txt"
FILE3="file3.txt"

# Test patterns
PATTERN="hello"

# Flags to test
FLAGS=(
    "-f 1"       # Select field 1
    "-f 2,3"     # Select fields 2 and 3
    "-f 3 -d','"       # Use comma as delimiter
    "-f 2 -s"         # Only print lines with delimiter
)

# Variables to count successful and failed tests
successful_tests=0
failed_tests=0

# Test each flag with both your my_cut and standard cut
for FLAG in "${FLAGS[@]}"; do
    echo "Testing flag: $FLAG"

    echo "Testing with ./my_cut:"
    result_my_cut1=$(cat $FILE1 | ./my_cut $FLAG)
    result_my_cut2=$(cat $FILE2 | ./my_cut $FLAG)
    result_my_cut3=$(cat $FILE3 | ./my_cut $FLAG)
    result_cut1=$(cut $FLAG < $FILE1)
    result_cut2=$(cut $FLAG < $FILE2)
    result_cut3=$(cut $FLAG < $FILE3)
    
    if [[ $result_my_cut1 != $result_cut1 || 
          $result_my_cut2 != $result_cut2 || 
          $result_my_cut3 != $result_cut3 ]]; then
        echo "Failed test for $FLAG"
        echo "Expected:"
        echo "$result_cut1"
        echo "$result_cut2"
        echo "$result_cut3"
        echo "Got:"
        echo "$result_my_cut1"
        echo "$result_my_cut2"
        echo "$result_my_cut3"
        ((failed_tests++))
    else
        ((successful_tests++))
    fi

    echo "------------------------------------------"
done

echo "Successful tests: $successful_tests"
echo "Failed tests: $failed_tests"
