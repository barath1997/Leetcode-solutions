func addBinary(a string, b string) string {
    
	result := make([]string,max(len(a),len(b))+1)
	resultIndex := len(result)-1
	
   // filling empty spaces with 0
   maxLen :=  max(len(a),len(b))
   if maxLen == len(a) {
	   iterations := maxLen - len(b)
	   temp := ""
	   for i:=0;i<iterations;i++ {
		  temp += "0"
	   }
	   b = temp+b
   } else {
	   iterations := maxLen - len(a)
	   temp := ""
	   for i:=0;i<iterations;i++ {
		  temp += "0"
	   }
	   a = temp+a
   }

    i,j := len(a)-1,len(b)-1
	carry := "0"
	for ;(i>=0 || j>=0) ;i,j = i-1,j-1  {
		first,second := string(a[i]),string(b[j])
	   if first == "1" && second == "1" && carry == "1" {
		   carry = "1"
		   result[resultIndex] = "1"
		   resultIndex--
	   } else if first == "1" && second == "1" && carry == "0" {
		   carry = "1"
		   result[resultIndex] = "0"
		   resultIndex--
	   } else if first != second {
		   if carry == "0" {
			   result[resultIndex] = "1"
			   resultIndex--
		   } else {
			   result[resultIndex] = "0"
			   resultIndex--
			   carry = "1"
		   }
	   } else {
		   if carry == "0" {
			   result[resultIndex] = "0"
			   resultIndex--
		   } else {
			   result[resultIndex] = "1"
			   resultIndex--
			   carry = "0"
		   }
	   }
		
	}

	if carry == "1" {
	   result[0] = carry
	}

	res := ""
	for _,v := range result {
	   res+=v
	}
	
	return res
}