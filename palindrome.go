package main
import "fmt"

func main(){
	ans:= checkPalindrome(121)
	fmt.Println(ans)
}

func checkPalindrome(x int) bool{
	var temp int = x;
	var rev int = 0;

	for(temp!=0){
		rev = (rev*10)+ (temp%10)
		temp /= 10;
	}

	if rev ==x {
		return  true
	} else{
		return false
	}

}