#include <string>
#include <iostream>
using namespace std;

// A function to check if a string str is palindrome
bool isThisAPalindrome(string str) {
  int l = 0;
  int h = str.length() - 1;
    while (h > l){
        if (str[l++] != str[h--]){
           return false;
        }
    }
    return true;
  }


int main(){
  string input;
  cout << "Please input a string to check if it is a palindrome:\n";
  cin >> input;
  string output = isThisAPalindrome(input)?"yes":"nope";
  cout << output;
}
