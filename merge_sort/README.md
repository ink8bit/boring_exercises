# Merge sort

Time complexity: `O(n log n)`
Space complexity: `O(n)`

The general idea of **merge sort** is that you divide your list down in two half size lists and _recursively_ invoke **merge sort** on those smaller lists.

Because **merge sort** uses a _recursive_ approach, you have to have a base case for it. The base case is when you have a list of one element, so you can simply return that list (because it is already _sorted_). If you have more than one lement in your list, you should merge your separate sorted lists together.
