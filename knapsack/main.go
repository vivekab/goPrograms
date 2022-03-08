package main

import "fmt"

func main() {
    n := 0
    c := 0
    fmt.Scanf("%d",&n)
    fmt.Scanf("%d",&c)
    wt := make([]int,n)
    val := make([]int,n)
    for i := range val{
        fmt.Scanf("%d",&val[i])
    }
    for i := range wt{
        fmt.Scanf("%d",&wt[i])
    }
    t := make([][]int,n+1)
    for i := range t{
        t[i] = make([]int,c+1)
    }
    fmt.Println(wt,val,n,c)
    fmt.Println(t)
    for i:=1;i<=n;i++{
        for j:=1;j<len(t[i]);j++{
            if wt[i-1]<=j{
                t[i][j] = max(val[i-1]+t[i-1][j-wt[i-1]],t[i-1][j])
            }else{
                t[i][j] = t[i-1][j]
            }
            
        }
    }
    return
}
func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}