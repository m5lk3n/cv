# https://github.com/charmbracelet/vhs

Output cv.gif

Require ./cv
Require echo

Set Theme "SpaceGray Eighties"
Set Shell "bash"
Set FontSize 14
Set WindowBar Colorful
Set Padding 2
Set Width 1000
Set Height 800

Type "./cv basics" Sleep 500ms  Enter
Sleep 2000ms

Type "./cv experience" Sleep 500ms  Enter
Sleep 4500ms

Type "./cv education" Sleep 500ms  Enter
Sleep 2000ms

Type "./cv version" Sleep 500ms  Enter
Sleep 5000ms
