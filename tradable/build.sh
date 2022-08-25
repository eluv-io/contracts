#
# Generate go bindings for tradable contracts
#

# Must use:
# - abigen v1.9.19 (go-ethereum rev 3e0641923d78bf1905e596a3a41a54277540bec7)
# - solc 0.5.4
abigen --version | grep -q 1.9.19-stable || badreq=true
solc --version | grep -q 0.5.4+commit.9549d8ff || badreq=true 
if test $badreq; then
  echo "Must use abigen 1.9.19 and solc 0.5.4"
  exit
fi

abigen --sol elv_tradable_full.sol --pkg=elv_tradable --out ../build/tradable/elv_tradable_full.go && echo "Go bindings generated successfully" 



