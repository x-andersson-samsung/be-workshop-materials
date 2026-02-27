package exercise4

//Write a stateful server.
// Server should store 2 numbers: `reg1` `reg2`

// Each number can be modified using a PUT request to `/values/reg1` or `/values/reg2`
// Each number can be read using a GET request to `/values/reg1` or `/values/reg2`

// Support 4 operations:
// `/add` - returns `reg1` + `reg2`
// `/sub` - returns `reg1` - `reg2`
// `/mul` - returns `reg1` * `reg2`
// `/div` - returns `reg1` / `reg2` with an error if reg2 is 0.

// Write a client for your server.
