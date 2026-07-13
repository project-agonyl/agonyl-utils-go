# pcdatafile

`pcdatafile` owns the single 28-byte Zone Server PC class-data row. The format contains fourteen little-endian 16-bit values.

`Value` and `SetValue` use zero-based field indexes so callers can assign runtime or UI names without moving path or presentation metadata into the codec.
