Paginated API

Problem Description

A third-party API that we're using has a paginated API that returns results in
chunks of N. This is implemented below in the 'FetchPage' function.

We don't think that API is very useful, and would prefer an implementation
where only one call fo 'Fetch' will return a given number of results,
abstracting away the need to do pagination.

Your task will be to implement Fetch()