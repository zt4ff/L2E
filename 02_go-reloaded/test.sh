#!/bin/bash

PASS=0
FAIL=0

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color
BOLD='\033[1m'

check() {
    local test_num="$1"
    local description="$2"
    local expected="$3"

    if [ ! -f "result.txt" ]; then
        echo -e "${RED}✗ Test $test_num FAILED${NC} — result.txt not found"
        ((FAIL++))
        return
    fi

    actual=$(cat result.txt)

    if [ "$actual" = "$expected" ]; then
        echo -e "${GREEN}✓ Test $test_num PASSED${NC} — $description"
        ((PASS++))
    else
        echo -e "${RED}✗ Test $test_num FAILED${NC} — $description"
        echo -e "  ${YELLOW}Expected:${NC} $expected"
        echo -e "  ${YELLOW}Got:     ${NC} $actual"
        ((FAIL++))
    fi
}

echo -e "${BOLD}========================================${NC}"
echo -e "${BOLD}        AUDIT TEST RUNNER               ${NC}"
echo -e "${BOLD}========================================${NC}"
echo ""

# ── Test 1 ──────────────────────────────────────────────────────────────────
echo -e "${BOLD}[Test 1]${NC} low/cap/up modifiers"
cat > sample.txt << 'EOF'
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
EOF
go run . sample.txt result.txt 2>/dev/null
check 1 "low/cap/up modifiers" \
    "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

# ── Test 2 ──────────────────────────────────────────────────────────────────
echo -e "${BOLD}[Test 2]${NC} bin/hex number conversion"
cat > sample.txt << 'EOF'
I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure
EOF
go run . sample.txt result.txt 2>/dev/null
check 2 "bin/hex number conversion" \
    "I have to pack 5 outfits. Packed 26 just to be sure"

# ── Test 3 ──────────────────────────────────────────────────────────────────
echo -e "${BOLD}[Test 3]${NC} punctuation spacing"
cat > sample.txt << 'EOF'
Don not be sad ,because sad backwards is das . And das not good
EOF
go run . sample.txt result.txt 2>/dev/null
check 3 "punctuation spacing" \
    "Don not be sad, because sad backwards is das. And das not good"

# ── Test 4 ──────────────────────────────────────────────────────────────────
echo -e "${BOLD}[Test 4]${NC} cap with count + a/an correction + punctuation"
cat > sample.txt << 'EOF'
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
EOF
go run . sample.txt result.txt 2>/dev/null
check 4 "cap with count + a/an correction + punctuation" \
    "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

# ── Summary ─────────────────────────────────────────────────────────────────
echo ""
echo -e "${BOLD}========================================${NC}"
TOTAL=$((PASS + FAIL))
if [ "$FAIL" -eq 0 ]; then
    echo -e "${GREEN}${BOLD}  ALL TESTS PASSED ($PASS/$TOTAL)${NC}"
else
    echo -e "${RED}${BOLD}  $FAIL/$TOTAL TESTS FAILED${NC} — $PASS/$TOTAL passed"
fi
echo -e "${BOLD}========================================${NC}"

# Exit with non-zero if any test failed
[ "$FAIL" -eq 0 ]