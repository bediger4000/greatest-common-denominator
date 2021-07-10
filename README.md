# Daily Coding Problem: Problem #931 [Easy]

This problem was asked by Amazon.

Given n numbers, find the greatest common denominator between them.

For example, given the numbers [42, 56, 14], return 14.

## Analysis

This requires some knowledge of algebra,
namely the [Fundamental Theorem of Arithmetic](https://en.wikipedia.org/wiki/Fundamental_theorem_of_arithmetic),
That is,
there's a unique prime factorization for every number greater than 1.

Once you realize that theorem applies,
you get the unique prime factorization of all the numbers,
then count the occurances of each prime number in all the facotorizations.

This would require generating prime numbers up to the magnitude
of the largest of the n numbers.

