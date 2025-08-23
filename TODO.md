# TODO

## Techdept

* \[ ] add `README.md` with issue description

  See `can_reach_corner/README.md`.

  ```md
  # Issue no. with title

  Link to issue description.

  Summary.
  ```

* \[ ] migrate table tests to structs

  ```go
  tests := []struct {
    name string
    in   int
    out  int
  }
  ```

* \[ ] make table tests parallel at the top

  ```go
  func TestMe(t *testing.T) {
    t.Run("Exmaple", func (t *testing.T) {
      t.Parallel()
      // Impl details
    })
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

* \[ ] rm `test` from `*_test.go` package

  ```go
  package mysolution

  import "testing"

  func TestMySolution () {
    // Impl details
  }
  ```
