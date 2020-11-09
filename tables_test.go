package tables

var table0 = Table{Table{}}
var table3 = Table{Table{1, "a", 3.1}}
var tableSimple = Table{Table{12,2,3,3.1},Table{4,"something",6},Table{"a","b","c"},Table{1,2},Table{12,13,14},Table{1,2,4}}

func Example0() {
	table0.Print()
	//Output:
	//
}

func Example3() {
	table3.Print()
	//Output:
	//╭───┬───┬─────╮
	//│ 1 │ a │ 3.1 │
	//╰───┴───┴─────╯
}

func ExampleSimple() {
	tableSimple.Print()
	// Output:
	//╭────┬───┬───┬──────╮
	//│ 12 │ 2 │ 3 │ 3.1  │
	//├───┬┴───┴───┴──┬───┤
	//│ 4 │ something │ 6 │
	//│ a │     b     │ c │
	//├───┴─────┬─────┴───┤
	//│    1    │    2    │
	//├─────┬───┴─┬───────┤
	//│ 12  │ 13  │  14   │
	//│  1  │  2  │   4   │
	//╰─────┴─────┴───────╯
}

