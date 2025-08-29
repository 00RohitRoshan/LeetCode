func removeDuplicates(s string) string {
    b:=[]byte(s)
    
     strptr := -1
    for i:=0;i<len(b);i++{
        if(strptr==-1 || b[i]!=b[strptr]){
            strptr++
            b[strptr]=b[i]
        }else{
            strptr--
        }
    }

    // ans:=""
    // for i:=0;i<=strptr;i++{
    //     ans+=string(s[i])
    // }

    return string(b[:strptr+1])
}