# go-euler

A few solutions (in Go) for problems presented in http://projecteuler.net. Not all solutions are yet available.

Please note a few things:

* The authors of http://projecteuler.net specifically ask users not to post solutions outside their site (so it won't spoil the fun for people who haven't yet been able to solve a given problem.) Unfortunately, the cat is way out of the bag on this one. It's even hard to Google for common algorithms pertaining to certain problems without hitting on the solution.

* These programs are not "production clean" and not optimized in any way, in most cases. I'm writing them just for fun, during my spare time. In some cases, I keep them simple so other people can easily follow the logic. In others, even though a more elegant solution exists, I use brute force if the problem is simple enough. TL;DR: This is **not** how I usually write Go programs!

* There's no attempt to use a common library between the programs. You may find an implementation of a function that works for problem X, but was slightly altered to also work for problem X+10. I usually don't "backport" the fixes, as long as the original implementation works for the original problem.

* In many solutions seen on the Project Euler website, people resort to all kinds of language features to simplify their lives. In many cases, this completely kills the fun of a problem. You'll find situations in which I did a sum of two big numbers by hand (since this was the goal of the problem), and in others I used math/big (since dealing with large numbers was just the means to a goal.)

* Remember, these programs are ugly and may contain bugs. If you find something you dislike, send me a PR.

Have fun
