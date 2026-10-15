func plusMinus(arr []int32) {
    // Write your code here
    l := len(arr)
    p := 0
    n := 0
    z := 0
    
    for i:=0; i<len(arr); i++{
        if arr[i]>0{
            p++
        }else if arr[i]<0 {
            n++
        }else if arr[i] == 0{
            z++
        }
    }
    fmt.Println(float64(p)/float64(l))
    fmt.Println(float64(n)/float64(l))
    fmt.Println(float64(z)/float64(l))
}
