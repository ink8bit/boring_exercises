# Quick sort

**Quick sort** is one of the most useful and powerful sorting algorithms out there, and it's not terribly difficult to conceptualize. Above I mentioned that occasionally JavaScript doesn't mergesort for `Array.prototype.sort`. In those other cases, it's usually some variant on **quick sort**.

It's another **divide-and-conquer**, **recursive** algorithm but it takes a slightly _different_ approach. The basic gist is that you take the last element in the list and call that the **pivot**. Everything that's smaller than the pivot gets put into a **left** list and everything that's greater get's put in a **right** list. You then call **quick sort** on the left and right lists independently (hence the recursion). After those two sorts come back, you concatenate the sorted left list, the pivot, and then the right list (in that order). The base case is when you have a list of length 1 or 0, where you just return the list given to you.

Another **Big O** of `O(n log n)` but takes up less memory than **merge sort** so it is often favored.

> However it does really poorly if you pass it a sorted list. Think about it. It would always have a pivot of the biggest number which defeats the effectiveness of the **divide-and-conquer** approach as one side will always contain all the elements. Hence not good for lists you expect may already be sorted.

There are some tricks to employ to get around that like checking the beginning, middle, and end numbers and swapping them to try to get the best pivot.
