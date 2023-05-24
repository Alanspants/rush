package main

import "fmt"

func Trifunc(n int){
    for i:=n;i>0;i--{
        for j:=n;j>0;j--{
        if j>13-i {
            if j>=10 {
            fmt.Printf("   ")
            }else{
                 fmt.Printf("  ")
                }
        }else{
            if  i <=3 && j>=10{
               fmt.Printf("%-3d",j)
            }else{
            fmt.Printf("%-2d",j)
            }
            }
        }

    fmt.Println()
    }
}

func main(){
    Trifunc(12)
}