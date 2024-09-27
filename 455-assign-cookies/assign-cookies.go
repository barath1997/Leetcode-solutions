func findContentChildren(g []int, s []int) int {
    // brute force solution
    /*sort.Ints(g)
    sort.Ints(s)
    gn := len(g)
    satisfiedChildren := 0
    for i:=0;i<gn;i++ {
        for j :=0;j<len(s);j++ {
            if s[j] - g[i] >= 0 {
              satisfiedChildren++
			  if j != len(s)-1 {
              s = append(s[:j],s[j+1:]...)
			  } else {
				s = s[:j]
			  }
			  break
            }
        }
    }

	return satisfiedChildren*/

    // optimised solution : 
    slices.Sort(g); slices.Sort(s)
    child := 0
    cookieSize := 0
    totalChildren,totalCookies := len(g),len(s)
    for cookieSize < totalCookies && child < totalChildren {
         if s[cookieSize] >= g[child] {
            child++
         }
         cookieSize++
    }
    return child
}

