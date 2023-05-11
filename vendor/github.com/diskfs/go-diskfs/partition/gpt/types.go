package gpt

// Type constants for the GUID for type of partition, see https://en.wikipedia.org/wiki/GUID_Partition_Table#Partition_entries
type Type string

// List of GUID partition types
const (
	Unused                   Type = "00000000-0000-0000-0000-000000000000"
	EFISystemPartition       Type = "C12A7328-F81F-11D2-BA4B-00A0C93EC93B"
	MBRPartitionScheme       Type = "024DEE41-33E7-11D3-9D69-0008C781F39F"
	IntelFastFlash           Type = "D3BFE2DE-3DAF-11DF-BA40-E3A556D89593"
	BIOSBoot                 Type = "21686148-6449-6E6F-744E-656564454649"
	SonyBootPartition        Type = "F4019732-066E-4E12-8273-346C5641494F"
	LenovoBootPartition      Type = "BFBFAFE7-A34F-448A-9A5B-6213EB736C22"
	PowerPCPRePBoot          Type = "9E1A2D38-C612-4316-AA26-8B49521E5A8B"
	ONIEBoot                 Type = "7412F7D5-A156-4B13-81DC-867174929325"
	ONIEConfig               Type = "D4E6E2CD-4469-46F3-B5CB-1BFF57AFC149"
	MicrosoftReserved        Type = "E3C9E316-0B5C-4DB8-817D-F92DF00215AE"
	MicrosoftBasicData       Type = "EBD0A0A2-B9E5-4433-87C0-68B6B72699C7"
	MicrosoftLDMMetadata     Type = "5808C8AA-7E8F-42E0-85D2-E1E90434CFB3"
	MicrosoftLDMData         Type = "AF9B60A0-1431-4F62-BC68-3311714A69AD"
	MicrosoftWindowsRecovery Type = "DE94BBA4-06D1-4D40-A16A-BFD50179D6AC"
	IBMGeneralParallelFs     Type = "37AFFC90-EF7D-4E96-91C3-2D7AE055B174"
	MicrosoftStorageSpaces   Type = "E75CAF8F-F680-4CEE-AFA3-B001E56EFC2D"
	HPUXData                 Type = "75894C1E-3AEB-11D3-B7C1-7B03A0000000"
	HPUXService              Type = "E2A1E728-32E3-11D6-A682-7B03A0000000"
	LinuxSwap                Type = "0657FD6D-A4AB-43C4-84E5-0933C84B4F4F"
	LinuxFilesystem          Type = "0FC63DAF-8483-4772-8E79-3D69D8477DE4"
	LinuxServerData          Type = "3B8F8425-20E0-4F3B-907F-1A25A76F98E8"
	LinuxRootX86             Type = "44479540-F297-41B2-9AF7-D131D5F0458A"
	LinuxRootArm             Type = "69DAD710-2CE4-4E3C-B16C-21A1D49ABED3"
	LinuxRootX86_64          Type = "4F68BCE3-E8CD-4DB1-96E7-FBCAF984B709"
	LinuxRootArm64           Type = "B921B045-1DF0-41C3-AF44-4C6F280D3FAE"
	LinuxRootIA64            Type = "993D8D3D-F80E-4225-855A-9DAF8ED7EA97"
	LinuxReserved            Type = "8DA63339-0007-60C0-C436-083AC8230908"
	LinuxHome                Type = "933AC7E1-2EB4-4F13-B844-0E14E2AEF915"
	LinuxRAID                Type = "A19D880F-05FC-4D3B-A006-743F0F84911E"
	LinuxExtendedBoot        Type = "BC13C2FF-59E6-4262-A352-B275FD6F7172"
	LinuxLVM                 Type = "E6D6D379-F507-44C2-A23C-238F2A3DF928"
	LinuxDMCrypt             Type = "7FFEC5C9-2D00-49B7-8941-3EA10A5586B7"
	LinuxLUKS                Type = "CA7D7CCB-63ED-4C53-861C-1742536059CC"
	FreeBSDData              Type = "516E7CB4-6ECF-11D6-8FF8-00022D09712B"
	FreeBSDBoot              Type = "83BD6B9D-7F41-11DC-BE0B-001560B84F0F"
	FreeBSDSwap              Type = "516E7CB5-6ECF-11D6-8FF8-00022D09712B"
	FreeBSDNANDFS            Type = "74BA7DD9-A689-11E1-BD04-00E081286ACF"
	FreeBSDUFS               Type = "516E7CB6-6ECF-11D6-8FF8-00022D09712B"
	FreeBSDZFS               Type = "516E7CBA-6ECF-11D6-8FF8-00022D09712B"
	FreeBSDVinum             Type = "516E7CB8-6ECF-11D6-8FF8-00022D09712B"
	AppleHFS                 Type = "48465300-0000-11AA-AA11-00306543ECAC"
	AppleUFS                 Type = "55465300-0000-11AA-AA11-00306543ECAC"
	AppleAPFS                Type = "7C3457EF-0000-11AA-AA11-00306543ECAC"
	AppleRAID                Type = "52414944-0000-11AA-AA11-00306543ECAC"
	AppleRAIDOffline         Type = "52414944-5F4F-11AA-AA11-00306543ECAC"
	AppleBoot                Type = "426F6F74-0000-11AA-AA11-00306543ECAC"
	AppleLabel               Type = "4C616265-6C00-11AA-AA11-00306543ECAC"
	AppleTVRecovery          Type = "5265636F-7665-11AA-AA11-00306543ECAC"
	AppleCoreStorage         Type = "53746F72-6167-11AA-AA11-00306543ECAC"
	DragonflyLabel32         Type = "9D087404-1CA5-11DC-8817-01301BB8A9F5"
	DragonflySwap            Type = "9D58FDBD-1CA5-11DC-8817-01301BB8A9F5"
	DragonflyUFS             Type = "9D94CE7C-1CA5-11DC-8817-01301BB8A9F5"
	DragonflyVINUM           Type = "9DD4478F-1CA5-11DC-8817-01301BB8A9F5"
	DragonflyCCD             Type = "DBD5211B-1CA5-11DC-8817-01301BB8A9F5"
	DragonflyLabel64         Type = "3D48CE54-1D16-11DC-8696-01301BB8A9F5"
	DragonflyLegacy          Type = "BD215AB2-1D16-11DC-8696-01301BB8A9F5"
	DragonflyHAMMER          Type = "61DC63AC-6E38-11DC-8513-01301BB8A9F5"
	DragonflyHAMMER2         Type = "5CBB9AD1-862D-11DC-A94D-01301BB8A9F5"
	SolarisBoot              Type = "6A82CB45-1DD2-11B2-99A6-080020736631"
	SolarisRoot              Type = "6A85CF4D-1DD2-11B2-99A6-080020736631"
	SolarisUsrAndAppleZFS    Type = "6A898CC3-1DD2-11B2-99A6-080020736631"
	SolarisSwap              Type = "6A87C46F-1DD2-11B2-99A6-080020736631"
	SolarisBackup            Type = "6A8B642B-1DD2-11B2-99A6-080020736631"
	SolarisVar               Type = "6A8EF2E9-1DD2-11B2-99A6-080020736631"
	SolarisHome              Type = "6A90BA39-1DD2-11B2-99A6-080020736631"
	SolarisAlternateSector   Type = "6A9283A5-1DD2-11B2-99A6-080020736631"
	SolarisReserved1         Type = "6A945A3B-1DD2-11B2-99A6-080020736631"
	SolarisReserved2         Type = "6A9630D1-1DD2-11B2-99A6-080020736631"
	SolarisReserved3         Type = "6A980767-1DD2-11B2-99A6-080020736631"
	SolarisReserved4         Type = "6A96237F-1DD2-11B2-99A6-080020736631"
	SolarisReserved5         Type = "6A8D2AC7-1DD2-11B2-99A6-080020736631"
	NetBSDSwap               Type = "49F48D32-B10E-11DC-B99B-0019D1879648"
	NetBSDFFS                Type = "49F48D5A-B10E-11DC-B99B-0019D1879648"
	NetBSDLFS                Type = "49F48D82-B10E-11DC-B99B-0019D1879648"
	NetBSDConcatenated       Type = "2DB519C4-B10E-11DC-B99B-0019D1879648"
	NetBSDEncrypted          Type = "2DB519EC-B10E-11DC-B99B-0019D1879648"
	NetBSDRAID               Type = "49F48DAA-B10E-11DC-B99B-0019D1879648"
	ChromeOSFirmware         Type = "CAB6E88E-ABF3-4102-A07A-D4BB9BE3C1D3"
	ChromeOSKernel           Type = "FE3A2A5D-4F32-41A7-B725-ACCC3285A309"
	ChromeOSRootFs           Type = "3CB8E202-3B7E-47DD-8A3C-7FF2A13CFCEC"
	ChromeOSReserved         Type = "2E0A753D-9E48-43B0-8337-B15192CB1B5E"
	MidnightBSDData          Type = "85D5E45A-237C-11E1-B4B3-E89A8F7FC3A7"
	MidnightBSDBoot          Type = "85D5E45E-237C-11E1-B4B3-E89A8F7FC3A7"
	MidnightBSDSwap          Type = "85D5E45B-237C-11E1-B4B3-E89A8F7FC3A7"
	MidnightBSDUFS           Type = "0394EF8B-237E-11E1-B4B3-E89A8F7FC3A7"
	MidnightBSDZFS           Type = "85D5E45D-237C-11E1-B4B3-E89A8F7FC3A7"
	MidnightBSDVinum         Type = "85D5E45C-237C-11E1-B4B3-E89A8F7FC3A7"
	CephJournal              Type = "45B0969E-9B03-4F30-B4C6-B4B80CEFF106"
	CephEncryptedJournal     Type = "45B0969E-9B03-4F30-B4C6-5EC00CEFF106"
	CephOSD                  Type = "4FBD7E29-9D25-41B8-AFD0-062C0CEFF05D"
	CephCryptOSD             Type = "4FBD7E29-9D25-41B8-AFD0-5EC00CEFF05D"
	CephDiskInCreation       Type = "89C57F98-2FE5-4DC0-89C1-F3AD0CEFF2BE"
	CephCryptDiskInCreation  Type = "89C57F98-2FE5-4DC0-89C1-5EC00CEFF2BE"
	VMwareVMFS               Type = "AA31E02A-400F-11DB-9590-000C2911D1B8"
	VMwareDiagnostic         Type = "9D275380-40AD-11DB-BF97-000C2911D1B8"
	VMwareVirtualSAN         Type = "381CFCCC-7288-11E0-92EE-000C2911D0B2"
	VMwareVirsto             Type = "77719A0C-A4A0-11E3-A47E-000C29745A24"
	VMwareReserved           Type = "9198EFFC-31C0-11DB-8F78-000C2911D1B8"
	OpenBSDData              Type = "824CC7A0-36A8-11E3-890A-952519AD3F61"
	QNX6FileSystem           Type = "CEF5A9AD-73BC-4601-89F3-CDEEEEE321A1"
	Plan9Partition           Type = "C91818F9-8025-47AF-89D2-F030D7000C2C"
	HiFiveUnleashedFSBL      Type = "5B193300-FC78-40CD-8002-E86C45580B47"
	HiFiveUnleashedBBL       Type = "2E54B353-1271-4842-806F-E436D6AF6985"
)