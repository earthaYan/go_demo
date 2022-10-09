package mathClass
import "fmt"
func Sub(x,y int) int {
    return x - y
}
func InCrementSelf(){
    a:=1
    var b int
    a++
    b=a
    fmt.Println(b)
}
func GetGrade(){
    var grade string = "B"
    var marks int = 90
    switch marks {
        case 90: grade = "A"
        case 80: grade = "B"
        case 50,60,70 : grade = "C"
        default: grade = "D"  
    }
    switch{
        case grade == "A" :
            fmt.Printf("优秀!\n" )    
        case grade == "B", grade == "C" :
            fmt.Printf("良好\n" )      
        case grade == "D" :
            fmt.Printf("及格\n" )      
        case grade == "F":
            fmt.Printf("不及格\n" )
        default:
            fmt.Printf("差\n" );
    }
    fmt.Printf("你的等级是 %s\n", grade );      
}

func GetType(){
    var x interface{}

    switch i := x.(type) {
        case nil:  
            fmt.Printf(" x 的类型 :%T",i)                
        case int:  
            fmt.Printf("x 是 int 型")                      
        case float64:
            fmt.Printf("x 是 float64 型")          
        case func(int) float64:
            fmt.Printf("x 是 func(int) 型")                      
        case bool, string:
            fmt.Printf("x 是 bool 或 string 型" )      
        default:
            fmt.Printf("未知型")    
    }     
}
func GetSum(){
    sum:=0
    for i:=0;i<=10;i++{
        sum+=i
    }
    fmt.Println(sum)
}
func GetSumOnlyCondition(){
    sum:=1
    for ;sum<=10;{
        sum+=sum
    }
    for sum<=10{
        sum=2*sum
    }
    fmt.Println(sum)
}
func LoopNest(){
    var i,j int
    for i=2;i<100;i++{
        for j=2;j<=(i/j);j++{
            if i%j==0{
                break
            }
        }
        if(j>(i/j)){
            fmt.Printf("%d  是素数\n", i);
        }
    }
}
func LoopBreakWithLabel(){
    fmt.Println("---- break label ----")
       for i := 1; i <= 3; i++ {
          fmt.Printf("i: %d\n", i)
          ref: 
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                break ref
            }
    }
}