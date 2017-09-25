package main

import (
  "io"
  "errors"
    "fmt"
  "os"
)

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
    fi, err := os.Stat(src)
    if err != nil {
        return err
    }

    switch mode := fi.Mode(); {
      case mode.IsDir():
          return errors.New("I cannot copy directories. Give me only regular files.")
      case mode.IsRegular():
          break;
      default:
          return errors.New("Please sir, give me a regular file.")
    }

    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    return nil
}

func main() {
  if len(os.Args) <= 2 {
      fmt.Println("usage: cp source_file target_file")
      os.Exit(1)
  }

  source := os.Args[1]
  dst := os.Args[2]

  err := Copy(source, dst)

  if err != nil {
    fmt.Println(err)
  }
}
