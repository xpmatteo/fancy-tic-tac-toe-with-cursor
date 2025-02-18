This is a fancy version of Tic-Tac-Toe with a twist in the rules:
- the X player wins if the board ends in a "winning" position for either X or O
- the O player wins if no "winning" position is possible

This was almost all driven by Cursor, with TDD.

Claude Sonnet 3.5 did almost all the implementation.  It remained on top of the logic until I asked it 
to detect incomplete board configurations that will inevitably lead to a draw, and hence a victory for the O player.  
Then I had to guide it much more closely; it could not even find good example on its own.

The implementation is Go with WASM.

# How to

Run the tests with `make test`.  Compile with `make`.  Serve with `make serve` and then open localhost:8080
