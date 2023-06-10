# Go Slice Utilities

### Module Name: slugo

This is a liblary for go that adds some slice-array functions which already exist in Javascript & Php. They are coded for general usage, so you could use them with any type of slices unless your action is non-logical. 

## Functions In slugo:

### 1 - [FilterStructSlices](#FilterStructSlices)

### 2 - [PopSlice](#PopSlice)

### 3 - [ReverseSlice](#ReverseSlice)

### 4 - [ShuffleSlice](#ShuffleSlice)

### 5 - [ReduceSlice](#ReduceSlice)

### 6 - [ReduceStructSlices](#ReduceStructSlices)

This is a sample struct slice which we'll use on our samples:

```go

type Person struct{
    Name string
    Surname string
    Age int
    Weight float32
    Height float64
    IsHealthy bool
}

PersonsSlice := []Person{
    {Name: "bob", Surname: "carter", Age: 36, Weight: 90.5, Height: 1.83, IsHealthy: true},
    {Name: "samantha", Surname: "sumner", Age: 21, Weight: 65.1, Height: 1.72, IsHealthy: true},
    {Name: "rick", Surname: "kingsley", Age: 52, Weight: 56.3, Height: 1.69, IsHealthy: false},
    {Name: "john", Surname: "doe", Age: 47, Weight: 83.0, Height: 1.92, IsHealthy: true},
    {Name: "jasmine", Surname: "johnson", Age: 29, Weight: 58, Height: 1.65, IsHealthy: false},
    {Name: "jeremy", Surname: "flavor", Age: 18, Weight: 78.5, Height: 1.88, IsHealthy: true},
    {Name: "abigail", Surname: "harris", Age: 38, Weight: 64.7, Height: 1.60, IsHealthy: true},
    {Name: "stephen", Surname: "shark", Age: 63, Weight: 55.8, Height: 1.65, IsHealthy: false},
}

```

## Attention

Because of the synthax of the go, i need to make some functions mutating and some functions are returning. I explain how to use all functions down below:

## <a id="FilterStructSlices">FilterStructSlices</a>

This function is inspired from `.filter()` function in javascript and structured for work the same way, except that function is mutating rather than the javascript version. It makes type checkings automatically. 

It has 4 arguments and all of them are mandatory.

#### Arguments are:

1 - A Reference Of Struct Slice.

2 - The Field Of Struct which you want to compare for.

3 - The operator of the comparison. You can do this comparisons: "==", "<", ">", "<=", ">=". You should be aware of, with strings and booleans you could do only "==" comparison.

4 - The comparison value. The type of the value should be same with the field's type. If you want to compare various types of numbers except
"int" type, you should convert that value first, for example `float32(1.83)` or `float64(1.65)` etc.

Here is some examples:

```go

slugo.FilterStructSlices(&PersonsSlice, "Name", "==", "bob");
slugo.FilterStructSlices(&PersonsSlice, "Age", "<", 40);
slugo.FilterStructSlices(&PersonsSlice, "Weight", "<=", float32(63.5));
slugo.FilterStructSlices(&PersonsSlice, "Height", ">=", float64(1.75));
slugo.FilterStructSlices(&PersonsSlice, "IsHealthy", "==", false);

```

## <a id="PopSlice">PopSlice</a>

This Function is inspired from `.pop()` function on javascript, it makes the same thing, and it's also mutating. It also takes a slice reference on it's argument.

```go

slugo.PopSlice(&PersonsSlice);

```

## <a id="ReverseSlice">ReverseSlice</a>

This Function is inspired from `.reverse()` function of javascript, it reverses the order of the Slice members and it's also mutating, it takes a slice on it's argument.

```go

slugo.ReverseSlice(PersonsSlice);

```

## <a id="ShuffleSlice">ShuffleSlice</a>

This Function is inspired from `shuffle()` function of php, it shuffles the order of the Slice members and it's also mutating, it takes a slice on it's argument.

```go

slugo.ShuffleSlice(PersonsSlice);

```

## <a id="ReduceSlice">ReduceSlice</a>

This Function is one of the functions inspired from `.reduce()` function in javascript. That function reduces the types of `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64` with arithmetic calculations, which is addition, subtraction, multiplication, division.

#### Arguments Are:

1 - Slice Which you want to reduce. Attention, it shouldn't be a field of a struct, for that i created <a id="ReduceStructSlices">ReduceStructSlices</a> function. It's NOT mutating, it returns the calculation value with the same type of the slice's type.

2 - The calculation type which you want to do. It should be string. Possible calculations: "+", "-", "/", "*"

```go

simpleIntSlice := []int{10, 20, 30, 40, 50, 60}

slugo.ReduceSlice(simpleIntSlice, "+");

```

That Function uses individual functions for reducing each type, the list of functions is:

```go

slugo.ReduceIntSlice()
slugo.ReduceInt8Slice()
slugo.ReduceInt16Slice()
slugo.ReduceInt32Slice()
slugo.ReduceInt64Slice()
slugo.ReduceUintSlice()
slugo.ReduceUint8Slice()
slugo.ReduceUint16Slice()
slugo.ReduceUint32Slice()
slugo.ReduceUint64Slice()
slugo.ReduceFloat32Slice()
slugo.ReduceFloat64Slice()

```

If you want to be more less affected about performance you can directly use the function that fits your use case.

## <a id="ReduceStructSlices">ReduceStructSlices</a>

Slice the field of a struct slice. This also inspired from `.reduce()`function of javascript. That function reduces the fields of a struct slice which has `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64` types with arithmetic calculations, also is addition, subtraction, multiplication, division. This function also returns the reduced value which has the same type of reduced field.

#### Arguments Are:

1 - Slice of a Struct Which you want to reduce. 

2 - The field which you want to slice, it should be a string.

3 - The calculation type which you want to do. It should be a string. Possible calculations: "+", "-", "/", "*"

```go

allAgesAddition := slugo.ReduceStructSlice(PersonsSlice, "Age", "+")

```

Also this function uses individual functions to reduce slice fields for each type, the list of That Functions:

```go

slugo.ReduceUintStructSlice()
slugo.ReduceUint8StructSlice()
slugo.ReduceUint16StructSlice()
slugo.ReduceUint32StructSlice()
slugo.ReduceUint64StructSlice()
slugo.ReduceintStructSlice()
slugo.Reduceint8StructSlice()
slugo.Reduceint16StructSlice()
slugo.Reduceint32StructSlice()
slugo.Reduceint64StructSlice()
slugo.ReduceFloat32StructSlice()
slugo.ReduceFloat64StructSlice()

```

If you want to be more less affected about performance you can directly use the function that fits your use case.


