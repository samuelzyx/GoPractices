The next step is to support removal from the container:

REMOVE <value> should remove a single occurrence of the specified value from the container. If the value has multiple occurrences, only one of them should be removed.
Previous level functionality should remain the same. To pass to the next level you have to pass all tests at this level.

Given a list of queries, return the list of answers for these queries.

Example

For
queries = [
    ["ADD", "1"],
    ["ADD", "2"],
    ["ADD", "2"],
    ["ADD", "3"],
    ["EXISTS", "1"],
    ["EXISTS", "2"],
    ["EXISTS", "3"],
    ["REMOVE", "2"],
    ["REMOVE", "1"],
    ["EXISTS", "2"],
    ["EXISTS", "1"]
]

the output should be solution(queries) = ["", "", "", "", "true", "true", "true", "true", "true", "true", "false"].

Explanation:

Add 1, 2, 2, 3 -> [1, 2, 2, 3]
Numbers 1, 2, 3 exist in the container
Remove 2, 1 -> [2, 3]
Number 2 exists in the container, number 1 doesn't exist