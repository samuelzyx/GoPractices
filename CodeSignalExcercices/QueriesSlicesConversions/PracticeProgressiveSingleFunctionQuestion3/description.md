The next step is to support the operation for finding the nearest integer in the container greater than the provided one:

GET_NEXT <value> should return the minimal integer in the container that is strictly greater than the provided value. In case there is no such integer in the container, return empty string.
Previous levels functionality should remain the same. To pass to the next level you have to pass all tests at this level.

Given a list of queries, return the list of answers for these queries.

Example

For
queries = [
    ["ADD", "1"],
    ["ADD", "2"],
    ["ADD", "2"],
    ["ADD", "4"],
    ["GET_NEXT", "1"],
    ["GET_NEXT", "2"],
    ["GET_NEXT", "3"],
    ["GET_NEXT", "4"],
    ["REMOVE", "2"],
    ["GET_NEXT", "1"],
    ["GET_NEXT", "2"],
    ["GET_NEXT", "3"],
    ["GET_NEXT", "4"]
]
the output should be solution(queries) = ["", "", "", "", "2", "4", "4", "", "true", "2", "4", "4", ""].

Explanation:

Add 1, 2, 2, 4 -> [1, 2, 2, 4]
Get Next 1 -> "2"
Get Next 2 -> "4"
Get Next 3 -> "4"
Get Next 4 -> ""
Remove 2 -> [1, 2, 4]
Get Next 1 -> "2"
Get Next 2 -> "4"
Get Next 3 -> "4"
Get Next 4 -> ""