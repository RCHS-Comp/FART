# FART

Fast Arithmetic Resolution Tool

A tool to help solve the area and perimeter for triangles.

## Usage

FART [FileName]

**Note: By default FART will try to open Questions.txt**

### File format

If the line starts with a `#` then it will be ignored.

The file format goes as follows:

```
[TriangleType] [Value1],[Value2],[Value3]
```

The available triangle types are:
* Scalene
* Isosceles
* Equilateral

When inputing the values, only input the needed values.

Here are some examples:
* Scalene 82,58,100.44
* Isosceles 47,76
* Equilateral 51

## Building

To build FART, run `go build Main.go Calculate.go FileOps.go`

You can rename the `Main` file that will be output to FART if you wish, it has no effect on execution.
