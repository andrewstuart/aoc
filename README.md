This is my repo for aoc2021.

I'm using [my fork of aocdl](https://github.com/andrewstuart/advent-of-code-downloader) to simplify the creation of
each day's input file and tests. e.g. `mkdir -p 2021/day12 && aocdl  -output
../input -force -story-output a.html` plus the necessary config in
~/.aocdlconfig:

```json
{
  "session-cookie": "<secret>",
  "story-out": "story.html",
  "test-output": "../test",
  "test-template-output": "./main_test.go",
  "test-template": "../../test_template.go.tpl",
  "template": "../../main_template.go.tpl",
  "template-output": "./main.go"
}
```
