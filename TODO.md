# TODO

## Techdept

* \[ ] add `README.md` with issue description

  See `can_reach_corner/README.md`.

* \[ ] migrate table tests from structs to maps

  ```go
  tests := map[string]struct {
    input  int
    output int
  }
  ```

* \[ ] make table tests parallel at the top

  ```go
  func TestMe(t *testing.T) {
    t.Parallel()
    // declare tests
    // run tests
  }
  ```
