

✅ 1. Declaration & Syntax
var arr [3]int = [3]int{1, 2, 3} // array
var s []int = []int{1, 2, 3}     // slice


✅ 2. Size Flexibility
| Feature   | Array | Slice                |
| --------- | ----- | -------------------- |
| Size      | Fixed | Dynamic (resizable)  |
| Can grow? | ❌ No  | ✅ Yes (via `append`) |

✅ 3. Memory & Reference Behavior
array
func modifyArr(a [3]int) { a[0] = 100 }
slice
func modifySlice(s []int) { s[0] = 100 }

✅ 4. Built-in Operations
| Operation    | Array              | Slice |
| ------------ | ------------------ | ----- |
| `append()`   | ❌                  | ✅     |
| `len()`      | ✅                  | ✅     |
| `cap()`      | ❌ (not meaningful) | ✅     |
| `range` loop | ✅                  | ✅     |



func main() {
    arr := [3]int{1, 2, 3} // array
    s := []int{1, 2, 3} // slice

    fmt.Println(arr) // [1 2 3]
    fmt.Println(s)   // [1 2 3]

    // Modify slice
    s = append(s, 4)
    fmt.Println(s) // [1 2 3 4]
}


| Feature          | Array               | Slice                     |
| ---------------- | ------------------- | ------------------------- |
| Syntax           | `[n]T`              | `[]T`                     |
| Size             | Fixed               | Dynamic                   |
| Resizable        | ❌ No                | ✅ Yes (`append`)          |
| Passed by        | Value               | Reference                 |
| Memory efficient | ❌ Copies full array | ✅ Shares underlying array |
| Typical usage    | Rare                | Common                    |


