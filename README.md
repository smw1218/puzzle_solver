# Puzzle Solver

This small app solves a puzzle I got for my birthday which has
9 square pieces. Each piece has 4 Hot Air baloon halves on each
edge. The halves are assymetric with one half having the basket
and the other half without that matches.

I realized that this puzzle was about brute force rather than 
reasoning out a solution. I decided to use some code to do the
tedious part.

I think the total possibilities is quite large. Since each piece
can have 4 orientations and there are 9 pieces so there are 4 * 9!
combinations or about 1.4 million. Some of these solutions are 
duplicates in just a different orientation, but it's still a ton of
combinations.

I came up with a technique to reduce the solution space. As pieces
are placed, the adjacent pieces constrain the solution and there are
no possible pieces that can be placed. At that point, no more pieces
need to be tried to eliminate the already placed pieces.

The specific technique that I employed was to try each piece as the
center piece. Then place pieces in a spiral out from that center.
