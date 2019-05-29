package main
import "fmt"
import "sort"
import "sync"

func main (){

var matriz1=[]int{5, 3, 2, 7, 13,15, 6, 18, 1,4,9,17,16,19,11,8,23,14,20,12}
	fmt.Println(matriz1)
	  dividir(matriz1)
}
func dividir(s []int){
	long:=len(s)/4
		var wg sync.WaitGroup
		wg.Add(4)
		go func() {
		   defer wg.Done()
			b:=s[:long]
			fmt.Println("cuarta:",b)
			sort.Ints(b)
			impSlice(b)
		}()
		go func() {
		 defer wg.Done()
			a:=s[long:long*2]
			fmt.Println("cuarta:",a)
			sort.Ints(a)
			impSlice(a)
		}()
		go func() {
		 defer wg.Done()
			c:=s[long*2:long*3]
			fmt.Println("cuarta:",c)
			sort.Ints(c)
			impSlice(c)
			
		}()
		go func() {
		 defer wg.Done()
			d:=s[long*3:long*4]
			fmt.Println("cuarta:",d)
			sort.Ints(d)
			impSlice(d)
		}()
		wg.Wait()
		sort.Ints(s)
		fmt.Println("toda la matriz en orden: ",s)
		
}
func impSlice(s []int) {
	fmt.Printf("ordenado: %d \n", s)
}

//go run c:\ego\matriz.go