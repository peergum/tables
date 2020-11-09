# tables
This Go package helps you present rows and columns in tables in a nice text way

## features
* no separation between consecutive rows with same number of columns 
* max size of rows determined independently of number of columns
* all columns are centered

## examples
### example 1
```
Table{
    Table{1, "a", 3.1}
}
```
```
╭───┬───┬─────╮
│ 1 │ a │ 3.1 │
╰───┴───┴─────╯
```

### Example 2
```
Table{
 	Table{12, 2, 3, 3.1},
 	Table{4, "something", 6},
 	Table{"a", "b", "c"},
 	Table{1, 2},
 	Table{12, 13, 14},
 	Table{1, 2, 4},
}
```
```
╭────┬───┬───┬──────╮
│ 12 │ 2 │ 3 │ 3.1  │
├───┬┴───┴───┴──┬───┤
│ 4 │ something │ 6 │
│ a │     b     │ c │
├───┴─────┬─────┴───┤
│    1    │    2    │
├─────┬───┴─┬───────┤
│ 12  │ 13  │  14   │
│  1  │  2  │   4   │
╰─────┴─────┴───────╯
```

## future developments
* different alignment modes (left, right, decimal justification)
* forced column sizes
* embedded tables

