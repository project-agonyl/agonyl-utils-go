# npcfavorfile

`npcfavorfile` owns the 8-byte `ZoneData/npc/FavIndex.dat` row layout. Each row contains a signed favor index and a 32-bit NPC type.

Files must be row aligned. Both fields have byte-preserving setters.
