# skilldelayfile

`skilldelayfile` owns the 5-byte `SkillDelay.dat` row layout. The last byte is both the high byte of the recovered delay word and the recovered delay-info byte; its setters intentionally preserve that packed behavior.

Files must be row aligned.
