package oui // generated code - do not edit

type ouiDB struct {
	ouis    map[string]uint
	vendors map[uint]string
}

var database = ouiDB{
	ouis: map[string]uint{
		"000000": 0,    // Xerox
		"000001": 0,    // Xerox
		"000002": 0,    // Xerox
		"000003": 0,    // Xerox
		"000004": 0,    // Xerox
		"000005": 0,    // Xerox
		"000006": 0,    // Xerox
		"000007": 0,    // Xerox
		"000008": 0,    // Xerox
		"000009": 0,    // Xerox
		"00000b": 1,    // Matrix
		"00000c": 2,    // Cisco
		"00000d": 3,    // Fibronics
		"00000e": 4,    // Fujitsu
		"00000f": 5,    // Next
		"000010": 6,    // Sytek
		"000012": 7,    // Information
		"000013": 8,    // Camex
		"000014": 9,    // Netronix
		"000015": 10,   // Datapoint
		"000017": 11,   // Oracle
		"000018": 12,   // Webster
		"00001a": 13,   // AMD
		"00001b": 14,   // Novell
		"00001c": 15,   // Bell
		"00001d": 16,   // Cabletron
		"00001e": 17,   // Telsist
		"00001f": 18,   // Telco
		"000021": 19,   // Sureman
		"000022": 20,   // Visual
		"000023": 21,   // ABB
		"000024": 22,   // Connect
		"000025": 23,   // Ramtek
		"000026": 24,   // SHA-KEN
		"000028": 25,   // Prodigy
		"000029": 26,   // IMC
		"00002c": 27,   // Autotote
		"00002d": 28,   // Chromatics
		"00002f": 29,   // Timeplex
		"000031": 30,   // Qpsx
		"000032": 31,   // Marconi
		"000035": 32,   // Spectragraphics
		"000036": 33,   // Atari
		"000038": 34,   // CSS
		"000039": 35,   // Toshiba
		"00003a": 36,   // Chyron
		"00003c": 37,   // Auspex
		"00003d": 38,   // Unisys
		"00003e": 39,   // Simpact
		"00003f": 40,   // Syntrex
		"000040": 41,   // Applicon
		"000041": 42,   // ICE
		"000043": 43,   // Micro
		"000044": 44,   // Castelle
		"000046": 45,   // Olivetti
		"000048": 46,   // Epson
		"000049": 47,   // Apricot
		"00004c": 48,   // NEC
		"00004d": 49,   // DCI
		"00004e": 50,   // Ampex
		"00004f": 51,   // Logicraft
		"000050": 52,   // Radisys
		"000051": 53,   // HOB
		"000052": 54,   // Intrusion.com
		"000053": 55,   // Compucorp
		"000054": 56,   // Schneider
		"000057": 57,   // Scitex
		"000058": 58,   // Racore
		"000059": 59,   // Hellige
		"00005a": 60,   // SysKonnect
		"00005c": 61,   // Telematics
		"00005e": 62,   // ICANN
		"000060": 63,   // Kontron
		"000061": 64,   // Gateway
		"000066": 65,   // Talaris
		"000069": 66,   // Concord
		"00006c": 67,   // Private
		"00006d": 68,   // Cray
		"00006e": 69,   // Artisoft
		"00006f": 70,   // Madge
		"000070": 71,   // HCL
		"000071": 72,   // Adra
		"000072": 73,   // Miniware
		"000073": 74,   // Siecor
		"000074": 75,   // Ricoh
		"000075": 76,   // Nortel
		"000077": 77,   // Interphase
		"000078": 78,   // Labtam
		"000079": 79,   // Networth
		"00007a": 80,   // Dana
		"00007c": 81,   // Ampere
		"00007d": 11,   // Oracle
		"00007e": 82,   // Clustrix
		"00007f": 83,   // Linotype-Hell
		"000080": 68,   // Cray
		"000081": 84,   // Bay
		"000083": 85,   // Tadpole
		"000084": 86,   // Supernet
		"000085": 87,   // Canon
		"000086": 88,   // Megahertz
		"000087": 89,   // Hitachi
		"000088": 90,   // Brocade
		"000089": 91,   // Cayman
		"00008b": 92,   // Infotron
		"00008c": 93,   // Alloy
		"00008d": 94,   // Cryptek
		"00008e": 95,   // Solbourne
		"00008f": 96,   // Raytheon
		"000090": 97,   // Microcom
		"000091": 98,   // Anritsu
		"000093": 99,   // Proteon
		"000094": 100,  // Asante
		"000095": 101,  // Sony
		"000096": 31,   // Marconi
		"000097": 102,  // Dell
		"000098": 103,  // Crosscomm
		"000099": 104,  // MTX
		"00009a": 105,  // RC
		"00009b": 7,    // Information
		"00009f": 106,  // Ameristar
		"0000a2": 84,   // Bay
		"0000a4": 107,  // Acorn
		"0000a8": 108,  // Stratus
		"0000a9": 109,  // Network
		"0000aa": 0,    // Xerox
		"0000ac": 110,  // Conware
		"0000ae": 111,  // Dassault
		"0000af": 112,  // Canberra
		"0000b2": 113,  // Televideo
		"0000b3": 114,  // Cimlinc
		"0000b4": 115,  // Edimax
		"0000b5": 116,  // Datability
		"0000b6": 117,  // Micro-Matic
		"0000b7": 118,  // Dove
		"0000b8": 119,  // Seikosha
		"0000ba": 120,  // Siig
		"0000bb": 121,  // TRI-Data
		"0000bf": 122,  // Symmetric
		"0000c0": 123,  // WD
		"0000c1": 70,   // Madge
		"0000c3": 124,  // Harris
		"0000c5": 125,  // ARRIS
		"0000c6": 126,  // EON
		"0000c7": 127,  // Arix
		"0000c8": 128,  // Altos
		"0000c9": 129,  // Emulex
		"0000ca": 125,  // ARRIS
		"0000cb": 130,  // Compu-Shack
		"0000cc": 131,  // Densan
		"0000ce": 132,  // Megadata
		"0000d0": 133,  // Develcon
		"0000d1": 134,  // Adaptec
		"0000d2": 135,  // SBE
		"0000d3": 136,  // Wang
		"0000d5": 137,  // Micrognosis
		"0000d8": 14,   // Novell
		"0000da": 138,  // Atex
		"0000dd": 139,  // TCL
		"0000de": 140,  // Cetia
		"0000e0": 141,  // Quadram
		"0000e1": 142,  // Grid
		"0000e2": 143,  // Acer
		"0000e5": 144,  // Sigmex
		"0000e8": 145,  // Accton
		"0000e9": 146,  // Isicad
		"0000ea": 147,  // Upnod
		"0000eb": 148,  // Matsushita
		"0000ec": 149,  // Microprocess
		"0000ed": 150,  // April
		"0000ef": 151,  // KTI
		"0000f0": 152,  // Samsung
		"0000f1": 153,  // Magna
		"0000f2": 154,  // Spider
		"0000f9": 155,  // Quotron
		"0000fa": 156,  // Microsage
		"0000fc": 157,  // Meiko
		"0000ff": 158,  // Camtec
		"000100": 159,  // Equip'Trans
		"000101": 67,   // Private
		"000102": 160,  // 3Com
		"000103": 160,  // 3Com
		"000104": 161,  // DVICO
		"000107": 162,  // Leiser
		"000108": 163,  // AVLAB
		"00010a": 164,  // CIS
		"00010e": 165,  // Bri-Link
		"00010f": 90,   // Brocade
		"000110": 166,  // Gotham
		"000111": 167,  // iDigm
		"000113": 168,  // Olympus
		"000115": 169,  // Extratech
		"000116": 170,  // Netspect
		"000119": 171,  // RTUnet
		"00011b": 172,  // Unizone
		"00011d": 173,  // Centillium
		"00011e": 174,  // Precidia
		"00011f": 105,  // RC
		"000121": 175,  // WatchGuard
		"000122": 176,  // Trend
		"000123": 56,   // Schneider
		"000124": 143,  // Acer
		"000126": 177,  // PAC
		"000127": 178,  // OPEN
		"000128": 179,  // EnjoyWeb
		"000129": 180,  // DFI
		"00012b": 181,  // TELENET
		"00012c": 182,  // Aravox
		"00012d": 183,  // Komodo
		"00012f": 184,  // Twinhead
		"000130": 185,  // Extreme
		"000133": 186,  // KYOWA
		"000134": 187,  // Selectron
		"000135": 188,  // KDC
		"000136": 189,  // CyberTAN
		"000138": 190,  // XAVi
		"00013a": 191,  // Shelcad
		"00013b": 192,  // BNA
		"00013c": 193,  // TIW
		"00013d": 194,  // RiscStation
		"000140": 195,  // Sendtek
		"000142": 2,    // Cisco
		"000143": 2,    // Cisco
		"000144": 102,  // Dell
		"000145": 196,  // Winsystems
		"000147": 197,  // Zhone
		"000148": 198,  // X-traWeb
		"000149": 199,  // TDT
		"00014a": 101,  // Sony
		"00014b": 200,  // Ennovate
		"00014e": 201,  // WIN
		"00014f": 202,  // Adtran
		"000150": 203,  // Gilat
		"000151": 204,  // Ensemble
		"000152": 205,  // Chromatek
		"000154": 206,  // G3M
		"000155": 207,  // Promise
		"000156": 208,  // Firewiredirect.com
		"000157": 209,  // Syswave
		"000158": 210,  // Electro
		"000159": 211,  // S1
		"00015c": 212,  // Cadant
		"00015d": 11,   // Oracle
		"00015e": 213,  // Best
		"000160": 214,  // ELMEX
		"000162": 215,  // Cygnet
		"000163": 2,    // Cisco
		"000164": 2,    // Cisco
		"000165": 216,  // AirSwitch
		"000168": 217,  // Vitana
		"000169": 218,  // Celestix
		"00016a": 219,  // Alitec
		"00016b": 220,  // LightChip
		"00016c": 221,  // Foxconn
		"00016d": 222,  // CarrierComm
		"00016e": 223,  // Conklin
		"00016f": 224,  // Inkel
		"000172": 225,  // TechnoLand
		"000173": 226,  // AMCC
		"000174": 227,  // CyberOptics
		"000175": 228,  // Radiant
		"000177": 229,  // EDSL
		"000178": 230,  // MARGI
		"000179": 231,  // Wireless
		"00017c": 232,  // AG-E
		"00017d": 233,  // ThermoQuest
		"000180": 234,  // AOpen
		"000181": 76,   // Nortel
		"000182": 235,  // Dica
		"000187": 236,  // I2SE
		"000188": 237,  // LXCO
		"000189": 238,  // Refraction
		"00018a": 239,  // ROI
		"00018b": 240,  // NetLinks
		"00018d": 241,  // AudeSi
		"00018e": 242,  // Logitec
		"00018f": 243,  // Kenetec
		"000190": 244,  // SMK-M
		"000195": 245,  // Sena
		"000196": 2,    // Cisco
		"000197": 2,    // Cisco
		"000199": 246,  // HeiSei
		"00019a": 247,  // LEUNIG
		"00019d": 248,  // E-Control
		"00019e": 249,  // ESS
		"00019f": 250,  // ReadyNet
		"0001a0": 251,  // Infinilink
		"0001a1": 252,  // Mag-Tek
		"0001a2": 253,  // Logical
		"0001a4": 254,  // Microlink
		"0001a5": 255,  // Nextcomm
		"0001a7": 256,  // Unex
		"0001a8": 257,  // Welltech
		"0001a9": 258,  // BMW
		"0001aa": 259,  // Airspan
		"0001ac": 260,  // Sitara
		"0001ae": 261,  // Trex
		"0001b0": 262,  // Fulltek
		"0001b3": 263,  // Precision
		"0001b4": 264,  // Wayport
		"0001b5": 265,  // Turin
		"0001b7": 266,  // Centos
		"0001b8": 267,  // Netsensity
		"0001b9": 268,  // SKF
		"0001ba": 269,  // IC-Net
		"0001bb": 270,  // Frequentis
		"0001bc": 271,  // Brains
		"0001bd": 272,  // Peterson
		"0001be": 273,  // Gigalink
		"0001bf": 274,  // Teleforce
		"0001c0": 275,  // CompuLab
		"0001c2": 276,  // ARK
		"0001c3": 277,  // Acromag
		"0001c4": 278,  // NeoWave
		"0001c5": 279,  // Simpler
		"0001c6": 280,  // Quarry
		"0001c7": 2,    // Cisco
		"0001c8": 281,  // Conrad
		"0001c9": 2,    // Cisco
		"0001cb": 282,  // EVR
		"0001cd": 283,  // ARtem
		"0001d0": 284,  // VitalPoint
		"0001d1": 285,  // CoNet
		"0001d2": 286,  // inXtron
		"0001d3": 287,  // PAXCOMM
		"0001d6": 288,  // manroland
		"0001d7": 289,  // F5
		"0001d8": 290,  // Teltronics
		"0001d9": 291,  // Sigma
		"0001da": 292,  // WINCOMM
		"0001db": 293,  // Freecom
		"0001dc": 294,  // Activetelco
		"0001dd": 295,  // Avail
		"0001de": 296,  // Trango
		"0001df": 297,  // ISDN
		"0001e0": 298,  // Fast
		"0001e1": 299,  // Kinpo
		"0001e3": 300,  // Siemens
		"0001e4": 301,  // Sitera
		"0001e5": 86,   // Supernet
		"0001e6": 302,  // HP
		"0001e7": 302,  // HP
		"0001e8": 303,  // Force10
		"0001ea": 304,  // Cirilium
		"0001eb": 305,  // C-CoM
		"0001ec": 306,  // Ericsson
		"0001ed": 307,  // SETA
		"0001ee": 308,  // Comtrol
		"0001ef": 309,  // Camtel
		"0001f0": 310,  // Tridium
		"0001f3": 311,  // QPS
		"0001f4": 312,  // Enterasys
		"0001f8": 313,  // Texio
		"0001f9": 314,  // TeraGlobal
		"0001fa": 315,  // Horoscas
		"0001fb": 316,  // DoTop
		"0001fc": 317,  // Keyence
		"000200": 318,  // Net
		"000201": 319,  // IFM
		"000202": 320,  // Amino
		"000204": 321,  // Bodmann
		"000208": 322,  // Unify
		"00020a": 323,  // Gefran
		"00020b": 324,  // Native
		"00020c": 325,  // Metro-Optix
		"00020d": 326,  // Micronpc.com
		"00020f": 327,  // Aatr
		"000210": 328,  // Fenecom
		"000212": 329,  // SierraCom
		"000213": 330,  // S.D.E.L
		"000214": 331,  // Dtvro
		"000216": 2,    // Cisco
		"000217": 2,    // Cisco
		"000219": 332,  // Paralon
		"00021a": 333,  // Zuma
		"00021b": 334,  // Kollmorgen-Servotronix
		"00021f": 335,  // Aculab
		"000222": 336,  // Chromisys
		"000223": 337,  // ClickTV
		"000224": 338,  // C-CoR
		"000226": 339,  // XESystems
		"000227": 340,  // ESD
		"000228": 341,  // Necsom
		"000229": 342,  // Adtec
		"00022a": 343,  // Asound
		"00022b": 344,  // SAXA
		"00022d": 345,  // Agere
		"00022f": 346,  // P-Cube
		"000230": 347,  // Intersoft
		"000231": 348,  // Ingersoll-Rand
		"000232": 349,  // Avision
		"000233": 350,  // Mantra
		"000234": 351,  // Imperial
		"000235": 352,  // Paragon
		"000236": 353,  // INIT
		"000237": 354,  // Cosmo
		"000238": 355,  // Serome
		"000239": 356,  // Visicom
		"00023b": 306,  // Ericsson
		"00023c": 357,  // Creative
		"00023d": 2,    // Cisco
		"00023f": 358,  // Compal
		"000240": 359,  // Seedek
		"000241": 360,  // Amer.com
		"000242": 361,  // Videoframe
		"000243": 362,  // Raysis
		"000244": 363,  // SURECOM
		"000245": 364,  // Lampus
		"000246": 365,  // All-Win
		"000248": 366,  // Pilz
		"00024a": 2,    // Cisco
		"00024b": 2,    // Cisco
		"00024c": 367,  // SiByte
		"000250": 368,  // Geyser
		"000251": 369,  // Soma
		"000252": 370,  // Carrier
		"000253": 113,  // Televideo
		"000254": 371,  // WorldGate
		"000255": 372,  // IBM
		"000257": 97,   // Microcom
		"00025a": 373,  // Catena
		"00025d": 374,  // Calix
		"00025e": 375,  // High
		"00025f": 76,   // Nortel
		"000260": 376,  // Accordion
		"000261": 377,  // Tilgin
		"000264": 378,  // AudioRamp.com
		"000265": 379,  // Virditech
		"000266": 380,  // Thermalogic
		"000269": 381,  // Nadatel
		"00026b": 382,  // BCM
		"00026f": 383,  // Senao
		"000270": 384,  // Crewave
		"000271": 197,  // Zhone
		"000272": 385,  // CC&C
		"000273": 386,  // Coriolis
		"000274": 387,  // Tommy
		"000275": 388,  // SMART
		"000276": 389,  // Primax
		"000278": 152,  // Samsung
		"00027a": 390,  // IOI
		"00027c": 391,  // Trilithic
		"00027d": 2,    // Cisco
		"00027e": 2,    // Cisco
		"00027f": 392,  // ask-technologies.com
		"000281": 70,   // Madge
		"000282": 393,  // ViaClix
		"000285": 394,  // Riverstone
		"000286": 395,  // Occam
		"000287": 396,  // Adapcom
		"000289": 397,  // DNE
		"00028b": 398,  // VDSL
		"00028d": 399,  // Movita
		"00028f": 400,  // Globetek
		"000290": 401,  // Woorigisool
		"000295": 402,  // IP.Access
		"000296": 403,  // Lectron
		"000297": 404,  // C-CoR.net
		"000298": 405,  // Broadframe
		"000299": 406,  // Apex
		"00029b": 407,  // Kreatel
		"00029c": 160,  // 3Com
		"00029d": 408,  // Merix
		"0002a0": 409,  // Flatstack
		"0002a2": 410,  // Hilscher
		"0002a4": 411,  // AddPac
		"0002a5": 302,  // HP
		"0002a6": 412,  // Effinet
		"0002a7": 413,  // Vivace
		"0002a9": 414,  // RACOM
		"0002aa": 415,  // PLcom
		"0002ad": 416,  // HOYA
		"0002ae": 417,  // Scannex
		"0002af": 418,  // TeleCruz
		"0002b0": 419,  // Hokubu
		"0002b1": 98,   // Anritsu
		"0002b2": 420,  // Cablevision
		"0002b3": 421,  // Intel
		"0002b4": 422,  // Daphne
		"0002b5": 423,  // Avnet
		"0002b6": 424,  // Acrosser
		"0002b9": 2,    // Cisco
		"0002ba": 2,    // Cisco
		"0002bd": 425,  // Bionet
		"0002be": 426,  // Totsu
		"0002bf": 427,  // dotRocket
		"0002c1": 428,  // Innovative
		"0002c3": 429,  // Arelnet
		"0002c7": 430,  // Alpsalpine
		"0002c8": 431,  // Technocom
		"0002c9": 432,  // Mellanox
		"0002ca": 433,  // EndPoints
		"0002cb": 434,  // TriState
		"0002cc": 435,  // M.C.C.I
		"0002cd": 436,  // TeleDream
		"0002ce": 437,  // FoxJet
		"0002cf": 438,  // ZyGate
		"0002d0": 439,  // Comdial
		"0002d1": 440,  // Vivotek
		"0002d2": 441,  // Workstation
		"0002d3": 442,  // NetBotz
		"0002d5": 443,  // ACR
		"0002d6": 444,  // NICE
		"0002d7": 445,  // EMPEG
		"0002d8": 446,  // BRECIS
		"0002da": 447,  // ExiO
		"0002db": 448,  // Netsec
		"0002dd": 449,  // Bromax
		"0002de": 450,  // Astrodesign
		"0002e0": 451,  // ETAS
		"0002e3": 452,  // LITE-ON
		"0002e5": 453,  // Timeware
		"0002e7": 454,  // CAB
		"0002e8": 455,  // E.D.&A
		"0002eb": 456,  // Pico
		"0002ee": 457,  // Nokia
		"0002f1": 458,  // Pinetron
		"0002f2": 459,  // eDevice
		"0002f4": 460,  // PCTEL
		"0002f6": 461,  // Equipe
		"0002f7": 462,  // ARM
		"0002f8": 463,  // SEAKR
		"0002fc": 2,    // Cisco
		"0002fd": 2,    // Cisco
		"0002fe": 464,  // Viditec
		"000300": 465,  // Barracuda
		"000301": 466,  // Exfo
		"000302": 467,  // Charles
		"000303": 468,  // JAMA
		"000308": 469,  // AM
		"000309": 470,  // Texcel
		"00030a": 471,  // Argus
		"00030b": 472,  // Hunter
		"00030c": 473,  // Telesoft
		"00030d": 474,  // Uniwill
		"00030e": 475,  // Core
		"000310": 476,  // E-Globaledge
		"000311": 43,   // Micro
		"000312": 477,  // TRsystems
		"000315": 478,  // Cidco
		"000316": 479,  // Nobell
		"000317": 480,  // Merlin
		"000318": 481,  // Cyras
		"000319": 482,  // Infineon
		"00031b": 483,  // Cellvision
		"00031e": 484,  // Optranet
		"00031f": 485,  // Condev
		"000320": 486,  // Xpeed
		"000321": 487,  // Reco
		"000322": 488,  // IDIS
		"000323": 489,  // Cornet
		"000325": 490,  // Arima
		"000327": 491,  // ACT'L
		"000329": 492,  // F3
		"00032a": 493,  // UniData
		"00032d": 494,  // IBASE
		"000330": 495,  // Imagenics
		"000331": 2,    // Cisco
		"000332": 2,    // Cisco
		"000333": 496,  // Digitel
		"000334": 497,  // Omega
		"000335": 498,  // Mirae
		"000336": 499,  // Zetes
		"000337": 500,  // Vaone
		"000338": 501,  // Oak
		"000339": 502,  // Eurologic
		"00033b": 503,  // TAMI
		"00033c": 504,  // Daiden
		"00033f": 505,  // BigBand
		"000342": 76,   // Nortel
		"000344": 506,  // Tietech
		"000345": 507,  // Routrek
		"000347": 421,  // Intel
		"00034a": 508,  // RIAS
		"00034b": 76,   // Nortel
		"00034d": 509,  // Chiaro
		"000350": 510,  // Bticino
		"000351": 511,  // Diebold
		"000352": 512,  // Colubris
		"000353": 513,  // Mitac
		"000357": 514,  // Intervoice-Brite
		"000359": 515,  // DigitalSis
		"00035a": 516,  // Photron
		"00035b": 517,  // BridgeWave
		"000361": 518,  // Widcomm
		"000362": 519,  // Vodtel
		"000363": 520,  // Miraesys
		"000367": 521,  // Jasmine
		"000368": 522,  // Embedone
		"00036a": 523,  // Mainnet
		"00036b": 2,    // Cisco
		"00036c": 2,    // Cisco
		"00036d": 524,  // Runtop
		"00036f": 525,  // Telsey
		"000370": 526,  // NXTV
		"000371": 527,  // Acomz
		"000372": 528,  // Ulan
		"000375": 529,  // NetMedia
		"000376": 530,  // Graphtec
		"000378": 531,  // HUMAX
		"000379": 532,  // Proscend
		"00037d": 533,  // Stellcom
		"00037e": 534,  // PORTech
		"00037f": 535,  // Atheros
		"000381": 536,  // Ingenico
		"000382": 537,  // A-One
		"000383": 538,  // Metera
		"000384": 539,  // Aeta
		"000385": 540,  // Actelis
		"000388": 541,  // Fastfame
		"000389": 542,  // Plantronics
		"00038e": 543,  // Atoga
		"00038f": 544,  // Weinschel
		"000393": 545,  // Apple
		"000397": 546,  // FireBrick
		"000398": 547,  // Wisi
		"00039a": 548,  // SiConnect
		"00039b": 549,  // NetChip
		"00039c": 550,  // OptiMight
		"00039d": 551,  // Qisda
		"00039f": 2,    // Cisco
		"0003a0": 2,    // Cisco
		"0003a2": 552,  // Catapult
		"0003a3": 553,  // MAVIX
		"0003a4": 554,  // Imation
		"0003a5": 555,  // Medea
		"0003a6": 556,  // Traxit
		"0003a7": 557,  // Unixtar
		"0003a8": 558,  // IDOT
		"0003aa": 559,  // Watlow
		"0003af": 560,  // Paragea
		"0003b0": 561,  // Xsense
		"0003b1": 562,  // Hospira
		"0003b2": 563,  // Radware
		"0003b4": 564,  // Macrotek
		"0003b5": 565,  // Entra
		"0003b6": 566,  // QSI
		"0003b7": 567,  // ZACCESS
		"0003b8": 568,  // NetKit
		"0003ba": 11,   // Oracle
		"0003bb": 569,  // Signal
		"0003bc": 570,  // CoT
		"0003bd": 571,  // OmniCluster
		"0003be": 572,  // Netility
		"0003c0": 573,  // RFTNC
		"0003c5": 574,  // Mobotix
		"0003c6": 575,  // ICUE
		"0003c9": 576,  // TECOM
		"0003ca": 577,  // MTS
		"0003cb": 578,  // SystemGear
		"0003cc": 579,  // Momentum
		"0003cd": 580,  // Clovertech
		"0003ce": 581,  // ETEN
		"0003cf": 582,  // Muxcom
		"0003d0": 583,  // KOANKEISO
		"0003d1": 584,  // Takaya
		"0003d2": 585,  // Crossbeam
		"0003d4": 586,  // Alloptic
		"0003d5": 587,  // Advanced
		"0003d6": 588,  // RADVision
		"0003d8": 589,  // iMPath
		"0003d9": 590,  // Secheron
		"0003db": 591,  // Apogee
		"0003df": 592,  // Desana
		"0003e0": 125,  // ARRIS
		"0003e1": 593,  // Winmate
		"0003e2": 594,  // Comspace
		"0003e3": 2,    // Cisco
		"0003e4": 2,    // Cisco
		"0003e5": 595,  // Hermstedt
		"0003e6": 596,  // Entone
		"0003e7": 597,  // Logostek
		"0003e8": 598,  // Wavesight
		"0003eb": 599,  // Atrica
		"0003ec": 600,  // ICG
		"0003ee": 601,  // MKNet
		"0003ef": 602,  // Oneline
		"0003f2": 603,  // Seneca
		"0003f4": 604,  // NetBurner
		"0003f5": 605,  // Chip2Chip
		"0003f6": 606,  // Allegro
		"0003f7": 607,  // Plast-Control
		"0003f8": 608,  // SanCastle
		"0003f9": 609,  // Pleiades
		"0003fa": 610,  // TiMetra
		"0003fb": 611,  // ENEGATE
		"0003fd": 2,    // Cisco
		"0003fe": 2,    // Cisco
		"0003ff": 612,  // Microsoft
		"000400": 613,  // Lexmark
		"000402": 614,  // Nexsan
		"000403": 615,  // Nexsi
		"000405": 616,  // ACN
		"000408": 617,  // Sanko
		"000409": 618,  // Cratos
		"00040a": 619,  // Sage
		"00040b": 160,  // 3Com
		"00040d": 620,  // Avaya
		"00040e": 621,  // AVM
		"000410": 622,  // Spinnaker
		"000411": 623,  // Inkra
		"000412": 624,  // WaveSmith
		"000413": 625,  // snom
		"000415": 626,  // Rasteme
		"000417": 627,  // Elau
		"000419": 628,  // Fibercycle
		"00041b": 629,  // Bridgeworks
		"00041c": 630,  // ipDialog
		"00041f": 101,  // Sony
		"000421": 631,  // Ocular
		"000422": 632,  // Studio
		"000423": 421,  // Intel
		"000425": 633,  // Atmel
		"000426": 634,  // Autosys
		"000427": 2,    // Cisco
		"000428": 2,    // Cisco
		"000429": 635,  // Pixord
		"00042a": 231,  // Wireless
		"00042c": 636,  // Minet
		"00042d": 637,  // Sarian
		"00042e": 638,  // Netous
		"00042f": 639,  // International
		"000430": 640,  // Netgem
		"000431": 641,  // GlobalStreams
		"000433": 642,  // Cyberboard
		"000434": 643,  // Accelent
		"000435": 644,  // InfiNet
		"000436": 645,  // ELANsat
		"000438": 76,   // Nortel
		"000439": 646,  // Rosco
		"00043b": 647,  // Lava
		"00043c": 648,  // SONOS
		"00043d": 649,  // Indel
		"00043e": 650,  // Telencomm
		"000440": 651,  // cyberPIXIE
		"000442": 652,  // Nact
		"000443": 653,  // Agilent
		"000446": 654,  // CYZENTECH
		"000447": 655,  // Acrowave
		"000448": 656,  // Polaroid
		"000449": 657,  // Mapletree
		"00044a": 658,  // iPolicy
		"00044b": 659,  // Nvidia
		"00044c": 660,  // Jenoptik
		"00044d": 2,    // Cisco
		"00044e": 2,    // Cisco
		"000451": 661,  // Medrad
		"000452": 662,  // RocketLogix
		"000453": 663,  // YottaYotta
		"000455": 664,  // ANTARA.net
		"000456": 665,  // Cambium
		"000459": 666,  // Veristar
		"00045b": 667,  // Techsan
		"00045c": 668,  // Mobiwave
		"00045f": 669,  // Avalue
		"000460": 670,  // Knilink
		"000461": 671,  // EPOX
		"000464": 672,  // Pulse-Link
		"000466": 673,  // ARMITEL
		"000468": 674,  // Vivity
		"000469": 675,  // Innocom
		"00046a": 676,  // Navini
		"00046c": 677,  // Cyber
		"00046d": 2,    // Cisco
		"00046e": 2,    // Cisco
		"000470": 678,  // ipUnplugged
		"000471": 679,  // IPrad
		"000472": 680,  // Telelynx
		"000473": 681,  // Photonex
		"000474": 682,  // Legrand
		"000475": 160,  // 3Com
		"000476": 160,  // 3Com
		"000477": 683,  // Scalant
		"000479": 684,  // Radius
		"00047b": 685,  // Schlumberger
		"00047c": 686,  // Skidata
		"00047d": 687,  // Motorola
		"000480": 90,   // Brocade
		"000482": 688,  // Medialogic
		"000483": 689,  // Deltron
		"000484": 690,  // Amann
		"000485": 691,  // PicoLight
		"000486": 692,  // ITTC
		"000489": 693,  // YAFO
		"00048b": 694,  // Poscon
		"00048c": 695,  // Nayna
		"00048d": 696,  // Teo
		"00048e": 697,  // Ohm
		"00048f": 698,  // TD
		"000491": 699,  // Technovision
		"000494": 700,  // Breezecom
		"000496": 185,  // Extreme
		"000498": 701,  // Mahi
		"000499": 702,  // Chino
		"00049a": 2,    // Cisco
		"00049b": 2,    // Cisco
		"00049c": 703,  // Surgient
		"00049d": 704,  // Ipanema
		"00049e": 705,  // Wirelink
		"0004a3": 706,  // Microchip
		"0004a4": 707,  // NetEnabled
		"0004a7": 708,  // FabiaTech
		"0004a8": 709,  // Broadmax
		"0004a9": 710,  // SandStream
		"0004aa": 711,  // Jetstream
		"0004ab": 712,  // Mavenir
		"0004ac": 372,  // IBM
		"0004ad": 713,  // Malibu
		"0004ae": 714,  // Sullair
		"0004b0": 715,  // ELESIGN
		"0004b1": 569,  // Signal
		"0004b3": 716,  // Videotek
		"0004b4": 717,  // Ciac
		"0004b5": 718,  // Equitrac
		"0004b6": 719,  // Stratex
		"0004b8": 720,  // Kumahira
		"0004bb": 721,  // Bardac
		"0004bc": 722,  // Giantec
		"0004bd": 125,  // ARRIS
		"0004be": 723,  // OptXCon
		"0004bf": 724,  // VersaLogic
		"0004c0": 2,    // Cisco
		"0004c1": 2,    // Cisco
		"0004c2": 725,  // Magnipix
		"0004c5": 726,  // ASE
		"0004c6": 727,  // Yamaha
		"0004c7": 728,  // NetMount
		"0004c9": 43,   // Micro
		"0004ca": 729,  // FreeMs
		"0004cb": 730,  // Tdsoft
		"0004cd": 731,  // Extenway
		"0004cf": 732,  // Seagate
		"0004d1": 733,  // Drew
		"0004d3": 734,  // Toyokeiki
		"0004d4": 735,  // Proview
		"0004d6": 736,  // Takagi
		"0004d8": 737,  // IPWireless
		"0004d9": 738,  // Titan
		"0004da": 739,  // Relax
		"0004dc": 76,   // Nortel
		"0004dd": 2,    // Cisco
		"0004de": 2,    // Cisco
		"0004e0": 740,  // Procket
		"0004e2": 741,  // SMC
		"0004e3": 145,  // Accton
		"0004e4": 742,  // Daeryung
		"0004e5": 743,  // Glonet
		"0004e7": 744,  // Lightpointe
		"0004e8": 745,  // IER
		"0004e9": 746,  // Infiniswitch
		"0004ea": 302,  // HP
		"0004eb": 747,  // Paxonet
		"0004ec": 748,  // Memobox
		"0004ef": 749,  // Polestar
		"0004f0": 639,  // International
		"0004f1": 750,  // WhereNet
		"0004f2": 751,  // Polycom
		"0004f4": 752,  // Infinite
		"0004f5": 753,  // SnowShore
		"0004f6": 754,  // Amphus
		"0004f9": 755,  // Xtera
		"0004fa": 756,  // NBS
		"0004fb": 757,  // Commtech
		"0004fc": 108,  // Stratus
		"0004fe": 758,  // Pelago
		"0004ff": 759,  // Acronet
		"000500": 2,    // Cisco
		"000501": 2,    // Cisco
		"000502": 545,  // Apple
		"000503": 760,  // Iconag
		"000506": 761,  // Reddo
		"000508": 762,  // Inetcam
		"00050a": 763,  // ICS
		"00050b": 764,  // SICOM
		"00050d": 765,  // Midstream
		"00050e": 766,  // 3ware
		"000511": 767,  // Complementary
		"000512": 768,  // Zebra
		"000514": 769,  // KDT
		"000515": 770,  // Nuark
		"000517": 771,  // Shellcomm
		"000518": 772,  // Jupiters
		"000519": 300,  // Siemens
		"00051a": 160,  // 3Com
		"00051c": 773,  // Xnet
		"00051d": 774,  // Airocon
		"00051e": 90,   // Brocade
		"000520": 775,  // Smartronix
		"000522": 776,  // LEA*D
		"000525": 777,  // Puretek
		"000526": 778,  // IPAS
		"00052b": 779,  // HORIBA
		"00052d": 780,  // Zoltrix
		"00052e": 781,  // Cinta
		"000530": 782,  // Andiamo
		"000531": 2,    // Cisco
		"000532": 2,    // Cisco
		"000533": 90,   // Brocade
		"000534": 783,  // Northstar
		"000536": 784,  // Danam
		"000537": 785,  // Nets
		"000538": 786,  // Merilus
		"00053a": 787,  // Willowglen
		"00053c": 788,  // Xircom
		"00053d": 345,  // Agere
		"00053f": 789,  // VisionTek
		"000540": 790,  // FAST
		"000541": 587,  // Advanced
		"000542": 791,  // Otari
		"000544": 792,  // Valley
		"000547": 793,  // Starent
		"000548": 794,  // Disco
		"00054d": 795,  // Brans
		"00054e": 796,  // Philips
		"00054f": 797,  // Garmin
		"000552": 798,  // Xycotec
		"000553": 799,  // DVC
		"000556": 800,  // 360
		"000558": 801,  // Synchronous
		"00055b": 467,  // Charles
		"00055c": 802,  // Kowa
		"00055d": 803,  // D-Link
		"00055e": 2,    // Cisco
		"00055f": 2,    // Cisco
		"000560": 804,  // Leader
		"000563": 805,  // J-Works
		"000565": 806,  // Tailyn
		"000566": 807,  // Secui.com
		"000568": 808,  // Piltofish
		"000569": 809,  // VMware
		"00056b": 810,  // C.P
		"00056d": 811,  // Pacific
		"00056f": 812,  // Innomedia
		"000570": 813,  // Baydel
		"000571": 814,  // Seiwa
		"000572": 815,  // Deonet
		"000573": 2,    // Cisco
		"000574": 2,    // Cisco
		"000575": 816,  // CDS-Electronics
		"000576": 817,  // NSM
		"000578": 67,   // Private
		"00057a": 818,  // Overture
		"00057d": 819,  // Sun
		"00057f": 820,  // Acqis
		"000580": 821,  // FibroLAN
		"000581": 822,  // Snell
		"000582": 823,  // ClearCube
		"000583": 824,  // ImageCom
		"000584": 825,  // AbsoluteValue
		"000585": 826,  // Juniper
		"000586": 827,  // Lucent
		"000587": 828,  // Locusorporated
		"000588": 829,  // Sensoria
		"00058a": 830,  // Netcom
		"00058b": 831,  // IPmental
		"00058c": 832,  // Opentech
		"00058e": 833,  // Flextronics
		"00058f": 834,  // CLCsoft
		"000590": 835,  // Swissvoice
		"000592": 836,  // Pultek
		"000594": 837,  // HMS
		"000595": 838,  // Alesis
		"000596": 839,  // Genotech
		"00059a": 2,    // Cisco
		"00059b": 2,    // Cisco
		"00059c": 840,  // Kleinknecht
		"00059e": 841,  // Zinwell
		"00059f": 842,  // Yotta
		"0005a1": 843,  // Zenocom
		"0005a2": 844,  // CELOX
		"0005a3": 845,  // QEI
		"0005a5": 846,  // Kott
		"0005a6": 847,  // Extron
		"0005a7": 848,  // HYPERCHIP
		"0005a8": 849,  // Wyle
		"0005a9": 850,  // Princeton
		"0005aa": 851,  // Moore
		"0005ad": 852,  // Topspin
		"0005b0": 853,  // Korea
		"0005b1": 854,  // ASB
		"0005b2": 855,  // Medison
		"0005b3": 856,  // Asahi-Engineering
		"0005b4": 857,  // Aceex
		"0005b5": 858,  // Broadcom
		"0005b7": 859,  // Arbor
		"0005b9": 860,  // Airvana
		"0005bb": 861,  // Myspace
		"0005bd": 862,  // Roax
		"0005bf": 863,  // JustEzy
		"0005c2": 864,  // Soronti
		"0005c4": 865,  // Telect
		"0005c6": 866,  // Triz
		"0005c7": 867,  // I/F-CoM
		"0005c8": 868,  // Verytech
		"0005c9": 869,  // LG
		"0005ca": 870,  // Hitron
		"0005cb": 871,  // ROIS
		"0005cc": 872,  // Sumtel
		"0005cd": 873,  // D&M
		"0005d0": 874,  // Solinet
		"0005d1": 875,  // Metavector
		"0005d2": 876,  // DAP
		"0005d3": 877,  // eProduction
		"0005d4": 878,  // FutureSmart
		"0005d8": 879,  // Arescom
		"0005dc": 2,    // Cisco
		"0005dd": 2,    // Cisco
		"0005e0": 880,  // Empirix
		"0005e3": 881,  // LightSand
		"0005e5": 882,  // Renishaw
		"0005e6": 883,  // Egenera
		"0005e8": 884,  // TurboWave
		"0005ea": 885,  // Rednix
		"0005ec": 886,  // Mosaic
		"0005ee": 887,  // Vanderbilt
		"0005f0": 888,  // Satec
		"0005f1": 889,  // Vrcom
		"0005f3": 890,  // Webyn
		"0005f5": 891,  // Geospace
		"0005f9": 892,  // TOA
		"0005fa": 893,  // IPOptical
		"0005fb": 894,  // ShareGate
		"0005fd": 895,  // PacketLight
		"0005ff": 896,  // SNS
		"000600": 35,   // Toshiba
		"000601": 897,  // Otanikeiki
		"000602": 898,  // Cirkitech
		"000604": 899,  // @Track
		"000605": 900,  // Inncom
		"000606": 901,  // RapidWAN
		"000609": 902,  // Crossport
		"00060a": 903,  // Blue2space
		"00060c": 904,  // Melco
		"00060e": 905,  // IGYS
		"00060f": 906,  // Narad
		"000610": 907,  // Abeona
		"000612": 908,  // Accusys
		"000614": 909,  // Prism
		"000617": 910,  // Redswitch
		"000618": 911,  // DigiPower
		"000619": 912,  // Connection
		"00061a": 913,  // Zetari
		"00061e": 914,  // Maxan
		"000621": 915,  // Hinox
		"000624": 916,  // Gentner
		"000626": 917,  // MWE
		"000627": 918,  // Uniwide
		"000628": 2,    // Cisco
		"000629": 372,  // IBM
		"00062a": 2,    // Cisco
		"00062b": 919,  // Intraserver
		"00062c": 920,  // Bivio
		"00062d": 921,  // TouchStar
		"00062f": 922,  // Pivotech
		"000631": 374,  // Calix
		"000632": 923,  // Mesco
		"000635": 924,  // PacketAir
		"000639": 925,  // Newtec
		"00063b": 926,  // Arcturus
		"00063c": 927,  // Intrinsyc
		"00063e": 928,  // Opthos
		"00063f": 929,  // Everex
		"000641": 930,  // Itcn
		"000642": 931,  // Genetel
		"000643": 932,  // SONO
		"000648": 933,  // Seedsware
		"00064a": 934,  // Honeywell
		"00064b": 935,  // Alexon
		"00064c": 936,  // Invicta
		"00064d": 937,  // Sencore
		"00064f": 938,  // PRO-NETS
		"000650": 939,  // Tiburon
		"000651": 940,  // Aspen
		"000652": 2,    // Cisco
		"000653": 2,    // Cisco
		"000654": 941,  // Winpresa
		"000655": 942,  // Yipee
		"000656": 943,  // Tactel
		"00065a": 944,  // Strix
		"00065b": 102,  // Dell
		"00065c": 945,  // Malachite
		"00065e": 946,  // Photuris
		"000660": 947,  // NADEX
		"000662": 948,  // MBM
		"000663": 949,  // Human
		"000664": 950,  // Fostex
		"000666": 951,  // Roving
		"000668": 952,  // Vicon
		"000669": 953,  // Datasound
		"00066a": 954,  // InfiniCon
		"00066b": 955,  // Sysmex
		"00066c": 956,  // Robinson
		"00066e": 957,  // Delta
		"000670": 958,  // Upponetti
		"000671": 959,  // Softing
		"000672": 960,  // Netezza
		"000675": 961,  // Banderacom
		"000676": 962,  // Novra
		"000677": 963,  // Sick
		"000678": 873,  // D&M
		"000679": 964,  // Konami
		"00067a": 965,  // JMP
		"00067c": 2,    // Cisco
		"00067d": 966,  // Takasago
		"00067e": 967,  // WinCom
		"00067f": 968,  // Digeo
		"000681": 969,  // Goepel
		"000682": 970,  // Convedia
		"000683": 971,  // Bravara
		"000684": 972,  // Biacore
		"000685": 973,  // NetNearU
		"000686": 974,  // ZARDCOM
		"000687": 975,  // Omnitron
		"000688": 976,  // Telways
		"000689": 977,  // yLez
		"00068b": 978,  // AirRunner
		"00068c": 160,  // 3Com
		"00068d": 979,  // SEPATON
		"00068e": 980,  // HID
		"00068f": 981,  // Telemonitor
		"000690": 982,  // Euracom
		"000692": 983,  // Intruvert
		"000693": 984,  // Flexus
		"000694": 985,  // Mobillian
		"000695": 986,  // Ensure
		"000696": 987,  // Advent
		"000698": 988,  // egnite
		"00069c": 989,  // Transmode
		"00069d": 990,  // Petards
		"00069e": 991,  // UNIQA
		"00069f": 992,  // Kuokoa
		"0006a1": 993,  // Celsian
		"0006a2": 994,  // Microtune
		"0006a3": 995,  // Bitran
		"0006a4": 996,  // INNOWELL
		"0006a5": 997,  // PINON
		"0006a7": 998,  // Primarion
		"0006a8": 999,  // KC
		"0006ab": 1000, // W-Link
		"0006ac": 347,  // Intersoft
		"0006ad": 1001, // KB
		"0006af": 1002, // Xalted
		"0006b1": 1003, // Sonicwall
		"0006b2": 1004, // Linxtek
		"0006b3": 1005, // Diagraph
		"0006b4": 1006, // Vorne
		"0006b7": 1007, // TELEM
		"0006b8": 1008, // Bandspeed
		"0006b9": 1009, // A5TEK
		"0006ba": 1010, // Westwave
		"0006bb": 1011, // ATI
		"0006bc": 1012, // Macrolink
		"0006bd": 1013, // BNTECHNOLOGY
		"0006bf": 1014, // Accella
		"0006c1": 2,    // Cisco
		"0006c2": 1015, // Smartmatic
		"0006c4": 1016, // Piolink
		"0006c5": 1017, // INNOVI
		"0006c6": 1018, // lesswire
		"0006c7": 1019, // RFNET
		"0006cb": 1020, // Jotron
		"0006cc": 1021, // JMI
		"0006ce": 1022, // Dateno
		"0006d0": 1023, // Elgar
		"0006d1": 1024, // Tahoe
		"0006d5": 1025, // Diamond
		"0006d6": 2,    // Cisco
		"0006d7": 2,    // Cisco
		"0006da": 1026, // ITRAN
		"0006db": 1027, // ICHIPS
		"0006dc": 1028, // Syabas
		"0006de": 1029, // Flash
		"0006df": 1030, // AIDONIC
		"0006e0": 1031, // MAT
		"0006e2": 1032, // Ceemax
		"0006e4": 1033, // Citel
		"0006e9": 1034, // Intime
		"0006ec": 124,  // Harris
		"0006ed": 1035, // Inara
		"0006ef": 1036, // Maxxan
		"0006f0": 968,  // Digeo
		"0006f1": 1037, // Optillion
		"0006f2": 1038, // Platys
		"0006f3": 1039, // AcceLight
		"0006f4": 1040, // Prime
		"0006f5": 430,  // Alpsalpine
		"0006f6": 2,    // Cisco
		"0006f7": 430,  // Alpsalpine
		"0006f8": 1041, // Boeing
		"0006fc": 1042, // Fnet
		"0006fe": 1043, // Ambrado
		"0006ff": 1044, // Sheba
		"000701": 1045, // Racal-Datacom
		"000704": 430,  // Alpsalpine
		"000706": 1046, // Sanritz
		"000707": 1047, // Interalia
		"000708": 1048, // Bitrage
		"00070c": 1049, // SVA-Intrusion
		"00070d": 2,    // Cisco
		"00070e": 2,    // Cisco
		"00070f": 1050, // Fujant
		"000710": 1051, // Adax
		"000711": 1052, // Acterna
		"000714": 1053, // Brightcom
		"000718": 1054, // iCanTek
		"000719": 1055, // Mobiis
		"00071a": 1056, // Finedigital
		"00071c": 1057, // AT&T
		"00071e": 1058, // Tri-M
		"000720": 1059, // Trutzschler
		"000722": 1060, // Nielsen
		"000724": 1061, // Telemax
		"000725": 1062, // Bematech
		"000727": 1063, // Zi
		"00072a": 1064, // Innovance
		"00072c": 1065, // Fabricom
		"00072d": 1066, // CNSystems
		"00072f": 1067, // Intransa
		"000731": 1068, // Ophir-Spiricon
		"000732": 1069, // AAEON
		"000733": 1070, // DANCONTROL
		"000734": 1071, // ONStor
		"000735": 1072, // Flarion
		"000737": 1073, // Soriya
		"000738": 1074, // Young
		"00073a": 1075, // Inventel
		"00073b": 1076, // Tenovis
		"000740": 1077, // Buffalo
		"000742": 1078, // Ormazabal
		"000743": 1079, // Chelsio
		"000744": 1080, // Unico
		"000745": 1081, // Radlan
		"000746": 1082, // TURCK
		"000747": 1083, // Mecalc
		"000749": 1084, // CENiX
		"00074b": 1085, // Daihen
		"00074c": 1086, // Beicom
		"00074d": 768,  // Zebra
		"00074e": 1087, // IPFRONT
		"00074f": 2,    // Cisco
		"000750": 2,    // Cisco
		"000751": 1088, // m-u-t
		"000755": 1089, // Lafon
		"000757": 1090, // Topcall
		"000758": 1091, // DragonWave
		"000759": 1092, // Boris
		"00075d": 1093, // Celleritas
		"000761": 1094, // 29530
		"000767": 1095, // Yuxing
		"000768": 1096, // Danfoss
		"00076a": 1097, // NEXTEYE
		"00076b": 1098, // Stralfors
		"00076c": 1099, // Daehanet
		"00076d": 1100, // Flexlight
		"00076e": 1101, // Sinetica
		"00076f": 1102, // Synoptics
		"000770": 1103, // Ubiquoss
		"000777": 1104, // Motah
		"000778": 1105, // GERSTEL
		"00077d": 2,    // Cisco
		"00077e": 1106, // Elrest
		"00077f": 1107, // J
		"000780": 1108, // Bluegiga
		"000781": 1109, // Itron
		"000782": 11,   // Oracle
		"000784": 2,    // Cisco
		"000785": 2,    // Cisco
		"000786": 231,  // Wireless
		"000788": 1110, // Clipcomm
		"000789": 1111, // Allradio
		"00078b": 1112, // Wegener
		"00078d": 1113, // NetEngines
		"000790": 1058, // Tri-M
		"000792": 1114, // Sutron
		"000795": 1115, // Elitegroup
		"000796": 1116, // LSI
		"000797": 1117, // Netpower
		"00079a": 1118, // Verint
		"00079b": 1119, // Aurora
		"00079c": 1120, // Golden
		"00079d": 1121, // Musashi
		"00079e": 1122, // Ilinx
		"0007a0": 1123, // e-Watch
		"0007a2": 1124, // Opteon
		"0007a3": 1125, // Ositis
		"0007a5": 1126, // Y.D.K
		"0007a7": 1127, // A-Z
		"0007a9": 1128, // Novasonics
		"0007ab": 152,  // Samsung
		"0007ac": 1129, // Eolring
		"0007ae": 1130, // Britestream
		"0007b1": 1131, // Equator
		"0007b3": 2,    // Cisco
		"0007b4": 2,    // Cisco
		"0007b6": 1132, // Telecom
		"0007b8": 1133, // Corvalent
		"0007b9": 1134, // Ginganet
		"0007ba": 1135, // UTStarcom
		"0007bb": 1136, // Candera
		"0007bc": 1137, // Identix
		"0007bd": 1138, // Radionet
		"0007be": 1139, // DataLogic
		"0007bf": 1140, // Armillaire
		"0007c0": 1141, // NetZerver
		"0007c1": 818,  // Overture
		"0007c3": 1142, // Thomson
		"0007c4": 1143, // JEAN
		"0007c5": 1144, // Gcom
		"0007c7": 1145, // Synectics
		"0007c8": 1146, // Brain21
		"0007cd": 1147, // Kumoh
		"0007ce": 1148, // Cabletime
		"0007cf": 1149, // Anoto
		"0007d6": 1150, // Commil
		"0007d7": 1151, // Caporis
		"0007d8": 870,  // Hitron
		"0007d9": 1152, // Splicecom
		"0007db": 1153, // Kirana
		"0007dc": 1154, // Atek
		"0007dd": 1155, // Cradle
		"0007de": 1156, // eCopilt
		"0007df": 1157, // Vbrick
		"0007e0": 1158, // Palm
		"0007e1": 1159, // WIS
		"0007e2": 1160, // Bitworks
		"0007e3": 1161, // Navcom
		"0007e4": 1162, // SoftRadio
		"0007e5": 1163, // Coup
		"0007e7": 1164, // FreeWave
		"0007e8": 1165, // EdgeWave
		"0007e9": 421,  // Intel
		"0007ea": 1166, // Massana
		"0007eb": 2,    // Cisco
		"0007ec": 2,    // Cisco
		"0007ed": 1167, // Altera
		"0007f0": 1168, // LogiSync
		"0007f1": 1169, // TeraBurst
		"0007f2": 1170, // IOA
		"0007f3": 1171, // Thinkengine
		"0007f4": 1172, // Eletex
		"0007f5": 1173, // Bridgeco
		"0007f6": 1174, // Qqest
		"0007f7": 1175, // Galtronics
		"0007f8": 1176, // ITDevices
		"0007f9": 1177, // Sensaphone
		"0007fa": 1178, // ITT
		"0007fc": 1179, // Adept
		"0007fd": 1180, // LANergy
		"0007fe": 1181, // Rigaku
		"0007ff": 1182, // Gluon
		"000800": 1183, // Multitech
		"000802": 302,  // HP
		"000804": 1184, // ICA
		"000805": 1185, // Techno-Holon
		"000806": 1186, // Raonet
		"000809": 1187, // Systemonic
		"00080a": 1188, // Espera-Werke
		"00080d": 35,   // Toshiba
		"00080e": 125,  // ARRIS
		"000810": 1189, // Key
		"000811": 1190, // VOIX
		"000812": 1191, // GM-2
		"000813": 1192, // Diskbank
		"000814": 1193, // TIL
		"000815": 1194, // CATS
		"000817": 1195, // EmergeCore
		"000818": 1196, // Pixelworks
		"000819": 1197, // Banksys
		"00081b": 1198, // Windigo
		"00081c": 1199, // @pos.com
		"00081d": 1200, // Ipsilorporated
		"00081e": 1201, // Repeatit
		"000820": 2,    // Cisco
		"000821": 2,    // Cisco
		"000822": 1202, // InPro
		"000823": 1203, // Texa
		"000828": 1204, // Koei
		"000829": 1205, // Tokyo
		"00082b": 1206, // Wooksung
		"00082c": 1207, // Homag
		"00082e": 1208, // Multitone
		"00082f": 2,    // Cisco
		"000830": 2,    // Cisco
		"000831": 2,    // Cisco
		"000832": 2,    // Cisco
		"00084e": 1209, // DivergeNet
		"00084f": 1210, // Qualstar
		"000852": 1211, // Davolink
		"000854": 9,    // Netronix
		"000856": 1212, // Gamatronic
		"000857": 1213, // Polaris
		"000858": 1214, // Novatechnology
		"00085a": 1215, // IntiGate
		"00085b": 1216, // Hanbit
		"00085d": 1217, // Mitel
		"00085e": 1218, // PCO
		"000860": 1219, // LodgeNet
		"000861": 1220, // SoftEnergy
		"000863": 1221, // Entrisphere
		"000865": 1222, // Jascom
		"000868": 1223, // PurOptix
		"000869": 1224, // Command-e
		"00086a": 1225, // Securiton
		"00086b": 1226, // Mipsys
		"00086e": 1227, // Hyglo
		"000870": 1228, // Rasvia
		"000871": 1229, // NORTHDATA
		"000872": 1230, // Sorenson
		"000874": 102,  // Dell
		"000875": 1231, // Acorp
		"000876": 1232, // SDSystem
		"000877": 1233, // Liebert-Hiross
		"000879": 1234, // CEM
		"00087a": 1235, // Wipotec
		"00087c": 2,    // Cisco
		"00087d": 2,    // Cisco
		"00087e": 1236, // Bon
		"00087f": 1237, // SPAUN
		"000882": 291,  // Sigma
		"000883": 302,  // HP
		"000889": 1238, // Dish
		"00088a": 1239, // Minds@Work
		"00088b": 1240, // Tropic
		"00088d": 1241, // Sigma-Links
		"00088e": 1242, // Nihon
		"000890": 1243, // Avilinks
		"000891": 1244, // Lyan
		"000892": 1245, // EM
		"000896": 1246, // Printronix
		"000897": 1247, // Quake
		"000899": 1248, // Netbind
		"00089b": 1249, // ICP
		"00089c": 1250, // Elecs
		"00089d": 1251, // UHD-Elektronik
		"00089f": 1252, // EFM
		"0008a1": 1253, // CNet
		"0008a2": 1254, // ADI
		"0008a3": 2,    // Cisco
		"0008a4": 2,    // Cisco
		"0008a5": 1255, // Peninsula
		"0008a7": 1256, // iLogic
		"0008a8": 1257, // Systec
		"0008a9": 1258, // SangSang
		"0008aa": 1259, // Karam
		"0008ab": 1260, // EnerLinx.com
		"0008ac": 1261, // BST
		"0008ad": 1262, // Toyo-Linx
		"0008af": 1263, // Novatec
		"0008b1": 1264, // ProQuent
		"0008b3": 1265, // Fastwel
		"0008b4": 1266, // Syspol
		"0008b6": 1267, // RouteFree
		"0008b7": 1268, // HIT
		"0008b9": 1269, // Kaonmedia
		"0008ba": 1270, // Erskine
		"0008bb": 1271, // NetExcell
		"0008bc": 1272, // Ilevo
		"0008bd": 1273, // Tepg-US
		"0008c0": 1274, // ASA
		"0008c1": 1275, // Avistar
		"0008c2": 2,    // Cisco
		"0008c3": 1276, // Contex
		"0008c4": 1277, // Hikari
		"0008c5": 1278, // Liontech
		"0008c7": 302,  // HP
		"0008c8": 1279, // Soneticom
		"0008ca": 1280, // TwinHan
		"0008cc": 1281, // Remotec
		"0008cd": 1282, // With-Net
		"0008ce": 1283, // IPMobileNet
		"0008d0": 1121, // Musashi
		"0008d1": 1284, // Karel
		"0008d2": 1285, // ZOOM
		"0008d4": 1286, // IneoQuest
		"0008d6": 1287, // HASSNET
		"0008d7": 1288, // HOW
		"0008d9": 1289, // Mitadenshi
		"0008da": 1290, // SofaWare
		"0008db": 1291, // Corrigent
		"0008dc": 1292, // Wiznet
		"0008dd": 1293, // Telena
		"0008de": 1294, // 3UP
		"0008df": 1295, // Alistel
		"0008e0": 1296, // ATO
		"0008e1": 1297, // Barix
		"0008e2": 2,    // Cisco
		"0008e3": 2,    // Cisco
		"0008e4": 1298, // Envenergy
		"0008e5": 1299, // IDK
		"0008e6": 1300, // Littlefeet
		"0008e9": 1301, // NextGig
		"0008eb": 1302, // ROMWin
		"0008ef": 1303, // Dibal
		"0008f1": 1304, // Voltaire
		"0008f2": 1305, // C&S
		"0008f3": 1306, // Wany
		"0008f4": 1307, // Bluetake
		"0008f5": 1308, // YESTECHNOLOGY
		"0008fb": 1309, // SonoSite
		"0008fc": 1310, // Gigaphoton
		"0008fd": 1311, // BlueKorea
		"0008ff": 1312, // Trilogy
		"000900": 1313, // TMT
		"000902": 1314, // Redline
		"000903": 1315, // Panasas
		"000904": 1316, // MONDIAL
		"000905": 1317, // iTEC
		"000906": 1318, // Esteem
		"000908": 1319, // VTech
		"00090a": 1320, // SnedFar
		"00090c": 1321, // Mayekawa
		"00090d": 804,  // Leader
		"00090e": 1322, // Helix
		"00090f": 1323, // Fortinet
		"000911": 2,    // Cisco
		"000912": 2,    // Cisco
		"000913": 1324, // SystemK
		"000914": 1325, // Computrols
		"000915": 1326, // CAS
		"000917": 1327, // WEM
		"00091c": 1328, // CacheVision
		"00091d": 1329, // Proteam
		"00091e": 1330, // Firstech
		"00091f": 1331, // A&D
		"000920": 1332, // EpoX
		"000921": 1333, // Planmeca
		"000924": 1334, // Telebau
		"000926": 1335, // Yoda
		"000927": 734,  // Toyokeiki
		"000928": 1336, // Telecore
		"000929": 1337, // Sanyo
		"00092a": 1338, // MYTECS
		"00092b": 1339, // iQstor
		"00092c": 1340, // Hitpoint
		"00092d": 1341, // HTC
		"00092f": 1342, // Akom
		"000930": 1343, // AeroConcierge
		"000932": 1344, // Omnilux
		"000933": 1345, // Ophit
		"000934": 1346, // Dream-Multimedia-Tv
		"000935": 1347, // Sandvine
		"000936": 1348, // Ipetronik
		"000938": 1349, // Allot
		"000939": 1350, // ShibaSoku
		"00093b": 1351, // Hyundai
		"00093d": 1352, // Newisys
		"00093e": 1353, // C&I
		"000940": 1354, // AGFEO
		"000942": 231,  // Wireless
		"000943": 2,    // Cisco
		"000944": 2,    // Cisco
		"000945": 1355, // Palmmicro
		"000946": 1356, // Cluster
		"000947": 1357, // Aztek
		"000949": 1358, // Glyph
		"00094a": 1359, // Homenet
		"00094b": 1360, // FillFactory
		"00094d": 1361, // Braintree
		"00094e": 1362, // Bartech
		"00094f": 1363, // elmegt
		"000952": 1364, // Auerswald
		"000957": 1365, // Supercaller
		"000959": 1366, // Sitecsoft
		"00095a": 1367, // Racewood
		"00095b": 1368, // Netgear
		"00095d": 1369, // Dialogue
		"00095f": 1370, // Telebyte
		"000960": 1371, // YOZAN
		"000962": 1372, // Sonitor
		"000964": 1373, // Hi-Techniques
		"000965": 1374, // HyunJu
		"000966": 1375, // Trimble
		"000967": 1376, // Tachyon
		"000968": 1377, // Technoventure
		"00096a": 1378, // Cloverleaf
		"00096b": 372,  // IBM
		"00096d": 1379, // Powernet
		"00096e": 1380, // Giant
		"000970": 1381, // Vibration
		"000972": 1382, // Securebase
		"000973": 1383, // Lenten
		"000974": 1384, // Innopia
		"000975": 1385, // fSONA
		"00097b": 2,    // Cisco
		"00097c": 2,    // Cisco
		"00097d": 1386, // SecWell
		"00097e": 1387, // IMI
		"000981": 1388, // Newport
		"000983": 1389, // GlobalTop
		"000986": 1390, // Metalink
		"000988": 1391, // Nudian
		"000989": 1392, // VividLogic
		"00098a": 1393, // EqualLogic
		"00098b": 1394, // Entropic
		"00098e": 1395, // ipcas
		"00098f": 1396, // Cetacean
		"000990": 1397, // ACKSYS
		"000992": 1398, // InterEpoch
		"000993": 1399, // Visteon
		"000994": 1400, // Cronyx
		"000995": 1401, // Castle
		"000996": 1402, // RDI
		"000997": 76,   // Nortel
		"000998": 1403, // Capinfo
		"00099a": 1404, // Elmo
		"00099d": 1405, // Haliplex
		"00099e": 1406, // Testech
		"00099f": 1407, // Videx
		"0009a0": 1408, // Microtechno
		"0009a1": 1409, // Telewise
		"0009a2": 1410, // Interface
		"0009a3": 1411, // Leadfly
		"0009a4": 1412, // HARTEC
		"0009a8": 1413, // Eastmode
		"0009a9": 1414, // Ikanos
		"0009ab": 1415, // Netcontrol
		"0009ac": 1416, // Lanvoice
		"0009af": 1417, // e-generis
		"0009b1": 1418, // Kanematsu
		"0009b2": 1419, // L&F
		"0009b3": 1420, // MCM
		"0009b5": 1421, // 3J
		"0009b6": 2,    // Cisco
		"0009b7": 2,    // Cisco
		"0009b8": 1422, // Entise
		"0009bb": 1423, // MathStar
		"0009bc": 1424, // Utility
		"0009bd": 1425, // Epygi
		"0009be": 1426, // Mamiya-OP
		"0009bf": 1427, // Nintendo
		"0009c0": 1428, // 6WIND
		"0009c1": 1429, // Proces-Data
		"0009c2": 1430, // Onity
		"0009c3": 1431, // Netas
		"0009c4": 1432, // Medicore
		"0009c5": 1433, // KINGENE
		"0009c6": 1434, // Visionics
		"0009c7": 1435, // Movistec
		"0009c9": 1436, // BlueWINC
		"0009ca": 1437, // iMaxNetworksLimited
		"0009cb": 1438, // HBrain
		"0009cc": 1439, // Moog
		"0009cf": 1440, // iAd
		"0009d0": 1441, // Solacom
		"0009d1": 1442, // Seranoa
		"0009d4": 1443, // Transtech
		"0009d5": 569,  // Signal
		"0009d8": 1444, // Falt
		"0009d9": 1445, // Neoscale
		"0009db": 1446, // eSpace
		"0009dc": 1447, // Galaxis
		"0009dd": 1448, // Mavin
		"0009df": 1449, // Vestel
		"0009e1": 1450, // Gemtek
		"0009e2": 1451, // Sinbon
		"0009e8": 2,    // Cisco
		"0009e9": 2,    // Cisco
		"0009ea": 1452, // YEM
		"0009eb": 1453, // HuMANDATA
		"0009ec": 1454, // Daktronics
		"0009ed": 1455, // CipherOptics
		"0009ef": 1456, // Vocera
		"0009f0": 1457, // Shimizu
		"0009f2": 1458, // Cohu
		"0009f3": 1459, // WELL
		"0009f4": 1460, // Alcon
		"0009f7": 1461, // SED
		"0009f8": 1462, // Unimo
		"0009fc": 1463, // IPFLEX
		"0009fd": 1464, // Ubinetics
		"0009fe": 1465, // Daisy
		"000a00": 1466, // Mediatek
		"000a01": 1467, // SOHOware
		"000a02": 1468, // Annso
		"000a04": 160,  // 3Com
		"000a05": 1469, // Widax
		"000a06": 1470, // Teledex
		"000a07": 1471, // WebWayOne
		"000a0a": 1472, // SUNIX
		"000a0b": 1473, // Sealevel
		"000a0c": 1474, // Scientific
		"000a0d": 1475, // Amphenol
		"000a0e": 1476, // Invivo
		"000a11": 1477, // ExPet
		"000a12": 1478, // Azylex
		"000a16": 1479, // Lassen
		"000a17": 1480, // Nestar
		"000a18": 1481, // Vichel
		"000a1a": 1482, // Imerge
		"000a1b": 1483, // Stream
		"000a1d": 1484, // Optical
		"000a1e": 1485, // Red-M
		"000a20": 1486, // SVA
		"000a22": 1487, // Amperion
		"000a23": 1488, // Parama
		"000a24": 1489, // Octave
		"000a25": 1490, // Ceragon
		"000a27": 545,  // Apple
		"000a28": 687,  // Motorola
		"000a2a": 566,  // QSI
		"000a2b": 1491, // Etherstuff
		"000a2d": 1492, // Cabot
		"000a2e": 1493, // Maple
		"000a2f": 1494, // Artnix
		"000a30": 1399, // Visteon
		"000a31": 1495, // HCV
		"000a32": 1496, // Xsido
		"000a33": 129,  // Emulex
		"000a34": 1497, // Identicard
		"000a35": 1498, // Xilinx
		"000a37": 1499, // Procera
		"000a38": 1500, // Apani
		"000a3a": 1501, // J-THREE
		"000a3c": 1502, // Enerpoint
		"000a41": 2,    // Cisco
		"000a42": 2,    // Cisco
		"000a45": 1503, // Audio-Technica
		"000a48": 1504, // Albatron
		"000a49": 289,  // F5
		"000a4a": 1505, // Targa
		"000a4b": 1506, // DataPower
		"000a4d": 1507, // Noritz
		"000a4e": 1508, // UNITEK
		"000a50": 1509, // Remotek
		"000a51": 1510, // GyroSignal
		"000a52": 1511, // AsiaRF
		"000a53": 1512, // Intronicsorporated
		"000a55": 1513, // MARKEM
		"000a57": 302,  // HP
		"000a5a": 1514, // GreenNET
		"000a5b": 1515, // Power-One
		"000a5e": 160,  // 3Com
		"000a5f": 1516, // almedio
		"000a60": 1517, // Autostar
		"000a61": 1518, // Cellinx
		"000a62": 1519, // Crinis
		"000a63": 1520, // DHD
		"000a64": 1521, // Eracom
		"000a65": 1522, // GentechMedia
		"000a67": 1523, // OngCorp
		"000a68": 1524, // Solarflare
		"000a6c": 1525, // Walchem
		"000a6e": 1526, // Harmonic
		"000a6f": 1527, // ZyFLEX
		"000a71": 1528, // Avrio
		"000a72": 1529, // Stec
		"000a74": 1530, // Manticom
		"000a75": 1531, // Caterpillar
		"000a77": 1532, // Bluewire
		"000a78": 1533, // Olitec
		"000a7c": 1534, // Tecton
		"000a7d": 1535, // Valo
		"000a7f": 1536, // Teradon
		"000a80": 1537, // Telkonet
		"000a84": 1538, // Rainsun
		"000a85": 1539, // PLAT'C2
		"000a86": 1540, // Lenze
		"000a89": 1541, // Creval
		"000a8a": 2,    // Cisco
		"000a8b": 2,    // Cisco
		"000a8c": 1542, // Guardware
		"000a8d": 1543, // Eurotherm
		"000a8e": 1544, // Invacom
		"000a8f": 1545, // Aska
		"000a91": 1546, // HemoCue
		"000a92": 1547, // Presonus
		"000a93": 1548, // W2
		"000a95": 545,  // Apple
		"000a96": 1549, // Mewtel
		"000a97": 1550, // SONICblue
		"000a9a": 1551, // Aiptek
		"000a9c": 1552, // Server
		"000a9f": 1553, // Pannaway
		"000aa2": 1554, // Systek
		"000aa5": 1555, // Maxlink
		"000aa6": 1556, // Hochiki
		"000aa7": 1557, // FEI
		"000aa8": 1558, // ePipe
		"000aaa": 1559, // AltiGen
		"000aac": 1560, // TerraTec
		"000aad": 1561, // Stargames
		"000aaf": 1562, // Pipal
		"000ab0": 1563, // LOYTEC
		"000ab1": 1564, // GENETEC
		"000ab5": 1565, // Digital
		"000ab6": 1566, // Compunetix
		"000ab7": 2,    // Cisco
		"000ab8": 2,    // Cisco
		"000ab9": 1567, // Astera
		"000aba": 1568, // Arcon
		"000abc": 1569, // Seabridge
		"000abe": 1570, // OPNET
		"000ac1": 1571, // Futuretel
		"000ac8": 1572, // ZPSYS
		"000ac9": 1573, // Zambeel
		"000acc": 1574, // Winnow
		"000acd": 1575, // Sunrich
		"000ace": 1576, // Radiantech
		"000ad1": 1577, // MWS
		"000ad2": 1578, // JEPICO
		"000ad3": 1579, // INITECH
		"000ad4": 1580, // CoreBell
		"000ad5": 1581, // Brainchild
		"000ad6": 1582, // BeamReach
		"000ad8": 1583, // IPCserv
		"000ad9": 101,  // Sony
		"000ada": 1584, // Vindicator
		"000adb": 1585, // Trilliant
		"000adc": 1586, // RuggedCom
		"000add": 1587, // Allworx
		"000ade": 1588, // Happy
		"000adf": 1589, // Gennum
		"000ae1": 1590, // EG
		"000ae2": 1591, // Binatone
		"000ae4": 1592, // Wistron
		"000ae5": 1593, // ScottCare
		"000ae6": 1115, // Elitegroup
		"000ae9": 1594, // AirVast
		"000aeb": 1595, // TP-Link
		"000aed": 1596, // HARTING
		"000af0": 1597, // Shin-OH
		"000af2": 1598, // NeoAxiom
		"000af3": 2,    // Cisco
		"000af4": 2,    // Cisco
		"000af5": 1599, // Airgo
		"000af7": 858,  // Broadcom
		"000af9": 1600, // HiConnect
		"000afb": 1601, // Ambri
		"000afd": 1602, // Kentec
		"000afe": 1603, // NovaPal
		"000b01": 1604, // Daiichi
		"000b02": 1605, // Dallmeier
		"000b03": 1606, // Taekwang
		"000b04": 1607, // Volktek
		"000b06": 125,  // ARRIS
		"000b07": 1608, // Voxpath
		"000b0b": 1609, // Corrent
		"000b0c": 1610, // Agile
		"000b0d": 1611, // Air2U
		"000b0e": 1612, // Trapeze
		"000b13": 1613, // Zetron
		"000b14": 1614, // ViewSonic
		"000b15": 1615, // Platypus
		"000b18": 67,   // Private
		"000b19": 1616, // Vernier
		"000b1b": 1617, // Systronix
		"000b1c": 1618, // SIBCO
		"000b20": 1619, // Hirata
		"000b21": 1620, // G-Star
		"000b23": 300,  // Siemens
		"000b24": 1621, // AirLogic
		"000b25": 1622, // Aeluros
		"000b26": 1623, // Wetek
		"000b27": 1624, // Scion
		"000b28": 1625, // Quatech
		"000b29": 1626, // LS
		"000b2a": 1627, // HOWTEL
		"000b2b": 1628, // Hostnet
		"000b2c": 1629, // Eiki
		"000b2d": 1096, // Danfoss
		"000b2e": 1630, // Cal-Comp
		"000b2f": 1631, // bplan
		"000b32": 1632, // Vormetric
		"000b33": 1633, // Vivato
		"000b36": 1634, // Productivity
		"000b38": 1635, // Knurr
		"000b3a": 1636, // Pesa
		"000b3b": 1637, // devolo
		"000b3e": 1638, // BittWare
		"000b3f": 1639, // Anthology
		"000b40": 1640, // Cambridge
		"000b42": 1641, // Commax
		"000b43": 1642, // Microscan
		"000b45": 2,    // Cisco
		"000b46": 2,    // Cisco
		"000b48": 1643, // sofrel
		"000b4a": 1644, // Visimetrics
		"000b4b": 1645, // Visiowave
		"000b4d": 1646, // Emuzed
		"000b4f": 1647, // Verifone
		"000b50": 1648, // Oxygnet
		"000b51": 1649, // Micetek
		"000b52": 1650, // Joymax
		"000b53": 1651, // INITIUM
		"000b54": 1652, // BiTMICRO
		"000b55": 1653, // ADInstruments
		"000b56": 1654, // Cybernetics
		"000b57": 1655, // Silicon
		"000b59": 1656, // ScriptPro
		"000b5a": 1657, // HyperEdge
		"000b5b": 1658, // Rincon
		"000b5c": 1659, // Newtech
		"000b5d": 4,    // Fujitsu
		"000b5e": 1660, // Audio
		"000b5f": 2,    // Cisco
		"000b60": 2,    // Cisco
		"000b62": 1661, // ib-mohnen
		"000b63": 1662, // Kaleidescape
		"000b66": 1663, // Teralink
		"000b67": 1664, // Topview
		"000b68": 1665, // Addvalue
		"000b6a": 1666, // Asiarock
		"000b6b": 1592, // Wistron
		"000b6c": 1667, // Sychip
		"000b70": 1668, // Load
		"000b71": 1669, // Litchfield
		"000b72": 1670, // Lawo
		"000b73": 1671, // Kodeos
		"000b74": 1672, // Kingwave
		"000b75": 1673, // Iosoft
		"000b76": 1674, // ET&T
		"000b77": 1675, // Cogent
		"000b78": 1676, // Taifatech
		"000b79": 1677, // X-CoM
		"000b7b": 1678, // Test-Um
		"000b7c": 1679, // Telex
		"000b7f": 1680, // Align
		"000b80": 1681, // Lycium
		"000b81": 1682, // Kaparel
		"000b82": 1683, // Grandstream
		"000b84": 1684, // Bodet
		"000b85": 2,    // Cisco
		"000b86": 1685, // Aruba
		"000b88": 1686, // Vidisco
		"000b8a": 1687, // MITEQ
		"000b8b": 1688, // Kerajet
		"000b8c": 833,  // Flextronics
		"000b8d": 1689, // Avvio
		"000b8e": 1690, // Ascent
		"000b8f": 1691, // Akita
		"000b92": 1692, // Ascom
		"000b98": 1693, // NiceTechVision
		"000b99": 1694, // SensAble
		"000b9c": 1695, // TriBeam
		"000b9d": 1696, // TwinMOS
		"000b9e": 1697, // Yasing
		"000ba1": 1698, // Fujikura
		"000ba3": 300,  // Siemens
		"000ba7": 1699, // Maranti
		"000ba8": 1700, // Hanback
		"000ba9": 1701, // CloudShield
		"000baa": 1702, // Aiphone
		"000bab": 1703, // Advantech
		"000bac": 160,  // 3Com
		"000bad": 1704, // PC-PoS
		"000baf": 1705, // WOOJU
		"000bb2": 1706, // Smallbig
		"000bb3": 1707, // RiT
		"000bb5": 1708, // nStor
		"000bb6": 1709, // Metalligence
		"000bb7": 43,   // Micro
		"000bb8": 1710, // Kihoku
		"000bb9": 1711, // Imsys
		"000bba": 1526, // Harmonic
		"000bbb": 1712, // Etin
		"000bbd": 1713, // Connexionz
		"000bbe": 2,    // Cisco
		"000bbf": 2,    // Cisco
		"000bc2": 1714, // Corinex
		"000bc3": 1715, // Multiplex
		"000bc4": 1716, // BIOTRONIK
		"000bc5": 741,  // SMC
		"000bc6": 1717, // ISAC
		"000bc8": 1718, // AirFlow
		"000bcc": 1719, // Jusan
		"000bcd": 302,  // HP
		"000bce": 1720, // Free2move
		"000bd1": 1721, // Aeronix
		"000bd2": 1722, // Remopro
		"000bd3": 1723, // cd3o
		"000bd5": 1724, // Nvergence
		"000bda": 1725, // EyeCross
		"000bdb": 102,  // Dell
		"000bdc": 1726, // Akcp
		"000bde": 1727, // TELDIX
		"000be0": 1728, // SercoNet
		"000be2": 1729, // Lumenera
		"000be4": 1730, // Hosiden
		"000be5": 1731, // HIMS
		"000be6": 1732, // Datel
		"000be7": 1733, // Comflux
		"000be8": 1734, // Aoip
		"000be9": 1735, // Actel
		"000bea": 1736, // Zultys
		"000beb": 1737, // Systegra
		"000bed": 1738, // ELM
		"000bee": 1739, // Inc.jetorporated
		"000bef": 1740, // Code
		"000bf0": 1741, // MoTEX
		"000bf2": 1742, // Chih-Kan
		"000bf3": 1743, // BAE
		"000bf4": 67,   // Private
		"000bf6": 1744, // Nitgen
		"000bf7": 1745, // Nidek
		"000bf8": 1746, // Infinera
		"000bf9": 1747, // Gemstone
		"000bfb": 1748, // D-NET
		"000bfc": 2,    // Cisco
		"000bfd": 2,    // Cisco
		"000c01": 1749, // Abatron
		"000c02": 21,   // ABB
		"000c04": 1750, // Tecnova
		"000c06": 1751, // Nixvue
		"000c07": 1752, // Iftest
		"000c08": 1753, // HUMEX
		"000c0b": 1754, // Broadbus
		"000c0c": 1755, // Appro
		"000c0e": 1756, // XtremeSpectrum
		"000c0f": 1757, // Techno-One
		"000c10": 1758, // PNI
		"000c12": 1759, // Micro-Optronic-Messtechnik
		"000c13": 1760, // MediaQ
		"000c15": 1761, // CyberPower
		"000c19": 1762, // Telio
		"000c1b": 1763, // ORACOM
		"000c1c": 1764, // MicroWeb
		"000c1f": 1765, // Glimmerglass
		"000c24": 1766, // Anator
		"000c26": 1767, // Weintek
		"000c27": 1768, // Sammy
		"000c28": 1769, // Rifatron
		"000c29": 809,  // VMware
		"000c2a": 1770, // OCTTEL
		"000c2b": 1771, // ELIAS
		"000c2c": 1772, // Enwiser
		"000c2d": 1773, // FullWave
		"000c2f": 1774, // SeorimTechnology
		"000c30": 2,    // Cisco
		"000c31": 2,    // Cisco
		"000c33": 1775, // Compucase
		"000c34": 1776, // Vixen
		"000c36": 1777, // S-Takaya
		"000c37": 1778, // Geomation
		"000c38": 1779, // TelcoBridges
		"000c3a": 1780, // Oxance
		"000c3c": 1781, // MediaChorus
		"000c3d": 1782, // Glsystech
		"000c41": 1783, // Linksys
		"000c42": 1784, // Routerboard.com
		"000c43": 1785, // Ralink
		"000c45": 1786, // Animation
		"000c48": 1787, // QoStek
		"000c4c": 1788, // Arcor&Co
		"000c4e": 1789, // Winbest
		"000c50": 732,  // Seagate
		"000c51": 1474, // Scientific
		"000c52": 1790, // Roll
		"000c53": 67,   // Private
		"000c54": 1791, // Pedestal
		"000c55": 254,  // Microlink
		"000c56": 1792, // Megatel
		"000c57": 1793, // MACKIE
		"000c58": 1794, // M&S
		"000c59": 1795, // Indyme
		"000c5b": 1796, // Hanwang
		"000c5d": 1797, // Chic
		"000c5f": 1798, // Avtec
		"000c60": 1799, // ACM
		"000c62": 21,   // ABB
		"000c63": 1800, // Zenith
		"000c66": 1801, // Pronto
		"000c68": 1802, // SigmaTel
		"000c6a": 1803, // Mbari
		"000c6c": 1804, // Eve
		"000c6d": 1805, // Edwards
		"000c6e": 1806, // ASUS
		"000c70": 1807, // ACC
		"000c71": 1808, // Wybron
		"000c72": 1809, // Tempearl
		"000c73": 1810, // Telson
		"000c74": 1811, // Rivertec
		"000c76": 1812, // Micro-Star
		"000c78": 1813, // In-Tech
		"000c7a": 1814, // DaTARIUS
		"000c7e": 1815, // Tellium
		"000c7f": 1816, // synertronixx
		"000c80": 1817, // Opelcomm
		"000c81": 56,   // Schneider
		"000c82": 109,  // Network
		"000c83": 253,  // Logical
		"000c84": 1818, // Eazix
		"000c85": 2,    // Cisco
		"000c86": 2,    // Cisco
		"000c87": 13,   // AMD
		"000c8a": 1819, // Bose
		"000c8b": 22,   // Connect
		"000c8c": 1820, // Kodicom
		"000c8e": 1821, // Mentor
		"000c90": 1822, // Octasic
		"000c91": 1823, // Riverhead
		"000c92": 1824, // WolfVision
		"000c93": 1825, // Xeline
		"000c94": 1826, // United
		"000c95": 1827, // PrimeNet
		"000c96": 1828, // OQO
		"000c98": 1829, // LETEK
		"000c9a": 1830, // Hitech
		"000c9b": 1831, // EE
		"000c9d": 1832, // UbeeAirWalk
		"000c9e": 1833, // MemoryLink
		"000c9f": 1834, // NKE
		"000ca0": 1835, // StorCase
		"000ca1": 1836, // SIGMACOM
		"000ca3": 1837, // Rancho
		"000ca6": 1838, // Mintera
		"000ca7": 1839, // Metro
		"000ca8": 1840, // Garuda
		"000ca9": 1841, // Ebtron
		"000cab": 1842, // Commend
		"000cad": 1843, // BTU
		"000cae": 1844, // Ailocom
		"000cb1": 1845, // Salland
		"000cb2": 1846, // UNION
		"000cb3": 1847, // ROUND
		"000cb4": 1848, // AutoCell
		"000cb8": 1849, // Medion
		"000cb9": 1850, // LEA
		"000cba": 1851, // Jamex
		"000cbb": 1852, // Iskraemeco
		"000cbc": 1853, // Iscutum
		"000cc0": 1854, // Genera
		"000cc1": 1855, // Eaton
		"000cc2": 1856, // ControlNet
		"000cc3": 1857, // BeWAN
		"000cc4": 1858, // Tiptel
		"000cc5": 1859, // Nextlink
		"000cc6": 1860, // Ka-Ro
		"000cc7": 1861, // Intelligent
		"000cca": 1862, // HGST
		"000ccc": 1863, // Aeroscout
		"000cce": 2,    // Cisco
		"000ccf": 2,    // Cisco
		"000cd0": 1864, // Symetrix
		"000cd1": 1865, // SFOM
		"000cd5": 1866, // Passave
		"000cd6": 1867, // Partner
		"000cd7": 1868, // Nallatech
		"000cd9": 1869, // Itcare
		"000cda": 1870, // FreeHand
		"000cdb": 90,   // Brocade
		"000cdc": 1871, // BECS
		"000cdd": 1872, // AOS
		"000cdf": 1873, // JAI
		"000ce2": 1874, // Rolls-Royce
		"000ce4": 1875, // NeuroCom
		"000ce5": 125,  // ARRIS
		"000ce6": 1323, // Fortinet
		"000ce7": 1876, // MediaTek
		"000ceb": 1877, // CNMP
		"000cee": 1878, // jp-embedded
		"000cef": 1879, // Open
		"000cf1": 421,  // Intel
		"000cf5": 1880, // InfoExpress
		"000cf6": 1881, // Sitecom
		"000cf7": 76,   // Nortel
		"000cf8": 76,   // Nortel
		"000cfa": 1565, // Digital
		"000cfc": 1882, // S2io
		"000cfe": 1883, // Grand
		"000d00": 1884, // Seaway
		"000d03": 1885, // Matrics
		"000d05": 1886, // cybernet
		"000d06": 1887, // Compulogic
		"000d08": 1888, // AboveCable
		"000d09": 1889, // Yuehua
		"000d0b": 1077, // Buffalo
		"000d0d": 1890, // ITSupported
		"000d0e": 1891, // Inqnet
		"000d0f": 1892, // Finlux
		"000d10": 1893, // Embedtronics
		"000d12": 1894, // AXELL
		"000d16": 1895, // UHS
		"000d17": 1896, // Turbo
		"000d18": 1897, // Mega-Trend
		"000d1b": 1898, // Kyoto
		"000d21": 1899, // WISCORE
		"000d22": 1900, // Unitronics
		"000d25": 1901, // Sanden
		"000d26": 1902, // Primagraphics
		"000d28": 2,    // Cisco
		"000d29": 2,    // Cisco
		"000d2a": 1903, // Scanmatic
		"000d2c": 1904, // Net2Edge
		"000d2f": 1905, // AINTech
		"000d31": 1906, // Compellent
		"000d32": 1907, // DispenseSource
		"000d33": 1908, // Prediwave
		"000d35": 177,  // PAC
		"000d37": 1909, // Wiplug
		"000d38": 1910, // Nissin
		"000d39": 109,  // Network
		"000d3a": 612,  // Microsoft
		"000d3b": 1911, // Microelectronics
		"000d3d": 1912, // Hammerhead
		"000d3e": 1913, // APLUX
		"000d41": 300,  // Siemens
		"000d47": 1914, // Collex
		"000d48": 1915, // AEWIN
		"000d4b": 1916, // Roku
		"000d4c": 1917, // Outline
		"000d4d": 1918, // Ninelanes
		"000d4e": 1919, // NDR
		"000d4f": 1920, // Kenwood
		"000d50": 1921, // Galazar
		"000d51": 1922, // DivR
		"000d54": 160,  // 3Com
		"000d55": 1923, // SANYCOM
		"000d56": 102,  // Dell
		"000d58": 67,   // Private
		"000d59": 1924, // Amity
		"000d5a": 1925, // Tiesse
		"000d5c": 1926, // Bosch
		"000d5d": 1927, // Raritan
		"000d5f": 1928, // Minds
		"000d60": 372,  // IBM
		"000d61": 1929, // Giga-Byte
		"000d65": 2,    // Cisco
		"000d66": 2,    // Cisco
		"000d67": 306,  // Ericsson
		"000d68": 1930, // Vinci
		"000d69": 1931, // TMT&D
		"000d6a": 1932, // Redwood
		"000d6b": 1933, // Mita-Teknik
		"000d6c": 1934, // M-Audio
		"000d6e": 1935, // K-Patents
		"000d6f": 1936, // Ember
		"000d70": 1937, // Datamax
		"000d71": 1938, // boca
		"000d72": 1939, // 2Wire
		"000d77": 1940, // FalconStor
		"000d79": 1941, // Dynamic
		"000d7b": 1942, // Consensys
		"000d7c": 1943, // Codian
		"000d7d": 1944, // Afco
		"000d7e": 1945, // Axiowave
		"000d7f": 1946, // MIDAS
		"000d81": 1947, // Pepperl+Fuchs
		"000d82": 1948, // Phsnet
		"000d84": 1949, // Makus
		"000d85": 1950, // Tapwave
		"000d87": 1115, // Elitegroup
		"000d88": 803,  // D-Link
		"000d89": 1951, // Bils
		"000d8a": 1952, // Winners
		"000d8b": 1953, // T&D
		"000d8d": 1954, // Prosoft
		"000d8e": 1955, // Koden
		"000d90": 1956, // Factum
		"000d92": 1957, // ARIMA
		"000d93": 545,  // Apple
		"000d94": 1958, // AFAR
		"000d95": 1959, // Opti-cell
		"000d96": 1960, // Vtera
		"000d97": 1961, // ABB./Tropos
		"000d9a": 1962, // Infotec
		"000d9b": 1963, // Heraeus
		"000d9d": 302,  // HP
		"000da2": 1964, // Infrant
		"000da3": 1965, // Emerging
		"000da5": 1966, // Fabric7
		"000da7": 67,   // Private
		"000da8": 1967, // TElectronics
		"000da9": 1968, // Ingeteam
		"000daa": 1969, // S.A.Tehnology
		"000dad": 1970, // Dataprobe
		"000daf": 1971, // Plexus
		"000db0": 1972, // Olym-tech
		"000db2": 1973, // Ammasso
		"000db4": 1974, // Stormshield
		"000db5": 1975, // Globalsat
		"000db6": 858,  // Broadcom
		"000db8": 1976, // Schiller
		"000dbc": 2,    // Cisco
		"000dbd": 2,    // Cisco
		"000dc0": 1977, // Spagat
		"000dc1": 1978, // SafeWeb
		"000dc2": 67,   // Private
		"000dc3": 1979, // First
		"000dc4": 1980, // Emcore
		"000dc6": 1981, // DigiRose
		"000dc7": 1982, // Cosmic
		"000dc8": 1983, // AirMagnet
		"000dca": 1984, // Tait
		"000dcb": 1985, // Petcomkorea
		"000dcc": 1986, // NEOSMART
		"000dce": 1987, // Dynavac
		"000dcf": 1988, // Cidra
		"000dd1": 1989, // Stryker
		"000dd4": 1990, // Symantec
		"000dd5": 1991, // O'Rite
		"000dd6": 1992, // ITI
		"000dd7": 1993, // Bright
		"000dd8": 1994, // BBN
		"000ddb": 1995, // Airwave
		"000ddc": 1996, // VAC
		"000dde": 1997, // Joyteck
		"000de0": 1998, // ICPDAS
		"000de1": 1999, // Control
		"000de4": 2000, // DIGINICS
		"000de6": 2001, // Youngbo
		"000de8": 2002, // Nasaco
		"000deb": 2003, // CompXs
		"000dec": 2,    // Cisco
		"000ded": 2,    // Cisco
		"000df0": 2004, // Qcom
		"000df1": 2005, // Ionix
		"000df2": 67,   // Private
		"000df3": 2006, // Asmax
		"000df4": 2007, // Watertek
		"000df5": 1967, // TElectronics
		"000df9": 2008, // NDS
		"000dfb": 2009, // Komax
		"000dfc": 2010, // ITFOR
		"000e00": 2011, // Atrie
		"000e01": 2012, // ASIP
		"000e03": 129,  // Emulex
		"000e04": 2013, // CMA/Microdialysis
		"000e07": 101,  // Sony
		"000e08": 1783, // Linksys
		"000e0b": 2014, // Netac
		"000e0c": 421,  // Intel
		"000e0f": 2015, // Ermme
		"000e10": 2016, // C-guys
		"000e13": 2017, // Accu-Sort
		"000e14": 2018, // Visionary
		"000e15": 2019, // Tadlys
		"000e17": 67,   // Private
		"000e18": 2020, // MyA
		"000e19": 2021, // LogicaCMG
		"000e1a": 2022, // JPS
		"000e1b": 2023, // IAV
		"000e1c": 2024, // Hach
		"000e1d": 2025, // ARION
		"000e1e": 2026, // QLogic
		"000e22": 67,   // Private
		"000e23": 2027, // Incipient
		"000e24": 2028, // Huwell
		"000e25": 2029, // Hannae
		"000e26": 2030, // Gincom
		"000e27": 2031, // Crere
		"000e29": 2032, // Shester
		"000e2a": 67,   // Private
		"000e2b": 2033, // Safari
		"000e2c": 2034, // Netcodec
		"000e2e": 115,  // Edimax
		"000e30": 2035, // AERAS
		"000e33": 2036, // Shuko
		"000e35": 421,  // Intel
		"000e36": 2037, // HEINESYS
		"000e38": 2,    // Cisco
		"000e39": 2,    // Cisco
		"000e3b": 2038, // Hawking
		"000e3c": 2039, // Transact
		"000e3f": 864,  // Soronti
		"000e40": 76,   // Nortel
		"000e43": 2040, // G-Tek
		"000e4c": 2041, // Bermai
		"000e4d": 2042, // Numesa
		"000e4e": 2043, // Waveplus
		"000e4f": 2044, // Trajet
		"000e52": 2045, // Optium
		"000e53": 2046, // AV
		"000e55": 2047, // Auvitran
		"000e58": 2048, // Sonos
		"000e59": 2049, // Sagemcom
		"000e5a": 2050, // TELEFIELD
		"000e5c": 125,  // ARRIS
		"000e5e": 2051, // Raisecom
		"000e5f": 2052, // activ-net
		"000e61": 2053, // Microtrol
		"000e62": 76,   // Nortel
		"000e64": 2054, // Elphel
		"000e65": 2055, // TransCore
		"000e6a": 160,  // 3Com
		"000e6b": 2056, // Janitza
		"000e6d": 2057, // Murata
		"000e70": 2058, // in2
		"000e72": 2059, // CTS
		"000e73": 2060, // Tpack
		"000e77": 2061, // Decru
		"000e78": 2062, // Amtelco
		"000e79": 2063, // Ample
		"000e7a": 2064, // GemWon
		"000e7b": 35,   // Toshiba
		"000e7e": 2065, // ionSign
		"000e7f": 302,  // HP
		"000e80": 1142, // Thomson
		"000e81": 2066, // Devicescape
		"000e82": 2067, // Infinity
		"000e83": 2,    // Cisco
		"000e84": 2,    // Cisco
		"000e85": 2068, // Catalyst
		"000e86": 2069, // Alcatel
		"000e88": 2070, // Videotron
		"000e89": 2071, // Clematic
		"000e8a": 2072, // Avara
		"000e8b": 2073, // Astarte
		"000e8c": 300,  // Siemens
		"000e8e": 2074, // SparkLAN
		"000e8f": 2075, // Sercomm
		"000e90": 2076, // Ponico
		"000e94": 2077, // Maas
		"000e97": 2078, // Ultracker
		"000e9a": 2079, // BOE
		"000e9c": 2080, // Benchmark
		"000e9e": 2081, // Topfield
		"000ea0": 2082, // NetKlass
		"000ea2": 2083, // McAfee
		"000ea4": 2084, // Quantum
		"000ea5": 2085, // BLIP
		"000ea6": 1806, // ASUS
		"000ea7": 2086, // Endace
		"000eaa": 2087, // Scalent
		"000eab": 68,   // Cray
		"000eac": 2088, // Mintron
		"000ead": 2089, // Metanoia
		"000eae": 2090, // Gawell
		"000eaf": 2091, // Castel
		"000eb1": 2092, // Newcotech
		"000eb3": 302,  // HP
		"000eb5": 2093, // Ecastle
		"000eb6": 2094, // Riverbed
		"000eb7": 2095, // Knovative
		"000eb8": 2096, // Iiga
		"000eb9": 2097, // HASHIMOTO
		"000ebb": 2098, // Everbee
		"000ebd": 2099, // Burdick
		"000ebe": 2100, // B&B
		"000ebf": 2101, // Remsdaq
		"000ec0": 76,   // Nortel
		"000ec1": 2102, // MYNAH
		"000ec2": 2103, // Lowrance
		"000ec6": 2104, // Asix
		"000ec8": 2105, // Zoran
		"000ec9": 2106, // YOKO
		"000eca": 2107, // WTSS
		"000ecb": 2108, // VineSys
		"000ecc": 2109, // Tableau
		"000ecd": 2110, // Skov
		"000ed0": 2111, // Privaris
		"000ed2": 2112, // Filtronic
		"000ed3": 2113, // Epicenter
		"000ed5": 2114, // CoPAN
		"000ed6": 2,    // Cisco
		"000ed7": 2,    // Cisco
		"000ed9": 2115, // Aksys
		"000edb": 2116, // XiNCOM
		"000edc": 2117, // Tellion
		"000edd": 2118, // Shure
		"000ede": 2119, // REMEC
		"000edf": 2120, // PLX
		"000ee0": 2121, // Mcharge
		"000ee1": 2122, // ExtremeSpeed
		"000ee2": 2123, // Custom
		"000ee3": 2124, // Chiyu
		"000ee4": 2079, // BOE
		"000ee5": 2125, // bitWallet
		"000ee6": 2126, // Adimos
		"000ee7": 2127, // AAC
		"000ee8": 2128, // Zioncom
		"000eeb": 2129, // SandmartinElectronics
		"000eec": 2130, // Orban
		"000eed": 457,  // Nokia
		"000eef": 67,   // Private
		"000ef0": 2131, // Festo
		"000ef1": 2132, // Ezquest
		"000ef2": 2133, // Infinico
		"000ef3": 2134, // Smartlabs
		"000ef4": 2135, // Kasda
		"000ef5": 2136, // iPAC
		"000efa": 2137, // Optoway
		"000efb": 2138, // Macey
		"000efd": 2139, // Fujinon
		"000efe": 2140, // EndRun
		"000eff": 2141, // Megasolution
		"000f00": 2142, // Legra
		"000f01": 2143, // Digitalks
		"000f02": 2144, // Digicube
		"000f03": 2145, // CoM&C
		"000f04": 2146, // cim-usa
		"000f06": 76,   // Nortel
		"000f07": 2147, // Mangrove
		"000f08": 2148, // Indagon
		"000f09": 67,   // Private
		"000f0b": 2149, // Kentima
		"000f0c": 2150, // Synchronic
		"000f0d": 2151, // Hunt
		"000f0e": 2152, // WaveSplitter
		"000f10": 2153, // RDM
		"000f12": 2154, // Panasonic
		"000f13": 2155, // Nisca
		"000f14": 2156, // Mindray
		"000f15": 2157, // Icotera
		"000f1b": 2158, // Ego
		"000f1f": 102,  // Dell
		"000f20": 302,  // HP
		"000f22": 2159, // Helius
		"000f23": 2,    // Cisco
		"000f24": 2,    // Cisco
		"000f26": 2160, // WorldAccxx
		"000f27": 2161, // TEAL
		"000f28": 2162, // Itronix
		"000f29": 2163, // Augmentix
		"000f2a": 2164, // Cableware
		"000f2b": 2165, // Greenbell
		"000f2c": 2166, // Uplogix
		"000f2e": 2167, // Megapower
		"000f2f": 2168, // W-Linx
		"000f33": 2169, // DUALi
		"000f34": 2,    // Cisco
		"000f35": 2,    // Cisco
		"000f37": 2170, // Xambala
		"000f38": 2171, // Netstar
		"000f3a": 2172, // Hisharp
		"000f3c": 2173, // Endeleo
		"000f3d": 803,  // D-Link
		"000f3e": 2174, // CardioNet
		"000f41": 2175, // Zipher
		"000f42": 2176, // Xalyo
		"000f43": 2177, // Wasabi
		"000f44": 2178, // Tivella
		"000f45": 2179, // Stretch
		"000f46": 2180, // Sinar
		"000f47": 2181, // Robox
		"000f48": 2182, // Polypix
		"000f49": 2183, // Northover
		"000f4a": 2184, // Kyushu-kyohan
		"000f4b": 11,   // Oracle
		"000f4c": 2185, // Elextech
		"000f4d": 2186, // TalkSwitch
		"000f4e": 2187, // Cellink
		"000f50": 2188, // StreamScale
		"000f51": 2189, // Azul
		"000f53": 1524, // Solarflare
		"000f54": 2190, // Entrelogic
		"000f55": 2191, // Datawire
		"000f57": 2192, // CABLELOGIC
		"000f58": 2193, // Adder
		"000f59": 2194, // Phonak
		"000f5a": 2195, // Peribit
		"000f5d": 2196, // Genexis
		"000f5e": 2197, // Veo
		"000f5f": 2198, // Nicety
		"000f60": 2199, // Lifetron
		"000f61": 302,  // HP
		"000f63": 2200, // Obzerv
		"000f64": 2201, // D&R
		"000f65": 2202, // icube
		"000f66": 1783, // Linksys
		"000f6a": 76,   // Nortel
		"000f6b": 2203, // GateWare
		"000f6c": 2204, // ADDI-DATA
		"000f6d": 2205, // Midas
		"000f6e": 2206, // BBox
		"000f6f": 2207, // FTA
		"000f70": 2208, // Wintec
		"000f71": 2209, // Sanmei
		"000f72": 2210, // Sandburst
		"000f74": 2211, // Qamcom
		"000f77": 2212, // Dentum
		"000f78": 2213, // Datacap
		"000f7c": 2214, // ACTi
		"000f7d": 2215, // Xirrus
		"000f7e": 2216, // Ablerex
		"000f7f": 2217, // UBSTORAGE
		"000f83": 2218, // Brainium
		"000f84": 2219, // Astute
		"000f85": 2220, // ADDO-Japan
		"000f86": 2221, // BlackBerry
		"000f87": 2222, // Maxcess
		"000f88": 2223, // AMETEK
		"000f8a": 2224, // WideView
		"000f8c": 2225, // Gigawavetech
		"000f8f": 2,    // Cisco
		"000f90": 2,    // Cisco
		"000f91": 2226, // Aerotelecom
		"000f92": 2227, // Microhard
		"000f93": 2228, // Landis+Gyr
		"000f94": 2196, // Genexis
		"000f96": 18,   // Telco
		"000f97": 2229, // Avanex
		"000f98": 2230, // Avamax
		"000f9a": 2231, // Synchrony
		"000f9c": 2232, // Panduit
		"000f9d": 2233, // DisplayLink
		"000f9e": 2234, // Murrelektronik
		"000f9f": 125,  // ARRIS
		"000fa1": 2235, // Gigabit
		"000fa2": 2236, // 2xWireless
		"000fa3": 2237, // Alpha
		"000fa5": 2238, // BWA
		"000fa7": 2239, // Raptor
		"000fa8": 2240, // Photometrics
		"000faa": 2241, // Nexus
		"000fab": 2242, // Kyushu
		"000fad": 2243, // FMN
		"000fae": 2244, // E2O
		"000faf": 2245, // Dialog
		"000fb0": 358,  // Compal
		"000fb1": 2246, // Cognio
		"000fb3": 2247, // Actiontec
		"000fb4": 2248, // Timespace
		"000fb5": 1368, // Netgear
		"000fb6": 2249, // Europlex
		"000fb7": 2250, // Cavium
		"000fb8": 2251, // CallURL
		"000fba": 2252, // Tevebox
		"000fbc": 2253, // Onkey
		"000fbd": 2254, // MRV
		"000fbe": 2255, // e-w/you
		"000fc0": 2256, // DELCOMp
		"000fc1": 2257, // WAVE
		"000fc2": 2258, // Uniwell
		"000fc3": 2259, // PalmPalm
		"000fc4": 2260, // NST
		"000fc5": 2261, // KeyMed
		"000fc6": 2262, // Eurocom
		"000fc8": 2263, // Chantry
		"000fc9": 2264, // Allnet
		"000fcb": 160,  // 3Com
		"000fcc": 125,  // ARRIS
		"000fcd": 76,   // Nortel
		"000fce": 2265, // Kikusui
		"000fcf": 2266, // DataWind
		"000fd0": 2267, // Astri
		"000fd2": 2268, // EWA
		"000fd3": 2269, // Digium
		"000fd4": 2270, // Soundcraft
		"000fd6": 2271, // Sarotech
		"000fd8": 2272, // Force
		"000fda": 2273, // Yazaki
		"000fdb": 2274, // Westell
		"000fdd": 2275, // Sordin
		"000fde": 101,  // Sony
		"000fdf": 2276, // SOLOMON
		"000fe0": 2277, // NComputing
		"000fe4": 2278, // Pantech
		"000fe6": 2279, // MBTech
		"000fe7": 2280, // Lutron
		"000fe8": 2281, // Lobos
		"000fe9": 2282, // GW
		"000fea": 1929, // Giga-Byte
		"000fec": 2283, // ARKUS
		"000fed": 2284, // Anam
		"000fee": 2285, // XTecorporated
		"000ff0": 2286, // Sunray
		"000ff1": 2287, // nex-G
		"000ff2": 2288, // Loud
		"000ff5": 2289, // GN&S
		"000ff7": 2,    // Cisco
		"000ff8": 2,    // Cisco
		"000ff9": 2290, // Valcretec
		"000ffa": 2291, // Optinel
		"000ffe": 2292, // G-PRO
		"000fff": 2293, // Control4
		"001001": 1033, // Citel
		"001002": 2294, // Actia
		"001003": 2295, // Imatron
		"001007": 2,    // Cisco
		"001008": 2296, // Vienna
		"001009": 2297, // Horanet
		"00100b": 2,    // Cisco
		"00100c": 2298, // ITO
		"00100d": 2,    // Cisco
		"001010": 2299, // Initio
		"001011": 2,    // Cisco
		"001014": 2,    // Cisco
		"001015": 2300, // OOmon
		"001016": 2301, // T.Sqware
		"001018": 858,  // Broadcom
		"00101a": 2302, // PictureTel
		"00101b": 489,  // Cornet
		"00101d": 2303, // Winbond
		"00101e": 148,  // Matsushita
		"00101f": 2,    // Cisco
		"001021": 2304, // Encanto
		"001025": 2305, // Grayhill
		"001026": 2306, // Accelerated
		"001029": 2,    // Cisco
		"00102c": 2307, // Lasat
		"00102d": 89,   // Hitachi
		"00102f": 2,    // Cisco
		"001030": 2308, // EION
		"001031": 2309, // Objective
		"001032": 2310, // Alta
		"001033": 2311, // Accesslan
		"001034": 2312, // GNP
		"001035": 1115, // Elitegroup
		"001037": 2313, // CYQ've
		"001038": 43,   // Micro
		"001039": 2314, // Vectron
		"00103d": 2315, // Phasecom
		"00103e": 2316, // Netschools
		"00103f": 2317, // Tollgrade
		"001040": 2318, // Intermec
		"001042": 2319, // Alacritech
		"001043": 2320, // A2
		"001044": 2321, // InnoLabs
		"001045": 76,   // Nortel
		"001049": 2322, // ShoreTel
		"00104a": 2323, // Parvus
		"00104b": 160,  // 3Com
		"00104d": 2324, // Surtec
		"00104e": 2325, // Ceologic
		"00104f": 11,   // Oracle
		"001050": 2326, // Rion
		"001051": 2327, // Cmicro
		"001052": 2328, // Mettler-Toledo
		"001053": 2329, // Computer
		"001054": 2,    // Cisco
		"001056": 2330, // Sodick
		"001057": 2331, // Rebel.com
		"001058": 2332, // ArrowPoint
		"001059": 2333, // Diablo
		"00105a": 160,  // 3Com
		"00105e": 2334, // Spirent
		"001060": 2335, // Billionton
		"001061": 2336, // Hostlink
		"001064": 2337, // Dnpg
		"001065": 2338, // Radyne
		"001067": 306,  // Ericsson
		"001069": 2339, // Helioss
		"00106b": 2340, // Sonus
		"00106c": 2341, // EDNT
		"00106f": 2342, // Trenton
		"001071": 2343, // Advanet
		"001072": 2344, // GVN
		"001073": 2345, // Technobox
		"001074": 2346, // Aten
		"001075": 2347, // Segate
		"001076": 2348, // EUREM
		"001078": 2349, // Nuera
		"001079": 2,    // Cisco
		"00107a": 2350, // AmbiCom
		"00107b": 2,    // Cisco
		"00107c": 2351, // P-CoM
		"00107d": 1119, // Aurora
		"00107e": 2352, // BACHMANN
		"00107f": 2353, // Crestron
		"001080": 2354, // Metawave
		"001081": 2355, // DPS
		"001083": 302,  // HP
		"001084": 2356, // K-BOT
		"001085": 1213, // Polaris
		"001086": 2357, // ATTO
		"001087": 2358, // Xstreamis
		"001088": 2359, // American
		"001089": 2360, // WebSonic
		"00108a": 2361, // TeraLogic
		"00108c": 4,    // Fujitsu
		"00108f": 2239, // Raptor
		"001090": 2362, // Cimetrics
		"001092": 2363, // Netcore
		"001093": 2364, // CMS
		"001095": 1142, // Thomson
		"001096": 2365, // Tracewell
		"001098": 2366, // Starnet
		"001099": 2367, // InnoMedia
		"00109a": 2368, // Netline
		"00109b": 129,  // Emulex
		"00109c": 2369, // M-System
		"00109d": 2370, // Clarinet
		"00109e": 2371, // Aware
		"00109f": 2372, // Pavo
		"0010a0": 2373, // Innovex
		"0010a2": 2374, // TNS
		"0010a3": 2375, // Omnitronix
		"0010a4": 788,  // Xircom
		"0010a6": 2,    // Cisco
		"0010a7": 256,  // Unex
		"0010a8": 2376, // Reliance
		"0010a9": 2377, // Adhoc
		"0010aa": 2378, // MEDIA4
		"0010ac": 2379, // Imci
		"0010af": 2380, // TAC
		"0010b0": 2381, // Meridian
		"0010b1": 2382, // FOR-A
		"0010b4": 2383, // Atmosphere
		"0010b5": 145,  // Accton
		"0010b6": 2384, // Entrata
		"0010b7": 2385, // Coyote
		"0010b9": 2386, // Maxtor
		"0010ba": 2387, // Martinho-Davis
		"0010be": 2388, // March
		"0010c0": 2389, // ARMA
		"0010c2": 2390, // Willnet
		"0010c3": 2391, // CSI-Control
		"0010c5": 2392, // Protocol
		"0010c8": 2393, // Communications
		"0010c9": 2394, // Mitsubishi
		"0010ca": 18,   // Telco
		"0010ce": 2395, // Volamp
		"0010cf": 2396, // Fiberlane
		"0010d0": 2397, // Witcom
		"0010d3": 2398, // Grips
		"0010d4": 2399, // Storage
		"0010d6": 2400, // Exelis
		"0010d7": 2401, // Argosy
		"0010d8": 2402, // Calista
		"0010da": 2403, // Kollmorgen
		"0010db": 826,  // Juniper
		"0010dc": 1812, // Micro-Star
		"0010df": 2404, // Rise
		"0010e0": 11,   // Oracle
		"0010e1": 2405, // S.I
		"0010e2": 2406, // ArrayComm
		"0010e3": 302,  // HP
		"0010e4": 2407, // NSI
		"0010e7": 700,  // Breezecom
		"0010e8": 2408, // Telocityorporated
		"0010e9": 2409, // Raidtec
		"0010ea": 1179, // Adept
		"0010eb": 2410, // Selsius
		"0010ec": 2411, // Rpcg
		"0010ed": 2412, // Sundance
		"0010ee": 2413, // CTI
		"0010ef": 2414, // Dbtel
		"0010f1": 2415, // I-O
		"0010f2": 2416, // Antec
		"0010f3": 2417, // Nexcom
		"0010f4": 2418, // Vertical
		"0010f5": 2419, // Amherst
		"0010f6": 2,    // Cisco
		"0010f7": 2420, // IRIICHI
		"0010f8": 313,  // Texio
		"0010f9": 2421, // Unique
		"0010fa": 545,  // Apple
		"0010fb": 2422, // Zida
		"0010fc": 2423, // Broadband
		"0010fd": 2424, // Cocom
		"0010ff": 2,    // Cisco
		"001100": 56,   // Schneider
		"001101": 2425, // CET
		"001104": 2426, // Telexy
		"001105": 2427, // Sunplus
		"001106": 300,  // Siemens
		"001107": 2428, // RGB
		"001109": 1812, // Micro-Star
		"00110a": 302,  // HP
		"00110b": 2429, // Franklin
		"00110d": 2430, // SANBlaze
		"00110f": 2431, // netplat
		"001110": 2432, // Maxanna
		"001111": 421,  // Intel
		"001114": 2433, // EverFocus
		"001115": 2434, // EPIN
		"001117": 2435, // Cesnet
		"001119": 2436, // Solteras
		"00111a": 125,  // ARRIS
		"00111c": 2437, // Pleora
		"00111d": 2438, // Hectrix
		"00111f": 2439, // Doremi
		"001120": 2,    // Cisco
		"001121": 2,    // Cisco
		"001122": 2440, // CIMSYS
		"001123": 2441, // Appointech
		"001124": 545,  // Apple
		"001125": 372,  // IBM
		"001126": 2442, // Venstar
		"001127": 2443, // TASI
		"001128": 2444, // Streamit
		"00112a": 2445, // Niko
		"00112b": 2446, // NetModule
		"00112c": 2447, // IZT
		"00112d": 2448, // iPulse
		"00112e": 2449, // Ceicom
		"00112f": 1806, // ASUS
		"001131": 2450, // Unatech
		"001132": 2451, // Synology
		"001133": 300,  // Siemens
		"001134": 2452, // MediaCell
		"001135": 2453, // Grandeye
		"001138": 2454, // Taishin
		"00113a": 2455, // Shinboram
		"00113b": 2456, // Micronet
		"00113c": 2457, // Micronas
		"00113e": 2458, // JL
		"001140": 2459, // Nanometrics
		"001141": 2460, // GoodMan
		"001142": 2461, // e-SMARTCOM
		"001143": 102,  // Dell
		"001144": 2462, // Assurance
		"001145": 2463, // ValuePoint
		"001146": 2464, // Telecard-Pribor
		"001147": 2465, // Secom-Industry
		"001149": 2466, // Proliphix
		"00114a": 2467, // KAYABA
		"00114b": 2468, // Francotyp-Postalia
		"001150": 2469, // Belkin
		"001151": 2470, // Mykotronx
		"001152": 2471, // Eidsvoll
		"001154": 2472, // Webpro
		"001155": 2473, // Sevis
		"001158": 76,   // Nortel
		"001159": 2474, // Matisse
		"00115b": 1115, // Elitegroup
		"00115c": 2,    // Cisco
		"00115d": 2,    // Cisco
		"001160": 2475, // ARTDIO
		"001161": 2476, // NetStreams
		"001164": 2477, // ACARD
		"001165": 2478, // ZNYX
		"001166": 2479, // Taelim
		"001168": 2480, // HomeLogic
		"00116a": 2481, // Domo
		"00116e": 2482, // Peplink
		"00116f": 2483, // Netforyou
		"001171": 2484, // DEXTER
		"001172": 2485, // Cotron
		"001174": 2486, // Mojo
		"001175": 421,  // Intel
		"001176": 2487, // Intellambda
		"001177": 2488, // Coaxial
		"001178": 2489, // Chiron
		"001179": 2490, // Singular
		"00117a": 2491, // Singim
		"00117c": 2492, // e-zy.net
		"00117e": 2493, // Midmark
		"001180": 125,  // ARRIS
		"001181": 2494, // InterEnergy
		"001185": 302,  // HP
		"001186": 1040, // Prime
		"001187": 2495, // Category
		"001188": 312,  // Enterasys
		"001189": 2496, // Aerotech
		"00118a": 2497, // Viewtran
		"001192": 2,    // Cisco
		"001193": 2,    // Cisco
		"001195": 803,  // D-Link
		"001196": 2498, // Actuality
		"001197": 2499, // Monitoring
		"001199": 2500, // 2wcom
		"00119b": 2501, // Telesynergy
		"00119d": 2502, // Diginfo
		"00119f": 457,  // Nokia
		"0011a0": 2503, // Vtech
		"0011a3": 2504, // LanReady
		"0011a4": 2505, // JStream
		"0011a5": 2506, // Fortuna
		"0011a6": 2507, // Sypixx
		"0011a8": 2508, // Quest
		"0011a9": 2509, // MOIMSTONE
		"0011aa": 2510, // Uniclass
		"0011ab": 2511, // Trustable
		"0011ac": 2512, // Simtec
		"0011ae": 125,  // ARRIS
		"0011af": 2513, // Medialink-i
		"0011b0": 2514, // Fortelink
		"0011b1": 2515, // BlueExpert
		"0011b2": 2516, // 2001
		"0011b3": 2517, // Yoshimiya
		"0011b6": 1879, // Open
		"0011ba": 2518, // Elexol
		"0011bb": 2,    // Cisco
		"0011bc": 2,    // Cisco
		"0011c0": 2519, // Aday
		"0011c5": 2520, // TEN
		"0011c6": 732,  // Seagate
		"0011c8": 2521, // Powercom
		"0011c9": 2522, // MTT
		"0011cb": 2523, // Jacobsons
		"0011cd": 2524, // Axsun
		"0011ce": 2525, // Ubisense
		"0011d4": 2526, // NetEnrich
		"0011d6": 2527, // HandEra
		"0011d7": 2528, // eWerks
		"0011d8": 1806, // ASUS
		"0011d9": 2529, // TiVo
		"0011da": 2530, // Vivaas
		"0011db": 2531, // Land-Cellular
		"0011de": 2532, // Eurilogic
		"0011e0": 2533, // U-MEDIA
		"0011e3": 1142, // Thomson
		"0011e4": 2534, // Danelec
		"0011e5": 2535, // KCodes
		"0011e8": 2536, // Tixi.com
		"0011e9": 2537, // Starnex
		"0011ea": 2538, // IWICS
		"0011ec": 2539, // Avix
		"0011ee": 2540, // Estari
		"0011f0": 2541, // Wideful
		"0011f1": 2542, // QinetiQ
		"0011f3": 2543, // NeoMedia
		"0011f4": 2544, // woori-net
		"0011f5": 2545, // Askey
		"0011f8": 2546, // AIRAYA
		"0011f9": 76,   // Nortel
		"0011fa": 2547, // Rane
		"0011fb": 2548, // Heidelberg
		"0011fc": 1596, // HARTING
		"0011fd": 2549, // Korg
		"001200": 2,    // Cisco
		"001201": 2,    // Cisco
		"001203": 2550, // ActivNetworks
		"001204": 2551, // u10
		"001205": 2552, // Terrasat
		"001206": 2553, // iQuest
		"001209": 2554, // Fastrax
		"00120b": 2555, // Chinasys
		"00120c": 2556, // CE-Infosys
		"00120e": 2557, // AboCom
		"001210": 2558, // WideRay
		"001212": 2559, // PLUS
		"001213": 2560, // Metrohm
		"001215": 2561, // iStor
		"001217": 1783, // Linksys
		"001218": 2562, // ARUZE
		"00121c": 2563, // Parrot
		"00121d": 2564, // Netfabric
		"00121e": 826,  // Juniper
		"001220": 2565, // Cadco
		"001222": 2566, // Skardin
		"001223": 2567, // Pixim
		"001224": 2568, // NexQL
		"001225": 125,  // ARRIS
		"001228": 2569, // Data
		"001229": 2570, // BroadEasy
		"00122b": 2571, // Virbiage
		"00122d": 2572, // SiNett
		"001232": 2573, // LeWiz
		"001235": 2574, // Andrew
		"001236": 2575, // ConSentry
		"001238": 2576, // SetaBox
		"00123a": 2577, // Posystech
		"00123d": 2578, // GES
		"00123e": 2579, // ERUNE
		"00123f": 102,  // Dell
		"001240": 2580, // Amoi
		"001243": 2,    // Cisco
		"001244": 2,    // Cisco
		"001246": 2581, // T.O.M
		"001247": 152,  // Samsung
		"001248": 102,  // Dell
		"00124c": 2582, // BBWM
		"00124d": 2583, // Inducon
		"00124f": 2584, // nVent
		"001251": 2585, // Silink
		"001252": 2586, // Citronix
		"001253": 2587, // AudioDev
		"001254": 2588, // Spectra
		"001255": 2589, // NetEffect
		"001256": 869,  // LG
		"001257": 2590, // LeapComm
		"001259": 2591, // Thermo
		"00125a": 612,  // Microsoft
		"00125b": 2592, // Kaimei
		"00125d": 2593, // CyberNet
		"00125e": 2594, // Caen
		"00125f": 2595, // AWIND
		"001261": 2596, // Adaptix
		"001262": 457,  // Nokia
		"001264": 2597, // daum
		"001265": 2598, // Enerdyne
		"001267": 2154, // Panasonic
		"001269": 2599, // Value
		"00126a": 2600, // OPTOELECTRONICS
		"00126b": 2601, // Ascalade
		"00126f": 2602, // Rayson
		"001272": 2603, // Redux
		"001273": 2604, // Stoke
		"001275": 2605, // Sentilla
		"001277": 2606, // Korenix
		"001279": 302,  // HP
		"00127a": 2607, // Sanyu
		"00127c": 2608, // Swegon
		"00127d": 2609, // MobileAria
		"00127f": 2,    // Cisco
		"001280": 2,    // Cisco
		"001282": 2610, // Qovia
		"001283": 76,   // Nortel
		"001285": 2611, // Gizmondo
		"001286": 2612, // Endevco
		"001288": 1939, // 2Wire
		"00128a": 125,  // ARRIS
		"00128b": 2613, // Sensory
		"00128f": 2614, // Montilio
		"001292": 2615, // Griffin
		"001295": 2616, // Aiware
		"001296": 2617, // Addlogix
		"001297": 2618, // O2Micro
		"00129a": 2619, // IRT
		"00129b": 2620, // E2S
		"00129c": 2621, // Yulinet
		"00129e": 2622, // Surf
		"00129f": 2623, // RAE
		"0012a1": 2624, // BluePacket
		"0012a2": 2625, // Vita
		"0012a4": 2626, // ThingMagic
		"0012a7": 2627, // ISR
		"0012a8": 2628, // intec
		"0012a9": 160,  // 3Com
		"0012aa": 2629, // IEE
		"0012ab": 2630, // WiLife
		"0012ac": 2631, // Ontimetek
		"0012ad": 2632, // Vivavis
		"0012af": 2633, // ELPRO
		"0012b2": 2634, // Avolites
		"0012b5": 2635, // Vialta
		"0012ba": 2636, // FSI
		"0012bc": 2637, // Echolab
		"0012bd": 2638, // Avantec
		"0012be": 2639, // Astek
		"0012bf": 2640, // Arcadyan
		"0012c0": 2641, // HotLava
		"0012c2": 406,  // Apex
		"0012c4": 2642, // Viseon
		"0012c5": 2643, // V-Show
		"0012c8": 2644, // Perfect
		"0012c9": 125,  // ARRIS
		"0012cb": 34,   // CSS
		"0012cc": 2645, // Bitatek
		"0012cd": 2646, // ASEM
		"0012cf": 145,  // Accton
		"0012d0": 2647, // Gossen-Metrawatt-GmbH
		"0012d3": 2648, // Zetta
		"0012d4": 850,  // Princeton
		"0012d7": 2649, // Invento
		"0012d9": 2,    // Cisco
		"0012da": 2,    // Cisco
		"0012dc": 2650, // SunCorp
		"0012df": 2651, // Novomatic
		"0012e0": 2652, // Codan
		"0012e1": 2653, // Alliant
		"0012e2": 2654, // ALAXALA
		"0012e6": 2655, // Spectec
		"0012e9": 2656, // Abbey
		"0012ea": 2657, // Trane
		"0012eb": 2658, // PDH
		"0012ee": 101,  // Sony
		"0012ef": 2659, // OneAccess
		"0012f0": 421,  // Intel
		"0012f1": 2660, // Ifotec
		"0012f2": 90,   // Brocade
		"0012f3": 2661, // ConnectBlue
		"0012f4": 2662, // Belco
		"0012f6": 2663, // MDK
		"0012fa": 2664, // THX
		"0012fb": 152,  // Samsung
		"0012fe": 2665, // Lenovo
		"0012ff": 2666, // Lely
		"001300": 2667, // IT-Factory
		"001302": 421,  // Intel
		"001303": 2668, // GateConnect
		"001304": 2669, // Flaircomm
		"001305": 2670, // Epicom
		"001307": 2671, // Paravirtual
		"00130a": 76,   // Nortel
		"001310": 1783, // Linksys
		"001311": 125,  // ARRIS
		"001312": 2672, // Amedia
		"001314": 2673, // Asiamajor
		"001315": 101,  // Sony
		"001318": 2674, // DGSTATION
		"001319": 2,    // Cisco
		"00131a": 2,    // Cisco
		"00131c": 2675, // LiteTouch
		"00131d": 2676, // Scanvaegt
		"001320": 421,  // Intel
		"001321": 302,  // HP
		"001322": 2677, // DAQ
		"001323": 2678, // Cap
		"001324": 56,   // Schneider
		"001325": 2679, // Cortina
		"001326": 2680, // ECM
		"001329": 2681, // VSST
		"00132d": 2682, // iWise
		"00132f": 2683, // Interactek
		"001333": 2684, // BaudTec
		"001334": 2685, // Arkados
		"001338": 2686, // Fresenius-Vial
		"00133a": 2687, // VadaTech
		"00133c": 2688, // Quintron
		"00133e": 2689, // MetaSwitch
		"001342": 2690, // Vision
		"001343": 148,  // Matsushita
		"001344": 2691, // Fargo
		"001345": 1855, // Eaton
		"001346": 803,  // D-Link
		"001348": 2692, // Artila
		"001349": 2693, // ZyXEL
		"00134a": 2694, // Engim
		"00134b": 2695, // ToGoldenNet
		"00134c": 2696, // YDT
		"00134d": 2697, // Inepro
		"00134e": 2698, // Valox
		"001352": 2699, // Naztec
		"001354": 2700, // Zcomax
		"001357": 2701, // Soyal
		"001358": 2702, // Realm
		"001359": 2703, // ProTelevision
		"00135c": 2704, // OnSite
		"00135d": 2705, // NTTPC
		"00135e": 2706, // EAB/RWI/K
		"00135f": 2,    // Cisco
		"001360": 2,    // Cisco
		"001361": 2707, // Biospace
		"001363": 2708, // Verascape
		"001364": 2709, // Paradigm
		"001365": 76,   // Nortel
		"001366": 2710, // Neturity
		"001367": 2711, // Narayon
		"001368": 2712, // Saab
		"001369": 2713, // Honda
		"00136b": 2714, // E-TEC
		"00136c": 2715, // TomTom
		"00136d": 2716, // Tentaculus
		"00136e": 2717, // Techmetro
		"00136f": 2718, // PacketMotion
		"001370": 457,  // Nokia
		"001371": 125,  // ARRIS
		"001372": 102,  // Dell
		"001373": 2719, // BLwave
		"001374": 535,  // Atheros
		"001376": 2720, // Tabor
		"001377": 152,  // Samsung
		"001378": 2721, // Qsan
		"00137a": 2722, // Netvox
		"00137b": 2723, // Movon
		"00137c": 2724, // Kaicom
		"00137d": 2725, // Dynalab
		"00137e": 2726, // CorEdge
		"00137f": 2,    // Cisco
		"001380": 2,    // Cisco
		"001381": 2727, // CHIPS
		"001382": 2728, // Cetacea
		"001383": 2729, // Application
		"001385": 2730, // Add-On
		"001386": 2731, // ABB/Totalflow
		"001387": 2732, // 27M
		"00138b": 2733, // Phantom
		"00138c": 2734, // Kumyoung
		"00138d": 2735, // Kinghold
		"00138f": 1666, // Asiarock
		"001390": 2736, // Termtek
		"001391": 2737, // Ouen
		"001392": 2738, // Ruckus
		"001393": 2739, // Panta
		"001394": 2740, // Infohand
		"001395": 2741, // Congatec
		"001397": 11,   // Oracle
		"001398": 2742, // TrafficSim
		"001399": 2743, // STAC
		"00139b": 2744, // ioIMAGE
		"00139c": 2745, // Exavera
		"00139e": 2746, // Ciara
		"0013a0": 2747, // ALGOSYSTEM
		"0013a1": 2748, // Crow
		"0013a2": 2749, // MaxStream
		"0013a3": 300,  // Siemens
		"0013a4": 2750, // KeyEye
		"0013a5": 2751, // General
		"0013a6": 2752, // Extricom
		"0013a8": 2753, // Tanisys
		"0013a9": 101,  // Sony
		"0013ab": 2754, // Telemotive
		"0013ac": 2755, // Sunmyung
		"0013ad": 2756, // Sendo
		"0013ae": 2757, // Radiance
		"0013af": 2758, // NUMA
		"0013b0": 2759, // Jablotron
		"0013b2": 2760, // Carallon
		"0013b3": 2761, // Ecom
		"0013b5": 2762, // Wavesat
		"0013b8": 2763, // RyCo
		"0013b9": 2764, // BM
		"0013ba": 2765, // ReadyLinks
		"0013bb": 2766, // Smartvue
		"0013bc": 2767, // Artimi
		"0013bd": 2768, // Hymatom
		"0013c2": 2769, // WACOM
		"0013c3": 2,    // Cisco
		"0013c4": 2,    // Cisco
		"0013c6": 2770, // OpenGear
		"0013c7": 2771, // IONOS
		"0013ca": 2772, // ATX
		"0013cd": 2773, // MTI
		"0013ce": 421,  // Intel
		"0013cf": 2774, // 4Access
		"0013d3": 1812, // Micro-Star
		"0013d4": 1806, // ASUS
		"0013d5": 1586, // RuggedCom
		"0013d7": 2775, // SPIDCOM
		"0013da": 2776, // Diskware
		"0013dc": 2777, // Ibtek
		"0013de": 2778, // Adapt4
		"0013df": 2779, // Ryvor
		"0013e0": 2057, // Murata
		"0013e1": 2780, // Iprobe
		"0013e2": 2781, // GeoVision
		"0013e3": 2782, // CoVi
		"0013e4": 2783, // Yangjae
		"0013e5": 2784, // Tenosys
		"0013e6": 2785, // Technolution
		"0013e7": 2786, // Halcro
		"0013e8": 421,  // Intel
		"0013e9": 2787, // VeriWave
		"0013ea": 2788, // Kamstrup
		"0013eb": 2789, // Sysmaster
		"0013ed": 2790, // Psia
		"0013f1": 2791, // AMOD
		"0013f2": 2792, // Klas
		"0013f3": 2793, // Giga-byte
		"0013f4": 2794, // Psitek
		"0013f5": 2795, // Akimbi
		"0013f6": 2796, // Cintech
		"0013f7": 741,  // SMC
		"0013f9": 2797, // Cavera
		"0013fa": 2798, // LifeSize
		"0013fc": 2799, // SiCortex
		"0013fd": 457,  // Nokia
		"0013fe": 2800, // Grandtec
		"001401": 2801, // Rivertree
		"001402": 2802, // kk-electronic
		"001403": 2803, // Renasis
		"001404": 125,  // ARRIS
		"001405": 2804, // OpenIB
		"001406": 2805, // Go
		"001408": 2806, // Eka
		"00140a": 2807, // WEPIO
		"00140b": 1979, // First
		"00140d": 76,   // Nortel
		"00140e": 76,   // Nortel
		"001412": 2808, // S-TEC
		"001416": 2809, // Scosche
		"001418": 2810, // C4Line
		"001419": 2811, // Sidsa
		"00141a": 2812, // Deicy
		"00141b": 2,    // Cisco
		"00141c": 2,    // Cisco
		"00141f": 2813, // SunKwang
		"001422": 102,  // Dell
		"001426": 2814, // NL
		"001427": 2815, // JazzMutant
		"001428": 2816, // Vocollect
		"00142a": 1115, // Elitegroup
		"00142b": 2817, // Edata
		"00142c": 2818, // Koncept
		"00142d": 2819, // Toradex
		"00142f": 2820, // Savvius
		"001430": 2821, // ViPowER
		"001431": 2822, // PDL
		"001433": 2823, // Empower
		"001434": 2824, // Keri
		"001435": 2825, // CityCom
		"001437": 2826, // GSTeletech
		"001438": 302,  // HP
		"00143b": 2827, // Sensovation
		"00143d": 2828, // Aevoe
		"00143e": 2829, // AirLink
		"00143f": 2830, // Hotway
		"001440": 2831, // ATOMIC
		"001442": 2832, // Atto
		"001443": 2833, // Consultronics
		"001444": 2834, // Grundfos
		"001446": 2835, // SuperVision
		"001447": 2836, // BOAZ
		"00144b": 2837, // Hifn
		"00144d": 1861, // Intelligent
		"00144e": 2838, // Srisa
		"00144f": 11,   // Oracle
		"001450": 2839, // Heim
		"001451": 545,  // Apple
		"001452": 2840, // Calculex
		"001453": 1703, // Advantech
		"001454": 2841, // Symwave
		"001455": 2842, // Coder
		"001456": 2843, // Edge
		"001457": 2844, // T-Vips
		"001459": 2845, // Moram
		"00145b": 2846, // SeekerNet
		"00145d": 2847, // WJ
		"00145e": 372,  // IBM
		"00145f": 2848, // Aditec
		"001460": 2849, // Kyocera
		"001461": 2850, // Corona
		"001462": 2851, // Digiwell
		"001464": 2852, // Cryptosoft
		"001467": 2853, // ArrowSpan
		"001468": 2854, // CelPlan
		"001469": 2,    // Cisco
		"00146a": 2,    // Cisco
		"00146b": 2855, // Anagran
		"00146c": 1368, // Netgear
		"00146d": 2856, // RF
		"00146f": 2857, // Kohler
		"001470": 2858, // Prokom
		"001473": 2859, // Bookham
		"001474": 2860, // K40
		"001475": 2861, // Wiline
		"001476": 2862, // MultiCom
		"001477": 1585, // Trilliant
		"001478": 1595, // TP-Link
		"00147a": 2863, // Eubus
		"00147b": 2864, // Iteris
		"00147c": 160,  // 3Com
		"00147e": 2865, // InnerWireless
		"001481": 2866, // Multilink
		"001482": 1119, // Aurora
		"001483": 2867, // eXS
		"001484": 2868, // Cermate
		"001485": 1929, // Giga-Byte
		"001488": 2869, // Akorri
		"00148b": 2870, // Globo
		"00148f": 2871, // Protronic
		"001490": 2872, // ASP
		"001491": 2873, // Daniels
		"001492": 2874, // Liteon
		"001493": 2875, // Systimax
		"001494": 2876, // ESU
		"001495": 1939, // 2Wire
		"001496": 2877, // Phonic
		"001497": 2878, // ZHIYUAN
		"001499": 2879, // Helicomm
		"00149a": 125,  // ARRIS
		"00149b": 2880, // Nokota
		"00149c": 2881, // HF
		"00149e": 2882, // UbONE
		"0014a0": 2883, // Accsense
		"0014a1": 801,  // Synchronous
		"0014a3": 2884, // Vitelec
		"0014a5": 1450, // Gemtek
		"0014a6": 2885, // Teranetics
		"0014a7": 457,  // Nokia
		"0014a8": 2,    // Cisco
		"0014a9": 2,    // Cisco
		"0014ab": 2886, // Senhai
		"0014ae": 2887, // Wizlogics
		"0014b2": 2888, // mCubelogics
		"0014b3": 2889, // CoreStar
		"0014b5": 2890, // Physiometrix
		"0014b6": 2891, // Enswer
		"0014b8": 2892, // Hill-Rom
		"0014bd": 2893, // IncNETWORKS
		"0014be": 2894, // Wink
		"0014bf": 1783, // Linksys
		"0014c2": 302,  // HP
		"0014c3": 732,  // Seagate
		"0014c4": 2895, // Vitelcom
		"0014c5": 2896, // Alive
		"0014c6": 2897, // Quixant
		"0014c7": 76,   // Nortel
		"0014c8": 2898, // Contemporary
		"0014c9": 90,   // Brocade
		"0014cb": 2899, // LifeSync
		"0014cc": 2900, // Zetec
		"0014cd": 2901, // DigitalZone
		"0014ce": 2902, // NF
		"0014cf": 2903, // INVISIO
		"0014d0": 2904, // BTI
		"0014d1": 2905, // TRENDnet
		"0014d3": 2906, // Sepsa
		"0014d4": 2907, // K
		"0014d6": 2908, // Jeongmin
		"0014d7": 2909, // Datastore
		"0014d8": 2910, // bio-logic
		"0014dd": 2911, // Covergence
		"0014df": 2912, // HI-P
		"0014e0": 2913, // LET'S
		"0014e2": 2914, // datacom
		"0014e3": 2915, // mm-lab
		"0014e4": 2916, // Infinias
		"0014e5": 2917, // Alticast
		"0014e7": 2918, // Stolinx
		"0014e8": 125,  // ARRIS
		"0014e9": 2919, // Nortech
		"0014eb": 2920, // AwarePoint
		"0014ed": 2921, // Airak
		"0014ee": 123,  // WD
		"0014ef": 2922, // TZero
		"0014f1": 2,    // Cisco
		"0014f2": 2,    // Cisco
		"0014f3": 2923, // ViXS
		"0014f6": 826,  // Juniper
		"0014f7": 2924, // CREVIS
		"0014fb": 2925, // Technical
		"0014fc": 2926, // Extandon
		"0014fd": 2927, // Thecus
		"0014fe": 2928, // Artech
		"001500": 421,  // Intel
		"001501": 2929, // LexBox
		"001502": 2930, // BETA
		"001505": 2247, // Actiontec
		"001509": 2931, // Plus
		"00150a": 2932, // Sonoa
		"00150c": 621,  // AVM
		"00150e": 2933, // Openbrain
		"00150f": 2934, // mingjong
		"001510": 2935, // Techsphere
		"001515": 2936, // Leipold+Co.GmbH
		"001516": 2937, // Uriel
		"001517": 421,  // Intel
		"00151a": 472,  // Hunter
		"00151b": 2938, // Isilon
		"00151c": 2939, // Leneco
		"00151d": 2940, // M2I
		"001520": 2941, // Radiocrafts
		"001521": 2942, // Horoquartz
		"001523": 2943, // Meteor
		"001524": 2944, // Numatics
		"001526": 2945, // Remote
		"001529": 2946, // N3
		"00152a": 457,  // Nokia
		"00152b": 2,    // Cisco
		"00152c": 2,    // Cisco
		"00152d": 2947, // TenX
		"00152e": 2948, // PacketHop
		"00152f": 125,  // ARRIS
		"001530": 102,  // Dell
		"001531": 2949, // Kocom
		"001533": 2950, // Nadam
		"001535": 2951, // OTE
		"001536": 2952, // Powertech
		"001537": 2953, // Ventus
		"001538": 2954, // RFID
		"00153c": 2955, // Kprotech
		"001540": 76,   // Nortel
		"001541": 2956, // StrataLight
		"001544": 2957, // CoM.s.a.t
		"001545": 2958, // SEECODE
		"001547": 2959, // AiZen
		"001548": 2960, // Cube
		"00154a": 2961, // Wanshih
		"00154c": 2962, // Saunders
		"00154d": 2963, // Netronome
		"00154e": 2964, // IEC
		"001550": 2965, // Nits
		"001551": 2966, // RadioPulse
		"001552": 2967, // Wi-Gear
		"001553": 2968, // Cytyc
		"001555": 2969, // DFM
		"001556": 2049, // Sagemcom
		"001557": 45,   // Olivetti
		"001558": 221,  // Foxconn
		"001559": 2970, // Securaplane
		"00155b": 2971, // Sampo
		"00155d": 612,  // Microsoft
		"00155f": 2972, // GreenPeak
		"001560": 302,  // HP
		"001561": 2973, // JJPlus
		"001562": 2,    // Cisco
		"001563": 2,    // Cisco
		"001566": 2974, // A-First
		"001567": 2975, // RADWIN
		"001568": 2976, // Dilithium
		"00156a": 2977, // DG2L
		"00156b": 2978, // Perfisans
		"00156d": 2979, // Ubiquiti
		"00156f": 2980, // Xiranet
		"001570": 768,  // Zebra
		"001571": 2981, // Nolan
		"001572": 2982, // Red-Lemon
		"001573": 2983, // NewSoft
		"001575": 2984, // Nevis
		"00157b": 2985, // Leuze
		"00157c": 2986, // Dave
		"00157d": 2987, // Posdata
		"00157f": 2988, // ChuanG
		"001580": 2989, // U-WAY
		"001581": 2990, // MAKUS
		"001583": 2991, // IVT
		"001589": 2992, // D-MAX
		"00158a": 363,  // SURECOM
		"00158d": 2993, // Jennic
		"00158e": 2994, // Plustek.Inc
		"001590": 2995, // Hectronic
		"001591": 2996, // RLW
		"001593": 2997, // U4EA
		"001594": 2998, // Bixolon
		"001596": 125,  // ARRIS
		"001599": 152,  // Samsung
		"00159a": 125,  // ARRIS
		"00159b": 76,   // Nortel
		"00159f": 2999, // Terascala
		"0015a0": 457,  // Nokia
		"0015a1": 3000, // ECA-Sinters
		"0015a2": 125,  // ARRIS
		"0015a3": 125,  // ARRIS
		"0015a4": 125,  // ARRIS
		"0015a5": 49,   // DCI
		"0015a6": 1565, // Digital
		"0015a7": 3001, // Robatech
		"0015a8": 125,  // ARRIS
		"0015aa": 3002, // Rextechnik
		"0015ac": 3003, // Capelon
		"0015ad": 3004, // Accedian
		"0015af": 3005, // Azurewave
		"0015b0": 3006, // Autotelenet
		"0015b1": 3007, // Ambient
		"0015b2": 587,  // Advanced
		"0015b3": 3008, // Caretech
		"0015b6": 3009, // ShinMaywa
		"0015b7": 35,   // Toshiba
		"0015b8": 1024, // Tahoe
		"0015b9": 152,  // Samsung
		"0015ba": 3010, // iba
		"0015bc": 3011, // Develco
		"0015be": 3012, // Iqua
		"0015bf": 3013, // technicob
		"0015c1": 101,  // Sony
		"0015c4": 3014, // Flovel
		"0015c5": 102,  // Dell
		"0015c6": 2,    // Cisco
		"0015c7": 2,    // Cisco
		"0015c8": 3015, // FlexiPanel
		"0015c9": 3016, // Gumstix
		"0015ca": 3017, // TeraRecon
		"0015cb": 2622, // Surf
		"0015cc": 3018, // Uquest
		"0015cd": 3019, // Exartech
		"0015ce": 125,  // ARRIS
		"0015cf": 125,  // ARRIS
		"0015d0": 125,  // ARRIS
		"0015d1": 125,  // ARRIS
		"0015d2": 3020, // Xantech
		"0015d4": 3021, // Emitor
		"0015d5": 3022, // Nicevt
		"0015d7": 3023, // Reti
		"0015d8": 3024, // Interlink
		"0015d9": 3025, // PKC
		"0015db": 3026, // Canesta
		"0015dc": 3027, // KT&C
		"0015de": 457,  // Nokia
		"0015e0": 306,  // Ericsson
		"0015e1": 3028, // Picochip
		"0015e3": 3029, // Dream
		"0015e5": 3030, // Cheertek
		"0015e8": 76,   // Nortel
		"0015e9": 803,  // D-Link
		"0015ea": 3031, // Tellumat
		"0015eb": 3032, // ZTE
		"0015f0": 3033, // EGO
		"0015f1": 3034, // KYLINK
		"0015f2": 1806, // ASUS
		"0015f3": 3035, // Peltor
		"0015f4": 3036, // Eventide
		"0015f6": 3037, // Science
		"0015f7": 3038, // Wintecronics
		"0015f8": 3039, // Kingtronics
		"0015f9": 2,    // Cisco
		"0015fa": 2,    // Cisco
		"001601": 1077, // Buffalo
		"001602": 3040, // Ceyon
		"001603": 3041, // CoOLKSKY
		"001604": 3042, // Sigpro
		"001606": 3043, // Ideal
		"001607": 3044, // Curves
		"001608": 3045, // Sequans
		"001609": 3046, // Unitech
		"00160a": 3047, // SWEEX
		"00160b": 3048, // TVWorks
		"00160e": 3049, // Optica
		"001610": 3050, // Carina
		"001612": 3051, // Otsuka
		"001613": 3052, // LibreStream
		"001615": 3053, // Nittan
		"001616": 3054, // Browan
		"001617": 3055, // MSI
		"001618": 3056, // HIVION
		"00161a": 3057, // Dametric
		"00161b": 2456, // Micronet
		"00161c": 3058, // e:cue
		"00161e": 3059, // Woojinnet
		"00161f": 3060, // SUNWAVETEC
		"001620": 101,  // Sony
		"001622": 3061, // BBH
		"001624": 3062, // Teneros
		"001625": 3063, // Impinj
		"001626": 125,  // ARRIS
		"001628": 3064, // Magicard
		"001629": 3065, // Nivus
		"00162c": 3066, // Xanboo
		"00162d": 3067, // STNet
		"00162f": 3068, // Geutebruck
		"001630": 3069, // Vativ
		"001631": 3070, // Xteam
		"001632": 152,  // Samsung
		"001634": 3071, // Mathtech
		"001635": 302,  // HP
		"001636": 3072, // Quanta
		"001637": 3073, // CITEL
		"001638": 576,  // TECOM
		"001639": 3074, // Ubiquam
		"00163a": 3075, // Yves
		"00163e": 3076, // Xensource
		"00163f": 3077, // CReTE
		"001640": 3078, // Asmobile
		"001642": 3079, // Pangolin
		"001643": 3080, // Sunhillo
		"001644": 452,  // LITE-ON
		"001646": 2,    // Cisco
		"001647": 2,    // Cisco
		"001648": 3081, // SSD
		"001649": 3082, // SetOne
		"00164a": 1381, // Vibration
		"00164e": 457,  // Nokia
		"001651": 3083, // Exeo
		"001652": 3084, // Hoatech
		"001654": 3085, // Flex-P
		"001655": 3086, // FUHO
		"001656": 1427, // Nintendo
		"001657": 3087, // Aegate
		"001658": 3088, // Fusiontech
		"00165c": 3089, // Trackflow
		"00165d": 3090, // AirDefense
		"001660": 76,   // Nortel
		"001662": 3091, // Liyuh
		"001663": 3092, // KBT
		"001664": 3093, // Prod-El
		"001666": 3094, // Quantier
		"001668": 3095, // Eishin
		"001669": 2254, // MRV
		"00166a": 3096, // TPS
		"00166b": 152,  // Samsung
		"00166c": 152,  // Samsung
		"00166d": 3097, // Yulong
		"00166e": 3098, // Arbitron
		"00166f": 421,  // Intel
		"001670": 3099, // SKNET
		"001672": 3100, // Zenway
		"001673": 3101, // Bury
		"001674": 3102, // EuroCB
		"001675": 125,  // ARRIS
		"001676": 421,  // Intel
		"001679": 3103, // eOn
		"00167c": 3104, // iRex
		"00167e": 3105, // Diboss
		"001683": 3106, // WEBIO
		"001684": 3107, // Donjin
		"001688": 3108, // ServerEngines
		"001689": 3109, // Pilkor
		"00168a": 3110, // id-Confirm
		"00168b": 3111, // Paralan
		"00168d": 3112, // KORWIN
		"00168e": 3113, // Vimicro
		"001690": 3114, // J-TEK
		"001691": 3115, // Moser-Baer
		"001692": 3116, // Scientific-Atlanta
		"001693": 3117, // PowerLink
		"001694": 3118, // Sennheiser
		"001695": 3119, // AVC
		"001696": 3120, // QDI
		"001697": 48,   // NEC
		"00169a": 3121, // Quadrics
		"00169c": 2,    // Cisco
		"00169d": 2,    // Cisco
		"00169f": 3122, // Vimtron
		"0016a0": 3123, // Auto-Maskin
		"0016a1": 3124, // 3Leaf
		"0016a2": 3125, // CentraLite
		"0016a3": 1968, // Ingeteam
		"0016a4": 3126, // Ezurio
		"0016a8": 3127, // CWT
		"0016a9": 3128, // 2EI
		"0016aa": 3129, // Kei
		"0016ab": 3130, // Dansensor
		"0016ac": 3131, // Toho
		"0016ad": 3132, // BT-Links
		"0016ae": 1075, // Inventel
		"0016b0": 3133, // VK
		"0016b1": 3134, // KBS
		"0016b2": 3135, // DriveCam
		"0016b3": 3136, // Photonicbridges
		"0016b4": 67,   // Private
		"0016b5": 125,  // ARRIS
		"0016b6": 1783, // Linksys
		"0016b8": 101,  // Sony
		"0016ba": 3137, // Weathernews
		"0016bb": 3138, // Law-Chain
		"0016bc": 457,  // Nokia
		"0016be": 3139, // InfRANET
		"0016c0": 3140, // Semtech
		"0016c1": 3141, // Eleksen
		"0016c2": 1798, // Avtec
		"0016c3": 3142, // BA
		"0016c4": 3143, // SiRF
		"0016c7": 2,    // Cisco
		"0016c8": 2,    // Cisco
		"0016ca": 76,   // Nortel
		"0016cb": 545,  // Apple
		"0016cc": 3144, // Xcute
		"0016d2": 3145, // Caspian
		"0016d3": 1592, // Wistron
		"0016d4": 358,  // Compal
		"0016d5": 3146, // Synccom
		"0016d6": 3147, // TDA
		"0016d7": 3148, // Sunways
		"0016d8": 3149, // Senea
		"0016da": 3150, // Futronic
		"0016db": 152,  // Samsung
		"0016dc": 3151, // Archos
		"0016dd": 3152, // Gigabeam
		"0016de": 790,  // FAST
		"0016df": 3153, // Lundinova
		"0016e0": 160,  // 3Com
		"0016e1": 3154, // SiliconStor
		"0016e3": 2545, // Askey
		"0016e6": 1929, // Giga-Byte
		"0016ea": 421,  // Intel
		"0016eb": 421,  // Intel
		"0016ec": 1115, // Elitegroup
		"0016ed": 1424, // Utility
		"0016ee": 3155, // Royaldigital
		"0016f0": 102,  // Dell
		"0016f1": 3156, // OmniSense
		"0016f4": 3157, // Eidicom
		"0016f7": 3158, // L-3
		"0016f8": 3159, // Aviqtech
		"0016fc": 3160, // Tohken
		"0016fd": 3161, // Jaty
		"0016fe": 430,  // Alpsalpine
		"001700": 125,  // ARRIS
		"001701": 3162, // KDE
		"001704": 3163, // Shinco
		"001705": 3164, // Methode
		"001706": 3165, // Techfaithwireless
		"001707": 3166, // InGrid
		"001708": 302,  // HP
		"001709": 3167, // Exalt
		"00170b": 3168, // Contela
		"00170c": 3169, // Twig
		"00170d": 3170, // Dust
		"00170e": 2,    // Cisco
		"00170f": 2,    // Cisco
		"001710": 3171, // Casa
		"001712": 3172, // ISCO
		"001715": 3173, // Qstik
		"001716": 3174, // Qno
		"001718": 3175, // Vansco
		"00171a": 3176, // Winegard
		"00171d": 3177, // Digit
		"00171f": 3178, // IMV
		"001722": 3179, // Hanazeder
		"001726": 3180, // m2c
		"001728": 3181, // Selex
		"001729": 3182, // Ubicod
		"00172a": 3183, // Proware
		"00172b": 3184, // Global
		"00172e": 3185, // FXC
		"00172f": 3186, // NeuLion
		"001730": 3187, // Automation
		"001731": 1806, // ASUS
		"001733": 3188, // SFR
		"001736": 3189, // iiTron
		"00173a": 3190, // Cloudastructure
		"00173b": 2,    // Cisco
		"00173c": 185,  // Extreme
		"00173d": 3191, // Neology
		"00173e": 3192, // LeucotronEquipamentos
		"00173f": 2469, // Belkin
		"001741": 3193, // Defidev
		"001742": 4,    // Fujitsu
		"001744": 3194, // Araneo
		"001745": 3195, // INNOTZ
		"001746": 3196, // Freedom9
		"001747": 1375, // Trimble
		"00174a": 3197, // Socomec
		"00174b": 457,  // Nokia
		"00174c": 3198, // Millipore
		"00174e": 3199, // Parama-tech
		"00174f": 3200, // iCatch
		"001751": 3201, // Online
		"001752": 3202, // DAGS
		"001753": 3203, // nFore
		"001756": 1930, // Vinci
		"001757": 3204, // RIX
		"001758": 3205, // ThruVision
		"001759": 2,    // Cisco
		"00175a": 2,    // Cisco
		"00175c": 3206, // Sharp
		"00175e": 3207, // Zed-3
		"00175f": 3208, // XENOLINK
		"001761": 67,   // Private
		"001762": 3209, // Solar
		"001764": 3210, // ATMedia
		"001765": 76,   // Nortel
		"001766": 3211, // Accense
		"001767": 3212, // Earforce
		"001768": 3213, // Zinwave
		"001769": 3214, // Cymphonix
		"00176a": 3215, // Avago
		"00176b": 3216, // Kiyon
		"00176c": 3217, // Pivot3
		"00176d": 475,  // Core
		"00176f": 3218, // PAX
		"001770": 3219, // Arti
		"001771": 3220, // APD
		"001773": 3221, // Laketune
		"001774": 3222, // Elesta
		"001777": 3223, // Obsidian
		"001779": 3224, // QuickTel
		"00177b": 3225, // Azalea
		"00177d": 3226, // IDT
		"00177e": 3227, // Meshcom
		"001782": 3228, // LoBenn
		"001784": 125,  // ARRIS
		"001785": 3229, // Sparr
		"001786": 3230, // wisembed
		"001787": 3231, // Brother
		"001789": 3232, // Zenitron
		"00178a": 3233, // Darts
		"00178b": 3234, // Teledyne
		"00178d": 3235, // Checkpoint
		"001791": 3236, // LinTech
		"001793": 3237, // Tigi
		"001794": 2,    // Cisco
		"001795": 2,    // Cisco
		"001796": 3238, // Rittmeyer
		"001798": 3239, // Azonic
		"001799": 3240, // SmarTire
		"00179a": 803,  // D-Link
		"00179d": 3241, // Kelman
		"00179e": 3242, // Sirit
		"00179f": 3243, // Apricorn
		"0017a1": 3244, // 3soft
		"0017a2": 3245, // Camrivox
		"0017a4": 302,  // HP
		"0017a5": 1785, // Ralink
		"0017a6": 3246, // Yosin
		"0017a8": 3247, // EDM
		"0017a9": 3248, // Sentivision
		"0017aa": 3249, // elab-experience
		"0017ab": 1427, // Nintendo
		"0017ad": 3250, // AceNet
		"0017ae": 3251, // GAI-Tronics
		"0017af": 3252, // Enermet
		"0017b0": 457,  // Nokia
		"0017b5": 3253, // Peerless
		"0017b6": 3254, // Aquantia
		"0017b7": 3255, // Tonze
		"0017b8": 3256, // Novatron
		"0017ba": 3257, // Sedo
		"0017bb": 3258, // Syrinx
		"0017bd": 3259, // Tibetsystem
		"0017bf": 3260, // Coherent
		"0017c0": 3261, // PureTech
		"0017c3": 3262, // KTF
		"0017c5": 3263, // SonicWALL
		"0017c7": 3264, // MARA
		"0017c8": 2849, // Kyocera
		"0017c9": 152,  // Samsung
		"0017ca": 551,  // Qisda
		"0017cb": 826,  // Juniper
		"0017ce": 3265, // Screen
		"0017cf": 3266, // iMCA-GmbH
		"0017d0": 3267, // Opticom
		"0017d1": 76,   // Nortel
		"0017d2": 3268, // Thinlinx
		"0017d3": 3269, // Etymotic
		"0017d5": 152,  // Samsung
		"0017d9": 3270, // AAI
		"0017db": 3271, // Canko
		"0017df": 2,    // Cisco
		"0017e0": 2,    // Cisco
		"0017e1": 3272, // DACOS
		"0017e2": 125,  // ARRIS
		"0017ed": 3273, // WooJooIT
		"0017ee": 125,  // ARRIS
		"0017ef": 372,  // IBM
		"0017f1": 3274, // Renu
		"0017f2": 545,  // Apple
		"0017f3": 124,  // Harris
		"0017f7": 1234, // CEM
		"0017f8": 3275, // Motech
		"0017fa": 612,  // Microsoft
		"0017fb": 3276, // FA
		"0017fc": 3277, // Suprema
		"0017ff": 3278, // PLAYLINE
		"001800": 3279, // Unigrand
		"001801": 2247, // Actiontec
		"001802": 2237, // Alpha
		"001806": 3280, // Hokkei
		"001807": 3281, // Fanstel
		"001808": 3282, // SightLogix
		"001809": 3283, // Cresyn
		"00180e": 3284, // Avega
		"00180f": 457,  // Nokia
		"001811": 3285, // Neuros
		"001813": 101,  // Sony
		"001814": 3286, // Mitutoyo
		"001815": 3287, // GZ
		"001816": 3288, // Ubixon
		"001818": 2,    // Cisco
		"001819": 2,    // Cisco
		"00181c": 3289, // Exterity
		"00181d": 3290, // Asia
		"00181e": 3291, // GDX
		"00181f": 1355, // Palmmicro
		"001820": 3292, // w5networks
		"001821": 3293, // Sindoricoh
		"001823": 957,  // Delta
		"001824": 3294, // Kimaldi
		"001825": 67,   // Private
		"001828": 3295, // e2v
		"001829": 3296, // Gatsometer
		"00182b": 3297, // Softier
		"00182c": 3298, // Ascend
		"00182e": 3299, // XStreamHD
		"001836": 3300, // REJ
		"001838": 3301, // PanAccess
		"001839": 1783, // Linksys
		"00183a": 2274, // Westell
		"00183b": 3302, // CENITS
		"00183c": 3303, // Encore
		"00183e": 3304, // Digilent
		"00183f": 1939, // 2Wire
		"001841": 375,  // High
		"001842": 457,  // Nokia
		"001843": 3305, // Dawevision
		"001845": 3306, // Pulsar-Telecom
		"001847": 3250, // AceNet
		"001848": 3307, // Vecima
		"001849": 2584, // nVent
		"00184a": 3308, // Catcher
		"00184c": 3309, // Bogen
		"00184d": 1368, // Netgear
		"00184e": 3310, // Lianhe
		"001851": 3311, // SWsoft
		"001853": 3312, // Atera
		"001854": 3313, // Argard
		"001856": 3314, // EyeFi
		"001858": 3315, // TagMaster
		"00185a": 3316, // uControl
		"00185c": 3317, // EDSLAB
		"00185d": 3318, // Taiguen
		"00185e": 3319, // Nexterm
		"00185f": 2380, // TAC
		"001861": 3320, // Ooma
		"001862": 732,  // Seagate
		"001863": 3321, // Veritech
		"001864": 1855, // Eaton
		"001865": 300,  // Siemens
		"001869": 3322, // Kingjim
		"00186c": 3323, // Neonode
		"00186e": 160,  // 3Com
		"001871": 302,  // HP
		"001872": 3324, // Expertise
		"001873": 2,    // Cisco
		"001874": 2,    // Cisco
		"001876": 3325, // WowWee
		"001877": 3326, // Amplex
		"001878": 3327, // Mackware
		"001879": 3328, // dSys
		"00187a": 3329, // Wiremold
		"00187b": 3330, // 4NSYS
		"00187c": 3331, // Intercross
		"00187d": 3332, // Armorlink
		"00187f": 3333, // Zodianet
		"001881": 3334, // Buyang
		"001882": 3335, // Huawei
		"001883": 3336, // FORMOSA21
		"001885": 687,  // Motorola
		"001886": 3337, // EL-Tech
		"001887": 3338, // Metasystem
		"001889": 3339, // WinNet
		"00188a": 3340, // Infinova
		"00188b": 102,  // Dell
		"00188d": 457,  // Nokia
		"00188e": 3341, // Ekahau
		"00188f": 3342, // Montgomery
		"001890": 3343, // RadioCOM
		"001892": 3344, // ads-tec
		"001894": 3345, // NPCore
		"001895": 3346, // Hansun
		"001897": 3347, // JESS-LINK
		"001898": 3348, // Kingstate
		"00189b": 1142, // Thomson
		"00189c": 3349, // Weldex
		"00189d": 3350, // Navcast
		"00189e": 3351, // OMNIKEY
		"00189f": 3352, // Lenntek
		"0018a1": 3353, // Tiqit
		"0018a2": 3354, // XIP
		"0018a3": 3355, // Zippy
		"0018a4": 125,  // ARRIS
		"0018a5": 3356, // ADigit
		"0018a6": 3357, // Persistent
		"0018a8": 3358, // AnNeal
		"0018ae": 3359, // TVT
		"0018af": 152,  // Samsung
		"0018b0": 76,   // Nortel
		"0018b1": 372,  // IBM
		"0018b6": 3360, // S3C
		"0018b9": 2,    // Cisco
		"0018ba": 2,    // Cisco
		"0018be": 3361, // ANSA
		"0018c0": 125,  // ARRIS
		"0018c2": 3362, // Firetide
		"0018c3": 3363, // CS
		"0018c4": 3364, // Raba
		"0018c5": 457,  // Nokia
		"0018c8": 3365, // ISONAS
		"0018c9": 3366, // EOps
		"0018ca": 3367, // Viprinet
		"0018cb": 3368, // Tecobest
		"0018cd": 3369, // Erae
		"0018ce": 3370, // Dreamtech
		"0018d0": 3371, // AtRoad
		"0018d1": 300,  // Siemens
		"0018d3": 3372, // Teamcast
		"0018d5": 3373, // Reigncom
		"0018d6": 3374, // Swirlnet
		"0018db": 3375, // EPL
		"0018dc": 3376, // Prostar
		"0018dd": 3377, // Silicondust
		"0018de": 421,  // Intel
		"0018df": 3378, // Morey
		"0018e0": 3379, // Anaveo
		"0018e1": 3380, // Verkerk
		"0018e3": 3381, // Visualgate
		"0018e4": 3382, // Yiguang
		"0018e5": 3383, // Adhoco
		"0018e7": 3384, // Cameo
		"0018e8": 3385, // Hacetron
		"0018e9": 3386, // Numata
		"0018ea": 3387, // Alltec
		"0018ec": 3388, // Welding
		"0018ef": 3389, // Escape
		"0018f0": 3390, // JOYTOTO
		"0018f3": 1806, // ASUS
		"0018f7": 3391, // Kameleon
		"0018f8": 1783, // Linksys
		"0018f9": 3392, // VVOND
		"0018fb": 3393, // Compro
		"0018fc": 3394, // Altec
		"0018fd": 3395, // Optimal
		"0018fe": 302,  // HP
		"0018ff": 3396, // PowerQuattro
		"001901": 3397, // F1MEDIA
		"001903": 3398, // Bigfoot
		"001904": 3399, // WB
		"001906": 2,    // Cisco
		"001907": 2,    // Cisco
		"001908": 3400, // Duaxes
		"00190a": 3401, // Hasware
		"00190c": 3303, // Encore
		"00190e": 3402, // Atech
		"00190f": 3403, // Advansus
		"001912": 3404, // Welcat
		"001914": 3405, // Winix
		"001915": 576,  // TECOM
		"001916": 3406, // PayTec
		"001917": 3407, // Posiflex
		"001919": 3408, // ASTEL
		"00191a": 3409, // Irlink
		"00191b": 3410, // Sputnik
		"00191c": 3411, // Sensicast
		"00191d": 1427, // Nintendo
		"00191e": 3412, // Beyondwiz
		"00191f": 254,  // Microlink
		"001921": 1115, // Elitegroup
		"001924": 3413, // LBNL
		"001925": 3414, // Intelicis
		"001926": 3415, // BitsGen
		"001927": 3416, // ImCoSys
		"001928": 300,  // Siemens
		"00192c": 125,  // ARRIS
		"00192d": 457,  // Nokia
		"00192f": 2,    // Cisco
		"001930": 2,    // Cisco
		"001931": 3417, // Balluff
		"001932": 3418, // Gude
		"001933": 944,  // Strix
		"001937": 3419, // CommerceGuard
		"001938": 3420, // UMB
		"001939": 3421, // Gigamips
		"00193a": 3422, // Oesolutions
		"00193b": 3423, // LigoWave
		"00193c": 3424, // HighPoint
		"00193f": 1402, // RDI
		"001940": 3425, // Rackable
		"001942": 3426, // ON
		"001943": 3427, // Belden
		"001948": 3428, // AireSpider
		"00194a": 3429, // Testo
		"00194b": 2049, // Sagemcom
		"00194e": 3430, // Ultra
		"00194f": 457,  // Nokia
		"001951": 3431, // NETCONS
		"001952": 3432, // ACOGITO
		"001953": 3433, // Chainleader
		"001954": 3434, // Leaf
		"001955": 2,    // Cisco
		"001956": 2,    // Cisco
		"001959": 3435, // Staccato
		"00195b": 803,  // D-Link
		"00195c": 3436, // Innotech
		"00195e": 125,  // ARRIS
		"00195f": 3437, // Valemount
		"001960": 3438, // DoCoMo
		"001962": 3439, // Commerciant
		"001963": 101,  // Sony
		"001964": 3440, // Doorking
		"001966": 1666, // Asiarock
		"001969": 76,   // Nortel
		"00196a": 3441, // MikroM
		"00196b": 3442, // Danpex
		"00196c": 3443, // Etrovision
		"00196e": 3444, // Metacom
		"00196f": 3445, // SensoPart
		"001970": 3446, // Z-Com
		"001972": 1971, // Plexus
		"001973": 3447, // Zeugma
		"001974": 3448, // 16063
		"001976": 3449, // Xipher
		"001977": 185,  // Extreme
		"001978": 3450, // Datum
		"001979": 457,  // Nokia
		"00197a": 3451, // MAZeT
		"00197b": 3452, // Picotest
		"00197c": 3453, // Riedel
		"00197f": 542,  // Plantronics
		"001980": 3454, // Gridpoint
		"001981": 3455, // Vivox
		"001982": 3456, // SmarDTV
		"001984": 3457, // ESTIC
		"001987": 2154, // Panasonic
		"001988": 3458, // Wi2Wi
		"001989": 3459, // Sonitrol
		"00198c": 3460, // iXSea
		"00198e": 3461, // Demant
		"001991": 3462, // avinfo
		"001992": 202,  // Adtran
		"001994": 3463, // Jorjin
		"001996": 3464, // TurboChef
		"001998": 3465, // Sato
		"001999": 4,    // Fujitsu
		"00199a": 3466, // EDO-EVI
		"00199c": 3467, // Ctring
		"00199d": 3468, // Vizio
		"00199e": 3469, // Nifty
		"00199f": 3470, // DKT
		"0019a1": 869,  // LG
		"0019a2": 3471, // Ordyn
		"0019a3": 3472, // asteel
		"0019a4": 3473, // Austar
		"0019a5": 3474, // RadarFind
		"0019a6": 125,  // ARRIS
		"0019a7": 3475, // ITU-T
		"0019a8": 3476, // WiQuest
		"0019a9": 2,    // Cisco
		"0019aa": 2,    // Cisco
		"0019ab": 3477, // Raycom
		"0019ac": 3478, // GSP
		"0019ad": 3479, // Bobst
		"0019af": 3480, // Rigol
		"0019b1": 3481, // Arrow7
		"0019b2": 3482, // XYnetsoft
		"0019b3": 3483, // Stanford
		"0019b4": 3484, // Intellio
		"0019b7": 457,  // Nokia
		"0019b9": 102,  // Dell
		"0019bb": 302,  // HP
		"0019be": 3485, // Altai
		"0019bf": 3486, // Citiway
		"0019c0": 125,  // ARRIS
		"0019c1": 430,  // Alpsalpine
		"0019c2": 3487, // Equustek
		"0019c3": 3488, // Qualitrol
		"0019c4": 3489, // Infocrypt
		"0019c5": 101,  // Sony
		"0019c6": 3032, // ZTE
		"0019c7": 1640, // Cambridge
		"0019c8": 3490, // AnyDATA
		"0019ca": 3491, // Broadata
		"0019cb": 2693, // ZyXEL
		"0019cc": 3492, // RCG
		"0019cf": 3493, // Salicru
		"0019d0": 3494, // Cathexis
		"0019d1": 421,  // Intel
		"0019d2": 421,  // Intel
		"0019d4": 3495, // ICX
		"0019d7": 3496, // Fortunetek
		"0019d8": 3497, // Maxfor
		"0019d9": 3498, // Zeutschel
		"0019db": 1812, // Micro-Star
		"0019dc": 3499, // ENENSYS
		"0019dd": 3500, // FEI-Zyfer
		"0019de": 3501, // Mobitek
		"0019df": 1142, // Thomson
		"0019e0": 1595, // TP-Link
		"0019e1": 76,   // Nortel
		"0019e2": 826,  // Juniper
		"0019e3": 545,  // Apple
		"0019e4": 1939, // 2Wire
		"0019e7": 2,    // Cisco
		"0019e8": 2,    // Cisco
		"0019ea": 3502, // TeraMage
		"0019eb": 3503, // Pyronix
		"0019ec": 3504, // Sagamore
		"0019ed": 3505, // Axesstel
		"0019f0": 3506, // Unionman
		"0019f3": 3507, // Cetis
		"0019f4": 3508, // Convergens
		"0019f5": 3509, // Imagination
		"0019f6": 3510, // Acconet
		"0019f7": 3511, // Onset
		"0019f9": 3512, // TDK-Lambda
		"0019fb": 3513, // BSkyB
		"0019fd": 1427, // Nintendo
		"0019ff": 3514, // Finnzymes
		"001a00": 1,    // Matrix
		"001a03": 3515, // Angel
		"001a04": 3516, // Interay
		"001a05": 3517, // Optibase
		"001a06": 3518, // OpVista
		"001a08": 3519, // Simoco
		"001a0b": 3520, // Bona
		"001a0d": 3521, // HandHeld
		"001a11": 3522, // Google
		"001a12": 3523, // Essilor
		"001a16": 457,  // Nokia
		"001a17": 3524, // Teak
		"001a19": 2329, // Computer
		"001a1a": 3525, // Gentex/Electro-Acoustic
		"001a1b": 125,  // ARRIS
		"001a1c": 3526, // GT&T
		"001a1e": 1685, // Aruba
		"001a20": 3527, // CMOTECH
		"001a26": 3528, // Deltanode
		"001a27": 3529, // Ubistar
		"001a2a": 2640, // Arcadyan
		"001a2b": 3530, // Ayecom
		"001a2c": 3531, // SATEC
		"001a2f": 2,    // Cisco
		"001a30": 2,    // Cisco
		"001a33": 3532, // ASI
		"001a34": 3533, // Konka
		"001a35": 3534, // BARTEC
		"001a36": 3535, // Aipermon
		"001a37": 3536, // Lear
		"001a38": 3537, // Sanmina-SCI
		"001a39": 3538, // Merten&CoKG
		"001a3a": 3539, // Dongahelecomm
		"001a3c": 3540, // Technowave
		"001a3e": 3541, // Faster
		"001a3f": 3542, // Intelbras
		"001a40": 3543, // A-Four
		"001a41": 3544, // INOCOVA
		"001a42": 3545, // Techcity
		"001a44": 3546, // JWTrading
		"001a47": 3547, // Agami
		"001a48": 3548, // Takacom
		"001a4a": 3549, // Qumranet
		"001a4b": 302,  // HP
		"001a4c": 3550, // Crossbow
		"001a4d": 1929, // Giga-Byte
		"001a4f": 621,  // AVM
		"001a50": 3551, // PheeNet
		"001a53": 3552, // Zylaya
		"001a55": 3553, // ACA-Digital
		"001a56": 3554, // ViewTel
		"001a59": 3555, // Ircona
		"001a5b": 3556, // NetCare
		"001a5c": 3557, // Euchner+Co
		"001a5d": 3558, // Mobinnova
		"001a5e": 3559, // Thincom
		"001a5f": 3560, // KitWorks.fi
		"001a60": 3561, // Wave
		"001a61": 3562, // PacStar
		"001a63": 3563, // Elster
		"001a64": 372,  // IBM
		"001a65": 3564, // Seluxit
		"001a66": 125,  // ARRIS
		"001a68": 3565, // Weltec
		"001a6a": 3566, // Tranzas
		"001a6c": 2,    // Cisco
		"001a6d": 2,    // Cisco
		"001a6e": 3567, // Impro
		"001a70": 1783, // Linksys
		"001a71": 3568, // Diostech
		"001a73": 1450, // Gemtek
		"001a74": 3569, // Procare
		"001a75": 101,  // Sony
		"001a77": 125,  // ARRIS
		"001a78": 3570, // ubtos
		"001a79": 3571, // Telecomunication
		"001a7b": 3572, // Teleco
		"001a7d": 3573, // cyber-blue
		"001a80": 101,  // Sony
		"001a81": 3574, // Zelax
		"001a82": 3575, // PROBA
		"001a83": 3576, // Pegasus
		"001a87": 3577, // Canhold
		"001a88": 3578, // Venergy
		"001a89": 457,  // Nokia
		"001a8a": 152,  // Samsung
		"001a8c": 3579, // Sophos
		"001a8e": 3580, // 3Way
		"001a8f": 76,   // Nortel
		"001a91": 3581, // FusionDynamic
		"001a92": 1806, // ASUS
		"001a94": 3582, // Votronic
		"001a97": 3583, // fitivision
		"001a9c": 3584, // RightHand
		"001a9f": 3585, // A-Link
		"001aa0": 102,  // Dell
		"001aa1": 2,    // Cisco
		"001aa2": 2,    // Cisco
		"001aa3": 3586, // Delorme
		"001aaa": 3587, // Analogic
		"001aac": 3588, // Corelatus
		"001aad": 125,  // ARRIS
		"001aae": 3589, // Savant
		"001aaf": 3590, // Blusens
		"001ab0": 569,  // Signal
		"001ab2": 677,  // Cyber
		"001ab3": 3591, // Visionite
		"001ab4": 3592, // FFEI
		"001ab7": 3593, // Ethos
		"001ab8": 3594, // Anseri
		"001ab9": 3595, // PMC
		"001abb": 3596, // Fontal
		"001abc": 2997, // U4EA
		"001abd": 3597, // Impatica
		"001ac0": 3598, // Joybien
		"001ac1": 160,  // 3Com
		"001ac2": 3599, // YEC
		"001ac3": 3116, // Scientific-Atlanta
		"001ac4": 1939, // 2Wire
		"001ac5": 3600, // Keysight
		"001ac7": 3601, // Unipoint
		"001ac8": 3602, // ISL
		"001ac9": 3603, // Suzuken
		"001aca": 3604, // Tilera
		"001acb": 3605, // Autocom
		"001acd": 3606, // Tidel
		"001ace": 3607, // Yupiteru
		"001ad0": 3608, // Albis
		"001ad1": 2691, // Fargo
		"001ad3": 3609, // Vamp
		"001ad4": 3610, // iPOX
		"001ad8": 3611, // AlsterAero
		"001ada": 3612, // Biz-2-Me
		"001adb": 125,  // ARRIS
		"001adc": 457,  // Nokia
		"001add": 3613, // PePWave
		"001ade": 125,  // ARRIS
		"001adf": 3614, // Interactivetv
		"001ae2": 2,    // Cisco
		"001ae3": 2,    // Cisco
		"001ae4": 3615, // Medicis
		"001ae5": 3616, // Mvox
		"001ae7": 1357, // Aztek
		"001ae9": 1427, // Nintendo
		"001aec": 3617, // Keumbee
		"001aed": 3618, // IncOTEC
		"001aee": 3619, // Shenztech
		"001aef": 3620, // Loopcomm
		"001af3": 3621, // Samyoung
		"001af4": 3622, // Handreamnet
		"001af5": 3623, // Pentaone
		"001af6": 3624, // Woven
		"001af9": 3625, // AeroVIronment
		"001afb": 3626, // Joby
		"001afc": 3627, // ModusLink
		"001afd": 3628, // Evolis
		"001afe": 3629, // Sofacreal
		"001aff": 3630, // Wizyoung
		"001b00": 3631, // Neopost
		"001b02": 3632, // ED
		"001b03": 3633, // Action
		"001b05": 3634, // YMC
		"001b07": 3635, // Mendocino
		"001b0b": 3636, // Phidgets
		"001b0c": 2,    // Cisco
		"001b0d": 2,    // Cisco
		"001b0f": 3637, // Petratec
		"001b11": 803,  // D-Link
		"001b12": 3638, // Apprion
		"001b13": 3639, // Icron
		"001b15": 3640, // Voxtel
		"001b16": 3641, // Celtro
		"001b1b": 300,  // Siemens
		"001b1c": 3260, // Coherent
		"001b1d": 3642, // Phoenix
		"001b20": 3643, // TPine
		"001b21": 421,  // Intel
		"001b23": 3644, // SimpleComTools
		"001b24": 3072, // Quanta
		"001b25": 76,   // Nortel
		"001b28": 3645, // Polygon
		"001b29": 3646, // Avantis
		"001b2a": 2,    // Cisco
		"001b2b": 2,    // Cisco
		"001b2c": 3647, // ATRON
		"001b2d": 3648, // Med-Eng
		"001b2e": 3649, // Sinkyo
		"001b2f": 1368, // Netgear
		"001b30": 3650, // Solitech
		"001b32": 2026, // QLogic
		"001b33": 457,  // Nokia
		"001b36": 3651, // Tsubata
		"001b37": 3652, // Computec
		"001b38": 358,  // Compal
		"001b39": 3653, // Proxicast
		"001b3a": 3654, // SIMS
		"001b3b": 3655, // Yi-Qing
		"001b3d": 3656, // EuroTel
		"001b3e": 3657, // Curtis
		"001b44": 3658, // SanDisk
		"001b45": 21,   // ABB
		"001b46": 3659, // Blueone
		"001b47": 3660, // Futarque
		"001b4a": 3661, // W&W
		"001b4b": 3662, // SANION
		"001b4c": 3663, // Signtech
		"001b4d": 3664, // Areca
		"001b4f": 620,  // Avaya
		"001b51": 3665, // Vector
		"001b52": 125,  // ARRIS
		"001b53": 2,    // Cisco
		"001b54": 2,    // Cisco
		"001b56": 3666, // Tehuti
		"001b57": 3667, // Semindia
		"001b59": 101,  // Sony
		"001b5b": 1939, // 2Wire
		"001b5c": 3668, // Azuretec
		"001b5d": 3669, // Vololink
		"001b5e": 3670, // BPL
		"001b5f": 3671, // Alien
		"001b60": 3672, // Navigon
		"001b63": 545,  // Apple
		"001b64": 3673, // IsaacLandKorea
		"001b66": 3118, // Sennheiser
		"001b67": 2,    // Cisco
		"001b68": 3674, // Modnnet
		"001b69": 3675, // Equaline
		"001b6b": 3676, // Swyx
		"001b6d": 3677, // Midtronics
		"001b6e": 3600, // Keysight
		"001b6f": 3678, // Teletrak
		"001b71": 3679, // Telular
		"001b74": 3680, // MiraLink
		"001b75": 3681, // Hypermedia
		"001b76": 3682, // Ripcode
		"001b77": 421,  // Intel
		"001b78": 302,  // HP
		"001b7a": 1427, // Nintendo
		"001b7b": 3683, // Tintometer
		"001b7e": 3684, // Beckmann
		"001b80": 3685, // LORD
		"001b83": 3686, // Finsoft
		"001b84": 3687, // Scan
		"001b87": 3688, // Deepsound
		"001b8a": 3689, // 2M
		"001b8c": 3690, // JMicron
		"001b8d": 3691, // Electronic
		"001b8f": 2,    // Cisco
		"001b90": 2,    // Cisco
		"001b91": 3692, // Efkon
		"001b92": 3693, // l-acoustics
		"001b97": 3694, // Violin
		"001b98": 152,  // Samsung
		"001b9b": 3695, // Hose-McCann
		"001b9e": 2545, // Askey
		"001b9f": 3696, // Calyptech
		"001ba0": 3697, // Awox
		"001ba1": 3698, // Amic
		"001ba5": 3699, // MyungMin
		"001ba6": 3700, // intotech
		"001ba7": 3701, // Lorica
		"001ba8": 3702, // UBI&MOBI
		"001ba9": 3231, // Brother
		"001baa": 3703, // XenICs
		"001bab": 3704, // Telchemyorporated
		"001bad": 3705, // iControl
		"001baf": 457,  // Nokia
		"001bb0": 3706, // Bharat
		"001bb1": 1592, // Wistron
		"001bb2": 3707, // Intellect
		"001bb3": 3708, // Condalo
		"001bb4": 3709, // Airvod
		"001bb5": 3710, // Cherry
		"001bb6": 3711, // Bird
		"001bb8": 3712, // Blueway
		"001bb9": 1115, // Elitegroup
		"001bba": 76,   // Nortel
		"001bbb": 3713, // RFTech
		"001bbf": 2049, // Sagemcom
		"001bc0": 826,  // Juniper
		"001bc1": 3714, // HOLUX
		"001bc3": 3715, // Mobisolution
		"001bc4": 3716, // Ultratec
		"001bc7": 3717, // StarVedia
		"001bc8": 3718, // Miura
		"001bcb": 3719, // Pempek
		"001bcd": 3720, // Daviscomms
		"001bcf": 3721, // Dataupia
		"001bd0": 3722, // Identec
		"001bd1": 3723, // Sogestmatic
		"001bd3": 2154, // Panasonic
		"001bd4": 2,    // Cisco
		"001bd5": 2,    // Cisco
		"001bd8": 3724, // FLIR
		"001bda": 1135, // UTStarcom
		"001bdc": 3725, // Vencer
		"001bdd": 125,  // ARRIS
		"001bde": 3726, // Renkus-Heinz
		"001be0": 3727, // TELENOT
		"001be1": 3728, // ViaLogy
		"001be2": 3729, // AhnLab
		"001be5": 3730, // 802automation
		"001be6": 3731, // VR
		"001be7": 3732, // Postek
		"001be8": 3733, // Ultratronik
		"001be9": 858,  // Broadcom
		"001bea": 1427, // Nintendo
		"001beb": 3734, // DMP
		"001bec": 3735, // Netio
		"001bed": 90,   // Brocade
		"001bee": 457,  // Nokia
		"001bf2": 3736, // Kworld
		"001bf4": 3737, // Kenwin
		"001bf6": 3738, // CoNWISE
		"001bf8": 3739, // Digitrax
		"001bfb": 430,  // Alpsalpine
		"001bfc": 1806, // ASUS
		"001bfd": 3740, // Dignsys
		"001bfe": 3741, // Zavio
		"001c04": 3742, // Airgain
		"001c06": 300,  // Siemens
		"001c07": 3743, // Cwlinux
		"001c08": 3744, // Echo360
		"001c09": 3745, // SAE
		"001c0c": 3746, // TANITA
		"001c0d": 3747, // G-Technology
		"001c0e": 2,    // Cisco
		"001c0f": 2,    // Cisco
		"001c10": 1783, // Linksys
		"001c11": 125,  // ARRIS
		"001c12": 125,  // ARRIS
		"001c13": 3748, // Optsys
		"001c14": 809,  // VMware
		"001c15": 3749, // iPhotonix
		"001c17": 76,   // Nortel
		"001c1b": 3750, // Hyperstone
		"001c1c": 3751, // Center
		"001c1e": 3752, // emtrion
		"001c21": 3753, // Nucsafe
		"001c23": 102,  // Dell
		"001c27": 3754, // Sunell
		"001c28": 3755, // Sphairon
		"001c2a": 3756, // Envisacor
		"001c2b": 3757, // Alertme.com
		"001c2c": 3758, // Synapse
		"001c2d": 3759, // FlexRadio
		"001c2f": 3760, // Pfister
		"001c32": 3761, // Telian
		"001c33": 1114, // Sutron
		"001c35": 457,  // Nokia
		"001c36": 3762, // iNEWiT
		"001c37": 3763, // Callpod
		"001c38": 3764, // Bio-Rad
		"001c3a": 3765, // Element
		"001c3b": 3766, // AmRoad
		"001c3d": 3767, // WaveStorm
		"001c3e": 3768, // ECKey
		"001c40": 3769, // VDG-Security
		"001c42": 3770, // Parallels
		"001c43": 152,  // Samsung
		"001c46": 3771, // Qtum
		"001c48": 3772, // WiDeFi
		"001c49": 3773, // Zoltan
		"001c4a": 621,  // AVM
		"001c4b": 3774, // Gener8
		"001c4e": 3775, // TASA
		"001c4f": 3776, // Macab
		"001c51": 3777, // Celeno
		"001c54": 3778, // Hillstone
		"001c56": 3779, // Pado
		"001c57": 2,    // Cisco
		"001c58": 2,    // Cisco
		"001c5b": 3780, // Chubb
		"001c5f": 3781, // Winland
		"001c62": 869,  // LG
		"001c63": 3782, // Truen
		"001c64": 2228, // Landis+Gyr
		"001c65": 3783, // JoeScan
		"001c66": 3784, // Ucamp
		"001c67": 3785, // Pumpkin
		"001c6a": 3786, // Weiss
		"001c6b": 3787, // CoVAX
		"001c6c": 3788, // 30805
		"001c6d": 3789, // Kyohritsu
		"001c6e": 3790, // Newbury
		"001c6f": 3791, // Emfit
		"001c70": 3792, // Novacomm
		"001c71": 3793, // Emergent
		"001c73": 3794, // Arista
		"001c74": 3795, // Syswan
		"001c75": 3796, // Segnet
		"001c77": 3797, // Prodys
		"001c7b": 3798, // Castlenet
		"001c7c": 3799, // Perq
		"001c7d": 3800, // Excelpoint
		"001c7e": 35,   // Toshiba
		"001c82": 3801, // Genew
		"001c85": 3802, // Eunicorn
		"001c86": 3803, // Cranite
		"001c87": 3804, // Uriver
		"001c88": 3805, // Transystem
		"001c89": 2272, // Force
		"001c8a": 3806, // Cirrascale
		"001c8c": 3807, // Dial
		"001c8f": 587,  // Advanced
		"001c90": 3808, // Empacket
		"001c91": 3809, // Gefen
		"001c92": 3810, // Tervela
		"001c93": 3811, // ExaDigm
		"001c95": 3812, // Opticomm
		"001c96": 3813, // Linkwise
		"001c97": 3814, // Enzytek
		"001c98": 3815, // Lucky
		"001c99": 3816, // Shunra
		"001c9a": 457,  // Nokia
		"001c9b": 3817, // FEIG
		"001c9c": 76,   // Nortel
		"001c9d": 3818, // Liecthi
		"001c9f": 3819, // Razorstream
		"001ca1": 3820, // Akamai
		"001ca3": 3821, // Terra
		"001ca4": 101,  // Sony
		"001ca5": 3822, // Zygo
		"001ca6": 3823, // Win4NET
		"001caa": 3824, // Bellon
		"001cac": 3825, // Qniq
		"001cae": 3826, // WiChorus
		"001caf": 3827, // Plato
		"001cb0": 2,    // Cisco
		"001cb1": 2,    // Cisco
		"001cb2": 3828, // BPT
		"001cb3": 545,  // Apple
		"001cb8": 3829, // CBC
		"001cba": 3830, // VerScient
		"001cbb": 3831, // MusicianLink
		"001cbc": 3832, // CastGrabber
		"001cbe": 1427, // Nintendo
		"001cbf": 421,  // Intel
		"001cc0": 421,  // Intel
		"001cc1": 125,  // ARRIS
		"001cc3": 125,  // ARRIS
		"001cc4": 302,  // HP
		"001cc5": 160,  // 3Com
		"001cc6": 3833, // ProStor
		"001cc9": 3834, // Kaise
		"001ccc": 2221, // BlackBerry
		"001ccd": 3835, // Alektrona
		"001ccf": 3836, // Limetek
		"001cd0": 3837, // Circleone
		"001cd3": 3838, // ZP
		"001cd4": 457,  // Nokia
		"001cd5": 3839, // ZeeVee
		"001cd6": 457,  // Nokia
		"001cd9": 1389, // GlobalTop
		"001cda": 3840, // Exegin
		"001cdb": 3841, // Carpoint
		"001cdc": 2123, // Custom
		"001cdd": 3842, // Cowbell
		"001cdf": 2469, // Belkin
		"001ce2": 3843, // Attero
		"001ce3": 3844, // Optimedical
		"001ce5": 3845, // MBS
		"001ce6": 3846, // Innes
		"001ce8": 3847, // Cummins
		"001ce9": 3848, // Galaxy
		"001cea": 3116, // Scientific-Atlanta
		"001ceb": 76,   // Nortel
		"001cec": 3849, // Mobilesoft
		"001ced": 3850, // Environnement
		"001cee": 3206, // Sharp
		"001cef": 389,  // Primax
		"001cf0": 803,  // D-Link
		"001cf1": 3851, // SUPoX
		"001cf2": 3852, // Tenlon
		"001cf4": 3853, // Media
		"001cf5": 3854, // Wiseblue
		"001cf6": 2,    // Cisco
		"001cf7": 3855, // AudioScience
		"001cf8": 3856, // Parade
		"001cf9": 2,    // Cisco
		"001cfa": 3857, // Alarm.com
		"001cfb": 125,  // ARRIS
		"001cfd": 3858, // Universal
		"001cfe": 3859, // Quartics
		"001cff": 3860, // Napera
		"001d00": 3861, // Brivo
		"001d03": 3862, // Design
		"001d06": 3863, // HM
		"001d09": 102,  // Dell
		"001d0c": 3864, // MobileCompia
		"001d0d": 101,  // Sony
		"001d0e": 3865, // Agapha
		"001d0f": 1595, // TP-Link
		"001d12": 3866, // Rohm
		"001d13": 3867, // NextGTV
		"001d16": 3188, // SFR
		"001d19": 2640, // Arcadyan
		"001d1b": 3868, // Sangean
		"001d1d": 3869, // Inter-M
		"001d20": 3870, // Comtrend
		"001d23": 3871, // Sensus
		"001d25": 152,  // Samsung
		"001d26": 3872, // Rockridgesound
		"001d27": 3873, // NAC-Intercom
		"001d28": 101,  // Sony
		"001d29": 3874, // Doro
		"001d2c": 3875, // Wavetrend
		"001d2d": 3876, // Pylone
		"001d2e": 2738, // Ruckus
		"001d2f": 3877, // QuantumVision
		"001d32": 3878, // Longkay
		"001d33": 3879, // Maverick
		"001d34": 3880, // SYRIS
		"001d35": 3881, // Viconics
		"001d38": 732,  // Seagate
		"001d39": 3882, // Moohadigital
		"001d3b": 457,  // Nokia
		"001d3c": 3883, // Muscle
		"001d3d": 3884, // Avidyne
		"001d3f": 3885, // Mitron
		"001d42": 76,   // Nortel
		"001d44": 3886, // Krohne
		"001d45": 2,    // Cisco
		"001d46": 2,    // Cisco
		"001d47": 3887, // Covote
		"001d4e": 3888, // TCM
		"001d4f": 545,  // Apple
		"001d50": 3889, // Spinetix
		"001d53": 3890, // S&O
		"001d55": 3891, // ZANTAZ
		"001d56": 3892, // Kramer
		"001d58": 3893, // CQ
		"001d5a": 1939, // 2Wire
		"001d5c": 3894, // Tom
		"001d60": 1806, // ASUS
		"001d61": 3895, // BIJ
		"001d62": 3896, // InPhase
		"001d67": 3897, // Amec
		"001d6a": 2237, // Alpha
		"001d6b": 125,  // ARRIS
		"001d6c": 3898, // ClariPhy
		"001d6d": 3899, // Confidant
		"001d6e": 457,  // Nokia
		"001d6f": 3900, // Chainzone
		"001d70": 2,    // Cisco
		"001d71": 2,    // Cisco
		"001d72": 1592, // Wistron
		"001d73": 1077, // Buffalo
		"001d75": 3901, // Radioscape
		"001d76": 3902, // Eyeheight
		"001d77": 3903, // NSGate
		"001d79": 3904, // Signamax
		"001d7d": 1929, // Giga-Byte
		"001d7e": 1783, // Linksys
		"001d7f": 3905, // Tekron
		"001d83": 3906, // Emitech
		"001d84": 64,   // Gateway
		"001d86": 3907, // Shinwa
		"001d88": 3908, // Clearwire
		"001d89": 3909, // VaultStor
		"001d8a": 3910, // TechTrex
		"001d8e": 3911, // Alereon
		"001d8f": 3912, // PureWave
		"001d91": 3913, // Digitize
		"001d92": 1812, // Micro-Star
		"001d93": 3914, // Modacom
		"001d94": 3915, // Climax
		"001d95": 1029, // Flash
		"001d97": 3916, // Alertus
		"001d98": 457,  // Nokia
		"001d9a": 3917, // Godex
		"001d9d": 3918, // Artjoy
		"001d9e": 3919, // Axion
		"001da1": 2,    // Cisco
		"001da2": 2,    // Cisco
		"001da3": 3920, // SabiOso
		"001da5": 3399, // WB
		"001da8": 3921, // Takahata
		"001da9": 3922, // Castles
		"001daa": 3923, // DrayTek
		"001dac": 3924, // Gigamon
		"001dad": 3925, // Sinotech
		"001daf": 76,   // Nortel
		"001db1": 3926, // Crescendo
		"001db5": 826,  // Juniper
		"001db6": 3927, // BestComm
		"001db7": 3928, // Tendril
		"001db8": 3929, // Intoto
		"001dba": 101,  // Sony
		"001dbc": 1427, // Nintendo
		"001dbd": 3930, // Versamed
		"001dbe": 125,  // ARRIS
		"001dbf": 3931, // Radiient
		"001dc2": 3932, // Xortec
		"001dc4": 3933, // AIOI
		"001dc6": 3934, // SNR
		"001dc8": 3935, // Navionics
		"001dc9": 3936, // GainSpan
		"001dca": 3937, // PAV
		"001dcd": 125,  // ARRIS
		"001dce": 125,  // ARRIS
		"001dcf": 125,  // ARRIS
		"001dd0": 125,  // ARRIS
		"001dd1": 125,  // ARRIS
		"001dd2": 125,  // ARRIS
		"001dd3": 125,  // ARRIS
		"001dd4": 125,  // ARRIS
		"001dd5": 125,  // ARRIS
		"001dd6": 125,  // ARRIS
		"001dd7": 3938, // Algolith
		"001dd8": 612,  // Microsoft
		"001ddb": 3939, // C-BEL
		"001ddf": 3940, // Sunitec
		"001de0": 421,  // Intel
		"001de1": 421,  // Intel
		"001de2": 3941, // Radionor
		"001de3": 3942, // Intuicom
		"001de5": 2,    // Cisco
		"001de6": 2,    // Cisco
		"001de9": 457,  // Nokia
		"001deb": 3943, // DINEC
		"001dec": 3944, // Marusys
		"001def": 3945, // Trimm
		"001df0": 3946, // Vidient
		"001df1": 3947, // Intego
		"001df2": 3948, // Netflix
		"001df4": 3949, // Magellan
		"001df5": 3950, // Sunshine
		"001df6": 152,  // Samsung
		"001df9": 3951, // Cybiotronics
		"001dfb": 3952, // NETCLEUS
		"001dfc": 3953, // Ksic
		"001dfd": 457,  // Nokia
		"001dfe": 1158, // Palm
		"001e03": 3954, // LiComm
		"001e04": 3955, // Hanson
		"001e06": 3956, // Wibrain
		"001e07": 3957, // Winy
		"001e08": 3958, // Centec
		"001e09": 3959, // ZEFATEK
		"001e0a": 3960, // Syba
		"001e0b": 302,  // HP
		"001e0d": 3961, // Micran
		"001e0f": 3962, // Briot
		"001e10": 3335, // Huawei
		"001e11": 3963, // Elelux
		"001e12": 3964, // Ecolab
		"001e13": 2,    // Cisco
		"001e14": 2,    // Cisco
		"001e16": 3965, // Keytronix
		"001e17": 3966, // STN
		"001e19": 3967, // Gtri
		"001e1f": 76,   // Nortel
		"001e20": 3968, // Intertain
		"001e21": 551,  // Qisda
		"001e26": 3969, // Digifriends
		"001e27": 3970, // SBN
		"001e28": 3971, // Lumexis
		"001e29": 3972, // Hypertherm
		"001e2a": 1368, // Netgear
		"001e2c": 3973, // CyVerse
		"001e2d": 3974, // Stim
		"001e2f": 3975, // DiMoto
		"001e30": 3976, // Shireen
		"001e31": 3977, // Infomark
		"001e32": 3978, // Zensys
		"001e33": 3979, // Inventec
		"001e34": 3980, // CryptoMetrics
		"001e35": 1427, // Nintendo
		"001e36": 3981, // Ipte
		"001e38": 3982, // Bluecard
		"001e39": 3983, // Comsys
		"001e3a": 457,  // Nokia
		"001e3b": 457,  // Nokia
		"001e3d": 430,  // Alpsalpine
		"001e3e": 3984, // KMW
		"001e3f": 3985, // TrellisWare
		"001e42": 3986, // Teltonika
		"001e44": 3987, // Santec
		"001e45": 101,  // Sony
		"001e46": 125,  // ARRIS
		"001e48": 3988, // Wi-Links
		"001e49": 2,    // Cisco
		"001e4a": 2,    // Cisco
		"001e4f": 102,  // Dell
		"001e50": 3989, // Battistoni
		"001e52": 545,  // Apple
		"001e53": 3990, // Further
		"001e55": 3991, // CoWON
		"001e57": 3992, // ALCOMA
		"001e58": 803,  // D-Link
		"001e5a": 125,  // ARRIS
		"001e5b": 3993, // Unitron
		"001e5e": 3994, // Computime
		"001e5f": 3995, // KwikByte
		"001e61": 3996, // ITEC
		"001e62": 3997, // Siemon
		"001e63": 3998, // Vibro-Meter
		"001e64": 421,  // Intel
		"001e65": 421,  // Intel
		"001e67": 421,  // Intel
		"001e68": 3072, // Quanta
		"001e69": 1142, // Thomson
		"001e6c": 3999, // Opaque
		"001e6f": 4000, // Magna-Power
		"001e72": 4001, // PCS
		"001e73": 3032, // ZTE
		"001e74": 2049, // Sagemcom
		"001e75": 869,  // LG
		"001e77": 4002, // Air2App
		"001e78": 4003, // Owitek
		"001e79": 2,    // Cisco
		"001e7a": 2,    // Cisco
		"001e7c": 4004, // Taiwick
		"001e7d": 152,  // Samsung
		"001e7e": 76,   // Nortel
		"001e80": 2157, // Icotera
		"001e81": 4005, // CNB
		"001e82": 3658, // SanDisk
		"001e84": 4006, // Pika
		"001e85": 4007, // Lagotek
		"001e86": 4008, // MEL
		"001e87": 4009, // Realease
		"001e89": 4010, // CRFS
		"001e8a": 4011, // eCopy
		"001e8c": 1806, // ASUS
		"001e8d": 125,  // ARRIS
		"001e8e": 4012, // Hunkeler
		"001e8f": 87,   // Canon
		"001e90": 1115, // Elitegroup
		"001e91": 4013, // KIMIN
		"001e93": 4014, // CiriTech
		"001e94": 4015, // Supercom
		"001e95": 4016, // Sigmalink
		"001e96": 4017, // Sepura
		"001e98": 4018, // GreenLine
		"001e99": 4019, // Vantanol
		"001e9b": 4020, // San-Eisha
		"001e9c": 4021, // Fidustron
		"001e9d": 4022, // Recall
		"001e9f": 4023, // Visioneering
		"001ea0": 4024, // XLN-t
		"001ea1": 4025, // Brunata
		"001ea2": 4026, // Symx
		"001ea3": 457,  // Nokia
		"001ea4": 457,  // Nokia
		"001ea5": 4027, // ROBOTOUS
		"001ea7": 2247, // Actiontec
		"001ea9": 1427, // Nintendo
		"001eaa": 4028, // E-Senza
		"001eab": 4029, // TeleWell
		"001eac": 4030, // Armadeus
		"001ead": 4031, // Wingtech
		"001eb0": 4032, // ImesD
		"001eb1": 4033, // Cryptsoft
		"001eb2": 869,  // LG
		"001eb4": 4034, // Unifat
		"001eb7": 4035, // TBTech
		"001eb8": 4036, // Aloys
		"001ebb": 4037, // Bluelight
		"001ebd": 2,    // Cisco
		"001ebe": 2,    // Cisco
		"001ec0": 706,  // Microchip
		"001ec1": 160,  // 3Com
		"001ec2": 545,  // Apple
		"001ec3": 4038, // Kozio
		"001ec4": 4039, // Celio
		"001ec6": 4040, // Obvius
		"001ec7": 1939, // 2Wire
		"001ec9": 102,  // Dell
		"001eca": 76,   // Nortel
		"001ecc": 4041, // Cdvi
		"001ecd": 4042, // KYLAND
		"001ece": 4043, // BISA
		"001ecf": 796,  // Philips
		"001ed0": 4044, // Ingespace
		"001ed4": 4045, // Doble
		"001ed5": 4046, // Tekon-Automatics
		"001edc": 101,  // Sony
		"001ede": 4047, // BYD
		"001ee0": 4048, // Urmet
		"001ee1": 152,  // Samsung
		"001ee2": 152,  // Samsung
		"001ee3": 4049, // T&W
		"001ee5": 1783, // Linksys
		"001ee7": 4050, // Epic
		"001ee8": 4051, // Mytek
		"001ee9": 4052, // Stoneridge
		"001eeb": 4053, // Talk-A-Phone
		"001eec": 358,  // Compal
		"001eed": 4054, // Adventiq
		"001eee": 4055, // ETL
		"001eef": 4056, // Cantronic
		"001ef0": 4057, // Gigafin
		"001ef1": 4058, // Servimat
		"001ef3": 4059, // From2
		"001ef6": 2,    // Cisco
		"001ef7": 2,    // Cisco
		"001ef8": 4060, // Emfinity
		"001efa": 4061, // PROTEI
		"001eff": 4062, // Mueller-Elektronik
		"001f00": 457,  // Nokia
		"001f01": 457,  // Nokia
		"001f02": 4063, // Pixelmetrix
		"001f03": 4064, // NUM
		"001f04": 4065, // Granch
		"001f05": 4066, // iTAS
		"001f07": 4067, // AZTEQ
		"001f08": 4068, // Risco
		"001f09": 4069, // Jastec
		"001f0a": 76,   // Nortel
		"001f0f": 4070, // Select
		"001f11": 4071, // Openmoko
		"001f12": 826,  // Juniper
		"001f14": 4072, // NexG
		"001f15": 4073, // Bioscrypt
		"001f16": 1592, // Wistron
		"001f17": 4074, // IDX
		"001f18": 4075, // Hakusan.Mfg
		"001f19": 4076, // BEN-RI
		"001f1a": 4077, // Prominvest
		"001f1b": 4078, // RoyalTek
		"001f1e": 4079, // Astec
		"001f1f": 115,  // Edimax
		"001f20": 4080, // Logitech
		"001f23": 4081, // Interacoustics
		"001f24": 4082, // Digitview
		"001f25": 3845, // MBS
		"001f26": 2,    // Cisco
		"001f27": 2,    // Cisco
		"001f29": 302,  // HP
		"001f2a": 4083, // Accm
		"001f2c": 4084, // Starbridge
		"001f2f": 4085, // Berker
		"001f30": 4086, // Travelping
		"001f31": 4087, // Radiocomp
		"001f32": 1427, // Nintendo
		"001f33": 1368, // Netgear
		"001f35": 4088, // AIR802
		"001f38": 4089, // Positron
		"001f3b": 421,  // Intel
		"001f3c": 421,  // Intel
		"001f3d": 4090, // Qbit
		"001f3f": 621,  // AVM
		"001f40": 4091, // Speakercraft
		"001f41": 2738, // Ruckus
		"001f42": 4092, // Etherstack
		"001f45": 312,  // Enterasys
		"001f46": 76,   // Nortel
		"001f48": 4093, // Mojix
		"001f4c": 4094, // Roseman
		"001f4d": 4095, // Segnetics
		"001f4f": 4096, // Thinkware
		"001f50": 4097, // Swissdis
		"001f51": 4098, // HD
		"001f54": 4099, // Lorex
		"001f5b": 545,  // Apple
		"001f5c": 457,  // Nokia
		"001f5d": 457,  // Nokia
		"001f5e": 4100, // Dyna
		"001f5f": 4101, // Blatand
		"001f60": 4102, // Compass
		"001f61": 4103, // Talent
		"001f66": 4104, // Planar
		"001f67": 89,   // Hitachi
		"001f69": 4105, // Pingood
		"001f6a": 4106, // PacketFlux
		"001f6b": 869,  // LG
		"001f6c": 2,    // Cisco
		"001f6d": 2,    // Cisco
		"001f6e": 2503, // Vtech
		"001f70": 4107, // Botik
		"001f71": 4108, // xG
		"001f73": 4109, // Teraview
		"001f76": 1621, // AirLogic
		"001f79": 4110, // Lodam
		"001f7a": 4111, // WiWide
		"001f7b": 4112, // TechNexion
		"001f7c": 4113, // Witelcom
		"001f7e": 125,  // ARRIS
		"001f7f": 4114, // Phabrix
		"001f80": 4115, // Lucas
		"001f82": 1630, // Cal-Comp
		"001f83": 4116, // Teleplan
		"001f86": 4117, // digEcor
		"001f87": 4118, // Skydigital
		"001f89": 4119, // Signalion
		"001f8c": 4120, // CCS
		"001f90": 2247, // Actiontec
		"001f92": 687,  // Motorola
		"001f93": 4121, // Xiotech
		"001f94": 4122, // Lascar
		"001f95": 2049, // Sagemcom
		"001f96": 4123, // Aprotech
		"001f98": 4124, // Daiichi-Dentsu
		"001f99": 4125, // SERONICS
		"001f9a": 76,   // Nortel
		"001f9b": 4126, // Posbro
		"001f9c": 4127, // Ledco
		"001f9d": 2,    // Cisco
		"001f9e": 2,    // Cisco
		"001fa0": 4128, // A10
		"001fa1": 4129, // Gtran
		"001fa3": 4049, // T&W
		"001fa5": 4130, // Blue-White
		"001fa7": 101,  // Sony
		"001faa": 4131, // Taseon
		"001fac": 4132, // Goodmill
		"001faf": 4133, // NextIO
		"001fb0": 4134, // TimeIPS
		"001fb1": 4135, // Cybertech
		"001fb3": 1939, // 2Wire
		"001fb4": 4136, // SmartShare
		"001fb7": 4137, // WiMate
		"001fb9": 4138, // Paltronics
		"001fba": 4139, // Boyoung
		"001fbb": 4140, // Xenatech
		"001fbc": 4141, // EVGA
		"001fbd": 2849, // Kyocera
		"001fc1": 4142, // Hanlong
		"001fc3": 4143, // SmartSynch
		"001fc4": 125,  // ARRIS
		"001fc5": 1427, // Nintendo
		"001fc6": 1806, // ASUS
		"001fc8": 4144, // Up-Today
		"001fc9": 2,    // Cisco
		"001fca": 2,    // Cisco
		"001fcb": 4145, // NIW
		"001fcc": 152,  // Samsung
		"001fcd": 152,  // Samsung
		"001fce": 4146, // Qtech
		"001fcf": 3055, // MSI
		"001fd0": 1929, // Giga-Byte
		"001fd1": 4147, // Optex
		"001fd3": 4148, // RIVA
		"001fd4": 4149, // 4IPNET
		"001fd7": 4150, // Telerad
		"001fd8": 4151, // A-Trust
		"001fd9": 4152, // RSD
		"001fda": 76,   // Nortel
		"001fdd": 4153, // GDI
		"001fde": 457,  // Nokia
		"001fdf": 457,  // Nokia
		"001fe0": 4154, // EdgeVelocity
		"001fe3": 869,  // LG
		"001fe4": 101,  // Sony
		"001fe5": 4155, // In-Circuit
		"001fe6": 4156, // Alphion
		"001fe7": 4157, // Simet
		"001fe8": 4158, // KURUSUGAWA
		"001fe9": 4159, // Printrex
		"001fec": 3758, // Synapse
		"001fed": 4160, // Tecan
		"001fee": 4161, // ubisys
		"001fef": 4162, // Shinsei
		"001ff2": 4163, // VIA
		"001ff3": 545,  // Apple
		"001ff8": 300,  // Siemens
		"001ffa": 4164, // Coretree
		"001ffc": 4165, // Riccius+Sohn
		"001ffd": 4166, // Indigo
		"001fff": 4167, // Respironics
		"002000": 613,  // Lexmark
		"002001": 4168, // DSP
		"002002": 4169, // Seritech
		"002004": 4170, // Yamatake-Honeywell
		"002005": 4171, // Simple
		"002006": 4172, // Garrett
		"002007": 4173, // SFA
		"002008": 4174, // Cable
		"00200a": 4175, // Source-Comm
		"00200b": 4176, // Octagon
		"00200c": 4177, // Adastra
		"00200e": 4178, // NSSLGlobal
		"00200f": 4179, // EBRAINS
		"002011": 4180, // Canopus
		"002013": 4181, // Diversified
		"002015": 4182, // Actis
		"002017": 4183, // Orbotech
		"002018": 164,  // CIS
		"002019": 4184, // Ohler
		"00201a": 2254, // MRV
		"00201c": 4185, // Excel
		"00201d": 4186, // Katana
		"00201e": 4187, // Netquest
		"002020": 4188, // Megatron
		"002021": 4189, // Algorithms
		"002022": 4190, // NMS
		"002023": 4191, // T.C
		"002025": 1999, // Control
		"002026": 4192, // Amkly
		"002029": 4193, // Teleprocessing
		"00202c": 4194, // Welltronix
		"00202d": 4195, // Taiyo
		"00202f": 4196, // Zeta
		"002033": 3758, // Synapse
		"002035": 372,  // IBM
		"002036": 4197, // BMC
		"002037": 732,  // Seagate
		"002039": 4198, // Scinets
		"00203b": 4199, // Wisdm
		"00203c": 4200, // Eurotime
		"00203e": 4201, // LogiCan
		"00203f": 4202, // Juki
		"002040": 125,  // ARRIS
		"002042": 4203, // Datametrics
		"002043": 4204, // Neuron
		"002044": 4205, // Genitech
		"002045": 4206, // ION
		"002046": 4207, // Ciprico
		"002047": 4208, // Steinbrecher
		"002048": 31,   // Marconi
		"002049": 4209, // Comtron
		"00204a": 4210, // Pronet
		"00204b": 4211, // Autocomputer
		"00204c": 3885, // Mitron
		"00204d": 4212, // Inovis
		"002050": 853,  // Korea
		"002051": 4213, // Verilink
		"002052": 4214, // Ragula
		"002054": 4215, // Sycamore
		"002055": 4216, // Altech
		"002056": 4217, // Neoproducts
		"002059": 4218, // Miro
		"00205b": 4219, // Kentrox
		"00205d": 4220, // Nanomatic
		"00205f": 4221, // Gammadata
		"002061": 4222, // GarrettCom
		"002067": 67,   // Private
		"002068": 4223, // Isdyne
		"002069": 4224, // Isdn
		"00206a": 4225, // Osaka
		"00206c": 4226, // Evergreen
		"00206e": 4227, // Xact
		"00206f": 4228, // Flowpoint
		"002070": 4229, // Hynet
		"002071": 4230, // IBR
		"002073": 4231, // Fusion
		"002074": 4232, // Sungwoon
		"002076": 4233, // Reudo
		"002077": 4234, // Kardios
		"002078": 524,  // Runtop
		"002079": 4235, // Mikron
		"00207a": 4236, // WiSE
		"00207b": 421,  // Intel
		"00207c": 4237, // Autec
		"00207e": 4238, // Finecom
		"002080": 4239, // Synergy
		"002081": 738,  // Titan
		"002082": 4240, // Oneac
		"002083": 4241, // Presticom
		"002085": 1855, // Eaton
		"002086": 4242, // Microtech
		"002087": 4243, // Memotec
		"00208a": 4244, // Sonix
		"00208b": 4245, // Lapis
		"00208c": 3848, // Galaxy
		"00208d": 4246, // CMD
		"002091": 4247, // J125
		"002092": 4248, // Chess
		"002093": 4249, // Landings
		"002094": 4250, // Cubix
		"002095": 4251, // Riva
		"002096": 4252, // Invensys
		"002098": 2995, // Hectronic
		"00209b": 4253, // Ersat
		"00209f": 4254, // Mercury
		"0020a1": 4255, // Dovatron
		"0020a3": 1526, // Harmonic
		"0020a4": 4256, // Multipoint
		"0020a5": 4257, // API
		"0020a7": 4258, // Pairgain
		"0020a8": 4259, // Sast
		"0020ab": 43,   // Micro
		"0020ad": 4260, // Linq
		"0020af": 160,  // 3Com
		"0020b1": 4261, // Comtech
		"0020b6": 1610, // Agile
		"0020b9": 4262, // Metricom
		"0020bb": 4263, // ZAX
		"0020bf": 4264, // Aehr
		"0020c0": 4265, // Pulse
		"0020c1": 344,  // SAXA
		"0020c3": 4266, // Counter
		"0020c4": 4267, // Inet
		"0020c5": 4268, // Eagle
		"0020c6": 4269, // Nectec
		"0020c8": 4270, // Larscom
		"0020c9": 4271, // Victron
		"0020cb": 4272, // Pretec
		"0020cc": 1565, // Digital
		"0020cd": 4273, // Hybrid
		"0020d0": 4274, // Versalynx
		"0020d4": 16,   // Cabletron
		"0020d5": 4275, // Vipa
		"0020d6": 700,  // Breezecom
		"0020d8": 76,   // Nortel
		"0020d9": 2154, // Panasonic
		"0020db": 773,  // Xnet
		"0020dd": 4276, // Cybertec
		"0020e0": 2247, // Actiontec
		"0020e1": 4277, // Alamar
		"0020e4": 4278, // Hsing
		"0020e8": 4279, // Datatrek
		"0020e9": 4280, // Dantel
		"0020ea": 4281, // Efficient
		"0020ec": 4282, // Techware
		"0020ed": 1929, // Giga-Byte
		"0020ee": 4283, // Gtech
		"0020ef": 4284, // USC
		"0020f2": 11,   // Oracle
		"0020f3": 4285, // Raynet
		"0020f4": 4286, // Spectrix
		"0020f5": 4287, // Pandatel
		"0020f7": 4288, // Cyberdata
		"0020f8": 4289, // Carrera
		"0020f9": 4290, // Paralink
		"0020fa": 4291, // GDE
		"0020fb": 4292, // Octel
		"0020fc": 4293, // Matrox
		"0020fd": 4294, // ITV
		"0020ff": 4295, // Symmetrical
		"002100": 1450, // Gemtek
		"002101": 4296, // Aplicaciones
		"002102": 4297, // UpdateLogic
		"002103": 4298, // GHI
		"002104": 4299, // Gigaset
		"002107": 4300, // Seowonintech
		"002108": 457,  // Nokia
		"002109": 457,  // Nokia
		"00210a": 4301, // byd:sign
		"00210c": 4302, // Cymtec
		"00210f": 4303, // Cernium
		"002110": 4304, // Clearbox
		"002111": 4305, // Uniphone
		"002114": 4306, // Hylab
		"002116": 4307, // Transcon
		"002117": 4308, // Tellord
		"002118": 4309, // Athena
		"002119": 152,  // Samsung
		"00211a": 4310, // LInTech
		"00211b": 2,    // Cisco
		"00211c": 2,    // Cisco
		"00211d": 4311, // Dataline
		"00211e": 125,  // ARRIS
		"002120": 4312, // Sequel
		"002121": 4313, // VRmagic
		"002122": 4314, // Chip-pro
		"002124": 4315, // Optos
		"002127": 1595, // TP-Link
		"002128": 11,   // Oracle
		"002129": 1783, // Linksys
		"00212a": 4316, // Audiovox
		"00212d": 4317, // Scimolex
		"00212e": 4318, // dresden-elektronik
		"002131": 4319, // Blynke
		"002132": 4320, // Masterclock
		"002134": 4321, // Brandywine
		"002136": 125,  // ARRIS
		"002138": 4322, // Cepheid
		"002139": 4323, // Escherlogic
		"00213a": 4324, // Winchester
		"00213b": 4325, // Berkshire
		"00213c": 4326, // AliphCom
		"00213e": 2715, // TomTom
		"00213f": 4327, // A-Team
		"002140": 4328, // EN
		"002141": 4329, // Radlive
		"002143": 125,  // ARRIS
		"002145": 4330, // Semptian
		"002146": 3537, // Sanmina-SCI
		"002147": 1427, // Nintendo
		"00214c": 152,  // Samsung
		"00214f": 430,  // Alpsalpine
		"002150": 4331, // Eyeview
		"002151": 4332, // Millinet
		"002153": 4333, // SeaMicro
		"002154": 4334, // D-TACQ
		"002155": 2,    // Cisco
		"002156": 2,    // Cisco
		"002159": 826,  // Juniper
		"00215a": 302,  // HP
		"00215b": 4335, // SenseAnywhere
		"00215c": 421,  // Intel
		"00215d": 421,  // Intel
		"00215e": 372,  // IBM
		"00215f": 4336, // IHSE
		"002160": 4337, // Hidea
		"002161": 4338, // Yournet
		"002162": 76,   // Nortel
		"002163": 2545, // Askey
		"002165": 4339, // Presstek
		"002166": 4340, // NovAtel
		"002168": 4341, // iVeia
		"002169": 4342, // Prologix
		"00216a": 421,  // Intel
		"00216b": 421,  // Intel
		"00216c": 4343, // Odva
		"00216d": 4344, // Soltech
		"00216f": 4345, // SymCom
		"002170": 102,  // Dell
		"002179": 4346, // IOGEAR
		"00217a": 4347, // Sejin
		"00217b": 4348, // Bastec
		"00217c": 1939, // 2Wire
		"00217f": 4349, // Intraco
		"002180": 125,  // ARRIS
		"002182": 4350, // SandLinks
		"002185": 1812, // Micro-Star
		"002187": 4351, // Imacs
		"002188": 4352, // EMC
		"002189": 4353, // AppTech
		"00218b": 4354, // Wescon
		"00218c": 4355, // TopControl
		"00218e": 4356, // Mekics
		"002190": 4357, // Goliath
		"002191": 803,  // D-Link
		"002194": 4358, // Ping
		"002197": 1115, // Elitegroup
		"002199": 4359, // Vacon
		"00219b": 102,  // Dell
		"00219c": 4360, // Honeywld
		"00219d": 4361, // Adesys
		"00219e": 101,  // Sony
		"00219f": 4362, // Satel
		"0021a0": 2,    // Cisco
		"0021a1": 2,    // Cisco
		"0021a2": 4363, // EKE-Electronics
		"0021a3": 4364, // Micromint
		"0021a4": 4365, // Dbii
		"0021a6": 4366, // Videotec
		"0021a8": 4367, // Telephonics
		"0021aa": 457,  // Nokia
		"0021ab": 457,  // Nokia
		"0021b1": 1565, // Digital
		"0021b2": 4368, // Fiberblaze
		"0021b5": 4369, // Galvanic
		"0021b7": 613,  // Lexmark
		"0021b8": 4370, // Inphi
		"0021bc": 4371, // Zala
		"0021bd": 1427, // Nintendo
		"0021c2": 4372, // GL
		"0021c3": 4373, // CoRNELL
		"0021c4": 4374, // Consilium
		"0021c5": 4375, // 3DSP
		"0021c7": 4376, // Russound
		"0021c8": 4377, // LOHUIS
		"0021cc": 833,  // Flextronics
		"0021cd": 4378, // LiveTV
		"0021ce": 4379, // NTC-Metrotek
		"0021d1": 152,  // Samsung
		"0021d2": 152,  // Samsung
		"0021d5": 4380, // X2E
		"0021d7": 2,    // Cisco
		"0021d8": 2,    // Cisco
		"0021d9": 4381, // Sekonic
		"0021dd": 783,  // Northstar
		"0021e0": 4382, // CommAgility
		"0021e1": 76,   // Nortel
		"0021e3": 4383, // SerialTek
		"0021e4": 4384, // I-WIN
		"0021e7": 4385, // Informatics
		"0021e8": 2057, // Murata
		"0021e9": 545,  // Apple
		"0021eb": 4386, // ESP
		"0021ec": 4387, // Solutronic
		"0021ed": 4388, // Telegesis
		"0021ef": 4389, // Kapsys
		"0021f0": 4390, // EW3
		"0021f2": 4391, // EASY3CALL
		"0021f3": 4392, // Si14
		"0021f4": 4393, // INRange
		"0021f6": 11,   // Oracle
		"0021f8": 4394, // Enseo
		"0021f9": 4395, // WIRECOM
		"0021fa": 4396, // A4SP
		"0021fb": 869,  // LG
		"0021fc": 457,  // Nokia
		"0021fe": 457,  // Nokia
		"002200": 372,  // IBM
		"002201": 2115, // Aksys
		"002203": 4397, // Glensound
		"002204": 4398, // Koratek
		"002205": 4399, // WeLink
		"002206": 4400, // Cyberdyne
		"002208": 4401, // Certicom
		"00220a": 4402, // OnLive
		"00220c": 2,    // Cisco
		"00220d": 2,    // Cisco
		"00220f": 4403, // MoCA
		"002210": 125,  // ARRIS
		"002211": 4404, // Rohati
		"002212": 4405, // CAI
		"002213": 4406, // PCI
		"002215": 1806, // ASUS
		"002217": 4407, // Neat
		"002218": 3820, // Akamai
		"002219": 102,  // Dell
		"00221b": 4408, // Morega
		"00221c": 67,   // Private
		"00221d": 4409, // Freegene
		"00221f": 4410, // eSang
		"002220": 513,  // Mitac
		"002223": 4411, // TimeKeeping
		"002226": 4412, // Avaak
		"002227": 4413, // uv-electronic
		"002229": 4414, // Compumedics
		"00222a": 4415, // SoundEar
		"00222b": 4416, // Nucomm
		"00222c": 4417, // Ceton
		"00222d": 741,  // SMC
		"00222e": 4418, // maintech
		"002230": 4419, // FutureLogic
		"002231": 4420, // SMT&C
		"002234": 4421, // Corventis
		"002235": 4422, // Strukton
		"002238": 4423, // Logiplus
		"00223b": 4424, // Communication
		"00223d": 4425, // JumpGen
		"00223e": 4426, // IRTrans
		"00223f": 1368, // Netgear
		"002241": 545,  // Apple
		"002242": 4427, // Alacron
		"002243": 3005, // Azurewave
		"002247": 4428, // DAC
		"002248": 612,  // Microsoft
		"00224a": 4429, // Raylase
		"00224b": 4430, // Airtech
		"00224c": 1427, // Nintendo
		"00224d": 513,  // Mitac
		"00224e": 4431, // SEEnergy
		"00224f": 4432, // Byzoro
		"002251": 4433, // Lumasense
		"002253": 4434, // Entorian
		"002255": 2,    // Cisco
		"002256": 2,    // Cisco
		"002257": 160,  // 3Com
		"00225b": 4435, // Teradici
		"00225c": 4436, // Multimedia
		"00225e": 4437, // Uwin
		"00225f": 2874, // Liteon
		"002260": 4438, // AFREEY
		"002264": 302,  // HP
		"002265": 457,  // Nokia
		"002266": 457,  // Nokia
		"002267": 76,   // Nortel
		"00226a": 934,  // Honeywell
		"00226b": 1783, // Linksys
		"00226c": 4439, // LinkSprite
		"00226e": 4440, // Gowell
		"00226f": 4441, // 3onedata
		"002270": 4442, // ABK
		"002273": 4443, // Techway
		"002274": 4444, // FamilyPhone
		"002275": 2469, // Belkin
		"00227b": 591,  // Apogee
		"00227f": 2738, // Ruckus
		"002280": 4445, // A2B
		"002281": 4446, // Daintree
		"002283": 826,  // Juniper
		"002285": 4447, // Nomus
		"002286": 4448, // Astron
		"002288": 4449, // Sagrad
		"00228c": 4450, // Photon
		"00228d": 4451, // GBS
		"00228e": 4452, // TV-Numeric
		"00228f": 4453, // Cnrs
		"002290": 2,    // Cisco
		"002291": 2,    // Cisco
		"002292": 4454, // Cinetal
		"002293": 3032, // ZTE
		"002294": 2849, // Kyocera
		"002296": 4455, // LinoWave
		"002298": 101,  // Sony
		"002299": 4333, // SeaMicro
		"00229a": 4456, // Lastar
		"00229b": 4457, // AverLogic
		"00229c": 4458, // Verismo
		"00229d": 4459, // Pyung-HWA
		"0022a0": 4460, // Aptiv
		"0022a2": 4461, // Xtramus
		"0022a4": 1939, // 2Wire
		"0022a6": 101,  // Sony
		"0022a7": 4462, // Tyco
		"0022a8": 4463, // Ouman
		"0022a9": 869,  // LG
		"0022aa": 1427, // Nintendo
		"0022ad": 4464, // Telesis
		"0022ae": 4465, // Mattel
		"0022b0": 803,  // D-Link
		"0022b1": 4466, // Elbit
		"0022b2": 4467, // 4RF
		"0022b4": 125,  // ARRIS
		"0022b5": 4468, // Novita
		"0022b8": 4469, // Norcott
		"0022bb": 4470, // beyerdynamic
		"0022bd": 2,    // Cisco
		"0022be": 2,    // Cisco
		"0022c3": 4471, // Zeeport
		"0022c4": 4472, // epro
		"0022c5": 4473, // InfORSON
		"0022c6": 4474, // Sutus
		"0022c9": 4475, // Lenord
		"0022cb": 4476, // IONODES
		"0022cc": 4477, // SciLog
		"0022cd": 4478, // Ared
		"0022cf": 4479, // Planex
		"0022d0": 4480, // Polar
		"0022d3": 4481, // Hub-Tech
		"0022d4": 4482, // ComWorth
		"0022d6": 4483, // Cypak
		"0022d7": 1427, // Nintendo
		"0022d9": 4484, // Fortex
		"0022da": 4485, // Anatek
		"0022db": 4486, // Translogic
		"0022dd": 4487, // Protecta
		"0022e1": 4488, // ZORT
		"0022e3": 4489, // Amerigon
		"0022e4": 4490, // Apass
		"0022e5": 4491, // Fisher-Rosemount
		"0022e8": 4492, // Applition
		"0022e9": 4493, // ProVision
		"0022ea": 4494, // Rustelcom
		"0022ec": 4495, // Idealbt
		"0022ee": 4496, // Algo
		"0022ef": 4497, // iWDL
		"0022f1": 67,   // Private
		"0022f2": 4498, // SunPower
		"0022f3": 3206, // Sharp
		"0022f4": 4499, // AMPAK
		"0022f6": 4500, // Syracuse
		"0022f7": 4501, // Conceptronic
		"0022f8": 4502, // PIMA
		"0022f9": 4503, // Pollin
		"0022fa": 421,  // Intel
		"0022fb": 421,  // Intel
		"0022fc": 457,  // Nokia
		"0022fd": 457,  // Nokia
		"0022ff": 4504, // Nivis
		"002300": 4505, // Cayee
		"002301": 4506, // Witron
		"002304": 2,    // Cisco
		"002305": 2,    // Cisco
		"002306": 430,  // Alpsalpine
		"002308": 2640, // Arcadyan
		"002309": 4507, // Janam
		"00230a": 4508, // ARBURG
		"00230b": 125,  // ARRIS
		"00230c": 4509, // Clover
		"00230d": 76,   // Nortel
		"00230e": 4510, // Gorba
		"00230f": 4511, // Hirsch
		"002310": 4512, // LNC
		"002311": 4513, // Gloscom
		"002312": 545,  // Apple
		"002313": 4514, // Qool
		"002314": 421,  // Intel
		"002315": 421,  // Intel
		"002316": 4515, // Kisan
		"002317": 4516, // Lasercraft
		"002318": 35,   // Toshiba
		"002319": 4517, // Sielox
		"00231a": 4518, // ITF
		"00231c": 4519, // Fourier
		"00231d": 4520, // Deltacom
		"00231f": 4521, // Guangda
		"002320": 4522, // Nicira
		"002321": 4523, // Avitech
		"002323": 4524, // Zylin
		"002324": 2292, // G-PRO
		"002325": 4525, // IOLAN
		"002326": 4,    // Fujitsu
		"002327": 4526, // Shouyo
		"002329": 4527, // DDRdrive
		"00232b": 4528, // IRD
		"00232c": 4529, // Senticare
		"00232d": 4530, // SandForce
		"00232e": 4531, // Kedah
		"002330": 4532, // Dizipia
		"002331": 1427, // Nintendo
		"002332": 545,  // Apple
		"002333": 2,    // Cisco
		"002334": 2,    // Cisco
		"002335": 4533, // Linkflex
		"002338": 4534, // OJ-Electronics
		"002339": 152,  // Samsung
		"00233a": 152,  // Samsung
		"00233b": 4535, // C-Matic
		"00233c": 4536, // Alflex
		"00233d": 4537, // Laird
		"00233f": 4538, // Purechoice
		"002340": 4539, // MiXTelematics
		"002341": 887,  // Vanderbilt
		"002343": 4540, // TEM
		"002345": 101,  // Sony
		"002346": 4541, // Vestac
		"002348": 2049, // Sagemcom
		"00234a": 67,   // Private
		"00234b": 4542, // Inyuan
		"00234c": 4543, // KTC
		"002351": 1939, // 2Wire
		"002354": 1806, // ASUS
		"002357": 4544, // Pitronot
		"002358": 4545, // Systel
		"002359": 2080, // Benchmark
		"00235a": 358,  // Compal
		"00235b": 4546, // Gulfstream
		"00235c": 4547, // Aprius
		"00235d": 2,    // Cisco
		"00235e": 2,    // Cisco
		"002360": 4548, // Lookit
		"002361": 4549, // Unigen
		"002368": 768,  // Zebra
		"002369": 1783, // Linksys
		"00236a": 4550, // SmartRG
		"00236b": 4551, // Xembedded
		"00236c": 545,  // Apple
		"00236d": 4552, // ResMed
		"00236e": 4553, // Burster
		"002370": 822,  // Snell
		"002373": 4554, // GridIron
		"002374": 125,  // ARRIS
		"002375": 125,  // ARRIS
		"002376": 1341, // HTC
		"002377": 4555, // Isotek
		"00237a": 4556, // RIM
		"00237b": 4557, // Whdi
		"00237c": 4558, // Neotion
		"00237d": 302,  // HP
		"00237e": 3563, // Elster
		"00237f": 542,  // Plantronics
		"002380": 4559, // Nanoteq
		"002381": 4560, // Lengda
		"002383": 4561, // InMage
		"002384": 4562, // GGH
		"002385": 4563, // Antipode
		"002387": 4564, // ThinkFlood
		"00238a": 4565, // Ciena
		"00238b": 3072, // Quanta
		"00238c": 67,   // Private
		"002390": 4566, // Algolware
		"002391": 4567, // Maxian
		"002392": 4568, // Proteus
		"002393": 4569, // Ajinextek
		"002394": 4570, // Samjeon
		"002395": 125,  // ARRIS
		"002396": 4571, // Andes
		"002397": 2274, // Westell
		"002399": 152,  // Samsung
		"00239b": 3563, // Elster
		"00239c": 826,  // Juniper
		"00239d": 4572, // Mapower
		"0023a1": 176,  // Trend
		"0023a2": 125,  // ARRIS
		"0023a3": 125,  // ARRIS
		"0023a5": 4573, // SageTV
		"0023a6": 4574, // E-Mon
		"0023a8": 4575, // Marshall
		"0023aa": 4576, // HFR
		"0023ab": 2,    // Cisco
		"0023ac": 2,    // Cisco
		"0023ad": 4577, // Xmark
		"0023ae": 102,  // Dell
		"0023af": 125,  // ARRIS
		"0023b0": 4578, // CoMXION
		"0023b1": 4579, // Longcheer
		"0023b3": 4580, // Lyyn
		"0023b4": 457,  // Nokia
		"0023b5": 4581, // Ortana
		"0023b7": 4582, // Q-Light
		"0023ba": 4583, // Chroma
		"0023bc": 4584, // EQ-Sys
		"0023bf": 4585, // Mainpine
		"0023c0": 4586, // Broadway
		"0023c2": 4587, // SAMSUNG
		"0023c3": 4588, // LogMeIn
		"0023c6": 741,  // SMC
		"0023c7": 4589, // AVSystem
		"0023c8": 4590, // Team-R
		"0023cc": 1427, // Nintendo
		"0023cd": 1595, // TP-Link
		"0023cf": 4591, // Cummins-Allison
		"0023d1": 4592, // TRG
		"0023d2": 4593, // Inhand
		"0023d5": 4594, // WAREMA
		"0023d6": 152,  // Samsung
		"0023d7": 152,  // Samsung
		"0023d8": 4595, // Ball-It
		"0023d9": 4596, // Banner
		"0023db": 4597, // saxnet
		"0023dc": 4598, // Benein
		"0023de": 4599, // Ansync
		"0023df": 545,  // Apple
		"0023e3": 4600, // Microtronic
		"0023e4": 4601, // IPnect
		"0023e5": 4602, // IPaXiom
		"0023e7": 4603, // Hinke
		"0023e8": 4604, // Demco
		"0023e9": 289,  // F5
		"0023ea": 2,    // Cisco
		"0023eb": 2,    // Cisco
		"0023ec": 4605, // Algorithmix
		"0023ed": 125,  // ARRIS
		"0023ee": 125,  // ARRIS
		"0023f1": 101,  // Sony
		"0023f2": 4606, // TVLogic
		"0023f3": 4607, // Glocom
		"0023f4": 4608, // Masternaut
		"0023f6": 4609, // Softwell
		"0023f7": 67,   // Private
		"0023f8": 2693, // ZyXEL
		"0023f9": 4610, // Double-Take
		"0023fe": 4611, // Biodevices
		"002400": 76,   // Nortel
		"002401": 803,  // D-Link
		"002402": 4612, // Op-Tection
		"002403": 457,  // Nokia
		"002404": 457,  // Nokia
		"002406": 4613, // Pointmobile
		"002409": 4614, // Toro
		"00240b": 4615, // Virtual
		"00240c": 4616, // DELEC
		"00240d": 4617, // OnePath
		"002410": 4618, // NUETEQ
		"002411": 4619, // PharmaSmart
		"002412": 4620, // Benign
		"002413": 2,    // Cisco
		"002414": 2,    // Cisco
		"002419": 67,   // Private
		"00241b": 4621, // iWOW
		"00241c": 4622, // FuGang
		"00241d": 1929, // Giga-Byte
		"00241e": 1427, // Nintendo
		"00241f": 4623, // DCT-Delta
		"002420": 4624, // NetUP
		"002421": 1812, // Micro-Star
		"002423": 3005, // Azurewave
		"002427": 4625, // SSI
		"002428": 4626, // EnergyICT
		"00242e": 4627, // Datastrip
		"00242f": 4628, // Micron
		"002430": 4629, // Ruby
		"002431": 4630, // Uni-v
		"002432": 4631, // Neostar
		"002433": 430,  // Alpsalpine
		"002434": 4632, // Lectrosonics
		"002435": 4633, // Wide
		"002436": 545,  // Apple
		"002438": 90,   // Brocade
		"00243a": 4634, // Ludl
		"00243b": 4635, // CSSI
		"00243c": 4636, // S.A.A.A
		"00243f": 4637, // Storwize
		"002442": 4638, // Axona
		"002443": 76,   // Nortel
		"002444": 1427, // Nintendo
		"002445": 202,  // Adtran
		"002446": 4639, // MMB
		"002447": 4640, // Kaztek
		"00244a": 4641, // Voyant
		"00244b": 4642, // Perceptron
		"00244d": 4643, // Hokkaido
		"00244e": 4644, // RadChips
		"00244f": 4645, // Asantron
		"002450": 2,    // Cisco
		"002451": 2,    // Cisco
		"002452": 1655, // Silicon
		"002454": 152,  // Samsung
		"002455": 4646, // MuLogic
		"002456": 1939, // 2Wire
		"00245b": 4647, // Raidon
		"00245c": 4648, // Design-Com
		"00245e": 4649, // Hivision
		"002462": 4650, // Rayzone
		"002463": 4651, // Phybridge
		"002464": 4652, // Bridge
		"002465": 4653, // Elentec
		"002466": 3993, // Unitron
		"002467": 4654, // AOC
		"002468": 4655, // Sumavision
		"00246b": 4656, // Covia
		"00246c": 1685, // Aruba
		"00246d": 4657, // Weinzierl
		"00246f": 4658, // Onda
		"002473": 160,  // 3Com
		"002476": 4659, // TAP.tv
		"002477": 4660, // Tibbo
		"002478": 4661, // Mag
		"00247b": 2247, // Actiontec
		"00247c": 457,  // Nokia
		"00247d": 457,  // Nokia
		"00247f": 76,   // Nortel
		"002480": 4662, // Meteocontrol
		"002481": 302,  // HP
		"002482": 2738, // Ruckus
		"002483": 869,  // LG
		"002485": 4663, // ConteXtream
		"002486": 4664, // DesignArt
		"00248a": 4665, // Kaga
		"00248b": 4666, // Hybus
		"00248c": 1806, // ASUS
		"00248d": 101,  // Sony
		"00248f": 4667, // DO-Monix
		"002490": 152,  // Samsung
		"002491": 152,  // Samsung
		"002492": 687,  // Motorola
		"002493": 125,  // ARRIS
		"002495": 125,  // ARRIS
		"002496": 4668, // Ginzinger
		"002497": 2,    // Cisco
		"002498": 2,    // Cisco
		"002499": 4669, // Aquila
		"00249d": 4670, // NES
		"00249e": 4671, // ADC-Elektronik
		"0024a0": 125,  // ARRIS
		"0024a1": 125,  // ARRIS
		"0024a3": 4672, // Sonim
		"0024a4": 4673, // Siklu
		"0024a5": 1077, // Buffalo
		"0024aa": 4674, // Dycor
		"0024ab": 4675, // A7
		"0024ae": 4676, // Idemia
		"0024af": 1238, // Dish
		"0024b0": 4677, // Esab
		"0024b1": 4678, // Coulomb
		"0024b2": 1368, // Netgear
		"0024b3": 4679, // Graf-Syteco
		"0024b4": 4680, // ESCATRONIC
		"0024b5": 76,   // Nortel
		"0024b6": 732,  // Seagate
		"0024b7": 4681, // GridPoint
		"0024bb": 4682, // CENTRAL
		"0024bc": 4683, // HuRob
		"0024bd": 4684, // Hainzl
		"0024be": 101,  // Sony
		"0024bf": 4685, // Ciat
		"0024c1": 125,  // ARRIS
		"0024c2": 4686, // Asumo
		"0024c3": 2,    // Cisco
		"0024c4": 2,    // Cisco
		"0024c6": 4687, // Hager
		"0024c7": 4688, // Mobilarm
		"0024ca": 4689, // Tobii
		"0024cb": 4690, // Autonet
		"0024ce": 4691, // Exeltech
		"0024d1": 1142, // Thomson
		"0024d2": 2545, // Askey
		"0024d3": 4692, // QUALICA
		"0024d5": 4693, // Winward
		"0024d6": 421,  // Intel
		"0024d7": 421,  // Intel
		"0024d9": 4694, // BICOM
		"0024da": 4695, // Innovar
		"0024dc": 826,  // Juniper
		"0024dd": 4696, // Centrak
		"0024de": 4697, // GLOBAL
		"0024df": 4698, // Digitalbox
		"0024e0": 4699, // DS
		"0024e1": 4700, // Convey
		"0024e4": 4701, // Withings
		"0024e5": 4702, // Seer
		"0024e7": 4703, // Plaster
		"0024e8": 102,  // Dell
		"0024e9": 152,  // Samsung
		"0024eb": 4704, // ClearPath
		"0024ee": 4705, // Wynmax
		"0024ef": 101,  // Sony
		"0024f0": 4706, // Seanodes
		"0024f3": 1427, // Nintendo
		"0024f4": 4707, // Kaminario
		"0024f6": 4708, // Miyoshi
		"0024f7": 2,    // Cisco
		"0024f8": 2925, // Technical
		"0024f9": 2,    // Cisco
		"0024fb": 67,   // Private
		"0024fc": 4709, // QuoPin
		"0024fd": 3004, // Accedian
		"0024fe": 621,  // AVM
		"0024ff": 2026, // QLogic
		"002500": 545,  // Apple
		"002502": 4710, // NaturalPoint
		"002503": 372,  // IBM
		"002504": 4711, // Valiant
		"002507": 4712, // ASTAK
		"00250b": 4713, // Centrofactor
		"00250c": 4714, // Senet
		"002511": 1115, // Elitegroup
		"002512": 3032, // ZTE
		"002515": 3188, // SFR
		"002517": 4715, // Venntis
		"002519": 4716, // Viaas
		"00251c": 4717, // EDT
		"00251e": 4718, // Rotel
		"002521": 4719, // Logitek
		"002522": 4720, // ASRock
		"002523": 4721, // OCP
		"002524": 4722, // Lightcomm
		"002525": 4723, // CTERA
		"002526": 4724, // Genuine
		"002527": 4725, // Bitrode
		"00252c": 4726, // Entourage
		"00252d": 4727, // Kiryung
		"00252f": 4728, // Energy
		"002530": 4729, // Aetas
		"002533": 4730, // Wittenstein
		"002535": 4731, // Minimax
		"002537": 4732, // Runcom
		"002538": 152,  // Samsung
		"002539": 4733, // IfTA
		"00253a": 4734, // CEVA
		"00253c": 1939, // 2Wire
		"002540": 4735, // Quasar
		"002542": 4736, // Pittasoft
		"002543": 4737, // Moneytech
		"002544": 4738, // LoJack
		"002545": 2,    // Cisco
		"002546": 2,    // Cisco
		"002547": 457,  // Nokia
		"002548": 457,  // Nokia
		"002549": 4739, // Jeorich
		"00254a": 4740, // RingCube
		"00254b": 545,  // Apple
		"00254d": 4741, // Singapore
		"002550": 2094, // Riverbed
		"002551": 4742, // SE-Elektronic
		"002552": 4743, // VXi
		"002554": 4744, // Pixel8
		"002557": 2221, // BlackBerry
		"002558": 4745, // Mpedia
		"002559": 4746, // Syphan
		"00255a": 4747, // Tantalus
		"00255b": 4748, // CoachComm
		"00255c": 48,   // NEC
		"00255d": 4749, // Morningstar
		"00255f": 4750, // SenTec
		"002562": 4751, // interbro
		"002563": 4752, // Luxtera
		"002564": 102,  // Dell
		"002565": 4753, // Vizimax
		"002566": 152,  // Samsung
		"002567": 152,  // Samsung
		"002568": 3335, // Huawei
		"002569": 2049, // Sagemcom
		"002570": 4754, // Eastern
		"002572": 4755, // Nemo-Q
		"002573": 4756, // ST
		"002575": 4757, // FiberPlex
		"002576": 4758, // Neli
		"002577": 4759, // D-BOX
		"00257b": 4760, // STJ
		"00257f": 4761, // CallTechSolution
		"002581": 4762, // x-star
		"002582": 4763, // Maksat
		"002583": 2,    // Cisco
		"002584": 2,    // Cisco
		"002586": 1595, // TP-Link
		"002587": 4764, // Vitality
		"002588": 4765, // Genie
		"002589": 4766, // Hills
		"00258a": 4767, // Pole/Zero
		"00258b": 432,  // Mellanox
		"00258d": 4768, // Haier
		"002591": 4769, // NEXTEK
		"002594": 4770, // Eurodesign
		"002597": 4771, // Kalki
		"00259a": 4772, // CEStronics
		"00259c": 1783, // Linksys
		"00259d": 67,   // Private
		"00259e": 3335, // Huawei
		"00259f": 4773, // TechnoDigital
		"0025a0": 1427, // Nintendo
		"0025a1": 4774, // Enalasys
		"0025a7": 4775, // itron
		"0025a8": 63,   // Kontron
		"0025ac": 4776, // I-Tech
		"0025ae": 612,  // Microsoft
		"0025af": 4777, // CoMFILE
		"0025b0": 4778, // Schmartz
		"0025b1": 4779, // Maya-Creation
		"0025b3": 302,  // HP
		"0025b4": 2,    // Cisco
		"0025b5": 2,    // Cisco
		"0025b7": 4780, // Costar
		"0025b8": 1610, // Agile
		"0025b9": 4781, // Cypress
		"0025bb": 4782, // INNERINT
		"0025bc": 545,  // Apple
		"0025be": 4783, // Tektrap
		"0025c0": 4784, // ZillionTV
		"0025c2": 4785, // RingBell
		"0025c3": 4786, // 21168
		"0025c4": 2738, // Ruckus
		"0025c6": 4787, // kasercorp
		"0025c7": 4788, // altek
		"0025c8": 4789, // S-Access
		"0025ca": 1626, // LS
		"0025ce": 4790, // InnerSpace
		"0025cf": 457,  // Nokia
		"0025d0": 457,  // Nokia
		"0025d2": 4791, // InpegVision
		"0025d3": 3005, // Azurewave
		"0025d5": 4792, // Robonica
		"0025d6": 4793, // Kroger
		"0025d7": 4794, // Cedo
		"0025d9": 4795, // DataFab
		"0025db": 1011, // ATI
		"0025de": 4796, // Probits
		"0025df": 67,   // Private
		"0025e2": 4797, // Everspring
		"0025e3": 4798, // Hanshinit
		"0025e4": 4799, // OMNI-WiFi
		"0025e5": 869,  // LG
		"0025e7": 101,  // Sony
		"0025e8": 4800, // Idaho
		"0025ea": 4801, // Iphion
		"0025ec": 4802, // Humanware
		"0025ed": 4803, // NuVo
		"0025ee": 4804, // Avtex
		"0025ef": 4805, // I-TEC
		"0025f0": 4806, // Suga
		"0025f1": 125,  // ARRIS
		"0025f2": 125,  // ARRIS
		"0025f6": 4807, // netTALK.com
		"0025f9": 4808, // GMK
		"0025fe": 4809, // Pilot
		"002601": 4810, // Cutera
		"002605": 4811, // CC
		"002606": 4812, // RAUMFELD
		"002607": 4813, // Enabling
		"002608": 545,  // Apple
		"002609": 4814, // Phyllis
		"00260a": 2,    // Cisco
		"00260b": 2,    // Cisco
		"00260c": 4815, // Dataram
		"00260d": 4254, // Mercury
		"00260e": 4816, // Ablaze
		"00260f": 4817, // Linn
		"002610": 4818, // Apacewave
		"002611": 4819, // Licera
		"002614": 4820, // Ktnf
		"002615": 4821, // Teracom
		"002616": 4822, // Rosemount
		"002618": 1806, // ASUS
		"002619": 4823, // FRC
		"00261c": 4824, // Neovia
		"002620": 4825, // ISGUS
		"002621": 4826, // InteliCloud
		"002622": 358,  // Compal
		"002623": 4827, // JRD
		"002624": 1142, // Thomson
		"002625": 4828, // MediaSputnik
		"002627": 4829, // Truesell
		"00262a": 4830, // Proxense
		"00262b": 4831, // Wongs
		"00262d": 1592, // Wistron
		"002631": 4832, // Commtact
		"002634": 4833, // Infineta
		"002635": 4834, // Bluetechnix
		"002636": 125,  // ARRIS
		"002637": 152,  // Samsung
		"002639": 4835, // T.M
		"00263a": 4836, // Digitec
		"00263b": 4837, // Onbnetech
		"00263c": 4838, // Bachmann
		"00263d": 4839, // MIA
		"00263e": 1612, // Trapeze
		"00263f": 4840, // LIOS
		"002641": 125,  // ARRIS
		"002642": 125,  // ARRIS
		"002643": 430,  // Alpsalpine
		"002647": 4841, // WFE
		"002648": 3906, // Emitech
		"00264a": 545,  // Apple
		"00264d": 2640, // Arcadyan
		"00264e": 4842, // r2p
		"002650": 1939, // 2Wire
		"002651": 2,    // Cisco
		"002652": 2,    // Cisco
		"002653": 4843, // DaySequerra
		"002654": 160,  // 3Com
		"002655": 302,  // HP
		"002656": 4844, // Sansonic
		"002658": 4845, // T-Platforms
		"002659": 1427, // Nintendo
		"00265a": 803,  // D-Link
		"00265b": 870,  // Hitron
		"00265d": 152,  // Samsung
		"00265f": 152,  // Samsung
		"002660": 4846, // Logiways
		"002661": 4847, // Irumtek
		"002662": 2247, // Actiontec
		"002665": 4848, // ProtectedLogic
		"002666": 1252, // EFM
		"002667": 4849, // Carecom
		"002668": 457,  // Nokia
		"002669": 457,  // Nokia
		"00266a": 4850, // Essensium
		"00266c": 3979, // Inventec
		"00266d": 4851, // MobileAccess
		"00266e": 4852, // Nissho-denki
		"00266f": 4853, // Coordiwise
		"002671": 4854, // AUTOVISION
		"002673": 75,   // Ricoh
		"002675": 4855, // Aztech
		"002676": 4856, // CoMMidt
		"002677": 4857, // Deif
		"002679": 4858, // Euphonic
		"00267c": 4859, // Metz-Werke
		"00267e": 2563, // Parrot
		"00267f": 4860, // Zenterio
		"002680": 4861, // SIL3
		"002681": 4862, // Interspiro
		"002682": 1450, // Gemtek
		"002683": 4863, // Ajoho
		"002688": 826,  // Juniper
		"00268c": 4864, // StarLeaf
		"00268e": 2310, // Alta
		"00268f": 4865, // MTA
		"002691": 2049, // Sagemcom
		"002693": 4866, // QVidium
		"002694": 4867, // Senscient
		"002696": 4868, // NOOLIX
		"002697": 2237, // Alpha
		"002698": 2,    // Cisco
		"002699": 2,    // Cisco
		"00269b": 4869, // SOKRAT
		"00269d": 4870, // M2Mnet
		"00269e": 3072, // Quanta
		"00269f": 67,   // Private
		"0026a0": 4871, // moblic
		"0026a1": 4872, // Megger
		"0026a2": 4873, // Instrumentation
		"0026a5": 4874, // Microrobot
		"0026a6": 4875, // Trixell
		"0026a9": 4876, // Strong
		"0026ab": 46,   // Epson
		"0026ad": 4877, // Arada
		"0026af": 4878, // Duelco
		"0026b0": 545,  // Apple
		"0026b2": 4879, // Setrix
		"0026b3": 4880, // Thales
		"0026b4": 4881, // Ford
		"0026b6": 2545, // Askey
		"0026b7": 4882, // Kingston
		"0026b8": 2247, // Actiontec
		"0026b9": 102,  // Dell
		"0026ba": 125,  // ARRIS
		"0026bb": 545,  // Apple
		"0026c0": 4883, // EnergyHub
		"0026c1": 4884, // Artray
		"0026c2": 4885, // SCDI
		"0026c3": 4886, // Insightek
		"0026c6": 421,  // Intel
		"0026c7": 421,  // Intel
		"0026c9": 4887, // Proventix
		"0026ca": 2,    // Cisco
		"0026cb": 2,    // Cisco
		"0026cc": 457,  // Nokia
		"0026cd": 4888, // PurpleComm
		"0026d0": 4889, // Semihalf
		"0026d2": 4890, // Pcube
		"0026d4": 4891, // IRCA
		"0026d9": 125,  // ARRIS
		"0026df": 4892, // TaiDoc
		"0026e0": 4893, // Asiteq
		"0026e2": 869,  // LG
		"0026e3": 4894, // DTI
		"0026e6": 4895, // Visionhitech
		"0026e8": 2057, // Murata
		"0026e9": 4896, // SP
		"0026ea": 4897, // Cheerchip
		"0026ed": 3032, // ZTE
		"0026ee": 4898, // TKM
		"0026f0": 4899, // cTrixs
		"0026f2": 1368, // Netgear
		"0026f3": 741,  // SMC
		"0026f4": 4900, // Nesslab
		"0026f5": 4901, // XRPLUS
		"0026f7": 4902, // Nivetti
		"0026fa": 4903, // BandRich
		"0026fc": 4904, // AcSiP
		"0026fe": 4905, // MKD
		"0026ff": 2221, // BlackBerry
		"002701": 4906, // IncOstartec
		"002702": 4907, // SolarEdge
		"002703": 1406, // Testech
		"002705": 4908, // Sectronic
		"002706": 4909, // Yoisys
		"002709": 1427, // Nintendo
		"00270b": 4910, // Adura
		"00270c": 2,    // Cisco
		"00270d": 2,    // Cisco
		"00270e": 421,  // Intel
		"00270f": 4911, // Envisionnovation
		"002710": 421,  // Intel
		"002711": 4912, // LanPro
		"002712": 4913, // MaxVision
		"002714": 4914, // Grainmustards
		"002716": 4915, // Adachi-Syokai
		"002719": 1595, // TP-Link
		"00271a": 4916, // Geenovo
		"00271c": 4254, // Mercury
		"00271e": 4917, // Xagyl
		"00271f": 4918, // MIPRO
		"002722": 2979, // Ubiquiti
		"002790": 2,    // Cisco
		"0027e3": 2,    // Cisco
		"0027f8": 90,   // Brocade
		"00289f": 4330, // Semptian
		"0028f8": 421,  // Intel
		"0029c2": 2,    // Cisco
		"002a10": 2,    // Cisco
		"002a6a": 2,    // Cisco
		"002aaf": 4919, // LARsys-Automation
		"002b67": 4920, // LCFC
		"002cc8": 2,    // Cisco
		"002d76": 4921, // TITECH
		"002db3": 4499, // AMPAK
		"002ec7": 3335, // Huawei
		"002f5c": 2,    // Cisco
		"002fd9": 4922, // Fiberhome
		"003000": 4923, // Allwell
		"003001": 4924, // SMP
		"003002": 4925, // Expand
		"003003": 4926, // Phasys
		"003004": 4927, // Leadtek
		"003006": 4928, // Superpower
		"003007": 4929, // Opti
		"003009": 4930, // Tachion
		"00300a": 4855, // Aztech
		"00300b": 4931, // mPHASE
		"00300c": 4932, // Congruency
		"00300d": 4933, // MMC
		"003010": 4934, // Visionetics
		"003011": 837,  // HMS
		"003012": 1565, // Digital
		"003013": 48,   // NEC
		"003014": 4935, // Divio
		"003016": 4936, // Ishida
		"003019": 2,    // Cisco
		"00301a": 4937, // Smartbridges
		"00301b": 4938, // Shuttle
		"00301d": 4939, // Skystream
		"00301e": 160,  // 3Com
		"00301f": 1484, // Optical
		"003020": 4940, // TSI
		"003021": 4278, // Hsing
		"003023": 1675, // Cogent
		"003024": 2,    // Cisco
		"003025": 4941, // Checkout
		"003027": 4942, // Kerbango
		"003029": 4943, // Opicom
		"00302b": 4944, // Inalp
		"00302c": 4945, // Sylantro
		"003030": 4946, // Harmonix
		"003031": 4947, // Lightwave
		"003032": 4948, // MagicRam
		"003034": 4949, // SET
		"003035": 4950, // Corning
		"003038": 4951, // XCP
		"00303a": 4952, // Maatel
		"00303b": 4953, // PowerCom
		"00303c": 4954, // Onnto
		"00303d": 4955, // IVA
		"00303e": 4956, // Radcom
		"00303f": 4957, // TurboComm
		"003040": 2,    // Cisco
		"003043": 4958, // Idream
		"003044": 4959, // CradlePoint
		"003046": 4960, // Controlled
		"003049": 4961, // Bryant
		"00304b": 4962, // Orbacom
		"00304c": 4963, // Appian
		"00304d": 4964, // ESI
		"00304f": 4965, // PLANET
		"003050": 4966, // Versa
		"003052": 4967, // Elastic
		"003053": 4968, // Basler
		"003054": 3798, // Castlenet
		"003056": 837,  // HMS
		"003057": 4969, // QTelNet
		"003059": 63,   // Kontron
		"00305a": 4970, // Telgen
		"00305b": 4971, // Toko
		"00305c": 4972, // SMAR
		"00305d": 4973, // Digitra
		"00305f": 4974, // Hasselblad
		"003060": 4975, // Powerfile
		"003061": 4976, // MobyTEL
		"003063": 4977, // Santera
		"003064": 4978, // Adlink
		"003065": 545,  // Apple
		"003066": 4979, // RFM
		"003068": 1654, // Cybernetics
		"003069": 4980, // Impacct
		"00306b": 4981, // Cmos
		"00306c": 4982, // Hitex
		"00306d": 827,  // Lucent
		"00306e": 302,  // HP
		"00306f": 4983, // Seyeon
		"003070": 4984, // 1Net
		"003071": 2,    // Cisco
		"003072": 4985, // Intellibyte
		"003074": 4986, // Equiinet
		"003075": 4987, // Adtech
		"003076": 4988, // Akamba
		"003077": 4989, // Onprem
		"003078": 2,    // Cisco
		"003079": 4990, // Cqos
		"00307a": 587,  // Advanced
		"00307b": 2,    // Cisco
		"00307c": 4991, // Adid
		"00307e": 4992, // Redflex
		"00307f": 4993, // Irlan
		"003080": 2,    // Cisco
		"003083": 4994, // Ivron
		"003085": 2,    // Cisco
		"003088": 306,  // Ericsson
		"00308b": 4995, // Brix
		"00308c": 2084, // Quantum
		"00308d": 4996, // Pinnacle
		"00308f": 4997, // MICRILOR
		"003090": 4998, // Cyra
		"003092": 63,   // Kontron
		"003093": 4999, // Sonnet
		"003094": 2,    // Cisco
		"003096": 2,    // Cisco
		"00309b": 5000, // Smartware
		"00309e": 5001, // Workbit
		"00309f": 5002, // Amber
		"0030a1": 5003, // WEBGATE
		"0030a2": 5004, // Lightner
		"0030a3": 2,    // Cisco
		"0030a6": 5005, // Vianet
		"0030a7": 5006, // Schweitzer
		"0030a8": 5007, // OL'E
		"0030a9": 5008, // Netiverse
		"0030ab": 957,  // Delta
		"0030ad": 5009, // Shanghai
		"0030af": 934,  // Honeywell
		"0030b0": 5010, // Convergenet
		"0030b1": 5011, // TrunkNet
		"0030b4": 5012, // Intersil
		"0030b6": 2,    // Cisco
		"0030b7": 5013, // Teletrol
		"0030b8": 5014, // RiverDelta
		"0030b9": 5015, // Ectel
		"0030bb": 5016, // CacheFlow
		"0030bc": 5017, // Optronic
		"0030be": 5018, // City-Net
		"0030bf": 5019, // Multidata
		"0030c0": 5020, // Lara
		"0030c1": 302,  // HP
		"0030c2": 5021, // Comone
		"0030c6": 1999, // Control
		"0030c7": 5022, // Macromate
		"0030c9": 5023, // LuxN
		"0030cc": 5024, // Tenor
		"0030cd": 5025, // Conexant
		"0030ce": 5026, // Zaffire
		"0030cf": 5027, // TWO
		"0030d0": 5028, // Tellabs
		"0030d1": 5029, // Inova
		"0030d2": 201,  // WIN
		"0030d3": 653,  // Agilent
		"0030d4": 5030, // AAE
		"0030d5": 5031, // DResearch
		"0030d7": 428,  // Innovative
		"0030d8": 5032, // Sitek
		"0030d9": 5033, // Datacore
		"0030da": 3870, // Comtrend
		"0030db": 5034, // Mindready
		"0030dc": 5035, // Rightech
		"0030dd": 5036, // Indigita
		"0030e2": 5037, // Garnet
		"0030e3": 5038, // Sedona
		"0030e8": 5039, // Ensim
		"0030ea": 5040, // TeraForce
		"0030eb": 5041, // Turbonet
		"0030ec": 5042, // Borgardt
		"0030ee": 5043, // DSG
		"0030ef": 5044, // Neon
		"0030f0": 5045, // Uniform
		"0030f1": 145,  // Accton
		"0030f2": 2,    // Cisco
		"0030f4": 5046, // Stardot
		"0030f6": 5047, // Securelogix
		"0030f7": 5048, // Ramix
		"0030f8": 5049, // Dynapro
		"0030f9": 5050, // Sollae
		"0030fa": 5051, // Telica
		"0030fb": 5052, // AZS
		"0030fc": 5053, // Terawave
		"0030fe": 5054, // DSA
		"0030ff": 4795, // DataFab
		"003146": 826,  // Juniper
		"003192": 1595, // TP-Link
		"003217": 2,    // Cisco
		"00323a": 5055, // so-logic
		"00336c": 5056, // SynapSense
		"0034da": 869,  // LG
		"0034f1": 5057, // Radicom
		"0034fe": 3335, // Huawei
		"00351a": 2,    // Cisco
		"003532": 5058, // Electro-Metrics
		"003676": 125,  // ARRIS
		"0036fe": 2835, // SuperVision
		"00376d": 2057, // Murata
		"0037b7": 2049, // Sagemcom
		"0038df": 2,    // Cisco
		"003a7d": 2,    // Cisco
		"003a98": 2,    // Cisco
		"003a99": 2,    // Cisco
		"003a9a": 2,    // Cisco
		"003a9b": 2,    // Cisco
		"003a9c": 2,    // Cisco
		"003aaf": 5059, // BlueBit
		"003c10": 2,    // Cisco
		"003c84": 1655, // Silicon
		"003cc5": 5060, // WONWOO
		"003d41": 5061, // Hatteland
		"003de1": 3335, // Huawei
		"003de8": 869,  // LG
		"003ee1": 545,  // Apple
		"004002": 5062, // Perle
		"004004": 5063, // ICM
		"004005": 5064, // ANI
		"004006": 2971, // Sampo
		"00400a": 5065, // Pivotal
		"00400b": 2,    // Cisco
		"00400e": 4243, // Memotec
		"00400f": 5066, // Datacom
		"004010": 5067, // Sonic
		"004012": 5068, // Windata
		"004014": 5069, // Comsoft
		"004018": 5070, // Adobe
		"004019": 5071, // Aeon
		"00401b": 5072, // Printer
		"00401c": 5073, // AST
		"00401d": 5074, // Invisible
		"00401e": 5075, // ICC
		"00401f": 5076, // Colorgraph
		"004020": 5077, // CommScope
		"004022": 5078, // Klever
		"004023": 5079, // Logic
		"004024": 5080, // Compac
		"004026": 1077, // Buffalo
		"004028": 5081, // Netcomm
		"004029": 5082, // Compex
		"00402b": 5083, // Trigem
		"00402e": 263,  // Precision
		"004030": 5084, // GK
		"004032": 1565, // Digital
		"004033": 5085, // Addtron
		"004034": 5086, // Bustek
		"004035": 5087, // Opcom
		"004036": 5088, // Minim
		"004037": 5089, // SEA-Ilan
		"00403a": 5090, // Impact
		"00403b": 5091, // Synerjet
		"00403c": 5092, // Forks
		"00403d": 5093, // Teradata
		"00403f": 5094, // Ssangyong
		"004041": 1698, // Fujikura
		"004042": 5095, // N.A.T
		"004044": 5096, // Qnix
		"004045": 184,  // Twinhead
		"004046": 5097, // UDC
		"00404b": 1493, // Maple
		"00404c": 5098, // Hypertec
		"00404e": 5099, // Fluent
		"004050": 5100, // Ironicsorporated
		"004052": 5101, // Star
		"004053": 5102, // Ampro
		"004055": 5103, // Metronix
		"004058": 5104, // UKG
		"00405b": 5105, // Funasset
		"00405c": 5106, // Future
		"00405d": 5107, // Star-TEK
		"00405f": 5108, // AFE
		"004060": 5109, // Comendec
		"004061": 5110, // Datatech
		"004062": 5111, // E-Systems./Garland
		"004063": 4163, // VIA
		"004066": 5112, // APRESIA
		"004067": 5113, // Omnibyte
		"004068": 5114, // Extended
		"004069": 5115, // Lemcom
		"00406b": 5116, // Sysgen
		"00406c": 5117, // Copernique
		"00406d": 5118, // Lanco
		"00406e": 5119, // Corollary
		"00406f": 5120, // Sync
		"004070": 5121, // Interware
		"004071": 5122, // ATM
		"004077": 5123, // Maxton
		"004079": 5124, // Juko
		"00407c": 5125, // Qume
		"00407d": 5126, // Extension
		"00407e": 4226, // Evergreen
		"00407f": 3724, // FLIR
		"004080": 5127, // Athenix
		"004084": 934,  // Honeywell
		"004087": 5128, // Ubitrex
		"004088": 5129, // Mobius
		"004089": 5130, // Meidensha
		"00408b": 5131, // Raylan
		"00408c": 5132, // Axis
		"004090": 5133, // Ansel
		"004092": 2872, // ASP
		"004093": 5134, // Paxdata
		"004094": 5135, // Shographics
		"004096": 2,    // Cisco
		"004098": 5136, // Dressler
		"004099": 5137, // Newgen
		"00409b": 5138, // HAL
		"00409c": 5139, // Transware
		"00409d": 5140, // DigiBoard
		"00409e": 5141, // Concurrent
		"00409f": 18,   // Telco
		"0040a0": 5142, // Goldstar
		"0040a2": 5143, // Kingstar
		"0040a3": 5144, // Microunity
		"0040a4": 5145, // Rose
		"0040a6": 68,   // Cray
		"0040a8": 5146, // IMF
		"0040a9": 5066, // Datacom
		"0040af": 1565, // Digital
		"0040b0": 5147, // Bytex
		"0040b1": 5148, // Codonics
		"0040b2": 5149, // Systemforschung
		"0040b3": 5150, // ParTech
		"0040b5": 5151, // Video
		"0040b6": 5152, // Computerm
		"0040b7": 5153, // Stealth
		"0040b9": 5154, // Macq
		"0040ba": 2653, // Alliant
		"0040bc": 5155, // Algorithmics
		"0040bd": 5156, // Starlight
		"0040be": 1041, // Boeing
		"0040c5": 5157, // Micom
		"0040c6": 5158, // Fibernet
		"0040c7": 4629, // Ruby
		"0040c8": 5159, // Milan
		"0040c9": 5160, // Ncube
		"0040cb": 5161, // Lanwan
		"0040ce": 5162, // NET-Source
		"0040d0": 513,  // Mitac
		"0040d2": 5163, // Pagine
		"0040d3": 5164, // Kimpsion
		"0040da": 5165, // Telspec
		"0040dc": 5166, // Tritec
		"0040dd": 5167, // Hong
		"0040df": 5168, // Digalog
		"0040e0": 5169, // Atomwide
		"0040e1": 5170, // Marner
		"0040e3": 5171, // Quin
		"0040e4": 5172, // E-M
		"0040e5": 5173, // Sybus
		"0040e6": 5174, // C.A.E.N
		"0040e9": 5175, // Accord
		"0040ee": 5176, // Optimem
		"0040ef": 5177, // Hypercom
		"0040f0": 5178, // MicroBrain
		"0040f1": 5179, // Chuo
		"0040f3": 5180, // Netcor
		"0040f4": 3384, // Cameo
		"0040f6": 5181, // Katron
		"0040f7": 656,  // Polaroid
		"0040f9": 5182, // Combinet
		"0040fa": 5183, // Microboards
		"0040fb": 5184, // Cascade
		"0040fd": 5185, // LXE
		"0040fe": 5186, // Symplex
		"0040ff": 5187, // Telebit
		"0041d2": 2,    // Cisco
		"004238": 421,  // Intel
		"004252": 5188, // RLX
		"00425a": 2,    // Cisco
		"004268": 2,    // Cisco
		"004279": 3940, // Sunitec
		"00451d": 2,    // Cisco
		"0045e2": 189,  // CyberTAN
		"00464b": 3335, // Huawei
		"004a77": 3032, // ZTE
		"004e01": 102,  // Dell
		"004e35": 302,  // HP
		"004f1a": 3335, // Huawei
		"005000": 5189, // Nexo
		"005001": 5190, // Yamashita
		"005002": 5191, // Omnisec
		"005003": 5192, // Xrite
		"005004": 160,  // 3Com
		"005006": 2380, // TAC
		"005007": 300,  // Siemens
		"00500a": 5193, // Iris
		"00500b": 2,    // Cisco
		"00500c": 5194, // e-Tek
		"00500e": 5195, // Chromatis
		"00500f": 2,    // Cisco
		"005012": 5196, // CBL
		"005014": 2,    // Cisco
		"005018": 5197, // AMIT
		"00501a": 5198, // IQinVision
		"00501c": 5199, // Jatom
		"00501f": 5200, // MRG
		"005020": 5201, // Mediastar
		"005021": 5202, // EIS
		"005022": 5203, // Zonet
		"005024": 5204, // Navic
		"005026": 5205, // Cosystems
		"005027": 5206, // Genicom
		"005028": 5207, // Aval
		"00502a": 2,    // Cisco
		"00502b": 5208, // Genrad
		"00502c": 5209, // Soyo
		"00502d": 5210, // Accel
		"00502e": 5211, // Cambex
		"00502f": 5212, // TollBridge
		"005031": 5213, // Aeroflex
		"005032": 5214, // Picazo
		"005033": 5215, // Mayan
		"005036": 5216, // Netcam
		"005037": 5217, // Koga
		"005039": 5218, // Mariner
		"00503a": 5219, // Datong
		"00503b": 5220, // Mediafire
		"00503e": 2,    // Cisco
		"005040": 2154, // Panasonic
		"005041": 5221, // Coretronic
		"005044": 5222, // Asaca
		"005045": 5223, // Rioworks
		"005046": 5224, // Menicx
		"005047": 67,   // Private
		"005048": 5225, // Infolibria
		"005049": 859,  // Arbor
		"00504d": 1205, // Tokyo
		"00504f": 5226, // Olencom
		"005050": 2,    // Cisco
		"005052": 5227, // Tiara
		"005053": 2,    // Cisco
		"005054": 2,    // Cisco
		"005055": 5228, // Doms
		"005056": 809,  // VMware
		"005058": 5229, // Sangoma
		"005059": 5230, // iBAHN
		"00505c": 5231, // Tundo
		"005062": 5232, // Kouwell
		"005064": 5233, // CAE
		"005065": 3512, // TDK-Lambda
		"005067": 5234, // Aerocomm
		"005068": 3691, // Electronic
		"005069": 5235, // PixStream
		"00506a": 5236, // Edeva
		"00506b": 5237, // SPX-Ateg
		"00506c": 5238, // Beijer
		"00506d": 5239, // Videojet
		"00506e": 5240, // Corder
		"00506f": 5241, // G-Connect
		"005070": 5242, // Chaintech
		"005071": 5243, // Aiwa
		"005072": 5244, // Corvis
		"005073": 2,    // Cisco
		"005075": 5245, // Kestrel
		"005076": 372,  // IBM
		"005077": 5246, // Prolific
		"005079": 67,   // Private
		"00507a": 486,  // Xpeed
		"00507b": 5247, // Merlot
		"00507c": 5248, // Videocon
		"00507d": 5249, // IFP
		"00507e": 5250, // Newer
		"00507f": 3923, // DrayTek
		"005080": 2,    // Cisco
		"005082": 5251, // Foresson
		"005083": 5252, // Gilbarco
		"005084": 2084, // Quantum
		"005086": 5253, // Telkom
		"005088": 5254, // Amano
		"00508b": 302,  // HP
		"00508c": 5255, // RSI
		"00508d": 5256, // Abit
		"00508e": 5257, // Optimation
		"005090": 5258, // Dctri
		"005091": 5259, // Netaccess
		"005093": 1041, // Boeing
		"005094": 125,  // ARRIS
		"005095": 5260, // Peracom
		"005096": 5261, // Salix
		"005098": 5262, // Globaloop
		"005099": 160,  // 3Com
		"00509a": 5263, // TAG
		"00509b": 5264, // Switchcore
		"00509c": 5265, // Beta
		"00509f": 5266, // Horizon
		"0050a0": 957,  // Delta
		"0050a2": 2,    // Cisco
		"0050a3": 5267, // TransMedia
		"0050a4": 5268, // IO
		"0050a6": 5269, // Optronics
		"0050a7": 2,    // Cisco
		"0050a8": 5270, // OpenCon
		"0050ab": 5271, // NALTEC
		"0050ac": 1493, // Maple
		"0050ae": 5272, // FDK
		"0050af": 5273, // Intergon
		"0050b2": 5274, // BRODEL
		"0050b3": 5275, // Voiceboard
		"0050b7": 5276, // Boser
		"0050b8": 5029, // Inova
		"0050b9": 5277, // Xitron
		"0050ba": 803,  // D-Link
		"0050bb": 2364, // CMS
		"0050bd": 2,    // Cisco
		"0050bf": 1709, // Metalligence
		"0050c0": 5278, // Gatan
		"0050c1": 5279, // Gemflex
		"0050c4": 5280, // IMD
		"0050c5": 5281, // ADS
		"0050c7": 67,   // Private
		"0050c8": 5282, // Addonics
		"0050cb": 5283, // Jetter
		"0050cd": 5284, // Digianswer
		"0050ce": 869,  // LG
		"0050d0": 5285, // Minerva
		"0050d1": 2,    // Cisco
		"0050d2": 5286, // CMC
		"0050d5": 5287, // AD
		"0050d7": 5288, // Telstrat
		"0050d8": 5289, // Unicorn
		"0050da": 160,  // 3Com
		"0050de": 5290, // Signum
		"0050df": 5291, // AirFiber
		"0050e1": 5292, // NS
		"0050e2": 2,    // Cisco
		"0050e3": 125,  // ARRIS
		"0050e4": 545,  // Apple
		"0050e6": 5293, // Hakusan
		"0050e8": 5294, // Nomadix
		"0050ea": 5295, // XEL
		"0050eb": 5296, // Alpha-TOP
		"0050ec": 5297, // Olicom
		"0050ed": 5298, // Anda
		"0050f0": 2,    // Cisco
		"0050f1": 5299, // Maxlinear
		"0050f2": 612,  // Microsoft
		"0050f4": 5300, // Sigmatek
		"0050f6": 5301, // PAN-International
		"0050f8": 5302, // Entrega
		"0050f9": 5303, // Sensormatic
		"0050fa": 5304, // Oxtel
		"0050fb": 5305, // VSK
		"0050fc": 115,  // Edimax
		"0050fd": 5306, // Visioncomm
		"0050ff": 5307, // Hakko
		"0051ed": 869,  // LG
		"00549f": 620,  // Avaya
		"0054bd": 5308, // Swelaser
		"00562b": 2,    // Cisco
		"0056cd": 545,  // Apple
		"0057c1": 869,  // LG
		"0057d2": 2,    // Cisco
		"005907": 2665, // Lenovo
		"0059dc": 2,    // Cisco
		"005a13": 3335, // Huawei
		"005b94": 545,  // Apple
		"005d03": 1498, // Xilinx
		"005d73": 2,    // Cisco
		"005f67": 1595, // TP-Link
		"005f86": 2,    // Cisco
		"005fbf": 35,   // Toshiba
		"006000": 5309, // Xycom
		"006001": 5310, // InnoSys
		"006006": 5311, // Sotec
		"006008": 160,  // 3Com
		"006009": 2,    // Cisco
		"00600a": 5312, // Sord
		"00600b": 5313, // LOGWARE
		"00600c": 5314, // Eurotech
		"00600e": 5315, // Wavenet
		"00600f": 2274, // Westell
		"006014": 5316, // Edec
		"006015": 5317, // NET2NET
		"006016": 5318, // Clariion
		"006017": 5319, // Tokimec
		"00601b": 5320, // Mesa
		"00601c": 5321, // Telxon
		"00601d": 827,  // Lucent
		"00601e": 5322, // Softlab
		"00601f": 5323, // Stallion
		"006021": 5324, // DSC
		"006022": 5325, // Vicom
		"006024": 5326, // Gradient
		"006028": 5327, // Macrovision
		"00602a": 5328, // Symicron
		"00602d": 5329, // Alerton
		"00602e": 5330, // Cyclades
		"00602f": 2,    // Cisco
		"006031": 5331, // HRK
		"006032": 5332, // I-Cube
		"006038": 76,   // Nortel
		"006039": 5333, // SanCom
		"00603b": 5334, // AMTEC
		"00603c": 5335, // Hagiwara
		"00603d": 5336, // 3CX
		"00603e": 2,    // Cisco
		"006040": 5337, // Netro
		"006042": 5338, // TKS
		"006043": 5339, // iDirect
		"006044": 5340, // Litton/Poly-Scientific
		"006045": 5341, // Pathlight
		"006046": 5342, // Vmetro
		"006047": 2,    // Cisco
		"006048": 102,  // Dell
		"006049": 5343, // Vina
		"00604b": 5344, // Safe-Com
		"00604c": 2049, // Sagemcom
		"00604d": 4933, // MMC
		"00604e": 5345, // Cycle
		"006050": 5346, // Internix
		"006052": 5347, // PERIPHERALS
		"006054": 5348, // Controlware
		"006057": 2057, // Murata
		"006059": 2925, // Technical
		"00605a": 5349, // Celcore
		"00605b": 5350, // IntraServer
		"00605c": 2,    // Cisco
		"00605d": 5351, // Scanivalve
		"006061": 5352, // Whistle
		"006062": 5353, // Telesync
		"006064": 5081, // Netcomm
		"006068": 5354, // Dialogic
		"006069": 90,   // Brocade
		"00606b": 5355, // Synclayer
		"00606c": 879,  // Arescom
		"006070": 2,    // Cisco
		"006073": 5356, // Redcreek
		"006074": 5357, // QSC
		"006075": 5358, // Pentek
		"006077": 5359, // Prisa
		"00607a": 5360, // DVS
		"00607b": 5361, // Fore
		"00607c": 5362, // WaveAccess
		"00607d": 5363, // Sentient
		"00607e": 5364, // Gigalabs
		"00607f": 1119, // Aurora
		"006081": 5365, // TV/CoM
		"006082": 5366, // Novalink
		"006083": 2,    // Cisco
		"006089": 5367, // Xata
		"00608a": 5368, // Citadel
		"00608b": 5369, // ConferTech
		"00608c": 160,  // 3Com
		"00608d": 5370, // Unipulse
		"00608e": 5371, // HE
		"00608f": 5372, // Tekram
		"006090": 5373, // Artiza
		"006092": 5374, // Micro/Sys
		"006093": 5375, // Varian
		"006094": 372,  // IBM
		"006095": 5376, // Accu-Time
		"006097": 160,  // 3Com
		"006098": 5377, // HT
		"006099": 135,  // SBE
		"00609b": 5378, // AstroNova
		"00609c": 5379, // Perkin-Elmer
		"00609f": 5380, // Phast
		"0060a1": 5381, // VPNet
		"0060a3": 5382, // Continuum
		"0060a4": 5383, // GEW
		"0060a7": 5384, // MICROSENS
		"0060a8": 5385, // Tidomat
		"0060ab": 4270, // Larscom
		"0060ac": 5386, // Resilience
		"0060ad": 5387, // MegaChips
		"0060b0": 302,  // HP
		"0060b1": 5388, // Input/Output
		"0060b3": 5389, // Z-CoM
		"0060b5": 5390, // KEBA
		"0060b6": 5391, // Land
		"0060b7": 5392, // Channelmatic
		"0060b8": 5393, // CoRELIS
		"0060ba": 5394, // Sahara
		"0060bb": 16,   // Cabletron
		"0060bc": 5395, // KeunYoung
		"0060bd": 5396, // Enginuity
		"0060be": 5397, // Webtronics
		"0060bf": 5398, // Macraigor
		"0060c0": 5399, // Nera
		"0060c1": 5400, // WaveSpan
		"0060c2": 5401, // MPL
		"0060c3": 5402, // Netvision
		"0060c5": 5403, // Ancot
		"0060c6": 5404, // DCS
		"0060c7": 5405, // Amati
		"0060c9": 1856, // ControlNet
		"0060ca": 1526, // Harmonic
		"0060cc": 5406, // Emtrakorporated
		"0060cd": 5407, // VideoServer
		"0060ce": 5408, // Acclaim
		"0060cf": 5409, // Alteon
		"0060d0": 5410, // Snmp
		"0060d1": 5184, // Cascade
		"0060d3": 1057, // AT&T
		"0060d4": 5411, // Eldat
		"0060d6": 4340, // NovAtel
		"0060d8": 5412, // Elmic
		"0060d9": 5413, // Transys
		"0060dd": 5414, // Myricom
		"0060de": 5415, // Kayser-Threde
		"0060df": 90,   // Brocade
		"0060e0": 5416, // Axiom
		"0060e1": 5417, // Orckit
		"0060e2": 2508, // Quest
		"0060e4": 5418, // Compuserve
		"0060e6": 5419, // Shomiti
		"0060e7": 5420, // Randata
		"0060e8": 89,   // Hitachi
		"0060e9": 5421, // Atop
		"0060ea": 5422, // StreamLogic
		"0060eb": 5423, // Fourthtrack
		"0060ee": 5424, // Apollo
		"0060ef": 5425, // Flytech
		"0060f1": 5426, // EXP
		"0060f2": 5427, // Lasergraphics
		"0060f4": 5428, // ADVANCED
		"0060f6": 5429, // Nextest
		"0060f7": 5430, // Datafusion
		"0060f8": 5431, // Loran
		"0060fb": 5432, // Packeteer
		"0060fd": 5433, // NetICs
		"0060ff": 5434, // QuVis
		"006151": 3335, // Huawei
		"006171": 545,  // Apple
		"00620b": 858,  // Broadcom
		"0062ec": 2,    // Cisco
		"0063de": 5435, // Cloudwalk
		"006440": 2,    // Cisco
		"0064af": 1238, // Dish
		"006619": 3335, // Huawei
		"00664b": 3335, // Huawei
		"006762": 4922, // Fiberhome
		"00682b": 3335, // Huawei
		"0068eb": 302,  // HP
		"00692d": 5436, // Sunnovo
		"006b6f": 3335, // Huawei
		"006b9e": 3468, // Vizio
		"006bf1": 2,    // Cisco
		"006cbc": 2,    // Cisco
		"006d52": 545,  // Apple
		"006dfb": 5437, // Vutrix
		"006e02": 5438, // Xovis
		"006f64": 152,  // Samsung
		"007147": 5439, // Amazon
		"0071c2": 5440, // Pegatron
		"007204": 152,  // Samsung
		"007263": 2363, // Netcore
		"007278": 2,    // Cisco
		"0073e0": 152,  // Samsung
		"00749c": 5441, // Ruijie
		"007532": 5442, // Inid
		"0075e1": 5443, // Ampt
		"00763d": 5444, // Veea
		"007686": 2,    // Cisco
		"00778d": 2,    // Cisco
		"007888": 2,    // Cisco
		"00789e": 2049, // Sagemcom
		"007b18": 5445, // SENTRY
		"007c2d": 152,  // Samsung
		"007d60": 545,  // Apple
		"007e95": 2,    // Cisco
		"007f28": 2247, // Actiontec
		"008000": 1183, // Multitech
		"008001": 5446, // Periphonics
		"008002": 5447, // Satelcom
		"008003": 5448, // Hytec
		"008004": 5449, // Antlow
		"008005": 5450, // Cactus
		"008006": 5451, // Compuadd
		"008008": 5452, // Dynatech
		"008009": 5453, // Jupiter
		"00800a": 5454, // Japan
		"00800b": 5455, // CSK
		"00800c": 5456, // Videcom
		"00800e": 5457, // Atlantix
		"008010": 5458, // Commodore
		"008013": 5459, // Thomas-Conrad
		"008014": 5460, // Esprit
		"008015": 5461, // Seiko
		"008017": 5462, // PFU
		"008019": 5463, // Dayna
		"00801b": 5464, // Kodiak
		"00801e": 5465, // Xinetron
		"008020": 109,  // Network
		"008022": 5466, // Scan-Optics
		"008024": 5467, // Kalpana
		"008026": 109,  // Network
		"008027": 5468, // Adaptive
		"008028": 5469, // Tradpost
		"008029": 4268, // Eagle
		"00802d": 5470, // Xylogics
		"008030": 2241, // Nexus
		"008031": 5471, // Basys
		"008032": 5472, // Access
		"008037": 306,  // Ericsson
		"00803a": 5473, // Varityper
		"00803b": 5474, // APT
		"00803c": 5475, // TVS
		"00803d": 5476, // Surigiken
		"00803e": 5477, // Synernetics
		"00803f": 5478, // Tatung
		"008043": 5479, // Networld
		"008044": 5480, // Systech
		"008047": 5481, // IN-NET
		"008048": 5082, // Compex
		"00804a": 5482, // PRO-LOG
		"00804b": 4268, // Eagle
		"00804c": 5483, // Contec
		"00804e": 406,  // Apex
		"00804f": 5484, // Daikin
		"008050": 5485, // Ziatech
		"008051": 5486, // Fibermux
		"008053": 5487, // Intellicom
		"008054": 5488, // Frontier
		"008055": 5489, // Fermilab
		"008056": 5490, // SPHINX
		"008057": 5491, // Adsoft
		"008058": 5072, // Printer
		"00805b": 5492, // Condor
		"00805c": 5493, // Agilis
		"00805d": 5494, // Canstar
		"00805f": 302,  // HP
		"008061": 5495, // Litton
		"008062": 1410, // Interface
		"008064": 5496, // Wyse
		"008065": 5497, // Cybergraphic
		"008069": 5498, // Computone
		"00806a": 5499, // ERI
		"00806d": 5500, // Century
		"00806f": 5501, // Onelan
		"008071": 5502, // SAI
		"008072": 5503, // Microplex
		"008075": 5504, // Parsytec
		"008076": 5505, // Mcnc
		"008077": 3231, // Brother
		"00807a": 5506, // Aitech
		"00807b": 5507, // Artel
		"00807c": 5508, // Fibercom
		"00807d": 5509, // Equinox
		"00807f": 5510, // DY-4
		"008080": 5511, // Datamedia
		"008083": 5512, // Amdahl
		"008085": 5513, // H-Three
		"008089": 5514, // Tecnetics
		"00808b": 5515, // Dacoll
		"00808c": 5516, // Netscout
		"00808e": 5517, // Radstone
		"008090": 5518, // Microtek
		"008092": 5519, // Silex
		"008093": 5520, // Xyron
		"008098": 5521, // TDK
		"008099": 1855, // Eaton
		"00809a": 5522, // Novus
		"00809b": 5523, // Justsystem
		"00809c": 5524, // Luxcom
		"00809d": 5525, // Commscraft
		"00809e": 5526, // Datus
		"00809f": 5527, // ALE
		"0080a0": 302,  // HP
		"0080a1": 5528, // Microtest
		"0080a2": 357,  // Creative
		"0080a3": 5529, // Lantronix
		"0080a4": 5530, // Liberty
		"0080a5": 5531, // Speed
		"0080a6": 5532, // Republic
		"0080a7": 934,  // Honeywell
		"0080a8": 5533, // Vitacom
		"0080a9": 5534, // Clearpoint
		"0080aa": 5535, // Maxpeed
		"0080ac": 5536, // Imlogix
		"0080ad": 5537, // Cnet
		"0080af": 5538, // Allumer
		"0080b1": 5539, // Softcom
		"0080b4": 5540, // Sophia
		"0080b5": 1826, // United
		"0080b7": 5541, // Stellar
		"0080ba": 5542, // Specialix
		"0080bc": 89,   // Hitachi
		"0080be": 5543, // Aries
		"0080c1": 5544, // Lanex
		"0080c4": 5545, // Document
		"0080c7": 788,  // Xircom
		"0080c8": 803,  // D-Link
		"0080ca": 830,  // Netcom
		"0080cd": 5546, // Micronics
		"0080d1": 5547, // Kimtron
		"0080d2": 5548, // Shinnihondenko
		"0080d3": 5549, // Shiva
		"0080d4": 5550, // Chase
		"0080d5": 5551, // Cadre
		"0080d6": 5552, // Nuvotech
		"0080d7": 5553, // Fantum
		"0080db": 5554, // Graphon
		"0080dc": 5555, // Picker
		"0080dd": 5556, // GMX/Gimix
		"0080e0": 5557, // XTP
		"0080e2": 5558, // T.D.I
		"0080e5": 5559, // NetApp
		"0080e6": 5560, // Peer
		"0080e9": 70,   // Madge
		"0080ec": 5561, // Supercomputing
		"0080ed": 5562, // IQ
		"0080ef": 5563, // Rational
		"0080f0": 2154, // Panasonic
		"0080f1": 5564, // Opus
		"0080f2": 3477, // Raycom
		"0080f3": 5565, // SUN
		"0080f5": 5566, // Quantel
		"0080f7": 1800, // Zenith
		"0080f8": 5567, // Mizar
		"0080f9": 5568, // Heurikon
		"0080fa": 5569, // RWT
		"0080fb": 5570, // BVM
		"0080fc": 5571, // Avatar
		"0080fe": 5572, // Azure
		"0081c4": 2,    // Cisco
		"0084ed": 67,   // Private
		"0086a0": 67,   // Private
		"008701": 152,  // Samsung
		"008731": 2,    // Cisco
		"008764": 2,    // Cisco
		"008865": 545,  // Apple
		"0088ba": 5573, // NC&C
		"008a55": 3335, // Huawei
		"008a76": 545,  // Apple
		"008a96": 2,    // Cisco
		"008b43": 5574, // Rftech
		"008bfc": 5575, // mixi
		"008cfa": 3979, // Inventec
		"008e73": 2,    // Cisco
		"008ef2": 1368, // Netgear
		"009001": 5576, // Nishimu
		"009002": 5577, // Allgon
		"009003": 5578, // Aplio
		"009004": 160,  // 3Com
		"009005": 5579, // Protech
		"009007": 5580, // Domex
		"009008": 5581, // HanA
		"00900a": 5582, // Proton
		"00900b": 5583, // Lanner
		"00900c": 2,    // Cisco
		"00900e": 5584, // Handlink
		"009010": 5585, // Simulation
		"009011": 5586, // WAVTrace
		"009013": 5587, // Samsan
		"009015": 5588, // Centigram
		"009016": 5589, // ZAC
		"009017": 5590, // Zypcom
		"009019": 5591, // Hermes
		"00901a": 5592, // Unisphere
		"00901c": 5593, // mps
		"00901d": 5594, // PEC
		"009021": 2,    // Cisco
		"009022": 5595, // Ivex
		"009023": 5596, // Zilog
		"009024": 5597, // Pipelinks
		"009027": 421,  // Intel
		"009029": 5598, // Crypto
		"00902b": 2,    // Cisco
		"00902e": 5599, // Namco
		"00902f": 2363, // Netcore
		"009030": 5600, // Honeywell-Dating
		"009031": 5601, // Mysticom
		"009033": 5602, // Innovaphone
		"009034": 5603, // Imagic
		"009036": 5604, // ens
		"009037": 5605, // Acucomm
		"009038": 5606, // Fountain
		"009039": 5607, // Shasta
		"00903d": 5608, // Biopac
		"00903f": 5609, // WorldCast
		"009040": 300,  // Siemens
		"009042": 5610, // ECCS
		"009045": 31,   // Marconi
		"009046": 5611, // Dexdyne
		"009048": 5612, // Zeal
		"009049": 5613, // Entridia
		"00904b": 1450, // Gemtek
		"00904c": 5614, // Epigram
		"00904e": 5615, // Delem
		"009050": 5616, // Teleste
		"009051": 5617, // Ultimate
		"009053": 5618, // Daewoo
		"009056": 5619, // Telestream
		"009057": 5620, // AANetcom
		"009058": 3430, // Ultra
		"00905c": 5621, // Edmi
		"00905e": 5622, // Rauland-Borg
		"00905f": 2,    // Cisco
		"009061": 811,  // Pacific
		"009063": 3260, // Coherent
		"009064": 1142, // Thomson
		"009065": 5623, // Finisar
		"009066": 5624, // Troika
		"009067": 5625, // WalkAbout
		"009068": 5626, // DVT
		"009069": 826,  // Juniper
		"00906a": 5627, // Turnstone
		"00906d": 2,    // Cisco
		"00906e": 5628, // Praxon
		"00906f": 2,    // Cisco
		"009070": 5629, // NEO
		"009072": 5630, // Simrad
		"009073": 5631, // Gaio
		"009074": 5632, // Argon
		"009079": 5633, // ClearOne
		"00907a": 5634, // Spectralink
		"00907b": 5635, // E-Tech
		"00907c": 5636, // Digitalcast
		"00907d": 5637, // Lake
		"00907e": 5638, // Vetronix
		"00907f": 175,  // WatchGuard
		"009080": 5639, // NOT
		"009081": 5640, // Aloha
		"009083": 1896, // Turbo
		"009085": 1120, // Golden
		"009086": 2,    // Cisco
		"009087": 5641, // Itis
		"00908a": 5642, // Bayly
		"00908c": 5643, // Etrend
		"00908d": 5644, // Vickers
		"009090": 5645, // I-BUS
		"009091": 5646, // DigitalScape
		"009092": 2,    // Cisco
		"009093": 5647, // EIZO
		"009094": 5648, // Osprey
		"009096": 2545, // Askey
		"009097": 4215, // Sycamore
		"00909b": 5649, // Markem-Imaje
		"00909c": 125,  // ARRIS
		"00909f": 5650, // Digi-Data
		"0090a0": 5651, // 8X8
		"0090a2": 189,  // CyberTAN
		"0090a3": 5652, // Corecess
		"0090a4": 5653, // Altiga
		"0090a6": 2,    // Cisco
		"0090a7": 5654, // Clientec
		"0090a8": 5655, // NineTiles
		"0090a9": 123,  // WD
		"0090ab": 2,    // Cisco
		"0090ac": 5656, // Optivision
		"0090ad": 5657, // Aspect
		"0090b0": 5658, // Vadem
		"0090b1": 2,    // Cisco
		"0090b2": 5659, // Avici
		"0090b3": 5660, // Agranat
		"0090b4": 5661, // Willowbrook
		"0090b5": 5662, // Nikon
		"0090b6": 5663, // Fibex
		"0090ba": 5664, // Valid
		"0090bc": 5665, // Telemann
		"0090bd": 5666, // Omnia
		"0090bf": 2,    // Cisco
		"0090c4": 5667, // Javelin
		"0090c6": 5668, // Optim
		"0090c7": 5669, // Icom
		"0090c8": 5670, // Waverider
		"0090c9": 5671, // DPAC
		"0090cc": 4479, // Planex
		"0090cf": 76,   // Nortel
		"0090d1": 5672, // Leichu
		"0090d5": 5673, // Euphonix
		"0090d7": 5674, // NetBoost
		"0090d8": 5675, // Whitecross
		"0090d9": 2,    // Cisco
		"0090da": 5676, // Dynarc
		"0090dd": 5677, // MIHARU
		"0090de": 5678, // Cardkey
		"0090e0": 5679, // Systran
		"0090e3": 5680, // Avex
		"0090e5": 5681, // Teknema
		"0090e6": 5682, // ALi
		"0090e8": 5683, // Moxa
		"0090e9": 5684, // Janz
		"0090ea": 2237, // Alpha
		"0090ec": 5685, // Pyrescom
		"0090ee": 5686, // Personal
		"0090ef": 5687, // Integrix
		"0090f2": 2,    // Cisco
		"0090f3": 5657, // Aspect
		"0090f5": 5688, // Clevo
		"0090f6": 5689, // Escalate
		"0090f7": 5690, // Nbase
		"0090f9": 5691, // Imagine
		"0090fa": 129,  // Emulex
		"0090fb": 5692, // Portwell
		"0090fd": 5693, // CopperCom
		"0090fe": 5694, // Elecom
		"0090ff": 5695, // Tellus
		"00919e": 421,  // Intel
		"0091eb": 5696, // Renesas
		"009337": 421,  // Intel
		"009363": 5697, // Uni-Link
		"0094a1": 289,  // F5
		"0094ec": 3335, // Huawei
		"00991d": 3335, // Huawei
		"009acd": 3335, // Huawei
		"009ad2": 2,    // Cisco
		"009c02": 302,  // HP
		"009d6b": 2057, // Murata
		"009e1e": 2,    // Cisco
		"009ec8": 5698, // Xiaomi
		"00a000": 5699, // Centillion
		"00a003": 300,  // Siemens
		"00a004": 1117, // Netpower
		"00a007": 5700, // Apexx
		"00a008": 5701, // Netcorp
		"00a00a": 259,  // Airspan
		"00a00b": 5702, // Computex
		"00a00c": 5703, // Kingmax
		"00a00e": 5516, // Netscout
		"00a00f": 2423, // Broadband
		"00a011": 5704, // Mutoh
		"00a012": 18,   // Telco
		"00a013": 5705, // Teltrend
		"00a014": 5706, // Csir
		"00a015": 849,  // Wyle
		"00a016": 5707, // Micropolis
		"00a01b": 5708, // Premisys
		"00a01c": 5709, // Nascent
		"00a01e": 5710, // EST
		"00a01f": 5711, // Tricord
		"00a020": 5712, // Citicorp/TTI
		"00a024": 160,  // 3Com
		"00a025": 5713, // Redcom
		"00a026": 5714, // Teldat
		"00a027": 5715, // Firepower
		"00a029": 5716, // Coulter
		"00a02a": 5717, // Trancell
		"00a02b": 5718, // Transitions
		"00a02c": 5719, // interWAVE
		"00a02e": 5720, // Brand
		"00a030": 5721, // Captor/SA
		"00a031": 5722, // Hazeltine
		"00a034": 5723, // Axel
		"00a035": 5724, // Cylink
		"00a038": 5725, // Email
		"00a039": 5726, // Ross
		"00a03a": 5727, // Kubotek
		"00a03d": 5728, // Opto-22
		"00a040": 545,  // Apple
		"00a041": 5729, // Inficon
		"00a042": 5730, // Spur
		"00a043": 2359, // American
		"00a046": 57,   // Scitex
		"00a048": 5731, // Questech
		"00a049": 5732, // Digitech
		"00a04e": 5733, // Voelker
		"00a04f": 5734, // Ameritec
		"00a051": 5735, // Angia
		"00a052": 5736, // Stanilite
		"00a054": 67,   // Private
		"00a056": 5737, // Micropross
		"00a057": 5738, // LANCOM
		"00a058": 5739, // Glory
		"00a05b": 5740, // Marquip
		"00a05f": 5741, // BTG
		"00a063": 5742, // JRL
		"00a064": 5743, // KVB/Analect
		"00a065": 1990, // Symantec
		"00a066": 5744, // ISA
		"00a067": 109,  // Network
		"00a068": 5745, // BHP
		"00a069": 5746, // Symmetricom
		"00a06a": 4213, // Verilink
		"00a06e": 5747, // Austron
		"00a070": 5748, // Coastcom
		"00a072": 5749, // Ovation
		"00a073": 5750, // CoM21
		"00a074": 5751, // Perception
		"00a075": 4628, // Micron
		"00a078": 31,   // Marconi
		"00a079": 5752, // Alps
		"00a07b": 5753, // Dawn
		"00a07d": 5754, // Seeq
		"00a07e": 5755, // Avid
		"00a07f": 5756, // GSM-Syntel
		"00a084": 5757, // Dataplex
		"00a085": 67,   // Private
		"00a087": 5758, // Microsemi
		"00a088": 5759, // Essential
		"00a089": 5760, // Xpoint
		"00a08a": 5761, // Brooktrout
		"00a08b": 5762, // Aston
		"00a08d": 5763, // Jacomo
		"00a08f": 5764, // Desknet
		"00a090": 5765, // TimeStep
		"00a091": 5766, // Applicom
		"00a094": 5767, // Comsat
		"00a095": 5768, // Acacia
		"00a098": 5559, // NetApp
		"00a099": 5769, // K-NET
		"00a09b": 30,   // Qpsx
		"00a09c": 5770, // Xyplex
		"00a09e": 5771, // Ictv
		"00a09f": 5772, // Commvision
		"00a0a4": 11,   // Oracle
		"00a0a6": 5773, // M.I
		"00a0a7": 5774, // Vorax
		"00a0a8": 5775, // Renex
		"00a0a9": 5776, // Navtel
		"00a0ad": 31,   // Marconi
		"00a0ae": 5777, // Nucom
		"00a0af": 5778, // WMS
		"00a0b3": 5779, // Zykronix
		"00a0b5": 5780, // 3H
		"00a0b7": 5781, // Cordant
		"00a0b8": 5559, // NetApp
		"00a0b9": 4268, // Eagle
		"00a0ba": 5782, // Patton
		"00a0bb": 5783, // Hilan
		"00a0bc": 5784, // Viasatorporated
		"00a0bd": 4776, // I-Tech
		"00a0c2": 5785, // R.A
		"00a0c3": 5786, // Unicomputer
		"00a0c4": 5787, // Cristie
		"00a0c5": 2693, // ZyXEL
		"00a0c6": 5788, // Qualcomm
		"00a0c8": 202,  // Adtran
		"00a0c9": 421,  // Intel
		"00a0cc": 5789, // Lite-ON
		"00a0ce": 5790, // Ecessa
		"00a0cf": 5791, // Sotas
		"00a0d1": 3979, // Inventec
		"00a0d3": 5792, // Instem
		"00a0d4": 5793, // Radiolan
		"00a0d6": 135,  // SBE
		"00a0d9": 5794, // Convex
		"00a0da": 5795, // INTEGRATED
		"00a0dc": 5796, // O.N
		"00a0dd": 5797, // Azonix
		"00a0de": 727,  // Yamaha
		"00a0df": 5798, // STS
		"00a0e0": 5799, // Tennyson
		"00a0e2": 5800, // Keisokugiken
		"00a0e3": 5801, // XKL
		"00a0e4": 5802, // Optiquest
		"00a0e5": 5803, // NHC
		"00a0e6": 5354, // Dialogic
		"00a0e8": 5804, // Reuters
		"00a0ea": 5805, // Ethercom
		"00a0eb": 3303, // Encore
		"00a0ec": 5806, // Transmitton
		"00a0ee": 5807, // Nashoba
		"00a0ef": 5808, // Lucidata
		"00a0f1": 2773, // MTI
		"00a0f2": 5809, // Infotek
		"00a0f3": 5810, // Staubli
		"00a0f4": 5811, // GE
		"00a0f5": 5812, // Radguard
		"00a0f6": 5813, // AutoGas
		"00a0f7": 5814, // V.I
		"00a0f8": 768,  // Zebra
		"00a0f9": 5815, // Bintec
		"00a0fa": 31,   // Marconi
		"00a0fb": 5816, // Toray
		"00a0fe": 5817, // Boston
		"00a265": 5818, // M2Motive
		"00a289": 2,    // Cisco
		"00a2da": 5819, // INAT
		"00a2ee": 2,    // Cisco
		"00a388": 3513, // BSkyB
		"00a38e": 2,    // Cisco
		"00a3d1": 2,    // Cisco
		"00a45f": 3335, // Huawei
		"00a509": 5820, // WigWag
		"00a554": 421,  // Intel
		"00a5bf": 2,    // Cisco
		"00a6ca": 2,    // Cisco
		"00a742": 2,    // Cisco
		"00aa00": 421,  // Intel
		"00aa01": 421,  // Intel
		"00aa02": 421,  // Intel
		"00aa6e": 2,    // Cisco
		"00aa70": 869,  // LG
		"00ab48": 5821, // eero
		"00ace0": 125,  // ARRIS
		"00ad24": 803,  // D-Link
		"00add5": 3335, // Huawei
		"00aecd": 5822, // Pensando
		"00aefa": 2057, // Murata
		"00af1f": 2,    // Cisco
		"00b017": 5823, // InfoGear
		"00b01c": 5824, // Westport
		"00b01e": 5825, // Rantic
		"00b02a": 5826, // ORSYS
		"00b02d": 5827, // ViaGate
		"00b03b": 5828, // HiQ
		"00b048": 31,   // Marconi
		"00b04a": 2,    // Cisco
		"00b052": 535,  // Atheros
		"00b064": 2,    // Cisco
		"00b069": 5829, // Honewell
		"00b086": 5830, // LocSoft
		"00b08e": 2,    // Cisco
		"00b091": 5831, // Transmeta
		"00b094": 5832, // Alaris
		"00b09a": 5833, // Morrow
		"00b0ae": 5746, // Symmetricom
		"00b0b3": 2358, // Xstreamis
		"00b0c2": 2,    // Cisco
		"00b0ce": 5834, // Viveris
		"00b0d0": 102,  // Dell
		"00b0db": 5835, // Nextcell
		"00b0e1": 2,    // Cisco
		"00b0ec": 5836, // Eacem
		"00b0ee": 5837, // Ajile
		"00b0f0": 5838, // Caly
		"00b0f5": 5839, // NetWorth
		"00b1e3": 2,    // Cisco
		"00b342": 5840, // MacroSAN
		"00b362": 545,  // Apple
		"00b56d": 5841, // David
		"00b5d0": 152,  // Samsung
		"00b5d6": 5842, // Omnibit
		"00b600": 5843, // VOIM
		"00b670": 2,    // Cisco
		"00b69f": 5844, // Latch
		"00b771": 2,    // Cisco
		"00b7a8": 5845, // Heinzinger
		"00b8b3": 2,    // Cisco
		"00b8b6": 687,  // Motorola
		"00bb01": 5846, // Octothorpe
		"00bb1c": 3335, // Huawei
		"00bb3a": 5439, // Amazon
		"00bb60": 421,  // Intel
		"00bb8e": 5847, // HME
		"00bbc1": 87,   // Canon
		"00bbf0": 5848, // Ungermann-Bass
		"00bc60": 2,    // Cisco
		"00bd27": 5849, // Exar
		"00bd3a": 457,  // Nokia
		"00bd3e": 3468, // Vizio
		"00be3b": 3335, // Huawei
		"00be43": 102,  // Dell
		"00be75": 2,    // Cisco
		"00be9e": 4922, // Fiberhome
		"00bf15": 5850, // Genetec
		"00bf61": 152,  // Samsung
		"00bf77": 2,    // Cisco
		"00c000": 5851, // Lanoptics
		"00c002": 2075, // Sercomm
		"00c003": 5852, // Globalnet
		"00c005": 5853, // Livingston
		"00c009": 5854, // KT
		"00c00e": 5855, // Psitech
		"00c00f": 2084, // Quantum
		"00c012": 5856, // Netspan
		"00c013": 5857, // Netrix
		"00c017": 5858, // NetAlly
		"00c018": 5859, // Lanart
		"00c019": 5860, // Leap
		"00c01b": 5861, // Socket
		"00c01c": 3024, // Interlink
		"00c01f": 5862, // S.E.R.C.E.L
		"00c020": 5863, // Arco
		"00c021": 5864, // Netexpress
		"00c022": 5865, // Lasermaster
		"00c023": 5866, // Tutankhamon
		"00c025": 5867, // Dataproducts
		"00c026": 5868, // Lans
		"00c027": 5869, // Cipher
		"00c028": 5870, // Jasco
		"00c02c": 5871, // Centrum
		"00c02e": 5872, // Netwiz
		"00c02f": 5873, // Okuma
		"00c030": 5874, // Integrated
		"00c031": 3862, // Design
		"00c032": 5875, // I-Cubed
		"00c035": 5876, // Quintar
		"00c036": 5877, // Raytech
		"00c037": 5878, // Dynatem
		"00c040": 5879, // Ecci
		"00c042": 5880, // Datalux
		"00c043": 5881, // Stratacom
		"00c044": 5882, // Emcom
		"00c045": 5883, // Isolation
		"00c047": 5884, // Unimicro
		"00c04d": 5885, // Mitec
		"00c04e": 308,  // Comtrol
		"00c04f": 102,  // Dell
		"00c052": 5886, // Burr-Brown
		"00c053": 5657, // Aspect
		"00c056": 5887, // Somelec
		"00c057": 5888, // Myco
		"00c058": 5889, // Dataexpert
		"00c059": 5890, // Denso
		"00c05a": 5891, // Semaphore
		"00c05c": 5892, // Elonex
		"00c05d": 5893, // L&N
		"00c05e": 5894, // Vari-Lite
		"00c05f": 5895, // Fine-PAL
		"00c061": 5896, // Solectek
		"00c062": 5897, // Impulse
		"00c065": 5898, // Scope
		"00c066": 5899, // Docupoint
		"00c06c": 5900, // Svec
		"00c06d": 5901, // Boca
		"00c06e": 5902, // Haft
		"00c06f": 5903, // Komatsu
		"00c071": 5904, // Areanex
		"00c072": 5905, // KNX
		"00c073": 5906, // Xedia
		"00c075": 5907, // Xante
		"00c078": 2329, // Computer
		"00c079": 5908, // Fonsys
		"00c07b": 3298, // Ascend
		"00c07e": 5909, // Kubota
		"00c080": 2171, // Netstar
		"00c081": 5910, // Metrodata
		"00c082": 851,  // Moore
		"00c087": 5911, // Uunet
		"00c08a": 5912, // Lauterbach
		"00c08c": 5913, // Performance
		"00c08f": 2154, // Panasonic
		"00c093": 2310, // Alta
		"00c094": 5914, // VMX
		"00c095": 2478, // ZNYX
		"00c096": 5915, // Tamura
		"00c097": 5916, // Archipel
		"00c098": 5917, // Chuntex
		"00c099": 5918, // Yoshiki
		"00c09a": 5919, // Photonics
		"00c09b": 5028, // Tellabs
		"00c09e": 5920, // Cache
		"00c09f": 3072, // Quanta
		"00c0a2": 5921, // Intermedium
		"00c0a3": 5922, // Dual
		"00c0a4": 5923, // Unigraf
		"00c0a7": 5924, // Seel
		"00c0a8": 5925, // GVC
		"00c0ab": 18,   // Telco
		"00c0ac": 5926, // Gambit
		"00c0ad": 5927, // Marben
		"00c0af": 5928, // Teklogix
		"00c0b0": 5929, // GCC
		"00c0b2": 5930, // Norand
		"00c0b4": 5931, // Myson
		"00c0b6": 5932, // HVE
		"00c0b7": 5933, // APC
		"00c0b9": 5934, // Funk
		"00c0ba": 5935, // Netvantage
		"00c0bd": 5936, // Inex
		"00c0c1": 5937, // Quad/Graphics
		"00c0c2": 752,  // Infinite
		"00c0ca": 5938, // Alfa
		"00c0cb": 1999, // Control
		"00c0cc": 5939, // Telesciences
		"00c0cd": 5940, // Comelta
		"00c0d1": 5941, // Comtree
		"00c0d2": 5942, // Syntellect
		"00c0d4": 5943, // Axon
		"00c0d6": 5944, // J1
		"00c0da": 5945, // Nice
		"00c0db": 5946, // IPC
		"00c0dc": 5947, // EOS
		"00c0dd": 2026, // QLogic
		"00c0de": 5948, // Zcomm
		"00c0df": 5949, // KYE
		"00c0e0": 5324, // DSC
		"00c0e1": 5067, // Sonic
		"00c0e2": 5950, // Calcomp
		"00c0e3": 5951, // Ositech
		"00c0e4": 300,  // Siemens
		"00c0e5": 5952, // Gespac
		"00c0e6": 4213, // Verilink
		"00c0e7": 5953, // Fiberdata
		"00c0e8": 5954, // Plexcom
		"00c0e9": 5955, // OAK
		"00c0ea": 5956, // Array
		"00c0ec": 5957, // Dauphin
		"00c0ee": 2849, // Kyocera
		"00c0ef": 5256, // Abit
		"00c0f0": 4882, // Kingston
		"00c0f2": 5958, // Transition
		"00c0f3": 109,  // Network
		"00c0f5": 5959, // Metacomp
		"00c0f6": 5960, // Celan
		"00c0f7": 5961, // Engage
		"00c0fa": 5962, // Canary
		"00c0fb": 587,  // Advanced
		"00c0fd": 5963, // Prosum
		"00c0fe": 5964, // Aptec
		"00c14f": 5965, // DDL
		"00c164": 2,    // Cisco
		"00c1b1": 2,    // Cisco
		"00c2c6": 421,  // Intel
		"00c30a": 5698, // Xiaomi
		"00c3f4": 152,  // Samsung
		"00c52c": 826,  // Juniper
		"00c610": 545,  // Apple
		"00c88b": 2,    // Cisco
		"00cae5": 2,    // Cisco
		"00cb00": 67,   // Private
		"00cb51": 2049, // Sagemcom
		"00cc34": 826,  // Juniper
		"00cc3f": 3858, // Universal
		"00ccfc": 2,    // Cisco
		"00cdfe": 545,  // Apple
		"00d001": 5966, // VST
		"00d002": 5967, // Ditech
		"00d003": 5968, // Comda
		"00d004": 5969, // Pentacom
		"00d006": 2,    // Cisco
		"00d008": 5970, // Mactell
		"00d009": 4278, // Hsing
		"00d00b": 5971, // RHK
		"00d00e": 5972, // Pluris
		"00d010": 5973, // Convergent
		"00d012": 5974, // Gateworks
		"00d014": 5975, // Root
		"00d01b": 5976, // Mimaki
		"00d01c": 5977, // SBS
		"00d01e": 5978, // Pingtel
		"00d01f": 5979, // Senetas
		"00d021": 5980, // Regent
		"00d022": 5981, // Incredible
		"00d023": 5982, // Infortrend
		"00d024": 5983, // Cognex
		"00d025": 5984, // Xrosstech
		"00d028": 1526, // Harmonic
		"00d02a": 5985, // Voxent
		"00d02b": 5986, // Jetcell
		"00d02d": 5987, // Resideo
		"00d02f": 5988, // Vlsi
		"00d030": 5989, // Safetran
		"00d034": 5990, // Ormec
		"00d035": 5991, // Behavior
		"00d037": 125,  // ARRIS
		"00d038": 5992, // Fivemere
		"00d039": 5993, // Utilicom
		"00d03a": 5994, // Zoneworx
		"00d03b": 2690, // Vision
		"00d03c": 5995, // Vieo
		"00d03d": 5996, // Galileo
		"00d03e": 5997, // Rocketchips
		"00d03f": 2359, // American
		"00d040": 5998, // Sysmate
		"00d041": 5999, // Amigo
		"00d044": 6000, // Alidian
		"00d045": 6001, // Kvaser
		"00d046": 6002, // Dolby
		"00d047": 6003, // XN
		"00d048": 6004, // Ecton
		"00d049": 6005, // Impresstek
		"00d04a": 6006, // Presence
		"00d04e": 6007, // Logibag
		"00d04f": 6008, // Bitronics
		"00d052": 3298, // Ascend
		"00d053": 6009, // Connected
		"00d055": 6010, // Kathrein-Werke
		"00d056": 6011, // Somat
		"00d057": 6012, // Ultrak
		"00d058": 2,    // Cisco
		"00d05a": 6013, // Symbionics
		"00d05d": 6014, // Intelliworxx
		"00d05e": 6015, // Stratabeam
		"00d05f": 6016, // Valcom
		"00d060": 2154, // Panasonic
		"00d061": 6017, // Tremon
		"00d062": 6018, // Digigram
		"00d063": 2,    // Cisco
		"00d064": 6019, // Multitel
		"00d066": 6020, // Wintriss
		"00d067": 6021, // Campio
		"00d068": 6022, // Iwill
		"00d069": 6023, // Technologic
		"00d06a": 6024, // Linkup
		"00d06c": 6025, // Sharewave
		"00d06d": 6026, // Acrison
		"00d071": 6027, // Echelon
		"00d072": 6028, // Broadlogic
		"00d074": 6029, // Taqua
		"00d077": 827,  // Lucent
		"00d079": 2,    // Cisco
		"00d07a": 6030, // Amaquest
		"00d07b": 6031, // Comcam
		"00d07c": 6032, // Koyo
		"00d07d": 6033, // Cosine
		"00d07e": 6034, // Keycorp
		"00d07f": 6035, // Strategy
		"00d080": 6036, // Exabyte
		"00d082": 6037, // Iowave
		"00d083": 6038, // Invertex
		"00d084": 6039, // Nexcomm
		"00d086": 6040, // Foveon
		"00d087": 6041, // Microfirst
		"00d088": 125,  // ARRIS
		"00d089": 6042, // Dynacolor
		"00d08c": 6043, // Genoa
		"00d08f": 6044, // Ardent
		"00d090": 2,    // Cisco
		"00d091": 6045, // Smartsan
		"00d096": 160,  // 3Com
		"00d097": 2,    // Cisco
		"00d09a": 6046, // Filanet
		"00d09b": 6047, // Spectel
		"00d09c": 6048, // Kapadia
		"00d09d": 6049, // Veris
		"00d09e": 1939, // 2Wire
		"00d09f": 6050, // Novtek
		"00d0a4": 6051, // Alantro
		"00d0a6": 6052, // Lanbird
		"00d0aa": 5550, // Chase
		"00d0ac": 6053, // Commscope
		"00d0ad": 6054, // TL
		"00d0ae": 6055, // Oresis
		"00d0af": 6056, // Cutler-Hammer
		"00d0b0": 6057, // Bitswitch
		"00d0b1": 497,  // Omega
		"00d0b2": 4121, // Xiotech
		"00d0b4": 6058, // Katsujima
		"00d0b6": 6059, // Crescent
		"00d0b7": 421,  // Intel
		"00d0b8": 6060, // Iomega
		"00d0b9": 5518, // Microtek
		"00d0ba": 2,    // Cisco
		"00d0bb": 2,    // Cisco
		"00d0bc": 2,    // Cisco
		"00d0be": 6061, // Emutec
		"00d0bf": 5065, // Pivotal
		"00d0c0": 2,    // Cisco
		"00d0c2": 6062, // Balthazar
		"00d0c3": 6063, // Vivid
		"00d0c4": 6064, // Teratech
		"00d0c5": 6065, // Computational
		"00d0c7": 6066, // Pathway
		"00d0c8": 6067, // Prevas
		"00d0c9": 1703, // Advantech
		"00d0ca": 927,  // Intrinsyc
		"00d0cb": 6068, // Dasan
		"00d0cd": 6069, // Atan
		"00d0ce": 6070, // iSystem
		"00d0d1": 4215, // Sycamore
		"00d0d2": 6071, // Epilog
		"00d0d3": 2,    // Cisco
		"00d0d4": 6072, // V-Bits
		"00d0d5": 6073, // Grundig
		"00d0d7": 6074, // B2C2
		"00d0d8": 160,  // 3Com
		"00d0db": 6075, // Mcquay
		"00d0df": 6076, // Kuzumi
		"00d0e0": 6077, // Dooin
		"00d0e3": 6078, // ELE-Chem
		"00d0e4": 2,    // Cisco
		"00d0e5": 6079, // Solidum
		"00d0e6": 6080, // Ibond
		"00d0ea": 6081, // Nextone
		"00d0eb": 6082, // Lightera
		"00d0ec": 6083, // NAKAYO
		"00d0ed": 6084, // Xiox
		"00d0ee": 6085, // Dictaphone
		"00d0ef": 6086, // IGT
		"00d0f0": 6087, // Convision
		"00d0f1": 6088, // Sega
		"00d0f2": 6089, // Monterey
		"00d0f6": 457,  // Nokia
		"00d0f9": 6090, // Acute
		"00d0ff": 2,    // Cisco
		"00d11c": 6091, // Acetel
		"00d49e": 421,  // Intel
		"00d6fe": 2,    // Cisco
		"00d76d": 421,  // Intel
		"00d78f": 2,    // Cisco
		"00d861": 1812, // Micro-Star
		"00d8a2": 3335, // Huawei
		"00d9d1": 101,  // Sony
		"00da55": 2,    // Cisco
		"00db45": 6092, // Thamway
		"00db70": 545,  // Apple
		"00dbdf": 421,  // Intel
		"00dcb2": 185,  // Extreme
		"00dd00": 5848, // Ungermann-Bass
		"00dd01": 5848, // Ungermann-Bass
		"00dd02": 5848, // Ungermann-Bass
		"00dd03": 5848, // Ungermann-Bass
		"00dd04": 5848, // Ungermann-Bass
		"00dd05": 5848, // Ungermann-Bass
		"00dd06": 5848, // Ungermann-Bass
		"00dd07": 5848, // Ungermann-Bass
		"00dd08": 5848, // Ungermann-Bass
		"00dd09": 5848, // Ungermann-Bass
		"00dd0a": 5848, // Ungermann-Bass
		"00dd0b": 5848, // Ungermann-Bass
		"00dd0c": 5848, // Ungermann-Bass
		"00dd0d": 5848, // Ungermann-Bass
		"00dd0e": 5848, // Ungermann-Bass
		"00dd0f": 5848, // Ungermann-Bass
		"00defb": 2,    // Cisco
		"00df1d": 2,    // Cisco
		"00e000": 4,    // Fujitsu
		"00e002": 6093, // Crossroads
		"00e004": 6094, // PMC-Sierra
		"00e005": 2925, // Technical
		"00e009": 108,  // Stratus
		"00e00a": 6095, // Diba
		"00e00b": 6096, // Rooftop
		"00e00c": 687,  // Motorola
		"00e00d": 228,  // Radiant
		"00e011": 6097, // Uniden
		"00e012": 6098, // Pluto
		"00e013": 4754, // Eastern
		"00e014": 2,    // Cisco
		"00e015": 6099, // Heiwa
		"00e017": 6100, // EXXACT
		"00e018": 1806, // ASUS
		"00e01a": 6101, // Comtec
		"00e01b": 6102, // Sphere
		"00e01c": 6103, // Cradlepoint
		"00e01d": 6104, // WebTV
		"00e01e": 2,    // Cisco
		"00e01f": 6105, // AVIDIA
		"00e020": 6106, // Tecnomen
		"00e021": 6107, // Freegate
		"00e023": 6108, // Telrad
		"00e024": 6109, // Gadzoox
		"00e025": 6110, // dit
		"00e027": 6111, // DUX
		"00e028": 6112, // Aptix
		"00e02b": 185,  // Extreme
		"00e02c": 5073, // AST
		"00e02d": 6113, // InnoMediaLogic
		"00e02e": 6114, // SPC
		"00e02f": 6115, // Mcns
		"00e030": 6116, // Melita
		"00e033": 6117, // E.E.P.D
		"00e034": 2,    // Cisco
		"00e036": 6118, // Pioneer
		"00e037": 5500, // Century
		"00e038": 6119, // Proxima
		"00e039": 6120, // Paradyne
		"00e03a": 16,   // Cabletron
		"00e03b": 6121, // Prominet
		"00e03c": 6122, // AdvanSys
		"00e03d": 6123, // Focon
		"00e03e": 6124, // Alfatech
		"00e03f": 6125, // Jaton
		"00e040": 6126, // DeskStation
		"00e041": 6127, // Cspi
		"00e042": 6128, // Pacom
		"00e043": 6129, // VitalCom
		"00e044": 6130, // Lsics
		"00e045": 6131, // Touchwave
		"00e047": 6132, // Infocus
		"00e048": 6133, // SDL
		"00e049": 6134, // MICROWI
		"00e04a": 6135, // ZX
		"00e04e": 1337, // Sanyo
		"00e04f": 2,    // Cisco
		"00e051": 6136, // Talx
		"00e052": 90,   // Brocade
		"00e053": 6137, // Cellport
		"00e055": 6138, // Ingenieria
		"00e056": 6139, // Holontech
		"00e05c": 6140, // PHC
		"00e05d": 6141, // Unitec
		"00e05f": 6142, // e-Net
		"00e060": 6143, // Sherwood
		"00e061": 6144, // EdgePoint
		"00e062": 6145, // Host
		"00e063": 16,   // Cabletron
		"00e064": 152,  // Samsung
		"00e066": 6146, // ProMax
		"00e068": 6147, // Merrimac
		"00e069": 6148, // Jaycor
		"00e06a": 6149, // Kapsch
		"00e06c": 3430, // Ultra
		"00e06d": 6150, // Compuware
		"00e06f": 125,  // ARRIS
		"00e070": 6151, // DH
		"00e072": 6152, // Lynk
		"00e074": 6153, // Tiernan
		"00e075": 4213, // Verilink
		"00e077": 6154, // Webgear
		"00e078": 6155, // Berkeley
		"00e079": 6156, // A.T.N.R
		"00e07a": 6157, // Mikrodidakt
		"00e07b": 6158, // BAY
		"00e07c": 2328, // Mettler-Toledo
		"00e07d": 9,    // Netronix
		"00e081": 6159, // Tyan
		"00e082": 6160, // Anerma
		"00e083": 6161, // Jato
		"00e088": 6162, // LTX-Credence
		"00e089": 4206, // ION
		"00e08b": 2026, // QLogic
		"00e08c": 6163, // Neoparadigm
		"00e08d": 6164, // Pressure
		"00e08e": 6165, // Utstarcom
		"00e08f": 2,    // Cisco
		"00e091": 869,  // LG
		"00e092": 6166, // Admtek
		"00e093": 6167, // Ackfin
		"00e096": 6168, // Shimadzu
		"00e098": 2557, // AboCom
		"00e099": 6169, // Samson
		"00e09a": 4089, // Positron
		"00e09b": 5961, // Engage
		"00e09c": 6170, // MII
		"00e09d": 6171, // Sarnoff
		"00e09e": 2084, // Quantum
		"00e0a0": 6172, // Wiltron
		"00e0a2": 6173, // Microslate
		"00e0a3": 2,    // Cisco
		"00e0a6": 6174, // Telogy
		"00e0a8": 6175, // SAT
		"00e0aa": 6176, // Electrosonic
		"00e0ac": 6177, // Midsco
		"00e0ad": 6178, // EES
		"00e0ae": 6179, // Xaqti
		"00e0b0": 2,    // Cisco
		"00e0b2": 6180, // Telmax
		"00e0b3": 6181, // EtherWAN
		"00e0b5": 6044, // Ardent
		"00e0b6": 6182, // Entrada
		"00e0b7": 6183, // Cosworth
		"00e0b9": 6184, // Byas
		"00e0bb": 6185, // NBX
		"00e0bc": 6186, // Symon
		"00e0bd": 1410, // Interface
		"00e0be": 6187, // Genroco
		"00e0c5": 6188, // Bcom
		"00e0c6": 6189, // LINK2IT
		"00e0c9": 6190, // AutomatedLogic
		"00e0cb": 6191, // Reson
		"00e0cc": 6192, // Hero
		"00e0ce": 6193, // ARN
		"00e0d0": 6194, // Netspeed
		"00e0d1": 6195, // Telsis
		"00e0d2": 6196, // Versanet
		"00e0d3": 6197, // DATENTECHNIK
		"00e0d4": 6198, // Excellent
		"00e0d5": 129,  // Emulex
		"00e0d7": 3950, // Sunshine
		"00e0d8": 6199, // LANBit
		"00e0d9": 6200, // Tazmo
		"00e0db": 6201, // ViaVideo
		"00e0dc": 6202, // Nexware
		"00e0dd": 1800, // Zenith
		"00e0de": 6203, // Datax
		"00e0df": 6204, // DZS
		"00e0e0": 6205, // SI
		"00e0e1": 6206, // G2
		"00e0e2": 6207, // Innova
		"00e0e3": 6208, // SK-Elektronik
		"00e0e5": 6209, // Cinco
		"00e0e6": 6210, // IncAA
		"00e0ea": 6211, // Innovat
		"00e0eb": 6212, // Digicom
		"00e0ec": 6213, // Celestica
		"00e0ed": 6214, // Silicom
		"00e0ef": 6215, // Dionex
		"00e0f0": 6216, // Abler
		"00e0f1": 6217, // That
		"00e0f3": 6218, // WebSprint
		"00e0f4": 6219, // INSIDE
		"00e0f5": 6220, // Teles
		"00e0f6": 6221, // Decision
		"00e0f7": 2,    // Cisco
		"00e0f9": 2,    // Cisco
		"00e0fa": 6222, // TRL
		"00e0fb": 6223, // Leightronix
		"00e0fc": 3335, // Huawei
		"00e0fd": 6224, // A-Trend
		"00e0fe": 2,    // Cisco
		"00e16d": 2,    // Cisco
		"00e175": 6225, // AK-Systems
		"00e18c": 421,  // Intel
		"00e3b2": 152,  // Samsung
		"00e406": 3335, // Huawei
		"00e421": 101,  // Sony
		"00e666": 1957, // ARIMA
		"00e6d3": 6226, // Nixdorf
		"00e6e8": 6227, // Netzin
		"00e7e3": 3032, // ZTE
		"00e93a": 3005, // Azurewave
		"00eabd": 2,    // Cisco
		"00eb2d": 101,  // Sony
		"00ebd5": 2,    // Cisco
		"00ebd8": 6228, // Mercusys
		"00ec0a": 5698, // Xiaomi
		"00edb8": 2849, // Kyocera
		"00eeab": 2,    // Cisco
		"00eebd": 1341, // HTC
		"00f051": 6229, // KWB
		"00f28b": 2,    // Cisco
		"00f361": 5439, // Amazon
		"00f39f": 545,  // Apple
		"00f403": 6230, // Orbis
		"00f46f": 152,  // Samsung
		"00f48d": 2874, // Liteon
		"00f4b9": 545,  // Apple
		"00f620": 3522, // Google
		"00f663": 2,    // Cisco
		"00f76f": 545,  // Apple
		"00f81c": 3335, // Huawei
		"00f82c": 2,    // Cisco
		"00f871": 3461, // Demant
		"00fa21": 152,  // Samsung
		"00fa3b": 6231, // Cloos
		"00fc58": 6232, // WebSilicon
		"00fc8b": 5439, // Amazon
		"00fc8d": 870,  // Hitron
		"00fcba": 2,    // Cisco
		"00fd22": 2,    // Cisco
		"00fd45": 302,  // HP
		"00fd4c": 6233, // Nevatec
		"00fec8": 2,    // Cisco
		"020701": 1045, // Racal-Datacom
		"021c7c": 3799, // Perq
		"02608c": 160,  // 3Com
		"027001": 1045, // Racal-Datacom
		"02bb01": 5846, // Octothorpe
		"02c08c": 160,  // 3Com
		"02e6d3": 6226, // Nixdorf
		"04021f": 3335, // Huawei
		"0403d6": 1427, // Nintendo
		"04072e": 1319, // VTech
		"040973": 302,  // HP
		"0409a5": 4576, // HFR
		"040ae0": 6234, // Xmit
		"040cce": 545,  // Apple
		"040d84": 1655, // Silicon
		"040e3c": 302,  // HP
		"04106b": 5698, // Xiaomi
		"041552": 545,  // Apple
		"0415d9": 6235, // Viwone
		"04180f": 152,  // Samsung
		"0418b6": 67,   // Private
		"0418d6": 2979, // Ubiquiti
		"041a04": 6236, // WaveIP
		"041b6d": 869,  // LG
		"041b94": 6145, // Host
		"041bba": 152,  // Samsung
		"041dc7": 3032, // ZTE
		"041e64": 545,  // Apple
		"041e7a": 6237, // DSPWorks
		"042084": 3032, // ZTE
		"04209a": 2154, // Panasonic
		"042144": 3940, // Sunitec
		"0425c5": 3335, // Huawei
		"042665": 545,  // Apple
		"042728": 612,  // Microsoft
		"042758": 3335, // Huawei
		"042ae2": 2,    // Cisco
		"042bbb": 6238, // PicoCELA
		"042f56": 6239, // ATOCS
		"0432f4": 6240, // Partron
		"043389": 3335, // Huawei
		"0433c2": 421,  // Intel
		"0434f6": 687,  // Motorola
		"043855": 6241, // Scopus-Belgium
		"043f72": 432,  // Mellanox
		"044169": 6242, // GoPro
		"04421a": 1806, // ASUS
		"044665": 2057, // Murata
		"04489a": 545,  // Apple
		"04495d": 3335, // Huawei
		"044a50": 6243, // Ramaxel
		"044a6c": 3335, // Huawei
		"044ac6": 6244, // Aipon
		"044bed": 545,  // Apple
		"044e06": 306,  // Ericsson
		"044e5a": 125,  // ARRIS
		"044eaf": 869,  // LG
		"044f17": 531,  // HUMAX
		"044f4c": 3335, // Huawei
		"044f8b": 6245, // Adapteva
		"044faa": 2738, // Ruckus
		"0452c7": 1819, // Bose
		"0452f3": 545,  // Apple
		"045453": 545,  // Apple
		"0455ca": 6246, // BriView
		"0456e5": 421,  // Intel
		"04572f": 6247, // Sertel
		"045a95": 457,  // Nokia
		"045c06": 6248, // Zmodo
		"045c6c": 826,  // Juniper
		"045d4b": 101,  // Sony
		"045d56": 6249, // camtron
		"045fb9": 2,    // Cisco
		"046273": 2,    // Cisco
		"0463e0": 6250, // Nome
		"046565": 6251, // Testop
		"046865": 545,  // Apple
		"0469f8": 545,  // Apple
		"046b1b": 6252, // SysDINE
		"046c59": 421,  // Intel
		"046c9d": 2,    // Cisco
		"046d42": 6253, // Bryston
		"046e49": 6254, // TaiYear
		"0470bc": 6255, // Globalstar
		"047153": 6256, // Sernet
		"047295": 545,  // Apple
		"047503": 3335, // Huawei
		"0475f5": 6257, // Csst
		"04766e": 430,  // Alpsalpine
		"0476b0": 2,    // Cisco
		"047970": 3335, // Huawei
		"047aae": 3335, // Huawei
		"047d7b": 3072, // Quanta
		"047e4a": 6258, // moobox
		"047f0e": 6259, // Barrot
		"04819b": 3513, // BSkyB
		"0481ae": 6260, // Clack
		"04848a": 6261, // 7INOVA
		"04885f": 3335, // Huawei
		"0488e2": 6262, // Beats
		"048a15": 620,  // Avaya
		"048ae1": 833,  // Flextronics
		"048b42": 6263, // Skspruce
		"048c03": 6264, // ThinPAD
		"048c16": 3335, // Huawei
		"048c9a": 3335, // Huawei
		"048d38": 2363, // Netcore
		"049081": 5822, // Pensando
		"049162": 706,  // Microchip
		"049226": 1806, // ASUS
		"0492ee": 6265, // iway
		"04946b": 6266, // Tecno
		"049573": 3032, // ZTE
		"0495e6": 6267, // Tenda
		"0498f3": 430,  // Alpsalpine
		"0499b9": 545,  // Apple
		"049dfe": 6268, // Hivesystem
		"049f06": 6269, // Smobile
		"049f15": 67,   // Private
		"049f81": 5516, // Netscout
		"049fca": 3335, // Huawei
		"04a151": 1368, // Netgear
		"04a222": 2640, // Arcadyan
		"04a2f3": 4922, // Fiberhome
		"04a3f3": 6270, // Emicon
		"04a82a": 457,  // Nokia
		"04ab18": 5694, // Elecom
		"04ab6a": 6271, // Chun-il
		"04b0e7": 3335, // Huawei
		"04b167": 5698, // Xiaomi
		"04b1a1": 152,  // Samsung
		"04b3b6": 6272, // Seamap
		"04b429": 152,  // Samsung
		"04b466": 6273, // BSP
		"04b648": 6274, // Zenner
		"04b86a": 3513, // BSkyB
		"04b9e3": 152,  // Samsung
		"04ba1c": 3335, // Huawei
		"04ba8d": 152,  // Samsung
		"04bc9f": 374,  // Calix
		"04bd70": 3335, // Huawei
		"04bd88": 1685, // Aruba
		"04bd97": 2,    // Cisco
		"04bdbf": 152,  // Samsung
		"04bf6d": 2693, // ZyXEL
		"04bfa8": 6275, // ISB
		"04c06f": 3335, // Huawei
		"04c09c": 5028, // Tellabs
		"04c1b9": 4922, // Fiberhome
		"04c1d8": 3335, // Huawei
		"04c23e": 1341, // HTC
		"04c241": 457,  // Nokia
		"04c461": 2057, // Murata
		"04c5a4": 2,    // Cisco
		"04c807": 5698, // Xiaomi
		"04c880": 6276, // Samtec
		"04c991": 6277, // Phistek
		"04c9d9": 1238, // Dish
		"04cb1d": 6278, // Traka
		"04ccbc": 3335, // Huawei
		"04cd15": 1655, // Silicon
		"04ce14": 6279, // Wilocity
		"04cf25": 6280, // Manycolors
		"04cf4b": 421,  // Intel
		"04cf8c": 6281, // XIAOMI
		"04d13a": 5698, // Xiaomi
		"04d320": 6282, // Itel
		"04d395": 687,  // Motorola
		"04d3b0": 421,  // Intel
		"04d3b5": 3335, // Huawei
		"04d3cf": 545,  // Apple
		"04d437": 6283, // ZNV
		"04d4c4": 1806, // ASUS
		"04d590": 1323, // Fortinet
		"04d6aa": 152,  // Samsung
		"04d921": 6284, // Occuspace
		"04d9f5": 1806, // ASUS
		"04dad2": 2,    // Cisco
		"04db56": 545,  // Apple
		"04db8a": 6285, // Suntech
		"04dd4c": 6286, // Velocytech
		"04dedb": 6287, // Rockport
		"04e0c4": 6288, // Triumph-Adler
		"04e31a": 2049, // Sagemcom
		"04e536": 545,  // Apple
		"04e56e": 6289, // THUB
		"04e598": 5698, // Xiaomi
		"04e662": 6290, // Acroname
		"04e676": 4499, // AMPAK
		"04e77e": 6291, // We
		"04e795": 3335, // Huawei
		"04e9e5": 6292, // Pjrc.com
		"04ea56": 421,  // Intel
		"04eb40": 2,    // Cisco
		"04ecbb": 4922, // Fiberhome
		"04ecd8": 421,  // Intel
		"04ed33": 421,  // Intel
		"04ee91": 6293, // x-fabric
		"04f021": 5082, // Compex
		"04f03e": 3335, // Huawei
		"04f13e": 545,  // Apple
		"04f169": 3335, // Huawei
		"04f352": 3335, // Huawei
		"04f4bc": 6294, // Xena
		"04f7e4": 545,  // Apple
		"04f8c2": 2669, // Flaircomm
		"04f8f8": 6295, // Edgecore
		"04f938": 3335, // Huawei
		"04f993": 6296, // Infinix
		"04f9d9": 6297, // Speaker
		"04fa3f": 6298, // Opticore
		"04fe31": 152,  // Samsung
		"04fe7f": 2,    // Cisco
		"04fe8d": 3335, // Huawei
		"04fea1": 6299, // Fihonest
		"04ff08": 3335, // Huawei
		"080001": 6300, // Computervision
		"080002": 4652, // Bridge
		"080003": 587,  // Advanced
		"080004": 6301, // Cromemco
		"080005": 6302, // Symbolics
		"080006": 300,  // Siemens
		"080007": 545,  // Apple
		"080009": 302,  // HP
		"08000a": 1480, // Nestar
		"08000b": 38,   // Unisys
		"08000d": 639,  // International
		"08000e": 6303, // NCR
		"08000f": 1217, // Mitel
		"080011": 6304, // Tektronix
		"080013": 6305, // Exxon
		"080014": 6306, // Excelan
		"08001b": 102,  // Dell
		"08001d": 6307, // Able
		"08001e": 5424, // Apollo
		"08001f": 3206, // Sharp
		"080020": 11,   // Oracle
		"080021": 6308, // 3M
		"080022": 6309, // NBI
		"080023": 2154, // Panasonic
		"080024": 6310, // 10NET/DCA
		"080029": 6311, // Megatek
		"08002a": 886,  // Mosaic
		"08002d": 6312, // LAN-TEC
		"08002e": 6313, // Metaphor
		"08002f": 1040, // Prime
		"080030": 6314, // Cern
		"080032": 6315, // Tigan
		"080034": 6316, // Filenet
		"080035": 6317, // Microfive
		"080036": 6318, // Intergraph
		"080039": 154,  // Spider
		"08003a": 6319, // Orcatech
		"08003b": 6320, // Torus
		"08003e": 6321, // Codex
		"080040": 6322, // Ferranti
		"080042": 6323, // MACNICA
		"080043": 6324, // Pixel
		"080044": 5841, // David
		"080045": 5141, // Concurrent
		"080046": 101,  // Sony
		"080047": 6325, // Sequent
		"080049": 6326, // Univation
		"08004a": 6327, // Banyan
		"08004b": 6328, // Planning
		"08004c": 6329, // Hydra
		"08004d": 6330, // Corvus
		"08004e": 160,  // 3Com
		"08004f": 215,  // Cygnet
		"080050": 1465, // Daisy
		"080051": 6331, // ExperData
		"080052": 6332, // Insystec
		"08005a": 372,  // IBM
		"08005b": 6333, // VTA
		"08005d": 6334, // Gould
		"08005e": 6335, // Counterpoint
		"08005f": 6336, // Saber
		"080061": 6337, // Jarogate
		"080063": 6338, // Plessey
		"080064": 6339, // Sitasys
		"080065": 5208, // Genrad
		"080066": 6340, // Agfa
		"080067": 6341, // ComDesign
		"080068": 6342, // Ridge
		"08006a": 1057, // AT&T
		"08006b": 5210, // Accel
		"08006e": 6343, // Masscomp
		"080071": 6344, // Matra
		"080073": 6345, // Tecmar
		"080074": 6346, // Casio
		"080077": 6347, // TSL
		"080078": 6348, // Accell
		"08007a": 6349, // Indata
		"08007b": 1337, // Sanyo
		"08007c": 6350, // Vitalink
		"080081": 6351, // Astech
		"080082": 6352, // Veritas
		"080084": 6353, // Tomen
		"080085": 6354, // Elxsi
		"080087": 5770, // Xyplex
		"080088": 90,   // Brocade
		"080089": 6355, // Kinetics
		"08008a": 6356, // PerfTech
		"08008b": 6357, // Pyramid
		"08008c": 109,  // Network
		"08008d": 6358, // Xyvision
		"08008e": 6359, // Tandem
		"08008f": 6360, // Chipcom
		"080090": 6361, // Sonoma
		"080205": 3335, // Huawei
		"08028e": 1368, // Netgear
		"080581": 1916, // Roku
		"0805e2": 826,  // Juniper
		"0808c2": 152,  // Samsung
		"0808ea": 6362, // Amsc
		"0809b6": 6363, // Masimo
		"080d84": 6364, // GECO
		"080ffa": 6365, // KSP
		"08115e": 6366, // Bitel
		"081196": 421,  // Intel
		"0812a5": 5439, // Amazon
		"08152f": 152,  // Samsung
		"0816d5": 6367, // Goertek
		"081735": 2,    // Cisco
		"0817f4": 372,  // IBM
		"08181a": 3032, // ZTE
		"0819a6": 3335, // Huawei
		"081c6e": 5698, // Xiaomi
		"081f3f": 6368, // WondaLink
		"081f71": 1595, // TP-Link
		"081feb": 6369, // BinCube
		"081ff3": 2,    // Cisco
		"0821ef": 152,  // Samsung
		"0823b2": 6370, // Vivo
		"082522": 6371, // Advansee
		"082525": 5698, // Xiaomi
		"082697": 2693, // ZyXEL
		"082cb6": 545,  // Apple
		"082ced": 6372, // Technity
		"082e36": 3335, // Huawei
		"082e5f": 302,  // HP
		"082fe9": 3335, // Huawei
		"08318b": 3335, // Huawei
		"0831a4": 3335, // Huawei
		"083571": 6373, // CASwell
		"0835b2": 6374, // CoreEdge
		"0836c9": 1368, // Netgear
		"08373d": 152,  // Samsung
		"08379c": 6375, // Topaz
		"0838e6": 687,  // Motorola
		"083a5c": 6376, // Junilab
		"083af2": 6377, // Espressif
		"083d88": 152,  // Samsung
		"083e0c": 125,  // ARRIS
		"083e5d": 2049, // Sagemcom
		"083f3e": 6378, // WSH
		"083f76": 6379, // Intellian
		"083fbc": 3032, // ZTE
		"084027": 6380, // Gridstore
		"0840f3": 6267, // Tenda
		"084296": 6381, // Mobile
		"0845d1": 2,    // Cisco
		"084656": 6382, // VEO-Labs
		"08474c": 457,  // Nokia
		"084e1c": 6383, // H2A
		"084f0a": 3335, // Huawei
		"084fa9": 2,    // Cisco
		"084ff9": 2,    // Cisco
		"085531": 1784, // Routerboard.com
		"085700": 1595, // TP-Link
		"0857fb": 5439, // Amazon
		"085a11": 803,  // D-Link
		"085ae0": 6384, // Recovision
		"085b0e": 1323, // Fortinet
		"085bd6": 421,  // Intel
		"085bda": 6385, // CliniCare
		"085c1b": 3335, // Huawei
		"085ddd": 4254, // Mercury
		"08606e": 1806, // ASUS
		"086083": 3032, // ZTE
		"086266": 1806, // ASUS
		"086361": 3335, // Huawei
		"086698": 545,  // Apple
		"0868ea": 6386, // Eito
		"086a0a": 2545, // Askey
		"086ac5": 421,  // Intel
		"086ae5": 5439, // Amazon
		"086bd7": 1655, // Silicon
		"086d41": 545,  // Apple
		"087045": 545,  // Apple
		"087190": 421,  // Intel
		"087402": 545,  // Apple
		"087572": 6387, // Obelux
		"087695": 6388, // Auto
		"087808": 152,  // Samsung
		"08798c": 3335, // Huawei
		"087999": 6389, // AIM
		"087a4c": 3335, // Huawei
		"087baa": 6390, // Svyazkomplektservice
		"087c39": 5439, // Amazon
		"087cbe": 6391, // Quintic
		"087d21": 6392, // Altasec
		"087e64": 6393, // Technicolor
		"087f98": 6370, // Vivo
		"0881b2": 4080, // Logitech
		"0881f4": 826,  // Juniper
		"08849d": 5439, // Amazon
		"08855b": 63,   // Kontron
		"088620": 6266, // Tecno
		"08863b": 2469, // Belkin
		"0887c7": 545,  // Apple
		"088c2c": 152,  // Samsung
		"088dc8": 6394, // Ryowa
		"088e4f": 6395, // SF
		"088e90": 421,  // Intel
		"088edc": 545,  // Apple
		"088f2c": 5002, // Amber
		"0890ba": 6396, // Danlaw
		"089204": 102,  // Dell
		"089356": 3335, // Huawei
		"0894ef": 1592, // Wistron
		"08952a": 6393, // Technicolor
		"0896ad": 2,    // Cisco
		"0896d7": 621,  // AVM
		"089734": 302,  // HP
		"089798": 358,  // Compal
		"0899e8": 6397, // KEMAS
		"089ac7": 3032, // ZTE
		"089b4b": 6398, // iKuai
		"089bf1": 5821, // eero
		"089e01": 3072, // Quanta
		"089e08": 3522, // Google
		"089e84": 3335, // Huawei
		"08a5c8": 5436, // Sunnovo
		"08a6bc": 5439, // Amazon
		"08a7c0": 6393, // Technicolor
		"08a842": 3335, // Huawei
		"08a95a": 3005, // Azurewave
		"08aa55": 687,  // Motorola
		"08aa89": 3032, // ZTE
		"08acc4": 6399, // FMTech
		"08aed6": 152,  // Samsung
		"08af78": 6400, // Totus
		"08b055": 2545, // Askey
		"08b0a7": 6401, // Truebeyond
		"08b258": 826,  // Juniper
		"08b3af": 6370, // Vivo
		"08b49d": 6266, // Tecno
		"08b4b1": 3522, // Google
		"08b4cf": 6402, // Abicom
		"08b738": 6403, // Lite-On
		"08ba22": 6404, // Swaive
		"08bb3c": 833,  // Flextronics
		"08bd43": 1368, // Netgear
		"08be09": 6405, // Astrol
		"08be77": 6406, // Green
		"08beac": 115,  // Edimax
		"08bfa0": 152,  // Samsung
		"08c021": 3335, // Huawei
		"08c06c": 3335, // Huawei
		"08c0eb": 432,  // Mellanox
		"08c5e1": 152,  // Samsung
		"08c6b3": 4146, // Qtech
		"08c729": 545,  // Apple
		"08cc27": 687,  // Motorola
		"08cc68": 2,    // Cisco
		"08cca7": 2,    // Cisco
		"08d09f": 2,    // Cisco
		"08d23e": 421,  // Intel
		"08d29a": 6407, // Proformatique
		"08d34b": 6408, // Techman
		"08d40c": 421,  // Intel
		"08d42b": 152,  // Samsung
		"08d46a": 869,  // LG
		"08d59d": 2049, // Sagemcom
		"08d5c0": 6409, // Seers
		"08df1f": 1819, // Bose
		"08dfcb": 6410, // Systrome
		"08e672": 6411, // Jebsee
		"08e689": 545,  // Apple
		"08e7e5": 3335, // Huawei
		"08e84f": 3335, // Huawei
		"08e9f6": 4499, // AMPAK
		"08ea44": 185,  // Extreme
		"08eb74": 531,  // HUMAX
		"08eca9": 152,  // Samsung
		"08ecf5": 2,    // Cisco
		"08ed9d": 6266, // Tecno
		"08ee8b": 152,  // Samsung
		"08f1ea": 302,  // HP
		"08f458": 3335, // Huawei
		"08f4ab": 545,  // Apple
		"08f606": 3032, // ZTE
		"08f69c": 545,  // Apple
		"08f6f8": 6412, // GET
		"08f8bc": 545,  // Apple
		"08fa28": 3335, // Huawei
		"08fa79": 6370, // Vivo
		"08fbea": 4499, // AMPAK
		"08fc52": 6413, // OpenXS
		"08fc88": 152,  // Samsung
		"08fd0e": 152,  // Samsung
		"08ff44": 545,  // Apple
		"0c01db": 6296, // Infinix
		"0c0227": 6393, // Technicolor
		"0c02bd": 152,  // Samsung
		"0c0535": 826,  // Juniper
		"0c08b4": 531,  // HUMAX
		"0c0e76": 803,  // D-Link
		"0c1105": 6414, // Akuvox
		"0c1167": 2,    // Cisco
		"0c1262": 3032, // ZTE
		"0c130b": 6415, // Uniqoteq
		"0c1420": 152,  // Samsung
		"0c1539": 545,  // Apple
		"0c15c5": 6416, // SDTEC
		"0c1773": 3335, // Huawei
		"0c17f1": 6417, // Telecsys
		"0c191f": 6418, // Inform
		"0c19f8": 545,  // Apple
		"0c1c19": 6419, // Longconn
		"0c1c1a": 5821, // eero
		"0c1c20": 6420, // Kakao
		"0c1daf": 5698, // Xiaomi
		"0c1dc2": 6421, // SeAH
		"0c1ef7": 6422, // Omni-ID
		"0c2026": 6423, // noax
		"0c20d3": 6370, // Vivo
		"0c2138": 6424, // Hengstler
		"0c2724": 2,    // Cisco
		"0c2755": 6425, // Valuable
		"0c29ef": 102,  // Dell
		"0c2a86": 4922, // Fiberhome
		"0c2c54": 3335, // Huawei
		"0c2d89": 6426, // QiiQ
		"0c2fb0": 152,  // Samsung
		"0c3021": 545,  // Apple
		"0c31dc": 3335, // Huawei
		"0c354f": 457,  // Nokia
		"0c35fe": 4922, // Fiberhome
		"0c3747": 3032, // ZTE
		"0c3796": 6427, // Bizlink
		"0c37dc": 3335, // Huawei
		"0c383e": 6428, // Fanvil
		"0c3b50": 545,  // Apple
		"0c3e9f": 545,  // Apple
		"0c413e": 612,  // Microsoft
		"0c41e9": 3335, // Huawei
		"0c42a1": 432,  // Mellanox
		"0c4314": 1655, // Silicon
		"0c43f9": 5439, // Amazon
		"0c45ba": 3335, // Huawei
		"0c473d": 870,  // Hitron
		"0c47c9": 5439, // Amazon
		"0c4885": 869,  // LG
		"0c48c6": 6213, // Celestica
		"0c4b54": 1595, // TP-Link
		"0c4c39": 6429, // MitraStar
		"0c4de9": 545,  // Apple
		"0c4ec0": 5299, // Maxlinear
		"0c4f9b": 3335, // Huawei
		"0c5101": 545,  // Apple
		"0c5415": 421,  // Intel
		"0c54a5": 5440, // Pegatron
		"0c54b9": 457,  // Nokia
		"0c5521": 6430, // Axiros
		"0c57eb": 6431, // Mueller
		"0c599c": 826,  // Juniper
		"0c6046": 6370, // Vivo
		"0c6127": 2247, // Actiontec
		"0c6803": 2,    // Cisco
		"0c6abc": 4922, // Fiberhome
		"0c6e4f": 6432, // PrimeVOLT
		"0c6f9c": 6433, // Shaw
		"0c704a": 3335, // Huawei
		"0c715d": 152,  // Samsung
		"0c718c": 6434, // TCT
		"0c722c": 1595, // TP-Link
		"0c72d9": 3032, // ZTE
		"0c7329": 2075, // Sercomm
		"0c74c2": 545,  // Apple
		"0c75bd": 2,    // Cisco
		"0c771a": 545,  // Apple
		"0c7a15": 421,  // Intel
		"0c8063": 1595, // TP-Link
		"0c8112": 67,   // Private
		"0c8126": 826,  // Juniper
		"0c8268": 1595, // TP-Link
		"0c839a": 3335, // Huawei
		"0c83cc": 2237, // Alpha
		"0c8408": 3335, // Huawei
		"0c8447": 4922, // Fiberhome
		"0c8484": 6435, // Zenovia
		"0c8525": 2,    // Cisco
		"0c8610": 826,  // Juniper
		"0c8910": 152,  // Samsung
		"0c8a87": 6436, // AgLogica
		"0c8b7d": 3468, // Vizio
		"0c8bd3": 6282, // Itel
		"0c8bfd": 421,  // Intel
		"0c8c8f": 6437, // Kamo
		"0c8cdc": 6438, // Suunto
		"0c8dca": 152,  // Samsung
		"0c8e29": 2640, // Arcadyan
		"0c8fff": 3335, // Huawei
		"0c93fb": 6439, // BNS
		"0c9541": 6440, // Chipsea
		"0c96bf": 3335, // Huawei
		"0c96cd": 4254, // Mercury
		"0c9838": 5698, // Xiaomi
		"0c9a3c": 421,  // Intel
		"0c9a42": 6441, // FN-Link
		"0c9d92": 1806, // ASUS
		"0c9e91": 6442, // Sankosha
		"0ca2f4": 6443, // Chameleon
		"0ca694": 3940, // Sunitec
		"0ca8a7": 152,  // Samsung
		"0caaee": 6444, // Ansjer
		"0cac05": 6445, // Unitend
		"0cac8a": 2049, // Sagemcom
		"0caebd": 6446, // Edifier
		"0cb088": 6447, // AITelecom
		"0cb319": 152,  // Samsung
		"0cb459": 6448, // Marketech
		"0cb4ef": 6449, // Digience
		"0cb527": 3335, // Huawei
		"0cb6d2": 803,  // D-Link
		"0cb771": 125,  // ARRIS
		"0cb815": 6377, // Espressif
		"0cb8e8": 5696, // Renesas
		"0cb912": 6450, // JM-DATA
		"0cbc9f": 545,  // Apple
		"0cbd51": 6434, // TCT
		"0cbf15": 5850, // Genetec
		"0cc3a7": 6451, // Meritec
		"0cc413": 3522, // Google
		"0cc47e": 6452, // EUCAST
		"0cc66a": 457,  // Nokia
		"0cc6ac": 6453, // Dags
		"0cc6cc": 3335, // Huawei
		"0cc6fd": 5698, // Xiaomi
		"0cc731": 6454, // Currant
		"0ccb85": 687,  // Motorola
		"0ccc26": 6455, // Airenetworks
		"0ccdd3": 6456, // Eastriver
		"0ccdfb": 6457, // EDIC
		"0ccfd1": 6458, // SPRINGWAVE
		"0cd0f8": 2,    // Cisco
		"0cd292": 421,  // Intel
		"0cd502": 2274, // Westell
		"0cd696": 6459, // Amimon
		"0cd6bd": 3335, // Huawei
		"0cd746": 545,  // Apple
		"0cd7c2": 6460, // Axium
		"0cd996": 2,    // Cisco
		"0cd9c1": 1399, // Visteon
		"0cdc7e": 6377, // Espressif
		"0cdccc": 6461, // Inala
		"0cdd24": 421,  // Intel
		"0cddef": 457,  // Nokia
		"0cdfa4": 152,  // Samsung
		"0ce041": 6462, // iDruide
		"0ce0dc": 152,  // Samsung
		"0ce0e4": 542,  // Plantronics
		"0ce441": 545,  // Apple
		"0ce4a0": 3335, // Huawei
		"0ce5a3": 6463, // SharkNinja
		"0ce5d3": 6151, // DH
		"0ce725": 612,  // Microsoft
		"0ceac9": 125,  // ARRIS
		"0cec8d": 687,  // Motorola
		"0cee99": 5439, // Amazon
		"0cef7c": 6464, // AnaCom
		"0cf019": 6465, // Malgn
		"0cf0b4": 1975, // Globalsat
		"0cf346": 5698, // Xiaomi
		"0cf4d5": 2738, // Ruckus
		"0cf5a4": 2,    // Cisco
		"0cf893": 125,  // ARRIS
		"0cf9c0": 3513, // BSkyB
		"0cfc83": 6466, // Airoha
		"0cfe45": 101,  // Sony
		"100000": 67,   // Private
		"100020": 545,  // Apple
		"10005a": 372,  // IBM
		"1000fd": 6467, // LaonPeople
		"100177": 3335, // Huawei
		"1002b5": 421,  // Intel
		"100501": 5440, // Pegatron
		"1005b1": 125,  // ARRIS
		"1005ca": 2,    // Cisco
		"1005e1": 457,  // Nokia
		"100645": 2049, // Sagemcom
		"1006ed": 2,    // Cisco
		"1007b6": 152,  // Samsung
		"1009f9": 5439, // Amazon
		"100ba9": 421,  // Intel
		"100c24": 6468, // pomdevices
		"100c6b": 1368, // Netgear
		"100d32": 6469, // Embedian
		"100d7f": 1368, // Netgear
		"100e7e": 826,  // Juniper
		"101081": 3032, // ZTE
		"1010b6": 6470, // McCain
		"101212": 6370, // Vivo
		"101218": 6471, // Korins
		"101248": 6472, // ITG
		"101331": 6393, // Technicolor
		"1013ee": 6473, // Justec
		"101b54": 3335, // Huawei
		"101c0c": 545,  // Apple
		"101d51": 6474, // 8Mesh
		"101dc0": 152,  // Samsung
		"101f74": 302,  // HP
		"102279": 6475, // ZeroDesktop
		"1027be": 6476, // Tvip
		"1027f5": 1595, // TP-Link
		"102831": 6477, // Morion
		"102959": 545,  // Apple
		"1029ab": 152,  // Samsung
		"102ab3": 5698, // Xiaomi
		"102b41": 152,  // Samsung
		"102c6b": 4499, // AMPAK
		"102c83": 6478, // Ximea
		"102cef": 6479, // EMU
		"102d96": 6480, // Looxcie
		"102f6b": 612,  // Microsoft
		"103025": 545,  // Apple
		"103034": 6481, // Cara
		"103047": 152,  // Samsung
		"10321d": 3335, // Huawei
		"10327e": 3335, // Huawei
		"103378": 6482, // FLECTRON
		"1033bf": 6393, // Technicolor
		"10341b": 6483, // Spacelink
		"103917": 152,  // Samsung
		"1039e9": 826,  // Juniper
		"103b59": 152,  // Samsung
		"103c59": 3032, // ZTE
		"103d1c": 421,  // Intel
		"103dea": 6484, // HFC
		"103f44": 5698, // Xiaomi
		"1040f3": 545,  // Apple
		"10417f": 545,  // Apple
		"104369": 6485, // Soundmax
		"104400": 3335, // Huawei
		"1045be": 6486, // Norphonic
		"1045f8": 6487, // LNT-Automation
		"1046b4": 6488, // FormericaOE
		"104780": 3335, // Huawei
		"104a7d": 421,  // Intel
		"104d15": 6489, // Viaanix
		"104d77": 428,  // Innovative
		"104e89": 797,  // Garmin
		"104f58": 1685, // Aruba
		"104fa8": 101,  // Sony
		"105072": 2075, // Sercomm
		"105107": 421,  // Intel
		"105172": 3335, // Huawei
		"10521c": 6377, // Espressif
		"105403": 6490, // INTARSO
		"105611": 125,  // ARRIS
		"1056ca": 2482, // Peplink
		"105887": 4922, // Fiberhome
		"105917": 6491, // Tonal
		"105932": 1916, // Roku
		"105af7": 6492, // ADB
		"105c3b": 6493, // Perma-Pipe
		"105cbf": 6494, // DuroByte
		"105ddc": 3335, // Huawei
		"105f06": 2247, // Actiontec
		"105fd4": 6495, // Tendyron
		"10604b": 302,  // HP
		"1062c9": 6496, // Adatis
		"1062d0": 6393, // Technicolor
		"1062e5": 302,  // HP
		"1062eb": 803,  // D-Link
		"1063c8": 2874, // Liteon
		"106530": 102,  // Dell
		"1065a3": 6497, // Panamax
		"1065cf": 6498, // Iqsim
		"10683f": 869,  // LG
		"106f3f": 1077, // Buffalo
		"1070fd": 432,  // Mellanox
		"107100": 3335, // Huawei
		"1071b3": 2693, // ZyXEL
		"107636": 6499, // Earda
		"10768a": 6500, // EoCell
		"1077b0": 4922, // Fiberhome
		"1077b1": 152,  // Samsung
		"10785b": 2247, // Actiontec
		"1078d2": 1115, // Elitegroup
		"107a86": 6501, // U&U
		"107b44": 1806, // ASUS
		"107bce": 457,  // Nokia
		"107bef": 2693, // ZyXEL
		"107d1a": 102,  // Dell
		"1083d2": 6502, // Microseven
		"10868c": 125,  // ARRIS
		"1088ce": 4922, // Fiberhome
		"1089fb": 152,  // Samsung
		"108a1b": 6503, // RAONIX
		"108b6a": 6504, // Antailiye
		"108ccf": 2,    // Cisco
		"108eba": 6505, // Molekule
		"108ee0": 152,  // Samsung
		"108ffe": 3335, // Huawei
		"1091a8": 6377, // Espressif
		"109266": 152,  // Samsung
		"109397": 125,  // ARRIS
		"1093e9": 545,  // Apple
		"1094bb": 545,  // Apple
		"10954b": 6506, // Megabyte
		"109693": 5439, // Amazon
		"1097bd": 6377, // Espressif
		"109836": 102,  // Dell
		"1098c3": 2057, // Murata
		"109ab9": 6507, // Tosibox
		"109add": 545,  // Apple
		"109d7a": 3335, // Huawei
		"109fa9": 2247, // Actiontec
		"10a24e": 6508, // GOLD3LINK
		"10a4da": 3335, // Huawei
		"10a51d": 421,  // Intel
		"10a5d0": 2057, // Murata
		"10ae60": 67,   // Private
		"10aea5": 6509, // Duskrise
		"10b1f8": 3335, // Huawei
		"10b26b": 6510, // base
		"10b36f": 6511, // Bowei
		"10b3c6": 2,    // Cisco
		"10b3d5": 2,    // Cisco
		"10b3d6": 2,    // Cisco
		"10b713": 67,   // Private
		"10b7a8": 6512, // CableFree
		"10b7f6": 6513, // Plastoform
		"10b9c4": 545,  // Apple
		"10b9f7": 6514, // Niko-Servodan
		"10bc97": 6370, // Vivo
		"10bd18": 2,    // Cisco
		"10bd55": 6515, // Q-Lab
		"10bef5": 803,  // D-Link
		"10bf48": 1806, // ASUS
		"10c172": 3335, // Huawei
		"10c25a": 6393, // Technicolor
		"10c2ba": 6516, // UTT
		"10c37b": 1806, // ASUS
		"10c3ab": 3335, // Huawei
		"10c595": 2665, // Lenovo
		"10c61f": 3335, // Huawei
		"10c65e": 6517, // Adapt-IP
		"10c6fc": 797,  // Garmin
		"10c9ca": 6518, // Ace
		"10ca81": 6519, // Precia
		"10cc1b": 6520, // Liverock
		"10cd6e": 6521, // Fisys
		"10cdae": 620,  // Avaya
		"10cdb6": 5759, // Essential
		"10ce02": 5439, // Amazon
		"10ce45": 6522, // Miromico
		"10cee9": 545,  // Apple
		"10d07a": 4499, // AMPAK
		"10d0ab": 3032, // ZTE
		"10d38a": 152,  // Samsung
		"10d542": 152,  // Samsung
		"10d7b0": 2049, // Sagemcom
		"10da43": 1368, // Netgear
		"10dc4a": 4922, // Fiberhome
		"10ddb1": 545,  // Apple
		"10ddf4": 6523, // Maxway
		"10dee4": 6524, // automationNEXT
		"10dffc": 300,  // Siemens
		"10e4af": 6525, // APR
		"10e4c2": 152,  // Samsung
		"10e68f": 6526, // Kwangsung
		"10e6ae": 6527, // Source
		"10e77a": 6528, // STMicrolectronics
		"10e7c6": 302,  // HP
		"10e878": 457,  // Nokia
		"10e8a7": 1592, // Wistron
		"10e8ee": 6529, // PhaseSpace
		"10e953": 3335, // Huawei
		"10ec81": 152,  // Samsung
		"10f005": 421,  // Intel
		"10f163": 6530, // TNK
		"10f1f2": 869,  // LG
		"10f311": 2,    // Cisco
		"10f3db": 6531, // Gridco
		"10f681": 6370, // Vivo
		"10f920": 2,    // Cisco
		"10f96f": 869,  // LG
		"10f9ee": 457,  // Nokia
		"10face": 6532, // Reacheng
		"10fbf0": 6533, // KangSheng
		"10fc54": 6534, // Shany
		"10fcb6": 6535, // mirusystems
		"10feed": 1595, // TP-Link
		"1100aa": 67,   // Private
		"111111": 67,   // Private
		"140020": 6536, // LongSung
		"14007d": 3032, // ZTE
		"1400e9": 1217, // Mitel
		"140152": 152,  // Samsung
		"14019c": 6537, // Ubyon
		"1402ec": 302,  // HP
		"140467": 6538, // SNK
		"140708": 67,   // Private
		"1407e0": 6539, // Abrantix
		"1409b4": 3032, // ZTE
		"1409dc": 3335, // Huawei
		"140a29": 6540, // Tiinlab
		"140ac5": 5439, // Amazon
		"140c5b": 6541, // PLNetworks
		"140d4f": 833,  // Flextronics
		"140f42": 457,  // Nokia
		"14109f": 545,  // Apple
		"141114": 6266, // Tecno
		"141333": 3005, // Azurewave
		"141357": 6542, // ATP
		"1413fb": 3335, // Huawei
		"14144b": 5441, // Ruijie
		"14169d": 2,    // Cisco
		"14169e": 4031, // Wingtech
		"14172a": 4922, // Fiberhome
		"141877": 102,  // Dell
		"1418c3": 421,  // Intel
		"141aa3": 687,  // Motorola
		"141bbd": 6543, // Volex
		"141bf0": 6544, // Intellimedia
		"141f78": 152,  // Samsung
		"14205e": 545,  // Apple
		"142233": 4922, // Fiberhome
		"14223b": 3522, // Google
		"1422db": 5821, // eero
		"14230a": 3335, // Huawei
		"1423d7": 6545, // Eutronix
		"142475": 6546, // 4DReplay
		"142882": 6547, // Midicom
		"142971": 6548, // Nemoa
		"142bd2": 6549, // Armtel
		"142d8b": 6550, // Incipio
		"142df5": 6551, // Amphitech
		"142e5e": 2075, // Sercomm
		"143004": 3335, // Huawei
		"14307a": 6552, // Avermetrics
		"1430c6": 687,  // Motorola
		"1432d1": 152,  // Samsung
		"143365": 4540, // TEM
		"14358b": 6553, // Mediabridge
		"143605": 457,  // Nokia
		"1436c6": 2665, // Lenovo
		"14373b": 6554, // PROCOM
		"143aea": 6555, // Dynapower
		"143cc3": 3335, // Huawei
		"143e60": 457,  // Nokia
		"143ebf": 3032, // ZTE
		"143f27": 6556, // Noccela
		"143fa6": 101,  // Sony
		"143fc3": 6557, // SnapAV
		"144146": 934,  // Honeywell
		"1441e2": 6558, // Monaco
		"144658": 3335, // Huawei
		"1446e4": 6559, // Avistel
		"144920": 3335, // Huawei
		"1449bc": 3923, // DrayTek
		"1449e0": 152,  // Samsung
		"144c1a": 6560, // Max
		"144d67": 2128, // Zioncom
		"144e2a": 4565, // Ciena
		"144f8a": 421,  // Intel
		"145051": 3206, // Sharp
		"145120": 3335, // Huawei
		"145412": 6561, // Entis
		"145594": 3335, // Huawei
		"14563a": 3335, // Huawei
		"145645": 6562, // Savitech
		"14568e": 152,  // Samsung
		"14579f": 3335, // Huawei
		"1458d0": 302,  // HP
		"1459c0": 1368, // Netgear
		"145a05": 545,  // Apple
		"145a83": 6563, // Logi-D
		"145afc": 2874, // Liteon
		"145bd1": 125,  // ARRIS
		"145be1": 6564, // nyantec
		"145f94": 3335, // Huawei
		"146080": 3032, // ZTE
		"1460cb": 545,  // Apple
		"14612f": 620,  // Avaya
		"146a0b": 4781, // Cypress
		"146b9a": 3032, // ZTE
		"146e0a": 67,   // Private
		"147411": 4556, // RIM
		"14755b": 421,  // Intel
		"147590": 1595, // TP-Link
		"147740": 3335, // Huawei
		"147bac": 457,  // Nokia
		"147dc5": 2057, // Murata
		"147dda": 545,  // Apple
		"14857f": 421,  // Intel
		"148692": 1595, // TP-Link
		"14876a": 545,  // Apple
		"1488e6": 545,  // Apple
		"148919": 6565, // 2bps
		"14893e": 6566, // Vixtel
		"1489cb": 3335, // Huawei
		"1489fd": 152,  // Samsung
		"148a70": 5281, // ADS
		"148c4a": 3335, // Huawei
		"148f21": 797,  // Garmin
		"148fc6": 545,  // Apple
		"149090": 6567, // KongTop
		"149138": 5439, // Amazon
		"149182": 2469, // Belkin
		"14942f": 6568, // Usys
		"1495ce": 545,  // Apple
		"1496e5": 152,  // Samsung
		"149877": 545,  // Apple
		"14987d": 6393, // Technicolor
		"1499e2": 545,  // Apple
		"149a10": 612,  // Microsoft
		"149d09": 3335, // Huawei
		"149d99": 545,  // Apple
		"149ecf": 102,  // Dell
		"149f3c": 152,  // Samsung
		"149fe8": 2665, // Lenovo
		"14a0f8": 3335, // Huawei
		"14a2a0": 2,    // Cisco
		"14a32f": 3335, // Huawei
		"14a364": 152,  // Samsung
		"14a3b4": 3335, // Huawei
		"14a51a": 3335, // Huawei
		"14a72b": 6569, // currentoptronics
		"14a9d0": 289,  // F5
		"14a9e3": 6570, // MST
		"14ab02": 3335, // Huawei
		"14abc5": 421,  // Intel
		"14abf0": 125,  // ARRIS
		"14b126": 6571, // Industrial
		"14b1c8": 6572, // InfiniWing
		"14b31f": 102,  // Dell
		"14b457": 1655, // Silicon
		"14b484": 152,  // Samsung
		"14b73d": 6573, // ARCHEAN
		"14b7f8": 6393, // Technicolor
		"14b968": 3335, // Huawei
		"14bb6e": 152,  // Samsung
		"14bd61": 545,  // Apple
		"14c03e": 125,  // ARRIS
		"14c0a1": 6574, // UCloud
		"14c126": 457,  // Nokia
		"14c14e": 3522, // Google
		"14c213": 545,  // Apple
		"14c21d": 6575, // Sabtech
		"14c88b": 545,  // Apple
		"14c913": 869,  // LG
		"14c9cf": 6576, // Sigmastar
		"14caa0": 6577, // Hu
		"14cb19": 302,  // HP
		"14cb65": 612,  // Microsoft
		"14cc20": 1595, // TP-Link
		"14cf8d": 6578, // Ohsung
		"14cf92": 1595, // TP-Link
		"14cfe2": 125,  // ARRIS
		"14d00d": 545,  // Apple
		"14d11f": 3335, // Huawei
		"14d169": 3335, // Huawei
		"14d19e": 545,  // Apple
		"14d4fe": 125,  // ARRIS
		"14d64d": 803,  // D-Link
		"14d76e": 6579, // CoNCH
		"14dae9": 1806, // ASUS
		"14dd9c": 6370, // Vivo
		"14dda9": 1806, // ASUS
		"14dde5": 6580, // Mpmkvvcl
		"14de39": 3335, // Huawei
		"14e4ec": 6581, // mLogic
		"14e6e4": 1595, // TP-Link
		"14e9b2": 4922, // Fiberhome
		"14eb08": 3335, // Huawei
		"14eb33": 6582, // BSMediasoft
		"14ebb6": 1595, // TP-Link
		"14edbb": 1939, // 2Wire
		"14ede4": 6583, // Kaiam
		"14ee9d": 6584, // AirNav
		"14efcf": 6585, // Schreder
		"14f0c5": 6586, // Xtremio
		"14f42a": 152,  // Samsung
		"14f65a": 5698, // Xiaomi
		"14f6d8": 421,  // Intel
		"14fb70": 3335, // Huawei
		"14feaf": 6587, // Sagittar
		"14feb5": 102,  // Dell
		"18002d": 101,  // Sony
		"1800db": 6588, // Fitbit
		"1801f1": 5698, // Xiaomi
		"18022d": 3335, // Huawei
		"1802ae": 6370, // Vivo
		"180373": 102,  // Dell
		"1806ff": 143,  // Acer
		"180b52": 6589, // Nanotron
		"180c14": 6590, // iSonea
		"180cac": 87,   // Canon
		"180d2c": 3542, // Intelbras
		"180f76": 803,  // D-Link
		"18104e": 6591, // Cedint-UPM
		"181212": 6592, // Cepton
		"18132d": 3032, // ZTE
		"181456": 457,  // Nokia
		"1816c9": 152,  // Samsung
		"181714": 6593, // Daewoois
		"181725": 3384, // Cameo
		"18193f": 6594, // Tamtron
		"1819d6": 152,  // Samsung
		"181beb": 2247, // Actiontec
		"181dea": 421,  // Intel
		"181e78": 2049, // Sagemcom
		"181e95": 6595, // AuVerte
		"181eb0": 152,  // Samsung
		"182032": 545,  // Apple
		"18204c": 6596, // Kummler+Matter
		"1820a6": 619,  // Sage
		"1820d5": 125,  // ARRIS
		"182195": 152,  // Samsung
		"18227e": 152,  // Samsung
		"182649": 421,  // Intel
		"182666": 152,  // Samsung
		"182a44": 6597, // Hirose
		"182a57": 3335, // Huawei
		"182a7b": 1427, // Nintendo
		"182ad3": 826,  // Juniper
		"182b05": 6598, // 8D
		"182cb4": 6599, // Nectarsoft
		"182df7": 6600, // JY
		"183009": 6601, // Woojin
		"1831bf": 1806, // ASUS
		"1832a2": 6602, // Laon
		"18339d": 2,    // Cisco
		"183451": 545,  // Apple
		"1835d1": 125,  // ARRIS
		"1836fc": 6603, // Elecsys
		"183864": 6604, // CAP-Tech
		"183919": 6605, // Unicoi
		"18399c": 6606, // Skorpios
		"183a2d": 152,  // Samsung
		"183a48": 6607, // VostroNet
		"183cb7": 3335, // Huawei
		"183d5e": 3335, // Huawei
		"183da2": 421,  // Intel
		"183eef": 545,  // Apple
		"183f47": 152,  // Samsung
		"18421d": 67,   // Private
		"184462": 6608, // Riava
		"1844e6": 3032, // ZTE
		"184617": 152,  // Samsung
		"184859": 3798, // Castlenet
		"1848be": 5439, // Amazon
		"1848ca": 2057, // Murata
		"1848d8": 6609, // Fastback
		"184b0d": 2738, // Ruckus
		"184bdf": 6610, // Caavo
		"184cae": 6611, // Continental
		"184e16": 152,  // Samsung
		"184e94": 6612, // Messoa
		"184ecb": 152,  // Samsung
		"184f5d": 6613, // JRC
		"18502a": 6614, // Soarnex
		"185253": 635,  // Pixord
		"185282": 4922, // Fiberhome
		"185345": 457,  // Nokia
		"1854cf": 152,  // Samsung
		"1855e3": 545,  // Apple
		"185644": 3335, // Huawei
		"185680": 421,  // Intel
		"1856c3": 545,  // Apple
		"185869": 6615, // Sailer
		"185880": 2640, // Arcadyan
		"185936": 5698, // Xiaomi
		"1859f5": 2,    // Cisco
		"185a58": 102,  // Dell
		"185ae8": 6616, // Zenotech
		"185b00": 457,  // Nokia
		"185bb3": 152,  // Samsung
		"185d9a": 6617, // BobjGear
		"185e0b": 3032, // ZTE
		"185e0f": 421,  // Intel
		"186024": 302,  // HP
		"1861c7": 6618, // lemonbeat
		"18622c": 2049, // Sagemcom
		"186472": 1685, // Aruba
		"186590": 545,  // Apple
		"1866da": 102,  // Dell
		"1866e3": 6619, // Veros
		"1867b0": 152,  // Samsung
		"18686a": 3032, // ZTE
		"1869d4": 152,  // Samsung
		"186d99": 6620, // Adanis
		"18703b": 3335, // Huawei
		"18742e": 5439, // Amazon
		"187758": 6621, // Audoo
		"1878d4": 6622, // Verizon
		"187a3b": 1685, // Aruba
		"187a93": 6623, // AMICCOM
		"187c0b": 2738, // Ruckus
		"187eb9": 545,  // Apple
		"188090": 2,    // Cisco
		"1880ce": 6624, // Barberry
		"18810e": 545,  // Apple
		"18828c": 2640, // Arcadyan
		"188331": 152,  // Samsung
		"1883bf": 2640, // Arcadyan
		"188410": 6625, // CoreTrust
		"1886ac": 457,  // Nokia
		"188740": 5698, // Xiaomi
		"188796": 1341, // HTC
		"18895b": 152,  // Samsung
		"1889cf": 6266, // Tecno
		"1889df": 6626, // CerebrEX
		"188b45": 2,    // Cisco
		"188b9d": 2,    // Cisco
		"188ef9": 6627, // G2C
		"189088": 5821, // eero
		"1890d8": 2049, // Sagemcom
		"1892a4": 4565, // Ciena
		"18937f": 4499, // AMPAK
		"189552": 6628, // 1MORE
		"189a67": 6629, // CSE-Servelec
		"189c27": 125,  // ARRIS
		"189c5d": 2,    // Cisco
		"189e2c": 3335, // Huawei
		"189efc": 545,  // Apple
		"18a28a": 6630, // Essel-T
		"18a3e8": 4922, // Fiberhome
		"18a4a9": 6631, // Vanu
		"18a6f7": 1595, // TP-Link
		"18a905": 302,  // HP
		"18a99b": 102,  // Dell
		"18a9a6": 6632, // Nebra
		"18aa0f": 3335, // Huawei
		"18aa45": 6633, // Fon
		"18ab1d": 152,  // Samsung
		"18abf5": 3430, // Ultra
		"18ac9e": 6282, // Itel
		"18ad4d": 6634, // Polostar
		"18aebb": 300,  // Siemens
		"18af61": 545,  // Apple
		"18af8f": 545,  // Apple
		"18b169": 1003, // Sonicwall
		"18b3ba": 6635, // Netlogic
		"18b430": 6636, // Nest
		"18b591": 6637, // I-Storm
		"18b6cc": 6291, // We
		"18b79e": 6638, // Invoxia
		"18b81f": 125,  // ARRIS
		"18bb26": 6441, // FN-Link
		"18bb41": 3335, // Huawei
		"18bdad": 6639, // L-Tech
		"18be92": 957,  // Delta
		"18bfb3": 152,  // Samsung
		"18c04d": 1929, // Giga-Byte
		"18c086": 858,  // Broadcom
		"18c241": 6640, // SonicWall
		"18c2bf": 1077, // Buffalo
		"18c58a": 3335, // Huawei
		"18cc18": 421,  // Intel
		"18cc23": 6641, // Philio
		"18ce94": 152,  // Samsung
		"18cf24": 3335, // Huawei
		"18cf5e": 2874, // Liteon
		"18d071": 6068, // Dasan
		"18d225": 4922, // Fiberhome
		"18d276": 3335, // Huawei
		"18d5b6": 6642, // SMG
		"18d66a": 6643, // Inmarsat
		"18d6c7": 1595, // TP-Link
		"18d6cf": 6644, // Kurth
		"18d949": 6645, // Qvis
		"18d98f": 3335, // Huawei
		"18d9ef": 4938, // Shuttle
		"18dbf2": 102,  // Dell
		"18dc56": 3097, // Yulong
		"18ded7": 3335, // Huawei
		"18dfc1": 6646, // Aetheros
		"18e1ca": 6647, // wanze
		"18e215": 457,  // Nokia
		"18e29f": 6370, // Vivo
		"18e2c2": 152,  // Samsung
		"18e3bc": 6434, // TCT
		"18e728": 2,    // Cisco
		"18e777": 6370, // Vivo
		"18e7b0": 545,  // Apple
		"18e7f4": 545,  // Apple
		"18e80f": 6648, // Viking
		"18e829": 2979, // Ubiquiti
		"18e8dd": 6649, // Moduletek
		"18ece7": 1077, // Buffalo
		"18ee69": 545,  // Apple
		"18ef63": 2,    // Cisco
		"18f0e4": 5698, // Xiaomi
		"18f18e": 6650, // ChipER
		"18f1d8": 545,  // Apple
		"18f22c": 1595, // TP-Link
		"18f292": 6651, // Shannon
		"18f643": 545,  // Apple
		"18f87a": 6652, // i3
		"18f9c4": 1743, // BAE
		"18fb7b": 102,  // Dell
		"18fc26": 6653, // Qorvo
		"18fc9f": 6654, // Changhe
		"18fe34": 6377, // Espressif
		"18ff0f": 421,  // Intel
		"1c0042": 6655, // NARI
		"1c012d": 6656, // Ficer
		"1c0656": 6657, // IDY
		"1c08c1": 869,  // LG
		"1c1161": 4565, // Ciena
		"1c12b0": 5439, // Amazon
		"1c1338": 6658, // Kimball
		"1c1386": 3335, // Huawei
		"1c1448": 125,  // ARRIS
		"1c14b3": 6659, // Airwire
		"1c151f": 3335, // Huawei
		"1c17d3": 2,    // Cisco
		"1c19de": 6660, // eyevis
		"1c1ac0": 545,  // Apple
		"1c1adf": 612,  // Microsoft
		"1c1b0d": 1929, // Giga-Byte
		"1c1b68": 125,  // ARRIS
		"1c1bb5": 421,  // Intel
		"1c1d67": 3335, // Huawei
		"1c1d86": 2,    // Cisco
		"1c1fd4": 6661, // LifeBEAM
		"1c1ff1": 3335, // Huawei
		"1c20db": 3335, // Huawei
		"1c232c": 152,  // Samsung
		"1c234f": 6662, // EDMI
		"1c24cd": 2545, // Askey
		"1c24eb": 6663, // Burlywood
		"1c2704": 3032, // ZTE
		"1c28af": 1685, // Aruba
		"1c2a8b": 457,  // Nokia
		"1c330e": 6664, // PernixData
		"1c34da": 432,  // Mellanox
		"1c34f1": 1655, // Silicon
		"1c36bb": 545,  // Apple
		"1c37bf": 6665, // Cloudium
		"1c3929": 6578, // Ohsung
		"1c3947": 358,  // Compal
		"1c398a": 4922, // Fiberhome
		"1c3a4f": 6666, // AccuSpec
		"1c3a60": 2738, // Ruckus
		"1c3ade": 152,  // Samsung
		"1c3b8f": 6667, // Selve
		"1c3bf3": 1595, // TP-Link
		"1c3cd4": 3335, // Huawei
		"1c3d2f": 3335, // Huawei
		"1c4024": 102,  // Dell
		"1c4190": 3858, // Universal
		"1c4363": 3335, // Huawei
		"1c4419": 1595, // TP-Link
		"1c4586": 1427, // Nintendo
		"1c472f": 3335, // Huawei
		"1c497b": 1450, // Gemtek
		"1c4af7": 6668, // Amon
		"1c4bb9": 6642, // SMG
		"1c4bd6": 3005, // Azurewave
		"1c4c48": 6282, // Itel
		"1c4d66": 5439, // Amazon
		"1c4d70": 421,  // Intel
		"1c501e": 2427, // Sunplus
		"1c51b5": 6669, // Techaya
		"1c53f9": 3522, // Google
		"1c549e": 3858, // Universal
		"1c553a": 6670, // QianGua
		"1c568e": 2128, // Zioncom
		"1c56fe": 687,  // Motorola
		"1c57d8": 6671, // Kraftway
		"1c57dc": 545,  // Apple
		"1c599b": 3335, // Huawei
		"1c5a0b": 6672, // Tegile
		"1c5a3e": 152,  // Samsung
		"1c5a6b": 796,  // Philips
		"1c5cf2": 545,  // Apple
		"1c5f2b": 803,  // D-Link
		"1c60d2": 4922, // Fiberhome
		"1c60de": 4254, // Mercury
		"1c61b4": 1595, // TP-Link
		"1c62b8": 152,  // Samsung
		"1c6499": 3870, // Comtrend
		"1c659d": 2874, // Liteon
		"1c66aa": 152,  // Samsung
		"1c6758": 3335, // Huawei
		"1c697a": 6673, // EliteGroup
		"1c69a5": 2221, // BlackBerry
		"1c6a7a": 2,    // Cisco
		"1c6bca": 6674, // Mitsunami
		"1c6e4c": 6675, // Logistic
		"1c6e76": 6676, // Quarion
		"1c6ee6": 6677, // Nhnetworks
		"1c6f65": 1929, // Giga-Byte
		"1c7022": 2057, // Murata
		"1c7125": 545,  // Apple
		"1c721d": 102,  // Dell
		"1c7370": 6678, // Neotech
		"1c73e2": 3335, // Huawei
		"1c740d": 2693, // ZyXEL
		"1c7508": 358,  // Compal
		"1c76ca": 6679, // Terasic
		"1c76f2": 152,  // Samsung
		"1c7b21": 101,  // Sony
		"1c7c11": 6680, // EID
		"1c7cc7": 6681, // Coriant
		"1c7e51": 6682, // 3bumen.com
		"1c7ee5": 803,  // D-Link
		"1c7f2c": 3335, // Huawei
		"1c869a": 152,  // Samsung
		"1c86ad": 6683, // MCT
		"1c872c": 1806, // ASUS
		"1c87e3": 6266, // Tecno
		"1c8e5c": 3335, // Huawei
		"1c8e8e": 6684, // DB
		"1c90be": 306,  // Ericsson
		"1c9148": 545,  // Apple
		"1c9180": 545,  // Apple
		"1c937c": 125,  // ARRIS
		"1c93c4": 5439, // Amazon
		"1c955d": 6685, // I-LAX
		"1c959f": 6686, // Veethree
		"1c97c5": 6687, // Ynomia
		"1c98ec": 302,  // HP
		"1c994c": 2057, // Murata
		"1c9957": 421,  // Intel
		"1c9c26": 6688, // Zoovel
		"1c9c8c": 826,  // Juniper
		"1c9d72": 6393, // Technicolor
		"1c9dc2": 6377, // Espressif
		"1c9e46": 545,  // Apple
		"1c9ecc": 6393, // Technicolor
		"1ca681": 3335, // Huawei
		"1ca852": 6689, // Sensaio
		"1caa07": 2,    // Cisco
		"1cab01": 6690, // Innovolt
		"1caba7": 545,  // Apple
		"1cabc0": 870,  // Hitron
		"1cadd1": 6691, // Bosung
		"1caecb": 3335, // Huawei
		"1caf05": 152,  // Samsung
		"1caff7": 803,  // D-Link
		"1cb044": 2545, // Askey
		"1cb094": 1341, // HTC
		"1cb243": 6692, // TDC
		"1cb3c9": 545,  // Apple
		"1cb72c": 1806, // ASUS
		"1cb796": 3335, // Huawei
		"1cb857": 6693, // Becon
		"1cb9c4": 2738, // Ruckus
		"1cbd0e": 6694, // Amplified
		"1cbdb9": 803,  // D-Link
		"1cc035": 4479, // Planex
		"1cc10c": 421,  // Intel
		"1cc11a": 6695, // Wavetronix
		"1cc1de": 302,  // HP
		"1cc316": 6696, // MileSight
		"1cc63c": 2640, // Arcadyan
		"1ccb99": 6434, // TCT
		"1cccd6": 5698, // Xiaomi
		"1cd1ba": 4922, // Fiberhome
		"1cd1e0": 2,    // Cisco
		"1cd6be": 1592, // Wistron
		"1cda27": 6370, // Vivo
		"1cde57": 4922, // Fiberhome
		"1cdea7": 2,    // Cisco
		"1cdf0f": 2,    // Cisco
		"1ce165": 6697, // Marshal
		"1ce192": 551,  // Qisda
		"1ce504": 3335, // Huawei
		"1ce57f": 152,  // Samsung
		"1ce61d": 152,  // Samsung
		"1ce62b": 545,  // Apple
		"1ce639": 3335, // Huawei
		"1ce6ad": 3335, // Huawei
		"1ce6c7": 2,    // Cisco
		"1ce85d": 2,    // Cisco
		"1cea0b": 6295, // Edgecore
		"1cea1b": 457,  // Nokia
		"1cec72": 1111, // Allradio
		"1cefce": 6698, // bebro
		"1cf03e": 6699, // Wearhaus
		"1cf061": 6700, // SCAPS
		"1cf29a": 3522, // Google
		"1cf42b": 3335, // Huawei
		"1cf4ca": 67,   // Private
		"1cf5e7": 6701, // Turtle
		"1cfa68": 1595, // TP-Link
		"1cfe2b": 5439, // Amazon
		"20014f": 6702, // Linea
		"2002af": 2057, // Murata
		"20040f": 102,  // Dell
		"200505": 6703, // Radmax
		"2008ed": 3335, // Huawei
		"200bc7": 3335, // Huawei
		"200bcf": 1427, // Nintendo
		"200cc8": 1368, // Netgear
		"200f70": 6704, // Foxtech
		"20107a": 1450, // Gemtek
		"2013e0": 152,  // Samsung
		"2016b9": 421,  // Intel
		"2016d8": 2874, // Liteon
		"201742": 869,  // LG
		"201a06": 358,  // Compal
		"201bc9": 826,  // Juniper
		"201d03": 6705, // Elatec
		"201e88": 421,  // Intel
		"201f3b": 3522, // Google
		"201f54": 2051, // Raisecom
		"2021a5": 869,  // LG
		"202564": 5440, // Pegatron
		"202598": 6706, // Teleview
		"2025d2": 4922, // Fiberhome
		"202681": 6266, // Tecno
		"20283e": 3335, // Huawei
		"2028bc": 6707, // Visionscape
		"202ac5": 6708, // Petite-En
		"202bc1": 3335, // Huawei
		"202d07": 152,  // Samsung
		"202d23": 6709, // Collinear
		"20311c": 6370, // Vivo
		"2031eb": 6710, // Hdsn
		"20326c": 152,  // Samsung
		"2032c6": 545,  // Apple
		"2034fb": 5698, // Xiaomi
		"20365b": 6711, // Megafone
		"203706": 2,    // Cisco
		"2037a5": 545,  // Apple
		"2037bc": 6712, // Kuipers
		"203a07": 2,    // Cisco
		"203aef": 6713, // Sivantos
		"203b69": 6370, // Vivo
		"203cae": 545,  // Apple
		"203d66": 125,  // ARRIS
		"203db2": 3335, // Huawei
		"203dbd": 869,  // LG
		"204005": 6714, // feno
		"20443a": 56,   // Schneider
		"2046a1": 6715, // VECOW
		"204747": 102,  // Dell
		"2047da": 5698, // Xiaomi
		"2047ed": 3513, // BSkyB
		"204b22": 5436, // Sunnovo
		"204c03": 1685, // Aruba
		"204c9e": 2,    // Cisco
		"204e6b": 6716, // Axxana
		"204e71": 826,  // Juniper
		"204e7f": 1368, // Netgear
		"204ef6": 3005, // Azurewave
		"2050e7": 4499, // AMPAK
		"205383": 3335, // Huawei
		"2053ca": 6717, // Risk
		"205476": 101,  // Sony
		"2054fa": 3335, // Huawei
		"205531": 152,  // Samsung
		"205532": 6718, // Gotech
		"205721": 5261, // Salix
		"205869": 2738, // Ruckus
		"2059a0": 352,  // Paragon
		"205a00": 6719, // Coval
		"205b2a": 67,   // Private
		"205d47": 6370, // Vivo
		"205e64": 3335, // Huawei
		"205ef7": 152,  // Samsung
		"205f3d": 1640, // Cambridge
		"206274": 612,  // Microsoft
		"20635f": 6720, // Abeeway
		"206432": 152,  // Samsung
		"20658e": 3335, // Huawei
		"2066fd": 6721, // CoNSTELL8
		"20677c": 302,  // HP
		"2067b1": 6098, // Pluto
		"20689d": 2874, // Liteon
		"206980": 545,  // Apple
		"206a8a": 1592, // Wistron
		"206a94": 870,  // Hitron
		"206be7": 1595, // TP-Link
		"206c8a": 185,  // Extreme
		"206d31": 6722, // Firewalla
		"206e9c": 152,  // Samsung
		"20719e": 6395, // SF
		"207355": 125,  // ARRIS
		"207454": 6370, // Vivo
		"207600": 2247, // Actiontec
		"20768f": 545,  // Apple
		"207693": 2665, // Lenovo
		"2078cd": 545,  // Apple
		"2078f0": 545,  // Apple
		"207918": 421,  // Intel
		"207bd2": 6723, // ASIX
		"207c14": 6724, // Qotom
		"207d74": 545,  // Apple
		"208058": 4565, // Ciena
		"2082c0": 5698, // Xiaomi
		"20858c": 6725, // Assa
		"208756": 300,  // Siemens
		"2087ec": 3335, // Huawei
		"20896f": 4922, // Fiberhome
		"208984": 358,  // Compal
		"208986": 3032, // ZTE
		"208c47": 6726, // Tenstorrent
		"208c86": 3335, // Huawei
		"20918a": 6727, // Profalux
		"2091d9": 6728, // I'M
		"209a7d": 2049, // Sagemcom
		"209ae9": 6729, // Volacomm
		"209bcd": 545,  // Apple
		"209e79": 3858, // Universal
		"209ef7": 185,  // Extreme
		"20a171": 5439, // Amazon
		"20a2e4": 545,  // Apple
		"20a2e7": 6730, // Lee-Dickens
		"20a60c": 5698, // Xiaomi
		"20a680": 3335, // Huawei
		"20a6cd": 302,  // HP
		"20a783": 6731, // miControl
		"20a8b9": 300,  // Siemens
		"20a90e": 6434, // TCT
		"20a99b": 612,  // Microsoft
		"20aa25": 6732, // IP-NET
		"20aa4b": 1783, // Linksys
		"20ab37": 545,  // Apple
		"20ab48": 3335, // Huawei
		"20b001": 6393, // Technicolor
		"20b0f7": 6733, // Enclustra
		"20b399": 312,  // Enterasys
		"20b5c6": 6734, // Mimosa
		"20b730": 6735, // TeconGroup
		"20b780": 35,   // Toshiba
		"20b7c0": 6736, // OMICRON
		"20b868": 687,  // Motorola
		"20bbc0": 2,    // Cisco
		"20becd": 5821, // eero
		"20bfdb": 6737, // DVL
		"20c047": 6622, // Verizon
		"20c19b": 421,  // Intel
		"20c3a4": 6738, // RetailNext
		"20c6eb": 2154, // Panasonic
		"20c74f": 6739, // SensorPush
		"20c9d0": 545,  // Apple
		"20cec4": 6740, // Peraso
		"20cf30": 1806, // ASUS
		"20cfae": 2,    // Cisco
		"20d160": 67,   // Private
		"20d21f": 6741, // Wincal
		"20d25f": 6742, // SmartCap
		"20d276": 6282, // Itel
		"20d390": 152,  // Samsung
		"20d5bf": 152,  // Samsung
		"20d607": 457,  // Nokia
		"20d75a": 6743, // Posh
		"20d80b": 826,  // Juniper
		"20d906": 6744, // Iota
		"20da22": 3335, // Huawei
		"20dbab": 152,  // Samsung
		"20dce6": 1595, // TP-Link
		"20dcfd": 3335, // Huawei
		"20df73": 3335, // Huawei
		"20dfb9": 3522, // Google
		"20e09c": 457,  // Nokia
		"20e2a8": 545,  // Apple
		"20e52a": 1368, // Netgear
		"20e564": 125,  // ARRIS
		"20e791": 300,  // Siemens
		"20e7b6": 3858, // Universal
		"20e874": 545,  // Apple
		"20e882": 3032, // ZTE
		"20ed74": 6745, // Ability
		"20ee28": 545,  // Apple
		"20efbd": 1916, // Roku
		"20f17c": 3335, // Huawei
		"20f19e": 125,  // ARRIS
		"20f375": 125,  // ARRIS
		"20f3a3": 3335, // Huawei
		"20f44f": 457,  // Nokia
		"20f478": 5698, // Xiaomi
		"20f77c": 6370, // Vivo
		"20f85e": 957,  // Delta
		"20fdf1": 160,  // 3Com
		"20fe00": 5439, // Amazon
		"20ff36": 6746, // Iflytek
		"2400ba": 3335, // Huawei
		"24016f": 3335, // Huawei
		"2401c7": 2,    // Cisco
		"24050f": 6747, // MTN
		"240588": 3522, // Google
		"240917": 6748, // Devlin
		"240995": 3335, // Huawei
		"240a11": 6434, // TCT
		"240a63": 125,  // ARRIS
		"240a64": 3005, // Azurewave
		"240ac4": 6377, // Espressif
		"240d6c": 6749, // Smnd
		"240dc2": 6434, // TCT
		"241125": 6750, // Hutek
		"241145": 5698, // Xiaomi
		"241148": 6751, // Entropix
		"24166d": 3335, // Huawei
		"24169d": 2,    // Cisco
		"24181d": 152,  // Samsung
		"241a8c": 6752, // Squarehead
		"241ae6": 3335, // Huawei
		"241b7a": 545,  // Apple
		"241eeb": 545,  // Apple
		"241f2c": 6753, // Calsys
		"241fa0": 3335, // Huawei
		"2420c7": 2049, // Sagemcom
		"242124": 457,  // Nokia
		"2421ab": 101,  // Sony
		"24240e": 545,  // Apple
		"242642": 3206, // Sharp
		"2429fe": 2849, // Kyocera
		"242e02": 3335, // Huawei
		"242ffa": 35,   // Toshiba
		"2430f8": 3335, // Huawei
		"243154": 3335, // Huawei
		"243184": 3206, // Sharp
		"24336c": 67,   // Private
		"2436da": 2,    // Cisco
		"2437ef": 4352, // EMC
		"243a82": 6754, // Irts
		"243faa": 3335, // Huawei
		"2440ae": 6755, // NIIC
		"24418c": 421,  // Intel
		"2442bc": 6756, // Alinco
		"244427": 3335, // Huawei
		"24456b": 3335, // Huawei
		"2446c8": 687,  // Motorola
		"2446e4": 3335, // Huawei
		"24470e": 6757, // PentronicAB
		"244b03": 152,  // Samsung
		"244b81": 152,  // Samsung
		"244bfe": 1806, // ASUS
		"244c07": 3335, // Huawei
		"244cab": 6377, // Espressif
		"244ce3": 5439, // Amazon
		"244f1d": 6758, // iRule
		"2453bf": 6759, // Enernet
		"24586e": 3032, // ZTE
		"245880": 6760, // Vizeo
		"245a4c": 2979, // Ubiquiti
		"245ab5": 152,  // Samsung
		"245b83": 5696, // Renesas
		"245ba7": 545,  // Apple
		"245bf0": 2874, // Liteon
		"245cbf": 6761, // Ncse
		"245e48": 545,  // Apple
		"245ebe": 6762, // QNAP
		"245f9f": 3335, // Huawei
		"245fdf": 2849, // Kyocera
		"246081": 6763, // razberi
		"2462ab": 6377, // Espressif
		"2462ce": 1685, // Aruba
		"24649f": 3335, // Huawei
		"246511": 621,  // AVM
		"246880": 6764, // Braveridge
		"2468b0": 152,  // Samsung
		"24693e": 6765, // innodisk
		"24694a": 521,  // Jasmine
		"246968": 1595, // TP-Link
		"2469a5": 3335, // Huawei
		"246aab": 6766, // IT-IS
		"246c8a": 6767, // YUKAI
		"246e96": 102,  // Dell
		"246f28": 6377, // Espressif
		"246f8c": 3335, // Huawei
		"247152": 102,  // Dell
		"247260": 6768, // IOTTECH
		"2474f7": 6242, // GoPro
		"247703": 421,  // Intel
		"24792a": 2738, // Ruckus
		"247e12": 2,    // Cisco
		"247e51": 3032, // ZTE
		"247f20": 2049, // Sagemcom
		"247f3c": 3335, // Huawei
		"248000": 6769, // Westcontrol
		"24813b": 2,    // Cisco
		"2481aa": 6770, // KSH
		"2481c7": 3335, // Huawei
		"24828a": 6771, // Prowave
		"2486f4": 6772, // Ctek
		"248707": 6773, // SEnergy
		"248a07": 432,  // Mellanox
		"2491bb": 3335, // Huawei
		"24920e": 152,  // Samsung
		"2494cb": 125,  // ARRIS
		"249504": 3188, // SFR
		"249745": 3335, // Huawei
		"249eab": 3335, // Huawei
		"24a074": 545,  // Apple
		"24a160": 6377, // Espressif
		"24a2e1": 545,  // Apple
		"24a43c": 2979, // Ubiquiti
		"24a487": 3335, // Huawei
		"24a52c": 3335, // Huawei
		"24a534": 6774, // SynTrust
		"24a65e": 3032, // ZTE
		"24a799": 3335, // Huawei
		"24a7dc": 3513, // BSkyB
		"24a87d": 2154, // Panasonic
		"24ab81": 545,  // Apple
		"24b209": 620,  // Avaya
		"24b2de": 6377, // Espressif
		"24b657": 2,    // Cisco
		"24b6b8": 6775, // Friem
		"24b6fd": 102,  // Dell
		"24b88c": 6776, // Crenus
		"24b8d2": 6777, // Opzoon
		"24bcf8": 3335, // Huawei
		"24be05": 302,  // HP
		"24be18": 6778, // Dadoutek
		"24bf74": 67,   // Private
		"24c0b3": 6779, // RSF
		"24c44a": 3032, // ZTE
		"24c613": 152,  // Samsung
		"24c696": 152,  // Samsung
		"24c9a1": 2738, // Ruckus
		"24c9de": 6780, // Genoray
		"24cacb": 4922, // Fiberhome
		"24cbe7": 6781, // MYK
		"24cd8d": 2057, // Murata
		"24ce33": 5439, // Amazon
		"24d0df": 545,  // Apple
		"24d13f": 6782, // Mexus
		"24d2cc": 6783, // SmartDrive
		"24d3f2": 3032, // ZTE
		"24d76b": 6784, // Syntronic
		"24d7eb": 6377, // Espressif
		"24d81e": 6785, // MirWifi
		"24d921": 620,  // Avaya
		"24da33": 3335, // Huawei
		"24da9b": 687,  // Motorola
		"24dbac": 3335, // Huawei
		"24dbed": 152,  // Samsung
		"24dec6": 1685, // Aruba
		"24df6a": 3335, // Huawei
		"24e314": 545,  // Apple
		"24e4c8": 4922, // Fiberhome
		"24e853": 869,  // LG
		"24e927": 2715, // TomTom
		"24e9b3": 2,    // Cisco
		"24e9ca": 3335, // Huawei
		"24ea40": 6786, // Helmholz
		"24ec99": 2545, // Askey
		"24edfd": 300,  // Siemens
		"24ee9a": 421,  // Intel
		"24f094": 545,  // Apple
		"24f0ff": 6787, // GHT
		"24f128": 6788, // Telstra
		"24f27f": 302,  // HP
		"24f57e": 6789, // HWH
		"24f5a2": 2469, // Belkin
		"24f5aa": 152,  // Samsung
		"24f603": 3335, // Huawei
		"24f677": 545,  // Apple
		"24fb65": 3335, // Huawei
		"24fc4e": 826,  // Juniper
		"24fce5": 152,  // Samsung
		"24fd0d": 3542, // Intelbras
		"24fd52": 2874, // Liteon
		"24fd5b": 6790, // SmartThings
		"280244": 545,  // Apple
		"2802d8": 152,  // Samsung
		"2804e0": 6791, // Fermax
		"28052e": 6792, // Dematic
		"28068d": 6793, // ITL
		"280b5c": 545,  // Apple
		"280dfc": 101,  // Sony
		"280feb": 869,  // LG
		"28101b": 6794, // MagnaCom
		"28107b": 803,  // D-Link
		"2811a5": 1819, // Bose
		"2811a8": 421,  // Intel
		"2811ec": 3335, // Huawei
		"281471": 6795, // Lantis
		"28162e": 1939, // 2Wire
		"28167f": 5698, // Xiaomi
		"2816a8": 612,  // Microsoft
		"2816ad": 421,  // Intel
		"281709": 3335, // Huawei
		"2817ce": 6796, // Omnisense
		"281878": 612,  // Microsoft
		"281b04": 6797, // Zalliant
		"282373": 6798, // Digita
		"2824ff": 1592, // Wistron
		"2826a6": 6799, // PBR
		"2827bf": 152,  // Samsung
		"28285d": 2693, // ZyXEL
		"2829cc": 6800, // Corsa
		"2829d9": 6801, // GlobalBeiMing
		"282a87": 6282, // Itel
		"282b96": 3335, // Huawei
		"282cb2": 1595, // TP-Link
		"282d06": 4499, // AMPAK
		"2830ac": 6802, // Frontiir
		"283152": 3335, // Huawei
		"283166": 6370, // Vivo
		"2832c5": 531,  // HUMAX
		"283334": 3335, // Huawei
		"2834a2": 2,    // Cisco
		"283737": 545,  // Apple
		"28385c": 833,  // Flextronics
		"2838cf": 6803, // Gen2wave
		"283926": 189,  // CyberTAN
		"28395e": 152,  // Samsung
		"2839e7": 6804, // Preceno
		"283b82": 803,  // D-Link
		"283ce4": 3335, // Huawei
		"283dc2": 152,  // Samsung
		"283e76": 6805, // Common
		"283f69": 101,  // Sony
		"2841c6": 3335, // Huawei
		"2841ec": 3335, // Huawei
		"2847aa": 457,  // Nokia
		"284846": 6806, // GridCentric
		"2848e7": 3335, // Huawei
		"284c53": 6807, // Intune
		"284d92": 6808, // Luminator
		"285261": 2,    // Cisco
		"2852e0": 6809, // Layon
		"28534e": 3335, // Huawei
		"285471": 3335, // Huawei
		"28563a": 4922, // Fiberhome
		"285767": 1238, // Dish
		"285aeb": 545,  // Apple
		"285f2f": 6810, // RNware
		"285fdb": 3335, // Huawei
		"286094": 6811, // Capelec
		"286336": 300,  // Siemens
		"2863bd": 4460, // Aptiv
		"2864b0": 3335, // Huawei
		"2866e3": 3005, // Azurewave
		"2868d2": 3335, // Huawei
		"286ab8": 545,  // Apple
		"286aba": 545,  // Apple
		"286b35": 421,  // Intel
		"286c07": 6281, // XIAOMI
		"286d97": 6812, // SAMJIN
		"286ed4": 3335, // Huawei
		"286f7f": 2,    // Cisco
		"2872c5": 1015, // Smartmatic
		"2872f0": 4309, // Athena
		"287610": 6813, // IgniteNet
		"287777": 3032, // ZTE
		"2877f1": 545,  // Apple
		"287aee": 125,  // ARRIS
		"287b09": 3032, // ZTE
		"287fcf": 421,  // Intel
		"288023": 302,  // HP
		"288088": 1368, // Netgear
		"288335": 152,  // Samsung
		"2884fa": 3206, // Sharp
		"28852d": 6814, // Touch
		"2887ba": 1595, // TP-Link
		"288a1c": 826,  // Juniper
		"288cb8": 3032, // ZTE
		"28924a": 302,  // HP
		"2893fe": 2,    // Cisco
		"28940f": 2,    // Cisco
		"2897b8": 6815, // myenergi
		"28987b": 152,  // Samsung
		"28993a": 3794, // Arista
		"289afa": 6434, // TCT
		"289e97": 3335, // Huawei
		"289efc": 2049, // Sagemcom
		"28a02b": 545,  // Apple
		"28a183": 430,  // Alpsalpine
		"28a186": 6816, // enblink
		"28a1eb": 6817, // Etek
		"28a241": 6818, // exlar
		"28a24b": 826,  // Juniper
		"28a53f": 6370, // Vivo
		"28a6ac": 6819, // seca
		"28a6db": 3335, // Huawei
		"28ac9e": 2,    // Cisco
		"28affd": 2,    // Cisco
		"28b2bd": 421,  // Intel
		"28b371": 2738, // Ruckus
		"28b448": 3335, // Huawei
		"28b4fb": 6820, // Sprocomm
		"28b9d9": 52,   // Radisys
		"28ba18": 6821, // NextNav
		"28bab5": 152,  // Samsung
		"28bb59": 6822, // RNET
		"28bc18": 6823, // SourcingOverseas
		"28bc56": 6824, // EMAC
		"28bd89": 3522, // Google
		"28be03": 6434, // TCT
		"28be9b": 6393, // Technicolor
		"28bf89": 4922, // Fiberhome
		"28c0da": 826,  // Juniper
		"28c21f": 152,  // Samsung
		"28c2dd": 3005, // Azurewave
		"28c538": 545,  // Apple
		"28c63f": 421,  // Intel
		"28c68e": 1368, // Netgear
		"28c709": 545,  // Apple
		"28c718": 6825, // Altierre
		"28c7ce": 2,    // Cisco
		"28c825": 6826, // DellKing
		"28c87a": 125,  // ARRIS
		"28c87c": 3032, // ZTE
		"28c914": 6827, // Taimag
		"28cbeb": 6828, // One
		"28cc01": 152,  // Samsung
		"28cd1c": 6829, // Espotel
		"28cd4c": 6830, // Individual
		"28cf08": 6831, // Essys
		"28cfda": 545,  // Apple
		"28cfe9": 545,  // Apple
		"28d0cb": 1640, // Cambridge
		"28d0ea": 421,  // Intel
		"28d0f5": 5441, // Ruijie
		"28d1af": 457,  // Nokia
		"28d244": 4920, // LCFC
		"28d3ea": 3335, // Huawei
		"28d93e": 6832, // Telecor
		"28d997": 6833, // Yuduan
		"28de65": 1685, // Aruba
		"28dea8": 3032, // ZTE
		"28dee5": 3335, // Huawei
		"28def6": 6834, // bioMerieux
		"28dfeb": 421,  // Intel
		"28e02c": 545,  // Apple
		"28e14c": 545,  // Apple
		"28e31f": 5698, // Xiaomi
		"28e347": 2874, // Liteon
		"28e34e": 3335, // Huawei
		"28e476": 6835, // Pi-Coral
		"28e5b0": 3335, // Huawei
		"28e608": 6836, // Tokheim
		"28e794": 6837, // Microtime
		"28e7cf": 545,  // Apple
		"28ea2d": 545,  // Apple
		"28ec95": 545,  // Apple
		"28ed6a": 545,  // Apple
		"28ede0": 4499, // AMPAK
		"28ee52": 1595, // TP-Link
		"28ef01": 5439, // Amazon
		"28f033": 545,  // Apple
		"28f076": 545,  // Apple
		"28f10e": 102,  // Dell
		"28f49b": 6838, // Leetek
		"28f532": 6839, // ADD-Engineering
		"28faa0": 6370, // Vivo
		"28fbae": 3335, // Huawei
		"28fbd3": 6840, // Ragentek
		"28fede": 6841, // CoMESTA
		"28ff3c": 545,  // Apple
		"28ff3e": 3032, // ZTE
		"28ffb2": 35,   // Toshiba
		"2c002c": 6842, // Unowhy
		"2c0033": 6843, // EControls
		"2c00ab": 125,  // ARRIS
		"2c00f7": 6844, // XOS
		"2c01b5": 2,    // Cisco
		"2c029f": 6845, // 3ALogics
		"2c073c": 6846, // Devline
		"2c0786": 3335, // Huawei
		"2c081c": 6847, // OVH
		"2c088c": 531,  // HUMAX
		"2c094d": 2239, // Raptor
		"2c09cb": 6848, // Cobs
		"2c0bab": 3335, // Huawei
		"2c0be9": 2,    // Cisco
		"2c0da7": 421,  // Intel
		"2c0e3d": 152,  // Samsung
		"2c10c1": 1427, // Nintendo
		"2c1165": 1655, // Silicon
		"2c15bf": 152,  // Samsung
		"2c15e1": 6849, // Phicomm
		"2c18ae": 176,  // Trend
		"2c1a01": 3335, // Huawei
		"2c1a05": 2,    // Cisco
		"2c1a31": 6850, // Electronics
		"2c1db8": 125,  // ARRIS
		"2c1eea": 6851, // Aerodev
		"2c1f23": 545,  // Apple
		"2c200b": 545,  // Apple
		"2c2131": 826,  // Juniper
		"2c2172": 826,  // Juniper
		"2c21d7": 6852, // IMAX
		"2c233a": 302,  // HP
		"2c26c5": 3032, // ZTE
		"2c2768": 3335, // Huawei
		"2c27d7": 302,  // HP
		"2c2997": 612,  // Microsoft
		"2c2bf9": 869,  // LG
		"2c2d48": 6853, // bct
		"2c301a": 6393, // Technicolor
		"2c3033": 1368, // Netgear
		"2c3068": 2278, // Pantech
		"2c3124": 2,    // Cisco
		"2c3311": 2,    // Cisco
		"2c3358": 421,  // Intel
		"2c3361": 545,  // Apple
		"2c36a0": 6854, // Capisco
		"2c36f8": 2,    // Cisco
		"2c3796": 6855, // Cybo
		"2c3996": 2049, // Sagemcom
		"2c39c1": 4565, // Ciena
		"2c3a28": 6856, // Fagor
		"2c3a91": 3335, // Huawei
		"2c3ae8": 6377, // Espressif
		"2c3bfd": 6857, // Netstor
		"2c3ecf": 2,    // Cisco
		"2c3f38": 2,    // Cisco
		"2c3f3e": 6858, // Alge-Timing
		"2c4053": 152,  // Samsung
		"2c4138": 302,  // HP
		"2c41a1": 1819, // Bose
		"2c4205": 6859, // Lytx
		"2c43be": 5436, // Sunnovo
		"2c4401": 152,  // Samsung
		"2c44fd": 302,  // HP
		"2c4881": 6370, // Vivo
		"2c4a11": 4565, // Ciena
		"2c4cc6": 2057, // Murata
		"2c4d54": 1806, // ASUS
		"2c4dde": 6266, // Tecno
		"2c4f52": 2,    // Cisco
		"2c52af": 3335, // Huawei
		"2c53d7": 6860, // Sonova
		"2c542d": 2,    // Cisco
		"2c5491": 612,  // Microsoft
		"2c54cf": 869,  // LG
		"2c553c": 6861, // Gainspeed
		"2c55d3": 3335, // Huawei
		"2c56dc": 1806, // ASUS
		"2c5731": 4031, // Wingtech
		"2c5741": 2,    // Cisco
		"2c584f": 125,  // ARRIS
		"2c58e8": 3335, // Huawei
		"2c598a": 869,  // LG
		"2c59e5": 302,  // HP
		"2c5a05": 457,  // Nokia
		"2c5a0f": 2,    // Cisco
		"2c5aa3": 6862, // Promate
		"2c5be1": 6863, // Centripetal
		"2c5d93": 2738, // Ruckus
		"2c5ff3": 6864, // Pertronic
		"2c600c": 3072, // Quanta
		"2c61f6": 545,  // Apple
		"2c6289": 6865, // Regenersis
		"2c641f": 3468, // Vizio
		"2c6798": 6866, // IntalTech
		"2c6bf5": 826,  // Juniper
		"2c6dc1": 421,  // Intel
		"2c6e85": 421,  // Intel
		"2c7155": 6867, // HiveMotion
		"2c71ff": 5439, // Amazon
		"2c72c3": 6868, // Soundmatters
		"2c7360": 6499, // Earda
		"2c73a0": 2,    // Cisco
		"2c768a": 302,  // HP
		"2c780e": 3335, // Huawei
		"2c784c": 6869, // Iton
		"2c79d7": 2049, // Sagemcom
		"2c7b5a": 6870, // Milper
		"2c7e81": 125,  // ARRIS
		"2c7ecf": 6871, // Onzo
		"2c86d2": 2,    // Cisco
		"2c8a72": 1341, // HTC
		"2c8db1": 421,  // Intel
		"2c9127": 6872, // Eintechno
		"2c9464": 6873, // Cincoze
		"2c9569": 125,  // ARRIS
		"2c957f": 3032, // ZTE
		"2c9662": 6874, // Invenit
		"2c97b1": 3335, // Huawei
		"2c97ed": 101,  // Sony
		"2c9924": 125,  // ARRIS
		"2c9aa4": 6875, // Eolo
		"2c9d1e": 3335, // Huawei
		"2c9e5f": 125,  // ARRIS
		"2c9efc": 87,   // Canon
		"2c9ffb": 1592, // Wistron
		"2ca02f": 6876, // Veroguard
		"2ca042": 3335, // Huawei
		"2ca157": 6877, // acromate
		"2ca17d": 125,  // ARRIS
		"2ca2b4": 6878, // Fortify
		"2ca327": 6879, // Oraimo
		"2ca780": 6880, // True
		"2ca835": 4556, // RIM
		"2ca89c": 6881, // Creatz
		"2caa8e": 6882, // Wyze
		"2cab00": 3335, // Huawei
		"2cabeb": 2,    // Cisco
		"2cac44": 6883, // Conextop
		"2cae2b": 152,  // Samsung
		"2cb05d": 1368, // Netgear
		"2cb0df": 6884, // Soliton
		"2cb21a": 6849, // Phicomm
		"2cb43a": 545,  // Apple
		"2cb693": 563,  // Radware
		"2cb8ed": 6640, // SonicWall
		"2cbaba": 152,  // Samsung
		"2cbc87": 545,  // Apple
		"2cbe08": 545,  // Apple
		"2cbeeb": 6885, // Nothing
		"2cc260": 11,   // Oracle
		"2cc407": 6886, // machineQ
		"2cc546": 3335, // Huawei
		"2cc548": 6887, // IAdea
		"2cc5d3": 2738, // Ruckus
		"2cc81b": 1784, // Routerboard.com
		"2ccc15": 457,  // Nokia
		"2ccc44": 101,  // Sony
		"2ccd27": 6888, // Precor
		"2ccd69": 6889, // Aqavi.com
		"2cce1e": 6890, // Cloudtronics
		"2ccf58": 3335, // Huawei
		"2cd02d": 2,    // Cisco
		"2cd05a": 2874, // Liteon
		"2cd066": 5698, // Xiaomi
		"2cd1da": 3600, // Keysight
		"2cd26b": 6441, // FN-Link
		"2cd2e7": 457,  // Nokia
		"2cd444": 4,    // Fujitsu
		"2cdb07": 421,  // Intel
		"2cdcad": 1592, // Wistron
		"2cdcd7": 3005, // Azurewave
		"2cdd0c": 6891, // Discovergy
		"2cdde9": 3794, // Arista
		"2ce2a8": 6892, // DeviceDesign
		"2ce310": 6893, // Stratacache
		"2ce412": 2049, // Sagemcom
		"2ce6cc": 2738, // Ruckus
		"2cea7f": 102,  // Dell
		"2ceadc": 2545, // Askey
		"2cf05d": 1812, // Micro-Star
		"2cf0a2": 545,  // Apple
		"2cf0ee": 545,  // Apple
		"2cf1bb": 3032, // ZTE
		"2cf295": 3335, // Huawei
		"2cf432": 6377, // Espressif
		"2cf4c5": 620,  // Avaya
		"2cf7f1": 6894, // Seeed
		"2cf89b": 2,    // Cisco
		"2cfda1": 1806, // ASUS
		"2cfdab": 687,  // Motorola
		"2cfdb3": 6895, // Tonly
		"2cffee": 6370, // Vivo
		"30053f": 6896, // JTI
		"30055c": 3231, // Brother
		"30074d": 152,  // Samsung
		"3009c0": 687,  // Motorola
		"300c23": 3032, // ZTE
		"300d43": 612,  // Microsoft
		"300d9e": 5441, // Ruijie
		"300eb8": 869,  // LG
		"300ee3": 3254, // Aquantia
		"3010b3": 2874, // Liteon
		"3010e4": 545,  // Apple
		"301389": 300,  // Siemens
		"30142d": 6897, // Piciorgros
		"30144a": 1592, // Wistron
		"301518": 6898, // Ubiquitous
		"30168d": 6899, // ProLon
		"3017c8": 101,  // Sony
		"301966": 152,  // Samsung
		"301a28": 6900, // Mako
		"301a30": 6900, // Mako
		"302303": 2469, // Belkin
		"302432": 421,  // Intel
		"302478": 2049, // Sagemcom
		"3024a9": 302,  // HP
		"3027cf": 67,   // Private
		"302952": 3778, // Hillstone
		"302bdc": 6901, // Top-Unum
		"302de8": 6902, // JDA
		"30317d": 1730, // Hosiden
		"3032d4": 6903, // Hanilstm
		"303335": 6904, // Boosty
		"303422": 5821, // eero
		"3034d2": 6905, // Availink
		"3035ad": 545,  // Apple
		"3035c5": 3335, // Huawei
		"3037a6": 2,    // Cisco
		"3037b3": 3335, // Huawei
		"303855": 457,  // Nokia
		"303926": 101,  // Sony
		"303a64": 421,  // Intel
		"303fbb": 302,  // HP
		"304225": 6906, // Burg-Wachter
		"304240": 3032, // ZTE
		"304449": 6907, // PLATH
		"304596": 3335, // Huawei
		"30469a": 1368, // Netgear
		"30499e": 3335, // Huawei
		"304b07": 687,  // Motorola
		"304c7e": 2154, // Panasonic
		"304e1b": 3335, // Huawei
		"3051f8": 6908, // BYK-Gardner
		"30525a": 2260, // NST
		"3052cb": 2874, // Liteon
		"3053c1": 3283, // Cresyn
		"305696": 6296, // Infinix
		"305714": 545,  // Apple
		"30578e": 5821, // eero
		"3057ac": 6909, // Irlab
		"30595b": 6910, // streamnow
		"3059b7": 612,  // Microsoft
		"305a3a": 1806, // ASUS
		"305d38": 6911, // Beissbarth
		"306023": 125,  // ARRIS
		"306112": 3937, // PAV
		"306118": 6912, // Paradom
		"30636b": 545,  // Apple
		"3065ec": 1592, // Wistron
		"3066d0": 3335, // Huawei
		"30688c": 6913, // Reach
		"30694b": 4556, // RIM
		"306a85": 152,  // Samsung
		"306cbe": 6914, // Skymotion
		"306e5c": 6915, // Validus
		"306f07": 6916, // Nations
		"307350": 6917, // Inpeco
		"307467": 152,  // Samsung
		"307496": 3335, // Huawei
		"307512": 101,  // Sony
		"30766f": 869,  // LG
		"3077cb": 6918, // Maike
		"3078d3": 6919, // Virgilant
		"307c30": 4556, // RIM
		"307c4a": 3335, // Huawei
		"307c5e": 826,  // Juniper
		"307ecb": 3188, // SFR
		"308398": 6377, // Espressif
		"3083d2": 687,  // Motorola
		"3085a9": 1806, // ASUS
		"3085eb": 4922, // Fiberhome
		"308730": 3335, // Huawei
		"3087d9": 2738, // Ruckus
		"308af7": 3335, // Huawei
		"308bb2": 2,    // Cisco
		"308cfb": 6920, // Dropcam
		"308d99": 302,  // HP
		"309048": 545,  // Apple
		"3090ab": 545,  // Apple
		"30918f": 6393, // Technicolor
		"3093bc": 2049, // Sagemcom
		"309435": 6370, // Vivo
		"309610": 3335, // Huawei
		"3096fb": 152,  // Samsung
		"309935": 3032, // ZTE
		"309c23": 1812, // Micro-Star
		"309e1d": 6578, // Ohsung
		"309ffb": 6921, // Ardomus
		"30a176": 4922, // Fiberhome
		"30a1fa": 3335, // Huawei
		"30a2c2": 3335, // Huawei
		"30a8db": 101,  // Sony
		"30a998": 3335, // Huawei
		"30a9de": 869,  // LG
		"30aae4": 3335, // Huawei
		"30ab6a": 152,  // Samsung
		"30aea4": 6377, // Espressif
		"30afce": 6370, // Vivo
		"30b164": 6922, // Power
		"30b1b5": 2640, // Arcadyan
		"30b49e": 1595, // TP-Link
		"30b4b8": 869,  // LG
		"30b5c2": 1595, // TP-Link
		"30b5f1": 6923, // Aitexin
		"30b62d": 2486, // Mojo
		"30b64f": 826,  // Juniper
		"30b7d4": 870,  // Hitron
		"30b930": 3032, // ZTE
		"30bb7d": 6924, // OnePlus
		"30c3d9": 430,  // Alpsalpine
		"30c50f": 3335, // Huawei
		"30c6f7": 6377, // Espressif
		"30c7ae": 152,  // Samsung
		"30cbc7": 665,  // Cambium
		"30cbf8": 152,  // Samsung
		"30cc21": 3032, // ZTE
		"30cda7": 152,  // Samsung
		"30d042": 102,  // Dell
		"30d16b": 2874, // Liteon
		"30d17e": 3335, // Huawei
		"30d32d": 1637, // devolo
		"30d357": 6925, // Logosol
		"30d386": 3032, // ZTE
		"30d46a": 6926, // Autosales
		"30d53e": 545,  // Apple
		"30d587": 152,  // Samsung
		"30d659": 6927, // Merging
		"30d6c9": 152,  // Samsung
		"30d9d9": 545,  // Apple
		"30e090": 6928, // Linctronix
		"30e171": 302,  // HP
		"30e37a": 421,  // Intel
		"30e396": 3335, // Huawei
		"30e4db": 2,    // Cisco
		"30e98e": 3335, // Huawei
		"30ea26": 6929, // Sycada
		"30f31d": 3032, // ZTE
		"30f335": 3335, // Huawei
		"30f42f": 4386, // ESP
		"30f70d": 2,    // Cisco
		"30f7c5": 545,  // Apple
		"30f7d7": 6930, // Thread
		"30f94b": 3858, // Universal
		"30f9ed": 101,  // Sony
		"30fbb8": 3335, // Huawei
		"30fc68": 1595, // TP-Link
		"30fceb": 869,  // LG
		"30fd11": 6931, // Macrotech
		"30fd38": 3522, // Google
		"30fd65": 3335, // Huawei
		"30fe31": 457,  // Nokia
		"3400a3": 3335, // Huawei
		"340286": 421,  // Intel
		"34029b": 6932, // Plexonics
		"34074f": 6933, // AccelStor
		"3407fb": 306,  // Ericsson
		"340804": 803,  // D-Link
		"3408bc": 545,  // Apple
		"340a22": 6934, // TOP-Access
		"340a33": 803,  // D-Link
		"340a98": 3335, // Huawei
		"340ced": 6935, // Moduel
		"341290": 6936, // Treeview
		"341298": 545,  // Apple
		"3412f9": 3335, // Huawei
		"3413a8": 6937, // Mediplan
		"3413e8": 421,  // Intel
		"34145f": 152,  // Samsung
		"34159e": 545,  // Apple
		"3417eb": 102,  // Dell
		"341a35": 4922, // Fiberhome
		"341b22": 6938, // Grandbeing
		"341cf0": 5698, // Xiaomi
		"341e6b": 3335, // Huawei
		"341fe4": 125,  // ARRIS
		"3420e3": 2738, // Ruckus
		"3423ba": 152,  // Samsung
		"34243e": 3032, // ZTE
		"342606": 6939, // CarePredict
		"342840": 545,  // Apple
		"3428f0": 6940, // ATN
		"342912": 3335, // Huawei
		"3429ea": 6941, // MCD
		"342b70": 6942, // Arris
		"342cc4": 358,  // Compal
		"342d0d": 152,  // Samsung
		"342eb6": 3335, // Huawei
		"342eb7": 421,  // Intel
		"342f6e": 6943, // Anywire
		"342fbd": 1427, // Nintendo
		"343111": 152,  // Samsung
		"34317f": 2154, // Panasonic
		"34318f": 545,  // Apple
		"3431c4": 621,  // AVM
		"3432e6": 2154, // Panasonic
		"34363b": 545,  // Apple
		"343654": 3032, // ZTE
		"343759": 3032, // ZTE
		"343794": 6944, // Hamee
		"3438af": 6945, // Inlab
		"3438b7": 531,  // HUMAX
		"343d98": 6946, // JinQianMao
		"343dc4": 1077, // Buffalo
		"343ea4": 6947, // Ring
		"3440b5": 372,  // IBM
		"34415d": 421,  // Intel
		"3441a8": 6948, // ER-Telecom
		"344262": 545,  // Apple
		"34466f": 6949, // HiTEM
		"3446ec": 3335, // Huawei
		"3448ed": 102,  // Dell
		"34495b": 2049, // Sagemcom
		"344b3d": 4922, // Fiberhome
		"344b50": 3032, // ZTE
		"344ca4": 6950, // amazipoint
		"344cc8": 6951, // Echodyne
		"344dea": 3032, // ZTE
		"344df7": 869,  // LG
		"344f3f": 6952, // IO-Power
		"344f5c": 6953, // R&M
		"345184": 3335, // Huawei
		"3451c9": 545,  // Apple
		"3453d2": 2049, // Sagemcom
		"345760": 6429, // MitraStar
		"345840": 3335, // Huawei
		"345a06": 3206, // Sharp
		"345c40": 6954, // Cargt
		"345d10": 6955, // Wytek
		"345d9e": 2049, // Sagemcom
		"3460f9": 1595, // TP-Link
		"346178": 1041, // Boeing
		"346288": 2,    // Cisco
		"3462b4": 5696, // Renesas
		"3464a9": 302,  // HP
		"3466ea": 6956, // Vertu
		"34684a": 6957, // Teraworks
		"346987": 3032, // ZTE
		"346ac2": 3335, // Huawei
		"346b46": 2049, // Sagemcom
		"346bd3": 3335, // Huawei
		"346d9c": 370,  // Carrier
		"346e8a": 6958, // Ecosense
		"346e9d": 306,  // Ericsson
		"346f24": 3005, // Azurewave
		"346f90": 2,    // Cisco
		"347146": 3335, // Huawei
		"34732d": 2,    // Cisco
		"34735a": 102,  // Dell
		"3475c7": 620,  // Avaya
		"347839": 3032, // ZTE
		"347877": 6959, // O-Net
		"347916": 3335, // Huawei
		"347a60": 125,  // ARRIS
		"347c25": 545,  // Apple
		"347df6": 421,  // Intel
		"347e00": 3335, // Huawei
		"347e39": 457,  // Nokia
		"347e5c": 2048, // Sonos
		"347eca": 6960, // Nextwill
		"34800d": 2250, // Cavium
		"3480b3": 5698, // Xiaomi
		"348137": 6961, // Unicard
		"3481c4": 621,  // AVM
		"3482c5": 152,  // Samsung
		"3482de": 6962, // Kiio
		"348302": 6963, // iFORCOM
		"348446": 306,  // Ericsson
		"348584": 185,  // Extreme
		"34865d": 6377, // Espressif
		"348a12": 1685, // Aruba
		"348a7b": 152,  // Samsung
		"348aae": 2049, // Sagemcom
		"348b75": 647,  // Lava
		"348f27": 2738, // Ruckus
		"34916f": 6964, // UserGate
		"349342": 6965, // TTE
		"349454": 6377, // Espressif
		"3495db": 242,  // Logitec
		"349672": 1595, // TP-Link
		"34976f": 6966, // Rootech
		"3497f6": 1806, // ASUS
		"3498b5": 1368, // Netgear
		"34996f": 6967, // VPI
		"349b5b": 6968, // Maquet
		"349d90": 6969, // Heinzmann
		"349e34": 6970, // Evervictory
		"349f7b": 87,   // Canon
		"34a183": 6971, // AWare
		"34a2a2": 3335, // Huawei
		"34a395": 545,  // Apple
		"34a3bf": 6972, // Terewave
		"34a5b4": 6973, // Navtech
		"34a7ba": 6974, // Fischer
		"34a843": 2849, // Kyocera
		"34a84e": 2,    // Cisco
		"34a8eb": 545,  // Apple
		"34aa8b": 152,  // Samsung
		"34aa99": 457,  // Nokia
		"34ab37": 545,  // Apple
		"34ab95": 6377, // Espressif
		"34af2c": 1427, // Nintendo
		"34afb3": 5439, // Amazon
		"34b20a": 3335, // Huawei
		"34b354": 3335, // Huawei
		"34b472": 6377, // Espressif
		"34b571": 6975, // Plds
		"34b883": 2,    // Cisco
		"34b98d": 5698, // Xiaomi
		"34ba75": 6976, // Everest
		"34ba9a": 6977, // Asiatelco
		"34bb1f": 2221, // BlackBerry
		"34bb26": 687,  // Motorola
		"34bdc8": 2,    // Cisco
		"34be00": 152,  // Samsung
		"34bf90": 4922, // Fiberhome
		"34c059": 545,  // Apple
		"34c3ac": 152,  // Samsung
		"34c3d2": 6441, // FN-Link
		"34c69a": 6978, // Enecsys
		"34c731": 430,  // Alpsalpine
		"34c803": 457,  // Nokia
		"34c93d": 421,  // Intel
		"34c99d": 6979, // Eidolon
		"34c9f0": 6980, // LM
		"34cc28": 6981, // Nexpring
		"34cd6d": 6982, // CommSky
		"34cdbe": 3335, // Huawei
		"34ce00": 6281, // XIAOMI
		"34ce94": 6983, // Parsec
		"34cff6": 421,  // Intel
		"34d09b": 6984, // MobilMAX
		"34d270": 5439, // Amazon
		"34d693": 3335, // Huawei
		"34d7b4": 6985, // Tributary
		"34d954": 6986, // WiBotic
		"34dab7": 3032, // ZTE
		"34db9c": 2049, // Sagemcom
		"34dbfd": 2,    // Cisco
		"34de1a": 421,  // Intel
		"34de34": 3032, // ZTE
		"34df2a": 6987, // Fujikon
		"34e0cf": 3032, // ZTE
		"34e12d": 421,  // Intel
		"34e2fd": 545,  // Apple
		"34e6ad": 421,  // Intel
		"34e6d7": 102,  // Dell
		"34e70b": 6988, // HAN
		"34e894": 1595, // TP-Link
		"34e911": 6370, // Vivo
		"34e9fe": 6989, // Metis
		"34ed1b": 2,    // Cisco
		"34ef44": 1939, // 2Wire
		"34ef8b": 6990, // NTT
		"34efb6": 6295, // Edgecore
		"34f39a": 421,  // Intel
		"34f39b": 6991, // WizLAN
		"34f62d": 3206, // Sharp
		"34f64b": 421,  // Intel
		"34f6d2": 2154, // Panasonic
		"34f716": 1595, // TP-Link
		"34f8e7": 2,    // Cisco
		"34f968": 6992, // ATEK
		"34fa9f": 2738, // Ruckus
		"34fc6f": 6993, // Alcea
		"34fcb9": 302,  // HP
		"34fcef": 869,  // LG
		"34fd6a": 545,  // Apple
		"34fe77": 545,  // Apple
		"34fe9e": 4,    // Fujitsu
		"380025": 421,  // Intel
		"380118": 6994, // ULVAC
		"380195": 152,  // Samsung
		"3802de": 2075, // Sercomm
		"3806b4": 6995, // A.D.C
		"3807d4": 6996, // Zeppelin
		"3808fd": 6997, // Silca
		"380a0a": 6998, // Sky-City
		"380a94": 152,  // Samsung
		"380aab": 6999, // Formlabs
		"380b40": 152,  // Samsung
		"380dd4": 389,  // Primax
		"380e4d": 2,    // Cisco
		"380f4a": 545,  // Apple
		"3810f0": 1685, // Aruba
		"381428": 102,  // Dell
		"38144e": 4922, // Fiberhome
		"3816d1": 152,  // Samsung
		"381766": 7000, // Promzakaz
		"3817c3": 302,  // HP
		"3817e1": 6393, // Technicolor
		"38184c": 101,  // Sony
		"38192f": 457,  // Nokia
		"381a52": 46,   // Epson
		"381c1a": 2,    // Cisco
		"381c23": 5783, // Hilan
		"381d14": 7001, // Skydio
		"381dd9": 6441, // FN-Link
		"381ec7": 6440, // Chipsea
		"382028": 3335, // Huawei
		"382056": 2,    // Cisco
		"3820a8": 7002, // ColorTokens
		"3821c7": 1685, // Aruba
		"3822e2": 302,  // HP
		"38256b": 612,  // Microsoft
		"38262b": 7003, // UTran
		"3826cd": 7004, // Andtek
		"3829dd": 7005, // ONvocal
		"382a19": 7006, // Technica
		"382c4a": 1806, // ASUS
		"382dd1": 152,  // Samsung
		"382de8": 152,  // Samsung
		"3830f9": 869,  // LG
		"3831ac": 7007, // WEG
		"3835fb": 2049, // Sagemcom
		"38378b": 3335, // Huawei
		"38384b": 6370, // Vivo
		"383bc8": 1939, // 2Wire
		"383d5b": 4922, // Fiberhome
		"383f10": 7008, // DBL
		"383fb3": 6393, // Technicolor
		"38420b": 2048, // Sonos
		"38437d": 358,  // Compal
		"3843e5": 7009, // Grotech
		"38453b": 2738, // Ruckus
		"38454c": 7010, // Light
		"38458c": 7011, // MyCloud
		"384608": 3032, // ZTE
		"3847bc": 3335, // Huawei
		"38484c": 545,  // Apple
		"384b24": 300,  // Siemens
		"384b5b": 7012, // Ztron
		"384c4f": 3335, // Huawei
		"384c90": 125,  // ARRIS
		"384f49": 826,  // Juniper
		"384ff0": 3005, // Azurewave
		"38521a": 457,  // Nokia
		"385247": 3335, // Huawei
		"38539c": 545,  // Apple
		"38549b": 3032, // ZTE
		"38563d": 612,  // Microsoft
		"38580c": 7013, // Panaccess
		"385b44": 1655, // Silicon
		"385c76": 542,  // Plantronics
		"386077": 5440, // Pegatron
		"3861a5": 7014, // Grabango
		"3863bb": 302,  // HP
		"3865b2": 545,  // Apple
		"386645": 7015, // OOSIC
		"3866f0": 545,  // Apple
		"386893": 421,  // Intel
		"3868a4": 152,  // Samsung
		"3868dd": 3979, // Inventec
		"386a77": 152,  // Samsung
		"386bbb": 125,  // ARRIS
		"386e88": 3032, // ZTE
		"386ea2": 6370, // Vivo
		"38700c": 125,  // ARRIS
		"3871de": 545,  // Apple
		"3872c0": 3870, // Comtrend
		"3876d1": 7016, // Euronda
		"387862": 101,  // Sony
		"387a0e": 421,  // Intel
		"387a3c": 4922, // Fiberhome
		"387b47": 7017, // AKELA
		"3880df": 687,  // Motorola
		"388345": 1595, // TP-Link
		"388602": 7018, // Flexoptix
		"3887d5": 421,  // Intel
		"38881e": 3335, // Huawei
		"3888a4": 545,  // Apple
		"38892c": 545,  // Apple
		"388ab7": 7019, // ITC
		"388b59": 3522, // Google
		"388c50": 869,  // LG
		"388e7a": 7020, // Autoit
		"388ee7": 7021, // Fanhattan
		"388f30": 152,  // Samsung
		"389052": 3335, // Huawei
		"3890a5": 2,    // Cisco
		"3891fb": 7022, // Xenox
		"389461": 5696, // Renesas
		"389496": 152,  // Samsung
		"3894e0": 7023, // Syrotech
		"3894ed": 1368, // Netgear
		"3897a4": 5694, // Elecom
		"3898d8": 7024, // Meritech
		"3898e9": 3335, // Huawei
		"389af6": 152,  // Samsung
		"389d92": 46,   // Epson
		"38a4ed": 5698, // Xiaomi
		"38a659": 2049, // Sagemcom
		"38a6ce": 3513, // BSkyB
		"38a851": 1439, // Moog
		"38a86b": 7025, // Orga
		"38a95f": 7026, // Actifio
		"38a9ea": 67,   // Private
		"38aa3c": 152,  // Samsung
		"38ac3d": 7027, // Nephos
		"38afd0": 7028, // Nevro
		"38afd7": 4,    // Fujitsu
		"38b3f7": 3335, // Huawei
		"38b54d": 545,  // Apple
		"38b5d3": 7029, // SecuWorks
		"38b725": 1592, // Wistron
		"38b74d": 7030, // Fijowave
		"38b800": 1592, // Wistron
		"38bab0": 858,  // Broadcom
		"38baf8": 421,  // Intel
		"38bb3c": 620,  // Avaya
		"38bc01": 3335, // Huawei
		"38bc1a": 7031, // MEIZU
		"38beab": 7032, // AltoBeam
		"38bf2f": 7033, // Espec
		"38c096": 430,  // Alpsalpine
		"38c70a": 7034, // WiFiSong
		"38c7ba": 3363, // CS
		"38c986": 545,  // Apple
		"38cada": 545,  // Apple
		"38d40b": 152,  // Samsung
		"38d547": 1806, // ASUS
		"38d7ca": 7035, // 7HUGS
		"38d82f": 3032, // ZTE
		"38de60": 7036, // Mohlenhoff
		"38dead": 421,  // Intel
		"38e1aa": 3032, // ZTE
		"38e2dd": 3032, // ZTE
		"38e39f": 687,  // Motorola
		"38e60a": 5698, // Xiaomi
		"38e7d8": 1341, // HTC
		"38eaa7": 302,  // HP
		"38eb47": 3335, // Huawei
		"38ec0d": 545,  // Apple
		"38ece4": 152,  // Samsung
		"38ed18": 2,    // Cisco
		"38ee9d": 7037, // Anedo
		"38f0c8": 7038, // Mevo
		"38f135": 7039, // SensorTec-Canada
		"38f23e": 612,  // Microsoft
		"38f32e": 7040, // Skullcandy
		"38f33f": 7041, // Tatsuno
		"38f3ab": 4920, // LCFC
		"38f3fb": 7042, // Asperiq
		"38f557": 7043, // Jolata
		"38f597": 7044, // home2net
		"38f73d": 5439, // Amazon
		"38f7f1": 3335, // Huawei
		"38f85e": 531,  // HUMAX
		"38f889": 3335, // Huawei
		"38f8ca": 7045, // OWIN
		"38f9d3": 545,  // Apple
		"38fb14": 3335, // Huawei
		"38fc98": 421,  // Intel
		"38fdf8": 2,    // Cisco
		"38ff36": 2738, // Ruckus
		"3c01ef": 101,  // Sony
		"3c02b1": 7046, // Creation
		"3c0461": 125,  // ARRIS
		"3c04bf": 7047, // PRAVIS
		"3c0518": 152,  // Samsung
		"3c0630": 545,  // Apple
		"3c06a7": 1595, // TP-Link
		"3c0754": 545,  // Apple
		"3c0771": 101,  // Sony
		"3c08f6": 2,    // Cisco
		"3c0c48": 7048, // Servergy
		"3c0cdb": 3506, // Unionman
		"3c0e23": 2,    // Cisco
		"3c0fc1": 7049, // KBC
		"3c106f": 7050, // Albahith
		"3c10e6": 7051, // PHAZR
		"3c13cc": 2,    // Cisco
		"3c15c2": 545,  // Apple
		"3c15ea": 7052, // Tescom
		"3c15fb": 3335, // Huawei
		"3c1710": 2049, // Sagemcom
		"3c189f": 457,  // Nokia
		"3c195e": 152,  // Samsung
		"3c197d": 306,  // Ericsson
		"3c1a57": 7053, // Cardiopulmonary
		"3c1a79": 7054, // Huayuan
		"3c1a9e": 7055, // VitalThings
		"3c1cbe": 7056, // Jadak
		"3c1e04": 803,  // D-Link
		"3c20f6": 152,  // Samsung
		"3c219c": 421,  // Intel
		"3c22fb": 545,  // Apple
		"3c25d7": 457,  // Nokia
		"3c286d": 3522, // Google
		"3c2af4": 3231, // Brother
		"3c2c30": 102,  // Dell
		"3c2c99": 6295, // Edgecore
		"3c2ef9": 545,  // Apple
		"3c2eff": 545,  // Apple
		"3c2f3a": 7057, // SFORZATO
		"3c300c": 7058, // Dewar
		"3c306f": 3335, // Huawei
		"3c3178": 7059, // Qolsys
		"3c3556": 7060, // Cognitec
		"3c363d": 457,  // Nokia
		"3c36e4": 125,  // ARRIS
		"3c3786": 1368, // Netgear
		"3c3888": 7061, // ConnectQuest
		"3c38f4": 101,  // Sony
		"3c39c3": 7062, // JW
		"3c3a73": 620,  // Avaya
		"3c3f51": 7063, // 2CRSI
		"3c410e": 2,    // Cisco
		"3c438e": 125,  // ARRIS
		"3c457a": 3513, // BSkyB
		"3c46d8": 1595, // TP-Link
		"3c4711": 3335, // Huawei
		"3c4937": 7064, // ASSMANN
		"3c4a92": 302,  // HP
		"3c4cd0": 1490, // Ceragon
		"3c4dbe": 545,  // Apple
		"3c4e47": 7065, // Etronic
		"3c510e": 2,    // Cisco
		"3c5282": 302,  // HP
		"3c53d7": 7066, // Cedes
		"3c5447": 3335, // Huawei
		"3c5731": 2,    // Cisco
		"3c576c": 152,  // Samsung
		"3c57d5": 7067, // FiveCo
		"3c58c2": 421,  // Intel
		"3c5a37": 152,  // Samsung
		"3c5ab4": 3522, // Google
		"3c5cc4": 5439, // Amazon
		"3c5cf1": 5821, // eero
		"3c5ec3": 2,    // Cisco
		"3c5f01": 7068, // Synerchip
		"3c6104": 826,  // Juniper
		"3c6105": 6377, // Espressif
		"3c6200": 152,  // Samsung
		"3c62f0": 2075, // Sercomm
		"3c672c": 7069, // Sciovid
		"3c678c": 3335, // Huawei
		"3c6816": 4743, // VXi
		"3c6a9d": 7070, // Dexatek
		"3c6aa7": 421,  // Intel
		"3c6e63": 3885, // Mitron
		"3c6f45": 7071, // Fiberpro
		"3c6fea": 2154, // Panasonic
		"3c6ff7": 7072, // EnTek
		"3c7059": 7073, // MakerBot
		"3c71bf": 6377, // Espressif
		"3c7437": 4556, // RIM
		"3c754a": 125,  // ARRIS
		"3c7843": 3335, // Huawei
		"3c7873": 7074, // Airsonics
		"3c7a8a": 125,  // ARRIS
		"3c7ac4": 7075, // Chemtronics
		"3c7af0": 6282, // Itel
		"3c7c3f": 1806, // ASUS
		"3c7d0a": 545,  // Apple
		"3c7f6f": 7076, // Telechips
		"3c81d8": 2049, // Sagemcom
		"3c831e": 7077, // CKD
		"3c8375": 612,  // Microsoft
		"3c846a": 1595, // TP-Link
		"3c869a": 3335, // Huawei
		"3c86d1": 6370, // Vivo
		"3c8970": 7078, // Neosfar
		"3c8994": 3513, // BSkyB
		"3c89a6": 7079, // Kapelse
		"3c8ab0": 826,  // Juniper
		"3c8b7f": 2,    // Cisco
		"3c8bfe": 152,  // Samsung
		"3c8c93": 826,  // Juniper
		"3c8cf8": 2905, // TRENDnet
		"3c8d20": 3522, // Google
		"3c9066": 4550, // SmartRG
		"3c912b": 7080, // Vexata
		"3c9157": 3097, // Yulong
		"3c9174": 7081, // Along
		"3c9180": 2874, // Liteon
		"3c92dc": 7082, // Octopod
		"3c93f4": 3335, // Huawei
		"3c94d5": 826,  // Juniper
		"3c9509": 2874, // Liteon
		"3c970e": 1592, // Wistron
		"3c977e": 7083, // IPS
		"3c9872": 2075, // Sercomm
		"3c99f7": 7084, // Lansentechnology
		"3c9a77": 6393, // Technicolor
		"3c9bc6": 3335, // Huawei
		"3c9bd6": 3468, // Vizio
		"3c9c0f": 421,  // Intel
		"3c9d56": 3335, // Huawei
		"3c9ec7": 3513, // BSkyB
		"3ca067": 2874, // Liteon
		"3ca10d": 152,  // Samsung
		"3ca161": 3335, // Huawei
		"3ca31a": 7085, // Oilfind
		"3ca348": 6370, // Vivo
		"3ca37e": 3335, // Huawei
		"3ca581": 6370, // Vivo
		"3ca616": 6370, // Vivo
		"3ca6f6": 545,  // Apple
		"3ca72b": 2254, // MRV
		"3ca82a": 302,  // HP
		"3ca916": 3335, // Huawei
		"3ca9f4": 421,  // Intel
		"3caa3f": 7086, // iKey
		"3cab8e": 545,  // Apple
		"3cb15b": 620,  // Avaya
		"3cb233": 3335, // Huawei
		"3cb6b7": 6370, // Vivo
		"3cb72b": 7087, // PLUMgrid
		"3cb74b": 6393, // Technicolor
		"3cb87a": 67,   // Private
		"3cbbfd": 152,  // Samsung
		"3cbdc5": 2640, // Arcadyan
		"3cbdd8": 869,  // LG
		"3cbee1": 5662, // Nikon
		"3cbf60": 545,  // Apple
		"3cc12c": 7088, // AES
		"3cc1f6": 7089, // Melange
		"3cc243": 457,  // Nokia
		"3cc99e": 7090, // Huiyang
		"3cca87": 7091, // Iders
		"3ccb7c": 6434, // TCT
		"3ccd36": 545,  // Apple
		"3ccd5d": 3335, // Huawei
		"3ccd93": 869,  // LG
		"3cce73": 2,    // Cisco
		"3cd0f8": 545,  // Apple
		"3cd16e": 7092, // Telepower
		"3cd4d6": 7093, // WirelessWERX
		"3cd92b": 302,  // HP
		"3cda2a": 3032, // ZTE
		"3cda6d": 7094, // Tiandy
		"3cdcbc": 152,  // Samsung
		"3cdf1e": 2,    // Cisco
		"3cdfa9": 125,  // ARRIS
		"3cdfbd": 3335, // Huawei
		"3ce038": 7095, // Omnifi
		"3ce072": 545,  // Apple
		"3ce624": 869,  // LG
		"3ce824": 3335, // Huawei
		"3ce9f7": 421,  // Intel
		"3cea4f": 1939, // 2Wire
		"3ceaf9": 7096, // Jubix
		"3ceafb": 7097, // NSE
		"3cf011": 421,  // Intel
		"3cf392": 7098, // Virtualtek
		"3cf4f9": 7099, // Moda-InnoChips
		"3cf52c": 7100, // DSPECIALISTS
		"3cf652": 3032, // ZTE
		"3cf72a": 457,  // Nokia
		"3cf7a4": 152,  // Samsung
		"3cf7d1": 7101, // OMRON
		"3cf808": 3335, // Huawei
		"3cf862": 421,  // Intel
		"3cfa43": 3335, // Huawei
		"3cfb5c": 4922, // Fiberhome
		"3cfb96": 7102, // Emcraft
		"3cfdfe": 421,  // Intel
		"3cffd8": 3335, // Huawei
		"4000e0": 7103, // DerekLimited
		"400107": 3794, // Arista
		"40017a": 2,    // Cisco
		"4001c6": 160,  // 3Com
		"40040c": 7104, // A&T
		"400589": 7105, // T-Mobile
		"400634": 3335, // Huawei
		"4006d5": 2,    // Cisco
		"4007c0": 7106, // Railtec
		"400d10": 125,  // ARRIS
		"400e67": 7107, // Tremol
		"400e85": 152,  // Samsung
		"4011c3": 152,  // Samsung
		"4011dc": 7108, // Sonance
		"4012e4": 7109, // Compass-EOS
		"4014ad": 3335, // Huawei
		"40163b": 152,  // Samsung
		"40167e": 1806, // ASUS
		"40169f": 1595, // TP-Link
		"4017e2": 7110, // Intai
		"4018b1": 185,  // Extreme
		"4018d7": 775,  // Smartronix
		"401920": 2723, // Movon
		"401c83": 421,  // Intel
		"4025c2": 421,  // Intel
		"402619": 545,  // Apple
		"40270b": 7111, // Mobileeco
		"402814": 7112, // RFI
		"402b50": 125,  // ARRIS
		"402ba1": 101,  // Sony
		"402e28": 4539, // MiXTelematics
		"402f86": 869,  // LG
		"403004": 545,  // Apple
		"403067": 7113, // Conlog
		"40313c": 6281, // XIAOMI
		"40331a": 545,  // Apple
		"4035e6": 152,  // Samsung
		"403cfc": 545,  // Apple
		"403dec": 531,  // HUMAX
		"403f8c": 1595, // TP-Link
		"404022": 7114, // ZIV
		"404028": 7114, // ZIV
		"40406b": 7115, // Icomera
		"40406c": 7115, // Icomera
		"4040a7": 101,  // Sony
		"404229": 7116, // Layer3TV
		"4045da": 7117, // Spreadtrum
		"40498a": 7118, // Synapticon
		"404a03": 2693, // ZyXEL
		"404ad4": 7119, // Widex
		"404c77": 125,  // ARRIS
		"404d7f": 545,  // Apple
		"404d8e": 3335, // Huawei
		"404e36": 1341, // HTC
		"40516c": 7120, // Grandex
		"40520d": 456,  // Pico
		"4054e4": 7121, // Wearsafe
		"405539": 2,    // Cisco
		"405582": 457,  // Nokia
		"405662": 7122, // GuoTengShengHua
		"405a9b": 7123, // Anovo
		"405cfd": 102,  // Dell
		"405d82": 1368, // Netgear
		"405f7d": 6434, // TCT
		"405fbe": 4556, // RIM
		"40605a": 7124, // Hawkeye
		"406186": 1812, // Micro-Star
		"40618e": 7125, // Stella-Green
		"406231": 7126, // Gifa
		"4065a3": 2049, // Sagemcom
		"406aab": 4556, // RIM
		"406c8f": 545,  // Apple
		"406f2a": 2221, // BlackBerry
		"407009": 125,  // ARRIS
		"407074": 7127, // Life
		"4070f5": 545,  // Apple
		"407183": 826,  // Juniper
		"407496": 7128, // aFUN
		"4074e0": 421,  // Intel
		"4076a9": 3335, // Huawei
		"40786a": 687,  // Motorola
		"407a80": 457,  // Nokia
		"407b1b": 7129, // Mettle
		"407c7d": 457,  // Nokia
		"407d0f": 3335, // Huawei
		"40831d": 545,  // Apple
		"4083de": 768,  // Zebra
		"408493": 7130, // Clavister
		"408805": 687,  // Motorola
		"40882f": 185,  // Extreme
		"4089a8": 7131, // WiredIQ
		"408a9a": 7132, // TITENG
		"408b07": 2247, // Actiontec
		"408d5c": 1929, // Giga-Byte
		"408f9d": 826,  // Juniper
		"409151": 6377, // Espressif
		"409505": 7133, // Acoinfo
		"409558": 7134, // Aisino
		"4095bd": 7135, // NTmore
		"4097d1": 7136, // BK
		"40984c": 7137, // Casacom
		"40987b": 7134, // Aisino
		"4098ad": 545,  // Apple
		"409922": 3005, // Azurewave
		"409b21": 457,  // Nokia
		"409bcd": 803,  // D-Link
		"409c28": 545,  // Apple
		"409ca6": 7138, // Curvalux
		"409f38": 3005, // Azurewave
		"409f87": 7139, // Jide
		"40a108": 687,  // Motorola
		"40a2db": 5439, // Amazon
		"40a3cc": 421,  // Intel
		"40a677": 826,  // Juniper
		"40a6a4": 7140, // PassivSystems
		"40a6b7": 421,  // Intel
		"40a6d9": 545,  // Apple
		"40a6e8": 2,    // Cisco
		"40a8f0": 302,  // HP
		"40a9cf": 5439, // Amazon
		"40b034": 302,  // HP
		"40b076": 1806, // ASUS
		"40b0a1": 6016, // Valcom
		"40b0fa": 869,  // LG
		"40b2c8": 76,   // Nortel
		"40b31e": 3858, // Universal
		"40b395": 545,  // Apple
		"40b3cd": 7141, // Chiyoda
		"40b3fc": 7142, // Logital
		"40b4cd": 5439, // Amazon
		"40b4f0": 826,  // Juniper
		"40b5c1": 2,    // Cisco
		"40b6b1": 7143, // SUNGSAM
		"40b6e7": 3335, // Huawei
		"40b7f3": 125,  // ARRIS
		"40b837": 101,  // Sony
		"40b93c": 302,  // HP
		"40ba61": 1957, // ARIMA
		"40bc60": 545,  // Apple
		"40bc8b": 7144, // itelio
		"40bd9e": 7145, // Physio-Control
		"40c3c6": 7146, // SnapRoute
		"40c48c": 7147, // N-iTUS
		"40c711": 545,  // Apple
		"40c729": 2049, // Sagemcom
		"40c7c9": 7148, // Naviit
		"40ca63": 7149, // Seongji
		"40cba8": 3335, // Huawei
		"40cbc0": 545,  // Apple
		"40cd3a": 7150, // Z3
		"40ce24": 2,    // Cisco
		"40d25f": 6282, // Itel
		"40d28a": 1427, // Nintendo
		"40d32d": 545,  // Apple
		"40d357": 7151, // Ison
		"40d3ae": 152,  // Samsung
		"40d40e": 7152, // Biodata
		"40d4bd": 7153, // SK
		"40d63c": 7154, // Equitech
		"40dc9d": 7155, // Hajen
		"40dca5": 3335, // Huawei
		"40dead": 826,  // Juniper
		"40e230": 3005, // Azurewave
		"40e3d6": 1685, // Aruba
		"40e64b": 545,  // Apple
		"40e99b": 152,  // Samsung
		"40ec99": 421,  // Intel
		"40ecf8": 300,  // Siemens
		"40ee15": 2128, // Zioncom
		"40eedd": 3335, // Huawei
		"40ef4c": 6299, // Fihonest
		"40f02f": 2874, // Liteon
		"40f078": 2,    // Cisco
		"40f201": 2049, // Sagemcom
		"40f2e9": 372,  // IBM
		"40f308": 2057, // Murata
		"40f407": 1427, // Nintendo
		"40f413": 7156, // Rubezh
		"40f4ec": 2,    // Cisco
		"40f4fd": 3506, // Unionman
		"40f520": 6377, // Espressif
		"40f6bc": 5439, // Amazon
		"40f946": 545,  // Apple
		"40f9d5": 7157, // Tecore
		"40fc89": 125,  // ARRIS
		"40fe0d": 7158, // Maxio
		"440010": 545,  // Apple
		"440049": 5439, // Amazon
		"44004d": 3335, // Huawei
		"44032c": 421,  // Intel
		"4403a7": 2,    // Cisco
		"44070b": 3522, // Google
		"4409b8": 7159, // Salcomp
		"440cfd": 7160, // NetMan
		"441102": 6662, // EDMI
		"441319": 7161, // WKK
		"4413d0": 3032, // ZTE
		"441441": 7162, // AudioControl
		"441622": 612,  // Microsoft
		"441793": 6377, // Espressif
		"44184f": 7163, // Fitview
		"4418fd": 545,  // Apple
		"441b88": 545,  // Apple
		"441c12": 6393, // Technicolor
		"441c7f": 687,  // Motorola
		"441e98": 2738, // Ruckus
		"441ea1": 302,  // HP
		"44227c": 3335, // Huawei
		"4422f1": 7164, // S.FAC
		"4423aa": 7165, // Farmage
		"4425bb": 7166, // Bamboo
		"4427f3": 7167, // 70mai
		"442938": 7168, // NietZsche
		"442a60": 545,  // Apple
		"442aff": 7169, // E3
		"442b03": 2,    // Cisco
		"442c05": 4499, // AMPAK
		"443192": 302,  // HP
		"44322a": 620,  // Avaya
		"4432c8": 6393, // Technicolor
		"44348f": 7170, // MXT
		"4434a7": 125,  // ARRIS
		"44356f": 7171, // Neterix
		"443583": 545,  // Apple
		"443839": 7172, // Cumulus
		"443b32": 3542, // Intelbras
		"443d21": 7173, // Nuvolt
		"443e07": 7174, // Electrolux
		"443eb2": 7175, // DEOTRON
		"44422f": 6251, // Testop
		"444450": 7176, // OttoQ
		"4448b9": 6429, // MitraStar
		"4448c1": 302,  // HP
		"444a65": 7177, // Silverflare
		"444adb": 545,  // Apple
		"444b7e": 4922, // Fiberhome
		"444c0c": 545,  // Apple
		"444ca8": 3794, // Arista
		"444e1a": 152,  // Samsung
		"444f8e": 7178, // WiZ
		"4455b1": 3335, // Huawei
		"4455c4": 3335, // Huawei
		"44568d": 7179, // PNC
		"4456b7": 7180, // Spawn
		"445943": 3032, // ZTE
		"44599f": 7181, // Criticare
		"4459e3": 3335, // Huawei
		"445bed": 1685, // Aruba
		"445ce9": 152,  // Samsung
		"445ecd": 7182, // Razer
		"446132": 7183, // ecobee
		"44619c": 7184, // FONsystem
		"446246": 7185, // Comat
		"44650d": 5439, // Amazon
		"44657f": 374,  // Calix
		"44666e": 7186, // IP-Line
		"446747": 3335, // Huawei
		"446752": 1592, // Wistron
		"44680c": 7187, // Wacom
		"4468ab": 7188, // Juin
		"446a2e": 3335, // Huawei
		"446ab7": 125,  // ARRIS
		"446c24": 7189, // Reallin
		"446d57": 2874, // Liteon
		"446d6c": 152,  // Samsung
		"446ee5": 3335, // Huawei
		"446ff8": 7190, // Dyson
		"44700b": 7191, // Iffu
		"4473d6": 4080, // Logitech
		"44746c": 101,  // Sony
		"447654": 3335, // Huawei
		"44783e": 152,  // Samsung
		"447bc4": 7192, // DualShine
		"447c7f": 7193, // Innolight
		"447e76": 7194, // Trek
		"4480eb": 687,  // Motorola
		"4482e5": 3335, // Huawei
		"448312": 7195, // Star-Net
		"448500": 421,  // Intel
		"4486c1": 300,  // Siemens
		"448723": 7196, // Hoya
		"4487fc": 1115, // Elitegroup
		"4488cb": 7197, // Camco
		"448a5b": 1812, // Micro-Star
		"448c52": 7198, // KTIS
		"448dbf": 7199, // Rhino
		"448e12": 7200, // DT
		"448e81": 7201, // VIG
		"448f17": 152,  // Samsung
		"4490bb": 545,  // Apple
		"449160": 2057, // Murata
		"4494fc": 1368, // Netgear
		"44962b": 7202, // Aidon
		"449bc1": 3335, // Huawei
		"449cb5": 7203, // Alcomp
		"449ef9": 6370, // Vivo
		"449f46": 3335, // Huawei
		"449f7f": 7204, // DataCore
		"44a191": 3335, // Huawei
		"44a42d": 6434, // TCT
		"44a54e": 6653, // Qorvo
		"44a56e": 1368, // Netgear
		"44a689": 7205, // Promax
		"44a6e5": 7206, // Thinking
		"44a7cf": 2057, // Murata
		"44a842": 102,  // Dell
		"44a8c2": 7207, // Sewoo
		"44a8fc": 545,  // Apple
		"44aa27": 7208, // udworks
		"44aa50": 826,  // Juniper
		"44aae8": 7209, // Nanotec
		"44aaf5": 125,  // ARRIS
		"44ad19": 7210, // Xingfei
		"44adb1": 2049, // Sagemcom
		"44add9": 2,    // Cisco
		"44ae25": 2,    // Cisco
		"44ae44": 3335, // Huawei
		"44af28": 421,  // Intel
		"44b32d": 1595, // TP-Link
		"44b412": 7211, // Sius
		"44b433": 7212, // tide
		"44b462": 833,  // Flextronics
		"44b6be": 2,    // Cisco
		"44bb3b": 3522, // Google
		"44bdde": 7213, // BHTC
		"44c306": 7214, // SIFROM
		"44c346": 3335, // Huawei
		"44c4a9": 3267, // Opticom
		"44c65d": 545,  // Apple
		"44c7fc": 3335, // Huawei
		"44c9a2": 7215, // Greenwald
		"44cb8b": 869,  // LG
		"44cd0e": 833,  // Flextronics
		"44ce7d": 3188, // SFR
		"44d244": 46,   // Epson
		"44d3ca": 2,    // Cisco
		"44d453": 2049, // Sagemcom
		"44d454": 2049, // Sagemcom
		"44d4e0": 101,  // Sony
		"44d5a5": 7216, // AddOn
		"44d5cc": 5439, // Amazon
		"44d63d": 7217, // Talari
		"44d6e1": 7218, // Snuza
		"44d791": 3335, // Huawei
		"44d832": 3005, // Azurewave
		"44d884": 545,  // Apple
		"44d9e7": 2979, // Ubiquiti
		"44da30": 545,  // Apple
		"44dc4e": 6282, // Itel
		"44dc91": 4479, // Planex
		"44dccb": 3667, // Semindia
		"44e137": 125,  // ARRIS
		"44e49a": 7219, // Omnitronics
		"44e4d9": 2,    // Cisco
		"44e4ee": 1592, // Wistron
		"44e517": 421,  // Intel
		"44e66e": 545,  // Apple
		"44e968": 3335, // Huawei
		"44e9dd": 2049, // Sagemcom
		"44ea30": 152,  // Samsung
		"44ea4b": 7220, // Actlas
		"44eb2e": 430,  // Alpsalpine
		"44ecce": 826,  // Juniper
		"44ed57": 7221, // Longicorn
		"44ee02": 2773, // MTI
		"44f034": 1269, // Kaonmedia
		"44f09e": 545,  // Apple
		"44f21b": 545,  // Apple
		"44f436": 3032, // ZTE
		"44f459": 152,  // Samsung
		"44f477": 826,  // Juniper
		"44f4e7": 7222, // Cohesity
		"44fb42": 545,  // Apple
		"44fb5a": 3032, // ZTE
		"44fda3": 7223, // Everysight
		"44fe3b": 2640, // Arcadyan
		"44ffba": 3032, // ZTE
		"480031": 3335, // Huawei
		"480033": 6393, // Technicolor
		"4801c5": 6924, // OnePlus
		"48022a": 7224, // B-Link
		"480362": 7225, // Desay
		"48049f": 5694, // Elecom
		"4805e2": 3335, // Huawei
		"48062b": 67,   // Private
		"48066a": 7226, // Tempered
		"480c49": 6083, // NAKAYO
		"480eec": 1595, // TP-Link
		"480fcf": 302,  // HP
		"481249": 5524, // Luxcom
		"481258": 3335, // Huawei
		"48128f": 3335, // Huawei
		"48137e": 152,  // Samsung
		"481693": 3536, // Lear
		"48174c": 7227, // MicroPower
		"4818fa": 7228, // Nocsys
		"48210b": 5440, // Pegatron
		"482567": 7229, // Poly
		"48262c": 545,  // Apple
		"4826e8": 7230, // Tek-Air
		"482759": 7231, // Levven
		"4827c5": 3335, // Huawei
		"4827ea": 152,  // Samsung
		"48282f": 3032, // ZTE
		"482952": 2049, // Sagemcom
		"482ae3": 1592, // Wistron
		"482ca0": 5698, // Xiaomi
		"482cd0": 3335, // Huawei
		"482e72": 2,    // Cisco
		"482f6b": 1685, // Aruba
		"482fd7": 3335, // Huawei
		"48343d": 7232, // IEP
		"48352b": 545,  // Apple
		"48365f": 3038, // Wintecronics
		"483871": 3335, // Huawei
		"483974": 3183, // Proware
		"483b38": 545,  // Apple
		"483c0c": 3335, // Huawei
		"483e5e": 6256, // Sernet
		"483fda": 6377, // Espressif
		"483fe9": 3335, // Huawei
		"48435a": 3335, // Huawei
		"48437c": 545,  // Apple
		"4843dd": 5439, // Amazon
		"4844f7": 152,  // Samsung
		"484520": 421,  // Intel
		"4846c1": 6441, // FN-Link
		"4846f1": 7233, // Uros
		"4846fb": 3335, // Huawei
		"48474b": 3335, // Huawei
		"4849c7": 152,  // Samsung
		"484ae9": 302,  // HP
		"484baa": 545,  // Apple
		"484bd4": 6393, // Technicolor
		"484c29": 3335, // Huawei
		"484c86": 3335, // Huawei
		"484d7e": 102,  // Dell
		"484efc": 125,  // ARRIS
		"485073": 612,  // Microsoft
		"485169": 152,  // Samsung
		"4851b7": 421,  // Intel
		"4851c5": 421,  // Intel
		"4851cf": 3542, // Intelbras
		"485261": 7234, // Soreel
		"485519": 6377, // Espressif
		"48555f": 4922, // Fiberhome
		"485702": 3335, // Huawei
		"4857dd": 7235, // Facebook
		"485929": 869,  // LG
		"4859a4": 3032, // ZTE
		"485a3f": 7236, // Wisol
		"485aea": 4922, // Fiberhome
		"485b39": 1806, // ASUS
		"485d36": 6622, // Verizon
		"485d60": 3005, // Azurewave
		"48605f": 869,  // LG
		"4860bc": 545,  // Apple
		"4861ee": 152,  // Samsung
		"486276": 3335, // Huawei
		"48684a": 421,  // Intel
		"486dbb": 1449, // Vestel
		"486e73": 7237, // Pica8
		"486fd2": 7238, // StorSimple
		"4873cb": 6540, // Tiinlab
		"487412": 6924, // OnePlus
		"48746e": 545,  // Apple
		"487583": 7239, // Intellion
		"487604": 67,   // Private
		"487746": 374,  // Calix
		"48785e": 5439, // Amazon
		"48794d": 152,  // Samsung
		"487a55": 5527, // ALE
		"487aff": 6831, // Essys
		"487b6b": 3335, // Huawei
		"487d2e": 1595, // TP-Link
		"487e48": 6499, // Earda
		"4883c7": 2049, // Sagemcom
		"4886e8": 612,  // Microsoft
		"488759": 5698, // Xiaomi
		"488764": 6370, // Vivo
		"488803": 7240, // ManTechnology
		"48881e": 7241, // EthoSwitch
		"4888ca": 687,  // Motorola
		"4889e7": 421,  // Intel
		"488ad2": 4254, // Mercury
		"488b0a": 2,    // Cisco
		"488c63": 3335, // Huawei
		"488d36": 2640, // Arcadyan
		"488e42": 7242, // DIGALOG
		"488eef": 3335, // Huawei
		"488f5a": 1784, // Routerboard.com
		"48902f": 869,  // LG
		"489a42": 7243, // Technomate
		"489bd5": 185,  // Extreme
		"489d18": 7244, // Flashbay
		"489d24": 2221, // BlackBerry
		"489dd1": 152,  // Samsung
		"489ebd": 302,  // HP
		"48a0f8": 4922, // Fiberhome
		"48a195": 545,  // Apple
		"48a2e6": 5987, // Resideo
		"48a472": 421,  // Intel
		"48a516": 3335, // Huawei
		"48a5e7": 1427, // Nintendo
		"48a6b8": 2048, // Sonos
		"48a74e": 3032, // ZTE
		"48a91c": 545,  // Apple
		"48a9d2": 1592, // Wistron
		"48aa5d": 7245, // Store
		"48ad08": 3335, // Huawei
		"48b02d": 7246, // NVIDIA
		"48b253": 7247, // Marketaxess
		"48b25d": 3335, // Huawei
		"48b423": 5439, // Amazon
		"48b620": 7248, // ROLI
		"48b8a3": 545,  // Apple
		"48b8de": 7249, // Homewins
		"48b977": 7250, // PulseOn
		"48b9c2": 7251, // Teletics
		"48ba4e": 302,  // HP
		"48bd4a": 3335, // Huawei
		"48be2d": 7252, // Symanitron
		"48bf6b": 545,  // Apple
		"48bf74": 7253, // Baicells
		"48c093": 2215, // Xirrus
		"48c1ac": 542,  // Plantronics
		"48c3b0": 7254, // Pharos
		"48c58d": 3536, // Lear
		"48c796": 152,  // Samsung
		"48c8b6": 7255, // SysTec
		"48cac6": 3506, // Unionman
		"48cb6e": 7256, // Cello
		"48d0cf": 3858, // Universal
		"48d18e": 6989, // Metis
		"48d224": 2874, // Liteon
		"48d24f": 2049, // Sagemcom
		"48d343": 125,  // ARRIS
		"48d35d": 67,   // Private
		"48d539": 3335, // Huawei
		"48d54c": 7257, // Jeda
		"48d6d5": 3522, // Google
		"48d705": 545,  // Apple
		"48d855": 7258, // Telvent
		"48d890": 6441, // FN-Link
		"48d8fe": 7259, // ClarIDy
		"48db50": 3335, // Huawei
		"48dc2d": 3335, // Huawei
		"48dcfb": 457,  // Nokia
		"48dd0c": 5821, // eero
		"48dd9d": 6282, // Itel
		"48df37": 302,  // HP
		"48e1af": 7260, // Vity
		"48e695": 7261, // Insigma
		"48e7da": 3005, // Azurewave
		"48e9f1": 545,  // Apple
		"48eb30": 7262, // Eterna
		"48eb62": 2057, // Murata
		"48ee0c": 803,  // D-Link
		"48ee86": 1135, // UTStarcom
		"48ef61": 3335, // Huawei
		"48f07b": 430,  // Alpsalpine
		"48f17f": 421,  // Intel
		"48f230": 7263, // Ubizcore
		"48f317": 67,   // Private
		"48f7c0": 6393, // Technicolor
		"48f7f1": 457,  // Nokia
		"48f8b3": 1783, // Linksys
		"48f8db": 3335, // Huawei
		"48f8e1": 457,  // Nokia
		"48f925": 7264, // Maestronic
		"48f97c": 4922, // Fiberhome
		"48fcb6": 647,  // Lava
		"48fcb8": 7265, // Woodstream
		"48fd8e": 3335, // Huawei
		"48fda3": 5698, // Xiaomi
		"4c0082": 2,    // Cisco
		"4c0143": 5821, // eero
		"4c0220": 5698, // Xiaomi
		"4c034f": 421,  // Intel
		"4c09b4": 3032, // ZTE
		"4c09d4": 2640, // Arcadyan
		"4c0a3d": 7266, // Adnacom
		"4c0b3a": 6434, // TCT
		"4c0bbe": 612,  // Microsoft
		"4c0fc7": 6499, // Earda
		"4c11ae": 6377, // Espressif
		"4c1265": 125,  // ARRIS
		"4c1365": 7267, // Emplus
		"4c1480": 7268, // Noregon
		"4c16f1": 3032, // ZTE
		"4c16fc": 826,  // Juniper
		"4c1744": 5439, // Amazon
		"4c17eb": 2049, // Sagemcom
		"4c195d": 2049, // Sagemcom
		"4c1a95": 7269, // Novakon
		"4c1b86": 2640, // Arcadyan
		"4c1d96": 421,  // Intel
		"4c1fcc": 3335, // Huawei
		"4c20b8": 545,  // Apple
		"4c218c": 2154, // Panasonic
		"4c21d0": 101,  // Sony
		"4c2258": 7270, // Cozybit
		"4c2578": 457,  // Nokia
		"4c26e7": 7271, // Welgate
		"4c2e5e": 152,  // Samsung
		"4c2fd7": 3335, // Huawei
		"4c322d": 7272, // Teledata
		"4c3275": 545,  // Apple
		"4c3329": 7273, // Sweroam
		"4c334e": 7274, // Hightech
		"4c3488": 421,  // Intel
		"4c364e": 2154, // Panasonic
		"4c38d8": 125,  // ARRIS
		"4c3910": 7275, // Newtek
		"4c3b6c": 7276, // Garo
		"4c3b74": 7277, // VOGTEC
		"4c3bdf": 612,  // Microsoft
		"4c3c16": 152,  // Samsung
		"4c4088": 7278, // Sanshin
		"4c445b": 421,  // Intel
		"4c494f": 3032, // ZTE
		"4c49e3": 5698, // Xiaomi
		"4c4e03": 6434, // TCT
		"4c4e35": 2,    // Cisco
		"4c4fee": 6924, // OnePlus
		"4c5077": 3335, // Huawei
		"4c5262": 4,    // Fujitsu
		"4c52ec": 7279, // SOLARWATT
		"4c53fd": 5439, // Amazon
		"4c5499": 3335, // Huawei
		"4c5585": 7280, // Hamilton
		"4c55cc": 7281, // Zentri
		"4c569d": 545,  // Apple
		"4c57ca": 545,  // Apple
		"4c5d3c": 2,    // Cisco
		"4c5e0c": 1784, // Routerboard.com
		"4c60de": 1368, // Netgear
		"4c617e": 3335, // Huawei
		"4c6371": 5698, // Xiaomi
		"4c6641": 152,  // Samsung
		"4c6be8": 545,  // Apple
		"4c6d58": 826,  // Juniper
		"4c6e6e": 7282, // Comnect
		"4c710c": 2,    // Cisco
		"4c710d": 2,    // Cisco
		"4c72b9": 5440, // Pegatron
		"4c73a5": 7283, // Kove
		"4c7403": 7284, // BQ
		"4c74bf": 545,  // Apple
		"4c7525": 6377, // Espressif
		"4c7625": 102,  // Dell
		"4c7713": 5696, // Renesas
		"4c776d": 2,    // Cisco
		"4c77cb": 421,  // Intel
		"4c796e": 421,  // Intel
		"4c7975": 545,  // Apple
		"4c79ba": 421,  // Intel
		"4c7c5f": 545,  // Apple
		"4c7cd9": 545,  // Apple
		"4c7f62": 457,  // Nokia
		"4c8093": 421,  // Intel
		"4c82cf": 1238, // Dish
		"4c875d": 1819, // Bose
		"4c8b30": 2247, // Actiontec
		"4c8bef": 3335, // Huawei
		"4c8d53": 3335, // Huawei
		"4c8d79": 545,  // Apple
		"4c8ecc": 7285, // Silkan
		"4c8fa5": 4069, // Jastec
		"4c9614": 826,  // Juniper
		"4c962d": 7286, // Fresh
		"4c98ef": 7287, // Zeo
		"4c9eff": 2693, // ZyXEL
		"4ca003": 7288, // Vitec
		"4ca515": 7289, // Baikal
		"4ca56d": 152,  // Samsung
		"4ca64d": 2,    // Cisco
		"4ca928": 7290, // Insensi
		"4caa16": 3005, // Azurewave
		"4cab33": 7291, // KST
		"4cab4f": 545,  // Apple
		"4cabf8": 2545, // Askey
		"4cabfc": 3032, // ZTE
		"4cac0a": 3032, // ZTE
		"4cada8": 7292, // Panoptics
		"4cae13": 3335, // Huawei
		"4cae31": 7293, // ShengHai
		"4caea3": 302,  // HP
		"4cb16c": 3335, // Huawei
		"4cb199": 545,  // Apple
		"4cb1cd": 2738, // Ruckus
		"4cb21c": 7294, // Maxphotonics
		"4cb44a": 7295, // NANOWAVE
		"4cb4ea": 7296, // HRD
		"4cb81c": 7297, // SAM
		"4cb910": 545,  // Apple
		"4cb911": 2051, // Raisecom
		"4cb9c8": 7298, // Conet
		"4cb9ea": 7299, // iRobot
		"4cbaa3": 7300, // Bison
		"4cbad7": 869,  // LG
		"4cbb58": 7301, // Chicony
		"4cbc48": 2,    // Cisco
		"4cbca5": 152,  // Samsung
		"4cbce9": 869,  // LG
		"4cc00a": 6370, // Vivo
		"4cc206": 7302, // Somfy
		"4cc449": 2157, // Icotera
		"4cc53e": 2693, // ZyXEL
		"4cc602": 7303, // Radios
		"4cc7d6": 833,  // Flextronics
		"4cc94f": 457,  // Nokia
		"4cc95e": 152,  // Samsung
		"4cca53": 7304, // Skyera
		"4ccbf5": 3032, // ZTE
		"4ccc34": 687,  // Motorola
		"4ccc6a": 1812, // Micro-Star
		"4cce2d": 6396, // Danlaw
		"4cd08a": 531,  // HUMAX
		"4cd0cb": 3335, // Huawei
		"4cd1a1": 3335, // Huawei
		"4cd629": 3335, // Huawei
		"4cd637": 7305, // Qsono
		"4cd98f": 102,  // Dell
		"4cdd31": 152,  // Samsung
		"4cdf3d": 7306, // Team
		"4ce0db": 5698, // Xiaomi
		"4ce175": 2,    // Cisco
		"4ce176": 2,    // Cisco
		"4ce19e": 6266, // Tecno
		"4ce676": 1077, // Buffalo
		"4ce6c0": 545,  // Apple
		"4ce933": 7307, // RailComm
		"4ceb42": 421,  // Intel
		"4cebd6": 6377, // Espressif
		"4cecef": 7308, // Soraa
		"4cedde": 2545, // Askey
		"4cedfb": 1806, // ASUS
		"4cefc0": 5439, // Amazon
		"4cf202": 5698, // Xiaomi
		"4cf2bf": 1640, // Cambridge
		"4cf55b": 3335, // Huawei
		"4cf737": 7309, // SamJi
		"4cf95d": 3335, // Huawei
		"4cfaca": 1640, // Cambridge
		"4cfb45": 3335, // Huawei
		"4cfcaa": 7310, // Tesla
		"4cff12": 7311, // Fuze
		"500084": 300,  // Siemens
		"50016b": 3335, // Huawei
		"5001bb": 152,  // Samsung
		"5001d9": 3335, // Huawei
		"500291": 6377, // Espressif
		"5004b8": 3335, // Huawei
		"500604": 2,    // Cisco
		"5006ab": 2,    // Cisco
		"500959": 6393, // Technicolor
		"5009e5": 7312, // Drimsys
		"500a52": 7313, // Huiwan
		"500b32": 7314, // Foxda
		"500e6d": 7315, // TrafficCast
		"500f59": 6528, // STMicrolectronics
		"500f80": 2,    // Cisco
		"500ff5": 6267, // Tenda
		"5011eb": 7316, // SilverNet
		"501408": 7317, // AiNET
		"501479": 7299, // iRobot
		"5017ff": 2,    // Cisco
		"50184c": 7318, // Platina
		"501ac5": 612,  // Microsoft
		"501cb0": 2,    // Cisco
		"501cbf": 2,    // Cisco
		"501d93": 3335, // Huawei
		"501e2d": 7319, // StreamUnlimited
		"501fc6": 545,  // Apple
		"5021ec": 3335, // Huawei
		"502267": 7320, // PixeLINK
		"5023a2": 545,  // Apple
		"502690": 4,    // Fujitsu
		"5027c7": 7321, // TECHNART
		"502873": 3335, // Huawei
		"502a7e": 7322, // Smart
		"502b73": 6267, // Tenda
		"502b98": 7323, // Es-tech
		"502d1d": 457,  // Nokia
		"502da2": 421,  // Intel
		"502dfb": 7324, // IGShare
		"502e5c": 1341, // HTC
		"502ece": 7325, // Asahi
		"502f9b": 421,  // Intel
		"502fa8": 2,    // Cisco
		"503237": 545,  // Apple
		"50325f": 1655, // Silicon
		"503275": 152,  // Samsung
		"5033f0": 7326, // Yichen
		"503cc4": 2665, // Lenovo
		"503da1": 152,  // Samsung
		"503dc6": 5698, // Xiaomi
		"503de5": 2,    // Cisco
		"503eaa": 1595, // TP-Link
		"503f56": 7327, // Syncmold
		"503f98": 7328, // Cmitech
		"504061": 457,  // Nokia
		"50411c": 4499, // AMPAK
		"504348": 7329, // ThingsMatrix
		"50464a": 3335, // Huawei
		"50465d": 1806, // ASUS
		"5046ae": 4254, // Mercury
		"5049b0": 152,  // Samsung
		"504a5e": 6363, // Masimo
		"504a6e": 1368, // Netgear
		"504b5b": 7330, // CoNTROLtronic
		"504b9e": 3335, // Huawei
		"504edc": 4358, // Ping
		"504f94": 7331, // Loxone
		"50502a": 7332, // Egardia
		"505065": 7333, // TAKT
		"5050a4": 152,  // Samsung
		"50523b": 457,  // Nokia
		"505527": 869,  // LG
		"5056a8": 7334, // Jolla
		"5056bf": 152,  // Samsung
		"50578a": 545,  // Apple
		"50579c": 46,   // Epson
		"5057a8": 2,    // Cisco
		"505800": 7335, // WyTec
		"50584f": 7336, // waytotec
		"50586f": 3335, // Huawei
		"505967": 7337, // Intent
		"505bc2": 2874, // Liteon
		"505d7a": 3032, // ZTE
		"505dac": 3335, // Huawei
		"505fb5": 2545, // Askey
		"506028": 2215, // Xirrus
		"506184": 620,  // Avaya
		"5061bf": 2,    // Cisco
		"5061d6": 7338, // Indu-Sol
		"5061f6": 3858, // Universal
		"50642b": 6281, // XIAOMI
		"506441": 7339, // Greenlee
		"5065f3": 302,  // HP
		"506787": 7340, // Planet
		"5067ae": 2,    // Cisco
		"5067f0": 2693, // ZyXEL
		"50680a": 3335, // Huawei
		"506a03": 1368, // Netgear
		"506b4b": 432,  // Mellanox
		"506b8d": 7341, // Nutanix
		"506cbe": 7342, // InnosiliconTechnology
		"506e92": 7343, // Innocent
		"506f0c": 2049, // Sagemcom
		"506f77": 3335, // Huawei
		"507043": 3513, // BSkyB
		"5075f1": 125,  // ARRIS
		"507691": 7344, // Tekpea
		"5076af": 421,  // Intel
		"507705": 152,  // Samsung
		"5078b0": 3335, // Huawei
		"5078b3": 3032, // ZTE
		"507a55": 545,  // Apple
		"507ac5": 545,  // Apple
		"507b9d": 4920, // LCFC
		"507c6f": 421,  // Intel
		"507d02": 7345, // Biodit
		"507e5d": 2640, // Arcadyan
		"508140": 302,  // HP
		"5082d5": 545,  // Apple
		"508492": 421,  // Intel
		"508569": 152,  // Samsung
		"508789": 2,    // Cisco
		"5087b8": 7346, // Nuvyyo
		"5089d1": 3335, // Huawei
		"508a42": 7347, // Uptmate
		"508d6f": 7348, // CHAHOO
		"508e49": 5698, // Xiaomi
		"508f4c": 5698, // Xiaomi
		"5092b9": 152,  // Samsung
		"509551": 125,  // ARRIS
		"509839": 5698, // Xiaomi
		"509871": 7349, // Inventum
		"509a46": 7350, // Safetrust
		"509a4c": 102,  // Dell
		"509a88": 3335, // Huawei
		"509ea7": 152,  // Samsung
		"509f27": 3335, // Huawei
		"50a009": 5698, // Xiaomi
		"50a054": 7351, // Actineon
		"50a0a4": 457,  // Nokia
		"50a4c8": 152,  // Samsung
		"50a5dc": 125,  // ARRIS
		"50a67f": 545,  // Apple
		"50a715": 7352, // Aboundi
		"50a72b": 3335, // Huawei
		"50a733": 2738, // Ruckus
		"50ab3e": 7353, // Qibixx
		"50ad92": 7354, // NX
		"50add5": 7355, // Dynalec
		"50ae86": 7356, // Linkintec
		"50af4d": 3032, // ZTE
		"50b7c3": 152,  // Samsung
		"50b8a2": 7357, // ImTech
		"50bc96": 545,  // Apple
		"50bd5f": 1595, // TP-Link
		"50c271": 7358, // Securetech
		"50c3a2": 3203, // nFore
		"50c4dd": 1077, // Buffalo
		"50c58d": 826,  // Juniper
		"50c6ad": 4922, // Fiberhome
		"50c709": 826,  // Juniper
		"50c7bf": 1595, // TP-Link
		"50c8e5": 152,  // Samsung
		"50c9a0": 7359, // Skipper
		"50ccf8": 152,  // Samsung
		"50cd22": 620,  // Avaya
		"50ce75": 7360, // Measy
		"50cee3": 7361, // Gigafirm
		"50d065": 7362, // ESYLUX
		"50d213": 7363, // CviLux
		"50d274": 7364, // Steffes
		"50d4f7": 1595, // TP-Link
		"50d753": 7365, // CoNELCOM
		"50dad6": 5698, // Xiaomi
		"50dcd0": 7366, // Observint
		"50dce7": 5439, // Amazon
		"50dcfc": 7367, // Ecocom
		"50de06": 545,  // Apple
		"50df95": 6859, // Lytx
		"50e039": 2693, // ZyXEL
		"50e085": 421,  // Intel
		"50e0c7": 7368, // TurControlSystme
		"50e0ef": 457,  // Nokia
		"50e14a": 67,   // Private
		"50e24e": 3032, // ZTE
		"50e549": 1929, // Giga-Byte
		"50e7a0": 5696, // Renesas
		"50e7b7": 6370, // Vivo
		"50e971": 7369, // Jibo
		"50ead6": 545,  // Apple
		"50eb1a": 90,   // Brocade
		"50eb71": 421,  // Intel
		"50ebf6": 1806, // ASUS
		"50ed3c": 545,  // Apple
		"50f0d3": 152,  // Samsung
		"50f43c": 7370, // Leeo
		"50f4eb": 545,  // Apple
		"50f520": 152,  // Samsung
		"50f5da": 5439, // Amazon
		"50f722": 2,    // Cisco
		"50f7ed": 3335, // Huawei
		"50f8a5": 7371, // eWBM
		"50f908": 7372, // Wizardlab
		"50f958": 3335, // Huawei
		"50fa84": 1595, // TP-Link
		"50fb19": 6440, // Chipsea
		"50fc30": 7373, // Treehouse
		"50fc9f": 152,  // Samsung
		"50fef2": 7374, // Sify
		"50ff20": 7375, // Keenetic
		"540237": 7376, // Teltronic
		"5403f5": 7377, // EBN
		"540496": 7378, // Gigawave
		"5404a6": 1806, // ASUS
		"540536": 7379, // Vivago
		"5405db": 4920, // LCFC
		"540764": 3335, // Huawei
		"540910": 545,  // Apple
		"540955": 3032, // ZTE
		"54098d": 7380, // deister
		"540df9": 3335, // Huawei
		"540e2d": 6370, // Vivo
		"540f57": 1655, // Silicon
		"541031": 7381, // Smarto
		"5410ec": 706,  // Microchip
		"54115f": 7382, // Atamo
		"541310": 3335, // Huawei
		"541473": 4031, // Wingtech
		"5414f3": 421,  // Intel
		"541651": 5441, // Ruijie
		"5419c8": 6370, // Vivo
		"541b5d": 7383, // Techno-Innov
		"541e56": 826,  // Juniper
		"541f8d": 3032, // ZTE
		"541fd5": 7384, // Advantage
		"542018": 7385, // Tely
		"54211d": 3335, // Huawei
		"542160": 7386, // Alula
		"54219d": 152,  // Samsung
		"5422f8": 3032, // ZTE
		"5425ea": 3335, // Huawei
		"542696": 545,  // Apple
		"54271e": 3005, // Azurewave
		"542758": 687,  // Motorola
		"542a1b": 2048, // Sonos
		"542aa2": 2237, // Alpha
		"542b8d": 545,  // Apple
		"542cea": 7387, // Protectron
		"542f89": 7388, // Euclid
		"5433cb": 545,  // Apple
		"5434ef": 3335, // Huawei
		"5435df": 7389, // Symeo
		"543968": 7390, // Edgewater
		"5439df": 3335, // Huawei
		"543ad6": 152,  // Samsung
		"543b30": 7391, // duagon
		"543d37": 2738, // Ruckus
		"543e64": 4922, // Fiberhome
		"5440ad": 152,  // Samsung
		"544249": 101,  // Sony
		"544408": 457,  // Nokia
		"5444a3": 152,  // Samsung
		"544617": 3032, // ZTE
		"544741": 7392, // Xcheng
		"5447d3": 7393, // Tsat
		"5447e8": 7023, // Syrotech
		"544810": 102,  // Dell
		"54489c": 7394, // Cdoubles
		"544a00": 2,    // Cisco
		"544b8c": 826,  // Juniper
		"544e45": 67,   // Private
		"544e90": 545,  // Apple
		"54511b": 3335, // Huawei
		"545146": 7395, // AMG
		"545284": 3335, // Huawei
		"5453ed": 101,  // Sony
		"5454cf": 7396, // Probedigital
		"5455d5": 3335, // Huawei
		"545aa6": 6377, // Espressif
		"545ebd": 2814, // NL
		"545fa9": 4821, // Teracom
		"546009": 3522, // Google
		"5461ea": 7397, // Zaplox
		"5462e2": 545,  // Apple
		"5464d9": 2049, // Sagemcom
		"5465de": 125,  // ARRIS
		"5466f9": 7398, // ConMet
		"546751": 358,  // Compal
		"546990": 3335, // Huawei
		"546ceb": 421,  // Intel
		"546f71": 7399, // uAvionix
		"5471dd": 3335, // Huawei
		"54724f": 545,  // Apple
		"54725e": 3506, // Unionman
		"547398": 7400, // Toyo
		"547595": 1595, // TP-Link
		"5475d0": 2,    // Cisco
		"547787": 6499, // Earda
		"54778a": 302,  // HP
		"54781a": 2,    // Cisco
		"5478c9": 4499, // AMPAK
		"547975": 457,  // Nokia
		"547c69": 2,    // Cisco
		"547d40": 7401, // Powervision
		"547f54": 536,  // Ingenico
		"547fa8": 7402, // TELCO
		"547fbc": 7403, // iodyne
		"547fee": 2,    // Cisco
		"548028": 302,  // HP
		"54812d": 3218, // PAX
		"5481ad": 4268, // Eagle
		"54833a": 2693, // ZyXEL
		"5484dc": 3032, // ZTE
		"5486bc": 2,    // Cisco
		"54880e": 152,  // Samsung
		"5488de": 2,    // Cisco
		"548922": 7404, // Zelfy
		"548998": 3335, // Huawei
		"548aba": 2,    // Cisco
		"548ca0": 2874, // Liteon
		"548d5a": 421,  // Intel
		"549209": 3335, // Huawei
		"5492be": 152,  // Samsung
		"549963": 545,  // Apple
		"549b12": 152,  // Samsung
		"549b72": 306,  // Ericsson
		"549d85": 7405, // EnerAccess
		"549f13": 545,  // Apple
		"549f35": 102,  // Dell
		"549fc6": 2,    // Cisco
		"54a04f": 7406, // t-mac
		"54a050": 1806, // ASUS
		"54a274": 2,    // Cisco
		"54a3fa": 7407, // BQT
		"54a51b": 3335, // Huawei
		"54a65c": 6393, // Technicolor
		"54a6db": 3335, // Huawei
		"54a703": 1595, // TP-Link
		"54a9d4": 7408, // Minibar
		"54ab3a": 3072, // Quanta
		"54ae27": 545,  // Apple
		"54aed0": 7409, // DASAN
		"54af97": 1595, // TP-Link
		"54b121": 3335, // Huawei
		"54b203": 5440, // Pegatron
		"54b7e5": 2602, // Rayson
		"54b802": 152,  // Samsung
		"54b80a": 803,  // D-Link
		"54bad6": 3335, // Huawei
		"54bd79": 152,  // Samsung
		"54be53": 3032, // ZTE
		"54bef7": 5440, // Pegatron
		"54bf64": 102,  // Dell
		"54c33e": 4565, // Ciena
		"54c480": 3335, // Huawei
		"54c57a": 5436, // Sunnovo
		"54c80f": 1595, // TP-Link
		"54c9df": 6441, // FN-Link
		"54cd10": 2154, // Panasonic
		"54ce82": 3032, // ZTE
		"54cf8d": 3335, // Huawei
		"54d0ed": 7410, // AXIM
		"54d163": 7411, // MAX-Tech
		"54d17d": 152,  // Samsung
		"54d751": 7412, // Proximus
		"54d9c6": 3335, // Huawei
		"54d9e4": 7413, // Brilliantts
		"54dba2": 7414, // Fibrain
		"54dc1d": 3097, // Yulong
		"54df00": 7415, // Ulterius
		"54df24": 4922, // Fiberhome
		"54df63": 7416, // Intrakey
		"54e005": 4922, // Fiberhome
		"54e019": 6947, // Ring
		"54e032": 826,  // Juniper
		"54e140": 536,  // Ingenico
		"54e1ad": 4920, // LCFC
		"54e2e0": 125,  // ARRIS
		"54e43a": 545,  // Apple
		"54e4a9": 7417, // BHR
		"54e4bd": 6441, // FN-Link
		"54e61b": 545,  // Apple
		"54e6fc": 1595, // TP-Link
		"54eaa8": 545,  // Apple
		"54ec2f": 2738, // Ruckus
		"54eda3": 7418, // Navdy
		"54ee75": 1592, // Wistron
		"54effe": 7419, // Fullpower
		"54f201": 152,  // Samsung
		"54f294": 3335, // Huawei
		"54f607": 3335, // Huawei
		"54f6e2": 3335, // Huawei
		"54f82a": 7420, // u-blox
		"54f876": 21,   // ABB
		"54fa3e": 152,  // Samsung
		"54fb58": 7421, // WISEWARE
		"54fcf0": 152,  // Samsung
		"5800bb": 826,  // Juniper
		"5800e3": 2874, // Liteon
		"580528": 7422, // Labris
		"580943": 67,   // Private
		"5809e5": 7423, // Kivic
		"580a20": 2,    // Cisco
		"580ad4": 545,  // Apple
		"58108c": 3542, // Intelbras
		"5810b7": 6296, // Infinix
		"581122": 1806, // ASUS
		"581243": 4904, // AcSiP
		"5813d3": 1450, // Gemtek
		"581626": 620,  // Avaya
		"5816d7": 430,  // Alpsalpine
		"58170c": 101,  // Sony
		"5819f8": 125,  // ARRIS
		"581cbd": 7424, // Affinegy
		"581f28": 3335, // Huawei
		"581f67": 7425, // Open-m
		"581faa": 545,  // Apple
		"581fef": 7426, // Tuttnaer
		"582059": 5698, // Xiaomi
		"5820b1": 302,  // HP
		"582136": 7427, // KMB
		"5821e9": 7428, // Twpi
		"58238c": 6393, // Technicolor
		"582429": 3522, // Google
		"582575": 3335, // Huawei
		"58278c": 1077, // Buffalo
		"582af7": 3335, // Huawei
		"582bdb": 7429, // Pax
		"582d34": 7430, // Qingping
		"582f40": 1427, // Nintendo
		"582ff7": 2049, // Sagemcom
		"583112": 7431, // Drust
		"583277": 2376, // Reliance
		"58343b": 7432, // Glovast
		"583526": 7433, // Deeplet
		"58355d": 3335, // Huawei
		"58356b": 6266, // Tecno
		"5835d9": 2,    // Cisco
		"583879": 75,   // Ricoh
		"583bd9": 4922, // Fiberhome
		"583cc6": 7434, // Omneality
		"583f54": 869,  // LG
		"58404e": 545,  // Apple
		"584120": 1595, // TP-Link
		"5842e4": 7435, // Baxter
		"584498": 5698, // Xiaomi
		"58454c": 306,  // Ericsson
		"58468f": 7436, // Koncar
		"5846e1": 7435, // Baxter
		"584822": 101,  // Sony
		"5848c0": 7437, // Coflec
		"584925": 7169, // E3
		"5849ba": 7438, // Chitai
		"584d42": 7439, // Dragos
		"5850ab": 7440, // TLS
		"5855ca": 545,  // Apple
		"5856c2": 3335, // Huawei
		"5856e8": 125,  // ARRIS
		"5859c2": 185,  // Extreme
		"585ff6": 3032, // ZTE
		"58605f": 3335, // Huawei
		"5860d8": 125,  // ARRIS
		"586356": 6441, // FN-Link
		"5865e6": 3977, // Infomark
		"58696c": 5441, // Ruijie
		"586b14": 545,  // Apple
		"586c25": 421,  // Intel
		"586d8f": 1783, // Linksys
		"586ed6": 67,   // Private
		"5876ac": 6256, // Sernet
		"587a4d": 7441, // Stonesoft
		"587f57": 545,  // Apple
		"587f66": 3335, // Huawei
		"587fb7": 7442, // Sonar
		"587fc8": 7443, // S2M
		"5882a8": 612,  // Microsoft
		"58856e": 5357, // QSC
		"588694": 1252, // EFM
		"588a5a": 102,  // Dell
		"588bf3": 2693, // ZyXEL
		"588d09": 2,    // Cisco
		"588e81": 1655, // Silicon
		"589043": 2049, // Sagemcom
		"5891cf": 421,  // Intel
		"589396": 2738, // Ruckus
		"58946b": 421,  // Intel
		"5894a2": 7444, // KETEK
		"5894ae": 3335, // Huawei
		"5894b2": 7445, // BrainCo
		"58961d": 421,  // Intel
		"589630": 6393, // Technicolor
		"58971e": 2,    // Cisco
		"5897bd": 2,    // Cisco
		"589835": 6393, // Technicolor
		"589b0b": 7446, // Shineway
		"589b4a": 7447, // DWnet
		"589ec6": 4299, // Gigaset
		"58a023": 421,  // Intel
		"58a0cb": 7448, // TrackNet
		"58a2b5": 869,  // LG
		"58a639": 152,  // Samsung
		"58a76f": 7449, // iD
		"58a839": 421,  // Intel
		"58a87b": 6588, // Fitbit
		"58ac78": 2,    // Cisco
		"58ae2b": 3335, // Huawei
		"58aea8": 3335, // Huawei
		"58aef1": 4922, // Fiberhome
		"58b035": 545,  // Apple
		"58b0d4": 7450, // ZuniData
		"58b10f": 152,  // Samsung
		"58b42d": 7451, // YSTen
		"58b633": 2738, // Ruckus
		"58b961": 7452, // SOLEM
		"58bad4": 3335, // Huawei
		"58bc27": 2,    // Cisco
		"58bc8f": 7453, // Cognitive
		"58bda3": 1427, // Nintendo
		"58bdf9": 7454, // Sigrand
		"58be72": 3335, // Huawei
		"58bf25": 6377, // Espressif
		"58bfea": 2,    // Cisco
		"58c17a": 665,  // Cambium
		"58c232": 48,   // NEC
		"58c38b": 152,  // Samsung
		"58c583": 6282, // Itel
		"58c5cb": 152,  // Samsung
		"58cb52": 3522, // Google
		"58ce2a": 421,  // Intel
		"58cf4b": 7455, // Lufkin
		"58cf79": 6377, // Espressif
		"58d061": 3335, // Huawei
		"58d349": 545,  // Apple
		"58d50a": 2057, // Murata
		"58d56e": 803,  // D-Link
		"58d67a": 7456, // TCPlink
		"58d759": 3335, // Huawei
		"58d9c3": 687,  // Motorola
		"58d9d5": 6267, // Tenda
		"58db15": 6266, // Tecno
		"58db8d": 298,  // Fast
		"58e28f": 545,  // Apple
		"58e326": 4102, // Compass
		"58e636": 7457, // EVRsafe
		"58e6ba": 545,  // Apple
		"58e747": 7458, // Deltanet
		"58e808": 7459, // Autonics
		"58eafc": 7460, // ELL-IoT
		"58ece1": 1388, // Newport
		"58ef68": 2469, // Belkin
		"58f102": 7461, // BLU
		"58f2fc": 3335, // Huawei
		"58f387": 7462, // Hccp
		"58f39c": 2,    // Cisco
		"58f987": 3335, // Huawei
		"58f98e": 7463, // SECUDOS
		"58fb84": 421,  // Intel
		"58fb96": 2738, // Ruckus
		"58fcc6": 7464, // Tozo
		"58fd20": 7465, // Systemhouse
		"58fdb1": 869,  // LG
		"5c0272": 1655, // Silicon
		"5c0339": 3335, // Huawei
		"5c0947": 545,  // Apple
		"5c0979": 3335, // Huawei
		"5c0a5b": 152,  // Samsung
		"5c0cbb": 7466, // CELIZION
		"5c0ce6": 1427, // Nintendo
		"5c0e8b": 185,  // Extreme
		"5c0ffb": 320,  // Amino
		"5c10c5": 152,  // Samsung
		"5c1515": 7467, // Advan
		"5c15e1": 7468, // Aidc
		"5c16c7": 3794, // Arista
		"5c1720": 3335, // Huawei
		"5c17cf": 6924, // OnePlus
		"5c17d3": 7469, // LGE
		"5c18b5": 7470, // Talon
		"5c1a6f": 1640, // Cambridge
		"5c1cb9": 6370, // Vivo
		"5c1dd9": 545,  // Apple
		"5c20d0": 7471, // Asoni
		"5c2316": 7472, // Squirrels
		"5c2479": 7473, // Baltech
		"5c260a": 102,  // Dell
		"5c2623": 7474, // WaveLynx
		"5c2e59": 152,  // Samsung
		"5c2ed2": 7475, // ABC
		"5c3192": 2,    // Cisco
		"5c32c5": 4821, // Teracom
		"5c338e": 2237, // Alpha
		"5c353b": 358,  // Compal
		"5c35da": 7476, // There
		"5c3a3d": 3032, // ZTE
		"5c3b35": 7477, // Gehirn
		"5c3c27": 152,  // Samsung
		"5c415a": 7478, // Amazon.com
		"5c41e7": 7479, // Wiatec
		"5c43d2": 7480, // Hazemeyer
		"5c443e": 7040, // Skullcandy
		"5c4527": 826,  // Juniper
		"5c497d": 152,  // Samsung
		"5c4a26": 7481, // Enguity
		"5c4ca9": 3335, // Huawei
		"5c5015": 2,    // Cisco
		"5c50d9": 545,  // Apple
		"5c514f": 421,  // Intel
		"5c5181": 152,  // Samsung
		"5c5188": 687,  // Motorola
		"5c521e": 1427, // Nintendo
		"5c5230": 545,  // Apple
		"5c546d": 3335, // Huawei
		"5c5578": 7482, // iryx
		"5c56ed": 7483, // 3pleplay
		"5c571a": 125,  // ARRIS
		"5c57c8": 457,  // Nokia
		"5c5819": 7484, // Jingsheng
		"5c5948": 545,  // Apple
		"5c5ac7": 2,    // Cisco
		"5c5aea": 4881, // Ford
		"5c5b35": 7485, // Mist
		"5c5bc2": 7486, // YIK
		"5c5eab": 826,  // Juniper
		"5c5f67": 421,  // Intel
		"5c625a": 87,   // Canon
		"5c63bf": 1595, // TP-Link
		"5c63c9": 7487, // Intellithings
		"5c647a": 3335, // Huawei
		"5c648e": 2693, // ZyXEL
		"5c6984": 7488, // Nuvico
		"5c6a80": 2693, // ZyXEL
		"5c6b4f": 7489, // Hello
		"5c6f69": 858,  // Broadcom
		"5c7017": 545,  // Apple
		"5c70a3": 869,  // LG
		"5c710d": 2,    // Cisco
		"5c75af": 6588, // Fitbit
		"5c7695": 6393, // Technicolor
		"5c7776": 6434, // TCT
		"5c78f8": 3335, // Huawei
		"5c7d5e": 3335, // Huawei
		"5c7d7d": 6393, // Technicolor
		"5c80b6": 421,  // Intel
		"5c8382": 457,  // Nokia
		"5c838f": 2,    // Cisco
		"5c843c": 101,  // Sony
		"5c8486": 7490, // Brightsource
		"5c864a": 7491, // Secret
		"5c865c": 152,  // Samsung
		"5c8730": 545,  // Apple
		"5c8778": 7492, // Cybertelbridge
		"5c879c": 421,  // Intel
		"5c899a": 1595, // TP-Link
		"5c8a38": 302,  // HP
		"5c8d4e": 545,  // Apple
		"5c8f40": 6266, // Tecno
		"5c8fe0": 125,  // ARRIS
		"5c9157": 3335, // Huawei
		"5c91fd": 7493, // Jaewoncnc
		"5c925e": 2128, // Zioncom
		"5c93a2": 2874, // Liteon
		"5c95ae": 545,  // Apple
		"5c9656": 3005, // Azurewave
		"5c9666": 101,  // Sony
		"5c966a": 7494, // Rtnet
		"5c969d": 545,  // Apple
		"5c97f3": 545,  // Apple
		"5c9960": 152,  // Samsung
		"5c9aa1": 3335, // Huawei
		"5c9ad8": 4,    // Fujitsu
		"5ca1e0": 7495, // EmbedWay
		"5ca39d": 152,  // Samsung
		"5ca48a": 2,    // Cisco
		"5ca4a4": 4922, // Fiberhome
		"5ca5bc": 5821, // eero
		"5ca62d": 2,    // Cisco
		"5ca6e6": 1595, // TP-Link
		"5ca86a": 3335, // Huawei
		"5caafd": 2048, // Sonos
		"5cadcf": 545,  // Apple
		"5caf06": 869,  // LG
		"5cb00a": 3335, // Huawei
		"5cb066": 125,  // ARRIS
		"5cb13e": 2049, // Sagemcom
		"5cb395": 3335, // Huawei
		"5cb3f6": 7496, // Humanorporated
		"5cb43e": 3335, // Huawei
		"5cb524": 101,  // Sony
		"5cb559": 7497, // CNEX
		"5cb6cc": 7498, // NovaComm
		"5cb8cb": 7499, // Allis
		"5cb901": 302,  // HP
		"5cba2c": 302,  // HP
		"5cba37": 612,  // Microsoft
		"5cbd9a": 3335, // Huawei
		"5cc0a0": 3335, // Huawei
		"5cc1d7": 152,  // Samsung
		"5cc307": 3335, // Huawei
		"5cc336": 7500, // ittim
		"5cc5d4": 421,  // Intel
		"5cc6e9": 6446, // Edifier
		"5cc7d7": 7501, // Azroad
		"5cc9c0": 5696, // Renesas
		"5cca1a": 612,  // Microsoft
		"5cca32": 7502, // Theben
		"5ccad3": 6440, // Chipsea
		"5ccb99": 152,  // Samsung
		"5ccca0": 7503, // Gridwiz
		"5ccd5b": 421,  // Intel
		"5ccd7c": 7031, // MEIZU
		"5ccead": 7504, // CDYNE
		"5ccf7f": 6377, // Espressif
		"5cd06e": 5698, // Xiaomi
		"5cd20b": 7505, // Yytek
		"5cd2e4": 421,  // Intel
		"5cd41b": 7506, // UCZOON
		"5cd4ab": 7507, // Zektor
		"5cd61f": 7508, // Qardio
		"5cd89e": 3335, // Huawei
		"5cd998": 803,  // D-Link
		"5cdad4": 2057, // Murata
		"5cdc96": 2640, // Arcadyan
		"5cdf89": 2738, // Ruckus
		"5ce0c5": 421,  // Intel
		"5ce176": 2,    // Cisco
		"5ce223": 7509, // Delphin
		"5ce286": 76,   // Nortel
		"5ce28c": 2693, // ZyXEL
		"5ce2f4": 4904, // AcSiP
		"5ce30e": 125,  // ARRIS
		"5ce3b6": 4922, // Fiberhome
		"5ce42a": 421,  // Intel
		"5ce747": 3335, // Huawei
		"5ce7a0": 457,  // Nokia
		"5ce883": 3335, // Huawei
		"5ce8b7": 6879, // Oraimo
		"5ce8eb": 152,  // Samsung
		"5ceb68": 7510, // Cheerstar
		"5ced8c": 302,  // HP
		"5cedf4": 152,  // Samsung
		"5cf207": 7511, // Speco
		"5cf370": 385,  // CC&C
		"5cf3fc": 372,  // IBM
		"5cf4ab": 2693, // ZyXEL
		"5cf5da": 545,  // Apple
		"5cf6dc": 152,  // Samsung
		"5cf7c3": 7512, // Syntech
		"5cf7e6": 545,  // Apple
		"5cf8a1": 2057, // Murata
		"5cf938": 545,  // Apple
		"5cf96a": 3335, // Huawei
		"5cf9dd": 102,  // Dell
		"5cf9f0": 7513, // Atomos
		"5cfafb": 7514, // Acubit
		"5cfc66": 2,    // Cisco
		"5cff35": 1592, // Wistron
		"600194": 6377, // Espressif
		"600292": 5440, // Pegatron
		"6002b4": 1592, // Wistron
		"600308": 545,  // Apple
		"600417": 7515, // Posbank
		"6006e3": 545,  // Apple
		"600810": 3335, // Huawei
		"6009c3": 7420, // u-blox
		"600f77": 7516, // SilverPlus
		"60109e": 3335, // Huawei
		"601199": 7517, // Siama
		"60123c": 3335, // Huawei
		"60128b": 87,   // Canon
		"601466": 3032, // ZTE
		"6014b3": 189,  // CyberTAN
		"6015c7": 7518, // IdaTech
		"601888": 3032, // ZTE
		"601895": 102,  // Dell
		"60190c": 7519, // Rramac
		"601971": 125,  // ARRIS
		"601d91": 687,  // Motorola
		"601e02": 7520, // EltexAlatau
		"602103": 7521, // I4VINE
		"6021c0": 2057, // Murata
		"602232": 2979, // Ubiquiti
		"6026aa": 2,    // Cisco
		"6026ef": 1685, // Aruba
		"6029d5": 7522, // DAVOLINK
		"602e20": 3335, // Huawei
		"6030d4": 545,  // Apple
		"60313b": 5436, // Sunnovo
		"603197": 2693, // ZyXEL
		"6032b1": 1595, // TP-Link
		"6032f0": 7523, // Mplus
		"60334b": 545,  // Apple
		"603553": 7524, // Buwon
		"603573": 6499, // Earda
		"6035c0": 3188, // SFR
		"603696": 7525, // Sapling
		"6036dd": 421,  // Intel
		"60380e": 430,  // Alpsalpine
		"6038e0": 2469, // Belkin
		"60391f": 21,   // ABB
		"603a7c": 1595, // TP-Link
		"603aaf": 152,  // Samsung
		"603cee": 869,  // LG
		"603d26": 6393, // Technicolor
		"603d29": 3335, // Huawei
		"603e7b": 7526, // Gafachi
		"60447a": 7527, // Water-i.d
		"6045bd": 612,  // Microsoft
		"6045cb": 1806, // ASUS
		"6047d4": 7528, // FORICS
		"6049c1": 620,  // Avaya
		"604a1c": 7529, // SUYIN
		"604de1": 3335, // Huawei
		"604f5b": 3335, // Huawei
		"60512c": 6434, // TCT
		"6052d0": 7530, // FACTS
		"605317": 7531, // Sandstone
		"605375": 3335, // Huawei
		"6055f9": 6377, // Espressif
		"605661": 7532, // IXECLOUD
		"605699": 67,   // Private
		"605718": 421,  // Intel
		"60577d": 5821, // eero
		"605bb4": 3005, // Azurewave
		"605e4f": 3335, // Huawei
		"605f8d": 5821, // eero
		"6061df": 7533, // Z-meta
		"60634c": 803,  // D-Link
		"6063f9": 7534, // Ciholas
		"606453": 7535, // AOD
		"6064a1": 7536, // RADiflow
		"606720": 421,  // Intel
		"60684e": 152,  // Samsung
		"606944": 545,  // Apple
		"60699b": 7537, // isepos
		"606bbd": 152,  // Samsung
		"606bff": 1427, // Nintendo
		"606c63": 870,  // Hitron
		"606c66": 421,  // Intel
		"606ed0": 7538, // Seal
		"606ee8": 5698, // Xiaomi
		"6070c0": 545,  // Apple
		"60720b": 7461, // BLU
		"60735c": 2,    // Cisco
		"6073bc": 3032, // ZTE
		"607688": 7539, // Velodyne
		"6077e2": 152,  // Samsung
		"607ec9": 545,  // Apple
		"607ecd": 3335, // Huawei
		"607edd": 612,  // Microsoft
		"6081f9": 7540, // Helium
		"608334": 3335, // Huawei
		"608373": 545,  // Apple
		"60843b": 7541, // Soladigm
		"6084bd": 1077, // Buffalo
		"608a10": 706,  // Microchip
		"608b0e": 545,  // Apple
		"608c2b": 3955, // Hanson
		"608c4a": 545,  // Apple
		"608ce6": 125,  // ARRIS
		"608d26": 2640, // Arcadyan
		"608e08": 152,  // Samsung
		"608f5c": 152,  // Samsung
		"609084": 7542, // DSSD
		"6091f3": 6370, // Vivo
		"609217": 545,  // Apple
		"6092f5": 125,  // ARRIS
		"609316": 545,  // Apple
		"609620": 67,   // Private
		"6097dd": 7543, // MicroSys
		"609ac1": 545,  // Apple
		"609bb4": 3335, // Huawei
		"609c9f": 90,   // Brocade
		"609e64": 7544, // Vivonic
		"609f9d": 7545, // CloudSwitch
		"60a10a": 152,  // Samsung
		"60a37d": 545,  // Apple
		"60a423": 1655, // Silicon
		"60a44c": 1806, // ASUS
		"60a4b7": 1595, // TP-Link
		"60a4d0": 152,  // Samsung
		"60a5e2": 421,  // Intel
		"60a6c5": 3335, // Huawei
		"60a751": 3335, // Huawei
		"60a9b0": 7546, // Merchandising
		"60aaef": 3335, // Huawei
		"60ab14": 869,  // LG
		"60ab67": 5698, // Xiaomi
		"60abd2": 1819, // Bose
		"60acc8": 7547, // KunTeng
		"60af6d": 152,  // Samsung
		"60b387": 7548, // Synergics
		"60b606": 7549, // Phorus
		"60b617": 4922, // Fiberhome
		"60b76e": 3522, // Google
		"60b933": 7550, // Deutron
		"60ba18": 7551, // nextLAP
		"60beb4": 7552, // S-Bluetech
		"60beb5": 687,  // Motorola
		"60bec4": 545,  // Apple
		"60c397": 1939, // 2Wire
		"60c547": 545,  // Apple
		"60c5ad": 152,  // Samsung
		"60c5e6": 7040, // Skullcandy
		"60c658": 7553, // PHYTRONIX
		"60c798": 1647, // Verifone
		"60c980": 7554, // Trymus
		"60cbfb": 7555, // AirScape
		"60cda9": 7556, // Abloomy
		"60ce41": 3335, // Huawei
		"60ce86": 2075, // Sercomm
		"60ce92": 7557, // Refined
		"60d02c": 2738, // Ruckus
		"60d0a9": 152,  // Samsung
		"60d21c": 5436, // Sunnovo
		"60d248": 125,  // ARRIS
		"60d262": 7558, // Tzukuri
		"60d30a": 7559, // Quatius
		"60d755": 3335, // Huawei
		"60d9a0": 2665, // Lenovo
		"60d9c7": 545,  // Apple
		"60da23": 7560, // Estech
		"60db2a": 7561, // HNS
		"60db98": 374,  // Calix
		"60dd70": 545,  // Apple
		"60dd8e": 421,  // Intel
		"60de35": 7562, // GITSN
		"60de44": 3335, // Huawei
		"60def3": 3335, // Huawei
		"60e00e": 4162, // Shinsei
		"60e327": 1595, // TP-Link
		"60e32b": 421,  // Intel
		"60e3ac": 869,  // LG
		"60e6bc": 7563, // Sino-Telecom
		"60e6f0": 1592, // Wistron
		"60e701": 3335, // Huawei
		"60e78a": 7564, // Unisem
		"60e956": 7565, // Ayla
		"60eb69": 3072, // Quanta
		"60f189": 2057, // Murata
		"60f18a": 3335, // Huawei
		"60f262": 421,  // Intel
		"60f281": 7566, // Tranwo
		"60f2ef": 7567, // VisionVera
		"60f43a": 6446, // Edifier
		"60f445": 545,  // Apple
		"60f59c": 7568, // CRU-Dataport
		"60f673": 7569, // Terumo
		"60f677": 421,  // Intel
		"60f81d": 545,  // Apple
		"60f8f2": 7570, // Synaptec
		"60fa9d": 3335, // Huawei
		"60facd": 545,  // Apple
		"60fb42": 545,  // Apple
		"60fcf1": 67,   // Private
		"60fd56": 7571, // WOORISYSTEMS
		"60fe20": 1939, // 2Wire
		"60fec5": 545,  // Apple
		"60ff12": 152,  // Samsung
		"60ffdd": 7572, // C.E
		"64002d": 7573, // Powerlinq
		"64006a": 102,  // Dell
		"6400f1": 2,    // Cisco
		"6401fb": 2228, // Landis+Gyr
		"6402cb": 125,  // ARRIS
		"64037f": 152,  // Samsung
		"6405e4": 430,  // Alpsalpine
		"6407f6": 152,  // Samsung
		"640980": 5698, // Xiaomi
		"6409ac": 6434, // TCT
		"640bd7": 545,  // Apple
		"640d22": 869,  // LG
		"640de6": 7574, // Petra
		"640e36": 7575, // Taztag
		"640e94": 7576, // Pluribus
		"640f28": 1939, // 2Wire
		"641225": 2,    // Cisco
		"641236": 6393, // Technicolor
		"641269": 125,  // ARRIS
		"64136c": 3032, // ZTE
		"6413ab": 3335, // Huawei
		"641666": 6636, // Nest
		"64167f": 751,  // Polycom
		"64168d": 2,    // Cisco
		"6416f0": 3335, // Huawei
		"641759": 7577, // Intellivision
		"641a22": 7578, // Heliospectra
		"641aba": 7579, // Dryad
		"641cae": 152,  // Samsung
		"641cb0": 152,  // Samsung
		"64200c": 545,  // Apple
		"64209f": 377,  // Tilgin
		"6420e0": 7580, // T3
		"642315": 3335, // Huawei
		"642400": 7581, // Xorcom
		"64255e": 7366, // Observint
		"642753": 3335, // Huawei
		"642c0f": 6370, // Vivo
		"642cac": 3335, // Huawei
		"642db7": 7582, // Seungil
		"643150": 302,  // HP
		"64317e": 7583, // Dexin
		"643216": 7584, // Weidu
		"6432a8": 421,  // Intel
		"643409": 7585, // BITwave
		"643aea": 2,    // Cisco
		"643e8c": 3335, // Huawei
		"643f5f": 7586, // Exablaze
		"6444d5": 67,   // Private
		"6447e0": 7587, // Feitian
		"644bf0": 7588, // CalDigit
		"644c36": 421,  // Intel
		"644d70": 7589, // dSPACE
		"644f42": 7590, // JETTER
		"644f74": 7591, // LENUS
		"644fb0": 7592, // Hyunjin.com
		"6450d6": 7593, // Liquidtool
		"645106": 302,  // HP
		"645563": 7594, // Intelight
		"6455b1": 125,  // ARRIS
		"645601": 1595, // TP-Link
		"645a04": 7301, // Chicony
		"645a36": 545,  // Apple
		"645aed": 545,  // Apple
		"645cf3": 7595, // ParanTek
		"645d86": 421,  // Intel
		"645df4": 152,  // Samsung
		"645e10": 3335, // Huawei
		"645e2c": 7596, // IRay
		"646140": 3335, // Huawei
		"646184": 7597, // Velux
		"646223": 7598, // Cellient
		"64628a": 7599, // evon
		"64649b": 826,  // Juniper
		"6465c0": 7600, // Nuvon
		"646624": 2049, // Sagemcom
		"6466b3": 1595, // TP-Link
		"64680c": 3870, // Comtrend
		"6469bc": 7601, // Hytera
		"646a52": 620,  // Avaya
		"646a74": 7602, // Auth-Servers
		"646cb2": 152,  // Samsung
		"646d2f": 545,  // Apple
		"646d4e": 3335, // Huawei
		"646d6c": 3335, // Huawei
		"646e69": 2874, // Liteon
		"646e97": 1595, // TP-Link
		"646ee0": 421,  // Intel
		"647002": 1595, // TP-Link
		"647033": 545,  // Apple
		"6472d8": 7603, // GooWi
		"6473e2": 7604, // Arbiter
		"6476ba": 545,  // Apple
		"64777d": 870,  // Hitron
		"647791": 152,  // Samsung
		"647924": 3335, // Huawei
		"6479a7": 7605, // Phison
		"6479f0": 421,  // Intel
		"647bce": 152,  // Samsung
		"647d81": 7606, // Yokota
		"647fda": 7607, // TEKTELIC
		"648099": 421,  // Intel
		"648788": 826,  // Juniper
		"64899a": 869,  // LG
		"6489f1": 152,  // Samsung
		"648d9e": 2991, // IVT
		"64956c": 869,  // LG
		"649714": 5821, // eero
		"64995d": 7469, // LGE
		"649968": 4653, // Elentec
		"649a12": 7608, // P2
		"649abe": 545,  // Apple
		"649b24": 7609, // V
		"649c81": 5788, // Qualcomm
		"649ef3": 2,    // Cisco
		"64a0e7": 2,    // Cisco
		"64a198": 3335, // Huawei
		"64a200": 5698, // Xiaomi
		"64a28a": 3335, // Huawei
		"64a2f9": 6924, // OnePlus
		"64a341": 7610, // Wonderlan
		"64a3cb": 545,  // Apple
		"64a5c3": 545,  // Apple
		"64a651": 3335, // Huawei
		"64a769": 1341, // HTC
		"64a7dd": 620,  // Avaya
		"64a965": 7611, // Linkflow
		"64ae0c": 2,    // Cisco
		"64ae88": 7612, // Polytec
		"64b0a6": 545,  // Apple
		"64b0e8": 3335, // Huawei
		"64b310": 152,  // Samsung
		"64b370": 7613, // PowerComm
		"64b379": 67,   // Private
		"64b473": 5698, // Xiaomi
		"64b5c6": 1427, // Nintendo
		"64b64a": 7614, // ViVOtech
		"64b853": 152,  // Samsung
		"64b94e": 102,  // Dell
		"64b9e8": 545,  // Apple
		"64babd": 7615, // SDJ
		"64bc0c": 869,  // LG
		"64bc11": 7616, // CombiQ
		"64bc58": 421,  // Intel
		"64be63": 7617, // STORDIS
		"64bf6b": 3335, // Huawei
		"64c269": 5821, // eero
		"64c2de": 869,  // LG
		"64c354": 620,  // Avaya
		"64c394": 3335, // Huawei
		"64c3d6": 826,  // Juniper
		"64c6af": 7618, // AXERRA
		"64c753": 545,  // Apple
		"64c901": 7619, // INVENTEC
		"64c944": 7620, // LARK
		"64cb9f": 6266, // Tecno
		"64cba3": 4613, // Pointmobile
		"64cbe9": 869,  // LG
		"64cc22": 2640, // Arcadyan
		"64cc2e": 5698, // Xiaomi
		"64d0d6": 152,  // Samsung
		"64d154": 1784, // Routerboard.com
		"64d1a3": 1881, // Sitecom
		"64d2c4": 545,  // Apple
		"64d4bd": 430,  // Alpsalpine
		"64d4da": 421,  // Intel
		"64d69a": 421,  // Intel
		"64d7c0": 3335, // Huawei
		"64d814": 2,    // Cisco
		"64d912": 7621, // Solidica
		"64d989": 2,    // Cisco
		"64db18": 7622, // OpenPattern
		"64db43": 687,  // Motorola
		"64db81": 7623, // Syszone
		"64dde9": 5698, // Xiaomi
		"64de1c": 7624, // Kingnetic
		"64dfe9": 7625, // Ateme
		"64e0ab": 3506, // Unionman
		"64e161": 7626, // DEP
		"64e599": 1252, // EFM
		"64e682": 545,  // Apple
		"64e7d8": 152,  // Samsung
		"64e84f": 7627, // Serialway
		"64e881": 1685, // Aruba
		"64e950": 2,    // Cisco
		"64eb8c": 46,   // Epson
		"64ed57": 125,  // ARRIS
		"64ed62": 7628, // WOORI
		"64eeb7": 2363, // Netcore
		"64f50e": 7629, // Kinion
		"64f69d": 2,    // Cisco
		"64f705": 3335, // Huawei
		"64f81c": 3335, // Huawei
		"64f970": 7630, // Kenade
		"64f987": 7631, // Avvasi
		"64fb50": 7632, // RoomReady/Zdi
		"64fc8c": 7633, // Zonar
		"64ff0a": 1592, // Wistron
		"680235": 7634, // Konten
		"6802b8": 358,  // Compal
		"680571": 152,  // Samsung
		"6805ca": 421,  // Intel
		"680715": 421,  // Intel
		"680927": 545,  // Apple
		"680ae2": 1655, // Silicon
		"681324": 3335, // Huawei
		"681590": 2049, // Sagemcom
		"681605": 7635, // Systems
		"681729": 421,  // Intel
		"681ab2": 3032, // ZTE
		"681bef": 3335, // Huawei
		"681ca2": 7636, // Rosewill
		"681d64": 7637, // Sunwave
		"681e8b": 7638, // InfoSight
		"681fd8": 300,  // Siemens
		"68215f": 6295, // Edgecore
		"68228e": 826,  // Juniper
		"682719": 706,  // Microchip
		"682737": 152,  // Samsung
		"6828ba": 7639, // Dejai
		"6828f6": 7640, // Vubiq
		"6829dc": 7641, // Ficosa
		"682c7b": 2,    // Cisco
		"682f67": 545,  // Apple
		"6831fe": 7642, // Teladin
		"68332c": 7643, // Kenstel
		"6836b5": 7644, // DriveScale
		"6837e9": 5439, // Amazon
		"683943": 7500, // ittim
		"683a48": 6812, // SAMJIN
		"683b1e": 7645, // Countwise
		"683b78": 2,    // Cisco
		"683e02": 7646, // SIEMENS
		"683e26": 421,  // Intel
		"683e34": 7031, // MEIZU
		"683eec": 7647, // Ereca
		"68403c": 4922, // Fiberhome
		"684352": 7648, // Bhuu
		"684571": 3335, // Huawei
		"6845f1": 35,   // Toshiba
		"684898": 152,  // Samsung
		"684a76": 5821, // eero
		"684aae": 3335, // Huawei
		"684ae9": 152,  // Samsung
		"684f64": 102,  // Dell
		"6851b7": 7649, // PowerCloud
		"6852d6": 7650, // UGame
		"68536c": 7651, // SPnS
		"685388": 7652, // P&S
		"68545a": 421,  // Intel
		"6854c1": 7002, // ColorTokens
		"6854f5": 7653, // enLighted
		"6854fd": 5439, // Amazon
		"685811": 4922, // Fiberhome
		"685acf": 152,  // Samsung
		"685b35": 545,  // Apple
		"685b36": 2952, // Powertech
		"685d43": 421,  // Intel
		"685e6b": 7654, // PowerRay
		"68644b": 545,  // Apple
		"686725": 6377, // Espressif
		"68692e": 7655, // Zycoo
		"686975": 7656, // Angler
		"6869ca": 89,   // Hitachi
		"686e23": 7657, // Wi3
		"686e48": 7658, // Prophet
		"687251": 2979, // Ubiquiti
		"6872c3": 152,  // Samsung
		"6872dc": 7659, // CETORY.TV
		"68764f": 101,  // Sony
		"687724": 1595, // TP-Link
		"687848": 7660, // Westunitis
		"68784c": 76,   // Nortel
		"687924": 7661, // ELS-GmbH
		"6879ed": 3206, // Sharp
		"687d6b": 152,  // Samsung
		"687db4": 2,    // Cisco
		"687f74": 1783, // Linksys
		"6881e0": 3335, // Huawei
		"6882f2": 7662, // grandcentrix
		"68831a": 7663, // Pandora
		"688470": 7664, // eSSys
		"68847e": 4,    // Fujitsu
		"688540": 7665, // IGI
		"68856a": 7666, // OuterLink
		"6886a7": 2,    // Cisco
		"6886e7": 7667, // Orbotix
		"68876b": 7668, // INQ
		"6887c6": 2,    // Cisco
		"6888a1": 3858, // Universal
		"688975": 7669, // nuoxc
		"6889c1": 3335, // Huawei
		"688af0": 3032, // ZTE
		"688db6": 7670, // Aetek
		"688f2e": 870,  // Hitron
		"688f84": 3335, // Huawei
		"688fc9": 7671, // Zhuolian
		"689234": 2738, // Ruckus
		"68966a": 6578, // Ohsung
		"68967b": 545,  // Apple
		"689861": 7672, // Beacon
		"6899cd": 2,    // Cisco
		"689a87": 5439, // Amazon
		"689c5e": 4904, // AcSiP
		"689c70": 545,  // Apple
		"689ce2": 2,    // Cisco
		"689e0b": 2,    // Cisco
		"689e6a": 3335, // Huawei
		"689ff0": 3032, // ZTE
		"68a03e": 3335, // Huawei
		"68a0f6": 3335, // Huawei
		"68a3c4": 2874, // Liteon
		"68a828": 3335, // Huawei
		"68a86d": 545,  // Apple
		"68a878": 7673, // GeoWAN
		"68a8e1": 7187, // Wacom
		"68aad2": 7674, // Datecs
		"68ab09": 457,  // Nokia
		"68ab1e": 545,  // Apple
		"68ae20": 545,  // Apple
		"68af13": 7675, // Futura
		"68b094": 7676, // Inesa
		"68b43a": 7677, // WaterFurnace
		"68b599": 302,  // HP
		"68b691": 5439, // Amazon
		"68b6fc": 870,  // Hitron
		"68b983": 7678, // b-plus
		"68bc0c": 2,    // Cisco
		"68bdab": 2,    // Cisco
		"68bfc4": 152,  // Samsung
		"68c44d": 687,  // Motorola
		"68c63a": 6377, // Espressif
		"68ca00": 7679, // Octopus
		"68cae4": 2,    // Cisco
		"68cc6e": 3335, // Huawei
		"68d48b": 7680, // Hailo
		"68d79a": 2979, // Ubiquiti
		"68d93c": 545,  // Apple
		"68db54": 6849, // Phicomm
		"68db96": 7681, // OPWILL
		"68dbca": 545,  // Apple
		"68dbf5": 5439, // Amazon
		"68dce8": 7682, // PacketStorm
		"68dfdd": 5698, // Xiaomi
		"68e166": 67,   // Private
		"68e209": 3335, // Huawei
		"68e7c2": 152,  // Samsung
		"68e8eb": 7683, // Linktel
		"68ebae": 152,  // Samsung
		"68ec62": 7684, // YODO
		"68ec8a": 67,   // Private
		"68ecc5": 421,  // Intel
		"68ed43": 2221, // BlackBerry
		"68ef43": 545,  // Apple
		"68efbd": 2,    // Cisco
		"68f06d": 7081, // Along
		"68f0d0": 7685, // SkyBell
		"68f38e": 826,  // Juniper
		"68f728": 4920, // LCFC
		"68f895": 7686, // Redflow
		"68fb7e": 545,  // Apple
		"68fb95": 7687, // Generalplus
		"68fcca": 152,  // Samsung
		"68feda": 4922, // Fiberhome
		"68fef7": 545,  // Apple
		"68ff7b": 1595, // TP-Link
		"6c006b": 152,  // Samsung
		"6c02e0": 302,  // HP
		"6c0309": 2,    // Cisco
		"6c047a": 3335, // Huawei
		"6c05d5": 7688, // Ethertronics
		"6c06d6": 3335, // Huawei
		"6c09bf": 4922, // Fiberhome
		"6c09d6": 7689, // Digiquest
		"6c0d34": 457,  // Nokia
		"6c0e0d": 101,  // Sony
		"6c0f61": 7690, // Hypervolt
		"6c0f6a": 7691, // JDC
		"6c108b": 4399, // WeLink
		"6c13d5": 2,    // Cisco
		"6c1414": 7692, // BUJEON
		"6c146e": 3335, // Huawei
		"6c14f7": 7693, // Erhardt+Leimer
		"6c15f9": 7694, // Nautronix
		"6c160e": 7695, // ShotTracker
		"6c1632": 3335, // Huawei
		"6c1811": 7696, // Decatur
		"6c198f": 803,  // D-Link
		"6c19c0": 545,  // Apple
		"6c1a75": 3335, // Huawei
		"6c1b3f": 7697, // MiraeSignal
		"6c1deb": 7420, // u-blox
		"6c1ed7": 6370, // Vivo
		"6c2056": 2,    // Cisco
		"6c21a2": 4499, // AMPAK
		"6c23b9": 101,  // Sony
		"6c23cb": 7698, // Wattty
		"6c2408": 4920, // LCFC
		"6c2483": 612,  // Microsoft
		"6c24a6": 6370, // Vivo
		"6c2636": 3335, // Huawei
		"6c2779": 612,  // Microsoft
		"6c2995": 421,  // Intel
		"6c2b59": 102,  // Dell
		"6c2e33": 7699, // Accelink
		"6c2e85": 2049, // Sagemcom
		"6c2f2c": 152,  // Samsung
		"6c2f8a": 152,  // Samsung
		"6c310e": 2,    // Cisco
		"6c32de": 7700, // Indieon
		"6c33a9": 7701, // Magicjack
		"6c3491": 3335, // Huawei
		"6c3845": 4922, // Fiberhome
		"6c3a36": 7702, // Glowforge
		"6c3b6b": 1784, // Routerboard.com
		"6c3be5": 302,  // HP
		"6c3c53": 7703, // SoundHawk
		"6c3c7c": 87,   // Canon
		"6c3e6d": 545,  // Apple
		"6c4008": 545,  // Apple
		"6c410e": 2,    // Cisco
		"6c416a": 2,    // Cisco
		"6c42ab": 7704, // Subscriber
		"6c433c": 6266, // Tecno
		"6c4418": 7705, // Zappware
		"6c442a": 3335, // Huawei
		"6c4598": 7706, // Antex
		"6c4760": 3940, // Sunitec
		"6c49c1": 7707, // o2ones
		"6c4a39": 7708, // Bita
		"6c4a74": 7709, // Aerodisk
		"6c4a85": 545,  // Apple
		"6c4b90": 7710, // LiteON
		"6c4bb4": 531,  // HUMAX
		"6c4d73": 545,  // Apple
		"6c504d": 2,    // Cisco
		"6c51bf": 3335, // Huawei
		"6c54cd": 7711, // Lampex
		"6c558d": 3335, // Huawei
		"6c55e8": 6393, // Technicolor
		"6c5640": 7461, // BLU
		"6c5697": 5439, // Amazon
		"6c5779": 7712, // Aclima
		"6c5940": 4254, // Mercury
		"6c5ab0": 1595, // TP-Link
		"6c5cde": 7713, // SunReports
		"6c5d3a": 612,  // Microsoft
		"6c5e3b": 2,    // Cisco
		"6c5f1c": 2665, // Lenovo
		"6c6126": 7714, // Rinicom
		"6c61f4": 3188, // SFR
		"6c626d": 1812, // Micro-Star
		"6c639c": 125,  // ARRIS
		"6c67ef": 3335, // Huawei
		"6c6a77": 421,  // Intel
		"6c6c0f": 3335, // Huawei
		"6c6cd3": 2,    // Cisco
		"6c6d09": 7715, // Kyowa
		"6c6f18": 7716, // Stereotaxis
		"6c7039": 7717, // Novar
		"6c709f": 545,  // Apple
		"6c710d": 2,    // Cisco
		"6c71d2": 3335, // Huawei
		"6c71d9": 3005, // Azurewave
		"6c7220": 803,  // D-Link
		"6c72e7": 545,  // Apple
		"6c750d": 7718, // WiFiSONG
		"6c7637": 3335, // Huawei
		"6c7660": 2849, // Kyocera
		"6c81fe": 7719, // Mitsuba
		"6c8336": 152,  // Samsung
		"6c8686": 7720, // Technonia
		"6c8814": 421,  // Intel
		"6c8b2f": 3032, // ZTE
		"6c8bd3": 2,    // Cisco
		"6c8cdb": 7721, // Otus
		"6c8d77": 2,    // Cisco
		"6c8dc1": 545,  // Apple
		"6c8fb5": 612,  // Microsoft
		"6c90b1": 7722, // SanLogic
		"6c9106": 67,   // Private
		"6c92bf": 7723, // Inspur
		"6c9354": 7724, // Yaojin
		"6c9392": 7725, // BEKO
		"6c9466": 421,  // Intel
		"6c94f8": 545,  // Apple
		"6c9522": 7726, // Scalys
		"6c96cf": 545,  // Apple
		"6c98eb": 2094, // Riverbed
		"6c9961": 2049, // Sagemcom
		"6c9989": 2,    // Cisco
		"6c999d": 5439, // Amazon
		"6c9ac9": 7727, // Valentine
		"6c9b02": 457,  // Nokia
		"6c9bc0": 7728, // Chemoptics
		"6c9ced": 2,    // Cisco
		"6c9e7c": 4922, // Fiberhome
		"6ca0b4": 3513, // BSkyB
		"6ca100": 421,  // Intel
		"6ca401": 7729, // essensys
		"6ca4d1": 4922, // Fiberhome
		"6ca604": 125,  // ARRIS
		"6ca75f": 3032, // ZTE
		"6ca780": 457,  // Nokia
		"6ca7fa": 2001, // Youngbo
		"6ca849": 620,  // Avaya
		"6ca858": 4922, // Fiberhome
		"6ca906": 7730, // Telefield
		"6ca936": 2233, // DisplayLink
		"6ca96f": 7731, // TransPacket
		"6caab3": 2738, // Ruckus
		"6cab05": 2,    // Cisco
		"6cab31": 545,  // Apple
		"6cac60": 7732, // Venetex
		"6cad3f": 7733, // Hubbell
		"6cadf8": 3005, // Azurewave
		"6cae8b": 372,  // IBM
		"6caee3": 457,  // Nokia
		"6caef6": 5821, // eero
		"6cb0ce": 1368, // Netgear
		"6cb158": 1595, // TP-Link
		"6cb227": 101,  // Sony
		"6cb2ae": 2,    // Cisco
		"6cb4a7": 7734, // Landauer
		"6cb4fd": 3335, // Huawei
		"6cb56b": 531,  // HUMAX
		"6cb6ca": 7735, // DivUS
		"6cb749": 3335, // Huawei
		"6cb7f4": 152,  // Samsung
		"6cb881": 3032, // ZTE
		"6cb9c5": 957,  // Delta
		"6cbab8": 2049, // Sagemcom
		"6cbfb5": 7736, // Noon
		"6cc1d2": 125,  // ARRIS
		"6cc217": 302,  // HP
		"6cc26b": 545,  // Apple
		"6cc49f": 1685, // Aruba
		"6cc7ec": 152,  // Samsung
		"6cca08": 125,  // ARRIS
		"6ccdd6": 1368, // Netgear
		"6cce44": 6628, // 1MORE
		"6cd032": 869,  // LG
		"6cd146": 7737, // FRAMOS
		"6cd1e5": 3335, // Huawei
		"6cd2ba": 3032, // ZTE
		"6cd3ee": 7738, // Zimi
		"6cd68a": 869,  // LG
		"6cd704": 3335, // Huawei
		"6cd719": 4922, // Fiberhome
		"6cd94c": 6370, // Vivo
		"6cdc6a": 7739, // Promethean
		"6cdd30": 2,    // Cisco
		"6cddbc": 152,  // Samsung
		"6cddef": 7740, // EPCOMM
		"6ce01e": 7741, // Modcam
		"6ce0b0": 7742, // SOUND4
		"6ce5c9": 545,  // Apple
		"6ce85c": 545,  // Apple
		"6ce873": 1595, // TP-Link
		"6ce874": 3335, // Huawei
		"6ce8c6": 6499, // Earda
		"6ce907": 457,  // Nokia
		"6ce983": 7743, // Gastron
		"6cebb6": 3335, // Huawei
		"6ced51": 7744, // NEXCONTROL
		"6cf049": 1929, // Giga-Byte
		"6cf373": 152,  // Samsung
		"6cf37f": 1685, // Aruba
		"6cf5e8": 7745, // Mooredoll
		"6cf784": 5698, // Xiaomi
		"6cf97c": 7746, // Nanoptix
		"6cfa58": 620,  // Avaya
		"6cfa89": 2,    // Cisco
		"6cfaa7": 4499, // AMPAK
		"6cfdb9": 3183, // Proware
		"6cfe54": 421,  // Intel
		"6cffbe": 7747, // MPB
		"6cffce": 2049, // Sagemcom
		"7001b5": 2,    // Cisco
		"700258": 7748, // 01DB-Metravib
		"70037e": 6393, // Technicolor
		"70039f": 6377, // Espressif
		"70041d": 6377, // Espressif
		"700514": 869,  // LG
		"7006ac": 7749, // Eastcompeace
		"700777": 7750, // OnTarget
		"700971": 152,  // Samsung
		"700b01": 2049, // Sagemcom
		"700b4f": 2,    // Cisco
		"700bc0": 7751, // Dewav
		"700f6a": 2,    // Cisco
		"700fec": 7752, // Poindus
		"70105c": 2,    // Cisco
		"70106f": 302,  // HP
		"701124": 545,  // Apple
		"701135": 7753, // Livesecu
		"7014a6": 545,  // Apple
		"7018a7": 2,    // Cisco
		"70192f": 3335, // Huawei
		"701a04": 2874, // Liteon
		"701ab8": 421,  // Intel
		"701aed": 7754, // Advas
		"701ce7": 421,  // Intel
		"701d7f": 4261, // Comtech
		"701f3c": 152,  // Samsung
		"701f53": 2,    // Cisco
		"702393": 7755, // fos4X
		"702526": 457,  // Nokia
		"702559": 189,  // CyberTAN
		"70288b": 152,  // Samsung
		"702a7d": 7756, // EpSpot
		"702ad5": 152,  // Samsung
		"702b1d": 7757, // E-Domus
		"702c09": 1427, // Nintendo
		"702c1f": 7236, // Wisol
		"702dd1": 7758, // Newings
		"702e22": 3032, // ZTE
		"702f35": 3335, // Huawei
		"702f4b": 7759, // Steelcase
		"702f97": 7760, // Aava
		"703018": 620,  // Avaya
		"70305d": 1103, // Ubiquoss
		"703187": 7761, // ACX
		"703217": 421,  // Intel
		"703509": 2,    // Cisco
		"703811": 300,  // Siemens
		"7038b4": 7762, // Low
		"7038ee": 620,  // Avaya
		"703a0e": 1685, // Aruba
		"703a51": 5698, // Xiaomi
		"703acb": 3522, // Google
		"703c03": 7763, // RadiAnt
		"703c69": 545,  // Apple
		"703e97": 6869, // Iton
		"703eac": 545,  // Apple
		"7040ff": 3335, // Huawei
		"7047e9": 6370, // Vivo
		"70480f": 545,  // Apple
		"7048f7": 1427, // Nintendo
		"704a0e": 4499, // AMPAK
		"704ae4": 7764, // Rinstrum
		"704ca5": 1323, // Fortinet
		"704ced": 7765, // TMRG
		"704d7b": 1806, // ASUS
		"704e01": 7766, // Kwangwon
		"704e6b": 3335, // Huawei
		"704f57": 1595, // TP-Link
		"704fb8": 125,  // ARRIS
		"7050af": 3513, // BSkyB
		"7052c5": 620,  // Avaya
		"7052d8": 6282, // Itel
		"705425": 125,  // ARRIS
		"7054b4": 1449, // Vestel
		"7054d2": 5440, // Pegatron
		"7054f5": 3335, // Huawei
		"7055f8": 7767, // Cerebras
		"705681": 545,  // Apple
		"705812": 2154, // Panasonic
		"705896": 7768, // InShow
		"705a0f": 302,  // HP
		"705a9e": 6393, // Technicolor
		"705aac": 152,  // Samsung
		"705ab6": 358,  // Compal
		"705b2e": 7769, // M2Communication
		"705dcc": 1252, // EFM
		"705fa3": 5698, // Xiaomi
		"7060de": 7770, // LaVision
		"706173": 7771, // Calantec
		"70617b": 2,    // Cisco
		"7061be": 1592, // Wistron
		"7061ee": 7772, // Sunwoda
		"7062b8": 803,  // D-Link
		"70661b": 6860, // Sonova
		"706655": 3005, // Azurewave
		"70695a": 2,    // Cisco
		"706bb9": 2,    // Cisco
		"706d15": 2,    // Cisco
		"706dec": 7773, // Wifi-soft
		"706e6d": 2,    // Cisco
		"706f81": 67,   // Private
		"70700d": 545,  // Apple
		"70704c": 7774, // Purple
		"70708b": 2,    // Cisco
		"7070aa": 5439, // Amazon
		"7071b3": 7775, // Brain
		"7071bc": 5440, // Pegatron
		"70720d": 2665, // Lenovo
		"70723c": 3335, // Huawei
		"7072cf": 7776, // EdgeCore
		"7073cb": 545,  // Apple
		"707414": 2057, // Murata
		"707630": 125,  // ARRIS
		"7076f0": 7777, // LevelOne
		"7076ff": 7778, // Kerlink
		"70788b": 6370, // Vivo
		"707990": 3335, // Huawei
		"7079b3": 2,    // Cisco
		"707be8": 3335, // Huawei
		"707c18": 7779, // ADATA
		"707c69": 620,  // Avaya
		"707db9": 2,    // Cisco
		"707e43": 125,  // ARRIS
		"707ede": 7780, // Nastec
		"708105": 2,    // Cisco
		"7081eb": 545,  // Apple
		"70820e": 7781, // as
		"70828e": 7782, // OleumTech
		"7085c2": 4720, // ASRock
		"7085c4": 5441, // Ruijie
		"7085c6": 125,  // ARRIS
		"70879e": 7783, // Beken
		"7087a7": 2057, // Murata
		"708a09": 3335, // Huawei
		"708b78": 7784, // citygrow
		"708bcd": 1806, // ASUS
		"708cb6": 3335, // Huawei
		"708cbb": 7785, // Mimodisplaykorea
		"708d09": 457,  // Nokia
		"708f47": 6370, // Vivo
		"7090b7": 3335, // Huawei
		"70918f": 7786, // Weber-Stephen
		"7091f3": 3858, // Universal
		"709741": 2640, // Arcadyan
		"709756": 7787, // Happyelectronics
		"709bfc": 7788, // Bryton
		"709c8f": 7789, // Nero
		"709cd1": 421,  // Intel
		"709e29": 101,  // Sony
		"709e86": 7790, // X6D
		"709f2d": 3032, // ZTE
		"709fa9": 6266, // Tecno
		"70a2b3": 545,  // Apple
		"70a6cc": 421,  // Intel
		"70a741": 2979, // Ubiquiti
		"70a84c": 7791, // MONAD
		"70a8d3": 421,  // Intel
		"70a8e3": 3335, // Huawei
		"70aab2": 2221, // BlackBerry
		"70af25": 7792, // Nishiyama
		"70b13d": 152,  // Samsung
		"70b14e": 125,  // ARRIS
		"70b317": 2,    // Cisco
		"70b5e8": 102,  // Dell
		"70b7aa": 6370, // Vivo
		"70b8f6": 6377, // Espressif
		"70b921": 4922, // Fiberhome
		"70bbe9": 5698, // Xiaomi
		"70bc10": 612,  // Microsoft
		"70c7f2": 3335, // Huawei
		"70c833": 7793, // Wirepas
		"70c94e": 2874, // Liteon
		"70c9c6": 2,    // Cisco
		"70ca97": 2738, // Ruckus
		"70ca9b": 2,    // Cisco
		"70cd0d": 421,  // Intel
		"70cd60": 545,  // Apple
		"70ce8c": 152,  // Samsung
		"70cf49": 421,  // Intel
		"70d313": 3335, // Huawei
		"70d379": 2,    // Cisco
		"70d4f2": 4556, // RIM
		"70d57e": 7794, // Scalar
		"70d5e7": 7795, // Wellcore
		"70d6b6": 7796, // Metrum
		"70d923": 6370, // Vivo
		"70d931": 1640, // Cambridge
		"70da9c": 7797, // Tecsen
		"70db98": 2,    // Cisco
		"70dda1": 5028, // Tellabs
		"70ddef": 3335, // Huawei
		"70dee2": 545,  // Apple
		"70df2f": 2,    // Cisco
		"70dff7": 125,  // ARRIS
		"70e027": 7798, // Hongyu
		"70e139": 7799, // 3view
		"70e1fd": 833,  // Flextronics
		"70e284": 1592, // Wistron
		"70e422": 2,    // Cisco
		"70e72c": 545,  // Apple
		"70ea1a": 2,    // Cisco
		"70ea5a": 545,  // Apple
		"70ece4": 545,  // Apple
		"70ee50": 7800, // Netatmo
		"70eea3": 7801, // Eoptolink
		"70ef00": 545,  // Apple
		"70f087": 545,  // Apple
		"70f088": 1427, // Nintendo
		"70f096": 2,    // Cisco
		"70f196": 2247, // Actiontec
		"70f1a1": 2874, // Liteon
		"70f1e5": 7802, // Xetawave
		"70f220": 2247, // Actiontec
		"70f35a": 2,    // Cisco
		"70f754": 4499, // AMPAK
		"70f82b": 7447, // DWnet
		"70f927": 152,  // Samsung
		"70fc8c": 2659, // OneAccess
		"70fd45": 3335, // Huawei
		"70fd46": 152,  // Samsung
		"7403bd": 1077, // Buffalo
		"74042b": 2665, // Lenovo
		"7404f1": 421,  // Intel
		"7405a5": 1595, // TP-Link
		"7409ac": 7803, // Quext
		"740abc": 7804, // LightwaveRF
		"740ae1": 3335, // Huawei
		"740cee": 3335, // Huawei
		"740edb": 7805, // Optowiz
		"7411b2": 2,    // Cisco
		"7412bb": 4922, // Fiberhome
		"741575": 5698, // Xiaomi
		"7415e2": 7806, // Tri-Sen
		"741bb2": 545,  // Apple
		"741c27": 6282, // Itel
		"741e93": 4922, // Fiberhome
		"741f79": 7807, // Youngkook
		"7422bb": 3335, // Huawei
		"742344": 5698, // Xiaomi
		"7426ac": 2,    // Cisco
		"7426ff": 3032, // ZTE
		"74273c": 7808, // ChangYang
		"7427ea": 1115, // Elitegroup
		"742b0f": 7809, // Infinidat
		"742b62": 4,    // Fujitsu
		"742edb": 7810, // Perinet
		"742efc": 7811, // DirectPacket
		"742f68": 3005, // Azurewave
		"743170": 2640, // Arcadyan
		"7432c2": 7812, // Kyolis
		"743400": 7813, // MTG
		"74373b": 7814, // UNINET
		"7438b7": 87,   // Canon
		"743a65": 48,   // NEC
		"743aef": 1269, // Kaonmedia
		"743e2b": 2738, // Ruckus
		"743ecb": 7815, // Gentrice
		"7440be": 869,  // LG
		"74428b": 545,  // Apple
		"744401": 1368, // Netgear
		"74452d": 3335, // Huawei
		"74458a": 152,  // Samsung
		"7445ce": 3283, // Cresyn
		"744687": 7816, // Kingsignal
		"7446a0": 302,  // HP
		"744aa4": 3032, // ZTE
		"744ca1": 2874, // Liteon
		"744d28": 1784, // Routerboard.com
		"744d79": 7817, // Arrive
		"7451ba": 5698, // Xiaomi
		"745327": 7818, // Commsen
		"745612": 125,  // ARRIS
		"7458f3": 5439, // Amazon
		"745909": 3335, // Huawei
		"745933": 7819, // Danal
		"745aaa": 3335, // Huawei
		"745c9f": 6434, // TCT
		"745d68": 4922, // Fiberhome
		"745e1c": 6118, // Pioneer
		"745f90": 7820, // LAM
		"7460fa": 3335, // Huawei
		"7463c2": 3335, // Huawei
		"7463df": 7821, // VTS
		"74650c": 545,  // Apple
		"7465d1": 7822, // Atlinks
		"7467f7": 185,  // Extreme
		"746a3a": 7823, // Aperi
		"746a89": 7824, // Rezolt
		"746b82": 7825, // Movek
		"746f19": 7826, // Icarvisions
		"746f3d": 5483, // Contec
		"746ff7": 1592, // Wistron
		"747069": 3335, // Huawei
		"7470fd": 421,  // Intel
		"74721e": 7827, // Edison
		"7472f2": 7828, // Chipsip
		"74731d": 7829, // ifm
		"747336": 7830, // MICRODIGTAL
		"747446": 3522, // Google
		"747548": 5439, // Amazon
		"747818": 7831, // Jurumani
		"747827": 102,  // Dell
		"747a90": 2057, // Murata
		"747b7a": 7832, // ETH
		"747d24": 6849, // Phicomm
		"747db6": 7833, // Aliwei
		"748114": 545,  // Apple
		"7483c2": 2979, // Ubiquiti
		"7483ef": 3794, // Arista
		"74852a": 5440, // Pegatron
		"74860b": 2,    // Cisco
		"74867a": 102,  // Dell
		"7486e2": 102,  // Dell
		"7487a9": 7834, // OCT
		"7487bb": 4565, // Ciena
		"74882a": 3335, // Huawei
		"7488bb": 2,    // Cisco
		"748a0d": 125,  // ARRIS
		"748b29": 7835, // Micobiomed
		"748d08": 545,  // Apple
		"748e08": 7836, // Bestek
		"748ef8": 90,   // Brocade
		"748f3c": 545,  // Apple
		"74901f": 7837, // Ragile
		"749050": 5696, // Renesas
		"74911a": 2738, // Ruckus
		"7491bd": 7838, // Four
		"7493a4": 768,  // Zebra
		"74943d": 7839, // AgJunction
		"7495ec": 430,  // Alpsalpine
		"749637": 7840, // Todaair
		"749781": 3032, // ZTE
		"749975": 372,  // IBM
		"749ac0": 7841, // Cachengo
		"749be8": 870,  // Hitron
		"749d79": 2075, // Sercomm
		"749d8f": 3335, // Huawei
		"749ddc": 1939, // 2Wire
		"749ea5": 6578, // Ohsung
		"749eaf": 545,  // Apple
		"749ef5": 152,  // Samsung
		"74a02f": 2,    // Cisco
		"74a063": 3335, // Huawei
		"74a2e6": 2,    // Cisco
		"74a34a": 7738, // Zimi
		"74a528": 3335, // Huawei
		"74a722": 869,  // LG
		"74a78e": 3032, // ZTE
		"74a7ea": 5439, // Amazon
		"74acb9": 2979, // Ubiquiti
		"74ad98": 2,    // Cisco
		"74b472": 7842, // Ciesse
		"74b57e": 3032, // ZTE
		"74b587": 545,  // Apple
		"74b6b6": 5821, // eero
		"74b7e6": 7843, // Zegna-Daidong
		"74b9eb": 6946, // JinQianMao
		"74be08": 6992, // ATEK
		"74bfa1": 7844, // Hyunteck
		"74bfb7": 7845, // Nusoft
		"74bfc0": 87,   // Canon
		"74c14f": 3335, // Huawei
		"74c17d": 6296, // Infinix
		"74c246": 5439, // Amazon
		"74c63b": 3005, // Azurewave
		"74c99a": 306,  // Ericsson
		"74c9a3": 4922, // Fiberhome
		"74ca25": 7846, // Calxeda
		"74cbf3": 647,  // Lava
		"74cc39": 4922, // Fiberhome
		"74d02b": 1806, // ASUS
		"74d0dc": 306,  // Ericsson
		"74d21d": 3335, // Huawei
		"74d435": 1929, // Giga-Byte
		"74d637": 5439, // Amazon
		"74d654": 7847, // Gint
		"74d7ca": 2154, // Panasonic
		"74d83e": 421,  // Intel
		"74d850": 7848, // Evrisko
		"74da38": 115,  // Edimax
		"74da88": 1595, // TP-Link
		"74dada": 803,  // D-Link
		"74dbd1": 7849, // Ebay
		"74de2b": 2874, // Liteon
		"74dfbf": 2874, // Liteon
		"74e06e": 7850, // Ergophone
		"74e19a": 4922, // Fiberhome
		"74e1b6": 545,  // Apple
		"74e20c": 5439, // Amazon
		"74e277": 7851, // Vizmonet
		"74e28c": 612,  // Microsoft
		"74e2f5": 545,  // Apple
		"74e424": 7852, // Apiste
		"74e50b": 421,  // Intel
		"74e537": 7853, // Radspin
		"74e543": 2874, // Liteon
		"74e5f9": 421,  // Intel
		"74e60f": 6266, // Tecno
		"74e6b8": 869,  // LG
		"74e6e2": 102,  // Dell
		"74e7c6": 125,  // ARRIS
		"74e9bf": 3335, // Huawei
		"74ea3a": 1595, // TP-Link
		"74eae8": 125,  // ARRIS
		"74eb80": 152,  // Samsung
		"74ec42": 4922, // Fiberhome
		"74ecb2": 5439, // Amazon
		"74ecf1": 7854, // Acumen
		"74f06d": 3005, // Azurewave
		"74f07d": 7855, // BnCOM
		"74f2fa": 5698, // Xiaomi
		"74f612": 125,  // ARRIS
		"74f61c": 1341, // HTC
		"74f661": 56,   // Schneider
		"74f737": 7856, // KCE
		"74f91a": 7857, // Onface
		"74f9ca": 1427, // Nintendo
		"74fda0": 7858, // Compupal
		"74fe48": 1703, // Advantech
		"78009e": 152,  // Samsung
		"78028b": 545,  // Apple
		"7802b1": 2,    // Cisco
		"7802f8": 5698, // Xiaomi
		"78047a": 2843, // Edge
		"7804e3": 3335, // Huawei
		"78058c": 7859, // mMax
		"7806c9": 3335, // Huawei
		"780cb8": 421,  // Intel
		"780cf0": 2,    // Cisco
		"781100": 7860, // Quantumsolution
		"7811dc": 6281, // XIAOMI
		"7812b8": 7861, // Orantek
		"7817be": 3335, // Huawei
		"781881": 3005, // Azurewave
		"7818a8": 3335, // Huawei
		"78192e": 7862, // NASCENT
		"7819f7": 826,  // Juniper
		"781c5a": 3206, // Sharp
		"781d4a": 3032, // ZTE
		"781dba": 3335, // Huawei
		"781dfd": 7863, // Jabil
		"781fdb": 152,  // Samsung
		"782079": 7864, // ID
		"7820a5": 1427, // Nintendo
		"782184": 6377, // Espressif
		"78223d": 7865, // Affirmed
		"782327": 152,  // Samsung
		"7823ae": 125,  // ARRIS
		"7824af": 1806, // ASUS
		"782544": 7866, // Omnima
		"7825ad": 152,  // Samsung
		"7828ca": 2048, // Sonos
		"7829ed": 2545, // Askey
		"782b46": 421,  // Intel
		"782b64": 1819, // Bose
		"782bcb": 102,  // Dell
		"782d7e": 2905, // TRENDnet
		"782eef": 457,  // Nokia
		"782f17": 7867, // Xlab
		"78303b": 7868, // Stephen
		"7830e1": 7869, // UltraClenz
		"78312b": 3032, // ZTE
		"7831c1": 545,  // Apple
		"78321b": 803,  // D-Link
		"7835a0": 7870, // Zurn
		"783607": 2868, // Cermate
		"783690": 3097, // Yulong
		"783716": 152,  // Samsung
		"783a6c": 6266, // Tecno
		"783a84": 545,  // Apple
		"783ce3": 7871, // Kai-EE
		"783e53": 3513, // BSkyB
		"783f15": 7872, // EasySYNC
		"7840e4": 152,  // Samsung
		"784405": 7873, // FUJITU
		"784476": 2128, // Zioncom
		"7844fd": 1595, // TP-Link
		"784501": 7874, // Biamp
		"784558": 2979, // Ubiquiti
		"784561": 189,  // CyberTAN
		"7845b3": 3335, // Huawei
		"7845c4": 102,  // Dell
		"7846d4": 152,  // Samsung
		"78471d": 152,  // Samsung
		"784859": 302,  // HP
		"78491d": 7875, // Will-Burt
		"784b87": 2057, // Murata
		"784f43": 545,  // Apple
		"784f9b": 826,  // Juniper
		"785005": 7876, // MOKO
		"78507c": 826,  // Juniper
		"78510c": 7877, // LiveU
		"78521a": 152,  // Samsung
		"78524a": 7878, // Optonic
		"785364": 7879, // SHIFT
		"7853f2": 7880, // Roxton
		"78542e": 803,  // D-Link
		"785517": 7881, // SankyuElectronics
		"785773": 3335, // Huawei
		"785860": 3335, // Huawei
		"7858f3": 7882, // Vachen
		"78595e": 152,  // Samsung
		"785c72": 7883, // Hioso
		"785dc8": 869,  // LG
		"785ea2": 3940, // Sunitec
		"786256": 3335, // Huawei
		"7864c0": 545,  // Apple
		"786559": 2049, // Sagemcom
		"78670e": 1592, // Wistron
		"7867d7": 545,  // Apple
		"7868f7": 7451, // YSTen
		"786a1f": 125,  // ARRIS
		"786a89": 3335, // Huawei
		"786c1c": 545,  // Apple
		"787052": 7884, // Welotec
		"78719c": 125,  // ARRIS
		"78725d": 2,    // Cisco
		"787a6f": 7885, // Juice
		"787b8a": 545,  // Apple
		"787d48": 6282, // Itel
		"787d53": 185,  // Extreme
		"787df3": 7886, // Sterlite
		"787e61": 545,  // Apple
		"788102": 2075, // Sercomm
		"78843c": 101,  // Sony
		"7885f4": 3335, // Huawei
		"78886d": 545,  // Apple
		"788973": 5286, // CMC
		"788a20": 2979, // Ubiquiti
		"788c4d": 1795, // Indyme
		"788c54": 4358, // Ping
		"788c77": 613,  // Lexmark
		"788df7": 870,  // Hitron
		"7890a2": 3032, // ZTE
		"7891e9": 2051, // Raisecom
		"78923e": 457,  // Nokia
		"78929c": 421,  // Intel
		"7894b4": 2075, // Sercomm
		"7895eb": 6282, // Itel
		"789682": 3032, // ZTE
		"789684": 125,  // ARRIS
		"7898e8": 803,  // D-Link
		"7898fd": 7887, // Q9
		"78995c": 7888, // Nationz
		"789966": 7889, // Musilab
		"789ed0": 152,  // Samsung
		"789f70": 545,  // Apple
		"789f87": 300,  // Siemens
		"789faa": 3335, // Huawei
		"78a03f": 5439, // Amazon
		"78a051": 7890, // iiNet
		"78a106": 1595, // TP-Link
		"78a183": 7891, // Advidia
		"78a2a0": 1427, // Nintendo
		"78a3e4": 545,  // Apple
		"78a683": 7892, // Precidata
		"78a6e1": 90,   // Brocade
		"78a714": 1475, // Amphenol
		"78a7eb": 6628, // 1MORE
		"78a873": 152,  // Samsung
		"78abbb": 152,  // Samsung
		"78ac44": 102,  // Dell
		"78acbf": 7893, // Igneous
		"78acc0": 302,  // HP
		"78af08": 421,  // Intel
		"78af58": 7894, // Gimasi
		"78b213": 7447, // DWnet
		"78b46a": 3335, // Huawei
		"78b554": 3335, // Huawei
		"78b8d6": 768,  // Zebra
		"78bad0": 7895, // Shinybow
		"78baf9": 2,    // Cisco
		"78bb88": 7158, // Maxio
		"78bc1a": 2,    // Cisco
		"78bdbc": 152,  // Samsung
		"78bebd": 7896, // STULZ
		"78c1a7": 3032, // ZTE
		"78c3e9": 152,  // Samsung
		"78c5f8": 3335, // Huawei
		"78c881": 101,  // Sony
		"78ca04": 457,  // Nokia
		"78ca39": 545,  // Apple
		"78ca5e": 7897, // Elno
		"78cb33": 7898, // DHC
		"78cc2b": 7899, // Sinewy
		"78cd8e": 741,  // SMC
		"78cf2f": 3335, // Huawei
		"78cff9": 3335, // Huawei
		"78d004": 7900, // Neousys
		"78d129": 7901, // Vicos
		"78d162": 545,  // Apple
		"78d294": 1368, // Netgear
		"78d347": 306,  // Ericsson
		"78d34f": 7902, // Pace-O-Matic
		"78d3ed": 7903, // Norma
		"78d6b2": 35,   // Toshiba
		"78d6dc": 687,  // Motorola
		"78d6f0": 152,  // Samsung
		"78d71a": 4565, // Ciena
		"78d752": 3335, // Huawei
		"78d75f": 545,  // Apple
		"78da6e": 2,    // Cisco
		"78daa2": 7904, // Cynosure
		"78dab3": 7905, // GBO
		"78dd12": 2640, // Arcadyan
		"78dd33": 3335, // Huawei
		"78ddd6": 7906, // c-scape
		"78e103": 5439, // Amazon
		"78e22c": 3335, // Huawei
		"78e36d": 6377, // Espressif
		"78e3b5": 302,  // HP
		"78e3de": 545,  // Apple
		"78e7d1": 302,  // HP
		"78e8b6": 3032, // ZTE
		"78e980": 7907, // RainUs
		"78eb46": 3335, // Huawei
		"78ec74": 7908, // Kyland-USA
		"78ef4c": 7909, // Unetconvergence
		"78f09b": 3335, // Huawei
		"78f238": 152,  // Samsung
		"78f29e": 5440, // Pegatron
		"78f557": 3335, // Huawei
		"78f5fd": 3335, // Huawei
		"78f7be": 152,  // Samsung
		"78f7d0": 7910, // Silverbrook
		"78f882": 869,  // LG
		"78f944": 67,   // Private
		"78fbd8": 545,  // Apple
		"78fd94": 545,  // Apple
		"78fe3d": 826,  // Juniper
		"78fe41": 7911, // Socus
		"78ff57": 421,  // Intel
		"78ffca": 6266, // Tecno
		"7c004d": 3335, // Huawei
		"7c0191": 545,  // Apple
		"7c02bc": 7912, // Hansung
		"7c034c": 2049, // Sagemcom
		"7c035e": 5698, // Xiaomi
		"7c03ab": 5698, // Xiaomi
		"7c03d8": 2049, // Sagemcom
		"7c04d0": 545,  // Apple
		"7c0507": 5440, // Pegatron
		"7c051e": 7913, // Rafael
		"7c0623": 3430, // Ultra
		"7c092b": 7914, // Bekey
		"7c0a3f": 152,  // Samsung
		"7c0a50": 7915, // J-MEX
		"7c0bc6": 152,  // Samsung
		"7c0ece": 2,    // Cisco
		"7c10c9": 1806, // ASUS
		"7c11be": 545,  // Apple
		"7c11cb": 3335, // Huawei
		"7c11cd": 7916, // QianTang
		"7c131d": 6256, // Sernet
		"7c1689": 2049, // Sagemcom
		"7c18cd": 7917, // E-TRON
		"7c1a03": 7918, // 8Locations
		"7c1ac0": 3335, // Huawei
		"7c1b93": 3335, // Huawei
		"7c1c4e": 869,  // LG
		"7c1c68": 152,  // Samsung
		"7c1cf1": 3335, // Huawei
		"7c1dd9": 5698, // Xiaomi
		"7c1e52": 612,  // Microsoft
		"7c2048": 7919, // KoamTac
		"7c210d": 2,    // Cisco
		"7c210e": 2,    // Cisco
		"7c214a": 421,  // Intel
		"7c2302": 152,  // Samsung
		"7c240c": 7076, // Telechips
		"7c2499": 545,  // Apple
		"7c2586": 826,  // Juniper
		"7c2587": 7920, // chaowifi.com
		"7c25da": 6441, // FN-Link
		"7c2634": 125,  // ARRIS
		"7c2664": 2049, // Sagemcom
		"7c2a31": 421,  // Intel
		"7c2adb": 5698, // Xiaomi
		"7c2ebd": 3522, // Google
		"7c2edd": 152,  // Samsung
		"7c2f80": 4299, // Gigaset
		"7c310e": 2,    // Cisco
		"7c336e": 7921, // MEG
		"7c33f9": 3335, // Huawei
		"7c38ad": 152,  // Samsung
		"7c3953": 3032, // ZTE
		"7c3985": 3335, // Huawei
		"7c3d2b": 3335, // Huawei
		"7c3e74": 3335, // Huawei
		"7c3e9d": 7922, // Patech
		"7c41a2": 457,  // Nokia
		"7c438f": 7923, // E-Band
		"7c444c": 7924, // Entertainment
		"7c4685": 687,  // Motorola
		"7c49eb": 6281, // XIAOMI
		"7c4a82": 7925, // Portsmith
		"7c4ca5": 3513, // BSkyB
		"7c4f7d": 7926, // Sawwave
		"7c4fb5": 2640, // Arcadyan
		"7c5049": 545,  // Apple
		"7c5079": 421,  // Intel
		"7c50da": 67,   // Private
		"7c534a": 7927, // Metamako
		"7c55a7": 7928, // Kastle
		"7c55e7": 7929, // YSI
		"7c573c": 1685, // Aruba
		"7c574e": 7930, // CoBI
		"7c5a1c": 3579, // Sophos
		"7c5a67": 7931, // JNC
		"7c5cf8": 421,  // Intel
		"7c604a": 7932, // Avelon
		"7c6097": 3335, // Huawei
		"7c6166": 5439, // Amazon
		"7c6193": 1341, // HTC
		"7c6456": 152,  // Samsung
		"7c67a2": 421,  // Intel
		"7c696b": 7933, // Atmosic
		"7c69f6": 2,    // Cisco
		"7c6ab3": 7934, // IBC
		"7c6ac3": 7935, // GatesAir
		"7c6adb": 7936, // SafeTone
		"7c6b33": 7937, // Tenyu
		"7c6bf7": 7938, // NTI
		"7c6d62": 545,  // Apple
		"7c6df8": 545,  // Apple
		"7c70db": 421,  // Intel
		"7c726e": 306,  // Ericsson
		"7c72e4": 7939, // Unikey
		"7c73eb": 3335, // Huawei
		"7c7635": 421,  // Intel
		"7c7668": 3335, // Huawei
		"7c7673": 7940, // ENMAS
		"7c7716": 2693, // ZyXEL
		"7c787e": 152,  // Samsung
		"7c78b2": 6882, // Wyze
		"7c79e8": 7941, // PayRange
		"7c7a53": 7942, // Phytrex
		"7c7a91": 421,  // Intel
		"7c7d3d": 3335, // Huawei
		"7c7d41": 7943, // Jinmuyu
		"7c822d": 7944, // Nortec
		"7c8530": 457,  // Nokia
		"7c87ce": 6377, // Espressif
		"7c8956": 152,  // Samsung
		"7c8ac0": 7945, // EVBox
		"7c8ae1": 358,  // Compal
		"7c8bb5": 152,  // Samsung
		"7c8bca": 1595, // TP-Link
		"7c8fde": 7447, // DWnet
		"7c9122": 152,  // Samsung
		"7c942a": 3335, // Huawei
		"7c95b1": 185,  // Extreme
		"7c95f3": 2,    // Cisco
		"7c96d2": 6299, // Fihonest
		"7c97e1": 3335, // Huawei
		"7c9a1d": 545,  // Apple
		"7c9a54": 6393, // Technicolor
		"7c9ebd": 6377, // Espressif
		"7ca177": 3335, // Huawei
		"7ca1ae": 545,  // Apple
		"7ca23e": 3335, // Huawei
		"7ca29b": 7946, // D.SignT
		"7ca61d": 7947, // MHL
		"7ca62a": 302,  // HP
		"7ca96b": 7023, // Syrotech
		"7ca97d": 7948, // Objenious
		"7cab25": 7949, // Mesmo
		"7cab60": 545,  // Apple
		"7cad4f": 2,    // Cisco
		"7cad74": 2,    // Cisco
		"7cb03e": 7950, // OSRAM
		"7cb0c2": 421,  // Intel
		"7cb15d": 3335, // Huawei
		"7cb177": 7951, // Satelco
		"7cb25c": 5768, // Acacia
		"7cb27d": 421,  // Intel
		"7cb542": 7952, // ACES
		"7cb566": 421,  // Intel
		"7cb59b": 1595, // TP-Link
		"7cb733": 2545, // Askey
		"7cb77b": 2709, // Paradigm
		"7cbb6f": 7953, // Cosco
		"7cbb8a": 1427, // Nintendo
		"7cbf88": 7954, // Mobilicom
		"7cbfb1": 125,  // ARRIS
		"7cc180": 545,  // Apple
		"7cc225": 152,  // Samsung
		"7cc2c6": 1595, // TP-Link
		"7cc385": 3335, // Huawei
		"7cc3a1": 545,  // Apple
		"7cc4ef": 7955, // Devialet
		"7cc537": 545,  // Apple
		"7cc77e": 4922, // Fiberhome
		"7cc8d7": 7956, // Damalisk
		"7cc95a": 102,  // Dell
		"7ccb0d": 7957, // Antaira
		"7cccb8": 421,  // Intel
		"7ccd11": 7958, // MS-Magnet
		"7cd1c3": 545,  // Apple
		"7cd30a": 3979, // Inventec
		"7cd566": 5439, // Amazon
		"7cd661": 5698, // Xiaomi
		"7cd762": 7959, // Freestyle
		"7cd844": 7960, // Enmotus
		"7cd95c": 3522, // Google
		"7cd9a0": 3335, // Huawei
		"7cda84": 7961, // Dongnian
		"7cdb98": 2545, // Askey
		"7cdde9": 7962, // ATOM
		"7cdfa1": 6377, // Espressif
		"7ce044": 7963, // NEON
		"7ce2ca": 826,  // Juniper
		"7ce4aa": 67,   // Private
		"7ce524": 7964, // Quirky
		"7ce97c": 6282, // Itel
		"7ceb7f": 7965, // Dmet
		"7cebea": 7966, // Asct
		"7ced8d": 612,  // Microsoft
		"7cef8a": 7967, // Inhon
		"7cf05f": 545,  // Apple
		"7cf2dd": 7968, // Vence
		"7cf31b": 869,  // LG
		"7cf429": 7969, // NUUO
		"7cf854": 152,  // Samsung
		"7cf880": 2,    // Cisco
		"7cf90e": 152,  // Samsung
		"7cf9a0": 4922, // Fiberhome
		"7cfadf": 545,  // Apple
		"7cfc16": 545,  // Apple
		"7cfc3c": 1399, // Visteon
		"7cfd6b": 5698, // Xiaomi
		"7cfe28": 7970, // Salutron
		"7cfe90": 432,  // Mellanox
		"80000b": 421,  // Intel
		"800010": 1057, // AT&T
		"80006e": 545,  // Apple
		"80015c": 7971, // Synaptics
		"800184": 1341, // HTC
		"80029c": 1450, // Gemtek
		"8002df": 7972, // ORA
		"800384": 2738, // Ruckus
		"80045f": 545,  // Apple
		"800588": 5441, // Ruijie
		"8007a2": 7973, // Esson
		"800902": 3600, // Keysight
		"800a06": 7974, // CoMTEC
		"800c67": 545,  // Apple
		"800cf9": 5439, // Amazon
		"800dd7": 7975, // Latticework
		"800e24": 7976, // ForgetBox
		"801382": 3335, // Huawei
		"80177d": 76,   // Nortel
		"801844": 102,  // Dell
		"8018a7": 152,  // Samsung
		"801934": 421,  // Intel
		"8019fe": 7977, // JianLing
		"801daa": 620,  // Avaya
		"801f02": 115,  // Edimax
		"801f12": 706,  // Microchip
		"8020da": 2049, // Sagemcom
		"8020fd": 152,  // Samsung
		"80248f": 2,    // Cisco
		"802511": 6282, // Itel
		"802689": 803,  // D-Link
		"802994": 6393, // Technicolor
		"802aa8": 2979, // Ubiquiti
		"802afa": 7978, // Germaneers
		"802dbf": 2,    // Cisco
		"802de1": 7979, // Solarbridge
		"802e14": 7980, // azeti
		"803049": 2874, // Liteon
		"8030e0": 302,  // HP
		"8031f0": 152,  // Samsung
		"803253": 421,  // Intel
		"803428": 706,  // Microchip
		"803457": 7981, // OT
		"8035c1": 5698, // Xiaomi
		"803773": 1368, // Netgear
		"803896": 3206, // Sharp
		"8038bc": 3335, // Huawei
		"8038fb": 421,  // Intel
		"8038fd": 7982, // LeapFrog
		"8039e5": 7983, // Patlite
		"803a59": 1057, // AT&T
		"803af4": 4922, // Fiberhome
		"803b9a": 7984, // ghe-ces
		"803f5d": 7985, // Winstars
		"804126": 3335, // Huawei
		"8045dd": 421,  // Intel
		"804786": 152,  // Samsung
		"804971": 545,  // Apple
		"804a14": 545,  // Apple
		"804b50": 1655, // Silicon
		"804e70": 152,  // Samsung
		"804e81": 152,  // Samsung
		"804f58": 7986, // ThinkEco
		"80501b": 457,  // Nokia
		"8050f6": 6282, // Itel
		"8054d9": 3335, // Huawei
		"805719": 152,  // Samsung
		"8058f8": 687,  // Motorola
		"8059fd": 7987, // Noviga
		"805a04": 869,  // LG
		"805e4f": 6441, // FN-Link
		"805fc5": 545,  // Apple
		"806007": 4556, // RIM
		"806459": 7988, // Nimbus
		"80656d": 152,  // Samsung
		"80657c": 545,  // Apple
		"8065e9": 7989, // BenQ
		"806629": 7990, // Prescope
		"80691a": 2469, // Belkin
		"806933": 3335, // Huawei
		"806940": 7991, // Lexar
		"806a00": 2,    // Cisco
		"806c1b": 687,  // Motorola
		"806d71": 5439, // Amazon
		"806d97": 67,   // Private
		"806f1c": 3335, // Huawei
		"80711f": 826,  // Juniper
		"80717a": 3335, // Huawei
		"807215": 3513, // BSkyB
		"807264": 3335, // Huawei
		"80739f": 2849, // Kyocera
		"807459": 7992, // K's
		"80751f": 3513, // BSkyB
		"807693": 7993, // Newag
		"8077a4": 6266, // Tecno
		"807871": 2545, // Askey
		"80795d": 6296, // Infinix
		"807abf": 1341, // HTC
		"807b3e": 152,  // Samsung
		"807d14": 3335, // Huawei
		"807d1b": 7994, // Neosystem
		"807d3a": 6377, // Espressif
		"807ff8": 826,  // Juniper
		"808223": 545,  // Apple
		"808287": 7995, // ATCOM
		"8084a9": 7996, // oshkosh
		"808698": 7997, // Netronics
		"8086d9": 152,  // Samsung
		"8086f2": 421,  // Intel
		"808917": 1595, // TP-Link
		"808a8b": 6370, // Vivo
		"808abd": 152,  // Samsung
		"808af7": 7998, // Nanoleaf
		"808c97": 1269, // Kaonmedia
		"808db7": 302,  // HP
		"808f1d": 1595, // TP-Link
		"808fe8": 3542, // Intelbras
		"809133": 3005, // Azurewave
		"8091c0": 7999, // AgileMesh
		"80929f": 545,  // Apple
		"809393": 8000, // Xapt
		"809621": 2665, // Lenovo
		"8096b1": 125,  // ARRIS
		"809b20": 421,  // Intel
		"809fab": 4922, // Fiberhome
		"809ff5": 152,  // Samsung
		"80a1ab": 8001, // Intellisis
		"80a235": 6295, // Edgecore
		"80a589": 3005, // Azurewave
		"80a796": 8002, // Neurotek
		"80aaa4": 8003, // Usag
		"80acac": 826,  // Juniper
		"80ad16": 5698, // Xiaomi
		"80ad67": 2135, // Kasda
		"80b03d": 545,  // Apple
		"80b07b": 3032, // ZTE
		"80b234": 6393, // Technicolor
		"80b289": 8004, // Forworld
		"80b575": 3335, // Huawei
		"80b624": 8005, // IVS
		"80b655": 421,  // Intel
		"80b686": 3335, // Huawei
		"80b709": 8006, // Viptela
		"80b95c": 8007, // ELFTECH
		"80b97a": 5821, // eero
		"80baac": 8008, // TeleAdapt
		"80bae6": 8009, // Neets
		"80bbeb": 8010, // Satmap
		"80be05": 545,  // Apple
		"80c16e": 302,  // HP
		"80c5e6": 612,  // Microsoft
		"80c5f2": 3005, // Azurewave
		"80c6ab": 6393, // Technicolor
		"80c755": 2154, // Panasonic
		"80c7c5": 4922, // Fiberhome
		"80c862": 8011, // Openpeak
		"80cc12": 3335, // Huawei
		"80cc9c": 1368, // Netgear
		"80ce62": 302,  // HP
		"80ceb9": 152,  // Samsung
		"80cf41": 2665, // Lenovo
		"80cfa2": 3335, // Huawei
		"80d019": 8012, // Embed
		"80d04a": 6393, // Technicolor
		"80d065": 8013, // CKS
		"80d09b": 3335, // Huawei
		"80d21d": 3005, // Azurewave
		"80d2e5": 1427, // Nintendo
		"80d336": 6314, // Cern
		"80d433": 8014, // LzLabs
		"80d4a5": 3335, // Huawei
		"80d605": 545,  // Apple
		"80da13": 5821, // eero
		"80dabc": 6711, // Megafone
		"80e01d": 2,    // Cisco
		"80e1bf": 3335, // Huawei
		"80e540": 125,  // ARRIS
		"80e650": 545,  // Apple
		"80e82c": 302,  // HP
		"80e86f": 2,    // Cisco
		"80ea07": 1595, // TP-Link
		"80ea23": 1592, // Wistron
		"80ea96": 545,  // Apple
		"80eb77": 1592, // Wistron
		"80ed2c": 545,  // Apple
		"80ee73": 4938, // Shuttle
		"80f1f1": 8015, // Tech4home
		"80f25e": 8016, // Kyynel
		"80f3ef": 7235, // Facebook
		"80f503": 125,  // ARRIS
		"80f8eb": 8017, // RayTight
		"80fa5b": 5688, // Clevo
		"80fb06": 3335, // Huawei
		"80fd7a": 7461, // BLU
		"80ffa8": 8018, // Unidis
		"84002d": 5440, // Pegatron
		"8400d2": 101,  // Sony
		"840112": 1269, // Kaonmedia
		"840283": 531,  // HUMAX
		"840328": 826,  // Juniper
		"8406fa": 4922, // Fiberhome
		"840b2d": 152,  // Samsung
		"840b7c": 870,  // Hitron
		"840d8e": 6377, // Espressif
		"84100d": 687,  // Motorola
		"84119e": 152,  // Samsung
		"84139f": 3032, // ZTE
		"84144d": 421,  // Intel
		"8415d3": 3335, // Huawei
		"84160c": 858,  // Broadcom
		"8416f9": 1595, // TP-Link
		"841715": 8019, // GP
		"8417ef": 6393, // Technicolor
		"841826": 8020, // Osram
		"84183a": 2738, // Ruckus
		"841888": 826,  // Juniper
		"841b5e": 1368, // Netgear
		"841b77": 421,  // Intel
		"841c70": 3032, // ZTE
		"841e26": 8021, // KERNEL-I
		"841ea3": 2049, // Sagemcom
		"8421f1": 3335, // Huawei
		"842388": 2738, // Ruckus
		"84248d": 768,  // Zebra
		"842519": 152,  // Samsung
		"84253f": 8022, // silex
		"8425a4": 8023, // Tariox
		"8425db": 152,  // Samsung
		"84262b": 457,  // Nokia
		"84285a": 8024, // Saffron
		"842999": 545,  // Apple
		"842afd": 302,  // HP
		"842b2b": 102,  // Dell
		"842b50": 8025, // Huria
		"842bbc": 8026, // Modelleisenbahn
		"842e14": 1655, // Silicon
		"842e27": 152,  // Samsung
		"8430e5": 8027, // SkyHawke
		"843497": 302,  // HP
		"8437d5": 152,  // Samsung
		"843835": 545,  // Apple
		"843838": 152,  // Samsung
		"843a4b": 421,  // Intel
		"843a5b": 3979, // Inventec
		"843b10": 8028, // Lvswitches
		"843dc6": 2,    // Cisco
		"843e92": 3335, // Huawei
		"843f4e": 8029, // Tri-Tech
		"844076": 8030, // Drivenets
		"844167": 545,  // Apple
		"844464": 8031, // ServerU
		"8446fe": 3335, // Huawei
		"844765": 3335, // Huawei
		"844823": 8032, // WOXTER
		"844915": 8033, // vArmour
		"844f03": 8034, // Ablelink
		"845181": 152,  // Samsung
		"8454df": 3335, // Huawei
		"8455a5": 152,  // Samsung
		"845733": 612,  // Microsoft
		"845a81": 8035, // ffly4u
		"845b12": 3335, // Huawei
		"845c93": 8036, // Chabrier
		"845cf3": 421,  // Intel
		"845f04": 152,  // Samsung
		"846082": 67,   // Private
		"8461a0": 125,  // ARRIS
		"8462a6": 3102, // EuroCB
		"8463d6": 612,  // Microsoft
		"84683e": 421,  // Intel
		"846878": 545,  // Apple
		"846991": 457,  // Nokia
		"847127": 1655, // Silicon
		"84716a": 3335, // Huawei
		"847207": 8037, // I&C
		"84742a": 3032, // ZTE
		"847460": 3032, // ZTE
		"847637": 3335, // Huawei
		"847778": 8038, // Cochlear
		"84788b": 545,  // Apple
		"8478ac": 2,    // Cisco
		"847933": 8039, // profichip
		"847a88": 1341, // HTC
		"847ab6": 7032, // AltoBeam
		"847b57": 421,  // Intel
		"847beb": 102,  // Dell
		"84802d": 2,    // Cisco
		"848094": 8040, // Meter
		"848102": 4922, // Fiberhome
		"848336": 8041, // Newrun
		"848371": 620,  // Avaya
		"848433": 8042, // Paradox
		"848506": 545,  // Apple
		"8486f3": 8043, // Greenvity
		"8489ad": 545,  // Apple
		"848a8d": 2,    // Cisco
		"848c8d": 545,  // Apple
		"848d84": 8044, // Rajant
		"848e0c": 545,  // Apple
		"848e96": 8045, // Embertec
		"848edf": 101,  // Sony
		"848f69": 102,  // Dell
		"84930c": 8046, // Incoax
		"8493a0": 3335, // Huawei
		"84948c": 870,  // Hitron
		"849681": 8047, // Cathay
		"8496d8": 125,  // ARRIS
		"8497b8": 8048, // Memjet
		"849866": 152,  // Samsung
		"849ca6": 2640, // Arcadyan
		"849d64": 741,  // SMC
		"849fb5": 3335, // Huawei
		"84a06e": 2049, // Sagemcom
		"84a134": 545,  // Apple
		"84a1d1": 2049, // Sagemcom
		"84a3b5": 8049, // Propulsion
		"84a423": 2049, // Sagemcom
		"84a466": 152,  // Samsung
		"84a6c8": 421,  // Intel
		"84a788": 8050, // Perples
		"84a8e4": 3335, // Huawei
		"84a938": 4920, // LCFC
		"84a93e": 302,  // HP
		"84a9c4": 3335, // Huawei
		"84aa9c": 6429, // MitraStar
		"84ab1a": 545,  // Apple
		"84ab26": 6540, // Tiinlab
		"84ac16": 545,  // Apple
		"84ad58": 3335, // Huawei
		"84ad8d": 545,  // Apple
		"84afec": 1077, // Buffalo
		"84b153": 545,  // Apple
		"84b261": 2,    // Cisco
		"84b31b": 8051, // Kinexon
		"84b4db": 1655, // Silicon
		"84b517": 2,    // Cisco
		"84b541": 152,  // Samsung
		"84b59c": 826,  // Juniper
		"84b802": 2,    // Cisco
		"84b8b8": 687,  // Motorola
		"84ba20": 1655, // Silicon
		"84ba3b": 87,   // Canon
		"84bb69": 125,  // ARRIS
		"84be52": 3335, // Huawei
		"84c0ef": 152,  // Samsung
		"84c1c1": 826,  // Juniper
		"84c3e8": 8052, // Vaillant
		"84c5a6": 421,  // Intel
		"84c727": 8053, // Gnodal
		"84c78f": 8054, // APS
		"84c7ea": 101,  // Sony
		"84c8b1": 8055, // Incognito
		"84c9b2": 803,  // D-Link
		"84cc63": 3335, // Huawei
		"84cca8": 6377, // Espressif
		"84cfbf": 8056, // Fairphone
		"84d15a": 6434, // TCT
		"84d343": 374,  // Calix
		"84d3d5": 3335, // Huawei
		"84d47e": 1685, // Aruba
		"84d4c8": 7119, // Widex
		"84d608": 4031, // Wingtech
		"84d6c5": 4907, // SolarEdge
		"84d6d0": 5439, // Amazon
		"84d81b": 1595, // TP-Link
		"84d9c8": 8057, // Unipattern
		"84dbac": 3335, // Huawei
		"84dbfc": 457,  // Nokia
		"84ddb7": 8058, // Cilag
		"84df0c": 8059, // NET2GRID
		"84e058": 125,  // ARRIS
		"84e327": 806,  // Tailyn
		"84e629": 8060, // Bluwan
		"84e892": 2247, // Actiontec
		"84e986": 3335, // Huawei
		"84ea99": 8061, // Vieworks
		"84eaed": 1916, // Roku
		"84ebef": 2,    // Cisco
		"84ed33": 8062, // BBMC
		"84ef18": 421,  // Intel
		"84f129": 8063, // Metrascale
		"84f147": 2,    // Cisco
		"84f3eb": 6377, // Espressif
		"84f6fa": 8064, // Miovision
		"84f703": 6377, // Espressif
		"84f883": 8065, // Luminar
		"84fcac": 545,  // Apple
		"84fcfe": 545,  // Apple
		"84fd27": 1655, // Silicon
		"84fdd1": 421,  // Intel
		"84fe9e": 8066, // RTC
		"880118": 8067, // BLT
		"880355": 2640, // Arcadyan
		"88074b": 869,  // LG
		"880905": 8068, // MTMCommunications
		"8809af": 6363, // Masimo
		"881036": 8069, // Panodic
		"88108f": 3335, // Huawei
		"881196": 3335, // Huawei
		"88124e": 5788, // Qualcomm
		"8815c5": 3335, // Huawei
		"8818ae": 8070, // Tamron
		"881908": 545,  // Apple
		"881c95": 6282, // Itel
		"881dfc": 2,    // Cisco
		"881fa1": 545,  // Apple
		"882012": 8071, // LMI
		"8821e3": 8072, // Nebusens
		"88238c": 4922, // Fiberhome
		"882510": 1685, // Aruba
		"88252c": 2640, // Arcadyan
		"882593": 1595, // TP-Link
		"8828b3": 3335, // Huawei
		"882949": 5696, // Renesas
		"882950": 8073, // Netmoon
		"88299c": 152,  // Samsung
		"882bd7": 8074, // Addenergie
		"882e5a": 8075, // storONE
		"88308a": 2057, // Murata
		"88329b": 152,  // Samsung
		"8833be": 8076, // Ivenix
		"88354c": 8077, // Transics
		"883612": 8078, // SRC
		"88365f": 869,  // LG
		"88366c": 1252, // EFM
		"8836cf": 3335, // Huawei
		"883a30": 1685, // Aruba
		"883c1c": 4254, // Mercury
		"883d24": 3522, // Google
		"883f99": 300,  // Siemens
		"883fd3": 3335, // Huawei
		"884033": 3335, // Huawei
		"88403b": 3335, // Huawei
		"884067": 3977, // Infomark
		"8843e1": 2,    // Cisco
		"884477": 3335, // Huawei
		"8844f6": 457,  // Nokia
		"884604": 5698, // Xiaomi
		"88462a": 7076, // Telechips
		"884a18": 8079, // Opulinks
		"884a70": 7187, // Wacom
		"884b39": 300,  // Siemens
		"884ccf": 8080, // Pulzze
		"884d7c": 545,  // Apple
		"885046": 3536, // Lear
		"8851fb": 302,  // HP
		"8852eb": 5698, // Xiaomi
		"88532e": 421,  // Intel
		"885395": 545,  // Apple
		"8853d4": 3335, // Huawei
		"88541f": 3522, // Google
		"88571d": 7149, // Seongji
		"88576d": 8081, // XTA
		"8857ee": 1077, // Buffalo
		"885a85": 1592, // Wistron
		"885a92": 2,    // Cisco
		"885bdd": 185,  // Extreme
		"885dfb": 3032, // ZTE
		"8863df": 545,  // Apple
		"886440": 545,  // Apple
		"886639": 3335, // Huawei
		"88665a": 545,  // Apple
		"8866a5": 545,  // Apple
		"88693d": 3335, // Huawei
		"886ab1": 6370, // Vivo
		"886ae3": 2237, // Alpha
		"886b0f": 1108, // Bluegiga
		"886b44": 5436, // Sunnovo
		"886b6e": 545,  // Apple
		"886f29": 8082, // Pocketbook
		"886fd4": 102,  // Dell
		"88708c": 2665, // Lenovo
		"8871b1": 125,  // ARRIS
		"8871e5": 5439, // Amazon
		"887384": 35,   // Toshiba
		"887556": 2,    // Cisco
		"887598": 152,  // Samsung
		"887873": 421,  // Intel
		"88789c": 8083, // Game
		"88795b": 3533, // Konka
		"88797e": 687,  // Motorola
		"887a31": 8084, // Velankani
		"887e25": 185,  // Extreme
		"8881b9": 3335, // Huawei
		"888322": 152,  // Samsung
		"88835d": 6441, // FN-Link
		"888603": 3335, // Huawei
		"8886a0": 8085, // Simton
		"8886c2": 8086, // STABILO
		"888717": 87,   // Canon
		"8887dd": 8087, // DarbeeVision
		"88892f": 3335, // Huawei
		"888964": 8088, // GSI
		"888e68": 3335, // Huawei
		"888e7f": 5421, // Atop
		"889009": 826,  // Juniper
		"88908d": 2,    // Cisco
		"889166": 8089, // Viewcooper
		"8891dd": 8090, // Racktivity
		"889471": 90,   // Brocade
		"88947e": 4922, // Fiberhome
		"8894f9": 8091, // Gemicom
		"88964e": 125,  // ARRIS
		"889655": 8092, // Zitte
		"889765": 8093, // exands
		"889821": 8094, // Teraon
		"889b39": 152,  // Samsung
		"889d98": 8095, // Allied-telesisK.K
		"889e33": 6434, // TCT
		"889e68": 6393, // Technicolor
		"889f6f": 152,  // Samsung
		"88a0be": 3335, // Huawei
		"88a25e": 826,  // Juniper
		"88a2d7": 3335, // Huawei
		"88a303": 152,  // Samsung
		"88a479": 545,  // Apple
		"88a4c2": 4920, // LCFC
		"88a5bd": 8096, // Qpcom
		"88a6c6": 2049, // Sagemcom
		"88a73c": 6840, // Ragentek
		"88a9b7": 545,  // Apple
		"88acc0": 2693, // ZyXEL
		"88acc1": 8097, // Generiton
		"88ad43": 5440, // Pegatron
		"88add2": 152,  // Samsung
		"88ae07": 545,  // Apple
		"88ae1d": 358,  // Compal
		"88aedd": 6673, // EliteGroup
		"88b111": 421,  // Intel
		"88b1e1": 2486, // Mojo
		"88b291": 545,  // Apple
		"88b436": 67,   // Private
		"88b4a6": 687,  // Motorola
		"88b627": 8098, // Gembird
		"88b66b": 8099, // easynetworks
		"88b6ee": 1238, // Dish
		"88b945": 545,  // Apple
		"88ba7f": 8100, // Qfiednet
		"88bcc1": 3335, // Huawei
		"88bd45": 152,  // Samsung
		"88bd78": 2669, // Flaircomm
		"88bfe4": 3335, // Huawei
		"88c08b": 545,  // Apple
		"88c227": 3335, // Huawei
		"88c242": 8101, // Poynt
		"88c3b3": 8102, // Sovico
		"88c626": 4080, // Logitech
		"88c663": 545,  // Apple
		"88c9d0": 869,  // LG
		"88cb87": 545,  // Apple
		"88cefa": 3335, // Huawei
		"88cf98": 3335, // Huawei
		"88d039": 6895, // Tonly
		"88d199": 3725, // Vencer
		"88d274": 3032, // ZTE
		"88d37b": 8103, // FirmTek
		"88d5a8": 6282, // Itel
		"88d652": 8104, // AMERGINT
		"88d7bc": 7626, // DEP
		"88d7f6": 1806, // ASUS
		"88d82e": 421,  // Intel
		"88d98f": 826,  // Juniper
		"88dc96": 8105, // EnGenius
		"88dd79": 1304, // Voltaire
		"88de7c": 2545, // Askey
		"88dea9": 1916, // Roku
		"88e034": 3907, // Shinwa
		"88e056": 3335, // Huawei
		"88e0f3": 826,  // Juniper
		"88e3ab": 3335, // Huawei
		"88e603": 8106, // Avotek
		"88e64b": 826,  // Juniper
		"88e712": 8107, // Whirlpool
		"88e87f": 545,  // Apple
		"88e90f": 8108, // innomdlelab
		"88e917": 8109, // Tamaggo
		"88e9a4": 302,  // HP
		"88e9fe": 545,  // Apple
		"88ed1c": 8110, // Cudo
		"88ef16": 125,  // ARRIS
		"88f031": 2,    // Cisco
		"88f077": 2,    // Cisco
		"88f488": 8111, // cellon
		"88f490": 8112, // Jetmobile
		"88f56e": 3335, // Huawei
		"88f7bf": 6370, // Vivo
		"88f7c7": 6393, // Technicolor
		"88f872": 3335, // Huawei
		"88fca6": 1637, // devolo
		"88fd15": 8113, // Lineeye
		"8c006d": 545,  // Apple
		"8c02fa": 8114, // CoMMANDO
		"8c04ba": 102,  // Dell
		"8c04ff": 6393, // Technicolor
		"8c0551": 8115, // Koubachi
		"8c09f4": 125,  // ARRIS
		"8c0c87": 457,  // Nokia
		"8c0c90": 2738, // Ruckus
		"8c0ca3": 8116, // Amper
		"8c0d76": 3335, // Huawei
		"8c0f6f": 5440, // Pegatron
		"8c0fa0": 8117, // di-soric
		"8c0fc9": 3335, // Huawei
		"8c0ffa": 8118, // Hutec
		"8c10d4": 2049, // Sagemcom
		"8c14b4": 3032, // ZTE
		"8c15c7": 3335, // Huawei
		"8c1645": 4920, // LCFC
		"8c1759": 421,  // Intel
		"8c17b6": 3335, // Huawei
		"8c19b5": 2640, // Arcadyan
		"8c1abf": 152,  // Samsung
		"8c1d96": 421,  // Intel
		"8c210a": 1595, // TP-Link
		"8c2505": 3335, // Huawei
		"8c271d": 8119, // QuantHouse
		"8c278a": 2816, // Vocollect
		"8c2937": 545,  // Apple
		"8c2daa": 545,  // Apple
		"8c31e2": 8120, // Dayouplus
		"8c3330": 8121, // EmFirst
		"8c3446": 3335, // Huawei
		"8c34fd": 3335, // Huawei
		"8c3a7e": 3858, // Universal
		"8c3ae3": 869,  // LG
		"8c3bad": 1368, // Netgear
		"8c3c07": 8122, // Skiva
		"8c3c4a": 6083, // NAKAYO
		"8c41f2": 8123, // RDA
		"8c41f4": 8124, // IPmotion
		"8c426d": 3335, // Huawei
		"8c444f": 531,  // HUMAX
		"8c4500": 2057, // Murata
		"8c477f": 8125, // NambooSolution
		"8c47be": 102,  // Dell
		"8c4962": 1916, // Roku
		"8c497a": 185,  // Extreme
		"8c49b6": 6370, // Vivo
		"8c4b14": 6377, // Espressif
		"8c4cad": 8126, // Evoluzn
		"8c4cdc": 4479, // Planex
		"8c4db9": 8127, // Unmonday
		"8c4dea": 8128, // Cerio
		"8c53f7": 1331, // A&D
		"8c541d": 7469, // LGE
		"8c554a": 421,  // Intel
		"8c5646": 869,  // LG
		"8c56c5": 1427, // Nintendo
		"8c579b": 1592, // Wistron
		"8c5877": 545,  // Apple
		"8c5973": 2693, // ZyXEL
		"8c598b": 8129, // C
		"8c59c3": 6492, // ADB
		"8c5a25": 125,  // ARRIS
		"8c5ac1": 3335, // Huawei
		"8c5bf0": 125,  // ARRIS
		"8c5ca1": 8130, // d-broad
		"8c5d60": 8131, // UCI
		"8c5ebd": 3335, // Huawei
		"8c5fad": 4922, // Fiberhome
		"8c604f": 2,    // Cisco
		"8c6078": 8132, // Swissbit
		"8c60e7": 8133, // Mpgio
		"8c61a3": 125,  // ARRIS
		"8c6422": 101,  // Sony
		"8c64a2": 6924, // OnePlus
		"8c6794": 6370, // Vivo
		"8c683a": 3335, // Huawei
		"8c6878": 8134, // Nortek-AS
		"8c68c8": 3032, // ZTE
		"8c6a8d": 6393, // Technicolor
		"8c6ae4": 8135, // Viogem
		"8c6d77": 3335, // Huawei
		"8c705a": 421,  // Intel
		"8c71f8": 152,  // Samsung
		"8c736e": 4,    // Fujitsu
		"8c73a0": 4922, // Fiberhome
		"8c76c1": 8136, // Goden
		"8c7712": 152,  // Samsung
		"8c7967": 3032, // ZTE
		"8c79f5": 152,  // Samsung
		"8c7a15": 2738, // Ruckus
		"8c7a3d": 5698, // Xiaomi
		"8c7aaa": 545,  // Apple
		"8c7b9d": 545,  // Apple
		"8c7c92": 545,  // Apple
		"8c7cff": 90,   // Brocade
		"8c7eb3": 8137, // Lytro
		"8c7f3b": 125,  // ARRIS
		"8c8126": 8138, // Arcom
		"8c82a8": 7261, // Insigma
		"8c83df": 457,  // Nokia
		"8c83e1": 152,  // Samsung
		"8c83e8": 3335, // Huawei
		"8c8401": 67,   // Private
		"8c8590": 545,  // Apple
		"8c85c1": 1685, // Aruba
		"8c85e6": 8139, // Cleondris
		"8c861e": 545,  // Apple
		"8c897a": 8140, // Augtek
		"8c89a5": 1812, // Micro-Star
		"8c8caa": 4920, // LCFC
		"8c8d28": 421,  // Intel
		"8c8e76": 8141, // taskit
		"8c8ef2": 545,  // Apple
		"8c8fe9": 545,  // Apple
		"8c90d3": 457,  // Nokia
		"8c9236": 8142, // Aus.Linx
		"8c9351": 8143, // Jigowatts
		"8c941f": 2,    // Cisco
		"8c94cc": 3188, // SFR
		"8c94cf": 8144, // Encell
		"8c99e6": 6434, // TCT
		"8ca2fd": 8145, // Starry
		"8ca399": 8146, // Servercom
		"8ca6df": 1595, // TP-Link
		"8ca96f": 873,  // D&M
		"8ca982": 421,  // Intel
		"8caab5": 6377, // Espressif
		"8caace": 5698, // Xiaomi
		"8cae4c": 8147, // Plugable
		"8cae89": 8148, // Y-cam
		"8caedb": 8149, // Nagtech
		"8cb0e9": 152,  // Samsung
		"8cb64f": 2,    // Cisco
		"8cb82c": 8150, // IPitomy
		"8cb84a": 152,  // Samsung
		"8cb864": 4904, // AcSiP
		"8cb87e": 421,  // Intel
		"8cba25": 3506, // Unionman
		"8cbebe": 5698, // Xiaomi
		"8cbfa6": 152,  // Samsung
		"8cc121": 2154, // Panasonic
		"8cc5b4": 2049, // Sagemcom
		"8cc661": 8151, // Current
		"8cc681": 421,  // Intel
		"8cc7aa": 8152, // Radinet
		"8cc8cd": 152,  // Samsung
		"8ccda2": 8153, // ACTP
		"8ccde8": 1427, // Nintendo
		"8cce4e": 6377, // Espressif
		"8ccf09": 102,  // Dell
		"8ccf5c": 8154, // BEFEGA
		"8ccf8f": 7019, // ITC
		"8cd17b": 8155, // CG
		"8cd3a2": 8156, // VisSim
		"8cd48e": 6282, // Itel
		"8cd9d6": 5698, // Xiaomi
		"8cdb25": 8157, // ESG
		"8cdc02": 3032, // ZTE
		"8cdcd4": 302,  // HP
		"8cde52": 8158, // ISSC
		"8cde99": 8159, // Comlab
		"8cdee6": 152,  // Samsung
		"8cdf9d": 48,   // NEC
		"8ce081": 3032, // ZTE
		"8ce117": 3032, // ZTE
		"8ce38e": 8160, // Kioxia
		"8ce5c0": 152,  // Samsung
		"8ce5ef": 3335, // Huawei
		"8ce78c": 8161, // DK
		"8ce7b3": 8162, // Sonardyne
		"8cea1b": 6295, // Edgecore
		"8cea48": 152,  // Samsung
		"8cebc6": 3335, // Huawei
		"8cec4b": 102,  // Dell
		"8cec7b": 545,  // Apple
		"8ceec6": 8163, // Precepscion
		"8cf112": 687,  // Motorola
		"8cf228": 4254, // Mercury
		"8cf319": 300,  // Siemens
		"8cf5a3": 152,  // Samsung
		"8cf681": 1655, // Silicon
		"8cf710": 4499, // AMPAK
		"8cf773": 457,  // Nokia
		"8cf8c5": 421,  // Intel
		"8cf9c9": 8164, // MESADA
		"8cfaba": 545,  // Apple
		"8cfd18": 3335, // Huawei
		"8cfdde": 2049, // Sagemcom
		"8cfdf0": 5788, // Qualcomm
		"8cfe57": 545,  // Apple
		"8cfe74": 2738, // Ruckus
		"8cfeb4": 8165, // Vsoontech
		"9000db": 152,  // Samsung
		"90013b": 2049, // Sagemcom
		"900218": 3513, // BSkyB
		"900325": 3335, // Huawei
		"9003b7": 2563, // Parrot
		"900628": 152,  // Samsung
		"900917": 8166, // Far-sighted
		"9009d0": 2451, // Synology
		"900a39": 8167, // Wiio
		"900a84": 432,  // Mellanox
		"900bc1": 6820, // Sprocomm
		"900cb4": 8168, // Alinket
		"900cc8": 3522, // Google
		"900d66": 8169, // Digimore
		"900dcb": 125,  // ARRIS
		"901195": 5439, // Amazon
		"9012a1": 6291, // We
		"9016ba": 3335, // Huawei
		"90173f": 3335, // Huawei
		"90179b": 8170, // Nanomegas
		"9017ac": 3335, // Huawei
		"9017c8": 3335, // Huawei
		"90187c": 152,  // Samsung
		"901900": 8171, // SCS
		"901aca": 125,  // ARRIS
		"901b0e": 4,    // Fujitsu
		"901d27": 3032, // ZTE
		"901edd": 8172, // Great
		"9020c2": 1685, // Aruba
		"902106": 3513, // BSkyB
		"902155": 1341, // HTC
		"9023ec": 6905, // Availink
		"9027e4": 545,  // Apple
		"902b34": 1929, // Giga-Byte
		"902bd2": 3335, // Huawei
		"902e16": 4920, // LCFC
		"902e1c": 421,  // Intel
		"902e87": 8173, // LabJack
		"90342b": 8174, // Gatekeeper
		"9035ea": 1655, // Silicon
		"903809": 306,  // Ericsson
		"90380c": 6377, // Espressif
		"903a72": 2738, // Ruckus
		"903aa0": 457,  // Nokia
		"903ae6": 2563, // Parrot
		"903c92": 545,  // Apple
		"903cb3": 6295, // Edgecore
		"903d68": 8175, // G-Printec
		"903d6b": 8176, // Zicon
		"903e7f": 4922, // Fiberhome
		"903eab": 125,  // ARRIS
		"903fea": 3335, // Huawei
		"9043e2": 8177, // Cornami
		"9046b7": 8178, // Vadaro
		"904716": 8179, // Rorze
		"9049fa": 421,  // Intel
		"904c81": 302,  // HP
		"904d4a": 2049, // Sagemcom
		"904dc3": 8180, // Flonidan
		"904e2b": 3335, // Huawei
		"90505a": 8181, // unGlue
		"9050ca": 870,  // Hitron
		"905446": 8182, // TES
		"9055ae": 306,  // Ericsson
		"9055de": 4922, // Fiberhome
		"905682": 8183, // Lenbrook
		"905692": 8184, // Autotalks
		"9056fc": 6266, // Tecno
		"905851": 6393, // Technicolor
		"905c34": 8185, // Sirius
		"905c44": 358,  // Compal
		"905f2e": 6434, // TCT
		"905f8d": 8186, // modas
		"9060f1": 545,  // Apple
		"90610c": 8187, // Fida
		"9061ae": 421,  // Intel
		"90633b": 152,  // Samsung
		"90671c": 3335, // Huawei
		"9068c3": 687,  // Motorola
		"906976": 8188, // Withrobot
		"906cac": 1323, // Fortinet
		"906d05": 8189, // BXB
		"906f18": 67,   // Private
		"907240": 545,  // Apple
		"907282": 2049, // Sagemcom
		"90735a": 687,  // Motorola
		"90749d": 7596, // IRay
		"9077ee": 2,    // Cisco
		"907841": 421,  // Intel
		"9078b2": 5698, // Xiaomi
		"907990": 2080, // Benchmark
		"907a58": 7843, // Zegna-Daidong
		"907af1": 8190, // Wally
		"907e30": 8191, // Lars
		"907eba": 8192, // Utek
		"907f61": 7301, // Chicony
		"908060": 8193, // Nilfisk
		"90808f": 3335, // Huawei
		"90812a": 545,  // Apple
		"908158": 545,  // Apple
		"908175": 152,  // Samsung
		"90840d": 545,  // Apple
		"90848b": 8194, // HDR10+
		"90869b": 3032, // ZTE
		"908c43": 545,  // Apple
		"908d1d": 8195, // GH
		"908d6c": 545,  // Apple
		"908d6e": 102,  // Dell
		"908d78": 803,  // D-Link
		"90903c": 8196, // Trison
		"909497": 3335, // Huawei
		"9094e4": 803,  // D-Link
		"9096f3": 1077, // Buffalo
		"9097d5": 6377, // Espressif
		"9097f3": 152,  // Samsung
		"909838": 3335, // Huawei
		"909864": 8197, // Impex-Sat&Co
		"909a4a": 1595, // TP-Link
		"909c4a": 545,  // Apple
		"909d7d": 125,  // ARRIS
		"909f33": 1252, // EFM
		"90a25b": 545,  // Apple
		"90a2da": 8198, // Gheo
		"90a46a": 8199, // Sisnet
		"90a4de": 1592, // Wistron
		"90a5af": 3335, // Huawei
		"90a62f": 8200, // Naver
		"90a822": 5439, // Amazon
		"90a935": 8201, // JWEntertainment
		"90aac3": 870,  // Hitron
		"90ac3f": 8202, // BrightSign
		"90adf7": 6370, // Vivo
		"90adfc": 7076, // Telechips
		"90ae1b": 1595, // TP-Link
		"90afd1": 8203, // netKTI
		"90b0ed": 545,  // Apple
		"90b11c": 102,  // Dell
		"90b134": 125,  // ARRIS
		"90b144": 152,  // Samsung
		"90b21f": 545,  // Apple
		"90b4dd": 67,   // Private
		"90b622": 152,  // Samsung
		"90b686": 2057, // Murata
		"90b832": 185,  // Extreme
		"90b8d0": 8204, // Joyent
		"90b931": 545,  // Apple
		"90c115": 101,  // Sony
		"90c119": 457,  // Nokia
		"90c1c6": 545,  // Apple
		"90c54a": 6370, // Vivo
		"90c792": 125,  // ARRIS
		"90c7d8": 3032, // ZTE
		"90cc24": 7971, // Synaptics
		"90ccdf": 421,  // Intel
		"90cf15": 457,  // Nokia
		"90cf6f": 8205, // Dlogixs
		"90d74f": 8206, // Bookeen
		"90d852": 6101, // Comtec
		"90d8f3": 3032, // ZTE
		"90d92c": 8207, // HUG-Witschi
		"90da4e": 8208, // Avanu
		"90db46": 8209, // E-Lead
		"90dd5d": 545,  // Apple
		"90dffb": 8210, // Homerider
		"90e17b": 545,  // Apple
		"90e2ba": 421,  // Intel
		"90e6ba": 1806, // ASUS
		"90e7c4": 1341, // HTC
		"90e868": 3005, // Azurewave
		"90ec50": 8211, // C.O.B.O
		"90ec77": 8212, // silicom
		"90eec7": 152,  // Samsung
		"90ef68": 2693, // ZyXEL
		"90f052": 7031, // MEIZU
		"90f157": 797,  // Garmin
		"90f1aa": 152,  // Samsung
		"90f305": 531,  // HUMAX
		"90f3b7": 8213, // Kirisun
		"90f644": 3335, // Huawei
		"90f652": 1595, // TP-Link
		"90f891": 1269, // Kaonmedia
		"90f9b7": 3335, // Huawei
		"90fb5b": 620,  // Avaya
		"90fd61": 545,  // Apple
		"90fd73": 3032, // ZTE
		"90fd9f": 1655, // Silicon
		"940006": 8214, // jinyoung
		"940070": 457,  // Nokia
		"9400b0": 3335, // Huawei
		"940149": 8215, // AutoHotBox
		"9401c2": 152,  // Samsung
		"940230": 4080, // Logitech
		"94026b": 8216, // Optictimes
		"94049c": 3335, // Huawei
		"940853": 2874, // Liteon
		"9408c7": 3335, // Huawei
		"940937": 531,  // HUMAX
		"940b19": 3335, // Huawei
		"940b2d": 8217, // NetView
		"940bd5": 8218, // Himax
		"940c6d": 1595, // TP-Link
		"940c98": 545,  // Apple
		"940d2d": 3858, // Universal
		"940e6b": 3335, // Huawei
		"94103e": 2469, // Belkin
		"94147a": 6370, // Vivo
		"941625": 545,  // Apple
		"941700": 5698, // Xiaomi
		"941865": 1368, // Netgear
		"941882": 302,  // HP
		"94193a": 8219, // Elvaco
		"941c56": 2247, // Actiontec
		"941f3a": 8220, // Ambiq
		"942053": 457,  // Nokia
		"942197": 8221, // Stalmart
		"942533": 3335, // Huawei
		"942790": 6434, // TCT
		"942957": 8222, // Airpo
		"942a3f": 8223, // Diversey
		"942a6f": 2979, // Ubiquiti
		"942cb3": 531,  // HUMAX
		"942ddc": 152,  // Samsung
		"942e17": 56,   // Schneider
		"942e63": 8224, // Finsecur
		"94319b": 8225, // Alphatronics
		"9431cb": 6370, // Vivo
		"9433dd": 8226, // Taco
		"943469": 1655, // Silicon
		"94350a": 152,  // Samsung
		"9437f7": 3335, // Huawei
		"943a91": 5439, // Amazon
		"943af0": 457,  // Nokia
		"943bb1": 1269, // Kaonmedia
		"943cc6": 6377, // Espressif
		"943fc2": 302,  // HP
		"9440a2": 8227, // Anywave
		"9440c9": 302,  // HP
		"9441c1": 8228, // Mini-Cam
		"94434d": 4565, // Ciena
		"944444": 869,  // LG
		"944452": 2469, // Belkin
		"944696": 2684, // BaudTec
		"944996": 8229, // WiSilica
		"944a0c": 2075, // Sercomm
		"945047": 8230, // Rechnerbetriebsgruppe
		"945089": 8231, // SimonsVoss
		"945103": 152,  // Samsung
		"945493": 8232, // Rigado
		"9454df": 8233, // YST
		"9457a5": 302,  // HP
		"9458cb": 1427, // Nintendo
		"94592d": 8234, // EKE
		"945afc": 5439, // Amazon
		"945b7e": 8235, // Trilobit
		"945c9a": 545,  // Apple
		"945f34": 5696, // Renesas
		"946010": 3335, // Huawei
		"9460d5": 1685, // Aruba
		"94611e": 8236, // Wata
		"946124": 8237, // Pason
		"946269": 125,  // ARRIS
		"946372": 6370, // Vivo
		"9463d1": 152,  // Samsung
		"946424": 1685, // Aruba
		"94652d": 6924, // OnePlus
		"94659c": 421,  // Intel
		"9466e7": 8238, // WOM
		"946a77": 6393, // Technicolor
		"946ab0": 2640, // Arcadyan
		"9470d2": 8239, // Winfirm
		"9471ac": 6434, // TCT
		"94756e": 2542, // QinetiQ
		"9476b7": 152,  // Samsung
		"94772b": 3335, // Huawei
		"947bbe": 8240, // Ubicquia
		"947be7": 152,  // Samsung
		"9481a4": 8241, // Azuray
		"9483c4": 4372, // GL
		"94857a": 8242, // Evantage
		"9486cd": 8243, // Seoul
		"94877c": 125,  // ARRIS
		"9487e0": 5698, // Xiaomi
		"948bc1": 152,  // Samsung
		"948d50": 8244, // Beamex
		"948ed3": 3794, // Arista
		"948fcf": 125,  // ARRIS
		"949010": 3335, // Huawei
		"94917f": 2545, // Askey
		"9492bc": 7512, // Syntech
		"9492d2": 8245, // KCF
		"949426": 545,  // Apple
		"94944a": 8246, // Particle
		"9495a0": 3522, // Google
		"949869": 3032, // ZTE
		"949aa9": 612,  // Microsoft
		"949b2c": 185,  // Extreme
		"949d57": 2154, // Panasonic
		"949f3e": 2048, // Sonos
		"94a04e": 8247, // Bostex
		"94a1a2": 4499, // AMPAK
		"94a3ca": 8248, // KonnectONE
		"94a4f9": 3335, // Huawei
		"94a67e": 1368, // Netgear
		"94a7b7": 3032, // ZTE
		"94a7bc": 8249, // BodyMedia
		"94aa0a": 4922, // Fiberhome
		"94aab8": 8250, // Joview
		"94abfe": 457,  // Nokia
		"94acca": 8251, // trivum
		"94aef0": 2,    // Cisco
		"94b01f": 545,  // Apple
		"94b10a": 152,  // Samsung
		"94b271": 3335, // Huawei
		"94b2cc": 6118, // Pioneer
		"94b34f": 2738, // Ruckus
		"94b40f": 1685, // Aruba
		"94b555": 6377, // Espressif
		"94b819": 457,  // Nokia
		"94b86d": 421,  // Intel
		"94b8c5": 1586, // RuggedCom
		"94b97e": 6377, // Espressif
		"94b9b4": 8252, // Aptos
		"94bbae": 8253, // Husqvarna
		"94be46": 687,  // Motorola
		"94bf2d": 545,  // Apple
		"94bf80": 3032, // ZTE
		"94bf94": 826,  // Juniper
		"94bfc4": 2738, // Ruckus
		"94c038": 8254, // Tallac
		"94c150": 1939, // 2Wire
		"94c2bd": 8255, // Tecnobit
		"94c691": 6673, // EliteGroup
		"94c6eb": 8256, // NOVA
		"94c7af": 8257, // Raylios
		"94c962": 8258, // Teseq
		"94ccb9": 125,  // ARRIS
		"94cdac": 8259, // Creowave
		"94ce2c": 101,  // Sony
		"94ce31": 2059, // CTS
		"94d00d": 3335, // Huawei
		"94d019": 8260, // Cydle
		"94d299": 8261, // Techmation
		"94d2bc": 3335, // Huawei
		"94d469": 2,    // Cisco
		"94d505": 4922, // Fiberhome
		"94d6db": 8262, // NexFi
		"94d771": 152,  // Samsung
		"94d859": 6434, // TCT
		"94d93c": 8263, // Enelps
		"94d9b3": 1595, // TP-Link
		"94db49": 8264, // Sitcorp
		"94db56": 101,  // Sony
		"94dbc9": 3005, // Azurewave
		"94dbda": 3335, // Huawei
		"94dc4e": 8265, // AEV
		"94de0e": 8266, // SmartOptics
		"94de80": 1929, // Giga-Byte
		"94deb8": 1655, // Silicon
		"94df4e": 1592, // Wistron
		"94df58": 8267, // IJ
		"94e129": 152,  // Samsung
		"94e23c": 421,  // Intel
		"94e3ee": 3032, // ZTE
		"94e4ba": 3335, // Huawei
		"94e686": 6377, // Espressif
		"94e6f7": 421,  // Intel
		"94e70b": 421,  // Intel
		"94e7ea": 3335, // Huawei
		"94e8c5": 125,  // ARRIS
		"94e96a": 545,  // Apple
		"94e979": 2874, // Liteon
		"94e98c": 457,  // Nokia
		"94e9ee": 3335, // Huawei
		"94ea32": 545,  // Apple
		"94eb2c": 3522, // Google
		"94ebcd": 2221, // BlackBerry
		"94f128": 302,  // HP
		"94f278": 8268, // Elma
		"94f665": 2738, // Ruckus
		"94f692": 8269, // Geminico
		"94f6a3": 545,  // Apple
		"94f6d6": 545,  // Apple
		"94f7ad": 826,  // Juniper
		"94fb29": 768,  // Zebra
		"94fd1d": 8270, // WhereWhen
		"94fe22": 3335, // Huawei
		"94fef4": 2049, // Sagemcom
		"94ff3c": 1323, // Fortinet
		"98006a": 3032, // ZTE
		"980074": 2051, // Raisecom
		"9800c6": 545,  // Apple
		"9801a7": 545,  // Apple
		"980284": 8271, // Theobroma
		"98039b": 432,  // Mellanox
		"9803d8": 545,  // Apple
		"98063c": 152,  // Samsung
		"9809cf": 6924, // OnePlus
		"980c82": 152,  // Samsung
		"980ca5": 687,  // Motorola
		"980d2e": 1341, // HTC
		"980d51": 3335, // Huawei
		"980d67": 2693, // ZyXEL
		"980e24": 8272, // Phytium
		"980ee4": 67,   // Private
		"981082": 8273, // Nsolution
		"9810e8": 545,  // Apple
		"981333": 3032, // ZTE
		"9814d2": 8274, // Avonic
		"98192c": 6295, // Edgecore
		"981a35": 3335, // Huawei
		"981dfa": 152,  // Samsung
		"981e19": 2049, // Sagemcom
		"98208e": 8275, // Definium
		"9822ef": 2874, // Liteon
		"98234e": 8276, // Micromedia
		"9828a6": 358,  // Compal
		"9829a6": 358,  // Compal
		"982cbc": 421,  // Intel
		"982cbe": 1939, // 2Wire
		"982d68": 152,  // Samsung
		"982dba": 8277, // Fibergate
		"982ff8": 3335, // Huawei
		"983571": 8278, // Sub10
		"9835b8": 8279, // Assembled
		"9835ed": 3335, // Huawei
		"98387d": 8280, // Itronic
		"98398e": 152,  // Samsung
		"983b16": 4499, // AMPAK
		"983b67": 7447, // DWnet
		"983b8f": 421,  // Intel
		"983f60": 3335, // Huawei
		"9840bb": 102,  // Dell
		"98415c": 1427, // Nintendo
		"984246": 8281, // SOL
		"984265": 2049, // Sagemcom
		"9843da": 8282, // Intertech
		"9843fa": 421,  // Intel
		"9844ce": 3335, // Huawei
		"98460a": 545,  // Apple
		"984827": 1595, // TP-Link
		"984874": 3335, // Huawei
		"984914": 1592, // Wistron
		"9849e1": 1041, // Boeing
		"984b4a": 125,  // ARRIS
		"984be1": 302,  // HP
		"984fee": 421,  // Intel
		"98502e": 545,  // Apple
		"98523d": 3940, // Sunitec
		"98524a": 6393, // Technicolor
		"9852b1": 152,  // Samsung
		"98541b": 421,  // Intel
		"98588a": 8283, // SysGRATION
		"985aeb": 545,  // Apple
		"985bb0": 8284, // Kmdata
		"985d46": 8285, // PeopleNet
		"985d82": 3794, // Arista
		"985e1b": 8286, // ConversDigital
		"985f4f": 8287, // Tongfang
		"985fd3": 612,  // Microsoft
		"986022": 8288, // EMW
		"9860ca": 545,  // Apple
		"98672e": 7040, // Skullcandy
		"98698a": 545,  // Apple
		"986b3d": 125,  // ARRIS
		"986cf5": 3032, // ZTE
		"986dc8": 35,   // Toshiba
		"9873c4": 619,  // Sage
		"9874da": 6296, // Infinix
		"98751a": 3335, // Huawei
		"9876b6": 8289, // Adafruit
		"9877e7": 1269, // Kaonmedia
		"987a10": 306,  // Ericsson
		"987a14": 612,  // Microsoft
		"987e46": 8290, // Emizon
		"987ee3": 6370, // Vivo
		"9880ee": 152,  // Samsung
		"988217": 8291, // Disruptive
		"988389": 152,  // Samsung
		"9886b1": 8292, // Flyaudio
		"988b5d": 2049, // Sagemcom
		"988bad": 8293, // Corintech
		"988d46": 421,  // Intel
		"988e4a": 8294, // Noxus
		"988e79": 8295, // Qudelix
		"988ed4": 6282, // Itel
		"989096": 102,  // Dell
		"9893cc": 869,  // LG
		"9897cc": 1595, // TP-Link
		"9897d1": 6429, // MitraStar
		"989ab9": 3032, // ZTE
		"989c57": 3335, // Huawei
		"989d5d": 6393, // Technicolor
		"989e63": 545,  // Apple
		"98a404": 306,  // Ericsson
		"98a40e": 8296, // Snap
		"98a5f9": 545,  // Apple
		"98ad1d": 3335, // Huawei
		"98ae71": 8297, // VVDN
		"98af65": 421,  // Intel
		"98b039": 457,  // Nokia
		"98b08b": 152,  // Samsung
		"98b3ef": 3335, // Huawei
		"98b6e9": 1427, // Nintendo
		"98b8ba": 869,  // LG
		"98b8bc": 152,  // Samsung
		"98b8e3": 545,  // Apple
		"98ba39": 3874, // Doro
		"98bb99": 6849, // Phicomm
		"98bc57": 1486, // SVA
		"98bc99": 8298, // Edeltech
		"98be94": 372,  // IBM
		"98c5db": 306,  // Ericsson
		"98c845": 8299, // PacketAccess
		"98c8b8": 6370, // Vivo
		"98ca33": 545,  // Apple
		"98cb27": 8300, // Galore
		"98cba4": 2080, // Benchmark
		"98cdac": 6377, // Espressif
		"98cdb4": 8301, // Virident
		"98d293": 3522, // Google
		"98d6bb": 545,  // Apple
		"98d6f7": 869,  // LG
		"98d88c": 76,   // Nortel
		"98da92": 8302, // Vuzix
		"98dac4": 1595, // TP-Link
		"98dcd9": 8303, // UNITEC
		"98dd60": 545,  // Apple
		"98ddea": 6296, // Infinix
		"98ded0": 1595, // TP-Link
		"98e0d9": 545,  // Apple
		"98e165": 8304, // Accutome
		"98e476": 8305, // Zentan
		"98e743": 102,  // Dell
		"98e79a": 221,  // Foxconn
		"98e7f4": 302,  // HP
		"98e7f5": 3335, // Huawei
		"98e848": 8306, // Axiim
		"98e8fa": 1427, // Nintendo
		"98ed5c": 7310, // Tesla
		"98ed7e": 5821, // eero
		"98eecb": 1592, // Wistron
		"98ef9b": 6578, // Ohsung
		"98f058": 8307, // Lynxspringl
		"98f083": 3335, // Huawei
		"98f0ab": 545,  // Apple
		"98f170": 2057, // Murata
		"98f217": 3798, // Castlenet
		"98f2b3": 302,  // HP
		"98f428": 3032, // ZTE
		"98f4ab": 6377, // Espressif
		"98f537": 3032, // ZTE
		"98f5a9": 6578, // Ohsung
		"98f621": 5698, // Xiaomi
		"98f781": 125,  // ARRIS
		"98f7d7": 125,  // ARRIS
		"98f8c1": 3226, // IDT
		"98fa9b": 4920, // LCFC
		"98faa7": 8308, // Innonet
		"98fae3": 5698, // Xiaomi
		"98fb12": 1883, // Grand
		"98fc11": 1783, // Linksys
		"98fd74": 8309, // ACT
		"98fdb4": 389,  // Primax
		"98fe03": 306,  // Ericsson
		"98fe94": 545,  // Apple
		"98ff6a": 8310, // OTECTechnology
		"98ffd0": 2665, // Lenovo
		"9c0298": 152,  // Samsung
		"9c0473": 8311, // Tecmobile
		"9c04eb": 545,  // Apple
		"9c066e": 7601, // Hytera
		"9c0b05": 5821, // eero
		"9c1874": 457,  // Nokia
		"9c1c12": 1685, // Aruba
		"9c1c37": 7032, // AltoBeam
		"9c1d36": 3335, // Huawei
		"9c1e95": 2247, // Actiontec
		"9c1ea4": 5696, // Renesas
		"9c1fdd": 8312, // Accupix
		"9c207b": 545,  // Apple
		"9c216a": 1595, // TP-Link
		"9c220e": 8313, // TASCAN
		"9c2595": 152,  // Samsung
		"9c2840": 8314, // Discovery
		"9c28b3": 545,  // Apple
		"9c28ef": 3335, // Huawei
		"9c28f7": 5698, // Xiaomi
		"9c293f": 545,  // Apple
		"9c2976": 421,  // Intel
		"9c2a83": 152,  // Samsung
		"9c2ba6": 5441, // Ruijie
		"9c2ea1": 5698, // Xiaomi
		"9c2f4e": 3032, // ZTE
		"9c2f9d": 2874, // Liteon
		"9c31c3": 3513, // BSkyB
		"9c32ce": 87,   // Canon
		"9c3426": 125,  // ARRIS
		"9c35eb": 545,  // Apple
		"9c37f4": 3335, // Huawei
		"9c3aaf": 152,  // Samsung
		"9c3dcf": 1368, // Netgear
		"9c3e53": 545,  // Apple
		"9c3eaa": 8315, // EnvyLogic
		"9c40cd": 5355, // Synclayer
		"9c417c": 8316, // Hame
		"9c44a6": 8317, // SwiftTest
		"9c4a7b": 457,  // Nokia
		"9c4cae": 5320, // Mesa
		"9c4e20": 2,    // Cisco
		"9c4e36": 421,  // Intel
		"9c4e8e": 8318, // ALT
		"9c4ebf": 8319, // BoxCast
		"9c4f5f": 3522, // Google
		"9c4fcf": 6434, // TCT
		"9c4fda": 545,  // Apple
		"9c50d1": 2057, // Murata
		"9c50ee": 1640, // Cambridge
		"9c52f8": 3335, // Huawei
		"9c5440": 67,   // Private
		"9c54da": 7685, // SkyBell
		"9c5636": 3335, // Huawei
		"9c57ad": 2,    // Cisco
		"9c583c": 545,  // Apple
		"9c5a44": 358,  // Compal
		"9c5a81": 5698, // Xiaomi
		"9c5b96": 8320, // NMR
		"9c5c8e": 1806, // ASUS
		"9c5cf9": 101,  // Sony
		"9c5d12": 185,  // Extreme
		"9c5d95": 8321, // VTC
		"9c5fb0": 152,  // Samsung
		"9c611d": 2154, // Panasonic
		"9c62ab": 4655, // Sumavision
		"9c63ed": 3032, // ZTE
		"9c648b": 545,  // Apple
		"9c65b0": 152,  // Samsung
		"9c65f9": 4904, // AcSiP
		"9c685b": 8322, // Octonion
		"9c6865": 4922, // Fiberhome
		"9c6937": 6653, // Qorvo
		"9c69d1": 3335, // Huawei
		"9c6b37": 5696, // Renesas
		"9c6c15": 612,  // Microsoft
		"9c6f52": 3032, // ZTE
		"9c713a": 3335, // Huawei
		"9c7370": 3335, // Huawei
		"9c741a": 3335, // Huawei
		"9c746f": 3335, // Huawei
		"9c760e": 545,  // Apple
		"9c7613": 6947, // Ring
		"9c77aa": 8323, // Nadasnv
		"9c79ac": 8324, // Suntec
		"9c7a03": 4565, // Ciena
		"9c7bef": 302,  // HP
		"9c7da3": 3335, // Huawei
		"9c80df": 2640, // Arcadyan
		"9c823f": 3335, // Huawei
		"9c8281": 6370, // Vivo
		"9c83bf": 8325, // PRO-VISION
		"9c84bf": 545,  // Apple
		"9c8566": 4031, // Wingtech
		"9c88ad": 4922, // Fiberhome
		"9c8acb": 826,  // Juniper
		"9c8ba0": 545,  // Apple
		"9c8bf1": 8326, // Warehouse
		"9c8c6e": 152,  // Samsung
		"9c8cd8": 302,  // HP
		"9c8d7c": 430,  // Alpsalpine
		"9c8dd3": 8327, // Leonton
		"9c8e99": 302,  // HP
		"9c8e9c": 3335, // Huawei
		"9c8ecd": 8328, // Amcrest
		"9c8edc": 4821, // Teracom
		"9c9019": 8329, // Beyless
		"9c934e": 0,    // Xerox
		"9c93b0": 8330, // Megatronix
		"9c93e4": 67,   // Private
		"9c9567": 3335, // Huawei
		"9c95f8": 8331, // SmartDoor
		"9c9726": 6393, // Technicolor
		"9c9789": 6628, // 1MORE
		"9c99a0": 5698, // Xiaomi
		"9c99cd": 8332, // Voippartners
		"9c9c1d": 8333, // Starkey
		"9c9c1f": 6377, // Espressif
		"9c9d5d": 8334, // Raden
		"9c9e71": 3335, // Huawei
		"9ca134": 8335, // Nike
		"9ca2f4": 1595, // TP-Link
		"9ca513": 152,  // Samsung
		"9ca570": 5821, // eero
		"9ca577": 8336, // Osorno
		"9ca5c0": 6370, // Vivo
		"9ca615": 1595, // TP-Link
		"9ca69d": 8337, // Whaley
		"9ca9e4": 3032, // ZTE
		"9caa1b": 612,  // Microsoft
		"9cac6d": 3858, // Universal
		"9cadef": 8338, // Obihai
		"9caed3": 46,   // Epson
		"9caf6f": 6282, // Itel
		"9cafca": 2,    // Cisco
		"9cb206": 8339, // Procentec
		"9cb2b2": 3335, // Huawei
		"9cb2e8": 3335, // Huawei
		"9cb654": 302,  // HP
		"9cb6d0": 8340, // Rivet
		"9cb70d": 2874, // Liteon
		"9cb793": 8341, // Creatcomm
		"9cb8b4": 4499, // AMPAK
		"9cbcf0": 5698, // Xiaomi
		"9cbd6e": 8342, // DERA
		"9cbd9d": 8343, // SkyDisk
		"9cbee0": 8344, // Biosoundlab
		"9cc077": 8345, // PrintCounts
		"9cc0d2": 8346, // Conductix-Wampfler
		"9cc172": 3335, // Huawei
		"9cc2c4": 7723, // Inspur
		"9cc7a6": 621,  // AVM
		"9cc7d1": 3206, // Sharp
		"9cc8ae": 8347, // Becton
		"9cc8fc": 125,  // ARRIS
		"9cc950": 8348, // Baumer
		"9cc9eb": 1368, // Netgear
		"9ccad9": 457,  // Nokia
		"9ccc83": 826,  // Juniper
		"9cd24b": 3032, // ZTE
		"9cd332": 8349, // PLC
		"9cd35b": 152,  // Samsung
		"9cd36d": 1368, // Netgear
		"9cd48b": 8350, // Innolux
		"9cd57d": 2,    // Cisco
		"9cd643": 803,  // D-Link
		"9cd917": 687,  // Motorola
		"9cd9cb": 8351, // Lesira
		"9cda3e": 421,  // Intel
		"9cdb07": 8352, // Thum+Mahr
		"9cdc71": 302,  // HP
		"9ce063": 152,  // Samsung
		"9ce10e": 8353, // NCTech
		"9ce176": 2,    // Cisco
		"9ce230": 8354, // Julong
		"9ce33f": 545,  // Apple
		"9ce374": 3335, // Huawei
		"9ce635": 1427, // Nintendo
		"9ce65e": 545,  // Apple
		"9ce6e7": 152,  // Samsung
		"9ce7bd": 8355, // Winduskorea
		"9ce82b": 6370, // Vivo
		"9ce91c": 3032, // ZTE
		"9cebe8": 8356, // BizLink
		"9cec61": 3335, // Huawei
		"9cedfa": 8357, // EVUlution
		"9cf155": 457,  // Nokia
		"9cf387": 545,  // Apple
		"9cf48e": 545,  // Apple
		"9cfbd5": 6370, // Vivo
		"9cfc01": 545,  // Apple
		"9cfc28": 545,  // Apple
		"9cfcd1": 8358, // Aetheris
		"9cfce8": 421,  // Intel
		"9cfea1": 4922, // Fiberhome
		"9cffbe": 8359, // OTSL
		"9cffc2": 8360, // AVI
		"a002dc": 5439, // Amazon
		"a00460": 1368, // Netgear
		"a00798": 152,  // Samsung
		"a0086f": 3335, // Huawei
		"a0092e": 3032, // ZTE
		"a0094c": 8361, // CenturyLink
		"a009ed": 620,  // Avaya
		"a00abf": 8362, // Wieson
		"a00bba": 152,  // Samsung
		"a00f37": 2,    // Cisco
		"a01081": 152,  // Samsung
		"a01290": 620,  // Avaya
		"a013cb": 4922, // Fiberhome
		"a0143d": 2563, // Parrot
		"a0165c": 8363, // Triteka
		"a01828": 545,  // Apple
		"a01842": 3870, // Comtrend
		"a01b29": 2049, // Sagemcom
		"a01c87": 3506, // Unionman
		"a01c8d": 3335, // Huawei
		"a01d48": 302,  // HP
		"a01e0b": 8364, // MINIX
		"a020a6": 6377, // Espressif
		"a02195": 152,  // Samsung
		"a021b7": 1368, // Netgear
		"a022de": 6370, // Vivo
		"a0239f": 2,    // Cisco
		"a027b6": 152,  // Samsung
		"a02919": 102,  // Dell
		"a02942": 421,  // Intel
		"a02bb8": 302,  // HP
		"a02c36": 6441, // FN-Link
		"a031db": 3335, // Huawei
		"a03299": 2665, // Lenovo
		"a0341b": 8365, // Adero
		"a03679": 3335, // Huawei
		"a0369f": 421,  // Intel
		"a036fa": 8366, // Ettus
		"a039ee": 2049, // Sagemcom
		"a039f7": 869,  // LG
		"a03b1b": 8367, // Inspire
		"a03be3": 545,  // Apple
		"a03d6e": 2,    // Cisco
		"a03d6f": 2,    // Cisco
		"a04025": 8368, // Actioncable
		"a04041": 8369, // SAMWONFA
		"a0406f": 3335, // Huawei
		"a040a0": 1368, // Netgear
		"a0412d": 8370, // Lansen
		"a04147": 3335, // Huawei
		"a0423f": 6159, // Tyan
		"a0445c": 3335, // Huawei
		"a0481c": 302,  // HP
		"a04a5e": 612,  // Microsoft
		"a04cc1": 8371, // Helixtech
		"a04e01": 4682, // CENTRAL
		"a04e04": 457,  // Nokia
		"a04ea7": 545,  // Apple
		"a04ecf": 545,  // Apple
		"a04f85": 869,  // LG
		"a0510b": 421,  // Intel
		"a051c6": 620,  // Avaya
		"a0554f": 2,    // Cisco
		"a055de": 125,  // ARRIS
		"a056f3": 545,  // Apple
		"a057e3": 3335, // Huawei
		"a05950": 421,  // Intel
		"a05b21": 8372, // ENVINET
		"a05dc1": 8373, // TMCT
		"a05de7": 8374, // DIRECTV
		"a05e6b": 8375, // MELPER
		"a06090": 152,  // Samsung
		"a06260": 67,   // Private
		"a06391": 1368, // Netgear
		"a0648f": 2545, // Askey
		"a06518": 8376, // Vnpt
		"a06610": 4,    // Fujitsu
		"a0687e": 125,  // ARRIS
		"a06986": 8377, // Wellav
		"a06a00": 4213, // Verilink
		"a06a44": 3468, // Vizio
		"a06cec": 4556, // RIM
		"a06faa": 869,  // LG
		"a070b7": 3335, // Huawei
		"a071a9": 457,  // Nokia
		"a0722c": 531,  // HUMAX
		"a07332": 8378, // Cashmaster
		"a073fc": 8379, // Rancore
		"a07591": 152,  // Samsung
		"a075ea": 8380, // BoxLock
		"a0764e": 6377, // Espressif
		"a07751": 8381, // ASMedia
		"a07771": 8382, // Vialis
		"a07817": 545,  // Apple
		"a078ba": 2278, // Pantech
		"a08069": 421,  // Intel
		"a0821f": 152,  // Samsung
		"a082c7": 8383, // P.T.I
		"a084cb": 8384, // SonicSensory
		"a085fc": 612,  // Microsoft
		"a086c6": 5698, // Xiaomi
		"a08869": 421,  // Intel
		"a088b4": 421,  // Intel
		"a08c9b": 8385, // Xtreme
		"a08cf8": 3335, // Huawei
		"a08cfd": 302,  // HP
		"a08d16": 3335, // Huawei
		"a08e78": 2049, // Sagemcom
		"a090de": 8386, // Veedims
		"a09169": 869,  // LG
		"a091a2": 6924, // OnePlus
		"a091c8": 3032, // ZTE
		"a09351": 2,    // Cisco
		"a09805": 8387, // OpenVox
		"a0999b": 545,  // Apple
		"a09d91": 8388, // SoundBridge
		"a09e1a": 4480, // Polar
		"a09f7a": 803,  // D-Link
		"a0a0dc": 3335, // Huawei
		"a0a23c": 8389, // Gpms
		"a0a309": 545,  // Apple
		"a0a33b": 3335, // Huawei
		"a0a3b8": 8390, // Wiscloud
		"a0a3e2": 2247, // Actiontec
		"a0a3f0": 803,  // D-Link
		"a0a4c5": 421,  // Intel
		"a0a65c": 5561, // Supercomputing
		"a0a8cd": 421,  // Intel
		"a0aafd": 8391, // EraThink
		"a0ab1b": 803,  // D-Link
		"a0ac69": 152,  // Samsung
		"a0ada1": 8392, // JMR
		"a0afbd": 421,  // Intel
		"a0b3cc": 302,  // HP
		"a0b439": 2,    // Cisco
		"a0b4a5": 152,  // Samsung
		"a0b4bf": 644,  // InfiNet
		"a0b53c": 6393, // Technicolor
		"a0b549": 2640, // Arcadyan
		"a0b9ed": 8393, // Skytap
		"a0bdcd": 3513, // BSkyB
		"a0bfa5": 8394, // Coresys
		"a0c3de": 8395, // Triton
		"a0c562": 125,  // ARRIS
		"a0c589": 421,  // Intel
		"a0c9a0": 2057, // Murata
		"a0cbfd": 152,  // Samsung
		"a0cc2b": 2057, // Murata
		"a0cf5b": 2,    // Cisco
		"a0cff5": 3032, // ZTE
		"a0d05b": 152,  // Samsung
		"a0d0dc": 5439, // Amazon
		"a0d12a": 8396, // AXPRO
		"a0d2b1": 5439, // Amazon
		"a0d37a": 421,  // Intel
		"a0d3c1": 302,  // HP
		"a0d635": 8397, // WBS
		"a0d722": 152,  // Samsung
		"a0d795": 545,  // Apple
		"a0d7a0": 3335, // Huawei
		"a0d7f3": 152,  // Samsung
		"a0d807": 3335, // Huawei
		"a0d83d": 4922, // Fiberhome
		"a0dc04": 8398, // Becker-Antriebe
		"a0dd97": 8399, // PolarLink
		"a0dde5": 3206, // Sharp
		"a0de0f": 3335, // Huawei
		"a0df15": 3335, // Huawei
		"a0e0af": 2,    // Cisco
		"a0e201": 8400, // AVTrace
		"a0e453": 101,  // Sony
		"a0e4cb": 2693, // ZyXEL
		"a0e5e9": 8401, // enimai
		"a0e617": 8402, // Matis
		"a0e70b": 421,  // Intel
		"a0e7ae": 125,  // ARRIS
		"a0eb76": 8403, // AirCUVE
		"a0ec80": 3032, // ZTE
		"a0ecf9": 2,    // Cisco
		"a0edcd": 545,  // Apple
		"a0f3c1": 1595, // TP-Link
		"a0f419": 457,  // Nokia
		"a0f450": 1341, // HTC
		"a0f459": 6441, // FN-Link
		"a0f479": 3335, // Huawei
		"a0f849": 2,    // Cisco
		"a0f9e0": 8404, // Vivatel
		"a0fbc5": 545,  // Apple
		"a0ff70": 6393, // Technicolor
		"a400e2": 3335, // Huawei
		"a40130": 8405, // ABIsystems
		"a402b9": 421,  // Intel
		"a40450": 3203, // nFore
		"a4056e": 6540, // Tiinlab
		"a405d6": 125,  // ARRIS
		"a407b6": 152,  // Samsung
		"a40801": 5439, // Amazon
		"a408ea": 2057, // Murata
		"a408f5": 2049, // Sagemcom
		"a40bed": 8406, // Carry
		"a40cc3": 2,    // Cisco
		"a40e2b": 7235, // Facebook
		"a41115": 1926, // Bosch
		"a41162": 8407, // Arlo
		"a41194": 2665, // Lenovo
		"a4134e": 8408, // Luxul
		"a41588": 125,  // ARRIS
		"a416e7": 3335, // Huawei
		"a41752": 8409, // Hifocus
		"a4178b": 3335, // Huawei
		"a41875": 2,    // Cisco
		"a41908": 4922, // Fiberhome
		"a41a3a": 1595, // TP-Link
		"a41f72": 102,  // Dell
		"a4218a": 76,   // Nortel
		"a424b3": 8410, // FlatFrog
		"a424dd": 8411, // Cambrionix
		"a4251b": 620,  // Avaya
		"a42983": 1041, // Boeing
		"a429b7": 8412, // bluesky
		"a42a95": 803,  // D-Link
		"a42b8c": 1368, // Netgear
		"a42bb0": 1595, // TP-Link
		"a4307a": 152,  // Samsung
		"a43111": 7114, // ZIV
		"a43135": 545,  // Apple
		"a433d1": 8413, // Fibrlink
		"a433d7": 6429, // MitraStar
		"a434d9": 421,  // Intel
		"a4352d": 8414, // TRIZ
		"a438cc": 1427, // Nintendo
		"a43a69": 8415, // Vers
		"a43b0e": 3335, // Huawei
		"a44027": 3032, // ZTE
		"a4423b": 421,  // Intel
		"a4438c": 125,  // ARRIS
		"a444d1": 4031, // Wingtech
		"a44519": 5698, // Xiaomi
		"a4466b": 8416, // EOC
		"a446b4": 3335, // Huawei
		"a44ad3": 4756, // ST
		"a44bd5": 5698, // Xiaomi
		"a44c11": 2,    // Cisco
		"a44cc8": 102,  // Dell
		"a44e31": 421,  // Intel
		"a45046": 5698, // Xiaomi
		"a45055": 8417, // Busware.DE
		"a45129": 8418, // XAG
		"a4516f": 612,  // Microsoft
		"a4530e": 2,    // Cisco
		"a45590": 5698, // Xiaomi
		"a45602": 8419, // fenglian
		"a4561b": 8420, // MCOT
		"a45630": 2,    // Cisco
		"a456cc": 6393, // Technicolor
		"a45802": 8421, // Shin-IL
		"a45a1c": 8422, // smart-electronic
		"a45c27": 1427, // Nintendo
		"a45d36": 302,  // HP
		"a45e5a": 8423, // ACTIVIO
		"a45e60": 545,  // Apple
		"a45f9b": 8424, // Nexell
		"a46011": 1647, // Verifone
		"a46032": 2254, // MRV
		"a46191": 8425, // NamJunSa
		"a46706": 545,  // Apple
		"a468bc": 8426, // Oakley
		"a46bb6": 421,  // Intel
		"a46c2a": 2,    // Cisco
		"a46cf1": 152,  // Samsung
		"a46da4": 3335, // Huawei
		"a470d6": 687,  // Motorola
		"a47174": 3335, // Huawei
		"a47733": 3522, // Google
		"a47760": 457,  // Nokia
		"a477f3": 545,  // Apple
		"a47806": 2,    // Cisco
		"a47886": 620,  // Avaya
		"a479e4": 8427, // KLINFO
		"a47aa4": 125,  // ARRIS
		"a47acf": 8428, // Vibicom
		"a47b1a": 3335, // Huawei
		"a47b2c": 457,  // Nokia
		"a47b85": 8429, // ULTIMEDIA
		"a47b9d": 6377, // Espressif
		"a47c14": 8430, // ChargeStorm
		"a47c1f": 8431, // Cobham
		"a47cc9": 3335, // Huawei
		"a47e39": 3032, // ZTE
		"a481ee": 457,  // Nokia
		"a48269": 8432, // Datrium
		"a483e7": 545,  // Apple
		"a48431": 152,  // Samsung
		"a4856b": 8433, // Q
		"a48873": 2,    // Cisco
		"a48cc0": 8434, // JLG
		"a48cdb": 2665, // Lenovo
		"a48d3b": 3468, // Vizio
		"a48e0a": 8435, // DeLaval
		"a491b1": 6393, // Technicolor
		"a492cb": 457,  // Nokia
		"a4933f": 3335, // Huawei
		"a4934c": 2,    // Cisco
		"a49426": 8436, // Elgama-Elektronika
		"a49733": 2545, // Askey
		"a49813": 125,  // ARRIS
		"a49947": 3335, // Huawei
		"a49a58": 152,  // Samsung
		"a49b4f": 3335, // Huawei
		"a49bcd": 2,    // Cisco
		"a49d49": 8437, // Ketra
		"a49edb": 8438, // AutoCrib
		"a4a1c2": 306,  // Ericsson
		"a4a1e4": 8439, // Innotube
		"a4a46b": 3335, // Huawei
		"a4a4d3": 8440, // Bluebank
		"a4a6a9": 67,   // Private
		"a4aafe": 3335, // Huawei
		"a4ac0f": 3335, // Huawei
		"a4ad00": 8441, // Ragsdale
		"a4b197": 545,  // Apple
		"a4b1c1": 421,  // Intel
		"a4b1e9": 6393, // Technicolor
		"a4b239": 2,    // Cisco
		"a4b2a7": 8442, // Adaxys
		"a4b439": 2,    // Cisco
		"a4b61e": 3335, // Huawei
		"a4b805": 545,  // Apple
		"a4ba76": 3335, // Huawei
		"a4badb": 102,  // Dell
		"a4bb6d": 102,  // Dell
		"a4bdc4": 3335, // Huawei
		"a4be2b": 3335, // Huawei
		"a4bf01": 421,  // Intel
		"a4c0e1": 1427, // Nintendo
		"a4c337": 545,  // Apple
		"a4c361": 545,  // Apple
		"a4c3f0": 421,  // Intel
		"a4c494": 421,  // Intel
		"a4c54e": 3335, // Huawei
		"a4c64f": 3335, // Huawei
		"a4c69a": 152,  // Samsung
		"a4c74b": 3335, // Huawei
		"a4c7de": 1640, // Cambridge
		"a4caa0": 3335, // Huawei
		"a4cc32": 8443, // Inficomm
		"a4ceda": 2640, // Arcadyan
		"a4cf12": 6377, // Espressif
		"a4d094": 2632, // Vivavis
		"a4d18c": 545,  // Apple
		"a4d1d1": 8444, // ECOtality
		"a4d1d2": 545,  // Apple
		"a4d73c": 46,   // Epson
		"a4d795": 4031, // Wingtech
		"a4d856": 8445, // Gimbal
		"a4d931": 545,  // Apple
		"a4d990": 152,  // Samsung
		"a4da3f": 8446, // Bionics
		"a4db30": 2874, // Liteon
		"a4dcbe": 3335, // Huawei
		"a4e11a": 826,  // Juniper
		"a4e31b": 457,  // Nokia
		"a4e32e": 1655, // Silicon
		"a4e4b8": 2221, // BlackBerry
		"a4e57c": 6377, // Espressif
		"a4e597": 8447, // Gessler
		"a4e731": 457,  // Nokia
		"a4e7e4": 8448, // Connex
		"a4e975": 545,  // Apple
		"a4e9a3": 8449, // Honest
		"a4ea8e": 185,  // Extreme
		"a4ebd3": 152,  // Samsung
		"a4ed4e": 125,  // ARRIS
		"a4ee57": 46,   // Epson
		"a4ef15": 7032, // AltoBeam
		"a4ef52": 8450, // Telewave
		"a4f1e8": 545,  // Apple
		"a4f33b": 3032, // ZTE
		"a4f465": 6282, // Itel
		"a4f4c2": 8376, // Vnpt
		"a4ff95": 457,  // Nokia
		"a8016d": 5243, // Aiwa
		"a80180": 8451, // IMAGO
		"a802db": 3032, // ZTE
		"a8032a": 6377, // Espressif
		"a80577": 8452, // Netlist
		"a80600": 152,  // Samsung
		"a80c0d": 2,    // Cisco
		"a80c63": 3335, // Huawei
		"a811fc": 125,  // ARRIS
		"a81306": 6370, // Vivo
		"a81374": 2154, // Panasonic
		"a8154d": 1595, // TP-Link
		"a81559": 8453, // Breathometer
		"a816b2": 869,  // LG
		"a816d0": 152,  // Samsung
		"a81b18": 8454, // XTS
		"a81d16": 3005, // Azurewave
		"a81e84": 3072, // Quanta
		"a82066": 545,  // Apple
		"a82316": 457,  // Nokia
		"a823fe": 869,  // LG
		"a824b8": 457,  // Nokia
		"a825eb": 1640, // Cambridge
		"a826d9": 1341, // HTC
		"a82bb5": 6295, // Edgecore
		"a82bb9": 152,  // Samsung
		"a82bcd": 3335, // Huawei
		"a830bc": 152,  // Samsung
		"a8329a": 6212, // Digicom
		"a8346a": 152,  // Samsung
		"a83512": 3335, // Huawei
		"a8367a": 8455, // frogblue
		"a83759": 3335, // Huawei
		"a83944": 2247, // Actiontec
		"a83b5c": 3335, // Huawei
		"a83ccb": 8456, // Rossma
		"a84025": 8457, // Oxide
		"a84041": 8458, // Dragino
		"a84397": 8459, // Innogrit
		"a84481": 457,  // Nokia
		"a845cd": 8460, // Siselectron
		"a845e9": 8461, // Firich
		"a848fa": 6377, // Espressif
		"a8494d": 3335, // Huawei
		"a849a5": 8462, // Lisantech
		"a84a28": 545,  // Apple
		"a84b4d": 152,  // Samsung
		"a84d4a": 8463, // Audiowise
		"a84e3f": 870,  // Hitron
		"a84fb1": 2,    // Cisco
		"a85081": 3335, // Huawei
		"a8515b": 152,  // Samsung
		"a8537d": 7485, // Mist
		"a854a2": 8464, // Heimgard
		"a854b2": 1592, // Wistron
		"a8574e": 1595, // TP-Link
		"a85840": 1640, // Cambridge
		"a8587c": 8465, // Shoogee
		"a85ae0": 3335, // Huawei
		"a85b6c": 1926, // Bosch
		"a85b78": 545,  // Apple
		"a85bb7": 545,  // Apple
		"a85bf3": 8466, // Audivo
		"a85bf7": 1685, // Aruba
		"a85c2c": 545,  // Apple
		"a85e45": 1806, // ASUS
		"a85ee4": 8467, // 12Sided
		"a860b6": 545,  // Apple
		"a8610a": 8468, // Arduino
		"a861aa": 8469, // Cloudview
		"a862a2": 8470, // Jiwumedia
		"a8637d": 803,  // D-Link
		"a863df": 8471, // Displaire
		"a864f1": 421,  // Intel
		"a8667f": 545,  // Apple
		"a8671e": 8472, // Ratp
		"a8698c": 11,   // Oracle
		"a86a6f": 4556, // RIM
		"a86abb": 2049, // Sagemcom
		"a86ac1": 8473, // HanbitEDS
		"a86d5f": 2051, // Raisecom
		"a86daa": 421,  // Intel
		"a86e4e": 3335, // Huawei
		"a8705d": 125,  // ARRIS
		"a870a5": 8474, // UniComm
		"a87285": 3226, // IDT
		"a87484": 3032, // ZTE
		"a875d6": 8475, // FreeTek
		"a875e2": 8476, // Aventura
		"a87650": 152,  // Samsung
		"a8776f": 8477, // Zonoff
		"a8798d": 152,  // Samsung
		"a87b39": 457,  // Nokia
		"a87c01": 152,  // Samsung
		"a87d12": 3335, // Huawei
		"a87e33": 457,  // Nokia
		"a87eea": 421,  // Intel
		"a8817e": 545,  // Apple
		"a88195": 152,  // Samsung
		"a885d7": 8478, // Sangfor
		"a886dd": 545,  // Apple
		"a887b3": 152,  // Samsung
		"a88808": 545,  // Apple
		"a88940": 3335, // Huawei
		"a88c3e": 612,  // Microsoft
		"a88e24": 545,  // Apple
		"a8913d": 545,  // Apple
		"a8922c": 869,  // LG
		"a89675": 687,  // Motorola
		"a8968a": 545,  // Apple
		"a897cd": 125,  // ARRIS
		"a897dc": 372,  // IBM
		"a898c6": 8479, // Shinbo
		"a8995c": 8480, // aizo
		"a89969": 102,  // Dell
		"a89a93": 2049, // Sagemcom
		"a89ad7": 457,  // Nokia
		"a89b10": 8481, // inMotion
		"a89ca4": 8482, // Furrion
		"a89ced": 5698, // Xiaomi
		"a89d21": 2,    // Cisco
		"a89fba": 152,  // Samsung
		"a89fec": 125,  // ARRIS
		"a8a089": 8483, // Tactical
		"a8a159": 4720, // ASRock
		"a8a198": 6434, // TCT
		"a8a668": 3032, // ZTE
		"a8b088": 5821, // eero
		"a8b0ae": 8484, // Leoni
		"a8b13b": 302,  // HP
		"a8b1d4": 2,    // Cisco
		"a8b2da": 4,    // Fujitsu
		"a8b456": 2,    // Cisco
		"a8b57c": 1916, // Roku
		"a8b86e": 869,  // LG
		"a8b9b3": 6831, // Essys
		"a8bbcf": 545,  // Apple
		"a8bd27": 302,  // HP
		"a8bd3a": 3506, // Unionman
		"a8be27": 545,  // Apple
		"a8c092": 3335, // Huawei
		"a8c0ea": 8485, // Pepwave
		"a8c222": 8486, // TM-Research
		"a8c252": 3335, // Huawei
		"a8c266": 531,  // HUMAX
		"a8c83a": 3335, // Huawei
		"a8c87f": 8487, // Roqos
		"a8ca7b": 3335, // Huawei
		"a8cab9": 152,  // Samsung
		"a8ccc5": 2712, // Saab
		"a8ce90": 8488, // CVC
		"a8d081": 3335, // Huawei
		"a8d0e3": 5480, // Systech
		"a8d0e5": 826,  // Juniper
		"a8d3c8": 8489, // Topcon
		"a8d3f7": 2640, // Arcadyan
		"a8d4e0": 3335, // Huawei
		"a8d88a": 8490, // Wyconn
		"a8da0c": 8146, // Servercom
		"a8db03": 152,  // Samsung
		"a8e018": 457,  // Nokia
		"a8e3ee": 101,  // Sony
		"a8e539": 8491, // Moimstone
		"a8e544": 3335, // Huawei
		"a8e621": 5439, // Amazon
		"a8e705": 4922, // Fiberhome
		"a8e81e": 8492, // ATW
		"a8e824": 8493, // Inim
		"a8e978": 3335, // Huawei
		"a8eec6": 8494, // Muuselabs/SA
		"a8ef26": 8495, // Tritonwave
		"a8f266": 3335, // Huawei
		"a8f274": 152,  // Samsung
		"a8f5ac": 3335, // Huawei
		"a8f5dd": 125,  // ARRIS
		"a8f766": 8496, // ITE
		"a8f7e0": 4965, // PLANET
		"a8f94b": 8497, // Eltex
		"a8fad8": 545,  // Apple
		"a8fe9d": 545,  // Apple
		"a8ffba": 3335, // Huawei
		"ac00d0": 3032, // ZTE
		"ac02ca": 8498, // HI
		"ac02ef": 8499, // Comsis
		"ac0425": 8500, // ball-b
		"ac0613": 8501, // Senselogix
		"ac075f": 3335, // Huawei
		"ac0bfb": 6377, // Espressif
		"ac0d1b": 869,  // LG
		"ac1203": 421,  // Intel
		"ac139c": 202,  // Adtran
		"ac1461": 8502, // ATAW
		"ac14d2": 8503, // wi-daq
		"ac1585": 8504, // silergy
		"ac15f4": 545,  // Apple
		"ac162d": 302,  // HP
		"ac1826": 46,   // Epson
		"ac1d06": 545,  // Apple
		"ac1e92": 152,  // Samsung
		"ac1e9e": 5698, // Xiaomi
		"ac1f74": 545,  // Apple
		"ac202e": 870,  // Hitron
		"ac20aa": 8505, // DMATEK
		"ac2205": 358,  // Compal
		"ac220b": 1806, // ASUS
		"ac2316": 7485, // Mist
		"ac2334": 6296, // Infinix
		"ac293a": 545,  // Apple
		"ac2b6e": 421,  // Intel
		"ac2da3": 8506, // TXTR
		"ac2da9": 6266, // Tecno
		"ac2fa8": 8507, // Humannix
		"ac3328": 3335, // Huawei
		"ac35ee": 6441, // FN-Link
		"ac3613": 152,  // Samsung
		"ac3743": 1341, // HTC
		"ac37c9": 8508, // RAID
		"ac3870": 2665, // Lenovo
		"ac3a67": 2,    // Cisco
		"ac3a7a": 1916, // Roku
		"ac3b77": 2049, // Sagemcom
		"ac3c0b": 545,  // Apple
		"ac3c8e": 833,  // Flextronics
		"ac3cb4": 8509, // Nilan
		"ac4122": 8510, // Eclipse
		"ac4228": 8511, // Parta
		"ac4330": 4966, // Versa
		"ac44f2": 727,  // Yamaha
		"ac471b": 3335, // Huawei
		"ac4723": 8512, // Genelec
		"ac49db": 545,  // Apple
		"ac4a56": 2,    // Cisco
		"ac4a67": 2,    // Cisco
		"ac4b1e": 8513, // Integri-Sys.com
		"ac4bc8": 826,  // Juniper
		"ac4e91": 3335, // Huawei
		"ac4ffc": 8514, // SVS-VISTEK
		"ac5036": 6835, // Pi-Coral
		"ac5093": 153,  // Magna
		"ac512c": 6296, // Infinix
		"ac5135": 8515, // MPI
		"ac51ee": 1640, // Cambridge
		"ac562c": 647,  // Lava
		"ac567b": 5436, // Sunnovo
		"ac5a14": 152,  // Samsung
		"ac5afc": 421,  // Intel
		"ac5d5c": 6441, // FN-Link
		"ac5e14": 3335, // Huawei
		"ac5e8c": 8516, // Utillink
		"ac5f3e": 152,  // Samsung
		"ac5fea": 6924, // OnePlus
		"ac6089": 3335, // Huawei
		"ac60b6": 306,  // Ericsson
		"ac6123": 8517, // Drivven
		"ac6175": 3335, // Huawei
		"ac61b9": 8518, // WAMA
		"ac61ea": 545,  // Apple
		"ac63be": 5439, // Amazon
		"ac6417": 300,  // Siemens
		"ac6462": 3032, // ZTE
		"ac6490": 3335, // Huawei
		"ac64cf": 6441, // FN-Link
		"ac6706": 2738, // Ruckus
		"ac675d": 421,  // Intel
		"ac6784": 3522, // Google
		"ac67b2": 6377, // Espressif
		"ac6c90": 152,  // Samsung
		"ac6f4f": 8519, // Enspert
		"ac6fbb": 8520, // TATUNG
		"ac6fd9": 8521, // Valueplus
		"ac7236": 8522, // Lexking
		"ac7289": 421,  // Intel
		"ac74b1": 421,  // Intel
		"ac74c4": 8523, // Maytronics
		"ac751d": 3335, // Huawei
		"ac78d1": 826,  // Juniper
		"ac7a42": 8524, // iConnectivity
		"ac7a4d": 430,  // Alpsalpine
		"ac7a56": 2,    // Cisco
		"ac7ba1": 421,  // Intel
		"ac7e01": 3335, // Huawei
		"ac7e8a": 2,    // Cisco
		"ac7f3e": 545,  // Apple
		"ac80ae": 4922, // Fiberhome
		"ac80d6": 8525, // Hexatronic
		"ac8112": 1450, // Gemtek
		"ac81f3": 457,  // Nokia
		"ac8247": 421,  // Intel
		"ac83f0": 8526, // ImmediaTV
		"ac83f3": 4499, // AMPAK
		"ac84c6": 1595, // TP-Link
		"ac84c9": 2049, // Sagemcom
		"ac853d": 3335, // Huawei
		"ac87a3": 545,  // Apple
		"ac88fd": 545,  // Apple
		"ac8995": 3005, // Azurewave
		"ac8b9c": 8527, // Primera
		"ac8ba9": 2979, // Ubiquiti
		"ac8d14": 8528, // Smartrove
		"ac8d34": 3335, // Huawei
		"ac8ff8": 457,  // Nokia
		"ac9085": 545,  // Apple
		"ac9232": 3335, // Huawei
		"ac932f": 457,  // Nokia
		"ac9572": 8529, // Jovision
		"ac976c": 8530, // Greenliant
		"ac9929": 3335, // Huawei
		"ac9a96": 5299, // Maxlinear
		"ac9b0a": 101,  // Sony
		"ac9e17": 1806, // ASUS
		"aca016": 2,    // Cisco
		"aca22c": 8531, // Baycity
		"aca31e": 1685, // Aruba
		"aca88e": 3206, // Sharp
		"aca919": 8532, // TrekStor
		"aca9a0": 8533, // Audioengine
		"acabbf": 8534, // AthenTek
		"acae19": 1916, // Roku
		"acafb9": 152,  // Samsung
		"acb313": 125,  // ARRIS
		"acb3b5": 3335, // Huawei
		"acb57d": 2874, // Liteon
		"acb859": 8535, // Uniband
		"acbb61": 7451, // YSTen
		"acbc32": 545,  // Apple
		"acbcd9": 2,    // Cisco
		"acbd0b": 8536, // Leimac
		"acbd70": 3335, // Huawei
		"acbe75": 8537, // Ufine
		"acbeb6": 8538, // Visualedge
		"acc1ee": 5698, // Xiaomi
		"acc25d": 4922, // Fiberhome
		"acc33a": 152,  // Samsung
		"acc595": 8539, // Graphite
		"acc662": 6429, // MitraStar
		"acc73f": 8540, // Vitsmo
		"acc935": 8541, // Ness
		"acca54": 8542, // Telldus
		"acca8e": 8543, // ODA
		"accaba": 8544, // Midokura
		"accc8e": 5132, // Axis
		"accf23": 8545, // Hi-flying
		"accf5c": 545,  // Apple
		"accf85": 3335, // Huawei
		"acd074": 6377, // Espressif
		"acd364": 21,   // ABB
		"acd618": 6924, // OnePlus
		"acd9d6": 8546, // tci
		"acdb48": 125,  // ARRIS
		"acdcca": 3335, // Huawei
		"acde48": 67,   // Private
		"ace010": 2874, // Liteon
		"ace215": 3335, // Huawei
		"ace2d3": 302,  // HP
		"ace342": 3335, // Huawei
		"ace348": 8547, // MadgeTech
		"ace4b5": 545,  // Apple
		"ace5f0": 8548, // Doppler
		"ace87b": 3335, // Huawei
		"ace87e": 8549, // Bytemark
		"ace97f": 8550, // IoT
		"ace9aa": 8551, // Hay
		"aceb51": 3858, // Universal
		"acec80": 125,  // ARRIS
		"acec85": 5821, // eero
		"aced5c": 421,  // Intel
		"acee3b": 8552, // 6harmonics
		"acee9e": 152,  // Samsung
		"acf0b2": 8553, // Becker
		"acf108": 869,  // LG
		"acf1df": 803,  // D-Link
		"acf2c5": 2,    // Cisco
		"acf5e6": 2,    // Cisco
		"acf6f7": 869,  // LG
		"acf7f3": 5698, // Xiaomi
		"acf8cc": 125,  // ARRIS
		"acf970": 3335, // Huawei
		"acf97e": 8554, // Elesys
		"acfaa5": 8555, // digitron
		"acfdce": 421,  // Intel
		"acfdec": 545,  // Apple
		"acfe05": 6282, // Itel
		"b00073": 1592, // Wistron
		"b000b4": 2,    // Cisco
		"b00247": 4499, // AMPAK
		"b0027e": 8556, // Muller
		"b00594": 2874, // Liteon
		"b00875": 3335, // Huawei
		"b009d3": 8557, // Avizia
		"b009da": 6947, // Ring
		"b00ad5": 3032, // ZTE
		"b00cd1": 302,  // HP
		"b01266": 8558, // Futaba-Kikaku
		"b01408": 8559, // Lightspeed
		"b01656": 3335, // Huawei
		"b01886": 3456, // SmarDTV
		"b019c6": 545,  // Apple
		"b01c91": 8560, // Elim
		"b01f29": 8561, // Helvetia
		"b0227a": 302,  // HP
		"b02491": 3335, // Huawei
		"b024f3": 8562, // Progeny
		"b025aa": 67,   // Private
		"b02628": 858,  // Broadcom
		"b02680": 2,    // Cisco
		"b027cf": 185,  // Extreme
		"b02a1f": 4031, // Wingtech
		"b02a43": 3522, // Google
		"b03366": 6370, // Vivo
		"b033a6": 826,  // Juniper
		"b03495": 545,  // Apple
		"b0358d": 457,  // Nokia
		"b0359f": 421,  // Intel
		"b035b5": 545,  // Apple
		"b03795": 869,  // LG
		"b03956": 1368, // Netgear
		"b03ace": 3335, // Huawei
		"b03cdc": 421,  // Intel
		"b03e51": 3513, // BSkyB
		"b03eb0": 8563, // MICRODIA
		"b04089": 8564, // Senient
		"b0411d": 8565, // ITTIM
		"b0435d": 8566, // NuLEDs
		"b04502": 3335, // Huawei
		"b04519": 6434, // TCT
		"b04530": 3513, // BSkyB
		"b046fc": 6429, // MitraStar
		"b047bf": 152,  // Samsung
		"b0481a": 545,  // Apple
		"b0487a": 1595, // TP-Link
		"b04a6a": 152,  // Samsung
		"b04e26": 1595, // TP-Link
		"b04f13": 102,  // Dell
		"b0518e": 8567, // Holl
		"b05508": 3335, // Huawei
		"b05706": 8568, // Vallox
		"b05ada": 302,  // HP
		"b05b67": 3335, // Huawei
		"b05c16": 4922, // Fiberhome
		"b05cda": 302,  // HP
		"b05ce5": 457,  // Nokia
		"b05dd4": 125,  // ARRIS
		"b06088": 421,  // Intel
		"b061c7": 8569, // Ericsson-LG
		"b065bd": 545,  // Apple
		"b06a41": 3522, // Google
		"b06ebf": 1806, // ASUS
		"b06fe0": 152,  // Samsung
		"b0700d": 457,  // Nokia
		"b0702d": 545,  // Apple
		"b072bf": 2057, // Murata
		"b0735d": 3335, // Huawei
		"b0739c": 5439, // Amazon
		"b0754d": 457,  // Nokia
		"b075d5": 3032, // ZTE
		"b0761b": 3335, // Huawei
		"b077ac": 125,  // ARRIS
		"b07870": 8570, // Wi-NEXT
		"b07908": 8571, // Cummings
		"b0793c": 8572, // Revolv
		"b07994": 687,  // Motorola
		"b07b25": 102,  // Dell
		"b07d47": 2,    // Cisco
		"b07d64": 421,  // Intel
		"b07fb9": 1368, // Netgear
		"b081d8": 8573, // I-Sys
		"b083d6": 125,  // ARRIS
		"b083fe": 102,  // Dell
		"b08900": 3335, // Huawei
		"b08991": 7469, // LGE
		"b089c2": 8574, // Zyptonite
		"b08b92": 3032, // ZTE
		"b08bcf": 2,    // Cisco
		"b08bd0": 2,    // Cisco
		"b08c75": 545,  // Apple
		"b08e1a": 8575, // URadio
		"b09074": 8576, // Fulan
		"b0907e": 2,    // Cisco
		"b09134": 8577, // Taleo
		"b0935b": 125,  // ARRIS
		"b09575": 1595, // TP-Link
		"b0958e": 1595, // TP-Link
		"b0966c": 8578, // Lanbowan
		"b0973a": 8579, // E-Fuel
		"b0982b": 2049, // Sagemcom
		"b0989f": 869,  // LG
		"b098bc": 3335, // Huawei
		"b09928": 4,    // Fujitsu
		"b099d7": 152,  // Samsung
		"b09fba": 545,  // Apple
		"b0a10a": 5065, // Pivotal
		"b0a454": 8580, // Tripwire
		"b0a460": 421,  // Intel
		"b0a4f0": 3335, // Huawei
		"b0a651": 2,    // Cisco
		"b0a6f5": 8581, // Xaptum
		"b0a737": 1916, // Roku
		"b0a7b9": 1595, // TP-Link
		"b0a86e": 826,  // Juniper
		"b0aa77": 2,    // Cisco
		"b0acd2": 3032, // ZTE
		"b0acfa": 4,    // Fujitsu
		"b0adaa": 620,  // Avaya
		"b0ae25": 8582, // Varikorea
		"b0b194": 3032, // ZTE
		"b0b28f": 2049, // Sagemcom
		"b0b2dc": 2693, // ZyXEL
		"b0b3ad": 531,  // HUMAX
		"b0b5e8": 8583, // Ruroc
		"b0b867": 302,  // HP
		"b0b98a": 1368, // Netgear
		"b0bb8b": 8584, // Wavetel
		"b0bbe5": 2049, // Sagemcom
		"b0be76": 1595, // TP-Link
		"b0be83": 545,  // Apple
		"b0bf99": 8585, // Wizitdongdo
		"b0c090": 7301, // Chicony
		"b0c19e": 3032, // ZTE
		"b0c205": 8586, // Bionime
		"b0c287": 6393, // Technicolor
		"b0c387": 8587, // GOEFER
		"b0c46c": 8588, // Senseit
		"b0c4e7": 152,  // Samsung
		"b0c53c": 2,    // Cisco
		"b0c554": 803,  // D-Link
		"b0c559": 152,  // Samsung
		"b0c69a": 826,  // Juniper
		"b0c745": 1077, // Buffalo
		"b0c787": 3335, // Huawei
		"b0ca68": 545,  // Apple
		"b0ccfe": 3335, // Huawei
		"b0d09c": 152,  // Samsung
		"b0d2f5": 8589, // Vello
		"b0d7c5": 8590, // Logipix
		"b0d7cc": 8591, // Tridonic
		"b0d888": 2154, // Panasonic
		"b0da00": 8592, // Cera
		"b0daf9": 125,  // ARRIS
		"b0dd74": 8464, // Heimgard
		"b0de28": 545,  // Apple
		"b0df3a": 152,  // Samsung
		"b0dfc1": 6267, // Tenda
		"b0e03c": 6434, // TCT
		"b0e17e": 3335, // Huawei
		"b0e235": 5698, // Xiaomi
		"b0e2e5": 4922, // Fiberhome
		"b0e4d5": 3522, // Google
		"b0e50e": 8593, // NRG
		"b0e5ed": 3335, // Huawei
		"b0e5f9": 545,  // Apple
		"b0e754": 1939, // 2Wire
		"b0e892": 46,   // Epson
		"b0e9fe": 8594, // Woan
		"b0eabc": 2545, // Askey
		"b0eb57": 3335, // Huawei
		"b0ec71": 152,  // Samsung
		"b0ecdd": 3335, // Huawei
		"b0ece1": 67,   // Private
		"b0ee45": 3005, // Azurewave
		"b0ee7b": 1916, // Roku
		"b0f1a3": 8595, // Fengfan
		"b0f1d8": 545,  // Apple
		"b0f1ec": 4499, // AMPAK
		"b0f530": 870,  // Hitron
		"b0f7c4": 5439, // Amazon
		"b0faeb": 2,    // Cisco
		"b0fc0d": 5439, // Amazon
		"b0fc36": 189,  // CyberTAN
		"b0febd": 67,   // Private
		"b0fee5": 3335, // Huawei
		"b4009c": 8596, // CableWorld
		"b40216": 2,    // Cisco
		"b4055d": 7723, // Inspur
		"b407f9": 152,  // Samsung
		"b40832": 8597, // TC
		"b40931": 3335, // Huawei
		"b40ac6": 8598, // DEXON
		"b40b44": 8599, // Smartisan
		"b40e96": 8600, // Heran
		"b40edc": 8601, // LG-Ericsson
		"b40ede": 421,  // Intel
		"b40f3b": 6267, // Tenda
		"b40fb3": 6370, // Vivo
		"b41489": 2,    // Cisco
		"b414e6": 3335, // Huawei
		"b41513": 3335, // Huawei
		"b4157e": 8602, // Celona
		"b417a8": 7235, // Facebook
		"b418d1": 545,  // Apple
		"b41a1d": 152,  // Samsung
		"b41bb0": 545,  // Apple
		"b41c30": 3032, // ZTE
		"b41cab": 8603, // ICR
		"b41def": 8604, // Internet
		"b42046": 5821, // eero
		"b42200": 3231, // Brother
		"b42330": 1109, // Itron
		"b424e7": 8605, // Codetek
		"b42875": 8606, // Futecho
		"b428f1": 8607, // E-Prime
		"b42a0e": 6393, // Technicolor
		"b42d56": 185,  // Extreme
		"b42e99": 1929, // Giga-Byte
		"b42ef8": 8608, // Eline
		"b43052": 3335, // Huawei
		"b431b8": 8609, // Aviwest
		"b436d1": 5696, // Renesas
		"b43741": 8610, // Consert
		"b437d8": 803,  // D-Link
		"b43a28": 152,  // Samsung
		"b43ae2": 3335, // Huawei
		"b43d08": 8611, // GX
		"b43e3b": 8612, // Viableware
		"b440a4": 545,  // Apple
		"b4430d": 8613, // Broadlink
		"b44326": 3335, // Huawei
		"b44506": 102,  // Dell
		"b4475e": 620,  // Avaya
		"b447f5": 6499, // Earda
		"b44bd2": 545,  // Apple
		"b45062": 8614, // EmBestor
		"b451f9": 8615, // NB
		"b45253": 732,  // Seagate
		"b4527d": 101,  // Sony
		"b4527e": 101,  // Sony
		"b45570": 8616, // Borea
		"b456b9": 8617, // Teraspek
		"b456e3": 545,  // Apple
		"b45861": 8618, // CRemote
		"b45d50": 1685, // Aruba
		"b45f84": 3032, // ZTE
		"b4608c": 4922, // Fiberhome
		"b461ff": 8619, // Lumigon
		"b46238": 8620, // Exablox
		"b46293": 152,  // Samsung
		"b46921": 421,  // Intel
		"b4695f": 6434, // TCT
		"b46bfc": 421,  // Intel
		"b46c47": 2154, // Panasonic
		"b46d83": 421,  // Intel
		"b46e08": 3335, // Huawei
		"b47064": 152,  // Samsung
		"b47443": 152,  // Samsung
		"b47447": 8621, // CoreOS
		"b4749f": 2545, // Askey
		"b4750e": 2469, // Belkin
		"b47947": 7341, // Nutanix
		"b479a7": 152,  // Samsung
		"b479c8": 2738, // Ruckus
		"b47af1": 302,  // HP
		"b47c9c": 5439, // Amazon
		"b481bf": 8622, // Meta-Networks
		"b48255": 8623, // Research
		"b482c5": 8624, // Relay2
		"b482fe": 2545, // Askey
		"b485e1": 545,  // Apple
		"b48655": 3335, // Huawei
		"b48901": 3335, // Huawei
		"b48a5f": 826,  // Juniper
		"b48b19": 545,  // Apple
		"b48c9d": 3005, // Azurewave
		"b4944e": 8625, // WeTelecom
		"b49691": 421,  // Intel
		"b49842": 3032, // ZTE
		"b499ba": 302,  // HP
		"b49cdf": 545,  // Apple
		"b49d02": 152,  // Samsung
		"b49d0b": 7284, // BQ
		"b49db4": 3919, // Axion
		"b49ee6": 8626, // Shenzhen
		"b4a25c": 665,  // Cambium
		"b4a4e3": 2,    // Cisco
		"b4a5a9": 8627, // MODI
		"b4a5ef": 2075, // Sercomm
		"b4a898": 3335, // Huawei
		"b4a8b9": 2,    // Cisco
		"b4a94f": 4254, // Mercury
		"b4a95a": 620,  // Avaya
		"b4a984": 1990, // Symantec
		"b4a9fc": 3072, // Quanta
		"b4a9fe": 8628, // GHIA
		"b4aa4d": 8629, // Ensequence
		"b4ab2c": 8630, // MtM
		"b4ae2b": 612,  // Microsoft
		"b4b017": 620,  // Avaya
		"b4b024": 1595, // TP-Link
		"b4b055": 3335, // Huawei
		"b4b15a": 300,  // Siemens
		"b4b291": 869,  // LG
		"b4b362": 3032, // ZTE
		"b4b52f": 302,  // HP
		"b4b5af": 8631, // Minsung
		"b4b676": 421,  // Intel
		"b4b686": 302,  // HP
		"b4b742": 5439, // Amazon
		"b4b859": 1203, // Texa
		"b4b88d": 8632, // Thuh
		"b4ba02": 8633, // Agatel
		"b4bff6": 152,  // Samsung
		"b4c26a": 797,  // Garmin
		"b4c4fc": 5698, // Xiaomi
		"b4c6f8": 8634, // Axilspot
		"b4c799": 185,  // Extreme
		"b4cc04": 8635, // Piranti
		"b4cce9": 8636, // Prosyst
		"b4cd27": 3335, // Huawei
		"b4ce40": 152,  // Samsung
		"b4cef6": 1341, // HTC
		"b4d135": 8637, // Cloudistics
		"b4d286": 7076, // Telechips
		"b4d5bd": 421,  // Intel
		"b4d64e": 8638, // Caldero
		"b4d8a9": 8639, // BetterBots
		"b4dd15": 8640, // ControlThings
		"b4de31": 2,    // Cisco
		"b4dedf": 3032, // ZTE
		"b4df3b": 8641, // Chromlech
		"b4dffa": 8642, // Litemax
		"b4e01d": 8643, // Conception
		"b4e0cd": 8644, // Fusion-io
		"b4e10f": 102,  // Dell
		"b4e1c4": 612,  // Microsoft
		"b4e1eb": 67,   // Private
		"b4e3f9": 1655, // Silicon
		"b4e454": 5439, // Amazon
		"b4e62a": 869,  // LG
		"b4e62d": 6377, // Espressif
		"b4e782": 8645, // Vivalnk
		"b4e8c9": 8646, // XADA
		"b4e9b0": 2,    // Cisco
		"b4ec02": 430,  // Alpsalpine
		"b4ed54": 8647, // Wohler
		"b4eeb4": 2545, // Askey
		"b4ef39": 152,  // Samsung
		"b4f0ab": 545,  // Apple
		"b4f18c": 3335, // Huawei
		"b4f1da": 869,  // LG
		"b4f267": 358,  // Compal
		"b4f2e8": 125,  // ARRIS
		"b4f323": 8648, // Petatel
		"b4f58e": 3335, // Huawei
		"b4f61c": 545,  // Apple
		"b4f7a1": 869,  // LG
		"b4f81e": 8649, // Kinova
		"b4f949": 8650, // optilink
		"b4fa48": 545,  // Apple
		"b4fbe3": 7032, // AltoBeam
		"b4fbe4": 2979, // Ubiquiti
		"b4fbf9": 3335, // Huawei
		"b4fc75": 8651, // SEMA
		"b4ff98": 3335, // Huawei
		"b80018": 8652, // Htel
		"b802a4": 8653, // Aeonsemi
		"b80305": 421,  // Intel
		"b805ab": 3032, // ZTE
		"b80716": 6370, // Vivo
		"b808cf": 421,  // Intel
		"b808d7": 3335, // Huawei
		"b8098a": 545,  // Apple
		"b810d4": 6363, // Masimo
		"b8114b": 2,    // Cisco
		"b812da": 8028, // Lvswitches
		"b81332": 4499, // AMPAK
		"b8145c": 3335, // Huawei
		"b814db": 6578, // Ohsung
		"b81619": 125,  // ARRIS
		"b817c2": 545,  // Apple
		"b8186f": 8654, // Oriental
		"b81999": 8655, // Nesys
		"b81daa": 869,  // LG
		"b81f5e": 8656, // Apption
		"b8208e": 2154, // Panasonic
		"b8259a": 8657, // Thalmic
		"b825b5": 8658, // Trakm8
		"b827c5": 3335, // Huawei
		"b829f7": 8659, // Blaster
		"b82a72": 102,  // Dell
		"b82aa9": 545,  // Apple
		"b82b68": 3335, // Huawei
		"b82ca0": 5987, // Resideo
		"b82d28": 4499, // AMPAK
		"b831b5": 612,  // Microsoft
		"b836d8": 8660, // Videoswitch
		"b8374a": 545,  // Apple
		"b83861": 2,    // Cisco
		"b83a08": 6267, // Tenda
		"b83a5a": 1685, // Aruba
		"b83a7b": 8661, // Worldplay
		"b83a9d": 3857, // Alarm.com
		"b83bcc": 5698, // Xiaomi
		"b83e59": 1916, // Roku
		"b83fd2": 432,  // Mellanox
		"b8415f": 2872, // ASP
		"b841a4": 545,  // Apple
		"b843e4": 8662, // Vlatacom
		"b844ae": 6434, // TCT
		"b844d9": 545,  // Apple
		"b8477a": 6068, // Dasan
		"b847c6": 8663, // SanJet
		"b84fd5": 612,  // Microsoft
		"b85001": 185,  // Extreme
		"b853ac": 545,  // Apple
		"b85510": 2128, // Zioncom
		"b85600": 3335, // Huawei
		"b856bd": 1178, // ITT
		"b85776": 8664, // lignex1
		"b857d8": 152,  // Samsung
		"b85810": 8665, // Numera
		"b8599f": 432,  // Mellanox
		"b859ce": 6499, // Earda
		"b85a73": 152,  // Samsung
		"b85af7": 8666, // Ouya
		"b85afe": 8667, // Handaer
		"b85d0a": 545,  // Apple
		"b85dc3": 3335, // Huawei
		"b85e7b": 152,  // Samsung
		"b85f98": 5439, // Amazon
		"b85fb0": 3335, // Huawei
		"b8616f": 145,  // Accton
		"b8621f": 2,    // Cisco
		"b8634d": 545,  // Apple
		"b863bc": 8668, // ROBOTIS
		"b8653b": 8669, // Bolymin
		"b86685": 2049, // Sagemcom
		"b869c2": 3940, // Sunitec
		"b869f4": 1784, // Routerboard.com
		"b86a97": 6295, // Edgecore
		"b86b23": 35,   // Toshiba
		"b86ce8": 152,  // Samsung
		"b870f4": 358,  // Compal
		"b87447": 8670, // Convergence
		"b875c0": 8671, // PayPal
		"b87826": 1427, // Nintendo
		"b8782e": 545,  // Apple
		"b87ac9": 300,  // Siemens
		"b87bc5": 545,  // Apple
		"b87cf2": 185,  // Extreme
		"b87ee5": 3542, // Intelbras
		"b88198": 421,  // Intel
		"b881fa": 545,  // Apple
		"b88303": 302,  // HP
		"b8857b": 3335, // Huawei
		"b88584": 102,  // Dell
		"b88687": 2874, // Liteon
		"b8876e": 8672, // Yandex
		"b887c6": 8673, // Prudential
		"b888e3": 358,  // Compal
		"b88a60": 421,  // Intel
		"b88a72": 5696, // Renesas
		"b88aec": 1427, // Nintendo
		"b88d12": 545,  // Apple
		"b88e82": 3335, // Huawei
		"b88ec6": 8674, // Stateless
		"b88edf": 8675, // Zencheer
		"b88f14": 8676, // Analytica
		"b89047": 545,  // Apple
		"b891c9": 3622, // Handreamnet
		"b89436": 3335, // Huawei
		"b89470": 374,  // Calix
		"b894e7": 5698, // Xiaomi
		"b89674": 8677, // AllDSP
		"b898ad": 687,  // Motorola
		"b898b0": 8678, // Atlona
		"b89919": 8679, // 7signal
		"b899b0": 8680, // Cohere
		"b89a2a": 421,  // Intel
		"b89aed": 8681, // OceanServer
		"b89bc9": 741,  // SMC
		"b89ea6": 8682, // Spbec-Mining
		"b89f09": 1592, // Wistron
		"b8a14a": 2051, // Raisecom
		"b8a175": 1916, // Roku
		"b8a377": 2,    // Cisco
		"b8a386": 803,  // D-Link
		"b8a3e0": 8683, // BenRui
		"b8a44f": 5132, // Axis
		"b8ac6f": 102,  // Dell
		"b8ad3e": 8684, // Bluecom
		"b8ae6e": 1427, // Nintendo
		"b8aeed": 1115, // Elitegroup
		"b8af67": 302,  // HP
		"b8b1c7": 8685, // BT
		"b8b2eb": 8686, // Googol
		"b8b2f8": 545,  // Apple
		"b8b3dc": 8687, // Derek
		"b8b7d7": 8688, // 2GIG
		"b8b7f1": 1592, // Wistron
		"b8b81e": 421,  // Intel
		"b8ba72": 8689, // Cynove
		"b8bb6d": 8690, // ENERES
		"b8bbaf": 152,  // Samsung
		"b8bc1b": 3335, // Huawei
		"b8bc5b": 152,  // Samsung
		"b8bd79": 8691, // TrendPoint
		"b8bebf": 2,    // Cisco
		"b8bef4": 1637, // devolo
		"b8bf83": 421,  // Intel
		"b8c111": 545,  // Apple
		"b8c227": 8692, // PSTec
		"b8c253": 826,  // Juniper
		"b8c385": 3335, // Huawei
		"b8c46f": 8693, // Primmcon
		"b8c68e": 152,  // Samsung
		"b8c6aa": 6499, // Earda
		"b8c716": 4922, // Fiberhome
		"b8c75d": 545,  // Apple
		"b8c8eb": 6282, // Itel
		"b8ca3a": 102,  // Dell
		"b8cb29": 102,  // Dell
		"b8cd93": 8694, // Penetek
		"b8cda7": 8695, // Maxeler
		"b8cef6": 432,  // Mellanox
		"b8d309": 8696, // Cox
		"b8d43e": 6370, // Vivo
		"b8d4e7": 1685, // Aruba
		"b8d50b": 3940, // Sunitec
		"b8d526": 2693, // ZyXEL
		"b8d56b": 8697, // Mirka
		"b8d6f6": 3335, // Huawei
		"b8d7af": 2057, // Murata
		"b8d94d": 2049, // Sagemcom
		"b8d9ce": 152,  // Samsung
		"b8dae8": 3335, // Huawei
		"b8dc87": 8698, // IAI
		"b8dd71": 3032, // ZTE
		"b8df6b": 8699, // SpotCam
		"b8e3b1": 3335, // Huawei
		"b8e3ee": 3858, // Universal
		"b8e589": 8700, // Payter
		"b8e60c": 545,  // Apple
		"b8e625": 1939, // 2Wire
		"b8e779": 8701, // 9Solutions
		"b8e856": 545,  // Apple
		"b8e937": 2048, // Sonos
		"b8eaaa": 600,  // ICG
		"b8eca3": 2693, // ZyXEL
		"b8ee0e": 2049, // Sagemcom
		"b8ee65": 2874, // Liteon
		"b8ee79": 8702, // YWire
		"b8f009": 6377, // Espressif
		"b8f080": 8703, // SPS
		"b8f12a": 545,  // Apple
		"b8f255": 3858, // Universal
		"b8f5e7": 8704, // WayTools
		"b8f6b1": 545,  // Apple
		"b8f732": 8705, // Aryaka
		"b8f74a": 8706, // Rcntec
		"b8f853": 2640, // Arcadyan
		"b8f883": 1595, // TP-Link
		"b8f8be": 8684, // Bluecom
		"b8f934": 101,  // Sony
		"b8ff61": 545,  // Apple
		"b8ffb3": 6429, // MitraStar
		"bc0543": 621,  // AVM
		"bc062d": 7187, // Wacom
		"bc091b": 421,  // Intel
		"bc0963": 545,  // Apple
		"bc0f64": 421,  // Intel
		"bc0f9a": 803,  // D-Link
		"bc0fa7": 8707, // Ouster
		"bc1401": 870,  // Hitron
		"bc1485": 152,  // Samsung
		"bc14ef": 8708, // ITON
		"bc1665": 2,    // Cisco
		"bc1695": 3032, // ZTE
		"bc16f5": 2,    // Cisco
		"bc17b8": 421,  // Intel
		"bc1a67": 8709, // YF
		"bc1ae4": 3335, // Huawei
		"bc1d89": 687,  // Motorola
		"bc1e85": 3335, // Huawei
		"bc20a4": 152,  // Samsung
		"bc20ba": 7723, // Inspur
		"bc2228": 803,  // D-Link
		"bc22fb": 2856, // RF
		"bc25e0": 3335, // Huawei
		"bc2643": 8710, // Elprotronic
		"bc26c7": 2,    // Cisco
		"bc282c": 8711, // e-Smart
		"bc2ce6": 2,    // Cisco
		"bc2d98": 8712, // ThinGlobal
		"bc2e48": 125,  // ARRIS
		"bc2ef6": 3335, // Huawei
		"bc2f3d": 6370, // Vivo
		"bc305b": 102,  // Dell
		"bc307d": 1592, // Wistron
		"bc307e": 1592, // Wistron
		"bc30d9": 2640, // Arcadyan
		"bc3329": 101,  // Sony
		"bc33ac": 1655, // Silicon
		"bc35e5": 8713, // Hydro
		"bc3865": 8714, // Jwcnetworks
		"bc38d2": 8715, // Pandachip
		"bc39d9": 8716, // Z-TEC
		"bc3baf": 545,  // Apple
		"bc3d85": 3335, // Huawei
		"bc3e07": 870,  // Hitron
		"bc3e13": 8717, // Accordance
		"bc3ecb": 6370, // Vivo
		"bc3f4e": 8718, // Teleepoch
		"bc3f8f": 3335, // Huawei
		"bc4100": 8719, // CoDACO
		"bc428c": 430,  // Alpsalpine
		"bc4486": 152,  // Samsung
		"bc44b0": 8720, // Elastifile
		"bc455b": 152,  // Samsung
		"bc4699": 1595, // TP-Link
		"bc4760": 152,  // Samsung
		"bc4a56": 2,    // Cisco
		"bc4b79": 8721, // SensingTek
		"bc4cc4": 545,  // Apple
		"bc4dfb": 870,  // Hitron
		"bc4e5d": 8722, // ZhongMiao
		"bc51fe": 8723, // Swann
		"bc52b4": 457,  // Nokia
		"bc52b7": 545,  // Apple
		"bc542f": 421,  // Intel
		"bc5436": 545,  // Apple
		"bc5451": 152,  // Samsung
		"bc54f9": 8724, // Drogoo
		"bc5a56": 2,    // Cisco
		"bc5bd5": 125,  // ARRIS
		"bc5c4c": 5694, // Elecom
		"bc5ea1": 8725, // PsiKick
		"bc5ff4": 4720, // ASRock
		"bc5ff6": 4254, // Mercury
		"bc60a7": 101,  // Sony
		"bc6193": 5698, // Xiaomi
		"bc620e": 3335, // Huawei
		"bc644b": 125,  // ARRIS
		"bc671c": 2,    // Cisco
		"bc6778": 545,  // Apple
		"bc6784": 8726, // Environics
		"bc69cb": 2154, // Panasonic
		"bc6a16": 8727, // tdvine
		"bc6a44": 1842, // Commend
		"bc6ad1": 5698, // Xiaomi
		"bc6b4d": 457,  // Nokia
		"bc6c21": 545,  // Apple
		"bc6d05": 8728, // Dusun
		"bc6e64": 101,  // Sony
		"bc6ee2": 421,  // Intel
		"bc71c1": 8729, // XTrillion
		"bc72b1": 152,  // Samsung
		"bc7536": 430,  // Alpsalpine
		"bc7574": 3335, // Huawei
		"bc765e": 152,  // Samsung
		"bc7670": 3335, // Huawei
		"bc76c5": 3335, // Huawei
		"bc7737": 421,  // Intel
		"bc779f": 8730, // SBM
		"bc79ad": 152,  // Samsung
		"bc7abf": 152,  // Samsung
		"bc7e8b": 152,  // Samsung
		"bc7f7b": 3335, // Huawei
		"bc7fa4": 5698, // Xiaomi
		"bc811f": 8731, // Ingate
		"bc8199": 8732, // BASIC
		"bc8385": 612,  // Microsoft
		"bc851f": 152,  // Samsung
		"bc8893": 8733, // VILLBAU
		"bc8aa3": 8734, // NHN
		"bc8ccd": 152,  // Samsung
		"bc8d0e": 457,  // Nokia
		"bc903a": 1926, // Bosch
		"bc91b5": 6296, // Infinix
		"bc926b": 545,  // Apple
		"bc9789": 3335, // Huawei
		"bc97e1": 858,  // Broadcom
		"bc9889": 4922, // Fiberhome
		"bc98df": 687,  // Motorola
		"bc9911": 2693, // ZyXEL
		"bc9930": 3335, // Huawei
		"bc99bc": 8735, // FonSee
		"bc9a53": 3335, // Huawei
		"bc9b68": 6393, // Technicolor
		"bc9c31": 3335, // Huawei
		"bc9da5": 8736, // DASCOM
		"bc9fe4": 1685, // Aruba
		"bc9fef": 545,  // Apple
		"bca13a": 8737, // SES-imagotag
		"bca4e1": 8738, // Nabto
		"bca511": 1368, // Netgear
		"bca58b": 152,  // Samsung
		"bca5a9": 545,  // Apple
		"bca8a6": 421,  // Intel
		"bca920": 545,  // Apple
		"bca993": 665,  // Cambium
		"bca9d6": 8739, // Cyber-Rain
		"bcadab": 620,  // Avaya
		"bcaec5": 1806, // ASUS
		"bcaf87": 8740, // smartAC.com
		"bcb0e7": 3335, // Huawei
		"bcb181": 3206, // Sharp
		"bcb1f3": 152,  // Samsung
		"bcb22b": 8741, // EM-Tech
		"bcb852": 8742, // Cybera
		"bcb863": 545,  // Apple
		"bcbae1": 8743, // AREC
		"bcbd9e": 6282, // Itel
		"bcc00f": 4922, // Fiberhome
		"bcc342": 2154, // Panasonic
		"bcc493": 2,    // Cisco
		"bcc6db": 457,  // Nokia
		"bccab5": 125,  // ARRIS
		"bccd45": 8744, // Voismart
		"bcce25": 1427, // Nintendo
		"bccf4f": 2693, // ZyXEL
		"bccfcc": 1341, // HTC
		"bcd074": 545,  // Apple
		"bcd11f": 152,  // Samsung
		"bcd177": 1595, // TP-Link
		"bcd206": 3335, // Huawei
		"bcd295": 2,    // Cisco
		"bcd5b6": 8745, // d2d
		"bcd713": 8746, // Owl
		"bcd767": 67,   // Private
		"bcd7a5": 1685, // Aruba
		"bcd7d4": 1916, // Roku
		"bcd940": 8747, // ASR
		"bcddc2": 6377, // Espressif
		"bcdf58": 3522, // Google
		"bce09d": 8748, // Eoslink
		"bce143": 545,  // Apple
		"bce265": 3335, // Huawei
		"bce59f": 8749, // WATERWORLD
		"bce63f": 152,  // Samsung
		"bce67c": 665,  // Cambium
		"bce712": 2,    // Cisco
		"bce92f": 302,  // HP
		"bcea2b": 2825, // CityCom
		"bceafa": 302,  // HP
		"bcec5d": 545,  // Apple
		"bceca0": 358,  // Compal
		"bcee7b": 1806, // ASUS
		"bcf171": 421,  // Intel
		"bcf1f2": 2,    // Cisco
		"bcf292": 542,  // Plantronics
		"bcf2af": 1637, // devolo
		"bcf310": 185,  // Extreme
		"bcf45f": 3032, // ZTE
		"bcf5ac": 869,  // LG
		"bcf685": 803,  // D-Link
		"bcf9f2": 8750, // Teko
		"bcfe8c": 8751, // Altronic
		"bcfed9": 545,  // Apple
		"bcff4d": 6377, // Espressif
		"bcffac": 8489, // Topcon
		"bcffeb": 687,  // Motorola
		"c00380": 826,  // Juniper
		"c005c2": 125,  // ARRIS
		"c0060c": 3335, // Huawei
		"c006c3": 1595, // TP-Link
		"c0074a": 8752, // Brita
		"c00d7e": 8753, // Additech
		"c01173": 152,  // Samsung
		"c011a6": 8754, // Fort-Telecom
		"c014b8": 457,  // Nokia
		"c014fe": 2,    // Cisco
		"c0174d": 152,  // Samsung
		"c01803": 302,  // HP
		"c01850": 3072, // Quanta
		"c01ada": 545,  // Apple
		"c01e9b": 8755, // Pixavi
		"c02250": 8756, // Koss
		"c0238d": 152,  // Samsung
		"c02506": 621,  // AVM
		"c0255c": 2,    // Cisco
		"c02567": 8757, // Nexxt
		"c025a5": 102,  // Dell
		"c025e9": 1595, // TP-Link
		"c0288d": 4080, // Logitech
		"c02973": 8758, // Audyssey
		"c029f3": 8759, // XySystem
		"c02b31": 8272, // Phytium
		"c02c5c": 545,  // Apple
		"c02dee": 8760, // Cuff
		"c02e26": 67,   // Private
		"c02ff1": 8761, // Volta
		"c0335e": 612,  // Microsoft
		"c034b4": 8762, // Gigastone
		"c03580": 8763, // A&R
		"c035c5": 1954, // Prosoft
		"c03653": 5821, // eero
		"c03656": 4922, // Fiberhome
		"c038f9": 457,  // Nokia
		"c03c04": 2049, // Sagemcom
		"c03c59": 421,  // Intel
		"c03d03": 152,  // Samsung
		"c03dd9": 6429, // MitraStar
		"c03e0f": 3513, // BSkyB
		"c03eba": 102,  // Dell
		"c03f0e": 1368, // Netgear
		"c03f2a": 8764, // Biscotti
		"c03fd5": 1115, // Elitegroup
		"c03fdd": 3335, // Huawei
		"c04004": 8765, // Medicaroid
		"c041f6": 869,  // LG
		"c042d0": 826,  // Juniper
		"c04301": 8766, // Epec
		"c04442": 545,  // Apple
		"c04754": 6370, // Vivo
		"c048e6": 152,  // Samsung
		"c049ef": 6377, // Espressif
		"c04a00": 1595, // TP-Link
		"c04b13": 8767, // WonderSound
		"c04df7": 8768, // Serelec
		"c05627": 2469, // Belkin
		"c057bc": 620,  // Avaya
		"c058a7": 456,  // Pico
		"c06118": 1595, // TP-Link
		"c0626b": 2,    // Cisco
		"c06369": 8769, // Binxin
		"c06394": 545,  // Apple
		"c064c6": 457,  // Nokia
		"c064e4": 2,    // Cisco
		"c06599": 152,  // Samsung
		"c067af": 2,    // Cisco
		"c06b55": 687,  // Motorola
		"c06c6d": 8770, // MagneMotion
		"c07009": 3335, // Huawei
		"c074ad": 1683, // Grandstream
		"c07831": 3335, // Huawei
		"c07878": 833,  // Flextronics
		"c07bbc": 2,    // Cisco
		"c07cd1": 5440, // Pegatron
		"c0830a": 1939, // 2Wire
		"c083c9": 3335, // Huawei
		"c0847a": 545,  // Apple
		"c0847d": 4499, // AMPAK
		"c08488": 8771, // Finis
		"c0854c": 6840, // Ragentek
		"c087eb": 152,  // Samsung
		"c0885b": 8772, // SnD
		"c0886d": 8773, // Securosys
		"c08997": 152,  // Samsung
		"c089ab": 125,  // ARRIS
		"c08ade": 2738, // Ruckus
		"c08b05": 3335, // Huawei
		"c08c60": 2,    // Cisco
		"c08c71": 687,  // Motorola
		"c09296": 3032, // ZTE
		"c09435": 125,  // ARRIS
		"c094ad": 3032, // ZTE
		"c09727": 152,  // Samsung
		"c09879": 143,  // Acer
		"c09ad0": 545,  // Apple
		"c09c92": 8774, // Coby
		"c09f42": 545,  // Apple
		"c09fe1": 3032, // ZTE
		"c0a00d": 125,  // ARRIS
		"c0a0bb": 803,  // D-Link
		"c0a0c7": 8775, // Fairfield
		"c0a1a2": 8776, // MarqMetrix
		"c0a36e": 3513, // BSkyB
		"c0a39e": 8777, // EarthCam
		"c0a53e": 545,  // Apple
		"c0a600": 545,  // Apple
		"c0a8f0": 8778, // Adamson
		"c0ac54": 2049, // Sagemcom
		"c0b101": 3032, // ZTE
		"c0b339": 8779, // Comigo
		"c0b357": 5918, // Yoshiki
		"c0b47d": 3335, // Huawei
		"c0b5cd": 3335, // Huawei
		"c0b658": 545,  // Apple
		"c0b6f9": 421,  // Intel
		"c0b883": 421,  // Intel
		"c0b8b1": 8780, // BitBox
		"c0b8e6": 5441, // Ruijie
		"c0bc9a": 3335, // Huawei
		"c0bdc8": 152,  // Samsung
		"c0bdd1": 152,  // Samsung
		"c0bfa7": 826,  // Juniper
		"c0bfc0": 3335, // Huawei
		"c0c1c0": 1783, // Linksys
		"c0c3b6": 8781, // Automatic
		"c0c520": 2738, // Ruckus
		"c0c522": 125,  // ARRIS
		"c0c946": 8782, // Mitsuya
		"c0c9e3": 1595, // TP-Link
		"c0ccf8": 545,  // Apple
		"c0cecd": 545,  // Apple
		"c0cfa3": 357,  // Creative
		"c0d012": 545,  // Apple
		"c0d026": 3335, // Huawei
		"c0d044": 2049, // Sagemcom
		"c0d193": 3335, // Huawei
		"c0d2dd": 152,  // Samsung
		"c0d3c0": 152,  // Samsung
		"c0d46b": 3335, // Huawei
		"c0d682": 3794, // Arista
		"c0d7aa": 2640, // Arcadyan
		"c0d834": 8783, // xvtec
		"c0d962": 2545, // Askey
		"c0dcd7": 3335, // Huawei
		"c0dcda": 152,  // Samsung
		"c0df77": 281,  // Conrad
		"c0e018": 3335, // Huawei
		"c0e1be": 3335, // Huawei
		"c0e3a0": 5696, // Renesas
		"c0e3fb": 3335, // Huawei
		"c0e42d": 1595, // TP-Link
		"c0e434": 3005, // Azurewave
		"c0e862": 545,  // Apple
		"c0eae4": 1003, // Sonicwall
		"c0ee40": 4537, // Laird
		"c0eefb": 6924, // OnePlus
		"c0f1c4": 8784, // Pacidal
		"c0f2fb": 545,  // Apple
		"c0f4e6": 3335, // Huawei
		"c0f535": 4499, // AMPAK
		"c0f6c2": 3335, // Huawei
		"c0f6ec": 3335, // Huawei
		"c0f79d": 8785, // Powercode
		"c0f827": 8786, // Rapidmax
		"c0f87f": 2,    // Cisco
		"c0f945": 35,   // Toshiba
		"c0f9b0": 3335, // Huawei
		"c0f9d2": 8787, // arkona
		"c0fbc1": 6282, // Itel
		"c0fd84": 3032, // ZTE
		"c0ffa8": 3335, // Huawei
		"c0ffd4": 1368, // Netgear
		"c40049": 8788, // Kamama
		"c400ad": 1703, // Advantech
		"c40142": 8789, // MaxMedia
		"c4017c": 2738, // Ruckus
		"c401b1": 8790, // SeekTech
		"c401ce": 8791, // Presition
		"c402e1": 8792, // Khwahish
		"c403a8": 421,  // Intel
		"c40415": 1368, // Netgear
		"c40528": 3335, // Huawei
		"c40683": 3335, // Huawei
		"c4072f": 3335, // Huawei
		"c4084a": 457,  // Nokia
		"c409b7": 826,  // Juniper
		"c40acb": 2,    // Cisco
		"c40b31": 545,  // Apple
		"c40bcb": 5698, // Xiaomi
		"c40d96": 3335, // Huawei
		"c40e45": 8793, // ACK
		"c40f09": 5591, // Hermes
		"c4108a": 2738, // Ruckus
		"c41234": 545,  // Apple
		"c412f5": 803,  // D-Link
		"c413e2": 185,  // Extreme
		"c41411": 545,  // Apple
		"c4143c": 2,    // Cisco
		"c41688": 3335, // Huawei
		"c416c8": 3335, // Huawei
		"c416fa": 8794, // Prysm
		"c4170e": 3335, // Huawei
		"c418e9": 152,  // Samsung
		"c419ec": 8795, // Qualisys
		"c41c07": 152,  // Samsung
		"c41c9c": 8796, // JiQiDao
		"c41cff": 3468, // Vizio
		"c421c8": 2849, // Kyocera
		"c42360": 421,  // Intel
		"c4237a": 8797, // WhizNets
		"c42728": 3032, // ZTE
		"c4278c": 3335, // Huawei
		"c42795": 6393, // Technicolor
		"c42ad0": 545,  // Apple
		"c42b44": 3335, // Huawei
		"c42c03": 545,  // Apple
		"c432d1": 8798, // Farlink
		"c4345b": 3335, // Huawei
		"c4346b": 302,  // HP
		"c4366c": 869,  // LG
		"c436c0": 1077, // Buffalo
		"c436da": 8799, // Rusteletech
		"c43772": 8800, // Virtuozzo
		"c43875": 2048, // Sonos
		"c438d3": 8801, // Tagatec
		"c4393a": 741,  // SMC
		"c43a35": 6441, // FN-Link
		"c43a9f": 8802, // Siconix
		"c43abe": 101,  // Sony
		"c43c3c": 8803, // Cybelec
		"c43cea": 1077, // Buffalo
		"c43dc7": 1368, // Netgear
		"c44044": 8804, // RackTop
		"c4411e": 2469, // Belkin
		"c44202": 152,  // Samsung
		"c44268": 2353, // Crestron
		"c4438f": 869,  // LG
		"c4447d": 3335, // Huawei
		"c444a0": 2,    // Cisco
		"c4473f": 3335, // Huawei
		"c44ad0": 8805, // Fireflies
		"c44b44": 8806, // Omniprint
		"c44d84": 2,    // Cisco
		"c44e1f": 8807, // BlueN
		"c44f33": 6377, // Espressif
		"c45006": 152,  // Samsung
		"c45444": 3072, // Quanta
		"c455a6": 8808, // Cadac
		"c455c2": 8809, // Bach-Simpson
		"c456fe": 647,  // Lava
		"c4576e": 152,  // Samsung
		"c45a86": 3335, // Huawei
		"c45bbe": 6377, // Espressif
		"c45bf7": 8810, // ants
		"c45d83": 152,  // Samsung
		"c45e5c": 3335, // Huawei
		"c46044": 929,  // Everex
		"c4618b": 545,  // Apple
		"c462ea": 152,  // Samsung
		"c46354": 8811, // U-Raku
		"c463fb": 8812, // Neatframe
		"c46413": 2,    // Cisco
		"c464b7": 4922, // Fiberhome
		"c46516": 302,  // HP
		"c46699": 6370, // Vivo
		"c467b5": 8813, // Libratone
		"c467d1": 3335, // Huawei
		"c469f0": 3335, // Huawei
		"c46ab7": 5698, // Xiaomi
		"c46bb4": 8814, // myIDkey
		"c46df1": 8815, // DataGravity
		"c46e1f": 1595, // TP-Link
		"c470ab": 5441, // Ruijie
		"c47154": 1595, // TP-Link
		"c471fe": 2,    // Cisco
		"c47295": 2,    // Cisco
		"c4731e": 152,  // Samsung
		"c4741e": 3032, // ZTE
		"c47469": 8816, // BT9
		"c475ab": 421,  // Intel
		"c478a2": 3335, // Huawei
		"c47ba3": 8817, // NAVIS
		"c47d46": 4,    // Fujitsu
		"c47d4f": 2,    // Cisco
		"c47d9f": 152,  // Samsung
		"c47dcc": 768,  // Zebra
		"c47dfe": 8818, // A.N
		"c47f51": 8819, // Inventek
		"c48025": 3335, // Huawei
		"c4836f": 4565, // Ciena
		"c48466": 545,  // Apple
		"c48508": 421,  // Intel
		"c486e9": 3335, // Huawei
		"c488e5": 152,  // Samsung
		"c48a5a": 8820, // Jfcontrol
		"c4910c": 545,  // Apple
		"c491cf": 8408, // Luxul
		"c49300": 8821, // 8Devices
		"c49313": 8822, // 100fio
		"c49380": 8823, // Speedytel
		"c493d9": 152,  // Samsung
		"c49500": 5439, // Amazon
		"c49805": 8824, // Minieum
		"c49880": 545,  // Apple
		"c49886": 6653, // Qorvo
		"c49a02": 869,  // LG
		"c49ded": 612,  // Microsoft
		"c49f4c": 3335, // Huawei
		"c49ff3": 8825, // Mciao
		"c4a366": 3032, // ZTE
		"c4a402": 3335, // Huawei
		"c4a81d": 803,  // D-Link
		"c4abb2": 6370, // Vivo
		"c4ac59": 2057, // Murata
		"c4ad21": 8826, // MEDIAEDGE
		"c4ad34": 1784, // Routerboard.com
		"c4adf1": 8827, // GOPEACE
		"c4ae12": 152,  // Samsung
		"c4b239": 2,    // Cisco
		"c4b301": 545,  // Apple
		"c4b36a": 2,    // Cisco
		"c4b8b4": 3335, // Huawei
		"c4b9cd": 2,    // Cisco
		"c4bd6a": 268,  // SKF
		"c4bde5": 421,  // Intel
		"c4bed4": 620,  // Avaya
		"c4bf60": 6266, // Tecno
		"c4c0ae": 8828, // Midori
		"c4c138": 8829, // OWLink
		"c4c36b": 545,  // Apple
		"c4c563": 6266, // Tecno
		"c4c603": 2,    // Cisco
		"c4ca2b": 3794, // Arista
		"c4d0e3": 421,  // Intel
		"c4d438": 3335, // Huawei
		"c4d655": 8830, // Tercel
		"c4d738": 3335, // Huawei
		"c4d8f3": 8831, // iZotope
		"c4d987": 421,  // Intel
		"c4da26": 8832, // Noblex
		"c4dd57": 6377, // Espressif
		"c4de7b": 3335, // Huawei
		"c4e17c": 8833, // U2S
		"c4e287": 3335, // Huawei
		"c4e506": 8834, // Piper
		"c4e510": 8835, // Mechatro
		"c4e532": 2640, // Arcadyan
		"c4e7be": 8836, // SCSpro
		"c4e90a": 803,  // D-Link
		"c4e984": 1595, // TP-Link
		"c4ea1d": 6393, // Technicolor
		"c4eef5": 8837, // II-VI
		"c4f081": 3335, // Huawei
		"c4f0ec": 4922, // Fiberhome
		"c4f122": 8838, // Nexar
		"c4f174": 5821, // eero
		"c4f464": 8839, // Spica
		"c4f57c": 90,   // Brocade
		"c4f5a5": 8840, // Kumalift
		"c4f7d5": 2,    // Cisco
		"c4fbaa": 3335, // Huawei
		"c4fcef": 8841, // SambaNova
		"c4fde6": 8842, // Drtech
		"c4fee2": 6623, // AMICCOM
		"c4ff1f": 3335, // Huawei
		"c4ff22": 3335, // Huawei
		"c80084": 2,    // Cisco
		"c80210": 869,  // LG
		"c8028f": 8843, // Nova
		"c803f5": 2738, // Ruckus
		"c80718": 8844, // TDSi
		"c80739": 6083, // NAKAYO
		"c80873": 2738, // Ruckus
		"c808e9": 869,  // LG
		"c809a8": 421,  // Intel
		"c80aa9": 3072, // Quanta
		"c80cc8": 3335, // Huawei
		"c80d32": 8845, // Holoplot
		"c80e95": 8846, // OmniLync
		"c81451": 3335, // Huawei
		"c81479": 152,  // Samsung
		"c816a5": 6363, // Masimo
		"c81739": 6282, // Itel
		"c819f7": 152,  // Samsung
		"c81afe": 8847, // DLOGIC
		"c81b5c": 8848, // BCTech
		"c81cfe": 768,  // Zebra
		"c81ee7": 545,  // Apple
		"c81f66": 102,  // Dell
		"c81fbe": 3335, // Huawei
		"c81fea": 620,  // Avaya
		"c8208e": 8849, // Storagedata
		"c82158": 421,  // Intel
		"c8292a": 8850, // Barun
		"c82a14": 545,  // Apple
		"c82b96": 6377, // Espressif
		"c82e94": 8851, // Halfa
		"c83168": 8852, // eZEX
		"c8334b": 545,  // Apple
		"c833e5": 3335, // Huawei
		"c8348e": 421,  // Intel
		"c835b8": 306,  // Ericsson
		"c83870": 152,  // Samsung
		"c839ac": 3335, // Huawei
		"c83a35": 6267, // Tenda
		"c83a6b": 1916, // Roku
		"c83b45": 8853, // JRI
		"c83c85": 545,  // Apple
		"c83d97": 457,  // Nokia
		"c83dd4": 189,  // CyberTAN
		"c83ddc": 5698, // Xiaomi
		"c83dfc": 8854, // AlphaTheta
		"c83e9e": 3335, // Huawei
		"c83ea7": 8855, // KUNBUS
		"c83f26": 612,  // Microsoft
		"c83fb4": 125,  // ARRIS
		"c84029": 4922, // Fiberhome
		"c8418a": 152,  // Samsung
		"c84529": 8856, // IMK
		"c8458f": 8857, // Wyler
		"c84782": 8858, // Areson
		"c8478c": 7783, // Beken
		"c84c75": 2,    // Cisco
		"c84f86": 3579, // Sophos
		"c850ce": 3335, // Huawei
		"c850e9": 2051, // Raisecom
		"c85142": 152,  // Samsung
		"c85195": 3335, // Huawei
		"c85261": 125,  // ARRIS
		"c8544b": 2693, // ZyXEL
		"c854a4": 6296, // Infinix
		"c85663": 8859, // Sunflex
		"c85895": 687,  // Motorola
		"c858c0": 421,  // Intel
		"c85a9f": 3032, // ZTE
		"c85acf": 302,  // HP
		"c85b76": 4920, // LCFC
		"c85d38": 531,  // HUMAX
		"c86000": 1806, // ASUS
		"c863f1": 101,  // Sony
		"c863fc": 125,  // ARRIS
		"c864c7": 3032, // ZTE
		"c8665d": 185,  // Extreme
		"c8675e": 185,  // Extreme
		"c868de": 3335, // Huawei
		"c869cd": 545,  // Apple
		"c86c1e": 8860, // Display
		"c86c3d": 5439, // Amazon
		"c86c87": 2693, // ZyXEL
		"c86cb6": 8861, // Optcom
		"c86f1d": 545,  // Apple
		"c87248": 8862, // Aplicom
		"c8727e": 457,  // Nokia
		"c8755b": 8863, // Quantify
		"c87765": 1925, // Tiesse
		"c87b23": 1819, // Bose
		"c87b5b": 3032, // ZTE
		"c87cbc": 8864, // Valink
		"c87e75": 152,  // Samsung
		"c88314": 8865, // Tempo
		"c88439": 8866, // Sunrise
		"c88447": 8867, // Beautiful
		"c884a1": 2,    // Cisco
		"c884cf": 3335, // Huawei
		"c88550": 545,  // Apple
		"c88722": 8868, // Lumenpulse
		"c889f3": 545,  // Apple
		"c88be8": 6363, // Masimo
		"c88d83": 3335, // Huawei
		"c8903e": 8869, // Pakton
		"c891f9": 2049, // Sagemcom
		"c89346": 8870, // MXCHIP
		"c894bb": 3335, // Huawei
		"c89665": 612,  // Microsoft
		"c8979f": 457,  // Nokia
		"c899b2": 2640, // Arcadyan
		"c89c13": 8871, // Inspiremobile
		"c89c1d": 2,    // Cisco
		"c89cdc": 1115, // Elitegroup
		"c89d18": 3335, // Huawei
		"c89e43": 1368, // Netgear
		"c89f1a": 3335, // Huawei
		"c89f1d": 8626, // Shenzhen
		"c8a1ba": 8872, // Neul
		"c8a620": 8873, // Nebula
		"c8a729": 8874, // Systronics
		"c8a776": 3335, // Huawei
		"c8a823": 152,  // Samsung
		"c8a9fc": 8875, // Goyoo
		"c8aa21": 125,  // ARRIS
		"c8aacc": 67,   // Private
		"c8b1cd": 545,  // Apple
		"c8b1ee": 6653, // Qorvo
		"c8b21e": 6440, // Chipsea
		"c8b29b": 421,  // Intel
		"c8b373": 1783, // Linksys
		"c8b422": 2545, // Askey
		"c8b5ad": 302,  // HP
		"c8b5b7": 545,  // Apple
		"c8b6d3": 3335, // Huawei
		"c8b82f": 5821, // eero
		"c8ba94": 152,  // Samsung
		"c8bae9": 8876, // Qdis
		"c8bb81": 3335, // Huawei
		"c8bbd3": 8877, // Embrane
		"c8bc9c": 3335, // Huawei
		"c8bcc8": 545,  // Apple
		"c8bd4d": 152,  // Samsung
		"c8bd69": 152,  // Samsung
		"c8be19": 803,  // D-Link
		"c8be35": 185,  // Extreme
		"c8bffe": 3335, // Huawei
		"c8c2f5": 833,  // Flextronics
		"c8c2fa": 3335, // Huawei
		"c8c465": 3335, // Huawei
		"c8c64a": 833,  // Flextronics
		"c8c750": 687,  // Motorola
		"c8c791": 8878, // Zero1.tv
		"c8c9a3": 6377, // Espressif
		"c8ca63": 3335, // Huawei
		"c8ca79": 4565, // Ciena
		"c8cb9e": 421,  // Intel
		"c8cbb8": 302,  // HP
		"c8cd72": 2049, // Sagemcom
		"c8d083": 545,  // Apple
		"c8d10b": 457,  // Nokia
		"c8d12a": 3870, // Comtrend
		"c8d15e": 3335, // Huawei
		"c8d1d1": 8879, // AGAiT
		"c8d2c1": 8880, // Jetlun
		"c8d3a3": 803,  // D-Link
		"c8d3ff": 302,  // HP
		"c8d429": 8881, // Muehlbauer
		"c8d719": 1783, // Linksys
		"c8d7b0": 152,  // Samsung
		"c8d884": 3858, // Universal
		"c8d9d2": 302,  // HP
		"c8db26": 4080, // Logitech
		"c8ddc9": 2665, // Lenovo
		"c8de51": 8882, // IntegraOptics
		"c8dec9": 6681, // Coriant
		"c8df7c": 457,  // Nokia
		"c8e0eb": 545,  // Apple
		"c8e1a7": 6956, // Vertu
		"c8e265": 421,  // Intel
		"c8e600": 3335, // Huawei
		"c8e776": 8883, // PTCOM
		"c8e7d8": 4254, // Mercury
		"c8e7f0": 826,  // Juniper
		"c8eaf8": 3032, // ZTE
		"c8ee08": 8884, // Tangtop
		"c8ee75": 8885, // Pishion
		"c8f319": 869,  // LG
		"c8f406": 620,  // Avaya
		"c8f650": 545,  // Apple
		"c8f68d": 8886, // S.E.Technologies
		"c8f6c8": 4922, // Fiberhome
		"c8f733": 421,  // Intel
		"c8f750": 102,  // Dell
		"c8f946": 8887, // LOCOSYS
		"c8f9c8": 8888, // NewSharp
		"c8f9f9": 2,    // Cisco
		"c8fa84": 8889, // Trusonus
		"c8fe6a": 826,  // Juniper
		"c8ff28": 2874, // Liteon
		"c8ff77": 7190, // Dyson
		"cc03fa": 6393, // Technicolor
		"cc051b": 152,  // Samsung
		"cc0577": 3335, // Huawei
		"cc0677": 4922, // Fiberhome
		"cc07ab": 152,  // Samsung
		"cc07e4": 2665, // Lenovo
		"cc088d": 545,  // Apple
		"cc08e0": 545,  // Apple
		"cc08fb": 1595, // TP-Link
		"cc09c8": 8890, // Imaqliq
		"cc0cda": 8891, // Miljovakt
		"cc0df2": 687,  // Motorola
		"cc1531": 421,  // Intel
		"cc167e": 2,    // Cisco
		"cc187b": 8892, // Manzanita
		"cc1afa": 3032, // ZTE
		"cc1fc4": 8893, // InVue
		"cc208c": 3335, // Huawei
		"cc20e8": 545,  // Apple
		"cc2119": 152,  // Samsung
		"cc2218": 8894, // InnoDigital
		"cc25ef": 545,  // Apple
		"cc262d": 8895, // Verifi
		"cc29f5": 545,  // Apple
		"cc2d1b": 3188, // SFR
		"cc2d21": 6267, // Tenda
		"cc2d8c": 869,  // LG
		"cc2db7": 545,  // Apple
		"cc2de0": 1784, // Routerboard.com
		"cc2f71": 421,  // Intel
		"cc3080": 8896, // VAIO
		"cc3296": 3335, // Huawei
		"cc32e5": 1595, // TP-Link
		"cc33bb": 2049, // Sagemcom
		"cc3429": 1595, // TP-Link
		"cc3540": 6393, // Technicolor
		"cc355a": 8897, // SecuGen
		"cc37ab": 6295, // Edgecore
		"cc398c": 8898, // Shiningtek
		"cc3a61": 152,  // Samsung
		"cc3b27": 6266, // Tecno
		"cc3b58": 8899, // Curiouser
		"cc3d82": 421,  // Intel
		"cc3e5f": 302,  // HP
		"cc3f8a": 5903, // Komatsu
		"cc3fea": 1743, // BAE
		"cc40d0": 1368, // Netgear
		"cc4463": 545,  // Apple
		"cc4639": 8900, // WAAV
		"cc464e": 152,  // Samsung
		"cc46d6": 2,    // Cisco
		"cc4703": 8901, // Intercon
		"cc4792": 6723, // ASIX
		"cc47bd": 8902, // Rhombus
		"cc483a": 102,  // Dell
		"cc4b73": 4499, // AMPAK
		"cc4d38": 8903, // Carnegie
		"cc4e24": 90,   // Brocade
		"cc4eec": 531,  // HUMAX
		"cc500a": 4922, // Fiberhome
		"cc501c": 8904, // KVH
		"cc5076": 8905, // Ocom
		"cc50e3": 6377, // Espressif
		"cc53b5": 3335, // Huawei
		"cc5459": 8906, // OnTime
		"cc55ad": 4556, // RIM
		"cc5a53": 2,    // Cisco
		"cc5b31": 1427, // Nintendo
		"cc5c61": 3335, // Huawei
		"cc5d4e": 2693, // ZyXEL
		"cc5d78": 8907, // JTD
		"cc60c8": 612,  // Microsoft
		"cc61e5": 687,  // Motorola
		"cc64a6": 3335, // Huawei
		"cc65ad": 125,  // ARRIS
		"cc660a": 545,  // Apple
		"cc66b2": 457,  // Nokia
		"cc68b6": 1595, // TP-Link
		"cc69fa": 545,  // Apple
		"cc6da0": 1916, // Roku
		"cc6ea4": 152,  // Samsung
		"cc70ed": 2,    // Cisco
		"cc720f": 8908, // Viscount
		"cc7498": 8909, // Filmetrics
		"cc75e2": 125,  // ARRIS
		"cc7669": 8910, // Seetech
		"cc785f": 545,  // Apple
		"cc794a": 7461, // BLU
		"cc7b35": 3032, // ZTE
		"cc7b61": 8911, // Nikkiso
		"cc7d37": 125,  // ARRIS
		"cc7ee7": 2154, // Panasonic
		"cc7f75": 2,    // Cisco
		"cc7f76": 2,    // Cisco
		"cc812a": 6370, // Vivo
		"cc81da": 6849, // Phicomm
		"cc827f": 1703, // Advantech
		"cc82eb": 2849, // Kyocera
		"cc86ec": 1655, // Silicon
		"cc874a": 457,  // Nokia
		"cc8826": 869,  // LG
		"cc88c7": 1685, // Aruba
		"cc895e": 3335, // Huawei
		"cc89fd": 457,  // Nokia
		"cc8e71": 2,    // Cisco
		"cc9070": 2,    // Cisco
		"cc9470": 8912, // Kinestral
		"cc95d7": 3468, // Vizio
		"cc9635": 8913, // LVS
		"cc96a0": 3335, // Huawei
		"cc9891": 2,    // Cisco
		"cc9da2": 8497, // Eltex
		"cc9e00": 1427, // Nintendo
		"cc9ea2": 5439, // Amazon
		"cca223": 3335, // Huawei
		"cca3bd": 6282, // Itel
		"cca462": 125,  // ARRIS
		"cca614": 8914, // Aifa
		"cca7c1": 3522, // Google
		"ccab2c": 531,  // HUMAX
		"ccb0a8": 3335, // Huawei
		"ccb0da": 2874, // Liteon
		"ccb11a": 152,  // Samsung
		"ccb182": 3335, // Huawei
		"ccb255": 803,  // D-Link
		"ccb691": 8915, // NECMagnusCommunications
		"ccb8a8": 4499, // AMPAK
		"ccbbfe": 3335, // Huawei
		"ccbce3": 3335, // Huawei
		"ccbd35": 8916, // Steinel
		"ccbe59": 374,  // Calix
		"ccbe71": 8917, // OptiLogix
		"ccc079": 2057, // Murata
		"ccc2e0": 2051, // Raisecom
		"ccc3ea": 687,  // Motorola
		"ccc5e5": 102,  // Dell
		"ccc62b": 8918, // Tri-Systems
		"ccc760": 545,  // Apple
		"ccc95d": 545,  // Apple
		"cccc81": 3335, // Huawei
		"cccccc": 1655, // Silicon
		"cccd64": 8919, // SM-Electronic
		"ccce40": 8920, // Janteq
		"ccd083": 1685, // Aruba
		"ccd281": 545,  // Apple
		"ccd3c1": 1449, // Vestel
		"ccd42e": 2640, // Arcadyan
		"ccd4a1": 6429, // MitraStar
		"ccd539": 2,    // Cisco
		"ccd73c": 3335, // Huawei
		"ccd811": 8921, // Aiconn
		"ccd81f": 8922, // Maipu
		"ccd8c1": 2,    // Cisco
		"ccd9ac": 421,  // Intel
		"ccd9e9": 8923, // SCR
		"ccdb04": 8924, // DataRemote
		"ccdb93": 2,    // Cisco
		"ccdc55": 8925, // Dragonchip
		"cce0c3": 8926, // EXTEN
		"cce17f": 826,  // Juniper
		"cce194": 826,  // Juniper
		"cce1d5": 1077, // Buffalo
		"cce8ac": 8927, // SOYEA
		"ccea1c": 8928, // DCONWORKS
		"cced4d": 2,    // Cisco
		"cceddc": 6429, // MitraStar
		"ccef48": 2,    // Cisco
		"ccf407": 8929, // Eukrea
		"ccf411": 3522, // Google
		"ccf538": 8930, // 3isysnetworks
		"ccf67a": 8931, // Ayecka
		"ccf735": 5439, // Amazon
		"ccf841": 8932, // Lumewave
		"ccf954": 620,  // Avaya
		"ccf957": 7420, // u-blox
		"ccf9e4": 421,  // Intel
		"ccf9e8": 152,  // Samsung
		"ccfa00": 869,  // LG
		"ccfa66": 3335, // Huawei
		"ccfb65": 1427, // Nintendo
		"ccfcb1": 231,  // Wireless
		"ccfd17": 6434, // TCT
		"ccfe3c": 152,  // Samsung
		"ccff90": 3335, // Huawei
		"d0034b": 545,  // Apple
		"d003df": 152,  // Samsung
		"d00401": 687,  // Motorola
		"d00492": 4922, // Fiberhome
		"d004b0": 152,  // Samsung
		"d0052a": 2640, // Arcadyan
		"d005e4": 3335, // Huawei
		"d007ca": 826,  // Juniper
		"d00df7": 3335, // Huawei
		"d00f6d": 4049, // T&W
		"d01242": 8933, // BIOS
		"d0131e": 8934, // Sunrex
		"d013fd": 869,  // LG
		"d0154a": 3032, // ZTE
		"d015a6": 1685, // Aruba
		"d0167c": 5821, // eero
		"d016b4": 3335, // Huawei
		"d01769": 2057, // Murata
		"d0176a": 152,  // Samsung
		"d017c2": 1806, // ASUS
		"d0196a": 4565, // Ciena
		"d01aa7": 8935, // UniPrint
		"d01b1f": 6578, // Ohsung
		"d01b49": 152,  // Samsung
		"d01c3c": 6266, // Tecno
		"d021ac": 8936, // Yo
		"d021f9": 2979, // Ubiquiti
		"d022be": 152,  // Samsung
		"d023db": 545,  // Apple
		"d02516": 4254, // Mercury
		"d02544": 152,  // Samsung
		"d02598": 545,  // Apple
		"d02b20": 545,  // Apple
		"d02c45": 8937, // littleBits
		"d02db3": 3335, // Huawei
		"d03169": 152,  // Samsung
		"d03311": 545,  // Apple
		"d03742": 3097, // Yulong
		"d03745": 1595, // TP-Link
		"d039b3": 125,  // ARRIS
		"d039ea": 5559, // NetApp
		"d03c1f": 421,  // Intel
		"d03dc3": 8938, // AQ
		"d03e5c": 3335, // Huawei
		"d03e7d": 6440, // Chipsea
		"d03f27": 6882, // Wyze
		"d03faa": 545,  // Apple
		"d040ef": 2057, // Murata
		"d041c9": 4922, // Fiberhome
		"d0431e": 102,  // Dell
		"d047c1": 8268, // Elma
		"d048f3": 8939, // DATTUS
		"d0497c": 6924, // OnePlus
		"d04cc1": 8940, // SINTRONES
		"d04d2c": 1916, // Roku
		"d04dc6": 1685, // Aruba
		"d04f7e": 545,  // Apple
		"d05099": 4720, // ASRock
		"d05162": 101,  // Sony
		"d05349": 2874, // Liteon
		"d0542d": 1640, // Cambridge
		"d05509": 1427, // Nintendo
		"d056bf": 8941, // Amosense
		"d0574c": 2,    // Cisco
		"d0577b": 421,  // Intel
		"d05785": 2278, // Pantech
		"d05794": 2049, // Sagemcom
		"d058a8": 3032, // ZTE
		"d058fc": 3513, // BSkyB
		"d05919": 3032, // ZTE
		"d05995": 4922, // Fiberhome
		"d059c3": 8942, // CeraMicro
		"d059e4": 152,  // Samsung
		"d05a00": 6393, // Technicolor
		"d05ba8": 3032, // ZTE
		"d0608c": 3032, // ZTE
		"d063b4": 8943, // SolidRun
		"d06544": 545,  // Apple
		"d065ca": 3335, // Huawei
		"d0667b": 152,  // Samsung
		"d06726": 302,  // HP
		"d067e5": 102,  // Dell
		"d06a1f": 8944, // BSE
		"d06dc9": 2049, // Sagemcom
		"d06ede": 2049, // Sagemcom
		"d06f4a": 8945, // Topwell
		"d06f82": 3335, // Huawei
		"d071c4": 3032, // ZTE
		"d072dc": 2,    // Cisco
		"d0737f": 8946, // Mini-Circuits
		"d0768f": 374,  // Calix
		"d076e7": 1595, // TP-Link
		"d07714": 687,  // Motorola
		"d07880": 4922, // Fiberhome
		"d07ab5": 3335, // Huawei
		"d07d33": 3335, // Huawei
		"d07e28": 302,  // HP
		"d07e35": 421,  // Intel
		"d07fa0": 152,  // Samsung
		"d0817a": 545,  // Apple
		"d084b0": 2049, // Sagemcom
		"d087e2": 152,  // Samsung
		"d0880c": 545,  // Apple
		"d08999": 8947, // APCON
		"d08a55": 7040, // Skullcandy
		"d08a91": 6393, // Technicolor
		"d08cff": 8948, // Upwis
		"d08e79": 102,  // Dell
		"d0929e": 612,  // Microsoft
		"d092fa": 4922, // Fiberhome
		"d09380": 8949, // Ducere
		"d09466": 102,  // Dell
		"d095c7": 2278, // Pantech
		"d09b05": 8950, // Emtronix
		"d09c7a": 5698, // Xiaomi
		"d09cae": 6370, // Vivo
		"d09d0a": 8951, // Linkcom
		"d09dab": 6434, // TCT
		"d0a4b1": 8952, // Sonifex
		"d0a5a6": 2,    // Cisco
		"d0a637": 545,  // Apple
		"d0abd5": 421,  // Intel
		"d0aeec": 2237, // Alpha
		"d0afb6": 8953, // Linktop
		"d0b0cd": 8954, // Moen
		"d0b128": 152,  // Samsung
		"d0b214": 8955, // PoeWit
		"d0b2c4": 6393, // Technicolor
		"d0b45d": 3335, // Huawei
		"d0b60a": 8956, // Xingluo
		"d0b66f": 6256, // Sernet
		"d0bb61": 3032, // ZTE
		"d0bd01": 4699, // DS
		"d0be2c": 8957, // CNSLink
		"d0bf9c": 302,  // HP
		"d0c193": 8958, // Skybell
		"d0c1b1": 152,  // Samsung
		"d0c24e": 152,  // Samsung
		"d0c282": 2,    // Cisco
		"d0c31e": 8959, // JUNGJIN
		"d0c5d3": 3005, // Azurewave
		"d0c5d8": 8960, // Latecoere
		"d0c5f3": 545,  // Apple
		"d0c637": 421,  // Intel
		"d0c65b": 3335, // Huawei
		"d0c789": 2,    // Cisco
		"d0c7c0": 1595, // TP-Link
		"d0cde1": 8961, // Scientech
		"d0d003": 152,  // Samsung
		"d0d04b": 3335, // Huawei
		"d0d0fd": 2,    // Cisco
		"d0d212": 8962, // K2NET
		"d0d23c": 545,  // Apple
		"d0d2b0": 545,  // Apple
		"d0d3e0": 1685, // Aruba
		"d0d3fc": 8963, // Mios
		"d0d471": 8964, // MVTECH
		"d0d6cc": 8965, // Wintop
		"d0d783": 3335, // Huawei
		"d0db32": 457,  // Nokia
		"d0dbb7": 3171, // Casa
		"d0dd49": 826,  // Juniper
		"d0df9a": 2874, // Liteon
		"d0dfb2": 4765, // Genie
		"d0dfc7": 152,  // Samsung
		"d0e042": 2,    // Cisco
		"d0e140": 545,  // Apple
		"d0e347": 8966, // Yoga
		"d0e40b": 8967, // Wearable
		"d0e44a": 2057, // Murata
		"d0e54d": 125,  // ARRIS
		"d0e782": 3005, // Azurewave
		"d0e828": 228,  // Radiant
		"d0eb03": 8968, // Zhehua
		"d0eb9e": 8969, // Seowoo
		"d0ec35": 2,    // Cisco
		"d0efc1": 3335, // Huawei
		"d0f0db": 306,  // Ericsson
		"d0f3f5": 3335, // Huawei
		"d0f520": 2849, // Kyocera
		"d0f865": 6282, // Itel
		"d0f88c": 687,  // Motorola
		"d0f99b": 3032, // ZTE
		"d0fccc": 152,  // Samsung
		"d0ff98": 3335, // Huawei
		"d40057": 8970, // MC
		"d40129": 858,  // Broadcom
		"d4016d": 1595, // TP-Link
		"d4024a": 8971, // Delphian
		"d404cd": 125,  // ARRIS
		"d404ff": 826,  // Juniper
		"d40598": 125,  // ARRIS
		"d405de": 5821, // eero
		"d40aa9": 125,  // ARRIS
		"d40b1a": 1341, // HTC
		"d41090": 8972, // InfORM
		"d411a3": 152,  // Samsung
		"d411d6": 8973, // ShotSpotter
		"d41243": 4499, // AMPAK
		"d41296": 8974, // Anobit
		"d413f8": 2482, // Peplink
		"d41ad1": 2693, // ZyXEL
		"d41e35": 8975, // TOHO
		"d41f0c": 1873, // JAI
		"d4206d": 1341, // HTC
		"d420b0": 7485, // Mist
		"d42122": 2075, // Sercomm
		"d4223f": 2665, // Lenovo
		"d42493": 2282, // GW
		"d4258b": 421,  // Intel
		"d42751": 8976, // Infopia
		"d428b2": 8977, // ioBridge
		"d428d5": 6434, // TCT
		"d429ea": 8978, // Zimory
		"d42c0f": 125,  // ARRIS
		"d42c44": 2,    // Cisco
		"d42c46": 1077, // Buffalo
		"d42dc5": 2154, // Panasonic
		"d42f23": 8979, // Akenori
		"d4319d": 8980, // Sinwatec
		"d43260": 6242, // GoPro
		"d43266": 8981, // Fike
		"d4351d": 6393, // Technicolor
		"d4354a": 2654, // ALAXALA
		"d437d7": 3032, // ZTE
		"d4389c": 101,  // Sony
		"d439b8": 4565, // Ciena
		"d43a65": 8982, // IGRS
		"d43b04": 421,  // Intel
		"d43d67": 8983, // Carma
		"d43d7e": 1812, // Micro-Star
		"d43df3": 2693, // ZyXEL
		"d43fcb": 125,  // ARRIS
		"d440d0": 8984, // OCOSMOS
		"d440f0": 3335, // Huawei
		"d44649": 3335, // Huawei
		"d446e1": 545,  // Apple
		"d4475a": 8985, // ScreenBeam
		"d44ca7": 8986, // Informtekhnika
		"d44da4": 2057, // Murata
		"d44f67": 3335, // Huawei
		"d44f68": 8987, // Eidetic
		"d4522a": 8988, // TangoWiFi.com
		"d45297": 8989, // nSTREAMS
		"d452ee": 3513, // BSkyB
		"d45383": 2057, // Murata
		"d4548b": 421,  // Intel
		"d45763": 545,  // Apple
		"d45800": 4922, // Fiberhome
		"d45ab2": 8990, // Galleon
		"d45d42": 457,  // Nokia
		"d45d64": 1806, // ASUS
		"d45ddf": 5440, // Pegatron
		"d460e3": 2075, // Sercomm
		"d4612e": 3335, // Huawei
		"d4619d": 545,  // Apple
		"d461da": 545,  // Apple
		"d462ea": 3335, // Huawei
		"d463c6": 687,  // Motorola
		"d463de": 6370, // Vivo
		"d463fe": 2640, // Arcadyan
		"d46624": 2,    // Cisco
		"d466a8": 8991, // Riedo
		"d46761": 8992, // XonTel
		"d467e7": 4922, // Fiberhome
		"d4684d": 2738, // Ruckus
		"d468aa": 545,  // Apple
		"d469a5": 3718, // Miura
		"d46a35": 2,    // Cisco
		"d46a91": 6557, // SnapAV
		"d46aa8": 3335, // Huawei
		"d46ba6": 3335, // Huawei
		"d46c6d": 125,  // ARRIS
		"d46cda": 8993, // CSM
		"d46d50": 2,    // Cisco
		"d46d6d": 421,  // Intel
		"d46e0e": 1595, // TP-Link
		"d46e5c": 3335, // Huawei
		"d47208": 8994, // Bragi
		"d47226": 3032, // ZTE
		"d47415": 3335, // Huawei
		"d476a0": 1323, // Fortinet
		"d476ea": 3032, // ZTE
		"d47798": 2,    // Cisco
		"d47856": 620,  // Avaya
		"d4789b": 2,    // Cisco
		"d47954": 3335, // Huawei
		"d479c3": 8995, // Cameronet
		"d47ae2": 152,  // Samsung
		"d47b75": 1596, // HARTING
		"d47bb0": 2545, // Askey
		"d47dfc": 6266, // Tecno
		"d481ca": 8996, // iDevices
		"d481d7": 102,  // Dell
		"d4823e": 2401, // Argosy
		"d48564": 302,  // HP
		"d48660": 2640, // Arcadyan
		"d487d8": 152,  // Samsung
		"d4883f": 8997, // Hdpro
		"d48866": 3335, // Huawei
		"d48890": 152,  // Samsung
		"d48a39": 152,  // Samsung
		"d48cb5": 2,    // Cisco
		"d48dd9": 8998, // Meld
		"d48f33": 612,  // Microsoft
		"d48fa2": 3335, // Huawei
		"d48faa": 8999, // Sogecam
		"d4909c": 545,  // Apple
		"d490e0": 8489, // Topcon
		"d4910f": 5439, // Amazon
		"d49234": 48,   // NEC
		"d49390": 5688, // Clevo
		"d49398": 457,  // Nokia
		"d493a0": 9000, // Fidelix
		"d49400": 3335, // Huawei
		"d4945a": 354,  // Cosmo
		"d494e8": 3335, // Huawei
		"d4970b": 5698, // Xiaomi
		"d49a20": 545,  // Apple
		"d49aa0": 8376, // Vnpt
		"d49c28": 9001, // JayBird
		"d49cdd": 4499, // AMPAK
		"d49dc0": 152,  // Samsung
		"d49e05": 3032, // ZTE
		"d49fdd": 3335, // Huawei
		"d4a02a": 2,    // Cisco
		"d4a148": 3335, // Huawei
		"d4a33d": 545,  // Apple
		"d4a425": 9002, // SMAX
		"d4a499": 9003, // InView
		"d4ab82": 125,  // ARRIS
		"d4ad2d": 4922, // Fiberhome
		"d4ad71": 2,    // Cisco
		"d4adbd": 2,    // Cisco
		"d4adfc": 67,   // Private
		"d4ae05": 152,  // Samsung
		"d4ae52": 102,  // Dell
		"d4aff7": 3794, // Arista
		"d4b110": 3335, // Huawei
		"d4b27a": 125,  // ARRIS
		"d4b709": 3032, // ZTE
		"d4b7d0": 4565, // Ciena
		"d4b92f": 6393, // Technicolor
		"d4bbc8": 6370, // Vivo
		"d4bbe6": 3335, // Huawei
		"d4bd1e": 9004, // 5VT
		"d4bd4f": 2738, // Ruckus
		"d4bed9": 102,  // Dell
		"d4bf7f": 9005, // Upvel
		"d4c19e": 2738, // Ruckus
		"d4c1c8": 3032, // ZTE
		"d4c1fc": 457,  // Nokia
		"d4c3b0": 9006, // Gearlinx
		"d4c766": 9007, // Acentic
		"d4c8b0": 1040, // Prime
		"d4c93c": 2,    // Cisco
		"d4c94b": 687,  // Motorola
		"d4c9b2": 9008, // Quanergy
		"d4c9ef": 302,  // HP
		"d4ca6d": 1784, // Routerboard.com
		"d4ca6e": 7420, // u-blox
		"d4cbaf": 457,  // Nokia
		"d4ceb8": 9009, // Enatel
		"d4d252": 421,  // Intel
		"d4d2d6": 6441, // FN-Link
		"d4d2e5": 9010, // BKAV
		"d4d51b": 3335, // Huawei
		"d4d748": 2,    // Cisco
		"d4d919": 6242, // GoPro
		"d4dacd": 3513, // BSkyB
		"d4dc09": 7485, // Mist
		"d4dccd": 545,  // Apple
		"d4e08e": 9011, // ValueHD
		"d4e2cb": 6393, // Technicolor
		"d4e33f": 457,  // Nokia
		"d4e6b7": 152,  // Samsung
		"d4e880": 2,    // Cisco
		"d4e8b2": 152,  // Samsung
		"d4e90b": 9012, // CVT
		"d4ea0e": 620,  // Avaya
		"d4eb68": 2,    // Cisco
		"d4ec0c": 9013, // Harley-Davidson
		"d4ecab": 6370, // Vivo
		"d4ee07": 9014, // HIWIFI
		"d4f057": 1427, // Nintendo
		"d4f143": 9015, // IPROAD
		"d4f207": 9016, // DIAODIAOTechnology
		"d4f337": 9017, // Xunison
		"d4f46f": 545,  // Apple
		"d4f527": 300,  // Siemens
		"d4f547": 3522, // Google
		"d4f5ef": 302,  // HP
		"d4f756": 3032, // ZTE
		"d4f786": 4922, // Fiberhome
		"d4f829": 2049, // Sagemcom
		"d4f98d": 6377, // Espressif
		"d4f9a1": 3335, // Huawei
		"d4fb8e": 545,  // Apple
		"d4fc13": 4922, // Fiberhome
		"d8004d": 545,  // Apple
		"d80093": 9018, // Aurender
		"d8052e": 9019, // Skyviia
		"d807b6": 1595, // TP-Link
		"d80831": 152,  // Samsung
		"d808f5": 9020, // Arcadia
		"d809c3": 9021, // Cercacor
		"d809d6": 9022, // Zexelon
		"d80a60": 3335, // Huawei
		"d80b9a": 152,  // Samsung
		"d80d17": 1595, // TP-Link
		"d80de3": 9023, // FXI
		"d81068": 2057, // Murata
		"d8109f": 3335, // Huawei
		"d8150d": 1595, // TP-Link
		"d8160a": 9024, // Nippon
		"d816c1": 7751, // Dewav
		"d818d3": 826,  // Juniper
		"d8197a": 9025, // Nuheara
		"d819ce": 9026, // Telesquare
		"d81bfe": 9027, // Twinlinx
		"d81c14": 9028, // Compacta
		"d81c79": 545,  // Apple
		"d81d72": 545,  // Apple
		"d81fcc": 90,   // Brocade
		"d824bd": 2,    // Cisco
		"d82522": 125,  // ARRIS
		"d825b0": 9029, // Rockeetech
		"d8270c": 9030, // MaxTronic
		"d82916": 1690, // Ascent
		"d82918": 3335, // Huawei
		"d82a15": 9031, // Leitner
		"d82a7e": 457,  // Nokia
		"d82de1": 9032, // Tricascade
		"d83062": 545,  // Apple
		"d83134": 1916, // Roku
		"d831cf": 152,  // Samsung
		"d83214": 6267, // Tenda
		"d832e3": 5698, // Xiaomi
		"d833b7": 2049, // Sagemcom
		"d8365f": 3542, // Intelbras
		"d838fc": 2738, // Ruckus
		"d83af5": 9033, // Wideband
		"d83bbf": 421,  // Intel
		"d84008": 3335, // Huawei
		"d843ed": 3603, // Suzuken
		"d84732": 1595, // TP-Link
		"d847bb": 3335, // Huawei
		"d8490b": 3335, // Huawei
		"d8492f": 87,   // Canon
		"d84a2b": 3032, // ZTE
		"d84b2a": 9034, // Cognitas
		"d84c90": 545,  // Apple
		"d84f37": 9035, // Proxis
		"d84fb8": 869,  // LG
		"d850e6": 1806, // ASUS
		"d8539a": 826,  // Juniper
		"d854a2": 185,  // Extreme
		"d85575": 152,  // Samsung
		"d855a3": 3032, // ZTE
		"d857ef": 152,  // Samsung
		"d858d7": 9036, // CZ.NIC
		"d85982": 3335, // Huawei
		"d85b2a": 152,  // Samsung
		"d85d4c": 1595, // TP-Link
		"d85dfb": 67,   // Private
		"d85ed3": 1929, // Giga-Byte
		"d86162": 1592, // Wistron
		"d862db": 9037, // Eno
		"d86375": 5698, // Xiaomi
		"d866ee": 9038, // Boxin
		"d867d3": 3335, // Huawei
		"d867d9": 2,    // Cisco
		"d86852": 3335, // Huawei
		"d868c3": 152,  // Samsung
		"d86960": 9039, // Steinsvik
		"d86bf7": 1427, // Nintendo
		"d86c5a": 531,  // HUMAX
		"d86c63": 3522, // Google
		"d86ce9": 2049, // Sagemcom
		"d86d17": 3335, // Huawei
		"d87157": 2665, // Lenovo
		"d87495": 3032, // ZTE
		"d87533": 457,  // Nokia
		"d8760a": 9040, // Escort
		"d876ae": 3335, // Huawei
		"d8778b": 3542, // Intelbras
		"d878e5": 9041, // Kuhn
		"d87cdd": 9042, // Sanix
		"d87d7f": 2049, // Sagemcom
		"d87e76": 6282, // Itel
		"d87eb1": 9043, // x.o.ware
		"d88039": 706,  // Microchip
		"d881ce": 9044, // AHN
		"d88466": 185,  // Extreme
		"d887d5": 9045, // Leadcore
		"d888ce": 2856, // RF
		"d88a3b": 9046, // Unit-EM
		"d88adc": 3335, // Huawei
		"d88b4c": 9047, // KingTing
		"d88c73": 3032, // ZTE
		"d88c79": 3522, // Google
		"d88d5c": 4653, // Elentec
		"d88dc8": 9048, // Atil
		"d88f76": 545,  // Apple
		"d890e8": 152,  // Samsung
		"d8912a": 2693, // ZyXEL
		"d89403": 302,  // HP
		"d89685": 6242, // GoPro
		"d89695": 545,  // Apple
		"d897ba": 5440, // Pegatron
		"d89ac1": 457,  // Nokia
		"d89b3b": 3335, // Huawei
		"d89d67": 302,  // HP
		"d89db9": 9049, // eMegatech
		"d89e3f": 545,  // Apple
		"d89e61": 3335, // Huawei
		"d89ed4": 4922, // Fiberhome
		"d89ef3": 102,  // Dell
		"d8a011": 7178, // WiZ
		"d8a01d": 6377, // Espressif
		"d8a105": 9050, // Syslane
		"d8a25e": 545,  // Apple
		"d8a315": 6370, // Vivo
		"d8a35c": 152,  // Samsung
		"d8a491": 3335, // Huawei
		"d8a534": 9051, // Spectronix
		"d8a756": 2049, // Sagemcom
		"d8a8c8": 3032, // ZTE
		"d8aa59": 6895, // Tonly
		"d8addd": 9052, // Sonavation
		"d8ae90": 9053, // Itibia
		"d8aed0": 5009, // Shanghai
		"d8aff1": 2154, // Panasonic
		"d8b053": 5698, // Xiaomi
		"d8b122": 826,  // Juniper
		"d8b12a": 2154, // Panasonic
		"d8b190": 2,    // Cisco
		"d8b377": 1341, // HTC
		"d8b6b7": 3870, // Comtrend
		"d8b6c1": 9054, // NetworkAccountant
		"d8b8f6": 9055, // Nantworks
		"d8bb2c": 545,  // Apple
		"d8bbc1": 1812, // Micro-Star
		"d8be1f": 545,  // Apple
		"d8be65": 5439, // Amazon
		"d8bfc0": 6377, // Espressif
		"d8c068": 9056, // Netgenetech
		"d8c0a6": 3005, // Azurewave
		"d8c3fb": 9057, // Detracom
		"d8c46a": 2057, // Murata
		"d8c497": 3072, // Quanta
		"d8c4e9": 152,  // Samsung
		"d8c561": 9058, // CommFront
		"d8c678": 6429, // MitraStar
		"d8c691": 9059, // Hichan
		"d8c771": 3335, // Huawei
		"d8c7c8": 1685, // Aruba
		"d8c8e9": 6849, // Phicomm
		"d8cb8a": 1812, // Micro-Star
		"d8cc98": 3335, // Huawei
		"d8ce3a": 5698, // Xiaomi
		"d8cf9c": 545,  // Apple
		"d8cfbf": 687,  // Motorola
		"d8d090": 102,  // Dell
		"d8d1cb": 545,  // Apple
		"d8d385": 302,  // HP
		"d8d43c": 101,  // Sony
		"d8d723": 9060, // IDS
		"d8d775": 2049, // Sagemcom
		"d8dc40": 545,  // Apple
		"d8dd5f": 9061, // BALMUDA
		"d8de3a": 545,  // Apple
		"d8dece": 9062, // Isung
		"d8df0d": 9063, // beroNet
		"d8df7a": 2508, // Quest
		"d8e004": 9064, // Vodia
		"d8e0b8": 9065, // Bulat
		"d8e0e1": 152,  // Samsung
		"d8e2df": 612,  // Microsoft
		"d8e56d": 6434, // TCT
		"d8e72b": 5516, // Netscout
		"d8e743": 9066, // Wush
		"d8e952": 9067, // Keopsys
		"d8eb46": 3522, // Google
		"d8eb97": 2905, // TRENDnet
		"d8ec5e": 2469, // Belkin
		"d8ece5": 2693, // ZyXEL
		"d8ef42": 3335, // Huawei
		"d8f0f2": 9068, // Zeebo
		"d8f15b": 6377, // Espressif
		"d8f1f0": 9069, // Pepxim
		"d8f2ca": 421,  // Intel
		"d8f3bc": 2874, // Liteon
		"d8f883": 421,  // Intel
		"d8f8af": 9070, // Daontec
		"d8fb11": 9071, // Axacore
		"d8fb5e": 2545, // Askey
		"d8fbd6": 5439, // Amazon
		"d8fc93": 421,  // Intel
		"d8fe8f": 9072, // IDFone
		"d8fee3": 803,  // D-Link
		"dc0077": 1595, // TP-Link
		"dc028e": 3032, // ZTE
		"dc0398": 869,  // LG
		"dc052f": 9073, // National
		"dc0575": 300,  // Siemens
		"dc05ed": 9074, // Nabtesco
		"dc080f": 545,  // Apple
		"dc0914": 4053, // Talk-A-Phone
		"dc094c": 3335, // Huawei
		"dc0b34": 869,  // LG
		"dc0c5c": 545,  // Apple
		"dc0ea1": 358,  // Compal
		"dc16b2": 3335, // Huawei
		"dc1a01": 9075, // Ecoliv
		"dc1ac5": 6370, // Vivo
		"dc1ba1": 421,  // Intel
		"dc1ea3": 9076, // Accensus
		"dc2008": 9077, // ASD
		"dc2148": 421,  // Intel
		"dc215c": 421,  // Intel
		"dc21b9": 9078, // Sentec
		"dc21e2": 3335, // Huawei
		"dc2727": 3335, // Huawei
		"dc2834": 9079, // HAKKO
		"dc2919": 7032, // AltoBeam
		"dc2aa1": 9080, // MedHab
		"dc2b2a": 545,  // Apple
		"dc2b61": 545,  // Apple
		"dc2bca": 9081, // Zera
		"dc2c26": 6869, // Iton
		"dc2c6e": 1784, // Routerboard.com
		"dc2d3c": 3335, // Huawei
		"dc2e6a": 9082, // HCT
		"dc309c": 9083, // Heyrex
		"dc31d1": 6370, // Vivo
		"dc333d": 3335, // Huawei
		"dc3350": 9084, // TechSAT
		"dc3714": 545,  // Apple
		"dc3752": 5811, // GE
		"dc38e1": 826,  // Juniper
		"dc3979": 2,    // Cisco
		"dc3a5e": 1916, // Roku
		"dc3ef8": 457,  // Nokia
		"dc415f": 545,  // Apple
		"dc41a9": 421,  // Intel
		"dc446d": 9085, // Allwinner
		"dc44b6": 152,  // Samsung
		"dc4517": 125,  // ARRIS
		"dc48b2": 9086, // Baraja
		"dc4a3e": 302,  // HP
		"dc4ede": 9087, // Shinyei
		"dc4f22": 6377, // Espressif
		"dc5285": 545,  // Apple
		"dc5360": 421,  // Intel
		"dc537c": 358,  // Compal
		"dc5392": 545,  // Apple
		"dc543d": 6282, // Itel
		"dc54d7": 5439, // Amazon
		"dc56e7": 545,  // Apple
		"dc5726": 1515, // Power-One
		"dc58bc": 9088, // Thomas-Krenn.AG
		"dc5e36": 9089, // Paterson
		"dc663a": 9090, // Apacer
		"dc6672": 152,  // Samsung
		"dc680c": 302,  // HP
		"dc68eb": 1427, // Nintendo
		"dc6aea": 6296, // Infinix
		"dc6b12": 9091, // worldcns
		"dc6b1b": 3335, // Huawei
		"dc6f00": 9092, // Livescribe
		"dc7014": 67,   // Private
		"dc7137": 3032, // ZTE
		"dc7144": 152,  // Samsung
		"dc7196": 421,  // Intel
		"dc729b": 3335, // Huawei
		"dc7385": 3335, // Huawei
		"dc74a8": 152,  // Samsung
		"dc774c": 2,    // Cisco
		"dc7834": 9093, // Logicom
		"dc7b94": 2,    // Cisco
		"dc7fa4": 1939, // 2Wire
		"dc8084": 545,  // Apple
		"dc825b": 9094, // JANUS
		"dc82f6": 9095, // iPort
		"dc85de": 3005, // Azurewave
		"dc86d8": 545,  // Apple
		"dc8983": 152,  // Samsung
		"dc8b28": 421,  // Intel
		"dc8c1b": 6370, // Vivo
		"dc8c37": 2,    // Cisco
		"dc9088": 3335, // Huawei
		"dc9166": 3335, // Huawei
		"dc91bf": 5439, // Amazon
		"dc973a": 9096, // Verana
		"dc9840": 612,  // Microsoft
		"dc9914": 3335, // Huawei
		"dc99fe": 9097, // Armatura
		"dc9b1e": 9098, // Intercom
		"dc9b9c": 545,  // Apple
		"dc9bd6": 6434, // TCT
		"dc9c52": 9099, // Sapphire
		"dc9fa4": 457,  // Nokia
		"dc9fdb": 2979, // Ubiquiti
		"dca120": 457,  // Nokia
		"dca3ac": 9100, // RBcloudtech
		"dca4ca": 545,  // Apple
		"dca5f4": 2,    // Cisco
		"dca633": 125,  // ARRIS
		"dca782": 3335, // Huawei
		"dca904": 545,  // Apple
		"dca971": 421,  // Intel
		"dca989": 9101, // Macandc
		"dcad9e": 9102, // GreenPriz
		"dcae04": 9103, // CELOXICA
		"dcaeeb": 2738, // Ruckus
		"dcb082": 457,  // Nokia
		"dcb4ac": 833,  // Flextronics
		"dcb4c4": 612,  // Microsoft
		"dcb54f": 545,  // Apple
		"dcb72e": 5698, // Xiaomi
		"dcb7fc": 5752, // Alps
		"dcb808": 185,  // Extreme
		"dcbfe9": 687,  // Motorola
		"dcc101": 9104, // SOLiD
		"dcc422": 9105, // Systembase
		"dcc64b": 3335, // Huawei
		"dcc793": 457,  // Nokia
		"dccba8": 9106, // Explora
		"dccce6": 152,  // Samsung
		"dccd2f": 46,   // Epson
		"dccec1": 2,    // Cisco
		"dccf96": 152,  // Samsung
		"dcd0f7": 9107, // Bentek
		"dcd255": 299,  // Kinpo
		"dcd2fc": 3335, // Huawei
		"dcd321": 531,  // HUMAX
		"dcd3a2": 545,  // Apple
		"dcd444": 3335, // Huawei
		"dcd7a0": 3335, // Huawei
		"dcd916": 3335, // Huawei
		"dcda4f": 9108, // Getck
		"dcdc07": 9109, // TRP
		"dcdce2": 152,  // Samsung
		"dcdd24": 9110, // Energica
		"dcdeca": 9111, // Akyllor
		"dcdfd6": 3032, // ZTE
		"dce55b": 3522, // Google
		"dceb69": 6393, // Technicolor
		"dceb94": 2,    // Cisco
		"dced84": 9112, // Haverford
		"dcee06": 3335, // Huawei
		"dcef09": 1368, // Netgear
		"dcef80": 3335, // Huawei
		"dcefca": 2057, // Murata
		"dcf090": 9113, // Nubia
		"dcf110": 457,  // Nokia
		"dcf401": 102,  // Dell
		"dcf4ca": 545,  // Apple
		"dcf505": 3005, // Azurewave
		"dcf56e": 9114, // Wellysis
		"dcf719": 2,    // Cisco
		"dcf755": 9115, // Sitronik
		"dcf756": 152,  // Samsung
		"dcf858": 9116, // Lorent
		"dcf8b9": 3032, // ZTE
		"dcfb02": 1077, // Buffalo
		"dcfb48": 421,  // Intel
		"dcfe07": 5440, // Pegatron
		"dcfe18": 1595, // TP-Link
		"dcfe23": 2057, // Murata
		"e00084": 3335, // Huawei
		"e0036b": 152,  // Samsung
		"e005c5": 1595, // TP-Link
		"e0071b": 302,  // HP
		"e00af6": 2874, // Liteon
		"e00b28": 9117, // Inovonics
		"e00c7f": 1427, // Nintendo
		"e00ce5": 3335, // Huawei
		"e00db9": 9118, // Cree
		"e00eda": 2,    // Cisco
		"e00ee1": 6291, // We
		"e00ee4": 7447, // DWnet
		"e0107f": 2738, // Ruckus
		"e013b5": 6370, // Vivo
		"e0143e": 9119, // Modoosis
		"e01877": 4,    // Fujitsu
		"e0191d": 3335, // Huawei
		"e01954": 3032, // ZTE
		"e01995": 7341, // Nutanix
		"e019d8": 9120, // BH
		"e01c41": 185,  // Extreme
		"e01cee": 9121, // Bravo
		"e01cfc": 803,  // D-Link
		"e01d3b": 1640, // Cambridge
		"e01f88": 5698, // Xiaomi
		"e02202": 125,  // ARRIS
		"e023ff": 1323, // Fortinet
		"e0247f": 3335, // Huawei
		"e02481": 3335, // Huawei
		"e02630": 9122, // Intrigue
		"e02636": 76,   // Nortel
		"e02861": 3335, // Huawei
		"e02ae6": 4922, // Fiberhome
		"e02b96": 545,  // Apple
		"e02be9": 421,  // Intel
		"e02cb2": 2665, // Lenovo
		"e02cf3": 9123, // MRS
		"e02e3f": 3335, // Huawei
		"e02f6d": 2,    // Cisco
		"e030f9": 826,  // Juniper
		"e0319e": 9124, // Valve
		"e0338e": 545,  // Apple
		"e03676": 3335, // Huawei
		"e03717": 6393, // Technicolor
		"e037bf": 1592, // Wistron
		"e0383f": 3032, // ZTE
		"e039d7": 9125, // Plexxi
		"e03e44": 858,  // Broadcom
		"e03e7d": 9126, // data-Complex
		"e03f49": 1806, // ASUS
		"e04007": 3335, // Huawei
		"e04102": 3032, // ZTE
		"e04136": 6429, // MitraStar
		"e0469a": 1368, // Netgear
		"e046ee": 1368, // Netgear
		"e048af": 9127, // Premietech
		"e049ed": 9128, // Audeze
		"e04b45": 9129, // Hi-P
		"e04ba6": 3335, // Huawei
		"e05163": 2640, // Arcadyan
		"e056f4": 9130, // AxesNetwork
		"e05b70": 9131, // Innovid
		"e05f45": 545,  // Apple
		"e05fb9": 2,    // Cisco
		"e06066": 2075, // Sercomm
		"e06089": 9132, // Cloudleaf
		"e06267": 5698, // Xiaomi
		"e063da": 2979, // Ubiquiti
		"e063e5": 101,  // Sony
		"e06678": 545,  // Apple
		"e0686d": 9133, // Raybased
		"e0693a": 9134, // Innophase
		"e06995": 5440, // Pegatron
		"e069ba": 2,    // Cisco
		"e06cf6": 9135, // ESSENCORE
		"e06d17": 545,  // Apple
		"e06d18": 9136, // Pioneercorporation
		"e070ea": 302,  // HP
		"e0735f": 5777, // Nucom
		"e0750a": 430,  // Alpsalpine
		"e0757d": 687,  // Motorola
		"e076d0": 4499, // AMPAK
		"e07726": 3335, // Huawei
		"e079c4": 9137, // iRay
		"e07c13": 3032, // ZTE
		"e07c62": 5352, // Whistle
		"e07e5f": 5696, // Renesas
		"e08177": 9138, // GreenBytes
		"e087b1": 9139, // Nata-Info
		"e0885d": 6393, // Technicolor
		"e0897e": 545,  // Apple
		"e0899d": 2,    // Cisco
		"e08a7e": 9140, // Exponent
		"e08e3c": 4855, // Aztech
		"e08fec": 9141, // Repotec
		"e09153": 190,  // XAVi
		"e091f5": 1368, // Netgear
		"e0925c": 545,  // Apple
		"e092a7": 7587, // Feitian
		"e09467": 421,  // Intel
		"e09579": 9142, // ORTHOsoft
		"e09796": 3335, // Huawei
		"e097f2": 9143, // Atomax
		"e09806": 6377, // Espressif
		"e09861": 687,  // Motorola
		"e09971": 152,  // Samsung
		"e09d13": 152,  // Samsung
		"e09d31": 421,  // Intel
		"e09db8": 4479, // Planex
		"e09f2a": 6869, // Iton
		"e0a1d7": 3188, // SFR
		"e0a30f": 9144, // Pevco
		"e0a3ac": 3335, // Huawei
		"e0a509": 9145, // Bitmain
		"e0a670": 457,  // Nokia
		"e0a700": 9146, // Verkada
		"e0aa96": 152,  // Samsung
		"e0aab0": 9147, // Suntaili
		"e0abfe": 9148, // Orb
		"e0accb": 545,  // Apple
		"e0acf1": 2,    // Cisco
		"e0ae5e": 430,  // Alpsalpine
		"e0aea2": 3335, // Huawei
		"e0aeed": 9149, // Loenk
		"e0af4b": 7576, // Pluribus
		"e0b2f1": 6441, // FN-Link
		"e0b52d": 545,  // Apple
		"e0b55f": 545,  // Apple
		"e0b70a": 125,  // ARRIS
		"e0b7b1": 125,  // ARRIS
		"e0b9a5": 3005, // Azurewave
		"e0b9ba": 545,  // Apple
		"e0b9e5": 6393, // Technicolor
		"e0bab4": 9150, // Arrcus
		"e0bb0c": 9151, // Synertau
		"e0bb9e": 46,   // Epson
		"e0c286": 9152, // Aisai
		"e0c2b7": 6363, // Masimo
		"e0c377": 152,  // Samsung
		"e0c3f3": 3032, // ZTE
		"e0c6b3": 9153, // MilDef
		"e0c767": 545,  // Apple
		"e0c97a": 545,  // Apple
		"e0ca94": 2545, // Askey
		"e0cb1d": 67,   // Private
		"e0cb4e": 1806, // ASUS
		"e0cbee": 152,  // Samsung
		"e0cc7a": 3335, // Huawei
		"e0ccf8": 5698, // Xiaomi
		"e0cec3": 2545, // Askey
		"e0cf2d": 9154, // Gemintek
		"e0d045": 421,  // Intel
		"e0d083": 152,  // Samsung
		"e0d10a": 9155, // Katoudenkikougyousyo
		"e0d173": 2,    // Cisco
		"e0d31a": 9156, // EQUES
		"e0d462": 3335, // Huawei
		"e0d464": 421,  // Intel
		"e0d4e8": 421,  // Intel
		"e0d55e": 1929, // Giga-Byte
		"e0d848": 102,  // Dell
		"e0d9e3": 8497, // Eltex
		"e0da90": 3335, // Huawei
		"e0db10": 152,  // Samsung
		"e0db55": 102,  // Dell
		"e0dbd1": 6393, // Technicolor
		"e0dca0": 300,  // Siemens
		"e0dcff": 5698, // Xiaomi
		"e0ddc0": 6370, // Vivo
		"e0e0fc": 3335, // Huawei
		"e0e2e6": 6377, // Espressif
		"e0e37c": 3335, // Huawei
		"e0e62e": 6434, // TCT
		"e0e631": 9157, // SNB
		"e0e751": 1427, // Nintendo
		"e0e7bb": 9158, // Nureva
		"e0eb40": 545,  // Apple
		"e0ed1a": 9159, // vastriver
		"e0ee1b": 2154, // Panasonic
		"e0ef25": 9160, // Lintes
		"e0f211": 9161, // Digitalwatt
		"e0f379": 9162, // Vaddio
		"e0f442": 3335, // Huawei
		"e0f5c6": 545,  // Apple
		"e0f62d": 826,  // Juniper
		"e0f6b5": 1427, // Nintendo
		"e0f847": 545,  // Apple
		"e0f9be": 9163, // Cloudena
		"e0fff7": 9164, // Softiron
		"e4029b": 421,  // Intel
		"e40439": 2715, // TomTom
		"e405f8": 9165, // Bytedance
		"e4072b": 3335, // Huawei
		"e40eee": 3335, // Huawei
		"e4115b": 302,  // HP
		"e4121d": 152,  // Samsung
		"e4186b": 2693, // ZyXEL
		"e419c1": 3335, // Huawei
		"e41a2c": 9166, // ZPE
		"e41c4b": 9167, // V2
		"e41d2d": 432,  // Mellanox
		"e41f13": 372,  // IBM
		"e41f7b": 2,    // Cisco
		"e41fe9": 9168, // Dunkermotoren
		"e422a5": 542,  // Plantronics
		"e425e7": 545,  // Apple
		"e425e9": 9169, // Color-Chip
		"e42686": 7447, // DWnet
		"e4268b": 3335, // Huawei
		"e42771": 2134, // Smartlabs
		"e42aac": 612,  // Microsoft
		"e42b34": 545,  // Apple
		"e42b79": 457,  // Nokia
		"e42c56": 9170, // Lilee
		"e42d02": 6434, // TCT
		"e42f26": 4922, // Fiberhome
		"e42f56": 9171, // OptoMET
		"e42ff6": 9172, // Unicore
		"e432cb": 152,  // Samsung
		"e43493": 3335, // Huawei
		"e435c8": 3335, // Huawei
		"e435fb": 9173, // Sabre
		"e43883": 2979, // Ubiquiti
		"e4388c": 1565, // Digital
		"e43a65": 9174, // MofiNetwork
		"e43d1a": 858,  // Broadcom
		"e43ec6": 3335, // Huawei
		"e43ed7": 2640, // Arcadyan
		"e440e2": 152,  // Samsung
		"e44122": 6924, // OnePlus
		"e44164": 457,  // Nokia
		"e441e6": 9175, // Ottec
		"e442a6": 421,  // Intel
		"e4434b": 102,  // Dell
		"e444e5": 185,  // Extreme
		"e446da": 5698, // Xiaomi
		"e447b3": 3032, // ZTE
		"e44e2d": 2,    // Cisco
		"e44e76": 9176, // Championtech
		"e4509a": 9177, // HW
		"e450eb": 545,  // Apple
		"e454e8": 102,  // Dell
		"e45740": 125,  // ARRIS
		"e457a8": 9178, // Stuart
		"e458b8": 152,  // Samsung
		"e458e7": 152,  // Samsung
		"e45aa2": 6370, // Vivo
		"e45ad4": 8497, // Eltex
		"e45d37": 826,  // Juniper
		"e45d51": 3188, // SFR
		"e45d52": 620,  // Avaya
		"e45d75": 152,  // Samsung
		"e45e1b": 3522, // Google
		"e45e37": 421,  // Intel
		"e46059": 9179, // Pingtek
		"e46449": 125,  // ARRIS
		"e468a3": 3335, // Huawei
		"e46c21": 9180, // messMa
		"e46f13": 803,  // D-Link
		"e470b8": 421,  // Intel
		"e47185": 9181, // Securifi
		"e472e2": 3335, // Huawei
		"e475dc": 2640, // Arcadyan
		"e47684": 545,  // Apple
		"e47723": 3032, // ZTE
		"e47727": 3335, // Huawei
		"e4776b": 9182, // Aartesys
		"e477d4": 9183, // Minrray
		"e47b3f": 9184, // Beijing-Cloud
		"e47c65": 9185, // Sunstar
		"e47cf9": 152,  // Samsung
		"e47dbd": 152,  // Samsung
		"e47e66": 3335, // Huawei
		"e47e9a": 3032, // ZTE
		"e47fb2": 4,    // Fujitsu
		"e48184": 457,  // Nokia
		"e48210": 3335, // Huawei
		"e482cc": 9186, // Jumptronic
		"e48326": 3335, // Huawei
		"e48399": 125,  // ARRIS
		"e484d3": 5698, // Xiaomi
		"e48501": 9187, // Geberit
		"e48b7f": 545,  // Apple
		"e48d8c": 1784, // Routerboard.com
		"e48f1d": 3335, // Huawei
		"e4907e": 687,  // Motorola
		"e490fd": 545,  // Apple
		"e4922a": 9188, // DBG
		"e492e7": 9189, // Gridlink
		"e492fb": 152,  // Samsung
		"e496ae": 9190, // ALTOGRAPHICS
		"e498d1": 612,  // Microsoft
		"e498d6": 545,  // Apple
		"e49a79": 545,  // Apple
		"e49adc": 545,  // Apple
		"e49f1e": 125,  // ARRIS
		"e4a387": 1999, // Control
		"e4a471": 421,  // Intel
		"e4a7a0": 421,  // Intel
		"e4a7c5": 3335, // Huawei
		"e4a8b6": 3335, // Huawei
		"e4a8df": 358,  // Compal
		"e4aa5d": 2,    // Cisco
		"e4aaea": 2874, // Liteon
		"e4ab89": 6429, // MitraStar
		"e4afa1": 9191, // HES-SO
		"e4b021": 152,  // Samsung
		"e4b2fb": 545,  // Apple
		"e4b318": 421,  // Intel
		"e4b97a": 102,  // Dell
		"e4bd4b": 3032, // ZTE
		"e4beed": 2363, // Netcore
		"e4bffa": 6393, // Technicolor
		"e4c0e2": 2049, // Sagemcom
		"e4c2d1": 3335, // Huawei
		"e4c32a": 1595, // TP-Link
		"e4c62b": 9192, // Airware
		"e4c63d": 545,  // Apple
		"e4c6e6": 9193, // Mophie
		"e4c722": 2,    // Cisco
		"e4c801": 7461, // BLU
		"e4c90b": 9194, // Radwin
		"e4ca12": 3032, // ZTE
		"e4ce02": 9195, // WyreStorm
		"e4ce8f": 545,  // Apple
		"e4d124": 2486, // Mojo
		"e4d332": 1595, // TP-Link
		"e4d373": 3335, // Huawei
		"e4d3f1": 2,    // Cisco
		"e4dc43": 3335, // Huawei
		"e4dc5f": 9196, // Cofractal
		"e4dccc": 3335, // Huawei
		"e4e0a6": 545,  // Apple
		"e4e0c5": 152,  // Samsung
		"e4e130": 6434, // TCT
		"e4e409": 9197, // Leifheit
		"e4e4ab": 545,  // Apple
		"e4e749": 302,  // HP
		"e4ec10": 457,  // Nokia
		"e4eefd": 9198, // MR&D
		"e4f004": 102,  // Dell
		"e4f042": 3522, // Google
		"e4f14c": 67,   // Private
		"e4f1d4": 6370, // Vivo
		"e4f327": 9199, // Atol
		"e4f365": 9200, // Time-O-Matic
		"e4f3c4": 152,  // Samsung
		"e4f4c6": 1368, // Netgear
		"e4f75b": 125,  // ARRIS
		"e4f7a1": 9201, // Datafox
		"e4f89c": 421,  // Intel
		"e4f8ef": 152,  // Samsung
		"e4faed": 152,  // Samsung
		"e4fafd": 421,  // Intel
		"e4fb5d": 3335, // Huawei
		"e4fc82": 826,  // Juniper
		"e4fd45": 421,  // Intel
		"e4fda1": 3335, // Huawei
		"e4fed9": 6662, // EDMI
		"e80036": 9202, // Befs
		"e8018d": 4922, // Fiberhome
		"e8039a": 152,  // Samsung
		"e8040b": 545,  // Apple
		"e80410": 67,   // Private
		"e80462": 2,    // Cisco
		"e804f3": 9203, // Throughtek
		"e8056d": 76,   // Nortel
		"e805dc": 1647, // Verifone
		"e80688": 545,  // Apple
		"e8088b": 3335, // Huawei
		"e80c75": 9204, // Syncbak
		"e80fc8": 3858, // Universal
		"e81132": 152,  // Samsung
		"e81367": 9205, // AIRSOUND
		"e8136e": 3335, // Huawei
		"e8150e": 457,  // Nokia
		"e81a58": 6023, // Technologic
		"e81b4b": 9206, // amnimo
		"e81b69": 2075, // Sercomm
		"e81cba": 1323, // Fortinet
		"e81cd8": 545,  // Apple
		"e81da8": 2738, // Ruckus
		"e81e92": 3335, // Huawei
		"e820e2": 531,  // HUMAX
		"e82689": 1685, // Aruba
		"e82877": 9207, // TMY
		"e828c1": 8497, // Eltex
		"e828d5": 9208, // Cots
		"e82a44": 2874, // Liteon
		"e82aea": 421,  // Intel
		"e82c6d": 4550, // SmartRG
		"e82e0c": 9209, // NETINT
		"e831cd": 6377, // Espressif
		"e8330d": 9210, // Xaptec
		"e83381": 125,  // ARRIS
		"e83617": 545,  // Apple
		"e8361d": 9211, // Sense
		"e8377a": 2693, // ZyXEL
		"e83935": 302,  // HP
		"e839df": 2545, // Askey
		"e83a12": 152,  // Samsung
		"e83a97": 35,   // Toshiba
		"e83eb6": 4556, // RIM
		"e83efb": 9212, // Geodesic
		"e83efc": 125,  // ARRIS
		"e83f67": 3335, // Huawei
		"e84040": 2,    // Cisco
		"e840f2": 5440, // Pegatron
		"e843b6": 6762, // QNAP
		"e848b8": 1595, // TP-Link
		"e84c56": 9213, // Intercept
		"e84d74": 3335, // Huawei
		"e84dd0": 3335, // Huawei
		"e84dec": 0,    // Xerox
		"e84e06": 9214, // Edup
		"e84e84": 152,  // Samsung
		"e84ece": 1427, // Nintendo
		"e84f25": 2057, // Murata
		"e84fa7": 3335, // Huawei
		"e8508b": 152,  // Samsung
		"e8516e": 9215, // TSMART
		"e855b4": 5502, // SAI
		"e85659": 9216, // Advanced-Connectek
		"e856d6": 8353, // NCTech
		"e85a8b": 5698, // Xiaomi
		"e85ad1": 4922, // Fiberhome
		"e85b5b": 869,  // LG
		"e85bb7": 2063, // Ample
		"e85c0a": 2,    // Cisco
		"e85f02": 545,  // Apple
		"e8617e": 2874, // Liteon
		"e861be": 9217, // Melec
		"e86549": 2,    // Cisco
		"e865d4": 6267, // Tenda
		"e866c4": 9218, // Diamanti
		"e86819": 3335, // Huawei
		"e868e7": 6377, // Espressif
		"e86a64": 4920, // LCFC
		"e86d52": 125,  // ARRIS
		"e86d54": 3177, // Digit
		"e86dcb": 152,  // Samsung
		"e86de9": 3335, // Huawei
		"e86e44": 3032, // ZTE
		"e86ff2": 2247, // Actiontec
		"e874c7": 9219, // Sentinhealth
		"e8757f": 9220, // FIRS
		"e87865": 545,  // Apple
		"e87f6b": 152,  // Samsung
		"e87f95": 545,  // Apple
		"e8802e": 545,  // Apple
		"e880d8": 9221, // GNTEK
		"e88152": 545,  // Apple
		"e88175": 3032, // ZTE
		"e8825b": 125,  // ARRIS
		"e884a5": 421,  // Intel
		"e884c6": 3335, // Huawei
		"e8854b": 545,  // Apple
		"e8892c": 125,  // ARRIS
		"e88d28": 545,  // Apple
		"e88df5": 2478, // ZNYX
		"e88e60": 9222, // NSD
		"e8910f": 4922, // Fiberhome
		"e89120": 687,  // Motorola
		"e89218": 9223, // Arcontia
		"e892a4": 869,  // LG
		"e89309": 152,  // Samsung
		"e89363": 457,  // Nokia
		"e893f3": 9224, // Graphiant
		"e894f6": 1595, // TP-Link
		"e898c2": 9225, // ZETLAB
		"e8995a": 9226, // PiiGAB
		"e899c4": 1341, // HTC
		"e89a8f": 3072, // Quanta
		"e89d87": 35,   // Toshiba
		"e89f39": 457,  // Nokia
		"e89f6d": 6377, // Espressif
		"e89f80": 2469, // Belkin
		"e8a0cd": 1427, // Nintendo
		"e8a1f8": 3032, // ZTE
		"e8a245": 826,  // Juniper
		"e8a660": 3335, // Huawei
		"e8a6ca": 3335, // Huawei
		"e8a72f": 612,  // Microsoft
		"e8a730": 545,  // Apple
		"e8a7f2": 9227, // sTraffic
		"e8abf3": 3335, // Huawei
		"e8acad": 3032, // ZTE
		"e8ada6": 2049, // Sagemcom
		"e8aec5": 3794, // Arista
		"e8b1fc": 421,  // Intel
		"e8b2ac": 545,  // Apple
		"e8b2fe": 531,  // HUMAX
		"e8b4c8": 152,  // Samsung
		"e8b541": 3032, // ZTE
		"e8b5d0": 102,  // Dell
		"e8b6c2": 826,  // Juniper
		"e8b748": 2,    // Cisco
		"e8ba70": 2,    // Cisco
		"e8bdd1": 3335, // Huawei
		"e8be81": 2049, // Sagemcom
		"e8c1d7": 796,  // Philips
		"e8c2dd": 6296, // Infinix
		"e8c417": 4922, // Fiberhome
		"e8c57a": 9228, // Ufispace
		"e8c74f": 2874, // Liteon
		"e8c7cf": 1592, // Wistron
		"e8cba1": 457,  // Nokia
		"e8cbed": 6440, // Chipsea
		"e8cc18": 803,  // D-Link
		"e8cc32": 2456, // Micronet
		"e8cd2d": 3335, // Huawei
		"e8ce06": 8027, // SkyHawke
		"e8d099": 4922, // Fiberhome
		"e8d0fc": 2874, // Liteon
		"e8d11b": 2545, // Askey
		"e8d2ff": 2049, // Sagemcom
		"e8d322": 2,    // Cisco
		"e8d765": 3335, // Huawei
		"e8d819": 3005, // Azurewave
		"e8d8d1": 302,  // HP
		"e8da00": 9229, // Kivo
		"e8da20": 1427, // Nintendo
		"e8daaa": 9230, // VideoHome
		"e8db84": 6377, // Espressif
		"e8de27": 1595, // TP-Link
		"e8ded6": 9231, // Intrising
		"e8dff2": 9232, // PRF
		"e8e0b7": 35,   // Toshiba
		"e8e1e1": 1450, // Gemtek
		"e8e1e2": 9233, // Energotest
		"e8e5d6": 152,  // Samsung
		"e8e875": 9234, // iS5
		"e8e8b7": 2057, // Murata
		"e8ea4d": 3335, // Huawei
		"e8ea6a": 9235, // StarTech.com
		"e8eb1b": 706,  // Microchip
		"e8eb34": 2,    // Cisco
		"e8ebd3": 432,  // Mellanox
		"e8ed05": 125,  // ARRIS
		"e8edd6": 1323, // Fortinet
		"e8edf3": 2,    // Cisco
		"e8ef89": 9236, // OPMEX
		"e8f1b0": 2049, // Sagemcom
		"e8f2e2": 869,  // LG
		"e8f375": 457,  // Nokia
		"e8f408": 421,  // Intel
		"e8f654": 3335, // Huawei
		"e8f724": 302,  // HP
		"e8f9d4": 3335, // Huawei
		"e8fa23": 3335, // Huawei
		"e8fbe9": 545,  // Apple
		"e8fcaf": 1368, // Netgear
		"e8fd35": 3335, // Huawei
		"e8fd90": 9237, // Turbostor
		"ec0133": 9238, // Trinus
		"ec01d5": 2,    // Cisco
		"ec0273": 1685, // Aruba
		"ec086b": 1595, // TP-Link
		"ec08e5": 687,  // Motorola
		"ec0d9a": 432,  // Mellanox
		"ec0de4": 5439, // Amazon
		"ec107b": 152,  // Samsung
		"ec13b2": 9239, // Netonix
		"ec13db": 826,  // Juniper
		"ec14f6": 9240, // BioControl
		"ec172f": 1595, // TP-Link
		"ec1a59": 2469, // Belkin
		"ec1bbd": 1655, // Silicon
		"ec1c5d": 300,  // Siemens
		"ec1d7f": 3032, // ZTE
		"ec1d8b": 2,    // Cisco
		"ec1f72": 152,  // Samsung
		"ec219f": 9241, // VidaBox
		"ec21e5": 35,   // Toshiba
		"ec2280": 803,  // D-Link
		"ec233d": 3335, // Huawei
		"ec2368": 9242, // IntelliVoice
		"ec237b": 3032, // ZTE
		"ec2651": 545,  // Apple
		"ec26ca": 1595, // TP-Link
		"ec26fb": 9243, // Tecc
		"ec2a72": 102,  // Dell
		"ec2af0": 9244, // Ypsomed
		"ec2beb": 5439, // Amazon
		"ec2ce2": 545,  // Apple
		"ec2e98": 3005, // Azurewave
		"ec3091": 2,    // Cisco
		"ec316d": 9245, // Hansgrohe
		"ec354d": 4031, // Wingtech
		"ec3586": 545,  // Apple
		"ec363f": 9246, // Markov
		"ec3873": 826,  // Juniper
		"ec388f": 3335, // Huawei
		"ec3a52": 3335, // Huawei
		"ec3bf0": 9247, // NovelSat
		"ec3c88": 9248, // MCNEX
		"ec3cbb": 3335, // Huawei
		"ec3eb3": 2693, // ZyXEL
		"ec3ef7": 826,  // Juniper
		"ec4118": 6281, // XIAOMI
		"ec42b4": 9249, // ADC
		"ec438b": 9250, // Yaptv
		"ec43e6": 9251, // AWCER
		"ec43f6": 2693, // ZyXEL
		"ec4476": 2,    // Cisco
		"ec473c": 9252, // Redwire
		"ec4993": 9253, // Qihan
		"ec4d47": 3335, // Huawei
		"ec4f82": 374,  // Calix
		"ec50aa": 1685, // Aruba
		"ec5623": 3335, // Huawei
		"ec570d": 5108, // AFE
		"ec58ea": 2738, // Ruckus
		"ec59e7": 612,  // Microsoft
		"ec5a86": 3097, // Yulong
		"ec5c84": 2057, // Murata
		"ec6073": 1595, // TP-Link
		"ec60e0": 9254, // AVI-ON
		"ec63d7": 421,  // Intel
		"ec64e7": 9255, // MOCACARE
		"ec656e": 9256, // Things
		"ec65cc": 2154, // Panasonic
		"ec6c9a": 2640, // Arcadyan
		"ec6cb5": 3032, // ZTE
		"ec6f0b": 9257, // FADU
		"ec7097": 125,  // ARRIS
		"ec753e": 3335, // Huawei
		"ec75ed": 9258, // Citrix
		"ec7949": 4,    // Fujitsu
		"ec79f2": 9259, // Startel
		"ec7c2c": 3335, // Huawei
		"ec7c5c": 826,  // Juniper
		"ec7c74": 9260, // Justone
		"ec7cb6": 152,  // Samsung
		"ec7d11": 6370, // Vivo
		"ec7d9d": 9261, // CPI
		"ec7e91": 6282, // Itel
		"ec8009": 9262, // NovaSparks
		"ec8193": 4080, // Logitech
		"ec8263": 3032, // ZTE
		"ec8350": 612,  // Microsoft
		"ec836c": 9263, // RM
		"ec83d5": 9264, // GIRD
		"ec852f": 545,  // Apple
		"ec888f": 1595, // TP-Link
		"ec8892": 687,  // Motorola
		"ec8914": 3335, // Huawei
		"ec89f5": 2665, // Lenovo
		"ec8a4c": 3032, // ZTE
		"ec8ac4": 5439, // Amazon
		"ec8ac7": 4922, // Fiberhome
		"ec8c9a": 3335, // Huawei
		"ec8ca2": 2738, // Ruckus
		"ec8ead": 9265, // DLX
		"ec8eae": 9266, // Nagravision
		"ec8eb5": 302,  // HP
		"ec9365": 9267, // Mapper.ai
		"ec937d": 6393, // Technicolor
		"ec93ed": 9268, // DDoS-Guard
		"ec94cb": 6377, // Espressif
		"ec94d5": 826,  // Juniper
		"ec9a74": 302,  // HP
		"ec9b5b": 457,  // Nokia
		"ec9b8b": 302,  // HP
		"ec9bf3": 152,  // Samsung
		"eca1d1": 3335, // Huawei
		"eca29b": 9269, // Kemppi
		"eca81f": 6393, // Technicolor
		"eca86b": 1115, // Elitegroup
		"eca907": 545,  // Apple
		"eca940": 125,  // ARRIS
		"ecaa25": 152,  // Samsung
		"ecaaa0": 5440, // Pegatron
		"ecadb8": 545,  // Apple
		"ecade0": 803,  // D-Link
		"ecaf97": 9270, // GIT
		"ecb0e1": 4565, // Ciena
		"ecb106": 9271, // Acuro
		"ecb1d7": 302,  // HP
		"ecb4e8": 1592, // Wistron
		"ecb907": 9272, // CloudGenix
		"ecb970": 5441, // Ruijie
		"ecbafe": 9273, // Giroptic
		"ecbd09": 9274, // FUSION
		"ecbd1d": 2,    // Cisco
		"ecbe5f": 1449, // Vestel
		"ecbedd": 2049, // Sagemcom
		"ecc01b": 3335, // Huawei
		"ecc302": 531,  // HUMAX
		"ecc38a": 9275, // Accuenergy
		"ecc40d": 1427, // Nintendo
		"ecc5d2": 3335, // Huawei
		"ecc882": 2,    // Cisco
		"eccb30": 3335, // Huawei
		"ecce13": 2,    // Cisco
		"ecced7": 545,  // Apple
		"ecd00e": 9276, // MiraeRecognition
		"ecd09f": 5698, // Xiaomi
		"ecd925": 9277, // Rami
		"ecd950": 2619, // IRT
		"ecdb86": 9278, // API-K
		"ecde3d": 9279, // Lamprey
		"ecdf3a": 6370, // Vivo
		"ece09b": 152,  // Samsung
		"ece1a9": 2,    // Cisco
		"ece512": 9280, // tado
		"ece744": 9281, // Omntec
		"ece7a7": 421,  // Intel
		"ece915": 9282, // STI
		"ecebb8": 302,  // HP
		"ecf00e": 2557, // AboCom
		"ecf0fe": 3032, // ZTE
		"ecf22b": 6266, // Tecno
		"ecf236": 9283, // Neomontana
		"ecf35b": 457,  // Nokia
		"ecf40c": 2,    // Cisco
		"ecf451": 2640, // Arcadyan
		"ecf4bb": 102,  // Dell
		"ecfa03": 9284, // FCA
		"ecfaaa": 9285, // IMS
		"ecfabc": 6377, // Espressif
		"ecfaf4": 9286, // SenRa
		"ecfe7e": 9287, // BlueRadios
		"f0022b": 9288, // Chrontel
		"f00248": 9289, // SmarteBuilding
		"f0038c": 3005, // Azurewave
		"f008d1": 6377, // Espressif
		"f008f1": 152,  // Samsung
		"f00d5c": 6946, // JinQianMao
		"f00e1d": 6711, // Megafone
		"f00ebf": 9290, // ZettaHash
		"f00fec": 3335, // Huawei
		"f013c1": 9291, // Hannto
		"f015b9": 9292, // PlayFusion
		"f01628": 6393, // Technicolor
		"f0182b": 869,  // LG
		"f01898": 545,  // Apple
		"f01b6c": 6370, // Vivo
		"f01c13": 869,  // LG
		"f01c2d": 826,  // Juniper
		"f01d2d": 2,    // Cisco
		"f01dbc": 612,  // Microsoft
		"f01e34": 9293, // ORICO
		"f01faf": 102,  // Dell
		"f0219d": 1630, // Cal-Comp
		"f021e0": 5821, // eero
		"f0224e": 9294, // Esan
		"f023ae": 4499, // AMPAK
		"f02408": 65,   // Talaris
		"f02475": 545,  // Apple
		"f02572": 2,    // Cisco
		"f0258e": 3335, // Huawei
		"f025b7": 152,  // Samsung
		"f02624": 9295, // Wafa
		"f0264c": 9296, // Sigrist-Photometer
		"f0272d": 5439, // Amazon
		"f02745": 9297, // F-Secure
		"f02765": 2057, // Murata
		"f02929": 2,    // Cisco
		"f02a61": 9298, // Waldo
		"f02e51": 3171, // Casa
		"f02f4b": 545,  // Apple
		"f02f74": 1806, // ASUS
		"f02fa7": 3335, // Huawei
		"f02fd8": 9299, // Bi2-Vision
		"f0321a": 1933, // Mita-Teknik
		"f033e5": 3335, // Huawei
		"f03404": 6434, // TCT
		"f037a1": 9300, // Huike
		"f03965": 152,  // Samsung
		"f03a4b": 9301, // Bloombase
		"f03d03": 6266, // Tecno
		"f03d29": 9302, // Actility
		"f03e90": 2738, // Ruckus
		"f03f95": 3335, // Huawei
		"f0407b": 4922, // Fiberhome
		"f041c6": 9303, // Heat
		"f0421c": 421,  // Intel
		"f042f5": 3335, // Huawei
		"f04335": 9304, // DVN
		"f04347": 3335, // Huawei
		"f04a02": 2,    // Cisco
		"f04a2b": 9305, // PYRAMID
		"f04b3a": 826,  // Juniper
		"f04bf2": 9306, // JTECH
		"f04cd5": 5299, // Maxlinear
		"f04da2": 102,  // Dell
		"f04f7c": 5439, // Amazon
		"f05136": 6434, // TCT
		"f051ea": 6588, // Fitbit
		"f05501": 3335, // Huawei
		"f057a6": 421,  // Intel
		"f05849": 9307, // CareView
		"f05a09": 152,  // Samsung
		"f05b7b": 152,  // Samsung
		"f05c19": 1685, // Aruba
		"f05c77": 3522, // Google
		"f05cd5": 545,  // Apple
		"f05d89": 9308, // Dycon
		"f061c0": 1685, // Aruba
		"f063f9": 3335, // Huawei
		"f06426": 185,  // Extreme
		"f065ae": 152,  // Samsung
		"f065dd": 389,  // Primax
		"f06853": 5874, // Integrated
		"f06bca": 152,  // Samsung
		"f06c73": 457,  // Nokia
		"f06e0b": 612,  // Microsoft
		"f06f46": 9309, // Ubiik
		"f0704f": 152,  // Samsung
		"f0728c": 152,  // Samsung
		"f072ea": 3522, // Google
		"f07485": 9310, // NGD
		"f074e4": 9311, // Thundercomm
		"f0761c": 358,  // Compal
		"f0766f": 545,  // Apple
		"f07765": 9312, // Sourcefire
		"f077c3": 421,  // Intel
		"f077d0": 9313, // Xcellen
		"f07807": 545,  // Apple
		"f07816": 2,    // Cisco
		"f07959": 1806, // ASUS
		"f07960": 545,  // Apple
		"f07cc7": 826,  // Juniper
		"f07d68": 803,  // D-Link
		"f07f06": 2,    // Cisco
		"f08173": 5439, // Amazon
		"f08175": 2049, // Sagemcom
		"f08261": 2049, // Sagemcom
		"f084c9": 3032, // ZTE
		"f08620": 2640, // Arcadyan
		"f08a76": 152,  // Samsung
		"f08bfe": 9314, // Costel
		"f08cfb": 4922, // Fiberhome
		"f08edb": 9315, // VeloCloud
		"f0921c": 302,  // HP
		"f0933a": 9316, // NxtConect
		"f093c5": 9317, // Garland
		"f097e5": 9318, // Tamio
		"f09838": 3335, // Huawei
		"f0989d": 545,  // Apple
		"f09919": 797,  // Garmin
		"f099b6": 545,  // Apple
		"f099bf": 545,  // Apple
		"f09bb8": 3335, // Huawei
		"f09cbb": 9319, // RaonThink
		"f09ce9": 185,  // Extreme
		"f09e4a": 421,  // Intel
		"f09e63": 2,    // Cisco
		"f09fc2": 2979, // Ubiquiti
		"f09ffc": 3206, // Sharp
		"f0a225": 5439, // Amazon
		"f0a35a": 545,  // Apple
		"f0a764": 9320, // GST
		"f0a7b2": 9321, // Futaba
		"f0a951": 3335, // Huawei
		"f0a968": 6504, // Antailiye
		"f0aca4": 9322, // HBC-radiomatic
		"f0ad4e": 9323, // Globalscale
		"f0ae51": 9324, // Xi3
		"f0af85": 125,  // ARRIS
		"f0b022": 8975, // TOHO
		"f0b052": 2738, // Ruckus
		"f0b0e7": 545,  // Apple
		"f0b107": 306,  // Ericsson
		"f0b11d": 457,  // Nokia
		"f0b13f": 3335, // Huawei
		"f0b2b9": 421,  // Intel
		"f0b2e5": 2,    // Cisco
		"f0b31e": 3858, // Universal
		"f0b3ec": 545,  // Apple
		"f0b429": 5698, // Xiaomi
		"f0b479": 545,  // Apple
		"f0b4d2": 803,  // D-Link
		"f0b5b7": 8291, // Disruptive
		"f0b61e": 421,  // Intel
		"f0b6eb": 9325, // Poslab
		"f0b968": 6282, // Itel
		"f0bcc8": 9326, // MaxID
		"f0bcc9": 5462, // PFU
		"f0bdf1": 9327, // Sipod
		"f0bf97": 101,  // Sony
		"f0c1f1": 545,  // Apple
		"f0c371": 545,  // Apple
		"f0c42f": 3335, // Huawei
		"f0c478": 3335, // Huawei
		"f0c850": 3335, // Huawei
		"f0c88c": 9328, // LeddarTech
		"f0c8b5": 3335, // Huawei
		"f0cba1": 545,  // Apple
		"f0cd31": 152,  // Samsung
		"f0d08c": 6434, // TCT
		"f0d14f": 9329, // Linear
		"f0d1a9": 545,  // Apple
		"f0d1b8": 9330, // Ledvance
		"f0d2f1": 5439, // Amazon
		"f0d3a7": 9331, // CobaltRay
		"f0d3e7": 9332, // Sensometrix
		"f0d4e2": 102,  // Dell
		"f0d5bf": 421,  // Intel
		"f0d657": 9333, // Echosens
		"f0d7aa": 687,  // Motorola
		"f0d7dc": 9334, // Wesine
		"f0da7c": 9335, // RLH
		"f0db30": 9336, // Yottabyte
		"f0dbe2": 545,  // Apple
		"f0dbf8": 545,  // Apple
		"f0dce2": 545,  // Apple
		"f0def1": 1592, // Wistron
		"f0e4a2": 3335, // Huawei
		"f0e77e": 152,  // Samsung
		"f0ec39": 9337, // Essec
		"f0ee10": 152,  // Samsung
		"f0eebb": 9338, // VIPAR
		"f0ef86": 3522, // Google
		"f0f08f": 9339, // Nextek
		"f0f0a4": 5439, // Amazon
		"f0f249": 870,  // Hitron
		"f0f260": 9340, // Mobitec
		"f0f336": 1595, // TP-Link
		"f0f564": 152,  // Samsung
		"f0f5ae": 9341, // Adaptrum
		"f0f61c": 545,  // Apple
		"f0f6c1": 2048, // Sonos
		"f0f755": 2,    // Cisco
		"f0f7b3": 9342, // Phorm
		"f0f7e7": 3335, // Huawei
		"f0f842": 9343, // KEEBOX
		"f0f9f7": 9344, // IES
		"f0fac7": 3335, // Huawei
		"f0fcc8": 125,  // ARRIS
		"f0fda0": 9345, // Acurix
		"f0fee7": 3335, // Huawei
		"f40223": 3218, // PAX
		"f40228": 152,  // Samsung
		"f40270": 102,  // Dell
		"f40304": 3522, // Google
		"f4032a": 5439, // Amazon
		"f4032f": 9346, // Reduxio
		"f40343": 302,  // HP
		"f4044c": 9347, // ValenceTech
		"f40616": 545,  // Apple
		"f40669": 421,  // Intel
		"f4068d": 1637, // devolo
		"f409d8": 152,  // Samsung
		"f40a4a": 9348, // IndUSNET
		"f40b93": 2221, // BlackBerry
		"f40e01": 545,  // Apple
		"f40e22": 152,  // Samsung
		"f40e83": 125,  // ARRIS
		"f40f1b": 2,    // Cisco
		"f40f24": 545,  // Apple
		"f40f9b": 9349, // Wavelink
		"f412fa": 6377, // Espressif
		"f41535": 9350, // SPON
		"f41563": 289,  // F5
		"f419e2": 9351, // Volterra
		"f41ba1": 545,  // Apple
		"f41d6b": 3335, // Huawei
		"f41e26": 9352, // Simon-Kaloi
		"f41e5e": 9353, // RtBrick
		"f41f0b": 9354, // YAMABISHI
		"f41f88": 3032, // ZTE
		"f41fc2": 2,    // Cisco
		"f42012": 9355, // Cuciniale
		"f4239c": 6256, // Sernet
		"f42679": 421,  // Intel
		"f42833": 9356, // MMPC
		"f42853": 2128, // Zioncom
		"f42981": 6370, // Vivo
		"f42a7d": 1595, // TP-Link
		"f42b48": 9357, // Ubiqam
		"f42c56": 9358, // Senor
		"f42e7f": 1685, // Aruba
		"f4308b": 5698, // Xiaomi
		"f430b9": 302,  // HP
		"f431c3": 545,  // Apple
		"f434f0": 545,  // Apple
		"f437b7": 545,  // Apple
		"f438c1": 3335, // Huawei
		"f43909": 302,  // HP
		"f43d80": 9359, // FAG
		"f43e9d": 9360, // Benu
		"f44156": 9361, // Arrikto
		"f4419e": 3335, // Huawei
		"f4428f": 152,  // Samsung
		"f44450": 9362, // BND
		"f44588": 3335, // Huawei
		"f44637": 421,  // Intel
		"f44955": 9363, // MIMO
		"f449ef": 9364, // Emstone
		"f44c7f": 3335, // Huawei
		"f44d30": 1115, // Elitegroup
		"f44e05": 2,    // Cisco
		"f44e38": 9365, // Olibra
		"f44ee3": 421,  // Intel
		"f450eb": 7076, // Telechips
		"f45214": 432,  // Mellanox
		"f45595": 9366, // HENGBAO
		"f4559c": 3335, // Huawei
		"f4573e": 4922, // Fiberhome
		"f45c89": 545,  // Apple
		"f45f69": 9367, // Matsufu
		"f45ff7": 9368, // DQ
		"f4600d": 9369, // Panoptic
		"f460e2": 5698, // Xiaomi
		"f4631f": 3335, // Huawei
		"f46349": 9370, // Diffon
		"f4645d": 35,   // Toshiba
		"f465a6": 545,  // Apple
		"f46942": 2545, // Askey
		"f46abc": 9371, // Adonit
		"f46ad7": 612,  // Microsoft
		"f46bef": 2049, // Sagemcom
		"f46d04": 1806, // ASUS
		"f46d2f": 1595, // TP-Link
		"f46de2": 3032, // ZTE
		"f46e95": 185,  // Extreme
		"f46f4e": 9372, // Echowell
		"f46fed": 4922, // Fiberhome
		"f470ab": 6370, // Vivo
		"f47190": 152,  // Samsung
		"f47960": 3335, // Huawei
		"f47acc": 9373, // SolidFire
		"f47b09": 421,  // Intel
		"f47b5e": 152,  // Samsung
		"f47def": 152,  // Samsung
		"f47f35": 2,    // Cisco
		"f48139": 87,   // Canon
		"f483cd": 1595, // TP-Link
		"f4848d": 1595, // TP-Link
		"f485c6": 9374, // FDT
		"f48771": 9375, // Infoblox
		"f487c5": 3335, // Huawei
		"f48b32": 5698, // Xiaomi
		"f48c50": 421,  // Intel
		"f48ceb": 803,  // D-Link
		"f48e09": 457,  // Nokia
		"f48e38": 102,  // Dell
		"f48e92": 3335, // Huawei
		"f490ca": 9376, // Tensorcom
		"f492bf": 2979, // Ubiquiti
		"f49466": 9377, // CountMax
		"f49634": 421,  // Intel
		"f49651": 6083, // NAKAYO
		"f497c2": 9378, // Nebulon
		"f49c12": 9379, // Structab
		"f49f54": 152,  // Samsung
		"f49ff3": 3335, // Huawei
		"f4a475": 421,  // Intel
		"f4a4d6": 3335, // Huawei
		"f4a52a": 9380, // Hawa
		"f4a59d": 3335, // Huawei
		"f4a739": 826,  // Juniper
		"f4a80d": 1592, // Wistron
		"f4a997": 87,   // Canon
		"f4acc1": 2,    // Cisco
		"f4afe7": 545,  // Apple
		"f4b19c": 7032, // AltoBeam
		"f4b301": 421,  // Intel
		"f4b381": 9381, // WindowMaster
		"f4b52f": 826,  // Juniper
		"f4b5aa": 3032, // ZTE
		"f4b5bb": 1490, // Ceragon
		"f4b688": 542,  // Plantronics
		"f4b6e5": 9382, // TerraSem
		"f4b78d": 3335, // Huawei
		"f4b7b3": 6370, // Vivo
		"f4b8a7": 3032, // ZTE
		"f4bd9e": 2,    // Cisco
		"f4beec": 545,  // Apple
		"f4bf80": 3335, // Huawei
		"f4bfa8": 826,  // Juniper
		"f4c02f": 9383, // BlueBite
		"f4c114": 6393, // Technicolor
		"f4c248": 152,  // Samsung
		"f4c447": 9384, // Coagent
		"f4c6d7": 9385, // blackned
		"f4c714": 3335, // Huawei
		"f4c795": 9386, // WEY
		"f4c7c8": 9387, // Kelvin
		"f4ca24": 9388, // FreeBit
		"f4cb52": 3335, // Huawei
		"f4cc55": 826,  // Juniper
		"f4ce23": 421,  // Intel
		"f4ce46": 302,  // HP
		"f4ce48": 185,  // Extreme
		"f4cfa2": 6377, // Espressif
		"f4cfe2": 2,    // Cisco
		"f4d108": 421,  // Intel
		"f4d261": 9389, // SEMOCON
		"f4d488": 545,  // Apple
		"f4d9c6": 3506, // Unionman
		"f4d9fb": 152,  // Samsung
		"f4dbe3": 545,  // Apple
		"f4dbe6": 2,    // Cisco
		"f4dcf9": 3335, // Huawei
		"f4dd9e": 6242, // GoPro
		"f4de0c": 9390, // ESPOD
		"f4deaf": 3335, // Huawei
		"f4e2c6": 2979, // Ubiquiti
		"f4e3fb": 3335, // Huawei
		"f4e451": 3335, // Huawei
		"f4e4ad": 3032, // ZTE
		"f4e5f2": 3335, // Huawei
		"f4e9d4": 2026, // QLogic
		"f4ea67": 2,    // Cisco
		"f4eab5": 185,  // Extreme
		"f4eb38": 2049, // Sagemcom
		"f4ec38": 1595, // TP-Link
		"f4ee08": 102,  // Dell
		"f4ee14": 4254, // Mercury
		"f4f15a": 545,  // Apple
		"f4f197": 9391, // EMTAKE
		"f4f1e1": 687,  // Motorola
		"f4f26d": 1595, // TP-Link
		"f4f309": 152,  // Samsung
		"f4f3aa": 9392, // JBL
		"f4f524": 687,  // Motorola
		"f4f5a5": 457,  // Nokia
		"f4f5d8": 3522, // Google
		"f4f5db": 5698, // Xiaomi
		"f4f5e8": 3522, // Google
		"f4f646": 9393, // Dediprog
		"f4f647": 3032, // ZTE
		"f4f951": 545,  // Apple
		"f4fbb8": 3335, // Huawei
		"f4fcb1": 9394, // JJ
		"f4fd2b": 9395, // ZOYI
		"f4fefb": 152,  // Samsung
		"f800a1": 3335, // Huawei
		"f80113": 3335, // Huawei
		"f80332": 9396, // Khomp
		"f80377": 545,  // Apple
		"f8042e": 152,  // Samsung
		"f8084f": 2049, // Sagemcom
		"f80bbe": 125,  // ARRIS
		"f80bcb": 2,    // Cisco
		"f80cf3": 869,  // LG
		"f80d60": 87,   // Canon
		"f80dac": 302,  // HP
		"f80dea": 9397, // ZyCast
		"f80df0": 3032, // ZTE
		"f80df1": 9398, // Sontex
		"f80f41": 1592, // Wistron
		"f80f6f": 2,    // Cisco
		"f80ff9": 3522, // Google
		"f81037": 9399, // Atopia
		"f81093": 545,  // Apple
		"f81308": 457,  // Nokia
		"f814fe": 3506, // Unionman
		"f81547": 620,  // Avaya
		"f81654": 421,  // Intel
		"f81897": 1939, // 2Wire
		"f81a2b": 3522, // Google
		"f81a67": 1595, // TP-Link
		"f81d0f": 870,  // Hitron
		"f81d90": 9400, // Solidwintech
		"f81edf": 545,  // Apple
		"f81f32": 687,  // Motorola
		"f820a9": 3335, // Huawei
		"f82285": 4781, // Cypress
		"f823b2": 3335, // Huawei
		"f82441": 9401, // Yeelink
		"f82551": 46,   // Epson
		"f8272e": 9402, // Mercku
		"f82793": 545,  // Apple
		"f82819": 2874, // Liteon
		"f828c9": 3335, // Huawei
		"f829c0": 6905, // Availink
		"f82c18": 1939, // 2Wire
		"f82d7c": 545,  // Apple
		"f82dc0": 125,  // ARRIS
		"f82e3f": 3335, // Huawei
		"f82edb": 9403, // RTW
		"f82f5b": 9404, // eGauge
		"f82f65": 3335, // Huawei
		"f82f6a": 6282, // Itel
		"f8313e": 9405, // endeavour
		"f832e4": 1806, // ASUS
		"f83441": 421,  // Intel
		"f83553": 9406, // Magenta
		"f835dd": 1450, // Gemtek
		"f83869": 869,  // LG
		"f83880": 545,  // Apple
		"f83b1d": 6393, // Technicolor
		"f83b7e": 3335, // Huawei
		"f83cbf": 9407, // Botato
		"f83dff": 3335, // Huawei
		"f83e95": 3335, // Huawei
		"f83f51": 152,  // Samsung
		"f845ad": 3533, // Konka
		"f8461c": 101,  // Sony
		"f8462d": 9408, // SYNTEC
		"f84897": 89,   // Hitachi
		"f84a73": 9409, // Eumtech
		"f84a7f": 9410, // Innometriks
		"f84abf": 3335, // Huawei
		"f84cda": 3335, // Huawei
		"f84d33": 4922, // Fiberhome
		"f84d89": 545,  // Apple
		"f84e17": 101,  // Sony
		"f84e58": 152,  // Samsung
		"f84e73": 545,  // Apple
		"f84f57": 2,    // Cisco
		"f85063": 9411, // Verathon
		"f85128": 9412, // SimpliSafe
		"f8516d": 9413, // Denwa
		"f852df": 9414, // VNL
		"f85329": 3335, // Huawei
		"f854b8": 5439, // Amazon
		"f855cd": 1399, // Visteon
		"f856c3": 3032, // ZTE
		"f85971": 421,  // Intel
		"f85a00": 9415, // Sanford
		"f85b3b": 2545, // Askey
		"f85b9c": 9416, // SB
		"f85bc9": 9417, // M-Cube
		"f85c4d": 457,  // Nokia
		"f85e42": 6393, // Technicolor
		"f85ea0": 421,  // Intel
		"f85f2a": 457,  // Nokia
		"f860f0": 1685, // Aruba
		"f86214": 545,  // Apple
		"f862aa": 9418, // xn
		"f8633f": 421,  // Intel
		"f863d9": 125,  // ARRIS
		"f864b8": 3032, // ZTE
		"f8665a": 545,  // Apple
		"f866f2": 2,    // Cisco
		"f86bd9": 2,    // Cisco
		"f86d73": 9419, // Zengge
		"f86ecf": 9420, // Arcx
		"f86eee": 3335, // Huawei
		"f86fc1": 545,  // Apple
		"f872ea": 2,    // Cisco
		"f87394": 1368, // Netgear
		"f873a2": 620,  // Avaya
		"f87588": 3335, // Huawei
		"f875a4": 4920, // LCFC
		"f8769b": 9421, // Neopis
		"f877b8": 152,  // Samsung
		"f8790a": 125,  // ARRIS
		"f87a41": 2,    // Cisco
		"f87aef": 9422, // Rosonix
		"f87b20": 2,    // Cisco
		"f87b7a": 125,  // ARRIS
		"f87fa5": 9423, // Greatek
		"f8811a": 9424, // Overkiz
		"f88200": 9425, // CaptionCall
		"f88479": 7724, // Yaojin
		"f884f2": 152,  // Samsung
		"f885f9": 374,  // Calix
		"f887f1": 545,  // Apple
		"f88b37": 125,  // ARRIS
		"f88c1c": 9426, // Kaishun
		"f88c21": 1595, // TP-Link
		"f88def": 9427, // Tenebraex
		"f88e85": 3870, // Comtrend
		"f88ea1": 6295, // Edgecore
		"f88f07": 152,  // Samsung
		"f88fca": 3522, // Google
		"f89066": 9428, // Nain
		"f893f3": 9429, // Volans
		"f894c2": 421,  // Intel
		"f89522": 3335, // Huawei
		"f895c7": 869,  // LG
		"f895ea": 545,  // Apple
		"f89753": 3335, // Huawei
		"f897a9": 306,  // Ericsson
		"f897cf": 9430, // Daeshin-Information
		"f8983a": 9431, // Leeman
		"f898b9": 3335, // Huawei
		"f898ef": 3335, // Huawei
		"f89955": 9432, // Fortress
		"f89a78": 3335, // Huawei
		"f89d0d": 1999, // Control
		"f89dbb": 9433, // Tintri
		"f89e94": 421,  // Intel
		"f8a03d": 9434, // Dinstar
		"f8a097": 125,  // ARRIS
		"f8a26d": 87,   // Canon
		"f8a2d6": 2874, // Liteon
		"f8a34f": 3032, // ZTE
		"f8a45f": 5698, // Xiaomi
		"f8a5c5": 2,    // Cisco
		"f8a73a": 2,    // Cisco
		"f8a91f": 9435, // ZVISION
		"f8a963": 358,  // Compal
		"f8a9d0": 869,  // LG
		"f8aa3f": 7447, // DWnet
		"f8aa8a": 9436, // Axview
		"f8ab05": 2049, // Sagemcom
		"f8ac65": 421,  // Intel
		"f8ac6d": 9437, // Deltenna
		"f8af05": 3335, // Huawei
		"f8afdb": 4922, // Fiberhome
		"f8b132": 3335, // Huawei
		"f8b156": 102,  // Dell
		"f8b1dd": 545,  // Apple
		"f8b46a": 302,  // HP
		"f8b54d": 421,  // Intel
		"f8b7e2": 2,    // Cisco
		"f8b95a": 869,  // LG
		"f8bae6": 457,  // Nokia
		"f8bbbf": 5821, // eero
		"f8bc0e": 5821, // eero
		"f8bc12": 102,  // Dell
		"f8bc41": 9438, // Rosslare
		"f8be0d": 9439, // A2UICT
		"f8bf09": 3335, // Huawei
		"f8c001": 826,  // Juniper
		"f8c091": 9440, // Highgates
		"f8c116": 826,  // Juniper
		"f8c249": 67,   // Private
		"f8c288": 2,    // Cisco
		"f8c397": 9441, // NZXT
		"f8c39e": 3335, // Huawei
		"f8c3cc": 545,  // Apple
		"f8c678": 9442, // Carefusion
		"f8c96c": 4922, // Fiberhome
		"f8ca85": 48,   // NEC
		"f8cab8": 102,  // Dell
		"f8cc6e": 9443, // DEPO
		"f8ce72": 1592, // Wistron
		"f8cfc5": 687,  // Motorola
		"f8d027": 46,   // Epson
		"f8d0ac": 101,  // Sony
		"f8d0bd": 152,  // Samsung
		"f8d111": 1595, // TP-Link
		"f8d3a9": 9444, // AXAN
		"f8d478": 833,  // Flextronics
		"f8dadf": 9445, // EcoTech
		"f8dae2": 9446, // NDC
		"f8db4c": 9447, // PNY
		"f8db7f": 1341, // HTC
		"f8db88": 102,  // Dell
		"f8dc7a": 9448, // Variscite
		"f8df15": 3940, // Sunitec
		"f8dfa8": 3032, // ZTE
		"f8dfe1": 9449, // MyLight
		"f8e079": 687,  // Motorola
		"f8e43b": 6723, // ASIX
		"f8e44e": 9450, // Mcot
		"f8e4e3": 421,  // Intel
		"f8e4fb": 2247, // Actiontec
		"f8e57e": 2,    // Cisco
		"f8e61a": 152,  // Samsung
		"f8e71e": 2738, // Ruckus
		"f8e7a0": 6370, // Vivo
		"f8e811": 3335, // Huawei
		"f8e903": 803,  // D-Link
		"f8e94e": 545,  // Apple
		"f8eda5": 125,  // ARRIS
		"f8f014": 9451, // RackWare
		"f8f082": 8149, // Nagtech
		"f8f1b6": 687,  // Motorola
		"f8f1e6": 152,  // Samsung
		"f8f21e": 421,  // Intel
		"f8f25a": 9452, // G-Lab
		"f8f532": 125,  // ARRIS
		"f8f7b9": 3335, // Huawei
		"f8f7d3": 639,  // International
		"f8f7ff": 9453, // SYN-Tech
		"f8fb2f": 9454, // Santur
		"f8fce1": 5439, // Amazon
		"f8fe5c": 9455, // Reciprocal
		"f8ff0b": 3691, // Electronic
		"f8ff5f": 8626, // Shenzhen
		"f8ffc2": 545,  // Apple
		"fc0012": 35,   // Toshiba
		"fc019e": 9456, // Vievu
		"fc0296": 5698, // Xiaomi
		"fc039f": 152,  // Samsung
		"fc0647": 9457, // Cortland
		"fc06ed": 5818, // M2Motive
		"fc084a": 4,    // Fujitsu
		"fc0a81": 185,  // Extreme
		"fc0fe6": 101,  // Sony
		"fc0fe7": 706,  // Microchip
		"fc1186": 9458, // Logic3
		"fc1193": 3335, // Huawei
		"fc122c": 3335, // Huawei
		"fc15b4": 302,  // HP
		"fc1607": 9459, // Taian
		"fc1794": 9460, // InterCreative
		"fc183c": 545,  // Apple
		"fc1910": 152,  // Samsung
		"fc1999": 5698, // Xiaomi
		"fc1a11": 6370, // Vivo
		"fc1bd1": 3335, // Huawei
		"fc1bff": 9461, // V-ZUG
		"fc1ca1": 457,  // Nokia
		"fc1d2a": 6370, // Vivo
		"fc1d43": 545,  // Apple
		"fc1d84": 9462, // Autobase
		"fc1e16": 9463, // IPEVO
		"fc1f19": 152,  // Samsung
		"fc1fc0": 9464, // Eurecam
		"fc2325": 9465, // EosTek
		"fc253f": 545,  // Apple
		"fc29e3": 6296, // Infinix
		"fc29f3": 9466, // McPay
		"fc2a9c": 545,  // Apple
		"fc2bb2": 2247, // Actiontec
		"fc2d5e": 3032, // ZTE
		"fc2e2d": 9467, // Lorom
		"fc2f40": 7846, // Calxeda
		"fc2f6b": 9468, // Everspin
		"fc2faa": 457,  // Nokia
		"fc2fef": 6516, // UTT
		"fc3342": 826,  // Juniper
		"fc335f": 9469, // Polyera
		"fc3497": 1806, // ASUS
		"fc3598": 9470, // Favite
		"fc35e6": 1399, // Visteon
		"fc3964": 6282, // Itel
		"fc3ce9": 9471, // Tsingtong
		"fc3da5": 2640, // Arcadyan
		"fc3f7c": 3335, // Huawei
		"fc3fdb": 302,  // HP
		"fc4009": 3032, // ZTE
		"fc4203": 152,  // Samsung
		"fc4482": 421,  // Intel
		"fc449f": 3032, // ZTE
		"fc4596": 358,  // Compal
		"fc48ef": 3335, // Huawei
		"fc492d": 5439, // Amazon
		"fc4ae9": 3798, // Castlenet
		"fc4bbc": 2427, // Sunplus
		"fc4da6": 3335, // Huawei
		"fc4ea4": 545,  // Apple
		"fc51a4": 125,  // ARRIS
		"fc528d": 6393, // Technicolor
		"fc589a": 2,    // Cisco
		"fc58df": 9472, // Interphone
		"fc5a1d": 870,  // Hitron
		"fc5b26": 9473, // MikroBits
		"fc5b39": 2,    // Cisco
		"fc5c45": 2738, // Ruckus
		"fc61e9": 4922, // Fiberhome
		"fc62b9": 430,  // Alpsalpine
		"fc643a": 152,  // Samsung
		"fc64ba": 5698, // Xiaomi
		"fc65b3": 3335, // Huawei
		"fc65de": 5439, // Amazon
		"fc66cf": 545,  // Apple
		"fc6c31": 9474, // LXinstruments
		"fc6dc0": 9475, // BME
		"fc6dd1": 5112, // APRESIA
		"fc6fb7": 125,  // ARRIS
		"fc71fa": 2657, // Trane
		"fc73fb": 3335, // Huawei
		"fc7516": 803,  // D-Link
		"fc75e6": 3622, // Handreamnet
		"fc7692": 4330, // Semptian
		"fc7774": 421,  // Intel
		"fc777b": 870,  // Hitron
		"fc7c02": 6849, // Phicomm
		"fc7ff1": 1685, // Aruba
		"fc8399": 620,  // Avaya
		"fc83c6": 9476, // N-Radio
		"fc8596": 9477, // Axonne
		"fc862a": 3335, // Huawei
		"fc8743": 3335, // Huawei
		"fc8a3d": 3032, // ZTE
		"fc8d3d": 9478, // Leapfive
		"fc8e6e": 9479, // StreamCCTV
		"fc8e7e": 125,  // ARRIS
		"fc8f90": 152,  // Samsung
		"fc8fc4": 1861, // Intelligent
		"fc90fa": 9480, // Independent
		"fc9114": 6393, // Technicolor
		"fc923b": 457,  // Nokia
		"fc9257": 5696, // Renesas
		"fc9435": 3335, // Huawei
		"fc946c": 9481, // Ubivelox
		"fc94ce": 3032, // ZTE
		"fc94e3": 6393, // Technicolor
		"fc956a": 4176, // Octagon
		"fc9643": 826,  // Juniper
		"fc9947": 2,    // Cisco
		"fc9bc6": 4655, // Sumavision
		"fc9bd4": 9482, // EdgeQ
		"fc9c98": 8407, // Arlo
		"fc9fae": 9483, // Fidus
		"fc9fe1": 9484, // CoNWIN.Tech
		"fca13e": 152,  // Samsung
		"fca183": 5439, // Amazon
		"fca621": 152,  // Samsung
		"fca667": 5439, // Amazon
		"fca6cd": 4922, // Fiberhome
		"fca841": 620,  // Avaya
		"fca84a": 9485, // Sentinum
		"fca89a": 3940, // Sunitec
		"fca9b0": 9486, // Miartech
		"fca9dc": 5696, // Renesas
		"fcaa14": 1929, // Giga-Byte
		"fcaa81": 545,  // Apple
		"fcaab6": 152,  // Samsung
		"fcab90": 3335, // Huawei
		"fcad0f": 9487, // QTS
		"fcae34": 125,  // ARRIS
		"fcaf6a": 9488, // Qulsar
		"fcafac": 9489, // Socionext
		"fcafbe": 9490, // TireCheck
		"fcb3bc": 421,  // Intel
		"fcb4e6": 2545, // Askey
		"fcb58a": 9491, // Wapice
		"fcb662": 9492, // IC
		"fcb698": 1640, // Cambridge
		"fcb6d8": 545,  // Apple
		"fcbc9c": 9493, // Vimar
		"fcbcd1": 3335, // Huawei
		"fcbd67": 3794, // Arista
		"fcbe7b": 6370, // Vivo
		"fcc233": 1806, // ASUS
		"fcc23d": 633,  // Atmel
		"fcc2de": 2057, // Murata
		"fcc734": 152,  // Samsung
		"fcc897": 3032, // ZTE
		"fccac4": 9494, // LifeHealth
		"fccce4": 9495, // Ascon
		"fccf62": 372,  // IBM
		"fcd436": 687,  // Motorola
		"fcd6bd": 1926, // Bosch
		"fcd733": 1595, // TP-Link
		"fcd848": 545,  // Apple
		"fcd908": 5698, // Xiaomi
		"fcdb21": 9496, // Samsara
		"fcdb96": 9497, // Enervalley
		"fcdbb3": 2057, // Murata
		"fcdc4a": 9498, // G-Wearables
		"fcde90": 152,  // Samsung
		"fce186": 9499, // A3M
		"fce1fb": 5956, // Array
		"fce26c": 545,  // Apple
		"fce33c": 3335, // Huawei
		"fce557": 457,  // Nokia
		"fce66a": 6571, // Industrial
		"fce806": 6446, // Edifier
		"fce998": 545,  // Apple
		"fcecda": 2979, // Ubiquiti
		"fcedb9": 9500, // Arrayent
		"fceee6": 9501, // Formike
		"fcf136": 152,  // Samsung
		"fcf152": 101,  // Sony
		"fcf1cd": 9502, // Optex-FA
		"fcf528": 2693, // ZyXEL
		"fcf5c4": 6377, // Espressif
		"fcf647": 4922, // Fiberhome
		"fcf77b": 3335, // Huawei
		"fcf8ae": 421,  // Intel
		"fcf8b7": 9503, // TRONTEQ
		"fcfbfb": 2,    // Cisco
		"fcfc48": 545,  // Apple
	},
	vendors: map[uint]string{
		7748: "01DB-Metravib",
		8822: "100fio",
		6310: "10NET/DCA",
		8467: "12Sided",
		3448: "16063",
		6628: "1MORE",
		4984: "1Net",
		2516: "2001",
		4786: "21168",
		2732: "27M",
		1094: "29530",
		7063: "2CRSI",
		3128: "2EI",
		8688: "2GIG",
		3689: "2M",
		1939: "2Wire",
		6565: "2bps",
		2500: "2wcom",
		2236: "2xWireless",
		3788: "30805",
		800:  "360",
		6845: "3ALogics",
		5336: "3CX",
		160:  "3Com",
		4375: "3DSP",
		5780: "3H",
		1421: "3J",
		3124: "3Leaf",
		6308: "3M",
		1294: "3UP",
		3580: "3Way",
		6682: "3bumen.com",
		8930: "3isysnetworks",
		4441: "3onedata",
		7483: "3pleplay",
		3244: "3soft",
		7799: "3view",
		766:  "3ware",
		2774: "4Access",
		6546: "4DReplay",
		4149: "4IPNET",
		3330: "4NSYS",
		4467: "4RF",
		9004: "5VT",
		1428: "6WIND",
		8552: "6harmonics",
		7167: "70mai",
		7035: "7HUGS",
		6261: "7INOVA",
		8679: "7signal",
		3730: "802automation",
		6598: "8D",
		8821: "8Devices",
		7918: "8Locations",
		6474: "8Mesh",
		5651: "8X8",
		8701: "9Solutions",
		899:  "@Track",
		1199: "@pos.com",
		1331: "A&D",
		8763: "A&R",
		7104: "A&T",
		2974: "A-First",
		3543: "A-Four",
		3585: "A-Link",
		537:  "A-One",
		4327: "A-Team",
		6224: "A-Trend",
		4151: "A-Trust",
		1127: "A-Z",
		6995: "A.D.C",
		8818: "A.N",
		6156: "A.T.N.R",
		4128: "A10",
		2320: "A2",
		4445: "A2B",
		9439: "A2UICT",
		9499: "A3M",
		4396: "A4SP",
		1009: "A5TEK",
		4675: "A7",
		2127: "AAC",
		5030: "AAE",
		1069: "AAEON",
		3270: "AAI",
		5620: "AANetcom",
		21:   "ABB",
		1961: "ABB./Tropos",
		2731: "ABB/Totalflow",
		7475: "ABC",
		8405: "ABIsystems",
		4442: "ABK",
		3553: "ACA-Digital",
		2477: "ACARD",
		1807: "ACC",
		7952: "ACES",
		8793: "ACK",
		1397: "ACKSYS",
		1799: "ACM",
		616:  "ACN",
		3432: "ACOGITO",
		443:  "ACR",
		8309: "ACT",
		491:  "ACT'L",
		8423: "ACTIVIO",
		8153: "ACTP",
		2214: "ACTi",
		7761: "ACX",
		5287: "AD",
		7779: "ADATA",
		6492: "ADB",
		9249: "ADC",
		4671: "ADC-Elektronik",
		6839: "ADD-Engineering",
		2204: "ADDI-DATA",
		2220: "ADDO-Japan",
		1254: "ADI",
		1653: "ADInstruments",
		5281: "ADS",
		5428: "ADVANCED",
		3356: "ADigit",
		2035: "AERAS",
		7088: "AES",
		8265: "AEV",
		1915: "AEWIN",
		1958: "AFAR",
		5108: "AFE",
		4438: "AFREEY",
		232:  "AG-E",
		8879: "AGAiT",
		1354: "AGFEO",
		9044: "AHN",
		1030: "AIDONIC",
		6389: "AIM",
		1905: "AINTech",
		3933: "AIOI",
		4088: "AIR802",
		2546: "AIRAYA",
		9205: "AIRSOUND",
		6447: "AITelecom",
		6225: "AK-Systems",
		7017: "AKELA",
		2654: "ALAXALA",
		3992: "ALCOMA",
		5527: "ALE",
		2747: "ALGOSYSTEM",
		8318: "ALT",
		9190: "ALTOGRAPHICS",
		5682: "ALi",
		469:  "AM",
		226:  "AMCC",
		13:   "AMD",
		8104: "AMERGINT",
		2223: "AMETEK",
		7395: "AMG",
		6623: "AMICCOM",
		5197: "AMIT",
		2791: "AMOD",
		4499: "AMPAK",
		5334: "AMTEC",
		5064: "ANI",
		3361: "ANSA",
		664:  "ANTARA.net",
		4654: "AOC",
		7535: "AOD",
		1872: "AOS",
		234:  "AOpen",
		5933: "APC",
		8947: "APCON",
		3220: "APD",
		4257: "API",
		9278: "API-K",
		1913: "APLUX",
		6525: "APR",
		5112: "APRESIA",
		8054: "APS",
		5474: "APT",
		8938: "AQ",
		4508: "ARBURG",
		6573: "ARCHEAN",
		8743: "AREC",
		1957: "ARIMA",
		2025: "ARION",
		276:  "ARK",
		2283: "ARKUS",
		462:  "ARM",
		2389: "ARMA",
		673:  "ARMITEL",
		6193: "ARN",
		125:  "ARRIS",
		2475: "ARTDIO",
		2562: "ARUZE",
		283:  "ARtem",
		1274: "ASA",
		854:  "ASB",
		9077: "ASD",
		726:  "ASE",
		2646: "ASEM",
		3532: "ASI",
		2012: "ASIP",
		6723: "ASIX",
		8381: "ASMedia",
		2872: "ASP",
		8747: "ASR",
		4720: "ASRock",
		7064: "ASSMANN",
		5073: "AST",
		4712: "ASTAK",
		3408: "ASTEL",
		1806: "ASUS",
		1057: "AT&T",
		8502: "ATAW",
		7995: "ATCOM",
		6992: "ATEK",
		1011: "ATI",
		5122: "ATM",
		3210: "ATMedia",
		6940: "ATN",
		1296: "ATO",
		6239: "ATOCS",
		7962: "ATOM",
		2831: "ATOMIC",
		6542: "ATP",
		3647: "ATRON",
		2357: "ATTO",
		8492: "ATW",
		2772: "ATX",
		4854: "AUTOVISION",
		2046: "AV",
		3119: "AVC",
		8360: "AVI",
		9254: "AVI-ON",
		6105: "AVIDIA",
		163:  "AVLAB",
		621:  "AVM",
		4589: "AVSystem",
		8400: "AVTrace",
		9251: "AWCER",
		2595: "AWIND",
		6971: "AWare",
		9444: "AXAN",
		1894: "AXELL",
		7618: "AXERRA",
		7410: "AXIM",
		8396: "AXPRO",
		5052: "AZS",
		4067: "AZTEQ",
		9182: "Aartesys",
		327:  "Aatr",
		7760: "Aava",
		1749: "Abatron",
		2656: "Abbey",
		6720: "Abeeway",
		907:  "Abeona",
		6402: "Abicom",
		6745: "Ability",
		5256: "Abit",
		4816: "Ablaze",
		6307: "Able",
		8034: "Ablelink",
		6216: "Abler",
		2216: "Ablerex",
		7556: "Abloomy",
		2557: "AboCom",
		7352: "Aboundi",
		1888: "AboveCable",
		6539: "Abrantix",
		825:  "AbsoluteValue",
		4904: "AcSiP",
		5768: "Acacia",
		1039: "AcceLight",
		3004: "Accedian",
		5210: "Accel",
		6933: "AccelStor",
		643:  "Accelent",
		2306: "Accelerated",
		7699: "Accelink",
		6348: "Accell",
		1014: "Accella",
		3211: "Accense",
		9076: "Accensus",
		5472: "Access",
		2311: "Accesslan",
		5408: "Acclaim",
		4083: "Accm",
		3510: "Acconet",
		5175: "Accord",
		8717: "Accordance",
		376:  "Accordion",
		2883: "Accsense",
		145:  "Accton",
		2017: "Accu-Sort",
		5376: "Accu-Time",
		6666: "AccuSpec",
		9275: "Accuenergy",
		8312: "Accupix",
		908:  "Accusys",
		8304: "Accutome",
		6518: "Ace",
		3250: "AceNet",
		857:  "Aceex",
		9007: "Acentic",
		143:  "Acer",
		6091: "Acetel",
		6167: "Ackfin",
		7712: "Aclima",
		7133: "Acoinfo",
		527:  "Acomz",
		107:  "Acorn",
		1231: "Acorp",
		820:  "Acqis",
		6026: "Acrison",
		277:  "Acromag",
		6290: "Acroname",
		759:  "Acronet",
		424:  "Acrosser",
		655:  "Acrowave",
		1735: "Actel",
		540:  "Actelis",
		1052: "Acterna",
		2294: "Actia",
		7026: "Actifio",
		9302: "Actility",
		7351: "Actineon",
		3633: "Action",
		8368: "Actioncable",
		2247: "Actiontec",
		4182: "Actis",
		2550: "ActivNetworks",
		294:  "Activetelco",
		7220: "Actlas",
		2498: "Actuality",
		7514: "Acubit",
		5605: "Acucomm",
		335:  "Aculab",
		7854: "Acumen",
		9345: "Acurix",
		9271: "Acuro",
		6090: "Acute",
		4915: "Adachi-Syokai",
		8289: "Adafruit",
		8778: "Adamson",
		6620: "Adanis",
		396:  "Adapcom",
		6517: "Adapt-IP",
		2778: "Adapt4",
		134:  "Adaptec",
		6245: "Adapteva",
		5468: "Adaptive",
		2596: "Adaptix",
		9341: "Adaptrum",
		4177: "Adastra",
		6496: "Adatis",
		1051: "Adax",
		8442: "Adaxys",
		2519: "Aday",
		2730: "Add-On",
		7216: "AddOn",
		411:  "AddPac",
		8074: "Addenergie",
		2193: "Adder",
		8753: "Additech",
		2617: "Addlogix",
		5282: "Addonics",
		5085: "Addtron",
		1665: "Addvalue",
		1179: "Adept",
		8365: "Adero",
		4361: "Adesys",
		2377: "Adhoc",
		3383: "Adhoco",
		4991: "Adid",
		2126: "Adimos",
		2848: "Aditec",
		4978: "Adlink",
		6166: "Admtek",
		7266: "Adnacom",
		5070: "Adobe",
		9371: "Adonit",
		72:   "Adra",
		5491: "Adsoft",
		342:  "Adtec",
		4987: "Adtech",
		202:  "Adtran",
		4910: "Adura",
		7467: "Advan",
		6122: "AdvanSys",
		587:  "Advanced",
		9216: "Advanced-Connectek",
		2343: "Advanet",
		6371: "Advansee",
		3403: "Advansus",
		7384: "Advantage",
		1703: "Advantech",
		7754: "Advas",
		987:  "Advent",
		4054: "Adventiq",
		7891: "Advidia",
		3087: "Aegate",
		4264: "Aehr",
		1622: "Aeluros",
		5071: "Aeon",
		8653: "Aeonsemi",
		1343: "AeroConcierge",
		3625: "AeroVIronment",
		5234: "Aerocomm",
		6851: "Aerodev",
		7709: "Aerodisk",
		5213: "Aeroflex",
		1721: "Aeronix",
		1863: "Aeroscout",
		2496: "Aerotech",
		2226: "Aerotelecom",
		539:  "Aeta",
		4729: "Aetas",
		7670: "Aetek",
		8358: "Aetheris",
		6646: "Aetheros",
		2828: "Aevoe",
		1944: "Afco",
		7424: "Affinegy",
		7865: "Affirmed",
		7839: "AgJunction",
		6436: "AgLogica",
		3547: "Agami",
		3865: "Agapha",
		8633: "Agatel",
		345:  "Agere",
		6340: "Agfa",
		1610: "Agile",
		7999: "AgileMesh",
		653:  "Agilent",
		5493: "Agilis",
		5660: "Agranat",
		3729: "AhnLab",
		7317: "AiNET",
		2959: "AiZen",
		8921: "Aiconn",
		7468: "Aidc",
		7202: "Aidon",
		8914: "Aifa",
		1844: "Ailocom",
		3535: "Aipermon",
		1702: "Aiphone",
		6244: "Aipon",
		1551: "Aiptek",
		4002: "Air2App",
		1611: "Air2U",
		8403: "AirCUVE",
		3090: "AirDefense",
		5291: "AirFiber",
		1718: "AirFlow",
		2829: "AirLink",
		1621: "AirLogic",
		1983: "AirMagnet",
		6584: "AirNav",
		978:  "AirRunner",
		7555: "AirScape",
		216:  "AirSwitch",
		1594: "AirVast",
		2921: "Airak",
		3428: "AireSpider",
		6455: "Airenetworks",
		3742: "Airgain",
		1599: "Airgo",
		774:  "Airocon",
		6466: "Airoha",
		8222: "Airpo",
		7074: "Airsonics",
		259:  "Airspan",
		4430: "Airtech",
		860:  "Airvana",
		3709: "Airvod",
		9192: "Airware",
		1995: "Airwave",
		6659: "Airwire",
		9152: "Aisai",
		7134: "Aisino",
		5506: "Aitech",
		6923: "Aitexin",
		5243: "Aiwa",
		2616: "Aiware",
		5837: "Ajile",
		4569: "Ajinextek",
		4863: "Ajoho",
		3820: "Akamai",
		4988: "Akamba",
		1726: "Akcp",
		8979: "Akenori",
		2795: "Akimbi",
		1691: "Akita",
		1342: "Akom",
		2869: "Akorri",
		2115: "Aksys",
		6414: "Akuvox",
		9111: "Akyllor",
		2319: "Alacritech",
		4427: "Alacron",
		4277: "Alamar",
		6051: "Alantro",
		5832: "Alaris",
		3857: "Alarm.com",
		7050: "Albahith",
		1504: "Albatron",
		3608: "Albis",
		2069: "Alcatel",
		6993: "Alcea",
		7203: "Alcomp",
		1460: "Alcon",
		3835: "Alektrona",
		3911: "Alereon",
		3757: "Alertme.com",
		5329: "Alerton",
		3916: "Alertus",
		838:  "Alesis",
		935:  "Alexon",
		5938: "Alfa",
		6124: "Alfatech",
		4536: "Alflex",
		6858: "Alge-Timing",
		4496: "Algo",
		3938: "Algolith",
		4566: "Algolware",
		5155: "Algorithmics",
		4605: "Algorithmix",
		4189: "Algorithms",
		6000: "Alidian",
		3671: "Alien",
		1680: "Align",
		6756: "Alinco",
		8168: "Alinket",
		4326: "AliphCom",
		1295: "Alistel",
		219:  "Alitec",
		2896: "Alive",
		7833: "Aliwei",
		365:  "All-Win",
		8677: "AllDSP",
		606:  "Allegro",
		5577: "Allgon",
		2653: "Alliant",
		8095: "Allied-telesisK.K",
		7499: "Allis",
		2264: "Allnet",
		586:  "Alloptic",
		1349: "Allot",
		93:   "Alloy",
		1111: "Allradio",
		3387: "Alltec",
		5538: "Allumer",
		4923: "Allwell",
		9085: "Allwinner",
		1587: "Allworx",
		5640: "Aloha",
		7081: "Along",
		4036: "Aloys",
		2237: "Alpha",
		5296: "Alpha-TOP",
		8854: "AlphaTheta",
		8225: "Alphatronics",
		4156: "Alphion",
		5752: "Alps",
		430:  "Alpsalpine",
		3611: "AlsterAero",
		2310: "Alta",
		3485: "Altai",
		6392: "Altasec",
		3394: "Altec",
		4216: "Altech",
		5409: "Alteon",
		1167: "Altera",
		1559: "AltiGen",
		2917: "Alticast",
		6825: "Altierre",
		5653: "Altiga",
		7032: "AltoBeam",
		128:  "Altos",
		8751: "Altronic",
		7386: "Alula",
		3766: "AmRoad",
		690:  "Amann",
		5254: "Amano",
		6030: "Amaquest",
		5405: "Amati",
		5439: "Amazon",
		7478: "Amazon.com",
		5002: "Amber",
		2350: "AmbiCom",
		3007: "Ambient",
		8220: "Ambiq",
		1043: "Ambrado",
		1601: "Ambri",
		8328: "Amcrest",
		5512: "Amdahl",
		3897: "Amec",
		2672: "Amedia",
		360:  "Amer.com",
		2359: "American",
		4489: "Amerigon",
		106:  "Ameristar",
		5734: "Ameritec",
		2419: "Amherst",
		3698: "Amic",
		5999: "Amigo",
		6459: "Amimon",
		320:  "Amino",
		1924: "Amity",
		4192: "Amkly",
		1973: "Ammasso",
		2580: "Amoi",
		6668: "Amon",
		8941: "Amosense",
		8116: "Amper",
		81:   "Ampere",
		1487: "Amperion",
		50:   "Ampex",
		1475: "Amphenol",
		6551: "Amphitech",
		754:  "Amphus",
		2063: "Ample",
		3326: "Amplex",
		6694: "Amplified",
		5102: "Ampro",
		5443: "Ampt",
		6362: "Amsc",
		2062: "Amtelco",
		3358: "AnNeal",
		6464: "AnaCom",
		2855: "Anagran",
		3587: "Analogic",
		8676: "Analytica",
		2284: "Anam",
		4485: "Anatek",
		1766: "Anator",
		3379: "Anaveo",
		5403: "Ancot",
		5298: "Anda",
		4571: "Andes",
		782:  "Andiamo",
		2574: "Andrew",
		7004: "Andtek",
		7037: "Anedo",
		6160: "Anerma",
		3515: "Angel",
		5735: "Angia",
		7656: "Angler",
		1786: "Animation",
		1468: "Annso",
		8974: "Anobit",
		1149: "Anoto",
		7123: "Anovo",
		98:   "Anritsu",
		5133: "Ansel",
		3594: "Anseri",
		6444: "Ansjer",
		4599: "Ansync",
		6504: "Antailiye",
		7957: "Antaira",
		2416: "Antec",
		7706: "Antex",
		1639: "Anthology",
		4563: "Antipode",
		5449: "Antlow",
		3490: "AnyDATA",
		8227: "Anywave",
		6943: "Anywire",
		1734: "Aoip",
		9090: "Apacer",
		4818: "Apacewave",
		1500: "Apani",
		4490: "Apass",
		7823: "Aperi",
		406:  "Apex",
		5700: "Apexx",
		7852: "Apiste",
		4296: "Aplicaciones",
		8862: "Aplicom",
		5578: "Aplio",
		591:  "Apogee",
		5424: "Apollo",
		4353: "AppTech",
		4963: "Appian",
		545:  "Apple",
		2729: "Application",
		5766: "Applicom",
		41:   "Applicon",
		4492: "Applition",
		2441: "Appointech",
		3638: "Apprion",
		1755: "Appro",
		8656: "Apption",
		3243: "Apricorn",
		47:   "Apricot",
		150:  "April",
		4547: "Aprius",
		4123: "Aprotech",
		5964: "Aptec",
		4460: "Aptiv",
		6112: "Aptix",
		8252: "Aptos",
		6889: "Aqavi.com",
		3254: "Aquantia",
		4669: "Aquila",
		4877: "Arada",
		3194: "Araneo",
		182:  "Aravox",
		7604: "Arbiter",
		3098: "Arbitron",
		859:  "Arbor",
		9020: "Arcadia",
		2640: "Arcadyan",
		5916: "Archipel",
		3151: "Archos",
		5863: "Arco",
		8138: "Arcom",
		1568: "Arcon",
		9223: "Arcontia",
		1788: "Arcor&Co",
		926:  "Arcturus",
		9420: "Arcx",
		6044: "Ardent",
		6921: "Ardomus",
		8468: "Arduino",
		5904: "Areanex",
		3664: "Areca",
		4478: "Ared",
		429:  "Arelnet",
		879:  "Arescom",
		8858: "Areson",
		3313: "Argard",
		5632: "Argon",
		2401: "Argosy",
		471:  "Argus",
		5543: "Aries",
		490:  "Arima",
		3794: "Arista",
		127:  "Arix",
		2685: "Arkados",
		8407: "Arlo",
		4030: "Armadeus",
		9097: "Armatura",
		1140: "Armillaire",
		3332: "Armorlink",
		6549: "Armtel",
		5956: "Array",
		2406: "ArrayComm",
		9500: "Arrayent",
		9150: "Arrcus",
		9361: "Arrikto",
		6942: "Arris",
		7817: "Arrive",
		3481: "Arrow7",
		2332: "ArrowPoint",
		2853: "ArrowSpan",
		2928: "Artech",
		5507: "Artel",
		3219: "Arti",
		2692: "Artila",
		2767: "Artimi",
		69:   "Artisoft",
		5373: "Artiza",
		3918: "Artjoy",
		1494: "Artnix",
		4884: "Artray",
		1685: "Aruba",
		8705: "Aryaka",
		5222: "Asaca",
		7325: "Asahi",
		856:  "Asahi-Engineering",
		100:  "Asante",
		4645: "Asantron",
		2601: "Ascalade",
		3298: "Ascend",
		1690: "Ascent",
		1692: "Ascom",
		9495: "Ascon",
		7966: "Asct",
		3290: "Asia",
		1511: "AsiaRF",
		2673: "Asiamajor",
		1666: "Asiarock",
		6977: "Asiatelco",
		4893: "Asiteq",
		2104: "Asix",
		1545: "Aska",
		2545: "Askey",
		2006: "Asmax",
		3078: "Asmobile",
		7471: "Asoni",
		343:  "Asound",
		5657: "Aspect",
		940:  "Aspen",
		7042: "Asperiq",
		6725: "Assa",
		8279: "Assembled",
		2462: "Assurance",
		2073: "Astarte",
		4079: "Astec",
		6351: "Astech",
		2639: "Astek",
		1567: "Astera",
		5762: "Aston",
		2267: "Astri",
		5378: "AstroNova",
		450:  "Astrodesign",
		6405: "Astrol",
		4448: "Astron",
		2219: "Astute",
		4686: "Asumo",
		3371: "AtRoad",
		7382: "Atamo",
		6069: "Atan",
		33:   "Atari",
		3402: "Atech",
		1154: "Atek",
		7625: "Ateme",
		2346: "Aten",
		3312: "Atera",
		138:  "Atex",
		8534: "AthenTek",
		4309: "Athena",
		5127: "Athenix",
		535:  "Atheros",
		9048: "Atil",
		5457: "Atlantix",
		7822: "Atlinks",
		8678: "Atlona",
		633:  "Atmel",
		7933: "Atmosic",
		2383: "Atmosphere",
		543:  "Atoga",
		9199: "Atol",
		9143: "Atomax",
		7513: "Atomos",
		5169: "Atomwide",
		5421: "Atop",
		9399: "Atopia",
		599:  "Atrica",
		2011: "Atrie",
		3843: "Attero",
		2832: "Atto",
		6595: "AuVerte",
		241:  "AudeSi",
		9128: "Audeze",
		1660: "Audio",
		1503: "Audio-Technica",
		7162: "AudioControl",
		2587: "AudioDev",
		378:  "AudioRamp.com",
		3855: "AudioScience",
		8533: "Audioengine",
		4316: "Audiovox",
		8463: "Audiowise",
		8466: "Audivo",
		6621: "Audoo",
		8758: "Audyssey",
		1364: "Auerswald",
		2163: "Augmentix",
		8140: "Augtek",
		9018: "Aurender",
		1119: "Aurora",
		8142: "Aus.Linx",
		37:   "Auspex",
		3473: "Austar",
		5747: "Austron",
		4237: "Autec",
		7602: "Auth-Servers",
		6388: "Auto",
		3123: "Auto-Maskin",
		1848: "AutoCell",
		8438: "AutoCrib",
		5813: "AutoGas",
		8215: "AutoHotBox",
		9462: "Autobase",
		3605: "Autocom",
		4211: "Autocomputer",
		7020: "Autoit",
		6190: "AutomatedLogic",
		8781: "Automatic",
		3187: "Automation",
		4690: "Autonet",
		7459: "Autonics",
		6926: "Autosales",
		1517: "Autostar",
		634:  "Autosys",
		8184: "Autotalks",
		3006: "Autotelenet",
		27:   "Autotote",
		2047: "Auvitran",
		4412: "Avaak",
		3215: "Avago",
		295:  "Avail",
		6905: "Availink",
		5207: "Aval",
		669:  "Avalue",
		2230: "Avamax",
		2229: "Avanex",
		2638: "Avantec",
		3646: "Avantis",
		8208: "Avanu",
		2072: "Avara",
		5571: "Avatar",
		620:  "Avaya",
		3284: "Avega",
		7932: "Avelon",
		8476: "Aventura",
		4457: "AverLogic",
		6552: "Avermetrics",
		5680: "Avex",
		5659: "Avici",
		5755: "Avid",
		3884: "Avidyne",
		1243: "Avilinks",
		3159: "Aviqtech",
		349:  "Avision",
		1275: "Avistar",
		6559: "Avistel",
		4523: "Avitech",
		8609: "Aviwest",
		2539: "Avix",
		8557: "Avizia",
		423:  "Avnet",
		2634: "Avolites",
		8274: "Avonic",
		8106: "Avotek",
		1528: "Avrio",
		1798: "Avtec",
		4804: "Avtex",
		7631: "Avvasi",
		1689: "Avvio",
		2371: "Aware",
		2920: "AwarePoint",
		3697: "Awox",
		9071: "Axacore",
		5723: "Axel",
		9130: "AxesNetwork",
		3505: "Axesstel",
		8306: "Axiim",
		8634: "Axilspot",
		5416: "Axiom",
		3919: "Axion",
		1945: "Axiowave",
		6430: "Axiros",
		5132: "Axis",
		6460: "Axium",
		5943: "Axon",
		4638: "Axona",
		9477: "Axonne",
		2524: "Axsun",
		9436: "Axview",
		6716: "Axxana",
		8931: "Ayecka",
		3530: "Ayecom",
		7565: "Ayla",
		3225: "Azalea",
		3239: "Azonic",
		5797: "Azonix",
		7501: "Azroad",
		4855: "Aztech",
		1357: "Aztek",
		2189: "Azul",
		8241: "Azuray",
		5572: "Azure",
		3668: "Azuretec",
		3005: "Azurewave",
		1478: "Azylex",
		2100: "B&B",
		7224: "B-Link",
		6074: "B2C2",
		3142: "BA",
		2352: "BACHMANN",
		1743: "BAE",
		9061: "BALMUDA",
		3534: "BARTEC",
		8732: "BASIC",
		6158: "BAY",
		3061: "BBH",
		8062: "BBMC",
		1994: "BBN",
		2582: "BBWM",
		2206: "BBox",
		382:  "BCM",
		8848: "BCTech",
		1871: "BECS",
		8154: "BEFEGA",
		7725: "BEKO",
		4076: "BEN-RI",
		2930: "BETA",
		9120: "BH",
		5745: "BHP",
		7417: "BHR",
		7213: "BHTC",
		4694: "BICOM",
		3895: "BIJ",
		8933: "BIOS",
		1716: "BIOTRONIK",
		4043: "BISA",
		7585: "BITwave",
		7136: "BK",
		9010: "BKAV",
		2085: "BLIP",
		8067: "BLT",
		7461: "BLU",
		2719: "BLwave",
		2764: "BM",
		4197: "BMC",
		9475: "BME",
		258:  "BMW",
		192:  "BNA",
		9362: "BND",
		6439: "BNS",
		1013: "BNTECHNOLOGY",
		2836: "BOAZ",
		2079: "BOE",
		3670: "BPL",
		3828: "BPT",
		7284: "BQ",
		7407: "BQT",
		446:  "BRECIS",
		5274: "BRODEL",
		8944: "BSE",
		6582: "BSMediasoft",
		6273: "BSP",
		1261: "BST",
		3513: "BSkyB",
		8685: "BT",
		3132: "BT-Links",
		8816: "BT9",
		5741: "BTG",
		2904: "BTI",
		1843: "BTU",
		7692: "BUJEON",
		5570: "BVM",
		2238: "BWA",
		8189: "BXB",
		4047: "BYD",
		6908: "BYK-Gardner",
		8809: "Bach-Simpson",
		4838: "Bachmann",
		7253: "Baicells",
		7289: "Baikal",
		4595: "Ball-It",
		3417: "Balluff",
		7473: "Baltech",
		6062: "Balthazar",
		7166: "Bamboo",
		4903: "BandRich",
		961:  "Banderacom",
		1008: "Bandspeed",
		1197: "Banksys",
		4596: "Banner",
		6327: "Banyan",
		9086: "Baraja",
		6624: "Barberry",
		721:  "Bardac",
		1297: "Barix",
		465:  "Barracuda",
		6259: "Barrot",
		1362: "Bartech",
		8850: "Barun",
		4968: "Basler",
		4348: "Bastec",
		5471: "Basys",
		3989: "Battistoni",
		2684: "BaudTec",
		8348: "Baumer",
		7435: "Baxter",
		84:   "Bay",
		8531: "Baycity",
		813:  "Baydel",
		5642: "Bayly",
		6188: "Bcom",
		1857: "BeWAN",
		7672: "Beacon",
		1582: "BeamReach",
		8244: "Beamex",
		6262: "Beats",
		8867: "Beautiful",
		8553: "Becker",
		8398: "Becker-Antriebe",
		3684: "Beckmann",
		6693: "Becon",
		8347: "Becton",
		9202: "Befs",
		5991: "Behavior",
		1086: "Beicom",
		5238: "Beijer",
		9184: "Beijing-Cloud",
		6911: "Beissbarth",
		7783: "Beken",
		7914: "Bekey",
		2662: "Belco",
		3427: "Belden",
		2469: "Belkin",
		15:   "Bell",
		3824: "Bellon",
		1062: "Bematech",
		7989: "BenQ",
		8683: "BenRui",
		2080: "Benchmark",
		4598: "Benein",
		4620: "Benign",
		9107: "Bentek",
		9360: "Benu",
		6155: "Berkeley",
		4085: "Berker",
		4325: "Berkshire",
		2041: "Bermai",
		213:  "Best",
		3927: "BestComm",
		7836: "Bestek",
		5265: "Beta",
		8639: "BetterBots",
		8329: "Beyless",
		3412: "Beyondwiz",
		3706: "Bharat",
		7648: "Bhuu",
		9299: "Bi2-Vision",
		1652: "BiTMICRO",
		972:  "Biacore",
		7874: "Biamp",
		505:  "BigBand",
		3398: "Bigfoot",
		2335: "Billionton",
		1951: "Bils",
		6369: "BinCube",
		1591: "Binatone",
		5815: "Bintec",
		8769: "Binxin",
		3764: "Bio-Rad",
		9240: "BioControl",
		7152: "Biodata",
		4611: "Biodevices",
		7345: "Biodit",
		425:  "Bionet",
		8446: "Bionics",
		8586: "Bionime",
		5608: "Biopac",
		4073: "Bioscrypt",
		8344: "Biosoundlab",
		2707: "Biospace",
		3711: "Bird",
		8764: "Biscotti",
		7300: "Bison",
		8780: "BitBox",
		7708: "Bita",
		2645: "Bitatek",
		6366: "Bitel",
		9145: "Bitmain",
		1048: "Bitrage",
		995:  "Bitran",
		4725: "Bitrode",
		6008: "Bitronics",
		3415: "BitsGen",
		6057: "Bitswitch",
		1638: "BittWare",
		1160: "Bitworks",
		920:  "Bivio",
		2998: "Bixolon",
		3612: "Biz-2-Me",
		8356: "BizLink",
		6427: "Bizlink",
		2221: "BlackBerry",
		8659: "Blaster",
		4101: "Blatand",
		9301: "Bloombase",
		4130: "Blue-White",
		903:  "Blue2space",
		5059: "BlueBit",
		9383: "BlueBite",
		2515: "BlueExpert",
		1311: "BlueKorea",
		8807: "BlueN",
		2624: "BluePacket",
		9287: "BlueRadios",
		1436: "BlueWINC",
		8440: "Bluebank",
		3982: "Bluecard",
		8684: "Bluecom",
		1108: "Bluegiga",
		4037: "Bluelight",
		3659: "Blueone",
		1307: "Bluetake",
		4834: "Bluetechnix",
		3712: "Blueway",
		1532: "Bluewire",
		3590: "Blusens",
		8060: "Bluwan",
		4319: "Blynke",
		7855: "BnCOM",
		6617: "BobjGear",
		3479: "Bobst",
		5901: "Boca",
		1684: "Bodet",
		321:  "Bodmann",
		8249: "BodyMedia",
		1041: "Boeing",
		3309: "Bogen",
		8669: "Bolymin",
		1236: "Bon",
		3520: "Bona",
		8206: "Bookeen",
		2859: "Bookham",
		6904: "Boosty",
		8616: "Borea",
		5042: "Borgardt",
		1092: "Boris",
		1926: "Bosch",
		1819: "Bose",
		5276: "Boser",
		8247: "Bostex",
		5817: "Boston",
		6691: "Bosung",
		9407: "Botato",
		4107: "Botik",
		6511: "Bowei",
		8319: "BoxCast",
		8380: "BoxLock",
		9038: "Boxin",
		4139: "Boyoung",
		8994: "Bragi",
		7775: "Brain",
		1146: "Brain21",
		7445: "BrainCo",
		1581: "Brainchild",
		2218: "Brainium",
		271:  "Brains",
		1361: "Braintree",
		5720: "Brand",
		4321: "Brandywine",
		795:  "Brans",
		971:  "Bravara",
		6764: "Braveridge",
		9121: "Bravo",
		8453: "Breathometer",
		700:  "Breezecom",
		165:  "Bri-Link",
		6246: "BriView",
		4652: "Bridge",
		517:  "BridgeWave",
		1173: "Bridgeco",
		629:  "Bridgeworks",
		1993: "Bright",
		8202: "BrightSign",
		1053: "Brightcom",
		7490: "Brightsource",
		7413: "Brilliantts",
		3962: "Briot",
		8752: "Brita",
		1130: "Britestream",
		3861: "Brivo",
		4995: "Brix",
		2570: "BroadEasy",
		3491: "Broadata",
		2423: "Broadband",
		1754: "Broadbus",
		858:  "Broadcom",
		405:  "Broadframe",
		8613: "Broadlink",
		6028: "Broadlogic",
		709:  "Broadmax",
		4586: "Broadway",
		90:   "Brocade",
		449:  "Bromax",
		5761: "Brooktrout",
		3231: "Brother",
		3054: "Browan",
		4025: "Brunata",
		4961: "Bryant",
		6253: "Bryston",
		7788: "Bryton",
		510:  "Bticino",
		1077: "Buffalo",
		9065: "Bulat",
		2099: "Burdick",
		6906: "Burg-Wachter",
		6663: "Burlywood",
		5886: "Burr-Brown",
		4553: "Burster",
		3101: "Bury",
		5086: "Bustek",
		8417: "Busware.DE",
		7524: "Buwon",
		3334: "Buyang",
		6184: "Byas",
		9165: "Bytedance",
		8549: "Bytemark",
		5147: "Bytex",
		4432: "Byzoro",
		8129: "C",
		1353: "C&I",
		1305: "C&S",
		3939: "C-BEL",
		305:  "C-CoM",
		338:  "C-CoR",
		404:  "C-CoR.net",
		4535: "C-Matic",
		2016: "C-guys",
		5174: "C.A.E.N",
		7572: "C.E",
		8211: "C.O.B.O",
		810:  "C.P",
		2810: "C4Line",
		454:  "CAB",
		2192: "CABLELOGIC",
		5233: "CAE",
		4405: "CAI",
		6604: "CAP-Tech",
		1326: "CAS",
		6373: "CASwell",
		1194: "CATS",
		3829: "CBC",
		5196: "CBL",
		4811: "CC",
		385:  "CC&C",
		4120: "CCS",
		816:  "CDS-Electronics",
		7504: "CDYNE",
		2556: "CE-Infosys",
		7466: "CELIZION",
		844:  "CELOX",
		9103: "CELOXICA",
		1234: "CEM",
		3302: "CENITS",
		4682: "CENTRAL",
		1084: "CENiX",
		4772: "CEStronics",
		2425: "CET",
		7659: "CETORY.TV",
		4734: "CEVA",
		8155: "CG",
		7348: "CHAHOO",
		2727: "CHIPS",
		2440: "CIMSYS",
		164:  "CIS",
		3073: "CITEL",
		7077: "CKD",
		8013: "CKS",
		834:  "CLCsoft",
		2013: "CMA/Microdialysis",
		5286: "CMC",
		4246: "CMD",
		3527: "CMOTECH",
		2364: "CMS",
		4005: "CNB",
		7497: "CNEX",
		1877: "CNMP",
		8957: "CNSLink",
		1066: "CNSystems",
		1253: "CNet",
		9261: "CPI",
		3893: "CQ",
		2924: "CREVIS",
		4010: "CRFS",
		7568: "CRU-Dataport",
		3077: "CReTE",
		8618: "CRemote",
		3363: "CS",
		6629: "CSE-Servelec",
		2391: "CSI-Control",
		5455: "CSK",
		8993: "CSM",
		34:   "CSS",
		4635: "CSSI",
		4723: "CTERA",
		2413: "CTI",
		2059: "CTS",
		8488: "CVC",
		9012: "CVT",
		3127: "CWT",
		2313: "CYQ've",
		654:  "CYZENTECH",
		9036: "CZ.NIC",
		6610: "Caavo",
		4174: "Cable",
		6512: "CableFree",
		8596: "CableWorld",
		1148: "Cabletime",
		16:   "Cabletron",
		420:  "Cablevision",
		2164: "Cableware",
		1492: "Cabot",
		5920: "Cache",
		5016: "CacheFlow",
		1328: "CacheVision",
		7841: "Cachengo",
		5450: "Cactus",
		8808: "Cadac",
		212:  "Cadant",
		2565: "Cadco",
		5551: "Cadre",
		2594: "Caen",
		1630: "Cal-Comp",
		7588: "CalDigit",
		7771: "Calantec",
		5950: "Calcomp",
		2840: "Calculex",
		8638: "Caldero",
		2402: "Calista",
		374:  "Calix",
		4761: "CallTechSolution",
		2251: "CallURL",
		3763: "Callpod",
		6753: "Calsys",
		7846: "Calxeda",
		5838: "Caly",
		3696: "Calyptech",
		5211: "Cambex",
		665:  "Cambium",
		1640: "Cambridge",
		8411: "Cambrionix",
		7197: "Camco",
		3384: "Cameo",
		8995: "Cameronet",
		8:    "Camex",
		6021: "Campio",
		3245: "Camrivox",
		158:  "Camtec",
		309:  "Camtel",
		5962: "Canary",
		112:  "Canberra",
		1136: "Candera",
		3026: "Canesta",
		3577: "Canhold",
		3271: "Canko",
		87:   "Canon",
		4180: "Canopus",
		5494: "Canstar",
		4056: "Cantronic",
		2678: "Cap",
		6811: "Capelec",
		3003: "Capelon",
		1403: "Capinfo",
		6854: "Capisco",
		1151: "Caporis",
		9425: "CaptionCall",
		5721: "Captor/SA",
		6481: "Cara",
		2760: "Carallon",
		2174: "CardioNet",
		7053: "Cardiopulmonary",
		5678: "Cardkey",
		6939: "CarePredict",
		9307: "CareView",
		4849: "Carecom",
		9442: "Carefusion",
		3008: "Caretech",
		6954: "Cargt",
		3050: "Carina",
		8983: "Carma",
		8903: "Carnegie",
		3841: "Carpoint",
		4289: "Carrera",
		370:  "Carrier",
		222:  "CarrierComm",
		8406: "Carry",
		3171: "Casa",
		7137: "Casacom",
		5184: "Cascade",
		8378: "Cashmaster",
		6346: "Casio",
		3145: "Caspian",
		3832: "CastGrabber",
		2091: "Castel",
		44:   "Castelle",
		1401: "Castle",
		3798: "Castlenet",
		3922: "Castles",
		2068: "Catalyst",
		552:  "Catapult",
		3308: "Catcher",
		2495: "Category",
		373:  "Catena",
		1531: "Caterpillar",
		8047: "Cathay",
		3494: "Cathexis",
		2797: "Cavera",
		2250: "Cavium",
		4505: "Cayee",
		91:   "Cayman",
		7394: "Cdoubles",
		4041: "Cdvi",
		7066: "Cedes",
		6591: "Cedint-UPM",
		4794: "Cedo",
		1032: "Ceemax",
		2449: "Ceicom",
		2854: "CelPlan",
		5960: "Celan",
		5349: "Celcore",
		3777: "Celeno",
		6213: "Celestica",
		218:  "Celestix",
		4039: "Celio",
		1093: "Celleritas",
		7598: "Cellient",
		2187: "Cellink",
		1518: "Cellinx",
		7256: "Cello",
		6137: "Cellport",
		483:  "Cellvision",
		8602: "Celona",
		993:  "Celsian",
		3641: "Celtro",
		3958: "Centec",
		3751: "Center",
		5588: "Centigram",
		5699: "Centillion",
		173:  "Centillium",
		266:  "Centos",
		3125: "CentraLite",
		4696: "Centrak",
		6863: "Centripetal",
		4713: "Centrofactor",
		5871: "Centrum",
		5500: "Century",
		8361: "CenturyLink",
		2325: "Ceologic",
		4322: "Cepheid",
		6592: "Cepton",
		8592: "Cera",
		8942: "CeraMicro",
		1490: "Ceragon",
		9021: "Cercacor",
		6626: "CerebrEX",
		7767: "Cerebras",
		8128: "Cerio",
		2868: "Cermate",
		6314: "Cern",
		4303: "Cernium",
		4401: "Certicom",
		2435: "Cesnet",
		2728: "Cetacea",
		1396: "Cetacean",
		140:  "Cetia",
		3507: "Cetis",
		4417: "Ceton",
		3040: "Ceyon",
		8036: "Chabrier",
		3433: "Chainleader",
		5242: "Chaintech",
		3900: "Chainzone",
		6443: "Chameleon",
		9176: "Championtech",
		7808: "ChangYang",
		6654: "Changhe",
		5392: "Channelmatic",
		2263: "Chantry",
		8430: "ChargeStorm",
		467:  "Charles",
		5550: "Chase",
		4941: "Checkout",
		3235: "Checkpoint",
		4897: "Cheerchip",
		7510: "Cheerstar",
		3030: "Cheertek",
		1079: "Chelsio",
		7728: "Chemoptics",
		7075: "Chemtronics",
		3710: "Cherry",
		4248: "Chess",
		509:  "Chiaro",
		1797: "Chic",
		7301: "Chicony",
		1742: "Chih-Kan",
		2555: "Chinasys",
		702:  "Chino",
		4314: "Chip-pro",
		605:  "Chip2Chip",
		6650: "ChipER",
		6360: "Chipcom",
		6440: "Chipsea",
		7828: "Chipsip",
		2489: "Chiron",
		7438: "Chitai",
		7141: "Chiyoda",
		2124: "Chiyu",
		4583: "Chroma",
		205:  "Chromatek",
		28:   "Chromatics",
		5195: "Chromatis",
		336:  "Chromisys",
		8641: "Chromlech",
		9288: "Chrontel",
		2988: "ChuanG",
		3780: "Chubb",
		6271: "Chun-il",
		5917: "Chuntex",
		5179: "Chuo",
		36:   "Chyron",
		717:  "Ciac",
		2746: "Ciara",
		4685: "Ciat",
		478:  "Cidco",
		1988: "Cidra",
		4565: "Ciena",
		7842: "Ciesse",
		7534: "Ciholas",
		8058: "Cilag",
		2362: "Cimetrics",
		114:  "Cimlinc",
		6209: "Cinco",
		6873: "Cincoze",
		4454: "Cinetal",
		781:  "Cinta",
		2796: "Cintech",
		5869: "Cipher",
		1455: "CipherOptics",
		4207: "Ciprico",
		3837: "Circleone",
		4014: "CiriTech",
		304:  "Cirilium",
		898:  "Cirkitech",
		3806: "Cirrascale",
		2:    "Cisco",
		5368: "Citadel",
		1033: "Citel",
		5712: "Citicorp/TTI",
		3486: "Citiway",
		9258: "Citrix",
		2586: "Citronix",
		5018: "City-Net",
		2825: "CityCom",
		6260: "Clack",
		7259: "ClarIDy",
		3898: "ClariPhy",
		5318: "Clariion",
		2370: "Clarinet",
		7130: "Clavister",
		823:  "ClearCube",
		5633: "ClearOne",
		4704: "ClearPath",
		4304: "Clearbox",
		5534: "Clearpoint",
		3908: "Clearwire",
		2071: "Clematic",
		8139: "Cleondris",
		5688: "Clevo",
		337:  "ClickTV",
		5654: "Clientec",
		3915: "Climax",
		6385: "CliniCare",
		1110: "Clipcomm",
		6231: "Cloos",
		9272: "CloudGenix",
		1701: "CloudShield",
		7545: "CloudSwitch",
		3190: "Cloudastructure",
		9163: "Cloudena",
		8637: "Cloudistics",
		6665: "Cloudium",
		9132: "Cloudleaf",
		6890: "Cloudtronics",
		8469: "Cloudview",
		5435: "Cloudwalk",
		4509: "Clover",
		1378: "Cloverleaf",
		580:  "Clovertech",
		1356: "Cluster",
		82:   "Clustrix",
		2327: "Cmicro",
		7328: "Cmitech",
		4981: "Cmos",
		5537: "Cnet",
		4453: "Cnrs",
		7930: "CoBI",
		8719: "CoDACO",
		2145: "CoM&C",
		2957: "CoM.s.a.t",
		5750: "CoM21",
		6841: "CoMESTA",
		4777: "CoMFILE",
		8114: "CoMMANDO",
		4856: "CoMMidt",
		7974: "CoMTEC",
		4578: "CoMXION",
		6579: "CoNCH",
		7365: "CoNELCOM",
		6721: "CoNSTELL8",
		7330: "CoNTROLtronic",
		9484: "CoNWIN.Tech",
		3738: "CoNWISE",
		285:  "CoNet",
		3041: "CoOLKSKY",
		2114: "CoPAN",
		5393: "CoRELIS",
		4373: "CoRNELL",
		570:  "CoT",
		3787: "CoVAX",
		2782: "CoVi",
		3991: "CoWON",
		4748: "CoachComm",
		9384: "Coagent",
		5748: "Coastcom",
		2488: "Coaxial",
		9331: "CobaltRay",
		8431: "Cobham",
		6848: "Cobs",
		8774: "Coby",
		8038: "Cochlear",
		2424: "Cocom",
		2652: "Codan",
		1740: "Code",
		2842: "Coder",
		8605: "Codetek",
		6321: "Codex",
		1943: "Codian",
		5148: "Codonics",
		7437: "Coflec",
		9196: "Cofractal",
		1675: "Cogent",
		5983: "Cognex",
		2246: "Cognio",
		9034: "Cognitas",
		7060: "Cognitec",
		7453: "Cognitive",
		8680: "Cohere",
		3260: "Coherent",
		7222: "Cohesity",
		1458: "Cohu",
		1914: "Collex",
		6709: "Collinear",
		9169: "Color-Chip",
		7002: "ColorTokens",
		5076: "Colorgraph",
		512:  "Colubris",
		6341: "ComDesign",
		4482: "ComWorth",
		7185: "Comat",
		7616: "CombiQ",
		5182: "Combinet",
		6031: "Comcam",
		5968: "Comda",
		439:  "Comdial",
		5940: "Comelta",
		5109: "Comendec",
		1733: "Comflux",
		8779: "Comigo",
		8159: "Comlab",
		4382: "CommAgility",
		9058: "CommFront",
		5077: "CommScope",
		6982: "CommSky",
		1224: "Command-e",
		1641: "Commax",
		1842: "Commend",
		3419: "CommerceGuard",
		3439: "Commerciant",
		1150: "Commil",
		5458: "Commodore",
		6805: "Common",
		6053: "Commscope",
		5525: "Commscraft",
		7818: "Commsen",
		4832: "Commtact",
		757:  "Commtech",
		4424: "Communication",
		2393: "Communications",
		5772: "Commvision",
		7282: "Comnect",
		5021: "Comone",
		2003: "CompXs",
		5080: "Compac",
		9028: "Compacta",
		358:  "Compal",
		4102: "Compass",
		7109: "Compass-EOS",
		1906: "Compellent",
		5082: "Compex",
		767:  "Complementary",
		3393: "Compro",
		130:  "Compu-Shack",
		275:  "CompuLab",
		5451: "Compuadd",
		1775: "Compucase",
		55:   "Compucorp",
		1887: "Compulogic",
		4414: "Compumedics",
		1566: "Compunetix",
		7858: "Compupal",
		5418: "Compuserve",
		6065: "Computational",
		3652: "Computec",
		2329: "Computer",
		5152: "Computerm",
		6300: "Computervision",
		5702: "Computex",
		3994: "Computime",
		5498: "Computone",
		1325: "Computrols",
		6150: "Compuware",
		5767: "Comsat",
		8499: "Comsis",
		5069: "Comsoft",
		594:  "Comspace",
		3983: "Comsys",
		6101: "Comtec",
		4261: "Comtech",
		5941: "Comtree",
		3870: "Comtrend",
		308:  "Comtrol",
		4209: "Comtron",
		7398: "ConMet",
		2575: "ConSentry",
		8643: "Conception",
		4501: "Conceptronic",
		66:   "Concord",
		5141: "Concurrent",
		3708: "Condalo",
		485:  "Condev",
		5492: "Condor",
		8346: "Conductix-Wampfler",
		7298: "Conet",
		5025: "Conexant",
		6883: "Conextop",
		5369: "ConferTech",
		3899: "Confidant",
		2741: "Congatec",
		4932: "Congruency",
		223:  "Conklin",
		7113: "Conlog",
		22:   "Connect",
		2661: "ConnectBlue",
		7061: "ConnectQuest",
		6009: "Connected",
		912:  "Connection",
		8448: "Connex",
		1713: "Connexionz",
		281:  "Conrad",
		1942: "Consensys",
		8610: "Consert",
		4374: "Consilium",
		2833: "Consultronics",
		4663: "ConteXtream",
		5483: "Contec",
		3168: "Contela",
		2898: "Contemporary",
		1276: "Contex",
		6611: "Continental",
		5382: "Continuum",
		1999: "Control",
		2293: "Control4",
		1856: "ControlNet",
		8640: "ControlThings",
		4960: "Controlled",
		5348: "Controlware",
		970:  "Convedia",
		8670: "Convergence",
		5010: "Convergenet",
		3508: "Convergens",
		5973: "Convergent",
		8286: "ConversDigital",
		5794: "Convex",
		4700: "Convey",
		6087: "Convision",
		110:  "Conware",
		4853: "Coordiwise",
		5117: "Copernique",
		5693: "CopperCom",
		2726: "CorEdge",
		5781: "Cordant",
		5240: "Corder",
		475:  "Core",
		1580: "CoreBell",
		6374: "CoreEdge",
		8621: "CoreOS",
		2889: "CoreStar",
		6625: "CoreTrust",
		5652: "Corecess",
		3588: "Corelatus",
		8394: "Coresys",
		4164: "Coretree",
		5221: "Coretronic",
		6681: "Coriant",
		1714: "Corinex",
		8293: "Corintech",
		386:  "Coriolis",
		8177: "Cornami",
		489:  "Cornet",
		4950: "Corning",
		5119: "Corollary",
		2850: "Corona",
		1609: "Corrent",
		1291: "Corrigent",
		6800: "Corsa",
		2679: "Cortina",
		9457: "Cortland",
		1133: "Corvalent",
		4421: "Corventis",
		5244: "Corvis",
		6330: "Corvus",
		7953: "Cosco",
		6033: "Cosine",
		1982: "Cosmic",
		354:  "Cosmo",
		4780: "Costar",
		9314: "Costel",
		6183: "Cosworth",
		5205: "Cosystems",
		2485: "Cotron",
		9208: "Cots",
		4678: "Coulomb",
		5716: "Coulter",
		9377: "CountMax",
		4266: "Counter",
		6335: "Counterpoint",
		7645: "Countwise",
		1163: "Coup",
		6719: "Coval",
		2911: "Covergence",
		4656: "Covia",
		3887: "Covote",
		3842: "Cowbell",
		8696: "Cox",
		2385: "Coyote",
		7270: "Cozybit",
		4990: "Cqos",
		1155: "Cradle",
		4959: "CradlePoint",
		6103: "Cradlepoint",
		3803: "Cranite",
		618:  "Cratos",
		68:   "Cray",
		8341: "Creatcomm",
		7046: "Creation",
		357:  "Creative",
		6881: "Creatz",
		9118: "Cree",
		6776: "Crenus",
		8259: "Creowave",
		2031: "Crere",
		3926: "Crescendo",
		6059: "Crescent",
		2353: "Crestron",
		3283: "Cresyn",
		1541: "Creval",
		384:  "Crewave",
		1519: "Crinis",
		5787: "Cristie",
		7181: "Criticare",
		6301: "Cromemco",
		1400: "Cronyx",
		585:  "Crossbeam",
		3550: "Crossbow",
		103:  "Crosscomm",
		902:  "Crossport",
		6093: "Crossroads",
		2748: "Crow",
		94:   "Cryptek",
		5598: "Crypto",
		3980: "CryptoMetrics",
		2852: "Cryptosoft",
		4033: "Cryptsoft",
		5706: "Csir",
		6127: "Cspi",
		6257: "Csst",
		6772: "Ctek",
		3467: "Ctring",
		2960: "Cube",
		4250: "Cubix",
		9355: "Cuciniale",
		8110: "Cudo",
		8760: "Cuff",
		8571: "Cummings",
		3847: "Cummins",
		4591: "Cummins-Allison",
		7172: "Cumulus",
		8899: "Curiouser",
		6454: "Currant",
		8151: "Current",
		3657: "Curtis",
		7138: "Curvalux",
		3044: "Curves",
		2123: "Custom",
		4810: "Cutera",
		6056: "Cutler-Hammer",
		7363: "CviLux",
		3743: "Cwlinux",
		3973: "CyVerse",
		8803: "Cybelec",
		677:  "Cyber",
		8739: "Cyber-Rain",
		2593: "CyberNet",
		227:  "CyberOptics",
		1761: "CyberPower",
		189:  "CyberTAN",
		8742: "Cybera",
		642:  "Cyberboard",
		4288: "Cyberdata",
		4400: "Cyberdyne",
		5497: "Cybergraphic",
		1654: "Cybernetics",
		4276: "Cybertec",
		4135: "Cybertech",
		7492: "Cybertelbridge",
		3951: "Cybiotronics",
		6855: "Cybo",
		5330: "Cyclades",
		5345: "Cycle",
		8260: "Cydle",
		215:  "Cygnet",
		5724: "Cylink",
		3214: "Cymphonix",
		4302: "Cymtec",
		7904: "Cynosure",
		8689: "Cynove",
		4483: "Cypak",
		4781: "Cypress",
		4998: "Cyra",
		481:  "Cyras",
		2968: "Cytyc",
		873:  "D&M",
		2201: "D&R",
		4759: "D-BOX",
		803:  "D-Link",
		2992: "D-MAX",
		1748: "D-NET",
		4334: "D-TACQ",
		7946: "D.SignT",
		4428: "DAC",
		3272: "DACOS",
		3202: "DAGS",
		1070: "DANCONTROL",
		876:  "DAP",
		2677: "DAQ",
		7409: "DASAN",
		8736: "DASCOM",
		6197: "DATENTECHNIK",
		8939: "DATTUS",
		7522: "DAVOLINK",
		6684: "DB",
		9188: "DBG",
		7008: "DBL",
		49:   "DCI",
		8928: "DCONWORKS",
		5404: "DCS",
		4623: "DCT-Delta",
		5965: "DDL",
		4527: "DDRdrive",
		9268: "DDoS-Guard",
		2256: "DELCOMp",
		4616: "DELEC",
		7175: "DEOTRON",
		7626: "DEP",
		9443: "DEPO",
		8342: "DERA",
		8598: "DEXON",
		2484: "DEXTER",
		180:  "DFI",
		2969: "DFM",
		2977: "DG2L",
		2674: "DGSTATION",
		6151: "DH",
		7898: "DHC",
		1520: "DHD",
		9016: "DIAODIAOTechnology",
		7242: "DIGALOG",
		2000: "DIGINICS",
		3943: "DINEC",
		8374: "DIRECTV",
		8161: "DK",
		3470: "DKT",
		8847: "DLOGIC",
		9265: "DLX",
		8505: "DMATEK",
		3734: "DMP",
		397:  "DNE",
		4667: "DO-Monix",
		5671: "DPAC",
		2355: "DPS",
		9368: "DQ",
		5031: "DResearch",
		4699: "DS",
		5054: "DSA",
		5324: "DSC",
		5043: "DSG",
		4168: "DSP",
		7100: "DSPECIALISTS",
		6237: "DSPWorks",
		7542: "DSSD",
		7200: "DT",
		4894: "DTI",
		2169: "DUALi",
		6111: "DUX",
		799:  "DVC",
		161:  "DVICO",
		6737: "DVL",
		9304: "DVN",
		5360: "DVS",
		5626: "DVT",
		7447: "DWnet",
		5510: "DY-4",
		6204: "DZS",
		1814: "DaTARIUS",
		5515: "Dacoll",
		6778: "Dadoutek",
		1099: "Daehanet",
		742:  "Daeryung",
		9430: "Daeshin-Information",
		5618: "Daewoo",
		6593: "Daewoois",
		6453: "Dags",
		504:  "Daiden",
		1085: "Daihen",
		1604: "Daiichi",
		4124: "Daiichi-Dentsu",
		5484: "Daikin",
		4446: "Daintree",
		1465: "Daisy",
		1454: "Daktronics",
		1605: "Dallmeier",
		7956: "Damalisk",
		3057: "Dametric",
		80:   "Dana",
		7819: "Danal",
		784:  "Danam",
		2534: "Danelec",
		1096: "Danfoss",
		2873: "Daniels",
		6396: "Danlaw",
		3442: "Danpex",
		3130: "Dansensor",
		4280: "Dantel",
		9070: "Daontec",
		422:  "Daphne",
		8087: "DarbeeVision",
		3233: "Darts",
		6068: "Dasan",
		111:  "Dassault",
		2569: "Data",
		7204: "DataCore",
		4795: "DataFab",
		8815: "DataGravity",
		1139: "DataLogic",
		1506: "DataPower",
		8924: "DataRemote",
		2266: "DataWind",
		116:  "Datability",
		2213: "Datacap",
		5066: "Datacom",
		5033: "Datacore",
		5889: "Dataexpert",
		9201: "Datafox",
		5430: "Datafusion",
		4311: "Dataline",
		5880: "Datalux",
		1937: "Datamax",
		5511: "Datamedia",
		4203: "Datametrics",
		5757: "Dataplex",
		10:   "Datapoint",
		1970: "Dataprobe",
		5867: "Dataproducts",
		4815: "Dataram",
		953:  "Datasound",
		2909: "Datastore",
		4627: "Datastrip",
		5110: "Datatech",
		4279: "Datatrek",
		3721: "Dataupia",
		2191: "Datawire",
		6203: "Datax",
		7674: "Datecs",
		1732: "Datel",
		1022: "Dateno",
		5219: "Datong",
		8432: "Datrium",
		3450: "Datum",
		5526: "Datus",
		5957: "Dauphin",
		2986: "Dave",
		5841: "David",
		3720: "Daviscomms",
		1211: "Davolink",
		3305: "Dawevision",
		5753: "Dawn",
		4843: "DaySequerra",
		5463: "Dayna",
		8120: "Dayouplus",
		4365: "Dbii",
		2414: "Dbtel",
		5258: "Dctri",
		8435: "DeLaval",
		7696: "Decatur",
		6221: "Decision",
		2061: "Decru",
		9393: "Dediprog",
		7433: "Deeplet",
		3688: "Deepsound",
		3193: "Defidev",
		8275: "Definium",
		2812: "Deicy",
		4857: "Deif",
		7639: "Dejai",
		5615: "Delem",
		102:  "Dell",
		6826: "DellKing",
		3586: "Delorme",
		8971: "Delphian",
		7509: "Delphin",
		957:  "Delta",
		4520: "Deltacom",
		7458: "Deltanet",
		3528: "Deltanode",
		9437: "Deltenna",
		689:  "Deltron",
		3461: "Demant",
		6792: "Dematic",
		4604: "Demco",
		131:  "Densan",
		5890: "Denso",
		2212: "Dentum",
		9413: "Denwa",
		815:  "Deonet",
		8687: "Derek",
		7103: "DerekLimited",
		592:  "Desana",
		7225: "Desay",
		3862: "Design",
		4648: "Design-Com",
		4664: "DesignArt",
		6126: "DeskStation",
		5764: "Desknet",
		9057: "Detracom",
		7550: "Deutron",
		3011: "Develco",
		133:  "Develcon",
		7955: "Devialet",
		6892: "DeviceDesign",
		2066: "Devicescape",
		6748: "Devlin",
		6846: "Devline",
		7058: "Dewar",
		7751: "Dewav",
		7070: "Dexatek",
		5611: "Dexdyne",
		7583: "Dexin",
		3975: "DiMoto",
		2333: "Diablo",
		1005: "Diagraph",
		3807: "Dial",
		2245: "Dialog",
		5354: "Dialogic",
		1369: "Dialogue",
		9218: "Diamanti",
		1025: "Diamond",
		6095: "Diba",
		1303: "Dibal",
		3105: "Diboss",
		235:  "Dica",
		6085: "Dictaphone",
		511:  "Diebold",
		9370: "Diffon",
		5168: "Digalog",
		968:  "Digeo",
		5650: "Digi-Data",
		5140: "DigiBoard",
		911:  "DigiPower",
		1981: "DigiRose",
		5284: "Digianswer",
		6212: "Digicom",
		2144: "Digicube",
		6449: "Digience",
		3969: "Digifriends",
		6018: "Digigram",
		3304: "Digilent",
		8169: "Digimore",
		2502: "Diginfo",
		7689: "Digiquest",
		3177: "Digit",
		6798: "Digita",
		1565: "Digital",
		5646: "DigitalScape",
		515:  "DigitalSis",
		2901: "DigitalZone",
		4698: "Digitalbox",
		5636: "Digitalcast",
		2143: "Digitalks",
		9161: "Digitalwatt",
		4836: "Digitec",
		5732: "Digitech",
		496:  "Digitel",
		3913: "Digitize",
		4973: "Digitra",
		3739: "Digitrax",
		4082: "Digitview",
		2269: "Digium",
		2851: "Digiwell",
		3740: "Dignsys",
		2976: "Dilithium",
		9434: "Dinstar",
		6215: "Dionex",
		3568: "Diostech",
		7811: "DirectPacket",
		794:  "Disco",
		6891: "Discovergy",
		8314: "Discovery",
		1238: "Dish",
		1192: "Diskbank",
		2776: "Diskware",
		1907: "DispenseSource",
		8471: "Displaire",
		8860: "Display",
		2233: "DisplayLink",
		8291: "Disruptive",
		5967: "Ditech",
		1922: "DivR",
		7735: "DivUS",
		1209: "DivergeNet",
		8223: "Diversey",
		4181: "Diversified",
		4935: "Divio",
		4532: "Dizipia",
		8205: "Dlogixs",
		7965: "Dmet",
		2337: "Dnpg",
		3438: "DoCoMo",
		316:  "DoTop",
		4045: "Doble",
		5545: "Document",
		5899: "Docupoint",
		6002: "Dolby",
		5580: "Domex",
		2481: "Domo",
		5228: "Doms",
		3539: "Dongahelecomm",
		7961: "Dongnian",
		3107: "Donjin",
		6077: "Dooin",
		3440: "Doorking",
		8548: "Doppler",
		2439: "Doremi",
		3874: "Doro",
		4610: "Double-Take",
		4255: "Dovatron",
		118:  "Dove",
		8458: "Dragino",
		1091: "DragonWave",
		8925: "Dragonchip",
		7439: "Dragos",
		3923: "DrayTek",
		3029: "Dream",
		1346: "Dream-Multimedia-Tv",
		3370: "Dreamtech",
		5136: "Dressler",
		733:  "Drew",
		7312: "Drimsys",
		3135: "DriveCam",
		7644: "DriveScale",
		8030: "Drivenets",
		8517: "Drivven",
		8724: "Drogoo",
		6920: "Dropcam",
		8842: "Drtech",
		7431: "Drust",
		7579: "Dryad",
		331:  "Dtvro",
		5922: "Dual",
		7192: "DualShine",
		3400: "Duaxes",
		8949: "Ducere",
		4878: "Duelco",
		9168: "Dunkermotoren",
		6494: "DuroByte",
		6509: "Duskrise",
		3170: "Dust",
		8728: "Dusun",
		9308: "Dycon",
		4674: "Dycor",
		4100: "Dyna",
		6042: "Dynacolor",
		2725: "Dynalab",
		7355: "Dynalec",
		1941: "Dynamic",
		6555: "Dynapower",
		5049: "Dynapro",
		5676: "Dynarc",
		5452: "Dynatech",
		5878: "Dynatem",
		1987: "Dynavac",
		7190: "Dyson",
		7923: "E-Band",
		248:  "E-Control",
		7757: "E-Domus",
		8579: "E-Fuel",
		476:  "E-Globaledge",
		8209: "E-Lead",
		5172: "E-M",
		4574: "E-Mon",
		8607: "E-Prime",
		4028: "E-Senza",
		5111: "E-Systems./Garland",
		2714: "E-TEC",
		7917: "E-TRON",
		5635: "E-Tech",
		455:  "E.D.&A",
		6117: "E.E.P.D",
		2244: "E2O",
		2620: "E2S",
		7169: "E3",
		2706: "EAB/RWI/K",
		4391: "EASY3CALL",
		7377: "EBN",
		4179: "EBRAINS",
		3000: "ECA-Sinters",
		5610: "ECCS",
		3768: "ECKey",
		2680: "ECM",
		8444: "ECOtality",
		6843: "EControls",
		3632: "ED",
		6457: "EDIC",
		3247: "EDM",
		6662: "EDMI",
		2341: "EDNT",
		3466: "EDO-EVI",
		229:  "EDSL",
		3317: "EDSLAB",
		4717: "EDT",
		1831: "EE",
		6178: "EES",
		1252: "EFM",
		1590: "EG",
		3033: "EGO",
		6680: "EID",
		2308: "EION",
		5202: "EIS",
		5647: "EIZO",
		8234: "EKE",
		4363: "EKE-Electronics",
		3337: "EL-Tech",
		645:  "ELANsat",
		6078: "ELE-Chem",
		715:  "ELESIGN",
		8007: "ELFTECH",
		1771: "ELIAS",
		7460: "ELL-IoT",
		1738: "ELM",
		214:  "ELMEX",
		2633: "ELPRO",
		7661: "ELS-GmbH",
		1245: "EM",
		8741: "EM-Tech",
		6824: "EMAC",
		4352: "EMC",
		445:  "EMPEG",
		9391: "EMTAKE",
		6479: "EMU",
		8288: "EMW",
		4328: "EN",
		611:  "ENEGATE",
		3499: "ENENSYS",
		8690: "ENERES",
		7940: "ENMAS",
		8372: "ENVINET",
		8416: "EOC",
		126:  "EON",
		5947: "EOS",
		3366: "EOps",
		7740: "EPCOMM",
		2434: "EPIN",
		3375: "EPL",
		671:  "EPOX",
		4584: "EQ-Sys",
		9156: "EQUES",
		6948: "ER-Telecom",
		5499: "ERI",
		2579: "ERUNE",
		4680: "ESCATRONIC",
		340:  "ESD",
		8157: "ESG",
		4964: "ESI",
		4386: "ESP",
		9390: "ESPOD",
		249:  "ESS",
		9135: "ESSENCORE",
		5710: "EST",
		3457: "ESTIC",
		2876: "ESU",
		7362: "ESYLUX",
		1674: "ET&T",
		451:  "ETAS",
		581:  "ETEN",
		7832: "ETH",
		4055: "ETL",
		6452: "EUCAST",
		2348: "EUREM",
		7945: "EVBox",
		4141: "EVGA",
		282:  "EVR",
		7457: "EVRsafe",
		8357: "EVUlution",
		4390: "EW3",
		2268: "EWA",
		5426: "EXP",
		8926: "EXTEN",
		6100: "EXXACT",
		5836: "Eacem",
		4268: "Eagle",
		6499: "Earda",
		3212: "Earforce",
		8777: "EarthCam",
		7749: "Eastcompeace",
		4754: "Eastern",
		1413: "Eastmode",
		6456: "Eastriver",
		7872: "EasySYNC",
		1855: "Eaton",
		1818: "Eazix",
		7849: "Ebay",
		1841: "Ebtron",
		2093: "Ecastle",
		5879: "Ecci",
		5790: "Ecessa",
		6027: "Echelon",
		3744: "Echo360",
		6951: "Echodyne",
		2637: "Echolab",
		9333: "Echosens",
		9372: "Echowell",
		8510: "Eclipse",
		9445: "EcoTech",
		7367: "Ecocom",
		3964: "Ecolab",
		9075: "Ecoliv",
		2761: "Ecom",
		6958: "Ecosense",
		5015: "Ectel",
		6004: "Ecton",
		2817: "Edata",
		5316: "Edec",
		8298: "Edeltech",
		5236: "Edeva",
		2843: "Edge",
		7776: "EdgeCore",
		6144: "EdgePoint",
		9482: "EdgeQ",
		4154: "EdgeVelocity",
		1165: "EdgeWave",
		6295: "Edgecore",
		7390: "Edgewater",
		6446: "Edifier",
		115:  "Edimax",
		7827: "Edison",
		5621: "Edmi",
		9214: "Edup",
		1805: "Edwards",
		4281: "Efficient",
		412:  "Effinet",
		3692: "Efkon",
		7332: "Egardia",
		883:  "Egenera",
		2158: "Ego",
		8987: "Eidetic",
		3157: "Eidicom",
		6979: "Eidolon",
		2471: "Eidsvoll",
		1629: "Eiki",
		6872: "Eintechno",
		3095: "Eishin",
		6386: "Eito",
		2806: "Eka",
		3341: "Ekahau",
		4967: "Elastic",
		8720: "Elastifile",
		6705: "Elatec",
		627:  "Elau",
		4466: "Elbit",
		5411: "Eldat",
		5694: "Elecom",
		1250: "Elecs",
		6603: "Elecsys",
		210:  "Electro",
		5058: "Electro-Metrics",
		7174: "Electrolux",
		3691: "Electronic",
		6850: "Electronics",
		6176: "Electrosonic",
		3141: "Eleksen",
		3963: "Elelux",
		3765: "Element",
		4653: "Elentec",
		3222: "Elesta",
		8554: "Elesys",
		1172: "Eletex",
		2518: "Elexol",
		2185: "Elextech",
		8436: "Elgama-Elektronika",
		1023: "Elgar",
		8560: "Elim",
		8608: "Eline",
		6673: "EliteGroup",
		1115: "Elitegroup",
		8268: "Elma",
		5412: "Elmic",
		1404: "Elmo",
		7897: "Elno",
		5892: "Elonex",
		2054: "Elphel",
		8710: "Elprotronic",
		1106: "Elrest",
		3563: "Elster",
		8497: "Eltex",
		7520: "EltexAlatau",
		8219: "Elvaco",
		6354: "Elxsi",
		8614: "EmBestor",
		8121: "EmFirst",
		5725: "Email",
		8012: "Embed",
		7495: "EmbedWay",
		6469: "Embedian",
		522:  "Embedone",
		1893: "Embedtronics",
		1936: "Ember",
		8045: "Embertec",
		8877: "Embrane",
		5882: "Emcom",
		1980: "Emcore",
		7102: "Emcraft",
		1195: "EmergeCore",
		3793: "Emergent",
		1965: "Emerging",
		4060: "Emfinity",
		3791: "Emfit",
		6270: "Emicon",
		3906: "Emitech",
		3021: "Emitor",
		8290: "Emizon",
		3808: "Empacket",
		880:  "Empirix",
		7267: "Emplus",
		2823: "Empower",
		9364: "Emstone",
		5406: "Emtrakorporated",
		8950: "Emtronix",
		129:  "Emulex",
		6061: "Emutec",
		1646: "Emuzed",
		8105: "EnGenius",
		7072: "EnTek",
		4813: "Enabling",
		4774: "Enalasys",
		9009: "Enatel",
		2304: "Encanto",
		8144: "Encell",
		6733: "Enclustra",
		3303: "Encore",
		433:  "EndPoints",
		2140: "EndRun",
		2086: "Endace",
		2173: "Endeleo",
		2612: "Endevco",
		6978: "Enecsys",
		8263: "Enelps",
		7405: "EnerAccess",
		1260: "EnerLinx.com",
		2598: "Enerdyne",
		9110: "Energica",
		9233: "Energotest",
		4728: "Energy",
		4883: "EnergyHub",
		4626: "EnergyICT",
		3252: "Enermet",
		6759: "Enernet",
		1502: "Enerpoint",
		9497: "Enervalley",
		5961: "Engage",
		2694: "Engim",
		5396: "Enginuity",
		7481: "Enguity",
		179:  "EnjoyWeb",
		7960: "Enmotus",
		200:  "Ennovate",
		9037: "Eno",
		204:  "Ensemble",
		4394: "Enseo",
		8629: "Ensequence",
		5039: "Ensim",
		8519: "Enspert",
		986:  "Ensure",
		2891: "Enswer",
		312:  "Enterasys",
		7924: "Entertainment",
		6561: "Entis",
		1422: "Entise",
		596:  "Entone",
		4434: "Entorian",
		4726: "Entourage",
		565:  "Entra",
		6182: "Entrada",
		2384: "Entrata",
		5302: "Entrega",
		2190: "Entrelogic",
		5613: "Entridia",
		1221: "Entrisphere",
		1394: "Entropic",
		6751: "Entropix",
		1298: "Envenergy",
		8726: "Environics",
		3850: "Environnement",
		3756: "Envisacor",
		4911: "Envisionnovation",
		8315: "EnvyLogic",
		1772: "Enwiser",
		3814: "Enzytek",
		6500: "EoCell",
		6875: "Eolo",
		1129: "Eolring",
		7801: "Eoptolink",
		9465: "EosTek",
		8748: "Eoslink",
		7756: "EpSpot",
		8766: "Epec",
		4050: "Epic",
		2113: "Epicenter",
		2670: "Epicom",
		5614: "Epigram",
		6071: "Epilog",
		1332: "EpoX",
		46:   "Epson",
		1425: "Epygi",
		1393: "EqualLogic",
		3675: "Equaline",
		1131: "Equator",
		4986: "Equiinet",
		5509: "Equinox",
		159:  "Equip'Trans",
		461:  "Equipe",
		7154: "Equitech",
		718:  "Equitrac",
		3487: "Equustek",
		8391: "EraThink",
		1521: "Eracom",
		3369: "Erae",
		7647: "Ereca",
		7850: "Ergophone",
		7693: "Erhardt+Leimer",
		306:  "Ericsson",
		8569: "Ericsson-LG",
		2015: "Ermme",
		4253: "Ersat",
		1270: "Erskine",
		7323: "Es-tech",
		4677: "Esab",
		9294: "Esan",
		5689: "Escalate",
		3389: "Escape",
		4323: "Escherlogic",
		9040: "Escort",
		7033: "Espec",
		1188: "Espera-Werke",
		6829: "Espotel",
		6377: "Espressif",
		5460: "Esprit",
		9337: "Essec",
		6630: "Essel-T",
		4850: "Essensium",
		5759: "Essential",
		3523: "Essilor",
		7973: "Esson",
		6831: "Essys",
		2540: "Estari",
		7560: "Estech",
		1318: "Esteem",
		6817: "Etek",
		7262: "Eterna",
		6181: "EtherWAN",
		5805: "Ethercom",
		4092: "Etherstack",
		1491: "Etherstuff",
		7688: "Ethertronics",
		7241: "EthoSwitch",
		3593: "Ethos",
		1712: "Etin",
		5643: "Etrend",
		7065: "Etronic",
		3443: "Etrovision",
		8366: "Ettus",
		3269: "Etymotic",
		2863: "Eubus",
		3557: "Euchner+Co",
		7388: "Euclid",
		8929: "Eukrea",
		9409: "Eumtech",
		3802: "Eunicorn",
		4858: "Euphonic",
		5673: "Euphonix",
		982:  "Euracom",
		9464: "Eurecam",
		2532: "Eurilogic",
		3102: "EuroCB",
		3656: "EuroTel",
		2262: "Eurocom",
		4770: "Eurodesign",
		502:  "Eurologic",
		7016: "Euronda",
		2249: "Europlex",
		5314: "Eurotech",
		1543: "Eurotherm",
		4200: "Eurotime",
		6545: "Eutronix",
		8242: "Evantage",
		1804: "Eve",
		3036: "Eventide",
		2433: "EverFocus",
		2098: "Everbee",
		6976: "Everest",
		929:  "Everex",
		4226: "Evergreen",
		9468: "Everspin",
		4797: "Everspring",
		6970: "Evervictory",
		7223: "Everysight",
		3628: "Evolis",
		8126: "Evoluzn",
		7848: "Evrisko",
		1477: "ExPet",
		3811: "ExaDigm",
		7586: "Exablaze",
		8620: "Exablox",
		6036: "Exabyte",
		3167: "Exalt",
		5849: "Exar",
		3019: "Exartech",
		2745: "Exavera",
		4185: "Excel",
		6306: "Excelan",
		6198: "Excellent",
		3800: "Excelpoint",
		3840: "Exegin",
		2400: "Exelis",
		4691: "Exeltech",
		3083: "Exeo",
		466:  "Exfo",
		447:  "ExiO",
		4925: "Expand",
		6331: "ExperData",
		3324: "Expertise",
		9106: "Explora",
		9140: "Exponent",
		2926: "Extandon",
		5114: "Extended",
		5126: "Extension",
		731:  "Extenway",
		3289: "Exterity",
		169:  "Extratech",
		185:  "Extreme",
		2122: "ExtremeSpeed",
		2752: "Extricom",
		847:  "Extron",
		6305: "Exxon",
		1725: "EyeCross",
		3314: "EyeFi",
		3902: "Eyeheight",
		4331: "Eyeview",
		2132: "Ezquest",
		3126: "Ezurio",
		9297: "F-Secure",
		3397: "F1MEDIA",
		492:  "F3",
		289:  "F5",
		3276: "FA",
		7530: "FACTS",
		9257: "FADU",
		9359: "FAG",
		790:  "FAST",
		9284: "FCA",
		5272: "FDK",
		9374: "FDT",
		1557: "FEI",
		3500: "FEI-Zyfer",
		3817: "FEIG",
		3592: "FFEI",
		9220: "FIRS",
		6482: "FLECTRON",
		3724: "FLIR",
		2243: "FMN",
		6399: "FMTech",
		6441: "FN-Link",
		7184: "FONsystem",
		2382: "FOR-A",
		7528: "FORICS",
		3336: "FORMOSA21",
		7737: "FRAMOS",
		4823: "FRC",
		2636: "FSI",
		2207: "FTA",
		3086: "FUHO",
		7873: "FUJITU",
		9274: "FUSION",
		3185: "FXC",
		9023: "FXI",
		708:  "FabiaTech",
		1966: "Fabric7",
		1065: "Fabricom",
		7235: "Facebook",
		1956: "Factum",
		6856: "Fagor",
		8775: "Fairfield",
		8056: "Fairphone",
		1940: "FalconStor",
		1444: "Falt",
		4444: "FamilyPhone",
		7021: "Fanhattan",
		3281: "Fanstel",
		5553: "Fantum",
		6428: "Fanvil",
		8166: "Far-sighted",
		2691: "Fargo",
		8798: "Farlink",
		7165: "Farmage",
		298:  "Fast",
		6609: "Fastback",
		3541: "Faster",
		541:  "Fastfame",
		2554: "Fastrax",
		1265: "Fastwel",
		9470: "Favite",
		7587: "Feitian",
		328:  "Fenecom",
		8595: "Fengfan",
		6791: "Fermax",
		5489: "Fermilab",
		6322: "Ferranti",
		2131: "Festo",
		4757: "FiberPlex",
		4368: "Fiberblaze",
		5508: "Fibercom",
		628:  "Fibercycle",
		5953: "Fiberdata",
		8277: "Fibergate",
		4922: "Fiberhome",
		2396: "Fiberlane",
		5486: "Fibermux",
		5158: "Fibernet",
		7071: "Fiberpro",
		5663: "Fibex",
		7414: "Fibrain",
		8413: "Fibrlink",
		821:  "FibroLAN",
		3:    "Fibronics",
		6656: "Ficer",
		7641: "Ficosa",
		8187: "Fida",
		9000: "Fidelix",
		9483: "Fidus",
		4021: "Fidustron",
		6299: "Fihonest",
		7030: "Fijowave",
		8981: "Fike",
		6046: "Filanet",
		6316: "Filenet",
		1360: "FillFactory",
		8909: "Filmetrics",
		2112: "Filtronic",
		5895: "Fine-PAL",
		4238: "Finecom",
		1056: "Finedigital",
		8771: "Finis",
		5623: "Finisar",
		1892: "Finlux",
		3514: "Finnzymes",
		8224: "Finsecur",
		3686: "Finsoft",
		546:  "FireBrick",
		8805: "Fireflies",
		5715: "Firepower",
		3362: "Firetide",
		6722: "Firewalla",
		208:  "Firewiredirect.com",
		8461: "Firich",
		8103: "FirmTek",
		1979: "First",
		1330: "Firstech",
		6974: "Fischer",
		4491: "Fisher-Rosemount",
		6521: "Fisys",
		6588: "Fitbit",
		7163: "Fitview",
		7067: "FiveCo",
		5992: "Fivemere",
		2669: "Flaircomm",
		1072: "Flarion",
		1029: "Flash",
		7244: "Flashbay",
		8410: "FlatFrog",
		409:  "Flatstack",
		3085: "Flex-P",
		3759: "FlexRadio",
		3015: "FlexiPanel",
		1100: "Flexlight",
		7018: "Flexoptix",
		833:  "Flextronics",
		984:  "Flexus",
		8180: "Flonidan",
		3014: "Flovel",
		4228: "Flowpoint",
		5099: "Fluent",
		8292: "Flyaudio",
		5425: "Flytech",
		1042: "Fnet",
		6123: "Focon",
		6633: "Fon",
		8735: "FonSee",
		5908: "Fonsys",
		3596: "Fontal",
		2272: "Force",
		303:  "Force10",
		4881: "Ford",
		5361: "Fore",
		5251: "Foresson",
		7976: "ForgetBox",
		5092: "Forks",
		6488: "FormericaOE",
		9501: "Formike",
		6999: "Formlabs",
		8754: "Fort-Telecom",
		2514: "Fortelink",
		4484: "Fortex",
		6878: "Fortify",
		1323: "Fortinet",
		9432: "Fortress",
		2506: "Fortuna",
		3496: "Fortunetek",
		8004: "Forworld",
		950:  "Fostex",
		5606: "Fountain",
		7838: "Four",
		4519: "Fourier",
		5423: "Fourthtrack",
		6040: "Foveon",
		437:  "FoxJet",
		221:  "Foxconn",
		7314: "Foxda",
		6704: "Foxtech",
		2468: "Francotyp-Postalia",
		2429: "Franklin",
		1720: "Free2move",
		9388: "FreeBit",
		1870: "FreeHand",
		729:  "FreeMs",
		8475: "FreeTek",
		1164: "FreeWave",
		293:  "Freecom",
		3196: "Freedom9",
		6107: "Freegate",
		4409: "Freegene",
		7959: "Freestyle",
		270:  "Frequentis",
		2686: "Fresenius-Vial",
		7286: "Fresh",
		6775: "Friem",
		4059: "From2",
		5488: "Frontier",
		6802: "Frontiir",
		4622: "FuGang",
		1050: "Fujant",
		6987: "Fujikon",
		1698: "Fujikura",
		2139: "Fujinon",
		4:    "Fujitsu",
		8576: "Fulan",
		1773: "FullWave",
		7419: "Fullpower",
		262:  "Fulltek",
		5105: "Funasset",
		5934: "Funk",
		8482: "Furrion",
		3990: "Further",
		4231: "Fusion",
		8644: "Fusion-io",
		3581: "FusionDynamic",
		3088: "Fusiontech",
		9321: "Futaba",
		8558: "Futaba-Kikaku",
		3660: "Futarque",
		8606: "Futecho",
		3150: "Futronic",
		7675: "Futura",
		5106: "Future",
		4419: "FutureLogic",
		878:  "FutureSmart",
		1571: "Futuretel",
		7311: "Fuze",
		5241: "G-Connect",
		9452: "G-Lab",
		2292: "G-PRO",
		8175: "G-Printec",
		1620: "G-Star",
		3747: "G-Technology",
		2040: "G-Tek",
		9498: "G-Wearables",
		6206: "G2",
		6627: "G2C",
		206:  "G3M",
		3251: "GAI-Tronics",
		7905: "GBO",
		4451: "GBS",
		5929: "GCC",
		4291: "GDE",
		4153: "GDI",
		3291: "GDX",
		5811: "GE",
		6364: "GECO",
		1564: "GENETEC",
		1105: "GERSTEL",
		2578: "GES",
		6412: "GET",
		5383: "GEW",
		4562: "GGH",
		8195: "GH",
		4298: "GHI",
		8628: "GHIA",
		6787: "GHT",
		9264: "GIRD",
		9270: "GIT",
		7562: "GITSN",
		5084: "GK",
		4372: "GL",
		4697: "GLOBAL",
		1191: "GM-2",
		4808: "GMK",
		5556: "GMX/Gimix",
		2289: "GN&S",
		2312: "GNP",
		9221: "GNTEK",
		8587: "GOEFER",
		6508: "GOLD3LINK",
		8827: "GOPEACE",
		8019: "GP",
		8088: "GSI",
		5756: "GSM-Syntel",
		3478: "GSP",
		9320: "GST",
		2826: "GSTeletech",
		3526: "GT&T",
		5925: "GVC",
		2344: "GVN",
		2282: "GW",
		8611: "GX",
		3287: "GZ",
		6109: "Gadzoox",
		7526: "Gafachi",
		3936: "GainSpan",
		6861: "Gainspeed",
		5631: "Gaio",
		1447: "Galaxis",
		3848: "Galaxy",
		1921: "Galazar",
		5996: "Galileo",
		8990: "Galleon",
		8300: "Galore",
		1175: "Galtronics",
		4369: "Galvanic",
		1212: "Gamatronic",
		5926: "Gambit",
		8083: "Game",
		4221: "Gammadata",
		9317: "Garland",
		797:  "Garmin",
		5037: "Garnet",
		7276: "Garo",
		4172: "Garrett",
		4222: "GarrettCom",
		1840: "Garuda",
		7743: "Gastron",
		5278: "Gatan",
		2668: "GateConnect",
		2203: "GateWare",
		8174: "Gatekeeper",
		7935: "GatesAir",
		64:   "Gateway",
		5974: "Gateworks",
		3296: "Gatsometer",
		2090: "Gawell",
		1144: "Gcom",
		9006: "Gearlinx",
		9187: "Geberit",
		4916: "Geenovo",
		3809: "Gefen",
		323:  "Gefran",
		7477: "Gehirn",
		2064: "GemWon",
		8098: "Gembird",
		5279: "Gemflex",
		8091: "Gemicom",
		8269: "Geminico",
		9154: "Gemintek",
		1747: "Gemstone",
		1450: "Gemtek",
		6803: "Gen2wave",
		8512: "Genelec",
		3774: "Gener8",
		1854: "Genera",
		2751: "General",
		7687: "Generalplus",
		8097: "Generiton",
		5850: "Genetec",
		931:  "Genetel",
		3801: "Genew",
		2196: "Genexis",
		5206: "Genicom",
		4765: "Genie",
		4205: "Genitech",
		1589: "Gennum",
		6043: "Genoa",
		6780: "Genoray",
		839:  "Genotech",
		5208: "Genrad",
		6187: "Genroco",
		1522: "GentechMedia",
		3525: "Gentex/Electro-Acoustic",
		916:  "Gentner",
		7815: "Gentrice",
		4724: "Genuine",
		2781: "GeoVision",
		7673: "GeoWAN",
		9212: "Geodesic",
		1778: "Geomation",
		891:  "Geospace",
		7978: "Germaneers",
		5952: "Gespac",
		8447: "Gessler",
		9108: "Getck",
		3068: "Geutebruck",
		368:  "Geyser",
		8198: "Gheo",
		1380: "Giant",
		722:  "Giantec",
		7126: "Gifa",
		1929: "Giga-Byte",
		2793: "Giga-byte",
		3152: "Gigabeam",
		2235: "Gigabit",
		4057: "Gigafin",
		7361: "Gigafirm",
		5364: "Gigalabs",
		273:  "Gigalink",
		3421: "Gigamips",
		3924: "Gigamon",
		1310: "Gigaphoton",
		4299: "Gigaset",
		8762: "Gigastone",
		7378: "Gigawave",
		2225: "Gigawavetech",
		203:  "Gilat",
		5252: "Gilbarco",
		7894: "Gimasi",
		8445: "Gimbal",
		2030: "Gincom",
		1134: "Ginganet",
		7847: "Gint",
		4668: "Ginzinger",
		9273: "Giroptic",
		2611: "Gizmondo",
		4397: "Glensound",
		1765: "Glimmerglass",
		3184: "Global",
		6801: "GlobalBeiMing",
		641:  "GlobalStreams",
		1389: "GlobalTop",
		5852: "Globalnet",
		5262: "Globaloop",
		1975: "Globalsat",
		9323: "Globalscale",
		6255: "Globalstar",
		400:  "Globetek",
		2870: "Globo",
		4607: "Glocom",
		743:  "Glonet",
		5739: "Glory",
		4513: "Gloscom",
		7432: "Glovast",
		7702: "Glowforge",
		1782: "Glsystech",
		1182: "Gluon",
		1358: "Glyph",
		8053: "Gnodal",
		2805: "Go",
		6242: "GoPro",
		8136: "Goden",
		3917: "Godex",
		969:  "Goepel",
		6367: "Goertek",
		1120: "Golden",
		5142: "Goldstar",
		4357: "Goliath",
		7603: "GooWi",
		2460: "GoodMan",
		4132: "Goodmill",
		3522: "Google",
		8686: "Googol",
		4510: "Gorba",
		2647: "Gossen-Metrawatt-GmbH",
		6718: "Gotech",
		166:  "Gotham",
		6334: "Gould",
		4440: "Gowell",
		8875: "Goyoo",
		8389: "Gpms",
		7014: "Grabango",
		5326: "Gradient",
		4679: "Graf-Syteco",
		4914: "Grainmustards",
		4065: "Granch",
		1883: "Grand",
		6938: "Grandbeing",
		7120: "Grandex",
		2453: "Grandeye",
		1683: "Grandstream",
		2800: "Grandtec",
		9224: "Graphiant",
		8539: "Graphite",
		5554: "Graphon",
		530:  "Graphtec",
		2305: "Grayhill",
		8172: "Great",
		9423: "Greatek",
		6406: "Green",
		9138: "GreenBytes",
		4018: "GreenLine",
		1514: "GreenNET",
		2972: "GreenPeak",
		9102: "GreenPriz",
		2165: "Greenbell",
		7339: "Greenlee",
		8530: "Greenliant",
		8043: "Greenvity",
		7215: "Greenwald",
		142:  "Grid",
		6806: "GridCentric",
		4554: "GridIron",
		4681: "GridPoint",
		6531: "Gridco",
		9189: "Gridlink",
		3454: "Gridpoint",
		6380: "Gridstore",
		7503: "Gridwiz",
		2615: "Griffin",
		2398: "Grips",
		7009: "Grotech",
		2834: "Grundfos",
		6073: "Grundig",
		4283: "Gtech",
		4129: "Gtran",
		3967: "Gtri",
		4521: "Guangda",
		1542: "Guardware",
		3418: "Gude",
		4546: "Gulfstream",
		3016: "Gumstix",
		7122: "GuoTengShengHua",
		1510: "GyroSignal",
		5513: "H-Three",
		6383: "H2A",
		9079: "HAKKO",
		5138: "HAL",
		6988: "HAN",
		1412: "HARTEC",
		1596: "HARTING",
		2097: "HASHIMOTO",
		1287: "HASSNET",
		9322: "HBC-radiomatic",
		1438: "HBrain",
		71:   "HCL",
		9082: "HCT",
		1495: "HCV",
		4098: "HD",
		8194: "HDR10+",
		5371: "HE",
		2037: "HEINESYS",
		9366: "HENGBAO",
		9191: "HES-SO",
		2881: "HF",
		6484: "HFC",
		4576: "HFR",
		1862: "HGST",
		8498: "HI",
		2912: "HI-P",
		980:  "HID",
		1731: "HIMS",
		1268: "HIT",
		3056: "HIVION",
		9014: "HIWIFI",
		3863: "HM",
		5847: "HME",
		837:  "HMS",
		7561: "HNS",
		53:   "HOB",
		3714: "HOLUX",
		779:  "HORIBA",
		1288: "HOW",
		1627: "HOWTEL",
		416:  "HOYA",
		302:  "HP",
		7296: "HRD",
		5331: "HRK",
		5377: "HT",
		1341: "HTC",
		8207: "HUG-Witschi",
		531:  "HUMAX",
		1753: "HUMEX",
		5932: "HVE",
		9177: "HW",
		6789: "HWH",
		848:  "HYPERCHIP",
		3385: "Hacetron",
		2024: "Hach",
		5902: "Haft",
		4687: "Hager",
		5335: "Hagiwara",
		4768: "Haier",
		7680: "Hailo",
		4684: "Hainzl",
		7155: "Hajen",
		5307: "Hakko",
		5293: "Hakusan",
		4075: "Hakusan.Mfg",
		2786: "Halcro",
		8851: "Halfa",
		1405: "Haliplex",
		8316: "Hame",
		6944: "Hamee",
		7280: "Hamilton",
		1912: "Hammerhead",
		5581: "HanA",
		3179: "Hanazeder",
		1700: "Hanback",
		1216: "Hanbit",
		8473: "HanbitEDS",
		2527: "HandEra",
		3521: "HandHeld",
		8667: "Handaer",
		5584: "Handlink",
		3622: "Handreamnet",
		6903: "Hanilstm",
		4142: "Hanlong",
		2029: "Hannae",
		9291: "Hannto",
		9245: "Hansgrohe",
		4798: "Hanshinit",
		3955: "Hanson",
		3346: "Hansun",
		7912: "Hansung",
		1796: "Hanwang",
		1588: "Happy",
		7787: "Happyelectronics",
		9013: "Harley-Davidson",
		1526: "Harmonic",
		4946: "Harmonix",
		124:  "Harris",
		4974: "Hasselblad",
		3401: "Hasware",
		5061: "Hatteland",
		9112: "Haverford",
		9380: "Hawa",
		7124: "Hawkeye",
		2038: "Hawking",
		8551: "Hay",
		5722: "Hazeltine",
		7480: "Hazemeyer",
		7462: "Hccp",
		8997: "Hdpro",
		6710: "Hdsn",
		9303: "Heat",
		2438: "Hectrix",
		2995: "Hectronic",
		246:  "HeiSei",
		2548: "Heidelberg",
		2839: "Heim",
		8464: "Heimgard",
		5845: "Heinzinger",
		6969: "Heinzmann",
		6099: "Heiwa",
		2879: "Helicomm",
		7578: "Heliospectra",
		2339: "Helioss",
		7540: "Helium",
		2159: "Helius",
		1322: "Helix",
		8371: "Helixtech",
		59:   "Hellige",
		7489: "Hello",
		6786: "Helmholz",
		8561: "Helvetia",
		1546: "HemoCue",
		6424: "Hengstler",
		1963: "Heraeus",
		8600: "Heran",
		5591: "Hermes",
		595:  "Hermstedt",
		6192: "Hero",
		5568: "Heurikon",
		8525: "Hexatronic",
		9083: "Heyrex",
		9129: "Hi-P",
		1373: "Hi-Techniques",
		8545: "Hi-flying",
		1600: "HiConnect",
		5828: "HiQ",
		6949: "HiTEM",
		9059: "Hichan",
		4337: "Hidea",
		2837: "Hifn",
		8409: "Hifocus",
		375:  "High",
		3424: "HighPoint",
		9440: "Highgates",
		7274: "Hightech",
		1277: "Hikari",
		5783: "Hilan",
		2892: "Hill-Rom",
		4766: "Hills",
		3778: "Hillstone",
		410:  "Hilscher",
		8218: "Himax",
		4603: "Hinke",
		915:  "Hinox",
		7883: "Hioso",
		1619: "Hirata",
		6597: "Hirose",
		4511: "Hirsch",
		2172: "Hisharp",
		89:   "Hitachi",
		1830: "Hitech",
		4982: "Hitex",
		1340: "Hitpoint",
		870:  "Hitron",
		6867: "HiveMotion",
		6268: "Hivesystem",
		4649: "Hivision",
		3084: "Hoatech",
		1556: "Hochiki",
		4643: "Hokkaido",
		3280: "Hokkei",
		419:  "Hokubu",
		8567: "Holl",
		6139: "Holontech",
		8845: "Holoplot",
		1207: "Homag",
		2480: "HomeLogic",
		1359: "Homenet",
		8210: "Homerider",
		7249: "Homewins",
		2713: "Honda",
		8449: "Honest",
		5829: "Honewell",
		934:  "Honeywell",
		5600: "Honeywell-Dating",
		4360: "Honeywld",
		5167: "Hong",
		7798: "Hongyu",
		2297: "Horanet",
		5266: "Horizon",
		2942: "Horoquartz",
		315:  "Horoscas",
		3695: "Hose-McCann",
		1730: "Hosiden",
		562:  "Hospira",
		6145: "Host",
		2336: "Hostlink",
		1628: "Hostnet",
		2641: "HotLava",
		2830: "Hotway",
		7196: "Hoya",
		4278: "Hsing",
		8652: "Htel",
		6577: "Hu",
		1453: "HuMANDATA",
		4683: "HuRob",
		3335: "Huawei",
		7054: "Huayuan",
		4481: "Hub-Tech",
		7733: "Hubbell",
		9300: "Huike",
		7313: "Huiwan",
		7090: "Huiyang",
		949:  "Human",
		8507: "Humannix",
		7496: "Humanorporated",
		4802: "Humanware",
		4012: "Hunkeler",
		2151: "Hunt",
		472:  "Hunter",
		8025: "Huria",
		8253: "Husqvarna",
		8118: "Hutec",
		6750: "Hutek",
		2028: "Huwell",
		4273: "Hybrid",
		4666: "Hybus",
		6329: "Hydra",
		8713: "Hydro",
		1227: "Hyglo",
		4306: "Hylab",
		2768: "Hymatom",
		4229: "Hynet",
		1657: "HyperEdge",
		5177: "Hypercom",
		3681: "Hypermedia",
		3750: "Hyperstone",
		5098: "Hypertec",
		3972: "Hypertherm",
		7690: "Hypervolt",
		5448: "Hytec",
		7601: "Hytera",
		1374: "HyunJu",
		1351: "Hyundai",
		7592: "Hyunjin.com",
		7844: "Hyunteck",
		8037: "I&C",
		6728: "I'M",
		5645: "I-BUS",
		5332: "I-Cube",
		5875: "I-Cubed",
		6685: "I-LAX",
		2415: "I-O",
		6637: "I-Storm",
		8573: "I-Sys",
		4805: "I-TEC",
		4776: "I-Tech",
		4384: "I-WIN",
		867:  "I/F-CoM",
		236:  "I2SE",
		7521: "I4VINE",
		8698: "IAI",
		2023: "IAV",
		6887: "IAdea",
		494:  "IBASE",
		7934: "IBC",
		372:  "IBM",
		4230: "IBR",
		9492: "IC",
		269:  "IC-Net",
		1184: "ICA",
		62:   "ICANN",
		5075: "ICC",
		42:   "ICE",
		600:  "ICG",
		1027: "ICHIPS",
		5063: "ICM",
		1249: "ICP",
		1998: "ICPDAS",
		8603: "ICR",
		763:  "ICS",
		575:  "ICUE",
		3495: "ICX",
		7864: "ID",
		9072: "IDFone",
		488:  "IDIS",
		1299: "IDK",
		558:  "IDOT",
		9060: "IDS",
		3226: "IDT",
		4074: "IDX",
		6657: "IDY",
		2964: "IEC",
		2629: "IEE",
		7232: "IEP",
		745:  "IER",
		9344: "IES",
		319:  "IFM",
		5249: "IFP",
		7665: "IGI",
		8982: "IGRS",
		7324: "IGShare",
		6086: "IGT",
		905:  "IGYS",
		4336: "IHSE",
		8837: "II-VI",
		8267: "IJ",
		8451: "IMAGO",
		6852: "IMAX",
		26:   "IMC",
		5280: "IMD",
		5146: "IMF",
		1387: "IMI",
		8856: "IMK",
		9285: "IMS",
		3178: "IMV",
		5481: "IN-NET",
		5819: "INAT",
		353:  "INIT",
		1579: "INITECH",
		1651: "INITIUM",
		4782: "INNERINT",
		3195: "INNOTZ",
		1017: "INNOVI",
		996:  "INNOWELL",
		3544: "INOCOVA",
		7668: "INQ",
		4393: "INRange",
		6219: "INSIDE",
		6490: "INTARSO",
		5795: "INTEGRATED",
		7619: "INVENTEC",
		2903: "INVISIO",
		5268: "IO",
		6952: "IO-Power",
		1170: "IOA",
		4346: "IOGEAR",
		390:  "IOI",
		4525: "IOLAN",
		4206: "ION",
		4476: "IONODES",
		2771: "IONOS",
		6768: "IOTTECH",
		7186: "IP-Line",
		6732: "IP-NET",
		402:  "IP.Access",
		778:  "IPAS",
		5946: "IPC",
		1583: "IPCserv",
		9463: "IPEVO",
		1463: "IPFLEX",
		1087: "IPFRONT",
		1283: "IPMobileNet",
		893:  "IPOptical",
		9015: "IPROAD",
		7083: "IPS",
		737:  "IPWireless",
		4602: "IPaXiom",
		8150: "IPitomy",
		831:  "IPmental",
		8124: "IPmotion",
		4601: "IPnect",
		679:  "IPrad",
		5562: "IQ",
		5198: "IQinVision",
		4891: "IRCA",
		4528: "IRD",
		2420: "IRIICHI",
		2619: "IRT",
		4426: "IRTrans",
		7596: "IRay",
		5744: "ISA",
		1717: "ISAC",
		6275: "ISB",
		3172: "ISCO",
		297:  "ISDN",
		4825: "ISGUS",
		3602: "ISL",
		3365: "ISONAS",
		2627: "ISR",
		8158: "ISSC",
		2667: "IT-Factory",
		6766: "IT-IS",
		7019: "ITC",
		1176: "ITDevices",
		8496: "ITE",
		3996: "ITEC",
		4518: "ITF",
		2010: "ITFOR",
		6472: "ITG",
		1992: "ITI",
		6793: "ITL",
		2298: "ITO",
		8708: "ITON",
		1026: "ITRAN",
		1890: "ITSupported",
		1178: "ITT",
		692:  "ITTC",
		8565: "ITTIM",
		3475: "ITU-T",
		4294: "ITV",
		4955: "IVA",
		8005: "IVS",
		2991: "IVT",
		2538: "IWICS",
		7532: "IXECLOUD",
		2447: "IZT",
		6080: "Ibond",
		2777: "Ibtek",
		7826: "Icarvisions",
		5669: "Icom",
		7115: "Icomera",
		760:  "Iconag",
		2157: "Icotera",
		3639: "Icron",
		5771: "Ictv",
		7518: "IdaTech",
		4800: "Idaho",
		3043: "Ideal",
		4495: "Idealbt",
		4676: "Idemia",
		3722: "Identec",
		1497: "Identicard",
		1137: "Identix",
		7091: "Iders",
		4958: "Idream",
		4733: "IfTA",
		7191: "Iffu",
		6746: "Iflytek",
		2660: "Ifotec",
		1752: "Iftest",
		7893: "Igneous",
		6813: "IgniteNet",
		2096: "Iiga",
		1414: "Ikanos",
		1272: "Ilevo",
		1122: "Ilinx",
		3416: "ImCoSys",
		7357: "ImTech",
		4351: "Imacs",
		824:  "ImageCom",
		495:  "Imagenics",
		5603: "Imagic",
		3509: "Imagination",
		5691: "Imagine",
		8890: "Imaqliq",
		554:  "Imation",
		2295: "Imatron",
		2379: "Imci",
		1482: "Imerge",
		4032: "ImesD",
		5536: "Imlogix",
		8526: "ImmediaTV",
		4980: "Impacct",
		5090: "Impact",
		3597: "Impatica",
		351:  "Imperial",
		8197: "Impex-Sat&Co",
		3063: "Impinj",
		6005: "Impresstek",
		3567: "Impro",
		5897: "Impulse",
		1711: "Imsys",
		4155: "In-Circuit",
		1813: "In-Tech",
		3166: "InGrid",
		4561: "InMage",
		3896: "InPhase",
		1202: "InPro",
		7768: "InShow",
		9003: "InView",
		8893: "InVue",
		6461: "Inala",
		4944: "Inalp",
		1035: "Inara",
		1739: "Inc.jetorporated",
		6210: "IncAA",
		2893: "IncNETWORKS",
		3618: "IncOTEC",
		4906: "IncOstartec",
		2027: "Incipient",
		6550: "Incipio",
		8046: "Incoax",
		8055: "Incognito",
		5981: "Incredible",
		9348: "IndUSNET",
		2148: "Indagon",
		6349: "Indata",
		649:  "Indel",
		9480: "Independent",
		7700: "Indieon",
		5036: "Indigita",
		4166: "Indigo",
		6830: "Individual",
		7338: "Indu-Sol",
		2583: "Inducon",
		6571: "Industrial",
		1795: "Indyme",
		1286: "IneoQuest",
		2697: "Inepro",
		7676: "Inesa",
		4267: "Inet",
		762:  "Inetcam",
		5936: "Inex",
		8972: "InfORM",
		4473: "InfORSON",
		3139: "InfRANET",
		644:  "InfiNet",
		8443: "Inficomm",
		5729: "Inficon",
		482:  "Infineon",
		1746: "Infinera",
		4833: "Infineta",
		954:  "InfiniCon",
		6572: "InfiniWing",
		2916: "Infinias",
		2133: "Infinico",
		7809: "Infinidat",
		251:  "Infinilink",
		746:  "Infiniswitch",
		752:  "Infinite",
		2067: "Infinity",
		6296: "Infinix",
		3340: "Infinova",
		1880: "InfoExpress",
		5823: "InfoGear",
		7638: "InfoSight",
		9375: "Infoblox",
		3489: "Infocrypt",
		6132: "Infocus",
		2740: "Infohand",
		5225: "Infolibria",
		3977: "Infomark",
		8976: "Infopia",
		6418: "Inform",
		4385: "Informatics",
		7:    "Information",
		8986: "Informtekhnika",
		5982: "Infortrend",
		1962: "Infotec",
		5809: "Infotek",
		92:   "Infotron",
		1964: "Infrant",
		8731: "Ingate",
		536:  "Ingenico",
		6138: "Ingenieria",
		348:  "Ingersoll-Rand",
		4044: "Ingespace",
		1968: "Ingeteam",
		4593: "Inhand",
		7967: "Inhon",
		5442: "Inid",
		8493: "Inim",
		2299: "Initio",
		224:  "Inkel",
		623:  "Inkra",
		6945: "Inlab",
		6643: "Inmarsat",
		900:  "Inncom",
		4790: "InnerSpace",
		2865: "InnerWireless",
		3846: "Innes",
		8894: "InnoDigital",
		2321: "InnoLabs",
		2367: "InnoMedia",
		6113: "InnoMediaLogic",
		5310: "InnoSys",
		7343: "Innocent",
		675:  "Innocom",
		8459: "Innogrit",
		7193: "Innolight",
		8350: "Innolux",
		812:  "Innomedia",
		9410: "Innometriks",
		8308: "Innonet",
		9134: "Innophase",
		1384: "Innopia",
		7342: "InnosiliconTechnology",
		3436: "Innotech",
		8439: "Innotube",
		6207: "Innova",
		1064: "Innovance",
		5602: "Innovaphone",
		4695: "Innovar",
		6211: "Innovat",
		428:  "Innovative",
		2373: "Innovex",
		9131: "Innovid",
		6690: "Innovolt",
		5029: "Inova",
		4212: "Inovis",
		9117: "Inovonics",
		6917: "Inpeco",
		4791: "InpegVision",
		4370: "Inphi",
		5388: "Input/Output",
		1891: "Inqnet",
		7290: "Insensi",
		4886: "Insightek",
		7261: "Insigma",
		8367: "Inspire",
		8871: "Inspiremobile",
		7723: "Inspur",
		5792: "Instem",
		4873: "Instrumentation",
		6332: "Insystec",
		7110: "Intai",
		6866: "IntalTech",
		3947: "Intego",
		8882: "IntegraOptics",
		5874: "Integrated",
		8513: "Integri-Sys.com",
		5687: "Integrix",
		421:  "Intel",
		3542: "Intelbras",
		4826: "InteliCloud",
		3414: "Intelicis",
		7594: "Intelight",
		2487: "Intellambda",
		3707: "Intellect",
		9242: "IntelliVoice",
		6379: "Intellian",
		4985: "Intellibyte",
		5487: "Intellicom",
		1861: "Intelligent",
		6544: "Intellimedia",
		3484: "Intellio",
		7239: "Intellion",
		8001: "Intellisis",
		7487: "Intellithings",
		7577: "Intellivision",
		6014: "Intelliworxx",
		7337: "Intent",
		3869: "Inter-M",
		9460: "InterCreative",
		2494: "InterEnergy",
		1398: "InterEpoch",
		4081: "Interacoustics",
		2683: "Interactek",
		3614: "Interactivetv",
		1047: "Interalia",
		3516: "Interay",
		9213: "Intercept",
		9098: "Intercom",
		8901: "Intercon",
		3331: "Intercross",
		1410: "Interface",
		5273: "Intergon",
		6318: "Intergraph",
		3024: "Interlink",
		2318: "Intermec",
		5921: "Intermedium",
		639:  "International",
		8604: "Internet",
		5346: "Internix",
		77:   "Interphase",
		9472: "Interphone",
		5012: "Intersil",
		347:  "Intersoft",
		4862: "Interspiro",
		3968: "Intertain",
		8282: "Intertech",
		514:  "Intervoice-Brite",
		5121: "Interware",
		1215: "IntiGate",
		1034: "Intime",
		3929: "Intoto",
		5350: "IntraServer",
		4349: "Intraco",
		7416: "Intrakey",
		1067: "Intransa",
		919:  "Intraserver",
		9122: "Intrigue",
		927:  "Intrinsyc",
		9231: "Intrising",
		1512: "Intronicsorporated",
		54:   "Intrusion.com",
		983:  "Intruvert",
		3942: "Intuicom",
		6807: "Intune",
		1544: "Invacom",
		6874: "Invenit",
		4252: "Invensys",
		3979: "Inventec",
		8819: "Inventek",
		1075: "Inventel",
		2649: "Invento",
		7349: "Inventum",
		6038: "Invertex",
		936:  "Invicta",
		5074: "Invisible",
		1476: "Invivo",
		6638: "Invoxia",
		4542: "Inyuan",
		8550: "IoT",
		6060: "Iomega",
		2005: "Ionix",
		1673: "Iosoft",
		6744: "Iota",
		6037: "Iowave",
		704:  "Ipanema",
		1348: "Ipetronik",
		4801: "Iphion",
		2780: "Iprobe",
		1200: "Ipsilorporated",
		3981: "Ipte",
		6498: "Iqsim",
		3012: "Iqua",
		3555: "Ircona",
		5193: "Iris",
		6909: "Irlab",
		4993: "Irlan",
		3409: "Irlink",
		5100: "Ironicsorporated",
		6754: "Irts",
		4847: "Irumtek",
		3673: "IsaacLandKorea",
		1853: "Iscutum",
		4224: "Isdn",
		4223: "Isdyne",
		4936: "Ishida",
		146:  "Isicad",
		2938: "Isilon",
		1852: "Iskraemeco",
		5883: "Isolation",
		7151: "Ison",
		4555: "Isotek",
		9062: "Isung",
		1869: "Itcare",
		930:  "Itcn",
		6282: "Itel",
		2864: "Iteris",
		9053: "Itibia",
		5641: "Itis",
		6869: "Iton",
		1109: "Itron",
		8280: "Itronic",
		2162: "Itronix",
		8076: "Ivenix",
		5595: "Ivex",
		4994: "Ivron",
		6022: "Iwill",
		1107: "J",
		7915: "J-MEX",
		3114: "J-TEK",
		1501: "J-THREE",
		805:  "J-Works",
		5944: "J1",
		4247: "J125",
		1873: "JAI",
		468:  "JAMA",
		9094: "JANUS",
		9392: "JBL",
		6902: "JDA",
		7691: "JDC",
		1143: "JEAN",
		1578: "JEPICO",
		3347: "JESS-LINK",
		7590: "JETTER",
		9394: "JJ",
		2973: "JJPlus",
		2458: "JL",
		8434: "JLG",
		6450: "JM-DATA",
		1021: "JMI",
		965:  "JMP",
		8392: "JMR",
		3690: "JMicron",
		7931: "JNC",
		3390: "JOYTOTO",
		2022: "JPS",
		6613: "JRC",
		4827: "JRD",
		8853: "JRI",
		5742: "JRL",
		2505: "JStream",
		8907: "JTD",
		9306: "JTECH",
		6896: "JTI",
		8959: "JUNGJIN",
		7062: "JW",
		8201: "JWEntertainment",
		3546: "JWTrading",
		6600: "JY",
		7863: "Jabil",
		2759: "Jablotron",
		2523: "Jacobsons",
		5763: "Jacomo",
		7056: "Jadak",
		7493: "Jaewoncnc",
		1851: "Jamex",
		4507: "Janam",
		2056: "Janitza",
		8920: "Janteq",
		5684: "Janz",
		5454: "Japan",
		6337: "Jarogate",
		5870: "Jasco",
		1222: "Jascom",
		521:  "Jasmine",
		4069: "Jastec",
		6161: "Jato",
		5199: "Jatom",
		6125: "Jaton",
		3161: "Jaty",
		5667: "Javelin",
		9001: "JayBird",
		6148: "Jaycor",
		2815: "JazzMutant",
		6411: "Jebsee",
		7257: "Jeda",
		2993: "Jennic",
		660:  "Jenoptik",
		2908: "Jeongmin",
		4739: "Jeorich",
		5986: "Jetcell",
		8880: "Jetlun",
		8112: "Jetmobile",
		711:  "Jetstream",
		5283: "Jetter",
		8820: "Jfcontrol",
		8796: "JiQiDao",
		7977: "JianLing",
		7369: "Jibo",
		7139: "Jide",
		8143: "Jigowatts",
		6946: "JinQianMao",
		7484: "Jingsheng",
		7943: "Jinmuyu",
		8470: "Jiwumedia",
		3626: "Joby",
		3783: "JoeScan",
		7043: "Jolata",
		7334: "Jolla",
		3463: "Jorjin",
		1020: "Jotron",
		8250: "Joview",
		8529: "Jovision",
		3598: "Joybien",
		8204: "Joyent",
		1650: "Joymax",
		1997: "Joyteck",
		7096: "Jubix",
		7885: "Juice",
		7188: "Juin",
		4202: "Juki",
		5124: "Juko",
		8354: "Julong",
		4425: "JumpGen",
		9186: "Jumptronic",
		6376: "Junilab",
		826:  "Juniper",
		5453: "Jupiter",
		772:  "Jupiters",
		7831: "Jurumani",
		1719: "Jusan",
		863:  "JustEzy",
		6473: "Justec",
		9260: "Justone",
		5523: "Justsystem",
		8714: "Jwcnetworks",
		2907: "K",
		7992: "K's",
		2356: "K-BOT",
		5769: "K-NET",
		1935: "K-Patents",
		8962: "K2NET",
		2860: "K40",
		2467: "KAYABA",
		1001: "KB",
		7049: "KBC",
		3134: "KBS",
		3092: "KBT",
		999:  "KC",
		7856: "KCE",
		8245: "KCF",
		2535: "KCodes",
		188:  "KDC",
		3162: "KDE",
		769:  "KDT",
		5390: "KEBA",
		9343: "KEEBOX",
		6397: "KEMAS",
		8021: "KERNEL-I",
		7444: "KETEK",
		4013: "KIMIN",
		1433: "KINGENE",
		8427: "KLINFO",
		7427: "KMB",
		3984: "KMW",
		5905: "KNX",
		583:  "KOANKEISO",
		3112: "KORWIN",
		6770: "KSH",
		6365: "KSP",
		7291: "KST",
		5854: "KT",
		3027: "KT&C",
		4543: "KTC",
		3262: "KTF",
		151:  "KTI",
		7198: "KTIS",
		8855: "KUNBUS",
		4158: "KURUSUGAWA",
		5743: "KVB/Analect",
		8904: "KVH",
		6229: "KWB",
		5949: "KYE",
		4042: "KYLAND",
		3034: "KYLINK",
		186:  "KYOWA",
		1860: "Ka-Ro",
		4665: "Kaga",
		7871: "Kai-EE",
		6583: "Kaiam",
		2724: "Kaicom",
		2592: "Kaimei",
		3834: "Kaise",
		9426: "Kaishun",
		6420: "Kakao",
		1662: "Kaleidescape",
		4771: "Kalki",
		5467: "Kalpana",
		8788: "Kamama",
		3391: "Kameleon",
		4707: "Kaminario",
		6437: "Kamo",
		2788: "Kamstrup",
		1418: "Kanematsu",
		6533: "KangSheng",
		1269: "Kaonmedia",
		6048: "Kapadia",
		1682: "Kaparel",
		7079: "Kapelse",
		6149: "Kapsch",
		4389: "Kapsys",
		1259: "Karam",
		4234: "Kardios",
		1284: "Karel",
		2135: "Kasda",
		7928: "Kastle",
		4186: "Katana",
		6010: "Kathrein-Werke",
		9155: "Katoudenkikougyousyo",
		5181: "Katron",
		6058: "Katsujima",
		5415: "Kayser-Threde",
		4640: "Kaztek",
		4531: "Kedah",
		7375: "Keenetic",
		3129: "Kei",
		5800: "Keisokugiken",
		3241: "Kelman",
		9387: "Kelvin",
		9269: "Kemppi",
		7630: "Kenade",
		243:  "Kenetec",
		7643: "Kenstel",
		1602: "Kentec",
		2149: "Kentima",
		4219: "Kentrox",
		3737: "Kenwin",
		1920: "Kenwood",
		9067: "Keopsys",
		1688: "Kerajet",
		4942: "Kerbango",
		2824: "Keri",
		7778: "Kerlink",
		5245: "Kestrel",
		8437: "Ketra",
		3617: "Keumbee",
		5395: "KeunYoung",
		1189: "Key",
		2750: "KeyEye",
		2261: "KeyMed",
		6034: "Keycorp",
		317:  "Keyence",
		3600: "Keysight",
		3965: "Keytronix",
		9396: "Khomp",
		8792: "Khwahish",
		1710: "Kihoku",
		6962: "Kiio",
		2265: "Kikusui",
		3294: "Kimaldi",
		6658: "Kimball",
		5164: "Kimpsion",
		5547: "Kimtron",
		8912: "Kinestral",
		6355: "Kinetics",
		8051: "Kinexon",
		9047: "KingTing",
		2735: "Kinghold",
		3322: "Kingjim",
		5703: "Kingmax",
		7624: "Kingnetic",
		7816: "Kingsignal",
		5143: "Kingstar",
		3348: "Kingstate",
		4882: "Kingston",
		3039: "Kingtronics",
		1672: "Kingwave",
		7629: "Kinion",
		8649: "Kinova",
		299:  "Kinpo",
		8160: "Kioxia",
		1153: "Kirana",
		8213: "Kirisun",
		4727: "Kiryung",
		4515: "Kisan",
		3560: "KitWorks.fi",
		7423: "Kivic",
		9229: "Kivo",
		3216: "Kiyon",
		2792: "Klas",
		840:  "Kleinknecht",
		5078: "Klever",
		8284: "Kmdata",
		670:  "Knilink",
		2095: "Knovative",
		1635: "Knurr",
		7919: "KoamTac",
		2949: "Kocom",
		1955: "Koden",
		1671: "Kodeos",
		5464: "Kodiak",
		1820: "Kodicom",
		1204: "Koei",
		5217: "Koga",
		2857: "Kohler",
		2403: "Kollmorgen",
		334:  "Kollmorgen-Servotronix",
		5903: "Komatsu",
		2009: "Komax",
		183:  "Komodo",
		964:  "Konami",
		7436: "Koncar",
		2818: "Koncept",
		6567: "KongTop",
		3533: "Konka",
		8248: "KonnectONE",
		7634: "Konten",
		63:   "Kontron",
		4398: "Koratek",
		853:  "Korea",
		2606: "Korenix",
		2549: "Korg",
		6471: "Korins",
		8756: "Koss",
		846:  "Kott",
		8115: "Koubachi",
		5232: "Kouwell",
		7283: "Kove",
		802:  "Kowa",
		6032: "Koyo",
		4038: "Kozio",
		2955: "Kprotech",
		6671: "Kraftway",
		3892: "Kramer",
		407:  "Kreatel",
		4793: "Kroger",
		3886: "Krohne",
		3953: "Ksic",
		4820: "Ktnf",
		5909: "Kubota",
		5727: "Kubotek",
		9041: "Kuhn",
		6712: "Kuipers",
		720:  "Kumahira",
		8840: "Kumalift",
		6596: "Kummler+Matter",
		1147: "Kumoh",
		2734: "Kumyoung",
		7547: "KunTeng",
		992:  "Kuokoa",
		6644: "Kurth",
		6076: "Kuzumi",
		6001: "Kvaser",
		6526: "Kwangsung",
		7766: "Kwangwon",
		3995: "KwikByte",
		3736: "Kworld",
		7908: "Kyland-USA",
		2849: "Kyocera",
		3789: "Kyohritsu",
		7812: "Kyolis",
		1898: "Kyoto",
		7715: "Kyowa",
		2242: "Kyushu",
		2184: "Kyushu-kyohan",
		8016: "Kyynel",
		1419: "L&F",
		5893: "L&N",
		3158: "L-3",
		6639: "L-Tech",
		7820: "LAM",
		6312: "LAN-TEC",
		6199: "LANBit",
		5738: "LANCOM",
		1180: "LANergy",
		7620: "LARK",
		4919: "LARsys-Automation",
		3413: "LBNL",
		4920: "LCFC",
		1850: "LEA",
		776:  "LEA*D",
		7591: "LENUS",
		2913: "LET'S",
		1829: "LETEK",
		247:  "LEUNIG",
		869:  "LG",
		8601: "LG-Ericsson",
		7469: "LGE",
		6189: "LINK2IT",
		4840: "LIOS",
		452:  "LITE-ON",
		4310: "LInTech",
		6980: "LM",
		8071: "LMI",
		4512: "LNC",
		6487: "LNT-Automation",
		8887: "LOCOSYS",
		5313: "LOGWARE",
		4377: "LOHUIS",
		3685: "LORD",
		1563: "LOYTEC",
		1626: "LS",
		1116: "LSI",
		6162: "LTX-Credence",
		8913: "LVS",
		237:  "LXCO",
		5185: "LXE",
		9474: "LXinstruments",
		7770: "LaVision",
		8173: "LabJack",
		7422: "Labris",
		78:   "Labtam",
		1089: "Lafon",
		4007: "Lagotek",
		4537: "Laird",
		5637: "Lake",
		3221: "Laketune",
		7711: "Lampex",
		9279: "Lamprey",
		364:  "Lampus",
		4912: "LanPro",
		2504: "LanReady",
		5859: "Lanart",
		6052: "Lanbird",
		8578: "Lanbowan",
		5118: "Lanco",
		5391: "Land",
		2531: "Land-Cellular",
		7734: "Landauer",
		4249: "Landings",
		2228: "Landis+Gyr",
		5544: "Lanex",
		5583: "Lanner",
		5851: "Lanoptics",
		5868: "Lans",
		8370: "Lansen",
		7084: "Lansentechnology",
		6795: "Lantis",
		5529: "Lantronix",
		1416: "Lanvoice",
		5161: "Lanwan",
		6602: "Laon",
		6467: "LaonPeople",
		4245: "Lapis",
		5020: "Lara",
		8191: "Lars",
		4270: "Larscom",
		2307: "Lasat",
		4122: "Lascar",
		4516: "Lasercraft",
		5427: "Lasergraphics",
		5865: "Lasermaster",
		1479: "Lassen",
		4456: "Lastar",
		5844: "Latch",
		8960: "Latecoere",
		7975: "Latticework",
		5912: "Lauterbach",
		647:  "Lava",
		3138: "Law-Chain",
		1670: "Lawo",
		7116: "Layer3TV",
		6809: "Layon",
		2573: "LeWiz",
		9045: "Leadcore",
		804:  "Leader",
		1411: "Leadfly",
		4927: "Leadtek",
		3434: "Leaf",
		5860: "Leap",
		2590: "LeapComm",
		7982: "LeapFrog",
		9478: "Leapfive",
		3536: "Lear",
		403:  "Lectron",
		4632: "Lectrosonics",
		4127: "Ledco",
		9328: "LeddarTech",
		9330: "Ledvance",
		6730: "Lee-Dickens",
		9431: "Leeman",
		7370: "Leeo",
		6838: "Leetek",
		2142: "Legra",
		682:  "Legrand",
		5672: "Leichu",
		9197: "Leifheit",
		6223: "Leightronix",
		8536: "Leimac",
		2936: "Leipold+Co.GmbH",
		162:  "Leiser",
		9031: "Leitner",
		2666: "Lely",
		5115: "Lemcom",
		8183: "Lenbrook",
		2939: "Leneco",
		4560: "Lengda",
		3352: "Lenntek",
		4475: "Lenord",
		2665: "Lenovo",
		1383: "Lenten",
		1540: "Lenze",
		8484: "Leoni",
		8327: "Leonton",
		8351: "Lesira",
		3192: "LeucotronEquipamentos",
		2985: "Leuze",
		7777: "LevelOne",
		7231: "Levven",
		2929: "LexBox",
		7991: "Lexar",
		8522: "Lexking",
		613:  "Lexmark",
		3954: "LiComm",
		3310: "Lianhe",
		5530: "Liberty",
		8813: "Libratone",
		3052: "LibreStream",
		4819: "Licera",
		1233: "Liebert-Hiross",
		3818: "Liecthi",
		7127: "Life",
		6661: "LifeBEAM",
		9494: "LifeHealth",
		2798: "LifeSize",
		2899: "LifeSync",
		2199: "Lifetron",
		7010: "Light",
		220:  "LightChip",
		881:  "LightSand",
		4722: "Lightcomm",
		6082: "Lightera",
		5004: "Lightner",
		744:  "Lightpointe",
		8559: "Lightspeed",
		4947: "Lightwave",
		7804: "LightwaveRF",
		3423: "LigoWave",
		9170: "Lilee",
		3836: "Limetek",
		3236: "LinTech",
		6928: "Linctronix",
		6702: "Linea",
		9329: "Linear",
		8113: "Lineeye",
		4439: "LinkSprite",
		8951: "Linkcom",
		4533: "Linkflex",
		7611: "Linkflow",
		7356: "Linkintec",
		1783: "Linksys",
		7683: "Linktel",
		8953: "Linktop",
		6024: "Linkup",
		3813: "Linkwise",
		4817: "Linn",
		4455: "LinoWave",
		83:   "Linotype-Hell",
		4260: "Linq",
		9160: "Lintes",
		1004: "Linxtek",
		1278: "Liontech",
		7593: "Liquidtool",
		8462: "Lisantech",
		1669: "Litchfield",
		5789: "Lite-ON",
		6403: "Lite-On",
		7710: "LiteON",
		2675: "LiteTouch",
		8642: "Litemax",
		2874: "Liteon",
		1300: "Littlefeet",
		5495: "Litton",
		5340: "Litton/Poly-Scientific",
		4378: "LiveTV",
		7877: "LiveU",
		6520: "Liverock",
		9092: "Livescribe",
		7753: "Livesecu",
		5853: "Livingston",
		3091: "Liyuh",
		3228: "LoBenn",
		4738: "LoJack",
		1668: "Load",
		2281: "Lobos",
		5830: "LocSoft",
		828:  "Locusorporated",
		4110: "Lodam",
		1219: "LodgeNet",
		9149: "Loenk",
		4588: "LogMeIn",
		6563: "Logi-D",
		4201: "LogiCan",
		1168: "LogiSync",
		6007: "Logibag",
		5079: "Logic",
		9458: "Logic3",
		2021: "LogicaCMG",
		253:  "Logical",
		9093: "Logicom",
		51:   "Logicraft",
		8590: "Logipix",
		4423: "Logiplus",
		6675: "Logistic",
		7142: "Logital",
		242:  "Logitec",
		4080: "Logitech",
		4719: "Logitek",
		4846: "Logiways",
		6925: "Logosol",
		597:  "Logostek",
		6536: "LongSung",
		4579: "Longcheer",
		6419: "Longconn",
		7221: "Longicorn",
		3878: "Longkay",
		4548: "Lookit",
		3620: "Loopcomm",
		6480: "Looxcie",
		5431: "Loran",
		9116: "Lorent",
		4099: "Lorex",
		3701: "Lorica",
		9467: "Lorom",
		2288: "Loud",
		7762: "Low",
		2103: "Lowrance",
		7331: "Loxone",
		6130: "Lsics",
		4115: "Lucas",
		827:  "Lucent",
		5808: "Lucidata",
		3815: "Lucky",
		4634: "Ludl",
		7455: "Lufkin",
		4433: "Lumasense",
		1729: "Lumenera",
		8868: "Lumenpulse",
		8932: "Lumewave",
		3971: "Lumexis",
		8619: "Lumigon",
		8065: "Luminar",
		6808: "Luminator",
		3153: "Lundinova",
		2280: "Lutron",
		5023: "LuxN",
		5524: "Luxcom",
		4752: "Luxtera",
		8408: "Luxul",
		8028: "Lvswitches",
		1244: "Lyan",
		1681: "Lycium",
		6152: "Lynk",
		8307: "Lynxspringl",
		8137: "Lytro",
		6859: "Lytx",
		4580: "Lyyn",
		8014: "LzLabs",
		1794: "M&S",
		1934: "M-Audio",
		9417: "M-Cube",
		2369: "M-System",
		435:  "M.C.C.I",
		5773: "M.I",
		7769: "M2Communication",
		2940: "M2I",
		4870: "M2Mnet",
		5818: "M2Motive",
		1793: "MACKIE",
		6323: "MACNICA",
		2990: "MAKUS",
		3264: "MARA",
		230:  "MARGI",
		1513: "MARKEM",
		1031: "MAT",
		553:  "MAVIX",
		7411: "MAX-Tech",
		3451: "MAZeT",
		948:  "MBM",
		3845: "MBS",
		2279: "MBTech",
		8970: "MC",
		6941: "MCD",
		1420: "MCM",
		9248: "MCNEX",
		8420: "MCOT",
		6683: "MCT",
		2663: "MDK",
		2378: "MEDIA4",
		8826: "MEDIAEDGE",
		7921: "MEG",
		7031: "MEIZU",
		4008: "MEL",
		8375: "MELPER",
		8164: "MESADA",
		7947: "MHL",
		4839: "MIA",
		4997: "MICRILOR",
		8563: "MICRODIA",
		7830: "MICRODIGTAL",
		5384: "MICROSENS",
		6134: "MICROWI",
		1946: "MIDAS",
		5677: "MIHARU",
		6170: "MII",
		9363: "MIMO",
		8364: "MINIX",
		4918: "MIPRO",
		1687: "MITEQ",
		4905: "MKD",
		601:  "MKNet",
		4639: "MMB",
		4933: "MMC",
		9356: "MMPC",
		9255: "MOCACARE",
		8627: "MODI",
		2509: "MOIMSTONE",
		7876: "MOKO",
		7791: "MONAD",
		1316: "MONDIAL",
		7747: "MPB",
		8515: "MPI",
		5401: "MPL",
		9198: "MR&D",
		5200: "MRG",
		9123: "MRS",
		2254: "MRV",
		7958: "MS-Magnet",
		3055: "MSI",
		6570: "MST",
		4865: "MTA",
		7813: "MTG",
		2773: "MTI",
		8068: "MTMCommunications",
		6747: "MTN",
		577:  "MTS",
		2522: "MTT",
		104:  "MTX",
		8964: "MVTECH",
		917:  "MWE",
		1577: "MWS",
		8870: "MXCHIP",
		7170: "MXT",
		6781: "MYK",
		2102: "MYNAH",
		1338: "MYTECS",
		2077: "Maas",
		4952: "Maatel",
		3776: "Macab",
		9101: "Macandc",
		2138: "Macey",
		3327: "Mackware",
		5154: "Macq",
		5398: "Macraigor",
		5840: "MacroSAN",
		1012: "Macrolink",
		5022: "Macromate",
		6931: "Macrotech",
		564:  "Macrotek",
		5327: "Macrovision",
		5970: "Mactell",
		70:   "Madge",
		8547: "MadgeTech",
		7264: "Maestronic",
		4661: "Mag",
		252:  "Mag-Tek",
		3949: "Magellan",
		9406: "Magenta",
		4948: "MagicRam",
		3064: "Magicard",
		7701: "Magicjack",
		153:  "Magna",
		4000: "Magna-Power",
		6794: "MagnaCom",
		8770: "MagneMotion",
		725:  "Magnipix",
		701:  "Mahi",
		6918: "Maike",
		523:  "Mainnet",
		4585: "Mainpine",
		8922: "Maipu",
		7073: "MakerBot",
		6900: "Mako",
		4763: "Maksat",
		1949: "Makus",
		945:  "Malachite",
		6465: "Malgn",
		713:  "Malibu",
		1426: "Mamiya-OP",
		7240: "ManTechnology",
		2147: "Mangrove",
		1530: "Manticom",
		350:  "Mantra",
		6280: "Manycolors",
		8892: "Manzanita",
		1493: "Maple",
		657:  "Mapletree",
		4572: "Mapower",
		9267: "Mapper.ai",
		6968: "Maquet",
		1699: "Maranti",
		5927: "Marben",
		2388: "March",
		31:   "Marconi",
		5218: "Mariner",
		5649: "Markem-Imaje",
		7247: "Marketaxess",
		6448: "Marketech",
		9246: "Markov",
		5170: "Marner",
		8776: "MarqMetrix",
		5740: "Marquip",
		6697: "Marshal",
		4575: "Marshall",
		2387: "Martinho-Davis",
		3944: "Marusys",
		6363: "Masimo",
		1166: "Massana",
		6343: "Masscomp",
		4320: "Masterclock",
		4608: "Masternaut",
		1423: "MathStar",
		3071: "Mathtech",
		8402: "Matis",
		2474: "Matisse",
		6344: "Matra",
		1885: "Matrics",
		1:    "Matrix",
		4293: "Matrox",
		9367: "Matsufu",
		148:  "Matsushita",
		4465: "Mattel",
		712:  "Mavenir",
		3879: "Maverick",
		1448: "Mavin",
		6560: "Max",
		9326: "MaxID",
		8789: "MaxMedia",
		2749: "MaxStream",
		9030: "MaxTronic",
		4913: "MaxVision",
		914:  "Maxan",
		2432: "Maxanna",
		2222: "Maxcess",
		8695: "Maxeler",
		3497: "Maxfor",
		4567: "Maxian",
		7158: "Maxio",
		5299: "Maxlinear",
		1555: "Maxlink",
		5535: "Maxpeed",
		7294: "Maxphotonics",
		5123: "Maxton",
		2386: "Maxtor",
		6523: "Maxway",
		1036: "Maxxan",
		4779: "Maya-Creation",
		5215: "Mayan",
		1321: "Mayekawa",
		8523: "Maytronics",
		1803: "Mbari",
		2083: "McAfee",
		6470: "McCain",
		9466: "McPay",
		2121: "Mcharge",
		8825: "Mciao",
		5505: "Mcnc",
		6115: "Mcns",
		9450: "Mcot",
		6075: "Mcquay",
		7360: "Measy",
		1083: "Mecalc",
		8835: "Mechatro",
		3648: "Med-Eng",
		9080: "MedHab",
		555:  "Medea",
		3853: "Media",
		2452: "MediaCell",
		1781: "MediaChorus",
		1760: "MediaQ",
		4828: "MediaSputnik",
		1876: "MediaTek",
		6553: "Mediabridge",
		5220: "Mediafire",
		2513: "Medialink-i",
		688:  "Medialogic",
		5201: "Mediastar",
		1466: "Mediatek",
		8765: "Medicaroid",
		3615: "Medicis",
		1432: "Medicore",
		1849: "Medion",
		6937: "Mediplan",
		855:  "Medison",
		661:  "Medrad",
		1897: "Mega-Trend",
		5387: "MegaChips",
		6506: "Megabyte",
		132:  "Megadata",
		6711: "Megafone",
		88:   "Megahertz",
		2167: "Megapower",
		2141: "Megasolution",
		6311: "Megatek",
		1792: "Megatel",
		4188: "Megatron",
		8330: "Megatronix",
		4872: "Megger",
		5130: "Meidensha",
		157:  "Meiko",
		4356: "Mekics",
		7089: "Melange",
		904:  "Melco",
		8998: "Meld",
		9217: "Melec",
		6116: "Melita",
		432:  "Mellanox",
		8048: "Memjet",
		748:  "Memobox",
		1833: "MemoryLink",
		4243: "Memotec",
		3635: "Mendocino",
		5224: "Menicx",
		1821: "Mentor",
		7546: "Merchandising",
		9402: "Mercku",
		4254: "Mercury",
		6228: "Mercusys",
		6927: "Merging",
		2381: "Meridian",
		786:  "Merilus",
		6451: "Meritec",
		7024: "Meritech",
		408:  "Merix",
		480:  "Merlin",
		5247: "Merlot",
		6147: "Merrimac",
		3538: "Merten&CoKG",
		5320: "Mesa",
		923:  "Mesco",
		3227: "Meshcom",
		7949: "Mesmo",
		6612: "Messoa",
		8622: "Meta-Networks",
		2689: "MetaSwitch",
		3444: "Metacom",
		5959: "Metacomp",
		1390: "Metalink",
		1709: "Metalligence",
		7927: "Metamako",
		2089: "Metanoia",
		6313: "Metaphor",
		3338: "Metasystem",
		875:  "Metavector",
		2354: "Metawave",
		4662: "Meteocontrol",
		2943: "Meteor",
		8040: "Meter",
		538:  "Metera",
		3164: "Methode",
		6989: "Metis",
		8063: "Metrascale",
		4262: "Metricom",
		1839: "Metro",
		325:  "Metro-Optix",
		5910: "Metrodata",
		2560: "Metrohm",
		5103: "Metronix",
		7796: "Metrum",
		7129: "Mettle",
		2328: "Mettler-Toledo",
		4859: "Metz-Werke",
		7038: "Mevo",
		1549: "Mewtel",
		6782: "Mexus",
		4539: "MiXTelematics",
		9486: "Miartech",
		1649: "Micetek",
		7835: "Micobiomed",
		5157: "Micom",
		3961: "Micran",
		43:   "Micro",
		117:  "Micro-Matic",
		1759: "Micro-Optronic-Messtechnik",
		1812: "Micro-Star",
		5374: "Micro/Sys",
		5178: "MicroBrain",
		7227: "MicroPower",
		7543: "MicroSys",
		1764: "MicroWeb",
		5183: "Microboards",
		706:  "Microchip",
		97:   "Microcom",
		1911: "Microelectronics",
		6041: "Microfirst",
		6317: "Microfive",
		137:  "Micrognosis",
		2227: "Microhard",
		254:  "Microlink",
		8276: "Micromedia",
		4364: "Micromint",
		4628: "Micron",
		2457: "Micronas",
		2456: "Micronet",
		5546: "Micronics",
		326:  "Micronpc.com",
		5503: "Microplex",
		5707: "Micropolis",
		149:  "Microprocess",
		5737: "Micropross",
		4874: "Microrobot",
		156:  "Microsage",
		1642: "Microscan",
		5758: "Microsemi",
		6502: "Microseven",
		6173: "Microslate",
		612:  "Microsoft",
		4242: "Microtech",
		1408: "Microtechno",
		5518: "Microtek",
		5528: "Microtest",
		6837: "Microtime",
		2053: "Microtrol",
		4600: "Microtronic",
		994:  "Microtune",
		5144: "Microunity",
		2205: "Midas",
		6547: "Midicom",
		2493: "Midmark",
		8544: "Midokura",
		8828: "Midori",
		6177: "Midsco",
		765:  "Midstream",
		3677: "Midtronics",
		9473: "MikroBits",
		3441: "MikroM",
		6157: "Mikrodidakt",
		4235: "Mikron",
		9153: "MilDef",
		5159: "Milan",
		6696: "MileSight",
		8891: "Miljovakt",
		4332: "Millinet",
		3198: "Millipore",
		6870: "Milper",
		5976: "Mimaki",
		7785: "Mimodisplaykorea",
		6734: "Mimosa",
		2156: "Mindray",
		5034: "Mindready",
		1928: "Minds",
		1239: "Minds@Work",
		5285: "Minerva",
		636:  "Minet",
		8228: "Mini-Cam",
		8946: "Mini-Circuits",
		7408: "Minibar",
		8824: "Minieum",
		5088: "Minim",
		4731: "Minimax",
		73:   "Miniware",
		9183: "Minrray",
		8631: "Minsung",
		1838: "Mintera",
		2088: "Mintron",
		8963: "Mios",
		8064: "Miovision",
		1226: "Mipsys",
		6785: "MirWifi",
		3680: "MiraLink",
		498:  "Mirae",
		9276: "MiraeRecognition",
		7697: "MiraeSignal",
		520:  "Miraesys",
		8697: "Mirka",
		4218: "Miro",
		6522: "Miromico",
		7485: "Mist",
		1933: "Mita-Teknik",
		513:  "Mitac",
		1289: "Mitadenshi",
		5885: "Mitec",
		1217: "Mitel",
		6429: "MitraStar",
		3885: "Mitron",
		7719: "Mitsuba",
		2394: "Mitsubishi",
		6674: "Mitsunami",
		8782: "Mitsuya",
		3286: "Mitutoyo",
		3718: "Miura",
		4708: "Miyoshi",
		5567: "Mizar",
		4403: "MoCA",
		1741: "MoTEX",
		1055: "Mobiis",
		6984: "MobilMAX",
		4688: "Mobilarm",
		6381: "Mobile",
		4851: "MobileAccess",
		2609: "MobileAria",
		3864: "MobileCompia",
		7111: "Mobileeco",
		3849: "Mobilesoft",
		7954: "Mobilicom",
		985:  "Mobillian",
		3558: "Mobinnova",
		3715: "Mobisolution",
		9340: "Mobitec",
		3501: "Mobitek",
		5129: "Mobius",
		668:  "Mobiwave",
		574:  "Mobotix",
		4976: "MobyTEL",
		7099: "Moda-InnoChips",
		3914: "Modacom",
		7741: "Modcam",
		8026: "Modelleisenbahn",
		3674: "Modnnet",
		9119: "Modoosis",
		6935: "Moduel",
		6649: "Moduletek",
		3627: "ModusLink",
		8954: "Moen",
		9174: "MofiNetwork",
		7036: "Mohlenhoff",
		8491: "Moimstone",
		4093: "Mojix",
		2486: "Mojo",
		6505: "Molekule",
		579:  "Momentum",
		6558: "Monaco",
		4737: "Moneytech",
		2499: "Monitoring",
		6089: "Monterey",
		3342: "Montgomery",
		2614: "Montilio",
		1439: "Moog",
		3882: "Moohadigital",
		851:  "Moore",
		7745: "Mooredoll",
		9193: "Mophie",
		2845: "Moram",
		4408: "Morega",
		3378: "Morey",
		6477: "Morion",
		4749: "Morningstar",
		5833: "Morrow",
		886:  "Mosaic",
		3115: "Moser-Baer",
		1104: "Motah",
		3275: "Motech",
		687:  "Motorola",
		7825: "Movek",
		1435: "Movistec",
		399:  "Movita",
		2723: "Movon",
		5683: "Moxa",
		4745: "Mpedia",
		8133: "Mpgio",
		7523: "Mplus",
		6580: "Mpmkvvcl",
		8630: "MtM",
		4646: "MuLogic",
		8881: "Muehlbauer",
		6431: "Mueller",
		4062: "Mueller-Elektronik",
		8556: "Muller",
		2862: "MultiCom",
		5019: "Multidata",
		2866: "Multilink",
		4436: "Multimedia",
		1715: "Multiplex",
		4256: "Multipoint",
		1183: "Multitech",
		6019: "Multitel",
		1208: "Multitone",
		2057: "Murata",
		2234: "Murrelektronik",
		1121: "Musashi",
		3883: "Muscle",
		3831: "MusicianLink",
		7889: "Musilab",
		5704: "Mutoh",
		8494: "Muuselabs/SA",
		582:  "Muxcom",
		3616: "Mvox",
		2020: "MyA",
		7011: "MyCloud",
		9449: "MyLight",
		5888: "Myco",
		2470: "Mykotronx",
		5414: "Myricom",
		5931: "Myson",
		861:  "Myspace",
		5601: "Mysticom",
		4051: "Mytek",
		3699: "MyungMin",
		9476: "N-Radio",
		7147: "N-iTUS",
		5095: "N.A.T",
		2946: "N3",
		3873: "NAC-Intercom",
		947:  "NADEX",
		6083: "NAKAYO",
		5271: "NALTEC",
		7295: "NANOWAVE",
		6655: "NARI",
		7862: "NASCENT",
		8817: "NAVIS",
		8615: "NB",
		6309: "NBI",
		756:  "NBS",
		6185: "NBX",
		5573: "NC&C",
		6303: "NCR",
		8353: "NCTech",
		2277: "NComputing",
		9446: "NDC",
		1919: "NDR",
		2008: "NDS",
		48:   "NEC",
		8915: "NECMagnusCommunications",
		5629: "NEO",
		7963: "NEON",
		1986: "NEOSMART",
		4670: "NES",
		5162: "NET-Source",
		8059: "NET2GRID",
		5317: "NET2NET",
		3952: "NETCLEUS",
		3431: "NETCONS",
		9209: "NETINT",
		7744: "NEXCONTROL",
		4769: "NEXTEK",
		1097: "NEXTEYE",
		2902: "NF",
		9310: "NGD",
		5803: "NHC",
		8734: "NHN",
		444:  "NICE",
		6755: "NIIC",
		4145: "NIW",
		1834: "NKE",
		2814: "NL",
		8320: "NMR",
		4190: "NMS",
		4868: "NOOLIX",
		1229: "NORTHDATA",
		5639: "NOT",
		8256: "NOVA",
		3345: "NPCore",
		8593: "NRG",
		5292: "NS",
		9222: "NSD",
		7097: "NSE",
		3903: "NSGate",
		2407: "NSI",
		817:  "NSM",
		4178: "NSSLGlobal",
		2260: "NST",
		4379: "NTC-Metrotek",
		7938: "NTI",
		6990: "NTT",
		2705: "NTTPC",
		7135: "NTmore",
		4618: "NUETEQ",
		4064: "NUM",
		2758: "NUMA",
		7969: "NUUO",
		7246: "NVIDIA",
		7354: "NX",
		526:  "NXTV",
		9441: "NZXT",
		9074: "Nabtesco",
		8738: "Nabto",
		652:  "Nact",
		2950: "Nadam",
		8323: "Nadasnv",
		381:  "Nadatel",
		9266: "Nagravision",
		8149: "Nagtech",
		9428: "Nain",
		1868: "Nallatech",
		8425: "NamJunSa",
		8125: "NambooSolution",
		5599: "Namco",
		7998: "Nanoleaf",
		4220: "Nanomatic",
		8170: "Nanomegas",
		2459: "Nanometrics",
		7746: "Nanoptix",
		7209: "Nanotec",
		4559: "Nanoteq",
		6589: "Nanotron",
		9055: "Nantworks",
		3860: "Napera",
		906:  "Narad",
		2711: "Narayon",
		2002: "Nasaco",
		5709: "Nascent",
		5807: "Nashoba",
		7780: "Nastec",
		9139: "Nata-Info",
		9073: "National",
		6916: "Nations",
		7888: "Nationz",
		324:  "Native",
		4710: "NaturalPoint",
		7694: "Nautronix",
		3350: "Navcast",
		1161: "Navcom",
		7418: "Navdy",
		8200: "Naver",
		5204: "Navic",
		3672: "Navigon",
		7148: "Naviit",
		676:  "Navini",
		3935: "Navionics",
		6973: "Navtech",
		5776: "Navtel",
		695:  "Nayna",
		2699: "Naztec",
		5690: "Nbase",
		6761: "Ncse",
		5160: "Ncube",
		4407: "Neat",
		8812: "Neatframe",
		6632: "Nebra",
		8873: "Nebula",
		9378: "Nebulon",
		8072: "Nebusens",
		341:  "Necsom",
		6599: "Nectarsoft",
		4269: "Nectec",
		8009: "Neets",
		4758: "Neli",
		4755: "Nemo-Q",
		6548: "Nemoa",
		1598: "NeoAxiom",
		2543: "NeoMedia",
		278:  "NeoWave",
		3191: "Neology",
		9283: "Neomontana",
		5044: "Neon",
		3323: "Neonode",
		6163: "Neoparadigm",
		9421: "Neopis",
		3631: "Neopost",
		4217: "Neoproducts",
		1445: "Neoscale",
		7078: "Neosfar",
		4631: "Neostar",
		7994: "Neosystem",
		6678: "Neotech",
		4558: "Neotion",
		7900: "Neousys",
		4824: "Neovia",
		7027: "Nephos",
		5399: "Nera",
		7789: "Nero",
		8541: "Ness",
		4900: "Nesslab",
		6636: "Nest",
		1480: "Nestar",
		8655: "Nesys",
		318:  "Net",
		1904: "Net2Edge",
		5858: "NetAlly",
		5559: "NetApp",
		5674: "NetBoost",
		442:  "NetBotz",
		604:  "NetBurner",
		3556: "NetCare",
		549:  "NetChip",
		2589: "NetEffect",
		707:  "NetEnabled",
		1113: "NetEngines",
		2526: "NetEnrich",
		1271: "NetExcell",
		5433: "NetICs",
		568:  "NetKit",
		2082: "NetKlass",
		240:  "NetLinks",
		7160: "NetMan",
		529:  "NetMedia",
		2446: "NetModule",
		728:  "NetMount",
		973:  "NetNearU",
		2476: "NetStreams",
		4624: "NetUP",
		8217: "NetView",
		5839: "NetWorth",
		1141: "NetZerver",
		2014: "Netac",
		5259: "Netaccess",
		1431: "Netas",
		7800: "Netatmo",
		1248: "Netbind",
		5216: "Netcam",
		2034: "Netcodec",
		830:  "Netcom",
		5081: "Netcomm",
		1415: "Netcontrol",
		5180: "Netcor",
		2363: "Netcore",
		5701: "Netcorp",
		7171: "Neterix",
		5864: "Netexpress",
		960:  "Netezza",
		2564: "Netfabric",
		3948: "Netflix",
		2483: "Netforyou",
		1368: "Netgear",
		640:  "Netgem",
		9056: "Netgenetech",
		572:  "Netility",
		3735: "Netio",
		5008: "Netiverse",
		2368: "Netline",
		8452: "Netlist",
		6635: "Netlogic",
		8073: "Netmoon",
		9239: "Netonix",
		638:  "Netous",
		1117: "Netpower",
		4187: "Netquest",
		5857: "Netrix",
		5337: "Netro",
		7997: "Netronics",
		9:    "Netronix",
		2963: "Netronome",
		785:  "Nets",
		2316: "Netschools",
		5516: "Netscout",
		448:  "Netsec",
		267:  "Netsensity",
		5856: "Netspan",
		170:  "Netspect",
		6194: "Netspeed",
		2171: "Netstar",
		6857: "Netstor",
		2710: "Neturity",
		5935: "Netvantage",
		5402: "Netvision",
		2722: "Netvox",
		5872: "Netwiz",
		109:  "Network",
		9054: "NetworkAccountant",
		5479: "Networld",
		79:   "Networth",
		6227: "Netzin",
		3186: "NeuLion",
		8872: "Neul",
		1875: "NeuroCom",
		4204: "Neuron",
		3285: "Neuros",
		8002: "Neurotek",
		6233: "Nevatec",
		2984: "Nevis",
		7028: "Nevro",
		8888: "NewSharp",
		2983: "NewSoft",
		7993: "Newag",
		3790: "Newbury",
		2092: "Newcotech",
		5250: "Newer",
		5137: "Newgen",
		7758: "Newings",
		1352: "Newisys",
		1388: "Newport",
		8041: "Newrun",
		925:  "Newtec",
		1659: "Newtech",
		7275: "Newtek",
		8262: "NexFi",
		4072: "NexG",
		2568: "NexQL",
		8838: "Nexar",
		2417: "Nexcom",
		6039: "Nexcomm",
		8424: "Nexell",
		5189: "Nexo",
		6981: "Nexpring",
		614:  "Nexsan",
		615:  "Nexsi",
		5:    "Next",
		3867: "NextGTV",
		1301: "NextGig",
		4133: "NextIO",
		6821: "NextNav",
		5835: "Nextcell",
		255:  "Nextcomm",
		9339: "Nextek",
		3319: "Nexterm",
		5429: "Nextest",
		1859: "Nextlink",
		6081: "Nextone",
		6960: "Nextwill",
		2241: "Nexus",
		6202: "Nexware",
		8757: "Nexxt",
		6677: "Nhnetworks",
		5945: "Nice",
		1693: "NiceTechVision",
		2198: "Nicety",
		3022: "Nicevt",
		4522: "Nicira",
		1745: "Nidek",
		1060: "Nielsen",
		7168: "NietZsche",
		3469: "Nifty",
		1242: "Nihon",
		8335: "Nike",
		8911: "Nikkiso",
		2445: "Niko",
		6514: "Niko-Servodan",
		5662: "Nikon",
		8509: "Nilan",
		8193: "Nilfisk",
		7988: "Nimbus",
		5655: "NineTiles",
		1918: "Ninelanes",
		1427: "Nintendo",
		9024: "Nippon",
		2155: "Nisca",
		5576: "Nishimu",
		7792: "Nishiyama",
		4852: "Nissho-denki",
		1910: "Nissin",
		1744: "Nitgen",
		2965: "Nits",
		3053: "Nittan",
		4902: "Nivetti",
		4504: "Nivis",
		3065: "Nivus",
		6226: "Nixdorf",
		1751: "Nixvue",
		479:  "Nobell",
		8832: "Noblex",
		6556: "Noccela",
		7228: "Nocsys",
		457:  "Nokia",
		2880: "Nokota",
		2981: "Nolan",
		5294: "Nomadix",
		6250: "Nome",
		4447: "Nomus",
		7736: "Noon",
		5930: "Norand",
		4469: "Norcott",
		7268: "Noregon",
		1507: "Noritz",
		7903: "Norma",
		6486: "Norphonic",
		7944: "Nortec",
		2919: "Nortech",
		8134: "Nortek-AS",
		76:   "Nortel",
		2183: "Northover",
		783:  "Northstar",
		6885: "Nothing",
		4340: "NovAtel",
		8843: "Nova",
		7498: "NovaComm",
		1603: "NovaPal",
		9262: "NovaSparks",
		3792: "Novacomm",
		7269: "Novakon",
		5366: "Novalink",
		7717: "Novar",
		1128: "Novasonics",
		1263: "Novatec",
		1214: "Novatechnology",
		3256: "Novatron",
		9247: "NovelSat",
		14:   "Novell",
		7987: "Noviga",
		4468: "Novita",
		2651: "Novomatic",
		962:  "Novra",
		6050: "Novtek",
		5522: "Novus",
		8294: "Noxus",
		8273: "Nsolution",
		8566: "NuLEDs",
		4803: "NuVo",
		770:  "Nuark",
		9113: "Nubia",
		5777: "Nucom",
		4416: "Nucomm",
		3753: "Nucsafe",
		1391: "Nudian",
		2349: "Nuera",
		9025: "Nuheara",
		3386: "Numata",
		2944: "Numatics",
		8665: "Numera",
		2042: "Numesa",
		9158: "Nureva",
		7845: "Nusoft",
		7341: "Nutanix",
		7488: "Nuvico",
		7173: "Nuvolt",
		7600: "Nuvon",
		5552: "Nuvotech",
		7346: "Nuvyyo",
		1724: "Nvergence",
		659:  "Nvidia",
		9316: "NxtConect",
		1991: "O'Rite",
		6959: "O-Net",
		5796: "O.N",
		2618: "O2Micro",
		5955: "OAK",
		8984: "OCOSMOS",
		4721: "OCP",
		7834: "OCT",
		1770: "OCTTEL",
		8543: "ODA",
		4534: "OJ-Electronics",
		5007: "OL'E",
		6736: "OMICRON",
		4799: "OMNI-WiFi",
		3351: "OMNIKEY",
		7101: "OMRON",
		3426: "ON",
		1071: "ONStor",
		7005: "ONvocal",
		7015: "OOSIC",
		2300: "OOmon",
		178:  "OPEN",
		9236: "OPMEX",
		1570: "OPNET",
		2600: "OPTOELECTRONICS",
		7681: "OPWILL",
		1828: "OQO",
		7972: "ORA",
		1763: "ORACOM",
		9293: "ORICO",
		5826: "ORSYS",
		9142: "ORTHOsoft",
		7950: "OSRAM",
		7981: "OT",
		2951: "OTE",
		8310: "OTECTechnology",
		8359: "OTSL",
		6847: "OVH",
		7045: "OWIN",
		8829: "OWLink",
		501:  "Oak",
		8426: "Oakley",
		6387: "Obelux",
		8338: "Obihai",
		2309: "Objective",
		7948: "Objenious",
		7366: "Observint",
		3223: "Obsidian",
		4040: "Obvius",
		2200: "Obzerv",
		395:  "Occam",
		6284: "Occuspace",
		8681: "OceanServer",
		8905: "Ocom",
		4176: "Octagon",
		1822: "Octasic",
		1489: "Octave",
		4292: "Octel",
		8322: "Octonion",
		7082: "Octopod",
		7679: "Octopus",
		5846: "Octothorpe",
		631:  "Ocular",
		4343: "Odva",
		3422: "Oesolutions",
		4184: "Ohler",
		697:  "Ohm",
		6578: "Ohsung",
		7085: "Oilfind",
		5873: "Okuma",
		5226: "Olencom",
		7782: "OleumTech",
		9365: "Olibra",
		5297: "Olicom",
		1533: "Olitec",
		45:   "Olivetti",
		1972: "Olym-tech",
		168:  "Olympus",
		497:  "Omega",
		7434: "Omneality",
		6422: "Omni-ID",
		571:  "OmniCluster",
		8846: "OmniLync",
		3156: "OmniSense",
		5666: "Omnia",
		5842: "Omnibit",
		5113: "Omnibyte",
		7095: "Omnifi",
		1344: "Omnilux",
		7866: "Omnima",
		8806: "Omniprint",
		5191: "Omnisec",
		6796: "Omnisense",
		975:  "Omnitron",
		7219: "Omnitronics",
		2375: "Omnitronix",
		9281: "Omntec",
		4402: "OnLive",
		2704: "OnSite",
		7750: "OnTarget",
		8906: "OnTime",
		4837: "Onbnetech",
		4658: "Onda",
		6828: "One",
		2659: "OneAccess",
		4617: "OnePath",
		6924: "OnePlus",
		4240: "Oneac",
		5501: "Onelan",
		602:  "Oneline",
		7857: "Onface",
		1523: "OngCorp",
		1430: "Onity",
		2253: "Onkey",
		3201: "Online",
		4954: "Onnto",
		4989: "Onprem",
		3511: "Onset",
		2631: "Ontimetek",
		6871: "Onzo",
		3320: "Ooma",
		4612: "Op-Tection",
		3518: "OpVista",
		3999: "Opaque",
		5087: "Opcom",
		1817: "Opelcomm",
		1879: "Open",
		7425: "Open-m",
		5270: "OpenCon",
		2770: "OpenGear",
		2804: "OpenIB",
		7622: "OpenPattern",
		8387: "OpenVox",
		6413: "OpenXS",
		2933: "Openbrain",
		4071: "Openmoko",
		8011: "Openpeak",
		832:  "Opentech",
		1068: "Ophir-Spiricon",
		1345: "Ophit",
		4943: "Opicom",
		723:  "OptXCon",
		8861: "Optcom",
		1124: "Opteon",
		4147: "Optex",
		9502: "Optex-FA",
		928:  "Opthos",
		4929: "Opti",
		1959: "Opti-cell",
		8917: "OptiLogix",
		550:  "OptiMight",
		3517: "Optibase",
		3049: "Optica",
		1484: "Optical",
		3267: "Opticom",
		3812: "Opticomm",
		6298: "Opticore",
		8216: "Optictimes",
		1037: "Optillion",
		5668: "Optim",
		3395: "Optimal",
		5257: "Optimation",
		3844: "Optimedical",
		5176: "Optimem",
		2291: "Optinel",
		5802: "Optiquest",
		2045: "Optium",
		5656: "Optivision",
		5728: "Opto-22",
		9171: "OptoMET",
		7878: "Optonic",
		4315: "Optos",
		2137: "Optoway",
		7805: "Optowiz",
		484:  "Optranet",
		5017: "Optronic",
		5269: "Optronics",
		3748: "Optsys",
		8079: "Opulinks",
		5564: "Opus",
		6777: "Opzoon",
		11:   "Oracle",
		6879: "Oraimo",
		7861: "Orantek",
		9148: "Orb",
		4962: "Orbacom",
		2130: "Orban",
		6230: "Orbis",
		4183: "Orbotech",
		7667: "Orbotix",
		6319: "Orcatech",
		5417: "Orckit",
		3471: "Ordyn",
		6055: "Oresis",
		7025: "Orga",
		8654: "Oriental",
		1078: "Ormazabal",
		5990: "Ormec",
		4581: "Ortana",
		4225: "Osaka",
		5951: "Ositech",
		1125: "Ositis",
		8336: "Osorno",
		5648: "Osprey",
		8020: "Osram",
		897:  "Otanikeiki",
		791:  "Otari",
		3051: "Otsuka",
		9175: "Ottec",
		7176: "OttoQ",
		7721: "Otus",
		2737: "Ouen",
		4463: "Ouman",
		8707: "Ouster",
		7666: "OuterLink",
		1917: "Outline",
		8666: "Ouya",
		5749: "Ovation",
		9424: "Overkiz",
		818:  "Overture",
		4003: "Owitek",
		8746: "Owl",
		1780: "Oxance",
		8457: "Oxide",
		5304: "Oxtel",
		1648: "Oxygnet",
		7652: "P&S",
		2351: "P-CoM",
		346:  "P-Cube",
		8383: "P.T.I",
		7608: "P2",
		177:  "PAC",
		5301: "PAN-International",
		3937: "PAV",
		3218: "PAX",
		287:  "PAXCOMM",
		6799: "PBR",
		1704: "PC-PoS",
		4406: "PCI",
		1218: "PCO",
		4001: "PCS",
		460:  "PCTEL",
		2658: "PDH",
		2822: "PDL",
		5594: "PEC",
		5347: "PERIPHERALS",
		5462: "PFU",
		7051: "PHAZR",
		6140: "PHC",
		7553: "PHYTRONIX",
		4502: "PIMA",
		997:  "PINON",
		3025: "PKC",
		4965: "PLANET",
		1539: "PLAT'C2",
		6907: "PLATH",
		3278: "PLAYLINE",
		8349: "PLC",
		6541: "PLNetworks",
		7087: "PLUMgrid",
		2559: "PLUS",
		2120: "PLX",
		415:  "PLcom",
		3595: "PMC",
		6094: "PMC-Sierra",
		7179: "PNC",
		1758: "PNI",
		9447: "PNY",
		534:  "PORTech",
		7047: "PRAVIS",
		9232: "PRF",
		5482: "PRO-LOG",
		938:  "PRO-NETS",
		8325: "PRO-VISION",
		3575: "PROBA",
		6554: "PROCOM",
		4061: "PROTEI",
		8692: "PSTec",
		8883: "PTCOM",
		9305: "PYRAMID",
		3562: "PacStar",
		7902: "Pace-O-Matic",
		8784: "Pacidal",
		811:  "Pacific",
		8299: "PacketAccess",
		924:  "PacketAir",
		4106: "PacketFlux",
		2948: "PacketHop",
		895:  "PacketLight",
		2718: "PacketMotion",
		7682: "PacketStorm",
		5432: "Packeteer",
		6128: "Pacom",
		3779: "Pado",
		5163: "Pagine",
		4258: "Pairgain",
		8869: "Pakton",
		1158: "Palm",
		2259: "PalmPalm",
		1355: "Palmmicro",
		4138: "Paltronics",
		3301: "PanAccess",
		7013: "Panaccess",
		6497: "Panamax",
		1315: "Panasas",
		2154: "Panasonic",
		8715: "Pandachip",
		4287: "Pandatel",
		7663: "Pandora",
		2232: "Panduit",
		3079: "Pangolin",
		1553: "Pannaway",
		8069: "Panodic",
		9369: "Panoptic",
		7292: "Panoptics",
		2739: "Panta",
		2278: "Pantech",
		5150: "ParTech",
		3856: "Parade",
		2709: "Paradigm",
		6912: "Paradom",
		8042: "Paradox",
		6120: "Paradyne",
		560:  "Paragea",
		352:  "Paragon",
		3111: "Paralan",
		4290: "Paralink",
		3770: "Parallels",
		332:  "Paralon",
		1488: "Parama",
		3199: "Parama-tech",
		7595: "ParanTek",
		2671: "Paravirtual",
		2563: "Parrot",
		6983: "Parsec",
		5504: "Parsytec",
		8511: "Parta",
		8246: "Particle",
		1867: "Partner",
		6240: "Partron",
		2323: "Parvus",
		8237: "Pason",
		1866: "Passave",
		7140: "PassivSystems",
		7922: "Patech",
		9089: "Paterson",
		5341: "Pathlight",
		6066: "Pathway",
		7983: "Patlite",
		5782: "Patton",
		2372: "Pavo",
		7429: "Pax",
		5134: "Paxdata",
		747:  "Paxonet",
		8671: "PayPal",
		7941: "PayRange",
		3406: "PayTec",
		8700: "Payter",
		4890: "Pcube",
		3613: "PePWave",
		1791: "Pedestal",
		5560: "Peer",
		3253: "Peerless",
		3576: "Pegasus",
		5440: "Pegatron",
		758:  "Pelago",
		3035: "Peltor",
		3719: "Pempek",
		8694: "Penetek",
		1255: "Peninsula",
		5822: "Pensando",
		5969: "Pentacom",
		3623: "Pentaone",
		5358: "Pentek",
		6757: "PentronicAB",
		8285: "PeopleNet",
		2482: "Peplink",
		1947: "Pepperl+Fuchs",
		8485: "Pepwave",
		9069: "Pepxim",
		5260: "Peracom",
		6740: "Peraso",
		5751: "Perception",
		4642: "Perceptron",
		6356: "PerfTech",
		2644: "Perfect",
		2978: "Perfisans",
		5913: "Performance",
		2195: "Peribit",
		7810: "Perinet",
		5446: "Periphonics",
		5379: "Perkin-Elmer",
		5062: "Perle",
		6493: "Perma-Pipe",
		6664: "PernixData",
		8050: "Perples",
		3799: "Perq",
		3357: "Persistent",
		5686: "Personal",
		6864: "Pertronic",
		1636: "Pesa",
		990:  "Petards",
		8648: "Petatel",
		1985: "Petcomkorea",
		272:  "Peterson",
		6708: "Petite-En",
		7574: "Petra",
		3637: "Petratec",
		9144: "Pevco",
		3760: "Pfister",
		4114: "Phabrix",
		2733: "Phantom",
		4619: "PharmaSmart",
		7254: "Pharos",
		6529: "PhaseSpace",
		2315: "Phasecom",
		5380: "Phast",
		4926: "Phasys",
		3551: "PheeNet",
		6849: "Phicomm",
		3636: "Phidgets",
		6641: "Philio",
		796:  "Philips",
		7605: "Phison",
		6277: "Phistek",
		3642: "Phoenix",
		2194: "Phonak",
		2877: "Phonic",
		9342: "Phorm",
		7549: "Phorus",
		2240: "Photometrics",
		4450: "Photon",
		681:  "Photonex",
		3136: "Photonicbridges",
		5919: "Photonics",
		516:  "Photron",
		946:  "Photuris",
		1948: "Phsnet",
		4651: "Phybridge",
		4814: "Phyllis",
		7145: "Physio-Control",
		2890: "Physiometrix",
		8272: "Phytium",
		7942: "Phytrex",
		6835: "Pi-Coral",
		7237: "Pica8",
		5214: "Picazo",
		6897: "Piciorgros",
		5555: "Picker",
		456:  "Pico",
		6238: "PicoCELA",
		691:  "PicoLight",
		3028: "Picochip",
		3452: "Picotest",
		2302: "PictureTel",
		9226: "PiiGAB",
		4006: "Pika",
		3109: "Pilkor",
		4809: "Pilot",
		808:  "Piltofish",
		366:  "Pilz",
		458:  "Pinetron",
		4358: "Ping",
		4105: "Pingood",
		9179: "Pingtek",
		5978: "Pingtel",
		4996: "Pinnacle",
		1016: "Piolink",
		6118: "Pioneer",
		9136: "Pioneercorporation",
		1562: "Pipal",
		5597: "Pipelinks",
		8834: "Piper",
		8635: "Piranti",
		8885: "Pishion",
		4544: "Pitronot",
		4736: "Pittasoft",
		3217: "Pivot3",
		5065: "Pivotal",
		922:  "Pivotech",
		5235: "PixStream",
		8755: "Pixavi",
		7320: "PixeLINK",
		6324: "Pixel",
		4744: "Pixel8",
		4063: "Pixelmetrix",
		1196: "Pixelworks",
		2567: "Pixim",
		635:  "Pixord",
		6292: "Pjrc.com",
		4104: "Planar",
		7340: "Planet",
		4479: "Planex",
		1333: "Planmeca",
		6328: "Planning",
		542:  "Plantronics",
		607:  "Plast-Control",
		4703: "Plaster",
		6513: "Plastoform",
		7318: "Platina",
		3827: "Plato",
		1615: "Platypus",
		1038: "Platys",
		9292: "PlayFusion",
		6975: "Plds",
		609:  "Pleiades",
		2437: "Pleora",
		6338: "Plessey",
		5954: "Plexcom",
		6932: "Plexonics",
		1971: "Plexus",
		9125: "Plexxi",
		8147: "Plugable",
		7576: "Pluribus",
		5972: "Pluris",
		2931: "Plus",
		2994: "Plustek.Inc",
		6098: "Pluto",
		8082: "Pocketbook",
		8955: "PoeWit",
		7752: "Poindus",
		4613: "Pointmobile",
		4480: "Polar",
		8399: "PolarLink",
		1213: "Polaris",
		656:  "Polaroid",
		4767: "Pole/Zero",
		749:  "Polestar",
		4503: "Pollin",
		6634: "Polostar",
		7229: "Poly",
		751:  "Polycom",
		9469: "Polyera",
		3645: "Polygon",
		2182: "Polypix",
		7612: "Polytec",
		2076: "Ponico",
		7925: "Portsmith",
		5692: "Portwell",
		7515: "Posbank",
		4126: "Posbro",
		694:  "Poscon",
		2987: "Posdata",
		6743: "Posh",
		3407: "Posiflex",
		4089: "Positron",
		9325: "Poslab",
		3732: "Postek",
		2577: "Posystech",
		6922: "Power",
		1515: "Power-One",
		7649: "PowerCloud",
		4953: "PowerCom",
		7613: "PowerComm",
		3117: "PowerLink",
		3396: "PowerQuattro",
		7654: "PowerRay",
		8785: "Powercode",
		2521: "Powercom",
		4975: "Powerfile",
		7573: "Powerlinq",
		1379: "Powernet",
		2952: "Powertech",
		7401: "Powervision",
		8101: "Poynt",
		5628: "Praxon",
		6804: "Preceno",
		8163: "Precepscion",
		6519: "Precia",
		7892: "Precidata",
		174:  "Precidia",
		263:  "Precision",
		6888: "Precor",
		1908: "Prediwave",
		9127: "Premietech",
		5708: "Premisys",
		7990: "Prescope",
		6006: "Presence",
		8791: "Presition",
		1547: "Presonus",
		4339: "Presstek",
		6164: "Pressure",
		4241: "Presticom",
		4272: "Pretec",
		6067: "Prevas",
		1902: "Primagraphics",
		998:  "Primarion",
		389:  "Primax",
		1040: "Prime",
		1827: "PrimeNet",
		6432: "PrimeVOLT",
		8527: "Primera",
		8693: "Primmcon",
		850:  "Princeton",
		8345: "PrintCounts",
		5072: "Printer",
		4159: "Printrex",
		1246: "Printronix",
		5359: "Prisa",
		909:  "Prism",
		2111: "Privaris",
		67:   "Private",
		6899: "ProLon",
		6146: "ProMax",
		1264: "ProQuent",
		3833: "ProStor",
		2703: "ProTelevision",
		4493: "ProVision",
		7396: "Probedigital",
		4796: "Probits",
		3569: "Procare",
		8339: "Procentec",
		1499: "Procera",
		1429: "Proces-Data",
		740:  "Procket",
		3093: "Prod-El",
		25:   "Prodigy",
		1634: "Productivity",
		3797: "Prodys",
		6727: "Profalux",
		6407: "Proformatique",
		8562: "Progeny",
		2858: "Prokom",
		5246: "Prolific",
		2466: "Proliphix",
		4342: "Prologix",
		6862: "Promate",
		7205: "Promax",
		7739: "Promethean",
		6121: "Prominet",
		4077: "Prominvest",
		207:  "Promise",
		7000: "Promzakaz",
		4210: "Pronet",
		1801: "Pronto",
		7658: "Prophet",
		8049: "Propulsion",
		532:  "Proscend",
		1954: "Prosoft",
		3376: "Prostar",
		5963: "Prosum",
		8636: "Prosyst",
		1329: "Proteam",
		5579: "Protech",
		4487: "Protecta",
		4848: "ProtectedLogic",
		7387: "Protectron",
		99:   "Proteon",
		4568: "Proteus",
		2392: "Protocol",
		5582: "Proton",
		2871: "Protronic",
		4887: "Proventix",
		735:  "Proview",
		3183: "Proware",
		6771: "Prowave",
		4830: "Proxense",
		3653: "Proxicast",
		6119: "Proxima",
		7412: "Proximus",
		9035: "Proxis",
		8673: "Prudential",
		8794: "Prysm",
		8725: "PsiKick",
		2790: "Psia",
		5855: "Psitech",
		2794: "Psitek",
		3306: "Pulsar-Telecom",
		4265: "Pulse",
		672:  "Pulse-Link",
		7250: "PulseOn",
		836:  "Pultek",
		8080: "Pulzze",
		3785: "Pumpkin",
		1223: "PurOptix",
		3261: "PureTech",
		3912: "PureWave",
		4538: "Purechoice",
		777:  "Puretek",
		7774: "Purple",
		4888: "PurpleComm",
		3876: "Pylone",
		6357: "Pyramid",
		5685: "Pyrescom",
		3503: "Pyronix",
		4459: "Pyung-HWA",
		8433: "Q",
		6515: "Q-Lab",
		4582: "Q-Light",
		7887: "Q9",
		3120: "QDI",
		845:  "QEI",
		2026: "QLogic",
		6762: "QNAP",
		311:  "QPS",
		5357: "QSC",
		566:  "QSI",
		9487: "QTS",
		4969: "QTelNet",
		4692: "QUALICA",
		4866: "QVidium",
		2211: "Qamcom",
		7508: "Qardio",
		4090: "Qbit",
		2004: "Qcom",
		8876: "Qdis",
		8100: "Qfiednet",
		6670: "QianGua",
		7916: "QianTang",
		7353: "Qibixx",
		9253: "Qihan",
		6426: "QiiQ",
		2542: "QinetiQ",
		7430: "Qingping",
		551:  "Qisda",
		3825: "Qniq",
		5096: "Qnix",
		3174: "Qno",
		1787: "QoStek",
		7059: "Qolsys",
		4514: "Qool",
		6653: "Qorvo",
		6724: "Qotom",
		2610: "Qovia",
		8096: "Qpcom",
		30:   "Qpsx",
		1174: "Qqest",
		2721: "Qsan",
		7305: "Qsono",
		3173: "Qstik",
		4146: "Qtech",
		3771: "Qtum",
		5434: "QuVis",
		5937: "Quad/Graphics",
		141:  "Quadram",
		3121: "Quadrics",
		1247: "Quake",
		5788: "Qualcomm",
		8795: "Qualisys",
		3488: "Qualitrol",
		1210: "Qualstar",
		9008: "Quanergy",
		8119: "QuantHouse",
		3072: "Quanta",
		5566: "Quantel",
		3094: "Quantier",
		8863: "Quantify",
		2084: "Quantum",
		3877: "QuantumVision",
		7860: "Quantumsolution",
		6676: "Quarion",
		280:  "Quarry",
		3859: "Quartics",
		4735: "Quasar",
		1625: "Quatech",
		7559: "Quatius",
		8295: "Qudelix",
		2508: "Quest",
		5731: "Questech",
		7803: "Quext",
		3224: "QuickTel",
		5171: "Quin",
		5876: "Quintar",
		6391: "Quintic",
		2688: "Quintron",
		7964: "Quirky",
		2897: "Quixant",
		9488: "Qulsar",
		5125: "Qume",
		3549: "Qumranet",
		4709: "QuoPin",
		155:  "Quotron",
		6645: "Qvis",
		6953: "R&M",
		5785: "R.A",
		414:  "RACOM",
		588:  "RADVision",
		2975: "RADWIN",
		7536: "RADiflow",
		2623: "RAE",
		8508: "RAID",
		6503: "RAONIX",
		4812: "RAUMFELD",
		9100: "RBcloudtech",
		105:  "RC",
		3492: "RCG",
		8123: "RDA",
		1402: "RDI",
		2153: "RDM",
		3300: "REJ",
		2119: "REMEC",
		2856: "RF",
		7112: "RFI",
		2954: "RFID",
		4979: "RFM",
		1019: "RFNET",
		573:  "RFTNC",
		3713: "RFTech",
		2428: "RGB",
		5971: "RHK",
		508:  "RIAS",
		4556: "RIM",
		4148: "RIVA",
		3204: "RIX",
		9335: "RLH",
		2996: "RLW",
		5188: "RLX",
		9263: "RM",
		6822: "RNET",
		6810: "RNware",
		8668: "ROBOTIS",
		4027: "ROBOTOUS",
		239:  "ROI",
		871:  "ROIS",
		7248: "ROLI",
		1302: "ROMWin",
		1847: "ROUND",
		4152: "RSD",
		6779: "RSF",
		5255: "RSI",
		8066: "RTC",
		171:  "RTUnet",
		9403: "RTW",
		5569: "RWT",
		3364: "Raba",
		1045: "Racal-Datacom",
		1367: "Racewood",
		8804: "RackTop",
		9451: "RackWare",
		3425: "Rackable",
		8090: "Racktivity",
		58:   "Racore",
		4644: "RadChips",
		3474: "RadarFind",
		4956: "Radcom",
		8334: "Raden",
		5812: "Radguard",
		7763: "RadiAnt",
		2757: "Radiance",
		228:  "Radiant",
		1576: "Radiantech",
		5057: "Radicom",
		3931: "Radiient",
		8152: "Radinet",
		3343: "RadioCOM",
		2966: "RadioPulse",
		4087: "Radiocomp",
		2941: "Radiocrafts",
		5793: "Radiolan",
		1138: "Radionet",
		3941: "Radionor",
		7303: "Radios",
		3901: "Radioscape",
		52:   "Radisys",
		684:  "Radius",
		1081: "Radlan",
		4329: "Radlive",
		6703: "Radmax",
		7853: "Radspin",
		5517: "Radstone",
		563:  "Radware",
		9194: "Radwin",
		2338: "Radyne",
		7913: "Rafael",
		6840: "Ragentek",
		7837: "Ragile",
		8441: "Ragsdale",
		4214: "Ragula",
		4647: "Raidon",
		2409: "Raidtec",
		7307: "RailComm",
		7106: "Railtec",
		7907: "RainUs",
		1538: "Rainsun",
		2051: "Raisecom",
		8044: "Rajant",
		1785: "Ralink",
		6243: "Ramaxel",
		9277: "Rami",
		5048: "Ramix",
		23:   "Ramtek",
		1837: "Rancho",
		8379: "Rancore",
		5420: "Randata",
		2547: "Rane",
		5825: "Rantic",
		9319: "RaonThink",
		1186: "Raonet",
		901:  "RapidWAN",
		8786: "Rapidmax",
		2239: "Raptor",
		1927: "Raritan",
		626:  "Rasteme",
		1228: "Rasvia",
		5563: "Rational",
		8472: "Ratp",
		5622: "Rauland-Borg",
		8017: "RayTight",
		9133: "Raybased",
		3477: "Raycom",
		5131: "Raylan",
		4429: "Raylase",
		8257: "Raylios",
		4285: "Raynet",
		362:  "Raysis",
		2602: "Rayson",
		5877: "Raytech",
		96:   "Raytheon",
		4650: "Rayzone",
		7182: "Razer",
		3819: "Razorstream",
		8706: "Rcntec",
		6913: "Reach",
		6532: "Reacheng",
		2765: "ReadyLinks",
		250:  "ReadyNet",
		4009: "Realease",
		7189: "Reallin",
		2702: "Realm",
		2331: "Rebel.com",
		4022: "Recall",
		8230: "Rechnerbetriebsgruppe",
		9455: "Reciprocal",
		487:  "Reco",
		6384: "Recovision",
		2982: "Red-Lemon",
		1485: "Red-M",
		5713: "Redcom",
		5356: "Redcreek",
		761:  "Reddo",
		4992: "Redflex",
		7686: "Redflow",
		1314: "Redline",
		885:  "Rednix",
		910:  "Redswitch",
		2603: "Redux",
		9346: "Reduxio",
		9252: "Redwire",
		1932: "Redwood",
		7557: "Refined",
		238:  "Refraction",
		6865: "Regenersis",
		5980: "Regent",
		3373: "Reigncom",
		739:  "Relax",
		8624: "Relay2",
		2376: "Reliance",
		1722: "Remopro",
		2945: "Remote",
		1281: "Remotec",
		1509: "Remotek",
		2101: "Remsdaq",
		2803: "Renasis",
		5696: "Renesas",
		5775: "Renex",
		882:  "Renishaw",
		3726: "Renkus-Heinz",
		3274: "Renu",
		1201: "Repeatit",
		9141: "Repotec",
		5532: "Republic",
		4552: "ResMed",
		8623: "Research",
		5987: "Resideo",
		5386: "Resilience",
		6191: "Reson",
		4167: "Respironics",
		6738: "RetailNext",
		3023: "Reti",
		4233: "Reudo",
		5804: "Reuters",
		8572: "Revolv",
		3002: "Rextechnik",
		7824: "Rezolt",
		5574: "Rftech",
		7199: "Rhino",
		8902: "Rhombus",
		1707: "RiT",
		6608: "Riava",
		4165: "Riccius+Sohn",
		75:   "Ricoh",
		6342: "Ridge",
		3453: "Riedel",
		8991: "Riedo",
		1769: "Rifatron",
		8232: "Rigado",
		1181: "Rigaku",
		3584: "RightHand",
		5035: "Rightech",
		3480: "Rigol",
		1658: "Rincon",
		6947: "Ring",
		4785: "RingBell",
		4740: "RingCube",
		7714: "Rinicom",
		7764: "Rinstrum",
		2326: "Rion",
		5223: "Rioworks",
		3682: "Ripcode",
		194:  "RiscStation",
		4068: "Risco",
		2404: "Rise",
		6717: "Risk",
		3238: "Rittmeyer",
		4251: "Riva",
		5014: "RiverDelta",
		2094: "Riverbed",
		1823: "Riverhead",
		394:  "Riverstone",
		1811: "Rivertec",
		2801: "Rivertree",
		8340: "Rivet",
		862:  "Roax",
		3001: "Robatech",
		956:  "Robinson",
		4792: "Robonica",
		2181: "Robox",
		9029: "Rockeetech",
		662:  "RocketLogix",
		5997: "Rocketchips",
		6287: "Rockport",
		3872: "Rockridgesound",
		4404: "Rohati",
		3866: "Rohm",
		1916: "Roku",
		1790: "Roll",
		1874: "Rolls-Royce",
		6096: "Rooftop",
		7632: "RoomReady/Zdi",
		5975: "Root",
		6966: "Rootech",
		8487: "Roqos",
		8179: "Rorze",
		646:  "Rosco",
		5145: "Rose",
		4094: "Roseman",
		4822: "Rosemount",
		7636: "Rosewill",
		9422: "Rosonix",
		5726: "Ross",
		9438: "Rosslare",
		8456: "Rossma",
		4718: "Rotel",
		1267: "RouteFree",
		1784: "Routerboard.com",
		507:  "Routrek",
		951:  "Roving",
		7880: "Roxton",
		4078: "RoyalTek",
		3155: "Royaldigital",
		2411: "Rpcg",
		7519: "Rramac",
		9353: "RtBrick",
		7494: "Rtnet",
		7156: "Rubezh",
		4629: "Ruby",
		2738: "Ruckus",
		1586: "RuggedCom",
		5441: "Ruijie",
		4732: "Runcom",
		524:  "Runtop",
		8583: "Ruroc",
		4376: "Russound",
		4494: "Rustelcom",
		8799: "Rusteletech",
		2763: "RyCo",
		6394: "Ryowa",
		2779: "Ryvor",
		3890: "S&O",
		4789: "S-Access",
		7552: "S-Bluetech",
		2808: "S-TEC",
		1777: "S-Takaya",
		4636: "S.A.A.A",
		1969: "S.A.Tehnology",
		330:  "S.D.E.L",
		5862: "S.E.R.C.E.L",
		8886: "S.E.Technologies",
		7164: "S.FAC",
		2405: "S.I",
		211:  "S1",
		7443: "S2M",
		1882: "S2io",
		3360: "S3C",
		3745: "SAE",
		5502: "SAI",
		7297: "SAM",
		6812: "SAMJIN",
		4587: "SAMSUNG",
		8369: "SAMWONFA",
		2430: "SANBlaze",
		3662: "SANION",
		1923: "SANYCOM",
		6175: "SAT",
		3531: "SATEC",
		344:  "SAXA",
		9416: "SB",
		135:  "SBE",
		8730: "SBM",
		3970: "SBN",
		5977: "SBS",
		6700: "SCAPS",
		4885: "SCDI",
		8923: "SCR",
		8171: "SCS",
		8836: "SCSpro",
		7615: "SDJ",
		6133: "SDL",
		1232: "SDSystem",
		6416: "SDTEC",
		4742: "SE-Elektronic",
		5089: "SEA-Ilan",
		463:  "SEAKR",
		7463: "SECUDOS",
		1461: "SED",
		2958: "SEECODE",
		4431: "SEEnergy",
		8651: "SEMA",
		9389: "SEMOCON",
		5445: "SENTRY",
		979:  "SEPATON",
		4125: "SERONICS",
		8737: "SES-imagotag",
		4949: "SET",
		307:  "SETA",
		6773: "SEnergy",
		6395: "SF",
		4173: "SFA",
		1865: "SFOM",
		7057: "SFORZATO",
		3188: "SFR",
		24:   "SHA-KEN",
		7879: "SHIFT",
		6205: "SI",
		1618: "SIBCO",
		764:  "SICOM",
		7646: "SIEMENS",
		7214: "SIFROM",
		1836: "SIGMACOM",
		4861: "SIL3",
		3654: "SIMS",
		8940: "SINTRONES",
		7153: "SK",
		6208: "SK-Elektronik",
		268:  "SKF",
		3099: "SKNET",
		8919: "SM-Electronic",
		4972: "SMAR",
		388:  "SMART",
		9002: "SMAX",
		741:  "SMC",
		6642: "SMG",
		244:  "SMK-M",
		4924: "SMP",
		4420: "SMT&C",
		9157: "SNB",
		6538: "SNK",
		3934: "SNR",
		896:  "SNS",
		1467: "SOHOware",
		4869: "SOKRAT",
		8281: "SOL",
		7279: "SOLARWATT",
		7452: "SOLEM",
		2276: "SOLOMON",
		9104: "SOLiD",
		1550: "SONICblue",
		932:  "SONO",
		648:  "SONOS",
		7742: "SOUND4",
		8927: "SOYEA",
		4896: "SP",
		1237: "SPAUN",
		6114: "SPC",
		5490: "SPHINX",
		2775: "SPIDCOM",
		9350: "SPON",
		6458: "SPRINGWAVE",
		8703: "SPS",
		5237: "SPX-Ateg",
		7651: "SPnS",
		8078: "SRC",
		3081: "SSD",
		4625: "SSI",
		4756: "ST",
		8086: "STABILO",
		2743: "STAC",
		9282: "STI",
		4760: "STJ",
		6528: "STMicrolectronics",
		3966: "STN",
		3067: "STNet",
		7617: "STORDIS",
		5798: "STS",
		7896: "STULZ",
		5565: "SUN",
		7143: "SUNGSAM",
		1472: "SUNIX",
		3060: "SUNWAVETEC",
		3851: "SUPoX",
		363:  "SURECOM",
		7529: "SUYIN",
		1486: "SVA",
		1049: "SVA-Intrusion",
		8514: "SVS-VISTEK",
		3047: "SWEEX",
		3311: "SWsoft",
		9453: "SYN-Tech",
		9408: "SYNTEC",
		3880: "SYRIS",
		2712: "Saab",
		6336: "Saber",
		3920: "SabiOso",
		9173: "Sabre",
		6575: "Sabtech",
		2033: "Safari",
		5344: "Safe-Com",
		7936: "SafeTone",
		1978: "SafeWeb",
		5989: "Safetran",
		7350: "Safetrust",
		8024: "Saffron",
		3504: "Sagamore",
		619:  "Sage",
		4573: "SageTV",
		2049: "Sagemcom",
		6587: "Sagittar",
		4449: "Sagrad",
		5394: "Sahara",
		6615: "Sailer",
		7159: "Salcomp",
		3493: "Salicru",
		5261: "Salix",
		1845: "Salland",
		7970: "Salutron",
		7309: "SamJi",
		8841: "SambaNova",
		4570: "Samjeon",
		1768: "Sammy",
		2971: "Sampo",
		5587: "Samsan",
		9496: "Samsara",
		6169: "Samson",
		152:  "Samsung",
		6276: "Samtec",
		3621: "Samyoung",
		4020: "San-Eisha",
		608:  "SanCastle",
		5333: "SanCom",
		3658: "SanDisk",
		8663: "SanJet",
		7722: "SanLogic",
		4530: "SandForce",
		4350: "SandLinks",
		710:  "SandStream",
		2210: "Sandburst",
		1901: "Sanden",
		2129: "SandmartinElectronics",
		7531: "Sandstone",
		1347: "Sandvine",
		9415: "Sanford",
		1258: "SangSang",
		3868: "Sangean",
		8478: "Sangfor",
		5229: "Sangoma",
		9042: "Sanix",
		617:  "Sanko",
		6442: "Sankosha",
		7881: "SankyuElectronics",
		2209: "Sanmei",
		3537: "Sanmina-SCI",
		1046: "Sanritz",
		7278: "Sanshin",
		4844: "Sansonic",
		3987: "Santec",
		4977: "Santera",
		9454: "Santur",
		1337: "Sanyo",
		2607: "Sanyu",
		7525: "Sapling",
		9099: "Sapphire",
		637:  "Sarian",
		6171: "Sarnoff",
		2271: "Sarotech",
		4259: "Sast",
		888:  "Satec",
		4362: "Satel",
		7951: "Satelco",
		5447: "Satelcom",
		8010: "Satmap",
		3465: "Sato",
		2962: "Saunders",
		3589: "Savant",
		6562: "Savitech",
		2820: "Savvius",
		7926: "Sawwave",
		683:  "Scalant",
		7794: "Scalar",
		2087: "Scalent",
		7726: "Scalys",
		3687: "Scan",
		5466: "Scan-Optics",
		5351: "Scanivalve",
		1903: "Scanmatic",
		417:  "Scannex",
		2676: "Scanvaegt",
		1976: "Schiller",
		685:  "Schlumberger",
		4778: "Schmartz",
		56:   "Schneider",
		6585: "Schreder",
		5006: "Schweitzer",
		4477: "SciLog",
		3037: "Science",
		8961: "Scientech",
		1474: "Scientific",
		3116: "Scientific-Atlanta",
		4317: "Scimolex",
		4198: "Scinets",
		1624: "Scion",
		7069: "Sciovid",
		57:   "Scitex",
		5898: "Scope",
		6241: "Scopus-Belgium",
		2809: "Scosche",
		1593: "ScottCare",
		3265: "Screen",
		8985: "ScreenBeam",
		1656: "ScriptPro",
		6421: "SeAH",
		4333: "SeaMicro",
		1569: "Seabridge",
		732:  "Seagate",
		7538: "Seal",
		1473: "Sealevel",
		6272: "Seamap",
		4706: "Seanodes",
		1884: "Seaway",
		1386: "SecWell",
		590:  "Secheron",
		2465: "Secom-Industry",
		7491: "Secret",
		4908: "Sectronic",
		8897: "SecuGen",
		7029: "SecuWorks",
		807:  "Secui.com",
		2970: "Securaplane",
		1382: "Securebase",
		5047: "Securelogix",
		7358: "Securetech",
		9181: "Securifi",
		1225: "Securiton",
		8773: "Securosys",
		3257: "Sedo",
		5038: "Sedona",
		359:  "Seedek",
		933:  "Seedsware",
		6894: "Seeed",
		8790: "SeekTech",
		2846: "SeekerNet",
		5924: "Seel",
		5754: "Seeq",
		4702: "Seer",
		6409: "Seers",
		8910: "Seetech",
		6088: "Sega",
		2347: "Segate",
		3796: "Segnet",
		4095: "Segnetics",
		5461: "Seiko",
		119:  "Seikosha",
		814:  "Seiwa",
		4347: "Sejin",
		4381: "Sekonic",
		4070: "Select",
		187:  "Selectron",
		3181: "Selex",
		2410: "Selsius",
		3564: "Seluxit",
		6667: "Selve",
		5891: "Semaphore",
		4889: "Semihalf",
		3667: "Semindia",
		4330: "Semptian",
		3140: "Semtech",
		9286: "SenRa",
		4750: "SenTec",
		245:  "Sena",
		383:  "Senao",
		937:  "Sencore",
		2756: "Sendo",
		195:  "Sendtek",
		3149: "Senea",
		603:  "Seneca",
		4714: "Senet",
		5979: "Senetas",
		2886: "Senhai",
		8564: "Senient",
		3118: "Sennheiser",
		9358: "Senor",
		1694: "SensAble",
		6689: "Sensaio",
		1177: "Sensaphone",
		4867: "Senscient",
		9211: "Sense",
		4335: "SenseAnywhere",
		8588: "Senseit",
		8501: "Senselogix",
		3411: "Sensicast",
		8721: "SensingTek",
		3445: "SensoPart",
		9332: "Sensometrix",
		6739: "SensorPush",
		7039: "SensorTec-Canada",
		829:  "Sensoria",
		5303: "Sensormatic",
		2613: "Sensory",
		2827: "Sensovation",
		3871: "Sensus",
		9078: "Sentec",
		4529: "Senticare",
		5363: "Sentient",
		2605: "Sentilla",
		9219: "Sentinhealth",
		9485: "Sentinum",
		3248: "Sentivision",
		7149: "Seongji",
		1774: "SeorimTechnology",
		8243: "Seoul",
		4300: "Seowonintech",
		8969: "Seowoo",
		2906: "Sepsa",
		4017: "Sepura",
		3045: "Sequans",
		4312: "Sequel",
		6325: "Sequent",
		1442: "Seranoa",
		1728: "SercoNet",
		2075: "Sercomm",
		8768: "Serelec",
		4383: "SerialTek",
		7627: "Serialway",
		4169: "Seritech",
		6256: "Sernet",
		355:  "Serome",
		6247: "Sertel",
		1552: "Server",
		3108: "ServerEngines",
		8031: "ServerU",
		8146: "Servercom",
		7048: "Servergy",
		4058: "Servimat",
		3082: "SetOne",
		2576: "SetaBox",
		4879: "Setrix",
		7582: "Seungil",
		2473: "Sevis",
		7207: "Sewoo",
		4983: "Seyeon",
		5009: "Shanghai",
		6651: "Shannon",
		6534: "Shany",
		894:  "ShareGate",
		6025: "Sharewave",
		6463: "SharkNinja",
		3206: "Sharp",
		5607: "Shasta",
		6433: "Shaw",
		1044: "Sheba",
		191:  "Shelcad",
		771:  "Shellcomm",
		7293: "ShengHai",
		8626: "Shenzhen",
		3619: "Shenztech",
		6143: "Sherwood",
		2032: "Shester",
		1350: "ShibaSoku",
		6168: "Shimadzu",
		1457: "Shimizu",
		8421: "Shin-IL",
		1597: "Shin-OH",
		3009: "ShinMaywa",
		8479: "Shinbo",
		2455: "Shinboram",
		3163: "Shinco",
		7446: "Shineway",
		8898: "Shiningtek",
		5548: "Shinnihondenko",
		4162: "Shinsei",
		3907: "Shinwa",
		7895: "Shinybow",
		9087: "Shinyei",
		3976: "Shireen",
		5549: "Shiva",
		5135: "Shographics",
		5419: "Shomiti",
		8465: "Shoogee",
		2322: "ShoreTel",
		8973: "ShotSpotter",
		7695: "ShotTracker",
		4526: "Shouyo",
		2036: "Shuko",
		3816: "Shunra",
		2118: "Shure",
		4938: "Shuttle",
		4392: "Si14",
		367:  "SiByte",
		548:  "SiConnect",
		2799: "SiCortex",
		2572: "SiNett",
		3143: "SiRF",
		7517: "Siama",
		963:  "Sick",
		8802: "Siconix",
		2811: "Sidsa",
		74:   "Siecor",
		4517: "Sielox",
		300:  "Siemens",
		3997: "Siemon",
		329:  "SierraCom",
		7374: "Sify",
		3282: "SightLogix",
		291:  "Sigma",
		1241: "Sigma-Links",
		1802: "SigmaTel",
		4016: "Sigmalink",
		6576: "Sigmastar",
		5300: "Sigmatek",
		144:  "Sigmex",
		569:  "Signal",
		4119: "Signalion",
		3904: "Signamax",
		3663: "Signtech",
		5290: "Signum",
		3042: "Sigpro",
		7454: "Sigrand",
		9296: "Sigrist-Photometer",
		120:  "Siig",
		4673: "Siklu",
		6997: "Silca",
		5519: "Silex",
		6214: "Silicom",
		1655: "Silicon",
		3154: "SiliconStor",
		3377: "Silicondust",
		2585: "Silink",
		7285: "Silkan",
		7316: "SilverNet",
		7516: "SilverPlus",
		7910: "Silverbrook",
		7177: "Silverflare",
		4157: "Simet",
		3519: "Simoco",
		9352: "Simon-Kaloi",
		8231: "SimonsVoss",
		39:   "Simpact",
		4171: "Simple",
		3644: "SimpleComTools",
		279:  "Simpler",
		9412: "SimpliSafe",
		5630: "Simrad",
		2512: "Simtec",
		8085: "Simton",
		5585: "Simulation",
		2180: "Sinar",
		1451: "Sinbon",
		3293: "Sindoricoh",
		1101: "Sinetica",
		7899: "Sinewy",
		4741: "Singapore",
		2491: "Singim",
		2490: "Singular",
		3649: "Sinkyo",
		7563: "Sino-Telecom",
		3925: "Sinotech",
		8980: "Sinwatec",
		9327: "Sipod",
		3242: "Sirit",
		8185: "Sirius",
		8460: "Siselectron",
		8199: "Sisnet",
		260:  "Sitara",
		6339: "Sitasys",
		8264: "Sitcorp",
		1881: "Sitecom",
		1366: "Sitecsoft",
		5032: "Sitek",
		301:  "Sitera",
		9115: "Sitronik",
		7211: "Sius",
		6713: "Sivantos",
		2566: "Skardin",
		686:  "Skidata",
		7359: "Skipper",
		8122: "Skiva",
		6606: "Skorpios",
		2110: "Skov",
		6263: "Skspruce",
		7040: "Skullcandy",
		6998: "Sky-City",
		7685: "SkyBell",
		8343: "SkyDisk",
		8027: "SkyHawke",
		8958: "Skybell",
		4118: "Skydigital",
		7001: "Skydio",
		7304: "Skyera",
		6914: "Skymotion",
		4939: "Skystream",
		8393: "Skytap",
		9019: "Skyviia",
		1706: "Smallbig",
		3456: "SmarDTV",
		3240: "SmarTire",
		7322: "Smart",
		6742: "SmartCap",
		8331: "SmartDoor",
		6783: "SmartDrive",
		8266: "SmartOptics",
		4550: "SmartRG",
		4136: "SmartShare",
		4143: "SmartSynch",
		6790: "SmartThings",
		4937: "Smartbridges",
		9289: "SmarteBuilding",
		8599: "Smartisan",
		2134: "Smartlabs",
		1015: "Smartmatic",
		7381: "Smarto",
		775:  "Smartronix",
		8528: "Smartrove",
		6045: "Smartsan",
		2766: "Smartvue",
		5000: "Smartware",
		6749: "Smnd",
		6269: "Smobile",
		8772: "SnD",
		8296: "Snap",
		6557: "SnapAV",
		7146: "SnapRoute",
		1320: "SnedFar",
		822:  "Snell",
		5410: "Snmp",
		753:  "SnowShore",
		7218: "Snuza",
		6614: "Soarnex",
		9489: "Socionext",
		5861: "Socket",
		3197: "Socomec",
		7911: "Socus",
		2330: "Sodick",
		1290: "SofaWare",
		3629: "Sofacreal",
		1220: "SoftEnergy",
		1162: "SoftRadio",
		5539: "Softcom",
		3297: "Softier",
		959:  "Softing",
		9164: "Softiron",
		5322: "Softlab",
		4609: "Softwell",
		8999: "Sogecam",
		3723: "Sogestmatic",
		1441: "Solacom",
		7541: "Soladigm",
		3209: "Solar",
		4907: "SolarEdge",
		7979: "Solarbridge",
		1524: "Solarflare",
		95:   "Solbourne",
		5896: "Solectek",
		9373: "SolidFire",
		8943: "SolidRun",
		7621: "Solidica",
		6079: "Solidum",
		9400: "Solidwintech",
		874:  "Solinet",
		3650: "Solitech",
		6884: "Soliton",
		5050: "Sollae",
		4344: "Soltech",
		2436: "Solteras",
		4387: "Solutronic",
		369:  "Soma",
		6011: "Somat",
		5887: "Somelec",
		7302: "Somfy",
		7108: "Sonance",
		7442: "Sonar",
		8162: "Sonardyne",
		9052: "Sonavation",
		1279: "Soneticom",
		5067: "Sonic",
		8384: "SonicSensory",
		3263: "SonicWALL",
		6640: "SonicWall",
		1003: "Sonicwall",
		8952: "Sonifex",
		4672: "Sonim",
		1372: "Sonitor",
		3459: "Sonitrol",
		4244: "Sonix",
		4999: "Sonnet",
		1309: "SonoSite",
		2932: "Sonoa",
		6361: "Sonoma",
		2048: "Sonos",
		6860: "Sonova",
		9398: "Sontex",
		2340: "Sonus",
		101:  "Sony",
		5540: "Sophia",
		3579: "Sophos",
		7308: "Soraa",
		5312: "Sord",
		2275: "Sordin",
		7234: "Soreel",
		1230: "Sorenson",
		1073: "Soriya",
		864:  "Soronti",
		5791: "Sotas",
		5311: "Sotec",
		8388: "SoundBridge",
		4415: "SoundEar",
		7703: "SoundHawk",
		2270: "Soundcraft",
		6868: "Soundmatters",
		6485: "Soundmax",
		6527: "Source",
		4175: "Source-Comm",
		9312: "Sourcefire",
		6823: "SourcingOverseas",
		8102: "Sovico",
		2701: "Soyal",
		5209: "Soyo",
		6483: "Spacelink",
		1977: "Spagat",
		2074: "SparkLAN",
		3229: "Sparr",
		7180: "Spawn",
		8682: "Spbec-Mining",
		6297: "Speaker",
		4091: "Speakercraft",
		5542: "Specialix",
		7511: "Speco",
		2655: "Spectec",
		6047: "Spectel",
		2588: "Spectra",
		32:   "Spectragraphics",
		5634: "Spectralink",
		4286: "Spectrix",
		9051: "Spectronix",
		5531: "Speed",
		8823: "Speedytel",
		3755: "Sphairon",
		6102: "Sphere",
		8839: "Spica",
		154:  "Spider",
		3889: "Spinetix",
		622:  "Spinnaker",
		2334: "Spirent",
		1152: "Splicecom",
		8699: "SpotCam",
		7117: "Spreadtrum",
		6820: "Sprocomm",
		5730: "Spur",
		3410: "Sputnik",
		6752: "Squarehead",
		7472: "Squirrels",
		2838: "Srisa",
		5094: "Ssangyong",
		3435: "Staccato",
		5323: "Stallion",
		8221: "Stalmart",
		3483: "Stanford",
		5736: "Stanilite",
		5101: "Star",
		7195: "Star-Net",
		5107: "Star-TEK",
		4864: "StarLeaf",
		9235: "StarTech.com",
		3717: "StarVedia",
		4084: "Starbridge",
		5046: "Stardot",
		793:  "Starent",
		1561: "Stargames",
		8333: "Starkey",
		5156: "Starlight",
		2366: "Starnet",
		2537: "Starnex",
		8145: "Starry",
		9259: "Startel",
		8674: "Stateless",
		5810: "Staubli",
		5153: "Stealth",
		1529: "Stec",
		7759: "Steelcase",
		7364: "Steffes",
		4208: "Steinbrecher",
		8916: "Steinel",
		9039: "Steinsvik",
		7125: "Stella-Green",
		5541: "Stellar",
		533:  "Stellcom",
		7868: "Stephen",
		7716: "Stereotaxis",
		7886: "Sterlite",
		3974: "Stim",
		2604: "Stoke",
		2918: "Stolinx",
		4052: "Stoneridge",
		7441: "Stonesoft",
		1835: "StorCase",
		7238: "StorSimple",
		2399: "Storage",
		8849: "Storagedata",
		7245: "Store",
		1974: "Stormshield",
		4637: "Storwize",
		1098: "Stralfors",
		2956: "StrataLight",
		6015: "Stratabeam",
		6893: "Stratacache",
		5881: "Stratacom",
		6035: "Strategy",
		719:  "Stratex",
		108:  "Stratus",
		1483: "Stream",
		9479: "StreamCCTV",
		5422: "StreamLogic",
		2188: "StreamScale",
		7319: "StreamUnlimited",
		2444: "Streamit",
		2179: "Stretch",
		944:  "Strix",
		4876: "Strong",
		9379: "Structab",
		4422: "Strukton",
		1989: "Stryker",
		9178: "Stuart",
		632:  "Studio",
		8278: "Sub10",
		7704: "Subscriber",
		4806: "Suga",
		714:  "Sullair",
		4655: "Sumavision",
		872:  "Sumtel",
		819:  "Sun",
		2650: "SunCorp",
		2813: "SunKwang",
		4498: "SunPower",
		7713: "SunReports",
		2412: "Sundance",
		3754: "Sunell",
		8859: "Sunflex",
		4232: "Sungwoon",
		3080: "Sunhillo",
		3940: "Sunitec",
		2755: "Sunmyung",
		5436: "Sunnovo",
		2427: "Sunplus",
		2286: "Sunray",
		8934: "Sunrex",
		1575: "Sunrich",
		8866: "Sunrise",
		3950: "Sunshine",
		9185: "Sunstar",
		9147: "Suntaili",
		8324: "Suntec",
		6285: "Suntech",
		7637: "Sunwave",
		3148: "Sunways",
		7772: "Sunwoda",
		2835: "SuperVision",
		1365: "Supercaller",
		4015: "Supercom",
		5561: "Supercomputing",
		86:   "Supernet",
		4928: "Superpower",
		3277: "Suprema",
		19:   "Sureman",
		2622: "Surf",
		703:  "Surgient",
		5476: "Surigiken",
		2324: "Surtec",
		1114: "Sutron",
		4474: "Sutus",
		6438: "Suunto",
		3603: "Suzuken",
		5900: "Svec",
		6390: "Svyazkomplektservice",
		6404: "Swaive",
		8723: "Swann",
		2608: "Swegon",
		5308: "Swelaser",
		7273: "Sweroam",
		8317: "SwiftTest",
		3374: "Swirlnet",
		8132: "Swissbit",
		4097: "Swissdis",
		835:  "Swissvoice",
		5264: "Switchcore",
		3676: "Swyx",
		1028: "Syabas",
		3960: "Syba",
		5173: "Sybus",
		6929: "Sycada",
		4215: "Sycamore",
		1667: "Sychip",
		4945: "Sylantro",
		4345: "SymCom",
		7252: "Symanitron",
		1990: "Symantec",
		6013: "Symbionics",
		6302: "Symbolics",
		7389: "Symeo",
		1864: "Symetrix",
		5328: "Symicron",
		122:  "Symmetric",
		4295: "Symmetrical",
		5746: "Symmetricom",
		6186: "Symon",
		5186: "Symplex",
		2841: "Symwave",
		4026: "Symx",
		6774: "SynTrust",
		5056: "SynapSense",
		3758: "Synapse",
		7570: "Synaptec",
		7118: "Synapticon",
		7971: "Synaptics",
		5120: "Sync",
		9204: "Syncbak",
		3146: "Synccom",
		2150: "Synchronic",
		801:  "Synchronous",
		2231: "Synchrony",
		5355: "Synclayer",
		7327: "Syncmold",
		1145: "Synectics",
		7068: "Synerchip",
		7548: "Synergics",
		4239: "Synergy",
		5091: "Synerjet",
		5477: "Synernetics",
		9151: "Synertau",
		2451: "Synology",
		1102: "Synoptics",
		7512: "Syntech",
		5942: "Syntellect",
		40:   "Syntrex",
		6784: "Syntronic",
		4746: "Syphan",
		2507: "Sypixx",
		4500: "Syracuse",
		3258: "Syrinx",
		7023: "Syrotech",
		6252: "SysDINE",
		8283: "SysGRATION",
		60:   "SysKonnect",
		7255: "SysTec",
		5116: "Sysgen",
		9050: "Syslane",
		2789: "Sysmaster",
		5998: "Sysmate",
		955:  "Sysmex",
		1266: "Syspol",
		1257: "Systec",
		5480: "Systech",
		1737: "Systegra",
		1554: "Systek",
		4545: "Systel",
		578:  "SystemGear",
		1324: "SystemK",
		9105: "Systembase",
		5149: "Systemforschung",
		7465: "Systemhouse",
		1187: "Systemonic",
		7635: "Systems",
		2875: "Systimax",
		5679: "Systran",
		6410: "Systrome",
		8874: "Systronics",
		1617: "Systronix",
		3795: "Syswan",
		209:  "Syswave",
		7623: "Syszone",
		6:    "Sytek",
		1953: "T&D",
		4049: "T&W",
		7105: "T-Mobile",
		4845: "T-Platforms",
		2844: "T-Vips",
		4191: "T.C",
		5558: "T.D.I",
		4835: "T.M",
		2581: "T.O.M",
		2301: "T.Sqware",
		7580: "T3",
		2380: "TAC",
		5263: "TAG",
		7333: "TAKT",
		503:  "TAMI",
		3746: "TANITA",
		4659: "TAP.tv",
		3775: "TASA",
		8313: "TASCAN",
		2443: "TASI",
		8520: "TATUNG",
		4035: "TBTech",
		8597: "TC",
		139:  "TCL",
		3888: "TCM",
		7456: "TCPlink",
		6434: "TCT",
		698:  "TD",
		3147: "TDA",
		6692: "TDC",
		5521: "TDK",
		3512: "TDK-Lambda",
		8844: "TDSi",
		199:  "TDT",
		2161: "TEAL",
		7321: "TECHNART",
		576:  "TECOM",
		7607: "TEKTELIC",
		7402: "TELCO",
		1727: "TELDIX",
		2050: "TELEFIELD",
		1007: "TELEM",
		181:  "TELENET",
		3727: "TELENOT",
		4540: "TEM",
		2520: "TEN",
		8182: "TES",
		1967: "TElectronics",
		6289: "THUB",
		2664: "THX",
		1193: "TIL",
		4921: "TITECH",
		7132: "TITENG",
		193:  "TIW",
		4898: "TKM",
		5338: "TKS",
		6054: "TL",
		7440: "TLS",
		8486: "TM-Research",
		8373: "TMCT",
		7765: "TMRG",
		1313: "TMT",
		1931: "TMT&D",
		9207: "TMY",
		6530: "TNK",
		2374: "TNS",
		892:  "TOA",
		8975: "TOHO",
		6934: "TOP-Access",
		1595: "TP-Link",
		3096: "TPS",
		3643: "TPine",
		2905: "TRENDnet",
		4592: "TRG",
		121:  "TRI-Data",
		8414: "TRIZ",
		6222: "TRL",
		9503: "TRONTEQ",
		9109: "TRP",
		477:  "TRsystems",
		4940: "TSI",
		6347: "TSL",
		9215: "TSMART",
		6965: "TTE",
		1082: "TURCK",
		4452: "TV-Numeric",
		5365: "TV/CoM",
		4606: "TVLogic",
		5475: "TVS",
		3359: "TVT",
		3048: "TVWorks",
		5027: "TWO",
		8506: "TXTR",
		2922: "TZero",
		2109: "Tableau",
		2720: "Tabor",
		4930: "Tachion",
		1376: "Tachyon",
		8226: "Taco",
		943:  "Tactel",
		8483: "Tactical",
		2019: "Tadlys",
		85:   "Tadpole",
		1606: "Taekwang",
		2479: "Taelim",
		3315: "TagMaster",
		8801: "Tagatec",
		1024: "Tahoe",
		4892: "TaiDoc",
		6254: "TaiYear",
		9459: "Taian",
		1676: "Taifatech",
		3318: "Taiguen",
		806:  "Tailyn",
		6827: "Taimag",
		2454: "Taishin",
		1984: "Tait",
		4004: "Taiwick",
		4195: "Taiyo",
		3548: "Takacom",
		736:  "Takagi",
		3921: "Takahata",
		966:  "Takasago",
		584:  "Takaya",
		7217: "Talari",
		65:   "Talaris",
		4103: "Talent",
		8577: "Taleo",
		4053: "Talk-A-Phone",
		2186: "TalkSwitch",
		8254: "Tallac",
		7470: "Talon",
		6136: "Talx",
		8109: "Tamaggo",
		9318: "Tamio",
		8070: "Tamron",
		6594: "Tamtron",
		5915: "Tamura",
		6359: "Tandem",
		8988: "TangoWiFi.com",
		8884: "Tangtop",
		2753: "Tanisys",
		4747: "Tantalus",
		1950: "Tapwave",
		6029: "Taqua",
		1505: "Targa",
		8023: "Tariox",
		4131: "Taseon",
		7041: "Tatsuno",
		5478: "Tatung",
		6200: "Tazmo",
		7575: "Taztag",
		730:  "Tdsoft",
		3524: "Teak",
		7306: "Team",
		4590: "Team-R",
		3372: "Teamcast",
		4160: "Tecan",
		9243: "Tecc",
		8015: "Tech4home",
		4112: "TechNexion",
		9084: "TechSAT",
		3910: "TechTrex",
		6669: "Techaya",
		3545: "Techcity",
		3165: "Techfaithwireless",
		6408: "Techman",
		8261: "Techmation",
		2717: "Techmetro",
		7006: "Technica",
		2925: "Technical",
		6393: "Technicolor",
		6372: "Technity",
		1185: "Techno-Holon",
		7383: "Techno-Innov",
		1757: "Techno-One",
		4773: "TechnoDigital",
		225:  "TechnoLand",
		2345: "Technobox",
		431:  "Technocom",
		6023: "Technologic",
		2785: "Technolution",
		7243: "Technomate",
		7720: "Technonia",
		1377: "Technoventure",
		699:  "Technovision",
		3540: "Technowave",
		667:  "Techsan",
		2935: "Techsphere",
		4282: "Techware",
		4443: "Techway",
		6345: "Tecmar",
		8311: "Tecmobile",
		5514: "Tecnetics",
		6266: "Tecno",
		8255: "Tecnobit",
		6106: "Tecnomen",
		1750: "Tecnova",
		3368: "Tecobest",
		6735: "TeconGroup",
		7157: "Tecore",
		7797: "Tecsen",
		1534: "Tecton",
		6672: "Tegile",
		3666: "Tehuti",
		7230: "Tek-Air",
		5928: "Teklogix",
		5681: "Teknema",
		8750: "Teko",
		4046: "Tekon-Automatics",
		7344: "Tekpea",
		5372: "Tekram",
		3905: "Tekron",
		4783: "Tektrap",
		6304: "Tektronix",
		7642: "Teladin",
		3704: "Telchemyorporated",
		18:   "Telco",
		1779: "TelcoBridges",
		5714: "Teldat",
		8008: "TeleAdapt",
		418:  "TeleCruz",
		436:  "TeleDream",
		4029: "TeleWell",
		1334: "Telebau",
		5187: "Telebit",
		1370: "Telebyte",
		2464: "Telecard-Pribor",
		7076: "Telechips",
		3572: "Teleco",
		1132: "Telecom",
		3571: "Telecomunication",
		6832: "Telecor",
		1336: "Telecore",
		6417: "Telecsys",
		865:  "Telect",
		7272: "Teledata",
		1470: "Teledex",
		3234: "Teledyne",
		8718: "Teleepoch",
		7730: "Telefield",
		274:  "Teleforce",
		4388: "Telegesis",
		680:  "Telelynx",
		5665: "Telemann",
		61:   "Telematics",
		1061: "Telemax",
		981:  "Telemonitor",
		2754: "Telemotive",
		1293: "Telena",
		650:  "Telencomm",
		4367: "Telephonics",
		4116: "Teleplan",
		7092: "Telepower",
		4193: "Teleprocessing",
		4150: "Telerad",
		6220: "Teles",
		5939: "Telesciences",
		4464: "Telesis",
		473:  "Telesoft",
		9026: "Telesquare",
		5616: "Teleste",
		5619: "Telestream",
		5353: "Telesync",
		2501: "Telesynergy",
		7251: "Teletics",
		3678: "Teletrak",
		5013: "Teletrol",
		113:  "Televideo",
		6706: "Teleview",
		8450: "Telewave",
		1409: "Telewise",
		1679: "Telex",
		2426: "Telexy",
		4970: "Telgen",
		3761: "Telian",
		5051: "Telica",
		1762: "Telio",
		5253: "Telkom",
		1537: "Telkonet",
		5028: "Tellabs",
		8542: "Telldus",
		2117: "Tellion",
		1815: "Tellium",
		4308: "Tellord",
		3031: "Tellumat",
		5695: "Tellus",
		6180: "Telmax",
		2408: "Telocityorporated",
		6174: "Telogy",
		6108: "Telrad",
		525:  "Telsey",
		6195: "Telsis",
		17:   "Telsist",
		1810: "Telson",
		5165: "Telspec",
		6788: "Telstra",
		5288: "Telstrat",
		3986: "Teltonika",
		5705: "Teltrend",
		7376: "Teltronic",
		290:  "Teltronics",
		3679: "Telular",
		7258: "Telvent",
		976:  "Telways",
		5321: "Telxon",
		7385: "Tely",
		1809: "Tempearl",
		7226: "Tempered",
		8865: "Tempo",
		2947: "TenX",
		6267: "Tenda",
		3928: "Tendril",
		6495: "Tendyron",
		9427: "Tenebraex",
		3062: "Teneros",
		3852: "Tenlon",
		5799: "Tennyson",
		5024: "Tenor",
		2784: "Tenosys",
		1076: "Tenovis",
		9376: "Tensorcom",
		6726: "Tenstorrent",
		2716: "Tentaculus",
		7937: "Tenyu",
		696:  "Teo",
		1273: "Tepg-US",
		1169: "TeraBurst",
		5040: "TeraForce",
		314:  "TeraGlobal",
		2361: "TeraLogic",
		3502: "TeraMage",
		3017: "TeraRecon",
		4821: "Teracom",
		5093: "Teradata",
		4435: "Teradici",
		1536: "Teradon",
		1663: "Teralink",
		2885: "Teranetics",
		8094: "Teraon",
		2999: "Terascala",
		6679: "Terasic",
		8617: "Teraspek",
		6064: "Teratech",
		4109: "Teraview",
		5053: "Terawave",
		6957: "Teraworks",
		8830: "Tercel",
		6972: "Terewave",
		2736: "Termtek",
		3821: "Terra",
		9382: "TerraSem",
		1560: "TerraTec",
		2552: "Terrasat",
		7569: "Terumo",
		3810: "Tervela",
		7052: "Tescom",
		8258: "Teseq",
		7310: "Tesla",
		1678: "Test-Um",
		1406: "Testech",
		3429: "Testo",
		6251: "Testop",
		2252: "Tevebox",
		1203: "Texa",
		470:  "Texcel",
		313:  "Texio",
		4880: "Thales",
		8657: "Thalmic",
		6092: "Thamway",
		6217: "That",
		7502: "Theben",
		2927: "Thecus",
		8271: "Theobroma",
		7476: "There",
		380:  "Thermalogic",
		2591: "Thermo",
		233:  "ThermoQuest",
		8712: "ThinGlobal",
		6264: "ThinPAD",
		3559: "Thincom",
		2626: "ThingMagic",
		9256: "Things",
		7329: "ThingsMatrix",
		7986: "ThinkEco",
		4564: "ThinkFlood",
		1171: "Thinkengine",
		7206: "Thinking",
		4096: "Thinkware",
		3268: "Thinlinx",
		5459: "Thomas-Conrad",
		9088: "Thomas-Krenn.AG",
		1142: "Thomson",
		6930: "Thread",
		9203: "Throughtek",
		3205: "ThruVision",
		8632: "Thuh",
		8352: "Thum+Mahr",
		9311: "Thundercomm",
		610:  "TiMetra",
		2529: "TiVo",
		7094: "Tiandy",
		5227: "Tiara",
		4660: "Tibbo",
		3259: "Tibetsystem",
		939:  "Tiburon",
		3606: "Tidel",
		5385: "Tidomat",
		6153: "Tiernan",
		1925: "Tiesse",
		506:  "Tietech",
		6315: "Tigan",
		3237: "Tigi",
		6540: "Tiinlab",
		3604: "Tilera",
		377:  "Tilgin",
		9200: "Time-O-Matic",
		4134: "TimeIPS",
		4411: "TimeKeeping",
		5765: "TimeStep",
		29:   "Timeplex",
		2248: "Timespace",
		453:  "Timeware",
		3683: "Tintometer",
		9433: "Tintri",
		1858: "Tiptel",
		3353: "Tiqit",
		9490: "TireCheck",
		738:  "Titan",
		2178: "Tivella",
		2536: "Tixi.com",
		2695: "ToGoldenNet",
		4689: "Tobii",
		7840: "Todaair",
		3160: "Tohken",
		3131: "Toho",
		6836: "Tokheim",
		5319: "Tokimec",
		4971: "Toko",
		1205: "Tokyo",
		5212: "TollBridge",
		2317: "Tollgrade",
		3894: "Tom",
		2715: "TomTom",
		6353: "Tomen",
		387:  "Tommy",
		6491: "Tonal",
		8287: "Tongfang",
		6895: "Tonly",
		3255: "Tonze",
		6901: "Top-Unum",
		4355: "TopControl",
		6375: "Topaz",
		1090: "Topcall",
		8489: "Topcon",
		2081: "Topfield",
		852:  "Topspin",
		1664: "Topview",
		8945: "Topwell",
		2819: "Toradex",
		5816: "Toray",
		4614: "Toro",
		6320: "Torus",
		35:   "Toshiba",
		6507: "Tosibox",
		426:  "Totsu",
		6400: "Totus",
		6814: "Touch",
		921:  "TouchStar",
		6131: "Touchwave",
		7400: "Toyo",
		1262: "Toyo-Linx",
		734:  "Toyokeiki",
		7464: "Tozo",
		2060: "Tpack",
		2365: "Tracewell",
		7448: "TrackNet",
		3089: "Trackflow",
		5469: "Tradpost",
		7315: "TrafficCast",
		2742: "TrafficSim",
		2044: "Trajet",
		6278: "Traka",
		8658: "Trakm8",
		5717: "Trancell",
		2657: "Trane",
		296:  "Trango",
		2055: "TransCore",
		5267: "TransMedia",
		7731: "TransPacket",
		2039: "Transact",
		4307: "Transcon",
		8077: "Transics",
		5958: "Transition",
		5718: "Transitions",
		4486: "Translogic",
		5831: "Transmeta",
		5806: "Transmitton",
		989:  "Transmode",
		1443: "Transtech",
		5139: "Transware",
		5413: "Transys",
		3805: "Transystem",
		7566: "Tranwo",
		3566: "Tranzas",
		1612: "Trapeze",
		4086: "Travelping",
		556:  "Traxit",
		7373: "Treehouse",
		6936: "Treeview",
		7194: "Trek",
		8532: "TrekStor",
		3985: "TrellisWare",
		7107: "Tremol",
		6017: "Tremon",
		176:  "Trend",
		8691: "TrendPoint",
		2342: "Trenton",
		261:  "Trex",
		1058: "Tri-M",
		7806: "Tri-Sen",
		8918: "Tri-Systems",
		8029: "Tri-Tech",
		1695: "TriBeam",
		434:  "TriState",
		6985: "Tributary",
		9032: "Tricascade",
		5711: "Tricord",
		310:  "Tridium",
		8591: "Tridonic",
		5083: "Trigem",
		391:  "Trilithic",
		1585: "Trilliant",
		8235: "Trilobit",
		1312: "Trilogy",
		1375: "Trimble",
		3945: "Trimm",
		9238: "Trinus",
		8580: "Tripwire",
		8196: "Trison",
		5166: "Tritec",
		8363: "Triteka",
		8395: "Triton",
		8495: "Tritonwave",
		6288: "Triumph-Adler",
		4875: "Trixell",
		866:  "Triz",
		5624: "Troika",
		1240: "Tropic",
		6880: "True",
		6401: "Truebeyond",
		3782: "Truen",
		4829: "Truesell",
		5011: "TrunkNet",
		8889: "Trusonus",
		2511: "Trustable",
		1059: "Trutzschler",
		7554: "Trymus",
		7393: "Tsat",
		9471: "Tsingtong",
		3651: "Tsubata",
		5231: "Tundo",
		7368: "TurControlSystme",
		1896: "Turbo",
		3464: "TurboChef",
		4957: "TurboComm",
		884:  "TurboWave",
		5041: "Turbonet",
		9237: "Turbostor",
		265:  "Turin",
		5627: "Turnstone",
		6701: "Turtle",
		5866: "Tutankhamon",
		7426: "Tuttnaer",
		6476: "Tvip",
		3169: "Twig",
		1280: "TwinHan",
		1696: "TwinMOS",
		184:  "Twinhead",
		9027: "Twinlinx",
		7428: "Twpi",
		6159: "Tyan",
		4462: "Tyco",
		7558: "Tzukuri",
		6501: "U&U",
		2533: "U-MEDIA",
		8811: "U-Raku",
		2989: "U-WAY",
		8833: "U2S",
		2997: "U4EA",
		3702: "UBI&MOBI",
		2217: "UBSTORAGE",
		8131: "UCI",
		7506: "UCZOON",
		6574: "UCloud",
		5097: "UDC",
		7650: "UGame",
		1251: "UHD-Elektronik",
		1895: "UHS",
		5104: "UKG",
		8429: "ULTIMEDIA",
		6994: "ULVAC",
		3420: "UMB",
		7814: "UNINET",
		1846: "UNION",
		991:  "UNIQA",
		8303: "UNITEC",
		1508: "UNITEK",
		8575: "URadio",
		4284: "USC",
		1135: "UTStarcom",
		6516: "UTT",
		7003: "UTran",
		2882: "UbONE",
		1832: "UbeeAirWalk",
		3182: "Ubicod",
		8240: "Ubicquia",
		9309: "Ubiik",
		1464: "Ubinetics",
		9357: "Ubiqam",
		3074: "Ubiquam",
		2979: "Ubiquiti",
		6898: "Ubiquitous",
		1103: "Ubiquoss",
		2525: "Ubisense",
		3529: "Ubistar",
		5128: "Ubitrex",
		9481: "Ubivelox",
		3288: "Ubixon",
		7263: "Ubizcore",
		6537: "Ubyon",
		3784: "Ucamp",
		8537: "Ufine",
		9228: "Ufispace",
		528:  "Ulan",
		7415: "Ulterius",
		5617: "Ultimate",
		3430: "Ultra",
		7869: "UltraClenz",
		2078: "Ultracker",
		6012: "Ultrak",
		3716: "Ultratec",
		3733: "Ultratronik",
		2450: "Unatech",
		7909: "Unetconvergence",
		256:  "Unex",
		5848: "Ungermann-Bass",
		5697: "Uni-Link",
		4630: "Uni-v",
		8474: "UniComm",
		493:  "UniData",
		8935: "UniPrint",
		8535: "Uniband",
		6961: "Unicard",
		2510: "Uniclass",
		1080: "Unico",
		6605: "Unicoi",
		5786: "Unicomputer",
		9172: "Unicore",
		5289: "Unicorn",
		6097: "Uniden",
		8018: "Unidis",
		4034: "Unifat",
		5045: "Uniform",
		322:  "Unify",
		4549: "Unigen",
		5923: "Unigraf",
		3279: "Unigrand",
		7939: "Unikey",
		5884: "Unimicro",
		1462: "Unimo",
		3506: "Unionman",
		8057: "Unipattern",
		4305: "Uniphone",
		3601: "Unipoint",
		5370: "Unipulse",
		6415: "Uniqoteq",
		2421: "Unique",
		7564: "Unisem",
		5592: "Unisphere",
		38:   "Unisys",
		9046: "Unit-EM",
		6141: "Unitec",
		3046: "Unitech",
		1826: "United",
		6445: "Unitend",
		3993: "Unitron",
		1900: "Unitronics",
		6326: "Univation",
		3858: "Universal",
		2258: "Uniwell",
		918:  "Uniwide",
		474:  "Uniwill",
		557:  "Unixtar",
		172:  "Unizone",
		8127: "Unmonday",
		6842: "Unowhy",
		4144: "Up-Today",
		4297: "UpdateLogic",
		2166: "Uplogix",
		147:  "Upnod",
		958:  "Upponetti",
		7347: "Uptmate",
		9005: "Upvel",
		8948: "Upwis",
		3018: "Uquest",
		2937: "Uriel",
		3804: "Uriver",
		4048: "Urmet",
		7233: "Uros",
		8003: "Usag",
		6964: "UserGate",
		6568: "Usys",
		8192: "Utek",
		5993: "Utilicom",
		1424: "Utility",
		8516: "Utillink",
		6165: "Utstarcom",
		5911: "Uunet",
		4437: "Uwin",
		7609: "V",
		6072: "V-Bits",
		2643: "V-Show",
		9461: "V-ZUG",
		5814: "V.I",
		9167: "V2",
		1996: "VAC",
		8896: "VAIO",
		3769: "VDG-Security",
		398:  "VDSL",
		6715: "VECOW",
		6382: "VEO-Labs",
		4163: "VIA",
		7201: "VIG",
		8733: "VILLBAU",
		9338: "VIPAR",
		3133: "VK",
		5914: "VMX",
		809:  "VMware",
		9414: "VNL",
		7277: "VOGTEC",
		5843: "VOIM",
		1190: "VOIX",
		6967: "VPI",
		5381: "VPNet",
		3731: "VR",
		4313: "VRmagic",
		5305: "VSK",
		2681: "VSST",
		5966: "VST",
		6333: "VTA",
		8321: "VTC",
		7821: "VTS",
		1319: "VTech",
		8297: "VVDN",
		3392: "VVOND",
		4743: "VXi",
		7882: "Vachen",
		4359: "Vacon",
		2687: "VadaTech",
		8178: "Vadaro",
		9162: "Vaddio",
		5658: "Vadem",
		8052: "Vaillant",
		6016: "Valcom",
		2290: "Valcretec",
		3437: "Valemount",
		9347: "ValenceTech",
		7727: "Valentine",
		4711: "Valiant",
		5664: "Valid",
		6915: "Validus",
		8864: "Valink",
		792:  "Valley",
		8568: "Vallox",
		1535: "Valo",
		2698: "Valox",
		6425: "Valuable",
		2599: "Value",
		9011: "ValueHD",
		2463: "ValuePoint",
		8521: "Valueplus",
		9124: "Valve",
		3609: "Vamp",
		887:  "Vanderbilt",
		3175: "Vansco",
		4019: "Vantanol",
		6631: "Vanu",
		500:  "Vaone",
		5894: "Vari-Lite",
		5375: "Varian",
		8582: "Varikorea",
		9448: "Variscite",
		5473: "Varityper",
		3069: "Vativ",
		3909: "VaultStor",
		1157: "Vbrick",
		3307: "Vecima",
		3665: "Vector",
		2314: "Vectron",
		5444: "Veea",
		8386: "Veedims",
		6686: "Veethree",
		8084: "Velankani",
		8589: "Vello",
		9315: "VeloCloud",
		6286: "Velocytech",
		7539: "Velodyne",
		7597: "Velux",
		7968: "Vence",
		3725: "Vencer",
		3578: "Venergy",
		7732: "Venetex",
		4715: "Venntis",
		2442: "Venstar",
		2953: "Ventus",
		2197: "Veo",
		3830: "VerScient",
		9096: "Verana",
		2708: "Verascape",
		9411: "Verathon",
		2787: "VeriWave",
		8895: "Verifi",
		1647: "Verifone",
		4213: "Verilink",
		1118: "Verint",
		6049: "Veris",
		4458: "Verismo",
		666:  "Veristar",
		6352: "Veritas",
		3321: "Veritech",
		6622: "Verizon",
		9146: "Verkada",
		3380: "Verkerk",
		1616: "Vernier",
		6876: "Veroguard",
		6619: "Veros",
		8415: "Vers",
		4966: "Versa",
		724:  "VersaLogic",
		4274: "Versalynx",
		3930: "Versamed",
		6196: "Versanet",
		2418: "Vertical",
		6956: "Vertu",
		868:  "Verytech",
		4541: "Vestac",
		1449: "Vestel",
		5638: "Vetronix",
		7080: "Vexata",
		2821: "ViPowER",
		7614: "ViVOtech",
		2923: "ViXS",
		393:  "ViaClix",
		5827: "ViaGate",
		3728: "ViaLogy",
		6201: "ViaVideo",
		6489: "Viaanix",
		4716: "Viaas",
		8612: "Viableware",
		8382: "Vialis",
		2635: "Vialta",
		5005: "Vianet",
		5784: "Viasatorporated",
		8428: "Vibicom",
		1381: "Vibration",
		3998: "Vibro-Meter",
		1481: "Vichel",
		5644: "Vickers",
		5325: "Vicom",
		952:  "Vicon",
		3881: "Viconics",
		7901: "Vicos",
		4271: "Victron",
		9241: "VidaBox",
		5456: "Videcom",
		5151: "Video",
		9230: "VideoHome",
		5407: "VideoServer",
		5248: "Videocon",
		361:  "Videoframe",
		5239: "Videojet",
		8660: "Videoswitch",
		4366: "Videotec",
		716:  "Videotek",
		2070: "Videotron",
		1407: "Videx",
		3946: "Vidient",
		1686: "Vidisco",
		464:  "Viditec",
		2296: "Vienna",
		5995: "Vieo",
		9456: "Vievu",
		1614: "ViewSonic",
		3554: "ViewTel",
		8089: "Viewcooper",
		8061: "Vieworks",
		2497: "Viewtran",
		6648: "Viking",
		9493: "Vimar",
		3113: "Vimicro",
		3122: "Vimtron",
		5343: "Vina",
		1930: "Vinci",
		1584: "Vindicator",
		2108: "VineSys",
		8135: "Viogem",
		3694: "Violin",
		4275: "Vipa",
		3367: "Viprinet",
		8006: "Viptela",
		2571: "Virbiage",
		379:  "Virditech",
		6919: "Virgilant",
		8301: "Virident",
		4615: "Virtual",
		7098: "Virtualtek",
		8800: "Virtuozzo",
		8156: "VisSim",
		8908: "Viscount",
		2642: "Viseon",
		356:  "Visicom",
		1644: "Visimetrics",
		2690: "Vision",
		789:  "VisionTek",
		7567: "VisionVera",
		2018: "Visionary",
		5306: "Visioncomm",
		4023: "Visioneering",
		4934: "Visionetics",
		4895: "Visionhitech",
		1434: "Visionics",
		3591: "Visionite",
		6707: "Visionscape",
		1645: "Visiowave",
		1399: "Visteon",
		20:   "Visual",
		8538: "Visualedge",
		3381: "Visualgate",
		2625: "Vita",
		5533: "Vitacom",
		6129: "VitalCom",
		284:  "VitalPoint",
		7055: "VitalThings",
		6350: "Vitalink",
		4764: "Vitality",
		217:  "Vitana",
		7288: "Vitec",
		2895: "Vitelcom",
		2884: "Vitelec",
		8540: "Vitsmo",
		7260: "Vity",
		2530: "Vivaas",
		413:  "Vivace",
		7379: "Vivago",
		8645: "Vivalnk",
		8404: "Vivatel",
		1633: "Vivato",
		2632: "Vivavis",
		5834: "Viveris",
		6063: "Vivid",
		1392: "VividLogic",
		674:  "Vivity",
		6370: "Vivo",
		7544: "Vivonic",
		440:  "Vivotek",
		3455: "Vivox",
		6235: "Viwone",
		1776: "Vixen",
		6566: "Vixtel",
		6760: "Vizeo",
		4753: "Vizimax",
		3468: "Vizio",
		7851: "Vizmonet",
		8662: "Vlatacom",
		5988: "Vlsi",
		5342: "Vmetro",
		8376: "Vnpt",
		1456: "Vocera",
		2816: "Vocollect",
		9064: "Vodia",
		519:  "Vodtel",
		5733: "Voelker",
		5275: "Voiceboard",
		8332: "Voippartners",
		8744: "Voismart",
		6729: "Volacomm",
		2395: "Volamp",
		9429: "Volans",
		6543: "Volex",
		1607: "Volktek",
		3669: "Vololink",
		8761: "Volta",
		1304: "Voltaire",
		9351: "Volterra",
		5774: "Vorax",
		1632: "Vormetric",
		1006: "Vorne",
		6607: "VostroNet",
		3582: "Votronic",
		5985: "Voxent",
		1608: "Voxpath",
		3640: "Voxtel",
		4641: "Voyant",
		889:  "Vrcom",
		8165: "Vsoontech",
		2503: "Vtech",
		1960: "Vtera",
		7640: "Vubiq",
		5437: "Vutrix",
		8302: "Vuzix",
		3661: "W&W",
		1000: "W-Link",
		2168: "W-Linx",
		1548: "W2",
		8900: "WAAV",
		2769: "WACOM",
		8518: "WAMA",
		4594: "WAREMA",
		8749: "WATERWORLD",
		2257: "WAVE",
		5586: "WAVTrace",
		3399: "WB",
		8397: "WBS",
		123:  "WD",
		5003: "WEBGATE",
		3106: "WEBIO",
		7007: "WEG",
		1459: "WELL",
		1327: "WEM",
		2807: "WEPIO",
		9386: "WEY",
		4841: "WFE",
		201:  "WIN",
		292:  "WINCOMM",
		4395: "WIRECOM",
		1159: "WIS",
		1899: "WISCORE",
		7421: "WISEWARE",
		2847: "WJ",
		7161: "WKK",
		5778: "WMS",
		8238: "WOM",
		5060: "WONWOO",
		1705: "WOOJU",
		7628: "WOORI",
		7571: "WOORISYSTEMS",
		8032: "WOXTER",
		6378: "WSH",
		2107: "WTSS",
		7187: "Wacom",
		9295: "Wafa",
		1525: "Walchem",
		9298: "Waldo",
		5625: "WalkAbout",
		8190: "Wally",
		136:  "Wang",
		2961: "Wanshih",
		1306: "Wany",
		9491: "Wapice",
		8326: "Warehouse",
		2177: "Wasabi",
		8236: "Wata",
		175:  "WatchGuard",
		7527: "Water-i.d",
		7677: "WaterFurnace",
		2007: "Watertek",
		559:  "Watlow",
		7698: "Wattty",
		3561: "Wave",
		5362: "WaveAccess",
		6236: "WaveIP",
		7474: "WaveLynx",
		624:  "WaveSmith",
		5400: "WaveSpan",
		2152: "WaveSplitter",
		3767: "WaveStorm",
		9349: "Wavelink",
		5315: "Wavenet",
		2043: "Waveplus",
		5670: "Waverider",
		2762: "Wavesat",
		598:  "Wavesight",
		8584: "Wavetel",
		3875: "Wavetrend",
		6695: "Wavetronix",
		8704: "WayTools",
		264:  "Wayport",
		6291: "We",
		4399: "WeLink",
		8625: "WeTelecom",
		8967: "Wearable",
		6699: "Wearhaus",
		7121: "Wearsafe",
		3137: "Weathernews",
		6232: "WebSilicon",
		2360: "WebSonic",
		6218: "WebSprint",
		6104: "WebTV",
		1471: "WebWayOne",
		7786: "Weber-Stephen",
		6154: "Webgear",
		2472: "Webpro",
		12:   "Webster",
		5397: "Webtronics",
		890:  "Webyn",
		1112: "Wegener",
		7584: "Weidu",
		544:  "Weinschel",
		1767: "Weintek",
		4657: "Weinzierl",
		3786: "Weiss",
		3404: "Welcat",
		3349: "Weldex",
		3388: "Welding",
		7271: "Welgate",
		8377: "Wellav",
		7795: "Wellcore",
		257:  "Welltech",
		4194: "Welltronix",
		9114: "Wellysis",
		7884: "Welotec",
		3565: "Weltec",
		4354: "Wescon",
		9334: "Wesine",
		6769: "Westcontrol",
		2274: "Westell",
		5824: "Westport",
		7660: "Westunitis",
		1010: "Westwave",
		1623: "Wetek",
		8337: "Whaley",
		4557: "Whdi",
		750:  "WhereNet",
		8270: "WhereWhen",
		8107: "Whirlpool",
		5352: "Whistle",
		5675: "Whitecross",
		8797: "WhizNets",
		2967: "Wi-Gear",
		3988: "Wi-Links",
		8570: "Wi-NEXT",
		3458: "Wi2Wi",
		7657: "Wi3",
		6986: "WiBotic",
		3826: "WiChorus",
		3772: "WiDeFi",
		7718: "WiFiSONG",
		7034: "WiFiSong",
		2630: "WiLife",
		4137: "WiMate",
		3476: "WiQuest",
		4236: "WiSE",
		8229: "WiSilica",
		4111: "WiWide",
		7178: "WiZ",
		7479: "Wiatec",
		3956: "Wibrain",
		1469: "Widax",
		518:  "Widcomm",
		4633: "Wide",
		2558: "WideRay",
		2224: "WideView",
		9033: "Wideband",
		2541: "Wideful",
		7119: "Widex",
		8362: "Wieson",
		7773: "Wifi-soft",
		5820: "WigWag",
		8167: "Wiio",
		2861: "Wiline",
		7875: "Will-Burt",
		2390: "Willnet",
		5661: "Willowbrook",
		787:  "Willowglen",
		6279: "Wilocity",
		6172: "Wiltron",
		3823: "Win4NET",
		967:  "WinCom",
		3339: "WinNet",
		1789: "Winbest",
		2303: "Winbond",
		6741: "Wincal",
		4324: "Winchester",
		5068: "Windata",
		1198: "Windigo",
		9381: "WindowMaster",
		8355: "Winduskorea",
		3176: "Winegard",
		8239: "Winfirm",
		4031: "Wingtech",
		3405: "Winix",
		2894: "Wink",
		3781: "Winland",
		593:  "Winmate",
		1952: "Winners",
		1574: "Winnow",
		941:  "Winpresa",
		7985: "Winstars",
		196:  "Winsystems",
		2208: "Wintec",
		3038: "Wintecronics",
		8965: "Wintop",
		6020: "Wintriss",
		4693: "Winward",
		3957: "Winy",
		1909: "Wiplug",
		1235: "Wipotec",
		7131: "WiredIQ",
		231:  "Wireless",
		7093: "WirelessWERX",
		705:  "Wirelink",
		3329: "Wiremold",
		7793: "Wirepas",
		8390: "Wiscloud",
		4199: "Wisdm",
		3854: "Wiseblue",
		547:  "Wisi",
		7236: "Wisol",
		1592: "Wistron",
		2397: "Witcom",
		4113: "Witelcom",
		1282: "With-Net",
		4701: "Withings",
		8188: "Withrobot",
		4506: "Witron",
		4730: "Wittenstein",
		6991: "WizLAN",
		7372: "Wizardlab",
		8585: "Wizitdongdo",
		2887: "Wizlogics",
		1292: "Wiznet",
		3630: "Wizyoung",
		8594: "Woan",
		8647: "Wohler",
		1824: "WolfVision",
		6368: "WondaLink",
		8767: "WonderSound",
		7610: "Wonderlan",
		4831: "Wongs",
		3273: "WooJooIT",
		7265: "Woodstream",
		6601: "Woojin",
		3059: "Woojinnet",
		1206: "Wooksung",
		401:  "Woorigisool",
		5001: "Workbit",
		441:  "Workstation",
		2160: "WorldAccxx",
		5609: "WorldCast",
		371:  "WorldGate",
		8661: "Worldplay",
		3624: "Woven",
		3325: "WowWee",
		9066: "Wush",
		7335: "WyTec",
		1808: "Wybron",
		8490: "Wyconn",
		849:  "Wyle",
		8857: "Wyler",
		4705: "Wynmax",
		9195: "WyreStorm",
		5496: "Wyse",
		6955: "Wytek",
		6882: "Wyze",
		1677: "X-CoM",
		198:  "X-traWeb",
		4380: "X2E",
		7790: "X6D",
		8646: "XADA",
		8418: "XAG",
		190:  "XAVi",
		4951: "XCP",
		5295: "XEL",
		3208: "XENOLINK",
		339:  "XESystems",
		6281: "XIAOMI",
		3354: "XIP",
		5801: "XKL",
		4024: "XLN-t",
		6003: "XN",
		6844: "XOS",
		4901: "XRPLUS",
		3299: "XStreamHD",
		8081: "XTA",
		5557: "XTP",
		8454: "XTS",
		2285: "XTecorporated",
		8729: "XTrillion",
		3482: "XYnetsoft",
		4227: "Xact",
		4917: "Xagyl",
		1002: "Xalted",
		2176: "Xalyo",
		2170: "Xambala",
		3066: "Xanboo",
		5907: "Xante",
		3020: "Xantech",
		8000: "Xapt",
		9210: "Xaptec",
		8581: "Xaptum",
		6179: "Xaqti",
		5367: "Xata",
		9313: "Xcellen",
		7392: "Xcheng",
		3144: "Xcute",
		5906: "Xedia",
		1825: "Xeline",
		4551: "Xembedded",
		3703: "XenICs",
		6294: "Xena",
		4140: "Xenatech",
		7022: "Xenox",
		3076: "Xensource",
		0:    "Xerox",
		7802: "Xetawave",
		9324: "Xi3",
		2116: "XiNCOM",
		5698: "Xiaomi",
		1498: "Xilinx",
		6478: "Ximea",
		5465: "Xinetron",
		7210: "Xingfei",
		8956: "Xingluo",
		4121: "Xiotech",
		6084: "Xiox",
		3449: "Xipher",
		2980: "Xiranet",
		788:  "Xircom",
		2215: "Xirrus",
		5277: "Xitron",
		7867: "Xlab",
		4577: "Xmark",
		6234: "Xmit",
		773:  "Xnet",
		8992: "XonTel",
		7581: "Xorcom",
		3932: "Xortec",
		5438: "Xovis",
		486:  "Xpeed",
		5760: "Xpoint",
		5192: "Xrite",
		5984: "Xrosstech",
		561:  "Xsense",
		1496: "Xsido",
		2358: "Xstreamis",
		3070: "Xteam",
		755:  "Xtera",
		4461: "Xtramus",
		8385: "Xtreme",
		1756: "XtremeSpectrum",
		6586: "Xtremio",
		9017: "Xunison",
		8759: "XySystem",
		5309: "Xycom",
		798:  "Xycotec",
		5470: "Xylogics",
		5770: "Xyplex",
		5520: "Xyron",
		6358: "Xyvision",
		8148: "Y-cam",
		1126: "Y.D.K",
		693:  "YAFO",
		9354: "YAMABISHI",
		2696: "YDT",
		3599: "YEC",
		1452: "YEM",
		1308: "YESTECHNOLOGY",
		8709: "YF",
		7486: "YIK",
		3634: "YMC",
		7684: "YODO",
		2106: "YOKO",
		1371: "YOZAN",
		7929: "YSI",
		8233: "YST",
		7451: "YSTen",
		6767: "YUKAI",
		8702: "YWire",
		727:  "Yamaha",
		5190: "Yamashita",
		4170: "Yamatake-Honeywell",
		8672: "Yandex",
		2783: "Yangjae",
		7724: "Yaojin",
		9250: "Yaptv",
		1697: "Yasing",
		2273: "Yazaki",
		9401: "Yeelink",
		3655: "Yi-Qing",
		7326: "Yichen",
		3382: "Yiguang",
		942:  "Yipee",
		6687: "Ynomia",
		8936: "Yo",
		1335: "Yoda",
		8966: "Yoga",
		4909: "Yoisys",
		7606: "Yokota",
		5918: "Yoshiki",
		2517: "Yoshimiya",
		3246: "Yosin",
		842:  "Yotta",
		663:  "YottaYotta",
		9336: "Yottabyte",
		1074: "Young",
		2001: "Youngbo",
		7807: "Youngkook",
		4338: "Yournet",
		9244: "Ypsomed",
		6833: "Yuduan",
		1889: "Yuehua",
		2621: "Yulinet",
		3097: "Yulong",
		3607: "Yupiteru",
		1095: "Yuxing",
		3075: "Yves",
		7505: "Yytek",
		5389: "Z-CoM",
		3446: "Z-Com",
		8716: "Z-TEC",
		7533: "Z-meta",
		7150: "Z3",
		5589: "ZAC",
		567:  "ZACCESS",
		3891: "ZANTAZ",
		974:  "ZARDCOM",
		4263: "ZAX",
		3959: "ZEFATEK",
		9225: "ZETLAB",
		2878: "ZHIYUAN",
		7114: "ZIV",
		6283: "ZNV",
		2478: "ZNYX",
		1285: "ZOOM",
		4488: "ZORT",
		9395: "ZOYI",
		3838: "ZP",
		9166: "ZPE",
		1572: "ZPSYS",
		3032: "ZTE",
		9435: "ZVISION",
		6135: "ZX",
		5026: "Zaffire",
		4371: "Zala",
		6797: "Zalliant",
		1573: "Zambeel",
		7397: "Zaplox",
		7705: "Zappware",
		3741: "Zavio",
		2700: "Zcomax",
		5948: "Zcomm",
		5612: "Zeal",
		768:  "Zebra",
		3207: "Zed-3",
		3839: "ZeeVee",
		9068: "Zeebo",
		4471: "Zeeport",
		7843: "Zegna-Daidong",
		7507: "Zektor",
		3574: "Zelax",
		7404: "Zelfy",
		8675: "Zencheer",
		9419: "Zengge",
		1800: "Zenith",
		3232: "Zenitron",
		6274: "Zenner",
		843:  "Zenocom",
		6616: "Zenotech",
		6435: "Zenovia",
		3978: "Zensys",
		8305: "Zentan",
		4860: "Zenterio",
		7281: "Zentri",
		3100: "Zenway",
		7287: "Zeo",
		6996: "Zeppelin",
		9081: "Zera",
		8878: "Zero1.tv",
		6475: "ZeroDesktop",
		4196: "Zeta",
		913:  "Zetari",
		2900: "Zetec",
		499:  "Zetes",
		1613: "Zetron",
		2648: "Zetta",
		9290: "ZettaHash",
		3447: "Zeugma",
		3498: "Zeutschel",
		9022: "Zexelon",
		8968: "Zhehua",
		197:  "Zhone",
		8722: "ZhongMiao",
		7671: "Zhuolian",
		1063: "Zi",
		5485: "Ziatech",
		8176: "Zicon",
		2422: "Zida",
		4784: "ZillionTV",
		5596: "Zilog",
		7738: "Zimi",
		8978: "Zimory",
		3213: "Zinwave",
		841:  "Zinwell",
		2128: "Zioncom",
		2175: "Zipher",
		3355: "Zippy",
		8092: "Zitte",
		6248: "Zmodo",
		3333: "Zodianet",
		3773: "Zoltan",
		780:  "Zoltrix",
		7633: "Zonar",
		5203: "Zonet",
		5994: "Zoneworx",
		8477: "Zonoff",
		6688: "Zoovel",
		2105: "Zoran",
		7012: "Ztron",
		1736: "Zultys",
		333:  "Zuma",
		7450: "ZuniData",
		7870: "Zurn",
		9397: "ZyCast",
		1527: "ZyFLEX",
		438:  "ZyGate",
		2693: "ZyXEL",
		7655: "Zycoo",
		3822: "Zygo",
		5779: "Zykronix",
		3552: "Zylaya",
		4524: "Zylin",
		5590: "Zypcom",
		8574: "Zyptonite",
		7128: "aFUN",
		6877: "acromate",
		2052: "activ-net",
		3344: "ads-tec",
		8480: "aizo",
		1516: "almedio",
		4788: "altek",
		6950: "amazipoint",
		9206: "amnimo",
		8810: "ants",
		8787: "arkona",
		7781: "as",
		392:  "ask-technologies.com",
		3472: "asteel",
		6524: "automationNEXT",
		3462: "avinfo",
		7980: "azeti",
		7678: "b-plus",
		8500: "ball-b",
		6510: "base",
		6853: "bct",
		6698: "bebro",
		9063: "beroNet",
		4470: "beyerdynamic",
		2910: "bio-logic",
		6834: "bioMerieux",
		2125: "bitWallet",
		9385: "blackned",
		8412: "bluesky",
		1938: "boca",
		1631: "bplan",
		4301: "byd:sign",
		7906: "c-scape",
		4899: "cTrixs",
		6249: "camtron",
		1723: "cd3o",
		8111: "cellon",
		7920: "chaowifi.com",
		2146: "cim-usa",
		7784: "citygrow",
		6569: "currentoptronics",
		3573: "cyber-blue",
		651:  "cyberPIXIE",
		1886: "cybernet",
		8130: "d-broad",
		8745: "d2d",
		7589: "dSPACE",
		3328: "dSys",
		9126: "data-Complex",
		2914: "datacom",
		2597: "daum",
		7380: "deister",
		1637: "devolo",
		8117: "di-soric",
		4117: "digEcor",
		8555: "digitron",
		6110: "dit",
		427:  "dotRocket",
		4318: "dresden-elektronik",
		7391: "duagon",
		6142: "e-Net",
		2461: "e-SMARTCOM",
		8711: "e-Smart",
		5194: "e-Tek",
		1123: "e-Watch",
		1417: "e-generis",
		2255: "e-w/you",
		2492: "e-zy.net",
		3295: "e2v",
		3058: "e:cue",
		1156: "eCopilt",
		4011: "eCopy",
		459:  "eDevice",
		9404: "eGauge",
		9049: "eMegatech",
		3103: "eOn",
		1558: "ePipe",
		877:  "eProduction",
		7664: "eSSys",
		4410: "eSang",
		1446: "eSpace",
		7371: "eWBM",
		2528: "eWerks",
		2867: "eXS",
		8852: "eZEX",
		8099: "easynetworks",
		7183: "ecobee",
		5821: "eero",
		988:  "egnite",
		3249: "elab-experience",
		1363: "elmegt",
		3752: "emtrion",
		7653: "enLighted",
		6816: "enblink",
		9405: "endeavour",
		8401: "enimai",
		5604: "ens",
		4472: "epro",
		7729: "essensys",
		7599: "evon",
		8093: "exands",
		6818: "exlar",
		6660: "eyevis",
		1385: "fSONA",
		8419: "fenglian",
		6714: "feno",
		8035: "ffly4u",
		3583: "fitivision",
		7755: "fos4X",
		8455: "frogblue",
		7984: "ghe-ces",
		7662: "grandcentrix",
		7044: "home2net",
		6652: "i3",
		1440: "iAd",
		5230: "iBAHN",
		1054: "iCanTek",
		3200: "iCatch",
		8524: "iConnectivity",
		3705: "iControl",
		7449: "iD",
		8996: "iDevices",
		167:  "iDigm",
		5339: "iDirect",
		6462: "iDruide",
		6963: "iFORCOM",
		7086: "iKey",
		6398: "iKuai",
		1256: "iLogic",
		3266: "iMCA-GmbH",
		589:  "iMPath",
		1437: "iMaxNetworksLimited",
		3762: "iNEWiT",
		2136: "iPAC",
		3610: "iPOX",
		3749: "iPhotonix",
		658:  "iPolicy",
		9095: "iPort",
		2448: "iPulse",
		1339: "iQstor",
		2553: "iQuest",
		9137: "iRay",
		3104: "iRex",
		7299: "iRobot",
		6758: "iRule",
		9234: "iS5",
		6590: "iSonea",
		2561: "iStor",
		6070: "iSystem",
		4066: "iTAS",
		1317: "iTEC",
		4341: "iVeia",
		4497: "iWDL",
		4621: "iWOW",
		2682: "iWise",
		3460: "iXSea",
		8831: "iZotope",
		1661: "ib-mohnen",
		3010: "iba",
		2202: "icube",
		3110: "id-Confirm",
		7829: "ifm",
		7890: "iiNet",
		3189: "iiTron",
		2058: "in2",
		8481: "inMotion",
		286:  "inXtron",
		6765: "innodisk",
		8108: "innomdlelab",
		2628: "intec",
		5719: "interWAVE",
		4751: "interbro",
		3700: "intotech",
		8977: "ioBridge",
		2744: "ioIMAGE",
		7403: "iodyne",
		2065: "ionSign",
		630:  "ipDialog",
		678:  "ipUnplugged",
		1395: "ipcas",
		7482: "iryx",
		7537: "isepos",
		7144: "itelio",
		4775: "itron",
		7500: "ittim",
		6265: "iway",
		8214: "jinyoung",
		1878: "jp-embedded",
		4787: "kasercorp",
		2802: "kk-electronic",
		3693: "l-acoustics",
		6618: "lemonbeat",
		1018: "lesswire",
		8664: "lignex1",
		8937: "littleBits",
		1088: "m-u-t",
		3180: "m2c",
		2888: "mCubelogics",
		6581: "mLogic",
		7859: "mMax",
		4931: "mPHASE",
		6886: "machineQ",
		4418: "maintech",
		288:  "manroland",
		9180: "messMa",
		6731: "miControl",
		2934: "mingjong",
		6535: "mirusystems",
		5575: "mixi",
		2915: "mm-lab",
		4871: "moblic",
		8186: "modas",
		6258: "moobox",
		5593: "mps",
		8814: "myIDkey",
		6815: "myenergi",
		3203: "nFore",
		8989: "nSTREAMS",
		1708: "nStor",
		2584: "nVent",
		8203: "netKTI",
		4807: "netTALK.com",
		2431: "netplat",
		2287: "nex-G",
		7551: "nextLAP",
		6423: "noax",
		7669: "nuoxc",
		6564: "nyantec",
		7707: "o2ones",
		8650: "optilink",
		7996: "oshkosh",
		6468: "pomdevices",
		8039: "profichip",
		4842: "r2p",
		6763: "razberi",
		9227: "sTraffic",
		4597: "saxnet",
		6819: "seca",
		8504: "silergy",
		8022: "silex",
		8212: "silicom",
		8422: "smart-electronic",
		8740: "smartAC.com",
		625:  "snom",
		5055: "so-logic",
		1643: "sofrel",
		8075: "storONE",
		6910: "streamnow",
		1816: "synertronixx",
		7406: "t-mac",
		9280: "tado",
		8141: "taskit",
		8546: "tci",
		8727: "tdvine",
		3013: "technicob",
		7212: "tide",
		8251: "trivum",
		7420: "u-blox",
		2551: "u10",
		7399: "uAvionix",
		3316: "uControl",
		4161: "ubisys",
		3570: "ubtos",
		7208: "udworks",
		8181: "unGlue",
		4413: "uv-electronic",
		8033: "vArmour",
		9159: "vastriver",
		3292: "w5networks",
		6647: "wanze",
		7336: "waytotec",
		8503: "wi-daq",
		3230: "wisembed",
		2544: "woori-net",
		9091: "worldcns",
		6293: "x-fabric",
		4762: "x-star",
		9043: "x.o.ware",
		4108: "xG",
		9418: "xn",
		8783: "xvtec",
		977:  "yLez",
	},
}
