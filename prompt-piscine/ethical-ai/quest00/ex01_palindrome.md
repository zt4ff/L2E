# The Palindrome Challenge

## Do it yourself

### Pseudocode

```txt
FUNCTION checkPalindrome
REVERSE input - store in new variable called reversedInput
IF input == reversedInput
RETURN true
ELSE RETURN false
```

### Solution with python

```py
def checkPalindrome(input):
    reversedInput = input[::-1]
    if input == reversedInput:
        return True
    return False
```

### Tests

```py
print(checkPalindrome("racecar")) # True
print(checkPalindrome("hello")) # False
print(checkPalindrome("A man a plan a canal Panama")) # False
```

### Comments

The problem felt straight forward to me. A palindrome is a letter that equals itself if reversed. So I thought to reverse the letter and compare them. Python don't have a reserve function, or any I can think of, so I used the split operator `[start:end:step]` and using -1 counts from the last character in the string, hence, reversed.

## Reflection

### What did you learn from solving it before asking AI?

Before asking AI, I learnt what a palindrome is. I also learnt to think throughly with my brain.

### How is your understanding different now?

After sending the prompt to AI, I saw that I was missing spacing and case-sensitivity edge cases. Though the time complexity of my solution is `O(n)`, there are other and better ways to go about it.

### Could you now write similar functions (e.g., reverse a string) without help?

Yes I can, my original solution utilized reversing a string
