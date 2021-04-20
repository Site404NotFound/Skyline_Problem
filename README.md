# Skyline_Problem
Skyline Problem solved using Divide & Conquer method (written in Golang).

## Time Complexity
``O(n log n)``

## Coding Problem:
We want to compute the skyline of a city given buildings positions and heights.

The input is a List of Buildings.
A Building (class provided) consist of the left side, the right side and the height.

The output is a List of SkylinePoints.
A SkylinePoint (class provided) consist of an x value and a height. The height corresponds
to the height of the skyline to the right of the x value until the next SkylinePoint.

Given a List of Buildings return a List of skyline points that represents the skyline
Modify the method computeSkyline() to contain your algorithm.

## Example:
### Input
- Building(1, 3, 2) (Shown as A below)
- Building(2, 3, 4) (Shown as B below)
- Building(4, 6, 3) (Shown as C below)
- Building(6, 7, 3) (Shown as D below)

### Visual Representation
Only for reference, no need to display in solution.  I displayed it anyway using a somewhat inelegant approach.
```
     B
     B   C C D
   A A   C C D
   A A   C C D
 0 1 2 3 4 5 6 7
 ```

### Output:
- SkylinePoint(1, 2)
- SkylinePoint(2, 4)
- SkylinePoint(3, 0)
- SkylinePoint(4, 3)
- SkylinePoint(7, 0)

## Example Output
![Results Example](/example_images/results.png)
