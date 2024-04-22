Your task is to implement a simple container of integer numbers. As a first step, consider implementing the following two operations:

ADD <value> should add the specified integer value to the container. Returns an empty string.
EXISTS <value> should check whether the specific integer value exists in the container. Returns "true" if the value exists, "false" otherwise.
The container is supposed to be empty at the beginning of execution.

Given a list of queries, return the list of answers for these queries. To pass to the next level you have to pass all tests at this level.

Example

For
queries = [
    ["ADD", "1"],
    ["ADD", "2"],
    ["ADD", "5"],
    ["ADD", "2"],
    ["EXISTS", "2"],
    ["EXISTS", "5"],
    ["EXISTS", "1"],
    ["EXISTS", "4"],
    ["EXISTS", "3"],
    ["EXISTS", "0"]
]
the output should be solution(queries) = ["", "", "", "", "true", "true", "true", "false", "false", "false"].

Explanation:

Add 1, 2, 5, 2 -> [1, 2, 5, 2]
Numbers 2, 5, 1 exist in the container
Numbers 4, 3, 0 don't exist in the container