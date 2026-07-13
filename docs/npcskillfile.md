# npcskillfile

`npcskillfile` owns the 18-byte `ZoneData/npc/NPCSkill` row layout. Records retain all bytes and expose the NPC type, skill family, attack type, target ranges, cooldown, and effect words.

Files must be row aligned. Setters update only their documented byte ranges.
