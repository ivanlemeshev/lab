#!/bin/bash

# ==========================================
# BASH SCRIPTING BEST PRACTICES CHEAT SHEET
# ==========================================

# 1. THE "STRICT MODE" BOILERPLATE
# --------------------------------
# set -e: Exit immediately if a command exits with a non-zero status.
# set -u: Treat unset variables as an error and exit immediately.
# set -o pipefail: The return value of a pipeline is the status of
#                  the last command to exit with a non-zero status.
set -euo pipefail
# Optional: set -x  # Uncomment this to print commands as they run (debugging)

# 2. TRAPS & CLEANUP (Best Practice)
# ----------------------------------
# Always define a cleanup function for temporary files or graceful exits.
cleanup() {
  echo ""
  echo "--- Script Finished (Cleanup trigger) ---"
  # rm -f /tmp/tempfile # Example of removing temp files
}
# Trap EXIT signal (happens on success or failure)
trap cleanup EXIT

# 3. FUNCTIONS
# ------------
# - Use 'local' variables to avoid global scope pollution.
# - Return values are usually exit codes (0-255), not strings.
# - To return a string, 'echo' it and capture output.
log_msg() {
  local level="${1:-INFO}"              # Default to INFO if $1 is missing
  local msg="${2:?Message is required}" # Error if $2 is missing

  # Print to stderr (>2) so it doesn't mess up piped output
  echo "[$level] $msg" >&2
}

main() {
  echo "=== Starting Script ==="

  # 4. VARIABLES & QUOTING
  # ----------------------
  # ALWAYS quote variables to prevent word splitting/globbing.
  local name="Dev User"
  local file_path="my file.txt" # Spaces require quotes

  # Parameter Expansion / Defaults
  # ${var:-default} -> Use default if var is unset/null
  # ${var:=default} -> Set var to default if unset/null
  log_msg "INFO" "Hello, ${name:-Unknown}"

  # 5. ARGUMENTS
  # ------------
  # $0 = Script name, $1 = First arg, $# = Arg count, $@ = All args
  log_msg "INFO" "Script name: $0"
  log_msg "INFO" "Arguments passed: $#"

  # Safely iterating over arguments
  for arg in "$@"; do
    log_msg "ARG" "Processing: $arg"
  done

  # 6. CONDITIONALS
  # ---------------
  # Use [[ ]] for strings/files (safer than [ ]).
  # Use (( )) for arithmetic.

  local count=10

  # String check
  if [[ "$name" == "Dev User" ]]; then
    echo "User match confirmed."
  fi

  # Arithmetic check
  if ((count > 5)); then
    echo "Count is greater than 5."
  fi

  # File check
  if [[ -f "$0" ]]; then
    echo "I exist as a file ($0)."
  fi

  # Combined logic: && (AND), || (OR)
  if [[ -n "$name" ]] && ((count == 10)); then
    echo "Name is not empty AND count is 10."
  fi

  # 7. ARRAYS
  # ---------
  local fruits=("Apple" "Banana" "Cherry")

  # Add an element
  fruits+=("Date")

  echo "First fruit: ${fruits[0]}"
  echo "All fruits: ${fruits[@]}"
  echo "Total fruits: ${#fruits[@]}"

  # 8. LOOPS
  # --------
  echo "--- Looping over array ---"
  for fruit in "${fruits[@]}"; do
    echo "Fruit: $fruit"
  done

  # C-Style loop
  echo "--- C-Style Loop ---"
  for ((i = 0; i < 3; i++)); do
    echo "Index: $i"
  done

  # While loop (Reading line by line safely)
  # Using a Here-String (<<<) to simulate a file for this demo
  echo "--- Reading Lines ---"
  while IFS= read -r line; do
    echo "Read line: $line"
  done <<<"Line 1
Line 2
Line 3"

  # 9. ARITHMETIC
  # -------------
  # Bash does integer math only.
  ((count = count + 5))
  echo "New count: $count"       # 15
  echo "Quick math: $((10 / 2))" # 5

  # 10. INPUT (Interactive)
  # -----------------------
  # -p provides a prompt.
  # -r prevents backslash interpretation.
  # checking if input is a terminal (-t 0) to avoid hanging in automation
  if [[ -t 0 ]]; then
    echo "--- Interactive Check ---"
    # read -r -p "Press Enter to continue..." dummy_var
    echo "Skipping actual 'read' so this script runs automatically."
  fi
}

# Run main and pass all arguments to it
main "$@"
