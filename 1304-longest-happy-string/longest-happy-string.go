func longestDiverseString(a int, b int, c int) string {
    res := ""
    currA, currB, currC := 0, 0, 0
    totalLen := a+b+c
    i:=0

    for i<totalLen {
        if (currA != 2 && a>=b && a>=c) || (a>0 && (currB==2 || currC==2) )  {
            res+="a"
            currA++
            currB=0
            currC=0
            a--
        } else if (currB != 2 && b>=c && b>=a) || (b>0 && (currA==2 || currC==2) ) {
            res+="b"
            currB++
            currA=0
            currC=0
            b--
        } else if (currC != 2 && c>=a && c>=b) || (c>0 && (currA==2 || currB==2) )  {
            res+="c"
            currC++
            currB=0
            currA=0
            c--
        }

        i++
    }
 
    return res
}