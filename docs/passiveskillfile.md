# passiveskillfile

`passiveskillfile` owns the 68-byte `PsvSkill.dat` level-row layout. It exposes six metadata dwords and eleven effect dwords while retaining the complete row bytes.

Files must be row aligned. Indexed effect setters validate their bounds.
