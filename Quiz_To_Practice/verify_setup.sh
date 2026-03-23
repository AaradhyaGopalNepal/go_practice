#!/bin/bash

# Verification script for Go Quiz Challenge Setup

echo "======================================"
echo "Go Quiz Challenge - Setup Verification"
echo "======================================"
echo ""

# Check Go installation
echo "1. Checking Go installation..."
if command -v go &> /dev/null; then
    go_version=$(go version)
    echo "   ✓ Go is installed: $go_version"
else
    echo "   ✗ Go is not installed. Please install Go first."
    exit 1
fi
echo ""

# Check directory structure
echo "2. Checking directory structure..."
all_dirs_exist=true
for week in {1..8}; do
    if [ $week -eq 1 ]; then dir="week1_basics"
    elif [ $week -eq 2 ]; then dir="week2_functions"
    elif [ $week -eq 3 ]; then dir="week3_collections"
    elif [ $week -eq 4 ]; then dir="week4_structs"
    elif [ $week -eq 5 ]; then dir="week5_interfaces"
    elif [ $week -eq 6 ]; then dir="week6_concurrency"
    elif [ $week -eq 7 ]; then dir="week7_advanced_concurrency"
    elif [ $week -eq 8 ]; then dir="week8_projects"
    fi

    if [ -d "$dir" ]; then
        echo "   ✓ $dir exists"
    else
        echo "   ✗ $dir is missing"
        all_dirs_exist=false
    fi
done
echo ""

# Check go.mod
echo "3. Checking go.mod..."
if [ -f "go.mod" ]; then
    echo "   ✓ go.mod exists"
    if grep -q "github.com/stretchr/testify" go.mod; then
        echo "   ✓ testify package is included"
    else
        echo "   ✗ testify package is missing"
    fi
else
    echo "   ✗ go.mod is missing"
fi
echo ""

# Check README
echo "4. Checking documentation..."
if [ -f "README.md" ]; then
    echo "   ✓ README.md exists"
else
    echo "   ✗ README.md is missing"
fi
echo ""

# Count challenge files
echo "5. Checking challenge files..."
go_files=$(find . -name "*.go" -type f | grep -E "(week[1-8])" | grep -v "_test.go" | wc -l)
test_files=$(find . -name "*_test.go" -type f | grep -E "(week[1-8])" | wc -l)
echo "   ✓ Found $go_files challenge files"
echo "   ✓ Found $test_files test files"
echo ""

# Check if tests can be compiled (they will be skipped but should compile)
echo "6. Checking if tests compile..."
echo "   Running: go test ./... -run=XXX (compile only)..."
if go test ./... -run=XXX &> /dev/null; then
    echo "   ✓ All tests compile successfully"
    echo "   Note: Tests are skipped by default. Remove t.Skip() when ready to test."
else
    echo "   ⚠ Some files have compilation issues."
    echo "   This is expected if you haven't started implementing yet."
fi
echo ""

echo "======================================"
echo "Setup Verification Complete!"
echo ""

if [ "$all_dirs_exist" = true ]; then
    echo "✅ Your Go Quiz Challenge is properly set up!"
    echo ""
    echo "Next steps:"
    echo "1. Start with Week 1 challenges"
    echo "2. Read the README.md for detailed instructions"
    echo "3. Uncomment code blocks as you work through them"
    echo "4. Remove t.Skip() in tests when ready to test"
    echo ""
    echo "Happy coding! 🚀"
else
    echo "⚠ Some issues were found. Please check the output above."
fi
echo "======================================"