## 785A - Anton and Polyhedrons

### Description
Given `n` polyhedron names, sum the total number of faces across all polyhedrons.

Each polyhedron has a fixed number of faces:
- Tetrahedron: 4
- Cube: 6
- Octahedron: 8
- Dodecahedron: 12
- Icosahedron: 20

### Mathematical Logic
For each polyhedron, add its face count to the sum.

### Complexity
- Time: `O(n)` - processing each polyhedron once.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
