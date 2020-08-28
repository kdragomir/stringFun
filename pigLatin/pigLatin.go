package main

import (
  "fmt"
  "unicode"
)
func doesItStartWithAVowel(word string) (result bool){
  switch word[0]{
  case 'a','e','i','o','u':
    return true
  }
  return
}

func seperateTheWords(sentence string) (arrayofWords []string){
  var arrayStuff []string
  currentWord:= ""
  for i := 0; i < len(sentence); i++{
    if unicode.IsSpace(rune(sentence[i])){
      arrayStuff = append(arrayStuff,currentWord)
      currentWord = ""
    }else if i == len(sentence)-1 {
      currentWord = currentWord + string(sentence[i])
      arrayStuff=append(arrayStuff, currentWord)
    }else{
      currentWord = currentWord + string(sentence[i])    }
  }
  return arrayStuff
}

func pigify(sentence string)(pigLatinSentence string){
  var wordsArray []string = seperateTheWords(sentence)
  for i:= 0; i< len(wordsArray); i++{
    if doesItStartWithAVowel(wordsArray[i]){
      wordsArray[i] = wordsArray[i]+ "yay"
    }else{
    var temp string = wordsArray[i]
    manipulatedTempString := ""
    var tempMover rune = rune(wordsArray[i][0])
    for letter:=0; letter <= len(wordsArray[i]); letter++{
      if letter ==len(wordsArray[i]){
        manipulatedTempString = manipulatedTempString+ string(tempMover)+ "ay"
      }else if letter !=0 {
        manipulatedTempString = manipulatedTempString + string(temp[letter])
      }}
      wordsArray[i]= manipulatedTempString
    }
    pigLatinSentence = pigLatinSentence +" "+ wordsArray[i]
    fmt.Println(pigLatinSentence)
  }
  return pigLatinSentence
}

func main(){
  fmt.Println(pigify("this is a test string"))
}
