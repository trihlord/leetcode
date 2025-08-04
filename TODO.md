# TODO

## Techdept

* \[ ] add `README.md` with issue description

  See `can_reach_corner/README.md`.

* \[ ] migrate table tests to structs

  ```go
  tests := []struct {
    name   string
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

* \[ ] move go modules to main

  ```shell
  problem_solution_module
  ├── main.go
  ├── main_test.go
  ├── go.mod
  └── README.md
  ```
