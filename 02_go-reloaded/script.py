#!/usr/bin/env python3
import os
import subprocess

passed = 0
failed = 0


def write_file(filename, content):
    with open(filename, "w") as f:
        f.write(content)


def run(input_file, output_file):
    subprocess.run(["go", "run", ".", input_file, output_file], stderr=subprocess.DEVNULL)


def check(test_num, description, expected):
    global passed, failed

    if not os.path.exists("result.txt"):
        print(f"FAIL Test {test_num}, result.txt not found")
        failed += 1
        return

    with open("result.txt") as f:
        actual = f.read().strip()

    if actual == expected:
        print(f"PASS Test {test_num} - {description}")
        passed += 1
    else:
        print(f"FAIL Test {test_num} - {description}")
        print(f"  Expected: {expected}")
        print(f"  Got:      {actual}")
        failed += 1


# Test  1
write_file("sample.txt",
    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: "
    "how (cap) did you get in my house (up, 2) ?\n"
)
run("sample.txt", "result.txt")
check(1, "low/cap/up modifiers",
    "If I make you breakfast in bed just say thank you instead of: "
    "How did you get in MY HOUSE?"
)

# Test 2
write_file("sample.txt",
    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure\n"
)
run("sample.txt", "result.txt")
check(2, "bin/hex number conversion",
    "I have to pack 5 outfits. Packed 26 just to be sure"
)

# Test 3
write_file("sample.txt",
    "Don not be sad ,because sad backwards is das . And das not good\n"
)
run("sample.txt", "result.txt")
check(3, "punctuation spacing",
    "Don not be sad, because sad backwards is das. And das not good"
)

# Test 4
write_file("sample.txt",
    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist "
    "who carries a raincoat . '\n"
)
run("sample.txt", "result.txt")
check(4, "cap with count + a/an correction + punctuation",
    "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"
)

total = passed + failed
print(f"\n{passed}/{total} tests passed")

if failed > 0:
    exit(1)