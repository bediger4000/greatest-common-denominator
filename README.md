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

[Code](primes.go) for this version.
Turns out this is painful.

You can also observe that the GCD of numbers a, b, c is GCD(GCD(a,b),c)
[Euclid's Algorithm](https://en.wikipedia.org/wiki/Greatest_common_divisor#Euclid's_algorithm)
is a convenient way to find the GCD of 2 integers.
It's possible to walk an array of numbers,
calculating the GCD of the GCD-so-far with the next number,
and arrive at the GCD of the entire list of numbers.

The [code](gcd.go) for this method is a lot smaller and cleaner.

## Interview Analysis

Each of the methods I found use some other piece of code as a building block.
I had Euclid's Algorithm and a prime number generator that happened
to generate primes in order.
This made it possible to do these algorithms in a small amount of time.
These aren't easy to get correct in a short amount of time.
The interviewer really can't expect many candidates to know
Euclid's Algorithm off the top of their heads,
or to have a prime number generator at their recall.

That's a big drawback for this problem.
If the candidate knows Euclid's Algorithm,
this problem is almost trivial,
but it's quite a slog otherwise.

I'm not sure that merely knowing an algorithm from number theory
makes this an "[Easy]" problem.
If you don't know a GCD algorithm for 2 numbers,
and realize you can GCD-together all the numbers,
this is at least a "{Medium]".
Very few programmers are algebraists as well,
but even those that remember the Fundamental Theorem of Arithmetic
are merely going to spend a lot of time trying to get the
rest of the code correct.
