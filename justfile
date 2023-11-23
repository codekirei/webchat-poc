default:
  @just --list

configure:
  @cp -r configs/. .

clean:
  ./bin/clean
  go clean

# DOCS: https://just.systems/man/en/chapter_1.html
