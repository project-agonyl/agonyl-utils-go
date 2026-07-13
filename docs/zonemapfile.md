# zonemapfile

`zonemapfile` owns the Zone Server `.map` binary layout: a 23-byte header, zero or more 6-byte warps, a fixed 256×256 cell mesh, and any trailing opaque bytes.

`Read` and `Write` preserve the opaque header and trailer. Editing is limited to the two-byte name, decoded warp words, and the documented movement, PK-level, and warp-index cell bits.
