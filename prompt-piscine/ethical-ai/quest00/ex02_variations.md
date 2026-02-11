# Strengthening Through Variations

## Modified palindrome

```py
def checkPalindrome(input):
    formattedInput = ''.join(c.lower() for c in input if c.isalnum())
    left, right = 0, len(formattedInput) - 1
    while left < right:
        if formattedInput[left] != formattedInput[right]:
            return left
        left += 1
        right -= 1
    return True
```

## Reflection

While my solution works and handles the edge case listed, AI listed other efficient ways to do it which is particular about not needing to create a new variable for the string.
