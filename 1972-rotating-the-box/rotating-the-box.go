// Rock # = 35
// Block * = 42
// Empty . = 46

func rotateTheBox(box [][]byte) [][]byte {
    for _,row := range box {
        e:= len(row)-1        
        for i:= len(row)-1; i > -1; i-- {
            if row[i] == 46 {
                continue
            } else if row[i] == 35 {
                stone:= row[i]
                row[e] = stone
                if e!=i {row[i] = 46}
                e--
            } else {
                e = i-1
            }
        }   
    }

    col := len(box)
    row := len(box[0])
    rBox := make([][]byte, row)   

    for ind := range rBox {
        rBox[ind] = make([]byte, col)
    }

    for i:=0; i < row; i++ {
        for j:= col-1; j > -1; j-- {
            rBox[i][col-j-1] = box[j][i]
        }
    }
    
    return rBox 
}