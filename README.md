# ConvexPolygon
Task: There are N points on the plane in the order of their circumvention. Need to find out if they form a convex polygon in this order.

When you walk the polygon counterclockwise, the turn will always be to the left, and when you walk clockwise, to the right.
If one edge of the polygon corresponds to the AB{x1; y1} vector , followed by an edge corresponding to the BC{x2; y2} vector, then the direction of turn in this pair of consecutive edges will be fit the sign of the expression x1 * y2 - y1 * x2.
To left turn, this value will be positive, and to right turn, it will be negative (if the Y axis is directed upward). A zero value means that the vectors are collinear.

So, it is necessary to check if all pairs of vectors lying on neighboring edges turn in the same direction.
