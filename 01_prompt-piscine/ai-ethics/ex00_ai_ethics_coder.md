# Learning to Use AI Ethically as a Coder

## Part A: The Critical Distinction

### How have you used AI for coding so far?

I use AI as a tool for studying before.

### Do you ask AI for solutions before trying yourself?

No, I don't

### Can you explain code you've submitted without AI's help?

Yes, I can

### What would happen if AI was suddenly unavailable during an exam or interview?

Nothing would happen, I'd perform effectively as I ought to

### Identify your current pattern: Which learner are you now?

I am a Learner B: "AI is my learning amplifier"

## Part B: The Wrong Way vs. The Right Way

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

### Reflection

#### What did you learn by struggling first?

Before asking AI, I learnt what a palindrome is. I also learnt to think throughly with my brain.

#### How is your understanding different than if you'd just asked for the solution?

After sending the prompt to AI, I saw that I was missing spacing and case-sensitivity edge cases. Though the time complexity of my solution is `O(n)`, there are other and better ways to go about it.

#### Can you now implement similar functions (reverse a string, find duplicates) without AI?

Yes I can, my original solution utilized reversing a string

#### What mental model did you build?

I built a mental model of resilient and deep thinking.

## Part C: Testing Your Understanding

### Modified palindrome

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

### Reflection

While my solution works and handles the edge case listed, AI listed other efficient ways to do it which is particular about not needing to create a new variable for the string.

## Part D: The Fairness Contract

I will use AI when:

- After I've attempted a problem for at least 20 minutes

- To understand why my solution works/doesn't

- To explore alternatives after I have a working solution

I will NOT use AI when:

- I haven't tried the problem myself

- I'm taking an assessment or test

- I need to build fundamentals

I know I'm using AI fairly when:

- I can explain my code without looking at AI's response

- I could solve similar problems without AI

- I feel more confident in my abilities

---

Signature: Oluwasegun Kayode

Date: 09/02/2026

## Part E: Real-World Scenario Analysis

### Interview: "Explain how you'd implement a caching system." If you always relied on AI, can you answer?

If I've always relied on AI to design systems, I won't able to think through designing a caching system without considering tradeoffs, edge cases, etc. It would be hard to speak concurrently through such situation.

### Production bug at 2 AM: AI is unavailable. Can you debug code you don't fully understand?

It would be hard to fix the problem I didn't understand how the project is developed, the tradeoffs and the structure of the project. It would take a lot of time figuring out the structure of the code rather than the bug itself.

### New tech with little documentation: If you never learned to read docs and experiment, what happens?

To learn such library, I'd need to read the official documentation of the library. Understnd the fundamentals of the library and practise building a solution or project with the new library.

### Reflection

Using AI fairly would allow me understand concept without absolute dependency on AI to do everyday tasks.

## Part F: Building Irreplaceable Skills

# Building Irreplaceable Skills

| Skill                                | Description                               | Rating | Improvement Plan                                                                                                         |
| ------------------------------------ | ----------------------------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------ |
| Problem Decomposition                | Breaking down problems logically          | 3/5    | Practise problem solving on leetcode                                                                                     |
| Systems Thinking                     | Understanding how components interact     | 3/5    | Pickup system designs and architecture course and study it, including systems case studies                               |
| Critical Evaluation                  | Knowing when code is wrong or inefficient | 4/5    | Learn and practise software development best practises                                                                   |
| Debugging Mindset                    | Investigating unexpected behavior         | 4/5    | Learn and practise software testing                                                                                      |
| Conceptual Understanding (the "why") | Knowing WHY, not just HOW                 | 3/5    | Reading documentation to understand the design decisions and tradeoff behind my primary programing language / frameworks |

## 3 Specific action to take to improve my Conceptual Understanding this week

- Reading documentation to understand the design decisions and tradeoff behind my primary programing language / frameworks
- Rebuild projects I've built with my primary language with other languages/framework to understand the tradeoffs
- Teach and explain concepts to others
