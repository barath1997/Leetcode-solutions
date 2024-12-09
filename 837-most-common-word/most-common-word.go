import "strings"

func mostCommonWord(paragraph string, banned []string) string {
    
    // split the parapgrpah into words using multiple delimeters
    para := strings.FieldsFunc(paragraph,split)

    wordFreq := make(map[string]int)

    bannedList := make(map[string]bool,len(banned))
    
    // add banned words as lowercase into map , it is done so that search opearation takes O(1)
    for _,v := range banned {
        bannedList[strings.ToLower(string(v))] = true
    }
    
    // iterating over paragrapgh
    for _,v := range para {

      // if word is in bannedlist or is a punctuation then skip it dont add it to wordFreq Map
      if bannedList[strings.ToLower(string(v))] ||  checkWordIsPunctuation(string(v)) {
         continue
      }

      // add only good words to map and increase their frequencey
      wordFreq[strings.ToLower(v)]++
    }
    
    result , maxFreq := "",-1

    // check for maxFreqency among the good words
    for word,freq := range wordFreq {
        if  freq > maxFreq {
            maxFreq = freq
            result = word
        }
    }
    
    // return the word which has maxFreqency among the good words
    return result
    
}

func checkWordIsPunctuation(str string) bool {
    if str == " " || str == "!" || str == "?" || str == "'" || str == "," || str == ";" || str == "." {
        return true
    }

    return false
}

func split(str1 rune) bool{
    str := string(str1)
    return str == " " || str == "!" || str == "?" || str == "'" || str == "," || str == ";" || str == "."
}