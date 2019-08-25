# Quick sort

Time complexity: `O(n log n)`

**Quick sort** is an exmaple of _divide-and-conquer_ algorithm. The basic idea is that you take the last element in the list and call that the `pivot`. Everything that's smaller than the pivot gets put into a **left** list and everything that's greater get's put in a **right** list. After that you should call **quick sort** on the **left** and **right** lists independently using recursion.

The base case is when you have a list of one element - you have to return just this list.

After **left** and **right** lists are sorted, you can grab a `pivot` and concatenate those three in the following way: `left` + `pivot` + `right`.

> NOTE: Obviously you shouldn't use it when your list is sorted.
