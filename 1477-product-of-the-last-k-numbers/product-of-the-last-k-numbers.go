type ProductOfNumbers struct {
    stream []int 
}


func Constructor() ProductOfNumbers {
    s := ProductOfNumbers{
        stream : make([]int,0),
    }
    return s
}


func (this *ProductOfNumbers) Add(num int)  {
    this.stream = append(this.stream, num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    res := 1
    for i:=0; i < k; i++ {
        res = res* this.stream[len(this.stream)-1-i]
    }
    return res
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */