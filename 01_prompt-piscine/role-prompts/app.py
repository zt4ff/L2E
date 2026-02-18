def checkPalindrome(input):
    reversedInput = input[-1:0:0]
    if input == reversedInput:
        return True
    return False