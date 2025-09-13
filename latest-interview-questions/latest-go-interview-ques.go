package main
import "fmt"

func Add(ch1 chan int,a,b int) {
   ch1 <- a + b;
}

func Sub(ch2 chan int,a,b int) {
   ch2 <- a - b;
}

func Mul(a,b int,wg *sync.WaitGroup) int {
   defer wg.Done()
   return a * b;
}

func main(){
    ch1 := make(chan int)
    ch2 := make(chan int )
    go Add(ch1 , 4,7);
    go Sub(ch2, 6,2);
    a := <-ch1;
    b := <-ch2;
    var wg sync.WaitGroup
    wg.Add(1)
    mul := go Mul(a,b,&wg);
    fmt.Println("MUl : ",mul);
    wg.Wait();
}
