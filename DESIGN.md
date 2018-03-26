# Djavul

The Djavul project is divided into a frontend and a backend component, where the frontend is responsible for rendering, audio playback, window creation and event handling, and the backend is responsible for handling the core logic of the game engine.

The frontend and backend components communicate using IPC (either through network sockets or named pipes). Currently, the frontend listens for incoming connections and the backend connects to the frontend. In the future, the direction of communication may be reversed to support multiplayer games and observer mode.

## Frontend

The [Djavul frontend](https://github.com/sanctuary/djavul/tree/master/cmd/djavul-frontend) is platform independent and handles:

* rendering with [Pixel](https://github.com/faiface/pixel),
* window creation and event handling with [PixelGL](https://github.com/faiface/pixel) (based on [GLFW](http://www.glfw.org/) and [OpenGL](https://www.opengl.org/)),
* and audio playback with [Beep](https://github.com/faiface/beep) (based on [ALSA](https://www.alsa-project.org/) on Linux, [OpenAL](https://www.openal.org/) on Darwin and FreeBSD, [WinMM](https://msdn.microsoft.com/en-us/library/windows/desktop/dd743834(v=vs.85).aspx) on Windows, [android/media/AudioManager](https://developer.android.com/reference/android/media/AudioManager.html) on Android, and [AudioContext](https://www.w3.org/TR/webaudio/#AudioContext) on web platforms).

## Backend

The [Djavul backend](https://github.com/sanctuary/djavul/tree/master/dll/djavul) is currently platform dependent (see [#1](https://github.com/sanctuary/djavul/issues/1) for the issue tracking cross-platform support) and runs on Windows or Linux through [Wine](https://www.winehq.org/).

The backend is split into two binaries, the `djavul.exe` executable and the `djavul.dll` shared library.

### djavul.exe

The `djavul.exe` executable is producted by applying a binary patch to the original `diablo.exe` executable (version 1.09b). See https://github.com/sanctuary/djavul-patch for further instructions.

In essence, the only difference between `djavul.exe` and `diablo.exe` is that the import section of `diablo.exe` has been amended to include the shared library `djavul.dll`, and hooks have been added to the assembly to invoke the exported functions of `djavul.dll` at key locations. Using this technique, large components of the Diablo 1 game engine have be rewritten in Go (e.g. the [dungeon generation of Cathedral](https://github.com/sanctuary/djavul/tree/master/cmd/l1)).

For technical background and historical anecdotes of how the Djavul.patch is created, see [[1](https://github.com/sanctuary/djavul/blob/5662c93cf2e45b0cdb863b99e686f3c7450c0dbc/dlls/djavul/README.md)].

### djavul.dll

The `djavul.dll` shared library is essentially a wrapper around the dedicated packages of the [d1](https://github.com/sanctuary/djavul/tree/master/d1) directory, each of which is directly associated with a source file of `diablo.exe`, using the same naming convention as the [notes](https://github.com/sanctuary/notes) repository.

Other than being a wrapper for the `d1` packages, the `djavul.dll` shared library handles the ICP communiction with the Djavul frontend.

It may also contain a few easter eggs and development features.
