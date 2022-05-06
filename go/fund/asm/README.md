### Arrays, slices, and maps


#### Array
- an array in go is a fixed-length data type tht contains a contiguous block of elements of the same type

#### Pointer Array
- you can have an array of pointers.

#### passing array between function
- passing an array between functions can be an expensive operation in terms of memory and performance. When you pass variable is an array. this mean the entire array , regardless if its size ,is copied and passed to the function


#### Slice
- A slice is a data structure that provides a way for you to work with and manage collec- tions of data. Slices are built around the concept of dynamic arrays that can grow and shrink as you see fit. They’re flexible in terms of growth because they have their own built-in function called append, which can grow a slice quickly with efficiency. You can also reduce the size of a slice by slicing out a part of the underlying memory. Slices give you all the benefits of indexing, iteration, and garbage collection optimizations because the underlying memory is allocated in contiguous blocks.

1. internals
- slide are tiny object that abstract and manipulate an underlying array. they're three flied data structures that contain the metadata that Go needs to manipulate the underlying arrays (pointer, length, capacity)

#### Map
- a map data structure that provides  you with an unordered collection of key/value pairs

#### Wrapping up
 Arrays are the building blocks for both slices and maps.

 Slices are the idiomatic way in Go you work with collections of data. Maps are
the way you work with key/value pairs of data.

 The built-in function make allows you to create slices and maps with initial
length and capacity. Slice and map literals can be used as well and support set-
ting initial values for use.

 Slices have a capacity restriction, but can be extended using the built-in func-
tion append.

 Maps don’t have a capacity or any restriction on growth.

 The built-in function len can be used to retrieve the length of a slice or map.

 The built-in function cap only works on slices.

 Through the use of composition, you can create multidimensional arrays and
slices. You can also create maps with values that are slices and other maps. A
slice can’t be used as a map key.

 Passing a slice or map to a function is cheap and doesn’t make a copy of the
underlying data structure.