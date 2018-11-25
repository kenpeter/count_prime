package main
import "fmt"

func main() {
	input := 10 
	countPrimes(input)
}

func countPrimes(n int) []int {
    // init arr with len, cap
	// need to have len+1 to cancel 0 index
    notP := make([]int, n+1, n+1)  
    
    // all of them not prime
    for key, _ := range notP {
        notP[key] = -1 
    }
	notP[0] = 0   
 
    cnt := 0
    
    // why starts at 2, because 1 * other, still others...
	// need to 2*2, 2*3, ... 
    // we need to conver the n as well 
    for i:=2; i<=n; i++ {
        if notP[i] == -1 {
            cnt = cnt + 1
            
            // look at the algorithm, 2*2, 2*3, 2*4, ....
            for j:=2; i*j<=n; j++ {
                // this not prime 
                notP[i*j] = i*j 
            }
        }
    }

	// orig
	fmt.Println(notP)

	// type + {}
	arr := []int{} 
	// e.g. start at 1, end at 10
	for i:=1; i<=n; i++ {
		if notP[i] == -1 {
			arr = append(arr, i)
		}	
	}

	// print prime
	fmt.Println(arr)

    return notP 
}
