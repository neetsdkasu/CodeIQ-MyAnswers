package main;import"fmt";func f(n,a,b,c,d,e int64){if n>0{fmt.Println(a+b+c+d+e);f(n-1,d+e,a+d,b,e,c)}};func main(){f(100,1,0,0,0,0)}