# ZoneData codec coverage

This manifest maps every binary loader in `a3-core/internal/zone_server/zonedata` to its filesystem pattern and reusable codec. Text formats are listed for completeness but remain outside binary-codec scope.

| a3-core loader | Relative path pattern | Codec package | Status | Release |
|---|---|---|---|---|
| `LoadMapFile` | `ZoneData/map/*.map` | `zonemapfile` | Added | v0.5.0 |
| `LoadNPCSkillTableFile` | `ZoneData/npc/NPCSkill` | `npcskillfile` | Added | v0.5.0 |
| `LoadNPCFavorIndexFile` | `ZoneData/npc/FavIndex.dat` | `npcfavorfile` | Added | v0.5.0 |
| `LoadPCDataTableRowFile` | `ZoneData/pc/{0,1,2,3}` | `pcdatafile` | Added | v0.5.0 |
| `LoadSkillDataTableFiles` | `ZoneData/skill/{0,1,2,3}` | `skilldatafile` | Added | v0.5.0 |
| `LoadSkillDelayFile` | `ZoneData/skill/SkillDelay.dat` | `skilldelayfile` | Added | v0.5.0 |
| `LoadPassiveSkillDataFile` | `ZoneData/skill/PsvSkill.dat` | `passiveskillfile` | Added | v0.5.0 |
| `LoadHiredSoldierSkillDataTableFiles` | `ZoneData/skill/HSST{0,1,2,3}` | `hiredsoldierskillfile` | Added | v0.5.0 |
| `LoadNPCDataFile` | `ZoneData/npc/<id>` | `npcfile` | Existing | Before v0.5.0 |
| `LoadNPCDropTableFile` | `ZoneData/npc/<id>.itm` | `dropfile` | Existing | Before v0.5.0 |
| `LoadNPCSpawnFile` | `ZoneData/map/<id>.n_ndt` | `spawnlist` | Existing | Before v0.5.0 |
| `LoadItemTable0BaseFile` | `ZoneData/item/0` | `itemfile` | Existing | Before v0.5.0 |
| `LoadItemTable0ExtendedFile` | `ZoneData/item/0ex` | `itemfile` | Existing | Before v0.5.0 |
| `LoadItemTable1File` | `ZoneData/item/1` | `itemfile` | Existing | Before v0.5.0 |
| `LoadItemTable2File` | `ZoneData/item/2` | `itemfile` | Existing | Before v0.5.0 |
| `LoadItemTable3File` | `ZoneData/item/3` | `itemfile` | Existing | Before v0.5.0 |
| `LoadItemCombinationTableFile` | `ZoneData/item/ItemCombinationData` | `itemcombinationdata` | Existing | Before v0.5.0 |
| `LoadCashItemTableFile` | `ZoneData/item/CashItem.dat` | `cashitemfile` | Planned | v0.6.0 |
| `LoadSetItemTableFiles` | `ZoneData/item/SetItem*` | `setitemfile` | Planned | v0.6.0 |
| `LoadPresentItemSetTableFile` | `ZoneData/item/PresentItemSet.dat` | `presentitemsetfile` | Planned | v0.6.0 |
| `LoadPetTableFile` | `ZoneData/item/pet` | `petfile` | Planned | v0.6.0 |
| `LoadShueCombinationTableFile` | `ZoneData/item/ShueCombinationData` | `shuecombinationfile` | Planned | v0.6.0 |
| `LoadLotteryPresentTableFile` | `Event/LotteryItem.dat` | `lotteryfile` | Planned | v0.6.0 |
| `LoadDerbyGiftTableFile` | `Event/DerbyGift.dat` | `derbygiftfile` | Planned | v0.6.0 |
| `LoadA3PresentRecordsFile` | `Event/A3Present.dat` | `a3presentfile` | Planned | v0.6.0 |
| `LoadQuestExRecordFile` | `ZoneData/quest/*` | `questexfile` | Planned | v0.7.0 |
| `LoadSQuestScenarioFile` | `ZoneData/SQuest/*` | `squestfile` | Planned | v0.7.0 |
| `LoadSQuestQuizTableFile` | `ZoneData/SQuest/QuizTable.dat` | `squestquizfile` | Planned | v0.7.0 |
| `LoadPartyQuestConfigFile` | `ZoneData/PQuest/*` | `partyquestconfigfile` | Planned | v0.7.0 |
| `LoadPartyQuestPortalFile` | `ZoneData/PQuest/*Portal*` | `partyquestportalfile` | Planned | v0.7.0 |
| `LoadA3MessageTableFile` | `A3Msg_Zone_Tw.dat`, event siblings | `messagefile` | Planned | v0.7.0 |
| `LoadTowerConfigFile` | `Tower/*` | `towerconfigfile` | Planned | v0.7.0 |
| `LoadTowerTreasureTableFile` | `Tower/*` | `towertreasurefile` | Planned | v0.7.0 |
| `LoadOXQuizConfigFile` | `OXQuiz/*` | `oxquizconfigfile` | Planned | v0.7.0 |
| `LoadOXQuizTableFile` | `OXQuiz/*` | `oxquizfile` | Planned | v0.7.0 |
| `LoadTyrBaseFile` | `ZoneData/tyr/*/Base.tyr` | `tyrbasefile` | Planned | v0.7.0 |
| `LoadTyrPortalFile` | `ZoneData/tyr/*/WarpPortal.tyr` | `tyrportalfile` | Planned | v0.7.0 |
| `LoadTyrUpgradeFile` | `ZoneData/tyr/*/Upgrade.tyr` | `tyrupgradefile` | Planned | v0.7.0 |
| `LoadTyrStartPointFile` | `ZoneData/tyr/*/StartPoint.tyr` | `tyrstartpointfile` | Planned | v0.7.0 |
| `LoadTyrGiftFile` | `ZoneData/tyr/*/Gift.tyr` | `tyrgiftfile` | Planned | v0.7.0 |
| `LoadTyrNPCRegenFile` | `ZoneData/tyr/*/NPCRegen.tyr` | `tyrnpcregenfile` | Planned | v0.7.0 |
| `LoadTyrSkillLayerFile` | `ZoneData/tyr/*/SkillLayer.tyr` | `tyrskilllayerfile` | Planned | v0.7.0 |
| `LoadNPCShopInfoFile` | `ZoneData/shop/*.txt` | None | Text, out of scope | — |
