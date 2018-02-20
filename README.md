# catch-signal

*catch-signal* illustrates a bug in the Windows PowerShell (possibly in connection with MINGW).

## Usage

* `PS C:\> .\catch-signal.exe > output`
* Press Ctrl-C
* Observe that the program exits immediately, rather than waiting five seconds
