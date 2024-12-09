import "strings"

func mostCommonWord(paragraph string, banned []string) string {

    para := strings.FieldsFunc(paragraph,split)

    wordFreq := make(map[string]int)

    for _,v := range para {
      if checkBannedWord(v,banned) ||  checkWordIsPunctuation(string(v)) {
         continue
      }
      wordFreq[strings.ToLower(v)]++
    }
    
    result , maxFreq := "",-1
    for word,freq := range wordFreq {
        if  freq > maxFreq {
            maxFreq = freq
            result = word
        }
    }

    return result
    
}

func checkWordIsPunctuation(str string) bool {
    if str == " " || str == "!" || str == "?" || str == "'" || str == "," || str == ";" || str == "." {
        return true
    }

    return false
}

func checkBannedWord(word string, bannedList []string) bool {
      
      for _,v := range bannedList {
        if strings.ToLower(v) == strings.ToLower(word) {
            return true
        }
      }

      return false

}

func split(str1 rune) bool{
    str := string(str1)
    return str == " " || str == "!" || str == "?" || str == "'" || str == "," || str == ";" || str == "."
}