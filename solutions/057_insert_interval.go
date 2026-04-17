package main

import "fmt"

// #57 Insert Interval [Medium]
// Tags: Array

func insert(intervals [][]int, newInterval []int) [][]int {
    
    insertI := -1
    for i := 1; i < len(intervals); i++{
     if insertI == -1 {
      if (intervals[i][0] > newInterval[0]) {
        insertI = i
        intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
        break
      }
     }
    }
    
    i := 1
    for i < len(intervals){
      c := intervals[i]
      p := intervals[i-1]

      if (p[0] <= c[0]) && (c[0] <= p[1]) {
        if p[1] < c[1]{
          p[1] = c[1]
        }
        intervals = append(intervals[:i],intervals[i+1:]...)
      } else {
        i ++
      }
    }

    fmt.Println(insertI,intervals)
    return nil
}

func main() {
	fmt.Println(insert([][]int{{1,3},{6,9}}, []int{2,5})) // → [[1,5],[6,9]]
	fmt.Println(insert([][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}, []int{4,8})) // → [[1,2],[3,10],[12,16]]
}
