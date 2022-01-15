// generated code - do not edit
package oui

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
		"000daa": 1968, // S.A.Tehnology
		"000dad": 1969, // Dataprobe
		"000daf": 1970, // Plexus
		"000db0": 1971, // Olym-tech
		"000db2": 1972, // Ammasso
		"000db4": 1973, // Stormshield
		"000db5": 1974, // Globalsat
		"000db6": 858,  // Broadcom
		"000db8": 1975, // Schiller
		"000dbc": 2,    // Cisco
		"000dbd": 2,    // Cisco
		"000dc0": 1976, // Spagat
		"000dc1": 1977, // SafeWeb
		"000dc2": 67,   // Private
		"000dc3": 1978, // First
		"000dc4": 1979, // Emcore
		"000dc6": 1980, // DigiRose
		"000dc7": 1981, // Cosmic
		"000dc8": 1982, // AirMagnet
		"000dca": 1983, // Tait
		"000dcb": 1984, // Petcomkorea
		"000dcc": 1985, // NEOSMART
		"000dce": 1986, // Dynavac
		"000dcf": 1987, // Cidra
		"000dd1": 1988, // Stryker
		"000dd4": 1989, // Symantec
		"000dd5": 1990, // O'Rite
		"000dd6": 1991, // ITI
		"000dd7": 1992, // Bright
		"000dd8": 1993, // BBN
		"000ddb": 1994, // Airwave
		"000ddc": 1995, // VAC
		"000dde": 1996, // Joyteck
		"000de0": 1997, // ICPDAS
		"000de1": 1998, // Control
		"000de4": 1999, // DIGINICS
		"000de6": 2000, // Youngbo
		"000de8": 2001, // Nasaco
		"000deb": 2002, // CompXs
		"000dec": 2,    // Cisco
		"000ded": 2,    // Cisco
		"000df0": 2003, // Qcom
		"000df1": 2004, // Ionix
		"000df2": 67,   // Private
		"000df3": 2005, // Asmax
		"000df4": 2006, // Watertek
		"000df5": 1967, // TElectronics
		"000df9": 2007, // NDS
		"000dfb": 2008, // Komax
		"000dfc": 2009, // ITFOR
		"000e00": 2010, // Atrie
		"000e01": 2011, // ASIP
		"000e03": 129,  // Emulex
		"000e04": 2012, // CMA/Microdialysis
		"000e07": 101,  // Sony
		"000e08": 1783, // Linksys
		"000e0b": 2013, // Netac
		"000e0c": 421,  // Intel
		"000e0f": 2014, // Ermme
		"000e10": 2015, // C-guys
		"000e13": 2016, // Accu-Sort
		"000e14": 2017, // Visionary
		"000e15": 2018, // Tadlys
		"000e17": 67,   // Private
		"000e18": 2019, // MyA
		"000e19": 2020, // LogicaCMG
		"000e1a": 2021, // JPS
		"000e1b": 2022, // IAV
		"000e1c": 2023, // Hach
		"000e1d": 2024, // ARION
		"000e1e": 2025, // QLogic
		"000e22": 67,   // Private
		"000e23": 2026, // Incipient
		"000e24": 2027, // Huwell
		"000e25": 2028, // Hannae
		"000e26": 2029, // Gincom
		"000e27": 2030, // Crere
		"000e29": 2031, // Shester
		"000e2a": 67,   // Private
		"000e2b": 2032, // Safari
		"000e2c": 2033, // Netcodec
		"000e2e": 115,  // Edimax
		"000e30": 2034, // AERAS
		"000e33": 2035, // Shuko
		"000e35": 421,  // Intel
		"000e36": 2036, // HEINESYS
		"000e38": 2,    // Cisco
		"000e39": 2,    // Cisco
		"000e3b": 2037, // Hawking
		"000e3c": 2038, // Transact
		"000e3f": 864,  // Soronti
		"000e40": 76,   // Nortel
		"000e43": 2039, // G-Tek
		"000e4c": 2040, // Bermai
		"000e4d": 2041, // Numesa
		"000e4e": 2042, // Waveplus
		"000e4f": 2043, // Trajet
		"000e52": 2044, // Optium
		"000e53": 2045, // AV
		"000e55": 2046, // Auvitran
		"000e58": 2047, // Sonos
		"000e59": 2048, // Sagemcom
		"000e5a": 2049, // TELEFIELD
		"000e5c": 125,  // ARRIS
		"000e5e": 2050, // Raisecom
		"000e5f": 2051, // activ-net
		"000e61": 2052, // Microtrol
		"000e62": 76,   // Nortel
		"000e64": 2053, // Elphel
		"000e65": 2054, // TransCore
		"000e6a": 160,  // 3Com
		"000e6b": 2055, // Janitza
		"000e6d": 2056, // Murata
		"000e70": 2057, // in2
		"000e72": 2058, // CTS
		"000e73": 2059, // Tpack
		"000e77": 2060, // Decru
		"000e78": 2061, // Amtelco
		"000e79": 2062, // Ample
		"000e7a": 2063, // GemWon
		"000e7b": 35,   // Toshiba
		"000e7e": 2064, // ionSign
		"000e7f": 302,  // HP
		"000e80": 1142, // Thomson
		"000e81": 2065, // Devicescape
		"000e82": 2066, // Infinity
		"000e83": 2,    // Cisco
		"000e84": 2,    // Cisco
		"000e85": 2067, // Catalyst
		"000e86": 2068, // Alcatel
		"000e88": 2069, // Videotron
		"000e89": 2070, // Clematic
		"000e8a": 2071, // Avara
		"000e8b": 2072, // Astarte
		"000e8c": 300,  // Siemens
		"000e8e": 2073, // SparkLAN
		"000e8f": 2074, // Sercomm
		"000e90": 2075, // Ponico
		"000e94": 2076, // Maas
		"000e97": 2077, // Ultracker
		"000e9a": 2078, // BOE
		"000e9c": 2079, // Benchmark
		"000e9e": 2080, // Topfield
		"000ea0": 2081, // NetKlass
		"000ea2": 2082, // McAfee
		"000ea4": 2083, // Quantum
		"000ea5": 2084, // BLIP
		"000ea6": 1806, // ASUS
		"000ea7": 2085, // Endace
		"000eaa": 2086, // Scalent
		"000eab": 68,   // Cray
		"000eac": 2087, // Mintron
		"000ead": 2088, // Metanoia
		"000eae": 2089, // Gawell
		"000eaf": 2090, // Castel
		"000eb1": 2091, // Newcotech
		"000eb3": 302,  // HP
		"000eb5": 2092, // Ecastle
		"000eb6": 2093, // Riverbed
		"000eb7": 2094, // Knovative
		"000eb8": 2095, // Iiga
		"000eb9": 2096, // HASHIMOTO
		"000ebb": 2097, // Everbee
		"000ebd": 2098, // Burdick
		"000ebe": 2099, // B&B
		"000ebf": 2100, // Remsdaq
		"000ec0": 76,   // Nortel
		"000ec1": 2101, // MYNAH
		"000ec2": 2102, // Lowrance
		"000ec6": 2103, // Asix
		"000ec8": 2104, // Zoran
		"000ec9": 2105, // YOKO
		"000eca": 2106, // WTSS
		"000ecb": 2107, // VineSys
		"000ecc": 2108, // Tableau
		"000ecd": 2109, // Skov
		"000ed0": 2110, // Privaris
		"000ed2": 2111, // Filtronic
		"000ed3": 2112, // Epicenter
		"000ed5": 2113, // CoPAN
		"000ed6": 2,    // Cisco
		"000ed7": 2,    // Cisco
		"000ed9": 2114, // Aksys
		"000edb": 2115, // XiNCOM
		"000edc": 2116, // Tellion
		"000edd": 2117, // Shure
		"000ede": 2118, // REMEC
		"000edf": 2119, // PLX
		"000ee0": 2120, // Mcharge
		"000ee1": 2121, // ExtremeSpeed
		"000ee2": 2122, // Custom
		"000ee3": 2123, // Chiyu
		"000ee4": 2078, // BOE
		"000ee5": 2124, // bitWallet
		"000ee6": 2125, // Adimos
		"000ee7": 2126, // AAC
		"000ee8": 2127, // Zioncom
		"000eeb": 2128, // SandmartinElectronics
		"000eec": 2129, // Orban
		"000eed": 457,  // Nokia
		"000eef": 67,   // Private
		"000ef0": 2130, // Festo
		"000ef1": 2131, // Ezquest
		"000ef2": 2132, // Infinico
		"000ef3": 2133, // Smartlabs
		"000ef4": 2134, // Kasda
		"000ef5": 2135, // iPAC
		"000efa": 2136, // Optoway
		"000efb": 2137, // Macey
		"000efd": 2138, // Fujinon
		"000efe": 2139, // EndRun
		"000eff": 2140, // Megasolution
		"000f00": 2141, // Legra
		"000f01": 2142, // Digitalks
		"000f02": 2143, // Digicube
		"000f03": 2144, // CoM&C
		"000f04": 2145, // cim-usa
		"000f06": 76,   // Nortel
		"000f07": 2146, // Mangrove
		"000f08": 2147, // Indagon
		"000f09": 67,   // Private
		"000f0b": 2148, // Kentima
		"000f0c": 2149, // Synchronic
		"000f0d": 2150, // Hunt
		"000f0e": 2151, // WaveSplitter
		"000f10": 2152, // RDM
		"000f12": 2153, // Panasonic
		"000f13": 2154, // Nisca
		"000f14": 2155, // Mindray
		"000f15": 2156, // Icotera
		"000f1b": 2157, // Ego
		"000f1f": 102,  // Dell
		"000f20": 302,  // HP
		"000f22": 2158, // Helius
		"000f23": 2,    // Cisco
		"000f24": 2,    // Cisco
		"000f26": 2159, // WorldAccxx
		"000f27": 2160, // TEAL
		"000f28": 2161, // Itronix
		"000f29": 2162, // Augmentix
		"000f2a": 2163, // Cableware
		"000f2b": 2164, // Greenbell
		"000f2c": 2165, // Uplogix
		"000f2e": 2166, // Megapower
		"000f2f": 2167, // W-Linx
		"000f33": 2168, // DUALi
		"000f34": 2,    // Cisco
		"000f35": 2,    // Cisco
		"000f37": 2169, // Xambala
		"000f38": 2170, // Netstar
		"000f3a": 2171, // Hisharp
		"000f3c": 2172, // Endeleo
		"000f3d": 803,  // D-Link
		"000f3e": 2173, // CardioNet
		"000f41": 2174, // Zipher
		"000f42": 2175, // Xalyo
		"000f43": 2176, // Wasabi
		"000f44": 2177, // Tivella
		"000f45": 2178, // Stretch
		"000f46": 2179, // Sinar
		"000f47": 2180, // Robox
		"000f48": 2181, // Polypix
		"000f49": 2182, // Northover
		"000f4a": 2183, // Kyushu-kyohan
		"000f4b": 11,   // Oracle
		"000f4c": 2184, // Elextech
		"000f4d": 2185, // TalkSwitch
		"000f4e": 2186, // Cellink
		"000f50": 2187, // StreamScale
		"000f51": 2188, // Azul
		"000f53": 1524, // Solarflare
		"000f54": 2189, // Entrelogic
		"000f55": 2190, // Datawire
		"000f57": 2191, // CABLELOGIC
		"000f58": 2192, // Adder
		"000f59": 2193, // Phonak
		"000f5a": 2194, // Peribit
		"000f5d": 2195, // Genexis
		"000f5e": 2196, // Veo
		"000f5f": 2197, // Nicety
		"000f60": 2198, // Lifetron
		"000f61": 302,  // HP
		"000f63": 2199, // Obzerv
		"000f64": 2200, // D&R
		"000f65": 2201, // icube
		"000f66": 1783, // Linksys
		"000f6a": 76,   // Nortel
		"000f6b": 2202, // GateWare
		"000f6c": 2203, // ADDI-DATA
		"000f6d": 2204, // Midas
		"000f6e": 2205, // BBox
		"000f6f": 2206, // FTA
		"000f70": 2207, // Wintec
		"000f71": 2208, // Sanmei
		"000f72": 2209, // Sandburst
		"000f74": 2210, // Qamcom
		"000f77": 2211, // Dentum
		"000f78": 2212, // Datacap
		"000f7c": 2213, // ACTi
		"000f7d": 2214, // Xirrus
		"000f7e": 2215, // Ablerex
		"000f7f": 2216, // UBSTORAGE
		"000f83": 2217, // Brainium
		"000f84": 2218, // Astute
		"000f85": 2219, // ADDO-Japan
		"000f86": 2220, // BlackBerry
		"000f87": 2221, // Maxcess
		"000f88": 2222, // AMETEK
		"000f8a": 2223, // WideView
		"000f8c": 2224, // Gigawavetech
		"000f8f": 2,    // Cisco
		"000f90": 2,    // Cisco
		"000f91": 2225, // Aerotelecom
		"000f92": 2226, // Microhard
		"000f93": 2227, // Landis+Gyr
		"000f94": 2195, // Genexis
		"000f96": 18,   // Telco
		"000f97": 2228, // Avanex
		"000f98": 2229, // Avamax
		"000f9a": 2230, // Synchrony
		"000f9c": 2231, // Panduit
		"000f9d": 2232, // DisplayLink
		"000f9e": 2233, // Murrelektronik
		"000f9f": 125,  // ARRIS
		"000fa1": 2234, // Gigabit
		"000fa2": 2235, // 2xWireless
		"000fa3": 2236, // Alpha
		"000fa5": 2237, // BWA
		"000fa7": 2238, // Raptor
		"000fa8": 2239, // Photometrics
		"000faa": 2240, // Nexus
		"000fab": 2241, // Kyushu
		"000fad": 2242, // FMN
		"000fae": 2243, // E2O
		"000faf": 2244, // Dialog
		"000fb0": 358,  // Compal
		"000fb1": 2245, // Cognio
		"000fb3": 2246, // Actiontec
		"000fb4": 2247, // Timespace
		"000fb5": 1368, // Netgear
		"000fb6": 2248, // Europlex
		"000fb7": 2249, // Cavium
		"000fb8": 2250, // CallURL
		"000fba": 2251, // Tevebox
		"000fbc": 2252, // Onkey
		"000fbd": 2253, // MRV
		"000fbe": 2254, // e-w/you
		"000fc0": 2255, // DELCOMp
		"000fc1": 2256, // WAVE
		"000fc2": 2257, // Uniwell
		"000fc3": 2258, // PalmPalm
		"000fc4": 2259, // NST
		"000fc5": 2260, // KeyMed
		"000fc6": 2261, // Eurocom
		"000fc8": 2262, // Chantry
		"000fc9": 2263, // Allnet
		"000fcb": 160,  // 3Com
		"000fcc": 125,  // ARRIS
		"000fcd": 76,   // Nortel
		"000fce": 2264, // Kikusui
		"000fcf": 2265, // DataWind
		"000fd0": 2266, // Astri
		"000fd2": 2267, // EWA
		"000fd3": 2268, // Digium
		"000fd4": 2269, // Soundcraft
		"000fd6": 2270, // Sarotech
		"000fd8": 2271, // Force
		"000fda": 2272, // Yazaki
		"000fdb": 2273, // Westell
		"000fdd": 2274, // Sordin
		"000fde": 101,  // Sony
		"000fdf": 2275, // SOLOMON
		"000fe0": 2276, // NComputing
		"000fe4": 2277, // Pantech
		"000fe6": 2278, // MBTech
		"000fe7": 2279, // Lutron
		"000fe8": 2280, // Lobos
		"000fe9": 2281, // GW
		"000fea": 1929, // Giga-Byte
		"000fec": 2282, // ARKUS
		"000fed": 2283, // Anam
		"000fee": 2284, // XTecorporated
		"000ff0": 2285, // Sunray
		"000ff1": 2286, // nex-G
		"000ff2": 2287, // Loud
		"000ff5": 2288, // GN&S
		"000ff7": 2,    // Cisco
		"000ff8": 2,    // Cisco
		"000ff9": 2289, // Valcretec
		"000ffa": 2290, // Optinel
		"000ffe": 2291, // G-PRO
		"000fff": 2292, // Control4
		"001001": 1033, // Citel
		"001002": 2293, // Actia
		"001003": 2294, // Imatron
		"001007": 2,    // Cisco
		"001008": 2295, // Vienna
		"001009": 2296, // Horanet
		"00100b": 2,    // Cisco
		"00100c": 2297, // ITO
		"00100d": 2,    // Cisco
		"001010": 2298, // Initio
		"001011": 2,    // Cisco
		"001014": 2,    // Cisco
		"001015": 2299, // OOmon
		"001016": 2300, // T.Sqware
		"001018": 858,  // Broadcom
		"00101a": 2301, // PictureTel
		"00101b": 489,  // Cornet
		"00101d": 2302, // Winbond
		"00101e": 148,  // Matsushita
		"00101f": 2,    // Cisco
		"001021": 2303, // Encanto
		"001025": 2304, // Grayhill
		"001026": 2305, // Accelerated
		"001029": 2,    // Cisco
		"00102c": 2306, // Lasat
		"00102d": 89,   // Hitachi
		"00102f": 2,    // Cisco
		"001030": 2307, // EION
		"001031": 2308, // Objective
		"001032": 2309, // Alta
		"001033": 2310, // Accesslan
		"001034": 2311, // GNP
		"001035": 1115, // Elitegroup
		"001037": 2312, // CYQ've
		"001038": 43,   // Micro
		"001039": 2313, // Vectron
		"00103d": 2314, // Phasecom
		"00103e": 2315, // Netschools
		"00103f": 2316, // Tollgrade
		"001040": 2317, // Intermec
		"001042": 2318, // Alacritech
		"001043": 2319, // A2
		"001044": 2320, // InnoLabs
		"001045": 76,   // Nortel
		"001049": 2321, // ShoreTel
		"00104a": 2322, // Parvus
		"00104b": 160,  // 3Com
		"00104d": 2323, // Surtec
		"00104e": 2324, // Ceologic
		"00104f": 11,   // Oracle
		"001050": 2325, // Rion
		"001051": 2326, // Cmicro
		"001052": 2327, // Mettler-Toledo
		"001053": 2328, // Computer
		"001054": 2,    // Cisco
		"001056": 2329, // Sodick
		"001057": 2330, // Rebel.com
		"001058": 2331, // ArrowPoint
		"001059": 2332, // Diablo
		"00105a": 160,  // 3Com
		"00105e": 2333, // Spirent
		"001060": 2334, // Billionton
		"001061": 2335, // Hostlink
		"001064": 2336, // Dnpg
		"001065": 2337, // Radyne
		"001067": 306,  // Ericsson
		"001069": 2338, // Helioss
		"00106b": 2339, // Sonus
		"00106c": 2340, // EDNT
		"00106f": 2341, // Trenton
		"001071": 2342, // Advanet
		"001072": 2343, // GVN
		"001073": 2344, // Technobox
		"001074": 2345, // Aten
		"001075": 2346, // Segate
		"001076": 2347, // EUREM
		"001078": 2348, // Nuera
		"001079": 2,    // Cisco
		"00107a": 2349, // AmbiCom
		"00107b": 2,    // Cisco
		"00107c": 2350, // P-CoM
		"00107d": 1119, // Aurora
		"00107e": 2351, // BACHMANN
		"00107f": 2352, // Crestron
		"001080": 2353, // Metawave
		"001081": 2354, // DPS
		"001083": 302,  // HP
		"001084": 2355, // K-BOT
		"001085": 1213, // Polaris
		"001086": 2356, // ATTO
		"001087": 2357, // Xstreamis
		"001088": 2358, // American
		"001089": 2359, // WebSonic
		"00108a": 2360, // TeraLogic
		"00108c": 4,    // Fujitsu
		"00108f": 2238, // Raptor
		"001090": 2361, // Cimetrics
		"001092": 2362, // Netcore
		"001093": 2363, // CMS
		"001095": 1142, // Thomson
		"001096": 2364, // Tracewell
		"001098": 2365, // Starnet
		"001099": 2366, // InnoMedia
		"00109a": 2367, // Netline
		"00109b": 129,  // Emulex
		"00109c": 2368, // M-System
		"00109d": 2369, // Clarinet
		"00109e": 2370, // Aware
		"00109f": 2371, // Pavo
		"0010a0": 2372, // Innovex
		"0010a2": 2373, // TNS
		"0010a3": 2374, // Omnitronix
		"0010a4": 788,  // Xircom
		"0010a6": 2,    // Cisco
		"0010a7": 256,  // Unex
		"0010a8": 2375, // Reliance
		"0010a9": 2376, // Adhoc
		"0010aa": 2377, // MEDIA4
		"0010ac": 2378, // Imci
		"0010af": 2379, // TAC
		"0010b0": 2380, // Meridian
		"0010b1": 2381, // FOR-A
		"0010b4": 2382, // Atmosphere
		"0010b5": 145,  // Accton
		"0010b6": 2383, // Entrata
		"0010b7": 2384, // Coyote
		"0010b9": 2385, // Maxtor
		"0010ba": 2386, // Martinho-Davis
		"0010be": 2387, // March
		"0010c0": 2388, // ARMA
		"0010c2": 2389, // Willnet
		"0010c3": 2390, // CSI-Control
		"0010c5": 2391, // Protocol
		"0010c8": 2392, // Communications
		"0010c9": 2393, // Mitsubishi
		"0010ca": 18,   // Telco
		"0010ce": 2394, // Volamp
		"0010cf": 2395, // Fiberlane
		"0010d0": 2396, // Witcom
		"0010d3": 2397, // Grips
		"0010d4": 2398, // Storage
		"0010d6": 2399, // Exelis
		"0010d7": 2400, // Argosy
		"0010d8": 2401, // Calista
		"0010da": 2402, // Kollmorgen
		"0010db": 826,  // Juniper
		"0010dc": 1812, // Micro-Star
		"0010df": 2403, // Rise
		"0010e0": 11,   // Oracle
		"0010e1": 2404, // S.I
		"0010e2": 2405, // ArrayComm
		"0010e3": 302,  // HP
		"0010e4": 2406, // NSI
		"0010e7": 700,  // Breezecom
		"0010e8": 2407, // Telocityorporated
		"0010e9": 2408, // Raidtec
		"0010ea": 1179, // Adept
		"0010eb": 2409, // Selsius
		"0010ec": 2410, // Rpcg
		"0010ed": 2411, // Sundance
		"0010ee": 2412, // CTI
		"0010ef": 2413, // Dbtel
		"0010f1": 2414, // I-O
		"0010f2": 2415, // Antec
		"0010f3": 2416, // Nexcom
		"0010f4": 2417, // Vertical
		"0010f5": 2418, // Amherst
		"0010f6": 2,    // Cisco
		"0010f7": 2419, // IRIICHI
		"0010f8": 313,  // Texio
		"0010f9": 2420, // Unique
		"0010fa": 545,  // Apple
		"0010fb": 2421, // Zida
		"0010fc": 2422, // Broadband
		"0010fd": 2423, // Cocom
		"0010ff": 2,    // Cisco
		"001100": 56,   // Schneider
		"001101": 2424, // CET
		"001104": 2425, // Telexy
		"001105": 2426, // Sunplus
		"001106": 300,  // Siemens
		"001107": 2427, // RGB
		"001109": 1812, // Micro-Star
		"00110a": 302,  // HP
		"00110b": 2428, // Franklin
		"00110d": 2429, // SANBlaze
		"00110f": 2430, // netplat
		"001110": 2431, // Maxanna
		"001111": 421,  // Intel
		"001114": 2432, // EverFocus
		"001115": 2433, // EPIN
		"001117": 2434, // Cesnet
		"001119": 2435, // Solteras
		"00111a": 125,  // ARRIS
		"00111c": 2436, // Pleora
		"00111d": 2437, // Hectrix
		"00111f": 2438, // Doremi
		"001120": 2,    // Cisco
		"001121": 2,    // Cisco
		"001122": 2439, // CIMSYS
		"001123": 2440, // Appointech
		"001124": 545,  // Apple
		"001125": 372,  // IBM
		"001126": 2441, // Venstar
		"001127": 2442, // TASI
		"001128": 2443, // Streamit
		"00112a": 2444, // Niko
		"00112b": 2445, // NetModule
		"00112c": 2446, // IZT
		"00112d": 2447, // iPulse
		"00112e": 2448, // Ceicom
		"00112f": 1806, // ASUS
		"001131": 2449, // Unatech
		"001132": 2450, // Synology
		"001133": 300,  // Siemens
		"001134": 2451, // MediaCell
		"001135": 2452, // Grandeye
		"001138": 2453, // Taishin
		"00113a": 2454, // Shinboram
		"00113b": 2455, // Micronet
		"00113c": 2456, // Micronas
		"00113e": 2457, // JL
		"001140": 2458, // Nanometrics
		"001141": 2459, // GoodMan
		"001142": 2460, // e-SMARTCOM
		"001143": 102,  // Dell
		"001144": 2461, // Assurance
		"001145": 2462, // ValuePoint
		"001146": 2463, // Telecard-Pribor
		"001147": 2464, // Secom-Industry
		"001149": 2465, // Proliphix
		"00114a": 2466, // KAYABA
		"00114b": 2467, // Francotyp-Postalia
		"001150": 2468, // Belkin
		"001151": 2469, // Mykotronx
		"001152": 2470, // Eidsvoll
		"001154": 2471, // Webpro
		"001155": 2472, // Sevis
		"001158": 76,   // Nortel
		"001159": 2473, // Matisse
		"00115b": 1115, // Elitegroup
		"00115c": 2,    // Cisco
		"00115d": 2,    // Cisco
		"001160": 2474, // ARTDIO
		"001161": 2475, // NetStreams
		"001164": 2476, // ACARD
		"001165": 2477, // ZNYX
		"001166": 2478, // Taelim
		"001168": 2479, // HomeLogic
		"00116a": 2480, // Domo
		"00116e": 2481, // Peplink
		"00116f": 2482, // Netforyou
		"001171": 2483, // DEXTER
		"001172": 2484, // Cotron
		"001174": 2485, // Mojo
		"001175": 421,  // Intel
		"001176": 2486, // Intellambda
		"001177": 2487, // Coaxial
		"001178": 2488, // Chiron
		"001179": 2489, // Singular
		"00117a": 2490, // Singim
		"00117c": 2491, // e-zy.net
		"00117e": 2492, // Midmark
		"001180": 125,  // ARRIS
		"001181": 2493, // InterEnergy
		"001185": 302,  // HP
		"001186": 1040, // Prime
		"001187": 2494, // Category
		"001188": 312,  // Enterasys
		"001189": 2495, // Aerotech
		"00118a": 2496, // Viewtran
		"001192": 2,    // Cisco
		"001193": 2,    // Cisco
		"001195": 803,  // D-Link
		"001196": 2497, // Actuality
		"001197": 2498, // Monitoring
		"001199": 2499, // 2wcom
		"00119b": 2500, // Telesynergy
		"00119d": 2501, // Diginfo
		"00119f": 457,  // Nokia
		"0011a0": 2502, // Vtech
		"0011a3": 2503, // LanReady
		"0011a4": 2504, // JStream
		"0011a5": 2505, // Fortuna
		"0011a6": 2506, // Sypixx
		"0011a8": 2507, // Quest
		"0011a9": 2508, // MOIMSTONE
		"0011aa": 2509, // Uniclass
		"0011ab": 2510, // Trustable
		"0011ac": 2511, // Simtec
		"0011ae": 125,  // ARRIS
		"0011af": 2512, // Medialink-i
		"0011b0": 2513, // Fortelink
		"0011b1": 2514, // BlueExpert
		"0011b2": 2515, // 2001
		"0011b3": 2516, // Yoshimiya
		"0011b6": 1879, // Open
		"0011ba": 2517, // Elexol
		"0011bb": 2,    // Cisco
		"0011bc": 2,    // Cisco
		"0011c0": 2518, // Aday
		"0011c5": 2519, // TEN
		"0011c6": 732,  // Seagate
		"0011c8": 2520, // Powercom
		"0011c9": 2521, // MTT
		"0011cb": 2522, // Jacobsons
		"0011cd": 2523, // Axsun
		"0011ce": 2524, // Ubisense
		"0011d4": 2525, // NetEnrich
		"0011d6": 2526, // HandEra
		"0011d7": 2527, // eWerks
		"0011d8": 1806, // ASUS
		"0011d9": 2528, // TiVo
		"0011da": 2529, // Vivaas
		"0011db": 2530, // Land-Cellular
		"0011de": 2531, // Eurilogic
		"0011e0": 2532, // U-MEDIA
		"0011e3": 1142, // Thomson
		"0011e4": 2533, // Danelec
		"0011e5": 2534, // KCodes
		"0011e8": 2535, // Tixi.com
		"0011e9": 2536, // Starnex
		"0011ea": 2537, // IWICS
		"0011ec": 2538, // Avix
		"0011ee": 2539, // Estari
		"0011f0": 2540, // Wideful
		"0011f1": 2541, // QinetiQ
		"0011f3": 2542, // NeoMedia
		"0011f4": 2543, // woori-net
		"0011f5": 2544, // Askey
		"0011f8": 2545, // AIRAYA
		"0011f9": 76,   // Nortel
		"0011fa": 2546, // Rane
		"0011fb": 2547, // Heidelberg
		"0011fc": 1596, // HARTING
		"0011fd": 2548, // Korg
		"001200": 2,    // Cisco
		"001201": 2,    // Cisco
		"001203": 2549, // ActivNetworks
		"001204": 2550, // u10
		"001205": 2551, // Terrasat
		"001206": 2552, // iQuest
		"001209": 2553, // Fastrax
		"00120b": 2554, // Chinasys
		"00120c": 2555, // CE-Infosys
		"00120e": 2556, // AboCom
		"001210": 2557, // WideRay
		"001212": 2558, // PLUS
		"001213": 2559, // Metrohm
		"001215": 2560, // iStor
		"001217": 1783, // Linksys
		"001218": 2561, // ARUZE
		"00121c": 2562, // Parrot
		"00121d": 2563, // Netfabric
		"00121e": 826,  // Juniper
		"001220": 2564, // Cadco
		"001222": 2565, // Skardin
		"001223": 2566, // Pixim
		"001224": 2567, // NexQL
		"001225": 125,  // ARRIS
		"001228": 2568, // Data
		"001229": 2569, // BroadEasy
		"00122b": 2570, // Virbiage
		"00122d": 2571, // SiNett
		"001232": 2572, // LeWiz
		"001235": 2573, // Andrew
		"001236": 2574, // ConSentry
		"001238": 2575, // SetaBox
		"00123a": 2576, // Posystech
		"00123d": 2577, // GES
		"00123e": 2578, // ERUNE
		"00123f": 102,  // Dell
		"001240": 2579, // Amoi
		"001243": 2,    // Cisco
		"001244": 2,    // Cisco
		"001246": 2580, // T.O.M
		"001247": 152,  // Samsung
		"001248": 102,  // Dell
		"00124c": 2581, // BBWM
		"00124d": 2582, // Inducon
		"00124f": 2583, // nVent
		"001251": 2584, // Silink
		"001252": 2585, // Citronix
		"001253": 2586, // AudioDev
		"001254": 2587, // Spectra
		"001255": 2588, // NetEffect
		"001256": 869,  // LG
		"001257": 2589, // LeapComm
		"001259": 2590, // Thermo
		"00125a": 612,  // Microsoft
		"00125b": 2591, // Kaimei
		"00125d": 2592, // CyberNet
		"00125e": 2593, // Caen
		"00125f": 2594, // AWIND
		"001261": 2595, // Adaptix
		"001262": 457,  // Nokia
		"001264": 2596, // daum
		"001265": 2597, // Enerdyne
		"001267": 2153, // Panasonic
		"001269": 2598, // Value
		"00126a": 2599, // OPTOELECTRONICS
		"00126b": 2600, // Ascalade
		"00126f": 2601, // Rayson
		"001272": 2602, // Redux
		"001273": 2603, // Stoke
		"001275": 2604, // Sentilla
		"001277": 2605, // Korenix
		"001279": 302,  // HP
		"00127a": 2606, // Sanyu
		"00127c": 2607, // Swegon
		"00127d": 2608, // MobileAria
		"00127f": 2,    // Cisco
		"001280": 2,    // Cisco
		"001282": 2609, // Qovia
		"001283": 76,   // Nortel
		"001285": 2610, // Gizmondo
		"001286": 2611, // Endevco
		"001288": 1939, // 2Wire
		"00128a": 125,  // ARRIS
		"00128b": 2612, // Sensory
		"00128f": 2613, // Montilio
		"001292": 2614, // Griffin
		"001295": 2615, // Aiware
		"001296": 2616, // Addlogix
		"001297": 2617, // O2Micro
		"00129a": 2618, // IRT
		"00129b": 2619, // E2S
		"00129c": 2620, // Yulinet
		"00129e": 2621, // Surf
		"00129f": 2622, // RAE
		"0012a1": 2623, // BluePacket
		"0012a2": 2624, // Vita
		"0012a4": 2625, // ThingMagic
		"0012a7": 2626, // ISR
		"0012a8": 2627, // intec
		"0012a9": 160,  // 3Com
		"0012aa": 2628, // IEE
		"0012ab": 2629, // WiLife
		"0012ac": 2630, // Ontimetek
		"0012ad": 2631, // Vivavis
		"0012af": 2632, // ELPRO
		"0012b2": 2633, // Avolites
		"0012b5": 2634, // Vialta
		"0012ba": 2635, // FSI
		"0012bc": 2636, // Echolab
		"0012bd": 2637, // Avantec
		"0012be": 2638, // Astek
		"0012bf": 2639, // Arcadyan
		"0012c0": 2640, // HotLava
		"0012c2": 406,  // Apex
		"0012c4": 2641, // Viseon
		"0012c5": 2642, // V-Show
		"0012c8": 2643, // Perfect
		"0012c9": 125,  // ARRIS
		"0012cb": 34,   // CSS
		"0012cc": 2644, // Bitatek
		"0012cd": 2645, // ASEM
		"0012cf": 145,  // Accton
		"0012d0": 2646, // Gossen-Metrawatt-GmbH
		"0012d3": 2647, // Zetta
		"0012d4": 850,  // Princeton
		"0012d7": 2648, // Invento
		"0012d9": 2,    // Cisco
		"0012da": 2,    // Cisco
		"0012dc": 2649, // SunCorp
		"0012df": 2650, // Novomatic
		"0012e0": 2651, // Codan
		"0012e1": 2652, // Alliant
		"0012e2": 2653, // ALAXALA
		"0012e6": 2654, // Spectec
		"0012e9": 2655, // Abbey
		"0012ea": 2656, // Trane
		"0012eb": 2657, // PDH
		"0012ee": 101,  // Sony
		"0012ef": 2658, // OneAccess
		"0012f0": 421,  // Intel
		"0012f1": 2659, // Ifotec
		"0012f2": 90,   // Brocade
		"0012f3": 2660, // ConnectBlue
		"0012f4": 2661, // Belco
		"0012f6": 2662, // MDK
		"0012fa": 2663, // THX
		"0012fb": 152,  // Samsung
		"0012fe": 2664, // Lenovo
		"0012ff": 2665, // Lely
		"001300": 2666, // IT-Factory
		"001302": 421,  // Intel
		"001303": 2667, // GateConnect
		"001304": 2668, // Flaircomm
		"001305": 2669, // Epicom
		"001307": 2670, // Paravirtual
		"00130a": 76,   // Nortel
		"001310": 1783, // Linksys
		"001311": 125,  // ARRIS
		"001312": 2671, // Amedia
		"001314": 2672, // Asiamajor
		"001315": 101,  // Sony
		"001318": 2673, // DGSTATION
		"001319": 2,    // Cisco
		"00131a": 2,    // Cisco
		"00131c": 2674, // LiteTouch
		"00131d": 2675, // Scanvaegt
		"001320": 421,  // Intel
		"001321": 302,  // HP
		"001322": 2676, // DAQ
		"001323": 2677, // Cap
		"001324": 56,   // Schneider
		"001325": 2678, // Cortina
		"001326": 2679, // ECM
		"001329": 2680, // VSST
		"00132d": 2681, // iWise
		"00132f": 2682, // Interactek
		"001333": 2683, // BaudTec
		"001334": 2684, // Arkados
		"001338": 2685, // Fresenius-Vial
		"00133a": 2686, // VadaTech
		"00133c": 2687, // Quintron
		"00133e": 2688, // MetaSwitch
		"001342": 2689, // Vision
		"001343": 148,  // Matsushita
		"001344": 2690, // Fargo
		"001345": 1855, // Eaton
		"001346": 803,  // D-Link
		"001348": 2691, // Artila
		"001349": 2692, // ZyXEL
		"00134a": 2693, // Engim
		"00134b": 2694, // ToGoldenNet
		"00134c": 2695, // YDT
		"00134d": 2696, // Inepro
		"00134e": 2697, // Valox
		"001352": 2698, // Naztec
		"001354": 2699, // Zcomax
		"001357": 2700, // Soyal
		"001358": 2701, // Realm
		"001359": 2702, // ProTelevision
		"00135c": 2703, // OnSite
		"00135d": 2704, // NTTPC
		"00135e": 2705, // EAB/RWI/K
		"00135f": 2,    // Cisco
		"001360": 2,    // Cisco
		"001361": 2706, // Biospace
		"001363": 2707, // Verascape
		"001364": 2708, // Paradigm
		"001365": 76,   // Nortel
		"001366": 2709, // Neturity
		"001367": 2710, // Narayon
		"001368": 2711, // Saab
		"001369": 2712, // Honda
		"00136b": 2713, // E-TEC
		"00136c": 2714, // TomTom
		"00136d": 2715, // Tentaculus
		"00136e": 2716, // Techmetro
		"00136f": 2717, // PacketMotion
		"001370": 457,  // Nokia
		"001371": 125,  // ARRIS
		"001372": 102,  // Dell
		"001373": 2718, // BLwave
		"001374": 535,  // Atheros
		"001376": 2719, // Tabor
		"001377": 152,  // Samsung
		"001378": 2720, // Qsan
		"00137a": 2721, // Netvox
		"00137b": 2722, // Movon
		"00137c": 2723, // Kaicom
		"00137d": 2724, // Dynalab
		"00137e": 2725, // CorEdge
		"00137f": 2,    // Cisco
		"001380": 2,    // Cisco
		"001381": 2726, // CHIPS
		"001382": 2727, // Cetacea
		"001383": 2728, // Application
		"001385": 2729, // Add-On
		"001386": 2730, // ABB/Totalflow
		"001387": 2731, // 27M
		"00138b": 2732, // Phantom
		"00138c": 2733, // Kumyoung
		"00138d": 2734, // Kinghold
		"00138f": 1666, // Asiarock
		"001390": 2735, // Termtek
		"001391": 2736, // Ouen
		"001392": 2737, // Ruckus
		"001393": 2738, // Panta
		"001394": 2739, // Infohand
		"001395": 2740, // Congatec
		"001397": 11,   // Oracle
		"001398": 2741, // TrafficSim
		"001399": 2742, // STAC
		"00139b": 2743, // ioIMAGE
		"00139c": 2744, // Exavera
		"00139e": 2745, // Ciara
		"0013a0": 2746, // ALGOSYSTEM
		"0013a1": 2747, // Crow
		"0013a2": 2748, // MaxStream
		"0013a3": 300,  // Siemens
		"0013a4": 2749, // KeyEye
		"0013a5": 2750, // General
		"0013a6": 2751, // Extricom
		"0013a8": 2752, // Tanisys
		"0013a9": 101,  // Sony
		"0013ab": 2753, // Telemotive
		"0013ac": 2754, // Sunmyung
		"0013ad": 2755, // Sendo
		"0013ae": 2756, // Radiance
		"0013af": 2757, // NUMA
		"0013b0": 2758, // Jablotron
		"0013b2": 2759, // Carallon
		"0013b3": 2760, // Ecom
		"0013b5": 2761, // Wavesat
		"0013b8": 2762, // RyCo
		"0013b9": 2763, // BM
		"0013ba": 2764, // ReadyLinks
		"0013bb": 2765, // Smartvue
		"0013bc": 2766, // Artimi
		"0013bd": 2767, // Hymatom
		"0013c2": 2768, // WACOM
		"0013c3": 2,    // Cisco
		"0013c4": 2,    // Cisco
		"0013c6": 2769, // OpenGear
		"0013c7": 2770, // IONOS
		"0013ca": 2771, // ATX
		"0013cd": 2772, // MTI
		"0013ce": 421,  // Intel
		"0013cf": 2773, // 4Access
		"0013d3": 1812, // Micro-Star
		"0013d4": 1806, // ASUS
		"0013d5": 1586, // RuggedCom
		"0013d7": 2774, // SPIDCOM
		"0013da": 2775, // Diskware
		"0013dc": 2776, // Ibtek
		"0013de": 2777, // Adapt4
		"0013df": 2778, // Ryvor
		"0013e0": 2056, // Murata
		"0013e1": 2779, // Iprobe
		"0013e2": 2780, // GeoVision
		"0013e3": 2781, // CoVi
		"0013e4": 2782, // Yangjae
		"0013e5": 2783, // Tenosys
		"0013e6": 2784, // Technolution
		"0013e7": 2785, // Halcro
		"0013e8": 421,  // Intel
		"0013e9": 2786, // VeriWave
		"0013ea": 2787, // Kamstrup
		"0013eb": 2788, // Sysmaster
		"0013ed": 2789, // Psia
		"0013f1": 2790, // AMOD
		"0013f2": 2791, // Klas
		"0013f3": 2792, // Giga-byte
		"0013f4": 2793, // Psitek
		"0013f5": 2794, // Akimbi
		"0013f6": 2795, // Cintech
		"0013f7": 741,  // SMC
		"0013f9": 2796, // Cavera
		"0013fa": 2797, // LifeSize
		"0013fc": 2798, // SiCortex
		"0013fd": 457,  // Nokia
		"0013fe": 2799, // Grandtec
		"001401": 2800, // Rivertree
		"001402": 2801, // kk-electronic
		"001403": 2802, // Renasis
		"001404": 125,  // ARRIS
		"001405": 2803, // OpenIB
		"001406": 2804, // Go
		"001408": 2805, // Eka
		"00140a": 2806, // WEPIO
		"00140b": 1978, // First
		"00140d": 76,   // Nortel
		"00140e": 76,   // Nortel
		"001412": 2807, // S-TEC
		"001416": 2808, // Scosche
		"001418": 2809, // C4Line
		"001419": 2810, // Sidsa
		"00141a": 2811, // Deicy
		"00141b": 2,    // Cisco
		"00141c": 2,    // Cisco
		"00141f": 2812, // SunKwang
		"001422": 102,  // Dell
		"001426": 2813, // NL
		"001427": 2814, // JazzMutant
		"001428": 2815, // Vocollect
		"00142a": 1115, // Elitegroup
		"00142b": 2816, // Edata
		"00142c": 2817, // Koncept
		"00142d": 2818, // Toradex
		"00142f": 2819, // Savvius
		"001430": 2820, // ViPowER
		"001431": 2821, // PDL
		"001433": 2822, // Empower
		"001434": 2823, // Keri
		"001435": 2824, // CityCom
		"001437": 2825, // GSTeletech
		"001438": 302,  // HP
		"00143b": 2826, // Sensovation
		"00143d": 2827, // Aevoe
		"00143e": 2828, // AirLink
		"00143f": 2829, // Hotway
		"001440": 2830, // ATOMIC
		"001442": 2831, // Atto
		"001443": 2832, // Consultronics
		"001444": 2833, // Grundfos
		"001446": 2834, // SuperVision
		"001447": 2835, // BOAZ
		"00144b": 2836, // Hifn
		"00144d": 1861, // Intelligent
		"00144e": 2837, // Srisa
		"00144f": 11,   // Oracle
		"001450": 2838, // Heim
		"001451": 545,  // Apple
		"001452": 2839, // Calculex
		"001453": 1703, // Advantech
		"001454": 2840, // Symwave
		"001455": 2841, // Coder
		"001456": 2842, // Edge
		"001457": 2843, // T-Vips
		"001459": 2844, // Moram
		"00145b": 2845, // SeekerNet
		"00145d": 2846, // WJ
		"00145e": 372,  // IBM
		"00145f": 2847, // Aditec
		"001460": 2848, // Kyocera
		"001461": 2849, // Corona
		"001462": 2850, // Digiwell
		"001464": 2851, // Cryptosoft
		"001467": 2852, // ArrowSpan
		"001468": 2853, // CelPlan
		"001469": 2,    // Cisco
		"00146a": 2,    // Cisco
		"00146b": 2854, // Anagran
		"00146c": 1368, // Netgear
		"00146d": 2855, // RF
		"00146f": 2856, // Kohler
		"001470": 2857, // Prokom
		"001473": 2858, // Bookham
		"001474": 2859, // K40
		"001475": 2860, // Wiline
		"001476": 2861, // MultiCom
		"001477": 1585, // Trilliant
		"001478": 1595, // TP-Link
		"00147a": 2862, // Eubus
		"00147b": 2863, // Iteris
		"00147c": 160,  // 3Com
		"00147e": 2864, // InnerWireless
		"001481": 2865, // Multilink
		"001482": 1119, // Aurora
		"001483": 2866, // eXS
		"001484": 2867, // Cermate
		"001485": 1929, // Giga-Byte
		"001488": 2868, // Akorri
		"00148b": 2869, // Globo
		"00148f": 2870, // Protronic
		"001490": 2871, // ASP
		"001491": 2872, // Daniels
		"001492": 2873, // Liteon
		"001493": 2874, // Systimax
		"001494": 2875, // ESU
		"001495": 1939, // 2Wire
		"001496": 2876, // Phonic
		"001497": 2877, // ZHIYUAN
		"001499": 2878, // Helicomm
		"00149a": 125,  // ARRIS
		"00149b": 2879, // Nokota
		"00149c": 2880, // HF
		"00149e": 2881, // UbONE
		"0014a0": 2882, // Accsense
		"0014a1": 801,  // Synchronous
		"0014a3": 2883, // Vitelec
		"0014a5": 1450, // Gemtek
		"0014a6": 2884, // Teranetics
		"0014a7": 457,  // Nokia
		"0014a8": 2,    // Cisco
		"0014a9": 2,    // Cisco
		"0014ab": 2885, // Senhai
		"0014ae": 2886, // Wizlogics
		"0014b2": 2887, // mCubelogics
		"0014b3": 2888, // CoreStar
		"0014b5": 2889, // Physiometrix
		"0014b6": 2890, // Enswer
		"0014b8": 2891, // Hill-Rom
		"0014bd": 2892, // IncNETWORKS
		"0014be": 2893, // Wink
		"0014bf": 1783, // Linksys
		"0014c2": 302,  // HP
		"0014c3": 732,  // Seagate
		"0014c4": 2894, // Vitelcom
		"0014c5": 2895, // Alive
		"0014c6": 2896, // Quixant
		"0014c7": 76,   // Nortel
		"0014c8": 2897, // Contemporary
		"0014c9": 90,   // Brocade
		"0014cb": 2898, // LifeSync
		"0014cc": 2899, // Zetec
		"0014cd": 2900, // DigitalZone
		"0014ce": 2901, // NF
		"0014cf": 2902, // INVISIO
		"0014d0": 2903, // BTI
		"0014d1": 2904, // TRENDnet
		"0014d3": 2905, // Sepsa
		"0014d4": 2906, // K
		"0014d6": 2907, // Jeongmin
		"0014d7": 2908, // Datastore
		"0014d8": 2909, // bio-logic
		"0014dd": 2910, // Covergence
		"0014df": 2911, // HI-P
		"0014e0": 2912, // LET'S
		"0014e2": 2913, // datacom
		"0014e3": 2914, // mm-lab
		"0014e4": 2915, // Infinias
		"0014e5": 2916, // Alticast
		"0014e7": 2917, // Stolinx
		"0014e8": 125,  // ARRIS
		"0014e9": 2918, // Nortech
		"0014eb": 2919, // AwarePoint
		"0014ed": 2920, // Airak
		"0014ee": 123,  // WD
		"0014ef": 2921, // TZero
		"0014f1": 2,    // Cisco
		"0014f2": 2,    // Cisco
		"0014f3": 2922, // ViXS
		"0014f6": 826,  // Juniper
		"0014f7": 2923, // CREVIS
		"0014fb": 2924, // Technical
		"0014fc": 2925, // Extandon
		"0014fd": 2926, // Thecus
		"0014fe": 2927, // Artech
		"001500": 421,  // Intel
		"001501": 2928, // LexBox
		"001502": 2929, // BETA
		"001505": 2246, // Actiontec
		"001509": 2930, // Plus
		"00150a": 2931, // Sonoa
		"00150c": 621,  // AVM
		"00150e": 2932, // Openbrain
		"00150f": 2933, // mingjong
		"001510": 2934, // Techsphere
		"001515": 2935, // Leipold+Co.GmbH
		"001516": 2936, // Uriel
		"001517": 421,  // Intel
		"00151a": 472,  // Hunter
		"00151b": 2937, // Isilon
		"00151c": 2938, // Leneco
		"00151d": 2939, // M2I
		"001520": 2940, // Radiocrafts
		"001521": 2941, // Horoquartz
		"001523": 2942, // Meteor
		"001524": 2943, // Numatics
		"001526": 2944, // Remote
		"001529": 2945, // N3
		"00152a": 457,  // Nokia
		"00152b": 2,    // Cisco
		"00152c": 2,    // Cisco
		"00152d": 2946, // TenX
		"00152e": 2947, // PacketHop
		"00152f": 125,  // ARRIS
		"001530": 102,  // Dell
		"001531": 2948, // Kocom
		"001533": 2949, // Nadam
		"001535": 2950, // OTE
		"001536": 2951, // Powertech
		"001537": 2952, // Ventus
		"001538": 2953, // RFID
		"00153c": 2954, // Kprotech
		"001540": 76,   // Nortel
		"001541": 2955, // StrataLight
		"001544": 2956, // CoM.s.a.t
		"001545": 2957, // SEECODE
		"001547": 2958, // AiZen
		"001548": 2959, // Cube
		"00154a": 2960, // Wanshih
		"00154c": 2961, // Saunders
		"00154d": 2962, // Netronome
		"00154e": 2963, // IEC
		"001550": 2964, // Nits
		"001551": 2965, // RadioPulse
		"001552": 2966, // Wi-Gear
		"001553": 2967, // Cytyc
		"001555": 2968, // DFM
		"001556": 2048, // Sagemcom
		"001557": 45,   // Olivetti
		"001558": 221,  // Foxconn
		"001559": 2969, // Securaplane
		"00155b": 2970, // Sampo
		"00155d": 612,  // Microsoft
		"00155f": 2971, // GreenPeak
		"001560": 302,  // HP
		"001561": 2972, // JJPlus
		"001562": 2,    // Cisco
		"001563": 2,    // Cisco
		"001566": 2973, // A-First
		"001567": 2974, // RADWIN
		"001568": 2975, // Dilithium
		"00156a": 2976, // DG2L
		"00156b": 2977, // Perfisans
		"00156d": 2978, // Ubiquiti
		"00156f": 2979, // Xiranet
		"001570": 768,  // Zebra
		"001571": 2980, // Nolan
		"001572": 2981, // Red-Lemon
		"001573": 2982, // NewSoft
		"001575": 2983, // Nevis
		"00157b": 2984, // Leuze
		"00157c": 2985, // Dave
		"00157d": 2986, // Posdata
		"00157f": 2987, // ChuanG
		"001580": 2988, // U-WAY
		"001581": 2989, // MAKUS
		"001583": 2990, // IVT
		"001589": 2991, // D-MAX
		"00158a": 363,  // SURECOM
		"00158d": 2992, // Jennic
		"00158e": 2993, // Plustek.Inc
		"001590": 2994, // Hectronic
		"001591": 2995, // RLW
		"001593": 2996, // U4EA
		"001594": 2997, // Bixolon
		"001596": 125,  // ARRIS
		"001599": 152,  // Samsung
		"00159a": 125,  // ARRIS
		"00159b": 76,   // Nortel
		"00159f": 2998, // Terascala
		"0015a0": 457,  // Nokia
		"0015a1": 2999, // ECA-Sinters
		"0015a2": 125,  // ARRIS
		"0015a3": 125,  // ARRIS
		"0015a4": 125,  // ARRIS
		"0015a5": 49,   // DCI
		"0015a6": 1565, // Digital
		"0015a7": 3000, // Robatech
		"0015a8": 125,  // ARRIS
		"0015aa": 3001, // Rextechnik
		"0015ac": 3002, // Capelon
		"0015ad": 3003, // Accedian
		"0015af": 3004, // Azurewave
		"0015b0": 3005, // Autotelenet
		"0015b1": 3006, // Ambient
		"0015b2": 587,  // Advanced
		"0015b3": 3007, // Caretech
		"0015b6": 3008, // ShinMaywa
		"0015b7": 35,   // Toshiba
		"0015b8": 1024, // Tahoe
		"0015b9": 152,  // Samsung
		"0015ba": 3009, // iba
		"0015bc": 3010, // Develco
		"0015be": 3011, // Iqua
		"0015bf": 3012, // technicob
		"0015c1": 101,  // Sony
		"0015c4": 3013, // Flovel
		"0015c5": 102,  // Dell
		"0015c6": 2,    // Cisco
		"0015c7": 2,    // Cisco
		"0015c8": 3014, // FlexiPanel
		"0015c9": 3015, // Gumstix
		"0015ca": 3016, // TeraRecon
		"0015cb": 2621, // Surf
		"0015cc": 3017, // Uquest
		"0015cd": 3018, // Exartech
		"0015ce": 125,  // ARRIS
		"0015cf": 125,  // ARRIS
		"0015d0": 125,  // ARRIS
		"0015d1": 125,  // ARRIS
		"0015d2": 3019, // Xantech
		"0015d4": 3020, // Emitor
		"0015d5": 3021, // Nicevt
		"0015d7": 3022, // Reti
		"0015d8": 3023, // Interlink
		"0015d9": 3024, // PKC
		"0015db": 3025, // Canesta
		"0015dc": 3026, // KT&C
		"0015de": 457,  // Nokia
		"0015e0": 306,  // Ericsson
		"0015e1": 3027, // Picochip
		"0015e3": 3028, // Dream
		"0015e5": 3029, // Cheertek
		"0015e8": 76,   // Nortel
		"0015e9": 803,  // D-Link
		"0015ea": 3030, // Tellumat
		"0015eb": 3031, // ZTE
		"0015f0": 3032, // EGO
		"0015f1": 3033, // KYLINK
		"0015f2": 1806, // ASUS
		"0015f3": 3034, // Peltor
		"0015f4": 3035, // Eventide
		"0015f6": 3036, // Science
		"0015f7": 3037, // Wintecronics
		"0015f8": 3038, // Kingtronics
		"0015f9": 2,    // Cisco
		"0015fa": 2,    // Cisco
		"001601": 1077, // Buffalo
		"001602": 3039, // Ceyon
		"001603": 3040, // CoOLKSKY
		"001604": 3041, // Sigpro
		"001606": 3042, // Ideal
		"001607": 3043, // Curves
		"001608": 3044, // Sequans
		"001609": 3045, // Unitech
		"00160a": 3046, // SWEEX
		"00160b": 3047, // TVWorks
		"00160e": 3048, // Optica
		"001610": 3049, // Carina
		"001612": 3050, // Otsuka
		"001613": 3051, // LibreStream
		"001615": 3052, // Nittan
		"001616": 3053, // Browan
		"001617": 3054, // MSI
		"001618": 3055, // HIVION
		"00161a": 3056, // Dametric
		"00161b": 2455, // Micronet
		"00161c": 3057, // e:cue
		"00161e": 3058, // Woojinnet
		"00161f": 3059, // SUNWAVETEC
		"001620": 101,  // Sony
		"001622": 3060, // BBH
		"001624": 3061, // Teneros
		"001625": 3062, // Impinj
		"001626": 125,  // ARRIS
		"001628": 3063, // Magicard
		"001629": 3064, // Nivus
		"00162c": 3065, // Xanboo
		"00162d": 3066, // STNet
		"00162f": 3067, // Geutebruck
		"001630": 3068, // Vativ
		"001631": 3069, // Xteam
		"001632": 152,  // Samsung
		"001634": 3070, // Mathtech
		"001635": 302,  // HP
		"001636": 3071, // Quanta
		"001637": 3072, // CITEL
		"001638": 576,  // TECOM
		"001639": 3073, // Ubiquam
		"00163a": 3074, // Yves
		"00163e": 3075, // Xensource
		"00163f": 3076, // CReTE
		"001640": 3077, // Asmobile
		"001642": 3078, // Pangolin
		"001643": 3079, // Sunhillo
		"001644": 452,  // LITE-ON
		"001646": 2,    // Cisco
		"001647": 2,    // Cisco
		"001648": 3080, // SSD
		"001649": 3081, // SetOne
		"00164a": 1381, // Vibration
		"00164e": 457,  // Nokia
		"001651": 3082, // Exeo
		"001652": 3083, // Hoatech
		"001654": 3084, // Flex-P
		"001655": 3085, // FUHO
		"001656": 1427, // Nintendo
		"001657": 3086, // Aegate
		"001658": 3087, // Fusiontech
		"00165c": 3088, // Trackflow
		"00165d": 3089, // AirDefense
		"001660": 76,   // Nortel
		"001662": 3090, // Liyuh
		"001663": 3091, // KBT
		"001664": 3092, // Prod-El
		"001666": 3093, // Quantier
		"001668": 3094, // Eishin
		"001669": 2253, // MRV
		"00166a": 3095, // TPS
		"00166b": 152,  // Samsung
		"00166c": 152,  // Samsung
		"00166d": 3096, // Yulong
		"00166e": 3097, // Arbitron
		"00166f": 421,  // Intel
		"001670": 3098, // SKNET
		"001672": 3099, // Zenway
		"001673": 3100, // Bury
		"001674": 3101, // EuroCB
		"001675": 125,  // ARRIS
		"001676": 421,  // Intel
		"001679": 3102, // eOn
		"00167c": 3103, // iRex
		"00167e": 3104, // Diboss
		"001683": 3105, // WEBIO
		"001684": 3106, // Donjin
		"001688": 3107, // ServerEngines
		"001689": 3108, // Pilkor
		"00168a": 3109, // id-Confirm
		"00168b": 3110, // Paralan
		"00168d": 3111, // KORWIN
		"00168e": 3112, // Vimicro
		"001690": 3113, // J-TEK
		"001691": 3114, // Moser-Baer
		"001692": 3115, // Scientific-Atlanta
		"001693": 3116, // PowerLink
		"001694": 3117, // Sennheiser
		"001695": 3118, // AVC
		"001696": 3119, // QDI
		"001697": 48,   // NEC
		"00169a": 3120, // Quadrics
		"00169c": 2,    // Cisco
		"00169d": 2,    // Cisco
		"00169f": 3121, // Vimtron
		"0016a0": 3122, // Auto-Maskin
		"0016a1": 3123, // 3Leaf
		"0016a2": 3124, // CentraLite
		"0016a4": 3125, // Ezurio
		"0016a8": 3126, // CWT
		"0016a9": 3127, // 2EI
		"0016aa": 3128, // Kei
		"0016ab": 3129, // Dansensor
		"0016ac": 3130, // Toho
		"0016ad": 3131, // BT-Links
		"0016ae": 1075, // Inventel
		"0016b0": 3132, // VK
		"0016b1": 3133, // KBS
		"0016b2": 3134, // DriveCam
		"0016b3": 3135, // Photonicbridges
		"0016b4": 67,   // Private
		"0016b5": 125,  // ARRIS
		"0016b6": 1783, // Linksys
		"0016b8": 101,  // Sony
		"0016ba": 3136, // Weathernews
		"0016bb": 3137, // Law-Chain
		"0016bc": 457,  // Nokia
		"0016be": 3138, // InfRANET
		"0016c0": 3139, // Semtech
		"0016c1": 3140, // Eleksen
		"0016c2": 1798, // Avtec
		"0016c3": 3141, // BA
		"0016c4": 3142, // SiRF
		"0016c7": 2,    // Cisco
		"0016c8": 2,    // Cisco
		"0016ca": 76,   // Nortel
		"0016cb": 545,  // Apple
		"0016cc": 3143, // Xcute
		"0016d2": 3144, // Caspian
		"0016d3": 1592, // Wistron
		"0016d4": 358,  // Compal
		"0016d5": 3145, // Synccom
		"0016d6": 3146, // TDA
		"0016d7": 3147, // Sunways
		"0016d8": 3148, // Senea
		"0016da": 3149, // Futronic
		"0016db": 152,  // Samsung
		"0016dc": 3150, // Archos
		"0016dd": 3151, // Gigabeam
		"0016de": 790,  // FAST
		"0016df": 3152, // Lundinova
		"0016e0": 160,  // 3Com
		"0016e1": 3153, // SiliconStor
		"0016e3": 2544, // Askey
		"0016e6": 1929, // Giga-Byte
		"0016ea": 421,  // Intel
		"0016eb": 421,  // Intel
		"0016ec": 1115, // Elitegroup
		"0016ed": 1424, // Utility
		"0016ee": 3154, // Royaldigital
		"0016f0": 102,  // Dell
		"0016f1": 3155, // OmniSense
		"0016f4": 3156, // Eidicom
		"0016f7": 3157, // L-3
		"0016f8": 3158, // Aviqtech
		"0016fc": 3159, // Tohken
		"0016fd": 3160, // Jaty
		"0016fe": 430,  // Alpsalpine
		"001700": 125,  // ARRIS
		"001701": 3161, // KDE
		"001704": 3162, // Shinco
		"001705": 3163, // Methode
		"001706": 3164, // Techfaithwireless
		"001707": 3165, // InGrid
		"001708": 302,  // HP
		"001709": 3166, // Exalt
		"00170b": 3167, // Contela
		"00170c": 3168, // Twig
		"00170d": 3169, // Dust
		"00170e": 2,    // Cisco
		"00170f": 2,    // Cisco
		"001710": 3170, // Casa
		"001712": 3171, // ISCO
		"001715": 3172, // Qstik
		"001716": 3173, // Qno
		"001718": 3174, // Vansco
		"00171a": 3175, // Winegard
		"00171d": 3176, // Digit
		"00171f": 3177, // IMV
		"001722": 3178, // Hanazeder
		"001726": 3179, // m2c
		"001728": 3180, // Selex
		"001729": 3181, // Ubicod
		"00172a": 3182, // Proware
		"00172b": 3183, // Global
		"00172e": 3184, // FXC
		"00172f": 3185, // NeuLion
		"001730": 3186, // Automation
		"001731": 1806, // ASUS
		"001733": 3187, // SFR
		"001736": 3188, // iiTron
		"00173a": 3189, // Cloudastructure
		"00173b": 2,    // Cisco
		"00173c": 185,  // Extreme
		"00173d": 3190, // Neology
		"00173e": 3191, // LeucotronEquipamentos
		"00173f": 2468, // Belkin
		"001741": 3192, // Defidev
		"001742": 4,    // Fujitsu
		"001744": 3193, // Araneo
		"001745": 3194, // INNOTZ
		"001746": 3195, // Freedom9
		"001747": 1375, // Trimble
		"00174a": 3196, // Socomec
		"00174b": 457,  // Nokia
		"00174c": 3197, // Millipore
		"00174e": 3198, // Parama-tech
		"00174f": 3199, // iCatch
		"001751": 3200, // Online
		"001752": 3201, // DAGS
		"001753": 3202, // nFore
		"001756": 1930, // Vinci
		"001757": 3203, // RIX
		"001758": 3204, // ThruVision
		"001759": 2,    // Cisco
		"00175a": 2,    // Cisco
		"00175c": 3205, // Sharp
		"00175e": 3206, // Zed-3
		"00175f": 3207, // XENOLINK
		"001761": 67,   // Private
		"001762": 3208, // Solar
		"001764": 3209, // ATMedia
		"001765": 76,   // Nortel
		"001766": 3210, // Accense
		"001767": 3211, // Earforce
		"001768": 3212, // Zinwave
		"001769": 3213, // Cymphonix
		"00176a": 3214, // Avago
		"00176b": 3215, // Kiyon
		"00176c": 3216, // Pivot3
		"00176d": 475,  // Core
		"00176f": 3217, // PAX
		"001770": 3218, // Arti
		"001771": 3219, // APD
		"001773": 3220, // Laketune
		"001774": 3221, // Elesta
		"001777": 3222, // Obsidian
		"001779": 3223, // QuickTel
		"00177b": 3224, // Azalea
		"00177d": 3225, // IDT
		"00177e": 3226, // Meshcom
		"001782": 3227, // LoBenn
		"001784": 125,  // ARRIS
		"001785": 3228, // Sparr
		"001786": 3229, // wisembed
		"001787": 3230, // Brother
		"001789": 3231, // Zenitron
		"00178a": 3232, // Darts
		"00178b": 3233, // Teledyne
		"00178d": 3234, // Checkpoint
		"001791": 3235, // LinTech
		"001793": 3236, // Tigi
		"001794": 2,    // Cisco
		"001795": 2,    // Cisco
		"001796": 3237, // Rittmeyer
		"001798": 3238, // Azonic
		"001799": 3239, // SmarTire
		"00179a": 803,  // D-Link
		"00179d": 3240, // Kelman
		"00179e": 3241, // Sirit
		"00179f": 3242, // Apricorn
		"0017a1": 3243, // 3soft
		"0017a2": 3244, // Camrivox
		"0017a4": 302,  // HP
		"0017a5": 1785, // Ralink
		"0017a6": 3245, // Yosin
		"0017a8": 3246, // EDM
		"0017a9": 3247, // Sentivision
		"0017aa": 3248, // elab-experience
		"0017ab": 1427, // Nintendo
		"0017ad": 3249, // AceNet
		"0017ae": 3250, // GAI-Tronics
		"0017af": 3251, // Enermet
		"0017b0": 457,  // Nokia
		"0017b5": 3252, // Peerless
		"0017b6": 3253, // Aquantia
		"0017b7": 3254, // Tonze
		"0017b8": 3255, // Novatron
		"0017ba": 3256, // Sedo
		"0017bb": 3257, // Syrinx
		"0017bd": 3258, // Tibetsystem
		"0017bf": 3259, // Coherent
		"0017c0": 3260, // PureTech
		"0017c3": 3261, // KTF
		"0017c5": 3262, // SonicWALL
		"0017c7": 3263, // MARA
		"0017c8": 2848, // Kyocera
		"0017c9": 152,  // Samsung
		"0017ca": 551,  // Qisda
		"0017cb": 826,  // Juniper
		"0017ce": 3264, // Screen
		"0017cf": 3265, // iMCA-GmbH
		"0017d0": 3266, // Opticom
		"0017d1": 76,   // Nortel
		"0017d2": 3267, // Thinlinx
		"0017d3": 3268, // Etymotic
		"0017d5": 152,  // Samsung
		"0017d9": 3269, // AAI
		"0017db": 3270, // Canko
		"0017df": 2,    // Cisco
		"0017e0": 2,    // Cisco
		"0017e1": 3271, // DACOS
		"0017e2": 125,  // ARRIS
		"0017ed": 3272, // WooJooIT
		"0017ee": 125,  // ARRIS
		"0017ef": 372,  // IBM
		"0017f1": 3273, // Renu
		"0017f2": 545,  // Apple
		"0017f3": 124,  // Harris
		"0017f7": 1234, // CEM
		"0017f8": 3274, // Motech
		"0017fa": 612,  // Microsoft
		"0017fb": 3275, // FA
		"0017fc": 3276, // Suprema
		"0017ff": 3277, // PLAYLINE
		"001800": 3278, // Unigrand
		"001801": 2246, // Actiontec
		"001802": 2236, // Alpha
		"001806": 3279, // Hokkei
		"001807": 3280, // Fanstel
		"001808": 3281, // SightLogix
		"001809": 3282, // Cresyn
		"00180e": 3283, // Avega
		"00180f": 457,  // Nokia
		"001811": 3284, // Neuros
		"001813": 101,  // Sony
		"001814": 3285, // Mitutoyo
		"001815": 3286, // GZ
		"001816": 3287, // Ubixon
		"001818": 2,    // Cisco
		"001819": 2,    // Cisco
		"00181c": 3288, // Exterity
		"00181d": 3289, // Asia
		"00181e": 3290, // GDX
		"00181f": 1355, // Palmmicro
		"001820": 3291, // w5networks
		"001821": 3292, // Sindoricoh
		"001823": 957,  // Delta
		"001824": 3293, // Kimaldi
		"001825": 67,   // Private
		"001828": 3294, // e2v
		"001829": 3295, // Gatsometer
		"00182b": 3296, // Softier
		"00182c": 3297, // Ascend
		"00182e": 3298, // XStreamHD
		"001836": 3299, // REJ
		"001838": 3300, // PanAccess
		"001839": 1783, // Linksys
		"00183a": 2273, // Westell
		"00183b": 3301, // CENITS
		"00183c": 3302, // Encore
		"00183e": 3303, // Digilent
		"00183f": 1939, // 2Wire
		"001841": 375,  // High
		"001842": 457,  // Nokia
		"001843": 3304, // Dawevision
		"001845": 3305, // Pulsar-Telecom
		"001847": 3249, // AceNet
		"001848": 3306, // Vecima
		"001849": 2583, // nVent
		"00184a": 3307, // Catcher
		"00184c": 3308, // Bogen
		"00184d": 1368, // Netgear
		"00184e": 3309, // Lianhe
		"001851": 3310, // SWsoft
		"001853": 3311, // Atera
		"001854": 3312, // Argard
		"001856": 3313, // EyeFi
		"001858": 3314, // TagMaster
		"00185a": 3315, // uControl
		"00185c": 3316, // EDSLAB
		"00185d": 3317, // Taiguen
		"00185e": 3318, // Nexterm
		"00185f": 2379, // TAC
		"001861": 3319, // Ooma
		"001862": 732,  // Seagate
		"001863": 3320, // Veritech
		"001864": 1855, // Eaton
		"001865": 300,  // Siemens
		"001869": 3321, // Kingjim
		"00186c": 3322, // Neonode
		"00186e": 160,  // 3Com
		"001871": 302,  // HP
		"001872": 3323, // Expertise
		"001873": 2,    // Cisco
		"001874": 2,    // Cisco
		"001876": 3324, // WowWee
		"001877": 3325, // Amplex
		"001878": 3326, // Mackware
		"001879": 3327, // dSys
		"00187a": 3328, // Wiremold
		"00187b": 3329, // 4NSYS
		"00187c": 3330, // Intercross
		"00187d": 3331, // Armorlink
		"00187f": 3332, // Zodianet
		"001881": 3333, // Buyang
		"001882": 3334, // Huawei
		"001883": 3335, // FORMOSA21
		"001885": 687,  // Motorola
		"001886": 3336, // EL-Tech
		"001887": 3337, // Metasystem
		"001889": 3338, // WinNet
		"00188a": 3339, // Infinova
		"00188b": 102,  // Dell
		"00188d": 457,  // Nokia
		"00188e": 3340, // Ekahau
		"00188f": 3341, // Montgomery
		"001890": 3342, // RadioCOM
		"001892": 3343, // ads-tec
		"001894": 3344, // NPCore
		"001895": 3345, // Hansun
		"001897": 3346, // JESS-LINK
		"001898": 3347, // Kingstate
		"00189b": 1142, // Thomson
		"00189c": 3348, // Weldex
		"00189d": 3349, // Navcast
		"00189e": 3350, // OMNIKEY
		"00189f": 3351, // Lenntek
		"0018a1": 3352, // Tiqit
		"0018a2": 3353, // XIP
		"0018a3": 3354, // Zippy
		"0018a4": 125,  // ARRIS
		"0018a5": 3355, // ADigit
		"0018a6": 3356, // Persistent
		"0018a8": 3357, // AnNeal
		"0018ae": 3358, // TVT
		"0018af": 152,  // Samsung
		"0018b0": 76,   // Nortel
		"0018b1": 372,  // IBM
		"0018b6": 3359, // S3C
		"0018b9": 2,    // Cisco
		"0018ba": 2,    // Cisco
		"0018be": 3360, // ANSA
		"0018c0": 125,  // ARRIS
		"0018c2": 3361, // Firetide
		"0018c3": 3362, // CS
		"0018c4": 3363, // Raba
		"0018c5": 457,  // Nokia
		"0018c8": 3364, // ISONAS
		"0018c9": 3365, // EOps
		"0018ca": 3366, // Viprinet
		"0018cb": 3367, // Tecobest
		"0018cd": 3368, // Erae
		"0018ce": 3369, // Dreamtech
		"0018d0": 3370, // AtRoad
		"0018d1": 300,  // Siemens
		"0018d3": 3371, // Teamcast
		"0018d5": 3372, // Reigncom
		"0018d6": 3373, // Swirlnet
		"0018db": 3374, // EPL
		"0018dc": 3375, // Prostar
		"0018dd": 3376, // Silicondust
		"0018de": 421,  // Intel
		"0018df": 3377, // Morey
		"0018e0": 3378, // Anaveo
		"0018e1": 3379, // Verkerk
		"0018e3": 3380, // Visualgate
		"0018e4": 3381, // Yiguang
		"0018e5": 3382, // Adhoco
		"0018e7": 3383, // Cameo
		"0018e8": 3384, // Hacetron
		"0018e9": 3385, // Numata
		"0018ea": 3386, // Alltec
		"0018ec": 3387, // Welding
		"0018ef": 3388, // Escape
		"0018f0": 3389, // JOYTOTO
		"0018f3": 1806, // ASUS
		"0018f7": 3390, // Kameleon
		"0018f8": 1783, // Linksys
		"0018f9": 3391, // VVOND
		"0018fb": 3392, // Compro
		"0018fc": 3393, // Altec
		"0018fd": 3394, // Optimal
		"0018fe": 302,  // HP
		"0018ff": 3395, // PowerQuattro
		"001901": 3396, // F1MEDIA
		"001903": 3397, // Bigfoot
		"001904": 3398, // WB
		"001906": 2,    // Cisco
		"001907": 2,    // Cisco
		"001908": 3399, // Duaxes
		"00190a": 3400, // Hasware
		"00190c": 3302, // Encore
		"00190e": 3401, // Atech
		"00190f": 3402, // Advansus
		"001912": 3403, // Welcat
		"001914": 3404, // Winix
		"001915": 576,  // TECOM
		"001916": 3405, // PayTec
		"001917": 3406, // Posiflex
		"001919": 3407, // ASTEL
		"00191a": 3408, // Irlink
		"00191b": 3409, // Sputnik
		"00191c": 3410, // Sensicast
		"00191d": 1427, // Nintendo
		"00191e": 3411, // Beyondwiz
		"00191f": 254,  // Microlink
		"001921": 1115, // Elitegroup
		"001924": 3412, // LBNL
		"001925": 3413, // Intelicis
		"001926": 3414, // BitsGen
		"001927": 3415, // ImCoSys
		"001928": 300,  // Siemens
		"00192c": 125,  // ARRIS
		"00192d": 457,  // Nokia
		"00192f": 2,    // Cisco
		"001930": 2,    // Cisco
		"001931": 3416, // Balluff
		"001932": 3417, // Gude
		"001933": 944,  // Strix
		"001937": 3418, // CommerceGuard
		"001938": 3419, // UMB
		"001939": 3420, // Gigamips
		"00193a": 3421, // Oesolutions
		"00193b": 3422, // LigoWave
		"00193c": 3423, // HighPoint
		"00193f": 1402, // RDI
		"001940": 3424, // Rackable
		"001942": 3425, // ON
		"001943": 3426, // Belden
		"001948": 3427, // AireSpider
		"00194a": 3428, // Testo
		"00194b": 2048, // Sagemcom
		"00194e": 3429, // Ultra
		"00194f": 457,  // Nokia
		"001951": 3430, // NETCONS
		"001952": 3431, // ACOGITO
		"001953": 3432, // Chainleader
		"001954": 3433, // Leaf
		"001955": 2,    // Cisco
		"001956": 2,    // Cisco
		"001959": 3434, // Staccato
		"00195b": 803,  // D-Link
		"00195c": 3435, // Innotech
		"00195e": 125,  // ARRIS
		"00195f": 3436, // Valemount
		"001960": 3437, // DoCoMo
		"001962": 3438, // Commerciant
		"001963": 101,  // Sony
		"001964": 3439, // Doorking
		"001966": 1666, // Asiarock
		"001969": 76,   // Nortel
		"00196a": 3440, // MikroM
		"00196b": 3441, // Danpex
		"00196c": 3442, // Etrovision
		"00196e": 3443, // Metacom
		"00196f": 3444, // SensoPart
		"001970": 3445, // Z-Com
		"001972": 1970, // Plexus
		"001973": 3446, // Zeugma
		"001974": 3447, // 16063
		"001976": 3448, // Xipher
		"001977": 185,  // Extreme
		"001978": 3449, // Datum
		"001979": 457,  // Nokia
		"00197a": 3450, // MAZeT
		"00197b": 3451, // Picotest
		"00197c": 3452, // Riedel
		"00197f": 542,  // Plantronics
		"001980": 3453, // Gridpoint
		"001981": 3454, // Vivox
		"001982": 3455, // SmarDTV
		"001984": 3456, // ESTIC
		"001987": 2153, // Panasonic
		"001988": 3457, // Wi2Wi
		"001989": 3458, // Sonitrol
		"00198c": 3459, // iXSea
		"00198e": 3460, // Demant
		"001991": 3461, // avinfo
		"001992": 202,  // Adtran
		"001994": 3462, // Jorjin
		"001996": 3463, // TurboChef
		"001998": 3464, // Sato
		"001999": 4,    // Fujitsu
		"00199a": 3465, // EDO-EVI
		"00199c": 3466, // Ctring
		"00199d": 3467, // Vizio
		"00199e": 3468, // Nifty
		"00199f": 3469, // DKT
		"0019a1": 869,  // LG
		"0019a2": 3470, // Ordyn
		"0019a3": 3471, // asteel
		"0019a4": 3472, // Austar
		"0019a5": 3473, // RadarFind
		"0019a6": 125,  // ARRIS
		"0019a7": 3474, // ITU-T
		"0019a8": 3475, // WiQuest
		"0019a9": 2,    // Cisco
		"0019aa": 2,    // Cisco
		"0019ab": 3476, // Raycom
		"0019ac": 3477, // GSP
		"0019ad": 3478, // Bobst
		"0019af": 3479, // Rigol
		"0019b1": 3480, // Arrow7
		"0019b2": 3481, // XYnetsoft
		"0019b3": 3482, // Stanford
		"0019b4": 3483, // Intellio
		"0019b7": 457,  // Nokia
		"0019b9": 102,  // Dell
		"0019bb": 302,  // HP
		"0019be": 3484, // Altai
		"0019bf": 3485, // Citiway
		"0019c0": 125,  // ARRIS
		"0019c1": 430,  // Alpsalpine
		"0019c2": 3486, // Equustek
		"0019c3": 3487, // Qualitrol
		"0019c4": 3488, // Infocrypt
		"0019c5": 101,  // Sony
		"0019c6": 3031, // ZTE
		"0019c7": 1640, // Cambridge
		"0019c8": 3489, // AnyDATA
		"0019ca": 3490, // Broadata
		"0019cb": 2692, // ZyXEL
		"0019cc": 3491, // RCG
		"0019cf": 3492, // Salicru
		"0019d0": 3493, // Cathexis
		"0019d1": 421,  // Intel
		"0019d2": 421,  // Intel
		"0019d4": 3494, // ICX
		"0019d7": 3495, // Fortunetek
		"0019d8": 3496, // Maxfor
		"0019d9": 3497, // Zeutschel
		"0019db": 1812, // Micro-Star
		"0019dc": 3498, // ENENSYS
		"0019dd": 3499, // FEI-Zyfer
		"0019de": 3500, // Mobitek
		"0019df": 1142, // Thomson
		"0019e0": 1595, // TP-Link
		"0019e1": 76,   // Nortel
		"0019e2": 826,  // Juniper
		"0019e3": 545,  // Apple
		"0019e4": 1939, // 2Wire
		"0019e7": 2,    // Cisco
		"0019e8": 2,    // Cisco
		"0019ea": 3501, // TeraMage
		"0019eb": 3502, // Pyronix
		"0019ec": 3503, // Sagamore
		"0019ed": 3504, // Axesstel
		"0019f0": 3505, // Unionman
		"0019f3": 3506, // Cetis
		"0019f4": 3507, // Convergens
		"0019f5": 3508, // Imagination
		"0019f6": 3509, // Acconet
		"0019f7": 3510, // Onset
		"0019f9": 3511, // TDK-Lambda
		"0019fb": 3512, // BSkyB
		"0019fd": 1427, // Nintendo
		"0019ff": 3513, // Finnzymes
		"001a00": 1,    // Matrix
		"001a03": 3514, // Angel
		"001a04": 3515, // Interay
		"001a05": 3516, // Optibase
		"001a06": 3517, // OpVista
		"001a08": 3518, // Simoco
		"001a0b": 3519, // Bona
		"001a0d": 3520, // HandHeld
		"001a11": 3521, // Google
		"001a12": 3522, // Essilor
		"001a16": 457,  // Nokia
		"001a17": 3523, // Teak
		"001a19": 2328, // Computer
		"001a1a": 3524, // Gentex/Electro-Acoustic
		"001a1b": 125,  // ARRIS
		"001a1c": 3525, // GT&T
		"001a1e": 1685, // Aruba
		"001a20": 3526, // CMOTECH
		"001a26": 3527, // Deltanode
		"001a27": 3528, // Ubistar
		"001a2a": 2639, // Arcadyan
		"001a2b": 3529, // Ayecom
		"001a2c": 3530, // SATEC
		"001a2f": 2,    // Cisco
		"001a30": 2,    // Cisco
		"001a33": 3531, // ASI
		"001a34": 3532, // Konka
		"001a35": 3533, // BARTEC
		"001a36": 3534, // Aipermon
		"001a37": 3535, // Lear
		"001a38": 3536, // Sanmina-SCI
		"001a39": 3537, // Merten&CoKG
		"001a3a": 3538, // Dongahelecomm
		"001a3c": 3539, // Technowave
		"001a3e": 3540, // Faster
		"001a3f": 3541, // Intelbras
		"001a40": 3542, // A-Four
		"001a41": 3543, // INOCOVA
		"001a42": 3544, // Techcity
		"001a44": 3545, // JWTrading
		"001a47": 3546, // Agami
		"001a48": 3547, // Takacom
		"001a4a": 3548, // Qumranet
		"001a4b": 302,  // HP
		"001a4c": 3549, // Crossbow
		"001a4d": 1929, // Giga-Byte
		"001a4f": 621,  // AVM
		"001a50": 3550, // PheeNet
		"001a53": 3551, // Zylaya
		"001a55": 3552, // ACA-Digital
		"001a56": 3553, // ViewTel
		"001a59": 3554, // Ircona
		"001a5b": 3555, // NetCare
		"001a5c": 3556, // Euchner+Co
		"001a5d": 3557, // Mobinnova
		"001a5e": 3558, // Thincom
		"001a5f": 3559, // KitWorks.fi
		"001a60": 3560, // Wave
		"001a61": 3561, // PacStar
		"001a63": 3562, // Elster
		"001a64": 372,  // IBM
		"001a65": 3563, // Seluxit
		"001a66": 125,  // ARRIS
		"001a68": 3564, // Weltec
		"001a6a": 3565, // Tranzas
		"001a6c": 2,    // Cisco
		"001a6d": 2,    // Cisco
		"001a6e": 3566, // Impro
		"001a70": 1783, // Linksys
		"001a71": 3567, // Diostech
		"001a73": 1450, // Gemtek
		"001a74": 3568, // Procare
		"001a75": 101,  // Sony
		"001a77": 125,  // ARRIS
		"001a78": 3569, // ubtos
		"001a79": 3570, // Telecomunication
		"001a7b": 3571, // Teleco
		"001a7d": 3572, // cyber-blue
		"001a80": 101,  // Sony
		"001a81": 3573, // Zelax
		"001a82": 3574, // PROBA
		"001a83": 3575, // Pegasus
		"001a87": 3576, // Canhold
		"001a88": 3577, // Venergy
		"001a89": 457,  // Nokia
		"001a8a": 152,  // Samsung
		"001a8c": 3578, // Sophos
		"001a8e": 3579, // 3Way
		"001a8f": 76,   // Nortel
		"001a91": 3580, // FusionDynamic
		"001a92": 1806, // ASUS
		"001a94": 3581, // Votronic
		"001a97": 3582, // fitivision
		"001a9c": 3583, // RightHand
		"001a9f": 3584, // A-Link
		"001aa0": 102,  // Dell
		"001aa1": 2,    // Cisco
		"001aa2": 2,    // Cisco
		"001aa3": 3585, // Delorme
		"001aaa": 3586, // Analogic
		"001aac": 3587, // Corelatus
		"001aad": 125,  // ARRIS
		"001aae": 3588, // Savant
		"001aaf": 3589, // Blusens
		"001ab0": 569,  // Signal
		"001ab2": 677,  // Cyber
		"001ab3": 3590, // Visionite
		"001ab4": 3591, // FFEI
		"001ab7": 3592, // Ethos
		"001ab8": 3593, // Anseri
		"001ab9": 3594, // PMC
		"001abb": 3595, // Fontal
		"001abc": 2996, // U4EA
		"001abd": 3596, // Impatica
		"001ac0": 3597, // Joybien
		"001ac1": 160,  // 3Com
		"001ac2": 3598, // YEC
		"001ac3": 3115, // Scientific-Atlanta
		"001ac4": 1939, // 2Wire
		"001ac5": 3599, // Keysight
		"001ac7": 3600, // Unipoint
		"001ac8": 3601, // ISL
		"001ac9": 3602, // Suzuken
		"001aca": 3603, // Tilera
		"001acb": 3604, // Autocom
		"001acd": 3605, // Tidel
		"001ace": 3606, // Yupiteru
		"001ad0": 3607, // Albis
		"001ad1": 2690, // Fargo
		"001ad3": 3608, // Vamp
		"001ad4": 3609, // iPOX
		"001ad8": 3610, // AlsterAero
		"001ada": 3611, // Biz-2-Me
		"001adb": 125,  // ARRIS
		"001adc": 457,  // Nokia
		"001add": 3612, // PePWave
		"001ade": 125,  // ARRIS
		"001adf": 3613, // Interactivetv
		"001ae2": 2,    // Cisco
		"001ae3": 2,    // Cisco
		"001ae4": 3614, // Medicis
		"001ae5": 3615, // Mvox
		"001ae7": 1357, // Aztek
		"001ae9": 1427, // Nintendo
		"001aec": 3616, // Keumbee
		"001aed": 3617, // IncOTEC
		"001aee": 3618, // Shenztech
		"001aef": 3619, // Loopcomm
		"001af3": 3620, // Samyoung
		"001af4": 3621, // Handreamnet
		"001af5": 3622, // Pentaone
		"001af6": 3623, // Woven
		"001af9": 3624, // AeroVIronment
		"001afb": 3625, // Joby
		"001afc": 3626, // ModusLink
		"001afd": 3627, // Evolis
		"001afe": 3628, // Sofacreal
		"001aff": 3629, // Wizyoung
		"001b00": 3630, // Neopost
		"001b02": 3631, // ED
		"001b03": 3632, // Action
		"001b05": 3633, // YMC
		"001b07": 3634, // Mendocino
		"001b0b": 3635, // Phidgets
		"001b0c": 2,    // Cisco
		"001b0d": 2,    // Cisco
		"001b0f": 3636, // Petratec
		"001b11": 803,  // D-Link
		"001b12": 3637, // Apprion
		"001b13": 3638, // Icron
		"001b15": 3639, // Voxtel
		"001b16": 3640, // Celtro
		"001b1b": 300,  // Siemens
		"001b1c": 3259, // Coherent
		"001b1d": 3641, // Phoenix
		"001b20": 3642, // TPine
		"001b21": 421,  // Intel
		"001b23": 3643, // SimpleComTools
		"001b24": 3071, // Quanta
		"001b25": 76,   // Nortel
		"001b28": 3644, // Polygon
		"001b29": 3645, // Avantis
		"001b2a": 2,    // Cisco
		"001b2b": 2,    // Cisco
		"001b2c": 3646, // ATRON
		"001b2d": 3647, // Med-Eng
		"001b2e": 3648, // Sinkyo
		"001b2f": 1368, // Netgear
		"001b30": 3649, // Solitech
		"001b32": 2025, // QLogic
		"001b33": 457,  // Nokia
		"001b36": 3650, // Tsubata
		"001b37": 3651, // Computec
		"001b38": 358,  // Compal
		"001b39": 3652, // Proxicast
		"001b3a": 3653, // SIMS
		"001b3b": 3654, // Yi-Qing
		"001b3d": 3655, // EuroTel
		"001b3e": 3656, // Curtis
		"001b44": 3657, // SanDisk
		"001b45": 21,   // ABB
		"001b46": 3658, // Blueone
		"001b47": 3659, // Futarque
		"001b4a": 3660, // W&W
		"001b4b": 3661, // SANION
		"001b4c": 3662, // Signtech
		"001b4d": 3663, // Areca
		"001b4f": 620,  // Avaya
		"001b51": 3664, // Vector
		"001b52": 125,  // ARRIS
		"001b53": 2,    // Cisco
		"001b54": 2,    // Cisco
		"001b56": 3665, // Tehuti
		"001b57": 3666, // Semindia
		"001b59": 101,  // Sony
		"001b5b": 1939, // 2Wire
		"001b5c": 3667, // Azuretec
		"001b5d": 3668, // Vololink
		"001b5e": 3669, // BPL
		"001b5f": 3670, // Alien
		"001b60": 3671, // Navigon
		"001b63": 545,  // Apple
		"001b64": 3672, // IsaacLandKorea
		"001b66": 3117, // Sennheiser
		"001b67": 2,    // Cisco
		"001b68": 3673, // Modnnet
		"001b69": 3674, // Equaline
		"001b6b": 3675, // Swyx
		"001b6d": 3676, // Midtronics
		"001b6e": 3599, // Keysight
		"001b6f": 3677, // Teletrak
		"001b71": 3678, // Telular
		"001b74": 3679, // MiraLink
		"001b75": 3680, // Hypermedia
		"001b76": 3681, // Ripcode
		"001b77": 421,  // Intel
		"001b78": 302,  // HP
		"001b7a": 1427, // Nintendo
		"001b7b": 3682, // Tintometer
		"001b7e": 3683, // Beckmann
		"001b80": 3684, // LORD
		"001b83": 3685, // Finsoft
		"001b84": 3686, // Scan
		"001b87": 3687, // Deepsound
		"001b8a": 3688, // 2M
		"001b8c": 3689, // JMicron
		"001b8d": 3690, // Electronic
		"001b8f": 2,    // Cisco
		"001b90": 2,    // Cisco
		"001b91": 3691, // Efkon
		"001b92": 3692, // l-acoustics
		"001b97": 3693, // Violin
		"001b98": 152,  // Samsung
		"001b9b": 3694, // Hose-McCann
		"001b9e": 2544, // Askey
		"001b9f": 3695, // Calyptech
		"001ba0": 3696, // Awox
		"001ba1": 3697, // Amic
		"001ba5": 3698, // MyungMin
		"001ba6": 3699, // intotech
		"001ba7": 3700, // Lorica
		"001ba8": 3701, // UBI&MOBI
		"001ba9": 3230, // Brother
		"001baa": 3702, // XenICs
		"001bab": 3703, // Telchemyorporated
		"001bad": 3704, // iControl
		"001baf": 457,  // Nokia
		"001bb0": 3705, // Bharat
		"001bb1": 1592, // Wistron
		"001bb2": 3706, // Intellect
		"001bb3": 3707, // Condalo
		"001bb4": 3708, // Airvod
		"001bb5": 3709, // Cherry
		"001bb6": 3710, // Bird
		"001bb8": 3711, // Blueway
		"001bb9": 1115, // Elitegroup
		"001bba": 76,   // Nortel
		"001bbb": 3712, // RFTech
		"001bbf": 2048, // Sagemcom
		"001bc0": 826,  // Juniper
		"001bc1": 3713, // HOLUX
		"001bc3": 3714, // Mobisolution
		"001bc4": 3715, // Ultratec
		"001bc7": 3716, // StarVedia
		"001bc8": 3717, // Miura
		"001bcb": 3718, // Pempek
		"001bcd": 3719, // Daviscomms
		"001bcf": 3720, // Dataupia
		"001bd0": 3721, // Identec
		"001bd1": 3722, // Sogestmatic
		"001bd3": 2153, // Panasonic
		"001bd4": 2,    // Cisco
		"001bd5": 2,    // Cisco
		"001bd8": 3723, // FLIR
		"001bda": 1135, // UTStarcom
		"001bdc": 3724, // Vencer
		"001bdd": 125,  // ARRIS
		"001bde": 3725, // Renkus-Heinz
		"001be0": 3726, // TELENOT
		"001be1": 3727, // ViaLogy
		"001be2": 3728, // AhnLab
		"001be5": 3729, // 802automation
		"001be6": 3730, // VR
		"001be7": 3731, // Postek
		"001be8": 3732, // Ultratronik
		"001be9": 858,  // Broadcom
		"001bea": 1427, // Nintendo
		"001beb": 3733, // DMP
		"001bec": 3734, // Netio
		"001bed": 90,   // Brocade
		"001bee": 457,  // Nokia
		"001bf2": 3735, // Kworld
		"001bf4": 3736, // Kenwin
		"001bf6": 3737, // CoNWISE
		"001bf8": 3738, // Digitrax
		"001bfb": 430,  // Alpsalpine
		"001bfc": 1806, // ASUS
		"001bfd": 3739, // Dignsys
		"001bfe": 3740, // Zavio
		"001c04": 3741, // Airgain
		"001c06": 300,  // Siemens
		"001c07": 3742, // Cwlinux
		"001c08": 3743, // Echo360
		"001c09": 3744, // SAE
		"001c0c": 3745, // TANITA
		"001c0d": 3746, // G-Technology
		"001c0e": 2,    // Cisco
		"001c0f": 2,    // Cisco
		"001c10": 1783, // Linksys
		"001c11": 125,  // ARRIS
		"001c12": 125,  // ARRIS
		"001c13": 3747, // Optsys
		"001c14": 809,  // VMware
		"001c15": 3748, // iPhotonix
		"001c17": 76,   // Nortel
		"001c1b": 3749, // Hyperstone
		"001c1c": 3750, // Center
		"001c1e": 3751, // emtrion
		"001c21": 3752, // Nucsafe
		"001c23": 102,  // Dell
		"001c27": 3753, // Sunell
		"001c28": 3754, // Sphairon
		"001c2a": 3755, // Envisacor
		"001c2b": 3756, // Alertme.com
		"001c2c": 3757, // Synapse
		"001c2d": 3758, // FlexRadio
		"001c2f": 3759, // Pfister
		"001c32": 3760, // Telian
		"001c33": 1114, // Sutron
		"001c35": 457,  // Nokia
		"001c36": 3761, // iNEWiT
		"001c37": 3762, // Callpod
		"001c38": 3763, // Bio-Rad
		"001c3a": 3764, // Element
		"001c3b": 3765, // AmRoad
		"001c3d": 3766, // WaveStorm
		"001c3e": 3767, // ECKey
		"001c40": 3768, // VDG-Security
		"001c42": 3769, // Parallels
		"001c43": 152,  // Samsung
		"001c46": 3770, // Qtum
		"001c48": 3771, // WiDeFi
		"001c49": 3772, // Zoltan
		"001c4a": 621,  // AVM
		"001c4b": 3773, // Gener8
		"001c4e": 3774, // TASA
		"001c4f": 3775, // Macab
		"001c51": 3776, // Celeno
		"001c54": 3777, // Hillstone
		"001c56": 3778, // Pado
		"001c57": 2,    // Cisco
		"001c58": 2,    // Cisco
		"001c5b": 3779, // Chubb
		"001c5f": 3780, // Winland
		"001c62": 869,  // LG
		"001c63": 3781, // Truen
		"001c64": 2227, // Landis+Gyr
		"001c65": 3782, // JoeScan
		"001c66": 3783, // Ucamp
		"001c67": 3784, // Pumpkin
		"001c6a": 3785, // Weiss
		"001c6b": 3786, // CoVAX
		"001c6c": 3787, // 30805
		"001c6d": 3788, // Kyohritsu
		"001c6e": 3789, // Newbury
		"001c6f": 3790, // Emfit
		"001c70": 3791, // Novacomm
		"001c71": 3792, // Emergent
		"001c73": 3793, // Arista
		"001c74": 3794, // Syswan
		"001c75": 3795, // Segnet
		"001c77": 3796, // Prodys
		"001c7b": 3797, // Castlenet
		"001c7c": 3798, // Perq
		"001c7d": 3799, // Excelpoint
		"001c7e": 35,   // Toshiba
		"001c82": 3800, // Genew
		"001c85": 3801, // Eunicorn
		"001c86": 3802, // Cranite
		"001c87": 3803, // Uriver
		"001c88": 3804, // Transystem
		"001c89": 2271, // Force
		"001c8a": 3805, // Cirrascale
		"001c8c": 3806, // Dial
		"001c8f": 587,  // Advanced
		"001c90": 3807, // Empacket
		"001c91": 3808, // Gefen
		"001c92": 3809, // Tervela
		"001c93": 3810, // ExaDigm
		"001c95": 3811, // Opticomm
		"001c96": 3812, // Linkwise
		"001c97": 3813, // Enzytek
		"001c98": 3814, // Lucky
		"001c99": 3815, // Shunra
		"001c9a": 457,  // Nokia
		"001c9b": 3816, // FEIG
		"001c9c": 76,   // Nortel
		"001c9d": 3817, // Liecthi
		"001c9f": 3818, // Razorstream
		"001ca1": 3819, // Akamai
		"001ca3": 3820, // Terra
		"001ca4": 101,  // Sony
		"001ca5": 3821, // Zygo
		"001ca6": 3822, // Win4NET
		"001caa": 3823, // Bellon
		"001cac": 3824, // Qniq
		"001cae": 3825, // WiChorus
		"001caf": 3826, // Plato
		"001cb0": 2,    // Cisco
		"001cb1": 2,    // Cisco
		"001cb2": 3827, // BPT
		"001cb3": 545,  // Apple
		"001cb8": 3828, // CBC
		"001cba": 3829, // VerScient
		"001cbb": 3830, // MusicianLink
		"001cbc": 3831, // CastGrabber
		"001cbe": 1427, // Nintendo
		"001cbf": 421,  // Intel
		"001cc0": 421,  // Intel
		"001cc1": 125,  // ARRIS
		"001cc3": 125,  // ARRIS
		"001cc4": 302,  // HP
		"001cc5": 160,  // 3Com
		"001cc6": 3832, // ProStor
		"001cc9": 3833, // Kaise
		"001ccc": 2220, // BlackBerry
		"001ccd": 3834, // Alektrona
		"001ccf": 3835, // Limetek
		"001cd0": 3836, // Circleone
		"001cd3": 3837, // ZP
		"001cd4": 457,  // Nokia
		"001cd5": 3838, // ZeeVee
		"001cd6": 457,  // Nokia
		"001cd9": 1389, // GlobalTop
		"001cda": 3839, // Exegin
		"001cdb": 3840, // Carpoint
		"001cdc": 2122, // Custom
		"001cdd": 3841, // Cowbell
		"001cdf": 2468, // Belkin
		"001ce2": 3842, // Attero
		"001ce3": 3843, // Optimedical
		"001ce5": 3844, // MBS
		"001ce6": 3845, // Innes
		"001ce8": 3846, // Cummins
		"001ce9": 3847, // Galaxy
		"001cea": 3115, // Scientific-Atlanta
		"001ceb": 76,   // Nortel
		"001cec": 3848, // Mobilesoft
		"001ced": 3849, // Environnement
		"001cee": 3205, // Sharp
		"001cef": 389,  // Primax
		"001cf0": 803,  // D-Link
		"001cf1": 3850, // SUPoX
		"001cf2": 3851, // Tenlon
		"001cf4": 3852, // Media
		"001cf5": 3853, // Wiseblue
		"001cf6": 2,    // Cisco
		"001cf7": 3854, // AudioScience
		"001cf8": 3855, // Parade
		"001cf9": 2,    // Cisco
		"001cfa": 3856, // Alarm.com
		"001cfb": 125,  // ARRIS
		"001cfd": 3857, // Universal
		"001cfe": 3858, // Quartics
		"001cff": 3859, // Napera
		"001d00": 3860, // Brivo
		"001d03": 3861, // Design
		"001d06": 3862, // HM
		"001d09": 102,  // Dell
		"001d0c": 3863, // MobileCompia
		"001d0d": 101,  // Sony
		"001d0e": 3864, // Agapha
		"001d0f": 1595, // TP-Link
		"001d12": 3865, // Rohm
		"001d13": 3866, // NextGTV
		"001d16": 3187, // SFR
		"001d19": 2639, // Arcadyan
		"001d1b": 3867, // Sangean
		"001d1d": 3868, // Inter-M
		"001d20": 3869, // Comtrend
		"001d23": 3870, // Sensus
		"001d25": 152,  // Samsung
		"001d26": 3871, // Rockridgesound
		"001d27": 3872, // NAC-Intercom
		"001d28": 101,  // Sony
		"001d29": 3873, // Doro
		"001d2c": 3874, // Wavetrend
		"001d2d": 3875, // Pylone
		"001d2e": 2737, // Ruckus
		"001d2f": 3876, // QuantumVision
		"001d32": 3877, // Longkay
		"001d33": 3878, // Maverick
		"001d34": 3879, // SYRIS
		"001d35": 3880, // Viconics
		"001d38": 732,  // Seagate
		"001d39": 3881, // Moohadigital
		"001d3b": 457,  // Nokia
		"001d3c": 3882, // Muscle
		"001d3d": 3883, // Avidyne
		"001d3f": 3884, // Mitron
		"001d42": 76,   // Nortel
		"001d44": 3885, // Krohne
		"001d45": 2,    // Cisco
		"001d46": 2,    // Cisco
		"001d47": 3886, // Covote
		"001d4e": 3887, // TCM
		"001d4f": 545,  // Apple
		"001d50": 3888, // Spinetix
		"001d53": 3889, // S&O
		"001d55": 3890, // ZANTAZ
		"001d56": 3891, // Kramer
		"001d58": 3892, // CQ
		"001d5a": 1939, // 2Wire
		"001d5c": 3893, // Tom
		"001d60": 1806, // ASUS
		"001d61": 3894, // BIJ
		"001d62": 3895, // InPhase
		"001d67": 3896, // Amec
		"001d6a": 2236, // Alpha
		"001d6b": 125,  // ARRIS
		"001d6c": 3897, // ClariPhy
		"001d6d": 3898, // Confidant
		"001d6e": 457,  // Nokia
		"001d6f": 3899, // Chainzone
		"001d70": 2,    // Cisco
		"001d71": 2,    // Cisco
		"001d72": 1592, // Wistron
		"001d73": 1077, // Buffalo
		"001d75": 3900, // Radioscape
		"001d76": 3901, // Eyeheight
		"001d77": 3902, // NSGate
		"001d79": 3903, // Signamax
		"001d7d": 1929, // Giga-Byte
		"001d7e": 1783, // Linksys
		"001d7f": 3904, // Tekron
		"001d83": 3905, // Emitech
		"001d84": 64,   // Gateway
		"001d86": 3906, // Shinwa
		"001d88": 3907, // Clearwire
		"001d89": 3908, // VaultStor
		"001d8a": 3909, // TechTrex
		"001d8e": 3910, // Alereon
		"001d8f": 3911, // PureWave
		"001d91": 3912, // Digitize
		"001d92": 1812, // Micro-Star
		"001d93": 3913, // Modacom
		"001d94": 3914, // Climax
		"001d95": 1029, // Flash
		"001d97": 3915, // Alertus
		"001d98": 457,  // Nokia
		"001d9a": 3916, // Godex
		"001d9d": 3917, // Artjoy
		"001d9e": 3918, // Axion
		"001da1": 2,    // Cisco
		"001da2": 2,    // Cisco
		"001da3": 3919, // SabiOso
		"001da5": 3398, // WB
		"001da8": 3920, // Takahata
		"001da9": 3921, // Castles
		"001daa": 3922, // DrayTek
		"001dac": 3923, // Gigamon
		"001dad": 3924, // Sinotech
		"001daf": 76,   // Nortel
		"001db1": 3925, // Crescendo
		"001db5": 826,  // Juniper
		"001db6": 3926, // BestComm
		"001db7": 3927, // Tendril
		"001db8": 3928, // Intoto
		"001dba": 101,  // Sony
		"001dbc": 1427, // Nintendo
		"001dbd": 3929, // Versamed
		"001dbe": 125,  // ARRIS
		"001dbf": 3930, // Radiient
		"001dc2": 3931, // Xortec
		"001dc4": 3932, // AIOI
		"001dc6": 3933, // SNR
		"001dc8": 3934, // Navionics
		"001dc9": 3935, // GainSpan
		"001dca": 3936, // PAV
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
		"001dd7": 3937, // Algolith
		"001dd8": 612,  // Microsoft
		"001ddb": 3938, // C-BEL
		"001ddf": 3939, // Sunitec
		"001de0": 421,  // Intel
		"001de1": 421,  // Intel
		"001de2": 3940, // Radionor
		"001de3": 3941, // Intuicom
		"001de5": 2,    // Cisco
		"001de6": 2,    // Cisco
		"001de9": 457,  // Nokia
		"001deb": 3942, // DINEC
		"001dec": 3943, // Marusys
		"001def": 3944, // Trimm
		"001df0": 3945, // Vidient
		"001df1": 3946, // Intego
		"001df2": 3947, // Netflix
		"001df4": 3948, // Magellan
		"001df5": 3949, // Sunshine
		"001df6": 152,  // Samsung
		"001df9": 3950, // Cybiotronics
		"001dfb": 3951, // NETCLEUS
		"001dfc": 3952, // Ksic
		"001dfd": 457,  // Nokia
		"001dfe": 1158, // Palm
		"001e03": 3953, // LiComm
		"001e04": 3954, // Hanson
		"001e06": 3955, // Wibrain
		"001e07": 3956, // Winy
		"001e08": 3957, // Centec
		"001e09": 3958, // ZEFATEK
		"001e0a": 3959, // Syba
		"001e0b": 302,  // HP
		"001e0d": 3960, // Micran
		"001e0f": 3961, // Briot
		"001e10": 3334, // Huawei
		"001e11": 3962, // Elelux
		"001e12": 3963, // Ecolab
		"001e13": 2,    // Cisco
		"001e14": 2,    // Cisco
		"001e16": 3964, // Keytronix
		"001e17": 3965, // STN
		"001e19": 3966, // Gtri
		"001e1f": 76,   // Nortel
		"001e20": 3967, // Intertain
		"001e21": 551,  // Qisda
		"001e26": 3968, // Digifriends
		"001e27": 3969, // SBN
		"001e28": 3970, // Lumexis
		"001e29": 3971, // Hypertherm
		"001e2a": 1368, // Netgear
		"001e2c": 3972, // CyVerse
		"001e2d": 3973, // Stim
		"001e2f": 3974, // DiMoto
		"001e30": 3975, // Shireen
		"001e31": 3976, // Infomark
		"001e32": 3977, // Zensys
		"001e33": 3978, // Inventec
		"001e34": 3979, // CryptoMetrics
		"001e35": 1427, // Nintendo
		"001e36": 3980, // Ipte
		"001e38": 3981, // Bluecard
		"001e39": 3982, // Comsys
		"001e3a": 457,  // Nokia
		"001e3b": 457,  // Nokia
		"001e3d": 430,  // Alpsalpine
		"001e3e": 3983, // KMW
		"001e3f": 3984, // TrellisWare
		"001e42": 3985, // Teltonika
		"001e44": 3986, // Santec
		"001e45": 101,  // Sony
		"001e46": 125,  // ARRIS
		"001e48": 3987, // Wi-Links
		"001e49": 2,    // Cisco
		"001e4a": 2,    // Cisco
		"001e4f": 102,  // Dell
		"001e50": 3988, // Battistoni
		"001e52": 545,  // Apple
		"001e53": 3989, // Further
		"001e55": 3990, // CoWON
		"001e57": 3991, // ALCOMA
		"001e58": 803,  // D-Link
		"001e5a": 125,  // ARRIS
		"001e5b": 3992, // Unitron
		"001e5e": 3993, // Computime
		"001e5f": 3994, // KwikByte
		"001e61": 3995, // ITEC
		"001e62": 3996, // Siemon
		"001e63": 3997, // Vibro-Meter
		"001e64": 421,  // Intel
		"001e65": 421,  // Intel
		"001e67": 421,  // Intel
		"001e68": 3071, // Quanta
		"001e69": 1142, // Thomson
		"001e6c": 3998, // Opaque
		"001e6f": 3999, // Magna-Power
		"001e72": 4000, // PCS
		"001e73": 3031, // ZTE
		"001e74": 2048, // Sagemcom
		"001e75": 869,  // LG
		"001e77": 4001, // Air2App
		"001e78": 4002, // Owitek
		"001e79": 2,    // Cisco
		"001e7a": 2,    // Cisco
		"001e7c": 4003, // Taiwick
		"001e7d": 152,  // Samsung
		"001e7e": 76,   // Nortel
		"001e80": 2156, // Icotera
		"001e81": 4004, // CNB
		"001e82": 3657, // SanDisk
		"001e84": 4005, // Pika
		"001e85": 4006, // Lagotek
		"001e86": 4007, // MEL
		"001e87": 4008, // Realease
		"001e89": 4009, // CRFS
		"001e8a": 4010, // eCopy
		"001e8c": 1806, // ASUS
		"001e8d": 125,  // ARRIS
		"001e8e": 4011, // Hunkeler
		"001e8f": 87,   // Canon
		"001e90": 1115, // Elitegroup
		"001e91": 4012, // KIMIN
		"001e93": 4013, // CiriTech
		"001e94": 4014, // Supercom
		"001e95": 4015, // Sigmalink
		"001e96": 4016, // Sepura
		"001e98": 4017, // GreenLine
		"001e99": 4018, // Vantanol
		"001e9b": 4019, // San-Eisha
		"001e9c": 4020, // Fidustron
		"001e9d": 4021, // Recall
		"001e9f": 4022, // Visioneering
		"001ea0": 4023, // XLN-t
		"001ea1": 4024, // Brunata
		"001ea2": 4025, // Symx
		"001ea3": 457,  // Nokia
		"001ea4": 457,  // Nokia
		"001ea5": 4026, // ROBOTOUS
		"001ea7": 2246, // Actiontec
		"001ea9": 1427, // Nintendo
		"001eaa": 4027, // E-Senza
		"001eab": 4028, // TeleWell
		"001eac": 4029, // Armadeus
		"001ead": 4030, // Wingtech
		"001eb0": 4031, // ImesD
		"001eb1": 4032, // Cryptsoft
		"001eb2": 869,  // LG
		"001eb4": 4033, // Unifat
		"001eb7": 4034, // TBTech
		"001eb8": 4035, // Aloys
		"001ebb": 4036, // Bluelight
		"001ebd": 2,    // Cisco
		"001ebe": 2,    // Cisco
		"001ec0": 706,  // Microchip
		"001ec1": 160,  // 3Com
		"001ec2": 545,  // Apple
		"001ec3": 4037, // Kozio
		"001ec4": 4038, // Celio
		"001ec6": 4039, // Obvius
		"001ec7": 1939, // 2Wire
		"001ec9": 102,  // Dell
		"001eca": 76,   // Nortel
		"001ecc": 4040, // Cdvi
		"001ecd": 4041, // KYLAND
		"001ece": 4042, // BISA
		"001ecf": 796,  // Philips
		"001ed0": 4043, // Ingespace
		"001ed4": 4044, // Doble
		"001ed5": 4045, // Tekon-Automatics
		"001edc": 101,  // Sony
		"001ede": 4046, // BYD
		"001ee0": 4047, // Urmet
		"001ee1": 152,  // Samsung
		"001ee2": 152,  // Samsung
		"001ee3": 4048, // T&W
		"001ee5": 1783, // Linksys
		"001ee7": 4049, // Epic
		"001ee8": 4050, // Mytek
		"001ee9": 4051, // Stoneridge
		"001eeb": 4052, // Talk-A-Phone
		"001eec": 358,  // Compal
		"001eed": 4053, // Adventiq
		"001eee": 4054, // ETL
		"001eef": 4055, // Cantronic
		"001ef0": 4056, // Gigafin
		"001ef1": 4057, // Servimat
		"001ef3": 4058, // From2
		"001ef6": 2,    // Cisco
		"001ef7": 2,    // Cisco
		"001ef8": 4059, // Emfinity
		"001efa": 4060, // PROTEI
		"001eff": 4061, // Mueller-Elektronik
		"001f00": 457,  // Nokia
		"001f01": 457,  // Nokia
		"001f02": 4062, // Pixelmetrix
		"001f03": 4063, // NUM
		"001f04": 4064, // Granch
		"001f05": 4065, // iTAS
		"001f07": 4066, // AZTEQ
		"001f08": 4067, // Risco
		"001f09": 4068, // Jastec
		"001f0a": 76,   // Nortel
		"001f0f": 4069, // Select
		"001f11": 4070, // Openmoko
		"001f12": 826,  // Juniper
		"001f14": 4071, // NexG
		"001f15": 4072, // Bioscrypt
		"001f16": 1592, // Wistron
		"001f17": 4073, // IDX
		"001f18": 4074, // Hakusan.Mfg
		"001f19": 4075, // BEN-RI
		"001f1a": 4076, // Prominvest
		"001f1b": 4077, // RoyalTek
		"001f1e": 4078, // Astec
		"001f1f": 115,  // Edimax
		"001f20": 4079, // Logitech
		"001f23": 4080, // Interacoustics
		"001f24": 4081, // Digitview
		"001f25": 3844, // MBS
		"001f26": 2,    // Cisco
		"001f27": 2,    // Cisco
		"001f29": 302,  // HP
		"001f2a": 4082, // Accm
		"001f2c": 4083, // Starbridge
		"001f2f": 4084, // Berker
		"001f30": 4085, // Travelping
		"001f31": 4086, // Radiocomp
		"001f32": 1427, // Nintendo
		"001f33": 1368, // Netgear
		"001f35": 4087, // AIR802
		"001f38": 4088, // Positron
		"001f3b": 421,  // Intel
		"001f3c": 421,  // Intel
		"001f3d": 4089, // Qbit
		"001f3f": 621,  // AVM
		"001f40": 4090, // Speakercraft
		"001f41": 2737, // Ruckus
		"001f42": 4091, // Etherstack
		"001f45": 312,  // Enterasys
		"001f46": 76,   // Nortel
		"001f48": 4092, // Mojix
		"001f4c": 4093, // Roseman
		"001f4d": 4094, // Segnetics
		"001f4f": 4095, // Thinkware
		"001f50": 4096, // Swissdis
		"001f51": 4097, // HD
		"001f54": 4098, // Lorex
		"001f5b": 545,  // Apple
		"001f5c": 457,  // Nokia
		"001f5d": 457,  // Nokia
		"001f5e": 4099, // Dyna
		"001f5f": 4100, // Blatand
		"001f60": 4101, // Compass
		"001f61": 4102, // Talent
		"001f66": 4103, // Planar
		"001f67": 89,   // Hitachi
		"001f69": 4104, // Pingood
		"001f6a": 4105, // PacketFlux
		"001f6b": 869,  // LG
		"001f6c": 2,    // Cisco
		"001f6d": 2,    // Cisco
		"001f6e": 2502, // Vtech
		"001f70": 4106, // Botik
		"001f71": 4107, // xG
		"001f73": 4108, // Teraview
		"001f76": 1621, // AirLogic
		"001f79": 4109, // Lodam
		"001f7a": 4110, // WiWide
		"001f7b": 4111, // TechNexion
		"001f7c": 4112, // Witelcom
		"001f7e": 125,  // ARRIS
		"001f7f": 4113, // Phabrix
		"001f80": 4114, // Lucas
		"001f82": 1630, // Cal-Comp
		"001f83": 4115, // Teleplan
		"001f86": 4116, // digEcor
		"001f87": 4117, // Skydigital
		"001f89": 4118, // Signalion
		"001f8c": 4119, // CCS
		"001f90": 2246, // Actiontec
		"001f92": 687,  // Motorola
		"001f93": 4120, // Xiotech
		"001f94": 4121, // Lascar
		"001f95": 2048, // Sagemcom
		"001f96": 4122, // Aprotech
		"001f98": 4123, // Daiichi-Dentsu
		"001f99": 4124, // SERONICS
		"001f9a": 76,   // Nortel
		"001f9b": 4125, // Posbro
		"001f9c": 4126, // Ledco
		"001f9d": 2,    // Cisco
		"001f9e": 2,    // Cisco
		"001fa0": 4127, // A10
		"001fa1": 4128, // Gtran
		"001fa3": 4048, // T&W
		"001fa5": 4129, // Blue-White
		"001fa7": 101,  // Sony
		"001faa": 4130, // Taseon
		"001fac": 4131, // Goodmill
		"001faf": 4132, // NextIO
		"001fb0": 4133, // TimeIPS
		"001fb1": 4134, // Cybertech
		"001fb3": 1939, // 2Wire
		"001fb4": 4135, // SmartShare
		"001fb7": 4136, // WiMate
		"001fb9": 4137, // Paltronics
		"001fba": 4138, // Boyoung
		"001fbb": 4139, // Xenatech
		"001fbc": 4140, // EVGA
		"001fbd": 2848, // Kyocera
		"001fc1": 4141, // Hanlong
		"001fc3": 4142, // SmartSynch
		"001fc4": 125,  // ARRIS
		"001fc5": 1427, // Nintendo
		"001fc6": 1806, // ASUS
		"001fc8": 4143, // Up-Today
		"001fc9": 2,    // Cisco
		"001fca": 2,    // Cisco
		"001fcb": 4144, // NIW
		"001fcc": 152,  // Samsung
		"001fcd": 152,  // Samsung
		"001fce": 4145, // Qtech
		"001fcf": 3054, // MSI
		"001fd0": 1929, // Giga-Byte
		"001fd1": 4146, // Optex
		"001fd3": 4147, // RIVA
		"001fd4": 4148, // 4IPNET
		"001fd7": 4149, // Telerad
		"001fd8": 4150, // A-Trust
		"001fd9": 4151, // RSD
		"001fda": 76,   // Nortel
		"001fdd": 4152, // GDI
		"001fde": 457,  // Nokia
		"001fdf": 457,  // Nokia
		"001fe0": 4153, // EdgeVelocity
		"001fe3": 869,  // LG
		"001fe4": 101,  // Sony
		"001fe5": 4154, // In-Circuit
		"001fe6": 4155, // Alphion
		"001fe7": 4156, // Simet
		"001fe8": 4157, // KURUSUGAWA
		"001fe9": 4158, // Printrex
		"001fec": 3757, // Synapse
		"001fed": 4159, // Tecan
		"001fee": 4160, // ubisys
		"001fef": 4161, // Shinsei
		"001ff2": 4162, // VIA
		"001ff3": 545,  // Apple
		"001ff8": 300,  // Siemens
		"001ffa": 4163, // Coretree
		"001ffc": 4164, // Riccius+Sohn
		"001ffd": 4165, // Indigo
		"001fff": 4166, // Respironics
		"002000": 613,  // Lexmark
		"002001": 4167, // DSP
		"002002": 4168, // Seritech
		"002004": 4169, // Yamatake-Honeywell
		"002005": 4170, // Simple
		"002006": 4171, // Garrett
		"002007": 4172, // SFA
		"002008": 4173, // Cable
		"00200a": 4174, // Source-Comm
		"00200b": 4175, // Octagon
		"00200c": 4176, // Adastra
		"00200e": 4177, // NSSLGlobal
		"00200f": 4178, // EBRAINS
		"002011": 4179, // Canopus
		"002013": 4180, // Diversified
		"002015": 4181, // Actis
		"002017": 4182, // Orbotech
		"002018": 164,  // CIS
		"002019": 4183, // Ohler
		"00201a": 2253, // MRV
		"00201c": 4184, // Excel
		"00201d": 4185, // Katana
		"00201e": 4186, // Netquest
		"002020": 4187, // Megatron
		"002021": 4188, // Algorithms
		"002022": 4189, // NMS
		"002023": 4190, // T.C
		"002025": 1998, // Control
		"002026": 4191, // Amkly
		"002029": 4192, // Teleprocessing
		"00202c": 4193, // Welltronix
		"00202d": 4194, // Taiyo
		"00202f": 4195, // Zeta
		"002033": 3757, // Synapse
		"002035": 372,  // IBM
		"002036": 4196, // BMC
		"002037": 732,  // Seagate
		"002039": 4197, // Scinets
		"00203b": 4198, // Wisdm
		"00203c": 4199, // Eurotime
		"00203e": 4200, // LogiCan
		"00203f": 4201, // Juki
		"002040": 125,  // ARRIS
		"002042": 4202, // Datametrics
		"002043": 4203, // Neuron
		"002044": 4204, // Genitech
		"002045": 4205, // ION
		"002046": 4206, // Ciprico
		"002047": 4207, // Steinbrecher
		"002048": 31,   // Marconi
		"002049": 4208, // Comtron
		"00204a": 4209, // Pronet
		"00204b": 4210, // Autocomputer
		"00204c": 3884, // Mitron
		"00204d": 4211, // Inovis
		"002050": 853,  // Korea
		"002051": 4212, // Verilink
		"002052": 4213, // Ragula
		"002054": 4214, // Sycamore
		"002055": 4215, // Altech
		"002056": 4216, // Neoproducts
		"002059": 4217, // Miro
		"00205b": 4218, // Kentrox
		"00205d": 4219, // Nanomatic
		"00205f": 4220, // Gammadata
		"002061": 4221, // GarrettCom
		"002067": 67,   // Private
		"002068": 4222, // Isdyne
		"002069": 4223, // Isdn
		"00206a": 4224, // Osaka
		"00206c": 4225, // Evergreen
		"00206e": 4226, // Xact
		"00206f": 4227, // Flowpoint
		"002070": 4228, // Hynet
		"002071": 4229, // IBR
		"002073": 4230, // Fusion
		"002074": 4231, // Sungwoon
		"002076": 4232, // Reudo
		"002077": 4233, // Kardios
		"002078": 524,  // Runtop
		"002079": 4234, // Mikron
		"00207a": 4235, // WiSE
		"00207b": 421,  // Intel
		"00207c": 4236, // Autec
		"00207e": 4237, // Finecom
		"002080": 4238, // Synergy
		"002081": 738,  // Titan
		"002082": 4239, // Oneac
		"002083": 4240, // Presticom
		"002085": 1855, // Eaton
		"002086": 4241, // Microtech
		"002087": 4242, // Memotec
		"00208a": 4243, // Sonix
		"00208b": 4244, // Lapis
		"00208c": 3847, // Galaxy
		"00208d": 4245, // CMD
		"002091": 4246, // J125
		"002092": 4247, // Chess
		"002093": 4248, // Landings
		"002094": 4249, // Cubix
		"002095": 4250, // Riva
		"002096": 4251, // Invensys
		"002098": 2994, // Hectronic
		"00209b": 4252, // Ersat
		"00209f": 4253, // Mercury
		"0020a1": 4254, // Dovatron
		"0020a3": 1526, // Harmonic
		"0020a4": 4255, // Multipoint
		"0020a5": 4256, // API
		"0020a7": 4257, // Pairgain
		"0020a8": 4258, // Sast
		"0020ab": 43,   // Micro
		"0020ad": 4259, // Linq
		"0020af": 160,  // 3Com
		"0020b1": 4260, // Comtech
		"0020b6": 1610, // Agile
		"0020b9": 4261, // Metricom
		"0020bb": 4262, // ZAX
		"0020bf": 4263, // Aehr
		"0020c0": 4264, // Pulse
		"0020c1": 344,  // SAXA
		"0020c3": 4265, // Counter
		"0020c4": 4266, // Inet
		"0020c5": 4267, // Eagle
		"0020c6": 4268, // Nectec
		"0020c8": 4269, // Larscom
		"0020c9": 4270, // Victron
		"0020cb": 4271, // Pretec
		"0020cc": 1565, // Digital
		"0020cd": 4272, // Hybrid
		"0020d0": 4273, // Versalynx
		"0020d4": 16,   // Cabletron
		"0020d5": 4274, // Vipa
		"0020d6": 700,  // Breezecom
		"0020d8": 76,   // Nortel
		"0020d9": 2153, // Panasonic
		"0020db": 773,  // Xnet
		"0020dd": 4275, // Cybertec
		"0020e0": 2246, // Actiontec
		"0020e1": 4276, // Alamar
		"0020e4": 4277, // Hsing
		"0020e8": 4278, // Datatrek
		"0020e9": 4279, // Dantel
		"0020ea": 4280, // Efficient
		"0020ec": 4281, // Techware
		"0020ed": 1929, // Giga-Byte
		"0020ee": 4282, // Gtech
		"0020ef": 4283, // USC
		"0020f2": 11,   // Oracle
		"0020f3": 4284, // Raynet
		"0020f4": 4285, // Spectrix
		"0020f5": 4286, // Pandatel
		"0020f7": 4287, // Cyberdata
		"0020f8": 4288, // Carrera
		"0020f9": 4289, // Paralink
		"0020fa": 4290, // GDE
		"0020fb": 4291, // Octel
		"0020fc": 4292, // Matrox
		"0020fd": 4293, // ITV
		"0020ff": 4294, // Symmetrical
		"002100": 1450, // Gemtek
		"002101": 4295, // Aplicaciones
		"002102": 4296, // UpdateLogic
		"002103": 4297, // GHI
		"002104": 4298, // Gigaset
		"002107": 4299, // Seowonintech
		"002108": 457,  // Nokia
		"002109": 457,  // Nokia
		"00210a": 4300, // byd:sign
		"00210c": 4301, // Cymtec
		"00210f": 4302, // Cernium
		"002110": 4303, // Clearbox
		"002111": 4304, // Uniphone
		"002114": 4305, // Hylab
		"002116": 4306, // Transcon
		"002117": 4307, // Tellord
		"002118": 4308, // Athena
		"002119": 152,  // Samsung
		"00211a": 4309, // LInTech
		"00211b": 2,    // Cisco
		"00211c": 2,    // Cisco
		"00211d": 4310, // Dataline
		"00211e": 125,  // ARRIS
		"002120": 4311, // Sequel
		"002121": 4312, // VRmagic
		"002122": 4313, // Chip-pro
		"002124": 4314, // Optos
		"002127": 1595, // TP-Link
		"002128": 11,   // Oracle
		"002129": 1783, // Linksys
		"00212a": 4315, // Audiovox
		"00212d": 4316, // Scimolex
		"00212e": 4317, // dresden-elektronik
		"002131": 4318, // Blynke
		"002132": 4319, // Masterclock
		"002134": 4320, // Brandywine
		"002136": 125,  // ARRIS
		"002138": 4321, // Cepheid
		"002139": 4322, // Escherlogic
		"00213a": 4323, // Winchester
		"00213b": 4324, // Berkshire
		"00213c": 4325, // AliphCom
		"00213e": 2714, // TomTom
		"00213f": 4326, // A-Team
		"002140": 4327, // EN
		"002141": 4328, // Radlive
		"002143": 125,  // ARRIS
		"002145": 4329, // Semptian
		"002146": 3536, // Sanmina-SCI
		"002147": 1427, // Nintendo
		"00214c": 152,  // Samsung
		"00214f": 430,  // Alpsalpine
		"002150": 4330, // Eyeview
		"002151": 4331, // Millinet
		"002153": 4332, // SeaMicro
		"002154": 4333, // D-TACQ
		"002155": 2,    // Cisco
		"002156": 2,    // Cisco
		"002159": 826,  // Juniper
		"00215a": 302,  // HP
		"00215b": 4334, // SenseAnywhere
		"00215c": 421,  // Intel
		"00215d": 421,  // Intel
		"00215e": 372,  // IBM
		"00215f": 4335, // IHSE
		"002160": 4336, // Hidea
		"002161": 4337, // Yournet
		"002162": 76,   // Nortel
		"002163": 2544, // Askey
		"002165": 4338, // Presstek
		"002166": 4339, // NovAtel
		"002168": 4340, // iVeia
		"002169": 4341, // Prologix
		"00216a": 421,  // Intel
		"00216b": 421,  // Intel
		"00216c": 4342, // Odva
		"00216d": 4343, // Soltech
		"00216f": 4344, // SymCom
		"002170": 102,  // Dell
		"002179": 4345, // IOGEAR
		"00217a": 4346, // Sejin
		"00217b": 4347, // Bastec
		"00217c": 1939, // 2Wire
		"00217f": 4348, // Intraco
		"002180": 125,  // ARRIS
		"002182": 4349, // SandLinks
		"002185": 1812, // Micro-Star
		"002187": 4350, // Imacs
		"002188": 4351, // EMC
		"002189": 4352, // AppTech
		"00218b": 4353, // Wescon
		"00218c": 4354, // TopControl
		"00218e": 4355, // Mekics
		"002190": 4356, // Goliath
		"002191": 803,  // D-Link
		"002194": 4357, // Ping
		"002197": 1115, // Elitegroup
		"002199": 4358, // Vacon
		"00219b": 102,  // Dell
		"00219c": 4359, // Honeywld
		"00219d": 4360, // Adesys
		"00219e": 101,  // Sony
		"00219f": 4361, // Satel
		"0021a0": 2,    // Cisco
		"0021a1": 2,    // Cisco
		"0021a2": 4362, // EKE-Electronics
		"0021a3": 4363, // Micromint
		"0021a4": 4364, // Dbii
		"0021a6": 4365, // Videotec
		"0021a8": 4366, // Telephonics
		"0021aa": 457,  // Nokia
		"0021ab": 457,  // Nokia
		"0021b1": 1565, // Digital
		"0021b2": 4367, // Fiberblaze
		"0021b5": 4368, // Galvanic
		"0021b7": 613,  // Lexmark
		"0021b8": 4369, // Inphi
		"0021bc": 4370, // Zala
		"0021bd": 1427, // Nintendo
		"0021c2": 4371, // GL
		"0021c3": 4372, // CoRNELL
		"0021c4": 4373, // Consilium
		"0021c5": 4374, // 3DSP
		"0021c7": 4375, // Russound
		"0021c8": 4376, // LOHUIS
		"0021cc": 833,  // Flextronics
		"0021cd": 4377, // LiveTV
		"0021ce": 4378, // NTC-Metrotek
		"0021d1": 152,  // Samsung
		"0021d2": 152,  // Samsung
		"0021d5": 4379, // X2E
		"0021d7": 2,    // Cisco
		"0021d8": 2,    // Cisco
		"0021d9": 4380, // Sekonic
		"0021dd": 783,  // Northstar
		"0021e0": 4381, // CommAgility
		"0021e1": 76,   // Nortel
		"0021e3": 4382, // SerialTek
		"0021e4": 4383, // I-WIN
		"0021e7": 4384, // Informatics
		"0021e8": 2056, // Murata
		"0021e9": 545,  // Apple
		"0021eb": 4385, // ESP
		"0021ec": 4386, // Solutronic
		"0021ed": 4387, // Telegesis
		"0021ef": 4388, // Kapsys
		"0021f0": 4389, // EW3
		"0021f2": 4390, // EASY3CALL
		"0021f3": 4391, // Si14
		"0021f4": 4392, // INRange
		"0021f6": 11,   // Oracle
		"0021f8": 4393, // Enseo
		"0021f9": 4394, // WIRECOM
		"0021fa": 4395, // A4SP
		"0021fb": 869,  // LG
		"0021fc": 457,  // Nokia
		"0021fe": 457,  // Nokia
		"002200": 372,  // IBM
		"002201": 2114, // Aksys
		"002203": 4396, // Glensound
		"002204": 4397, // Koratek
		"002205": 4398, // WeLink
		"002206": 4399, // Cyberdyne
		"002208": 4400, // Certicom
		"00220a": 4401, // OnLive
		"00220c": 2,    // Cisco
		"00220d": 2,    // Cisco
		"00220f": 4402, // MoCA
		"002210": 125,  // ARRIS
		"002211": 4403, // Rohati
		"002212": 4404, // CAI
		"002213": 4405, // PCI
		"002215": 1806, // ASUS
		"002217": 4406, // Neat
		"002218": 3819, // Akamai
		"002219": 102,  // Dell
		"00221b": 4407, // Morega
		"00221c": 67,   // Private
		"00221d": 4408, // Freegene
		"00221f": 4409, // eSang
		"002220": 513,  // Mitac
		"002223": 4410, // TimeKeeping
		"002226": 4411, // Avaak
		"002227": 4412, // uv-electronic
		"002229": 4413, // Compumedics
		"00222a": 4414, // SoundEar
		"00222b": 4415, // Nucomm
		"00222c": 4416, // Ceton
		"00222d": 741,  // SMC
		"00222e": 4417, // maintech
		"002230": 4418, // FutureLogic
		"002231": 4419, // SMT&C
		"002234": 4420, // Corventis
		"002235": 4421, // Strukton
		"002238": 4422, // Logiplus
		"00223b": 4423, // Communication
		"00223d": 4424, // JumpGen
		"00223e": 4425, // IRTrans
		"00223f": 1368, // Netgear
		"002241": 545,  // Apple
		"002242": 4426, // Alacron
		"002243": 3004, // Azurewave
		"002247": 4427, // DAC
		"002248": 612,  // Microsoft
		"00224a": 4428, // Raylase
		"00224b": 4429, // Airtech
		"00224c": 1427, // Nintendo
		"00224d": 513,  // Mitac
		"00224e": 4430, // SEEnergy
		"00224f": 4431, // Byzoro
		"002251": 4432, // Lumasense
		"002253": 4433, // Entorian
		"002255": 2,    // Cisco
		"002256": 2,    // Cisco
		"002257": 160,  // 3Com
		"00225b": 4434, // Teradici
		"00225c": 4435, // Multimedia
		"00225e": 4436, // Uwin
		"00225f": 2873, // Liteon
		"002260": 4437, // AFREEY
		"002264": 302,  // HP
		"002265": 457,  // Nokia
		"002266": 457,  // Nokia
		"002267": 76,   // Nortel
		"00226a": 934,  // Honeywell
		"00226b": 1783, // Linksys
		"00226c": 4438, // LinkSprite
		"00226e": 4439, // Gowell
		"00226f": 4440, // 3onedata
		"002270": 4441, // ABK
		"002273": 4442, // Techway
		"002274": 4443, // FamilyPhone
		"002275": 2468, // Belkin
		"00227b": 591,  // Apogee
		"00227f": 2737, // Ruckus
		"002280": 4444, // A2B
		"002281": 4445, // Daintree
		"002283": 826,  // Juniper
		"002285": 4446, // Nomus
		"002286": 4447, // Astron
		"002288": 4448, // Sagrad
		"00228c": 4449, // Photon
		"00228d": 4450, // GBS
		"00228e": 4451, // TV-Numeric
		"00228f": 4452, // Cnrs
		"002290": 2,    // Cisco
		"002291": 2,    // Cisco
		"002292": 4453, // Cinetal
		"002293": 3031, // ZTE
		"002294": 2848, // Kyocera
		"002296": 4454, // LinoWave
		"002298": 101,  // Sony
		"002299": 4332, // SeaMicro
		"00229a": 4455, // Lastar
		"00229b": 4456, // AverLogic
		"00229c": 4457, // Verismo
		"00229d": 4458, // Pyung-HWA
		"0022a0": 4459, // Aptiv
		"0022a2": 4460, // Xtramus
		"0022a4": 1939, // 2Wire
		"0022a6": 101,  // Sony
		"0022a7": 4461, // Tyco
		"0022a8": 4462, // Ouman
		"0022a9": 869,  // LG
		"0022aa": 1427, // Nintendo
		"0022ad": 4463, // Telesis
		"0022ae": 4464, // Mattel
		"0022b0": 803,  // D-Link
		"0022b1": 4465, // Elbit
		"0022b2": 4466, // 4RF
		"0022b4": 125,  // ARRIS
		"0022b5": 4467, // Novita
		"0022b8": 4468, // Norcott
		"0022bb": 4469, // beyerdynamic
		"0022bd": 2,    // Cisco
		"0022be": 2,    // Cisco
		"0022c3": 4470, // Zeeport
		"0022c4": 4471, // epro
		"0022c5": 4472, // InfORSON
		"0022c6": 4473, // Sutus
		"0022c9": 4474, // Lenord
		"0022cb": 4475, // IONODES
		"0022cc": 4476, // SciLog
		"0022cd": 4477, // Ared
		"0022cf": 4478, // Planex
		"0022d0": 4479, // Polar
		"0022d3": 4480, // Hub-Tech
		"0022d4": 4481, // ComWorth
		"0022d6": 4482, // Cypak
		"0022d7": 1427, // Nintendo
		"0022d9": 4483, // Fortex
		"0022da": 4484, // Anatek
		"0022db": 4485, // Translogic
		"0022dd": 4486, // Protecta
		"0022e1": 4487, // ZORT
		"0022e3": 4488, // Amerigon
		"0022e4": 4489, // Apass
		"0022e5": 4490, // Fisher-Rosemount
		"0022e8": 4491, // Applition
		"0022e9": 4492, // ProVision
		"0022ea": 4493, // Rustelcom
		"0022ec": 4494, // Idealbt
		"0022ee": 4495, // Algo
		"0022ef": 4496, // iWDL
		"0022f1": 67,   // Private
		"0022f2": 4497, // SunPower
		"0022f3": 3205, // Sharp
		"0022f4": 4498, // AMPAK
		"0022f6": 4499, // Syracuse
		"0022f7": 4500, // Conceptronic
		"0022f8": 4501, // PIMA
		"0022f9": 4502, // Pollin
		"0022fa": 421,  // Intel
		"0022fb": 421,  // Intel
		"0022fc": 457,  // Nokia
		"0022fd": 457,  // Nokia
		"0022ff": 4503, // Nivis
		"002300": 4504, // Cayee
		"002301": 4505, // Witron
		"002304": 2,    // Cisco
		"002305": 2,    // Cisco
		"002306": 430,  // Alpsalpine
		"002308": 2639, // Arcadyan
		"002309": 4506, // Janam
		"00230a": 4507, // ARBURG
		"00230b": 125,  // ARRIS
		"00230c": 4508, // Clover
		"00230d": 76,   // Nortel
		"00230e": 4509, // Gorba
		"00230f": 4510, // Hirsch
		"002310": 4511, // LNC
		"002311": 4512, // Gloscom
		"002312": 545,  // Apple
		"002313": 4513, // Qool
		"002314": 421,  // Intel
		"002315": 421,  // Intel
		"002316": 4514, // Kisan
		"002317": 4515, // Lasercraft
		"002318": 35,   // Toshiba
		"002319": 4516, // Sielox
		"00231a": 4517, // ITF
		"00231c": 4518, // Fourier
		"00231d": 4519, // Deltacom
		"00231f": 4520, // Guangda
		"002320": 4521, // Nicira
		"002321": 4522, // Avitech
		"002323": 4523, // Zylin
		"002324": 2291, // G-PRO
		"002325": 4524, // IOLAN
		"002326": 4,    // Fujitsu
		"002327": 4525, // Shouyo
		"002329": 4526, // DDRdrive
		"00232b": 4527, // IRD
		"00232c": 4528, // Senticare
		"00232d": 4529, // SandForce
		"00232e": 4530, // Kedah
		"002330": 4531, // Dizipia
		"002331": 1427, // Nintendo
		"002332": 545,  // Apple
		"002333": 2,    // Cisco
		"002334": 2,    // Cisco
		"002335": 4532, // Linkflex
		"002338": 4533, // OJ-Electronics
		"002339": 152,  // Samsung
		"00233a": 152,  // Samsung
		"00233b": 4534, // C-Matic
		"00233c": 4535, // Alflex
		"00233d": 4536, // Laird
		"00233f": 4537, // Purechoice
		"002340": 4538, // MiXTelematics
		"002341": 887,  // Vanderbilt
		"002343": 4539, // TEM
		"002345": 101,  // Sony
		"002346": 4540, // Vestac
		"002348": 2048, // Sagemcom
		"00234a": 67,   // Private
		"00234b": 4541, // Inyuan
		"00234c": 4542, // KTC
		"002351": 1939, // 2Wire
		"002354": 1806, // ASUS
		"002357": 4543, // Pitronot
		"002358": 4544, // Systel
		"002359": 2079, // Benchmark
		"00235a": 358,  // Compal
		"00235b": 4545, // Gulfstream
		"00235c": 4546, // Aprius
		"00235d": 2,    // Cisco
		"00235e": 2,    // Cisco
		"002360": 4547, // Lookit
		"002361": 4548, // Unigen
		"002368": 768,  // Zebra
		"002369": 1783, // Linksys
		"00236a": 4549, // SmartRG
		"00236b": 4550, // Xembedded
		"00236c": 545,  // Apple
		"00236d": 4551, // ResMed
		"00236e": 4552, // Burster
		"002370": 822,  // Snell
		"002373": 4553, // GridIron
		"002374": 125,  // ARRIS
		"002375": 125,  // ARRIS
		"002376": 1341, // HTC
		"002377": 4554, // Isotek
		"00237a": 4555, // RIM
		"00237b": 4556, // Whdi
		"00237c": 4557, // Neotion
		"00237d": 302,  // HP
		"00237e": 3562, // Elster
		"00237f": 542,  // Plantronics
		"002380": 4558, // Nanoteq
		"002381": 4559, // Lengda
		"002383": 4560, // InMage
		"002384": 4561, // GGH
		"002385": 4562, // Antipode
		"002387": 4563, // ThinkFlood
		"00238a": 4564, // Ciena
		"00238b": 3071, // Quanta
		"00238c": 67,   // Private
		"002390": 4565, // Algolware
		"002391": 4566, // Maxian
		"002392": 4567, // Proteus
		"002393": 4568, // Ajinextek
		"002394": 4569, // Samjeon
		"002395": 125,  // ARRIS
		"002396": 4570, // Andes
		"002397": 2273, // Westell
		"002399": 152,  // Samsung
		"00239b": 3562, // Elster
		"00239c": 826,  // Juniper
		"00239d": 4571, // Mapower
		"0023a1": 176,  // Trend
		"0023a2": 125,  // ARRIS
		"0023a3": 125,  // ARRIS
		"0023a5": 4572, // SageTV
		"0023a6": 4573, // E-Mon
		"0023a8": 4574, // Marshall
		"0023aa": 4575, // HFR
		"0023ab": 2,    // Cisco
		"0023ac": 2,    // Cisco
		"0023ad": 4576, // Xmark
		"0023ae": 102,  // Dell
		"0023af": 125,  // ARRIS
		"0023b0": 4577, // CoMXION
		"0023b1": 4578, // Longcheer
		"0023b3": 4579, // Lyyn
		"0023b4": 457,  // Nokia
		"0023b5": 4580, // Ortana
		"0023b7": 4581, // Q-Light
		"0023ba": 4582, // Chroma
		"0023bc": 4583, // EQ-Sys
		"0023bf": 4584, // Mainpine
		"0023c0": 4585, // Broadway
		"0023c2": 4586, // SAMSUNG
		"0023c3": 4587, // LogMeIn
		"0023c6": 741,  // SMC
		"0023c7": 4588, // AVSystem
		"0023c8": 4589, // Team-R
		"0023cc": 1427, // Nintendo
		"0023cd": 1595, // TP-Link
		"0023cf": 4590, // Cummins-Allison
		"0023d1": 4591, // TRG
		"0023d2": 4592, // Inhand
		"0023d5": 4593, // WAREMA
		"0023d6": 152,  // Samsung
		"0023d7": 152,  // Samsung
		"0023d8": 4594, // Ball-It
		"0023d9": 4595, // Banner
		"0023db": 4596, // saxnet
		"0023dc": 4597, // Benein
		"0023de": 4598, // Ansync
		"0023df": 545,  // Apple
		"0023e3": 4599, // Microtronic
		"0023e4": 4600, // IPnect
		"0023e5": 4601, // IPaXiom
		"0023e7": 4602, // Hinke
		"0023e8": 4603, // Demco
		"0023e9": 289,  // F5
		"0023ea": 2,    // Cisco
		"0023eb": 2,    // Cisco
		"0023ec": 4604, // Algorithmix
		"0023ed": 125,  // ARRIS
		"0023ee": 125,  // ARRIS
		"0023f1": 101,  // Sony
		"0023f2": 4605, // TVLogic
		"0023f3": 4606, // Glocom
		"0023f4": 4607, // Masternaut
		"0023f6": 4608, // Softwell
		"0023f7": 67,   // Private
		"0023f8": 2692, // ZyXEL
		"0023f9": 4609, // Double-Take
		"0023fe": 4610, // Biodevices
		"002400": 76,   // Nortel
		"002401": 803,  // D-Link
		"002402": 4611, // Op-Tection
		"002403": 457,  // Nokia
		"002404": 457,  // Nokia
		"002406": 4612, // Pointmobile
		"002409": 4613, // Toro
		"00240b": 4614, // Virtual
		"00240c": 4615, // DELEC
		"00240d": 4616, // OnePath
		"002410": 4617, // NUETEQ
		"002411": 4618, // PharmaSmart
		"002412": 4619, // Benign
		"002413": 2,    // Cisco
		"002414": 2,    // Cisco
		"002419": 67,   // Private
		"00241b": 4620, // iWOW
		"00241c": 4621, // FuGang
		"00241d": 1929, // Giga-Byte
		"00241e": 1427, // Nintendo
		"00241f": 4622, // DCT-Delta
		"002420": 4623, // NetUP
		"002421": 1812, // Micro-Star
		"002423": 3004, // Azurewave
		"002427": 4624, // SSI
		"002428": 4625, // EnergyICT
		"00242e": 4626, // Datastrip
		"00242f": 4627, // Micron
		"002430": 4628, // Ruby
		"002431": 4629, // Uni-v
		"002432": 4630, // Neostar
		"002433": 430,  // Alpsalpine
		"002434": 4631, // Lectrosonics
		"002435": 4632, // Wide
		"002436": 545,  // Apple
		"002438": 90,   // Brocade
		"00243a": 4633, // Ludl
		"00243b": 4634, // CSSI
		"00243c": 4635, // S.A.A.A
		"00243f": 4636, // Storwize
		"002442": 4637, // Axona
		"002443": 76,   // Nortel
		"002444": 1427, // Nintendo
		"002445": 202,  // Adtran
		"002446": 4638, // MMB
		"002447": 4639, // Kaztek
		"00244a": 4640, // Voyant
		"00244b": 4641, // Perceptron
		"00244d": 4642, // Hokkaido
		"00244e": 4643, // RadChips
		"00244f": 4644, // Asantron
		"002450": 2,    // Cisco
		"002451": 2,    // Cisco
		"002452": 1655, // Silicon
		"002454": 152,  // Samsung
		"002455": 4645, // MuLogic
		"002456": 1939, // 2Wire
		"00245b": 4646, // Raidon
		"00245c": 4647, // Design-Com
		"00245e": 4648, // Hivision
		"002462": 4649, // Rayzone
		"002463": 4650, // Phybridge
		"002464": 4651, // Bridge
		"002465": 4652, // Elentec
		"002466": 3992, // Unitron
		"002467": 4653, // AOC
		"002468": 4654, // Sumavision
		"00246b": 4655, // Covia
		"00246c": 1685, // Aruba
		"00246d": 4656, // Weinzierl
		"00246f": 4657, // Onda
		"002473": 160,  // 3Com
		"002476": 4658, // TAP.tv
		"002477": 4659, // Tibbo
		"002478": 4660, // Mag
		"00247b": 2246, // Actiontec
		"00247c": 457,  // Nokia
		"00247d": 457,  // Nokia
		"00247f": 76,   // Nortel
		"002480": 4661, // Meteocontrol
		"002481": 302,  // HP
		"002482": 2737, // Ruckus
		"002483": 869,  // LG
		"002485": 4662, // ConteXtream
		"002486": 4663, // DesignArt
		"00248a": 4664, // Kaga
		"00248b": 4665, // Hybus
		"00248c": 1806, // ASUS
		"00248d": 101,  // Sony
		"00248f": 4666, // DO-Monix
		"002490": 152,  // Samsung
		"002491": 152,  // Samsung
		"002492": 687,  // Motorola
		"002493": 125,  // ARRIS
		"002495": 125,  // ARRIS
		"002496": 4667, // Ginzinger
		"002497": 2,    // Cisco
		"002498": 2,    // Cisco
		"002499": 4668, // Aquila
		"00249d": 4669, // NES
		"00249e": 4670, // ADC-Elektronik
		"0024a0": 125,  // ARRIS
		"0024a1": 125,  // ARRIS
		"0024a3": 4671, // Sonim
		"0024a4": 4672, // Siklu
		"0024a5": 1077, // Buffalo
		"0024aa": 4673, // Dycor
		"0024ab": 4674, // A7
		"0024ae": 4675, // Idemia
		"0024af": 1238, // Dish
		"0024b0": 4676, // Esab
		"0024b1": 4677, // Coulomb
		"0024b2": 1368, // Netgear
		"0024b3": 4678, // Graf-Syteco
		"0024b4": 4679, // ESCATRONIC
		"0024b5": 76,   // Nortel
		"0024b6": 732,  // Seagate
		"0024b7": 4680, // GridPoint
		"0024bb": 4681, // CENTRAL
		"0024bc": 4682, // HuRob
		"0024bd": 4683, // Hainzl
		"0024be": 101,  // Sony
		"0024bf": 4684, // Ciat
		"0024c1": 125,  // ARRIS
		"0024c2": 4685, // Asumo
		"0024c3": 2,    // Cisco
		"0024c4": 2,    // Cisco
		"0024c6": 4686, // Hager
		"0024c7": 4687, // Mobilarm
		"0024ca": 4688, // Tobii
		"0024cb": 4689, // Autonet
		"0024ce": 4690, // Exeltech
		"0024d1": 1142, // Thomson
		"0024d2": 2544, // Askey
		"0024d3": 4691, // QUALICA
		"0024d5": 4692, // Winward
		"0024d6": 421,  // Intel
		"0024d7": 421,  // Intel
		"0024d9": 4693, // BICOM
		"0024da": 4694, // Innovar
		"0024dc": 826,  // Juniper
		"0024dd": 4695, // Centrak
		"0024de": 4696, // GLOBAL
		"0024df": 4697, // Digitalbox
		"0024e0": 4698, // DS
		"0024e1": 4699, // Convey
		"0024e4": 4700, // Withings
		"0024e5": 4701, // Seer
		"0024e7": 4702, // Plaster
		"0024e8": 102,  // Dell
		"0024e9": 152,  // Samsung
		"0024eb": 4703, // ClearPath
		"0024ee": 4704, // Wynmax
		"0024ef": 101,  // Sony
		"0024f0": 4705, // Seanodes
		"0024f3": 1427, // Nintendo
		"0024f4": 4706, // Kaminario
		"0024f6": 4707, // Miyoshi
		"0024f7": 2,    // Cisco
		"0024f8": 2924, // Technical
		"0024f9": 2,    // Cisco
		"0024fb": 67,   // Private
		"0024fc": 4708, // QuoPin
		"0024fd": 3003, // Accedian
		"0024fe": 621,  // AVM
		"0024ff": 2025, // QLogic
		"002500": 545,  // Apple
		"002502": 4709, // NaturalPoint
		"002503": 372,  // IBM
		"002504": 4710, // Valiant
		"002507": 4711, // ASTAK
		"00250b": 4712, // Centrofactor
		"00250c": 4713, // Senet
		"002511": 1115, // Elitegroup
		"002512": 3031, // ZTE
		"002515": 3187, // SFR
		"002517": 4714, // Venntis
		"002519": 4715, // Viaas
		"00251c": 4716, // EDT
		"00251e": 4717, // Rotel
		"002521": 4718, // Logitek
		"002522": 4719, // ASRock
		"002523": 4720, // OCP
		"002524": 4721, // Lightcomm
		"002525": 4722, // CTERA
		"002526": 4723, // Genuine
		"002527": 4724, // Bitrode
		"00252c": 4725, // Entourage
		"00252d": 4726, // Kiryung
		"00252f": 4727, // Energy
		"002530": 4728, // Aetas
		"002533": 4729, // Wittenstein
		"002535": 4730, // Minimax
		"002537": 4731, // Runcom
		"002538": 152,  // Samsung
		"002539": 4732, // IfTA
		"00253a": 4733, // CEVA
		"00253c": 1939, // 2Wire
		"002540": 4734, // Quasar
		"002542": 4735, // Pittasoft
		"002543": 4736, // Moneytech
		"002544": 4737, // LoJack
		"002545": 2,    // Cisco
		"002546": 2,    // Cisco
		"002547": 457,  // Nokia
		"002548": 457,  // Nokia
		"002549": 4738, // Jeorich
		"00254a": 4739, // RingCube
		"00254b": 545,  // Apple
		"00254d": 4740, // Singapore
		"002550": 2093, // Riverbed
		"002551": 4741, // SE-Elektronic
		"002552": 4742, // VXi
		"002554": 4743, // Pixel8
		"002557": 2220, // BlackBerry
		"002558": 4744, // Mpedia
		"002559": 4745, // Syphan
		"00255a": 4746, // Tantalus
		"00255b": 4747, // CoachComm
		"00255c": 48,   // NEC
		"00255d": 4748, // Morningstar
		"00255f": 4749, // SenTec
		"002562": 4750, // interbro
		"002563": 4751, // Luxtera
		"002564": 102,  // Dell
		"002565": 4752, // Vizimax
		"002566": 152,  // Samsung
		"002567": 152,  // Samsung
		"002568": 3334, // Huawei
		"002569": 2048, // Sagemcom
		"002570": 4753, // Eastern
		"002572": 4754, // Nemo-Q
		"002573": 4755, // ST
		"002575": 4756, // FiberPlex
		"002576": 4757, // Neli
		"002577": 4758, // D-BOX
		"00257b": 4759, // STJ
		"00257f": 4760, // CallTechSolution
		"002581": 4761, // x-star
		"002582": 4762, // Maksat
		"002583": 2,    // Cisco
		"002584": 2,    // Cisco
		"002586": 1595, // TP-Link
		"002587": 4763, // Vitality
		"002588": 4764, // Genie
		"002589": 4765, // Hills
		"00258a": 4766, // Pole/Zero
		"00258b": 432,  // Mellanox
		"00258d": 4767, // Haier
		"002591": 4768, // NEXTEK
		"002594": 4769, // Eurodesign
		"002597": 4770, // Kalki
		"00259a": 4771, // CEStronics
		"00259c": 1783, // Linksys
		"00259d": 67,   // Private
		"00259e": 3334, // Huawei
		"00259f": 4772, // TechnoDigital
		"0025a0": 1427, // Nintendo
		"0025a1": 4773, // Enalasys
		"0025a7": 4774, // itron
		"0025a8": 63,   // Kontron
		"0025ac": 4775, // I-Tech
		"0025ae": 612,  // Microsoft
		"0025af": 4776, // CoMFILE
		"0025b0": 4777, // Schmartz
		"0025b1": 4778, // Maya-Creation
		"0025b3": 302,  // HP
		"0025b4": 2,    // Cisco
		"0025b5": 2,    // Cisco
		"0025b7": 4779, // Costar
		"0025b8": 1610, // Agile
		"0025b9": 4780, // Cypress
		"0025bb": 4781, // INNERINT
		"0025bc": 545,  // Apple
		"0025be": 4782, // Tektrap
		"0025c0": 4783, // ZillionTV
		"0025c2": 4784, // RingBell
		"0025c3": 4785, // 21168
		"0025c4": 2737, // Ruckus
		"0025c6": 4786, // kasercorp
		"0025c7": 4787, // altek
		"0025c8": 4788, // S-Access
		"0025ca": 1626, // LS
		"0025ce": 4789, // InnerSpace
		"0025cf": 457,  // Nokia
		"0025d0": 457,  // Nokia
		"0025d2": 4790, // InpegVision
		"0025d3": 3004, // Azurewave
		"0025d5": 4791, // Robonica
		"0025d6": 4792, // Kroger
		"0025d7": 4793, // Cedo
		"0025d9": 4794, // DataFab
		"0025db": 1011, // ATI
		"0025de": 4795, // Probits
		"0025df": 67,   // Private
		"0025e2": 4796, // Everspring
		"0025e3": 4797, // Hanshinit
		"0025e4": 4798, // OMNI-WiFi
		"0025e5": 869,  // LG
		"0025e7": 101,  // Sony
		"0025e8": 4799, // Idaho
		"0025ea": 4800, // Iphion
		"0025ec": 4801, // Humanware
		"0025ed": 4802, // NuVo
		"0025ee": 4803, // Avtex
		"0025ef": 4804, // I-TEC
		"0025f0": 4805, // Suga
		"0025f1": 125,  // ARRIS
		"0025f2": 125,  // ARRIS
		"0025f6": 4806, // netTALK.com
		"0025f9": 4807, // GMK
		"0025fe": 4808, // Pilot
		"002601": 4809, // Cutera
		"002605": 4810, // CC
		"002606": 4811, // RAUMFELD
		"002607": 4812, // Enabling
		"002608": 545,  // Apple
		"002609": 4813, // Phyllis
		"00260a": 2,    // Cisco
		"00260b": 2,    // Cisco
		"00260c": 4814, // Dataram
		"00260d": 4253, // Mercury
		"00260e": 4815, // Ablaze
		"00260f": 4816, // Linn
		"002610": 4817, // Apacewave
		"002611": 4818, // Licera
		"002614": 4819, // Ktnf
		"002615": 4820, // Teracom
		"002616": 4821, // Rosemount
		"002618": 1806, // ASUS
		"002619": 4822, // FRC
		"00261c": 4823, // Neovia
		"002620": 4824, // ISGUS
		"002621": 4825, // InteliCloud
		"002622": 358,  // Compal
		"002623": 4826, // JRD
		"002624": 1142, // Thomson
		"002625": 4827, // MediaSputnik
		"002627": 4828, // Truesell
		"00262a": 4829, // Proxense
		"00262b": 4830, // Wongs
		"00262d": 1592, // Wistron
		"002631": 4831, // Commtact
		"002634": 4832, // Infineta
		"002635": 4833, // Bluetechnix
		"002636": 125,  // ARRIS
		"002637": 152,  // Samsung
		"002639": 4834, // T.M
		"00263a": 4835, // Digitec
		"00263b": 4836, // Onbnetech
		"00263c": 4837, // Bachmann
		"00263d": 4838, // MIA
		"00263e": 1612, // Trapeze
		"00263f": 4839, // LIOS
		"002641": 125,  // ARRIS
		"002642": 125,  // ARRIS
		"002643": 430,  // Alpsalpine
		"002647": 4840, // WFE
		"002648": 3905, // Emitech
		"00264a": 545,  // Apple
		"00264d": 2639, // Arcadyan
		"00264e": 4841, // r2p
		"002650": 1939, // 2Wire
		"002651": 2,    // Cisco
		"002652": 2,    // Cisco
		"002653": 4842, // DaySequerra
		"002654": 160,  // 3Com
		"002655": 302,  // HP
		"002656": 4843, // Sansonic
		"002658": 4844, // T-Platforms
		"002659": 1427, // Nintendo
		"00265a": 803,  // D-Link
		"00265b": 870,  // Hitron
		"00265d": 152,  // Samsung
		"00265f": 152,  // Samsung
		"002660": 4845, // Logiways
		"002661": 4846, // Irumtek
		"002662": 2246, // Actiontec
		"002665": 4847, // ProtectedLogic
		"002666": 1252, // EFM
		"002667": 4848, // Carecom
		"002668": 457,  // Nokia
		"002669": 457,  // Nokia
		"00266a": 4849, // Essensium
		"00266c": 3978, // Inventec
		"00266d": 4850, // MobileAccess
		"00266e": 4851, // Nissho-denki
		"00266f": 4852, // Coordiwise
		"002671": 4853, // AUTOVISION
		"002673": 75,   // Ricoh
		"002675": 4854, // Aztech
		"002676": 4855, // CoMMidt
		"002677": 4856, // Deif
		"002679": 4857, // Euphonic
		"00267c": 4858, // Metz-Werke
		"00267e": 2562, // Parrot
		"00267f": 4859, // Zenterio
		"002680": 4860, // SIL3
		"002681": 4861, // Interspiro
		"002682": 1450, // Gemtek
		"002683": 4862, // Ajoho
		"002688": 826,  // Juniper
		"00268c": 4863, // StarLeaf
		"00268e": 2309, // Alta
		"00268f": 4864, // MTA
		"002691": 2048, // Sagemcom
		"002693": 4865, // QVidium
		"002694": 4866, // Senscient
		"002696": 4867, // NOOLIX
		"002697": 2236, // Alpha
		"002698": 2,    // Cisco
		"002699": 2,    // Cisco
		"00269b": 4868, // SOKRAT
		"00269d": 4869, // M2Mnet
		"00269e": 3071, // Quanta
		"00269f": 67,   // Private
		"0026a0": 4870, // moblic
		"0026a1": 4871, // Megger
		"0026a2": 4872, // Instrumentation
		"0026a5": 4873, // Microrobot
		"0026a6": 4874, // Trixell
		"0026a9": 4875, // Strong
		"0026ab": 46,   // Epson
		"0026ad": 4876, // Arada
		"0026af": 4877, // Duelco
		"0026b0": 545,  // Apple
		"0026b2": 4878, // Setrix
		"0026b3": 4879, // Thales
		"0026b4": 4880, // Ford
		"0026b6": 2544, // Askey
		"0026b7": 4881, // Kingston
		"0026b8": 2246, // Actiontec
		"0026b9": 102,  // Dell
		"0026ba": 125,  // ARRIS
		"0026bb": 545,  // Apple
		"0026c0": 4882, // EnergyHub
		"0026c1": 4883, // Artray
		"0026c2": 4884, // SCDI
		"0026c3": 4885, // Insightek
		"0026c6": 421,  // Intel
		"0026c7": 421,  // Intel
		"0026c9": 4886, // Proventix
		"0026ca": 2,    // Cisco
		"0026cb": 2,    // Cisco
		"0026cc": 457,  // Nokia
		"0026cd": 4887, // PurpleComm
		"0026d0": 4888, // Semihalf
		"0026d2": 4889, // Pcube
		"0026d4": 4890, // IRCA
		"0026d9": 125,  // ARRIS
		"0026df": 4891, // TaiDoc
		"0026e0": 4892, // Asiteq
		"0026e2": 869,  // LG
		"0026e3": 4893, // DTI
		"0026e6": 4894, // Visionhitech
		"0026e8": 2056, // Murata
		"0026e9": 4895, // SP
		"0026ea": 4896, // Cheerchip
		"0026ed": 3031, // ZTE
		"0026ee": 4897, // TKM
		"0026f0": 4898, // cTrixs
		"0026f2": 1368, // Netgear
		"0026f3": 741,  // SMC
		"0026f4": 4899, // Nesslab
		"0026f5": 4900, // XRPLUS
		"0026f7": 4901, // Nivetti
		"0026fa": 4902, // BandRich
		"0026fc": 4903, // AcSiP
		"0026fe": 4904, // MKD
		"0026ff": 2220, // BlackBerry
		"002701": 4905, // IncOstartec
		"002702": 4906, // SolarEdge
		"002703": 1406, // Testech
		"002705": 4907, // Sectronic
		"002706": 4908, // Yoisys
		"002709": 1427, // Nintendo
		"00270b": 4909, // Adura
		"00270c": 2,    // Cisco
		"00270d": 2,    // Cisco
		"00270e": 421,  // Intel
		"00270f": 4910, // Envisionnovation
		"002710": 421,  // Intel
		"002711": 4911, // LanPro
		"002712": 4912, // MaxVision
		"002714": 4913, // Grainmustards
		"002716": 4914, // Adachi-Syokai
		"002719": 1595, // TP-Link
		"00271a": 4915, // Geenovo
		"00271c": 4253, // Mercury
		"00271e": 4916, // Xagyl
		"00271f": 4917, // MIPRO
		"002722": 2978, // Ubiquiti
		"002790": 2,    // Cisco
		"0027e3": 2,    // Cisco
		"0027f8": 90,   // Brocade
		"00289f": 4329, // Semptian
		"0028f8": 421,  // Intel
		"0029c2": 2,    // Cisco
		"002a10": 2,    // Cisco
		"002a6a": 2,    // Cisco
		"002aaf": 4918, // LARsys-Automation
		"002b67": 4919, // LCFC
		"002cc8": 2,    // Cisco
		"002d76": 4920, // TITECH
		"002db3": 4498, // AMPAK
		"002ec7": 3334, // Huawei
		"002f5c": 2,    // Cisco
		"002fd9": 4921, // Fiberhome
		"003000": 4922, // Allwell
		"003001": 4923, // SMP
		"003002": 4924, // Expand
		"003003": 4925, // Phasys
		"003004": 4926, // Leadtek
		"003006": 4927, // Superpower
		"003007": 4928, // Opti
		"003009": 4929, // Tachion
		"00300a": 4854, // Aztech
		"00300b": 4930, // mPHASE
		"00300c": 4931, // Congruency
		"00300d": 4932, // MMC
		"003010": 4933, // Visionetics
		"003011": 837,  // HMS
		"003012": 1565, // Digital
		"003013": 48,   // NEC
		"003014": 4934, // Divio
		"003016": 4935, // Ishida
		"003019": 2,    // Cisco
		"00301a": 4936, // Smartbridges
		"00301b": 4937, // Shuttle
		"00301d": 4938, // Skystream
		"00301e": 160,  // 3Com
		"00301f": 1484, // Optical
		"003020": 4939, // TSI
		"003021": 4277, // Hsing
		"003023": 1675, // Cogent
		"003024": 2,    // Cisco
		"003025": 4940, // Checkout
		"003027": 4941, // Kerbango
		"003029": 4942, // Opicom
		"00302b": 4943, // Inalp
		"00302c": 4944, // Sylantro
		"003030": 4945, // Harmonix
		"003031": 4946, // Lightwave
		"003032": 4947, // MagicRam
		"003034": 4948, // SET
		"003035": 4949, // Corning
		"003038": 4950, // XCP
		"00303a": 4951, // Maatel
		"00303b": 4952, // PowerCom
		"00303c": 4953, // Onnto
		"00303d": 4954, // IVA
		"00303e": 4955, // Radcom
		"00303f": 4956, // TurboComm
		"003040": 2,    // Cisco
		"003043": 4957, // Idream
		"003044": 4958, // CradlePoint
		"003046": 4959, // Controlled
		"003049": 4960, // Bryant
		"00304b": 4961, // Orbacom
		"00304c": 4962, // Appian
		"00304d": 4963, // ESI
		"00304f": 4964, // PLANET
		"003050": 4965, // Versa
		"003052": 4966, // Elastic
		"003053": 4967, // Basler
		"003054": 3797, // Castlenet
		"003056": 837,  // HMS
		"003057": 4968, // QTelNet
		"003059": 63,   // Kontron
		"00305a": 4969, // Telgen
		"00305b": 4970, // Toko
		"00305c": 4971, // SMAR
		"00305d": 4972, // Digitra
		"00305f": 4973, // Hasselblad
		"003060": 4974, // Powerfile
		"003061": 4975, // MobyTEL
		"003063": 4976, // Santera
		"003064": 4977, // Adlink
		"003065": 545,  // Apple
		"003066": 4978, // RFM
		"003068": 1654, // Cybernetics
		"003069": 4979, // Impacct
		"00306b": 4980, // Cmos
		"00306c": 4981, // Hitex
		"00306d": 827,  // Lucent
		"00306e": 302,  // HP
		"00306f": 4982, // Seyeon
		"003070": 4983, // 1Net
		"003071": 2,    // Cisco
		"003072": 4984, // Intellibyte
		"003074": 4985, // Equiinet
		"003075": 4986, // Adtech
		"003076": 4987, // Akamba
		"003077": 4988, // Onprem
		"003078": 2,    // Cisco
		"003079": 4989, // Cqos
		"00307a": 587,  // Advanced
		"00307b": 2,    // Cisco
		"00307c": 4990, // Adid
		"00307e": 4991, // Redflex
		"00307f": 4992, // Irlan
		"003080": 2,    // Cisco
		"003083": 4993, // Ivron
		"003085": 2,    // Cisco
		"003088": 306,  // Ericsson
		"00308b": 4994, // Brix
		"00308c": 2083, // Quantum
		"00308d": 4995, // Pinnacle
		"00308f": 4996, // MICRILOR
		"003090": 4997, // Cyra
		"003092": 63,   // Kontron
		"003093": 4998, // Sonnet
		"003094": 2,    // Cisco
		"003096": 2,    // Cisco
		"00309b": 4999, // Smartware
		"00309e": 5000, // Workbit
		"00309f": 5001, // Amber
		"0030a1": 5002, // WEBGATE
		"0030a2": 5003, // Lightner
		"0030a3": 2,    // Cisco
		"0030a6": 5004, // Vianet
		"0030a7": 5005, // Schweitzer
		"0030a8": 5006, // OL'E
		"0030a9": 5007, // Netiverse
		"0030ab": 957,  // Delta
		"0030ad": 5008, // Shanghai
		"0030af": 934,  // Honeywell
		"0030b0": 5009, // Convergenet
		"0030b1": 5010, // TrunkNet
		"0030b4": 5011, // Intersil
		"0030b6": 2,    // Cisco
		"0030b7": 5012, // Teletrol
		"0030b8": 5013, // RiverDelta
		"0030b9": 5014, // Ectel
		"0030bb": 5015, // CacheFlow
		"0030bc": 5016, // Optronic
		"0030be": 5017, // City-Net
		"0030bf": 5018, // Multidata
		"0030c0": 5019, // Lara
		"0030c1": 302,  // HP
		"0030c2": 5020, // Comone
		"0030c6": 1998, // Control
		"0030c7": 5021, // Macromate
		"0030c9": 5022, // LuxN
		"0030cc": 5023, // Tenor
		"0030cd": 5024, // Conexant
		"0030ce": 5025, // Zaffire
		"0030cf": 5026, // TWO
		"0030d0": 5027, // Tellabs
		"0030d1": 5028, // Inova
		"0030d2": 201,  // WIN
		"0030d3": 653,  // Agilent
		"0030d4": 5029, // AAE
		"0030d5": 5030, // DResearch
		"0030d7": 428,  // Innovative
		"0030d8": 5031, // Sitek
		"0030d9": 5032, // Datacore
		"0030da": 3869, // Comtrend
		"0030db": 5033, // Mindready
		"0030dc": 5034, // Rightech
		"0030dd": 5035, // Indigita
		"0030e2": 5036, // Garnet
		"0030e3": 5037, // Sedona
		"0030e8": 5038, // Ensim
		"0030ea": 5039, // TeraForce
		"0030eb": 5040, // Turbonet
		"0030ec": 5041, // Borgardt
		"0030ee": 5042, // DSG
		"0030ef": 5043, // Neon
		"0030f0": 5044, // Uniform
		"0030f1": 145,  // Accton
		"0030f2": 2,    // Cisco
		"0030f4": 5045, // Stardot
		"0030f6": 5046, // Securelogix
		"0030f7": 5047, // Ramix
		"0030f8": 5048, // Dynapro
		"0030f9": 5049, // Sollae
		"0030fa": 5050, // Telica
		"0030fb": 5051, // AZS
		"0030fc": 5052, // Terawave
		"0030fe": 5053, // DSA
		"0030ff": 4794, // DataFab
		"003146": 826,  // Juniper
		"003192": 1595, // TP-Link
		"003217": 2,    // Cisco
		"00323a": 5054, // so-logic
		"00336c": 5055, // SynapSense
		"0034da": 869,  // LG
		"0034f1": 5056, // Radicom
		"0034fe": 3334, // Huawei
		"00351a": 2,    // Cisco
		"003532": 5057, // Electro-Metrics
		"003676": 125,  // ARRIS
		"0036fe": 2834, // SuperVision
		"00376d": 2056, // Murata
		"0037b7": 2048, // Sagemcom
		"0038df": 2,    // Cisco
		"003a7d": 2,    // Cisco
		"003a98": 2,    // Cisco
		"003a99": 2,    // Cisco
		"003a9a": 2,    // Cisco
		"003a9b": 2,    // Cisco
		"003a9c": 2,    // Cisco
		"003aaf": 5058, // BlueBit
		"003c10": 2,    // Cisco
		"003c84": 1655, // Silicon
		"003cc5": 5059, // WONWOO
		"003d41": 5060, // Hatteland
		"003de1": 3334, // Huawei
		"003de8": 869,  // LG
		"003ee1": 545,  // Apple
		"004002": 5061, // Perle
		"004004": 5062, // ICM
		"004005": 5063, // ANI
		"004006": 2970, // Sampo
		"00400a": 5064, // Pivotal
		"00400b": 2,    // Cisco
		"00400e": 4242, // Memotec
		"00400f": 5065, // Datacom
		"004010": 5066, // Sonic
		"004012": 5067, // Windata
		"004014": 5068, // Comsoft
		"004018": 5069, // Adobe
		"004019": 5070, // Aeon
		"00401b": 5071, // Printer
		"00401c": 5072, // AST
		"00401d": 5073, // Invisible
		"00401e": 5074, // ICC
		"00401f": 5075, // Colorgraph
		"004020": 5076, // CommScope
		"004022": 5077, // Klever
		"004023": 5078, // Logic
		"004024": 5079, // Compac
		"004026": 1077, // Buffalo
		"004028": 5080, // Netcomm
		"004029": 5081, // Compex
		"00402b": 5082, // Trigem
		"00402e": 263,  // Precision
		"004030": 5083, // GK
		"004032": 1565, // Digital
		"004033": 5084, // Addtron
		"004034": 5085, // Bustek
		"004035": 5086, // Opcom
		"004036": 5087, // Minim
		"004037": 5088, // SEA-Ilan
		"00403a": 5089, // Impact
		"00403b": 5090, // Synerjet
		"00403c": 5091, // Forks
		"00403d": 5092, // Teradata
		"00403f": 5093, // Ssangyong
		"004041": 1698, // Fujikura
		"004042": 5094, // N.A.T
		"004044": 5095, // Qnix
		"004045": 184,  // Twinhead
		"004046": 5096, // UDC
		"00404b": 1493, // Maple
		"00404c": 5097, // Hypertec
		"00404e": 5098, // Fluent
		"004050": 5099, // Ironicsorporated
		"004052": 5100, // Star
		"004053": 5101, // Ampro
		"004055": 5102, // Metronix
		"004058": 5103, // UKG
		"00405b": 5104, // Funasset
		"00405c": 5105, // Future
		"00405d": 5106, // Star-TEK
		"00405f": 5107, // AFE
		"004060": 5108, // Comendec
		"004061": 5109, // Datatech
		"004062": 5110, // E-Systems./Garland
		"004063": 4162, // VIA
		"004066": 5111, // APRESIA
		"004067": 5112, // Omnibyte
		"004068": 5113, // Extended
		"004069": 5114, // Lemcom
		"00406b": 5115, // Sysgen
		"00406c": 5116, // Copernique
		"00406d": 5117, // Lanco
		"00406e": 5118, // Corollary
		"00406f": 5119, // Sync
		"004070": 5120, // Interware
		"004071": 5121, // ATM
		"004077": 5122, // Maxton
		"004079": 5123, // Juko
		"00407c": 5124, // Qume
		"00407d": 5125, // Extension
		"00407e": 4225, // Evergreen
		"00407f": 3723, // FLIR
		"004080": 5126, // Athenix
		"004084": 934,  // Honeywell
		"004087": 5127, // Ubitrex
		"004088": 5128, // Mobius
		"004089": 5129, // Meidensha
		"00408b": 5130, // Raylan
		"00408c": 5131, // Axis
		"004090": 5132, // Ansel
		"004092": 2871, // ASP
		"004093": 5133, // Paxdata
		"004094": 5134, // Shographics
		"004096": 2,    // Cisco
		"004098": 5135, // Dressler
		"004099": 5136, // Newgen
		"00409b": 5137, // HAL
		"00409c": 5138, // Transware
		"00409d": 5139, // DigiBoard
		"00409e": 5140, // Concurrent
		"00409f": 18,   // Telco
		"0040a0": 5141, // Goldstar
		"0040a2": 5142, // Kingstar
		"0040a3": 5143, // Microunity
		"0040a4": 5144, // Rose
		"0040a6": 68,   // Cray
		"0040a8": 5145, // IMF
		"0040a9": 5065, // Datacom
		"0040af": 1565, // Digital
		"0040b0": 5146, // Bytex
		"0040b1": 5147, // Codonics
		"0040b2": 5148, // Systemforschung
		"0040b3": 5149, // ParTech
		"0040b5": 5150, // Video
		"0040b6": 5151, // Computerm
		"0040b7": 5152, // Stealth
		"0040b9": 5153, // Macq
		"0040ba": 2652, // Alliant
		"0040bc": 5154, // Algorithmics
		"0040bd": 5155, // Starlight
		"0040be": 1041, // Boeing
		"0040c5": 5156, // Micom
		"0040c6": 5157, // Fibernet
		"0040c7": 4628, // Ruby
		"0040c8": 5158, // Milan
		"0040c9": 5159, // Ncube
		"0040cb": 5160, // Lanwan
		"0040ce": 5161, // NET-Source
		"0040d0": 513,  // Mitac
		"0040d2": 5162, // Pagine
		"0040d3": 5163, // Kimpsion
		"0040da": 5164, // Telspec
		"0040dc": 5165, // Tritec
		"0040dd": 5166, // Hong
		"0040df": 5167, // Digalog
		"0040e0": 5168, // Atomwide
		"0040e1": 5169, // Marner
		"0040e3": 5170, // Quin
		"0040e4": 5171, // E-M
		"0040e5": 5172, // Sybus
		"0040e6": 5173, // C.A.E.N
		"0040e9": 5174, // Accord
		"0040ee": 5175, // Optimem
		"0040ef": 5176, // Hypercom
		"0040f0": 5177, // MicroBrain
		"0040f1": 5178, // Chuo
		"0040f3": 5179, // Netcor
		"0040f4": 3383, // Cameo
		"0040f6": 5180, // Katron
		"0040f7": 656,  // Polaroid
		"0040f9": 5181, // Combinet
		"0040fa": 5182, // Microboards
		"0040fb": 5183, // Cascade
		"0040fd": 5184, // LXE
		"0040fe": 5185, // Symplex
		"0040ff": 5186, // Telebit
		"0041d2": 2,    // Cisco
		"004238": 421,  // Intel
		"004252": 5187, // RLX
		"00425a": 2,    // Cisco
		"004268": 2,    // Cisco
		"004279": 3939, // Sunitec
		"00451d": 2,    // Cisco
		"0045e2": 189,  // CyberTAN
		"00464b": 3334, // Huawei
		"004a77": 3031, // ZTE
		"004e01": 102,  // Dell
		"004e35": 302,  // HP
		"004f1a": 3334, // Huawei
		"005000": 5188, // Nexo
		"005001": 5189, // Yamashita
		"005002": 5190, // Omnisec
		"005003": 5191, // Xrite
		"005004": 160,  // 3Com
		"005006": 2379, // TAC
		"005007": 300,  // Siemens
		"00500a": 5192, // Iris
		"00500b": 2,    // Cisco
		"00500c": 5193, // e-Tek
		"00500e": 5194, // Chromatis
		"00500f": 2,    // Cisco
		"005012": 5195, // CBL
		"005014": 2,    // Cisco
		"005018": 5196, // AMIT
		"00501a": 5197, // IQinVision
		"00501c": 5198, // Jatom
		"00501f": 5199, // MRG
		"005020": 5200, // Mediastar
		"005021": 5201, // EIS
		"005022": 5202, // Zonet
		"005024": 5203, // Navic
		"005026": 5204, // Cosystems
		"005027": 5205, // Genicom
		"005028": 5206, // Aval
		"00502a": 2,    // Cisco
		"00502b": 5207, // Genrad
		"00502c": 5208, // Soyo
		"00502d": 5209, // Accel
		"00502e": 5210, // Cambex
		"00502f": 5211, // TollBridge
		"005031": 5212, // Aeroflex
		"005032": 5213, // Picazo
		"005033": 5214, // Mayan
		"005036": 5215, // Netcam
		"005037": 5216, // Koga
		"005039": 5217, // Mariner
		"00503a": 5218, // Datong
		"00503b": 5219, // Mediafire
		"00503e": 2,    // Cisco
		"005040": 2153, // Panasonic
		"005041": 5220, // Coretronic
		"005044": 5221, // Asaca
		"005045": 5222, // Rioworks
		"005046": 5223, // Menicx
		"005047": 67,   // Private
		"005048": 5224, // Infolibria
		"005049": 859,  // Arbor
		"00504d": 1205, // Tokyo
		"00504f": 5225, // Olencom
		"005050": 2,    // Cisco
		"005052": 5226, // Tiara
		"005053": 2,    // Cisco
		"005054": 2,    // Cisco
		"005055": 5227, // Doms
		"005056": 809,  // VMware
		"005058": 5228, // Sangoma
		"005059": 5229, // iBAHN
		"00505c": 5230, // Tundo
		"005062": 5231, // Kouwell
		"005064": 5232, // CAE
		"005065": 3511, // TDK-Lambda
		"005067": 5233, // Aerocomm
		"005068": 3690, // Electronic
		"005069": 5234, // PixStream
		"00506a": 5235, // Edeva
		"00506b": 5236, // SPX-Ateg
		"00506c": 5237, // Beijer
		"00506d": 5238, // Videojet
		"00506e": 5239, // Corder
		"00506f": 5240, // G-Connect
		"005070": 5241, // Chaintech
		"005071": 5242, // Aiwa
		"005072": 5243, // Corvis
		"005073": 2,    // Cisco
		"005075": 5244, // Kestrel
		"005076": 372,  // IBM
		"005077": 5245, // Prolific
		"005079": 67,   // Private
		"00507a": 486,  // Xpeed
		"00507b": 5246, // Merlot
		"00507c": 5247, // Videocon
		"00507d": 5248, // IFP
		"00507e": 5249, // Newer
		"00507f": 3922, // DrayTek
		"005080": 2,    // Cisco
		"005082": 5250, // Foresson
		"005083": 5251, // Gilbarco
		"005084": 2083, // Quantum
		"005086": 5252, // Telkom
		"005088": 5253, // Amano
		"00508b": 302,  // HP
		"00508c": 5254, // RSI
		"00508d": 5255, // Abit
		"00508e": 5256, // Optimation
		"005090": 5257, // Dctri
		"005091": 5258, // Netaccess
		"005093": 1041, // Boeing
		"005094": 125,  // ARRIS
		"005095": 5259, // Peracom
		"005096": 5260, // Salix
		"005098": 5261, // Globaloop
		"005099": 160,  // 3Com
		"00509a": 5262, // TAG
		"00509b": 5263, // Switchcore
		"00509c": 5264, // Beta
		"00509f": 5265, // Horizon
		"0050a0": 957,  // Delta
		"0050a2": 2,    // Cisco
		"0050a3": 5266, // TransMedia
		"0050a4": 5267, // IO
		"0050a6": 5268, // Optronics
		"0050a7": 2,    // Cisco
		"0050a8": 5269, // OpenCon
		"0050ab": 5270, // NALTEC
		"0050ac": 1493, // Maple
		"0050ae": 5271, // FDK
		"0050af": 5272, // Intergon
		"0050b2": 5273, // BRODEL
		"0050b3": 5274, // Voiceboard
		"0050b7": 5275, // Boser
		"0050b8": 5028, // Inova
		"0050b9": 5276, // Xitron
		"0050ba": 803,  // D-Link
		"0050bb": 2363, // CMS
		"0050bd": 2,    // Cisco
		"0050bf": 1709, // Metalligence
		"0050c0": 5277, // Gatan
		"0050c1": 5278, // Gemflex
		"0050c4": 5279, // IMD
		"0050c5": 5280, // ADS
		"0050c7": 67,   // Private
		"0050c8": 5281, // Addonics
		"0050cb": 5282, // Jetter
		"0050cd": 5283, // Digianswer
		"0050ce": 869,  // LG
		"0050d0": 5284, // Minerva
		"0050d1": 2,    // Cisco
		"0050d2": 5285, // CMC
		"0050d5": 5286, // AD
		"0050d7": 5287, // Telstrat
		"0050d8": 5288, // Unicorn
		"0050da": 160,  // 3Com
		"0050de": 5289, // Signum
		"0050df": 5290, // AirFiber
		"0050e1": 5291, // NS
		"0050e2": 2,    // Cisco
		"0050e3": 125,  // ARRIS
		"0050e4": 545,  // Apple
		"0050e6": 5292, // Hakusan
		"0050e8": 5293, // Nomadix
		"0050ea": 5294, // XEL
		"0050eb": 5295, // Alpha-TOP
		"0050ec": 5296, // Olicom
		"0050ed": 5297, // Anda
		"0050f0": 2,    // Cisco
		"0050f1": 5298, // Maxlinear
		"0050f2": 612,  // Microsoft
		"0050f4": 5299, // Sigmatek
		"0050f6": 5300, // PAN-International
		"0050f8": 5301, // Entrega
		"0050f9": 5302, // Sensormatic
		"0050fa": 5303, // Oxtel
		"0050fb": 5304, // VSK
		"0050fc": 115,  // Edimax
		"0050fd": 5305, // Visioncomm
		"0050ff": 5306, // Hakko
		"0051ed": 869,  // LG
		"00549f": 620,  // Avaya
		"0054bd": 5307, // Swelaser
		"00562b": 2,    // Cisco
		"0056cd": 545,  // Apple
		"0057c1": 869,  // LG
		"0057d2": 2,    // Cisco
		"005907": 2664, // Lenovo
		"0059dc": 2,    // Cisco
		"005a13": 3334, // Huawei
		"005b94": 545,  // Apple
		"005d03": 1498, // Xilinx
		"005d73": 2,    // Cisco
		"005f67": 1595, // TP-Link
		"005f86": 2,    // Cisco
		"005fbf": 35,   // Toshiba
		"006000": 5308, // Xycom
		"006001": 5309, // InnoSys
		"006006": 5310, // Sotec
		"006008": 160,  // 3Com
		"006009": 2,    // Cisco
		"00600a": 5311, // Sord
		"00600b": 5312, // LOGWARE
		"00600c": 5313, // Eurotech
		"00600e": 5314, // Wavenet
		"00600f": 2273, // Westell
		"006014": 5315, // Edec
		"006015": 5316, // NET2NET
		"006016": 5317, // Clariion
		"006017": 5318, // Tokimec
		"00601b": 5319, // Mesa
		"00601c": 5320, // Telxon
		"00601d": 827,  // Lucent
		"00601e": 5321, // Softlab
		"00601f": 5322, // Stallion
		"006021": 5323, // DSC
		"006022": 5324, // Vicom
		"006024": 5325, // Gradient
		"006028": 5326, // Macrovision
		"00602a": 5327, // Symicron
		"00602d": 5328, // Alerton
		"00602e": 5329, // Cyclades
		"00602f": 2,    // Cisco
		"006031": 5330, // HRK
		"006032": 5331, // I-Cube
		"006038": 76,   // Nortel
		"006039": 5332, // SanCom
		"00603b": 5333, // AMTEC
		"00603c": 5334, // Hagiwara
		"00603d": 5335, // 3CX
		"00603e": 2,    // Cisco
		"006040": 5336, // Netro
		"006042": 5337, // TKS
		"006043": 5338, // iDirect
		"006044": 5339, // Litton/Poly-Scientific
		"006045": 5340, // Pathlight
		"006046": 5341, // Vmetro
		"006047": 2,    // Cisco
		"006048": 102,  // Dell
		"006049": 5342, // Vina
		"00604b": 5343, // Safe-Com
		"00604c": 2048, // Sagemcom
		"00604d": 4932, // MMC
		"00604e": 5344, // Cycle
		"006050": 5345, // Internix
		"006052": 5346, // PERIPHERALS
		"006054": 5347, // Controlware
		"006057": 2056, // Murata
		"006059": 2924, // Technical
		"00605a": 5348, // Celcore
		"00605b": 5349, // IntraServer
		"00605c": 2,    // Cisco
		"00605d": 5350, // Scanivalve
		"006061": 5351, // Whistle
		"006062": 5352, // Telesync
		"006064": 5080, // Netcomm
		"006068": 5353, // Dialogic
		"006069": 90,   // Brocade
		"00606b": 5354, // Synclayer
		"00606c": 879,  // Arescom
		"006070": 2,    // Cisco
		"006073": 5355, // Redcreek
		"006074": 5356, // QSC
		"006075": 5357, // Pentek
		"006077": 5358, // Prisa
		"00607a": 5359, // DVS
		"00607b": 5360, // Fore
		"00607c": 5361, // WaveAccess
		"00607d": 5362, // Sentient
		"00607e": 5363, // Gigalabs
		"00607f": 1119, // Aurora
		"006081": 5364, // TV/CoM
		"006082": 5365, // Novalink
		"006083": 2,    // Cisco
		"006089": 5366, // Xata
		"00608a": 5367, // Citadel
		"00608b": 5368, // ConferTech
		"00608c": 160,  // 3Com
		"00608d": 5369, // Unipulse
		"00608e": 5370, // HE
		"00608f": 5371, // Tekram
		"006090": 5372, // Artiza
		"006092": 5373, // Micro/Sys
		"006093": 5374, // Varian
		"006094": 372,  // IBM
		"006095": 5375, // Accu-Time
		"006097": 160,  // 3Com
		"006098": 5376, // HT
		"006099": 135,  // SBE
		"00609b": 5377, // AstroNova
		"00609c": 5378, // Perkin-Elmer
		"00609f": 5379, // Phast
		"0060a1": 5380, // VPNet
		"0060a3": 5381, // Continuum
		"0060a4": 5382, // GEW
		"0060a7": 5383, // MICROSENS
		"0060a8": 5384, // Tidomat
		"0060ab": 4269, // Larscom
		"0060ac": 5385, // Resilience
		"0060ad": 5386, // MegaChips
		"0060b0": 302,  // HP
		"0060b1": 5387, // Input/Output
		"0060b3": 5388, // Z-CoM
		"0060b5": 5389, // KEBA
		"0060b6": 5390, // Land
		"0060b7": 5391, // Channelmatic
		"0060b8": 5392, // CoRELIS
		"0060ba": 5393, // Sahara
		"0060bb": 16,   // Cabletron
		"0060bc": 5394, // KeunYoung
		"0060bd": 5395, // Enginuity
		"0060be": 5396, // Webtronics
		"0060bf": 5397, // Macraigor
		"0060c0": 5398, // Nera
		"0060c1": 5399, // WaveSpan
		"0060c2": 5400, // MPL
		"0060c3": 5401, // Netvision
		"0060c5": 5402, // Ancot
		"0060c6": 5403, // DCS
		"0060c7": 5404, // Amati
		"0060c9": 1856, // ControlNet
		"0060ca": 1526, // Harmonic
		"0060cc": 5405, // Emtrakorporated
		"0060cd": 5406, // VideoServer
		"0060ce": 5407, // Acclaim
		"0060cf": 5408, // Alteon
		"0060d0": 5409, // Snmp
		"0060d1": 5183, // Cascade
		"0060d3": 1057, // AT&T
		"0060d4": 5410, // Eldat
		"0060d6": 4339, // NovAtel
		"0060d8": 5411, // Elmic
		"0060d9": 5412, // Transys
		"0060dd": 5413, // Myricom
		"0060de": 5414, // Kayser-Threde
		"0060df": 90,   // Brocade
		"0060e0": 5415, // Axiom
		"0060e1": 5416, // Orckit
		"0060e2": 2507, // Quest
		"0060e4": 5417, // Compuserve
		"0060e6": 5418, // Shomiti
		"0060e7": 5419, // Randata
		"0060e8": 89,   // Hitachi
		"0060e9": 5420, // Atop
		"0060ea": 5421, // StreamLogic
		"0060eb": 5422, // Fourthtrack
		"0060ee": 5423, // Apollo
		"0060ef": 5424, // Flytech
		"0060f1": 5425, // EXP
		"0060f2": 5426, // Lasergraphics
		"0060f4": 5427, // ADVANCED
		"0060f6": 5428, // Nextest
		"0060f7": 5429, // Datafusion
		"0060f8": 5430, // Loran
		"0060fb": 5431, // Packeteer
		"0060fd": 5432, // NetICs
		"0060ff": 5433, // QuVis
		"006151": 3334, // Huawei
		"006171": 545,  // Apple
		"00620b": 858,  // Broadcom
		"0062ec": 2,    // Cisco
		"0063de": 5434, // Cloudwalk
		"006440": 2,    // Cisco
		"0064af": 1238, // Dish
		"006619": 3334, // Huawei
		"00664b": 3334, // Huawei
		"006762": 4921, // Fiberhome
		"00682b": 3334, // Huawei
		"0068eb": 302,  // HP
		"00692d": 5435, // Sunnovo
		"006b6f": 3334, // Huawei
		"006b9e": 3467, // Vizio
		"006bf1": 2,    // Cisco
		"006cbc": 2,    // Cisco
		"006d52": 545,  // Apple
		"006dfb": 5436, // Vutrix
		"006e02": 5437, // Xovis
		"006f64": 152,  // Samsung
		"007147": 5438, // Amazon
		"0071c2": 5439, // Pegatron
		"007204": 152,  // Samsung
		"007263": 2362, // Netcore
		"007278": 2,    // Cisco
		"0073e0": 152,  // Samsung
		"00749c": 5440, // Ruijie
		"007532": 5441, // Inid
		"0075e1": 5442, // Ampt
		"00763d": 5443, // Veea
		"007686": 2,    // Cisco
		"00778d": 2,    // Cisco
		"007888": 2,    // Cisco
		"00789e": 2048, // Sagemcom
		"007b18": 5444, // SENTRY
		"007c2d": 152,  // Samsung
		"007d60": 545,  // Apple
		"007e95": 2,    // Cisco
		"007f28": 2246, // Actiontec
		"008000": 1183, // Multitech
		"008001": 5445, // Periphonics
		"008002": 5446, // Satelcom
		"008003": 5447, // Hytec
		"008004": 5448, // Antlow
		"008005": 5449, // Cactus
		"008006": 5450, // Compuadd
		"008008": 5451, // Dynatech
		"008009": 5452, // Jupiter
		"00800a": 5453, // Japan
		"00800b": 5454, // CSK
		"00800c": 5455, // Videcom
		"00800e": 5456, // Atlantix
		"008010": 5457, // Commodore
		"008013": 5458, // Thomas-Conrad
		"008014": 5459, // Esprit
		"008015": 5460, // Seiko
		"008017": 5461, // PFU
		"008019": 5462, // Dayna
		"00801b": 5463, // Kodiak
		"00801e": 5464, // Xinetron
		"008020": 109,  // Network
		"008022": 5465, // Scan-Optics
		"008024": 5466, // Kalpana
		"008026": 109,  // Network
		"008027": 5467, // Adaptive
		"008028": 5468, // Tradpost
		"008029": 4267, // Eagle
		"00802d": 5469, // Xylogics
		"008030": 2240, // Nexus
		"008031": 5470, // Basys
		"008032": 5471, // Access
		"008037": 306,  // Ericsson
		"00803a": 5472, // Varityper
		"00803b": 5473, // APT
		"00803c": 5474, // TVS
		"00803d": 5475, // Surigiken
		"00803e": 5476, // Synernetics
		"00803f": 5477, // Tatung
		"008043": 5478, // Networld
		"008044": 5479, // Systech
		"008047": 5480, // IN-NET
		"008048": 5081, // Compex
		"00804a": 5481, // PRO-LOG
		"00804b": 4267, // Eagle
		"00804c": 5482, // Contec
		"00804e": 406,  // Apex
		"00804f": 5483, // Daikin
		"008050": 5484, // Ziatech
		"008051": 5485, // Fibermux
		"008053": 5486, // Intellicom
		"008054": 5487, // Frontier
		"008055": 5488, // Fermilab
		"008056": 5489, // SPHINX
		"008057": 5490, // Adsoft
		"008058": 5071, // Printer
		"00805b": 5491, // Condor
		"00805c": 5492, // Agilis
		"00805d": 5493, // Canstar
		"00805f": 302,  // HP
		"008061": 5494, // Litton
		"008062": 1410, // Interface
		"008064": 5495, // Wyse
		"008065": 5496, // Cybergraphic
		"008069": 5497, // Computone
		"00806a": 5498, // ERI
		"00806d": 5499, // Century
		"00806f": 5500, // Onelan
		"008071": 5501, // SAI
		"008072": 5502, // Microplex
		"008075": 5503, // Parsytec
		"008076": 5504, // Mcnc
		"008077": 3230, // Brother
		"00807a": 5505, // Aitech
		"00807b": 5506, // Artel
		"00807c": 5507, // Fibercom
		"00807d": 5508, // Equinox
		"00807f": 5509, // DY-4
		"008080": 5510, // Datamedia
		"008083": 5511, // Amdahl
		"008085": 5512, // H-Three
		"008089": 5513, // Tecnetics
		"00808b": 5514, // Dacoll
		"00808c": 5515, // Netscout
		"00808e": 5516, // Radstone
		"008090": 5517, // Microtek
		"008092": 5518, // Silex
		"008093": 5519, // Xyron
		"008098": 5520, // TDK
		"008099": 1855, // Eaton
		"00809a": 5521, // Novus
		"00809b": 5522, // Justsystem
		"00809c": 5523, // Luxcom
		"00809d": 5524, // Commscraft
		"00809e": 5525, // Datus
		"00809f": 5526, // ALE
		"0080a0": 302,  // HP
		"0080a1": 5527, // Microtest
		"0080a2": 357,  // Creative
		"0080a3": 5528, // Lantronix
		"0080a4": 5529, // Liberty
		"0080a5": 5530, // Speed
		"0080a6": 5531, // Republic
		"0080a7": 934,  // Honeywell
		"0080a8": 5532, // Vitacom
		"0080a9": 5533, // Clearpoint
		"0080aa": 5534, // Maxpeed
		"0080ac": 5535, // Imlogix
		"0080ad": 5536, // Cnet
		"0080af": 5537, // Allumer
		"0080b1": 5538, // Softcom
		"0080b4": 5539, // Sophia
		"0080b5": 1826, // United
		"0080b7": 5540, // Stellar
		"0080ba": 5541, // Specialix
		"0080bc": 89,   // Hitachi
		"0080be": 5542, // Aries
		"0080c1": 5543, // Lanex
		"0080c4": 5544, // Document
		"0080c7": 788,  // Xircom
		"0080c8": 803,  // D-Link
		"0080ca": 830,  // Netcom
		"0080cd": 5545, // Micronics
		"0080d1": 5546, // Kimtron
		"0080d2": 5547, // Shinnihondenko
		"0080d3": 5548, // Shiva
		"0080d4": 5549, // Chase
		"0080d5": 5550, // Cadre
		"0080d6": 5551, // Nuvotech
		"0080d7": 5552, // Fantum
		"0080db": 5553, // Graphon
		"0080dc": 5554, // Picker
		"0080dd": 5555, // GMX/Gimix
		"0080e0": 5556, // XTP
		"0080e2": 5557, // T.D.I
		"0080e5": 5558, // NetApp
		"0080e6": 5559, // Peer
		"0080e9": 70,   // Madge
		"0080ec": 5560, // Supercomputing
		"0080ed": 5561, // IQ
		"0080ef": 5562, // Rational
		"0080f0": 2153, // Panasonic
		"0080f1": 5563, // Opus
		"0080f2": 3476, // Raycom
		"0080f3": 5564, // SUN
		"0080f5": 5565, // Quantel
		"0080f7": 1800, // Zenith
		"0080f8": 5566, // Mizar
		"0080f9": 5567, // Heurikon
		"0080fa": 5568, // RWT
		"0080fb": 5569, // BVM
		"0080fc": 5570, // Avatar
		"0080fe": 5571, // Azure
		"0081c4": 2,    // Cisco
		"0084ed": 67,   // Private
		"0086a0": 67,   // Private
		"008701": 152,  // Samsung
		"008731": 2,    // Cisco
		"008764": 2,    // Cisco
		"008865": 545,  // Apple
		"0088ba": 5572, // NC&C
		"008a55": 3334, // Huawei
		"008a76": 545,  // Apple
		"008a96": 2,    // Cisco
		"008b43": 5573, // Rftech
		"008bfc": 5574, // mixi
		"008cfa": 3978, // Inventec
		"008e73": 2,    // Cisco
		"008ef2": 1368, // Netgear
		"009001": 5575, // Nishimu
		"009002": 5576, // Allgon
		"009003": 5577, // Aplio
		"009004": 160,  // 3Com
		"009005": 5578, // Protech
		"009007": 5579, // Domex
		"009008": 5580, // HanA
		"00900a": 5581, // Proton
		"00900b": 5582, // Lanner
		"00900c": 2,    // Cisco
		"00900e": 5583, // Handlink
		"009010": 5584, // Simulation
		"009011": 5585, // WAVTrace
		"009013": 5586, // Samsan
		"009015": 5587, // Centigram
		"009016": 5588, // ZAC
		"009017": 5589, // Zypcom
		"009019": 5590, // Hermes
		"00901a": 5591, // Unisphere
		"00901c": 5592, // mps
		"00901d": 5593, // PEC
		"009021": 2,    // Cisco
		"009022": 5594, // Ivex
		"009023": 5595, // Zilog
		"009024": 5596, // Pipelinks
		"009027": 421,  // Intel
		"009029": 5597, // Crypto
		"00902b": 2,    // Cisco
		"00902e": 5598, // Namco
		"00902f": 2362, // Netcore
		"009030": 5599, // Honeywell-Dating
		"009031": 5600, // Mysticom
		"009033": 5601, // Innovaphone
		"009034": 5602, // Imagic
		"009036": 5603, // ens
		"009037": 5604, // Acucomm
		"009038": 5605, // Fountain
		"009039": 5606, // Shasta
		"00903d": 5607, // Biopac
		"00903f": 5608, // WorldCast
		"009040": 300,  // Siemens
		"009042": 5609, // ECCS
		"009045": 31,   // Marconi
		"009046": 5610, // Dexdyne
		"009048": 5611, // Zeal
		"009049": 5612, // Entridia
		"00904b": 1450, // Gemtek
		"00904c": 5613, // Epigram
		"00904e": 5614, // Delem
		"009050": 5615, // Teleste
		"009051": 5616, // Ultimate
		"009053": 5617, // Daewoo
		"009056": 5618, // Telestream
		"009057": 5619, // AANetcom
		"009058": 3429, // Ultra
		"00905c": 5620, // Edmi
		"00905e": 5621, // Rauland-Borg
		"00905f": 2,    // Cisco
		"009061": 811,  // Pacific
		"009063": 3259, // Coherent
		"009064": 1142, // Thomson
		"009065": 5622, // Finisar
		"009066": 5623, // Troika
		"009067": 5624, // WalkAbout
		"009068": 5625, // DVT
		"009069": 826,  // Juniper
		"00906a": 5626, // Turnstone
		"00906d": 2,    // Cisco
		"00906e": 5627, // Praxon
		"00906f": 2,    // Cisco
		"009070": 5628, // NEO
		"009072": 5629, // Simrad
		"009073": 5630, // Gaio
		"009074": 5631, // Argon
		"009079": 5632, // ClearOne
		"00907a": 5633, // Spectralink
		"00907b": 5634, // E-Tech
		"00907c": 5635, // Digitalcast
		"00907d": 5636, // Lake
		"00907e": 5637, // Vetronix
		"00907f": 175,  // WatchGuard
		"009080": 5638, // NOT
		"009081": 5639, // Aloha
		"009083": 1896, // Turbo
		"009085": 1120, // Golden
		"009086": 2,    // Cisco
		"009087": 5640, // Itis
		"00908a": 5641, // Bayly
		"00908c": 5642, // Etrend
		"00908d": 5643, // Vickers
		"009090": 5644, // I-BUS
		"009091": 5645, // DigitalScape
		"009092": 2,    // Cisco
		"009093": 5646, // EIZO
		"009094": 5647, // Osprey
		"009096": 2544, // Askey
		"009097": 4214, // Sycamore
		"00909b": 5648, // Markem-Imaje
		"00909c": 125,  // ARRIS
		"00909f": 5649, // Digi-Data
		"0090a0": 5650, // 8X8
		"0090a2": 189,  // CyberTAN
		"0090a3": 5651, // Corecess
		"0090a4": 5652, // Altiga
		"0090a6": 2,    // Cisco
		"0090a7": 5653, // Clientec
		"0090a8": 5654, // NineTiles
		"0090a9": 123,  // WD
		"0090ab": 2,    // Cisco
		"0090ac": 5655, // Optivision
		"0090ad": 5656, // Aspect
		"0090b0": 5657, // Vadem
		"0090b1": 2,    // Cisco
		"0090b2": 5658, // Avici
		"0090b3": 5659, // Agranat
		"0090b4": 5660, // Willowbrook
		"0090b5": 5661, // Nikon
		"0090b6": 5662, // Fibex
		"0090ba": 5663, // Valid
		"0090bc": 5664, // Telemann
		"0090bd": 5665, // Omnia
		"0090bf": 2,    // Cisco
		"0090c4": 5666, // Javelin
		"0090c6": 5667, // Optim
		"0090c7": 5668, // Icom
		"0090c8": 5669, // Waverider
		"0090c9": 5670, // DPAC
		"0090cc": 4478, // Planex
		"0090cf": 76,   // Nortel
		"0090d1": 5671, // Leichu
		"0090d5": 5672, // Euphonix
		"0090d7": 5673, // NetBoost
		"0090d8": 5674, // Whitecross
		"0090d9": 2,    // Cisco
		"0090da": 5675, // Dynarc
		"0090dd": 5676, // MIHARU
		"0090de": 5677, // Cardkey
		"0090e0": 5678, // Systran
		"0090e3": 5679, // Avex
		"0090e5": 5680, // Teknema
		"0090e6": 5681, // ALi
		"0090e8": 5682, // Moxa
		"0090e9": 5683, // Janz
		"0090ea": 2236, // Alpha
		"0090ec": 5684, // Pyrescom
		"0090ee": 5685, // Personal
		"0090ef": 5686, // Integrix
		"0090f2": 2,    // Cisco
		"0090f3": 5656, // Aspect
		"0090f5": 5687, // Clevo
		"0090f6": 5688, // Escalate
		"0090f7": 5689, // Nbase
		"0090f9": 5690, // Imagine
		"0090fa": 129,  // Emulex
		"0090fb": 5691, // Portwell
		"0090fd": 5692, // CopperCom
		"0090fe": 5693, // Elecom
		"0090ff": 5694, // Tellus
		"00919e": 421,  // Intel
		"0091eb": 5695, // Renesas
		"009337": 421,  // Intel
		"009363": 5696, // Uni-Link
		"0094a1": 289,  // F5
		"0094ec": 3334, // Huawei
		"00991d": 3334, // Huawei
		"009acd": 3334, // Huawei
		"009ad2": 2,    // Cisco
		"009c02": 302,  // HP
		"009d6b": 2056, // Murata
		"009e1e": 2,    // Cisco
		"009ec8": 5697, // Xiaomi
		"00a000": 5698, // Centillion
		"00a003": 300,  // Siemens
		"00a004": 1117, // Netpower
		"00a007": 5699, // Apexx
		"00a008": 5700, // Netcorp
		"00a00a": 259,  // Airspan
		"00a00b": 5701, // Computex
		"00a00c": 5702, // Kingmax
		"00a00e": 5515, // Netscout
		"00a00f": 2422, // Broadband
		"00a011": 5703, // Mutoh
		"00a012": 18,   // Telco
		"00a013": 5704, // Teltrend
		"00a014": 5705, // Csir
		"00a015": 849,  // Wyle
		"00a016": 5706, // Micropolis
		"00a01b": 5707, // Premisys
		"00a01c": 5708, // Nascent
		"00a01e": 5709, // EST
		"00a01f": 5710, // Tricord
		"00a020": 5711, // Citicorp/TTI
		"00a024": 160,  // 3Com
		"00a025": 5712, // Redcom
		"00a026": 5713, // Teldat
		"00a027": 5714, // Firepower
		"00a029": 5715, // Coulter
		"00a02a": 5716, // Trancell
		"00a02b": 5717, // Transitions
		"00a02c": 5718, // interWAVE
		"00a02e": 5719, // Brand
		"00a030": 5720, // Captor/SA
		"00a031": 5721, // Hazeltine
		"00a034": 5722, // Axel
		"00a035": 5723, // Cylink
		"00a038": 5724, // Email
		"00a039": 5725, // Ross
		"00a03a": 5726, // Kubotek
		"00a03d": 5727, // Opto-22
		"00a040": 545,  // Apple
		"00a041": 5728, // Inficon
		"00a042": 5729, // Spur
		"00a043": 2358, // American
		"00a046": 57,   // Scitex
		"00a048": 5730, // Questech
		"00a049": 5731, // Digitech
		"00a04e": 5732, // Voelker
		"00a04f": 5733, // Ameritec
		"00a051": 5734, // Angia
		"00a052": 5735, // Stanilite
		"00a054": 67,   // Private
		"00a056": 5736, // Micropross
		"00a057": 5737, // LANCOM
		"00a058": 5738, // Glory
		"00a05b": 5739, // Marquip
		"00a05f": 5740, // BTG
		"00a063": 5741, // JRL
		"00a064": 5742, // KVB/Analect
		"00a065": 1989, // Symantec
		"00a066": 5743, // ISA
		"00a067": 109,  // Network
		"00a068": 5744, // BHP
		"00a069": 5745, // Symmetricom
		"00a06a": 4212, // Verilink
		"00a06e": 5746, // Austron
		"00a070": 5747, // Coastcom
		"00a072": 5748, // Ovation
		"00a073": 5749, // CoM21
		"00a074": 5750, // Perception
		"00a075": 4627, // Micron
		"00a078": 31,   // Marconi
		"00a079": 5751, // Alps
		"00a07b": 5752, // Dawn
		"00a07d": 5753, // Seeq
		"00a07e": 5754, // Avid
		"00a07f": 5755, // GSM-Syntel
		"00a084": 5756, // Dataplex
		"00a085": 67,   // Private
		"00a087": 5757, // Microsemi
		"00a088": 5758, // Essential
		"00a089": 5759, // Xpoint
		"00a08a": 5760, // Brooktrout
		"00a08b": 5761, // Aston
		"00a08d": 5762, // Jacomo
		"00a08f": 5763, // Desknet
		"00a090": 5764, // TimeStep
		"00a091": 5765, // Applicom
		"00a094": 5766, // Comsat
		"00a095": 5767, // Acacia
		"00a098": 5558, // NetApp
		"00a099": 5768, // K-NET
		"00a09b": 30,   // Qpsx
		"00a09c": 5769, // Xyplex
		"00a09e": 5770, // Ictv
		"00a09f": 5771, // Commvision
		"00a0a4": 11,   // Oracle
		"00a0a6": 5772, // M.I
		"00a0a7": 5773, // Vorax
		"00a0a8": 5774, // Renex
		"00a0a9": 5775, // Navtel
		"00a0ad": 31,   // Marconi
		"00a0ae": 5776, // Nucom
		"00a0af": 5777, // WMS
		"00a0b3": 5778, // Zykronix
		"00a0b5": 5779, // 3H
		"00a0b7": 5780, // Cordant
		"00a0b8": 5558, // NetApp
		"00a0b9": 4267, // Eagle
		"00a0ba": 5781, // Patton
		"00a0bb": 5782, // Hilan
		"00a0bc": 5783, // Viasatorporated
		"00a0bd": 4775, // I-Tech
		"00a0c2": 5784, // R.A
		"00a0c3": 5785, // Unicomputer
		"00a0c4": 5786, // Cristie
		"00a0c5": 2692, // ZyXEL
		"00a0c6": 5787, // Qualcomm
		"00a0c8": 202,  // Adtran
		"00a0c9": 421,  // Intel
		"00a0cc": 5788, // Lite-ON
		"00a0ce": 5789, // Ecessa
		"00a0cf": 5790, // Sotas
		"00a0d1": 3978, // Inventec
		"00a0d3": 5791, // Instem
		"00a0d4": 5792, // Radiolan
		"00a0d6": 135,  // SBE
		"00a0d9": 5793, // Convex
		"00a0da": 5794, // INTEGRATED
		"00a0dc": 5795, // O.N
		"00a0dd": 5796, // Azonix
		"00a0de": 727,  // Yamaha
		"00a0df": 5797, // STS
		"00a0e0": 5798, // Tennyson
		"00a0e2": 5799, // Keisokugiken
		"00a0e3": 5800, // XKL
		"00a0e4": 5801, // Optiquest
		"00a0e5": 5802, // NHC
		"00a0e6": 5353, // Dialogic
		"00a0e8": 5803, // Reuters
		"00a0ea": 5804, // Ethercom
		"00a0eb": 3302, // Encore
		"00a0ec": 5805, // Transmitton
		"00a0ee": 5806, // Nashoba
		"00a0ef": 5807, // Lucidata
		"00a0f1": 2772, // MTI
		"00a0f2": 5808, // Infotek
		"00a0f3": 5809, // Staubli
		"00a0f4": 5810, // GE
		"00a0f5": 5811, // Radguard
		"00a0f6": 5812, // AutoGas
		"00a0f7": 5813, // V.I
		"00a0f8": 768,  // Zebra
		"00a0f9": 5814, // Bintec
		"00a0fa": 31,   // Marconi
		"00a0fb": 5815, // Toray
		"00a0fe": 5816, // Boston
		"00a265": 5817, // M2Motive
		"00a289": 2,    // Cisco
		"00a2da": 5818, // INAT
		"00a2ee": 2,    // Cisco
		"00a388": 3512, // BSkyB
		"00a38e": 2,    // Cisco
		"00a3d1": 2,    // Cisco
		"00a45f": 3334, // Huawei
		"00a509": 5819, // WigWag
		"00a554": 421,  // Intel
		"00a5bf": 2,    // Cisco
		"00a6ca": 2,    // Cisco
		"00a742": 2,    // Cisco
		"00aa00": 421,  // Intel
		"00aa01": 421,  // Intel
		"00aa02": 421,  // Intel
		"00aa6e": 2,    // Cisco
		"00aa70": 869,  // LG
		"00ab48": 5820, // eero
		"00ace0": 125,  // ARRIS
		"00ad24": 803,  // D-Link
		"00add5": 3334, // Huawei
		"00aecd": 5821, // Pensando
		"00aefa": 2056, // Murata
		"00af1f": 2,    // Cisco
		"00b017": 5822, // InfoGear
		"00b01c": 5823, // Westport
		"00b01e": 5824, // Rantic
		"00b02a": 5825, // ORSYS
		"00b02d": 5826, // ViaGate
		"00b03b": 5827, // HiQ
		"00b048": 31,   // Marconi
		"00b04a": 2,    // Cisco
		"00b052": 535,  // Atheros
		"00b064": 2,    // Cisco
		"00b069": 5828, // Honewell
		"00b086": 5829, // LocSoft
		"00b08e": 2,    // Cisco
		"00b091": 5830, // Transmeta
		"00b094": 5831, // Alaris
		"00b09a": 5832, // Morrow
		"00b0ae": 5745, // Symmetricom
		"00b0b3": 2357, // Xstreamis
		"00b0c2": 2,    // Cisco
		"00b0ce": 5833, // Viveris
		"00b0d0": 102,  // Dell
		"00b0db": 5834, // Nextcell
		"00b0e1": 2,    // Cisco
		"00b0ec": 5835, // Eacem
		"00b0ee": 5836, // Ajile
		"00b0f0": 5837, // Caly
		"00b0f5": 5838, // NetWorth
		"00b1e3": 2,    // Cisco
		"00b342": 5839, // MacroSAN
		"00b362": 545,  // Apple
		"00b56d": 5840, // David
		"00b5d0": 152,  // Samsung
		"00b5d6": 5841, // Omnibit
		"00b600": 5842, // VOIM
		"00b670": 2,    // Cisco
		"00b69f": 5843, // Latch
		"00b771": 2,    // Cisco
		"00b7a8": 5844, // Heinzinger
		"00b8b3": 2,    // Cisco
		"00b8b6": 687,  // Motorola
		"00bb01": 5845, // Octothorpe
		"00bb1c": 3334, // Huawei
		"00bb3a": 5438, // Amazon
		"00bb60": 421,  // Intel
		"00bb8e": 5846, // HME
		"00bbc1": 87,   // Canon
		"00bbf0": 5847, // Ungermann-Bass
		"00bc60": 2,    // Cisco
		"00bd27": 5848, // Exar
		"00bd3a": 457,  // Nokia
		"00bd3e": 3467, // Vizio
		"00be3b": 3334, // Huawei
		"00be43": 102,  // Dell
		"00be75": 2,    // Cisco
		"00be9e": 4921, // Fiberhome
		"00bf15": 5849, // Genetec
		"00bf61": 152,  // Samsung
		"00bf77": 2,    // Cisco
		"00c000": 5850, // Lanoptics
		"00c002": 2074, // Sercomm
		"00c003": 5851, // Globalnet
		"00c005": 5852, // Livingston
		"00c009": 5853, // KT
		"00c00e": 5854, // Psitech
		"00c00f": 2083, // Quantum
		"00c012": 5855, // Netspan
		"00c013": 5856, // Netrix
		"00c017": 5857, // NetAlly
		"00c018": 5858, // Lanart
		"00c019": 5859, // Leap
		"00c01b": 5860, // Socket
		"00c01c": 3023, // Interlink
		"00c01f": 5861, // S.E.R.C.E.L
		"00c020": 5862, // Arco
		"00c021": 5863, // Netexpress
		"00c022": 5864, // Lasermaster
		"00c023": 5865, // Tutankhamon
		"00c025": 5866, // Dataproducts
		"00c026": 5867, // Lans
		"00c027": 5868, // Cipher
		"00c028": 5869, // Jasco
		"00c02c": 5870, // Centrum
		"00c02e": 5871, // Netwiz
		"00c02f": 5872, // Okuma
		"00c030": 5873, // Integrated
		"00c031": 3861, // Design
		"00c032": 5874, // I-Cubed
		"00c035": 5875, // Quintar
		"00c036": 5876, // Raytech
		"00c037": 5877, // Dynatem
		"00c040": 5878, // Ecci
		"00c042": 5879, // Datalux
		"00c043": 5880, // Stratacom
		"00c044": 5881, // Emcom
		"00c045": 5882, // Isolation
		"00c047": 5883, // Unimicro
		"00c04d": 5884, // Mitec
		"00c04e": 308,  // Comtrol
		"00c04f": 102,  // Dell
		"00c052": 5885, // Burr-Brown
		"00c053": 5656, // Aspect
		"00c056": 5886, // Somelec
		"00c057": 5887, // Myco
		"00c058": 5888, // Dataexpert
		"00c059": 5889, // Denso
		"00c05a": 5890, // Semaphore
		"00c05c": 5891, // Elonex
		"00c05d": 5892, // L&N
		"00c05e": 5893, // Vari-Lite
		"00c05f": 5894, // Fine-PAL
		"00c061": 5895, // Solectek
		"00c062": 5896, // Impulse
		"00c065": 5897, // Scope
		"00c066": 5898, // Docupoint
		"00c06c": 5899, // Svec
		"00c06d": 5900, // Boca
		"00c06e": 5901, // Haft
		"00c06f": 5902, // Komatsu
		"00c071": 5903, // Areanex
		"00c072": 5904, // KNX
		"00c073": 5905, // Xedia
		"00c075": 5906, // Xante
		"00c078": 2328, // Computer
		"00c079": 5907, // Fonsys
		"00c07b": 3297, // Ascend
		"00c07e": 5908, // Kubota
		"00c080": 2170, // Netstar
		"00c081": 5909, // Metrodata
		"00c082": 851,  // Moore
		"00c087": 5910, // Uunet
		"00c08a": 5911, // Lauterbach
		"00c08c": 5912, // Performance
		"00c08f": 2153, // Panasonic
		"00c093": 2309, // Alta
		"00c094": 5913, // VMX
		"00c095": 2477, // ZNYX
		"00c096": 5914, // Tamura
		"00c097": 5915, // Archipel
		"00c098": 5916, // Chuntex
		"00c099": 5917, // Yoshiki
		"00c09a": 5918, // Photonics
		"00c09b": 5027, // Tellabs
		"00c09e": 5919, // Cache
		"00c09f": 3071, // Quanta
		"00c0a2": 5920, // Intermedium
		"00c0a3": 5921, // Dual
		"00c0a4": 5922, // Unigraf
		"00c0a7": 5923, // Seel
		"00c0a8": 5924, // GVC
		"00c0ab": 18,   // Telco
		"00c0ac": 5925, // Gambit
		"00c0ad": 5926, // Marben
		"00c0af": 5927, // Teklogix
		"00c0b0": 5928, // GCC
		"00c0b2": 5929, // Norand
		"00c0b4": 5930, // Myson
		"00c0b6": 5931, // HVE
		"00c0b7": 5932, // APC
		"00c0b9": 5933, // Funk
		"00c0ba": 5934, // Netvantage
		"00c0bd": 5935, // Inex
		"00c0c1": 5936, // Quad/Graphics
		"00c0c2": 752,  // Infinite
		"00c0ca": 5937, // Alfa
		"00c0cb": 1998, // Control
		"00c0cc": 5938, // Telesciences
		"00c0cd": 5939, // Comelta
		"00c0d1": 5940, // Comtree
		"00c0d2": 5941, // Syntellect
		"00c0d4": 5942, // Axon
		"00c0d6": 5943, // J1
		"00c0da": 5944, // Nice
		"00c0db": 5945, // IPC
		"00c0dc": 5946, // EOS
		"00c0dd": 2025, // QLogic
		"00c0de": 5947, // Zcomm
		"00c0df": 5948, // KYE
		"00c0e0": 5323, // DSC
		"00c0e1": 5066, // Sonic
		"00c0e2": 5949, // Calcomp
		"00c0e3": 5950, // Ositech
		"00c0e4": 300,  // Siemens
		"00c0e5": 5951, // Gespac
		"00c0e6": 4212, // Verilink
		"00c0e7": 5952, // Fiberdata
		"00c0e8": 5953, // Plexcom
		"00c0e9": 5954, // OAK
		"00c0ea": 5955, // Array
		"00c0ec": 5956, // Dauphin
		"00c0ee": 2848, // Kyocera
		"00c0ef": 5255, // Abit
		"00c0f0": 4881, // Kingston
		"00c0f2": 5957, // Transition
		"00c0f3": 109,  // Network
		"00c0f5": 5958, // Metacomp
		"00c0f6": 5959, // Celan
		"00c0f7": 5960, // Engage
		"00c0fa": 5961, // Canary
		"00c0fb": 587,  // Advanced
		"00c0fd": 5962, // Prosum
		"00c0fe": 5963, // Aptec
		"00c14f": 5964, // DDL
		"00c164": 2,    // Cisco
		"00c1b1": 2,    // Cisco
		"00c2c6": 421,  // Intel
		"00c30a": 5697, // Xiaomi
		"00c3f4": 152,  // Samsung
		"00c52c": 826,  // Juniper
		"00c610": 545,  // Apple
		"00c88b": 2,    // Cisco
		"00cae5": 2,    // Cisco
		"00cb00": 67,   // Private
		"00cb51": 2048, // Sagemcom
		"00cc34": 826,  // Juniper
		"00cc3f": 3857, // Universal
		"00ccfc": 2,    // Cisco
		"00cdfe": 545,  // Apple
		"00d001": 5965, // VST
		"00d002": 5966, // Ditech
		"00d003": 5967, // Comda
		"00d004": 5968, // Pentacom
		"00d006": 2,    // Cisco
		"00d008": 5969, // Mactell
		"00d009": 4277, // Hsing
		"00d00b": 5970, // RHK
		"00d00e": 5971, // Pluris
		"00d010": 5972, // Convergent
		"00d012": 5973, // Gateworks
		"00d014": 5974, // Root
		"00d01b": 5975, // Mimaki
		"00d01c": 5976, // SBS
		"00d01e": 5977, // Pingtel
		"00d01f": 5978, // Senetas
		"00d021": 5979, // Regent
		"00d022": 5980, // Incredible
		"00d023": 5981, // Infortrend
		"00d024": 5982, // Cognex
		"00d025": 5983, // Xrosstech
		"00d028": 1526, // Harmonic
		"00d02a": 5984, // Voxent
		"00d02b": 5985, // Jetcell
		"00d02d": 5986, // Resideo
		"00d02f": 5987, // Vlsi
		"00d030": 5988, // Safetran
		"00d034": 5989, // Ormec
		"00d035": 5990, // Behavior
		"00d037": 125,  // ARRIS
		"00d038": 5991, // Fivemere
		"00d039": 5992, // Utilicom
		"00d03a": 5993, // Zoneworx
		"00d03b": 2689, // Vision
		"00d03c": 5994, // Vieo
		"00d03d": 5995, // Galileo
		"00d03e": 5996, // Rocketchips
		"00d03f": 2358, // American
		"00d040": 5997, // Sysmate
		"00d041": 5998, // Amigo
		"00d044": 5999, // Alidian
		"00d045": 6000, // Kvaser
		"00d046": 6001, // Dolby
		"00d047": 6002, // XN
		"00d048": 6003, // Ecton
		"00d049": 6004, // Impresstek
		"00d04a": 6005, // Presence
		"00d04e": 6006, // Logibag
		"00d04f": 6007, // Bitronics
		"00d052": 3297, // Ascend
		"00d053": 6008, // Connected
		"00d055": 6009, // Kathrein-Werke
		"00d056": 6010, // Somat
		"00d057": 6011, // Ultrak
		"00d058": 2,    // Cisco
		"00d05a": 6012, // Symbionics
		"00d05d": 6013, // Intelliworxx
		"00d05e": 6014, // Stratabeam
		"00d05f": 6015, // Valcom
		"00d060": 2153, // Panasonic
		"00d061": 6016, // Tremon
		"00d062": 6017, // Digigram
		"00d063": 2,    // Cisco
		"00d064": 6018, // Multitel
		"00d066": 6019, // Wintriss
		"00d067": 6020, // Campio
		"00d068": 6021, // Iwill
		"00d069": 6022, // Technologic
		"00d06a": 6023, // Linkup
		"00d06c": 6024, // Sharewave
		"00d06d": 6025, // Acrison
		"00d071": 6026, // Echelon
		"00d072": 6027, // Broadlogic
		"00d074": 6028, // Taqua
		"00d077": 827,  // Lucent
		"00d079": 2,    // Cisco
		"00d07a": 6029, // Amaquest
		"00d07b": 6030, // Comcam
		"00d07c": 6031, // Koyo
		"00d07d": 6032, // Cosine
		"00d07e": 6033, // Keycorp
		"00d07f": 6034, // Strategy
		"00d080": 6035, // Exabyte
		"00d082": 6036, // Iowave
		"00d083": 6037, // Invertex
		"00d084": 6038, // Nexcomm
		"00d086": 6039, // Foveon
		"00d087": 6040, // Microfirst
		"00d088": 125,  // ARRIS
		"00d089": 6041, // Dynacolor
		"00d08c": 6042, // Genoa
		"00d08f": 6043, // Ardent
		"00d090": 2,    // Cisco
		"00d091": 6044, // Smartsan
		"00d096": 160,  // 3Com
		"00d097": 2,    // Cisco
		"00d09a": 6045, // Filanet
		"00d09b": 6046, // Spectel
		"00d09c": 6047, // Kapadia
		"00d09d": 6048, // Veris
		"00d09e": 1939, // 2Wire
		"00d09f": 6049, // Novtek
		"00d0a4": 6050, // Alantro
		"00d0a6": 6051, // Lanbird
		"00d0aa": 5549, // Chase
		"00d0ac": 6052, // Commscope
		"00d0ad": 6053, // TL
		"00d0ae": 6054, // Oresis
		"00d0af": 6055, // Cutler-Hammer
		"00d0b0": 6056, // Bitswitch
		"00d0b1": 497,  // Omega
		"00d0b2": 4120, // Xiotech
		"00d0b4": 6057, // Katsujima
		"00d0b6": 6058, // Crescent
		"00d0b7": 421,  // Intel
		"00d0b8": 6059, // Iomega
		"00d0b9": 5517, // Microtek
		"00d0ba": 2,    // Cisco
		"00d0bb": 2,    // Cisco
		"00d0bc": 2,    // Cisco
		"00d0be": 6060, // Emutec
		"00d0bf": 5064, // Pivotal
		"00d0c0": 2,    // Cisco
		"00d0c2": 6061, // Balthazar
		"00d0c3": 6062, // Vivid
		"00d0c4": 6063, // Teratech
		"00d0c5": 6064, // Computational
		"00d0c7": 6065, // Pathway
		"00d0c8": 6066, // Prevas
		"00d0c9": 1703, // Advantech
		"00d0ca": 927,  // Intrinsyc
		"00d0cb": 6067, // Dasan
		"00d0cd": 6068, // Atan
		"00d0ce": 6069, // iSystem
		"00d0d1": 4214, // Sycamore
		"00d0d2": 6070, // Epilog
		"00d0d3": 2,    // Cisco
		"00d0d4": 6071, // V-Bits
		"00d0d5": 6072, // Grundig
		"00d0d7": 6073, // B2C2
		"00d0d8": 160,  // 3Com
		"00d0db": 6074, // Mcquay
		"00d0df": 6075, // Kuzumi
		"00d0e0": 6076, // Dooin
		"00d0e3": 6077, // ELE-Chem
		"00d0e4": 2,    // Cisco
		"00d0e5": 6078, // Solidum
		"00d0e6": 6079, // Ibond
		"00d0ea": 6080, // Nextone
		"00d0eb": 6081, // Lightera
		"00d0ec": 6082, // NAKAYO
		"00d0ed": 6083, // Xiox
		"00d0ee": 6084, // Dictaphone
		"00d0ef": 6085, // IGT
		"00d0f0": 6086, // Convision
		"00d0f1": 6087, // Sega
		"00d0f2": 6088, // Monterey
		"00d0f6": 457,  // Nokia
		"00d0f9": 6089, // Acute
		"00d0ff": 2,    // Cisco
		"00d11c": 6090, // Acetel
		"00d6fe": 2,    // Cisco
		"00d76d": 421,  // Intel
		"00d78f": 2,    // Cisco
		"00d861": 1812, // Micro-Star
		"00d8a2": 3334, // Huawei
		"00d9d1": 101,  // Sony
		"00da55": 2,    // Cisco
		"00db45": 6091, // Thamway
		"00db70": 545,  // Apple
		"00dbdf": 421,  // Intel
		"00dcb2": 185,  // Extreme
		"00dd00": 5847, // Ungermann-Bass
		"00dd01": 5847, // Ungermann-Bass
		"00dd02": 5847, // Ungermann-Bass
		"00dd03": 5847, // Ungermann-Bass
		"00dd04": 5847, // Ungermann-Bass
		"00dd05": 5847, // Ungermann-Bass
		"00dd06": 5847, // Ungermann-Bass
		"00dd07": 5847, // Ungermann-Bass
		"00dd08": 5847, // Ungermann-Bass
		"00dd09": 5847, // Ungermann-Bass
		"00dd0a": 5847, // Ungermann-Bass
		"00dd0b": 5847, // Ungermann-Bass
		"00dd0c": 5847, // Ungermann-Bass
		"00dd0d": 5847, // Ungermann-Bass
		"00dd0e": 5847, // Ungermann-Bass
		"00dd0f": 5847, // Ungermann-Bass
		"00defb": 2,    // Cisco
		"00df1d": 2,    // Cisco
		"00e000": 4,    // Fujitsu
		"00e002": 6092, // Crossroads
		"00e004": 6093, // PMC-Sierra
		"00e005": 2924, // Technical
		"00e009": 108,  // Stratus
		"00e00a": 6094, // Diba
		"00e00b": 6095, // Rooftop
		"00e00c": 687,  // Motorola
		"00e00d": 228,  // Radiant
		"00e011": 6096, // Uniden
		"00e012": 6097, // Pluto
		"00e013": 4753, // Eastern
		"00e014": 2,    // Cisco
		"00e015": 6098, // Heiwa
		"00e017": 6099, // EXXACT
		"00e018": 1806, // ASUS
		"00e01a": 6100, // Comtec
		"00e01b": 6101, // Sphere
		"00e01c": 6102, // Cradlepoint
		"00e01d": 6103, // WebTV
		"00e01e": 2,    // Cisco
		"00e01f": 6104, // AVIDIA
		"00e020": 6105, // Tecnomen
		"00e021": 6106, // Freegate
		"00e023": 6107, // Telrad
		"00e024": 6108, // Gadzoox
		"00e025": 6109, // dit
		"00e027": 6110, // DUX
		"00e028": 6111, // Aptix
		"00e02b": 185,  // Extreme
		"00e02c": 5072, // AST
		"00e02d": 6112, // InnoMediaLogic
		"00e02e": 6113, // SPC
		"00e02f": 6114, // Mcns
		"00e030": 6115, // Melita
		"00e033": 6116, // E.E.P.D
		"00e034": 2,    // Cisco
		"00e036": 6117, // Pioneer
		"00e037": 5499, // Century
		"00e038": 6118, // Proxima
		"00e039": 6119, // Paradyne
		"00e03a": 16,   // Cabletron
		"00e03b": 6120, // Prominet
		"00e03c": 6121, // AdvanSys
		"00e03d": 6122, // Focon
		"00e03e": 6123, // Alfatech
		"00e03f": 6124, // Jaton
		"00e040": 6125, // DeskStation
		"00e041": 6126, // Cspi
		"00e042": 6127, // Pacom
		"00e043": 6128, // VitalCom
		"00e044": 6129, // Lsics
		"00e045": 6130, // Touchwave
		"00e047": 6131, // Infocus
		"00e048": 6132, // SDL
		"00e049": 6133, // MICROWI
		"00e04a": 6134, // ZX
		"00e04e": 1337, // Sanyo
		"00e04f": 2,    // Cisco
		"00e051": 6135, // Talx
		"00e052": 90,   // Brocade
		"00e053": 6136, // Cellport
		"00e055": 6137, // Ingenieria
		"00e056": 6138, // Holontech
		"00e05c": 6139, // PHC
		"00e05d": 6140, // Unitec
		"00e05f": 6141, // e-Net
		"00e060": 6142, // Sherwood
		"00e061": 6143, // EdgePoint
		"00e062": 6144, // Host
		"00e063": 16,   // Cabletron
		"00e064": 152,  // Samsung
		"00e066": 6145, // ProMax
		"00e068": 6146, // Merrimac
		"00e069": 6147, // Jaycor
		"00e06a": 6148, // Kapsch
		"00e06c": 3429, // Ultra
		"00e06d": 6149, // Compuware
		"00e06f": 125,  // ARRIS
		"00e070": 6150, // DH
		"00e072": 6151, // Lynk
		"00e074": 6152, // Tiernan
		"00e075": 4212, // Verilink
		"00e077": 6153, // Webgear
		"00e078": 6154, // Berkeley
		"00e079": 6155, // A.T.N.R
		"00e07a": 6156, // Mikrodidakt
		"00e07b": 6157, // BAY
		"00e07c": 2327, // Mettler-Toledo
		"00e07d": 9,    // Netronix
		"00e081": 6158, // Tyan
		"00e082": 6159, // Anerma
		"00e083": 6160, // Jato
		"00e088": 6161, // LTX-Credence
		"00e089": 4205, // ION
		"00e08b": 2025, // QLogic
		"00e08c": 6162, // Neoparadigm
		"00e08d": 6163, // Pressure
		"00e08e": 6164, // Utstarcom
		"00e08f": 2,    // Cisco
		"00e091": 869,  // LG
		"00e092": 6165, // Admtek
		"00e093": 6166, // Ackfin
		"00e096": 6167, // Shimadzu
		"00e098": 2556, // AboCom
		"00e099": 6168, // Samson
		"00e09a": 4088, // Positron
		"00e09b": 5960, // Engage
		"00e09c": 6169, // MII
		"00e09d": 6170, // Sarnoff
		"00e09e": 2083, // Quantum
		"00e0a0": 6171, // Wiltron
		"00e0a2": 6172, // Microslate
		"00e0a3": 2,    // Cisco
		"00e0a6": 6173, // Telogy
		"00e0a8": 6174, // SAT
		"00e0aa": 6175, // Electrosonic
		"00e0ac": 6176, // Midsco
		"00e0ad": 6177, // EES
		"00e0ae": 6178, // Xaqti
		"00e0b0": 2,    // Cisco
		"00e0b2": 6179, // Telmax
		"00e0b3": 6180, // EtherWAN
		"00e0b5": 6043, // Ardent
		"00e0b6": 6181, // Entrada
		"00e0b7": 6182, // Cosworth
		"00e0b9": 6183, // Byas
		"00e0bb": 6184, // NBX
		"00e0bc": 6185, // Symon
		"00e0bd": 1410, // Interface
		"00e0be": 6186, // Genroco
		"00e0c5": 6187, // Bcom
		"00e0c6": 6188, // LINK2IT
		"00e0c9": 6189, // AutomatedLogic
		"00e0cb": 6190, // Reson
		"00e0cc": 6191, // Hero
		"00e0ce": 6192, // ARN
		"00e0d0": 6193, // Netspeed
		"00e0d1": 6194, // Telsis
		"00e0d2": 6195, // Versanet
		"00e0d3": 6196, // DATENTECHNIK
		"00e0d4": 6197, // Excellent
		"00e0d5": 129,  // Emulex
		"00e0d7": 3949, // Sunshine
		"00e0d8": 6198, // LANBit
		"00e0d9": 6199, // Tazmo
		"00e0db": 6200, // ViaVideo
		"00e0dc": 6201, // Nexware
		"00e0dd": 1800, // Zenith
		"00e0de": 6202, // Datax
		"00e0df": 6203, // DZS
		"00e0e0": 6204, // SI
		"00e0e1": 6205, // G2
		"00e0e2": 6206, // Innova
		"00e0e3": 6207, // SK-Elektronik
		"00e0e5": 6208, // Cinco
		"00e0e6": 6209, // IncAA
		"00e0ea": 6210, // Innovat
		"00e0eb": 6211, // Digicom
		"00e0ec": 6212, // Celestica
		"00e0ed": 6213, // Silicom
		"00e0ef": 6214, // Dionex
		"00e0f0": 6215, // Abler
		"00e0f1": 6216, // That
		"00e0f3": 6217, // WebSprint
		"00e0f4": 6218, // INSIDE
		"00e0f5": 6219, // Teles
		"00e0f6": 6220, // Decision
		"00e0f7": 2,    // Cisco
		"00e0f9": 2,    // Cisco
		"00e0fa": 6221, // TRL
		"00e0fb": 6222, // Leightronix
		"00e0fc": 3334, // Huawei
		"00e0fd": 6223, // A-Trend
		"00e0fe": 2,    // Cisco
		"00e16d": 2,    // Cisco
		"00e175": 6224, // AK-Systems
		"00e18c": 421,  // Intel
		"00e3b2": 152,  // Samsung
		"00e406": 3334, // Huawei
		"00e421": 101,  // Sony
		"00e666": 1957, // ARIMA
		"00e6d3": 6225, // Nixdorf
		"00e6e8": 6226, // Netzin
		"00e7e3": 3031, // ZTE
		"00e93a": 3004, // Azurewave
		"00eabd": 2,    // Cisco
		"00eb2d": 101,  // Sony
		"00ebd5": 2,    // Cisco
		"00ebd8": 6227, // Mercusys
		"00ec0a": 5697, // Xiaomi
		"00edb8": 2848, // Kyocera
		"00eeab": 2,    // Cisco
		"00eebd": 1341, // HTC
		"00f051": 6228, // KWB
		"00f28b": 2,    // Cisco
		"00f361": 5438, // Amazon
		"00f39f": 545,  // Apple
		"00f403": 6229, // Orbis
		"00f46f": 152,  // Samsung
		"00f48d": 2873, // Liteon
		"00f4b9": 545,  // Apple
		"00f620": 3521, // Google
		"00f663": 2,    // Cisco
		"00f76f": 545,  // Apple
		"00f81c": 3334, // Huawei
		"00f82c": 2,    // Cisco
		"00f871": 3460, // Demant
		"00fa21": 152,  // Samsung
		"00fa3b": 6230, // Cloos
		"00fc58": 6231, // WebSilicon
		"00fc8b": 5438, // Amazon
		"00fc8d": 870,  // Hitron
		"00fcba": 2,    // Cisco
		"00fd22": 2,    // Cisco
		"00fd45": 302,  // HP
		"00fd4c": 6232, // Nevatec
		"00fec8": 2,    // Cisco
		"020701": 1045, // Racal-Datacom
		"021c7c": 3798, // Perq
		"02608c": 160,  // 3Com
		"027001": 1045, // Racal-Datacom
		"02bb01": 5845, // Octothorpe
		"02c08c": 160,  // 3Com
		"02e6d3": 6225, // Nixdorf
		"04021f": 3334, // Huawei
		"0403d6": 1427, // Nintendo
		"04072e": 1319, // VTech
		"040973": 302,  // HP
		"0409a5": 4575, // HFR
		"040ae0": 6233, // Xmit
		"040cce": 545,  // Apple
		"040d84": 1655, // Silicon
		"040e3c": 302,  // HP
		"04106b": 5697, // Xiaomi
		"041552": 545,  // Apple
		"0415d9": 6234, // Viwone
		"04180f": 152,  // Samsung
		"0418b6": 67,   // Private
		"0418d6": 2978, // Ubiquiti
		"041a04": 6235, // WaveIP
		"041b6d": 869,  // LG
		"041b94": 6144, // Host
		"041bba": 152,  // Samsung
		"041dc7": 3031, // ZTE
		"041e64": 545,  // Apple
		"041e7a": 6236, // DSPWorks
		"042084": 3031, // ZTE
		"04209a": 2153, // Panasonic
		"042144": 3939, // Sunitec
		"0425c5": 3334, // Huawei
		"042665": 545,  // Apple
		"042728": 612,  // Microsoft
		"042758": 3334, // Huawei
		"042ae2": 2,    // Cisco
		"042bbb": 6237, // PicoCELA
		"042f56": 6238, // ATOCS
		"0432f4": 6239, // Partron
		"043389": 3334, // Huawei
		"0433c2": 421,  // Intel
		"0434f6": 687,  // Motorola
		"043855": 6240, // Scopus-Belgium
		"043f72": 432,  // Mellanox
		"044169": 6241, // GoPro
		"04421a": 1806, // ASUS
		"044665": 2056, // Murata
		"04489a": 545,  // Apple
		"04495d": 3334, // Huawei
		"044a50": 6242, // Ramaxel
		"044a6c": 3334, // Huawei
		"044ac6": 6243, // Aipon
		"044bed": 545,  // Apple
		"044e06": 306,  // Ericsson
		"044e5a": 125,  // ARRIS
		"044eaf": 869,  // LG
		"044f17": 531,  // HUMAX
		"044f4c": 3334, // Huawei
		"044f8b": 6244, // Adapteva
		"044faa": 2737, // Ruckus
		"0452c7": 1819, // Bose
		"0452f3": 545,  // Apple
		"045453": 545,  // Apple
		"0455ca": 6245, // BriView
		"0456e5": 421,  // Intel
		"04572f": 6246, // Sertel
		"045a95": 457,  // Nokia
		"045c06": 6247, // Zmodo
		"045c6c": 826,  // Juniper
		"045d4b": 101,  // Sony
		"045d56": 6248, // camtron
		"045fb9": 2,    // Cisco
		"046273": 2,    // Cisco
		"0463e0": 6249, // Nome
		"046565": 6250, // Testop
		"046865": 545,  // Apple
		"0469f8": 545,  // Apple
		"046b1b": 6251, // SysDINE
		"046c59": 421,  // Intel
		"046c9d": 2,    // Cisco
		"046d42": 6252, // Bryston
		"046e49": 6253, // TaiYear
		"0470bc": 6254, // Globalstar
		"047153": 6255, // Sernet
		"047295": 545,  // Apple
		"047503": 3334, // Huawei
		"0475f5": 6256, // Csst
		"04766e": 430,  // Alpsalpine
		"0476b0": 2,    // Cisco
		"047970": 3334, // Huawei
		"047aae": 3334, // Huawei
		"047d7b": 3071, // Quanta
		"047e4a": 6257, // moobox
		"047f0e": 6258, // Barrot
		"04819b": 3512, // BSkyB
		"0481ae": 6259, // Clack
		"04848a": 6260, // 7INOVA
		"04885f": 3334, // Huawei
		"0488e2": 6261, // Beats
		"048a15": 620,  // Avaya
		"048ae1": 833,  // Flextronics
		"048b42": 6262, // Skspruce
		"048c03": 6263, // ThinPAD
		"048c16": 3334, // Huawei
		"048c9a": 3334, // Huawei
		"048d38": 2362, // Netcore
		"049081": 5821, // Pensando
		"049162": 706,  // Microchip
		"049226": 1806, // ASUS
		"0492ee": 6264, // iway
		"04946b": 6265, // Tecno
		"049573": 3031, // ZTE
		"0495e6": 6266, // Tenda
		"0498f3": 430,  // Alpsalpine
		"0499b9": 545,  // Apple
		"049dfe": 6267, // Hivesystem
		"049f06": 6268, // Smobile
		"049f15": 67,   // Private
		"049f81": 5515, // Netscout
		"049fca": 3334, // Huawei
		"04a151": 1368, // Netgear
		"04a222": 2639, // Arcadyan
		"04a2f3": 4921, // Fiberhome
		"04a3f3": 6269, // Emicon
		"04a82a": 457,  // Nokia
		"04ab18": 5693, // Elecom
		"04ab6a": 6270, // Chun-il
		"04b0e7": 3334, // Huawei
		"04b167": 5697, // Xiaomi
		"04b1a1": 152,  // Samsung
		"04b3b6": 6271, // Seamap
		"04b429": 152,  // Samsung
		"04b466": 6272, // BSP
		"04b648": 6273, // Zenner
		"04b86a": 3512, // BSkyB
		"04b9e3": 152,  // Samsung
		"04ba1c": 3334, // Huawei
		"04ba8d": 152,  // Samsung
		"04bc9f": 374,  // Calix
		"04bd70": 3334, // Huawei
		"04bd88": 1685, // Aruba
		"04bd97": 2,    // Cisco
		"04bdbf": 152,  // Samsung
		"04bf6d": 2692, // ZyXEL
		"04bfa8": 6274, // ISB
		"04c06f": 3334, // Huawei
		"04c09c": 5027, // Tellabs
		"04c1b9": 4921, // Fiberhome
		"04c1d8": 3334, // Huawei
		"04c23e": 1341, // HTC
		"04c241": 457,  // Nokia
		"04c461": 2056, // Murata
		"04c5a4": 2,    // Cisco
		"04c807": 5697, // Xiaomi
		"04c880": 6275, // Samtec
		"04c991": 6276, // Phistek
		"04c9d9": 1238, // Dish
		"04cb1d": 6277, // Traka
		"04ccbc": 3334, // Huawei
		"04cd15": 1655, // Silicon
		"04ce14": 6278, // Wilocity
		"04cf25": 6279, // Manycolors
		"04cf4b": 421,  // Intel
		"04cf8c": 6280, // XIAOMI
		"04d13a": 5697, // Xiaomi
		"04d320": 6281, // Itel
		"04d395": 687,  // Motorola
		"04d3b0": 421,  // Intel
		"04d3b5": 3334, // Huawei
		"04d3cf": 545,  // Apple
		"04d437": 6282, // ZNV
		"04d4c4": 1806, // ASUS
		"04d590": 1323, // Fortinet
		"04d6aa": 152,  // Samsung
		"04d921": 6283, // Occuspace
		"04d9f5": 1806, // ASUS
		"04dad2": 2,    // Cisco
		"04db56": 545,  // Apple
		"04db8a": 6284, // Suntech
		"04dd4c": 6285, // Velocytech
		"04dedb": 6286, // Rockport
		"04e0c4": 6287, // Triumph-Adler
		"04e31a": 2048, // Sagemcom
		"04e536": 545,  // Apple
		"04e56e": 6288, // THUB
		"04e598": 5697, // Xiaomi
		"04e662": 6289, // Acroname
		"04e676": 4498, // AMPAK
		"04e77e": 6290, // We
		"04e795": 3334, // Huawei
		"04e9e5": 6291, // Pjrc.com
		"04ea56": 421,  // Intel
		"04eb40": 2,    // Cisco
		"04ecbb": 4921, // Fiberhome
		"04ecd8": 421,  // Intel
		"04ed33": 421,  // Intel
		"04ee91": 6292, // x-fabric
		"04f021": 5081, // Compex
		"04f03e": 3334, // Huawei
		"04f13e": 545,  // Apple
		"04f169": 3334, // Huawei
		"04f352": 3334, // Huawei
		"04f4bc": 6293, // Xena
		"04f7e4": 545,  // Apple
		"04f8c2": 2668, // Flaircomm
		"04f8f8": 6294, // Edgecore
		"04f938": 3334, // Huawei
		"04f993": 6295, // Infinix
		"04f9d9": 6296, // Speaker
		"04fa3f": 6297, // Opticore
		"04fe31": 152,  // Samsung
		"04fe7f": 2,    // Cisco
		"04fe8d": 3334, // Huawei
		"04fea1": 6298, // Fihonest
		"04ff08": 3334, // Huawei
		"080001": 6299, // Computervision
		"080002": 4651, // Bridge
		"080003": 587,  // Advanced
		"080004": 6300, // Cromemco
		"080005": 6301, // Symbolics
		"080006": 300,  // Siemens
		"080007": 545,  // Apple
		"080009": 302,  // HP
		"08000a": 1480, // Nestar
		"08000b": 38,   // Unisys
		"08000d": 639,  // International
		"08000e": 6302, // NCR
		"08000f": 1217, // Mitel
		"080011": 6303, // Tektronix
		"080013": 6304, // Exxon
		"080014": 6305, // Excelan
		"08001b": 102,  // Dell
		"08001d": 6306, // Able
		"08001e": 5423, // Apollo
		"08001f": 3205, // Sharp
		"080020": 11,   // Oracle
		"080021": 6307, // 3M
		"080022": 6308, // NBI
		"080023": 2153, // Panasonic
		"080024": 6309, // 10NET/DCA
		"080029": 6310, // Megatek
		"08002a": 886,  // Mosaic
		"08002d": 6311, // LAN-TEC
		"08002e": 6312, // Metaphor
		"08002f": 1040, // Prime
		"080030": 6313, // Cern
		"080032": 6314, // Tigan
		"080034": 6315, // Filenet
		"080035": 6316, // Microfive
		"080036": 6317, // Intergraph
		"080039": 154,  // Spider
		"08003a": 6318, // Orcatech
		"08003b": 6319, // Torus
		"08003e": 6320, // Codex
		"080040": 6321, // Ferranti
		"080042": 6322, // MACNICA
		"080043": 6323, // Pixel
		"080044": 5840, // David
		"080045": 5140, // Concurrent
		"080046": 101,  // Sony
		"080047": 6324, // Sequent
		"080049": 6325, // Univation
		"08004a": 6326, // Banyan
		"08004b": 6327, // Planning
		"08004c": 6328, // Hydra
		"08004d": 6329, // Corvus
		"08004e": 160,  // 3Com
		"08004f": 215,  // Cygnet
		"080050": 1465, // Daisy
		"080051": 6330, // ExperData
		"080052": 6331, // Insystec
		"08005a": 372,  // IBM
		"08005b": 6332, // VTA
		"08005d": 6333, // Gould
		"08005e": 6334, // Counterpoint
		"08005f": 6335, // Saber
		"080061": 6336, // Jarogate
		"080063": 6337, // Plessey
		"080064": 6338, // Sitasys
		"080065": 5207, // Genrad
		"080066": 6339, // Agfa
		"080067": 6340, // ComDesign
		"080068": 6341, // Ridge
		"08006a": 1057, // AT&T
		"08006b": 5209, // Accel
		"08006e": 6342, // Masscomp
		"080071": 6343, // Matra
		"080073": 6344, // Tecmar
		"080074": 6345, // Casio
		"080077": 6346, // TSL
		"080078": 6347, // Accell
		"08007a": 6348, // Indata
		"08007b": 1337, // Sanyo
		"08007c": 6349, // Vitalink
		"080081": 6350, // Astech
		"080082": 6351, // Veritas
		"080084": 6352, // Tomen
		"080085": 6353, // Elxsi
		"080087": 5769, // Xyplex
		"080088": 90,   // Brocade
		"080089": 6354, // Kinetics
		"08008a": 6355, // PerfTech
		"08008b": 6356, // Pyramid
		"08008c": 109,  // Network
		"08008d": 6357, // Xyvision
		"08008e": 6358, // Tandem
		"08008f": 6359, // Chipcom
		"080090": 6360, // Sonoma
		"080205": 3334, // Huawei
		"08028e": 1368, // Netgear
		"080581": 1916, // Roku
		"0805e2": 826,  // Juniper
		"0808c2": 152,  // Samsung
		"0808ea": 6361, // Amsc
		"0809b6": 6362, // Masimo
		"080d84": 6363, // GECO
		"080ffa": 6364, // KSP
		"08115e": 6365, // Bitel
		"081196": 421,  // Intel
		"0812a5": 5438, // Amazon
		"08152f": 152,  // Samsung
		"0816d5": 6366, // Goertek
		"081735": 2,    // Cisco
		"0817f4": 372,  // IBM
		"08181a": 3031, // ZTE
		"0819a6": 3334, // Huawei
		"081c6e": 5697, // Xiaomi
		"081f3f": 6367, // WondaLink
		"081f71": 1595, // TP-Link
		"081feb": 6368, // BinCube
		"081ff3": 2,    // Cisco
		"0821ef": 152,  // Samsung
		"0823b2": 6369, // Vivo
		"082522": 6370, // Advansee
		"082525": 5697, // Xiaomi
		"082697": 2692, // ZyXEL
		"082cb6": 545,  // Apple
		"082ced": 6371, // Technity
		"082e36": 3334, // Huawei
		"082e5f": 302,  // HP
		"082fe9": 3334, // Huawei
		"08318b": 3334, // Huawei
		"0831a4": 3334, // Huawei
		"083571": 6372, // CASwell
		"0835b2": 6373, // CoreEdge
		"0836c9": 1368, // Netgear
		"08373d": 152,  // Samsung
		"08379c": 6374, // Topaz
		"0838e6": 687,  // Motorola
		"083a5c": 6375, // Junilab
		"083af2": 6376, // Espressif
		"083d88": 152,  // Samsung
		"083e0c": 125,  // ARRIS
		"083e5d": 2048, // Sagemcom
		"083f3e": 6377, // WSH
		"083f76": 6378, // Intellian
		"083fbc": 3031, // ZTE
		"084027": 6379, // Gridstore
		"0840f3": 6266, // Tenda
		"084296": 6380, // Mobile
		"0845d1": 2,    // Cisco
		"084656": 6381, // VEO-Labs
		"08474c": 457,  // Nokia
		"084e1c": 6382, // H2A
		"084f0a": 3334, // Huawei
		"084fa9": 2,    // Cisco
		"084ff9": 2,    // Cisco
		"085531": 1784, // Routerboard.com
		"085700": 1595, // TP-Link
		"0857fb": 5438, // Amazon
		"085a11": 803,  // D-Link
		"085ae0": 6383, // Recovision
		"085b0e": 1323, // Fortinet
		"085bd6": 421,  // Intel
		"085bda": 6384, // CliniCare
		"085c1b": 3334, // Huawei
		"085ddd": 4253, // Mercury
		"08606e": 1806, // ASUS
		"086083": 3031, // ZTE
		"086266": 1806, // ASUS
		"086361": 3334, // Huawei
		"086698": 545,  // Apple
		"0868ea": 6385, // Eito
		"086a0a": 2544, // Askey
		"086ac5": 421,  // Intel
		"086ae5": 5438, // Amazon
		"086bd7": 1655, // Silicon
		"086d41": 545,  // Apple
		"087045": 545,  // Apple
		"087190": 421,  // Intel
		"087402": 545,  // Apple
		"087572": 6386, // Obelux
		"087695": 6387, // Auto
		"087808": 152,  // Samsung
		"08798c": 3334, // Huawei
		"087999": 6388, // AIM
		"087a4c": 3334, // Huawei
		"087baa": 6389, // Svyazkomplektservice
		"087c39": 5438, // Amazon
		"087cbe": 6390, // Quintic
		"087d21": 6391, // Altasec
		"087e64": 6392, // Technicolor
		"087f98": 6369, // Vivo
		"0881b2": 4079, // Logitech
		"0881f4": 826,  // Juniper
		"08849d": 5438, // Amazon
		"08855b": 63,   // Kontron
		"088620": 6265, // Tecno
		"08863b": 2468, // Belkin
		"0887c7": 545,  // Apple
		"088c2c": 152,  // Samsung
		"088dc8": 6393, // Ryowa
		"088e4f": 6394, // SF
		"088e90": 421,  // Intel
		"088edc": 545,  // Apple
		"088f2c": 5001, // Amber
		"0890ba": 6395, // Danlaw
		"089204": 102,  // Dell
		"089356": 3334, // Huawei
		"0894ef": 1592, // Wistron
		"08952a": 6392, // Technicolor
		"0896ad": 2,    // Cisco
		"0896d7": 621,  // AVM
		"089734": 302,  // HP
		"089798": 358,  // Compal
		"0899e8": 6396, // KEMAS
		"089ac7": 3031, // ZTE
		"089b4b": 6397, // iKuai
		"089bf1": 5820, // eero
		"089e01": 3071, // Quanta
		"089e08": 3521, // Google
		"089e84": 3334, // Huawei
		"08a5c8": 5435, // Sunnovo
		"08a6bc": 5438, // Amazon
		"08a7c0": 6392, // Technicolor
		"08a842": 3334, // Huawei
		"08a95a": 3004, // Azurewave
		"08aa55": 687,  // Motorola
		"08aa89": 3031, // ZTE
		"08acc4": 6398, // FMTech
		"08aed6": 152,  // Samsung
		"08af78": 6399, // Totus
		"08b055": 2544, // Askey
		"08b0a7": 6400, // Truebeyond
		"08b258": 826,  // Juniper
		"08b3af": 6369, // Vivo
		"08b49d": 6265, // Tecno
		"08b4b1": 3521, // Google
		"08b4cf": 6401, // Abicom
		"08b738": 6402, // Lite-On
		"08ba22": 6403, // Swaive
		"08bb3c": 833,  // Flextronics
		"08bd43": 1368, // Netgear
		"08be09": 6404, // Astrol
		"08be77": 6405, // Green
		"08beac": 115,  // Edimax
		"08bfa0": 152,  // Samsung
		"08c021": 3334, // Huawei
		"08c06c": 3334, // Huawei
		"08c0eb": 432,  // Mellanox
		"08c5e1": 152,  // Samsung
		"08c6b3": 4145, // Qtech
		"08c729": 545,  // Apple
		"08cc27": 687,  // Motorola
		"08cc68": 2,    // Cisco
		"08cca7": 2,    // Cisco
		"08d09f": 2,    // Cisco
		"08d23e": 421,  // Intel
		"08d29a": 6406, // Proformatique
		"08d34b": 6407, // Techman
		"08d40c": 421,  // Intel
		"08d42b": 152,  // Samsung
		"08d46a": 869,  // LG
		"08d59d": 2048, // Sagemcom
		"08d5c0": 6408, // Seers
		"08df1f": 1819, // Bose
		"08dfcb": 6409, // Systrome
		"08e672": 6410, // Jebsee
		"08e689": 545,  // Apple
		"08e7e5": 3334, // Huawei
		"08e84f": 3334, // Huawei
		"08e9f6": 4498, // AMPAK
		"08ea44": 185,  // Extreme
		"08eb74": 531,  // HUMAX
		"08eca9": 152,  // Samsung
		"08ecf5": 2,    // Cisco
		"08ed9d": 6265, // Tecno
		"08ee8b": 152,  // Samsung
		"08f1ea": 302,  // HP
		"08f458": 3334, // Huawei
		"08f4ab": 545,  // Apple
		"08f606": 3031, // ZTE
		"08f69c": 545,  // Apple
		"08f6f8": 6411, // GET
		"08f8bc": 545,  // Apple
		"08fa28": 3334, // Huawei
		"08fa79": 6369, // Vivo
		"08fbea": 4498, // AMPAK
		"08fc52": 6412, // OpenXS
		"08fc88": 152,  // Samsung
		"08fd0e": 152,  // Samsung
		"08ff44": 545,  // Apple
		"0c01db": 6295, // Infinix
		"0c0227": 6392, // Technicolor
		"0c02bd": 152,  // Samsung
		"0c0535": 826,  // Juniper
		"0c08b4": 531,  // HUMAX
		"0c0e76": 803,  // D-Link
		"0c1105": 6413, // Akuvox
		"0c1167": 2,    // Cisco
		"0c1262": 3031, // ZTE
		"0c130b": 6414, // Uniqoteq
		"0c1420": 152,  // Samsung
		"0c1539": 545,  // Apple
		"0c15c5": 6415, // SDTEC
		"0c1773": 3334, // Huawei
		"0c17f1": 6416, // Telecsys
		"0c191f": 6417, // Inform
		"0c19f8": 545,  // Apple
		"0c1c19": 6418, // Longconn
		"0c1c1a": 5820, // eero
		"0c1c20": 6419, // Kakao
		"0c1daf": 5697, // Xiaomi
		"0c1dc2": 6420, // SeAH
		"0c1ef7": 6421, // Omni-ID
		"0c2026": 6422, // noax
		"0c20d3": 6369, // Vivo
		"0c2138": 6423, // Hengstler
		"0c2724": 2,    // Cisco
		"0c2755": 6424, // Valuable
		"0c29ef": 102,  // Dell
		"0c2a86": 4921, // Fiberhome
		"0c2c54": 3334, // Huawei
		"0c2d89": 6425, // QiiQ
		"0c2fb0": 152,  // Samsung
		"0c3021": 545,  // Apple
		"0c31dc": 3334, // Huawei
		"0c354f": 457,  // Nokia
		"0c35fe": 4921, // Fiberhome
		"0c3747": 3031, // ZTE
		"0c3796": 6426, // Bizlink
		"0c37dc": 3334, // Huawei
		"0c383e": 6427, // Fanvil
		"0c3b50": 545,  // Apple
		"0c3e9f": 545,  // Apple
		"0c413e": 612,  // Microsoft
		"0c41e9": 3334, // Huawei
		"0c42a1": 432,  // Mellanox
		"0c4314": 1655, // Silicon
		"0c43f9": 5438, // Amazon
		"0c45ba": 3334, // Huawei
		"0c473d": 870,  // Hitron
		"0c47c9": 5438, // Amazon
		"0c4885": 869,  // LG
		"0c48c6": 6212, // Celestica
		"0c4b54": 1595, // TP-Link
		"0c4c39": 6428, // MitraStar
		"0c4de9": 545,  // Apple
		"0c4ec0": 5298, // Maxlinear
		"0c4f9b": 3334, // Huawei
		"0c5101": 545,  // Apple
		"0c5415": 421,  // Intel
		"0c54a5": 5439, // Pegatron
		"0c54b9": 457,  // Nokia
		"0c5521": 6429, // Axiros
		"0c57eb": 6430, // Mueller
		"0c599c": 826,  // Juniper
		"0c6046": 6369, // Vivo
		"0c6127": 2246, // Actiontec
		"0c6803": 2,    // Cisco
		"0c6abc": 4921, // Fiberhome
		"0c6e4f": 6431, // PrimeVOLT
		"0c6f9c": 6432, // Shaw
		"0c704a": 3334, // Huawei
		"0c715d": 152,  // Samsung
		"0c718c": 6433, // TCT
		"0c722c": 1595, // TP-Link
		"0c72d9": 3031, // ZTE
		"0c7329": 2074, // Sercomm
		"0c74c2": 545,  // Apple
		"0c75bd": 2,    // Cisco
		"0c771a": 545,  // Apple
		"0c7a15": 421,  // Intel
		"0c8063": 1595, // TP-Link
		"0c8112": 67,   // Private
		"0c8126": 826,  // Juniper
		"0c8268": 1595, // TP-Link
		"0c839a": 3334, // Huawei
		"0c83cc": 2236, // Alpha
		"0c8408": 3334, // Huawei
		"0c8447": 4921, // Fiberhome
		"0c8484": 6434, // Zenovia
		"0c8525": 2,    // Cisco
		"0c8610": 826,  // Juniper
		"0c8910": 152,  // Samsung
		"0c8a87": 6435, // AgLogica
		"0c8b7d": 3467, // Vizio
		"0c8bd3": 6281, // Itel
		"0c8bfd": 421,  // Intel
		"0c8c8f": 6436, // Kamo
		"0c8cdc": 6437, // Suunto
		"0c8dca": 152,  // Samsung
		"0c8e29": 2639, // Arcadyan
		"0c8fff": 3334, // Huawei
		"0c93fb": 6438, // BNS
		"0c9541": 6439, // Chipsea
		"0c96bf": 3334, // Huawei
		"0c96cd": 4253, // Mercury
		"0c9838": 5697, // Xiaomi
		"0c9a3c": 421,  // Intel
		"0c9a42": 6440, // FN-Link
		"0c9d92": 1806, // ASUS
		"0c9e91": 6441, // Sankosha
		"0ca2f4": 6442, // Chameleon
		"0ca694": 3939, // Sunitec
		"0ca8a7": 152,  // Samsung
		"0caaee": 6443, // Ansjer
		"0cac05": 6444, // Unitend
		"0cac8a": 2048, // Sagemcom
		"0caebd": 6445, // Edifier
		"0cb088": 6446, // AITelecom
		"0cb319": 152,  // Samsung
		"0cb459": 6447, // Marketech
		"0cb4ef": 6448, // Digience
		"0cb527": 3334, // Huawei
		"0cb6d2": 803,  // D-Link
		"0cb771": 125,  // ARRIS
		"0cb815": 6376, // Espressif
		"0cb8e8": 5695, // Renesas
		"0cb912": 6449, // JM-DATA
		"0cbc9f": 545,  // Apple
		"0cbd51": 6433, // TCT
		"0cbf15": 5849, // Genetec
		"0cc3a7": 6450, // Meritec
		"0cc413": 3521, // Google
		"0cc47e": 6451, // EUCAST
		"0cc66a": 457,  // Nokia
		"0cc6ac": 6452, // Dags
		"0cc6cc": 3334, // Huawei
		"0cc6fd": 5697, // Xiaomi
		"0cc731": 6453, // Currant
		"0ccb85": 687,  // Motorola
		"0ccc26": 6454, // Airenetworks
		"0ccdd3": 6455, // Eastriver
		"0ccdfb": 6456, // EDIC
		"0ccfd1": 6457, // SPRINGWAVE
		"0cd0f8": 2,    // Cisco
		"0cd292": 421,  // Intel
		"0cd502": 2273, // Westell
		"0cd696": 6458, // Amimon
		"0cd6bd": 3334, // Huawei
		"0cd746": 545,  // Apple
		"0cd7c2": 6459, // Axium
		"0cd996": 2,    // Cisco
		"0cd9c1": 1399, // Visteon
		"0cdc7e": 6376, // Espressif
		"0cdccc": 6460, // Inala
		"0cdd24": 421,  // Intel
		"0cddef": 457,  // Nokia
		"0cdfa4": 152,  // Samsung
		"0ce041": 6461, // iDruide
		"0ce0dc": 152,  // Samsung
		"0ce0e4": 542,  // Plantronics
		"0ce441": 545,  // Apple
		"0ce4a0": 3334, // Huawei
		"0ce5a3": 6462, // SharkNinja
		"0ce5d3": 6150, // DH
		"0ce725": 612,  // Microsoft
		"0ceac9": 125,  // ARRIS
		"0cec8d": 687,  // Motorola
		"0cee99": 5438, // Amazon
		"0cef7c": 6463, // AnaCom
		"0cf019": 6464, // Malgn
		"0cf0b4": 1974, // Globalsat
		"0cf346": 5697, // Xiaomi
		"0cf4d5": 2737, // Ruckus
		"0cf5a4": 2,    // Cisco
		"0cf893": 125,  // ARRIS
		"0cf9c0": 3512, // BSkyB
		"0cfc83": 6465, // Airoha
		"0cfe45": 101,  // Sony
		"100000": 67,   // Private
		"100020": 545,  // Apple
		"10005a": 372,  // IBM
		"1000fd": 6466, // LaonPeople
		"100177": 3334, // Huawei
		"1002b5": 421,  // Intel
		"100501": 5439, // Pegatron
		"1005b1": 125,  // ARRIS
		"1005ca": 2,    // Cisco
		"1005e1": 457,  // Nokia
		"100645": 2048, // Sagemcom
		"1006ed": 2,    // Cisco
		"1007b6": 152,  // Samsung
		"1009f9": 5438, // Amazon
		"100ba9": 421,  // Intel
		"100c24": 6467, // pomdevices
		"100c6b": 1368, // Netgear
		"100d32": 6468, // Embedian
		"100d7f": 1368, // Netgear
		"100e7e": 826,  // Juniper
		"101081": 3031, // ZTE
		"1010b6": 6469, // McCain
		"101212": 6369, // Vivo
		"101218": 6470, // Korins
		"101248": 6471, // ITG
		"101331": 6392, // Technicolor
		"1013ee": 6472, // Justec
		"101b54": 3334, // Huawei
		"101c0c": 545,  // Apple
		"101d51": 6473, // 8Mesh
		"101dc0": 152,  // Samsung
		"101f74": 302,  // HP
		"102279": 6474, // ZeroDesktop
		"1027be": 6475, // Tvip
		"1027f5": 1595, // TP-Link
		"102831": 6476, // Morion
		"102959": 545,  // Apple
		"1029ab": 152,  // Samsung
		"102ab3": 5697, // Xiaomi
		"102b41": 152,  // Samsung
		"102c6b": 4498, // AMPAK
		"102c83": 6477, // Ximea
		"102cef": 6478, // EMU
		"102d96": 6479, // Looxcie
		"102f6b": 612,  // Microsoft
		"103025": 545,  // Apple
		"103034": 6480, // Cara
		"103047": 152,  // Samsung
		"10321d": 3334, // Huawei
		"10327e": 3334, // Huawei
		"103378": 6481, // FLECTRON
		"1033bf": 6392, // Technicolor
		"10341b": 6482, // Spacelink
		"103917": 152,  // Samsung
		"1039e9": 826,  // Juniper
		"103b59": 152,  // Samsung
		"103c59": 3031, // ZTE
		"103d1c": 421,  // Intel
		"103dea": 6483, // HFC
		"103f44": 5697, // Xiaomi
		"1040f3": 545,  // Apple
		"10417f": 545,  // Apple
		"104369": 6484, // Soundmax
		"104400": 3334, // Huawei
		"1045be": 6485, // Norphonic
		"1045f8": 6486, // LNT-Automation
		"1046b4": 6487, // FormericaOE
		"104780": 3334, // Huawei
		"104a7d": 421,  // Intel
		"104d15": 6488, // Viaanix
		"104d77": 428,  // Innovative
		"104e89": 797,  // Garmin
		"104f58": 1685, // Aruba
		"104fa8": 101,  // Sony
		"105072": 2074, // Sercomm
		"105107": 421,  // Intel
		"105172": 3334, // Huawei
		"10521c": 6376, // Espressif
		"105403": 6489, // INTARSO
		"105611": 125,  // ARRIS
		"1056ca": 2481, // Peplink
		"105887": 4921, // Fiberhome
		"105917": 6490, // Tonal
		"105932": 1916, // Roku
		"105af7": 6491, // ADB
		"105c3b": 6492, // Perma-Pipe
		"105cbf": 6493, // DuroByte
		"105ddc": 3334, // Huawei
		"105f06": 2246, // Actiontec
		"105fd4": 6494, // Tendyron
		"10604b": 302,  // HP
		"1062c9": 6495, // Adatis
		"1062d0": 6392, // Technicolor
		"1062e5": 302,  // HP
		"1062eb": 803,  // D-Link
		"1063c8": 2873, // Liteon
		"106530": 102,  // Dell
		"1065a3": 6496, // Panamax
		"1065cf": 6497, // Iqsim
		"10683f": 869,  // LG
		"106f3f": 1077, // Buffalo
		"1070fd": 432,  // Mellanox
		"107100": 3334, // Huawei
		"107636": 6498, // Earda
		"10768a": 6499, // EoCell
		"1077b0": 4921, // Fiberhome
		"1077b1": 152,  // Samsung
		"10785b": 2246, // Actiontec
		"1078d2": 1115, // Elitegroup
		"107a86": 6500, // U&U
		"107b44": 1806, // ASUS
		"107bce": 457,  // Nokia
		"107bef": 2692, // ZyXEL
		"107d1a": 102,  // Dell
		"1083d2": 6501, // Microseven
		"10868c": 125,  // ARRIS
		"1088ce": 4921, // Fiberhome
		"1089fb": 152,  // Samsung
		"108a1b": 6502, // RAONIX
		"108b6a": 6503, // Antailiye
		"108ccf": 2,    // Cisco
		"108eba": 6504, // Molekule
		"108ee0": 152,  // Samsung
		"108ffe": 3334, // Huawei
		"1091a8": 6376, // Espressif
		"109266": 152,  // Samsung
		"109397": 125,  // ARRIS
		"1093e9": 545,  // Apple
		"1094bb": 545,  // Apple
		"10954b": 6505, // Megabyte
		"109693": 5438, // Amazon
		"1097bd": 6376, // Espressif
		"109836": 102,  // Dell
		"1098c3": 2056, // Murata
		"109ab9": 6506, // Tosibox
		"109add": 545,  // Apple
		"109d7a": 3334, // Huawei
		"109fa9": 2246, // Actiontec
		"10a24e": 6507, // GOLD3LINK
		"10a4da": 3334, // Huawei
		"10a51d": 421,  // Intel
		"10a5d0": 2056, // Murata
		"10ae60": 67,   // Private
		"10aea5": 6508, // Duskrise
		"10b1f8": 3334, // Huawei
		"10b26b": 6509, // base
		"10b36f": 6510, // Bowei
		"10b3c6": 2,    // Cisco
		"10b3d5": 2,    // Cisco
		"10b3d6": 2,    // Cisco
		"10b713": 67,   // Private
		"10b7a8": 6511, // CableFree
		"10b7f6": 6512, // Plastoform
		"10b9c4": 545,  // Apple
		"10b9f7": 6513, // Niko-Servodan
		"10bc97": 6369, // Vivo
		"10bd18": 2,    // Cisco
		"10bd55": 6514, // Q-Lab
		"10bef5": 803,  // D-Link
		"10bf48": 1806, // ASUS
		"10c172": 3334, // Huawei
		"10c25a": 6392, // Technicolor
		"10c2ba": 6515, // UTT
		"10c37b": 1806, // ASUS
		"10c3ab": 3334, // Huawei
		"10c595": 2664, // Lenovo
		"10c61f": 3334, // Huawei
		"10c65e": 6516, // Adapt-IP
		"10c6fc": 797,  // Garmin
		"10c9ca": 6517, // Ace
		"10ca81": 6518, // Precia
		"10cc1b": 6519, // Liverock
		"10cd6e": 6520, // Fisys
		"10cdae": 620,  // Avaya
		"10cdb6": 5758, // Essential
		"10ce02": 5438, // Amazon
		"10ce45": 6521, // Miromico
		"10cee9": 545,  // Apple
		"10d07a": 4498, // AMPAK
		"10d0ab": 3031, // ZTE
		"10d38a": 152,  // Samsung
		"10d542": 152,  // Samsung
		"10d7b0": 2048, // Sagemcom
		"10da43": 1368, // Netgear
		"10dc4a": 4921, // Fiberhome
		"10ddb1": 545,  // Apple
		"10ddf4": 6522, // Maxway
		"10dee4": 6523, // automationNEXT
		"10dffc": 300,  // Siemens
		"10e4af": 6524, // APR
		"10e4c2": 152,  // Samsung
		"10e68f": 6525, // Kwangsung
		"10e6ae": 6526, // Source
		"10e77a": 6527, // STMicrolectronics
		"10e7c6": 302,  // HP
		"10e878": 457,  // Nokia
		"10e8a7": 1592, // Wistron
		"10e8ee": 6528, // PhaseSpace
		"10e953": 3334, // Huawei
		"10ec81": 152,  // Samsung
		"10f005": 421,  // Intel
		"10f163": 6529, // TNK
		"10f1f2": 869,  // LG
		"10f311": 2,    // Cisco
		"10f3db": 6530, // Gridco
		"10f681": 6369, // Vivo
		"10f920": 2,    // Cisco
		"10f96f": 869,  // LG
		"10f9ee": 457,  // Nokia
		"10face": 6531, // Reacheng
		"10fbf0": 6532, // KangSheng
		"10fc54": 6533, // Shany
		"10fcb6": 6534, // mirusystems
		"10feed": 1595, // TP-Link
		"1100aa": 67,   // Private
		"111111": 67,   // Private
		"140020": 6535, // LongSung
		"14007d": 3031, // ZTE
		"1400e9": 1217, // Mitel
		"140152": 152,  // Samsung
		"14019c": 6536, // Ubyon
		"1402ec": 302,  // HP
		"140467": 6537, // SNK
		"140708": 67,   // Private
		"1407e0": 6538, // Abrantix
		"1409b4": 3031, // ZTE
		"1409dc": 3334, // Huawei
		"140a29": 6539, // Tiinlab
		"140ac5": 5438, // Amazon
		"140c5b": 6540, // PLNetworks
		"140d4f": 833,  // Flextronics
		"140f42": 457,  // Nokia
		"14109f": 545,  // Apple
		"141114": 6265, // Tecno
		"141333": 3004, // Azurewave
		"141357": 6541, // ATP
		"1413fb": 3334, // Huawei
		"14144b": 5440, // Ruijie
		"14169d": 2,    // Cisco
		"14169e": 4030, // Wingtech
		"14172a": 4921, // Fiberhome
		"141877": 102,  // Dell
		"1418c3": 421,  // Intel
		"141aa3": 687,  // Motorola
		"141bbd": 6542, // Volex
		"141bf0": 6543, // Intellimedia
		"141f78": 152,  // Samsung
		"14205e": 545,  // Apple
		"142233": 4921, // Fiberhome
		"14223b": 3521, // Google
		"1422db": 5820, // eero
		"14230a": 3334, // Huawei
		"1423d7": 6544, // Eutronix
		"142475": 6545, // 4DReplay
		"142882": 6546, // Midicom
		"142971": 6547, // Nemoa
		"142bd2": 6548, // Armtel
		"142d8b": 6549, // Incipio
		"142df5": 6550, // Amphitech
		"142e5e": 2074, // Sercomm
		"143004": 3334, // Huawei
		"14307a": 6551, // Avermetrics
		"1430c6": 687,  // Motorola
		"1432d1": 152,  // Samsung
		"143365": 4539, // TEM
		"14358b": 6552, // Mediabridge
		"143605": 457,  // Nokia
		"1436c6": 2664, // Lenovo
		"14373b": 6553, // PROCOM
		"143aea": 6554, // Dynapower
		"143cc3": 3334, // Huawei
		"143e60": 457,  // Nokia
		"143ebf": 3031, // ZTE
		"143f27": 6555, // Noccela
		"143fa6": 101,  // Sony
		"143fc3": 6556, // SnapAV
		"144146": 934,  // Honeywell
		"1441e2": 6557, // Monaco
		"144658": 3334, // Huawei
		"1446e4": 6558, // Avistel
		"144920": 3334, // Huawei
		"1449bc": 3922, // DrayTek
		"1449e0": 152,  // Samsung
		"144c1a": 6559, // Max
		"144d67": 2127, // Zioncom
		"144e2a": 4564, // Ciena
		"144f8a": 421,  // Intel
		"145051": 3205, // Sharp
		"145120": 3334, // Huawei
		"145412": 6560, // Entis
		"145594": 3334, // Huawei
		"14563a": 3334, // Huawei
		"145645": 6561, // Savitech
		"14568e": 152,  // Samsung
		"14579f": 3334, // Huawei
		"1458d0": 302,  // HP
		"1459c0": 1368, // Netgear
		"145a05": 545,  // Apple
		"145a83": 6562, // Logi-D
		"145afc": 2873, // Liteon
		"145bd1": 125,  // ARRIS
		"145be1": 6563, // nyantec
		"145f94": 3334, // Huawei
		"146080": 3031, // ZTE
		"1460cb": 545,  // Apple
		"14612f": 620,  // Avaya
		"146a0b": 4780, // Cypress
		"146b9a": 3031, // ZTE
		"146e0a": 67,   // Private
		"147411": 4555, // RIM
		"14755b": 421,  // Intel
		"147590": 1595, // TP-Link
		"147740": 3334, // Huawei
		"147bac": 457,  // Nokia
		"147dc5": 2056, // Murata
		"147dda": 545,  // Apple
		"14857f": 421,  // Intel
		"148692": 1595, // TP-Link
		"14876a": 545,  // Apple
		"1488e6": 545,  // Apple
		"148919": 6564, // 2bps
		"14893e": 6565, // Vixtel
		"1489cb": 3334, // Huawei
		"1489fd": 152,  // Samsung
		"148a70": 5280, // ADS
		"148c4a": 3334, // Huawei
		"148f21": 797,  // Garmin
		"148fc6": 545,  // Apple
		"149090": 6566, // KongTop
		"149138": 5438, // Amazon
		"149182": 2468, // Belkin
		"14942f": 6567, // Usys
		"1495ce": 545,  // Apple
		"1496e5": 152,  // Samsung
		"149877": 545,  // Apple
		"14987d": 6392, // Technicolor
		"1499e2": 545,  // Apple
		"149a10": 612,  // Microsoft
		"149d09": 3334, // Huawei
		"149d99": 545,  // Apple
		"149ecf": 102,  // Dell
		"149f3c": 152,  // Samsung
		"149fe8": 2664, // Lenovo
		"14a0f8": 3334, // Huawei
		"14a2a0": 2,    // Cisco
		"14a32f": 3334, // Huawei
		"14a364": 152,  // Samsung
		"14a3b4": 3334, // Huawei
		"14a51a": 3334, // Huawei
		"14a72b": 6568, // currentoptronics
		"14a9d0": 289,  // F5
		"14a9e3": 6569, // MST
		"14ab02": 3334, // Huawei
		"14abc5": 421,  // Intel
		"14abf0": 125,  // ARRIS
		"14b126": 6570, // Industrial
		"14b1c8": 6571, // InfiniWing
		"14b31f": 102,  // Dell
		"14b457": 1655, // Silicon
		"14b484": 152,  // Samsung
		"14b73d": 6572, // ARCHEAN
		"14b7f8": 6392, // Technicolor
		"14b968": 3334, // Huawei
		"14bb6e": 152,  // Samsung
		"14bd61": 545,  // Apple
		"14c03e": 125,  // ARRIS
		"14c0a1": 6573, // UCloud
		"14c126": 457,  // Nokia
		"14c14e": 3521, // Google
		"14c213": 545,  // Apple
		"14c21d": 6574, // Sabtech
		"14c88b": 545,  // Apple
		"14c913": 869,  // LG
		"14c9cf": 6575, // Sigmastar
		"14caa0": 6576, // Hu
		"14cb19": 302,  // HP
		"14cb65": 612,  // Microsoft
		"14cc20": 1595, // TP-Link
		"14cf8d": 6577, // Ohsung
		"14cf92": 1595, // TP-Link
		"14cfe2": 125,  // ARRIS
		"14d00d": 545,  // Apple
		"14d11f": 3334, // Huawei
		"14d169": 3334, // Huawei
		"14d19e": 545,  // Apple
		"14d4fe": 125,  // ARRIS
		"14d64d": 803,  // D-Link
		"14d76e": 6578, // CoNCH
		"14dae9": 1806, // ASUS
		"14dd9c": 6369, // Vivo
		"14dda9": 1806, // ASUS
		"14dde5": 6579, // Mpmkvvcl
		"14de39": 3334, // Huawei
		"14e4ec": 6580, // mLogic
		"14e6e4": 1595, // TP-Link
		"14e9b2": 4921, // Fiberhome
		"14eb08": 3334, // Huawei
		"14eb33": 6581, // BSMediasoft
		"14ebb6": 1595, // TP-Link
		"14edbb": 1939, // 2Wire
		"14ede4": 6582, // Kaiam
		"14ee9d": 6583, // AirNav
		"14efcf": 6584, // Schreder
		"14f0c5": 6585, // Xtremio
		"14f42a": 152,  // Samsung
		"14f65a": 5697, // Xiaomi
		"14f6d8": 421,  // Intel
		"14fb70": 3334, // Huawei
		"14feaf": 6586, // Sagittar
		"14feb5": 102,  // Dell
		"18002d": 101,  // Sony
		"1800db": 6587, // Fitbit
		"1801f1": 5697, // Xiaomi
		"18022d": 3334, // Huawei
		"1802ae": 6369, // Vivo
		"180373": 102,  // Dell
		"1806ff": 143,  // Acer
		"180b52": 6588, // Nanotron
		"180c14": 6589, // iSonea
		"180cac": 87,   // Canon
		"180d2c": 3541, // Intelbras
		"180f76": 803,  // D-Link
		"18104e": 6590, // Cedint-UPM
		"181212": 6591, // Cepton
		"18132d": 3031, // ZTE
		"181456": 457,  // Nokia
		"1816c9": 152,  // Samsung
		"181714": 6592, // Daewoois
		"181725": 3383, // Cameo
		"18193f": 6593, // Tamtron
		"1819d6": 152,  // Samsung
		"181beb": 2246, // Actiontec
		"181dea": 421,  // Intel
		"181e78": 2048, // Sagemcom
		"181e95": 6594, // AuVerte
		"181eb0": 152,  // Samsung
		"182032": 545,  // Apple
		"18204c": 6595, // Kummler+Matter
		"1820a6": 619,  // Sage
		"1820d5": 125,  // ARRIS
		"182195": 152,  // Samsung
		"18227e": 152,  // Samsung
		"182649": 421,  // Intel
		"182666": 152,  // Samsung
		"182a44": 6596, // Hirose
		"182a57": 3334, // Huawei
		"182a7b": 1427, // Nintendo
		"182ad3": 826,  // Juniper
		"182b05": 6597, // 8D
		"182cb4": 6598, // Nectarsoft
		"182df7": 6599, // JY
		"183009": 6600, // Woojin
		"1831bf": 1806, // ASUS
		"1832a2": 6601, // Laon
		"18339d": 2,    // Cisco
		"183451": 545,  // Apple
		"1835d1": 125,  // ARRIS
		"1836fc": 6602, // Elecsys
		"183864": 6603, // CAP-Tech
		"183919": 6604, // Unicoi
		"18399c": 6605, // Skorpios
		"183a2d": 152,  // Samsung
		"183a48": 6606, // VostroNet
		"183cb7": 3334, // Huawei
		"183d5e": 3334, // Huawei
		"183da2": 421,  // Intel
		"183eef": 545,  // Apple
		"183f47": 152,  // Samsung
		"18421d": 67,   // Private
		"184462": 6607, // Riava
		"1844e6": 3031, // ZTE
		"184617": 152,  // Samsung
		"184859": 3797, // Castlenet
		"1848be": 5438, // Amazon
		"1848ca": 2056, // Murata
		"1848d8": 6608, // Fastback
		"184b0d": 2737, // Ruckus
		"184bdf": 6609, // Caavo
		"184cae": 6610, // Continental
		"184e16": 152,  // Samsung
		"184e94": 6611, // Messoa
		"184ecb": 152,  // Samsung
		"184f5d": 6612, // JRC
		"18502a": 6613, // Soarnex
		"185253": 635,  // Pixord
		"185282": 4921, // Fiberhome
		"185345": 457,  // Nokia
		"1854cf": 152,  // Samsung
		"1855e3": 545,  // Apple
		"185644": 3334, // Huawei
		"185680": 421,  // Intel
		"1856c3": 545,  // Apple
		"185869": 6614, // Sailer
		"185880": 2639, // Arcadyan
		"185936": 5697, // Xiaomi
		"1859f5": 2,    // Cisco
		"185a58": 102,  // Dell
		"185ae8": 6615, // Zenotech
		"185b00": 457,  // Nokia
		"185bb3": 152,  // Samsung
		"185d9a": 6616, // BobjGear
		"185e0b": 3031, // ZTE
		"185e0f": 421,  // Intel
		"186024": 302,  // HP
		"1861c7": 6617, // lemonbeat
		"18622c": 2048, // Sagemcom
		"186472": 1685, // Aruba
		"186590": 545,  // Apple
		"1866da": 102,  // Dell
		"1866e3": 6618, // Veros
		"1867b0": 152,  // Samsung
		"18686a": 3031, // ZTE
		"186d99": 6619, // Adanis
		"18703b": 3334, // Huawei
		"18742e": 5438, // Amazon
		"187758": 6620, // Audoo
		"1878d4": 6621, // Verizon
		"187a3b": 1685, // Aruba
		"187a93": 6622, // AMICCOM
		"187c0b": 2737, // Ruckus
		"187eb9": 545,  // Apple
		"188090": 2,    // Cisco
		"1880ce": 6623, // Barberry
		"18810e": 545,  // Apple
		"18828c": 2639, // Arcadyan
		"188331": 152,  // Samsung
		"1883bf": 2639, // Arcadyan
		"188410": 6624, // CoreTrust
		"1886ac": 457,  // Nokia
		"188740": 5697, // Xiaomi
		"188796": 1341, // HTC
		"18895b": 152,  // Samsung
		"1889cf": 6265, // Tecno
		"1889df": 6625, // CerebrEX
		"188b45": 2,    // Cisco
		"188b9d": 2,    // Cisco
		"188ef9": 6626, // G2C
		"189088": 5820, // eero
		"1890d8": 2048, // Sagemcom
		"1892a4": 4564, // Ciena
		"18937f": 4498, // AMPAK
		"189552": 6627, // 1MORE
		"189a67": 6628, // CSE-Servelec
		"189c27": 125,  // ARRIS
		"189c5d": 2,    // Cisco
		"189e2c": 3334, // Huawei
		"189efc": 545,  // Apple
		"18a28a": 6629, // Essel-T
		"18a3e8": 4921, // Fiberhome
		"18a4a9": 6630, // Vanu
		"18a6f7": 1595, // TP-Link
		"18a905": 302,  // HP
		"18a99b": 102,  // Dell
		"18a9a6": 6631, // Nebra
		"18aa0f": 3334, // Huawei
		"18aa45": 6632, // Fon
		"18ab1d": 152,  // Samsung
		"18abf5": 3429, // Ultra
		"18ac9e": 6281, // Itel
		"18ad4d": 6633, // Polostar
		"18aebb": 300,  // Siemens
		"18af61": 545,  // Apple
		"18af8f": 545,  // Apple
		"18b169": 1003, // Sonicwall
		"18b3ba": 6634, // Netlogic
		"18b430": 6635, // Nest
		"18b591": 6636, // I-Storm
		"18b6cc": 6290, // We
		"18b79e": 6637, // Invoxia
		"18b81f": 125,  // ARRIS
		"18bb26": 6440, // FN-Link
		"18bb41": 3334, // Huawei
		"18bdad": 6638, // L-Tech
		"18be92": 957,  // Delta
		"18bfb3": 152,  // Samsung
		"18c04d": 1929, // Giga-Byte
		"18c086": 858,  // Broadcom
		"18c241": 6639, // SonicWall
		"18c2bf": 1077, // Buffalo
		"18c58a": 3334, // Huawei
		"18cc18": 421,  // Intel
		"18cc23": 6640, // Philio
		"18ce94": 152,  // Samsung
		"18cf24": 3334, // Huawei
		"18cf5e": 2873, // Liteon
		"18d071": 6067, // Dasan
		"18d225": 4921, // Fiberhome
		"18d276": 3334, // Huawei
		"18d5b6": 6641, // SMG
		"18d66a": 6642, // Inmarsat
		"18d6c7": 1595, // TP-Link
		"18d6cf": 6643, // Kurth
		"18d949": 6644, // Qvis
		"18d98f": 3334, // Huawei
		"18d9ef": 4937, // Shuttle
		"18dbf2": 102,  // Dell
		"18dc56": 3096, // Yulong
		"18ded7": 3334, // Huawei
		"18dfc1": 6645, // Aetheros
		"18e1ca": 6646, // wanze
		"18e215": 457,  // Nokia
		"18e29f": 6369, // Vivo
		"18e2c2": 152,  // Samsung
		"18e3bc": 6433, // TCT
		"18e728": 2,    // Cisco
		"18e777": 6369, // Vivo
		"18e7b0": 545,  // Apple
		"18e7f4": 545,  // Apple
		"18e80f": 6647, // Viking
		"18e829": 2978, // Ubiquiti
		"18e8dd": 6648, // Moduletek
		"18ece7": 1077, // Buffalo
		"18ee69": 545,  // Apple
		"18ef63": 2,    // Cisco
		"18f0e4": 5697, // Xiaomi
		"18f18e": 6649, // ChipER
		"18f1d8": 545,  // Apple
		"18f22c": 1595, // TP-Link
		"18f292": 6650, // Shannon
		"18f643": 545,  // Apple
		"18f87a": 6651, // i3
		"18f9c4": 1743, // BAE
		"18fb7b": 102,  // Dell
		"18fc26": 6652, // Qorvo
		"18fc9f": 6653, // Changhe
		"18fe34": 6376, // Espressif
		"18ff0f": 421,  // Intel
		"1c0042": 6654, // NARI
		"1c012d": 6655, // Ficer
		"1c0656": 6656, // IDY
		"1c08c1": 869,  // LG
		"1c1161": 4564, // Ciena
		"1c12b0": 5438, // Amazon
		"1c1338": 6657, // Kimball
		"1c1386": 3334, // Huawei
		"1c1448": 125,  // ARRIS
		"1c14b3": 6658, // Airwire
		"1c151f": 3334, // Huawei
		"1c17d3": 2,    // Cisco
		"1c19de": 6659, // eyevis
		"1c1ac0": 545,  // Apple
		"1c1adf": 612,  // Microsoft
		"1c1b0d": 1929, // Giga-Byte
		"1c1b68": 125,  // ARRIS
		"1c1bb5": 421,  // Intel
		"1c1d67": 3334, // Huawei
		"1c1d86": 2,    // Cisco
		"1c1fd4": 6660, // LifeBEAM
		"1c1ff1": 3334, // Huawei
		"1c20db": 3334, // Huawei
		"1c232c": 152,  // Samsung
		"1c234f": 6661, // EDMI
		"1c24cd": 2544, // Askey
		"1c24eb": 6662, // Burlywood
		"1c2704": 3031, // ZTE
		"1c28af": 1685, // Aruba
		"1c2a8b": 457,  // Nokia
		"1c330e": 6663, // PernixData
		"1c34da": 432,  // Mellanox
		"1c34f1": 1655, // Silicon
		"1c36bb": 545,  // Apple
		"1c37bf": 6664, // Cloudium
		"1c3929": 6577, // Ohsung
		"1c3947": 358,  // Compal
		"1c398a": 4921, // Fiberhome
		"1c3a4f": 6665, // AccuSpec
		"1c3a60": 2737, // Ruckus
		"1c3ade": 152,  // Samsung
		"1c3b8f": 6666, // Selve
		"1c3bf3": 1595, // TP-Link
		"1c3cd4": 3334, // Huawei
		"1c3d2f": 3334, // Huawei
		"1c4024": 102,  // Dell
		"1c4190": 3857, // Universal
		"1c4363": 3334, // Huawei
		"1c4419": 1595, // TP-Link
		"1c4586": 1427, // Nintendo
		"1c472f": 3334, // Huawei
		"1c497b": 1450, // Gemtek
		"1c4af7": 6667, // Amon
		"1c4bb9": 6641, // SMG
		"1c4bd6": 3004, // Azurewave
		"1c4c48": 6281, // Itel
		"1c4d66": 5438, // Amazon
		"1c4d70": 421,  // Intel
		"1c501e": 2426, // Sunplus
		"1c51b5": 6668, // Techaya
		"1c53f9": 3521, // Google
		"1c549e": 3857, // Universal
		"1c553a": 6669, // QianGua
		"1c568e": 2127, // Zioncom
		"1c56fe": 687,  // Motorola
		"1c57d8": 6670, // Kraftway
		"1c57dc": 545,  // Apple
		"1c599b": 3334, // Huawei
		"1c5a0b": 6671, // Tegile
		"1c5a3e": 152,  // Samsung
		"1c5a6b": 796,  // Philips
		"1c5cf2": 545,  // Apple
		"1c5f2b": 803,  // D-Link
		"1c60d2": 4921, // Fiberhome
		"1c60de": 4253, // Mercury
		"1c62b8": 152,  // Samsung
		"1c6499": 3869, // Comtrend
		"1c659d": 2873, // Liteon
		"1c66aa": 152,  // Samsung
		"1c6758": 3334, // Huawei
		"1c697a": 6672, // EliteGroup
		"1c69a5": 2220, // BlackBerry
		"1c6a7a": 2,    // Cisco
		"1c6bca": 6673, // Mitsunami
		"1c6e4c": 6674, // Logistic
		"1c6e76": 6675, // Quarion
		"1c6ee6": 6676, // Nhnetworks
		"1c6f65": 1929, // Giga-Byte
		"1c7022": 2056, // Murata
		"1c7125": 545,  // Apple
		"1c721d": 102,  // Dell
		"1c7370": 6677, // Neotech
		"1c73e2": 3334, // Huawei
		"1c740d": 2692, // ZyXEL
		"1c7508": 358,  // Compal
		"1c76ca": 6678, // Terasic
		"1c7b21": 101,  // Sony
		"1c7c11": 6679, // EID
		"1c7cc7": 6680, // Coriant
		"1c7e51": 6681, // 3bumen.com
		"1c7ee5": 803,  // D-Link
		"1c7f2c": 3334, // Huawei
		"1c869a": 152,  // Samsung
		"1c86ad": 6682, // MCT
		"1c872c": 1806, // ASUS
		"1c87e3": 6265, // Tecno
		"1c8e5c": 3334, // Huawei
		"1c8e8e": 6683, // DB
		"1c90be": 306,  // Ericsson
		"1c9148": 545,  // Apple
		"1c9180": 545,  // Apple
		"1c937c": 125,  // ARRIS
		"1c93c4": 5438, // Amazon
		"1c955d": 6684, // I-LAX
		"1c959f": 6685, // Veethree
		"1c97c5": 6686, // Ynomia
		"1c98ec": 302,  // HP
		"1c994c": 2056, // Murata
		"1c9957": 421,  // Intel
		"1c9c26": 6687, // Zoovel
		"1c9c8c": 826,  // Juniper
		"1c9d72": 6392, // Technicolor
		"1c9dc2": 6376, // Espressif
		"1c9e46": 545,  // Apple
		"1c9ecc": 6392, // Technicolor
		"1ca681": 3334, // Huawei
		"1ca852": 6688, // Sensaio
		"1caa07": 2,    // Cisco
		"1cab01": 6689, // Innovolt
		"1caba7": 545,  // Apple
		"1cabc0": 870,  // Hitron
		"1cadd1": 6690, // Bosung
		"1caecb": 3334, // Huawei
		"1caf05": 152,  // Samsung
		"1caff7": 803,  // D-Link
		"1cb044": 2544, // Askey
		"1cb094": 1341, // HTC
		"1cb243": 6691, // TDC
		"1cb3c9": 545,  // Apple
		"1cb72c": 1806, // ASUS
		"1cb796": 3334, // Huawei
		"1cb857": 6692, // Becon
		"1cb9c4": 2737, // Ruckus
		"1cbd0e": 6693, // Amplified
		"1cbdb9": 803,  // D-Link
		"1cc035": 4478, // Planex
		"1cc10c": 421,  // Intel
		"1cc11a": 6694, // Wavetronix
		"1cc1de": 302,  // HP
		"1cc316": 6695, // MileSight
		"1cc63c": 2639, // Arcadyan
		"1ccb99": 6433, // TCT
		"1cccd6": 5697, // Xiaomi
		"1cd1ba": 4921, // Fiberhome
		"1cd1e0": 2,    // Cisco
		"1cd6be": 1592, // Wistron
		"1cda27": 6369, // Vivo
		"1cde57": 4921, // Fiberhome
		"1cdea7": 2,    // Cisco
		"1cdf0f": 2,    // Cisco
		"1ce165": 6696, // Marshal
		"1ce192": 551,  // Qisda
		"1ce504": 3334, // Huawei
		"1ce57f": 152,  // Samsung
		"1ce61d": 152,  // Samsung
		"1ce62b": 545,  // Apple
		"1ce639": 3334, // Huawei
		"1ce6ad": 3334, // Huawei
		"1ce6c7": 2,    // Cisco
		"1ce85d": 2,    // Cisco
		"1cea0b": 6294, // Edgecore
		"1cea1b": 457,  // Nokia
		"1cec72": 1111, // Allradio
		"1cefce": 6697, // bebro
		"1cf03e": 6698, // Wearhaus
		"1cf061": 6699, // SCAPS
		"1cf29a": 3521, // Google
		"1cf42b": 3334, // Huawei
		"1cf4ca": 67,   // Private
		"1cf5e7": 6700, // Turtle
		"1cfa68": 1595, // TP-Link
		"1cfe2b": 5438, // Amazon
		"20014f": 6701, // Linea
		"2002af": 2056, // Murata
		"20040f": 102,  // Dell
		"200505": 6702, // Radmax
		"2008ed": 3334, // Huawei
		"200bc7": 3334, // Huawei
		"200bcf": 1427, // Nintendo
		"200cc8": 1368, // Netgear
		"200f70": 6703, // Foxtech
		"20107a": 1450, // Gemtek
		"2013e0": 152,  // Samsung
		"2016b9": 421,  // Intel
		"2016d8": 2873, // Liteon
		"201742": 869,  // LG
		"201a06": 358,  // Compal
		"201bc9": 826,  // Juniper
		"201d03": 6704, // Elatec
		"201e88": 421,  // Intel
		"201f3b": 3521, // Google
		"201f54": 2050, // Raisecom
		"2021a5": 869,  // LG
		"202564": 5439, // Pegatron
		"202598": 6705, // Teleview
		"2025d2": 4921, // Fiberhome
		"202681": 6265, // Tecno
		"20283e": 3334, // Huawei
		"2028bc": 6706, // Visionscape
		"202ac5": 6707, // Petite-En
		"202bc1": 3334, // Huawei
		"202d07": 152,  // Samsung
		"202d23": 6708, // Collinear
		"20311c": 6369, // Vivo
		"2031eb": 6709, // Hdsn
		"20326c": 152,  // Samsung
		"2032c6": 545,  // Apple
		"2034fb": 5697, // Xiaomi
		"20365b": 6710, // Megafone
		"203706": 2,    // Cisco
		"2037a5": 545,  // Apple
		"2037bc": 6711, // Kuipers
		"203a07": 2,    // Cisco
		"203aef": 6712, // Sivantos
		"203b69": 6369, // Vivo
		"203cae": 545,  // Apple
		"203d66": 125,  // ARRIS
		"203db2": 3334, // Huawei
		"203dbd": 869,  // LG
		"204005": 6713, // feno
		"20443a": 56,   // Schneider
		"2046a1": 6714, // VECOW
		"204747": 102,  // Dell
		"2047da": 5697, // Xiaomi
		"2047ed": 3512, // BSkyB
		"204b22": 5435, // Sunnovo
		"204c03": 1685, // Aruba
		"204c9e": 2,    // Cisco
		"204e6b": 6715, // Axxana
		"204e71": 826,  // Juniper
		"204e7f": 1368, // Netgear
		"204ef6": 3004, // Azurewave
		"2050e7": 4498, // AMPAK
		"205383": 3334, // Huawei
		"2053ca": 6716, // Risk
		"205476": 101,  // Sony
		"2054fa": 3334, // Huawei
		"205531": 152,  // Samsung
		"205532": 6717, // Gotech
		"205721": 5260, // Salix
		"205869": 2737, // Ruckus
		"2059a0": 352,  // Paragon
		"205a00": 6718, // Coval
		"205b2a": 67,   // Private
		"205d47": 6369, // Vivo
		"205e64": 3334, // Huawei
		"205ef7": 152,  // Samsung
		"205f3d": 1640, // Cambridge
		"206274": 612,  // Microsoft
		"20635f": 6719, // Abeeway
		"206432": 152,  // Samsung
		"20658e": 3334, // Huawei
		"2066fd": 6720, // CoNSTELL8
		"20677c": 302,  // HP
		"2067b1": 6097, // Pluto
		"20689d": 2873, // Liteon
		"206980": 545,  // Apple
		"206a8a": 1592, // Wistron
		"206a94": 870,  // Hitron
		"206be7": 1595, // TP-Link
		"206c8a": 185,  // Extreme
		"206d31": 6721, // Firewalla
		"206e9c": 152,  // Samsung
		"20719e": 6394, // SF
		"207355": 125,  // ARRIS
		"207454": 6369, // Vivo
		"207600": 2246, // Actiontec
		"20768f": 545,  // Apple
		"207693": 2664, // Lenovo
		"2078cd": 545,  // Apple
		"2078f0": 545,  // Apple
		"207918": 421,  // Intel
		"207bd2": 6722, // ASIX
		"207c14": 6723, // Qotom
		"207d74": 545,  // Apple
		"208058": 4564, // Ciena
		"2082c0": 5697, // Xiaomi
		"20858c": 6724, // Assa
		"208756": 300,  // Siemens
		"2087ec": 3334, // Huawei
		"20896f": 4921, // Fiberhome
		"208984": 358,  // Compal
		"208986": 3031, // ZTE
		"208c47": 6725, // Tenstorrent
		"208c86": 3334, // Huawei
		"20918a": 6726, // Profalux
		"2091d9": 6727, // I'M
		"209a7d": 2048, // Sagemcom
		"209ae9": 6728, // Volacomm
		"209bcd": 545,  // Apple
		"209e79": 3857, // Universal
		"209ef7": 185,  // Extreme
		"20a171": 5438, // Amazon
		"20a2e4": 545,  // Apple
		"20a2e7": 6729, // Lee-Dickens
		"20a60c": 5697, // Xiaomi
		"20a680": 3334, // Huawei
		"20a6cd": 302,  // HP
		"20a783": 6730, // miControl
		"20a8b9": 300,  // Siemens
		"20a90e": 6433, // TCT
		"20a99b": 612,  // Microsoft
		"20aa25": 6731, // IP-NET
		"20aa4b": 1783, // Linksys
		"20ab37": 545,  // Apple
		"20ab48": 3334, // Huawei
		"20b001": 6392, // Technicolor
		"20b0f7": 6732, // Enclustra
		"20b399": 312,  // Enterasys
		"20b5c6": 6733, // Mimosa
		"20b730": 6734, // TeconGroup
		"20b780": 35,   // Toshiba
		"20b7c0": 6735, // OMICRON
		"20b868": 687,  // Motorola
		"20bbc0": 2,    // Cisco
		"20becd": 5820, // eero
		"20bfdb": 6736, // DVL
		"20c047": 6621, // Verizon
		"20c19b": 421,  // Intel
		"20c3a4": 6737, // RetailNext
		"20c6eb": 2153, // Panasonic
		"20c74f": 6738, // SensorPush
		"20c9d0": 545,  // Apple
		"20cec4": 6739, // Peraso
		"20cf30": 1806, // ASUS
		"20cfae": 2,    // Cisco
		"20d160": 67,   // Private
		"20d21f": 6740, // Wincal
		"20d25f": 6741, // SmartCap
		"20d276": 6281, // Itel
		"20d390": 152,  // Samsung
		"20d5bf": 152,  // Samsung
		"20d607": 457,  // Nokia
		"20d75a": 6742, // Posh
		"20d80b": 826,  // Juniper
		"20d906": 6743, // Iota
		"20da22": 3334, // Huawei
		"20dbab": 152,  // Samsung
		"20dce6": 1595, // TP-Link
		"20dcfd": 3334, // Huawei
		"20df73": 3334, // Huawei
		"20dfb9": 3521, // Google
		"20e09c": 457,  // Nokia
		"20e2a8": 545,  // Apple
		"20e52a": 1368, // Netgear
		"20e564": 125,  // ARRIS
		"20e791": 300,  // Siemens
		"20e7b6": 3857, // Universal
		"20e874": 545,  // Apple
		"20e882": 3031, // ZTE
		"20ed74": 6744, // Ability
		"20ee28": 545,  // Apple
		"20efbd": 1916, // Roku
		"20f17c": 3334, // Huawei
		"20f19e": 125,  // ARRIS
		"20f375": 125,  // ARRIS
		"20f3a3": 3334, // Huawei
		"20f44f": 457,  // Nokia
		"20f478": 5697, // Xiaomi
		"20f77c": 6369, // Vivo
		"20f85e": 957,  // Delta
		"20fdf1": 160,  // 3Com
		"20fe00": 5438, // Amazon
		"20ff36": 6745, // Iflytek
		"2400ba": 3334, // Huawei
		"24016f": 3334, // Huawei
		"2401c7": 2,    // Cisco
		"24050f": 6746, // MTN
		"240588": 3521, // Google
		"240917": 6747, // Devlin
		"240995": 3334, // Huawei
		"240a11": 6433, // TCT
		"240a63": 125,  // ARRIS
		"240a64": 3004, // Azurewave
		"240ac4": 6376, // Espressif
		"240d6c": 6748, // Smnd
		"240dc2": 6433, // TCT
		"241125": 6749, // Hutek
		"241145": 5697, // Xiaomi
		"241148": 6750, // Entropix
		"24166d": 3334, // Huawei
		"24169d": 2,    // Cisco
		"24181d": 152,  // Samsung
		"241a8c": 6751, // Squarehead
		"241ae6": 3334, // Huawei
		"241b7a": 545,  // Apple
		"241eeb": 545,  // Apple
		"241f2c": 6752, // Calsys
		"241fa0": 3334, // Huawei
		"2420c7": 2048, // Sagemcom
		"242124": 457,  // Nokia
		"2421ab": 101,  // Sony
		"24240e": 545,  // Apple
		"242642": 3205, // Sharp
		"2429fe": 2848, // Kyocera
		"242e02": 3334, // Huawei
		"242ffa": 35,   // Toshiba
		"2430f8": 3334, // Huawei
		"243154": 3334, // Huawei
		"243184": 3205, // Sharp
		"24336c": 67,   // Private
		"2436da": 2,    // Cisco
		"2437ef": 4351, // EMC
		"243a82": 6753, // Irts
		"243faa": 3334, // Huawei
		"2440ae": 6754, // NIIC
		"24418c": 421,  // Intel
		"2442bc": 6755, // Alinco
		"244427": 3334, // Huawei
		"24456b": 3334, // Huawei
		"2446c8": 687,  // Motorola
		"2446e4": 3334, // Huawei
		"24470e": 6756, // PentronicAB
		"244b03": 152,  // Samsung
		"244b81": 152,  // Samsung
		"244bfe": 1806, // ASUS
		"244c07": 3334, // Huawei
		"244cab": 6376, // Espressif
		"244ce3": 5438, // Amazon
		"244f1d": 6757, // iRule
		"2453bf": 6758, // Enernet
		"24586e": 3031, // ZTE
		"245880": 6759, // Vizeo
		"245a4c": 2978, // Ubiquiti
		"245ab5": 152,  // Samsung
		"245b83": 5695, // Renesas
		"245ba7": 545,  // Apple
		"245bf0": 2873, // Liteon
		"245cbf": 6760, // Ncse
		"245e48": 545,  // Apple
		"245ebe": 6761, // QNAP
		"245f9f": 3334, // Huawei
		"245fdf": 2848, // Kyocera
		"246081": 6762, // razberi
		"2462ab": 6376, // Espressif
		"2462ce": 1685, // Aruba
		"24649f": 3334, // Huawei
		"246511": 621,  // AVM
		"246880": 6763, // Braveridge
		"2468b0": 152,  // Samsung
		"24693e": 6764, // innodisk
		"24694a": 521,  // Jasmine
		"246968": 1595, // TP-Link
		"2469a5": 3334, // Huawei
		"246aab": 6765, // IT-IS
		"246c8a": 6766, // YUKAI
		"246e96": 102,  // Dell
		"246f28": 6376, // Espressif
		"246f8c": 3334, // Huawei
		"247152": 102,  // Dell
		"247260": 6767, // IOTTECH
		"2474f7": 6241, // GoPro
		"247703": 421,  // Intel
		"24792a": 2737, // Ruckus
		"247e12": 2,    // Cisco
		"247e51": 3031, // ZTE
		"247f20": 2048, // Sagemcom
		"247f3c": 3334, // Huawei
		"248000": 6768, // Westcontrol
		"24813b": 2,    // Cisco
		"2481aa": 6769, // KSH
		"2481c7": 3334, // Huawei
		"24828a": 6770, // Prowave
		"2486f4": 6771, // Ctek
		"248707": 6772, // SEnergy
		"248a07": 432,  // Mellanox
		"2491bb": 3334, // Huawei
		"24920e": 152,  // Samsung
		"2494cb": 125,  // ARRIS
		"249504": 3187, // SFR
		"249745": 3334, // Huawei
		"249eab": 3334, // Huawei
		"24a074": 545,  // Apple
		"24a160": 6376, // Espressif
		"24a2e1": 545,  // Apple
		"24a43c": 2978, // Ubiquiti
		"24a487": 3334, // Huawei
		"24a52c": 3334, // Huawei
		"24a534": 6773, // SynTrust
		"24a65e": 3031, // ZTE
		"24a799": 3334, // Huawei
		"24a7dc": 3512, // BSkyB
		"24a87d": 2153, // Panasonic
		"24ab81": 545,  // Apple
		"24b209": 620,  // Avaya
		"24b2de": 6376, // Espressif
		"24b657": 2,    // Cisco
		"24b6b8": 6774, // Friem
		"24b6fd": 102,  // Dell
		"24b88c": 6775, // Crenus
		"24b8d2": 6776, // Opzoon
		"24bcf8": 3334, // Huawei
		"24be05": 302,  // HP
		"24be18": 6777, // Dadoutek
		"24bf74": 67,   // Private
		"24c0b3": 6778, // RSF
		"24c44a": 3031, // ZTE
		"24c613": 152,  // Samsung
		"24c696": 152,  // Samsung
		"24c9a1": 2737, // Ruckus
		"24c9de": 6779, // Genoray
		"24cacb": 4921, // Fiberhome
		"24cbe7": 6780, // MYK
		"24cd8d": 2056, // Murata
		"24ce33": 5438, // Amazon
		"24d0df": 545,  // Apple
		"24d13f": 6781, // Mexus
		"24d2cc": 6782, // SmartDrive
		"24d3f2": 3031, // ZTE
		"24d76b": 6783, // Syntronic
		"24d7eb": 6376, // Espressif
		"24d81e": 6784, // MirWifi
		"24d921": 620,  // Avaya
		"24da33": 3334, // Huawei
		"24da9b": 687,  // Motorola
		"24dbac": 3334, // Huawei
		"24dbed": 152,  // Samsung
		"24dec6": 1685, // Aruba
		"24df6a": 3334, // Huawei
		"24e314": 545,  // Apple
		"24e4c8": 4921, // Fiberhome
		"24e853": 869,  // LG
		"24e927": 2714, // TomTom
		"24e9b3": 2,    // Cisco
		"24e9ca": 3334, // Huawei
		"24ea40": 6785, // Helmholz
		"24ec99": 2544, // Askey
		"24edfd": 300,  // Siemens
		"24ee9a": 421,  // Intel
		"24f094": 545,  // Apple
		"24f0ff": 6786, // GHT
		"24f128": 6787, // Telstra
		"24f27f": 302,  // HP
		"24f57e": 6788, // HWH
		"24f5a2": 2468, // Belkin
		"24f5aa": 152,  // Samsung
		"24f603": 3334, // Huawei
		"24f677": 545,  // Apple
		"24fb65": 3334, // Huawei
		"24fc4e": 826,  // Juniper
		"24fce5": 152,  // Samsung
		"24fd0d": 3541, // Intelbras
		"24fd52": 2873, // Liteon
		"24fd5b": 6789, // SmartThings
		"280244": 545,  // Apple
		"2802d8": 152,  // Samsung
		"2804e0": 6790, // Fermax
		"28052e": 6791, // Dematic
		"28068d": 6792, // ITL
		"280b5c": 545,  // Apple
		"280dfc": 101,  // Sony
		"280feb": 869,  // LG
		"28101b": 6793, // MagnaCom
		"28107b": 803,  // D-Link
		"2811a5": 1819, // Bose
		"2811a8": 421,  // Intel
		"2811ec": 3334, // Huawei
		"281471": 6794, // Lantis
		"28162e": 1939, // 2Wire
		"28167f": 5697, // Xiaomi
		"2816a8": 612,  // Microsoft
		"2816ad": 421,  // Intel
		"281709": 3334, // Huawei
		"2817ce": 6795, // Omnisense
		"281878": 612,  // Microsoft
		"281b04": 6796, // Zalliant
		"282373": 6797, // Digita
		"2824ff": 1592, // Wistron
		"2826a6": 6798, // PBR
		"2827bf": 152,  // Samsung
		"28285d": 2692, // ZyXEL
		"2829cc": 6799, // Corsa
		"2829d9": 6800, // GlobalBeiMing
		"282a87": 6281, // Itel
		"282b96": 3334, // Huawei
		"282cb2": 1595, // TP-Link
		"282d06": 4498, // AMPAK
		"2830ac": 6801, // Frontiir
		"283152": 3334, // Huawei
		"283166": 6369, // Vivo
		"2832c5": 531,  // HUMAX
		"283334": 3334, // Huawei
		"2834a2": 2,    // Cisco
		"283737": 545,  // Apple
		"28385c": 833,  // Flextronics
		"2838cf": 6802, // Gen2wave
		"283926": 189,  // CyberTAN
		"28395e": 152,  // Samsung
		"2839e7": 6803, // Preceno
		"283b82": 803,  // D-Link
		"283ce4": 3334, // Huawei
		"283e76": 6804, // Common
		"283f69": 101,  // Sony
		"2841c6": 3334, // Huawei
		"2841ec": 3334, // Huawei
		"2847aa": 457,  // Nokia
		"284846": 6805, // GridCentric
		"2848e7": 3334, // Huawei
		"284c53": 6806, // Intune
		"284d92": 6807, // Luminator
		"285261": 2,    // Cisco
		"2852e0": 6808, // Layon
		"28534e": 3334, // Huawei
		"285471": 3334, // Huawei
		"28563a": 4921, // Fiberhome
		"285767": 1238, // Dish
		"285aeb": 545,  // Apple
		"285f2f": 6809, // RNware
		"285fdb": 3334, // Huawei
		"286094": 6810, // Capelec
		"286336": 300,  // Siemens
		"2863bd": 4459, // Aptiv
		"2864b0": 3334, // Huawei
		"2866e3": 3004, // Azurewave
		"2868d2": 3334, // Huawei
		"286ab8": 545,  // Apple
		"286aba": 545,  // Apple
		"286c07": 6280, // XIAOMI
		"286d97": 6811, // SAMJIN
		"286ed4": 3334, // Huawei
		"286f7f": 2,    // Cisco
		"2872c5": 1015, // Smartmatic
		"2872f0": 4308, // Athena
		"287610": 6812, // IgniteNet
		"287777": 3031, // ZTE
		"2877f1": 545,  // Apple
		"287aee": 125,  // ARRIS
		"287b09": 3031, // ZTE
		"287fcf": 421,  // Intel
		"288023": 302,  // HP
		"288088": 1368, // Netgear
		"288335": 152,  // Samsung
		"2884fa": 3205, // Sharp
		"28852d": 6813, // Touch
		"2887ba": 1595, // TP-Link
		"288a1c": 826,  // Juniper
		"288cb8": 3031, // ZTE
		"28924a": 302,  // HP
		"2893fe": 2,    // Cisco
		"28940f": 2,    // Cisco
		"2897b8": 6814, // myenergi
		"28987b": 152,  // Samsung
		"28993a": 3793, // Arista
		"289afa": 6433, // TCT
		"289e97": 3334, // Huawei
		"289efc": 2048, // Sagemcom
		"28a02b": 545,  // Apple
		"28a183": 430,  // Alpsalpine
		"28a186": 6815, // enblink
		"28a1eb": 6816, // Etek
		"28a241": 6817, // exlar
		"28a24b": 826,  // Juniper
		"28a53f": 6369, // Vivo
		"28a6ac": 6818, // seca
		"28a6db": 3334, // Huawei
		"28ac9e": 2,    // Cisco
		"28affd": 2,    // Cisco
		"28b2bd": 421,  // Intel
		"28b371": 2737, // Ruckus
		"28b448": 3334, // Huawei
		"28b4fb": 6819, // Sprocomm
		"28b9d9": 52,   // Radisys
		"28ba18": 6820, // NextNav
		"28bab5": 152,  // Samsung
		"28bb59": 6821, // RNET
		"28bc18": 6822, // SourcingOverseas
		"28bc56": 6823, // EMAC
		"28bd89": 3521, // Google
		"28be03": 6433, // TCT
		"28be9b": 6392, // Technicolor
		"28bf89": 4921, // Fiberhome
		"28c0da": 826,  // Juniper
		"28c21f": 152,  // Samsung
		"28c2dd": 3004, // Azurewave
		"28c538": 545,  // Apple
		"28c63f": 421,  // Intel
		"28c68e": 1368, // Netgear
		"28c709": 545,  // Apple
		"28c718": 6824, // Altierre
		"28c7ce": 2,    // Cisco
		"28c825": 6825, // DellKing
		"28c87a": 125,  // ARRIS
		"28c87c": 3031, // ZTE
		"28c914": 6826, // Taimag
		"28cbeb": 6827, // One
		"28cc01": 152,  // Samsung
		"28cd1c": 6828, // Espotel
		"28cd4c": 6829, // Individual
		"28cf08": 6830, // Essys
		"28cfda": 545,  // Apple
		"28cfe9": 545,  // Apple
		"28d0cb": 1640, // Cambridge
		"28d0ea": 421,  // Intel
		"28d0f5": 5440, // Ruijie
		"28d1af": 457,  // Nokia
		"28d244": 4919, // LCFC
		"28d3ea": 3334, // Huawei
		"28d93e": 6831, // Telecor
		"28d997": 6832, // Yuduan
		"28de65": 1685, // Aruba
		"28dea8": 3031, // ZTE
		"28dee5": 3334, // Huawei
		"28def6": 6833, // bioMerieux
		"28dfeb": 421,  // Intel
		"28e02c": 545,  // Apple
		"28e14c": 545,  // Apple
		"28e31f": 5697, // Xiaomi
		"28e347": 2873, // Liteon
		"28e34e": 3334, // Huawei
		"28e476": 6834, // Pi-Coral
		"28e5b0": 3334, // Huawei
		"28e608": 6835, // Tokheim
		"28e794": 6836, // Microtime
		"28e7cf": 545,  // Apple
		"28ea2d": 545,  // Apple
		"28ec95": 545,  // Apple
		"28ed6a": 545,  // Apple
		"28ede0": 4498, // AMPAK
		"28ee52": 1595, // TP-Link
		"28ef01": 5438, // Amazon
		"28f033": 545,  // Apple
		"28f076": 545,  // Apple
		"28f10e": 102,  // Dell
		"28f49b": 6837, // Leetek
		"28f532": 6838, // ADD-Engineering
		"28faa0": 6369, // Vivo
		"28fbae": 3334, // Huawei
		"28fbd3": 6839, // Ragentek
		"28fede": 6840, // CoMESTA
		"28ff3c": 545,  // Apple
		"28ff3e": 3031, // ZTE
		"28ffb2": 35,   // Toshiba
		"2c002c": 6841, // Unowhy
		"2c0033": 6842, // EControls
		"2c00ab": 125,  // ARRIS
		"2c00f7": 6843, // XOS
		"2c01b5": 2,    // Cisco
		"2c029f": 6844, // 3ALogics
		"2c073c": 6845, // Devline
		"2c0786": 3334, // Huawei
		"2c081c": 6846, // OVH
		"2c088c": 531,  // HUMAX
		"2c094d": 2238, // Raptor
		"2c09cb": 6847, // Cobs
		"2c0bab": 3334, // Huawei
		"2c0be9": 2,    // Cisco
		"2c0da7": 421,  // Intel
		"2c0e3d": 152,  // Samsung
		"2c10c1": 1427, // Nintendo
		"2c1165": 1655, // Silicon
		"2c15bf": 152,  // Samsung
		"2c15e1": 6848, // Phicomm
		"2c18ae": 176,  // Trend
		"2c1a01": 3334, // Huawei
		"2c1a05": 2,    // Cisco
		"2c1a31": 6849, // Electronics
		"2c1db8": 125,  // ARRIS
		"2c1eea": 6850, // Aerodev
		"2c1f23": 545,  // Apple
		"2c200b": 545,  // Apple
		"2c2131": 826,  // Juniper
		"2c2172": 826,  // Juniper
		"2c21d7": 6851, // IMAX
		"2c233a": 302,  // HP
		"2c26c5": 3031, // ZTE
		"2c2768": 3334, // Huawei
		"2c27d7": 302,  // HP
		"2c2997": 612,  // Microsoft
		"2c2bf9": 869,  // LG
		"2c2d48": 6852, // bct
		"2c301a": 6392, // Technicolor
		"2c3033": 1368, // Netgear
		"2c3068": 2277, // Pantech
		"2c3124": 2,    // Cisco
		"2c3311": 2,    // Cisco
		"2c3358": 421,  // Intel
		"2c3361": 545,  // Apple
		"2c36a0": 6853, // Capisco
		"2c36f8": 2,    // Cisco
		"2c3796": 6854, // Cybo
		"2c3996": 2048, // Sagemcom
		"2c39c1": 4564, // Ciena
		"2c3a28": 6855, // Fagor
		"2c3a91": 3334, // Huawei
		"2c3ae8": 6376, // Espressif
		"2c3bfd": 6856, // Netstor
		"2c3ecf": 2,    // Cisco
		"2c3f38": 2,    // Cisco
		"2c3f3e": 6857, // Alge-Timing
		"2c4053": 152,  // Samsung
		"2c4138": 302,  // HP
		"2c41a1": 1819, // Bose
		"2c4205": 6858, // Lytx
		"2c43be": 5435, // Sunnovo
		"2c4401": 152,  // Samsung
		"2c44fd": 302,  // HP
		"2c4881": 6369, // Vivo
		"2c4a11": 4564, // Ciena
		"2c4cc6": 2056, // Murata
		"2c4d54": 1806, // ASUS
		"2c4dde": 6265, // Tecno
		"2c4f52": 2,    // Cisco
		"2c52af": 3334, // Huawei
		"2c53d7": 6859, // Sonova
		"2c542d": 2,    // Cisco
		"2c5491": 612,  // Microsoft
		"2c54cf": 869,  // LG
		"2c553c": 6860, // Gainspeed
		"2c55d3": 3334, // Huawei
		"2c56dc": 1806, // ASUS
		"2c5731": 4030, // Wingtech
		"2c5741": 2,    // Cisco
		"2c584f": 125,  // ARRIS
		"2c58e8": 3334, // Huawei
		"2c598a": 869,  // LG
		"2c59e5": 302,  // HP
		"2c5a05": 457,  // Nokia
		"2c5a0f": 2,    // Cisco
		"2c5aa3": 6861, // Promate
		"2c5be1": 6862, // Centripetal
		"2c5d93": 2737, // Ruckus
		"2c5ff3": 6863, // Pertronic
		"2c600c": 3071, // Quanta
		"2c61f6": 545,  // Apple
		"2c6289": 6864, // Regenersis
		"2c641f": 3467, // Vizio
		"2c6798": 6865, // IntalTech
		"2c6bf5": 826,  // Juniper
		"2c6dc1": 421,  // Intel
		"2c6e85": 421,  // Intel
		"2c7155": 6866, // HiveMotion
		"2c71ff": 5438, // Amazon
		"2c72c3": 6867, // Soundmatters
		"2c7360": 6498, // Earda
		"2c73a0": 2,    // Cisco
		"2c768a": 302,  // HP
		"2c780e": 3334, // Huawei
		"2c784c": 6868, // Iton
		"2c79d7": 2048, // Sagemcom
		"2c7b5a": 6869, // Milper
		"2c7e81": 125,  // ARRIS
		"2c7ecf": 6870, // Onzo
		"2c86d2": 2,    // Cisco
		"2c8a72": 1341, // HTC
		"2c8db1": 421,  // Intel
		"2c9127": 6871, // Eintechno
		"2c9464": 6872, // Cincoze
		"2c9569": 125,  // ARRIS
		"2c957f": 3031, // ZTE
		"2c9662": 6873, // Invenit
		"2c97b1": 3334, // Huawei
		"2c97ed": 101,  // Sony
		"2c9924": 125,  // ARRIS
		"2c9aa4": 6874, // Eolo
		"2c9d1e": 3334, // Huawei
		"2c9e5f": 125,  // ARRIS
		"2c9efc": 87,   // Canon
		"2c9ffb": 1592, // Wistron
		"2ca02f": 6875, // Veroguard
		"2ca042": 3334, // Huawei
		"2ca157": 6876, // acromate
		"2ca17d": 125,  // ARRIS
		"2ca2b4": 6877, // Fortify
		"2ca327": 6878, // Oraimo
		"2ca780": 6879, // True
		"2ca835": 4555, // RIM
		"2ca89c": 6880, // Creatz
		"2caa8e": 6881, // Wyze
		"2cab00": 3334, // Huawei
		"2cabeb": 2,    // Cisco
		"2cac44": 6882, // Conextop
		"2cae2b": 152,  // Samsung
		"2cb05d": 1368, // Netgear
		"2cb0df": 6883, // Soliton
		"2cb21a": 6848, // Phicomm
		"2cb43a": 545,  // Apple
		"2cb693": 563,  // Radware
		"2cb8ed": 6639, // SonicWall
		"2cbaba": 152,  // Samsung
		"2cbc87": 545,  // Apple
		"2cbe08": 545,  // Apple
		"2cbeeb": 6884, // Nothing
		"2cc260": 11,   // Oracle
		"2cc407": 6885, // machineQ
		"2cc546": 3334, // Huawei
		"2cc548": 6886, // IAdea
		"2cc5d3": 2737, // Ruckus
		"2cc81b": 1784, // Routerboard.com
		"2ccc15": 457,  // Nokia
		"2ccc44": 101,  // Sony
		"2ccd27": 6887, // Precor
		"2ccd69": 6888, // Aqavi.com
		"2cce1e": 6889, // Cloudtronics
		"2ccf58": 3334, // Huawei
		"2cd02d": 2,    // Cisco
		"2cd05a": 2873, // Liteon
		"2cd066": 5697, // Xiaomi
		"2cd1da": 3599, // Keysight
		"2cd26b": 6440, // FN-Link
		"2cd2e7": 457,  // Nokia
		"2cd444": 4,    // Fujitsu
		"2cdb07": 421,  // Intel
		"2cdcad": 1592, // Wistron
		"2cdcd7": 3004, // Azurewave
		"2cdd0c": 6890, // Discovergy
		"2cdde9": 3793, // Arista
		"2ce2a8": 6891, // DeviceDesign
		"2ce310": 6892, // Stratacache
		"2ce412": 2048, // Sagemcom
		"2ce6cc": 2737, // Ruckus
		"2cea7f": 102,  // Dell
		"2ceadc": 2544, // Askey
		"2cf05d": 1812, // Micro-Star
		"2cf0a2": 545,  // Apple
		"2cf0ee": 545,  // Apple
		"2cf1bb": 3031, // ZTE
		"2cf295": 3334, // Huawei
		"2cf432": 6376, // Espressif
		"2cf4c5": 620,  // Avaya
		"2cf7f1": 6893, // Seeed
		"2cf89b": 2,    // Cisco
		"2cfda1": 1806, // ASUS
		"2cfdab": 687,  // Motorola
		"2cfdb3": 6894, // Tonly
		"2cffee": 6369, // Vivo
		"30053f": 6895, // JTI
		"30055c": 3230, // Brother
		"30074d": 152,  // Samsung
		"3009c0": 687,  // Motorola
		"300c23": 3031, // ZTE
		"300d43": 612,  // Microsoft
		"300d9e": 5440, // Ruijie
		"300eb8": 869,  // LG
		"300ee3": 3253, // Aquantia
		"3010b3": 2873, // Liteon
		"3010e4": 545,  // Apple
		"301389": 300,  // Siemens
		"30142d": 6896, // Piciorgros
		"30144a": 1592, // Wistron
		"301518": 6897, // Ubiquitous
		"30168d": 6898, // ProLon
		"3017c8": 101,  // Sony
		"301966": 152,  // Samsung
		"301a28": 6899, // Mako
		"301a30": 6899, // Mako
		"302303": 2468, // Belkin
		"302432": 421,  // Intel
		"302478": 2048, // Sagemcom
		"3024a9": 302,  // HP
		"3027cf": 67,   // Private
		"302952": 3777, // Hillstone
		"302bdc": 6900, // Top-Unum
		"302de8": 6901, // JDA
		"30317d": 1730, // Hosiden
		"3032d4": 6902, // Hanilstm
		"303335": 6903, // Boosty
		"303422": 5820, // eero
		"3034d2": 6904, // Availink
		"3035ad": 545,  // Apple
		"3035c5": 3334, // Huawei
		"3037a6": 2,    // Cisco
		"3037b3": 3334, // Huawei
		"303855": 457,  // Nokia
		"303926": 101,  // Sony
		"303a64": 421,  // Intel
		"303fbb": 302,  // HP
		"304225": 6905, // Burg-Wachter
		"304240": 3031, // ZTE
		"304449": 6906, // PLATH
		"304596": 3334, // Huawei
		"30469a": 1368, // Netgear
		"30499e": 3334, // Huawei
		"304b07": 687,  // Motorola
		"304c7e": 2153, // Panasonic
		"304e1b": 3334, // Huawei
		"3051f8": 6907, // BYK-Gardner
		"30525a": 2259, // NST
		"3052cb": 2873, // Liteon
		"3053c1": 3282, // Cresyn
		"305696": 6295, // Infinix
		"305714": 545,  // Apple
		"30578e": 5820, // eero
		"3057ac": 6908, // Irlab
		"30595b": 6909, // streamnow
		"3059b7": 612,  // Microsoft
		"305a3a": 1806, // ASUS
		"305d38": 6910, // Beissbarth
		"306023": 125,  // ARRIS
		"306112": 3936, // PAV
		"306118": 6911, // Paradom
		"30636b": 545,  // Apple
		"3065ec": 1592, // Wistron
		"3066d0": 3334, // Huawei
		"30688c": 6912, // Reach
		"30694b": 4555, // RIM
		"306a85": 152,  // Samsung
		"306cbe": 6913, // Skymotion
		"306e5c": 6914, // Validus
		"306f07": 6915, // Nations
		"307350": 6916, // Inpeco
		"307467": 152,  // Samsung
		"307496": 3334, // Huawei
		"307512": 101,  // Sony
		"30766f": 869,  // LG
		"3077cb": 6917, // Maike
		"3078d3": 6918, // Virgilant
		"307c30": 4555, // RIM
		"307c4a": 3334, // Huawei
		"307c5e": 826,  // Juniper
		"307ecb": 3187, // SFR
		"308398": 6376, // Espressif
		"3083d2": 687,  // Motorola
		"3085a9": 1806, // ASUS
		"3085eb": 4921, // Fiberhome
		"308730": 3334, // Huawei
		"3087d9": 2737, // Ruckus
		"308af7": 3334, // Huawei
		"308bb2": 2,    // Cisco
		"308cfb": 6919, // Dropcam
		"308d99": 302,  // HP
		"309048": 545,  // Apple
		"3090ab": 545,  // Apple
		"30918f": 6392, // Technicolor
		"3093bc": 2048, // Sagemcom
		"309435": 6369, // Vivo
		"309610": 3334, // Huawei
		"3096fb": 152,  // Samsung
		"309935": 3031, // ZTE
		"309c23": 1812, // Micro-Star
		"309e1d": 6577, // Ohsung
		"309ffb": 6920, // Ardomus
		"30a176": 4921, // Fiberhome
		"30a1fa": 3334, // Huawei
		"30a2c2": 3334, // Huawei
		"30a8db": 101,  // Sony
		"30a998": 3334, // Huawei
		"30a9de": 869,  // LG
		"30aae4": 3334, // Huawei
		"30ab6a": 152,  // Samsung
		"30aea4": 6376, // Espressif
		"30afce": 6369, // Vivo
		"30b164": 6921, // Power
		"30b1b5": 2639, // Arcadyan
		"30b49e": 1595, // TP-Link
		"30b4b8": 869,  // LG
		"30b5c2": 1595, // TP-Link
		"30b5f1": 6922, // Aitexin
		"30b62d": 2485, // Mojo
		"30b64f": 826,  // Juniper
		"30b7d4": 870,  // Hitron
		"30b930": 3031, // ZTE
		"30bb7d": 6923, // OnePlus
		"30c3d9": 430,  // Alpsalpine
		"30c50f": 3334, // Huawei
		"30c6f7": 6376, // Espressif
		"30c7ae": 152,  // Samsung
		"30cbc7": 665,  // Cambium
		"30cbf8": 152,  // Samsung
		"30cc21": 3031, // ZTE
		"30cda7": 152,  // Samsung
		"30d042": 102,  // Dell
		"30d16b": 2873, // Liteon
		"30d17e": 3334, // Huawei
		"30d32d": 1637, // devolo
		"30d357": 6924, // Logosol
		"30d386": 3031, // ZTE
		"30d46a": 6925, // Autosales
		"30d53e": 545,  // Apple
		"30d587": 152,  // Samsung
		"30d659": 6926, // Merging
		"30d6c9": 152,  // Samsung
		"30d9d9": 545,  // Apple
		"30e090": 6927, // Linctronix
		"30e171": 302,  // HP
		"30e37a": 421,  // Intel
		"30e396": 3334, // Huawei
		"30e4db": 2,    // Cisco
		"30e98e": 3334, // Huawei
		"30ea26": 6928, // Sycada
		"30f31d": 3031, // ZTE
		"30f335": 3334, // Huawei
		"30f42f": 4385, // ESP
		"30f70d": 2,    // Cisco
		"30f7c5": 545,  // Apple
		"30f7d7": 6929, // Thread
		"30f94b": 3857, // Universal
		"30f9ed": 101,  // Sony
		"30fbb8": 3334, // Huawei
		"30fc68": 1595, // TP-Link
		"30fceb": 869,  // LG
		"30fd11": 6930, // Macrotech
		"30fd38": 3521, // Google
		"30fd65": 3334, // Huawei
		"30fe31": 457,  // Nokia
		"3400a3": 3334, // Huawei
		"340286": 421,  // Intel
		"34029b": 6931, // Plexonics
		"34074f": 6932, // AccelStor
		"3407fb": 306,  // Ericsson
		"340804": 803,  // D-Link
		"3408bc": 545,  // Apple
		"340a22": 6933, // TOP-Access
		"340a33": 803,  // D-Link
		"340a98": 3334, // Huawei
		"340ced": 6934, // Moduel
		"341290": 6935, // Treeview
		"341298": 545,  // Apple
		"3412f9": 3334, // Huawei
		"3413a8": 6936, // Mediplan
		"3413e8": 421,  // Intel
		"34145f": 152,  // Samsung
		"34159e": 545,  // Apple
		"3417eb": 102,  // Dell
		"341a35": 4921, // Fiberhome
		"341b22": 6937, // Grandbeing
		"341cf0": 5697, // Xiaomi
		"341e6b": 3334, // Huawei
		"341fe4": 125,  // ARRIS
		"3420e3": 2737, // Ruckus
		"3423ba": 152,  // Samsung
		"34243e": 3031, // ZTE
		"342606": 6938, // CarePredict
		"342840": 545,  // Apple
		"3428f0": 6939, // ATN
		"342912": 3334, // Huawei
		"3429ea": 6940, // MCD
		"342b70": 6941, // Arris
		"342cc4": 358,  // Compal
		"342d0d": 152,  // Samsung
		"342eb6": 3334, // Huawei
		"342eb7": 421,  // Intel
		"342f6e": 6942, // Anywire
		"342fbd": 1427, // Nintendo
		"343111": 152,  // Samsung
		"34317f": 2153, // Panasonic
		"34318f": 545,  // Apple
		"3431c4": 621,  // AVM
		"3432e6": 2153, // Panasonic
		"34363b": 545,  // Apple
		"343654": 3031, // ZTE
		"343759": 3031, // ZTE
		"343794": 6943, // Hamee
		"3438af": 6944, // Inlab
		"3438b7": 531,  // HUMAX
		"343d98": 6945, // JinQianMao
		"343dc4": 1077, // Buffalo
		"343ea4": 6946, // Ring
		"3440b5": 372,  // IBM
		"34415d": 421,  // Intel
		"3441a8": 6947, // ER-Telecom
		"344262": 545,  // Apple
		"34466f": 6948, // HiTEM
		"3446ec": 3334, // Huawei
		"3448ed": 102,  // Dell
		"34495b": 2048, // Sagemcom
		"344b3d": 4921, // Fiberhome
		"344b50": 3031, // ZTE
		"344ca4": 6949, // amazipoint
		"344cc8": 6950, // Echodyne
		"344dea": 3031, // ZTE
		"344df7": 869,  // LG
		"344f3f": 6951, // IO-Power
		"344f5c": 6952, // R&M
		"345184": 3334, // Huawei
		"3451c9": 545,  // Apple
		"3453d2": 2048, // Sagemcom
		"345760": 6428, // MitraStar
		"345840": 3334, // Huawei
		"345a06": 3205, // Sharp
		"345c40": 6953, // Cargt
		"345d10": 6954, // Wytek
		"345d9e": 2048, // Sagemcom
		"3460f9": 1595, // TP-Link
		"346178": 1041, // Boeing
		"346288": 2,    // Cisco
		"3462b4": 5695, // Renesas
		"3464a9": 302,  // HP
		"3466ea": 6955, // Vertu
		"34684a": 6956, // Teraworks
		"346987": 3031, // ZTE
		"346ac2": 3334, // Huawei
		"346b46": 2048, // Sagemcom
		"346bd3": 3334, // Huawei
		"346d9c": 370,  // Carrier
		"346e8a": 6957, // Ecosense
		"346e9d": 306,  // Ericsson
		"346f24": 3004, // Azurewave
		"346f90": 2,    // Cisco
		"347146": 3334, // Huawei
		"34732d": 2,    // Cisco
		"34735a": 102,  // Dell
		"3475c7": 620,  // Avaya
		"347839": 3031, // ZTE
		"347877": 6958, // O-Net
		"347916": 3334, // Huawei
		"347a60": 125,  // ARRIS
		"347c25": 545,  // Apple
		"347df6": 421,  // Intel
		"347e00": 3334, // Huawei
		"347e39": 457,  // Nokia
		"347e5c": 2047, // Sonos
		"347eca": 6959, // Nextwill
		"34800d": 2249, // Cavium
		"3480b3": 5697, // Xiaomi
		"348137": 6960, // Unicard
		"3481c4": 621,  // AVM
		"3482c5": 152,  // Samsung
		"3482de": 6961, // Kiio
		"348302": 6962, // iFORCOM
		"348446": 306,  // Ericsson
		"348584": 185,  // Extreme
		"34865d": 6376, // Espressif
		"348a12": 1685, // Aruba
		"348a7b": 152,  // Samsung
		"348aae": 2048, // Sagemcom
		"348b75": 647,  // Lava
		"348f27": 2737, // Ruckus
		"34916f": 6963, // UserGate
		"349342": 6964, // TTE
		"349454": 6376, // Espressif
		"3495db": 242,  // Logitec
		"349672": 1595, // TP-Link
		"34976f": 6965, // Rootech
		"3497f6": 1806, // ASUS
		"3498b5": 1368, // Netgear
		"34996f": 6966, // VPI
		"349b5b": 6967, // Maquet
		"349d90": 6968, // Heinzmann
		"349e34": 6969, // Evervictory
		"349f7b": 87,   // Canon
		"34a183": 6970, // AWare
		"34a2a2": 3334, // Huawei
		"34a395": 545,  // Apple
		"34a3bf": 6971, // Terewave
		"34a5b4": 6972, // Navtech
		"34a7ba": 6973, // Fischer
		"34a843": 2848, // Kyocera
		"34a84e": 2,    // Cisco
		"34a8eb": 545,  // Apple
		"34aa8b": 152,  // Samsung
		"34aa99": 457,  // Nokia
		"34ab37": 545,  // Apple
		"34ab95": 6376, // Espressif
		"34af2c": 1427, // Nintendo
		"34afb3": 5438, // Amazon
		"34b20a": 3334, // Huawei
		"34b354": 3334, // Huawei
		"34b472": 6376, // Espressif
		"34b571": 6974, // Plds
		"34b883": 2,    // Cisco
		"34b98d": 5697, // Xiaomi
		"34ba75": 6975, // Everest
		"34ba9a": 6976, // Asiatelco
		"34bb1f": 2220, // BlackBerry
		"34bb26": 687,  // Motorola
		"34bdc8": 2,    // Cisco
		"34be00": 152,  // Samsung
		"34bf90": 4921, // Fiberhome
		"34c059": 545,  // Apple
		"34c3ac": 152,  // Samsung
		"34c3d2": 6440, // FN-Link
		"34c69a": 6977, // Enecsys
		"34c731": 430,  // Alpsalpine
		"34c803": 457,  // Nokia
		"34c93d": 421,  // Intel
		"34c99d": 6978, // Eidolon
		"34c9f0": 6979, // LM
		"34cc28": 6980, // Nexpring
		"34cd6d": 6981, // CommSky
		"34cdbe": 3334, // Huawei
		"34ce00": 6280, // XIAOMI
		"34ce94": 6982, // Parsec
		"34cff6": 421,  // Intel
		"34d09b": 6983, // MobilMAX
		"34d270": 5438, // Amazon
		"34d693": 3334, // Huawei
		"34d7b4": 6984, // Tributary
		"34d954": 6985, // WiBotic
		"34dab7": 3031, // ZTE
		"34db9c": 2048, // Sagemcom
		"34dbfd": 2,    // Cisco
		"34de1a": 421,  // Intel
		"34de34": 3031, // ZTE
		"34df2a": 6986, // Fujikon
		"34e0cf": 3031, // ZTE
		"34e12d": 421,  // Intel
		"34e2fd": 545,  // Apple
		"34e6ad": 421,  // Intel
		"34e6d7": 102,  // Dell
		"34e70b": 6987, // HAN
		"34e894": 1595, // TP-Link
		"34e911": 6369, // Vivo
		"34e9fe": 6988, // Metis
		"34ed1b": 2,    // Cisco
		"34ef44": 1939, // 2Wire
		"34ef8b": 6989, // NTT
		"34efb6": 6294, // Edgecore
		"34f39a": 421,  // Intel
		"34f39b": 6990, // WizLAN
		"34f62d": 3205, // Sharp
		"34f64b": 421,  // Intel
		"34f6d2": 2153, // Panasonic
		"34f716": 1595, // TP-Link
		"34f8e7": 2,    // Cisco
		"34f968": 6991, // ATEK
		"34fa9f": 2737, // Ruckus
		"34fc6f": 6992, // Alcea
		"34fcb9": 302,  // HP
		"34fcef": 869,  // LG
		"34fd6a": 545,  // Apple
		"34fe77": 545,  // Apple
		"34fe9e": 4,    // Fujitsu
		"380025": 421,  // Intel
		"380118": 6993, // ULVAC
		"380195": 152,  // Samsung
		"3802de": 2074, // Sercomm
		"3806b4": 6994, // A.D.C
		"3807d4": 6995, // Zeppelin
		"3808fd": 6996, // Silca
		"380a0a": 6997, // Sky-City
		"380a94": 152,  // Samsung
		"380aab": 6998, // Formlabs
		"380b40": 152,  // Samsung
		"380dd4": 389,  // Primax
		"380e4d": 2,    // Cisco
		"380f4a": 545,  // Apple
		"3810f0": 1685, // Aruba
		"381428": 102,  // Dell
		"38144e": 4921, // Fiberhome
		"3816d1": 152,  // Samsung
		"381766": 6999, // Promzakaz
		"3817c3": 302,  // HP
		"3817e1": 6392, // Technicolor
		"38184c": 101,  // Sony
		"38192f": 457,  // Nokia
		"381a52": 46,   // Epson
		"381c1a": 2,    // Cisco
		"381c23": 5782, // Hilan
		"381d14": 7000, // Skydio
		"381dd9": 6440, // FN-Link
		"381ec7": 6439, // Chipsea
		"382028": 3334, // Huawei
		"382056": 2,    // Cisco
		"3820a8": 7001, // ColorTokens
		"3821c7": 1685, // Aruba
		"3822e2": 302,  // HP
		"38256b": 612,  // Microsoft
		"38262b": 7002, // UTran
		"3826cd": 7003, // Andtek
		"3829dd": 7004, // ONvocal
		"382a19": 7005, // Technica
		"382c4a": 1806, // ASUS
		"382dd1": 152,  // Samsung
		"382de8": 152,  // Samsung
		"3830f9": 869,  // LG
		"3831ac": 7006, // WEG
		"3835fb": 2048, // Sagemcom
		"38378b": 3334, // Huawei
		"38384b": 6369, // Vivo
		"383bc8": 1939, // 2Wire
		"383d5b": 4921, // Fiberhome
		"383f10": 7007, // DBL
		"383fb3": 6392, // Technicolor
		"38420b": 2047, // Sonos
		"38437d": 358,  // Compal
		"3843e5": 7008, // Grotech
		"38453b": 2737, // Ruckus
		"38454c": 7009, // Light
		"38458c": 7010, // MyCloud
		"384608": 3031, // ZTE
		"3847bc": 3334, // Huawei
		"38484c": 545,  // Apple
		"384b24": 300,  // Siemens
		"384b5b": 7011, // Ztron
		"384c4f": 3334, // Huawei
		"384c90": 125,  // ARRIS
		"384f49": 826,  // Juniper
		"384ff0": 3004, // Azurewave
		"38521a": 457,  // Nokia
		"385247": 3334, // Huawei
		"38539c": 545,  // Apple
		"38549b": 3031, // ZTE
		"38563d": 612,  // Microsoft
		"38580c": 7012, // Panaccess
		"385b44": 1655, // Silicon
		"385c76": 542,  // Plantronics
		"386077": 5439, // Pegatron
		"3861a5": 7013, // Grabango
		"3863bb": 302,  // HP
		"3865b2": 545,  // Apple
		"386645": 7014, // OOSIC
		"3866f0": 545,  // Apple
		"386893": 421,  // Intel
		"3868a4": 152,  // Samsung
		"3868dd": 3978, // Inventec
		"386a77": 152,  // Samsung
		"386bbb": 125,  // ARRIS
		"386e88": 3031, // ZTE
		"386ea2": 6369, // Vivo
		"38700c": 125,  // ARRIS
		"3871de": 545,  // Apple
		"3872c0": 3869, // Comtrend
		"3876d1": 7015, // Euronda
		"387862": 101,  // Sony
		"387a0e": 421,  // Intel
		"387a3c": 4921, // Fiberhome
		"387b47": 7016, // AKELA
		"3880df": 687,  // Motorola
		"388345": 1595, // TP-Link
		"388602": 7017, // Flexoptix
		"3887d5": 421,  // Intel
		"38881e": 3334, // Huawei
		"3888a4": 545,  // Apple
		"38892c": 545,  // Apple
		"388ab7": 7018, // ITC
		"388b59": 3521, // Google
		"388c50": 869,  // LG
		"388e7a": 7019, // Autoit
		"388ee7": 7020, // Fanhattan
		"389052": 3334, // Huawei
		"3890a5": 2,    // Cisco
		"3891fb": 7021, // Xenox
		"389461": 5695, // Renesas
		"389496": 152,  // Samsung
		"3894e0": 7022, // Syrotech
		"3894ed": 1368, // Netgear
		"3897a4": 5693, // Elecom
		"3898d8": 7023, // Meritech
		"3898e9": 3334, // Huawei
		"389af6": 152,  // Samsung
		"389d92": 46,   // Epson
		"38a4ed": 5697, // Xiaomi
		"38a659": 2048, // Sagemcom
		"38a6ce": 3512, // BSkyB
		"38a851": 1439, // Moog
		"38a86b": 7024, // Orga
		"38a95f": 7025, // Actifio
		"38a9ea": 67,   // Private
		"38aa3c": 152,  // Samsung
		"38ac3d": 7026, // Nephos
		"38afd0": 7027, // Nevro
		"38afd7": 4,    // Fujitsu
		"38b3f7": 3334, // Huawei
		"38b54d": 545,  // Apple
		"38b5d3": 7028, // SecuWorks
		"38b725": 1592, // Wistron
		"38b74d": 7029, // Fijowave
		"38b800": 1592, // Wistron
		"38bab0": 858,  // Broadcom
		"38baf8": 421,  // Intel
		"38bb3c": 620,  // Avaya
		"38bc01": 3334, // Huawei
		"38bc1a": 7030, // MEIZU
		"38beab": 7031, // AltoBeam
		"38bf2f": 7032, // Espec
		"38c096": 430,  // Alpsalpine
		"38c70a": 7033, // WiFiSong
		"38c7ba": 3362, // CS
		"38c986": 545,  // Apple
		"38cada": 545,  // Apple
		"38d40b": 152,  // Samsung
		"38d547": 1806, // ASUS
		"38d7ca": 7034, // 7HUGS
		"38d82f": 3031, // ZTE
		"38de60": 7035, // Mohlenhoff
		"38dead": 421,  // Intel
		"38e1aa": 3031, // ZTE
		"38e2dd": 3031, // ZTE
		"38e39f": 687,  // Motorola
		"38e60a": 5697, // Xiaomi
		"38e7d8": 1341, // HTC
		"38eaa7": 302,  // HP
		"38eb47": 3334, // Huawei
		"38ec0d": 545,  // Apple
		"38ece4": 152,  // Samsung
		"38ed18": 2,    // Cisco
		"38ee9d": 7036, // Anedo
		"38f0c8": 7037, // Mevo
		"38f135": 7038, // SensorTec-Canada
		"38f23e": 612,  // Microsoft
		"38f32e": 7039, // Skullcandy
		"38f33f": 7040, // Tatsuno
		"38f3ab": 4919, // LCFC
		"38f3fb": 7041, // Asperiq
		"38f557": 7042, // Jolata
		"38f597": 7043, // home2net
		"38f73d": 5438, // Amazon
		"38f7f1": 3334, // Huawei
		"38f85e": 531,  // HUMAX
		"38f889": 3334, // Huawei
		"38f8ca": 7044, // OWIN
		"38f9d3": 545,  // Apple
		"38fb14": 3334, // Huawei
		"38fc98": 421,  // Intel
		"38fdf8": 2,    // Cisco
		"38ff36": 2737, // Ruckus
		"3c01ef": 101,  // Sony
		"3c02b1": 7045, // Creation
		"3c0461": 125,  // ARRIS
		"3c04bf": 7046, // PRAVIS
		"3c0518": 152,  // Samsung
		"3c0630": 545,  // Apple
		"3c06a7": 1595, // TP-Link
		"3c0754": 545,  // Apple
		"3c0771": 101,  // Sony
		"3c08f6": 2,    // Cisco
		"3c0c48": 7047, // Servergy
		"3c0cdb": 3505, // Unionman
		"3c0e23": 2,    // Cisco
		"3c0fc1": 7048, // KBC
		"3c106f": 7049, // Albahith
		"3c10e6": 7050, // PHAZR
		"3c13cc": 2,    // Cisco
		"3c15c2": 545,  // Apple
		"3c15ea": 7051, // Tescom
		"3c15fb": 3334, // Huawei
		"3c1710": 2048, // Sagemcom
		"3c189f": 457,  // Nokia
		"3c195e": 152,  // Samsung
		"3c197d": 306,  // Ericsson
		"3c1a57": 7052, // Cardiopulmonary
		"3c1a79": 7053, // Huayuan
		"3c1a9e": 7054, // VitalThings
		"3c1cbe": 7055, // Jadak
		"3c1e04": 803,  // D-Link
		"3c20f6": 152,  // Samsung
		"3c219c": 421,  // Intel
		"3c22fb": 545,  // Apple
		"3c25d7": 457,  // Nokia
		"3c286d": 3521, // Google
		"3c2af4": 3230, // Brother
		"3c2c30": 102,  // Dell
		"3c2c99": 6294, // Edgecore
		"3c2ef9": 545,  // Apple
		"3c2eff": 545,  // Apple
		"3c2f3a": 7056, // SFORZATO
		"3c300c": 7057, // Dewar
		"3c306f": 3334, // Huawei
		"3c3178": 7058, // Qolsys
		"3c3556": 7059, // Cognitec
		"3c363d": 457,  // Nokia
		"3c36e4": 125,  // ARRIS
		"3c3786": 1368, // Netgear
		"3c3888": 7060, // ConnectQuest
		"3c38f4": 101,  // Sony
		"3c39c3": 7061, // JW
		"3c3a73": 620,  // Avaya
		"3c3f51": 7062, // 2CRSI
		"3c410e": 2,    // Cisco
		"3c438e": 125,  // ARRIS
		"3c457a": 3512, // BSkyB
		"3c46d8": 1595, // TP-Link
		"3c4711": 3334, // Huawei
		"3c4937": 7063, // ASSMANN
		"3c4a92": 302,  // HP
		"3c4cd0": 1490, // Ceragon
		"3c4dbe": 545,  // Apple
		"3c4e47": 7064, // Etronic
		"3c510e": 2,    // Cisco
		"3c5282": 302,  // HP
		"3c53d7": 7065, // Cedes
		"3c5447": 3334, // Huawei
		"3c5731": 2,    // Cisco
		"3c576c": 152,  // Samsung
		"3c57d5": 7066, // FiveCo
		"3c58c2": 421,  // Intel
		"3c5a37": 152,  // Samsung
		"3c5ab4": 3521, // Google
		"3c5cc4": 5438, // Amazon
		"3c5cf1": 5820, // eero
		"3c5ec3": 2,    // Cisco
		"3c5f01": 7067, // Synerchip
		"3c6104": 826,  // Juniper
		"3c6105": 6376, // Espressif
		"3c6200": 152,  // Samsung
		"3c62f0": 2074, // Sercomm
		"3c672c": 7068, // Sciovid
		"3c678c": 3334, // Huawei
		"3c6816": 4742, // VXi
		"3c6a9d": 7069, // Dexatek
		"3c6aa7": 421,  // Intel
		"3c6e63": 3884, // Mitron
		"3c6f45": 7070, // Fiberpro
		"3c6fea": 2153, // Panasonic
		"3c6ff7": 7071, // EnTek
		"3c7059": 7072, // MakerBot
		"3c71bf": 6376, // Espressif
		"3c7437": 4555, // RIM
		"3c754a": 125,  // ARRIS
		"3c7843": 3334, // Huawei
		"3c7873": 7073, // Airsonics
		"3c7a8a": 125,  // ARRIS
		"3c7ac4": 7074, // Chemtronics
		"3c7af0": 6281, // Itel
		"3c7c3f": 1806, // ASUS
		"3c7d0a": 545,  // Apple
		"3c7f6f": 7075, // Telechips
		"3c81d8": 2048, // Sagemcom
		"3c831e": 7076, // CKD
		"3c8375": 612,  // Microsoft
		"3c846a": 1595, // TP-Link
		"3c869a": 3334, // Huawei
		"3c86d1": 6369, // Vivo
		"3c8970": 7077, // Neosfar
		"3c8994": 3512, // BSkyB
		"3c89a6": 7078, // Kapelse
		"3c8ab0": 826,  // Juniper
		"3c8b7f": 2,    // Cisco
		"3c8bfe": 152,  // Samsung
		"3c8c93": 826,  // Juniper
		"3c8cf8": 2904, // TRENDnet
		"3c8d20": 3521, // Google
		"3c9066": 4549, // SmartRG
		"3c912b": 7079, // Vexata
		"3c9157": 3096, // Yulong
		"3c9174": 7080, // Along
		"3c9180": 2873, // Liteon
		"3c92dc": 7081, // Octopod
		"3c93f4": 3334, // Huawei
		"3c94d5": 826,  // Juniper
		"3c9509": 2873, // Liteon
		"3c970e": 1592, // Wistron
		"3c977e": 7082, // IPS
		"3c9872": 2074, // Sercomm
		"3c99f7": 7083, // Lansentechnology
		"3c9a77": 6392, // Technicolor
		"3c9bc6": 3334, // Huawei
		"3c9bd6": 3467, // Vizio
		"3c9c0f": 421,  // Intel
		"3c9d56": 3334, // Huawei
		"3c9ec7": 3512, // BSkyB
		"3ca067": 2873, // Liteon
		"3ca10d": 152,  // Samsung
		"3ca161": 3334, // Huawei
		"3ca31a": 7084, // Oilfind
		"3ca348": 6369, // Vivo
		"3ca37e": 3334, // Huawei
		"3ca581": 6369, // Vivo
		"3ca616": 6369, // Vivo
		"3ca6f6": 545,  // Apple
		"3ca72b": 2253, // MRV
		"3ca82a": 302,  // HP
		"3ca916": 3334, // Huawei
		"3ca9f4": 421,  // Intel
		"3caa3f": 7085, // iKey
		"3cab8e": 545,  // Apple
		"3cb15b": 620,  // Avaya
		"3cb233": 3334, // Huawei
		"3cb6b7": 6369, // Vivo
		"3cb72b": 7086, // PLUMgrid
		"3cb74b": 6392, // Technicolor
		"3cb87a": 67,   // Private
		"3cbbfd": 152,  // Samsung
		"3cbdc5": 2639, // Arcadyan
		"3cbdd8": 869,  // LG
		"3cbee1": 5661, // Nikon
		"3cbf60": 545,  // Apple
		"3cc12c": 7087, // AES
		"3cc1f6": 7088, // Melange
		"3cc243": 457,  // Nokia
		"3cc99e": 7089, // Huiyang
		"3cca87": 7090, // Iders
		"3ccb7c": 6433, // TCT
		"3ccd36": 545,  // Apple
		"3ccd5d": 3334, // Huawei
		"3ccd93": 869,  // LG
		"3cce73": 2,    // Cisco
		"3cd0f8": 545,  // Apple
		"3cd16e": 7091, // Telepower
		"3cd4d6": 7092, // WirelessWERX
		"3cd92b": 302,  // HP
		"3cda2a": 3031, // ZTE
		"3cda6d": 7093, // Tiandy
		"3cdcbc": 152,  // Samsung
		"3cdf1e": 2,    // Cisco
		"3cdfa9": 125,  // ARRIS
		"3cdfbd": 3334, // Huawei
		"3ce038": 7094, // Omnifi
		"3ce072": 545,  // Apple
		"3ce624": 869,  // LG
		"3ce824": 3334, // Huawei
		"3ce9f7": 421,  // Intel
		"3cea4f": 1939, // 2Wire
		"3ceaf9": 7095, // Jubix
		"3ceafb": 7096, // NSE
		"3cf011": 421,  // Intel
		"3cf392": 7097, // Virtualtek
		"3cf4f9": 7098, // Moda-InnoChips
		"3cf52c": 7099, // DSPECIALISTS
		"3cf652": 3031, // ZTE
		"3cf72a": 457,  // Nokia
		"3cf7a4": 152,  // Samsung
		"3cf7d1": 7100, // OMRON
		"3cf808": 3334, // Huawei
		"3cf862": 421,  // Intel
		"3cfa43": 3334, // Huawei
		"3cfb5c": 4921, // Fiberhome
		"3cfb96": 7101, // Emcraft
		"3cfdfe": 421,  // Intel
		"3cffd8": 3334, // Huawei
		"4000e0": 7102, // DerekLimited
		"400107": 3793, // Arista
		"40017a": 2,    // Cisco
		"4001c6": 160,  // 3Com
		"40040c": 7103, // A&T
		"400589": 7104, // T-Mobile
		"400634": 3334, // Huawei
		"4006d5": 2,    // Cisco
		"4007c0": 7105, // Railtec
		"400d10": 125,  // ARRIS
		"400e67": 7106, // Tremol
		"400e85": 152,  // Samsung
		"4011c3": 152,  // Samsung
		"4011dc": 7107, // Sonance
		"4012e4": 7108, // Compass-EOS
		"4014ad": 3334, // Huawei
		"40163b": 152,  // Samsung
		"40167e": 1806, // ASUS
		"40169f": 1595, // TP-Link
		"4017e2": 7109, // Intai
		"4018b1": 185,  // Extreme
		"4018d7": 775,  // Smartronix
		"401920": 2722, // Movon
		"401c83": 421,  // Intel
		"4025c2": 421,  // Intel
		"402619": 545,  // Apple
		"40270b": 7110, // Mobileeco
		"402814": 7111, // RFI
		"402b50": 125,  // ARRIS
		"402ba1": 101,  // Sony
		"402e28": 4538, // MiXTelematics
		"402f86": 869,  // LG
		"403004": 545,  // Apple
		"403067": 7112, // Conlog
		"40313c": 6280, // XIAOMI
		"40331a": 545,  // Apple
		"403cfc": 545,  // Apple
		"403dec": 531,  // HUMAX
		"403f8c": 1595, // TP-Link
		"404022": 7113, // ZIV
		"404028": 7113, // ZIV
		"40406b": 7114, // Icomera
		"40406c": 7114, // Icomera
		"4040a7": 101,  // Sony
		"404229": 7115, // Layer3TV
		"4045da": 7116, // Spreadtrum
		"40498a": 7117, // Synapticon
		"404a03": 2692, // ZyXEL
		"404ad4": 7118, // Widex
		"404c77": 125,  // ARRIS
		"404d7f": 545,  // Apple
		"404d8e": 3334, // Huawei
		"404e36": 1341, // HTC
		"40516c": 7119, // Grandex
		"40520d": 456,  // Pico
		"4054e4": 7120, // Wearsafe
		"405539": 2,    // Cisco
		"405582": 457,  // Nokia
		"405662": 7121, // GuoTengShengHua
		"405a9b": 7122, // Anovo
		"405cfd": 102,  // Dell
		"405d82": 1368, // Netgear
		"405f7d": 6433, // TCT
		"405fbe": 4555, // RIM
		"40605a": 7123, // Hawkeye
		"406186": 1812, // Micro-Star
		"40618e": 7124, // Stella-Green
		"406231": 7125, // Gifa
		"4065a3": 2048, // Sagemcom
		"406aab": 4555, // RIM
		"406c8f": 545,  // Apple
		"406f2a": 2220, // BlackBerry
		"407009": 125,  // ARRIS
		"407074": 7126, // Life
		"4070f5": 545,  // Apple
		"407183": 826,  // Juniper
		"407496": 7127, // aFUN
		"4074e0": 421,  // Intel
		"4076a9": 3334, // Huawei
		"40786a": 687,  // Motorola
		"407a80": 457,  // Nokia
		"407b1b": 7128, // Mettle
		"407c7d": 457,  // Nokia
		"407d0f": 3334, // Huawei
		"40831d": 545,  // Apple
		"4083de": 768,  // Zebra
		"408493": 7129, // Clavister
		"408805": 687,  // Motorola
		"40882f": 185,  // Extreme
		"4089a8": 7130, // WiredIQ
		"408a9a": 7131, // TITENG
		"408b07": 2246, // Actiontec
		"408d5c": 1929, // Giga-Byte
		"408f9d": 826,  // Juniper
		"409151": 6376, // Espressif
		"409505": 7132, // Acoinfo
		"409558": 7133, // Aisino
		"4095bd": 7134, // NTmore
		"4097d1": 7135, // BK
		"40984c": 7136, // Casacom
		"40987b": 7133, // Aisino
		"4098ad": 545,  // Apple
		"409922": 3004, // Azurewave
		"409b21": 457,  // Nokia
		"409bcd": 803,  // D-Link
		"409c28": 545,  // Apple
		"409ca6": 7137, // Curvalux
		"409f38": 3004, // Azurewave
		"409f87": 7138, // Jide
		"40a108": 687,  // Motorola
		"40a2db": 5438, // Amazon
		"40a3cc": 421,  // Intel
		"40a677": 826,  // Juniper
		"40a6a4": 7139, // PassivSystems
		"40a6b7": 421,  // Intel
		"40a6d9": 545,  // Apple
		"40a6e8": 2,    // Cisco
		"40a8f0": 302,  // HP
		"40a9cf": 5438, // Amazon
		"40b034": 302,  // HP
		"40b076": 1806, // ASUS
		"40b0a1": 6015, // Valcom
		"40b0fa": 869,  // LG
		"40b2c8": 76,   // Nortel
		"40b31e": 3857, // Universal
		"40b395": 545,  // Apple
		"40b3cd": 7140, // Chiyoda
		"40b3fc": 7141, // Logital
		"40b4cd": 5438, // Amazon
		"40b4f0": 826,  // Juniper
		"40b5c1": 2,    // Cisco
		"40b6b1": 7142, // SUNGSAM
		"40b6e7": 3334, // Huawei
		"40b7f3": 125,  // ARRIS
		"40b837": 101,  // Sony
		"40b93c": 302,  // HP
		"40ba61": 1957, // ARIMA
		"40bc60": 545,  // Apple
		"40bc8b": 7143, // itelio
		"40bd9e": 7144, // Physio-Control
		"40c3c6": 7145, // SnapRoute
		"40c48c": 7146, // N-iTUS
		"40c711": 545,  // Apple
		"40c729": 2048, // Sagemcom
		"40c7c9": 7147, // Naviit
		"40ca63": 7148, // Seongji
		"40cba8": 3334, // Huawei
		"40cbc0": 545,  // Apple
		"40cd3a": 7149, // Z3
		"40ce24": 2,    // Cisco
		"40d25f": 6281, // Itel
		"40d28a": 1427, // Nintendo
		"40d32d": 545,  // Apple
		"40d357": 7150, // Ison
		"40d3ae": 152,  // Samsung
		"40d40e": 7151, // Biodata
		"40d4bd": 7152, // SK
		"40d63c": 7153, // Equitech
		"40dc9d": 7154, // Hajen
		"40dca5": 3334, // Huawei
		"40dead": 826,  // Juniper
		"40e230": 3004, // Azurewave
		"40e3d6": 1685, // Aruba
		"40e64b": 545,  // Apple
		"40e99b": 152,  // Samsung
		"40ec99": 421,  // Intel
		"40ecf8": 300,  // Siemens
		"40ee15": 2127, // Zioncom
		"40eedd": 3334, // Huawei
		"40ef4c": 6298, // Fihonest
		"40f02f": 2873, // Liteon
		"40f078": 2,    // Cisco
		"40f201": 2048, // Sagemcom
		"40f2e9": 372,  // IBM
		"40f308": 2056, // Murata
		"40f407": 1427, // Nintendo
		"40f413": 7155, // Rubezh
		"40f4ec": 2,    // Cisco
		"40f4fd": 3505, // Unionman
		"40f520": 6376, // Espressif
		"40f6bc": 5438, // Amazon
		"40f946": 545,  // Apple
		"40f9d5": 7156, // Tecore
		"40fc89": 125,  // ARRIS
		"40fe0d": 7157, // Maxio
		"440010": 545,  // Apple
		"440049": 5438, // Amazon
		"44004d": 3334, // Huawei
		"44032c": 421,  // Intel
		"4403a7": 2,    // Cisco
		"44070b": 3521, // Google
		"4409b8": 7158, // Salcomp
		"440cfd": 7159, // NetMan
		"441102": 6661, // EDMI
		"441319": 7160, // WKK
		"4413d0": 3031, // ZTE
		"441441": 7161, // AudioControl
		"441622": 612,  // Microsoft
		"441793": 6376, // Espressif
		"44184f": 7162, // Fitview
		"4418fd": 545,  // Apple
		"441b88": 545,  // Apple
		"441c12": 6392, // Technicolor
		"441c7f": 687,  // Motorola
		"441e98": 2737, // Ruckus
		"441ea1": 302,  // HP
		"44227c": 3334, // Huawei
		"4422f1": 7163, // S.FAC
		"4423aa": 7164, // Farmage
		"4425bb": 7165, // Bamboo
		"4427f3": 7166, // 70mai
		"442938": 7167, // NietZsche
		"442a60": 545,  // Apple
		"442aff": 7168, // E3
		"442b03": 2,    // Cisco
		"442c05": 4498, // AMPAK
		"443192": 302,  // HP
		"44322a": 620,  // Avaya
		"4432c8": 6392, // Technicolor
		"44348f": 7169, // MXT
		"4434a7": 125,  // ARRIS
		"44356f": 7170, // Neterix
		"443583": 545,  // Apple
		"443839": 7171, // Cumulus
		"443b32": 3541, // Intelbras
		"443d21": 7172, // Nuvolt
		"443e07": 7173, // Electrolux
		"443eb2": 7174, // DEOTRON
		"44422f": 6250, // Testop
		"444450": 7175, // OttoQ
		"4448b9": 6428, // MitraStar
		"4448c1": 302,  // HP
		"444a65": 7176, // Silverflare
		"444adb": 545,  // Apple
		"444b7e": 4921, // Fiberhome
		"444c0c": 545,  // Apple
		"444ca8": 3793, // Arista
		"444e1a": 152,  // Samsung
		"444f8e": 7177, // WiZ
		"4455b1": 3334, // Huawei
		"4455c4": 3334, // Huawei
		"44568d": 7178, // PNC
		"4456b7": 7179, // Spawn
		"445943": 3031, // ZTE
		"44599f": 7180, // Criticare
		"4459e3": 3334, // Huawei
		"445bed": 1685, // Aruba
		"445ce9": 152,  // Samsung
		"445ecd": 7181, // Razer
		"446132": 7182, // ecobee
		"44619c": 7183, // FONsystem
		"446246": 7184, // Comat
		"44650d": 5438, // Amazon
		"44657f": 374,  // Calix
		"44666e": 7185, // IP-Line
		"446747": 3334, // Huawei
		"446752": 1592, // Wistron
		"44680c": 7186, // Wacom
		"4468ab": 7187, // Juin
		"446a2e": 3334, // Huawei
		"446ab7": 125,  // ARRIS
		"446c24": 7188, // Reallin
		"446d57": 2873, // Liteon
		"446d6c": 152,  // Samsung
		"446ee5": 3334, // Huawei
		"446ff8": 7189, // Dyson
		"44700b": 7190, // Iffu
		"4473d6": 4079, // Logitech
		"44746c": 101,  // Sony
		"447654": 3334, // Huawei
		"44783e": 152,  // Samsung
		"447bc4": 7191, // DualShine
		"447c7f": 7192, // Innolight
		"447e76": 7193, // Trek
		"4480eb": 687,  // Motorola
		"4482e5": 3334, // Huawei
		"448312": 7194, // Star-Net
		"448500": 421,  // Intel
		"4486c1": 300,  // Siemens
		"448723": 7195, // Hoya
		"4487fc": 1115, // Elitegroup
		"4488cb": 7196, // Camco
		"448a5b": 1812, // Micro-Star
		"448c52": 7197, // KTIS
		"448dbf": 7198, // Rhino
		"448e12": 7199, // DT
		"448e81": 7200, // VIG
		"448f17": 152,  // Samsung
		"4490bb": 545,  // Apple
		"449160": 2056, // Murata
		"4494fc": 1368, // Netgear
		"44962b": 7201, // Aidon
		"449bc1": 3334, // Huawei
		"449cb5": 7202, // Alcomp
		"449ef9": 6369, // Vivo
		"449f46": 3334, // Huawei
		"449f7f": 7203, // DataCore
		"44a191": 3334, // Huawei
		"44a42d": 6433, // TCT
		"44a54e": 6652, // Qorvo
		"44a56e": 1368, // Netgear
		"44a689": 7204, // Promax
		"44a6e5": 7205, // Thinking
		"44a7cf": 2056, // Murata
		"44a842": 102,  // Dell
		"44a8c2": 7206, // Sewoo
		"44a8fc": 545,  // Apple
		"44aa27": 7207, // udworks
		"44aa50": 826,  // Juniper
		"44aae8": 7208, // Nanotec
		"44aaf5": 125,  // ARRIS
		"44ad19": 7209, // Xingfei
		"44adb1": 2048, // Sagemcom
		"44add9": 2,    // Cisco
		"44ae25": 2,    // Cisco
		"44ae44": 3334, // Huawei
		"44af28": 421,  // Intel
		"44b32d": 1595, // TP-Link
		"44b412": 7210, // Sius
		"44b433": 7211, // tide
		"44b462": 833,  // Flextronics
		"44b6be": 2,    // Cisco
		"44bb3b": 3521, // Google
		"44bdde": 7212, // BHTC
		"44c306": 7213, // SIFROM
		"44c346": 3334, // Huawei
		"44c4a9": 3266, // Opticom
		"44c65d": 545,  // Apple
		"44c7fc": 3334, // Huawei
		"44c9a2": 7214, // Greenwald
		"44cb8b": 869,  // LG
		"44cd0e": 833,  // Flextronics
		"44ce7d": 3187, // SFR
		"44d244": 46,   // Epson
		"44d3ca": 2,    // Cisco
		"44d453": 2048, // Sagemcom
		"44d454": 2048, // Sagemcom
		"44d4e0": 101,  // Sony
		"44d5a5": 7215, // AddOn
		"44d5cc": 5438, // Amazon
		"44d63d": 7216, // Talari
		"44d6e1": 7217, // Snuza
		"44d791": 3334, // Huawei
		"44d832": 3004, // Azurewave
		"44d884": 545,  // Apple
		"44d9e7": 2978, // Ubiquiti
		"44da30": 545,  // Apple
		"44dc4e": 6281, // Itel
		"44dc91": 4478, // Planex
		"44dccb": 3666, // Semindia
		"44e137": 125,  // ARRIS
		"44e49a": 7218, // Omnitronics
		"44e4d9": 2,    // Cisco
		"44e4ee": 1592, // Wistron
		"44e517": 421,  // Intel
		"44e66e": 545,  // Apple
		"44e968": 3334, // Huawei
		"44e9dd": 2048, // Sagemcom
		"44ea30": 152,  // Samsung
		"44ea4b": 7219, // Actlas
		"44eb2e": 430,  // Alpsalpine
		"44ecce": 826,  // Juniper
		"44ed57": 7220, // Longicorn
		"44ee02": 2772, // MTI
		"44f034": 1269, // Kaonmedia
		"44f09e": 545,  // Apple
		"44f21b": 545,  // Apple
		"44f436": 3031, // ZTE
		"44f459": 152,  // Samsung
		"44f477": 826,  // Juniper
		"44f4e7": 7221, // Cohesity
		"44fb42": 545,  // Apple
		"44fb5a": 3031, // ZTE
		"44fda3": 7222, // Everysight
		"44fe3b": 2639, // Arcadyan
		"44ffba": 3031, // ZTE
		"480031": 3334, // Huawei
		"480033": 6392, // Technicolor
		"4801c5": 6923, // OnePlus
		"48022a": 7223, // B-Link
		"480362": 7224, // Desay
		"48049f": 5693, // Elecom
		"4805e2": 3334, // Huawei
		"48062b": 67,   // Private
		"48066a": 7225, // Tempered
		"480c49": 6082, // NAKAYO
		"480eec": 1595, // TP-Link
		"480fcf": 302,  // HP
		"481249": 5523, // Luxcom
		"481258": 3334, // Huawei
		"48128f": 3334, // Huawei
		"48137e": 152,  // Samsung
		"481693": 3535, // Lear
		"48174c": 7226, // MicroPower
		"4818fa": 7227, // Nocsys
		"48210b": 5439, // Pegatron
		"482567": 7228, // Poly
		"48262c": 545,  // Apple
		"4826e8": 7229, // Tek-Air
		"482759": 7230, // Levven
		"4827c5": 3334, // Huawei
		"4827ea": 152,  // Samsung
		"48282f": 3031, // ZTE
		"482952": 2048, // Sagemcom
		"482ae3": 1592, // Wistron
		"482ca0": 5697, // Xiaomi
		"482cd0": 3334, // Huawei
		"482e72": 2,    // Cisco
		"482f6b": 1685, // Aruba
		"482fd7": 3334, // Huawei
		"48343d": 7231, // IEP
		"48352b": 545,  // Apple
		"48365f": 3037, // Wintecronics
		"483871": 3334, // Huawei
		"483974": 3182, // Proware
		"483b38": 545,  // Apple
		"483c0c": 3334, // Huawei
		"483e5e": 6255, // Sernet
		"483fda": 6376, // Espressif
		"483fe9": 3334, // Huawei
		"48435a": 3334, // Huawei
		"48437c": 545,  // Apple
		"4843dd": 5438, // Amazon
		"4844f7": 152,  // Samsung
		"484520": 421,  // Intel
		"4846c1": 6440, // FN-Link
		"4846f1": 7232, // Uros
		"4846fb": 3334, // Huawei
		"48474b": 3334, // Huawei
		"4849c7": 152,  // Samsung
		"484ae9": 302,  // HP
		"484baa": 545,  // Apple
		"484bd4": 6392, // Technicolor
		"484c29": 3334, // Huawei
		"484c86": 3334, // Huawei
		"484d7e": 102,  // Dell
		"484efc": 125,  // ARRIS
		"485073": 612,  // Microsoft
		"485169": 152,  // Samsung
		"4851b7": 421,  // Intel
		"4851c5": 421,  // Intel
		"4851cf": 3541, // Intelbras
		"485261": 7233, // Soreel
		"485519": 6376, // Espressif
		"48555f": 4921, // Fiberhome
		"485702": 3334, // Huawei
		"4857dd": 7234, // Facebook
		"485929": 869,  // LG
		"4859a4": 3031, // ZTE
		"485a3f": 7235, // Wisol
		"485aea": 4921, // Fiberhome
		"485b39": 1806, // ASUS
		"485d36": 6621, // Verizon
		"485d60": 3004, // Azurewave
		"48605f": 869,  // LG
		"4860bc": 545,  // Apple
		"4861ee": 152,  // Samsung
		"486276": 3334, // Huawei
		"48684a": 421,  // Intel
		"486dbb": 1449, // Vestel
		"486e73": 7236, // Pica8
		"486fd2": 7237, // StorSimple
		"4873cb": 6539, // Tiinlab
		"487412": 6923, // OnePlus
		"48746e": 545,  // Apple
		"487583": 7238, // Intellion
		"487604": 67,   // Private
		"487746": 374,  // Calix
		"48785e": 5438, // Amazon
		"48794d": 152,  // Samsung
		"487a55": 5526, // ALE
		"487aff": 6830, // Essys
		"487b6b": 3334, // Huawei
		"487d2e": 1595, // TP-Link
		"487e48": 6498, // Earda
		"4883c7": 2048, // Sagemcom
		"4886e8": 612,  // Microsoft
		"488759": 5697, // Xiaomi
		"488764": 6369, // Vivo
		"488803": 7239, // ManTechnology
		"48881e": 7240, // EthoSwitch
		"4888ca": 687,  // Motorola
		"4889e7": 421,  // Intel
		"488ad2": 4253, // Mercury
		"488b0a": 2,    // Cisco
		"488c63": 3334, // Huawei
		"488d36": 2639, // Arcadyan
		"488e42": 7241, // DIGALOG
		"488eef": 3334, // Huawei
		"488f5a": 1784, // Routerboard.com
		"48902f": 869,  // LG
		"489a42": 7242, // Technomate
		"489bd5": 185,  // Extreme
		"489d18": 7243, // Flashbay
		"489d24": 2220, // BlackBerry
		"489dd1": 152,  // Samsung
		"489ebd": 302,  // HP
		"48a0f8": 4921, // Fiberhome
		"48a195": 545,  // Apple
		"48a2e6": 5986, // Resideo
		"48a472": 421,  // Intel
		"48a516": 3334, // Huawei
		"48a5e7": 1427, // Nintendo
		"48a6b8": 2047, // Sonos
		"48a74e": 3031, // ZTE
		"48a91c": 545,  // Apple
		"48a9d2": 1592, // Wistron
		"48aa5d": 7244, // Store
		"48ad08": 3334, // Huawei
		"48b02d": 7245, // NVIDIA
		"48b253": 7246, // Marketaxess
		"48b25d": 3334, // Huawei
		"48b423": 5438, // Amazon
		"48b620": 7247, // ROLI
		"48b8a3": 545,  // Apple
		"48b8de": 7248, // Homewins
		"48b977": 7249, // PulseOn
		"48b9c2": 7250, // Teletics
		"48ba4e": 302,  // HP
		"48bd4a": 3334, // Huawei
		"48be2d": 7251, // Symanitron
		"48bf6b": 545,  // Apple
		"48bf74": 7252, // Baicells
		"48c093": 2214, // Xirrus
		"48c1ac": 542,  // Plantronics
		"48c3b0": 7253, // Pharos
		"48c58d": 3535, // Lear
		"48c796": 152,  // Samsung
		"48c8b6": 7254, // SysTec
		"48cac6": 3505, // Unionman
		"48cb6e": 7255, // Cello
		"48d0cf": 3857, // Universal
		"48d18e": 6988, // Metis
		"48d224": 2873, // Liteon
		"48d24f": 2048, // Sagemcom
		"48d343": 125,  // ARRIS
		"48d35d": 67,   // Private
		"48d539": 3334, // Huawei
		"48d54c": 7256, // Jeda
		"48d6d5": 3521, // Google
		"48d705": 545,  // Apple
		"48d855": 7257, // Telvent
		"48d890": 6440, // FN-Link
		"48d8fe": 7258, // ClarIDy
		"48db50": 3334, // Huawei
		"48dc2d": 3334, // Huawei
		"48dcfb": 457,  // Nokia
		"48dd0c": 5820, // eero
		"48dd9d": 6281, // Itel
		"48df37": 302,  // HP
		"48e1af": 7259, // Vity
		"48e695": 7260, // Insigma
		"48e7da": 3004, // Azurewave
		"48e9f1": 545,  // Apple
		"48eb30": 7261, // Eterna
		"48eb62": 2056, // Murata
		"48ee0c": 803,  // D-Link
		"48ee86": 1135, // UTStarcom
		"48ef61": 3334, // Huawei
		"48f07b": 430,  // Alpsalpine
		"48f17f": 421,  // Intel
		"48f230": 7262, // Ubizcore
		"48f317": 67,   // Private
		"48f7c0": 6392, // Technicolor
		"48f7f1": 457,  // Nokia
		"48f8b3": 1783, // Linksys
		"48f8db": 3334, // Huawei
		"48f8e1": 457,  // Nokia
		"48f925": 7263, // Maestronic
		"48f97c": 4921, // Fiberhome
		"48fcb6": 647,  // Lava
		"48fcb8": 7264, // Woodstream
		"48fd8e": 3334, // Huawei
		"48fda3": 5697, // Xiaomi
		"4c0082": 2,    // Cisco
		"4c0143": 5820, // eero
		"4c0220": 5697, // Xiaomi
		"4c034f": 421,  // Intel
		"4c09b4": 3031, // ZTE
		"4c09d4": 2639, // Arcadyan
		"4c0a3d": 7265, // Adnacom
		"4c0b3a": 6433, // TCT
		"4c0bbe": 612,  // Microsoft
		"4c0fc7": 6498, // Earda
		"4c11ae": 6376, // Espressif
		"4c1265": 125,  // ARRIS
		"4c1365": 7266, // Emplus
		"4c1480": 7267, // Noregon
		"4c16f1": 3031, // ZTE
		"4c16fc": 826,  // Juniper
		"4c1744": 5438, // Amazon
		"4c17eb": 2048, // Sagemcom
		"4c195d": 2048, // Sagemcom
		"4c1a95": 7268, // Novakon
		"4c1b86": 2639, // Arcadyan
		"4c1d96": 421,  // Intel
		"4c1fcc": 3334, // Huawei
		"4c20b8": 545,  // Apple
		"4c218c": 2153, // Panasonic
		"4c21d0": 101,  // Sony
		"4c2258": 7269, // Cozybit
		"4c2578": 457,  // Nokia
		"4c26e7": 7270, // Welgate
		"4c2fd7": 3334, // Huawei
		"4c322d": 7271, // Teledata
		"4c3275": 545,  // Apple
		"4c3329": 7272, // Sweroam
		"4c334e": 7273, // Hightech
		"4c3488": 421,  // Intel
		"4c364e": 2153, // Panasonic
		"4c38d8": 125,  // ARRIS
		"4c3910": 7274, // Newtek
		"4c3b6c": 7275, // Garo
		"4c3b74": 7276, // VOGTEC
		"4c3bdf": 612,  // Microsoft
		"4c3c16": 152,  // Samsung
		"4c4088": 7277, // Sanshin
		"4c445b": 421,  // Intel
		"4c494f": 3031, // ZTE
		"4c49e3": 5697, // Xiaomi
		"4c4e03": 6433, // TCT
		"4c4e35": 2,    // Cisco
		"4c4fee": 6923, // OnePlus
		"4c5077": 3334, // Huawei
		"4c5262": 4,    // Fujitsu
		"4c52ec": 7278, // SOLARWATT
		"4c53fd": 5438, // Amazon
		"4c5499": 3334, // Huawei
		"4c5585": 7279, // Hamilton
		"4c55cc": 7280, // Zentri
		"4c569d": 545,  // Apple
		"4c57ca": 545,  // Apple
		"4c5d3c": 2,    // Cisco
		"4c5e0c": 1784, // Routerboard.com
		"4c60de": 1368, // Netgear
		"4c617e": 3334, // Huawei
		"4c6371": 5697, // Xiaomi
		"4c6641": 152,  // Samsung
		"4c6be8": 545,  // Apple
		"4c6d58": 826,  // Juniper
		"4c6e6e": 7281, // Comnect
		"4c710c": 2,    // Cisco
		"4c710d": 2,    // Cisco
		"4c72b9": 5439, // Pegatron
		"4c73a5": 7282, // Kove
		"4c7403": 7283, // BQ
		"4c74bf": 545,  // Apple
		"4c7525": 6376, // Espressif
		"4c7625": 102,  // Dell
		"4c7713": 5695, // Renesas
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
		"4c8b30": 2246, // Actiontec
		"4c8bef": 3334, // Huawei
		"4c8d53": 3334, // Huawei
		"4c8d79": 545,  // Apple
		"4c8ecc": 7284, // Silkan
		"4c8fa5": 4068, // Jastec
		"4c9614": 826,  // Juniper
		"4c962d": 7285, // Fresh
		"4c98ef": 7286, // Zeo
		"4c9eff": 2692, // ZyXEL
		"4ca003": 7287, // Vitec
		"4ca515": 7288, // Baikal
		"4ca56d": 152,  // Samsung
		"4ca64d": 2,    // Cisco
		"4ca928": 7289, // Insensi
		"4caa16": 3004, // Azurewave
		"4cab33": 7290, // KST
		"4cab4f": 545,  // Apple
		"4cabf8": 2544, // Askey
		"4cabfc": 3031, // ZTE
		"4cac0a": 3031, // ZTE
		"4cada8": 7291, // Panoptics
		"4cae13": 3334, // Huawei
		"4cae31": 7292, // ShengHai
		"4caea3": 302,  // HP
		"4cb16c": 3334, // Huawei
		"4cb199": 545,  // Apple
		"4cb1cd": 2737, // Ruckus
		"4cb21c": 7293, // Maxphotonics
		"4cb44a": 7294, // NANOWAVE
		"4cb4ea": 7295, // HRD
		"4cb81c": 7296, // SAM
		"4cb910": 545,  // Apple
		"4cb911": 2050, // Raisecom
		"4cb9c8": 7297, // Conet
		"4cb9ea": 7298, // iRobot
		"4cbaa3": 7299, // Bison
		"4cbad7": 869,  // LG
		"4cbb58": 7300, // Chicony
		"4cbc48": 2,    // Cisco
		"4cbca5": 152,  // Samsung
		"4cbce9": 869,  // LG
		"4cc00a": 6369, // Vivo
		"4cc206": 7301, // Somfy
		"4cc449": 2156, // Icotera
		"4cc53e": 2692, // ZyXEL
		"4cc602": 7302, // Radios
		"4cc7d6": 833,  // Flextronics
		"4cc94f": 457,  // Nokia
		"4cc95e": 152,  // Samsung
		"4cca53": 7303, // Skyera
		"4ccbf5": 3031, // ZTE
		"4ccc34": 687,  // Motorola
		"4ccc6a": 1812, // Micro-Star
		"4cce2d": 6395, // Danlaw
		"4cd08a": 531,  // HUMAX
		"4cd0cb": 3334, // Huawei
		"4cd1a1": 3334, // Huawei
		"4cd629": 3334, // Huawei
		"4cd637": 7304, // Qsono
		"4cd98f": 102,  // Dell
		"4cdd31": 152,  // Samsung
		"4cdf3d": 7305, // Team
		"4ce0db": 5697, // Xiaomi
		"4ce175": 2,    // Cisco
		"4ce176": 2,    // Cisco
		"4ce19e": 6265, // Tecno
		"4ce676": 1077, // Buffalo
		"4ce6c0": 545,  // Apple
		"4ce933": 7306, // RailComm
		"4ceb42": 421,  // Intel
		"4cebd6": 6376, // Espressif
		"4cecef": 7307, // Soraa
		"4cedde": 2544, // Askey
		"4cedfb": 1806, // ASUS
		"4cefc0": 5438, // Amazon
		"4cf202": 5697, // Xiaomi
		"4cf2bf": 1640, // Cambridge
		"4cf55b": 3334, // Huawei
		"4cf737": 7308, // SamJi
		"4cf95d": 3334, // Huawei
		"4cfaca": 1640, // Cambridge
		"4cfb45": 3334, // Huawei
		"4cfcaa": 7309, // Tesla
		"4cff12": 7310, // Fuze
		"500084": 300,  // Siemens
		"50016b": 3334, // Huawei
		"5001bb": 152,  // Samsung
		"5001d9": 3334, // Huawei
		"500291": 6376, // Espressif
		"5004b8": 3334, // Huawei
		"500604": 2,    // Cisco
		"5006ab": 2,    // Cisco
		"500959": 6392, // Technicolor
		"5009e5": 7311, // Drimsys
		"500a52": 7312, // Huiwan
		"500b32": 7313, // Foxda
		"500e6d": 7314, // TrafficCast
		"500f59": 6527, // STMicrolectronics
		"500f80": 2,    // Cisco
		"500ff5": 6266, // Tenda
		"5011eb": 7315, // SilverNet
		"501408": 7316, // AiNET
		"501479": 7298, // iRobot
		"5017ff": 2,    // Cisco
		"50184c": 7317, // Platina
		"501ac5": 612,  // Microsoft
		"501cb0": 2,    // Cisco
		"501cbf": 2,    // Cisco
		"501d93": 3334, // Huawei
		"501e2d": 7318, // StreamUnlimited
		"501fc6": 545,  // Apple
		"5021ec": 3334, // Huawei
		"502267": 7319, // PixeLINK
		"5023a2": 545,  // Apple
		"502690": 4,    // Fujitsu
		"5027c7": 7320, // TECHNART
		"502873": 3334, // Huawei
		"502a7e": 7321, // Smart
		"502b73": 6266, // Tenda
		"502b98": 7322, // Es-tech
		"502d1d": 457,  // Nokia
		"502da2": 421,  // Intel
		"502dfb": 7323, // IGShare
		"502e5c": 1341, // HTC
		"502ece": 7324, // Asahi
		"502f9b": 421,  // Intel
		"502fa8": 2,    // Cisco
		"503237": 545,  // Apple
		"50325f": 1655, // Silicon
		"503275": 152,  // Samsung
		"5033f0": 7325, // Yichen
		"503cc4": 2664, // Lenovo
		"503da1": 152,  // Samsung
		"503dc6": 5697, // Xiaomi
		"503de5": 2,    // Cisco
		"503eaa": 1595, // TP-Link
		"503f56": 7326, // Syncmold
		"503f98": 7327, // Cmitech
		"504061": 457,  // Nokia
		"50411c": 4498, // AMPAK
		"504348": 7328, // ThingsMatrix
		"50464a": 3334, // Huawei
		"50465d": 1806, // ASUS
		"5046ae": 4253, // Mercury
		"5049b0": 152,  // Samsung
		"504a5e": 6362, // Masimo
		"504a6e": 1368, // Netgear
		"504b5b": 7329, // CoNTROLtronic
		"504b9e": 3334, // Huawei
		"504edc": 4357, // Ping
		"504f94": 7330, // Loxone
		"50502a": 7331, // Egardia
		"505065": 7332, // TAKT
		"5050a4": 152,  // Samsung
		"50523b": 457,  // Nokia
		"505527": 869,  // LG
		"5056a8": 7333, // Jolla
		"5056bf": 152,  // Samsung
		"50578a": 545,  // Apple
		"50579c": 46,   // Epson
		"5057a8": 2,    // Cisco
		"505800": 7334, // WyTec
		"50584f": 7335, // waytotec
		"50586f": 3334, // Huawei
		"505967": 7336, // Intent
		"505bc2": 2873, // Liteon
		"505d7a": 3031, // ZTE
		"505dac": 3334, // Huawei
		"505fb5": 2544, // Askey
		"506028": 2214, // Xirrus
		"506184": 620,  // Avaya
		"5061bf": 2,    // Cisco
		"5061d6": 7337, // Indu-Sol
		"5061f6": 3857, // Universal
		"50642b": 6280, // XIAOMI
		"506441": 7338, // Greenlee
		"5065f3": 302,  // HP
		"506787": 7339, // Planet
		"5067ae": 2,    // Cisco
		"5067f0": 2692, // ZyXEL
		"50680a": 3334, // Huawei
		"506a03": 1368, // Netgear
		"506b4b": 432,  // Mellanox
		"506b8d": 7340, // Nutanix
		"506cbe": 7341, // InnosiliconTechnology
		"506e92": 7342, // Innocent
		"506f0c": 2048, // Sagemcom
		"506f77": 3334, // Huawei
		"507043": 3512, // BSkyB
		"5075f1": 125,  // ARRIS
		"507691": 7343, // Tekpea
		"5076af": 421,  // Intel
		"507705": 152,  // Samsung
		"5078b0": 3334, // Huawei
		"5078b3": 3031, // ZTE
		"507a55": 545,  // Apple
		"507ac5": 545,  // Apple
		"507b9d": 4919, // LCFC
		"507c6f": 421,  // Intel
		"507d02": 7344, // Biodit
		"507e5d": 2639, // Arcadyan
		"508140": 302,  // HP
		"5082d5": 545,  // Apple
		"508492": 421,  // Intel
		"508569": 152,  // Samsung
		"508789": 2,    // Cisco
		"5087b8": 7345, // Nuvyyo
		"5089d1": 3334, // Huawei
		"508a42": 7346, // Uptmate
		"508d6f": 7347, // CHAHOO
		"508e49": 5697, // Xiaomi
		"508f4c": 5697, // Xiaomi
		"5092b9": 152,  // Samsung
		"509551": 125,  // ARRIS
		"509839": 5697, // Xiaomi
		"509871": 7348, // Inventum
		"509a46": 7349, // Safetrust
		"509a4c": 102,  // Dell
		"509a88": 3334, // Huawei
		"509ea7": 152,  // Samsung
		"509f27": 3334, // Huawei
		"50a009": 5697, // Xiaomi
		"50a054": 7350, // Actineon
		"50a0a4": 457,  // Nokia
		"50a4c8": 152,  // Samsung
		"50a5dc": 125,  // ARRIS
		"50a67f": 545,  // Apple
		"50a715": 7351, // Aboundi
		"50a72b": 3334, // Huawei
		"50a733": 2737, // Ruckus
		"50ab3e": 7352, // Qibixx
		"50ad92": 7353, // NX
		"50add5": 7354, // Dynalec
		"50ae86": 7355, // Linkintec
		"50af4d": 3031, // ZTE
		"50b7c3": 152,  // Samsung
		"50b8a2": 7356, // ImTech
		"50bc96": 545,  // Apple
		"50bd5f": 1595, // TP-Link
		"50c271": 7357, // Securetech
		"50c3a2": 3202, // nFore
		"50c4dd": 1077, // Buffalo
		"50c58d": 826,  // Juniper
		"50c6ad": 4921, // Fiberhome
		"50c709": 826,  // Juniper
		"50c7bf": 1595, // TP-Link
		"50c8e5": 152,  // Samsung
		"50c9a0": 7358, // Skipper
		"50ccf8": 152,  // Samsung
		"50cd22": 620,  // Avaya
		"50ce75": 7359, // Measy
		"50cee3": 7360, // Gigafirm
		"50d065": 7361, // ESYLUX
		"50d213": 7362, // CviLux
		"50d274": 7363, // Steffes
		"50d4f7": 1595, // TP-Link
		"50d753": 7364, // CoNELCOM
		"50dad6": 5697, // Xiaomi
		"50dcd0": 7365, // Observint
		"50dce7": 5438, // Amazon
		"50dcfc": 7366, // Ecocom
		"50de06": 545,  // Apple
		"50df95": 6858, // Lytx
		"50e039": 2692, // ZyXEL
		"50e085": 421,  // Intel
		"50e0c7": 7367, // TurControlSystme
		"50e0ef": 457,  // Nokia
		"50e14a": 67,   // Private
		"50e24e": 3031, // ZTE
		"50e549": 1929, // Giga-Byte
		"50e7a0": 5695, // Renesas
		"50e7b7": 6369, // Vivo
		"50e971": 7368, // Jibo
		"50ead6": 545,  // Apple
		"50eb1a": 90,   // Brocade
		"50eb71": 421,  // Intel
		"50ebf6": 1806, // ASUS
		"50ed3c": 545,  // Apple
		"50f0d3": 152,  // Samsung
		"50f43c": 7369, // Leeo
		"50f4eb": 545,  // Apple
		"50f520": 152,  // Samsung
		"50f5da": 5438, // Amazon
		"50f722": 2,    // Cisco
		"50f7ed": 3334, // Huawei
		"50f8a5": 7370, // eWBM
		"50f908": 7371, // Wizardlab
		"50f958": 3334, // Huawei
		"50fa84": 1595, // TP-Link
		"50fb19": 6439, // Chipsea
		"50fc30": 7372, // Treehouse
		"50fc9f": 152,  // Samsung
		"50fef2": 7373, // Sify
		"50ff20": 7374, // Keenetic
		"540237": 7375, // Teltronic
		"5403f5": 7376, // EBN
		"540496": 7377, // Gigawave
		"5404a6": 1806, // ASUS
		"540536": 7378, // Vivago
		"5405db": 4919, // LCFC
		"540764": 3334, // Huawei
		"540910": 545,  // Apple
		"540955": 3031, // ZTE
		"54098d": 7379, // deister
		"540df9": 3334, // Huawei
		"540e2d": 6369, // Vivo
		"540f57": 1655, // Silicon
		"541031": 7380, // Smarto
		"5410ec": 706,  // Microchip
		"54115f": 7381, // Atamo
		"541310": 3334, // Huawei
		"541473": 4030, // Wingtech
		"5414f3": 421,  // Intel
		"541651": 5440, // Ruijie
		"5419c8": 6369, // Vivo
		"541b5d": 7382, // Techno-Innov
		"541e56": 826,  // Juniper
		"541f8d": 3031, // ZTE
		"541fd5": 7383, // Advantage
		"542018": 7384, // Tely
		"54211d": 3334, // Huawei
		"542160": 7385, // Alula
		"54219d": 152,  // Samsung
		"5422f8": 3031, // ZTE
		"5425ea": 3334, // Huawei
		"542696": 545,  // Apple
		"54271e": 3004, // Azurewave
		"542758": 687,  // Motorola
		"542a1b": 2047, // Sonos
		"542aa2": 2236, // Alpha
		"542b8d": 545,  // Apple
		"542cea": 7386, // Protectron
		"542f89": 7387, // Euclid
		"5433cb": 545,  // Apple
		"5434ef": 3334, // Huawei
		"5435df": 7388, // Symeo
		"543968": 7389, // Edgewater
		"5439df": 3334, // Huawei
		"543ad6": 152,  // Samsung
		"543b30": 7390, // duagon
		"543d37": 2737, // Ruckus
		"543e64": 4921, // Fiberhome
		"5440ad": 152,  // Samsung
		"544249": 101,  // Sony
		"544408": 457,  // Nokia
		"5444a3": 152,  // Samsung
		"544617": 3031, // ZTE
		"544741": 7391, // Xcheng
		"5447d3": 7392, // Tsat
		"5447e8": 7022, // Syrotech
		"544810": 102,  // Dell
		"54489c": 7393, // Cdoubles
		"544a00": 2,    // Cisco
		"544b8c": 826,  // Juniper
		"544e45": 67,   // Private
		"544e90": 545,  // Apple
		"54511b": 3334, // Huawei
		"545146": 7394, // AMG
		"545284": 3334, // Huawei
		"5453ed": 101,  // Sony
		"5454cf": 7395, // Probedigital
		"5455d5": 3334, // Huawei
		"545aa6": 6376, // Espressif
		"545ebd": 2813, // NL
		"545fa9": 4820, // Teracom
		"546009": 3521, // Google
		"5461ea": 7396, // Zaplox
		"5462e2": 545,  // Apple
		"5464d9": 2048, // Sagemcom
		"5465de": 125,  // ARRIS
		"5466f9": 7397, // ConMet
		"546751": 358,  // Compal
		"546990": 3334, // Huawei
		"546ceb": 421,  // Intel
		"546f71": 7398, // uAvionix
		"5471dd": 3334, // Huawei
		"54724f": 545,  // Apple
		"54725e": 3505, // Unionman
		"547398": 7399, // Toyo
		"547595": 1595, // TP-Link
		"5475d0": 2,    // Cisco
		"547787": 6498, // Earda
		"54778a": 302,  // HP
		"54781a": 2,    // Cisco
		"5478c9": 4498, // AMPAK
		"547975": 457,  // Nokia
		"547c69": 2,    // Cisco
		"547d40": 7400, // Powervision
		"547f54": 536,  // Ingenico
		"547fa8": 7401, // TELCO
		"547fbc": 7402, // iodyne
		"547fee": 2,    // Cisco
		"548028": 302,  // HP
		"54812d": 3217, // PAX
		"5481ad": 4267, // Eagle
		"54833a": 2692, // ZyXEL
		"5484dc": 3031, // ZTE
		"5486bc": 2,    // Cisco
		"54880e": 152,  // Samsung
		"5488de": 2,    // Cisco
		"548922": 7403, // Zelfy
		"548998": 3334, // Huawei
		"548aba": 2,    // Cisco
		"548ca0": 2873, // Liteon
		"548d5a": 421,  // Intel
		"549209": 3334, // Huawei
		"5492be": 152,  // Samsung
		"549963": 545,  // Apple
		"549b12": 152,  // Samsung
		"549b72": 306,  // Ericsson
		"549d85": 7404, // EnerAccess
		"549f13": 545,  // Apple
		"549f35": 102,  // Dell
		"549fc6": 2,    // Cisco
		"54a04f": 7405, // t-mac
		"54a050": 1806, // ASUS
		"54a274": 2,    // Cisco
		"54a3fa": 7406, // BQT
		"54a51b": 3334, // Huawei
		"54a65c": 6392, // Technicolor
		"54a6db": 3334, // Huawei
		"54a703": 1595, // TP-Link
		"54a9d4": 7407, // Minibar
		"54ab3a": 3071, // Quanta
		"54ae27": 545,  // Apple
		"54aed0": 7408, // DASAN
		"54af97": 1595, // TP-Link
		"54b121": 3334, // Huawei
		"54b203": 5439, // Pegatron
		"54b7e5": 2601, // Rayson
		"54b802": 152,  // Samsung
		"54b80a": 803,  // D-Link
		"54bad6": 3334, // Huawei
		"54bd79": 152,  // Samsung
		"54be53": 3031, // ZTE
		"54bef7": 5439, // Pegatron
		"54bf64": 102,  // Dell
		"54c33e": 4564, // Ciena
		"54c480": 3334, // Huawei
		"54c57a": 5435, // Sunnovo
		"54c80f": 1595, // TP-Link
		"54c9df": 6440, // FN-Link
		"54cd10": 2153, // Panasonic
		"54ce82": 3031, // ZTE
		"54cf8d": 3334, // Huawei
		"54d0ed": 7409, // AXIM
		"54d163": 7410, // MAX-Tech
		"54d17d": 152,  // Samsung
		"54d751": 7411, // Proximus
		"54d9c6": 3334, // Huawei
		"54d9e4": 7412, // Brilliantts
		"54dba2": 7413, // Fibrain
		"54dc1d": 3096, // Yulong
		"54df00": 7414, // Ulterius
		"54df24": 4921, // Fiberhome
		"54df63": 7415, // Intrakey
		"54e005": 4921, // Fiberhome
		"54e019": 6946, // Ring
		"54e032": 826,  // Juniper
		"54e140": 536,  // Ingenico
		"54e1ad": 4919, // LCFC
		"54e2e0": 125,  // ARRIS
		"54e43a": 545,  // Apple
		"54e4a9": 7416, // BHR
		"54e4bd": 6440, // FN-Link
		"54e61b": 545,  // Apple
		"54e6fc": 1595, // TP-Link
		"54eaa8": 545,  // Apple
		"54ec2f": 2737, // Ruckus
		"54eda3": 7417, // Navdy
		"54ee75": 1592, // Wistron
		"54effe": 7418, // Fullpower
		"54f201": 152,  // Samsung
		"54f294": 3334, // Huawei
		"54f607": 3334, // Huawei
		"54f6e2": 3334, // Huawei
		"54f82a": 7419, // u-blox
		"54f876": 21,   // ABB
		"54fa3e": 152,  // Samsung
		"54fb58": 7420, // WISEWARE
		"54fcf0": 152,  // Samsung
		"5800bb": 826,  // Juniper
		"5800e3": 2873, // Liteon
		"580528": 7421, // Labris
		"580943": 67,   // Private
		"5809e5": 7422, // Kivic
		"580a20": 2,    // Cisco
		"580ad4": 545,  // Apple
		"58108c": 3541, // Intelbras
		"5810b7": 6295, // Infinix
		"581122": 1806, // ASUS
		"581243": 4903, // AcSiP
		"5813d3": 1450, // Gemtek
		"581626": 620,  // Avaya
		"5816d7": 430,  // Alpsalpine
		"58170c": 101,  // Sony
		"5819f8": 125,  // ARRIS
		"581cbd": 7423, // Affinegy
		"581f28": 3334, // Huawei
		"581f67": 7424, // Open-m
		"581faa": 545,  // Apple
		"581fef": 7425, // Tuttnaer
		"582059": 5697, // Xiaomi
		"5820b1": 302,  // HP
		"582136": 7426, // KMB
		"5821e9": 7427, // Twpi
		"58238c": 6392, // Technicolor
		"582429": 3521, // Google
		"582575": 3334, // Huawei
		"58278c": 1077, // Buffalo
		"582af7": 3334, // Huawei
		"582bdb": 7428, // Pax
		"582d34": 7429, // Qingping
		"582f40": 1427, // Nintendo
		"582ff7": 2048, // Sagemcom
		"583112": 7430, // Drust
		"583277": 2375, // Reliance
		"58343b": 7431, // Glovast
		"583526": 7432, // Deeplet
		"58355d": 3334, // Huawei
		"58356b": 6265, // Tecno
		"5835d9": 2,    // Cisco
		"583879": 75,   // Ricoh
		"583bd9": 4921, // Fiberhome
		"583cc6": 7433, // Omneality
		"583f54": 869,  // LG
		"58404e": 545,  // Apple
		"584120": 1595, // TP-Link
		"5842e4": 7434, // Baxter
		"584498": 5697, // Xiaomi
		"58454c": 306,  // Ericsson
		"58468f": 7435, // Koncar
		"5846e1": 7434, // Baxter
		"584822": 101,  // Sony
		"5848c0": 7436, // Coflec
		"584925": 7168, // E3
		"5849ba": 7437, // Chitai
		"584d42": 7438, // Dragos
		"5850ab": 7439, // TLS
		"5855ca": 545,  // Apple
		"5856c2": 3334, // Huawei
		"5856e8": 125,  // ARRIS
		"5859c2": 185,  // Extreme
		"585ff6": 3031, // ZTE
		"58605f": 3334, // Huawei
		"5860d8": 125,  // ARRIS
		"586356": 6440, // FN-Link
		"5865e6": 3976, // Infomark
		"58696c": 5440, // Ruijie
		"586b14": 545,  // Apple
		"586c25": 421,  // Intel
		"586d8f": 1783, // Linksys
		"586ed6": 67,   // Private
		"587a4d": 7440, // Stonesoft
		"587f57": 545,  // Apple
		"587f66": 3334, // Huawei
		"587fb7": 7441, // Sonar
		"587fc8": 7442, // S2M
		"5882a8": 612,  // Microsoft
		"58856e": 5356, // QSC
		"588694": 1252, // EFM
		"588a5a": 102,  // Dell
		"588bf3": 2692, // ZyXEL
		"588d09": 2,    // Cisco
		"588e81": 1655, // Silicon
		"589043": 2048, // Sagemcom
		"5891cf": 421,  // Intel
		"589396": 2737, // Ruckus
		"58946b": 421,  // Intel
		"5894a2": 7443, // KETEK
		"5894ae": 3334, // Huawei
		"5894b2": 7444, // BrainCo
		"58961d": 421,  // Intel
		"589630": 6392, // Technicolor
		"58971e": 2,    // Cisco
		"5897bd": 2,    // Cisco
		"589835": 6392, // Technicolor
		"589b0b": 7445, // Shineway
		"589b4a": 7446, // DWnet
		"589ec6": 4298, // Gigaset
		"58a023": 421,  // Intel
		"58a0cb": 7447, // TrackNet
		"58a2b5": 869,  // LG
		"58a639": 152,  // Samsung
		"58a76f": 7448, // iD
		"58a839": 421,  // Intel
		"58a87b": 6587, // Fitbit
		"58ac78": 2,    // Cisco
		"58ae2b": 3334, // Huawei
		"58aea8": 3334, // Huawei
		"58aef1": 4921, // Fiberhome
		"58b035": 545,  // Apple
		"58b0d4": 7449, // ZuniData
		"58b10f": 152,  // Samsung
		"58b42d": 7450, // YSTen
		"58b633": 2737, // Ruckus
		"58b961": 7451, // SOLEM
		"58bad4": 3334, // Huawei
		"58bc27": 2,    // Cisco
		"58bc8f": 7452, // Cognitive
		"58bda3": 1427, // Nintendo
		"58bdf9": 7453, // Sigrand
		"58be72": 3334, // Huawei
		"58bf25": 6376, // Espressif
		"58bfea": 2,    // Cisco
		"58c17a": 665,  // Cambium
		"58c232": 48,   // NEC
		"58c38b": 152,  // Samsung
		"58c583": 6281, // Itel
		"58c5cb": 152,  // Samsung
		"58cb52": 3521, // Google
		"58ce2a": 421,  // Intel
		"58cf4b": 7454, // Lufkin
		"58cf79": 6376, // Espressif
		"58d061": 3334, // Huawei
		"58d349": 545,  // Apple
		"58d50a": 2056, // Murata
		"58d56e": 803,  // D-Link
		"58d67a": 7455, // TCPlink
		"58d759": 3334, // Huawei
		"58d9c3": 687,  // Motorola
		"58d9d5": 6266, // Tenda
		"58db15": 6265, // Tecno
		"58db8d": 298,  // Fast
		"58e28f": 545,  // Apple
		"58e326": 4101, // Compass
		"58e636": 7456, // EVRsafe
		"58e6ba": 545,  // Apple
		"58e747": 7457, // Deltanet
		"58e808": 7458, // Autonics
		"58eafc": 7459, // ELL-IoT
		"58ece1": 1388, // Newport
		"58ef68": 2468, // Belkin
		"58f102": 7460, // BLU
		"58f2fc": 3334, // Huawei
		"58f387": 7461, // Hccp
		"58f39c": 2,    // Cisco
		"58f987": 3334, // Huawei
		"58f98e": 7462, // SECUDOS
		"58fb84": 421,  // Intel
		"58fb96": 2737, // Ruckus
		"58fcc6": 7463, // Tozo
		"58fd20": 7464, // Systemhouse
		"58fdb1": 869,  // LG
		"5c0272": 1655, // Silicon
		"5c0339": 3334, // Huawei
		"5c0947": 545,  // Apple
		"5c0979": 3334, // Huawei
		"5c0a5b": 152,  // Samsung
		"5c0cbb": 7465, // CELIZION
		"5c0ce6": 1427, // Nintendo
		"5c0e8b": 185,  // Extreme
		"5c0ffb": 320,  // Amino
		"5c10c5": 152,  // Samsung
		"5c1515": 7466, // Advan
		"5c15e1": 7467, // Aidc
		"5c16c7": 3793, // Arista
		"5c1720": 3334, // Huawei
		"5c17cf": 6923, // OnePlus
		"5c17d3": 7468, // LGE
		"5c18b5": 7469, // Talon
		"5c1a6f": 1640, // Cambridge
		"5c1cb9": 6369, // Vivo
		"5c1dd9": 545,  // Apple
		"5c20d0": 7470, // Asoni
		"5c2316": 7471, // Squirrels
		"5c2479": 7472, // Baltech
		"5c260a": 102,  // Dell
		"5c2623": 7473, // WaveLynx
		"5c2e59": 152,  // Samsung
		"5c2ed2": 7474, // ABC
		"5c3192": 2,    // Cisco
		"5c32c5": 4820, // Teracom
		"5c338e": 2236, // Alpha
		"5c353b": 358,  // Compal
		"5c35da": 7475, // There
		"5c3a3d": 3031, // ZTE
		"5c3b35": 7476, // Gehirn
		"5c3c27": 152,  // Samsung
		"5c415a": 7477, // Amazon.com
		"5c41e7": 7478, // Wiatec
		"5c43d2": 7479, // Hazemeyer
		"5c443e": 7039, // Skullcandy
		"5c4527": 826,  // Juniper
		"5c497d": 152,  // Samsung
		"5c4a26": 7480, // Enguity
		"5c4ca9": 3334, // Huawei
		"5c5015": 2,    // Cisco
		"5c50d9": 545,  // Apple
		"5c514f": 421,  // Intel
		"5c5181": 152,  // Samsung
		"5c5188": 687,  // Motorola
		"5c521e": 1427, // Nintendo
		"5c5230": 545,  // Apple
		"5c546d": 3334, // Huawei
		"5c5578": 7481, // iryx
		"5c56ed": 7482, // 3pleplay
		"5c571a": 125,  // ARRIS
		"5c57c8": 457,  // Nokia
		"5c5819": 7483, // Jingsheng
		"5c5948": 545,  // Apple
		"5c5ac7": 2,    // Cisco
		"5c5aea": 4880, // Ford
		"5c5b35": 7484, // Mist
		"5c5bc2": 7485, // YIK
		"5c5eab": 826,  // Juniper
		"5c5f67": 421,  // Intel
		"5c625a": 87,   // Canon
		"5c63bf": 1595, // TP-Link
		"5c63c9": 7486, // Intellithings
		"5c647a": 3334, // Huawei
		"5c648e": 2692, // ZyXEL
		"5c6984": 7487, // Nuvico
		"5c6a80": 2692, // ZyXEL
		"5c6b4f": 7488, // Hello
		"5c6f69": 858,  // Broadcom
		"5c7017": 545,  // Apple
		"5c70a3": 869,  // LG
		"5c710d": 2,    // Cisco
		"5c75af": 6587, // Fitbit
		"5c7695": 6392, // Technicolor
		"5c7776": 6433, // TCT
		"5c78f8": 3334, // Huawei
		"5c7d5e": 3334, // Huawei
		"5c7d7d": 6392, // Technicolor
		"5c80b6": 421,  // Intel
		"5c8382": 457,  // Nokia
		"5c838f": 2,    // Cisco
		"5c843c": 101,  // Sony
		"5c8486": 7489, // Brightsource
		"5c864a": 7490, // Secret
		"5c865c": 152,  // Samsung
		"5c8730": 545,  // Apple
		"5c8778": 7491, // Cybertelbridge
		"5c879c": 421,  // Intel
		"5c899a": 1595, // TP-Link
		"5c8a38": 302,  // HP
		"5c8d4e": 545,  // Apple
		"5c8f40": 6265, // Tecno
		"5c8fe0": 125,  // ARRIS
		"5c9157": 3334, // Huawei
		"5c91fd": 7492, // Jaewoncnc
		"5c925e": 2127, // Zioncom
		"5c93a2": 2873, // Liteon
		"5c95ae": 545,  // Apple
		"5c9656": 3004, // Azurewave
		"5c9666": 101,  // Sony
		"5c966a": 7493, // Rtnet
		"5c969d": 545,  // Apple
		"5c97f3": 545,  // Apple
		"5c9960": 152,  // Samsung
		"5c9aa1": 3334, // Huawei
		"5c9ad8": 4,    // Fujitsu
		"5ca1e0": 7494, // EmbedWay
		"5ca39d": 152,  // Samsung
		"5ca48a": 2,    // Cisco
		"5ca4a4": 4921, // Fiberhome
		"5ca5bc": 5820, // eero
		"5ca62d": 2,    // Cisco
		"5ca6e6": 1595, // TP-Link
		"5ca86a": 3334, // Huawei
		"5caafd": 2047, // Sonos
		"5cadcf": 545,  // Apple
		"5caf06": 869,  // LG
		"5cb00a": 3334, // Huawei
		"5cb066": 125,  // ARRIS
		"5cb13e": 2048, // Sagemcom
		"5cb395": 3334, // Huawei
		"5cb3f6": 7495, // Humanorporated
		"5cb43e": 3334, // Huawei
		"5cb524": 101,  // Sony
		"5cb559": 7496, // CNEX
		"5cb6cc": 7497, // NovaComm
		"5cb8cb": 7498, // Allis
		"5cb901": 302,  // HP
		"5cba2c": 302,  // HP
		"5cba37": 612,  // Microsoft
		"5cbd9a": 3334, // Huawei
		"5cc0a0": 3334, // Huawei
		"5cc1d7": 152,  // Samsung
		"5cc307": 3334, // Huawei
		"5cc336": 7499, // ittim
		"5cc5d4": 421,  // Intel
		"5cc6e9": 6445, // Edifier
		"5cc7d7": 7500, // Azroad
		"5cc9c0": 5695, // Renesas
		"5cca1a": 612,  // Microsoft
		"5cca32": 7501, // Theben
		"5ccad3": 6439, // Chipsea
		"5ccb99": 152,  // Samsung
		"5ccca0": 7502, // Gridwiz
		"5ccd5b": 421,  // Intel
		"5ccd7c": 7030, // MEIZU
		"5ccead": 7503, // CDYNE
		"5ccf7f": 6376, // Espressif
		"5cd06e": 5697, // Xiaomi
		"5cd20b": 7504, // Yytek
		"5cd2e4": 421,  // Intel
		"5cd41b": 7505, // UCZOON
		"5cd4ab": 7506, // Zektor
		"5cd61f": 7507, // Qardio
		"5cd89e": 3334, // Huawei
		"5cd998": 803,  // D-Link
		"5cdad4": 2056, // Murata
		"5cdc96": 2639, // Arcadyan
		"5cdf89": 2737, // Ruckus
		"5ce0c5": 421,  // Intel
		"5ce176": 2,    // Cisco
		"5ce223": 7508, // Delphin
		"5ce286": 76,   // Nortel
		"5ce28c": 2692, // ZyXEL
		"5ce2f4": 4903, // AcSiP
		"5ce30e": 125,  // ARRIS
		"5ce3b6": 4921, // Fiberhome
		"5ce42a": 421,  // Intel
		"5ce747": 3334, // Huawei
		"5ce7a0": 457,  // Nokia
		"5ce883": 3334, // Huawei
		"5ce8b7": 6878, // Oraimo
		"5ce8eb": 152,  // Samsung
		"5ceb68": 7509, // Cheerstar
		"5ced8c": 302,  // HP
		"5cf207": 7510, // Speco
		"5cf370": 385,  // CC&C
		"5cf3fc": 372,  // IBM
		"5cf4ab": 2692, // ZyXEL
		"5cf5da": 545,  // Apple
		"5cf6dc": 152,  // Samsung
		"5cf7c3": 7511, // Syntech
		"5cf7e6": 545,  // Apple
		"5cf8a1": 2056, // Murata
		"5cf938": 545,  // Apple
		"5cf96a": 3334, // Huawei
		"5cf9dd": 102,  // Dell
		"5cf9f0": 7512, // Atomos
		"5cfafb": 7513, // Acubit
		"5cfc66": 2,    // Cisco
		"5cff35": 1592, // Wistron
		"600194": 6376, // Espressif
		"600292": 5439, // Pegatron
		"6002b4": 1592, // Wistron
		"600308": 545,  // Apple
		"600417": 7514, // Posbank
		"6006e3": 545,  // Apple
		"600810": 3334, // Huawei
		"6009c3": 7419, // u-blox
		"600f77": 7515, // SilverPlus
		"60109e": 3334, // Huawei
		"601199": 7516, // Siama
		"60123c": 3334, // Huawei
		"60128b": 87,   // Canon
		"601466": 3031, // ZTE
		"6014b3": 189,  // CyberTAN
		"6015c7": 7517, // IdaTech
		"601888": 3031, // ZTE
		"601895": 102,  // Dell
		"60190c": 7518, // Rramac
		"601971": 125,  // ARRIS
		"601d91": 687,  // Motorola
		"601e02": 7519, // EltexAlatau
		"602103": 7520, // I4VINE
		"6021c0": 2056, // Murata
		"602232": 2978, // Ubiquiti
		"6026aa": 2,    // Cisco
		"6026ef": 1685, // Aruba
		"6029d5": 7521, // DAVOLINK
		"602e20": 3334, // Huawei
		"6030d4": 545,  // Apple
		"60313b": 5435, // Sunnovo
		"603197": 2692, // ZyXEL
		"6032b1": 1595, // TP-Link
		"6032f0": 7522, // Mplus
		"60334b": 545,  // Apple
		"603553": 7523, // Buwon
		"603573": 6498, // Earda
		"6035c0": 3187, // SFR
		"603696": 7524, // Sapling
		"6036dd": 421,  // Intel
		"60380e": 430,  // Alpsalpine
		"6038e0": 2468, // Belkin
		"60391f": 21,   // ABB
		"603a7c": 1595, // TP-Link
		"603aaf": 152,  // Samsung
		"603cee": 869,  // LG
		"603d26": 6392, // Technicolor
		"603d29": 3334, // Huawei
		"603e7b": 7525, // Gafachi
		"60447a": 7526, // Water-i.d
		"6045bd": 612,  // Microsoft
		"6045cb": 1806, // ASUS
		"6047d4": 7527, // FORICS
		"6049c1": 620,  // Avaya
		"604a1c": 7528, // SUYIN
		"604de1": 3334, // Huawei
		"604f5b": 3334, // Huawei
		"60512c": 6433, // TCT
		"6052d0": 7529, // FACTS
		"605317": 7530, // Sandstone
		"605375": 3334, // Huawei
		"6055f9": 6376, // Espressif
		"605661": 7531, // IXECLOUD
		"605699": 67,   // Private
		"605718": 421,  // Intel
		"60577d": 5820, // eero
		"605bb4": 3004, // Azurewave
		"605e4f": 3334, // Huawei
		"605f8d": 5820, // eero
		"6061df": 7532, // Z-meta
		"60634c": 803,  // D-Link
		"6063f9": 7533, // Ciholas
		"606453": 7534, // AOD
		"6064a1": 7535, // RADiflow
		"606720": 421,  // Intel
		"60684e": 152,  // Samsung
		"606944": 545,  // Apple
		"60699b": 7536, // isepos
		"606bbd": 152,  // Samsung
		"606bff": 1427, // Nintendo
		"606c63": 870,  // Hitron
		"606c66": 421,  // Intel
		"606ed0": 7537, // Seal
		"606ee8": 5697, // Xiaomi
		"6070c0": 545,  // Apple
		"60720b": 7460, // BLU
		"60735c": 2,    // Cisco
		"6073bc": 3031, // ZTE
		"607688": 7538, // Velodyne
		"6077e2": 152,  // Samsung
		"607ec9": 545,  // Apple
		"607ecd": 3334, // Huawei
		"607edd": 612,  // Microsoft
		"6081f9": 7539, // Helium
		"608334": 3334, // Huawei
		"608373": 545,  // Apple
		"60843b": 7540, // Soladigm
		"6084bd": 1077, // Buffalo
		"608a10": 706,  // Microchip
		"608b0e": 545,  // Apple
		"608c2b": 3954, // Hanson
		"608c4a": 545,  // Apple
		"608ce6": 125,  // ARRIS
		"608d26": 2639, // Arcadyan
		"608e08": 152,  // Samsung
		"608f5c": 152,  // Samsung
		"609084": 7541, // DSSD
		"6091f3": 6369, // Vivo
		"609217": 545,  // Apple
		"6092f5": 125,  // ARRIS
		"609316": 545,  // Apple
		"609620": 67,   // Private
		"6097dd": 7542, // MicroSys
		"609ac1": 545,  // Apple
		"609bb4": 3334, // Huawei
		"609c9f": 90,   // Brocade
		"609e64": 7543, // Vivonic
		"609f9d": 7544, // CloudSwitch
		"60a10a": 152,  // Samsung
		"60a37d": 545,  // Apple
		"60a423": 1655, // Silicon
		"60a44c": 1806, // ASUS
		"60a4b7": 1595, // TP-Link
		"60a4d0": 152,  // Samsung
		"60a5e2": 421,  // Intel
		"60a6c5": 3334, // Huawei
		"60a751": 3334, // Huawei
		"60a9b0": 7545, // Merchandising
		"60aaef": 3334, // Huawei
		"60ab14": 869,  // LG
		"60ab67": 5697, // Xiaomi
		"60abd2": 1819, // Bose
		"60acc8": 7546, // KunTeng
		"60af6d": 152,  // Samsung
		"60b387": 7547, // Synergics
		"60b606": 7548, // Phorus
		"60b617": 4921, // Fiberhome
		"60b76e": 3521, // Google
		"60b933": 7549, // Deutron
		"60ba18": 7550, // nextLAP
		"60beb4": 7551, // S-Bluetech
		"60beb5": 687,  // Motorola
		"60bec4": 545,  // Apple
		"60c397": 1939, // 2Wire
		"60c547": 545,  // Apple
		"60c5ad": 152,  // Samsung
		"60c5e6": 7039, // Skullcandy
		"60c658": 7552, // PHYTRONIX
		"60c798": 1647, // Verifone
		"60c980": 7553, // Trymus
		"60cbfb": 7554, // AirScape
		"60cda9": 7555, // Abloomy
		"60ce41": 3334, // Huawei
		"60ce86": 2074, // Sercomm
		"60ce92": 7556, // Refined
		"60d02c": 2737, // Ruckus
		"60d0a9": 152,  // Samsung
		"60d21c": 5435, // Sunnovo
		"60d248": 125,  // ARRIS
		"60d262": 7557, // Tzukuri
		"60d30a": 7558, // Quatius
		"60d755": 3334, // Huawei
		"60d9a0": 2664, // Lenovo
		"60d9c7": 545,  // Apple
		"60da23": 7559, // Estech
		"60db2a": 7560, // HNS
		"60db98": 374,  // Calix
		"60dd70": 545,  // Apple
		"60dd8e": 421,  // Intel
		"60de35": 7561, // GITSN
		"60de44": 3334, // Huawei
		"60def3": 3334, // Huawei
		"60e00e": 4161, // Shinsei
		"60e327": 1595, // TP-Link
		"60e32b": 421,  // Intel
		"60e3ac": 869,  // LG
		"60e6bc": 7562, // Sino-Telecom
		"60e6f0": 1592, // Wistron
		"60e701": 3334, // Huawei
		"60e78a": 7563, // Unisem
		"60e956": 7564, // Ayla
		"60eb69": 3071, // Quanta
		"60f189": 2056, // Murata
		"60f18a": 3334, // Huawei
		"60f262": 421,  // Intel
		"60f281": 7565, // Tranwo
		"60f2ef": 7566, // VisionVera
		"60f43a": 6445, // Edifier
		"60f445": 545,  // Apple
		"60f59c": 7567, // CRU-Dataport
		"60f673": 7568, // Terumo
		"60f677": 421,  // Intel
		"60f81d": 545,  // Apple
		"60f8f2": 7569, // Synaptec
		"60fa9d": 3334, // Huawei
		"60facd": 545,  // Apple
		"60fb42": 545,  // Apple
		"60fcf1": 67,   // Private
		"60fd56": 7570, // WOORISYSTEMS
		"60fe20": 1939, // 2Wire
		"60fec5": 545,  // Apple
		"60ff12": 152,  // Samsung
		"60ffdd": 7571, // C.E
		"64002d": 7572, // Powerlinq
		"64006a": 102,  // Dell
		"6400f1": 2,    // Cisco
		"6401fb": 2227, // Landis+Gyr
		"6402cb": 125,  // ARRIS
		"64037f": 152,  // Samsung
		"6405e4": 430,  // Alpsalpine
		"6407f6": 152,  // Samsung
		"640980": 5697, // Xiaomi
		"6409ac": 6433, // TCT
		"640bd7": 545,  // Apple
		"640d22": 869,  // LG
		"640de6": 7573, // Petra
		"640e36": 7574, // Taztag
		"640e94": 7575, // Pluribus
		"640f28": 1939, // 2Wire
		"641225": 2,    // Cisco
		"641236": 6392, // Technicolor
		"641269": 125,  // ARRIS
		"64136c": 3031, // ZTE
		"6413ab": 3334, // Huawei
		"641666": 6635, // Nest
		"64167f": 751,  // Polycom
		"64168d": 2,    // Cisco
		"6416f0": 3334, // Huawei
		"641759": 7576, // Intellivision
		"641a22": 7577, // Heliospectra
		"641aba": 7578, // Dryad
		"641cae": 152,  // Samsung
		"641cb0": 152,  // Samsung
		"64200c": 545,  // Apple
		"64209f": 377,  // Tilgin
		"6420e0": 7579, // T3
		"642315": 3334, // Huawei
		"642400": 7580, // Xorcom
		"64255e": 7365, // Observint
		"642753": 3334, // Huawei
		"642c0f": 6369, // Vivo
		"642cac": 3334, // Huawei
		"642db7": 7581, // Seungil
		"643150": 302,  // HP
		"64317e": 7582, // Dexin
		"643216": 7583, // Weidu
		"6432a8": 421,  // Intel
		"643409": 7584, // BITwave
		"643aea": 2,    // Cisco
		"643e8c": 3334, // Huawei
		"643f5f": 7585, // Exablaze
		"6444d5": 67,   // Private
		"6447e0": 7586, // Feitian
		"644bf0": 7587, // CalDigit
		"644c36": 421,  // Intel
		"644d70": 7588, // dSPACE
		"644f42": 7589, // JETTER
		"644f74": 7590, // LENUS
		"644fb0": 7591, // Hyunjin.com
		"6450d6": 7592, // Liquidtool
		"645106": 302,  // HP
		"645563": 7593, // Intelight
		"6455b1": 125,  // ARRIS
		"645601": 1595, // TP-Link
		"645a04": 7300, // Chicony
		"645a36": 545,  // Apple
		"645aed": 545,  // Apple
		"645cf3": 7594, // ParanTek
		"645d86": 421,  // Intel
		"645e10": 3334, // Huawei
		"645e2c": 7595, // IRay
		"646140": 3334, // Huawei
		"646184": 7596, // Velux
		"646223": 7597, // Cellient
		"64628a": 7598, // evon
		"64649b": 826,  // Juniper
		"6465c0": 7599, // Nuvon
		"646624": 2048, // Sagemcom
		"6466b3": 1595, // TP-Link
		"64680c": 3869, // Comtrend
		"6469bc": 7600, // Hytera
		"646a52": 620,  // Avaya
		"646a74": 7601, // Auth-Servers
		"646cb2": 152,  // Samsung
		"646d2f": 545,  // Apple
		"646d4e": 3334, // Huawei
		"646d6c": 3334, // Huawei
		"646e69": 2873, // Liteon
		"646e97": 1595, // TP-Link
		"646ee0": 421,  // Intel
		"647002": 1595, // TP-Link
		"647033": 545,  // Apple
		"6472d8": 7602, // GooWi
		"6473e2": 7603, // Arbiter
		"6476ba": 545,  // Apple
		"64777d": 870,  // Hitron
		"647791": 152,  // Samsung
		"647924": 3334, // Huawei
		"6479a7": 7604, // Phison
		"6479f0": 421,  // Intel
		"647bce": 152,  // Samsung
		"647d81": 7605, // Yokota
		"647fda": 7606, // TEKTELIC
		"648099": 421,  // Intel
		"648788": 826,  // Juniper
		"64899a": 869,  // LG
		"6489f1": 152,  // Samsung
		"648d9e": 2990, // IVT
		"64956c": 869,  // LG
		"649714": 5820, // eero
		"64995d": 7468, // LGE
		"649968": 4652, // Elentec
		"649a12": 7607, // P2
		"649abe": 545,  // Apple
		"649b24": 7608, // V
		"649c81": 5787, // Qualcomm
		"649ef3": 2,    // Cisco
		"64a0e7": 2,    // Cisco
		"64a198": 3334, // Huawei
		"64a200": 5697, // Xiaomi
		"64a28a": 3334, // Huawei
		"64a2f9": 6923, // OnePlus
		"64a341": 7609, // Wonderlan
		"64a3cb": 545,  // Apple
		"64a5c3": 545,  // Apple
		"64a651": 3334, // Huawei
		"64a769": 1341, // HTC
		"64a7dd": 620,  // Avaya
		"64a965": 7610, // Linkflow
		"64ae0c": 2,    // Cisco
		"64ae88": 7611, // Polytec
		"64b0a6": 545,  // Apple
		"64b0e8": 3334, // Huawei
		"64b310": 152,  // Samsung
		"64b370": 7612, // PowerComm
		"64b379": 67,   // Private
		"64b473": 5697, // Xiaomi
		"64b5c6": 1427, // Nintendo
		"64b64a": 7613, // ViVOtech
		"64b853": 152,  // Samsung
		"64b94e": 102,  // Dell
		"64b9e8": 545,  // Apple
		"64babd": 7614, // SDJ
		"64bc0c": 869,  // LG
		"64bc11": 7615, // CombiQ
		"64bc58": 421,  // Intel
		"64be63": 7616, // STORDIS
		"64bf6b": 3334, // Huawei
		"64c2de": 869,  // LG
		"64c354": 620,  // Avaya
		"64c394": 3334, // Huawei
		"64c3d6": 826,  // Juniper
		"64c6af": 7617, // AXERRA
		"64c753": 545,  // Apple
		"64c901": 7618, // INVENTEC
		"64c944": 7619, // LARK
		"64cb9f": 6265, // Tecno
		"64cba3": 4612, // Pointmobile
		"64cbe9": 869,  // LG
		"64cc22": 2639, // Arcadyan
		"64cc2e": 5697, // Xiaomi
		"64d0d6": 152,  // Samsung
		"64d154": 1784, // Routerboard.com
		"64d1a3": 1881, // Sitecom
		"64d2c4": 545,  // Apple
		"64d4bd": 430,  // Alpsalpine
		"64d4da": 421,  // Intel
		"64d69a": 421,  // Intel
		"64d7c0": 3334, // Huawei
		"64d814": 2,    // Cisco
		"64d912": 7620, // Solidica
		"64d989": 2,    // Cisco
		"64db18": 7621, // OpenPattern
		"64db43": 687,  // Motorola
		"64db81": 7622, // Syszone
		"64dde9": 5697, // Xiaomi
		"64de1c": 7623, // Kingnetic
		"64dfe9": 7624, // Ateme
		"64e0ab": 3505, // Unionman
		"64e161": 7625, // DEP
		"64e599": 1252, // EFM
		"64e682": 545,  // Apple
		"64e7d8": 152,  // Samsung
		"64e84f": 7626, // Serialway
		"64e881": 1685, // Aruba
		"64e950": 2,    // Cisco
		"64eb8c": 46,   // Epson
		"64ed57": 125,  // ARRIS
		"64ed62": 7627, // WOORI
		"64eeb7": 2362, // Netcore
		"64f50e": 7628, // Kinion
		"64f69d": 2,    // Cisco
		"64f705": 3334, // Huawei
		"64f81c": 3334, // Huawei
		"64f970": 7629, // Kenade
		"64f987": 7630, // Avvasi
		"64fb50": 7631, // RoomReady/Zdi
		"64fc8c": 7632, // Zonar
		"64ff0a": 1592, // Wistron
		"680235": 7633, // Konten
		"6802b8": 358,  // Compal
		"680571": 152,  // Samsung
		"6805ca": 421,  // Intel
		"680715": 421,  // Intel
		"680927": 545,  // Apple
		"680ae2": 1655, // Silicon
		"681324": 3334, // Huawei
		"681590": 2048, // Sagemcom
		"681605": 7634, // Systems
		"681729": 421,  // Intel
		"681ab2": 3031, // ZTE
		"681bef": 3334, // Huawei
		"681ca2": 7635, // Rosewill
		"681d64": 7636, // Sunwave
		"681e8b": 7637, // InfoSight
		"681fd8": 300,  // Siemens
		"68215f": 6294, // Edgecore
		"68228e": 826,  // Juniper
		"682719": 706,  // Microchip
		"682737": 152,  // Samsung
		"6828ba": 7638, // Dejai
		"6828f6": 7639, // Vubiq
		"6829dc": 7640, // Ficosa
		"682c7b": 2,    // Cisco
		"682f67": 545,  // Apple
		"6831fe": 7641, // Teladin
		"68332c": 7642, // Kenstel
		"6836b5": 7643, // DriveScale
		"6837e9": 5438, // Amazon
		"683943": 7499, // ittim
		"683a48": 6811, // SAMJIN
		"683b1e": 7644, // Countwise
		"683b78": 2,    // Cisco
		"683e02": 7645, // SIEMENS
		"683e26": 421,  // Intel
		"683e34": 7030, // MEIZU
		"683eec": 7646, // Ereca
		"68403c": 4921, // Fiberhome
		"684352": 7647, // Bhuu
		"684571": 3334, // Huawei
		"6845f1": 35,   // Toshiba
		"684898": 152,  // Samsung
		"684a76": 5820, // eero
		"684aae": 3334, // Huawei
		"684ae9": 152,  // Samsung
		"684f64": 102,  // Dell
		"6851b7": 7648, // PowerCloud
		"6852d6": 7649, // UGame
		"68536c": 7650, // SPnS
		"685388": 7651, // P&S
		"68545a": 421,  // Intel
		"6854c1": 7001, // ColorTokens
		"6854f5": 7652, // enLighted
		"6854fd": 5438, // Amazon
		"685811": 4921, // Fiberhome
		"685acf": 152,  // Samsung
		"685b35": 545,  // Apple
		"685b36": 2951, // Powertech
		"685d43": 421,  // Intel
		"685e6b": 7653, // PowerRay
		"68644b": 545,  // Apple
		"686725": 6376, // Espressif
		"68692e": 7654, // Zycoo
		"686975": 7655, // Angler
		"6869ca": 89,   // Hitachi
		"686e23": 7656, // Wi3
		"686e48": 7657, // Prophet
		"687251": 2978, // Ubiquiti
		"6872c3": 152,  // Samsung
		"6872dc": 7658, // CETORY.TV
		"68764f": 101,  // Sony
		"687724": 1595, // TP-Link
		"687848": 7659, // Westunitis
		"68784c": 76,   // Nortel
		"687924": 7660, // ELS-GmbH
		"6879ed": 3205, // Sharp
		"687d6b": 152,  // Samsung
		"687db4": 2,    // Cisco
		"687f74": 1783, // Linksys
		"6881e0": 3334, // Huawei
		"6882f2": 7661, // grandcentrix
		"68831a": 7662, // Pandora
		"688470": 7663, // eSSys
		"68847e": 4,    // Fujitsu
		"688540": 7664, // IGI
		"68856a": 7665, // OuterLink
		"6886a7": 2,    // Cisco
		"6886e7": 7666, // Orbotix
		"68876b": 7667, // INQ
		"6887c6": 2,    // Cisco
		"6888a1": 3857, // Universal
		"688975": 7668, // nuoxc
		"6889c1": 3334, // Huawei
		"688af0": 3031, // ZTE
		"688db6": 7669, // Aetek
		"688f2e": 870,  // Hitron
		"688f84": 3334, // Huawei
		"688fc9": 7670, // Zhuolian
		"689234": 2737, // Ruckus
		"68966a": 6577, // Ohsung
		"68967b": 545,  // Apple
		"689861": 7671, // Beacon
		"6899cd": 2,    // Cisco
		"689a87": 5438, // Amazon
		"689c5e": 4903, // AcSiP
		"689c70": 545,  // Apple
		"689ce2": 2,    // Cisco
		"689e0b": 2,    // Cisco
		"689e6a": 3334, // Huawei
		"689ff0": 3031, // ZTE
		"68a03e": 3334, // Huawei
		"68a0f6": 3334, // Huawei
		"68a3c4": 2873, // Liteon
		"68a828": 3334, // Huawei
		"68a86d": 545,  // Apple
		"68a878": 7672, // GeoWAN
		"68a8e1": 7186, // Wacom
		"68aad2": 7673, // Datecs
		"68ab09": 457,  // Nokia
		"68ab1e": 545,  // Apple
		"68ae20": 545,  // Apple
		"68af13": 7674, // Futura
		"68b094": 7675, // Inesa
		"68b43a": 7676, // WaterFurnace
		"68b599": 302,  // HP
		"68b691": 5438, // Amazon
		"68b6fc": 870,  // Hitron
		"68b983": 7677, // b-plus
		"68bc0c": 2,    // Cisco
		"68bdab": 2,    // Cisco
		"68bfc4": 152,  // Samsung
		"68c44d": 687,  // Motorola
		"68c63a": 6376, // Espressif
		"68ca00": 7678, // Octopus
		"68cae4": 2,    // Cisco
		"68cc6e": 3334, // Huawei
		"68d48b": 7679, // Hailo
		"68d79a": 2978, // Ubiquiti
		"68d93c": 545,  // Apple
		"68db54": 6848, // Phicomm
		"68db96": 7680, // OPWILL
		"68dbca": 545,  // Apple
		"68dbf5": 5438, // Amazon
		"68dce8": 7681, // PacketStorm
		"68dfdd": 5697, // Xiaomi
		"68e166": 67,   // Private
		"68e209": 3334, // Huawei
		"68e7c2": 152,  // Samsung
		"68e8eb": 7682, // Linktel
		"68ebae": 152,  // Samsung
		"68ec62": 7683, // YODO
		"68ec8a": 67,   // Private
		"68ecc5": 421,  // Intel
		"68ed43": 2220, // BlackBerry
		"68ef43": 545,  // Apple
		"68efbd": 2,    // Cisco
		"68f06d": 7080, // Along
		"68f0d0": 7684, // SkyBell
		"68f38e": 826,  // Juniper
		"68f728": 4919, // LCFC
		"68f895": 7685, // Redflow
		"68fb7e": 545,  // Apple
		"68fb95": 7686, // Generalplus
		"68fcca": 152,  // Samsung
		"68feda": 4921, // Fiberhome
		"68fef7": 545,  // Apple
		"68ff7b": 1595, // TP-Link
		"6c006b": 152,  // Samsung
		"6c02e0": 302,  // HP
		"6c0309": 2,    // Cisco
		"6c047a": 3334, // Huawei
		"6c05d5": 7687, // Ethertronics
		"6c06d6": 3334, // Huawei
		"6c09bf": 4921, // Fiberhome
		"6c09d6": 7688, // Digiquest
		"6c0d34": 457,  // Nokia
		"6c0e0d": 101,  // Sony
		"6c0f6a": 7689, // JDC
		"6c108b": 4398, // WeLink
		"6c13d5": 2,    // Cisco
		"6c1414": 7690, // BUJEON
		"6c146e": 3334, // Huawei
		"6c14f7": 7691, // Erhardt+Leimer
		"6c15f9": 7692, // Nautronix
		"6c160e": 7693, // ShotTracker
		"6c1632": 3334, // Huawei
		"6c1811": 7694, // Decatur
		"6c198f": 803,  // D-Link
		"6c19c0": 545,  // Apple
		"6c1a75": 3334, // Huawei
		"6c1b3f": 7695, // MiraeSignal
		"6c1deb": 7419, // u-blox
		"6c1ed7": 6369, // Vivo
		"6c2056": 2,    // Cisco
		"6c21a2": 4498, // AMPAK
		"6c23b9": 101,  // Sony
		"6c23cb": 7696, // Wattty
		"6c2408": 4919, // LCFC
		"6c2483": 612,  // Microsoft
		"6c24a6": 6369, // Vivo
		"6c2636": 3334, // Huawei
		"6c2779": 612,  // Microsoft
		"6c2995": 421,  // Intel
		"6c2b59": 102,  // Dell
		"6c2e33": 7697, // Accelink
		"6c2e85": 2048, // Sagemcom
		"6c2f2c": 152,  // Samsung
		"6c2f8a": 152,  // Samsung
		"6c310e": 2,    // Cisco
		"6c32de": 7698, // Indieon
		"6c33a9": 7699, // Magicjack
		"6c3491": 3334, // Huawei
		"6c3845": 4921, // Fiberhome
		"6c3a36": 7700, // Glowforge
		"6c3b6b": 1784, // Routerboard.com
		"6c3be5": 302,  // HP
		"6c3c53": 7701, // SoundHawk
		"6c3c7c": 87,   // Canon
		"6c3e6d": 545,  // Apple
		"6c4008": 545,  // Apple
		"6c410e": 2,    // Cisco
		"6c416a": 2,    // Cisco
		"6c42ab": 7702, // Subscriber
		"6c433c": 6265, // Tecno
		"6c4418": 7703, // Zappware
		"6c442a": 3334, // Huawei
		"6c4598": 7704, // Antex
		"6c4760": 3939, // Sunitec
		"6c49c1": 7705, // o2ones
		"6c4a39": 7706, // Bita
		"6c4a74": 7707, // Aerodisk
		"6c4a85": 545,  // Apple
		"6c4b90": 7708, // LiteON
		"6c4bb4": 531,  // HUMAX
		"6c4d73": 545,  // Apple
		"6c504d": 2,    // Cisco
		"6c51bf": 3334, // Huawei
		"6c54cd": 7709, // Lampex
		"6c558d": 3334, // Huawei
		"6c55e8": 6392, // Technicolor
		"6c5640": 7460, // BLU
		"6c5697": 5438, // Amazon
		"6c5779": 7710, // Aclima
		"6c5940": 4253, // Mercury
		"6c5ab0": 1595, // TP-Link
		"6c5cde": 7711, // SunReports
		"6c5d3a": 612,  // Microsoft
		"6c5e3b": 2,    // Cisco
		"6c5f1c": 2664, // Lenovo
		"6c6126": 7712, // Rinicom
		"6c61f4": 3187, // SFR
		"6c626d": 1812, // Micro-Star
		"6c639c": 125,  // ARRIS
		"6c67ef": 3334, // Huawei
		"6c6a77": 421,  // Intel
		"6c6c0f": 3334, // Huawei
		"6c6cd3": 2,    // Cisco
		"6c6d09": 7713, // Kyowa
		"6c6f18": 7714, // Stereotaxis
		"6c7039": 7715, // Novar
		"6c709f": 545,  // Apple
		"6c710d": 2,    // Cisco
		"6c71d2": 3334, // Huawei
		"6c71d9": 3004, // Azurewave
		"6c7220": 803,  // D-Link
		"6c72e7": 545,  // Apple
		"6c750d": 7716, // WiFiSONG
		"6c7637": 3334, // Huawei
		"6c7660": 2848, // Kyocera
		"6c81fe": 7717, // Mitsuba
		"6c8336": 152,  // Samsung
		"6c8686": 7718, // Technonia
		"6c8814": 421,  // Intel
		"6c8b2f": 3031, // ZTE
		"6c8bd3": 2,    // Cisco
		"6c8cdb": 7719, // Otus
		"6c8d77": 2,    // Cisco
		"6c8dc1": 545,  // Apple
		"6c8fb5": 612,  // Microsoft
		"6c90b1": 7720, // SanLogic
		"6c9106": 67,   // Private
		"6c92bf": 7721, // Inspur
		"6c9354": 7722, // Yaojin
		"6c9392": 7723, // BEKO
		"6c9466": 421,  // Intel
		"6c94f8": 545,  // Apple
		"6c9522": 7724, // Scalys
		"6c96cf": 545,  // Apple
		"6c98eb": 2093, // Riverbed
		"6c9961": 2048, // Sagemcom
		"6c9989": 2,    // Cisco
		"6c999d": 5438, // Amazon
		"6c9ac9": 7725, // Valentine
		"6c9b02": 457,  // Nokia
		"6c9bc0": 7726, // Chemoptics
		"6c9ced": 2,    // Cisco
		"6c9e7c": 4921, // Fiberhome
		"6ca0b4": 3512, // BSkyB
		"6ca100": 421,  // Intel
		"6ca401": 7727, // essensys
		"6ca4d1": 4921, // Fiberhome
		"6ca604": 125,  // ARRIS
		"6ca75f": 3031, // ZTE
		"6ca780": 457,  // Nokia
		"6ca7fa": 2000, // Youngbo
		"6ca849": 620,  // Avaya
		"6ca858": 4921, // Fiberhome
		"6ca906": 7728, // Telefield
		"6ca936": 2232, // DisplayLink
		"6ca96f": 7729, // TransPacket
		"6caab3": 2737, // Ruckus
		"6cab05": 2,    // Cisco
		"6cab31": 545,  // Apple
		"6cac60": 7730, // Venetex
		"6cad3f": 7731, // Hubbell
		"6cadf8": 3004, // Azurewave
		"6cae8b": 372,  // IBM
		"6caee3": 457,  // Nokia
		"6caef6": 5820, // eero
		"6cb0ce": 1368, // Netgear
		"6cb158": 1595, // TP-Link
		"6cb227": 101,  // Sony
		"6cb2ae": 2,    // Cisco
		"6cb4a7": 7732, // Landauer
		"6cb4fd": 3334, // Huawei
		"6cb56b": 531,  // HUMAX
		"6cb6ca": 7733, // DivUS
		"6cb749": 3334, // Huawei
		"6cb7f4": 152,  // Samsung
		"6cb881": 3031, // ZTE
		"6cb9c5": 957,  // Delta
		"6cbab8": 2048, // Sagemcom
		"6cbfb5": 7734, // Noon
		"6cc1d2": 125,  // ARRIS
		"6cc217": 302,  // HP
		"6cc26b": 545,  // Apple
		"6cc49f": 1685, // Aruba
		"6cc7ec": 152,  // Samsung
		"6cca08": 125,  // ARRIS
		"6ccdd6": 1368, // Netgear
		"6cce44": 6627, // 1MORE
		"6cd032": 869,  // LG
		"6cd146": 7735, // FRAMOS
		"6cd1e5": 3334, // Huawei
		"6cd2ba": 3031, // ZTE
		"6cd3ee": 7736, // Zimi
		"6cd68a": 869,  // LG
		"6cd704": 3334, // Huawei
		"6cd719": 4921, // Fiberhome
		"6cd94c": 6369, // Vivo
		"6cdc6a": 7737, // Promethean
		"6cdd30": 2,    // Cisco
		"6cddbc": 152,  // Samsung
		"6cddef": 7738, // EPCOMM
		"6ce01e": 7739, // Modcam
		"6ce0b0": 7740, // SOUND4
		"6ce5c9": 545,  // Apple
		"6ce85c": 545,  // Apple
		"6ce873": 1595, // TP-Link
		"6ce874": 3334, // Huawei
		"6ce8c6": 6498, // Earda
		"6ce907": 457,  // Nokia
		"6ce983": 7741, // Gastron
		"6cebb6": 3334, // Huawei
		"6ced51": 7742, // NEXCONTROL
		"6cf049": 1929, // Giga-Byte
		"6cf373": 152,  // Samsung
		"6cf37f": 1685, // Aruba
		"6cf5e8": 7743, // Mooredoll
		"6cf784": 5697, // Xiaomi
		"6cf97c": 7744, // Nanoptix
		"6cfa58": 620,  // Avaya
		"6cfa89": 2,    // Cisco
		"6cfaa7": 4498, // AMPAK
		"6cfdb9": 3182, // Proware
		"6cfe54": 421,  // Intel
		"6cffbe": 7745, // MPB
		"6cffce": 2048, // Sagemcom
		"7001b5": 2,    // Cisco
		"700258": 7746, // 01DB-Metravib
		"70037e": 6392, // Technicolor
		"70039f": 6376, // Espressif
		"70041d": 6376, // Espressif
		"700514": 869,  // LG
		"7006ac": 7747, // Eastcompeace
		"700777": 7748, // OnTarget
		"700971": 152,  // Samsung
		"700b01": 2048, // Sagemcom
		"700b4f": 2,    // Cisco
		"700bc0": 7749, // Dewav
		"700f6a": 2,    // Cisco
		"700fec": 7750, // Poindus
		"70105c": 2,    // Cisco
		"70106f": 302,  // HP
		"701124": 545,  // Apple
		"701135": 7751, // Livesecu
		"7014a6": 545,  // Apple
		"7018a7": 2,    // Cisco
		"70192f": 3334, // Huawei
		"701a04": 2873, // Liteon
		"701ab8": 421,  // Intel
		"701aed": 7752, // Advas
		"701ce7": 421,  // Intel
		"701d7f": 4260, // Comtech
		"701f3c": 152,  // Samsung
		"701f53": 2,    // Cisco
		"702393": 7753, // fos4X
		"702526": 457,  // Nokia
		"702559": 189,  // CyberTAN
		"70288b": 152,  // Samsung
		"702a7d": 7754, // EpSpot
		"702ad5": 152,  // Samsung
		"702b1d": 7755, // E-Domus
		"702c09": 1427, // Nintendo
		"702c1f": 7235, // Wisol
		"702dd1": 7756, // Newings
		"702e22": 3031, // ZTE
		"702f35": 3334, // Huawei
		"702f4b": 7757, // Steelcase
		"702f97": 7758, // Aava
		"703018": 620,  // Avaya
		"70305d": 1103, // Ubiquoss
		"703187": 7759, // ACX
		"703217": 421,  // Intel
		"703509": 2,    // Cisco
		"703811": 300,  // Siemens
		"7038b4": 7760, // Low
		"7038ee": 620,  // Avaya
		"703a0e": 1685, // Aruba
		"703a51": 5697, // Xiaomi
		"703acb": 3521, // Google
		"703c03": 7761, // RadiAnt
		"703c69": 545,  // Apple
		"703e97": 6868, // Iton
		"703eac": 545,  // Apple
		"7040ff": 3334, // Huawei
		"7047e9": 6369, // Vivo
		"70480f": 545,  // Apple
		"7048f7": 1427, // Nintendo
		"704a0e": 4498, // AMPAK
		"704ae4": 7762, // Rinstrum
		"704ca5": 1323, // Fortinet
		"704ced": 7763, // TMRG
		"704d7b": 1806, // ASUS
		"704e01": 7764, // Kwangwon
		"704e6b": 3334, // Huawei
		"704f57": 1595, // TP-Link
		"704fb8": 125,  // ARRIS
		"7050af": 3512, // BSkyB
		"7052c5": 620,  // Avaya
		"7052d8": 6281, // Itel
		"705425": 125,  // ARRIS
		"7054b4": 1449, // Vestel
		"7054d2": 5439, // Pegatron
		"7054f5": 3334, // Huawei
		"7055f8": 7765, // Cerebras
		"705681": 545,  // Apple
		"705812": 2153, // Panasonic
		"705896": 7766, // InShow
		"705a0f": 302,  // HP
		"705a9e": 6392, // Technicolor
		"705aac": 152,  // Samsung
		"705ab6": 358,  // Compal
		"705b2e": 7767, // M2Communication
		"705dcc": 1252, // EFM
		"705fa3": 5697, // Xiaomi
		"7060de": 7768, // LaVision
		"706173": 7769, // Calantec
		"70617b": 2,    // Cisco
		"7061be": 1592, // Wistron
		"7061ee": 7770, // Sunwoda
		"7062b8": 803,  // D-Link
		"70661b": 6859, // Sonova
		"706655": 3004, // Azurewave
		"70695a": 2,    // Cisco
		"706bb9": 2,    // Cisco
		"706d15": 2,    // Cisco
		"706dec": 7771, // Wifi-soft
		"706e6d": 2,    // Cisco
		"706f81": 67,   // Private
		"70700d": 545,  // Apple
		"70704c": 7772, // Purple
		"70708b": 2,    // Cisco
		"7070aa": 5438, // Amazon
		"7071b3": 7773, // Brain
		"7071bc": 5439, // Pegatron
		"70720d": 2664, // Lenovo
		"70723c": 3334, // Huawei
		"7072cf": 7774, // EdgeCore
		"7073cb": 545,  // Apple
		"707414": 2056, // Murata
		"707630": 125,  // ARRIS
		"7076f0": 7775, // LevelOne
		"7076ff": 7776, // Kerlink
		"70788b": 6369, // Vivo
		"707990": 3334, // Huawei
		"7079b3": 2,    // Cisco
		"707be8": 3334, // Huawei
		"707c18": 7777, // ADATA
		"707c69": 620,  // Avaya
		"707db9": 2,    // Cisco
		"707e43": 125,  // ARRIS
		"707ede": 7778, // Nastec
		"708105": 2,    // Cisco
		"7081eb": 545,  // Apple
		"70820e": 7779, // as
		"70828e": 7780, // OleumTech
		"7085c2": 4719, // ASRock
		"7085c4": 5440, // Ruijie
		"7085c6": 125,  // ARRIS
		"70879e": 7781, // Beken
		"7087a7": 2056, // Murata
		"708a09": 3334, // Huawei
		"708b78": 7782, // citygrow
		"708bcd": 1806, // ASUS
		"708cb6": 3334, // Huawei
		"708cbb": 7783, // Mimodisplaykorea
		"708d09": 457,  // Nokia
		"708f47": 6369, // Vivo
		"7090b7": 3334, // Huawei
		"70918f": 7784, // Weber-Stephen
		"7091f3": 3857, // Universal
		"709741": 2639, // Arcadyan
		"709756": 7785, // Happyelectronics
		"709bfc": 7786, // Bryton
		"709c8f": 7787, // Nero
		"709cd1": 421,  // Intel
		"709e29": 101,  // Sony
		"709e86": 7788, // X6D
		"709f2d": 3031, // ZTE
		"709fa9": 6265, // Tecno
		"70a2b3": 545,  // Apple
		"70a6cc": 421,  // Intel
		"70a741": 2978, // Ubiquiti
		"70a84c": 7789, // MONAD
		"70a8d3": 421,  // Intel
		"70a8e3": 3334, // Huawei
		"70aab2": 2220, // BlackBerry
		"70af25": 7790, // Nishiyama
		"70b13d": 152,  // Samsung
		"70b14e": 125,  // ARRIS
		"70b317": 2,    // Cisco
		"70b5e8": 102,  // Dell
		"70b7aa": 6369, // Vivo
		"70b8f6": 6376, // Espressif
		"70b921": 4921, // Fiberhome
		"70bbe9": 5697, // Xiaomi
		"70bc10": 612,  // Microsoft
		"70c7f2": 3334, // Huawei
		"70c833": 7791, // Wirepas
		"70c94e": 2873, // Liteon
		"70c9c6": 2,    // Cisco
		"70ca97": 2737, // Ruckus
		"70ca9b": 2,    // Cisco
		"70cd0d": 421,  // Intel
		"70cd60": 545,  // Apple
		"70ce8c": 152,  // Samsung
		"70cf49": 421,  // Intel
		"70d313": 3334, // Huawei
		"70d379": 2,    // Cisco
		"70d4f2": 4555, // RIM
		"70d57e": 7792, // Scalar
		"70d5e7": 7793, // Wellcore
		"70d6b6": 7794, // Metrum
		"70d923": 6369, // Vivo
		"70d931": 1640, // Cambridge
		"70da9c": 7795, // Tecsen
		"70db98": 2,    // Cisco
		"70dda1": 5027, // Tellabs
		"70ddef": 3334, // Huawei
		"70dee2": 545,  // Apple
		"70df2f": 2,    // Cisco
		"70dff7": 125,  // ARRIS
		"70e027": 7796, // Hongyu
		"70e139": 7797, // 3view
		"70e1fd": 833,  // Flextronics
		"70e284": 1592, // Wistron
		"70e422": 2,    // Cisco
		"70e72c": 545,  // Apple
		"70ea1a": 2,    // Cisco
		"70ea5a": 545,  // Apple
		"70ece4": 545,  // Apple
		"70ee50": 7798, // Netatmo
		"70eea3": 7799, // Eoptolink
		"70ef00": 545,  // Apple
		"70f087": 545,  // Apple
		"70f088": 1427, // Nintendo
		"70f096": 2,    // Cisco
		"70f196": 2246, // Actiontec
		"70f1a1": 2873, // Liteon
		"70f1e5": 7800, // Xetawave
		"70f220": 2246, // Actiontec
		"70f35a": 2,    // Cisco
		"70f754": 4498, // AMPAK
		"70f82b": 7446, // DWnet
		"70f927": 152,  // Samsung
		"70fc8c": 2658, // OneAccess
		"70fd45": 3334, // Huawei
		"70fd46": 152,  // Samsung
		"7403bd": 1077, // Buffalo
		"74042b": 2664, // Lenovo
		"7404f1": 421,  // Intel
		"7405a5": 1595, // TP-Link
		"7409ac": 7801, // Quext
		"740abc": 7802, // LightwaveRF
		"740ae1": 3334, // Huawei
		"740cee": 3334, // Huawei
		"740edb": 7803, // Optowiz
		"7411b2": 2,    // Cisco
		"7412bb": 4921, // Fiberhome
		"741575": 5697, // Xiaomi
		"7415e2": 7804, // Tri-Sen
		"741bb2": 545,  // Apple
		"741c27": 6281, // Itel
		"741e93": 4921, // Fiberhome
		"741f79": 7805, // Youngkook
		"7422bb": 3334, // Huawei
		"742344": 5697, // Xiaomi
		"7426ac": 2,    // Cisco
		"7426ff": 3031, // ZTE
		"74273c": 7806, // ChangYang
		"7427ea": 1115, // Elitegroup
		"742b0f": 7807, // Infinidat
		"742b62": 4,    // Fujitsu
		"742edb": 7808, // Perinet
		"742efc": 7809, // DirectPacket
		"742f68": 3004, // Azurewave
		"743170": 2639, // Arcadyan
		"7432c2": 7810, // Kyolis
		"743400": 7811, // MTG
		"74373b": 7812, // UNINET
		"7438b7": 87,   // Canon
		"743a65": 48,   // NEC
		"743aef": 1269, // Kaonmedia
		"743e2b": 2737, // Ruckus
		"743ecb": 7813, // Gentrice
		"7440be": 869,  // LG
		"74428b": 545,  // Apple
		"744401": 1368, // Netgear
		"74452d": 3334, // Huawei
		"74458a": 152,  // Samsung
		"7445ce": 3282, // Cresyn
		"744687": 7814, // Kingsignal
		"7446a0": 302,  // HP
		"744aa4": 3031, // ZTE
		"744ca1": 2873, // Liteon
		"744d28": 1784, // Routerboard.com
		"744d79": 7815, // Arrive
		"7451ba": 5697, // Xiaomi
		"745327": 7816, // Commsen
		"745612": 125,  // ARRIS
		"7458f3": 5438, // Amazon
		"745909": 3334, // Huawei
		"745933": 7817, // Danal
		"745aaa": 3334, // Huawei
		"745c9f": 6433, // TCT
		"745d68": 4921, // Fiberhome
		"745e1c": 6117, // Pioneer
		"745f90": 7818, // LAM
		"7460fa": 3334, // Huawei
		"7463c2": 3334, // Huawei
		"7463df": 7819, // VTS
		"74650c": 545,  // Apple
		"7465d1": 7820, // Atlinks
		"7467f7": 185,  // Extreme
		"746a3a": 7821, // Aperi
		"746a89": 7822, // Rezolt
		"746b82": 7823, // Movek
		"746f19": 7824, // Icarvisions
		"746f3d": 5482, // Contec
		"746ff7": 1592, // Wistron
		"747069": 3334, // Huawei
		"7470fd": 421,  // Intel
		"74721e": 7825, // Edison
		"7472f2": 7826, // Chipsip
		"74731d": 7827, // ifm
		"747336": 7828, // MICRODIGTAL
		"747446": 3521, // Google
		"747548": 5438, // Amazon
		"747818": 7829, // Jurumani
		"747827": 102,  // Dell
		"747a90": 2056, // Murata
		"747b7a": 7830, // ETH
		"747d24": 6848, // Phicomm
		"747db6": 7831, // Aliwei
		"748114": 545,  // Apple
		"7483c2": 2978, // Ubiquiti
		"7483ef": 3793, // Arista
		"74852a": 5439, // Pegatron
		"74860b": 2,    // Cisco
		"74867a": 102,  // Dell
		"7486e2": 102,  // Dell
		"7487a9": 7832, // OCT
		"7487bb": 4564, // Ciena
		"74882a": 3334, // Huawei
		"7488bb": 2,    // Cisco
		"748a0d": 125,  // ARRIS
		"748b29": 7833, // Micobiomed
		"748d08": 545,  // Apple
		"748e08": 7834, // Bestek
		"748ef8": 90,   // Brocade
		"748f3c": 545,  // Apple
		"74901f": 7835, // Ragile
		"749050": 5695, // Renesas
		"74911a": 2737, // Ruckus
		"7491bd": 7836, // Four
		"7493a4": 768,  // Zebra
		"74943d": 7837, // AgJunction
		"7495ec": 430,  // Alpsalpine
		"749637": 7838, // Todaair
		"749781": 3031, // ZTE
		"749975": 372,  // IBM
		"749ac0": 7839, // Cachengo
		"749be8": 870,  // Hitron
		"749d79": 2074, // Sercomm
		"749d8f": 3334, // Huawei
		"749ddc": 1939, // 2Wire
		"749ea5": 6577, // Ohsung
		"749eaf": 545,  // Apple
		"749ef5": 152,  // Samsung
		"74a02f": 2,    // Cisco
		"74a063": 3334, // Huawei
		"74a2e6": 2,    // Cisco
		"74a34a": 7736, // Zimi
		"74a528": 3334, // Huawei
		"74a722": 869,  // LG
		"74a78e": 3031, // ZTE
		"74a7ea": 5438, // Amazon
		"74acb9": 2978, // Ubiquiti
		"74ad98": 2,    // Cisco
		"74b472": 7840, // Ciesse
		"74b57e": 3031, // ZTE
		"74b587": 545,  // Apple
		"74b6b6": 5820, // eero
		"74b7e6": 7841, // Zegna-Daidong
		"74b9eb": 6945, // JinQianMao
		"74be08": 6991, // ATEK
		"74bfa1": 7842, // Hyunteck
		"74bfb7": 7843, // Nusoft
		"74bfc0": 87,   // Canon
		"74c14f": 3334, // Huawei
		"74c17d": 6295, // Infinix
		"74c246": 5438, // Amazon
		"74c63b": 3004, // Azurewave
		"74c99a": 306,  // Ericsson
		"74c9a3": 4921, // Fiberhome
		"74ca25": 7844, // Calxeda
		"74cbf3": 647,  // Lava
		"74cc39": 4921, // Fiberhome
		"74d02b": 1806, // ASUS
		"74d0dc": 306,  // Ericsson
		"74d21d": 3334, // Huawei
		"74d435": 1929, // Giga-Byte
		"74d637": 5438, // Amazon
		"74d654": 7845, // Gint
		"74d7ca": 2153, // Panasonic
		"74d83e": 421,  // Intel
		"74d850": 7846, // Evrisko
		"74da38": 115,  // Edimax
		"74da88": 1595, // TP-Link
		"74dada": 803,  // D-Link
		"74dbd1": 7847, // Ebay
		"74de2b": 2873, // Liteon
		"74dfbf": 2873, // Liteon
		"74e06e": 7848, // Ergophone
		"74e19a": 4921, // Fiberhome
		"74e1b6": 545,  // Apple
		"74e20c": 5438, // Amazon
		"74e277": 7849, // Vizmonet
		"74e28c": 612,  // Microsoft
		"74e2f5": 545,  // Apple
		"74e424": 7850, // Apiste
		"74e50b": 421,  // Intel
		"74e537": 7851, // Radspin
		"74e543": 2873, // Liteon
		"74e5f9": 421,  // Intel
		"74e60f": 6265, // Tecno
		"74e6b8": 869,  // LG
		"74e6e2": 102,  // Dell
		"74e7c6": 125,  // ARRIS
		"74e9bf": 3334, // Huawei
		"74ea3a": 1595, // TP-Link
		"74eae8": 125,  // ARRIS
		"74eb80": 152,  // Samsung
		"74ec42": 4921, // Fiberhome
		"74ecb2": 5438, // Amazon
		"74ecf1": 7852, // Acumen
		"74f06d": 3004, // Azurewave
		"74f07d": 7853, // BnCOM
		"74f2fa": 5697, // Xiaomi
		"74f612": 125,  // ARRIS
		"74f61c": 1341, // HTC
		"74f661": 56,   // Schneider
		"74f737": 7854, // KCE
		"74f91a": 7855, // Onface
		"74f9ca": 1427, // Nintendo
		"74fda0": 7856, // Compupal
		"74fe48": 1703, // Advantech
		"78009e": 152,  // Samsung
		"78028b": 545,  // Apple
		"7802b1": 2,    // Cisco
		"7802f8": 5697, // Xiaomi
		"78047a": 2842, // Edge
		"7804e3": 3334, // Huawei
		"78058c": 7857, // mMax
		"7806c9": 3334, // Huawei
		"780cb8": 421,  // Intel
		"780cf0": 2,    // Cisco
		"781100": 7858, // Quantumsolution
		"7811dc": 6280, // XIAOMI
		"7812b8": 7859, // Orantek
		"7817be": 3334, // Huawei
		"781881": 3004, // Azurewave
		"7818a8": 3334, // Huawei
		"78192e": 7860, // NASCENT
		"7819f7": 826,  // Juniper
		"781c5a": 3205, // Sharp
		"781d4a": 3031, // ZTE
		"781dba": 3334, // Huawei
		"781dfd": 7861, // Jabil
		"781fdb": 152,  // Samsung
		"782079": 7862, // ID
		"7820a5": 1427, // Nintendo
		"782184": 6376, // Espressif
		"78223d": 7863, // Affirmed
		"782327": 152,  // Samsung
		"7823ae": 125,  // ARRIS
		"7824af": 1806, // ASUS
		"782544": 7864, // Omnima
		"7825ad": 152,  // Samsung
		"7828ca": 2047, // Sonos
		"7829ed": 2544, // Askey
		"782b46": 421,  // Intel
		"782b64": 1819, // Bose
		"782bcb": 102,  // Dell
		"782d7e": 2904, // TRENDnet
		"782eef": 457,  // Nokia
		"782f17": 7865, // Xlab
		"78303b": 7866, // Stephen
		"7830e1": 7867, // UltraClenz
		"78312b": 3031, // ZTE
		"7831c1": 545,  // Apple
		"78321b": 803,  // D-Link
		"7835a0": 7868, // Zurn
		"783607": 2867, // Cermate
		"783690": 3096, // Yulong
		"783716": 152,  // Samsung
		"783a6c": 6265, // Tecno
		"783a84": 545,  // Apple
		"783ce3": 7869, // Kai-EE
		"783e53": 3512, // BSkyB
		"783f15": 7870, // EasySYNC
		"7840e4": 152,  // Samsung
		"784405": 7871, // FUJITU
		"784476": 2127, // Zioncom
		"7844fd": 1595, // TP-Link
		"784501": 7872, // Biamp
		"784558": 2978, // Ubiquiti
		"784561": 189,  // CyberTAN
		"7845b3": 3334, // Huawei
		"7845c4": 102,  // Dell
		"7846d4": 152,  // Samsung
		"78471d": 152,  // Samsung
		"784859": 302,  // HP
		"78491d": 7873, // Will-Burt
		"784b87": 2056, // Murata
		"784f43": 545,  // Apple
		"784f9b": 826,  // Juniper
		"785005": 7874, // MOKO
		"78507c": 826,  // Juniper
		"78510c": 7875, // LiveU
		"78521a": 152,  // Samsung
		"78524a": 7876, // Optonic
		"785364": 7877, // SHIFT
		"7853f2": 7878, // Roxton
		"78542e": 803,  // D-Link
		"785517": 7879, // SankyuElectronics
		"785773": 3334, // Huawei
		"785860": 3334, // Huawei
		"7858f3": 7880, // Vachen
		"78595e": 152,  // Samsung
		"785c72": 7881, // Hioso
		"785dc8": 869,  // LG
		"785ea2": 3939, // Sunitec
		"786256": 3334, // Huawei
		"7864c0": 545,  // Apple
		"786559": 2048, // Sagemcom
		"78670e": 1592, // Wistron
		"7867d7": 545,  // Apple
		"7868f7": 7450, // YSTen
		"786a1f": 125,  // ARRIS
		"786a89": 3334, // Huawei
		"786c1c": 545,  // Apple
		"787052": 7882, // Welotec
		"78719c": 125,  // ARRIS
		"78725d": 2,    // Cisco
		"787a6f": 7883, // Juice
		"787b8a": 545,  // Apple
		"787d48": 6281, // Itel
		"787d53": 185,  // Extreme
		"787df3": 7884, // Sterlite
		"787e61": 545,  // Apple
		"788102": 2074, // Sercomm
		"78843c": 101,  // Sony
		"7885f4": 3334, // Huawei
		"78886d": 545,  // Apple
		"788973": 5285, // CMC
		"788a20": 2978, // Ubiquiti
		"788c4d": 1795, // Indyme
		"788c54": 4357, // Ping
		"788c77": 613,  // Lexmark
		"788df7": 870,  // Hitron
		"7890a2": 3031, // ZTE
		"7891e9": 2050, // Raisecom
		"78923e": 457,  // Nokia
		"78929c": 421,  // Intel
		"7894b4": 2074, // Sercomm
		"7895eb": 6281, // Itel
		"789682": 3031, // ZTE
		"789684": 125,  // ARRIS
		"7898e8": 803,  // D-Link
		"7898fd": 7885, // Q9
		"78995c": 7886, // Nationz
		"789966": 7887, // Musilab
		"789ed0": 152,  // Samsung
		"789f70": 545,  // Apple
		"789f87": 300,  // Siemens
		"789faa": 3334, // Huawei
		"78a03f": 5438, // Amazon
		"78a051": 7888, // iiNet
		"78a106": 1595, // TP-Link
		"78a183": 7889, // Advidia
		"78a2a0": 1427, // Nintendo
		"78a3e4": 545,  // Apple
		"78a683": 7890, // Precidata
		"78a6e1": 90,   // Brocade
		"78a714": 1475, // Amphenol
		"78a7eb": 6627, // 1MORE
		"78a873": 152,  // Samsung
		"78abbb": 152,  // Samsung
		"78ac44": 102,  // Dell
		"78acbf": 7891, // Igneous
		"78acc0": 302,  // HP
		"78af58": 7892, // Gimasi
		"78b213": 7446, // DWnet
		"78b46a": 3334, // Huawei
		"78b554": 3334, // Huawei
		"78b8d6": 768,  // Zebra
		"78bad0": 7893, // Shinybow
		"78baf9": 2,    // Cisco
		"78bb88": 7157, // Maxio
		"78bc1a": 2,    // Cisco
		"78bdbc": 152,  // Samsung
		"78bebd": 7894, // STULZ
		"78c1a7": 3031, // ZTE
		"78c3e9": 152,  // Samsung
		"78c5f8": 3334, // Huawei
		"78c881": 101,  // Sony
		"78ca04": 457,  // Nokia
		"78ca39": 545,  // Apple
		"78ca5e": 7895, // Elno
		"78cb33": 7896, // DHC
		"78cc2b": 7897, // Sinewy
		"78cd8e": 741,  // SMC
		"78cf2f": 3334, // Huawei
		"78cff9": 3334, // Huawei
		"78d004": 7898, // Neousys
		"78d129": 7899, // Vicos
		"78d162": 545,  // Apple
		"78d294": 1368, // Netgear
		"78d347": 306,  // Ericsson
		"78d34f": 7900, // Pace-O-Matic
		"78d3ed": 7901, // Norma
		"78d6b2": 35,   // Toshiba
		"78d6dc": 687,  // Motorola
		"78d6f0": 152,  // Samsung
		"78d71a": 4564, // Ciena
		"78d752": 3334, // Huawei
		"78d75f": 545,  // Apple
		"78da6e": 2,    // Cisco
		"78daa2": 7902, // Cynosure
		"78dab3": 7903, // GBO
		"78dd12": 2639, // Arcadyan
		"78dd33": 3334, // Huawei
		"78ddd6": 7904, // c-scape
		"78e103": 5438, // Amazon
		"78e22c": 3334, // Huawei
		"78e36d": 6376, // Espressif
		"78e3b5": 302,  // HP
		"78e3de": 545,  // Apple
		"78e7d1": 302,  // HP
		"78e8b6": 3031, // ZTE
		"78e980": 7905, // RainUs
		"78eb46": 3334, // Huawei
		"78ec74": 7906, // Kyland-USA
		"78ef4c": 7907, // Unetconvergence
		"78f09b": 3334, // Huawei
		"78f238": 152,  // Samsung
		"78f29e": 5439, // Pegatron
		"78f557": 3334, // Huawei
		"78f5fd": 3334, // Huawei
		"78f7be": 152,  // Samsung
		"78f7d0": 7908, // Silverbrook
		"78f882": 869,  // LG
		"78f944": 67,   // Private
		"78fbd8": 545,  // Apple
		"78fd94": 545,  // Apple
		"78fe3d": 826,  // Juniper
		"78fe41": 7909, // Socus
		"78ff57": 421,  // Intel
		"78ffca": 6265, // Tecno
		"7c004d": 3334, // Huawei
		"7c0191": 545,  // Apple
		"7c02bc": 7910, // Hansung
		"7c034c": 2048, // Sagemcom
		"7c035e": 5697, // Xiaomi
		"7c03ab": 5697, // Xiaomi
		"7c03d8": 2048, // Sagemcom
		"7c04d0": 545,  // Apple
		"7c0507": 5439, // Pegatron
		"7c051e": 7911, // Rafael
		"7c0623": 3429, // Ultra
		"7c092b": 7912, // Bekey
		"7c0a3f": 152,  // Samsung
		"7c0a50": 7913, // J-MEX
		"7c0bc6": 152,  // Samsung
		"7c0ece": 2,    // Cisco
		"7c10c9": 1806, // ASUS
		"7c11be": 545,  // Apple
		"7c11cb": 3334, // Huawei
		"7c11cd": 7914, // QianTang
		"7c131d": 6255, // Sernet
		"7c1689": 2048, // Sagemcom
		"7c18cd": 7915, // E-TRON
		"7c1a03": 7916, // 8Locations
		"7c1ac0": 3334, // Huawei
		"7c1b93": 3334, // Huawei
		"7c1c4e": 869,  // LG
		"7c1c68": 152,  // Samsung
		"7c1cf1": 3334, // Huawei
		"7c1dd9": 5697, // Xiaomi
		"7c1e52": 612,  // Microsoft
		"7c2048": 7917, // KoamTac
		"7c210d": 2,    // Cisco
		"7c210e": 2,    // Cisco
		"7c214a": 421,  // Intel
		"7c2302": 152,  // Samsung
		"7c240c": 7075, // Telechips
		"7c2499": 545,  // Apple
		"7c2586": 826,  // Juniper
		"7c2587": 7918, // chaowifi.com
		"7c25da": 6440, // FN-Link
		"7c2634": 125,  // ARRIS
		"7c2664": 2048, // Sagemcom
		"7c2a31": 421,  // Intel
		"7c2adb": 5697, // Xiaomi
		"7c2ebd": 3521, // Google
		"7c2edd": 152,  // Samsung
		"7c2f80": 4298, // Gigaset
		"7c310e": 2,    // Cisco
		"7c336e": 7919, // MEG
		"7c33f9": 3334, // Huawei
		"7c38ad": 152,  // Samsung
		"7c3953": 3031, // ZTE
		"7c3985": 3334, // Huawei
		"7c3d2b": 3334, // Huawei
		"7c3e74": 3334, // Huawei
		"7c3e9d": 7920, // Patech
		"7c41a2": 457,  // Nokia
		"7c438f": 7921, // E-Band
		"7c444c": 7922, // Entertainment
		"7c4685": 687,  // Motorola
		"7c49eb": 6280, // XIAOMI
		"7c4a82": 7923, // Portsmith
		"7c4ca5": 3512, // BSkyB
		"7c4f7d": 7924, // Sawwave
		"7c4fb5": 2639, // Arcadyan
		"7c5049": 545,  // Apple
		"7c5079": 421,  // Intel
		"7c50da": 67,   // Private
		"7c534a": 7925, // Metamako
		"7c55a7": 7926, // Kastle
		"7c55e7": 7927, // YSI
		"7c573c": 1685, // Aruba
		"7c574e": 7928, // CoBI
		"7c5a1c": 3578, // Sophos
		"7c5a67": 7929, // JNC
		"7c5cf8": 421,  // Intel
		"7c604a": 7930, // Avelon
		"7c6097": 3334, // Huawei
		"7c6166": 5438, // Amazon
		"7c6193": 1341, // HTC
		"7c6456": 152,  // Samsung
		"7c67a2": 421,  // Intel
		"7c696b": 7931, // Atmosic
		"7c69f6": 2,    // Cisco
		"7c6ab3": 7932, // IBC
		"7c6ac3": 7933, // GatesAir
		"7c6adb": 7934, // SafeTone
		"7c6b33": 7935, // Tenyu
		"7c6bf7": 7936, // NTI
		"7c6d62": 545,  // Apple
		"7c6df8": 545,  // Apple
		"7c70db": 421,  // Intel
		"7c726e": 306,  // Ericsson
		"7c72e4": 7937, // Unikey
		"7c73eb": 3334, // Huawei
		"7c7635": 421,  // Intel
		"7c7668": 3334, // Huawei
		"7c7673": 7938, // ENMAS
		"7c7716": 2692, // ZyXEL
		"7c787e": 152,  // Samsung
		"7c78b2": 6881, // Wyze
		"7c79e8": 7939, // PayRange
		"7c7a53": 7940, // Phytrex
		"7c7a91": 421,  // Intel
		"7c7d3d": 3334, // Huawei
		"7c7d41": 7941, // Jinmuyu
		"7c822d": 7942, // Nortec
		"7c8530": 457,  // Nokia
		"7c87ce": 6376, // Espressif
		"7c8956": 152,  // Samsung
		"7c8ac0": 7943, // EVBox
		"7c8ae1": 358,  // Compal
		"7c8bb5": 152,  // Samsung
		"7c8bca": 1595, // TP-Link
		"7c8fde": 7446, // DWnet
		"7c9122": 152,  // Samsung
		"7c942a": 3334, // Huawei
		"7c95b1": 185,  // Extreme
		"7c95f3": 2,    // Cisco
		"7c96d2": 6298, // Fihonest
		"7c97e1": 3334, // Huawei
		"7c9a1d": 545,  // Apple
		"7c9a54": 6392, // Technicolor
		"7c9ebd": 6376, // Espressif
		"7ca177": 3334, // Huawei
		"7ca1ae": 545,  // Apple
		"7ca23e": 3334, // Huawei
		"7ca29b": 7944, // D.SignT
		"7ca61d": 7945, // MHL
		"7ca62a": 302,  // HP
		"7ca96b": 7022, // Syrotech
		"7ca97d": 7946, // Objenious
		"7cab25": 7947, // Mesmo
		"7cab60": 545,  // Apple
		"7cad4f": 2,    // Cisco
		"7cad74": 2,    // Cisco
		"7cb03e": 7948, // OSRAM
		"7cb0c2": 421,  // Intel
		"7cb15d": 3334, // Huawei
		"7cb177": 7949, // Satelco
		"7cb25c": 5767, // Acacia
		"7cb27d": 421,  // Intel
		"7cb542": 7950, // ACES
		"7cb566": 421,  // Intel
		"7cb59b": 1595, // TP-Link
		"7cb733": 2544, // Askey
		"7cb77b": 2708, // Paradigm
		"7cbb6f": 7951, // Cosco
		"7cbb8a": 1427, // Nintendo
		"7cbf88": 7952, // Mobilicom
		"7cbfb1": 125,  // ARRIS
		"7cc180": 545,  // Apple
		"7cc225": 152,  // Samsung
		"7cc2c6": 1595, // TP-Link
		"7cc385": 3334, // Huawei
		"7cc3a1": 545,  // Apple
		"7cc4ef": 7953, // Devialet
		"7cc537": 545,  // Apple
		"7cc77e": 4921, // Fiberhome
		"7cc8d7": 7954, // Damalisk
		"7cc95a": 102,  // Dell
		"7ccb0d": 7955, // Antaira
		"7cccb8": 421,  // Intel
		"7ccd11": 7956, // MS-Magnet
		"7cd1c3": 545,  // Apple
		"7cd30a": 3978, // Inventec
		"7cd566": 5438, // Amazon
		"7cd661": 5697, // Xiaomi
		"7cd762": 7957, // Freestyle
		"7cd844": 7958, // Enmotus
		"7cd95c": 3521, // Google
		"7cd9a0": 3334, // Huawei
		"7cda84": 7959, // Dongnian
		"7cdb98": 2544, // Askey
		"7cdde9": 7960, // ATOM
		"7cdfa1": 6376, // Espressif
		"7ce044": 7961, // NEON
		"7ce2ca": 826,  // Juniper
		"7ce4aa": 67,   // Private
		"7ce524": 7962, // Quirky
		"7ce97c": 6281, // Itel
		"7ceb7f": 7963, // Dmet
		"7cebea": 7964, // Asct
		"7ced8d": 612,  // Microsoft
		"7cef8a": 7965, // Inhon
		"7cf05f": 545,  // Apple
		"7cf2dd": 7966, // Vence
		"7cf31b": 869,  // LG
		"7cf429": 7967, // NUUO
		"7cf854": 152,  // Samsung
		"7cf880": 2,    // Cisco
		"7cf90e": 152,  // Samsung
		"7cf9a0": 4921, // Fiberhome
		"7cfadf": 545,  // Apple
		"7cfc16": 545,  // Apple
		"7cfc3c": 1399, // Visteon
		"7cfd6b": 5697, // Xiaomi
		"7cfe28": 7968, // Salutron
		"7cfe90": 432,  // Mellanox
		"80000b": 421,  // Intel
		"800010": 1057, // AT&T
		"80006e": 545,  // Apple
		"80015c": 7969, // Synaptics
		"800184": 1341, // HTC
		"80029c": 1450, // Gemtek
		"8002df": 7970, // ORA
		"800384": 2737, // Ruckus
		"80045f": 545,  // Apple
		"800588": 5440, // Ruijie
		"8007a2": 7971, // Esson
		"800902": 3599, // Keysight
		"800a06": 7972, // CoMTEC
		"800c67": 545,  // Apple
		"800cf9": 5438, // Amazon
		"800dd7": 7973, // Latticework
		"800e24": 7974, // ForgetBox
		"801382": 3334, // Huawei
		"80177d": 76,   // Nortel
		"801844": 102,  // Dell
		"8018a7": 152,  // Samsung
		"801934": 421,  // Intel
		"8019fe": 7975, // JianLing
		"801daa": 620,  // Avaya
		"801f02": 115,  // Edimax
		"801f12": 706,  // Microchip
		"8020da": 2048, // Sagemcom
		"8020fd": 152,  // Samsung
		"80248f": 2,    // Cisco
		"802511": 6281, // Itel
		"802689": 803,  // D-Link
		"802994": 6392, // Technicolor
		"802aa8": 2978, // Ubiquiti
		"802afa": 7976, // Germaneers
		"802dbf": 2,    // Cisco
		"802de1": 7977, // Solarbridge
		"802e14": 7978, // azeti
		"803049": 2873, // Liteon
		"8030e0": 302,  // HP
		"8031f0": 152,  // Samsung
		"803253": 421,  // Intel
		"803428": 706,  // Microchip
		"803457": 7979, // OT
		"8035c1": 5697, // Xiaomi
		"803773": 1368, // Netgear
		"803896": 3205, // Sharp
		"8038bc": 3334, // Huawei
		"8038fb": 421,  // Intel
		"8038fd": 7980, // LeapFrog
		"8039e5": 7981, // Patlite
		"803a59": 1057, // AT&T
		"803af4": 4921, // Fiberhome
		"803b9a": 7982, // ghe-ces
		"803f5d": 7983, // Winstars
		"804126": 3334, // Huawei
		"8045dd": 421,  // Intel
		"804786": 152,  // Samsung
		"804971": 545,  // Apple
		"804a14": 545,  // Apple
		"804b50": 1655, // Silicon
		"804e70": 152,  // Samsung
		"804e81": 152,  // Samsung
		"804f58": 7984, // ThinkEco
		"80501b": 457,  // Nokia
		"8050f6": 6281, // Itel
		"8054d9": 3334, // Huawei
		"805719": 152,  // Samsung
		"8058f8": 687,  // Motorola
		"8059fd": 7985, // Noviga
		"805a04": 869,  // LG
		"805e4f": 6440, // FN-Link
		"805fc5": 545,  // Apple
		"806007": 4555, // RIM
		"806459": 7986, // Nimbus
		"80656d": 152,  // Samsung
		"80657c": 545,  // Apple
		"8065e9": 7987, // BenQ
		"806629": 7988, // Prescope
		"80691a": 2468, // Belkin
		"806933": 3334, // Huawei
		"806940": 7989, // Lexar
		"806a00": 2,    // Cisco
		"806c1b": 687,  // Motorola
		"806d71": 5438, // Amazon
		"806d97": 67,   // Private
		"806f1c": 3334, // Huawei
		"80711f": 826,  // Juniper
		"80717a": 3334, // Huawei
		"807215": 3512, // BSkyB
		"807264": 3334, // Huawei
		"80739f": 2848, // Kyocera
		"807459": 7990, // K's
		"80751f": 3512, // BSkyB
		"807693": 7991, // Newag
		"8077a4": 6265, // Tecno
		"807871": 2544, // Askey
		"80795d": 6295, // Infinix
		"807abf": 1341, // HTC
		"807b3e": 152,  // Samsung
		"807d14": 3334, // Huawei
		"807d1b": 7992, // Neosystem
		"807d3a": 6376, // Espressif
		"807ff8": 826,  // Juniper
		"808223": 545,  // Apple
		"808287": 7993, // ATCOM
		"8084a9": 7994, // oshkosh
		"808698": 7995, // Netronics
		"8086d9": 152,  // Samsung
		"8086f2": 421,  // Intel
		"808917": 1595, // TP-Link
		"808a8b": 6369, // Vivo
		"808abd": 152,  // Samsung
		"808af7": 7996, // Nanoleaf
		"808c97": 1269, // Kaonmedia
		"808db7": 302,  // HP
		"808f1d": 1595, // TP-Link
		"808fe8": 3541, // Intelbras
		"809133": 3004, // Azurewave
		"8091c0": 7997, // AgileMesh
		"80929f": 545,  // Apple
		"809393": 7998, // Xapt
		"809621": 2664, // Lenovo
		"8096b1": 125,  // ARRIS
		"809b20": 421,  // Intel
		"809fab": 4921, // Fiberhome
		"809ff5": 152,  // Samsung
		"80a1ab": 7999, // Intellisis
		"80a235": 6294, // Edgecore
		"80a589": 3004, // Azurewave
		"80a796": 8000, // Neurotek
		"80aaa4": 8001, // Usag
		"80acac": 826,  // Juniper
		"80ad16": 5697, // Xiaomi
		"80ad67": 2134, // Kasda
		"80b03d": 545,  // Apple
		"80b07b": 3031, // ZTE
		"80b234": 6392, // Technicolor
		"80b289": 8002, // Forworld
		"80b575": 3334, // Huawei
		"80b624": 8003, // IVS
		"80b655": 421,  // Intel
		"80b686": 3334, // Huawei
		"80b709": 8004, // Viptela
		"80b95c": 8005, // ELFTECH
		"80b97a": 5820, // eero
		"80baac": 8006, // TeleAdapt
		"80bae6": 8007, // Neets
		"80bbeb": 8008, // Satmap
		"80be05": 545,  // Apple
		"80c16e": 302,  // HP
		"80c3ba": 3117, // Sennheiser
		"80c5e6": 612,  // Microsoft
		"80c5f2": 3004, // Azurewave
		"80c6ab": 6392, // Technicolor
		"80c755": 2153, // Panasonic
		"80c7c5": 4921, // Fiberhome
		"80c862": 8009, // Openpeak
		"80cc12": 3334, // Huawei
		"80cc9c": 1368, // Netgear
		"80ce62": 302,  // HP
		"80ceb9": 152,  // Samsung
		"80cf41": 2664, // Lenovo
		"80cfa2": 3334, // Huawei
		"80d019": 8010, // Embed
		"80d04a": 6392, // Technicolor
		"80d065": 8011, // CKS
		"80d09b": 3334, // Huawei
		"80d21d": 3004, // Azurewave
		"80d2e5": 1427, // Nintendo
		"80d336": 6313, // Cern
		"80d433": 8012, // LzLabs
		"80d4a5": 3334, // Huawei
		"80d605": 545,  // Apple
		"80da13": 5820, // eero
		"80dabc": 6710, // Megafone
		"80e01d": 2,    // Cisco
		"80e1bf": 3334, // Huawei
		"80e540": 125,  // ARRIS
		"80e650": 545,  // Apple
		"80e82c": 302,  // HP
		"80e86f": 2,    // Cisco
		"80ea07": 1595, // TP-Link
		"80ea23": 1592, // Wistron
		"80ea96": 545,  // Apple
		"80eb77": 1592, // Wistron
		"80ed2c": 545,  // Apple
		"80ee73": 4937, // Shuttle
		"80f1f1": 8013, // Tech4home
		"80f25e": 8014, // Kyynel
		"80f3ef": 7234, // Facebook
		"80f503": 125,  // ARRIS
		"80f8eb": 8015, // RayTight
		"80fa5b": 5687, // Clevo
		"80fb06": 3334, // Huawei
		"80fd7a": 7460, // BLU
		"80ffa8": 8016, // Unidis
		"84002d": 5439, // Pegatron
		"8400d2": 101,  // Sony
		"840112": 1269, // Kaonmedia
		"840283": 531,  // HUMAX
		"840328": 826,  // Juniper
		"8406fa": 4921, // Fiberhome
		"840b2d": 152,  // Samsung
		"840b7c": 870,  // Hitron
		"840d8e": 6376, // Espressif
		"84100d": 687,  // Motorola
		"84119e": 152,  // Samsung
		"84139f": 3031, // ZTE
		"84144d": 421,  // Intel
		"8415d3": 3334, // Huawei
		"84160c": 858,  // Broadcom
		"8416f9": 1595, // TP-Link
		"841715": 8017, // GP
		"8417ef": 6392, // Technicolor
		"841826": 8018, // Osram
		"84183a": 2737, // Ruckus
		"841888": 826,  // Juniper
		"841b5e": 1368, // Netgear
		"841b77": 421,  // Intel
		"841c70": 3031, // ZTE
		"841e26": 8019, // KERNEL-I
		"841ea3": 2048, // Sagemcom
		"8421f1": 3334, // Huawei
		"842388": 2737, // Ruckus
		"84248d": 768,  // Zebra
		"842519": 152,  // Samsung
		"84253f": 8020, // silex
		"8425a4": 8021, // Tariox
		"8425db": 152,  // Samsung
		"84262b": 457,  // Nokia
		"84285a": 8022, // Saffron
		"842999": 545,  // Apple
		"842afd": 302,  // HP
		"842b2b": 102,  // Dell
		"842b50": 8023, // Huria
		"842bbc": 8024, // Modelleisenbahn
		"842e14": 1655, // Silicon
		"842e27": 152,  // Samsung
		"8430e5": 8025, // SkyHawke
		"843497": 302,  // HP
		"8437d5": 152,  // Samsung
		"843835": 545,  // Apple
		"843838": 152,  // Samsung
		"843a4b": 421,  // Intel
		"843a5b": 3978, // Inventec
		"843b10": 8026, // Lvswitches
		"843dc6": 2,    // Cisco
		"843e92": 3334, // Huawei
		"843f4e": 8027, // Tri-Tech
		"844076": 8028, // Drivenets
		"844167": 545,  // Apple
		"844464": 8029, // ServerU
		"8446fe": 3334, // Huawei
		"844765": 3334, // Huawei
		"844823": 8030, // WOXTER
		"844915": 8031, // vArmour
		"844f03": 8032, // Ablelink
		"845181": 152,  // Samsung
		"8454df": 3334, // Huawei
		"8455a5": 152,  // Samsung
		"845733": 612,  // Microsoft
		"845a81": 8033, // ffly4u
		"845b12": 3334, // Huawei
		"845c93": 8034, // Chabrier
		"845cf3": 421,  // Intel
		"845f04": 152,  // Samsung
		"846082": 67,   // Private
		"8461a0": 125,  // ARRIS
		"8462a6": 3101, // EuroCB
		"8463d6": 612,  // Microsoft
		"84683e": 421,  // Intel
		"846878": 545,  // Apple
		"846991": 457,  // Nokia
		"847127": 1655, // Silicon
		"84716a": 3334, // Huawei
		"847207": 8035, // I&C
		"84742a": 3031, // ZTE
		"847460": 3031, // ZTE
		"847637": 3334, // Huawei
		"847778": 8036, // Cochlear
		"84788b": 545,  // Apple
		"8478ac": 2,    // Cisco
		"847933": 8037, // profichip
		"847a88": 1341, // HTC
		"847ab6": 7031, // AltoBeam
		"847b57": 421,  // Intel
		"847beb": 102,  // Dell
		"84802d": 2,    // Cisco
		"848094": 8038, // Meter
		"848102": 4921, // Fiberhome
		"848336": 8039, // Newrun
		"848371": 620,  // Avaya
		"848433": 8040, // Paradox
		"848506": 545,  // Apple
		"8486f3": 8041, // Greenvity
		"8489ad": 545,  // Apple
		"848a8d": 2,    // Cisco
		"848c8d": 545,  // Apple
		"848d84": 8042, // Rajant
		"848e0c": 545,  // Apple
		"848e96": 8043, // Embertec
		"848edf": 101,  // Sony
		"848f69": 102,  // Dell
		"84930c": 8044, // Incoax
		"8493a0": 3334, // Huawei
		"84948c": 870,  // Hitron
		"849681": 8045, // Cathay
		"8496d8": 125,  // ARRIS
		"8497b8": 8046, // Memjet
		"849866": 152,  // Samsung
		"849ca6": 2639, // Arcadyan
		"849d64": 741,  // SMC
		"849fb5": 3334, // Huawei
		"84a06e": 2048, // Sagemcom
		"84a134": 545,  // Apple
		"84a1d1": 2048, // Sagemcom
		"84a3b5": 8047, // Propulsion
		"84a423": 2048, // Sagemcom
		"84a466": 152,  // Samsung
		"84a6c8": 421,  // Intel
		"84a788": 8048, // Perples
		"84a8e4": 3334, // Huawei
		"84a938": 4919, // LCFC
		"84a93e": 302,  // HP
		"84a9c4": 3334, // Huawei
		"84aa9c": 6428, // MitraStar
		"84ab1a": 545,  // Apple
		"84ab26": 6539, // Tiinlab
		"84ac16": 545,  // Apple
		"84ad58": 3334, // Huawei
		"84ad8d": 545,  // Apple
		"84afec": 1077, // Buffalo
		"84b153": 545,  // Apple
		"84b261": 2,    // Cisco
		"84b31b": 8049, // Kinexon
		"84b4db": 1655, // Silicon
		"84b517": 2,    // Cisco
		"84b541": 152,  // Samsung
		"84b59c": 826,  // Juniper
		"84b802": 2,    // Cisco
		"84b8b8": 687,  // Motorola
		"84ba20": 1655, // Silicon
		"84ba3b": 87,   // Canon
		"84bb69": 125,  // ARRIS
		"84be52": 3334, // Huawei
		"84c0ef": 152,  // Samsung
		"84c1c1": 826,  // Juniper
		"84c3e8": 8050, // Vaillant
		"84c5a6": 421,  // Intel
		"84c727": 8051, // Gnodal
		"84c78f": 8052, // APS
		"84c7ea": 101,  // Sony
		"84c8b1": 8053, // Incognito
		"84c9b2": 803,  // D-Link
		"84cc63": 3334, // Huawei
		"84cca8": 6376, // Espressif
		"84cfbf": 8054, // Fairphone
		"84d15a": 6433, // TCT
		"84d343": 374,  // Calix
		"84d3d5": 3334, // Huawei
		"84d47e": 1685, // Aruba
		"84d4c8": 7118, // Widex
		"84d608": 4030, // Wingtech
		"84d6c5": 4906, // SolarEdge
		"84d6d0": 5438, // Amazon
		"84d81b": 1595, // TP-Link
		"84d9c8": 8055, // Unipattern
		"84dbac": 3334, // Huawei
		"84dbfc": 457,  // Nokia
		"84ddb7": 8056, // Cilag
		"84df0c": 8057, // NET2GRID
		"84e058": 125,  // ARRIS
		"84e327": 806,  // Tailyn
		"84e629": 8058, // Bluwan
		"84e892": 2246, // Actiontec
		"84e986": 3334, // Huawei
		"84ea99": 8059, // Vieworks
		"84eaed": 1916, // Roku
		"84ebef": 2,    // Cisco
		"84ed33": 8060, // BBMC
		"84ef18": 421,  // Intel
		"84f129": 8061, // Metrascale
		"84f147": 2,    // Cisco
		"84f3eb": 6376, // Espressif
		"84f6fa": 8062, // Miovision
		"84f703": 6376, // Espressif
		"84f883": 8063, // Luminar
		"84fcac": 545,  // Apple
		"84fcfe": 545,  // Apple
		"84fd27": 1655, // Silicon
		"84fdd1": 421,  // Intel
		"84fe9e": 8064, // RTC
		"880118": 8065, // BLT
		"880355": 2639, // Arcadyan
		"88074b": 869,  // LG
		"880905": 8066, // MTMCommunications
		"8809af": 6362, // Masimo
		"881036": 8067, // Panodic
		"88108f": 3334, // Huawei
		"881196": 3334, // Huawei
		"88124e": 5787, // Qualcomm
		"8815c5": 3334, // Huawei
		"8818ae": 8068, // Tamron
		"881908": 545,  // Apple
		"881c95": 6281, // Itel
		"881dfc": 2,    // Cisco
		"881fa1": 545,  // Apple
		"882012": 8069, // LMI
		"8821e3": 8070, // Nebusens
		"88238c": 4921, // Fiberhome
		"882510": 1685, // Aruba
		"88252c": 2639, // Arcadyan
		"882593": 1595, // TP-Link
		"8828b3": 3334, // Huawei
		"882949": 5695, // Renesas
		"882950": 8071, // Netmoon
		"88299c": 152,  // Samsung
		"882bd7": 8072, // Addenergie
		"882e5a": 8073, // storONE
		"88308a": 2056, // Murata
		"88329b": 152,  // Samsung
		"8833be": 8074, // Ivenix
		"88354c": 8075, // Transics
		"883612": 8076, // SRC
		"88365f": 869,  // LG
		"88366c": 1252, // EFM
		"8836cf": 3334, // Huawei
		"883a30": 1685, // Aruba
		"883c1c": 4253, // Mercury
		"883d24": 3521, // Google
		"883f99": 300,  // Siemens
		"883fd3": 3334, // Huawei
		"884033": 3334, // Huawei
		"88403b": 3334, // Huawei
		"884067": 3976, // Infomark
		"8843e1": 2,    // Cisco
		"884477": 3334, // Huawei
		"8844f6": 457,  // Nokia
		"884604": 5697, // Xiaomi
		"88462a": 7075, // Telechips
		"884a18": 8077, // Opulinks
		"884a70": 7186, // Wacom
		"884b39": 300,  // Siemens
		"884ccf": 8078, // Pulzze
		"884d7c": 545,  // Apple
		"885046": 3535, // Lear
		"8851fb": 302,  // HP
		"8852eb": 5697, // Xiaomi
		"88532e": 421,  // Intel
		"885395": 545,  // Apple
		"8853d4": 3334, // Huawei
		"88541f": 3521, // Google
		"88571d": 7148, // Seongji
		"88576d": 8079, // XTA
		"8857ee": 1077, // Buffalo
		"885a85": 1592, // Wistron
		"885a92": 2,    // Cisco
		"885bdd": 185,  // Extreme
		"885dfb": 3031, // ZTE
		"8863df": 545,  // Apple
		"886440": 545,  // Apple
		"886639": 3334, // Huawei
		"88665a": 545,  // Apple
		"8866a5": 545,  // Apple
		"88693d": 3334, // Huawei
		"886ab1": 6369, // Vivo
		"886ae3": 2236, // Alpha
		"886b0f": 1108, // Bluegiga
		"886b44": 5435, // Sunnovo
		"886b6e": 545,  // Apple
		"886f29": 8080, // Pocketbook
		"886fd4": 102,  // Dell
		"88708c": 2664, // Lenovo
		"8871b1": 125,  // ARRIS
		"8871e5": 5438, // Amazon
		"887384": 35,   // Toshiba
		"887556": 2,    // Cisco
		"887598": 152,  // Samsung
		"887873": 421,  // Intel
		"88789c": 8081, // Game
		"88795b": 3532, // Konka
		"88797e": 687,  // Motorola
		"887a31": 8082, // Velankani
		"887e25": 185,  // Extreme
		"8881b9": 3334, // Huawei
		"888322": 152,  // Samsung
		"88835d": 6440, // FN-Link
		"888603": 3334, // Huawei
		"8886a0": 8083, // Simton
		"8886c2": 8084, // STABILO
		"888717": 87,   // Canon
		"8887dd": 8085, // DarbeeVision
		"88892f": 3334, // Huawei
		"888964": 8086, // GSI
		"888e68": 3334, // Huawei
		"888e7f": 5420, // Atop
		"889009": 826,  // Juniper
		"88908d": 2,    // Cisco
		"889166": 8087, // Viewcooper
		"8891dd": 8088, // Racktivity
		"889471": 90,   // Brocade
		"88947e": 4921, // Fiberhome
		"8894f9": 8089, // Gemicom
		"88964e": 125,  // ARRIS
		"889655": 8090, // Zitte
		"889765": 8091, // exands
		"889821": 8092, // Teraon
		"889b39": 152,  // Samsung
		"889d98": 8093, // Allied-telesisK.K
		"889e33": 6433, // TCT
		"889e68": 6392, // Technicolor
		"889f6f": 152,  // Samsung
		"88a0be": 3334, // Huawei
		"88a25e": 826,  // Juniper
		"88a2d7": 3334, // Huawei
		"88a303": 152,  // Samsung
		"88a479": 545,  // Apple
		"88a4c2": 4919, // LCFC
		"88a5bd": 8094, // Qpcom
		"88a6c6": 2048, // Sagemcom
		"88a73c": 6839, // Ragentek
		"88a9b7": 545,  // Apple
		"88acc0": 2692, // ZyXEL
		"88acc1": 8095, // Generiton
		"88ad43": 5439, // Pegatron
		"88add2": 152,  // Samsung
		"88ae07": 545,  // Apple
		"88ae1d": 358,  // Compal
		"88aedd": 6672, // EliteGroup
		"88b111": 421,  // Intel
		"88b1e1": 2485, // Mojo
		"88b291": 545,  // Apple
		"88b436": 67,   // Private
		"88b4a6": 687,  // Motorola
		"88b627": 8096, // Gembird
		"88b66b": 8097, // easynetworks
		"88b6ee": 1238, // Dish
		"88b945": 545,  // Apple
		"88ba7f": 8098, // Qfiednet
		"88bcc1": 3334, // Huawei
		"88bd45": 152,  // Samsung
		"88bd78": 2668, // Flaircomm
		"88bfe4": 3334, // Huawei
		"88c08b": 545,  // Apple
		"88c227": 3334, // Huawei
		"88c242": 8099, // Poynt
		"88c3b3": 8100, // Sovico
		"88c626": 4079, // Logitech
		"88c663": 545,  // Apple
		"88c9d0": 869,  // LG
		"88cb87": 545,  // Apple
		"88cefa": 3334, // Huawei
		"88cf98": 3334, // Huawei
		"88d039": 6894, // Tonly
		"88d199": 3724, // Vencer
		"88d274": 3031, // ZTE
		"88d37b": 8101, // FirmTek
		"88d5a8": 6281, // Itel
		"88d652": 8102, // AMERGINT
		"88d7bc": 7625, // DEP
		"88d7f6": 1806, // ASUS
		"88d82e": 421,  // Intel
		"88d98f": 826,  // Juniper
		"88dc96": 8103, // EnGenius
		"88dd79": 1304, // Voltaire
		"88de7c": 2544, // Askey
		"88dea9": 1916, // Roku
		"88e034": 3906, // Shinwa
		"88e056": 3334, // Huawei
		"88e0f3": 826,  // Juniper
		"88e3ab": 3334, // Huawei
		"88e603": 8104, // Avotek
		"88e64b": 826,  // Juniper
		"88e712": 8105, // Whirlpool
		"88e87f": 545,  // Apple
		"88e90f": 8106, // innomdlelab
		"88e917": 8107, // Tamaggo
		"88e9a4": 302,  // HP
		"88e9fe": 545,  // Apple
		"88ed1c": 8108, // Cudo
		"88ef16": 125,  // ARRIS
		"88f031": 2,    // Cisco
		"88f077": 2,    // Cisco
		"88f488": 8109, // cellon
		"88f490": 8110, // Jetmobile
		"88f56e": 3334, // Huawei
		"88f7bf": 6369, // Vivo
		"88f7c7": 6392, // Technicolor
		"88f872": 3334, // Huawei
		"88fca6": 1637, // devolo
		"88fd15": 8111, // Lineeye
		"8c006d": 545,  // Apple
		"8c02fa": 8112, // CoMMANDO
		"8c04ba": 102,  // Dell
		"8c04ff": 6392, // Technicolor
		"8c0551": 8113, // Koubachi
		"8c09f4": 125,  // ARRIS
		"8c0c87": 457,  // Nokia
		"8c0c90": 2737, // Ruckus
		"8c0ca3": 8114, // Amper
		"8c0d76": 3334, // Huawei
		"8c0f6f": 5439, // Pegatron
		"8c0fa0": 8115, // di-soric
		"8c0fc9": 3334, // Huawei
		"8c0ffa": 8116, // Hutec
		"8c10d4": 2048, // Sagemcom
		"8c14b4": 3031, // ZTE
		"8c15c7": 3334, // Huawei
		"8c1645": 4919, // LCFC
		"8c1759": 421,  // Intel
		"8c17b6": 3334, // Huawei
		"8c19b5": 2639, // Arcadyan
		"8c1abf": 152,  // Samsung
		"8c1d96": 421,  // Intel
		"8c210a": 1595, // TP-Link
		"8c2505": 3334, // Huawei
		"8c271d": 8117, // QuantHouse
		"8c278a": 2815, // Vocollect
		"8c2937": 545,  // Apple
		"8c2daa": 545,  // Apple
		"8c31e2": 8118, // Dayouplus
		"8c3330": 8119, // EmFirst
		"8c3446": 3334, // Huawei
		"8c34fd": 3334, // Huawei
		"8c3a7e": 3857, // Universal
		"8c3ae3": 869,  // LG
		"8c3bad": 1368, // Netgear
		"8c3c07": 8120, // Skiva
		"8c3c4a": 6082, // NAKAYO
		"8c41f2": 8121, // RDA
		"8c41f4": 8122, // IPmotion
		"8c426d": 3334, // Huawei
		"8c444f": 531,  // HUMAX
		"8c4500": 2056, // Murata
		"8c477f": 8123, // NambooSolution
		"8c47be": 102,  // Dell
		"8c4962": 1916, // Roku
		"8c497a": 185,  // Extreme
		"8c49b6": 6369, // Vivo
		"8c4b14": 6376, // Espressif
		"8c4cad": 8124, // Evoluzn
		"8c4cdc": 4478, // Planex
		"8c4db9": 8125, // Unmonday
		"8c4dea": 8126, // Cerio
		"8c53f7": 1331, // A&D
		"8c541d": 7468, // LGE
		"8c554a": 421,  // Intel
		"8c5646": 869,  // LG
		"8c56c5": 1427, // Nintendo
		"8c579b": 1592, // Wistron
		"8c5877": 545,  // Apple
		"8c5973": 2692, // ZyXEL
		"8c598b": 8127, // C
		"8c59c3": 6491, // ADB
		"8c5a25": 125,  // ARRIS
		"8c5ac1": 3334, // Huawei
		"8c5bf0": 125,  // ARRIS
		"8c5ca1": 8128, // d-broad
		"8c5d60": 8129, // UCI
		"8c5ebd": 3334, // Huawei
		"8c5fad": 4921, // Fiberhome
		"8c604f": 2,    // Cisco
		"8c6078": 8130, // Swissbit
		"8c60e7": 8131, // Mpgio
		"8c61a3": 125,  // ARRIS
		"8c6422": 101,  // Sony
		"8c64a2": 6923, // OnePlus
		"8c6794": 6369, // Vivo
		"8c683a": 3334, // Huawei
		"8c6878": 8132, // Nortek-AS
		"8c68c8": 3031, // ZTE
		"8c6a8d": 6392, // Technicolor
		"8c6ae4": 8133, // Viogem
		"8c6d77": 3334, // Huawei
		"8c705a": 421,  // Intel
		"8c71f8": 152,  // Samsung
		"8c736e": 4,    // Fujitsu
		"8c73a0": 4921, // Fiberhome
		"8c76c1": 8134, // Goden
		"8c7712": 152,  // Samsung
		"8c7967": 3031, // ZTE
		"8c79f5": 152,  // Samsung
		"8c7a15": 2737, // Ruckus
		"8c7a3d": 5697, // Xiaomi
		"8c7aaa": 545,  // Apple
		"8c7b9d": 545,  // Apple
		"8c7c92": 545,  // Apple
		"8c7cff": 90,   // Brocade
		"8c7eb3": 8135, // Lytro
		"8c7f3b": 125,  // ARRIS
		"8c8126": 8136, // Arcom
		"8c82a8": 7260, // Insigma
		"8c83df": 457,  // Nokia
		"8c83e1": 152,  // Samsung
		"8c83e8": 3334, // Huawei
		"8c8401": 67,   // Private
		"8c8590": 545,  // Apple
		"8c85c1": 1685, // Aruba
		"8c85e6": 8137, // Cleondris
		"8c861e": 545,  // Apple
		"8c897a": 8138, // Augtek
		"8c89a5": 1812, // Micro-Star
		"8c8caa": 4919, // LCFC
		"8c8d28": 421,  // Intel
		"8c8e76": 8139, // taskit
		"8c8ef2": 545,  // Apple
		"8c8fe9": 545,  // Apple
		"8c90d3": 457,  // Nokia
		"8c9236": 8140, // Aus.Linx
		"8c9351": 8141, // Jigowatts
		"8c941f": 2,    // Cisco
		"8c94cc": 3187, // SFR
		"8c94cf": 8142, // Encell
		"8c99e6": 6433, // TCT
		"8ca2fd": 8143, // Starry
		"8ca399": 8144, // Servercom
		"8ca6df": 1595, // TP-Link
		"8ca96f": 873,  // D&M
		"8ca982": 421,  // Intel
		"8caab5": 6376, // Espressif
		"8caace": 5697, // Xiaomi
		"8cae4c": 8145, // Plugable
		"8cae89": 8146, // Y-cam
		"8caedb": 8147, // Nagtech
		"8cb0e9": 152,  // Samsung
		"8cb64f": 2,    // Cisco
		"8cb82c": 8148, // IPitomy
		"8cb84a": 152,  // Samsung
		"8cb864": 4903, // AcSiP
		"8cb87e": 421,  // Intel
		"8cba25": 3505, // Unionman
		"8cbebe": 5697, // Xiaomi
		"8cbfa6": 152,  // Samsung
		"8cc121": 2153, // Panasonic
		"8cc5b4": 2048, // Sagemcom
		"8cc661": 8149, // Current
		"8cc681": 421,  // Intel
		"8cc7aa": 8150, // Radinet
		"8cc8cd": 152,  // Samsung
		"8ccda2": 8151, // ACTP
		"8ccde8": 1427, // Nintendo
		"8cce4e": 6376, // Espressif
		"8ccf09": 102,  // Dell
		"8ccf5c": 8152, // BEFEGA
		"8ccf8f": 7018, // ITC
		"8cd17b": 8153, // CG
		"8cd3a2": 8154, // VisSim
		"8cd48e": 6281, // Itel
		"8cd9d6": 5697, // Xiaomi
		"8cdb25": 8155, // ESG
		"8cdc02": 3031, // ZTE
		"8cdcd4": 302,  // HP
		"8cde52": 8156, // ISSC
		"8cde99": 8157, // Comlab
		"8cdee6": 152,  // Samsung
		"8cdf9d": 48,   // NEC
		"8ce081": 3031, // ZTE
		"8ce117": 3031, // ZTE
		"8ce38e": 8158, // Kioxia
		"8ce5c0": 152,  // Samsung
		"8ce5ef": 3334, // Huawei
		"8ce78c": 8159, // DK
		"8ce7b3": 8160, // Sonardyne
		"8cea1b": 6294, // Edgecore
		"8cea48": 152,  // Samsung
		"8cebc6": 3334, // Huawei
		"8cec4b": 102,  // Dell
		"8cec7b": 545,  // Apple
		"8ceec6": 8161, // Precepscion
		"8cf112": 687,  // Motorola
		"8cf228": 4253, // Mercury
		"8cf319": 300,  // Siemens
		"8cf5a3": 152,  // Samsung
		"8cf681": 1655, // Silicon
		"8cf710": 4498, // AMPAK
		"8cf773": 457,  // Nokia
		"8cf8c5": 421,  // Intel
		"8cf9c9": 8162, // MESADA
		"8cfaba": 545,  // Apple
		"8cfd18": 3334, // Huawei
		"8cfdde": 2048, // Sagemcom
		"8cfdf0": 5787, // Qualcomm
		"8cfe57": 545,  // Apple
		"8cfe74": 2737, // Ruckus
		"8cfeb4": 8163, // Vsoontech
		"9000db": 152,  // Samsung
		"90013b": 2048, // Sagemcom
		"900218": 3512, // BSkyB
		"900325": 3334, // Huawei
		"9003b7": 2562, // Parrot
		"900628": 152,  // Samsung
		"900917": 8164, // Far-sighted
		"9009d0": 2450, // Synology
		"900a39": 8165, // Wiio
		"900a84": 432,  // Mellanox
		"900bc1": 6819, // Sprocomm
		"900cb4": 8166, // Alinket
		"900cc8": 3521, // Google
		"900d66": 8167, // Digimore
		"900dcb": 125,  // ARRIS
		"901195": 5438, // Amazon
		"9012a1": 6290, // We
		"9016ba": 3334, // Huawei
		"90173f": 3334, // Huawei
		"90179b": 8168, // Nanomegas
		"9017ac": 3334, // Huawei
		"9017c8": 3334, // Huawei
		"90187c": 152,  // Samsung
		"901900": 8169, // SCS
		"901aca": 125,  // ARRIS
		"901b0e": 4,    // Fujitsu
		"901d27": 3031, // ZTE
		"901edd": 8170, // Great
		"9020c2": 1685, // Aruba
		"902106": 3512, // BSkyB
		"902155": 1341, // HTC
		"9023ec": 6904, // Availink
		"9027e4": 545,  // Apple
		"902b34": 1929, // Giga-Byte
		"902bd2": 3334, // Huawei
		"902e16": 4919, // LCFC
		"902e1c": 421,  // Intel
		"902e87": 8171, // LabJack
		"90342b": 8172, // Gatekeeper
		"9035ea": 1655, // Silicon
		"903809": 306,  // Ericsson
		"90380c": 6376, // Espressif
		"903a72": 2737, // Ruckus
		"903aa0": 457,  // Nokia
		"903ae6": 2562, // Parrot
		"903c92": 545,  // Apple
		"903cb3": 6294, // Edgecore
		"903d68": 8173, // G-Printec
		"903d6b": 8174, // Zicon
		"903e7f": 4921, // Fiberhome
		"903eab": 125,  // ARRIS
		"903fea": 3334, // Huawei
		"9043e2": 8175, // Cornami
		"9046b7": 8176, // Vadaro
		"904716": 8177, // Rorze
		"9049fa": 421,  // Intel
		"904c81": 302,  // HP
		"904d4a": 2048, // Sagemcom
		"904dc3": 8178, // Flonidan
		"904e2b": 3334, // Huawei
		"90505a": 8179, // unGlue
		"9050ca": 870,  // Hitron
		"905446": 8180, // TES
		"9055ae": 306,  // Ericsson
		"9055de": 4921, // Fiberhome
		"905682": 8181, // Lenbrook
		"905692": 8182, // Autotalks
		"9056fc": 6265, // Tecno
		"905851": 6392, // Technicolor
		"905c34": 8183, // Sirius
		"905c44": 358,  // Compal
		"905f2e": 6433, // TCT
		"905f8d": 8184, // modas
		"9060f1": 545,  // Apple
		"90610c": 8185, // Fida
		"9061ae": 421,  // Intel
		"90633b": 152,  // Samsung
		"90671c": 3334, // Huawei
		"9068c3": 687,  // Motorola
		"906976": 8186, // Withrobot
		"906cac": 1323, // Fortinet
		"906d05": 8187, // BXB
		"906f18": 67,   // Private
		"907240": 545,  // Apple
		"907282": 2048, // Sagemcom
		"90735a": 687,  // Motorola
		"90749d": 7595, // IRay
		"9077ee": 2,    // Cisco
		"907841": 421,  // Intel
		"9078b2": 5697, // Xiaomi
		"907990": 2079, // Benchmark
		"907a58": 7841, // Zegna-Daidong
		"907af1": 8188, // Wally
		"907e30": 8189, // Lars
		"907eba": 8190, // Utek
		"907f61": 7300, // Chicony
		"908060": 8191, // Nilfisk
		"90808f": 3334, // Huawei
		"90812a": 545,  // Apple
		"908158": 545,  // Apple
		"908175": 152,  // Samsung
		"90840d": 545,  // Apple
		"90848b": 8192, // HDR10+
		"90869b": 3031, // ZTE
		"908c43": 545,  // Apple
		"908d1d": 8193, // GH
		"908d6c": 545,  // Apple
		"908d6e": 102,  // Dell
		"908d78": 803,  // D-Link
		"90903c": 8194, // Trison
		"909497": 3334, // Huawei
		"9094e4": 803,  // D-Link
		"9096f3": 1077, // Buffalo
		"9097d5": 6376, // Espressif
		"9097f3": 152,  // Samsung
		"909838": 3334, // Huawei
		"909864": 8195, // Impex-Sat&Co
		"909a4a": 1595, // TP-Link
		"909c4a": 545,  // Apple
		"909d7d": 125,  // ARRIS
		"909f33": 1252, // EFM
		"90a25b": 545,  // Apple
		"90a2da": 8196, // Gheo
		"90a46a": 8197, // Sisnet
		"90a4de": 1592, // Wistron
		"90a5af": 3334, // Huawei
		"90a62f": 8198, // Naver
		"90a822": 5438, // Amazon
		"90a935": 8199, // JWEntertainment
		"90aac3": 870,  // Hitron
		"90ac3f": 8200, // BrightSign
		"90adf7": 6369, // Vivo
		"90adfc": 7075, // Telechips
		"90ae1b": 1595, // TP-Link
		"90afd1": 8201, // netKTI
		"90b0ed": 545,  // Apple
		"90b11c": 102,  // Dell
		"90b134": 125,  // ARRIS
		"90b144": 152,  // Samsung
		"90b21f": 545,  // Apple
		"90b4dd": 67,   // Private
		"90b622": 152,  // Samsung
		"90b686": 2056, // Murata
		"90b832": 185,  // Extreme
		"90b8d0": 8202, // Joyent
		"90b931": 545,  // Apple
		"90c115": 101,  // Sony
		"90c119": 457,  // Nokia
		"90c1c6": 545,  // Apple
		"90c54a": 6369, // Vivo
		"90c792": 125,  // ARRIS
		"90c7d8": 3031, // ZTE
		"90cc24": 7969, // Synaptics
		"90ccdf": 421,  // Intel
		"90cf15": 457,  // Nokia
		"90cf6f": 8203, // Dlogixs
		"90d74f": 8204, // Bookeen
		"90d852": 6100, // Comtec
		"90d8f3": 3031, // ZTE
		"90d92c": 8205, // HUG-Witschi
		"90da4e": 8206, // Avanu
		"90db46": 8207, // E-Lead
		"90dd5d": 545,  // Apple
		"90dffb": 8208, // Homerider
		"90e17b": 545,  // Apple
		"90e2ba": 421,  // Intel
		"90e6ba": 1806, // ASUS
		"90e7c4": 1341, // HTC
		"90e868": 3004, // Azurewave
		"90ec50": 8209, // C.O.B.O
		"90ec77": 8210, // silicom
		"90eec7": 152,  // Samsung
		"90ef68": 2692, // ZyXEL
		"90f052": 7030, // MEIZU
		"90f157": 797,  // Garmin
		"90f1aa": 152,  // Samsung
		"90f305": 531,  // HUMAX
		"90f3b7": 8211, // Kirisun
		"90f644": 3334, // Huawei
		"90f652": 1595, // TP-Link
		"90f891": 1269, // Kaonmedia
		"90f9b7": 3334, // Huawei
		"90fb5b": 620,  // Avaya
		"90fd61": 545,  // Apple
		"90fd73": 3031, // ZTE
		"90fd9f": 1655, // Silicon
		"940006": 8212, // jinyoung
		"940070": 457,  // Nokia
		"9400b0": 3334, // Huawei
		"940149": 8213, // AutoHotBox
		"9401c2": 152,  // Samsung
		"940230": 4079, // Logitech
		"94026b": 8214, // Optictimes
		"94049c": 3334, // Huawei
		"940853": 2873, // Liteon
		"9408c7": 3334, // Huawei
		"940937": 531,  // HUMAX
		"940b19": 3334, // Huawei
		"940b2d": 8215, // NetView
		"940bd5": 8216, // Himax
		"940c6d": 1595, // TP-Link
		"940c98": 545,  // Apple
		"940d2d": 3857, // Universal
		"940e6b": 3334, // Huawei
		"94103e": 2468, // Belkin
		"94147a": 6369, // Vivo
		"941625": 545,  // Apple
		"941700": 5697, // Xiaomi
		"941865": 1368, // Netgear
		"941882": 302,  // HP
		"94193a": 8217, // Elvaco
		"941c56": 2246, // Actiontec
		"941f3a": 8218, // Ambiq
		"942053": 457,  // Nokia
		"942197": 8219, // Stalmart
		"942533": 3334, // Huawei
		"942790": 6433, // TCT
		"942957": 8220, // Airpo
		"942a3f": 8221, // Diversey
		"942a6f": 2978, // Ubiquiti
		"942cb3": 531,  // HUMAX
		"942ddc": 152,  // Samsung
		"942e17": 56,   // Schneider
		"942e63": 8222, // Finsecur
		"94319b": 8223, // Alphatronics
		"9431cb": 6369, // Vivo
		"9433dd": 8224, // Taco
		"943469": 1655, // Silicon
		"94350a": 152,  // Samsung
		"9437f7": 3334, // Huawei
		"943a91": 5438, // Amazon
		"943af0": 457,  // Nokia
		"943bb1": 1269, // Kaonmedia
		"943cc6": 6376, // Espressif
		"943fc2": 302,  // HP
		"9440a2": 8225, // Anywave
		"9440c9": 302,  // HP
		"9441c1": 8226, // Mini-Cam
		"94434d": 4564, // Ciena
		"944444": 869,  // LG
		"944452": 2468, // Belkin
		"944696": 2683, // BaudTec
		"944996": 8227, // WiSilica
		"944a0c": 2074, // Sercomm
		"945047": 8228, // Rechnerbetriebsgruppe
		"945089": 8229, // SimonsVoss
		"945103": 152,  // Samsung
		"945493": 8230, // Rigado
		"9454df": 8231, // YST
		"9457a5": 302,  // HP
		"9458cb": 1427, // Nintendo
		"94592d": 8232, // EKE
		"945afc": 5438, // Amazon
		"945b7e": 8233, // Trilobit
		"945c9a": 545,  // Apple
		"945f34": 5695, // Renesas
		"946010": 3334, // Huawei
		"9460d5": 1685, // Aruba
		"94611e": 8234, // Wata
		"946124": 8235, // Pason
		"946269": 125,  // ARRIS
		"946372": 6369, // Vivo
		"9463d1": 152,  // Samsung
		"946424": 1685, // Aruba
		"94652d": 6923, // OnePlus
		"94659c": 421,  // Intel
		"9466e7": 8236, // WOM
		"946a77": 6392, // Technicolor
		"946ab0": 2639, // Arcadyan
		"9470d2": 8237, // Winfirm
		"9471ac": 6433, // TCT
		"94756e": 2541, // QinetiQ
		"9476b7": 152,  // Samsung
		"94772b": 3334, // Huawei
		"947bbe": 8238, // Ubicquia
		"947be7": 152,  // Samsung
		"9481a4": 8239, // Azuray
		"9483c4": 4371, // GL
		"94857a": 8240, // Evantage
		"9486cd": 8241, // Seoul
		"94877c": 125,  // ARRIS
		"9487e0": 5697, // Xiaomi
		"948bc1": 152,  // Samsung
		"948d50": 8242, // Beamex
		"948ed3": 3793, // Arista
		"948fcf": 125,  // ARRIS
		"949010": 3334, // Huawei
		"94917f": 2544, // Askey
		"9492bc": 7511, // Syntech
		"9492d2": 8243, // KCF
		"949426": 545,  // Apple
		"94944a": 8244, // Particle
		"9495a0": 3521, // Google
		"949869": 3031, // ZTE
		"949aa9": 612,  // Microsoft
		"949b2c": 185,  // Extreme
		"949d57": 2153, // Panasonic
		"949f3e": 2047, // Sonos
		"94a04e": 8245, // Bostex
		"94a1a2": 4498, // AMPAK
		"94a3ca": 8246, // KonnectONE
		"94a4f9": 3334, // Huawei
		"94a67e": 1368, // Netgear
		"94a7b7": 3031, // ZTE
		"94a7bc": 8247, // BodyMedia
		"94aa0a": 4921, // Fiberhome
		"94aab8": 8248, // Joview
		"94abfe": 457,  // Nokia
		"94acca": 8249, // trivum
		"94aef0": 2,    // Cisco
		"94b01f": 545,  // Apple
		"94b10a": 152,  // Samsung
		"94b271": 3334, // Huawei
		"94b2cc": 6117, // Pioneer
		"94b34f": 2737, // Ruckus
		"94b40f": 1685, // Aruba
		"94b555": 6376, // Espressif
		"94b819": 457,  // Nokia
		"94b86d": 421,  // Intel
		"94b8c5": 1586, // RuggedCom
		"94b97e": 6376, // Espressif
		"94b9b4": 8250, // Aptos
		"94bbae": 8251, // Husqvarna
		"94be46": 687,  // Motorola
		"94bf2d": 545,  // Apple
		"94bf80": 3031, // ZTE
		"94bf94": 826,  // Juniper
		"94bfc4": 2737, // Ruckus
		"94c038": 8252, // Tallac
		"94c150": 1939, // 2Wire
		"94c2bd": 8253, // Tecnobit
		"94c691": 6672, // EliteGroup
		"94c6eb": 8254, // NOVA
		"94c7af": 8255, // Raylios
		"94c962": 8256, // Teseq
		"94ccb9": 125,  // ARRIS
		"94cdac": 8257, // Creowave
		"94ce2c": 101,  // Sony
		"94ce31": 2058, // CTS
		"94d00d": 3334, // Huawei
		"94d019": 8258, // Cydle
		"94d299": 8259, // Techmation
		"94d2bc": 3334, // Huawei
		"94d469": 2,    // Cisco
		"94d505": 4921, // Fiberhome
		"94d6db": 8260, // NexFi
		"94d771": 152,  // Samsung
		"94d859": 6433, // TCT
		"94d93c": 8261, // Enelps
		"94d9b3": 1595, // TP-Link
		"94db49": 8262, // Sitcorp
		"94db56": 101,  // Sony
		"94dbc9": 3004, // Azurewave
		"94dbda": 3334, // Huawei
		"94dc4e": 8263, // AEV
		"94de0e": 8264, // SmartOptics
		"94de80": 1929, // Giga-Byte
		"94deb8": 1655, // Silicon
		"94df4e": 1592, // Wistron
		"94df58": 8265, // IJ
		"94e129": 152,  // Samsung
		"94e23c": 421,  // Intel
		"94e3ee": 3031, // ZTE
		"94e4ba": 3334, // Huawei
		"94e686": 6376, // Espressif
		"94e6f7": 421,  // Intel
		"94e70b": 421,  // Intel
		"94e7ea": 3334, // Huawei
		"94e8c5": 125,  // ARRIS
		"94e96a": 545,  // Apple
		"94e979": 2873, // Liteon
		"94e98c": 457,  // Nokia
		"94e9ee": 3334, // Huawei
		"94ea32": 545,  // Apple
		"94eb2c": 3521, // Google
		"94ebcd": 2220, // BlackBerry
		"94f128": 302,  // HP
		"94f278": 8266, // Elma
		"94f665": 2737, // Ruckus
		"94f692": 8267, // Geminico
		"94f6a3": 545,  // Apple
		"94f6d6": 545,  // Apple
		"94f7ad": 826,  // Juniper
		"94fb29": 768,  // Zebra
		"94fd1d": 8268, // WhereWhen
		"94fe22": 3334, // Huawei
		"94fef4": 2048, // Sagemcom
		"94ff3c": 1323, // Fortinet
		"98006a": 3031, // ZTE
		"980074": 2050, // Raisecom
		"9800c6": 545,  // Apple
		"9801a7": 545,  // Apple
		"980284": 8269, // Theobroma
		"98039b": 432,  // Mellanox
		"9803d8": 545,  // Apple
		"98063c": 152,  // Samsung
		"9809cf": 6923, // OnePlus
		"980c82": 152,  // Samsung
		"980ca5": 687,  // Motorola
		"980d2e": 1341, // HTC
		"980d51": 3334, // Huawei
		"980d67": 2692, // ZyXEL
		"980e24": 8270, // Phytium
		"980ee4": 67,   // Private
		"981082": 8271, // Nsolution
		"9810e8": 545,  // Apple
		"981333": 3031, // ZTE
		"9814d2": 8272, // Avonic
		"98192c": 6294, // Edgecore
		"981a35": 3334, // Huawei
		"981dfa": 152,  // Samsung
		"981e19": 2048, // Sagemcom
		"98208e": 8273, // Definium
		"9822ef": 2873, // Liteon
		"98234e": 8274, // Micromedia
		"9828a6": 358,  // Compal
		"9829a6": 358,  // Compal
		"982cbc": 421,  // Intel
		"982cbe": 1939, // 2Wire
		"982d68": 152,  // Samsung
		"982dba": 8275, // Fibergate
		"982ff8": 3334, // Huawei
		"983571": 8276, // Sub10
		"9835b8": 8277, // Assembled
		"9835ed": 3334, // Huawei
		"98387d": 8278, // Itronic
		"98398e": 152,  // Samsung
		"983b16": 4498, // AMPAK
		"983b67": 7446, // DWnet
		"983b8f": 421,  // Intel
		"983f60": 3334, // Huawei
		"9840bb": 102,  // Dell
		"98415c": 1427, // Nintendo
		"984246": 8279, // SOL
		"984265": 2048, // Sagemcom
		"9843da": 8280, // Intertech
		"9843fa": 421,  // Intel
		"9844ce": 3334, // Huawei
		"98460a": 545,  // Apple
		"984827": 1595, // TP-Link
		"984874": 3334, // Huawei
		"984914": 1592, // Wistron
		"9849e1": 1041, // Boeing
		"984b4a": 125,  // ARRIS
		"984be1": 302,  // HP
		"984fee": 421,  // Intel
		"98502e": 545,  // Apple
		"98523d": 3939, // Sunitec
		"98524a": 6392, // Technicolor
		"9852b1": 152,  // Samsung
		"98541b": 421,  // Intel
		"98588a": 8281, // SysGRATION
		"985aeb": 545,  // Apple
		"985bb0": 8282, // Kmdata
		"985d46": 8283, // PeopleNet
		"985d82": 3793, // Arista
		"985e1b": 8284, // ConversDigital
		"985f4f": 8285, // Tongfang
		"985fd3": 612,  // Microsoft
		"986022": 8286, // EMW
		"9860ca": 545,  // Apple
		"98672e": 7039, // Skullcandy
		"98698a": 545,  // Apple
		"986b3d": 125,  // ARRIS
		"986cf5": 3031, // ZTE
		"986dc8": 35,   // Toshiba
		"9873c4": 619,  // Sage
		"9874da": 6295, // Infinix
		"98751a": 3334, // Huawei
		"9876b6": 8287, // Adafruit
		"9877e7": 1269, // Kaonmedia
		"987a10": 306,  // Ericsson
		"987a14": 612,  // Microsoft
		"987e46": 8288, // Emizon
		"987ee3": 6369, // Vivo
		"9880ee": 152,  // Samsung
		"988217": 8289, // Disruptive
		"988389": 152,  // Samsung
		"9886b1": 8290, // Flyaudio
		"988b5d": 2048, // Sagemcom
		"988bad": 8291, // Corintech
		"988d46": 421,  // Intel
		"988e4a": 8292, // Noxus
		"988e79": 8293, // Qudelix
		"988ed4": 6281, // Itel
		"989096": 102,  // Dell
		"9893cc": 869,  // LG
		"9897cc": 1595, // TP-Link
		"9897d1": 6428, // MitraStar
		"989ab9": 3031, // ZTE
		"989c57": 3334, // Huawei
		"989d5d": 6392, // Technicolor
		"989e63": 545,  // Apple
		"98a404": 306,  // Ericsson
		"98a40e": 8294, // Snap
		"98a5f9": 545,  // Apple
		"98ad1d": 3334, // Huawei
		"98ae71": 8295, // VVDN
		"98af65": 421,  // Intel
		"98b039": 457,  // Nokia
		"98b08b": 152,  // Samsung
		"98b3ef": 3334, // Huawei
		"98b6e9": 1427, // Nintendo
		"98b8ba": 869,  // LG
		"98b8bc": 152,  // Samsung
		"98b8e3": 545,  // Apple
		"98ba39": 3873, // Doro
		"98bb99": 6848, // Phicomm
		"98bc57": 1486, // SVA
		"98bc99": 8296, // Edeltech
		"98be94": 372,  // IBM
		"98c5db": 306,  // Ericsson
		"98c845": 8297, // PacketAccess
		"98c8b8": 6369, // Vivo
		"98ca33": 545,  // Apple
		"98cb27": 8298, // Galore
		"98cba4": 2079, // Benchmark
		"98cdac": 6376, // Espressif
		"98cdb4": 8299, // Virident
		"98d293": 3521, // Google
		"98d6bb": 545,  // Apple
		"98d6f7": 869,  // LG
		"98d88c": 76,   // Nortel
		"98da92": 8300, // Vuzix
		"98dac4": 1595, // TP-Link
		"98dcd9": 8301, // UNITEC
		"98dd60": 545,  // Apple
		"98ddea": 6295, // Infinix
		"98ded0": 1595, // TP-Link
		"98e0d9": 545,  // Apple
		"98e165": 8302, // Accutome
		"98e476": 8303, // Zentan
		"98e743": 102,  // Dell
		"98e79a": 221,  // Foxconn
		"98e7f4": 302,  // HP
		"98e7f5": 3334, // Huawei
		"98e848": 8304, // Axiim
		"98e8fa": 1427, // Nintendo
		"98ed5c": 7309, // Tesla
		"98ed7e": 5820, // eero
		"98eecb": 1592, // Wistron
		"98ef9b": 6577, // Ohsung
		"98f058": 8305, // Lynxspringl
		"98f083": 3334, // Huawei
		"98f0ab": 545,  // Apple
		"98f170": 2056, // Murata
		"98f217": 3797, // Castlenet
		"98f2b3": 302,  // HP
		"98f428": 3031, // ZTE
		"98f4ab": 6376, // Espressif
		"98f537": 3031, // ZTE
		"98f5a9": 6577, // Ohsung
		"98f621": 5697, // Xiaomi
		"98f781": 125,  // ARRIS
		"98f7d7": 125,  // ARRIS
		"98f8c1": 3225, // IDT
		"98fa9b": 4919, // LCFC
		"98faa7": 8306, // Innonet
		"98fae3": 5697, // Xiaomi
		"98fb12": 1883, // Grand
		"98fc11": 1783, // Linksys
		"98fd74": 8307, // ACT
		"98fdb4": 389,  // Primax
		"98fe03": 306,  // Ericsson
		"98fe94": 545,  // Apple
		"98ff6a": 8308, // OTECTechnology
		"98ffd0": 2664, // Lenovo
		"9c0298": 152,  // Samsung
		"9c0473": 8309, // Tecmobile
		"9c04eb": 545,  // Apple
		"9c066e": 7600, // Hytera
		"9c0b05": 5820, // eero
		"9c1874": 457,  // Nokia
		"9c1c12": 1685, // Aruba
		"9c1c37": 7031, // AltoBeam
		"9c1d36": 3334, // Huawei
		"9c1e95": 2246, // Actiontec
		"9c1ea4": 5695, // Renesas
		"9c1fdd": 8310, // Accupix
		"9c207b": 545,  // Apple
		"9c216a": 1595, // TP-Link
		"9c220e": 8311, // TASCAN
		"9c2595": 152,  // Samsung
		"9c2840": 8312, // Discovery
		"9c28b3": 545,  // Apple
		"9c28ef": 3334, // Huawei
		"9c28f7": 5697, // Xiaomi
		"9c293f": 545,  // Apple
		"9c2976": 421,  // Intel
		"9c2a83": 152,  // Samsung
		"9c2ba6": 5440, // Ruijie
		"9c2ea1": 5697, // Xiaomi
		"9c2f4e": 3031, // ZTE
		"9c2f9d": 2873, // Liteon
		"9c31c3": 3512, // BSkyB
		"9c32ce": 87,   // Canon
		"9c3426": 125,  // ARRIS
		"9c35eb": 545,  // Apple
		"9c37f4": 3334, // Huawei
		"9c3aaf": 152,  // Samsung
		"9c3dcf": 1368, // Netgear
		"9c3e53": 545,  // Apple
		"9c3eaa": 8313, // EnvyLogic
		"9c40cd": 5354, // Synclayer
		"9c417c": 8314, // Hame
		"9c44a6": 8315, // SwiftTest
		"9c4a7b": 457,  // Nokia
		"9c4cae": 5319, // Mesa
		"9c4e20": 2,    // Cisco
		"9c4e36": 421,  // Intel
		"9c4e8e": 8316, // ALT
		"9c4ebf": 8317, // BoxCast
		"9c4fcf": 6433, // TCT
		"9c4fda": 545,  // Apple
		"9c50d1": 2056, // Murata
		"9c50ee": 1640, // Cambridge
		"9c52f8": 3334, // Huawei
		"9c5440": 67,   // Private
		"9c54da": 7684, // SkyBell
		"9c5636": 3334, // Huawei
		"9c57ad": 2,    // Cisco
		"9c583c": 545,  // Apple
		"9c5a44": 358,  // Compal
		"9c5a81": 5697, // Xiaomi
		"9c5b96": 8318, // NMR
		"9c5c8e": 1806, // ASUS
		"9c5cf9": 101,  // Sony
		"9c5d12": 185,  // Extreme
		"9c5d95": 8319, // VTC
		"9c5fb0": 152,  // Samsung
		"9c611d": 2153, // Panasonic
		"9c62ab": 4654, // Sumavision
		"9c63ed": 3031, // ZTE
		"9c648b": 545,  // Apple
		"9c65b0": 152,  // Samsung
		"9c65f9": 4903, // AcSiP
		"9c685b": 8320, // Octonion
		"9c6865": 4921, // Fiberhome
		"9c6937": 6652, // Qorvo
		"9c69d1": 3334, // Huawei
		"9c6b37": 5695, // Renesas
		"9c6c15": 612,  // Microsoft
		"9c6f52": 3031, // ZTE
		"9c713a": 3334, // Huawei
		"9c7370": 3334, // Huawei
		"9c741a": 3334, // Huawei
		"9c746f": 3334, // Huawei
		"9c760e": 545,  // Apple
		"9c7613": 6946, // Ring
		"9c77aa": 8321, // Nadasnv
		"9c79ac": 8322, // Suntec
		"9c7a03": 4564, // Ciena
		"9c7bef": 302,  // HP
		"9c7da3": 3334, // Huawei
		"9c80df": 2639, // Arcadyan
		"9c823f": 3334, // Huawei
		"9c8281": 6369, // Vivo
		"9c83bf": 8323, // PRO-VISION
		"9c84bf": 545,  // Apple
		"9c8566": 4030, // Wingtech
		"9c88ad": 4921, // Fiberhome
		"9c8acb": 826,  // Juniper
		"9c8ba0": 545,  // Apple
		"9c8bf1": 8324, // Warehouse
		"9c8c6e": 152,  // Samsung
		"9c8cd8": 302,  // HP
		"9c8d7c": 430,  // Alpsalpine
		"9c8dd3": 8325, // Leonton
		"9c8e99": 302,  // HP
		"9c8e9c": 3334, // Huawei
		"9c8ecd": 8326, // Amcrest
		"9c8edc": 4820, // Teracom
		"9c9019": 8327, // Beyless
		"9c934e": 0,    // Xerox
		"9c93b0": 8328, // Megatronix
		"9c93e4": 67,   // Private
		"9c9567": 3334, // Huawei
		"9c95f8": 8329, // SmartDoor
		"9c9726": 6392, // Technicolor
		"9c9789": 6627, // 1MORE
		"9c99a0": 5697, // Xiaomi
		"9c99cd": 8330, // Voippartners
		"9c9c1d": 8331, // Starkey
		"9c9c1f": 6376, // Espressif
		"9c9d5d": 8332, // Raden
		"9c9e71": 3334, // Huawei
		"9ca134": 8333, // Nike
		"9ca513": 152,  // Samsung
		"9ca570": 5820, // eero
		"9ca577": 8334, // Osorno
		"9ca5c0": 6369, // Vivo
		"9ca615": 1595, // TP-Link
		"9ca69d": 8335, // Whaley
		"9ca9e4": 3031, // ZTE
		"9caa1b": 612,  // Microsoft
		"9cac6d": 3857, // Universal
		"9cadef": 8336, // Obihai
		"9caed3": 46,   // Epson
		"9caf6f": 6281, // Itel
		"9cafca": 2,    // Cisco
		"9cb206": 8337, // Procentec
		"9cb2b2": 3334, // Huawei
		"9cb2e8": 3334, // Huawei
		"9cb654": 302,  // HP
		"9cb6d0": 8338, // Rivet
		"9cb70d": 2873, // Liteon
		"9cb793": 8339, // Creatcomm
		"9cb8b4": 4498, // AMPAK
		"9cbcf0": 5697, // Xiaomi
		"9cbd6e": 8340, // DERA
		"9cbd9d": 8341, // SkyDisk
		"9cbee0": 8342, // Biosoundlab
		"9cc077": 8343, // PrintCounts
		"9cc0d2": 8344, // Conductix-Wampfler
		"9cc172": 3334, // Huawei
		"9cc2c4": 7721, // Inspur
		"9cc7a6": 621,  // AVM
		"9cc7d1": 3205, // Sharp
		"9cc8ae": 8345, // Becton
		"9cc8fc": 125,  // ARRIS
		"9cc950": 8346, // Baumer
		"9cc9eb": 1368, // Netgear
		"9ccad9": 457,  // Nokia
		"9ccc83": 826,  // Juniper
		"9cd24b": 3031, // ZTE
		"9cd332": 8347, // PLC
		"9cd35b": 152,  // Samsung
		"9cd36d": 1368, // Netgear
		"9cd48b": 8348, // Innolux
		"9cd57d": 2,    // Cisco
		"9cd643": 803,  // D-Link
		"9cd917": 687,  // Motorola
		"9cd9cb": 8349, // Lesira
		"9cda3e": 421,  // Intel
		"9cdb07": 8350, // Thum+Mahr
		"9cdc71": 302,  // HP
		"9ce063": 152,  // Samsung
		"9ce10e": 8351, // NCTech
		"9ce176": 2,    // Cisco
		"9ce230": 8352, // Julong
		"9ce33f": 545,  // Apple
		"9ce374": 3334, // Huawei
		"9ce635": 1427, // Nintendo
		"9ce65e": 545,  // Apple
		"9ce6e7": 152,  // Samsung
		"9ce7bd": 8353, // Winduskorea
		"9ce82b": 6369, // Vivo
		"9ce91c": 3031, // ZTE
		"9cebe8": 8354, // BizLink
		"9cec61": 3334, // Huawei
		"9cedfa": 8355, // EVUlution
		"9cf155": 457,  // Nokia
		"9cf387": 545,  // Apple
		"9cf48e": 545,  // Apple
		"9cfbd5": 6369, // Vivo
		"9cfc01": 545,  // Apple
		"9cfc28": 545,  // Apple
		"9cfcd1": 8356, // Aetheris
		"9cfce8": 421,  // Intel
		"9cfea1": 4921, // Fiberhome
		"9cffbe": 8357, // OTSL
		"9cffc2": 8358, // AVI
		"a002dc": 5438, // Amazon
		"a00460": 1368, // Netgear
		"a00798": 152,  // Samsung
		"a0086f": 3334, // Huawei
		"a0092e": 3031, // ZTE
		"a0094c": 8359, // CenturyLink
		"a009ed": 620,  // Avaya
		"a00abf": 8360, // Wieson
		"a00bba": 152,  // Samsung
		"a00f37": 2,    // Cisco
		"a01081": 152,  // Samsung
		"a01290": 620,  // Avaya
		"a013cb": 4921, // Fiberhome
		"a0143d": 2562, // Parrot
		"a0165c": 8361, // Triteka
		"a01828": 545,  // Apple
		"a01842": 3869, // Comtrend
		"a01b29": 2048, // Sagemcom
		"a01c87": 3505, // Unionman
		"a01c8d": 3334, // Huawei
		"a01d48": 302,  // HP
		"a01e0b": 8362, // MINIX
		"a020a6": 6376, // Espressif
		"a02195": 152,  // Samsung
		"a021b7": 1368, // Netgear
		"a022de": 6369, // Vivo
		"a0239f": 2,    // Cisco
		"a027b6": 152,  // Samsung
		"a02919": 102,  // Dell
		"a02bb8": 302,  // HP
		"a02c36": 6440, // FN-Link
		"a031db": 3334, // Huawei
		"a03299": 2664, // Lenovo
		"a0341b": 8363, // Adero
		"a03679": 3334, // Huawei
		"a0369f": 421,  // Intel
		"a036fa": 8364, // Ettus
		"a039ee": 2048, // Sagemcom
		"a039f7": 869,  // LG
		"a03b1b": 8365, // Inspire
		"a03be3": 545,  // Apple
		"a03d6e": 2,    // Cisco
		"a03d6f": 2,    // Cisco
		"a04025": 8366, // Actioncable
		"a04041": 8367, // SAMWONFA
		"a0406f": 3334, // Huawei
		"a040a0": 1368, // Netgear
		"a0412d": 8368, // Lansen
		"a04147": 3334, // Huawei
		"a0423f": 6158, // Tyan
		"a0445c": 3334, // Huawei
		"a0481c": 302,  // HP
		"a04a5e": 612,  // Microsoft
		"a04cc1": 8369, // Helixtech
		"a04e01": 4681, // CENTRAL
		"a04e04": 457,  // Nokia
		"a04ea7": 545,  // Apple
		"a04ecf": 545,  // Apple
		"a04f85": 869,  // LG
		"a0510b": 421,  // Intel
		"a051c6": 620,  // Avaya
		"a0554f": 2,    // Cisco
		"a055de": 125,  // ARRIS
		"a056f3": 545,  // Apple
		"a057e3": 3334, // Huawei
		"a05950": 421,  // Intel
		"a05b21": 8370, // ENVINET
		"a05dc1": 8371, // TMCT
		"a05de7": 8372, // DIRECTV
		"a05e6b": 8373, // MELPER
		"a06090": 152,  // Samsung
		"a06260": 67,   // Private
		"a06391": 1368, // Netgear
		"a0648f": 2544, // Askey
		"a06518": 8374, // Vnpt
		"a06610": 4,    // Fujitsu
		"a0687e": 125,  // ARRIS
		"a06986": 8375, // Wellav
		"a06a00": 4212, // Verilink
		"a06a44": 3467, // Vizio
		"a06cec": 4555, // RIM
		"a06faa": 869,  // LG
		"a070b7": 3334, // Huawei
		"a071a9": 457,  // Nokia
		"a0722c": 531,  // HUMAX
		"a07332": 8376, // Cashmaster
		"a073fc": 8377, // Rancore
		"a07591": 152,  // Samsung
		"a075ea": 8378, // BoxLock
		"a0764e": 6376, // Espressif
		"a07751": 8379, // ASMedia
		"a07771": 8380, // Vialis
		"a07817": 545,  // Apple
		"a078ba": 2277, // Pantech
		"a08069": 421,  // Intel
		"a0821f": 152,  // Samsung
		"a082c7": 8381, // P.T.I
		"a084cb": 8382, // SonicSensory
		"a085fc": 612,  // Microsoft
		"a086c6": 5697, // Xiaomi
		"a08869": 421,  // Intel
		"a088b4": 421,  // Intel
		"a08c9b": 8383, // Xtreme
		"a08cf8": 3334, // Huawei
		"a08cfd": 302,  // HP
		"a08d16": 3334, // Huawei
		"a08e78": 2048, // Sagemcom
		"a090de": 8384, // Veedims
		"a09169": 869,  // LG
		"a091a2": 6923, // OnePlus
		"a091c8": 3031, // ZTE
		"a09351": 2,    // Cisco
		"a09805": 8385, // OpenVox
		"a0999b": 545,  // Apple
		"a09d91": 8386, // SoundBridge
		"a09e1a": 4479, // Polar
		"a09f7a": 803,  // D-Link
		"a0a0dc": 3334, // Huawei
		"a0a23c": 8387, // Gpms
		"a0a309": 545,  // Apple
		"a0a33b": 3334, // Huawei
		"a0a3b8": 8388, // Wiscloud
		"a0a3e2": 2246, // Actiontec
		"a0a3f0": 803,  // D-Link
		"a0a4c5": 421,  // Intel
		"a0a65c": 5560, // Supercomputing
		"a0a8cd": 421,  // Intel
		"a0aafd": 8389, // EraThink
		"a0ab1b": 803,  // D-Link
		"a0ac69": 152,  // Samsung
		"a0ada1": 8390, // JMR
		"a0afbd": 421,  // Intel
		"a0b3cc": 302,  // HP
		"a0b439": 2,    // Cisco
		"a0b4a5": 152,  // Samsung
		"a0b4bf": 644,  // InfiNet
		"a0b53c": 6392, // Technicolor
		"a0b549": 2639, // Arcadyan
		"a0b9ed": 8391, // Skytap
		"a0bdcd": 3512, // BSkyB
		"a0bfa5": 8392, // Coresys
		"a0c3de": 8393, // Triton
		"a0c562": 125,  // ARRIS
		"a0c589": 421,  // Intel
		"a0c9a0": 2056, // Murata
		"a0cbfd": 152,  // Samsung
		"a0cc2b": 2056, // Murata
		"a0cf5b": 2,    // Cisco
		"a0cff5": 3031, // ZTE
		"a0d05b": 152,  // Samsung
		"a0d0dc": 5438, // Amazon
		"a0d12a": 8394, // AXPRO
		"a0d2b1": 5438, // Amazon
		"a0d37a": 421,  // Intel
		"a0d3c1": 302,  // HP
		"a0d635": 8395, // WBS
		"a0d722": 152,  // Samsung
		"a0d795": 545,  // Apple
		"a0d7a0": 3334, // Huawei
		"a0d7f3": 152,  // Samsung
		"a0d807": 3334, // Huawei
		"a0d83d": 4921, // Fiberhome
		"a0dc04": 8396, // Becker-Antriebe
		"a0dd97": 8397, // PolarLink
		"a0dde5": 3205, // Sharp
		"a0de0f": 3334, // Huawei
		"a0df15": 3334, // Huawei
		"a0e0af": 2,    // Cisco
		"a0e201": 8398, // AVTrace
		"a0e453": 101,  // Sony
		"a0e4cb": 2692, // ZyXEL
		"a0e5e9": 8399, // enimai
		"a0e617": 8400, // Matis
		"a0e70b": 421,  // Intel
		"a0e7ae": 125,  // ARRIS
		"a0eb76": 8401, // AirCUVE
		"a0ec80": 3031, // ZTE
		"a0ecf9": 2,    // Cisco
		"a0edcd": 545,  // Apple
		"a0f3c1": 1595, // TP-Link
		"a0f419": 457,  // Nokia
		"a0f450": 1341, // HTC
		"a0f459": 6440, // FN-Link
		"a0f479": 3334, // Huawei
		"a0f849": 2,    // Cisco
		"a0f9e0": 8402, // Vivatel
		"a0fbc5": 545,  // Apple
		"a0ff70": 6392, // Technicolor
		"a400e2": 3334, // Huawei
		"a40130": 8403, // ABIsystems
		"a402b9": 421,  // Intel
		"a40450": 3202, // nFore
		"a4056e": 6539, // Tiinlab
		"a405d6": 125,  // ARRIS
		"a407b6": 152,  // Samsung
		"a40801": 5438, // Amazon
		"a408ea": 2056, // Murata
		"a408f5": 2048, // Sagemcom
		"a40bed": 8404, // Carry
		"a40cc3": 2,    // Cisco
		"a40e2b": 7234, // Facebook
		"a41115": 1926, // Bosch
		"a41162": 8405, // Arlo
		"a41194": 2664, // Lenovo
		"a4134e": 8406, // Luxul
		"a41588": 125,  // ARRIS
		"a416e7": 3334, // Huawei
		"a41752": 8407, // Hifocus
		"a4178b": 3334, // Huawei
		"a41875": 2,    // Cisco
		"a41908": 4921, // Fiberhome
		"a41a3a": 1595, // TP-Link
		"a41f72": 102,  // Dell
		"a4218a": 76,   // Nortel
		"a424b3": 8408, // FlatFrog
		"a424dd": 8409, // Cambrionix
		"a4251b": 620,  // Avaya
		"a42983": 1041, // Boeing
		"a429b7": 8410, // bluesky
		"a42a95": 803,  // D-Link
		"a42b8c": 1368, // Netgear
		"a42bb0": 1595, // TP-Link
		"a4307a": 152,  // Samsung
		"a43111": 7113, // ZIV
		"a43135": 545,  // Apple
		"a433d1": 8411, // Fibrlink
		"a433d7": 6428, // MitraStar
		"a434d9": 421,  // Intel
		"a4352d": 8412, // TRIZ
		"a438cc": 1427, // Nintendo
		"a43a69": 8413, // Vers
		"a43b0e": 3334, // Huawei
		"a44027": 3031, // ZTE
		"a4423b": 421,  // Intel
		"a4438c": 125,  // ARRIS
		"a444d1": 4030, // Wingtech
		"a44519": 5697, // Xiaomi
		"a4466b": 8414, // EOC
		"a446b4": 3334, // Huawei
		"a44ad3": 4755, // ST
		"a44bd5": 5697, // Xiaomi
		"a44c11": 2,    // Cisco
		"a44cc8": 102,  // Dell
		"a44e31": 421,  // Intel
		"a45046": 5697, // Xiaomi
		"a45055": 8415, // Busware.DE
		"a45129": 8416, // XAG
		"a4516f": 612,  // Microsoft
		"a4530e": 2,    // Cisco
		"a45590": 5697, // Xiaomi
		"a45602": 8417, // fenglian
		"a4561b": 8418, // MCOT
		"a45630": 2,    // Cisco
		"a456cc": 6392, // Technicolor
		"a45802": 8419, // Shin-IL
		"a45a1c": 8420, // smart-electronic
		"a45c27": 1427, // Nintendo
		"a45d36": 302,  // HP
		"a45e5a": 8421, // ACTIVIO
		"a45e60": 545,  // Apple
		"a45f9b": 8422, // Nexell
		"a46011": 1647, // Verifone
		"a46032": 2253, // MRV
		"a46191": 8423, // NamJunSa
		"a46706": 545,  // Apple
		"a468bc": 8424, // Oakley
		"a46bb6": 421,  // Intel
		"a46c2a": 2,    // Cisco
		"a46cf1": 152,  // Samsung
		"a46da4": 3334, // Huawei
		"a470d6": 687,  // Motorola
		"a47174": 3334, // Huawei
		"a47733": 3521, // Google
		"a47760": 457,  // Nokia
		"a477f3": 545,  // Apple
		"a47806": 2,    // Cisco
		"a47886": 620,  // Avaya
		"a479e4": 8425, // KLINFO
		"a47aa4": 125,  // ARRIS
		"a47acf": 8426, // Vibicom
		"a47b1a": 3334, // Huawei
		"a47b2c": 457,  // Nokia
		"a47b85": 8427, // ULTIMEDIA
		"a47b9d": 6376, // Espressif
		"a47c14": 8428, // ChargeStorm
		"a47c1f": 8429, // Cobham
		"a47cc9": 3334, // Huawei
		"a47e39": 3031, // ZTE
		"a481ee": 457,  // Nokia
		"a48269": 8430, // Datrium
		"a483e7": 545,  // Apple
		"a48431": 152,  // Samsung
		"a4856b": 8431, // Q
		"a48873": 2,    // Cisco
		"a48cc0": 8432, // JLG
		"a48cdb": 2664, // Lenovo
		"a48d3b": 3467, // Vizio
		"a48e0a": 8433, // DeLaval
		"a491b1": 6392, // Technicolor
		"a492cb": 457,  // Nokia
		"a4933f": 3334, // Huawei
		"a4934c": 2,    // Cisco
		"a49426": 8434, // Elgama-Elektronika
		"a49733": 2544, // Askey
		"a49813": 125,  // ARRIS
		"a49947": 3334, // Huawei
		"a49a58": 152,  // Samsung
		"a49b4f": 3334, // Huawei
		"a49bcd": 2,    // Cisco
		"a49d49": 8435, // Ketra
		"a49edb": 8436, // AutoCrib
		"a4a1c2": 306,  // Ericsson
		"a4a1e4": 8437, // Innotube
		"a4a46b": 3334, // Huawei
		"a4a4d3": 8438, // Bluebank
		"a4a6a9": 67,   // Private
		"a4aafe": 3334, // Huawei
		"a4ac0f": 3334, // Huawei
		"a4ad00": 8439, // Ragsdale
		"a4b197": 545,  // Apple
		"a4b1c1": 421,  // Intel
		"a4b1e9": 6392, // Technicolor
		"a4b239": 2,    // Cisco
		"a4b2a7": 8440, // Adaxys
		"a4b439": 2,    // Cisco
		"a4b61e": 3334, // Huawei
		"a4b805": 545,  // Apple
		"a4ba76": 3334, // Huawei
		"a4badb": 102,  // Dell
		"a4bb6d": 102,  // Dell
		"a4bdc4": 3334, // Huawei
		"a4be2b": 3334, // Huawei
		"a4bf01": 421,  // Intel
		"a4c0e1": 1427, // Nintendo
		"a4c337": 545,  // Apple
		"a4c361": 545,  // Apple
		"a4c3f0": 421,  // Intel
		"a4c494": 421,  // Intel
		"a4c54e": 3334, // Huawei
		"a4c64f": 3334, // Huawei
		"a4c69a": 152,  // Samsung
		"a4c74b": 3334, // Huawei
		"a4c7de": 1640, // Cambridge
		"a4caa0": 3334, // Huawei
		"a4cc32": 8441, // Inficomm
		"a4ceda": 2639, // Arcadyan
		"a4cf12": 6376, // Espressif
		"a4d094": 2631, // Vivavis
		"a4d18c": 545,  // Apple
		"a4d1d1": 8442, // ECOtality
		"a4d1d2": 545,  // Apple
		"a4d73c": 46,   // Epson
		"a4d795": 4030, // Wingtech
		"a4d856": 8443, // Gimbal
		"a4d931": 545,  // Apple
		"a4d990": 152,  // Samsung
		"a4da3f": 8444, // Bionics
		"a4db30": 2873, // Liteon
		"a4dcbe": 3334, // Huawei
		"a4e11a": 826,  // Juniper
		"a4e31b": 457,  // Nokia
		"a4e32e": 1655, // Silicon
		"a4e4b8": 2220, // BlackBerry
		"a4e57c": 6376, // Espressif
		"a4e597": 8445, // Gessler
		"a4e731": 457,  // Nokia
		"a4e7e4": 8446, // Connex
		"a4e975": 545,  // Apple
		"a4e9a3": 8447, // Honest
		"a4ea8e": 185,  // Extreme
		"a4ebd3": 152,  // Samsung
		"a4ed4e": 125,  // ARRIS
		"a4ee57": 46,   // Epson
		"a4ef15": 7031, // AltoBeam
		"a4ef52": 8448, // Telewave
		"a4f1e8": 545,  // Apple
		"a4f33b": 3031, // ZTE
		"a4f465": 6281, // Itel
		"a4f4c2": 8374, // Vnpt
		"a4ff95": 457,  // Nokia
		"a8016d": 5242, // Aiwa
		"a80180": 8449, // IMAGO
		"a802db": 3031, // ZTE
		"a8032a": 6376, // Espressif
		"a80577": 8450, // Netlist
		"a80600": 152,  // Samsung
		"a80c0d": 2,    // Cisco
		"a80c63": 3334, // Huawei
		"a811fc": 125,  // ARRIS
		"a81306": 6369, // Vivo
		"a81374": 2153, // Panasonic
		"a8154d": 1595, // TP-Link
		"a81559": 8451, // Breathometer
		"a816b2": 869,  // LG
		"a816d0": 152,  // Samsung
		"a81b18": 8452, // XTS
		"a81d16": 3004, // Azurewave
		"a81e84": 3071, // Quanta
		"a82066": 545,  // Apple
		"a82316": 457,  // Nokia
		"a823fe": 869,  // LG
		"a824b8": 457,  // Nokia
		"a825eb": 1640, // Cambridge
		"a826d9": 1341, // HTC
		"a82bb5": 6294, // Edgecore
		"a82bb9": 152,  // Samsung
		"a82bcd": 3334, // Huawei
		"a830bc": 152,  // Samsung
		"a8329a": 6211, // Digicom
		"a8346a": 152,  // Samsung
		"a83512": 3334, // Huawei
		"a8367a": 8453, // frogblue
		"a83759": 3334, // Huawei
		"a83944": 2246, // Actiontec
		"a83b5c": 3334, // Huawei
		"a83ccb": 8454, // Rossma
		"a84025": 8455, // Oxide
		"a84041": 8456, // Dragino
		"a84397": 8457, // Innogrit
		"a84481": 457,  // Nokia
		"a845cd": 8458, // Siselectron
		"a845e9": 8459, // Firich
		"a848fa": 6376, // Espressif
		"a8494d": 3334, // Huawei
		"a849a5": 8460, // Lisantech
		"a84a28": 545,  // Apple
		"a84b4d": 152,  // Samsung
		"a84d4a": 8461, // Audiowise
		"a84e3f": 870,  // Hitron
		"a84fb1": 2,    // Cisco
		"a85081": 3334, // Huawei
		"a8515b": 152,  // Samsung
		"a8537d": 7484, // Mist
		"a854a2": 8462, // Heimgard
		"a854b2": 1592, // Wistron
		"a8574e": 1595, // TP-Link
		"a85840": 1640, // Cambridge
		"a8587c": 8463, // Shoogee
		"a85ae0": 3334, // Huawei
		"a85b6c": 1926, // Bosch
		"a85b78": 545,  // Apple
		"a85bb7": 545,  // Apple
		"a85bf3": 8464, // Audivo
		"a85bf7": 1685, // Aruba
		"a85c2c": 545,  // Apple
		"a85e45": 1806, // ASUS
		"a85ee4": 8465, // 12Sided
		"a860b6": 545,  // Apple
		"a8610a": 8466, // Arduino
		"a861aa": 8467, // Cloudview
		"a862a2": 8468, // Jiwumedia
		"a8637d": 803,  // D-Link
		"a863df": 8469, // Displaire
		"a864f1": 421,  // Intel
		"a8667f": 545,  // Apple
		"a8671e": 8470, // Ratp
		"a8698c": 11,   // Oracle
		"a86a6f": 4555, // RIM
		"a86abb": 2048, // Sagemcom
		"a86ac1": 8471, // HanbitEDS
		"a86d5f": 2050, // Raisecom
		"a86daa": 421,  // Intel
		"a86e4e": 3334, // Huawei
		"a8705d": 125,  // ARRIS
		"a870a5": 8472, // UniComm
		"a87285": 3225, // IDT
		"a87484": 3031, // ZTE
		"a875d6": 8473, // FreeTek
		"a875e2": 8474, // Aventura
		"a87650": 152,  // Samsung
		"a8776f": 8475, // Zonoff
		"a87b39": 457,  // Nokia
		"a87c01": 152,  // Samsung
		"a87d12": 3334, // Huawei
		"a87e33": 457,  // Nokia
		"a87eea": 421,  // Intel
		"a8817e": 545,  // Apple
		"a88195": 152,  // Samsung
		"a885d7": 8476, // Sangfor
		"a886dd": 545,  // Apple
		"a887b3": 152,  // Samsung
		"a88808": 545,  // Apple
		"a88940": 3334, // Huawei
		"a88c3e": 612,  // Microsoft
		"a88e24": 545,  // Apple
		"a8913d": 545,  // Apple
		"a8922c": 869,  // LG
		"a89675": 687,  // Motorola
		"a8968a": 545,  // Apple
		"a897cd": 125,  // ARRIS
		"a897dc": 372,  // IBM
		"a898c6": 8477, // Shinbo
		"a8995c": 8478, // aizo
		"a89969": 102,  // Dell
		"a89a93": 2048, // Sagemcom
		"a89ad7": 457,  // Nokia
		"a89b10": 8479, // inMotion
		"a89ca4": 8480, // Furrion
		"a89ced": 5697, // Xiaomi
		"a89d21": 2,    // Cisco
		"a89fba": 152,  // Samsung
		"a89fec": 125,  // ARRIS
		"a8a089": 8481, // Tactical
		"a8a159": 4719, // ASRock
		"a8a198": 6433, // TCT
		"a8a668": 3031, // ZTE
		"a8b088": 5820, // eero
		"a8b0ae": 8482, // Leoni
		"a8b13b": 302,  // HP
		"a8b1d4": 2,    // Cisco
		"a8b2da": 4,    // Fujitsu
		"a8b456": 2,    // Cisco
		"a8b57c": 1916, // Roku
		"a8b86e": 869,  // LG
		"a8b9b3": 6830, // Essys
		"a8bbcf": 545,  // Apple
		"a8bd27": 302,  // HP
		"a8bd3a": 3505, // Unionman
		"a8be27": 545,  // Apple
		"a8c092": 3334, // Huawei
		"a8c0ea": 8483, // Pepwave
		"a8c222": 8484, // TM-Research
		"a8c252": 3334, // Huawei
		"a8c266": 531,  // HUMAX
		"a8c83a": 3334, // Huawei
		"a8c87f": 8485, // Roqos
		"a8ca7b": 3334, // Huawei
		"a8cab9": 152,  // Samsung
		"a8ccc5": 2711, // Saab
		"a8ce90": 8486, // CVC
		"a8d081": 3334, // Huawei
		"a8d0e3": 5479, // Systech
		"a8d0e5": 826,  // Juniper
		"a8d3c8": 8487, // Topcon
		"a8d3f7": 2639, // Arcadyan
		"a8d4e0": 3334, // Huawei
		"a8d88a": 8488, // Wyconn
		"a8da0c": 8144, // Servercom
		"a8db03": 152,  // Samsung
		"a8e018": 457,  // Nokia
		"a8e3ee": 101,  // Sony
		"a8e539": 8489, // Moimstone
		"a8e544": 3334, // Huawei
		"a8e621": 5438, // Amazon
		"a8e705": 4921, // Fiberhome
		"a8e81e": 8490, // ATW
		"a8e824": 8491, // Inim
		"a8e978": 3334, // Huawei
		"a8eec6": 8492, // Muuselabs/SA
		"a8ef26": 8493, // Tritonwave
		"a8f266": 3334, // Huawei
		"a8f274": 152,  // Samsung
		"a8f5ac": 3334, // Huawei
		"a8f5dd": 125,  // ARRIS
		"a8f766": 8494, // ITE
		"a8f7e0": 4964, // PLANET
		"a8f94b": 8495, // Eltex
		"a8fad8": 545,  // Apple
		"a8fe9d": 545,  // Apple
		"a8ffba": 3334, // Huawei
		"ac00d0": 3031, // ZTE
		"ac02ca": 8496, // HI
		"ac02ef": 8497, // Comsis
		"ac0425": 8498, // ball-b
		"ac0613": 8499, // Senselogix
		"ac075f": 3334, // Huawei
		"ac0bfb": 6376, // Espressif
		"ac0d1b": 869,  // LG
		"ac1203": 421,  // Intel
		"ac139c": 202,  // Adtran
		"ac1461": 8500, // ATAW
		"ac14d2": 8501, // wi-daq
		"ac1585": 8502, // silergy
		"ac15f4": 545,  // Apple
		"ac162d": 302,  // HP
		"ac1826": 46,   // Epson
		"ac1d06": 545,  // Apple
		"ac1e92": 152,  // Samsung
		"ac1e9e": 5697, // Xiaomi
		"ac1f74": 545,  // Apple
		"ac202e": 870,  // Hitron
		"ac20aa": 8503, // DMATEK
		"ac2205": 358,  // Compal
		"ac220b": 1806, // ASUS
		"ac2316": 7484, // Mist
		"ac2334": 6295, // Infinix
		"ac293a": 545,  // Apple
		"ac2b6e": 421,  // Intel
		"ac2da3": 8504, // TXTR
		"ac2da9": 6265, // Tecno
		"ac2fa8": 8505, // Humannix
		"ac3328": 3334, // Huawei
		"ac35ee": 6440, // FN-Link
		"ac3613": 152,  // Samsung
		"ac3743": 1341, // HTC
		"ac37c9": 8506, // RAID
		"ac3870": 2664, // Lenovo
		"ac3a67": 2,    // Cisco
		"ac3a7a": 1916, // Roku
		"ac3b77": 2048, // Sagemcom
		"ac3c0b": 545,  // Apple
		"ac3c8e": 833,  // Flextronics
		"ac3cb4": 8507, // Nilan
		"ac4122": 8508, // Eclipse
		"ac4228": 8509, // Parta
		"ac4330": 4965, // Versa
		"ac44f2": 727,  // Yamaha
		"ac471b": 3334, // Huawei
		"ac4723": 8510, // Genelec
		"ac49db": 545,  // Apple
		"ac4a56": 2,    // Cisco
		"ac4a67": 2,    // Cisco
		"ac4b1e": 8511, // Integri-Sys.com
		"ac4bc8": 826,  // Juniper
		"ac4e91": 3334, // Huawei
		"ac4ffc": 8512, // SVS-VISTEK
		"ac5036": 6834, // Pi-Coral
		"ac5093": 153,  // Magna
		"ac512c": 6295, // Infinix
		"ac5135": 8513, // MPI
		"ac51ee": 1640, // Cambridge
		"ac562c": 647,  // Lava
		"ac567b": 5435, // Sunnovo
		"ac5a14": 152,  // Samsung
		"ac5afc": 421,  // Intel
		"ac5d5c": 6440, // FN-Link
		"ac5e14": 3334, // Huawei
		"ac5e8c": 8514, // Utillink
		"ac5f3e": 152,  // Samsung
		"ac5fea": 6923, // OnePlus
		"ac6089": 3334, // Huawei
		"ac60b6": 306,  // Ericsson
		"ac6123": 8515, // Drivven
		"ac6175": 3334, // Huawei
		"ac61b9": 8516, // WAMA
		"ac61ea": 545,  // Apple
		"ac63be": 5438, // Amazon
		"ac6417": 300,  // Siemens
		"ac6462": 3031, // ZTE
		"ac6490": 3334, // Huawei
		"ac64cf": 6440, // FN-Link
		"ac6706": 2737, // Ruckus
		"ac675d": 421,  // Intel
		"ac6784": 3521, // Google
		"ac67b2": 6376, // Espressif
		"ac6c90": 152,  // Samsung
		"ac6f4f": 8517, // Enspert
		"ac6fbb": 8518, // TATUNG
		"ac6fd9": 8519, // Valueplus
		"ac7236": 8520, // Lexking
		"ac7289": 421,  // Intel
		"ac74b1": 421,  // Intel
		"ac74c4": 8521, // Maytronics
		"ac751d": 3334, // Huawei
		"ac78d1": 826,  // Juniper
		"ac7a42": 8522, // iConnectivity
		"ac7a4d": 430,  // Alpsalpine
		"ac7a56": 2,    // Cisco
		"ac7ba1": 421,  // Intel
		"ac7e01": 3334, // Huawei
		"ac7e8a": 2,    // Cisco
		"ac7f3e": 545,  // Apple
		"ac80ae": 4921, // Fiberhome
		"ac80d6": 8523, // Hexatronic
		"ac8112": 1450, // Gemtek
		"ac81f3": 457,  // Nokia
		"ac8247": 421,  // Intel
		"ac83f0": 8524, // ImmediaTV
		"ac83f3": 4498, // AMPAK
		"ac84c6": 1595, // TP-Link
		"ac84c9": 2048, // Sagemcom
		"ac853d": 3334, // Huawei
		"ac87a3": 545,  // Apple
		"ac88fd": 545,  // Apple
		"ac8995": 3004, // Azurewave
		"ac8b9c": 8525, // Primera
		"ac8ba9": 2978, // Ubiquiti
		"ac8d14": 8526, // Smartrove
		"ac8d34": 3334, // Huawei
		"ac8ff8": 457,  // Nokia
		"ac9085": 545,  // Apple
		"ac9232": 3334, // Huawei
		"ac932f": 457,  // Nokia
		"ac9572": 8527, // Jovision
		"ac976c": 8528, // Greenliant
		"ac9929": 3334, // Huawei
		"ac9a96": 5298, // Maxlinear
		"ac9b0a": 101,  // Sony
		"ac9e17": 1806, // ASUS
		"aca016": 2,    // Cisco
		"aca22c": 8529, // Baycity
		"aca31e": 1685, // Aruba
		"aca88e": 3205, // Sharp
		"aca919": 8530, // TrekStor
		"aca9a0": 8531, // Audioengine
		"acabbf": 8532, // AthenTek
		"acae19": 1916, // Roku
		"acafb9": 152,  // Samsung
		"acb313": 125,  // ARRIS
		"acb3b5": 3334, // Huawei
		"acb57d": 2873, // Liteon
		"acb859": 8533, // Uniband
		"acbb61": 7450, // YSTen
		"acbc32": 545,  // Apple
		"acbcd9": 2,    // Cisco
		"acbd0b": 8534, // Leimac
		"acbd70": 3334, // Huawei
		"acbe75": 8535, // Ufine
		"acbeb6": 8536, // Visualedge
		"acc1ee": 5697, // Xiaomi
		"acc25d": 4921, // Fiberhome
		"acc33a": 152,  // Samsung
		"acc595": 8537, // Graphite
		"acc662": 6428, // MitraStar
		"acc73f": 8538, // Vitsmo
		"acc935": 8539, // Ness
		"acca54": 8540, // Telldus
		"acca8e": 8541, // ODA
		"accaba": 8542, // Midokura
		"accc8e": 5131, // Axis
		"accf23": 8543, // Hi-flying
		"accf5c": 545,  // Apple
		"accf85": 3334, // Huawei
		"acd074": 6376, // Espressif
		"acd364": 21,   // ABB
		"acd618": 6923, // OnePlus
		"acd9d6": 8544, // tci
		"acdb48": 125,  // ARRIS
		"acdcca": 3334, // Huawei
		"acde48": 67,   // Private
		"ace010": 2873, // Liteon
		"ace215": 3334, // Huawei
		"ace2d3": 302,  // HP
		"ace342": 3334, // Huawei
		"ace348": 8545, // MadgeTech
		"ace4b5": 545,  // Apple
		"ace5f0": 8546, // Doppler
		"ace87b": 3334, // Huawei
		"ace87e": 8547, // Bytemark
		"ace97f": 8548, // IoT
		"ace9aa": 8549, // Hay
		"aceb51": 3857, // Universal
		"acec80": 125,  // ARRIS
		"acec85": 5820, // eero
		"aced5c": 421,  // Intel
		"acee3b": 8550, // 6harmonics
		"acee9e": 152,  // Samsung
		"acf0b2": 8551, // Becker
		"acf108": 869,  // LG
		"acf1df": 803,  // D-Link
		"acf2c5": 2,    // Cisco
		"acf5e6": 2,    // Cisco
		"acf6f7": 869,  // LG
		"acf7f3": 5697, // Xiaomi
		"acf8cc": 125,  // ARRIS
		"acf970": 3334, // Huawei
		"acf97e": 8552, // Elesys
		"acfaa5": 8553, // digitron
		"acfdce": 421,  // Intel
		"acfdec": 545,  // Apple
		"acfe05": 6281, // Itel
		"b00073": 1592, // Wistron
		"b000b4": 2,    // Cisco
		"b00247": 4498, // AMPAK
		"b0027e": 8554, // Muller
		"b00594": 2873, // Liteon
		"b00875": 3334, // Huawei
		"b009d3": 8555, // Avizia
		"b009da": 6946, // Ring
		"b00ad5": 3031, // ZTE
		"b00cd1": 302,  // HP
		"b01266": 8556, // Futaba-Kikaku
		"b01408": 8557, // Lightspeed
		"b01656": 3334, // Huawei
		"b01886": 3455, // SmarDTV
		"b019c6": 545,  // Apple
		"b01c91": 8558, // Elim
		"b01f29": 8559, // Helvetia
		"b0227a": 302,  // HP
		"b02491": 3334, // Huawei
		"b024f3": 8560, // Progeny
		"b025aa": 67,   // Private
		"b02628": 858,  // Broadcom
		"b02680": 2,    // Cisco
		"b027cf": 185,  // Extreme
		"b02a1f": 4030, // Wingtech
		"b02a43": 3521, // Google
		"b03366": 6369, // Vivo
		"b033a6": 826,  // Juniper
		"b03495": 545,  // Apple
		"b0358d": 457,  // Nokia
		"b0359f": 421,  // Intel
		"b035b5": 545,  // Apple
		"b03795": 869,  // LG
		"b03956": 1368, // Netgear
		"b03ace": 3334, // Huawei
		"b03cdc": 421,  // Intel
		"b03e51": 3512, // BSkyB
		"b03eb0": 8561, // MICRODIA
		"b04089": 8562, // Senient
		"b0411d": 8563, // ITTIM
		"b0435d": 8564, // NuLEDs
		"b04502": 3334, // Huawei
		"b04519": 6433, // TCT
		"b04530": 3512, // BSkyB
		"b046fc": 6428, // MitraStar
		"b047bf": 152,  // Samsung
		"b0481a": 545,  // Apple
		"b0487a": 1595, // TP-Link
		"b04e26": 1595, // TP-Link
		"b04f13": 102,  // Dell
		"b0518e": 8565, // Holl
		"b05508": 3334, // Huawei
		"b05706": 8566, // Vallox
		"b05ada": 302,  // HP
		"b05b67": 3334, // Huawei
		"b05c16": 4921, // Fiberhome
		"b05cda": 302,  // HP
		"b05ce5": 457,  // Nokia
		"b05dd4": 125,  // ARRIS
		"b06088": 421,  // Intel
		"b061c7": 8567, // Ericsson-LG
		"b065bd": 545,  // Apple
		"b06a41": 3521, // Google
		"b06ebf": 1806, // ASUS
		"b06fe0": 152,  // Samsung
		"b0700d": 457,  // Nokia
		"b0702d": 545,  // Apple
		"b072bf": 2056, // Murata
		"b0735d": 3334, // Huawei
		"b0739c": 5438, // Amazon
		"b0754d": 457,  // Nokia
		"b075d5": 3031, // ZTE
		"b0761b": 3334, // Huawei
		"b077ac": 125,  // ARRIS
		"b07870": 8568, // Wi-NEXT
		"b07908": 8569, // Cummings
		"b0793c": 8570, // Revolv
		"b07994": 687,  // Motorola
		"b07b25": 102,  // Dell
		"b07d47": 2,    // Cisco
		"b07d64": 421,  // Intel
		"b07fb9": 1368, // Netgear
		"b081d8": 8571, // I-Sys
		"b083d6": 125,  // ARRIS
		"b083fe": 102,  // Dell
		"b08900": 3334, // Huawei
		"b08991": 7468, // LGE
		"b089c2": 8572, // Zyptonite
		"b08b92": 3031, // ZTE
		"b08bcf": 2,    // Cisco
		"b08bd0": 2,    // Cisco
		"b08c75": 545,  // Apple
		"b08e1a": 8573, // URadio
		"b09074": 8574, // Fulan
		"b0907e": 2,    // Cisco
		"b09134": 8575, // Taleo
		"b0935b": 125,  // ARRIS
		"b09575": 1595, // TP-Link
		"b0958e": 1595, // TP-Link
		"b0966c": 8576, // Lanbowan
		"b0973a": 8577, // E-Fuel
		"b0982b": 2048, // Sagemcom
		"b0989f": 869,  // LG
		"b098bc": 3334, // Huawei
		"b09928": 4,    // Fujitsu
		"b099d7": 152,  // Samsung
		"b09fba": 545,  // Apple
		"b0a10a": 5064, // Pivotal
		"b0a454": 8578, // Tripwire
		"b0a460": 421,  // Intel
		"b0a4f0": 3334, // Huawei
		"b0a651": 2,    // Cisco
		"b0a6f5": 8579, // Xaptum
		"b0a737": 1916, // Roku
		"b0a7b9": 1595, // TP-Link
		"b0a86e": 826,  // Juniper
		"b0aa77": 2,    // Cisco
		"b0acd2": 3031, // ZTE
		"b0acfa": 4,    // Fujitsu
		"b0adaa": 620,  // Avaya
		"b0ae25": 8580, // Varikorea
		"b0b194": 3031, // ZTE
		"b0b28f": 2048, // Sagemcom
		"b0b2dc": 2692, // ZyXEL
		"b0b3ad": 531,  // HUMAX
		"b0b5e8": 8581, // Ruroc
		"b0b867": 302,  // HP
		"b0b98a": 1368, // Netgear
		"b0bb8b": 8582, // Wavetel
		"b0bbe5": 2048, // Sagemcom
		"b0be76": 1595, // TP-Link
		"b0be83": 545,  // Apple
		"b0bf99": 8583, // Wizitdongdo
		"b0c090": 7300, // Chicony
		"b0c19e": 3031, // ZTE
		"b0c205": 8584, // Bionime
		"b0c287": 6392, // Technicolor
		"b0c387": 8585, // GOEFER
		"b0c46c": 8586, // Senseit
		"b0c4e7": 152,  // Samsung
		"b0c53c": 2,    // Cisco
		"b0c554": 803,  // D-Link
		"b0c559": 152,  // Samsung
		"b0c69a": 826,  // Juniper
		"b0c745": 1077, // Buffalo
		"b0c787": 3334, // Huawei
		"b0ca68": 545,  // Apple
		"b0ccfe": 3334, // Huawei
		"b0d09c": 152,  // Samsung
		"b0d2f5": 8587, // Vello
		"b0d7c5": 8588, // Logipix
		"b0d7cc": 8589, // Tridonic
		"b0d888": 2153, // Panasonic
		"b0da00": 8590, // Cera
		"b0daf9": 125,  // ARRIS
		"b0dd74": 8462, // Heimgard
		"b0de28": 545,  // Apple
		"b0df3a": 152,  // Samsung
		"b0dfc1": 6266, // Tenda
		"b0e03c": 6433, // TCT
		"b0e17e": 3334, // Huawei
		"b0e235": 5697, // Xiaomi
		"b0e2e5": 4921, // Fiberhome
		"b0e4d5": 3521, // Google
		"b0e50e": 8591, // NRG
		"b0e5ed": 3334, // Huawei
		"b0e5f9": 545,  // Apple
		"b0e754": 1939, // 2Wire
		"b0e892": 46,   // Epson
		"b0e9fe": 8592, // Woan
		"b0eabc": 2544, // Askey
		"b0eb57": 3334, // Huawei
		"b0ec71": 152,  // Samsung
		"b0ecdd": 3334, // Huawei
		"b0ece1": 67,   // Private
		"b0ee45": 3004, // Azurewave
		"b0ee7b": 1916, // Roku
		"b0f1a3": 8593, // Fengfan
		"b0f1d8": 545,  // Apple
		"b0f1ec": 4498, // AMPAK
		"b0f530": 870,  // Hitron
		"b0f7c4": 5438, // Amazon
		"b0faeb": 2,    // Cisco
		"b0fc0d": 5438, // Amazon
		"b0fc36": 189,  // CyberTAN
		"b0febd": 67,   // Private
		"b0fee5": 3334, // Huawei
		"b4009c": 8594, // CableWorld
		"b40216": 2,    // Cisco
		"b4055d": 7721, // Inspur
		"b407f9": 152,  // Samsung
		"b40832": 8595, // TC
		"b40931": 3334, // Huawei
		"b40ac6": 8596, // DEXON
		"b40b44": 8597, // Smartisan
		"b40e96": 8598, // Heran
		"b40edc": 8599, // LG-Ericsson
		"b40ede": 421,  // Intel
		"b40f3b": 6266, // Tenda
		"b40fb3": 6369, // Vivo
		"b41489": 2,    // Cisco
		"b414e6": 3334, // Huawei
		"b41513": 3334, // Huawei
		"b4157e": 8600, // Celona
		"b417a8": 7234, // Facebook
		"b418d1": 545,  // Apple
		"b41a1d": 152,  // Samsung
		"b41bb0": 545,  // Apple
		"b41c30": 3031, // ZTE
		"b41cab": 8601, // ICR
		"b41def": 8602, // Internet
		"b42046": 5820, // eero
		"b42200": 3230, // Brother
		"b42330": 1109, // Itron
		"b424e7": 8603, // Codetek
		"b42875": 8604, // Futecho
		"b428f1": 8605, // E-Prime
		"b42a0e": 6392, // Technicolor
		"b42d56": 185,  // Extreme
		"b42e99": 1929, // Giga-Byte
		"b42ef8": 8606, // Eline
		"b43052": 3334, // Huawei
		"b431b8": 8607, // Aviwest
		"b436d1": 5695, // Renesas
		"b43741": 8608, // Consert
		"b437d8": 803,  // D-Link
		"b43a28": 152,  // Samsung
		"b43ae2": 3334, // Huawei
		"b43d08": 8609, // GX
		"b43e3b": 8610, // Viableware
		"b440a4": 545,  // Apple
		"b4430d": 8611, // Broadlink
		"b44326": 3334, // Huawei
		"b44506": 102,  // Dell
		"b4475e": 620,  // Avaya
		"b447f5": 6498, // Earda
		"b44bd2": 545,  // Apple
		"b45062": 8612, // EmBestor
		"b451f9": 8613, // NB
		"b45253": 732,  // Seagate
		"b4527d": 101,  // Sony
		"b4527e": 101,  // Sony
		"b45570": 8614, // Borea
		"b456b9": 8615, // Teraspek
		"b456e3": 545,  // Apple
		"b45861": 8616, // CRemote
		"b45d50": 1685, // Aruba
		"b45f84": 3031, // ZTE
		"b4608c": 4921, // Fiberhome
		"b461ff": 8617, // Lumigon
		"b46238": 8618, // Exablox
		"b46293": 152,  // Samsung
		"b46921": 421,  // Intel
		"b4695f": 6433, // TCT
		"b46bfc": 421,  // Intel
		"b46c47": 2153, // Panasonic
		"b46d83": 421,  // Intel
		"b46e08": 3334, // Huawei
		"b47443": 152,  // Samsung
		"b47447": 8619, // CoreOS
		"b4749f": 2544, // Askey
		"b4750e": 2468, // Belkin
		"b47947": 7340, // Nutanix
		"b479a7": 152,  // Samsung
		"b479c8": 2737, // Ruckus
		"b47af1": 302,  // HP
		"b47c9c": 5438, // Amazon
		"b481bf": 8620, // Meta-Networks
		"b48255": 8621, // Research
		"b482c5": 8622, // Relay2
		"b482fe": 2544, // Askey
		"b485e1": 545,  // Apple
		"b48655": 3334, // Huawei
		"b48901": 3334, // Huawei
		"b48a5f": 826,  // Juniper
		"b48b19": 545,  // Apple
		"b48c9d": 3004, // Azurewave
		"b4944e": 8623, // WeTelecom
		"b49691": 421,  // Intel
		"b49842": 3031, // ZTE
		"b499ba": 302,  // HP
		"b49cdf": 545,  // Apple
		"b49d02": 152,  // Samsung
		"b49d0b": 7283, // BQ
		"b49db4": 3918, // Axion
		"b49ee6": 8624, // Shenzhen
		"b4a25c": 665,  // Cambium
		"b4a4e3": 2,    // Cisco
		"b4a5a9": 8625, // MODI
		"b4a5ef": 2074, // Sercomm
		"b4a898": 3334, // Huawei
		"b4a8b9": 2,    // Cisco
		"b4a94f": 4253, // Mercury
		"b4a95a": 620,  // Avaya
		"b4a984": 1989, // Symantec
		"b4a9fc": 3071, // Quanta
		"b4a9fe": 8626, // GHIA
		"b4aa4d": 8627, // Ensequence
		"b4ab2c": 8628, // MtM
		"b4ae2b": 612,  // Microsoft
		"b4b017": 620,  // Avaya
		"b4b024": 1595, // TP-Link
		"b4b055": 3334, // Huawei
		"b4b15a": 300,  // Siemens
		"b4b291": 869,  // LG
		"b4b362": 3031, // ZTE
		"b4b52f": 302,  // HP
		"b4b5af": 8629, // Minsung
		"b4b676": 421,  // Intel
		"b4b686": 302,  // HP
		"b4b742": 5438, // Amazon
		"b4b859": 1203, // Texa
		"b4b88d": 8630, // Thuh
		"b4ba02": 8631, // Agatel
		"b4bff6": 152,  // Samsung
		"b4c26a": 797,  // Garmin
		"b4c4fc": 5697, // Xiaomi
		"b4c6f8": 8632, // Axilspot
		"b4c799": 185,  // Extreme
		"b4cc04": 8633, // Piranti
		"b4cce9": 8634, // Prosyst
		"b4cd27": 3334, // Huawei
		"b4ce40": 152,  // Samsung
		"b4cef6": 1341, // HTC
		"b4d135": 8635, // Cloudistics
		"b4d286": 7075, // Telechips
		"b4d5bd": 421,  // Intel
		"b4d64e": 8636, // Caldero
		"b4d8a9": 8637, // BetterBots
		"b4dd15": 8638, // ControlThings
		"b4de31": 2,    // Cisco
		"b4dedf": 3031, // ZTE
		"b4df3b": 8639, // Chromlech
		"b4dffa": 8640, // Litemax
		"b4e01d": 8641, // Conception
		"b4e0cd": 8642, // Fusion-io
		"b4e10f": 102,  // Dell
		"b4e1c4": 612,  // Microsoft
		"b4e1eb": 67,   // Private
		"b4e3f9": 1655, // Silicon
		"b4e454": 5438, // Amazon
		"b4e62a": 869,  // LG
		"b4e62d": 6376, // Espressif
		"b4e782": 8643, // Vivalnk
		"b4e8c9": 8644, // XADA
		"b4e9b0": 2,    // Cisco
		"b4ec02": 430,  // Alpsalpine
		"b4ed54": 8645, // Wohler
		"b4eeb4": 2544, // Askey
		"b4ef39": 152,  // Samsung
		"b4f0ab": 545,  // Apple
		"b4f18c": 3334, // Huawei
		"b4f1da": 869,  // LG
		"b4f267": 358,  // Compal
		"b4f2e8": 125,  // ARRIS
		"b4f323": 8646, // Petatel
		"b4f58e": 3334, // Huawei
		"b4f61c": 545,  // Apple
		"b4f7a1": 869,  // LG
		"b4f81e": 8647, // Kinova
		"b4f949": 8648, // optilink
		"b4fa48": 545,  // Apple
		"b4fbe3": 7031, // AltoBeam
		"b4fbe4": 2978, // Ubiquiti
		"b4fbf9": 3334, // Huawei
		"b4fc75": 8649, // SEMA
		"b4ff98": 3334, // Huawei
		"b80018": 8650, // Htel
		"b802a4": 8651, // Aeonsemi
		"b80305": 421,  // Intel
		"b805ab": 3031, // ZTE
		"b80716": 6369, // Vivo
		"b808cf": 421,  // Intel
		"b808d7": 3334, // Huawei
		"b8098a": 545,  // Apple
		"b810d4": 6362, // Masimo
		"b8114b": 2,    // Cisco
		"b812da": 8026, // Lvswitches
		"b81332": 4498, // AMPAK
		"b8145c": 3334, // Huawei
		"b814db": 6577, // Ohsung
		"b81619": 125,  // ARRIS
		"b817c2": 545,  // Apple
		"b8186f": 8652, // Oriental
		"b81999": 8653, // Nesys
		"b81daa": 869,  // LG
		"b81f5e": 8654, // Apption
		"b8208e": 2153, // Panasonic
		"b8259a": 8655, // Thalmic
		"b825b5": 8656, // Trakm8
		"b827c5": 3334, // Huawei
		"b829f7": 8657, // Blaster
		"b82a72": 102,  // Dell
		"b82aa9": 545,  // Apple
		"b82b68": 3334, // Huawei
		"b82ca0": 5986, // Resideo
		"b82d28": 4498, // AMPAK
		"b831b5": 612,  // Microsoft
		"b836d8": 8658, // Videoswitch
		"b8374a": 545,  // Apple
		"b83861": 2,    // Cisco
		"b83a08": 6266, // Tenda
		"b83a5a": 1685, // Aruba
		"b83a7b": 8659, // Worldplay
		"b83a9d": 3856, // Alarm.com
		"b83bcc": 5697, // Xiaomi
		"b83e59": 1916, // Roku
		"b83fd2": 432,  // Mellanox
		"b8415f": 2871, // ASP
		"b841a4": 545,  // Apple
		"b843e4": 8660, // Vlatacom
		"b844ae": 6433, // TCT
		"b844d9": 545,  // Apple
		"b8477a": 6067, // Dasan
		"b847c6": 8661, // SanJet
		"b84fd5": 612,  // Microsoft
		"b85001": 185,  // Extreme
		"b853ac": 545,  // Apple
		"b85510": 2127, // Zioncom
		"b85600": 3334, // Huawei
		"b856bd": 1178, // ITT
		"b85776": 8662, // lignex1
		"b857d8": 152,  // Samsung
		"b85810": 8663, // Numera
		"b8599f": 432,  // Mellanox
		"b859ce": 6498, // Earda
		"b85a73": 152,  // Samsung
		"b85af7": 8664, // Ouya
		"b85afe": 8665, // Handaer
		"b85d0a": 545,  // Apple
		"b85dc3": 3334, // Huawei
		"b85e7b": 152,  // Samsung
		"b85f98": 5438, // Amazon
		"b85fb0": 3334, // Huawei
		"b8616f": 145,  // Accton
		"b8621f": 2,    // Cisco
		"b8634d": 545,  // Apple
		"b863bc": 8666, // ROBOTIS
		"b8653b": 8667, // Bolymin
		"b86685": 2048, // Sagemcom
		"b869c2": 3939, // Sunitec
		"b869f4": 1784, // Routerboard.com
		"b86a97": 6294, // Edgecore
		"b86b23": 35,   // Toshiba
		"b86ce8": 152,  // Samsung
		"b870f4": 358,  // Compal
		"b87447": 8668, // Convergence
		"b875c0": 8669, // PayPal
		"b87826": 1427, // Nintendo
		"b8782e": 545,  // Apple
		"b87ac9": 300,  // Siemens
		"b87bc5": 545,  // Apple
		"b87cf2": 185,  // Extreme
		"b87ee5": 3541, // Intelbras
		"b88198": 421,  // Intel
		"b881fa": 545,  // Apple
		"b88303": 302,  // HP
		"b8857b": 3334, // Huawei
		"b88584": 102,  // Dell
		"b88687": 2873, // Liteon
		"b8876e": 8670, // Yandex
		"b887c6": 8671, // Prudential
		"b888e3": 358,  // Compal
		"b88a60": 421,  // Intel
		"b88a72": 5695, // Renesas
		"b88aec": 1427, // Nintendo
		"b88d12": 545,  // Apple
		"b88e82": 3334, // Huawei
		"b88ec6": 8672, // Stateless
		"b88edf": 8673, // Zencheer
		"b88f14": 8674, // Analytica
		"b89047": 545,  // Apple
		"b891c9": 3621, // Handreamnet
		"b89436": 3334, // Huawei
		"b89470": 374,  // Calix
		"b894e7": 5697, // Xiaomi
		"b89674": 8675, // AllDSP
		"b898ad": 687,  // Motorola
		"b898b0": 8676, // Atlona
		"b89919": 8677, // 7signal
		"b899b0": 8678, // Cohere
		"b89a2a": 421,  // Intel
		"b89aed": 8679, // OceanServer
		"b89bc9": 741,  // SMC
		"b89ea6": 8680, // Spbec-Mining
		"b89f09": 1592, // Wistron
		"b8a14a": 2050, // Raisecom
		"b8a175": 1916, // Roku
		"b8a377": 2,    // Cisco
		"b8a386": 803,  // D-Link
		"b8a3e0": 8681, // BenRui
		"b8a44f": 5131, // Axis
		"b8ac6f": 102,  // Dell
		"b8ad3e": 8682, // Bluecom
		"b8ae6e": 1427, // Nintendo
		"b8aeed": 1115, // Elitegroup
		"b8af67": 302,  // HP
		"b8b1c7": 8683, // BT
		"b8b2eb": 8684, // Googol
		"b8b2f8": 545,  // Apple
		"b8b3dc": 8685, // Derek
		"b8b7d7": 8686, // 2GIG
		"b8b7f1": 1592, // Wistron
		"b8b81e": 421,  // Intel
		"b8ba72": 8687, // Cynove
		"b8bb6d": 8688, // ENERES
		"b8bbaf": 152,  // Samsung
		"b8bc1b": 3334, // Huawei
		"b8bc5b": 152,  // Samsung
		"b8bd79": 8689, // TrendPoint
		"b8bebf": 2,    // Cisco
		"b8bef4": 1637, // devolo
		"b8bf83": 421,  // Intel
		"b8c111": 545,  // Apple
		"b8c227": 8690, // PSTec
		"b8c253": 826,  // Juniper
		"b8c385": 3334, // Huawei
		"b8c46f": 8691, // Primmcon
		"b8c68e": 152,  // Samsung
		"b8c6aa": 6498, // Earda
		"b8c716": 4921, // Fiberhome
		"b8c75d": 545,  // Apple
		"b8c8eb": 6281, // Itel
		"b8ca3a": 102,  // Dell
		"b8cb29": 102,  // Dell
		"b8cd93": 8692, // Penetek
		"b8cda7": 8693, // Maxeler
		"b8cef6": 432,  // Mellanox
		"b8d309": 8694, // Cox
		"b8d43e": 6369, // Vivo
		"b8d4e7": 1685, // Aruba
		"b8d50b": 3939, // Sunitec
		"b8d526": 2692, // ZyXEL
		"b8d56b": 8695, // Mirka
		"b8d6f6": 3334, // Huawei
		"b8d7af": 2056, // Murata
		"b8d94d": 2048, // Sagemcom
		"b8d9ce": 152,  // Samsung
		"b8dae8": 3334, // Huawei
		"b8dc87": 8696, // IAI
		"b8dd71": 3031, // ZTE
		"b8df6b": 8697, // SpotCam
		"b8e3b1": 3334, // Huawei
		"b8e3ee": 3857, // Universal
		"b8e589": 8698, // Payter
		"b8e60c": 545,  // Apple
		"b8e625": 1939, // 2Wire
		"b8e779": 8699, // 9Solutions
		"b8e856": 545,  // Apple
		"b8e937": 2047, // Sonos
		"b8eaaa": 600,  // ICG
		"b8eca3": 2692, // ZyXEL
		"b8ee0e": 2048, // Sagemcom
		"b8ee65": 2873, // Liteon
		"b8ee79": 8700, // YWire
		"b8f009": 6376, // Espressif
		"b8f080": 8701, // SPS
		"b8f12a": 545,  // Apple
		"b8f255": 3857, // Universal
		"b8f5e7": 8702, // WayTools
		"b8f6b1": 545,  // Apple
		"b8f732": 8703, // Aryaka
		"b8f74a": 8704, // Rcntec
		"b8f853": 2639, // Arcadyan
		"b8f883": 1595, // TP-Link
		"b8f8be": 8682, // Bluecom
		"b8f934": 101,  // Sony
		"b8ff61": 545,  // Apple
		"b8ffb3": 6428, // MitraStar
		"bc0543": 621,  // AVM
		"bc062d": 7186, // Wacom
		"bc091b": 421,  // Intel
		"bc0963": 545,  // Apple
		"bc0f64": 421,  // Intel
		"bc0f9a": 803,  // D-Link
		"bc0fa7": 8705, // Ouster
		"bc1401": 870,  // Hitron
		"bc1485": 152,  // Samsung
		"bc14ef": 8706, // ITON
		"bc1665": 2,    // Cisco
		"bc1695": 3031, // ZTE
		"bc16f5": 2,    // Cisco
		"bc17b8": 421,  // Intel
		"bc1a67": 8707, // YF
		"bc1ae4": 3334, // Huawei
		"bc1d89": 687,  // Motorola
		"bc1e85": 3334, // Huawei
		"bc20a4": 152,  // Samsung
		"bc20ba": 7721, // Inspur
		"bc2228": 803,  // D-Link
		"bc22fb": 2855, // RF
		"bc25e0": 3334, // Huawei
		"bc2643": 8708, // Elprotronic
		"bc26c7": 2,    // Cisco
		"bc282c": 8709, // e-Smart
		"bc2ce6": 2,    // Cisco
		"bc2d98": 8710, // ThinGlobal
		"bc2e48": 125,  // ARRIS
		"bc2ef6": 3334, // Huawei
		"bc2f3d": 6369, // Vivo
		"bc305b": 102,  // Dell
		"bc307d": 1592, // Wistron
		"bc307e": 1592, // Wistron
		"bc30d9": 2639, // Arcadyan
		"bc3329": 101,  // Sony
		"bc33ac": 1655, // Silicon
		"bc35e5": 8711, // Hydro
		"bc3865": 8712, // Jwcnetworks
		"bc38d2": 8713, // Pandachip
		"bc39d9": 8714, // Z-TEC
		"bc3baf": 545,  // Apple
		"bc3d85": 3334, // Huawei
		"bc3e07": 870,  // Hitron
		"bc3e13": 8715, // Accordance
		"bc3ecb": 6369, // Vivo
		"bc3f4e": 8716, // Teleepoch
		"bc3f8f": 3334, // Huawei
		"bc4100": 8717, // CoDACO
		"bc428c": 430,  // Alpsalpine
		"bc4486": 152,  // Samsung
		"bc44b0": 8718, // Elastifile
		"bc455b": 152,  // Samsung
		"bc4699": 1595, // TP-Link
		"bc4760": 152,  // Samsung
		"bc4a56": 2,    // Cisco
		"bc4b79": 8719, // SensingTek
		"bc4cc4": 545,  // Apple
		"bc4dfb": 870,  // Hitron
		"bc4e5d": 8720, // ZhongMiao
		"bc51fe": 8721, // Swann
		"bc52b4": 457,  // Nokia
		"bc52b7": 545,  // Apple
		"bc542f": 421,  // Intel
		"bc5436": 545,  // Apple
		"bc5451": 152,  // Samsung
		"bc54f9": 8722, // Drogoo
		"bc5a56": 2,    // Cisco
		"bc5bd5": 125,  // ARRIS
		"bc5c4c": 5693, // Elecom
		"bc5ea1": 8723, // PsiKick
		"bc5ff4": 4719, // ASRock
		"bc5ff6": 4253, // Mercury
		"bc60a7": 101,  // Sony
		"bc6193": 5697, // Xiaomi
		"bc620e": 3334, // Huawei
		"bc644b": 125,  // ARRIS
		"bc671c": 2,    // Cisco
		"bc6778": 545,  // Apple
		"bc6784": 8724, // Environics
		"bc69cb": 2153, // Panasonic
		"bc6a16": 8725, // tdvine
		"bc6a44": 1842, // Commend
		"bc6ad1": 5697, // Xiaomi
		"bc6b4d": 457,  // Nokia
		"bc6c21": 545,  // Apple
		"bc6d05": 8726, // Dusun
		"bc6e64": 101,  // Sony
		"bc6ee2": 421,  // Intel
		"bc71c1": 8727, // XTrillion
		"bc72b1": 152,  // Samsung
		"bc7536": 430,  // Alpsalpine
		"bc7574": 3334, // Huawei
		"bc765e": 152,  // Samsung
		"bc7670": 3334, // Huawei
		"bc76c5": 3334, // Huawei
		"bc7737": 421,  // Intel
		"bc779f": 8728, // SBM
		"bc79ad": 152,  // Samsung
		"bc7abf": 152,  // Samsung
		"bc7e8b": 152,  // Samsung
		"bc7f7b": 3334, // Huawei
		"bc7fa4": 5697, // Xiaomi
		"bc811f": 8729, // Ingate
		"bc8199": 8730, // BASIC
		"bc8385": 612,  // Microsoft
		"bc851f": 152,  // Samsung
		"bc8893": 8731, // VILLBAU
		"bc8aa3": 8732, // NHN
		"bc8ccd": 152,  // Samsung
		"bc8d0e": 457,  // Nokia
		"bc903a": 1926, // Bosch
		"bc91b5": 6295, // Infinix
		"bc926b": 545,  // Apple
		"bc9789": 3334, // Huawei
		"bc97e1": 858,  // Broadcom
		"bc9889": 4921, // Fiberhome
		"bc98df": 687,  // Motorola
		"bc9911": 2692, // ZyXEL
		"bc9930": 3334, // Huawei
		"bc99bc": 8733, // FonSee
		"bc9a53": 3334, // Huawei
		"bc9b68": 6392, // Technicolor
		"bc9c31": 3334, // Huawei
		"bc9da5": 8734, // DASCOM
		"bc9fe4": 1685, // Aruba
		"bc9fef": 545,  // Apple
		"bca13a": 8735, // SES-imagotag
		"bca4e1": 8736, // Nabto
		"bca511": 1368, // Netgear
		"bca58b": 152,  // Samsung
		"bca5a9": 545,  // Apple
		"bca8a6": 421,  // Intel
		"bca920": 545,  // Apple
		"bca993": 665,  // Cambium
		"bca9d6": 8737, // Cyber-Rain
		"bcadab": 620,  // Avaya
		"bcaec5": 1806, // ASUS
		"bcaf87": 8738, // smartAC.com
		"bcb0e7": 3334, // Huawei
		"bcb181": 3205, // Sharp
		"bcb1f3": 152,  // Samsung
		"bcb22b": 8739, // EM-Tech
		"bcb852": 8740, // Cybera
		"bcb863": 545,  // Apple
		"bcbae1": 8741, // AREC
		"bcbd9e": 6281, // Itel
		"bcc00f": 4921, // Fiberhome
		"bcc342": 2153, // Panasonic
		"bcc493": 2,    // Cisco
		"bcc6db": 457,  // Nokia
		"bccab5": 125,  // ARRIS
		"bccd45": 8742, // Voismart
		"bcce25": 1427, // Nintendo
		"bccf4f": 2692, // ZyXEL
		"bccfcc": 1341, // HTC
		"bcd074": 545,  // Apple
		"bcd11f": 152,  // Samsung
		"bcd177": 1595, // TP-Link
		"bcd206": 3334, // Huawei
		"bcd295": 2,    // Cisco
		"bcd5b6": 8743, // d2d
		"bcd713": 8744, // Owl
		"bcd767": 67,   // Private
		"bcd7a5": 1685, // Aruba
		"bcd7d4": 1916, // Roku
		"bcd940": 8745, // ASR
		"bcddc2": 6376, // Espressif
		"bcdf58": 3521, // Google
		"bce09d": 8746, // Eoslink
		"bce143": 545,  // Apple
		"bce265": 3334, // Huawei
		"bce59f": 8747, // WATERWORLD
		"bce63f": 152,  // Samsung
		"bce67c": 665,  // Cambium
		"bce712": 2,    // Cisco
		"bce92f": 302,  // HP
		"bcea2b": 2824, // CityCom
		"bceafa": 302,  // HP
		"bcec5d": 545,  // Apple
		"bceca0": 358,  // Compal
		"bcee7b": 1806, // ASUS
		"bcf171": 421,  // Intel
		"bcf1f2": 2,    // Cisco
		"bcf292": 542,  // Plantronics
		"bcf2af": 1637, // devolo
		"bcf310": 185,  // Extreme
		"bcf45f": 3031, // ZTE
		"bcf5ac": 869,  // LG
		"bcf685": 803,  // D-Link
		"bcf9f2": 8748, // Teko
		"bcfe8c": 8749, // Altronic
		"bcfed9": 545,  // Apple
		"bcff4d": 6376, // Espressif
		"bcffac": 8487, // Topcon
		"bcffeb": 687,  // Motorola
		"c00380": 826,  // Juniper
		"c005c2": 125,  // ARRIS
		"c0060c": 3334, // Huawei
		"c006c3": 1595, // TP-Link
		"c0074a": 8750, // Brita
		"c00d7e": 8751, // Additech
		"c01173": 152,  // Samsung
		"c011a6": 8752, // Fort-Telecom
		"c014b8": 457,  // Nokia
		"c014fe": 2,    // Cisco
		"c0174d": 152,  // Samsung
		"c01803": 302,  // HP
		"c01850": 3071, // Quanta
		"c01ada": 545,  // Apple
		"c01e9b": 8753, // Pixavi
		"c02250": 8754, // Koss
		"c0238d": 152,  // Samsung
		"c02506": 621,  // AVM
		"c0255c": 2,    // Cisco
		"c02567": 8755, // Nexxt
		"c025a5": 102,  // Dell
		"c025e9": 1595, // TP-Link
		"c0288d": 4079, // Logitech
		"c02973": 8756, // Audyssey
		"c029f3": 8757, // XySystem
		"c02b31": 8270, // Phytium
		"c02c5c": 545,  // Apple
		"c02dee": 8758, // Cuff
		"c02e26": 67,   // Private
		"c02ff1": 8759, // Volta
		"c0335e": 612,  // Microsoft
		"c034b4": 8760, // Gigastone
		"c03580": 8761, // A&R
		"c035c5": 1954, // Prosoft
		"c03653": 5820, // eero
		"c03656": 4921, // Fiberhome
		"c038f9": 457,  // Nokia
		"c03c04": 2048, // Sagemcom
		"c03c59": 421,  // Intel
		"c03d03": 152,  // Samsung
		"c03dd9": 6428, // MitraStar
		"c03e0f": 3512, // BSkyB
		"c03eba": 102,  // Dell
		"c03f0e": 1368, // Netgear
		"c03f2a": 8762, // Biscotti
		"c03fd5": 1115, // Elitegroup
		"c03fdd": 3334, // Huawei
		"c04004": 8763, // Medicaroid
		"c041f6": 869,  // LG
		"c042d0": 826,  // Juniper
		"c04301": 8764, // Epec
		"c04442": 545,  // Apple
		"c04754": 6369, // Vivo
		"c048e6": 152,  // Samsung
		"c049ef": 6376, // Espressif
		"c04a00": 1595, // TP-Link
		"c04b13": 8765, // WonderSound
		"c04df7": 8766, // Serelec
		"c05627": 2468, // Belkin
		"c057bc": 620,  // Avaya
		"c058a7": 456,  // Pico
		"c06118": 1595, // TP-Link
		"c0626b": 2,    // Cisco
		"c06369": 8767, // Binxin
		"c06394": 545,  // Apple
		"c064c6": 457,  // Nokia
		"c064e4": 2,    // Cisco
		"c06599": 152,  // Samsung
		"c067af": 2,    // Cisco
		"c06b55": 687,  // Motorola
		"c06c6d": 8768, // MagneMotion
		"c07009": 3334, // Huawei
		"c074ad": 1683, // Grandstream
		"c07831": 3334, // Huawei
		"c07878": 833,  // Flextronics
		"c07bbc": 2,    // Cisco
		"c07cd1": 5439, // Pegatron
		"c0830a": 1939, // 2Wire
		"c083c9": 3334, // Huawei
		"c0847a": 545,  // Apple
		"c0847d": 4498, // AMPAK
		"c08488": 8769, // Finis
		"c0854c": 6839, // Ragentek
		"c087eb": 152,  // Samsung
		"c0885b": 8770, // SnD
		"c0886d": 8771, // Securosys
		"c08997": 152,  // Samsung
		"c089ab": 125,  // ARRIS
		"c08ade": 2737, // Ruckus
		"c08b05": 3334, // Huawei
		"c08c60": 2,    // Cisco
		"c08c71": 687,  // Motorola
		"c09296": 3031, // ZTE
		"c09435": 125,  // ARRIS
		"c094ad": 3031, // ZTE
		"c09727": 152,  // Samsung
		"c09879": 143,  // Acer
		"c09ad0": 545,  // Apple
		"c09c92": 8772, // Coby
		"c09f42": 545,  // Apple
		"c09fe1": 3031, // ZTE
		"c0a00d": 125,  // ARRIS
		"c0a0bb": 803,  // D-Link
		"c0a0c7": 8773, // Fairfield
		"c0a1a2": 8774, // MarqMetrix
		"c0a36e": 3512, // BSkyB
		"c0a39e": 8775, // EarthCam
		"c0a53e": 545,  // Apple
		"c0a600": 545,  // Apple
		"c0a8f0": 8776, // Adamson
		"c0ac54": 2048, // Sagemcom
		"c0b101": 3031, // ZTE
		"c0b339": 8777, // Comigo
		"c0b357": 5917, // Yoshiki
		"c0b47d": 3334, // Huawei
		"c0b5cd": 3334, // Huawei
		"c0b658": 545,  // Apple
		"c0b6f9": 421,  // Intel
		"c0b883": 421,  // Intel
		"c0b8b1": 8778, // BitBox
		"c0b8e6": 5440, // Ruijie
		"c0bc9a": 3334, // Huawei
		"c0bdc8": 152,  // Samsung
		"c0bdd1": 152,  // Samsung
		"c0bfa7": 826,  // Juniper
		"c0bfc0": 3334, // Huawei
		"c0c1c0": 1783, // Linksys
		"c0c3b6": 8779, // Automatic
		"c0c520": 2737, // Ruckus
		"c0c522": 125,  // ARRIS
		"c0c946": 8780, // Mitsuya
		"c0c9e3": 1595, // TP-Link
		"c0ccf8": 545,  // Apple
		"c0cecd": 545,  // Apple
		"c0cfa3": 357,  // Creative
		"c0d012": 545,  // Apple
		"c0d026": 3334, // Huawei
		"c0d044": 2048, // Sagemcom
		"c0d193": 3334, // Huawei
		"c0d2dd": 152,  // Samsung
		"c0d3c0": 152,  // Samsung
		"c0d46b": 3334, // Huawei
		"c0d682": 3793, // Arista
		"c0d7aa": 2639, // Arcadyan
		"c0d834": 8781, // xvtec
		"c0d962": 2544, // Askey
		"c0dcd7": 3334, // Huawei
		"c0dcda": 152,  // Samsung
		"c0df77": 281,  // Conrad
		"c0e018": 3334, // Huawei
		"c0e1be": 3334, // Huawei
		"c0e3a0": 5695, // Renesas
		"c0e3fb": 3334, // Huawei
		"c0e42d": 1595, // TP-Link
		"c0e434": 3004, // Azurewave
		"c0e862": 545,  // Apple
		"c0eae4": 1003, // Sonicwall
		"c0ee40": 4536, // Laird
		"c0eefb": 6923, // OnePlus
		"c0f1c4": 8782, // Pacidal
		"c0f2fb": 545,  // Apple
		"c0f4e6": 3334, // Huawei
		"c0f535": 4498, // AMPAK
		"c0f6c2": 3334, // Huawei
		"c0f6ec": 3334, // Huawei
		"c0f79d": 8783, // Powercode
		"c0f827": 8784, // Rapidmax
		"c0f87f": 2,    // Cisco
		"c0f945": 35,   // Toshiba
		"c0f9b0": 3334, // Huawei
		"c0f9d2": 8785, // arkona
		"c0fbc1": 6281, // Itel
		"c0fd84": 3031, // ZTE
		"c0ffa8": 3334, // Huawei
		"c0ffd4": 1368, // Netgear
		"c40049": 8786, // Kamama
		"c400ad": 1703, // Advantech
		"c40142": 8787, // MaxMedia
		"c4017c": 2737, // Ruckus
		"c401b1": 8788, // SeekTech
		"c401ce": 8789, // Presition
		"c402e1": 8790, // Khwahish
		"c403a8": 421,  // Intel
		"c40415": 1368, // Netgear
		"c40528": 3334, // Huawei
		"c40683": 3334, // Huawei
		"c4072f": 3334, // Huawei
		"c4084a": 457,  // Nokia
		"c409b7": 826,  // Juniper
		"c40acb": 2,    // Cisco
		"c40b31": 545,  // Apple
		"c40bcb": 5697, // Xiaomi
		"c40d96": 3334, // Huawei
		"c40e45": 8791, // ACK
		"c40f09": 5590, // Hermes
		"c4108a": 2737, // Ruckus
		"c41234": 545,  // Apple
		"c412f5": 803,  // D-Link
		"c413e2": 185,  // Extreme
		"c41411": 545,  // Apple
		"c4143c": 2,    // Cisco
		"c41688": 3334, // Huawei
		"c416c8": 3334, // Huawei
		"c416fa": 8792, // Prysm
		"c4170e": 3334, // Huawei
		"c418e9": 152,  // Samsung
		"c419ec": 8793, // Qualisys
		"c41c07": 152,  // Samsung
		"c41c9c": 8794, // JiQiDao
		"c41cff": 3467, // Vizio
		"c421c8": 2848, // Kyocera
		"c42360": 421,  // Intel
		"c4237a": 8795, // WhizNets
		"c42728": 3031, // ZTE
		"c4278c": 3334, // Huawei
		"c42795": 6392, // Technicolor
		"c42ad0": 545,  // Apple
		"c42b44": 3334, // Huawei
		"c42c03": 545,  // Apple
		"c432d1": 8796, // Farlink
		"c4345b": 3334, // Huawei
		"c4346b": 302,  // HP
		"c4366c": 869,  // LG
		"c436c0": 1077, // Buffalo
		"c436da": 8797, // Rusteletech
		"c43772": 8798, // Virtuozzo
		"c43875": 2047, // Sonos
		"c438d3": 8799, // Tagatec
		"c4393a": 741,  // SMC
		"c43a35": 6440, // FN-Link
		"c43a9f": 8800, // Siconix
		"c43abe": 101,  // Sony
		"c43c3c": 8801, // Cybelec
		"c43cea": 1077, // Buffalo
		"c43dc7": 1368, // Netgear
		"c44044": 8802, // RackTop
		"c4411e": 2468, // Belkin
		"c44202": 152,  // Samsung
		"c44268": 2352, // Crestron
		"c4438f": 869,  // LG
		"c4447d": 3334, // Huawei
		"c444a0": 2,    // Cisco
		"c4473f": 3334, // Huawei
		"c44ad0": 8803, // Fireflies
		"c44b44": 8804, // Omniprint
		"c44d84": 2,    // Cisco
		"c44e1f": 8805, // BlueN
		"c44f33": 6376, // Espressif
		"c45006": 152,  // Samsung
		"c45444": 3071, // Quanta
		"c455a6": 8806, // Cadac
		"c455c2": 8807, // Bach-Simpson
		"c456fe": 647,  // Lava
		"c4576e": 152,  // Samsung
		"c45a86": 3334, // Huawei
		"c45bbe": 6376, // Espressif
		"c45bf7": 8808, // ants
		"c45d83": 152,  // Samsung
		"c45e5c": 3334, // Huawei
		"c46044": 929,  // Everex
		"c4618b": 545,  // Apple
		"c462ea": 152,  // Samsung
		"c46354": 8809, // U-Raku
		"c463fb": 8810, // Neatframe
		"c46413": 2,    // Cisco
		"c464b7": 4921, // Fiberhome
		"c46516": 302,  // HP
		"c46699": 6369, // Vivo
		"c467b5": 8811, // Libratone
		"c467d1": 3334, // Huawei
		"c469f0": 3334, // Huawei
		"c46ab7": 5697, // Xiaomi
		"c46bb4": 8812, // myIDkey
		"c46df1": 8813, // DataGravity
		"c46e1f": 1595, // TP-Link
		"c470ab": 5440, // Ruijie
		"c47154": 1595, // TP-Link
		"c471fe": 2,    // Cisco
		"c47295": 2,    // Cisco
		"c4731e": 152,  // Samsung
		"c4741e": 3031, // ZTE
		"c47469": 8814, // BT9
		"c475ab": 421,  // Intel
		"c478a2": 3334, // Huawei
		"c47ba3": 8815, // NAVIS
		"c47d46": 4,    // Fujitsu
		"c47d4f": 2,    // Cisco
		"c47d9f": 152,  // Samsung
		"c47dcc": 768,  // Zebra
		"c47dfe": 8816, // A.N
		"c47f51": 8817, // Inventek
		"c48025": 3334, // Huawei
		"c4836f": 4564, // Ciena
		"c48466": 545,  // Apple
		"c48508": 421,  // Intel
		"c486e9": 3334, // Huawei
		"c488e5": 152,  // Samsung
		"c48a5a": 8818, // Jfcontrol
		"c4910c": 545,  // Apple
		"c491cf": 8406, // Luxul
		"c49300": 8819, // 8Devices
		"c49313": 8820, // 100fio
		"c49380": 8821, // Speedytel
		"c493d9": 152,  // Samsung
		"c49500": 5438, // Amazon
		"c49805": 8822, // Minieum
		"c49880": 545,  // Apple
		"c49886": 6652, // Qorvo
		"c49a02": 869,  // LG
		"c49ded": 612,  // Microsoft
		"c49f4c": 3334, // Huawei
		"c49ff3": 8823, // Mciao
		"c4a366": 3031, // ZTE
		"c4a402": 3334, // Huawei
		"c4a81d": 803,  // D-Link
		"c4abb2": 6369, // Vivo
		"c4ac59": 2056, // Murata
		"c4ad21": 8824, // MEDIAEDGE
		"c4ad34": 1784, // Routerboard.com
		"c4adf1": 8825, // GOPEACE
		"c4ae12": 152,  // Samsung
		"c4b239": 2,    // Cisco
		"c4b301": 545,  // Apple
		"c4b36a": 2,    // Cisco
		"c4b8b4": 3334, // Huawei
		"c4b9cd": 2,    // Cisco
		"c4bd6a": 268,  // SKF
		"c4bde5": 421,  // Intel
		"c4bed4": 620,  // Avaya
		"c4bf60": 6265, // Tecno
		"c4c0ae": 8826, // Midori
		"c4c138": 8827, // OWLink
		"c4c36b": 545,  // Apple
		"c4c563": 6265, // Tecno
		"c4c603": 2,    // Cisco
		"c4ca2b": 3793, // Arista
		"c4d0e3": 421,  // Intel
		"c4d438": 3334, // Huawei
		"c4d655": 8828, // Tercel
		"c4d738": 3334, // Huawei
		"c4d8f3": 8829, // iZotope
		"c4d987": 421,  // Intel
		"c4da26": 8830, // Noblex
		"c4dd57": 6376, // Espressif
		"c4de7b": 3334, // Huawei
		"c4e17c": 8831, // U2S
		"c4e287": 3334, // Huawei
		"c4e506": 8832, // Piper
		"c4e510": 8833, // Mechatro
		"c4e532": 2639, // Arcadyan
		"c4e7be": 8834, // SCSpro
		"c4e90a": 803,  // D-Link
		"c4e984": 1595, // TP-Link
		"c4ea1d": 6392, // Technicolor
		"c4eef5": 8835, // II-VI
		"c4f081": 3334, // Huawei
		"c4f0ec": 4921, // Fiberhome
		"c4f122": 8836, // Nexar
		"c4f174": 5820, // eero
		"c4f464": 8837, // Spica
		"c4f57c": 90,   // Brocade
		"c4f5a5": 8838, // Kumalift
		"c4f7d5": 2,    // Cisco
		"c4fbaa": 3334, // Huawei
		"c4fcef": 8839, // SambaNova
		"c4fde6": 8840, // Drtech
		"c4fee2": 6622, // AMICCOM
		"c4ff1f": 3334, // Huawei
		"c4ff22": 3334, // Huawei
		"c80084": 2,    // Cisco
		"c80210": 869,  // LG
		"c8028f": 8841, // Nova
		"c803f5": 2737, // Ruckus
		"c80718": 8842, // TDSi
		"c80739": 6082, // NAKAYO
		"c80873": 2737, // Ruckus
		"c808e9": 869,  // LG
		"c809a8": 421,  // Intel
		"c80aa9": 3071, // Quanta
		"c80cc8": 3334, // Huawei
		"c80d32": 8843, // Holoplot
		"c80e95": 8844, // OmniLync
		"c81451": 3334, // Huawei
		"c81479": 152,  // Samsung
		"c816a5": 6362, // Masimo
		"c81739": 6281, // Itel
		"c819f7": 152,  // Samsung
		"c81afe": 8845, // DLOGIC
		"c81b5c": 8846, // BCTech
		"c81cfe": 768,  // Zebra
		"c81ee7": 545,  // Apple
		"c81f66": 102,  // Dell
		"c81fbe": 3334, // Huawei
		"c81fea": 620,  // Avaya
		"c8208e": 8847, // Storagedata
		"c82158": 421,  // Intel
		"c8292a": 8848, // Barun
		"c82a14": 545,  // Apple
		"c82b96": 6376, // Espressif
		"c82e94": 8849, // Halfa
		"c83168": 8850, // eZEX
		"c8334b": 545,  // Apple
		"c833e5": 3334, // Huawei
		"c8348e": 421,  // Intel
		"c835b8": 306,  // Ericsson
		"c83870": 152,  // Samsung
		"c839ac": 3334, // Huawei
		"c83a35": 6266, // Tenda
		"c83a6b": 1916, // Roku
		"c83b45": 8851, // JRI
		"c83c85": 545,  // Apple
		"c83d97": 457,  // Nokia
		"c83dd4": 189,  // CyberTAN
		"c83ddc": 5697, // Xiaomi
		"c83dfc": 8852, // AlphaTheta
		"c83e9e": 3334, // Huawei
		"c83ea7": 8853, // KUNBUS
		"c83f26": 612,  // Microsoft
		"c83fb4": 125,  // ARRIS
		"c84029": 4921, // Fiberhome
		"c8418a": 152,  // Samsung
		"c84529": 8854, // IMK
		"c8458f": 8855, // Wyler
		"c84782": 8856, // Areson
		"c8478c": 7781, // Beken
		"c84c75": 2,    // Cisco
		"c84f86": 3578, // Sophos
		"c850ce": 3334, // Huawei
		"c850e9": 2050, // Raisecom
		"c85142": 152,  // Samsung
		"c85195": 3334, // Huawei
		"c85261": 125,  // ARRIS
		"c8544b": 2692, // ZyXEL
		"c854a4": 6295, // Infinix
		"c85663": 8857, // Sunflex
		"c85895": 687,  // Motorola
		"c858c0": 421,  // Intel
		"c85a9f": 3031, // ZTE
		"c85acf": 302,  // HP
		"c85b76": 4919, // LCFC
		"c85d38": 531,  // HUMAX
		"c86000": 1806, // ASUS
		"c863f1": 101,  // Sony
		"c863fc": 125,  // ARRIS
		"c864c7": 3031, // ZTE
		"c8665d": 185,  // Extreme
		"c8675e": 185,  // Extreme
		"c868de": 3334, // Huawei
		"c869cd": 545,  // Apple
		"c86c1e": 8858, // Display
		"c86c3d": 5438, // Amazon
		"c86c87": 2692, // ZyXEL
		"c86cb6": 8859, // Optcom
		"c86f1d": 545,  // Apple
		"c87248": 8860, // Aplicom
		"c8727e": 457,  // Nokia
		"c8755b": 8861, // Quantify
		"c87765": 1925, // Tiesse
		"c87b23": 1819, // Bose
		"c87b5b": 3031, // ZTE
		"c87cbc": 8862, // Valink
		"c87e75": 152,  // Samsung
		"c88314": 8863, // Tempo
		"c88439": 8864, // Sunrise
		"c88447": 8865, // Beautiful
		"c884a1": 2,    // Cisco
		"c884cf": 3334, // Huawei
		"c88550": 545,  // Apple
		"c88722": 8866, // Lumenpulse
		"c889f3": 545,  // Apple
		"c88be8": 6362, // Masimo
		"c88d83": 3334, // Huawei
		"c8903e": 8867, // Pakton
		"c891f9": 2048, // Sagemcom
		"c89346": 8868, // MXCHIP
		"c894bb": 3334, // Huawei
		"c89665": 612,  // Microsoft
		"c8979f": 457,  // Nokia
		"c899b2": 2639, // Arcadyan
		"c89c13": 8869, // Inspiremobile
		"c89c1d": 2,    // Cisco
		"c89cdc": 1115, // Elitegroup
		"c89d18": 3334, // Huawei
		"c89e43": 1368, // Netgear
		"c89f1a": 3334, // Huawei
		"c89f1d": 8624, // Shenzhen
		"c8a1ba": 8870, // Neul
		"c8a620": 8871, // Nebula
		"c8a729": 8872, // Systronics
		"c8a776": 3334, // Huawei
		"c8a823": 152,  // Samsung
		"c8a9fc": 8873, // Goyoo
		"c8aa21": 125,  // ARRIS
		"c8aacc": 67,   // Private
		"c8b1cd": 545,  // Apple
		"c8b1ee": 6652, // Qorvo
		"c8b21e": 6439, // Chipsea
		"c8b29b": 421,  // Intel
		"c8b373": 1783, // Linksys
		"c8b422": 2544, // Askey
		"c8b5ad": 302,  // HP
		"c8b5b7": 545,  // Apple
		"c8b6d3": 3334, // Huawei
		"c8b82f": 5820, // eero
		"c8ba94": 152,  // Samsung
		"c8bae9": 8874, // Qdis
		"c8bb81": 3334, // Huawei
		"c8bbd3": 8875, // Embrane
		"c8bc9c": 3334, // Huawei
		"c8bcc8": 545,  // Apple
		"c8bd4d": 152,  // Samsung
		"c8bd69": 152,  // Samsung
		"c8be19": 803,  // D-Link
		"c8bffe": 3334, // Huawei
		"c8c2f5": 833,  // Flextronics
		"c8c2fa": 3334, // Huawei
		"c8c465": 3334, // Huawei
		"c8c64a": 833,  // Flextronics
		"c8c750": 687,  // Motorola
		"c8c791": 8876, // Zero1.tv
		"c8c9a3": 6376, // Espressif
		"c8ca63": 3334, // Huawei
		"c8ca79": 4564, // Ciena
		"c8cb9e": 421,  // Intel
		"c8cbb8": 302,  // HP
		"c8cd72": 2048, // Sagemcom
		"c8d083": 545,  // Apple
		"c8d10b": 457,  // Nokia
		"c8d12a": 3869, // Comtrend
		"c8d15e": 3334, // Huawei
		"c8d1d1": 8877, // AGAiT
		"c8d2c1": 8878, // Jetlun
		"c8d3a3": 803,  // D-Link
		"c8d3ff": 302,  // HP
		"c8d429": 8879, // Muehlbauer
		"c8d719": 1783, // Linksys
		"c8d7b0": 152,  // Samsung
		"c8d884": 3857, // Universal
		"c8d9d2": 302,  // HP
		"c8db26": 4079, // Logitech
		"c8ddc9": 2664, // Lenovo
		"c8de51": 8880, // IntegraOptics
		"c8dec9": 6680, // Coriant
		"c8df7c": 457,  // Nokia
		"c8e0eb": 545,  // Apple
		"c8e1a7": 6955, // Vertu
		"c8e265": 421,  // Intel
		"c8e600": 3334, // Huawei
		"c8e776": 8881, // PTCOM
		"c8e7d8": 4253, // Mercury
		"c8e7f0": 826,  // Juniper
		"c8eaf8": 3031, // ZTE
		"c8ee08": 8882, // Tangtop
		"c8ee75": 8883, // Pishion
		"c8f319": 869,  // LG
		"c8f406": 620,  // Avaya
		"c8f650": 545,  // Apple
		"c8f68d": 8884, // S.E.Technologies
		"c8f6c8": 4921, // Fiberhome
		"c8f733": 421,  // Intel
		"c8f750": 102,  // Dell
		"c8f946": 8885, // LOCOSYS
		"c8f9c8": 8886, // NewSharp
		"c8f9f9": 2,    // Cisco
		"c8fa84": 8887, // Trusonus
		"c8fe6a": 826,  // Juniper
		"c8ff28": 2873, // Liteon
		"c8ff77": 7189, // Dyson
		"cc03fa": 6392, // Technicolor
		"cc051b": 152,  // Samsung
		"cc0577": 3334, // Huawei
		"cc0677": 4921, // Fiberhome
		"cc07ab": 152,  // Samsung
		"cc07e4": 2664, // Lenovo
		"cc088d": 545,  // Apple
		"cc08e0": 545,  // Apple
		"cc08fb": 1595, // TP-Link
		"cc09c8": 8888, // Imaqliq
		"cc0cda": 8889, // Miljovakt
		"cc0df2": 687,  // Motorola
		"cc1531": 421,  // Intel
		"cc167e": 2,    // Cisco
		"cc187b": 8890, // Manzanita
		"cc1afa": 3031, // ZTE
		"cc1fc4": 8891, // InVue
		"cc208c": 3334, // Huawei
		"cc20e8": 545,  // Apple
		"cc2119": 152,  // Samsung
		"cc2218": 8892, // InnoDigital
		"cc25ef": 545,  // Apple
		"cc262d": 8893, // Verifi
		"cc29f5": 545,  // Apple
		"cc2d1b": 3187, // SFR
		"cc2d21": 6266, // Tenda
		"cc2d8c": 869,  // LG
		"cc2db7": 545,  // Apple
		"cc2de0": 1784, // Routerboard.com
		"cc2f71": 421,  // Intel
		"cc3080": 8894, // VAIO
		"cc3296": 3334, // Huawei
		"cc32e5": 1595, // TP-Link
		"cc33bb": 2048, // Sagemcom
		"cc3429": 1595, // TP-Link
		"cc3540": 6392, // Technicolor
		"cc355a": 8895, // SecuGen
		"cc37ab": 6294, // Edgecore
		"cc398c": 8896, // Shiningtek
		"cc3a61": 152,  // Samsung
		"cc3b27": 6265, // Tecno
		"cc3b58": 8897, // Curiouser
		"cc3d82": 421,  // Intel
		"cc3e5f": 302,  // HP
		"cc3f8a": 5902, // Komatsu
		"cc3fea": 1743, // BAE
		"cc40d0": 1368, // Netgear
		"cc4463": 545,  // Apple
		"cc4639": 8898, // WAAV
		"cc464e": 152,  // Samsung
		"cc46d6": 2,    // Cisco
		"cc4703": 8899, // Intercon
		"cc47bd": 8900, // Rhombus
		"cc483a": 102,  // Dell
		"cc4b73": 4498, // AMPAK
		"cc4d38": 8901, // Carnegie
		"cc4e24": 90,   // Brocade
		"cc4eec": 531,  // HUMAX
		"cc500a": 4921, // Fiberhome
		"cc501c": 8902, // KVH
		"cc5076": 8903, // Ocom
		"cc50e3": 6376, // Espressif
		"cc53b5": 3334, // Huawei
		"cc5459": 8904, // OnTime
		"cc55ad": 4555, // RIM
		"cc5a53": 2,    // Cisco
		"cc5b31": 1427, // Nintendo
		"cc5c61": 3334, // Huawei
		"cc5d4e": 2692, // ZyXEL
		"cc5d78": 8905, // JTD
		"cc60c8": 612,  // Microsoft
		"cc61e5": 687,  // Motorola
		"cc64a6": 3334, // Huawei
		"cc65ad": 125,  // ARRIS
		"cc660a": 545,  // Apple
		"cc66b2": 457,  // Nokia
		"cc68b6": 1595, // TP-Link
		"cc69fa": 545,  // Apple
		"cc6da0": 1916, // Roku
		"cc6ea4": 152,  // Samsung
		"cc70ed": 2,    // Cisco
		"cc720f": 8906, // Viscount
		"cc7498": 8907, // Filmetrics
		"cc75e2": 125,  // ARRIS
		"cc7669": 8908, // Seetech
		"cc785f": 545,  // Apple
		"cc794a": 7460, // BLU
		"cc7b35": 3031, // ZTE
		"cc7b61": 8909, // Nikkiso
		"cc7d37": 125,  // ARRIS
		"cc7ee7": 2153, // Panasonic
		"cc7f75": 2,    // Cisco
		"cc7f76": 2,    // Cisco
		"cc812a": 6369, // Vivo
		"cc81da": 6848, // Phicomm
		"cc82eb": 2848, // Kyocera
		"cc86ec": 1655, // Silicon
		"cc874a": 457,  // Nokia
		"cc8826": 869,  // LG
		"cc88c7": 1685, // Aruba
		"cc895e": 3334, // Huawei
		"cc89fd": 457,  // Nokia
		"cc8e71": 2,    // Cisco
		"cc9070": 2,    // Cisco
		"cc9470": 8910, // Kinestral
		"cc95d7": 3467, // Vizio
		"cc9635": 8911, // LVS
		"cc96a0": 3334, // Huawei
		"cc9891": 2,    // Cisco
		"cc9da2": 8495, // Eltex
		"cc9e00": 1427, // Nintendo
		"cc9ea2": 5438, // Amazon
		"cca223": 3334, // Huawei
		"cca3bd": 6281, // Itel
		"cca462": 125,  // ARRIS
		"cca614": 8912, // Aifa
		"cca7c1": 3521, // Google
		"ccab2c": 531,  // HUMAX
		"ccb0a8": 3334, // Huawei
		"ccb0da": 2873, // Liteon
		"ccb11a": 152,  // Samsung
		"ccb182": 3334, // Huawei
		"ccb255": 803,  // D-Link
		"ccb691": 8913, // NECMagnusCommunications
		"ccb8a8": 4498, // AMPAK
		"ccbbfe": 3334, // Huawei
		"ccbce3": 3334, // Huawei
		"ccbd35": 8914, // Steinel
		"ccbe59": 374,  // Calix
		"ccbe71": 8915, // OptiLogix
		"ccc079": 2056, // Murata
		"ccc2e0": 2050, // Raisecom
		"ccc3ea": 687,  // Motorola
		"ccc5e5": 102,  // Dell
		"ccc62b": 8916, // Tri-Systems
		"ccc760": 545,  // Apple
		"ccc95d": 545,  // Apple
		"cccc81": 3334, // Huawei
		"cccccc": 1655, // Silicon
		"cccd64": 8917, // SM-Electronic
		"ccce40": 8918, // Janteq
		"ccd083": 1685, // Aruba
		"ccd281": 545,  // Apple
		"ccd3c1": 1449, // Vestel
		"ccd42e": 2639, // Arcadyan
		"ccd4a1": 6428, // MitraStar
		"ccd539": 2,    // Cisco
		"ccd73c": 3334, // Huawei
		"ccd811": 8919, // Aiconn
		"ccd81f": 8920, // Maipu
		"ccd8c1": 2,    // Cisco
		"ccd9ac": 421,  // Intel
		"ccd9e9": 8921, // SCR
		"ccdb04": 8922, // DataRemote
		"ccdb93": 2,    // Cisco
		"ccdc55": 8923, // Dragonchip
		"cce0c3": 8924, // EXTEN
		"cce17f": 826,  // Juniper
		"cce194": 826,  // Juniper
		"cce1d5": 1077, // Buffalo
		"cce8ac": 8925, // SOYEA
		"ccea1c": 8926, // DCONWORKS
		"cced4d": 2,    // Cisco
		"cceddc": 6428, // MitraStar
		"ccef48": 2,    // Cisco
		"ccf407": 8927, // Eukrea
		"ccf411": 3521, // Google
		"ccf538": 8928, // 3isysnetworks
		"ccf67a": 8929, // Ayecka
		"ccf735": 5438, // Amazon
		"ccf841": 8930, // Lumewave
		"ccf954": 620,  // Avaya
		"ccf957": 7419, // u-blox
		"ccf9e4": 421,  // Intel
		"ccf9e8": 152,  // Samsung
		"ccfa00": 869,  // LG
		"ccfa66": 3334, // Huawei
		"ccfb65": 1427, // Nintendo
		"ccfcb1": 231,  // Wireless
		"ccfd17": 6433, // TCT
		"ccfe3c": 152,  // Samsung
		"ccff90": 3334, // Huawei
		"d0034b": 545,  // Apple
		"d003df": 152,  // Samsung
		"d00401": 687,  // Motorola
		"d00492": 4921, // Fiberhome
		"d004b0": 152,  // Samsung
		"d0052a": 2639, // Arcadyan
		"d005e4": 3334, // Huawei
		"d007ca": 826,  // Juniper
		"d00df7": 3334, // Huawei
		"d00f6d": 4048, // T&W
		"d01242": 8931, // BIOS
		"d0131e": 8932, // Sunrex
		"d013fd": 869,  // LG
		"d0154a": 3031, // ZTE
		"d015a6": 1685, // Aruba
		"d0167c": 5820, // eero
		"d016b4": 3334, // Huawei
		"d01769": 2056, // Murata
		"d0176a": 152,  // Samsung
		"d017c2": 1806, // ASUS
		"d0196a": 4564, // Ciena
		"d01aa7": 8933, // UniPrint
		"d01b1f": 6577, // Ohsung
		"d01b49": 152,  // Samsung
		"d01c3c": 6265, // Tecno
		"d021ac": 8934, // Yo
		"d021f9": 2978, // Ubiquiti
		"d022be": 152,  // Samsung
		"d023db": 545,  // Apple
		"d02516": 4253, // Mercury
		"d02544": 152,  // Samsung
		"d02598": 545,  // Apple
		"d02b20": 545,  // Apple
		"d02c45": 8935, // littleBits
		"d02db3": 3334, // Huawei
		"d03169": 152,  // Samsung
		"d03311": 545,  // Apple
		"d03742": 3096, // Yulong
		"d03745": 1595, // TP-Link
		"d039b3": 125,  // ARRIS
		"d039ea": 5558, // NetApp
		"d03c1f": 421,  // Intel
		"d03dc3": 8936, // AQ
		"d03e5c": 3334, // Huawei
		"d03e7d": 6439, // Chipsea
		"d03f27": 6881, // Wyze
		"d03faa": 545,  // Apple
		"d040ef": 2056, // Murata
		"d041c9": 4921, // Fiberhome
		"d0431e": 102,  // Dell
		"d047c1": 8266, // Elma
		"d048f3": 8937, // DATTUS
		"d0497c": 6923, // OnePlus
		"d04cc1": 8938, // SINTRONES
		"d04d2c": 1916, // Roku
		"d04dc6": 1685, // Aruba
		"d04f7e": 545,  // Apple
		"d05099": 4719, // ASRock
		"d05162": 101,  // Sony
		"d05349": 2873, // Liteon
		"d0542d": 1640, // Cambridge
		"d05509": 1427, // Nintendo
		"d056bf": 8939, // Amosense
		"d0574c": 2,    // Cisco
		"d0577b": 421,  // Intel
		"d05785": 2277, // Pantech
		"d05794": 2048, // Sagemcom
		"d058a8": 3031, // ZTE
		"d058fc": 3512, // BSkyB
		"d05919": 3031, // ZTE
		"d05995": 4921, // Fiberhome
		"d059c3": 8940, // CeraMicro
		"d059e4": 152,  // Samsung
		"d05a00": 6392, // Technicolor
		"d05ba8": 3031, // ZTE
		"d0608c": 3031, // ZTE
		"d063b4": 8941, // SolidRun
		"d06544": 545,  // Apple
		"d065ca": 3334, // Huawei
		"d0667b": 152,  // Samsung
		"d06726": 302,  // HP
		"d067e5": 102,  // Dell
		"d06a1f": 8942, // BSE
		"d06dc9": 2048, // Sagemcom
		"d06ede": 2048, // Sagemcom
		"d06f4a": 8943, // Topwell
		"d06f82": 3334, // Huawei
		"d071c4": 3031, // ZTE
		"d072dc": 2,    // Cisco
		"d0737f": 8944, // Mini-Circuits
		"d0768f": 374,  // Calix
		"d076e7": 1595, // TP-Link
		"d07714": 687,  // Motorola
		"d07880": 4921, // Fiberhome
		"d07ab5": 3334, // Huawei
		"d07d33": 3334, // Huawei
		"d07e28": 302,  // HP
		"d07e35": 421,  // Intel
		"d07fa0": 152,  // Samsung
		"d0817a": 545,  // Apple
		"d084b0": 2048, // Sagemcom
		"d087e2": 152,  // Samsung
		"d0880c": 545,  // Apple
		"d08999": 8945, // APCON
		"d08a55": 7039, // Skullcandy
		"d08a91": 6392, // Technicolor
		"d08cff": 8946, // Upwis
		"d08e79": 102,  // Dell
		"d0929e": 612,  // Microsoft
		"d092fa": 4921, // Fiberhome
		"d09380": 8947, // Ducere
		"d09466": 102,  // Dell
		"d095c7": 2277, // Pantech
		"d09b05": 8948, // Emtronix
		"d09c7a": 5697, // Xiaomi
		"d09cae": 6369, // Vivo
		"d09d0a": 8949, // Linkcom
		"d09dab": 6433, // TCT
		"d0a4b1": 8950, // Sonifex
		"d0a5a6": 2,    // Cisco
		"d0a637": 545,  // Apple
		"d0abd5": 421,  // Intel
		"d0aeec": 2236, // Alpha
		"d0afb6": 8951, // Linktop
		"d0b0cd": 8952, // Moen
		"d0b128": 152,  // Samsung
		"d0b214": 8953, // PoeWit
		"d0b2c4": 6392, // Technicolor
		"d0b45d": 3334, // Huawei
		"d0b60a": 8954, // Xingluo
		"d0b66f": 6255, // Sernet
		"d0bb61": 3031, // ZTE
		"d0bd01": 4698, // DS
		"d0be2c": 8955, // CNSLink
		"d0bf9c": 302,  // HP
		"d0c193": 8956, // Skybell
		"d0c1b1": 152,  // Samsung
		"d0c24e": 152,  // Samsung
		"d0c282": 2,    // Cisco
		"d0c31e": 8957, // JUNGJIN
		"d0c5d3": 3004, // Azurewave
		"d0c5d8": 8958, // Latecoere
		"d0c5f3": 545,  // Apple
		"d0c637": 421,  // Intel
		"d0c65b": 3334, // Huawei
		"d0c789": 2,    // Cisco
		"d0c7c0": 1595, // TP-Link
		"d0cde1": 8959, // Scientech
		"d0d003": 152,  // Samsung
		"d0d04b": 3334, // Huawei
		"d0d0fd": 2,    // Cisco
		"d0d212": 8960, // K2NET
		"d0d23c": 545,  // Apple
		"d0d2b0": 545,  // Apple
		"d0d3e0": 1685, // Aruba
		"d0d3fc": 8961, // Mios
		"d0d471": 8962, // MVTECH
		"d0d6cc": 8963, // Wintop
		"d0d783": 3334, // Huawei
		"d0db32": 457,  // Nokia
		"d0dbb7": 3170, // Casa
		"d0dd49": 826,  // Juniper
		"d0df9a": 2873, // Liteon
		"d0dfb2": 4764, // Genie
		"d0dfc7": 152,  // Samsung
		"d0e042": 2,    // Cisco
		"d0e140": 545,  // Apple
		"d0e347": 8964, // Yoga
		"d0e40b": 8965, // Wearable
		"d0e44a": 2056, // Murata
		"d0e54d": 125,  // ARRIS
		"d0e782": 3004, // Azurewave
		"d0e828": 228,  // Radiant
		"d0eb03": 8966, // Zhehua
		"d0eb9e": 8967, // Seowoo
		"d0ec35": 2,    // Cisco
		"d0efc1": 3334, // Huawei
		"d0f0db": 306,  // Ericsson
		"d0f3f5": 3334, // Huawei
		"d0f520": 2848, // Kyocera
		"d0f865": 6281, // Itel
		"d0f88c": 687,  // Motorola
		"d0f99b": 3031, // ZTE
		"d0fccc": 152,  // Samsung
		"d0ff98": 3334, // Huawei
		"d40057": 8968, // MC
		"d40129": 858,  // Broadcom
		"d4016d": 1595, // TP-Link
		"d4024a": 8969, // Delphian
		"d404cd": 125,  // ARRIS
		"d404ff": 826,  // Juniper
		"d40598": 125,  // ARRIS
		"d405de": 5820, // eero
		"d40aa9": 125,  // ARRIS
		"d40b1a": 1341, // HTC
		"d41090": 8970, // InfORM
		"d411a3": 152,  // Samsung
		"d411d6": 8971, // ShotSpotter
		"d41243": 4498, // AMPAK
		"d41296": 8972, // Anobit
		"d413f8": 2481, // Peplink
		"d41ad1": 2692, // ZyXEL
		"d41e35": 8973, // TOHO
		"d41f0c": 1873, // JAI
		"d4206d": 1341, // HTC
		"d420b0": 7484, // Mist
		"d42122": 2074, // Sercomm
		"d4223f": 2664, // Lenovo
		"d42493": 2281, // GW
		"d4258b": 421,  // Intel
		"d42751": 8974, // Infopia
		"d428b2": 8975, // ioBridge
		"d428d5": 6433, // TCT
		"d429ea": 8976, // Zimory
		"d42c0f": 125,  // ARRIS
		"d42c44": 2,    // Cisco
		"d42c46": 1077, // Buffalo
		"d42dc5": 2153, // Panasonic
		"d42f23": 8977, // Akenori
		"d4319d": 8978, // Sinwatec
		"d43260": 6241, // GoPro
		"d43266": 8979, // Fike
		"d4351d": 6392, // Technicolor
		"d4354a": 2653, // ALAXALA
		"d437d7": 3031, // ZTE
		"d4389c": 101,  // Sony
		"d439b8": 4564, // Ciena
		"d43a65": 8980, // IGRS
		"d43b04": 421,  // Intel
		"d43d67": 8981, // Carma
		"d43d7e": 1812, // Micro-Star
		"d43df3": 2692, // ZyXEL
		"d43fcb": 125,  // ARRIS
		"d440d0": 8982, // OCOSMOS
		"d440f0": 3334, // Huawei
		"d44649": 3334, // Huawei
		"d446e1": 545,  // Apple
		"d4475a": 8983, // ScreenBeam
		"d44ca7": 8984, // Informtekhnika
		"d44da4": 2056, // Murata
		"d44f67": 3334, // Huawei
		"d44f68": 8985, // Eidetic
		"d4522a": 8986, // TangoWiFi.com
		"d45297": 8987, // nSTREAMS
		"d452ee": 3512, // BSkyB
		"d45383": 2056, // Murata
		"d4548b": 421,  // Intel
		"d45763": 545,  // Apple
		"d45800": 4921, // Fiberhome
		"d45ab2": 8988, // Galleon
		"d45d42": 457,  // Nokia
		"d45d64": 1806, // ASUS
		"d45ddf": 5439, // Pegatron
		"d460e3": 2074, // Sercomm
		"d4612e": 3334, // Huawei
		"d4619d": 545,  // Apple
		"d461da": 545,  // Apple
		"d462ea": 3334, // Huawei
		"d463c6": 687,  // Motorola
		"d463de": 6369, // Vivo
		"d463fe": 2639, // Arcadyan
		"d46624": 2,    // Cisco
		"d466a8": 8989, // Riedo
		"d46761": 8990, // XonTel
		"d467e7": 4921, // Fiberhome
		"d4684d": 2737, // Ruckus
		"d468aa": 545,  // Apple
		"d469a5": 3717, // Miura
		"d46a35": 2,    // Cisco
		"d46a91": 6556, // SnapAV
		"d46aa8": 3334, // Huawei
		"d46ba6": 3334, // Huawei
		"d46c6d": 125,  // ARRIS
		"d46cda": 8991, // CSM
		"d46d50": 2,    // Cisco
		"d46d6d": 421,  // Intel
		"d46e0e": 1595, // TP-Link
		"d46e5c": 3334, // Huawei
		"d47208": 8992, // Bragi
		"d47226": 3031, // ZTE
		"d47415": 3334, // Huawei
		"d476a0": 1323, // Fortinet
		"d476ea": 3031, // ZTE
		"d47798": 2,    // Cisco
		"d47856": 620,  // Avaya
		"d4789b": 2,    // Cisco
		"d47954": 3334, // Huawei
		"d479c3": 8993, // Cameronet
		"d47ae2": 152,  // Samsung
		"d47b75": 1596, // HARTING
		"d47bb0": 2544, // Askey
		"d47dfc": 6265, // Tecno
		"d481ca": 8994, // iDevices
		"d481d7": 102,  // Dell
		"d4823e": 2400, // Argosy
		"d48564": 302,  // HP
		"d48660": 2639, // Arcadyan
		"d487d8": 152,  // Samsung
		"d4883f": 8995, // Hdpro
		"d48866": 3334, // Huawei
		"d48890": 152,  // Samsung
		"d48a39": 152,  // Samsung
		"d48cb5": 2,    // Cisco
		"d48dd9": 8996, // Meld
		"d48f33": 612,  // Microsoft
		"d48fa2": 3334, // Huawei
		"d48faa": 8997, // Sogecam
		"d4909c": 545,  // Apple
		"d490e0": 8487, // Topcon
		"d4910f": 5438, // Amazon
		"d49234": 48,   // NEC
		"d49390": 5687, // Clevo
		"d49398": 457,  // Nokia
		"d493a0": 8998, // Fidelix
		"d49400": 3334, // Huawei
		"d4945a": 354,  // Cosmo
		"d494e8": 3334, // Huawei
		"d4970b": 5697, // Xiaomi
		"d49a20": 545,  // Apple
		"d49aa0": 8374, // Vnpt
		"d49c28": 8999, // JayBird
		"d49cdd": 4498, // AMPAK
		"d49dc0": 152,  // Samsung
		"d49e05": 3031, // ZTE
		"d49fdd": 3334, // Huawei
		"d4a02a": 2,    // Cisco
		"d4a148": 3334, // Huawei
		"d4a33d": 545,  // Apple
		"d4a425": 9000, // SMAX
		"d4a499": 9001, // InView
		"d4ab82": 125,  // ARRIS
		"d4ad2d": 4921, // Fiberhome
		"d4ad71": 2,    // Cisco
		"d4adbd": 2,    // Cisco
		"d4adfc": 67,   // Private
		"d4ae05": 152,  // Samsung
		"d4ae52": 102,  // Dell
		"d4aff7": 3793, // Arista
		"d4b110": 3334, // Huawei
		"d4b27a": 125,  // ARRIS
		"d4b709": 3031, // ZTE
		"d4b7d0": 4564, // Ciena
		"d4b92f": 6392, // Technicolor
		"d4bbc8": 6369, // Vivo
		"d4bbe6": 3334, // Huawei
		"d4bd1e": 9002, // 5VT
		"d4bd4f": 2737, // Ruckus
		"d4bed9": 102,  // Dell
		"d4bf7f": 9003, // Upvel
		"d4c19e": 2737, // Ruckus
		"d4c1c8": 3031, // ZTE
		"d4c1fc": 457,  // Nokia
		"d4c3b0": 9004, // Gearlinx
		"d4c766": 9005, // Acentic
		"d4c8b0": 1040, // Prime
		"d4c93c": 2,    // Cisco
		"d4c94b": 687,  // Motorola
		"d4c9b2": 9006, // Quanergy
		"d4c9ef": 302,  // HP
		"d4ca6d": 1784, // Routerboard.com
		"d4ca6e": 7419, // u-blox
		"d4cbaf": 457,  // Nokia
		"d4ceb8": 9007, // Enatel
		"d4d252": 421,  // Intel
		"d4d2d6": 6440, // FN-Link
		"d4d2e5": 9008, // BKAV
		"d4d51b": 3334, // Huawei
		"d4d748": 2,    // Cisco
		"d4d919": 6241, // GoPro
		"d4dacd": 3512, // BSkyB
		"d4dc09": 7484, // Mist
		"d4dccd": 545,  // Apple
		"d4e08e": 9009, // ValueHD
		"d4e2cb": 6392, // Technicolor
		"d4e33f": 457,  // Nokia
		"d4e6b7": 152,  // Samsung
		"d4e880": 2,    // Cisco
		"d4e8b2": 152,  // Samsung
		"d4e90b": 9010, // CVT
		"d4ea0e": 620,  // Avaya
		"d4eb68": 2,    // Cisco
		"d4ec0c": 9011, // Harley-Davidson
		"d4ecab": 6369, // Vivo
		"d4ee07": 9012, // HIWIFI
		"d4f057": 1427, // Nintendo
		"d4f143": 9013, // IPROAD
		"d4f207": 9014, // DIAODIAOTechnology
		"d4f337": 9015, // Xunison
		"d4f46f": 545,  // Apple
		"d4f527": 300,  // Siemens
		"d4f547": 3521, // Google
		"d4f5ef": 302,  // HP
		"d4f756": 3031, // ZTE
		"d4f786": 4921, // Fiberhome
		"d4f829": 2048, // Sagemcom
		"d4f98d": 6376, // Espressif
		"d4f9a1": 3334, // Huawei
		"d4fb8e": 545,  // Apple
		"d4fc13": 4921, // Fiberhome
		"d8004d": 545,  // Apple
		"d80093": 9016, // Aurender
		"d8052e": 9017, // Skyviia
		"d807b6": 1595, // TP-Link
		"d80831": 152,  // Samsung
		"d808f5": 9018, // Arcadia
		"d809c3": 9019, // Cercacor
		"d809d6": 9020, // Zexelon
		"d80a60": 3334, // Huawei
		"d80b9a": 152,  // Samsung
		"d80d17": 1595, // TP-Link
		"d80de3": 9021, // FXI
		"d8109f": 3334, // Huawei
		"d8150d": 1595, // TP-Link
		"d8160a": 9022, // Nippon
		"d816c1": 7749, // Dewav
		"d818d3": 826,  // Juniper
		"d8197a": 9023, // Nuheara
		"d819ce": 9024, // Telesquare
		"d81bfe": 9025, // Twinlinx
		"d81c14": 9026, // Compacta
		"d81c79": 545,  // Apple
		"d81d72": 545,  // Apple
		"d81fcc": 90,   // Brocade
		"d824bd": 2,    // Cisco
		"d82522": 125,  // ARRIS
		"d825b0": 9027, // Rockeetech
		"d8270c": 9028, // MaxTronic
		"d82916": 1690, // Ascent
		"d82918": 3334, // Huawei
		"d82a15": 9029, // Leitner
		"d82a7e": 457,  // Nokia
		"d82de1": 9030, // Tricascade
		"d83062": 545,  // Apple
		"d83134": 1916, // Roku
		"d831cf": 152,  // Samsung
		"d83214": 6266, // Tenda
		"d832e3": 5697, // Xiaomi
		"d833b7": 2048, // Sagemcom
		"d8365f": 3541, // Intelbras
		"d838fc": 2737, // Ruckus
		"d83af5": 9031, // Wideband
		"d83bbf": 421,  // Intel
		"d84008": 3334, // Huawei
		"d843ed": 3602, // Suzuken
		"d84732": 1595, // TP-Link
		"d847bb": 3334, // Huawei
		"d8490b": 3334, // Huawei
		"d8492f": 87,   // Canon
		"d84a2b": 3031, // ZTE
		"d84b2a": 9032, // Cognitas
		"d84c90": 545,  // Apple
		"d84f37": 9033, // Proxis
		"d84fb8": 869,  // LG
		"d850e6": 1806, // ASUS
		"d8539a": 826,  // Juniper
		"d854a2": 185,  // Extreme
		"d85575": 152,  // Samsung
		"d855a3": 3031, // ZTE
		"d857ef": 152,  // Samsung
		"d858d7": 9034, // CZ.NIC
		"d85982": 3334, // Huawei
		"d85b2a": 152,  // Samsung
		"d85d4c": 1595, // TP-Link
		"d85dfb": 67,   // Private
		"d85ed3": 1929, // Giga-Byte
		"d86162": 1592, // Wistron
		"d862db": 9035, // Eno
		"d86375": 5697, // Xiaomi
		"d866ee": 9036, // Boxin
		"d867d3": 3334, // Huawei
		"d867d9": 2,    // Cisco
		"d86852": 3334, // Huawei
		"d868c3": 152,  // Samsung
		"d86960": 9037, // Steinsvik
		"d86bf7": 1427, // Nintendo
		"d86c5a": 531,  // HUMAX
		"d86c63": 3521, // Google
		"d86ce9": 2048, // Sagemcom
		"d86d17": 3334, // Huawei
		"d87157": 2664, // Lenovo
		"d87495": 3031, // ZTE
		"d87533": 457,  // Nokia
		"d8760a": 9038, // Escort
		"d876ae": 3334, // Huawei
		"d8778b": 3541, // Intelbras
		"d878e5": 9039, // Kuhn
		"d87cdd": 9040, // Sanix
		"d87d7f": 2048, // Sagemcom
		"d87e76": 6281, // Itel
		"d87eb1": 9041, // x.o.ware
		"d88039": 706,  // Microchip
		"d881ce": 9042, // AHN
		"d88466": 185,  // Extreme
		"d887d5": 9043, // Leadcore
		"d888ce": 2855, // RF
		"d88a3b": 9044, // Unit-EM
		"d88adc": 3334, // Huawei
		"d88b4c": 9045, // KingTing
		"d88c73": 3031, // ZTE
		"d88c79": 3521, // Google
		"d88d5c": 4652, // Elentec
		"d88dc8": 9046, // Atil
		"d88f76": 545,  // Apple
		"d890e8": 152,  // Samsung
		"d8912a": 2692, // ZyXEL
		"d89403": 302,  // HP
		"d89685": 6241, // GoPro
		"d89695": 545,  // Apple
		"d897ba": 5439, // Pegatron
		"d89ac1": 457,  // Nokia
		"d89b3b": 3334, // Huawei
		"d89d67": 302,  // HP
		"d89db9": 9047, // eMegatech
		"d89e3f": 545,  // Apple
		"d89e61": 3334, // Huawei
		"d89ed4": 4921, // Fiberhome
		"d89ef3": 102,  // Dell
		"d8a011": 7177, // WiZ
		"d8a01d": 6376, // Espressif
		"d8a105": 9048, // Syslane
		"d8a25e": 545,  // Apple
		"d8a315": 6369, // Vivo
		"d8a35c": 152,  // Samsung
		"d8a491": 3334, // Huawei
		"d8a534": 9049, // Spectronix
		"d8a756": 2048, // Sagemcom
		"d8a8c8": 3031, // ZTE
		"d8aa59": 6894, // Tonly
		"d8addd": 9050, // Sonavation
		"d8ae90": 9051, // Itibia
		"d8aed0": 5008, // Shanghai
		"d8aff1": 2153, // Panasonic
		"d8b053": 5697, // Xiaomi
		"d8b122": 826,  // Juniper
		"d8b12a": 2153, // Panasonic
		"d8b190": 2,    // Cisco
		"d8b377": 1341, // HTC
		"d8b6b7": 3869, // Comtrend
		"d8b6c1": 9052, // NetworkAccountant
		"d8b8f6": 9053, // Nantworks
		"d8bb2c": 545,  // Apple
		"d8bbc1": 1812, // Micro-Star
		"d8be1f": 545,  // Apple
		"d8be65": 5438, // Amazon
		"d8bfc0": 6376, // Espressif
		"d8c068": 9054, // Netgenetech
		"d8c0a6": 3004, // Azurewave
		"d8c3fb": 9055, // Detracom
		"d8c46a": 2056, // Murata
		"d8c497": 3071, // Quanta
		"d8c4e9": 152,  // Samsung
		"d8c561": 9056, // CommFront
		"d8c678": 6428, // MitraStar
		"d8c691": 9057, // Hichan
		"d8c771": 3334, // Huawei
		"d8c7c8": 1685, // Aruba
		"d8c8e9": 6848, // Phicomm
		"d8cb8a": 1812, // Micro-Star
		"d8cc98": 3334, // Huawei
		"d8ce3a": 5697, // Xiaomi
		"d8cf9c": 545,  // Apple
		"d8cfbf": 687,  // Motorola
		"d8d090": 102,  // Dell
		"d8d1cb": 545,  // Apple
		"d8d385": 302,  // HP
		"d8d43c": 101,  // Sony
		"d8d723": 9058, // IDS
		"d8d775": 2048, // Sagemcom
		"d8dc40": 545,  // Apple
		"d8dd5f": 9059, // BALMUDA
		"d8de3a": 545,  // Apple
		"d8dece": 9060, // Isung
		"d8df0d": 9061, // beroNet
		"d8df7a": 2507, // Quest
		"d8e004": 9062, // Vodia
		"d8e0b8": 9063, // Bulat
		"d8e0e1": 152,  // Samsung
		"d8e56d": 6433, // TCT
		"d8e72b": 5515, // Netscout
		"d8e743": 9064, // Wush
		"d8e952": 9065, // Keopsys
		"d8eb46": 3521, // Google
		"d8eb97": 2904, // TRENDnet
		"d8ec5e": 2468, // Belkin
		"d8ece5": 2692, // ZyXEL
		"d8ef42": 3334, // Huawei
		"d8f0f2": 9066, // Zeebo
		"d8f15b": 6376, // Espressif
		"d8f1f0": 9067, // Pepxim
		"d8f2ca": 421,  // Intel
		"d8f3bc": 2873, // Liteon
		"d8f883": 421,  // Intel
		"d8f8af": 9068, // Daontec
		"d8fb11": 9069, // Axacore
		"d8fb5e": 2544, // Askey
		"d8fbd6": 5438, // Amazon
		"d8fc93": 421,  // Intel
		"d8fe8f": 9070, // IDFone
		"d8fee3": 803,  // D-Link
		"dc0077": 1595, // TP-Link
		"dc028e": 3031, // ZTE
		"dc0398": 869,  // LG
		"dc052f": 9071, // National
		"dc0575": 300,  // Siemens
		"dc05ed": 9072, // Nabtesco
		"dc080f": 545,  // Apple
		"dc0914": 4052, // Talk-A-Phone
		"dc094c": 3334, // Huawei
		"dc0b34": 869,  // LG
		"dc0c5c": 545,  // Apple
		"dc0ea1": 358,  // Compal
		"dc16b2": 3334, // Huawei
		"dc1a01": 9073, // Ecoliv
		"dc1ac5": 6369, // Vivo
		"dc1ba1": 421,  // Intel
		"dc1ea3": 9074, // Accensus
		"dc2008": 9075, // ASD
		"dc2148": 421,  // Intel
		"dc215c": 421,  // Intel
		"dc21b9": 9076, // Sentec
		"dc21e2": 3334, // Huawei
		"dc2727": 3334, // Huawei
		"dc2834": 9077, // HAKKO
		"dc2919": 7031, // AltoBeam
		"dc2aa1": 9078, // MedHab
		"dc2b2a": 545,  // Apple
		"dc2b61": 545,  // Apple
		"dc2bca": 9079, // Zera
		"dc2c26": 6868, // Iton
		"dc2c6e": 1784, // Routerboard.com
		"dc2d3c": 3334, // Huawei
		"dc2e6a": 9080, // HCT
		"dc309c": 9081, // Heyrex
		"dc31d1": 6369, // Vivo
		"dc333d": 3334, // Huawei
		"dc3350": 9082, // TechSAT
		"dc3714": 545,  // Apple
		"dc3752": 5810, // GE
		"dc38e1": 826,  // Juniper
		"dc3979": 2,    // Cisco
		"dc3a5e": 1916, // Roku
		"dc3ef8": 457,  // Nokia
		"dc415f": 545,  // Apple
		"dc41a9": 421,  // Intel
		"dc446d": 9083, // Allwinner
		"dc44b6": 152,  // Samsung
		"dc4517": 125,  // ARRIS
		"dc48b2": 9084, // Baraja
		"dc4a3e": 302,  // HP
		"dc4ede": 9085, // Shinyei
		"dc4f22": 6376, // Espressif
		"dc5285": 545,  // Apple
		"dc5360": 421,  // Intel
		"dc537c": 358,  // Compal
		"dc5392": 545,  // Apple
		"dc543d": 6281, // Itel
		"dc54d7": 5438, // Amazon
		"dc56e7": 545,  // Apple
		"dc5726": 1515, // Power-One
		"dc58bc": 9086, // Thomas-Krenn.AG
		"dc5e36": 9087, // Paterson
		"dc663a": 9088, // Apacer
		"dc6672": 152,  // Samsung
		"dc680c": 302,  // HP
		"dc68eb": 1427, // Nintendo
		"dc6aea": 6295, // Infinix
		"dc6b12": 9089, // worldcns
		"dc6b1b": 3334, // Huawei
		"dc6f00": 9090, // Livescribe
		"dc7014": 67,   // Private
		"dc7137": 3031, // ZTE
		"dc7144": 152,  // Samsung
		"dc7196": 421,  // Intel
		"dc729b": 3334, // Huawei
		"dc7385": 3334, // Huawei
		"dc74a8": 152,  // Samsung
		"dc774c": 2,    // Cisco
		"dc7834": 9091, // Logicom
		"dc7b94": 2,    // Cisco
		"dc7fa4": 1939, // 2Wire
		"dc8084": 545,  // Apple
		"dc825b": 9092, // JANUS
		"dc82f6": 9093, // iPort
		"dc85de": 3004, // Azurewave
		"dc86d8": 545,  // Apple
		"dc8983": 152,  // Samsung
		"dc8b28": 421,  // Intel
		"dc8c1b": 6369, // Vivo
		"dc8c37": 2,    // Cisco
		"dc9088": 3334, // Huawei
		"dc9166": 3334, // Huawei
		"dc91bf": 5438, // Amazon
		"dc973a": 9094, // Verana
		"dc9840": 612,  // Microsoft
		"dc9914": 3334, // Huawei
		"dc99fe": 9095, // Armatura
		"dc9b1e": 9096, // Intercom
		"dc9b9c": 545,  // Apple
		"dc9bd6": 6433, // TCT
		"dc9c52": 9097, // Sapphire
		"dc9fa4": 457,  // Nokia
		"dc9fdb": 2978, // Ubiquiti
		"dca120": 457,  // Nokia
		"dca3ac": 9098, // RBcloudtech
		"dca4ca": 545,  // Apple
		"dca5f4": 2,    // Cisco
		"dca633": 125,  // ARRIS
		"dca782": 3334, // Huawei
		"dca904": 545,  // Apple
		"dca971": 421,  // Intel
		"dca989": 9099, // Macandc
		"dcad9e": 9100, // GreenPriz
		"dcae04": 9101, // CELOXICA
		"dcaeeb": 2737, // Ruckus
		"dcb082": 457,  // Nokia
		"dcb4ac": 833,  // Flextronics
		"dcb4c4": 612,  // Microsoft
		"dcb54f": 545,  // Apple
		"dcb72e": 5697, // Xiaomi
		"dcb7fc": 5751, // Alps
		"dcb808": 185,  // Extreme
		"dcbfe9": 687,  // Motorola
		"dcc101": 9102, // SOLiD
		"dcc422": 9103, // Systembase
		"dcc64b": 3334, // Huawei
		"dcc793": 457,  // Nokia
		"dccba8": 9104, // Explora
		"dccce6": 152,  // Samsung
		"dccd2f": 46,   // Epson
		"dccec1": 2,    // Cisco
		"dccf96": 152,  // Samsung
		"dcd0f7": 9105, // Bentek
		"dcd255": 299,  // Kinpo
		"dcd2fc": 3334, // Huawei
		"dcd321": 531,  // HUMAX
		"dcd3a2": 545,  // Apple
		"dcd444": 3334, // Huawei
		"dcd7a0": 3334, // Huawei
		"dcd916": 3334, // Huawei
		"dcda4f": 9106, // Getck
		"dcdc07": 9107, // TRP
		"dcdce2": 152,  // Samsung
		"dcdd24": 9108, // Energica
		"dcdeca": 9109, // Akyllor
		"dcdfd6": 3031, // ZTE
		"dce55b": 3521, // Google
		"dceb69": 6392, // Technicolor
		"dceb94": 2,    // Cisco
		"dced84": 9110, // Haverford
		"dcee06": 3334, // Huawei
		"dcef09": 1368, // Netgear
		"dcef80": 3334, // Huawei
		"dcefca": 2056, // Murata
		"dcf090": 9111, // Nubia
		"dcf110": 457,  // Nokia
		"dcf401": 102,  // Dell
		"dcf4ca": 545,  // Apple
		"dcf505": 3004, // Azurewave
		"dcf56e": 9112, // Wellysis
		"dcf719": 2,    // Cisco
		"dcf755": 9113, // Sitronik
		"dcf756": 152,  // Samsung
		"dcf858": 9114, // Lorent
		"dcf8b9": 3031, // ZTE
		"dcfb02": 1077, // Buffalo
		"dcfb48": 421,  // Intel
		"dcfe07": 5439, // Pegatron
		"dcfe18": 1595, // TP-Link
		"dcfe23": 2056, // Murata
		"e00084": 3334, // Huawei
		"e005c5": 1595, // TP-Link
		"e0071b": 302,  // HP
		"e00af6": 2873, // Liteon
		"e00b28": 9115, // Inovonics
		"e00c7f": 1427, // Nintendo
		"e00ce5": 3334, // Huawei
		"e00db9": 9116, // Cree
		"e00eda": 2,    // Cisco
		"e00ee1": 6290, // We
		"e00ee4": 7446, // DWnet
		"e0107f": 2737, // Ruckus
		"e013b5": 6369, // Vivo
		"e0143e": 9117, // Modoosis
		"e01877": 4,    // Fujitsu
		"e0191d": 3334, // Huawei
		"e01954": 3031, // ZTE
		"e01995": 7340, // Nutanix
		"e019d8": 9118, // BH
		"e01c41": 185,  // Extreme
		"e01cee": 9119, // Bravo
		"e01cfc": 803,  // D-Link
		"e01d3b": 1640, // Cambridge
		"e01f88": 5697, // Xiaomi
		"e02202": 125,  // ARRIS
		"e023ff": 1323, // Fortinet
		"e0247f": 3334, // Huawei
		"e02481": 3334, // Huawei
		"e02630": 9120, // Intrigue
		"e02636": 76,   // Nortel
		"e02861": 3334, // Huawei
		"e02ae6": 4921, // Fiberhome
		"e02b96": 545,  // Apple
		"e02be9": 421,  // Intel
		"e02cb2": 2664, // Lenovo
		"e02cf3": 9121, // MRS
		"e02e3f": 3334, // Huawei
		"e02f6d": 2,    // Cisco
		"e030f9": 826,  // Juniper
		"e0319e": 9122, // Valve
		"e0338e": 545,  // Apple
		"e03676": 3334, // Huawei
		"e03717": 6392, // Technicolor
		"e037bf": 1592, // Wistron
		"e0383f": 3031, // ZTE
		"e039d7": 9123, // Plexxi
		"e03e44": 858,  // Broadcom
		"e03e7d": 9124, // data-Complex
		"e03f49": 1806, // ASUS
		"e04007": 3334, // Huawei
		"e04102": 3031, // ZTE
		"e04136": 6428, // MitraStar
		"e0469a": 1368, // Netgear
		"e048af": 9125, // Premietech
		"e049ed": 9126, // Audeze
		"e04b45": 9127, // Hi-P
		"e04ba6": 3334, // Huawei
		"e05163": 2639, // Arcadyan
		"e056f4": 9128, // AxesNetwork
		"e05b70": 9129, // Innovid
		"e05f45": 545,  // Apple
		"e05fb9": 2,    // Cisco
		"e06066": 2074, // Sercomm
		"e06089": 9130, // Cloudleaf
		"e06267": 5697, // Xiaomi
		"e063da": 2978, // Ubiquiti
		"e063e5": 101,  // Sony
		"e06678": 545,  // Apple
		"e0686d": 9131, // Raybased
		"e0693a": 9132, // Innophase
		"e06995": 5439, // Pegatron
		"e069ba": 2,    // Cisco
		"e06cf6": 9133, // ESSENCORE
		"e06d17": 545,  // Apple
		"e06d18": 9134, // Pioneercorporation
		"e070ea": 302,  // HP
		"e0735f": 5776, // Nucom
		"e0750a": 430,  // Alpsalpine
		"e0757d": 687,  // Motorola
		"e076d0": 4498, // AMPAK
		"e07726": 3334, // Huawei
		"e079c4": 9135, // iRay
		"e07c13": 3031, // ZTE
		"e07c62": 5351, // Whistle
		"e07e5f": 5695, // Renesas
		"e08177": 9136, // GreenBytes
		"e087b1": 9137, // Nata-Info
		"e0885d": 6392, // Technicolor
		"e0897e": 545,  // Apple
		"e0899d": 2,    // Cisco
		"e08a7e": 9138, // Exponent
		"e08e3c": 4854, // Aztech
		"e08fec": 9139, // Repotec
		"e09153": 190,  // XAVi
		"e091f5": 1368, // Netgear
		"e0925c": 545,  // Apple
		"e092a7": 7586, // Feitian
		"e09467": 421,  // Intel
		"e09579": 9140, // ORTHOsoft
		"e09796": 3334, // Huawei
		"e097f2": 9141, // Atomax
		"e09806": 6376, // Espressif
		"e09861": 687,  // Motorola
		"e09971": 152,  // Samsung
		"e09d13": 152,  // Samsung
		"e09d31": 421,  // Intel
		"e09db8": 4478, // Planex
		"e09f2a": 6868, // Iton
		"e0a1d7": 3187, // SFR
		"e0a30f": 9142, // Pevco
		"e0a3ac": 3334, // Huawei
		"e0a509": 9143, // Bitmain
		"e0a670": 457,  // Nokia
		"e0a700": 9144, // Verkada
		"e0aa96": 152,  // Samsung
		"e0aab0": 9145, // Suntaili
		"e0abfe": 9146, // Orb
		"e0accb": 545,  // Apple
		"e0acf1": 2,    // Cisco
		"e0ae5e": 430,  // Alpsalpine
		"e0aea2": 3334, // Huawei
		"e0aeed": 9147, // Loenk
		"e0af4b": 7575, // Pluribus
		"e0b2f1": 6440, // FN-Link
		"e0b52d": 545,  // Apple
		"e0b55f": 545,  // Apple
		"e0b70a": 125,  // ARRIS
		"e0b7b1": 125,  // ARRIS
		"e0b9a5": 3004, // Azurewave
		"e0b9ba": 545,  // Apple
		"e0b9e5": 6392, // Technicolor
		"e0bab4": 9148, // Arrcus
		"e0bb0c": 9149, // Synertau
		"e0bb9e": 46,   // Epson
		"e0c286": 9150, // Aisai
		"e0c2b7": 6362, // Masimo
		"e0c377": 152,  // Samsung
		"e0c3f3": 3031, // ZTE
		"e0c6b3": 9151, // MilDef
		"e0c767": 545,  // Apple
		"e0c97a": 545,  // Apple
		"e0ca94": 2544, // Askey
		"e0cb1d": 67,   // Private
		"e0cb4e": 1806, // ASUS
		"e0cbee": 152,  // Samsung
		"e0cc7a": 3334, // Huawei
		"e0ccf8": 5697, // Xiaomi
		"e0cec3": 2544, // Askey
		"e0cf2d": 9152, // Gemintek
		"e0d045": 421,  // Intel
		"e0d083": 152,  // Samsung
		"e0d10a": 9153, // Katoudenkikougyousyo
		"e0d173": 2,    // Cisco
		"e0d31a": 9154, // EQUES
		"e0d462": 3334, // Huawei
		"e0d464": 421,  // Intel
		"e0d4e8": 421,  // Intel
		"e0d55e": 1929, // Giga-Byte
		"e0d848": 102,  // Dell
		"e0d9e3": 8495, // Eltex
		"e0da90": 3334, // Huawei
		"e0db10": 152,  // Samsung
		"e0db55": 102,  // Dell
		"e0dbd1": 6392, // Technicolor
		"e0dca0": 300,  // Siemens
		"e0dcff": 5697, // Xiaomi
		"e0ddc0": 6369, // Vivo
		"e0e0fc": 3334, // Huawei
		"e0e2e6": 6376, // Espressif
		"e0e37c": 3334, // Huawei
		"e0e62e": 6433, // TCT
		"e0e631": 9155, // SNB
		"e0e751": 1427, // Nintendo
		"e0e7bb": 9156, // Nureva
		"e0eb40": 545,  // Apple
		"e0ed1a": 9157, // vastriver
		"e0ee1b": 2153, // Panasonic
		"e0ef25": 9158, // Lintes
		"e0f211": 9159, // Digitalwatt
		"e0f379": 9160, // Vaddio
		"e0f442": 3334, // Huawei
		"e0f5c6": 545,  // Apple
		"e0f62d": 826,  // Juniper
		"e0f6b5": 1427, // Nintendo
		"e0f847": 545,  // Apple
		"e0f9be": 9161, // Cloudena
		"e0fff7": 9162, // Softiron
		"e4029b": 421,  // Intel
		"e40439": 2714, // TomTom
		"e405f8": 9163, // Bytedance
		"e4072b": 3334, // Huawei
		"e40eee": 3334, // Huawei
		"e4115b": 302,  // HP
		"e4121d": 152,  // Samsung
		"e4186b": 2692, // ZyXEL
		"e419c1": 3334, // Huawei
		"e41a2c": 9164, // ZPE
		"e41c4b": 9165, // V2
		"e41d2d": 432,  // Mellanox
		"e41f13": 372,  // IBM
		"e41f7b": 2,    // Cisco
		"e41fe9": 9166, // Dunkermotoren
		"e422a5": 542,  // Plantronics
		"e425e7": 545,  // Apple
		"e425e9": 9167, // Color-Chip
		"e42686": 7446, // DWnet
		"e4268b": 3334, // Huawei
		"e42771": 2133, // Smartlabs
		"e42aac": 612,  // Microsoft
		"e42b34": 545,  // Apple
		"e42b79": 457,  // Nokia
		"e42c56": 9168, // Lilee
		"e42d02": 6433, // TCT
		"e42f26": 4921, // Fiberhome
		"e42f56": 9169, // OptoMET
		"e42ff6": 9170, // Unicore
		"e432cb": 152,  // Samsung
		"e43493": 3334, // Huawei
		"e435c8": 3334, // Huawei
		"e435fb": 9171, // Sabre
		"e43883": 2978, // Ubiquiti
		"e4388c": 1565, // Digital
		"e43a65": 9172, // MofiNetwork
		"e43d1a": 858,  // Broadcom
		"e43ec6": 3334, // Huawei
		"e43ed7": 2639, // Arcadyan
		"e440e2": 152,  // Samsung
		"e44122": 6923, // OnePlus
		"e44164": 457,  // Nokia
		"e441e6": 9173, // Ottec
		"e442a6": 421,  // Intel
		"e4434b": 102,  // Dell
		"e444e5": 185,  // Extreme
		"e446da": 5697, // Xiaomi
		"e447b3": 3031, // ZTE
		"e44e2d": 2,    // Cisco
		"e44e76": 9174, // Championtech
		"e4509a": 9175, // HW
		"e450eb": 545,  // Apple
		"e454e8": 102,  // Dell
		"e45740": 125,  // ARRIS
		"e457a8": 9176, // Stuart
		"e458b8": 152,  // Samsung
		"e458e7": 152,  // Samsung
		"e45aa2": 6369, // Vivo
		"e45ad4": 8495, // Eltex
		"e45d37": 826,  // Juniper
		"e45d51": 3187, // SFR
		"e45d52": 620,  // Avaya
		"e45d75": 152,  // Samsung
		"e45e1b": 3521, // Google
		"e45e37": 421,  // Intel
		"e46059": 9177, // Pingtek
		"e46449": 125,  // ARRIS
		"e468a3": 3334, // Huawei
		"e46c21": 9178, // messMa
		"e46f13": 803,  // D-Link
		"e470b8": 421,  // Intel
		"e47185": 9179, // Securifi
		"e472e2": 3334, // Huawei
		"e475dc": 2639, // Arcadyan
		"e47684": 545,  // Apple
		"e47723": 3031, // ZTE
		"e47727": 3334, // Huawei
		"e4776b": 9180, // Aartesys
		"e477d4": 9181, // Minrray
		"e47b3f": 9182, // Beijing-Cloud
		"e47c65": 9183, // Sunstar
		"e47cf9": 152,  // Samsung
		"e47dbd": 152,  // Samsung
		"e47e66": 3334, // Huawei
		"e47e9a": 3031, // ZTE
		"e47fb2": 4,    // Fujitsu
		"e48184": 457,  // Nokia
		"e48210": 3334, // Huawei
		"e482cc": 9184, // Jumptronic
		"e48326": 3334, // Huawei
		"e48399": 125,  // ARRIS
		"e484d3": 5697, // Xiaomi
		"e48501": 9185, // Geberit
		"e48b7f": 545,  // Apple
		"e48d8c": 1784, // Routerboard.com
		"e48f1d": 3334, // Huawei
		"e4907e": 687,  // Motorola
		"e490fd": 545,  // Apple
		"e4922a": 9186, // DBG
		"e492e7": 9187, // Gridlink
		"e492fb": 152,  // Samsung
		"e496ae": 9188, // ALTOGRAPHICS
		"e498d1": 612,  // Microsoft
		"e498d6": 545,  // Apple
		"e49a79": 545,  // Apple
		"e49adc": 545,  // Apple
		"e49f1e": 125,  // ARRIS
		"e4a387": 1998, // Control
		"e4a471": 421,  // Intel
		"e4a7a0": 421,  // Intel
		"e4a7c5": 3334, // Huawei
		"e4a8b6": 3334, // Huawei
		"e4a8df": 358,  // Compal
		"e4aa5d": 2,    // Cisco
		"e4aaea": 2873, // Liteon
		"e4ab89": 6428, // MitraStar
		"e4afa1": 9189, // HES-SO
		"e4b021": 152,  // Samsung
		"e4b2fb": 545,  // Apple
		"e4b318": 421,  // Intel
		"e4b97a": 102,  // Dell
		"e4bd4b": 3031, // ZTE
		"e4beed": 2362, // Netcore
		"e4bffa": 6392, // Technicolor
		"e4c0e2": 2048, // Sagemcom
		"e4c2d1": 3334, // Huawei
		"e4c32a": 1595, // TP-Link
		"e4c62b": 9190, // Airware
		"e4c63d": 545,  // Apple
		"e4c6e6": 9191, // Mophie
		"e4c722": 2,    // Cisco
		"e4c801": 7460, // BLU
		"e4c90b": 9192, // Radwin
		"e4ca12": 3031, // ZTE
		"e4ce02": 9193, // WyreStorm
		"e4ce8f": 545,  // Apple
		"e4d124": 2485, // Mojo
		"e4d332": 1595, // TP-Link
		"e4d373": 3334, // Huawei
		"e4d3f1": 2,    // Cisco
		"e4dc43": 3334, // Huawei
		"e4dc5f": 9194, // Cofractal
		"e4dccc": 3334, // Huawei
		"e4e0a6": 545,  // Apple
		"e4e0c5": 152,  // Samsung
		"e4e130": 6433, // TCT
		"e4e409": 9195, // Leifheit
		"e4e4ab": 545,  // Apple
		"e4e749": 302,  // HP
		"e4ec10": 457,  // Nokia
		"e4eefd": 9196, // MR&D
		"e4f004": 102,  // Dell
		"e4f042": 3521, // Google
		"e4f14c": 67,   // Private
		"e4f1d4": 6369, // Vivo
		"e4f327": 9197, // Atol
		"e4f365": 9198, // Time-O-Matic
		"e4f3c4": 152,  // Samsung
		"e4f4c6": 1368, // Netgear
		"e4f75b": 125,  // ARRIS
		"e4f7a1": 9199, // Datafox
		"e4f89c": 421,  // Intel
		"e4f8ef": 152,  // Samsung
		"e4faed": 152,  // Samsung
		"e4fafd": 421,  // Intel
		"e4fb5d": 3334, // Huawei
		"e4fc82": 826,  // Juniper
		"e4fd45": 421,  // Intel
		"e4fda1": 3334, // Huawei
		"e4fed9": 6661, // EDMI
		"e80036": 9200, // Befs
		"e8018d": 4921, // Fiberhome
		"e8039a": 152,  // Samsung
		"e8040b": 545,  // Apple
		"e80410": 67,   // Private
		"e80462": 2,    // Cisco
		"e804f3": 9201, // Throughtek
		"e8056d": 76,   // Nortel
		"e805dc": 1647, // Verifone
		"e80688": 545,  // Apple
		"e8088b": 3334, // Huawei
		"e80c75": 9202, // Syncbak
		"e80fc8": 3857, // Universal
		"e81132": 152,  // Samsung
		"e81367": 9203, // AIRSOUND
		"e8136e": 3334, // Huawei
		"e8150e": 457,  // Nokia
		"e81a58": 6022, // Technologic
		"e81b4b": 9204, // amnimo
		"e81b69": 2074, // Sercomm
		"e81cba": 1323, // Fortinet
		"e81cd8": 545,  // Apple
		"e81da8": 2737, // Ruckus
		"e81e92": 3334, // Huawei
		"e820e2": 531,  // HUMAX
		"e82689": 1685, // Aruba
		"e82877": 9205, // TMY
		"e828c1": 8495, // Eltex
		"e828d5": 9206, // Cots
		"e82a44": 2873, // Liteon
		"e82aea": 421,  // Intel
		"e82c6d": 4549, // SmartRG
		"e82e0c": 9207, // NETINT
		"e831cd": 6376, // Espressif
		"e8330d": 9208, // Xaptec
		"e83381": 125,  // ARRIS
		"e83617": 545,  // Apple
		"e8361d": 9209, // Sense
		"e8377a": 2692, // ZyXEL
		"e83935": 302,  // HP
		"e839df": 2544, // Askey
		"e83a12": 152,  // Samsung
		"e83a97": 35,   // Toshiba
		"e83eb6": 4555, // RIM
		"e83efb": 9210, // Geodesic
		"e83efc": 125,  // ARRIS
		"e83f67": 3334, // Huawei
		"e84040": 2,    // Cisco
		"e840f2": 5439, // Pegatron
		"e843b6": 6761, // QNAP
		"e848b8": 1595, // TP-Link
		"e84c56": 9211, // Intercept
		"e84d74": 3334, // Huawei
		"e84dd0": 3334, // Huawei
		"e84dec": 0,    // Xerox
		"e84e06": 9212, // Edup
		"e84e84": 152,  // Samsung
		"e84ece": 1427, // Nintendo
		"e84f25": 2056, // Murata
		"e84fa7": 3334, // Huawei
		"e8508b": 152,  // Samsung
		"e8516e": 9213, // TSMART
		"e855b4": 5501, // SAI
		"e85659": 9214, // Advanced-Connectek
		"e856d6": 8351, // NCTech
		"e85a8b": 5697, // Xiaomi
		"e85ad1": 4921, // Fiberhome
		"e85b5b": 869,  // LG
		"e85bb7": 2062, // Ample
		"e85c0a": 2,    // Cisco
		"e85f02": 545,  // Apple
		"e8617e": 2873, // Liteon
		"e861be": 9215, // Melec
		"e86549": 2,    // Cisco
		"e865d4": 6266, // Tenda
		"e866c4": 9216, // Diamanti
		"e86819": 3334, // Huawei
		"e868e7": 6376, // Espressif
		"e86a64": 4919, // LCFC
		"e86d52": 125,  // ARRIS
		"e86d54": 3176, // Digit
		"e86dcb": 152,  // Samsung
		"e86de9": 3334, // Huawei
		"e86e44": 3031, // ZTE
		"e86ff2": 2246, // Actiontec
		"e874c7": 9217, // Sentinhealth
		"e8757f": 9218, // FIRS
		"e87865": 545,  // Apple
		"e87f6b": 152,  // Samsung
		"e87f95": 545,  // Apple
		"e8802e": 545,  // Apple
		"e880d8": 9219, // GNTEK
		"e88152": 545,  // Apple
		"e88175": 3031, // ZTE
		"e8825b": 125,  // ARRIS
		"e884a5": 421,  // Intel
		"e884c6": 3334, // Huawei
		"e8854b": 545,  // Apple
		"e8892c": 125,  // ARRIS
		"e88d28": 545,  // Apple
		"e88df5": 2477, // ZNYX
		"e88e60": 9220, // NSD
		"e8910f": 4921, // Fiberhome
		"e89120": 687,  // Motorola
		"e89218": 9221, // Arcontia
		"e892a4": 869,  // LG
		"e89309": 152,  // Samsung
		"e89363": 457,  // Nokia
		"e893f3": 9222, // Graphiant
		"e894f6": 1595, // TP-Link
		"e898c2": 9223, // ZETLAB
		"e8995a": 9224, // PiiGAB
		"e899c4": 1341, // HTC
		"e89a8f": 3071, // Quanta
		"e89d87": 35,   // Toshiba
		"e89f39": 457,  // Nokia
		"e89f6d": 6376, // Espressif
		"e89f80": 2468, // Belkin
		"e8a0cd": 1427, // Nintendo
		"e8a1f8": 3031, // ZTE
		"e8a245": 826,  // Juniper
		"e8a660": 3334, // Huawei
		"e8a6ca": 3334, // Huawei
		"e8a72f": 612,  // Microsoft
		"e8a730": 545,  // Apple
		"e8a7f2": 9225, // sTraffic
		"e8abf3": 3334, // Huawei
		"e8acad": 3031, // ZTE
		"e8ada6": 2048, // Sagemcom
		"e8aec5": 3793, // Arista
		"e8b1fc": 421,  // Intel
		"e8b2ac": 545,  // Apple
		"e8b2fe": 531,  // HUMAX
		"e8b4c8": 152,  // Samsung
		"e8b541": 3031, // ZTE
		"e8b5d0": 102,  // Dell
		"e8b6c2": 826,  // Juniper
		"e8b748": 2,    // Cisco
		"e8ba70": 2,    // Cisco
		"e8bdd1": 3334, // Huawei
		"e8be81": 2048, // Sagemcom
		"e8c1d7": 796,  // Philips
		"e8c2dd": 6295, // Infinix
		"e8c417": 4921, // Fiberhome
		"e8c57a": 9226, // Ufispace
		"e8c74f": 2873, // Liteon
		"e8c7cf": 1592, // Wistron
		"e8cba1": 457,  // Nokia
		"e8cbed": 6439, // Chipsea
		"e8cc18": 803,  // D-Link
		"e8cc32": 2455, // Micronet
		"e8cd2d": 3334, // Huawei
		"e8ce06": 8025, // SkyHawke
		"e8d099": 4921, // Fiberhome
		"e8d0fc": 2873, // Liteon
		"e8d11b": 2544, // Askey
		"e8d2ff": 2048, // Sagemcom
		"e8d322": 2,    // Cisco
		"e8d765": 3334, // Huawei
		"e8d819": 3004, // Azurewave
		"e8d8d1": 302,  // HP
		"e8da00": 9227, // Kivo
		"e8da20": 1427, // Nintendo
		"e8daaa": 9228, // VideoHome
		"e8db84": 6376, // Espressif
		"e8de27": 1595, // TP-Link
		"e8ded6": 9229, // Intrising
		"e8dff2": 9230, // PRF
		"e8e0b7": 35,   // Toshiba
		"e8e1e1": 1450, // Gemtek
		"e8e1e2": 9231, // Energotest
		"e8e5d6": 152,  // Samsung
		"e8e875": 9232, // iS5
		"e8e8b7": 2056, // Murata
		"e8ea4d": 3334, // Huawei
		"e8ea6a": 9233, // StarTech.com
		"e8eb1b": 706,  // Microchip
		"e8eb34": 2,    // Cisco
		"e8ebd3": 432,  // Mellanox
		"e8ed05": 125,  // ARRIS
		"e8edd6": 1323, // Fortinet
		"e8edf3": 2,    // Cisco
		"e8ef89": 9234, // OPMEX
		"e8f1b0": 2048, // Sagemcom
		"e8f2e2": 869,  // LG
		"e8f375": 457,  // Nokia
		"e8f408": 421,  // Intel
		"e8f654": 3334, // Huawei
		"e8f724": 302,  // HP
		"e8f9d4": 3334, // Huawei
		"e8fa23": 3334, // Huawei
		"e8fbe9": 545,  // Apple
		"e8fcaf": 1368, // Netgear
		"e8fd35": 3334, // Huawei
		"e8fd90": 9235, // Turbostor
		"ec0133": 9236, // Trinus
		"ec01d5": 2,    // Cisco
		"ec0273": 1685, // Aruba
		"ec086b": 1595, // TP-Link
		"ec08e5": 687,  // Motorola
		"ec0d9a": 432,  // Mellanox
		"ec0de4": 5438, // Amazon
		"ec107b": 152,  // Samsung
		"ec13b2": 9237, // Netonix
		"ec13db": 826,  // Juniper
		"ec14f6": 9238, // BioControl
		"ec172f": 1595, // TP-Link
		"ec1a59": 2468, // Belkin
		"ec1bbd": 1655, // Silicon
		"ec1c5d": 300,  // Siemens
		"ec1d7f": 3031, // ZTE
		"ec1d8b": 2,    // Cisco
		"ec1f72": 152,  // Samsung
		"ec219f": 9239, // VidaBox
		"ec21e5": 35,   // Toshiba
		"ec2280": 803,  // D-Link
		"ec233d": 3334, // Huawei
		"ec2368": 9240, // IntelliVoice
		"ec237b": 3031, // ZTE
		"ec2651": 545,  // Apple
		"ec26ca": 1595, // TP-Link
		"ec26fb": 9241, // Tecc
		"ec2a72": 102,  // Dell
		"ec2af0": 9242, // Ypsomed
		"ec2beb": 5438, // Amazon
		"ec2ce2": 545,  // Apple
		"ec2e98": 3004, // Azurewave
		"ec3091": 2,    // Cisco
		"ec316d": 9243, // Hansgrohe
		"ec354d": 4030, // Wingtech
		"ec3586": 545,  // Apple
		"ec363f": 9244, // Markov
		"ec3873": 826,  // Juniper
		"ec388f": 3334, // Huawei
		"ec3a52": 3334, // Huawei
		"ec3bf0": 9245, // NovelSat
		"ec3c88": 9246, // MCNEX
		"ec3cbb": 3334, // Huawei
		"ec3eb3": 2692, // ZyXEL
		"ec3ef7": 826,  // Juniper
		"ec4118": 6280, // XIAOMI
		"ec42b4": 9247, // ADC
		"ec438b": 9248, // Yaptv
		"ec43e6": 9249, // AWCER
		"ec43f6": 2692, // ZyXEL
		"ec4476": 2,    // Cisco
		"ec473c": 9250, // Redwire
		"ec4993": 9251, // Qihan
		"ec4d47": 3334, // Huawei
		"ec4f82": 374,  // Calix
		"ec50aa": 1685, // Aruba
		"ec5623": 3334, // Huawei
		"ec570d": 5107, // AFE
		"ec58ea": 2737, // Ruckus
		"ec59e7": 612,  // Microsoft
		"ec5a86": 3096, // Yulong
		"ec5c84": 2056, // Murata
		"ec6073": 1595, // TP-Link
		"ec60e0": 9252, // AVI-ON
		"ec63d7": 421,  // Intel
		"ec64e7": 9253, // MOCACARE
		"ec656e": 9254, // Things
		"ec65cc": 2153, // Panasonic
		"ec6c9a": 2639, // Arcadyan
		"ec6cb5": 3031, // ZTE
		"ec6f0b": 9255, // FADU
		"ec7097": 125,  // ARRIS
		"ec753e": 3334, // Huawei
		"ec75ed": 9256, // Citrix
		"ec7949": 4,    // Fujitsu
		"ec79f2": 9257, // Startel
		"ec7c2c": 3334, // Huawei
		"ec7c5c": 826,  // Juniper
		"ec7c74": 9258, // Justone
		"ec7cb6": 152,  // Samsung
		"ec7d11": 6369, // Vivo
		"ec7d9d": 9259, // CPI
		"ec7e91": 6281, // Itel
		"ec8009": 9260, // NovaSparks
		"ec8193": 4079, // Logitech
		"ec8263": 3031, // ZTE
		"ec8350": 612,  // Microsoft
		"ec836c": 9261, // RM
		"ec83d5": 9262, // GIRD
		"ec852f": 545,  // Apple
		"ec888f": 1595, // TP-Link
		"ec8892": 687,  // Motorola
		"ec8914": 3334, // Huawei
		"ec89f5": 2664, // Lenovo
		"ec8a4c": 3031, // ZTE
		"ec8ac4": 5438, // Amazon
		"ec8ac7": 4921, // Fiberhome
		"ec8c9a": 3334, // Huawei
		"ec8ca2": 2737, // Ruckus
		"ec8ead": 9263, // DLX
		"ec8eae": 9264, // Nagravision
		"ec8eb5": 302,  // HP
		"ec9365": 9265, // Mapper.ai
		"ec937d": 6392, // Technicolor
		"ec93ed": 9266, // DDoS-Guard
		"ec94cb": 6376, // Espressif
		"ec94d5": 826,  // Juniper
		"ec9a74": 302,  // HP
		"ec9b5b": 457,  // Nokia
		"ec9b8b": 302,  // HP
		"ec9bf3": 152,  // Samsung
		"eca1d1": 3334, // Huawei
		"eca29b": 9267, // Kemppi
		"eca81f": 6392, // Technicolor
		"eca86b": 1115, // Elitegroup
		"eca907": 545,  // Apple
		"eca940": 125,  // ARRIS
		"ecaa25": 152,  // Samsung
		"ecaaa0": 5439, // Pegatron
		"ecadb8": 545,  // Apple
		"ecade0": 803,  // D-Link
		"ecaf97": 9268, // GIT
		"ecb0e1": 4564, // Ciena
		"ecb106": 9269, // Acuro
		"ecb1d7": 302,  // HP
		"ecb4e8": 1592, // Wistron
		"ecb907": 9270, // CloudGenix
		"ecb970": 5440, // Ruijie
		"ecbafe": 9271, // Giroptic
		"ecbd09": 9272, // FUSION
		"ecbd1d": 2,    // Cisco
		"ecbe5f": 1449, // Vestel
		"ecbedd": 2048, // Sagemcom
		"ecc01b": 3334, // Huawei
		"ecc302": 531,  // HUMAX
		"ecc38a": 9273, // Accuenergy
		"ecc40d": 1427, // Nintendo
		"ecc5d2": 3334, // Huawei
		"ecc882": 2,    // Cisco
		"eccb30": 3334, // Huawei
		"ecce13": 2,    // Cisco
		"ecced7": 545,  // Apple
		"ecd00e": 9274, // MiraeRecognition
		"ecd09f": 5697, // Xiaomi
		"ecd925": 9275, // Rami
		"ecd950": 2618, // IRT
		"ecdb86": 9276, // API-K
		"ecde3d": 9277, // Lamprey
		"ecdf3a": 6369, // Vivo
		"ece09b": 152,  // Samsung
		"ece1a9": 2,    // Cisco
		"ece512": 9278, // tado
		"ece744": 9279, // Omntec
		"ece7a7": 421,  // Intel
		"ece915": 9280, // STI
		"ecebb8": 302,  // HP
		"ecf00e": 2556, // AboCom
		"ecf0fe": 3031, // ZTE
		"ecf22b": 6265, // Tecno
		"ecf236": 9281, // Neomontana
		"ecf35b": 457,  // Nokia
		"ecf40c": 2,    // Cisco
		"ecf451": 2639, // Arcadyan
		"ecf4bb": 102,  // Dell
		"ecfa03": 9282, // FCA
		"ecfaaa": 9283, // IMS
		"ecfabc": 6376, // Espressif
		"ecfaf4": 9284, // SenRa
		"ecfe7e": 9285, // BlueRadios
		"f0022b": 9286, // Chrontel
		"f00248": 9287, // SmarteBuilding
		"f0038c": 3004, // Azurewave
		"f008d1": 6376, // Espressif
		"f008f1": 152,  // Samsung
		"f00d5c": 6945, // JinQianMao
		"f00e1d": 6710, // Megafone
		"f00ebf": 9288, // ZettaHash
		"f00fec": 3334, // Huawei
		"f013c1": 9289, // Hannto
		"f015b9": 9290, // PlayFusion
		"f01628": 6392, // Technicolor
		"f0182b": 869,  // LG
		"f01898": 545,  // Apple
		"f01b6c": 6369, // Vivo
		"f01c13": 869,  // LG
		"f01c2d": 826,  // Juniper
		"f01d2d": 2,    // Cisco
		"f01dbc": 612,  // Microsoft
		"f01e34": 9291, // ORICO
		"f01faf": 102,  // Dell
		"f0219d": 1630, // Cal-Comp
		"f021e0": 5820, // eero
		"f0224e": 9292, // Esan
		"f023ae": 4498, // AMPAK
		"f02408": 65,   // Talaris
		"f02475": 545,  // Apple
		"f02572": 2,    // Cisco
		"f0258e": 3334, // Huawei
		"f025b7": 152,  // Samsung
		"f02624": 9293, // Wafa
		"f0264c": 9294, // Sigrist-Photometer
		"f0272d": 5438, // Amazon
		"f02745": 9295, // F-Secure
		"f02765": 2056, // Murata
		"f02929": 2,    // Cisco
		"f02a61": 9296, // Waldo
		"f02e51": 3170, // Casa
		"f02f4b": 545,  // Apple
		"f02f74": 1806, // ASUS
		"f02fa7": 3334, // Huawei
		"f02fd8": 9297, // Bi2-Vision
		"f0321a": 1933, // Mita-Teknik
		"f033e5": 3334, // Huawei
		"f03404": 6433, // TCT
		"f037a1": 9298, // Huike
		"f03965": 152,  // Samsung
		"f03a4b": 9299, // Bloombase
		"f03d03": 6265, // Tecno
		"f03d29": 9300, // Actility
		"f03e90": 2737, // Ruckus
		"f03f95": 3334, // Huawei
		"f0407b": 4921, // Fiberhome
		"f041c6": 9301, // Heat
		"f0421c": 421,  // Intel
		"f042f5": 3334, // Huawei
		"f04335": 9302, // DVN
		"f04347": 3334, // Huawei
		"f04a02": 2,    // Cisco
		"f04a2b": 9303, // PYRAMID
		"f04b3a": 826,  // Juniper
		"f04bf2": 9304, // JTECH
		"f04cd5": 5298, // Maxlinear
		"f04da2": 102,  // Dell
		"f04f7c": 5438, // Amazon
		"f05136": 6433, // TCT
		"f051ea": 6587, // Fitbit
		"f05501": 3334, // Huawei
		"f057a6": 421,  // Intel
		"f05849": 9305, // CareView
		"f05a09": 152,  // Samsung
		"f05b7b": 152,  // Samsung
		"f05c19": 1685, // Aruba
		"f05c77": 3521, // Google
		"f05cd5": 545,  // Apple
		"f05d89": 9306, // Dycon
		"f061c0": 1685, // Aruba
		"f063f9": 3334, // Huawei
		"f06426": 185,  // Extreme
		"f065ae": 152,  // Samsung
		"f065dd": 389,  // Primax
		"f06853": 5873, // Integrated
		"f06bca": 152,  // Samsung
		"f06c73": 457,  // Nokia
		"f06e0b": 612,  // Microsoft
		"f06f46": 9307, // Ubiik
		"f0704f": 152,  // Samsung
		"f0728c": 152,  // Samsung
		"f072ea": 3521, // Google
		"f07485": 9308, // NGD
		"f074e4": 9309, // Thundercomm
		"f0761c": 358,  // Compal
		"f0766f": 545,  // Apple
		"f07765": 9310, // Sourcefire
		"f077c3": 421,  // Intel
		"f077d0": 9311, // Xcellen
		"f07807": 545,  // Apple
		"f07816": 2,    // Cisco
		"f07959": 1806, // ASUS
		"f07960": 545,  // Apple
		"f07cc7": 826,  // Juniper
		"f07d68": 803,  // D-Link
		"f07f06": 2,    // Cisco
		"f08173": 5438, // Amazon
		"f08175": 2048, // Sagemcom
		"f08261": 2048, // Sagemcom
		"f084c9": 3031, // ZTE
		"f08620": 2639, // Arcadyan
		"f08a76": 152,  // Samsung
		"f08bfe": 9312, // Costel
		"f08cfb": 4921, // Fiberhome
		"f08edb": 9313, // VeloCloud
		"f0921c": 302,  // HP
		"f0933a": 9314, // NxtConect
		"f093c5": 9315, // Garland
		"f097e5": 9316, // Tamio
		"f09838": 3334, // Huawei
		"f0989d": 545,  // Apple
		"f09919": 797,  // Garmin
		"f099b6": 545,  // Apple
		"f099bf": 545,  // Apple
		"f09bb8": 3334, // Huawei
		"f09cbb": 9317, // RaonThink
		"f09ce9": 185,  // Extreme
		"f09e4a": 421,  // Intel
		"f09e63": 2,    // Cisco
		"f09fc2": 2978, // Ubiquiti
		"f09ffc": 3205, // Sharp
		"f0a225": 5438, // Amazon
		"f0a35a": 545,  // Apple
		"f0a764": 9318, // GST
		"f0a7b2": 9319, // Futaba
		"f0a951": 3334, // Huawei
		"f0a968": 6503, // Antailiye
		"f0aca4": 9320, // HBC-radiomatic
		"f0ad4e": 9321, // Globalscale
		"f0ae51": 9322, // Xi3
		"f0af85": 125,  // ARRIS
		"f0b022": 8973, // TOHO
		"f0b052": 2737, // Ruckus
		"f0b0e7": 545,  // Apple
		"f0b107": 306,  // Ericsson
		"f0b11d": 457,  // Nokia
		"f0b13f": 3334, // Huawei
		"f0b2b9": 421,  // Intel
		"f0b2e5": 2,    // Cisco
		"f0b31e": 3857, // Universal
		"f0b3ec": 545,  // Apple
		"f0b429": 5697, // Xiaomi
		"f0b479": 545,  // Apple
		"f0b4d2": 803,  // D-Link
		"f0b5b7": 8289, // Disruptive
		"f0b61e": 421,  // Intel
		"f0b6eb": 9323, // Poslab
		"f0b968": 6281, // Itel
		"f0bcc8": 9324, // MaxID
		"f0bcc9": 5461, // PFU
		"f0bdf1": 9325, // Sipod
		"f0bf97": 101,  // Sony
		"f0c1f1": 545,  // Apple
		"f0c371": 545,  // Apple
		"f0c42f": 3334, // Huawei
		"f0c478": 3334, // Huawei
		"f0c850": 3334, // Huawei
		"f0c88c": 9326, // LeddarTech
		"f0c8b5": 3334, // Huawei
		"f0cba1": 545,  // Apple
		"f0cd31": 152,  // Samsung
		"f0d08c": 6433, // TCT
		"f0d14f": 9327, // Linear
		"f0d1a9": 545,  // Apple
		"f0d1b8": 9328, // Ledvance
		"f0d2f1": 5438, // Amazon
		"f0d3a7": 9329, // CobaltRay
		"f0d3e7": 9330, // Sensometrix
		"f0d4e2": 102,  // Dell
		"f0d5bf": 421,  // Intel
		"f0d657": 9331, // Echosens
		"f0d7aa": 687,  // Motorola
		"f0d7dc": 9332, // Wesine
		"f0da7c": 9333, // RLH
		"f0db30": 9334, // Yottabyte
		"f0dbe2": 545,  // Apple
		"f0dbf8": 545,  // Apple
		"f0dce2": 545,  // Apple
		"f0def1": 1592, // Wistron
		"f0e4a2": 3334, // Huawei
		"f0e77e": 152,  // Samsung
		"f0ec39": 9335, // Essec
		"f0ee10": 152,  // Samsung
		"f0eebb": 9336, // VIPAR
		"f0ef86": 3521, // Google
		"f0f08f": 9337, // Nextek
		"f0f0a4": 5438, // Amazon
		"f0f249": 870,  // Hitron
		"f0f260": 9338, // Mobitec
		"f0f336": 1595, // TP-Link
		"f0f564": 152,  // Samsung
		"f0f5ae": 9339, // Adaptrum
		"f0f61c": 545,  // Apple
		"f0f6c1": 2047, // Sonos
		"f0f755": 2,    // Cisco
		"f0f7b3": 9340, // Phorm
		"f0f7e7": 3334, // Huawei
		"f0f842": 9341, // KEEBOX
		"f0f9f7": 9342, // IES
		"f0fac7": 3334, // Huawei
		"f0fcc8": 125,  // ARRIS
		"f0fda0": 9343, // Acurix
		"f0fee7": 3334, // Huawei
		"f40223": 3217, // PAX
		"f40228": 152,  // Samsung
		"f40270": 102,  // Dell
		"f40304": 3521, // Google
		"f4032a": 5438, // Amazon
		"f4032f": 9344, // Reduxio
		"f40343": 302,  // HP
		"f4044c": 9345, // ValenceTech
		"f40616": 545,  // Apple
		"f40669": 421,  // Intel
		"f4068d": 1637, // devolo
		"f409d8": 152,  // Samsung
		"f40a4a": 9346, // IndUSNET
		"f40b93": 2220, // BlackBerry
		"f40e01": 545,  // Apple
		"f40e22": 152,  // Samsung
		"f40e83": 125,  // ARRIS
		"f40f1b": 2,    // Cisco
		"f40f24": 545,  // Apple
		"f40f9b": 9347, // Wavelink
		"f412fa": 6376, // Espressif
		"f41535": 9348, // SPON
		"f41563": 289,  // F5
		"f419e2": 9349, // Volterra
		"f41ba1": 545,  // Apple
		"f41d6b": 3334, // Huawei
		"f41e26": 9350, // Simon-Kaloi
		"f41e5e": 9351, // RtBrick
		"f41f0b": 9352, // YAMABISHI
		"f41f88": 3031, // ZTE
		"f41fc2": 2,    // Cisco
		"f42012": 9353, // Cuciniale
		"f4239c": 6255, // Sernet
		"f42679": 421,  // Intel
		"f42833": 9354, // MMPC
		"f42853": 2127, // Zioncom
		"f42981": 6369, // Vivo
		"f42a7d": 1595, // TP-Link
		"f42b48": 9355, // Ubiqam
		"f42c56": 9356, // Senor
		"f42e7f": 1685, // Aruba
		"f4308b": 5697, // Xiaomi
		"f430b9": 302,  // HP
		"f431c3": 545,  // Apple
		"f434f0": 545,  // Apple
		"f437b7": 545,  // Apple
		"f438c1": 3334, // Huawei
		"f43909": 302,  // HP
		"f43d80": 9357, // FAG
		"f43e9d": 9358, // Benu
		"f44156": 9359, // Arrikto
		"f4419e": 3334, // Huawei
		"f4428f": 152,  // Samsung
		"f44450": 9360, // BND
		"f44588": 3334, // Huawei
		"f44637": 421,  // Intel
		"f44955": 9361, // MIMO
		"f449ef": 9362, // Emstone
		"f44c7f": 3334, // Huawei
		"f44d30": 1115, // Elitegroup
		"f44e05": 2,    // Cisco
		"f44e38": 9363, // Olibra
		"f44ee3": 421,  // Intel
		"f450eb": 7075, // Telechips
		"f45214": 432,  // Mellanox
		"f45595": 9364, // HENGBAO
		"f4559c": 3334, // Huawei
		"f4573e": 4921, // Fiberhome
		"f45c89": 545,  // Apple
		"f45f69": 9365, // Matsufu
		"f45ff7": 9366, // DQ
		"f4600d": 9367, // Panoptic
		"f460e2": 5697, // Xiaomi
		"f4631f": 3334, // Huawei
		"f46349": 9368, // Diffon
		"f4645d": 35,   // Toshiba
		"f465a6": 545,  // Apple
		"f46942": 2544, // Askey
		"f46abc": 9369, // Adonit
		"f46ad7": 612,  // Microsoft
		"f46bef": 2048, // Sagemcom
		"f46d04": 1806, // ASUS
		"f46d2f": 1595, // TP-Link
		"f46de2": 3031, // ZTE
		"f46e95": 185,  // Extreme
		"f46f4e": 9370, // Echowell
		"f46fed": 4921, // Fiberhome
		"f470ab": 6369, // Vivo
		"f47190": 152,  // Samsung
		"f47960": 3334, // Huawei
		"f47acc": 9371, // SolidFire
		"f47b09": 421,  // Intel
		"f47b5e": 152,  // Samsung
		"f47def": 152,  // Samsung
		"f47f35": 2,    // Cisco
		"f48139": 87,   // Canon
		"f483cd": 1595, // TP-Link
		"f4848d": 1595, // TP-Link
		"f485c6": 9372, // FDT
		"f48771": 9373, // Infoblox
		"f487c5": 3334, // Huawei
		"f48b32": 5697, // Xiaomi
		"f48c50": 421,  // Intel
		"f48ceb": 803,  // D-Link
		"f48e09": 457,  // Nokia
		"f48e38": 102,  // Dell
		"f48e92": 3334, // Huawei
		"f490ca": 9374, // Tensorcom
		"f492bf": 2978, // Ubiquiti
		"f49466": 9375, // CountMax
		"f49634": 421,  // Intel
		"f49651": 6082, // NAKAYO
		"f497c2": 9376, // Nebulon
		"f49c12": 9377, // Structab
		"f49f54": 152,  // Samsung
		"f49ff3": 3334, // Huawei
		"f4a475": 421,  // Intel
		"f4a4d6": 3334, // Huawei
		"f4a52a": 9378, // Hawa
		"f4a59d": 3334, // Huawei
		"f4a739": 826,  // Juniper
		"f4a80d": 1592, // Wistron
		"f4a997": 87,   // Canon
		"f4acc1": 2,    // Cisco
		"f4afe7": 545,  // Apple
		"f4b19c": 7031, // AltoBeam
		"f4b301": 421,  // Intel
		"f4b381": 9379, // WindowMaster
		"f4b52f": 826,  // Juniper
		"f4b5aa": 3031, // ZTE
		"f4b5bb": 1490, // Ceragon
		"f4b688": 542,  // Plantronics
		"f4b6e5": 9380, // TerraSem
		"f4b78d": 3334, // Huawei
		"f4b7b3": 6369, // Vivo
		"f4b8a7": 3031, // ZTE
		"f4bd9e": 2,    // Cisco
		"f4beec": 545,  // Apple
		"f4bf80": 3334, // Huawei
		"f4bfa8": 826,  // Juniper
		"f4c02f": 9381, // BlueBite
		"f4c114": 6392, // Technicolor
		"f4c248": 152,  // Samsung
		"f4c447": 9382, // Coagent
		"f4c6d7": 9383, // blackned
		"f4c714": 3334, // Huawei
		"f4c795": 9384, // WEY
		"f4c7c8": 9385, // Kelvin
		"f4ca24": 9386, // FreeBit
		"f4cb52": 3334, // Huawei
		"f4cc55": 826,  // Juniper
		"f4ce23": 421,  // Intel
		"f4ce46": 302,  // HP
		"f4ce48": 185,  // Extreme
		"f4cfa2": 6376, // Espressif
		"f4cfe2": 2,    // Cisco
		"f4d108": 421,  // Intel
		"f4d261": 9387, // SEMOCON
		"f4d488": 545,  // Apple
		"f4d9c6": 3505, // Unionman
		"f4d9fb": 152,  // Samsung
		"f4dbe3": 545,  // Apple
		"f4dbe6": 2,    // Cisco
		"f4dcf9": 3334, // Huawei
		"f4dd9e": 6241, // GoPro
		"f4de0c": 9388, // ESPOD
		"f4deaf": 3334, // Huawei
		"f4e204": 9389, // Traqueur
		"f4e2c6": 2978, // Ubiquiti
		"f4e3fb": 3334, // Huawei
		"f4e451": 3334, // Huawei
		"f4e4ad": 3031, // ZTE
		"f4e5f2": 3334, // Huawei
		"f4e9d4": 2025, // QLogic
		"f4ea67": 2,    // Cisco
		"f4eab5": 185,  // Extreme
		"f4eb38": 2048, // Sagemcom
		"f4ec38": 1595, // TP-Link
		"f4ee08": 102,  // Dell
		"f4ee14": 4253, // Mercury
		"f4f15a": 545,  // Apple
		"f4f197": 9390, // EMTAKE
		"f4f1e1": 687,  // Motorola
		"f4f26d": 1595, // TP-Link
		"f4f309": 152,  // Samsung
		"f4f3aa": 9391, // JBL
		"f4f524": 687,  // Motorola
		"f4f5a5": 457,  // Nokia
		"f4f5d8": 3521, // Google
		"f4f5db": 5697, // Xiaomi
		"f4f5e8": 3521, // Google
		"f4f646": 9392, // Dediprog
		"f4f647": 3031, // ZTE
		"f4f951": 545,  // Apple
		"f4fbb8": 3334, // Huawei
		"f4fcb1": 9393, // JJ
		"f4fd2b": 9394, // ZOYI
		"f4fefb": 152,  // Samsung
		"f800a1": 3334, // Huawei
		"f80113": 3334, // Huawei
		"f80332": 9395, // Khomp
		"f80377": 545,  // Apple
		"f8042e": 152,  // Samsung
		"f8084f": 2048, // Sagemcom
		"f80bbe": 125,  // ARRIS
		"f80bcb": 2,    // Cisco
		"f80cf3": 869,  // LG
		"f80d60": 87,   // Canon
		"f80dac": 302,  // HP
		"f80dea": 9396, // ZyCast
		"f80df0": 3031, // ZTE
		"f80df1": 9397, // Sontex
		"f80f41": 1592, // Wistron
		"f80f6f": 2,    // Cisco
		"f80ff9": 3521, // Google
		"f81037": 9398, // Atopia
		"f81093": 545,  // Apple
		"f81308": 457,  // Nokia
		"f814fe": 3505, // Unionman
		"f81547": 620,  // Avaya
		"f81654": 421,  // Intel
		"f81897": 1939, // 2Wire
		"f81a2b": 3521, // Google
		"f81a67": 1595, // TP-Link
		"f81d0f": 870,  // Hitron
		"f81d90": 9399, // Solidwintech
		"f81edf": 545,  // Apple
		"f81f32": 687,  // Motorola
		"f820a9": 3334, // Huawei
		"f82285": 4780, // Cypress
		"f823b2": 3334, // Huawei
		"f82441": 9400, // Yeelink
		"f82551": 46,   // Epson
		"f8272e": 9401, // Mercku
		"f82793": 545,  // Apple
		"f82819": 2873, // Liteon
		"f828c9": 3334, // Huawei
		"f829c0": 6904, // Availink
		"f82c18": 1939, // 2Wire
		"f82d7c": 545,  // Apple
		"f82dc0": 125,  // ARRIS
		"f82e3f": 3334, // Huawei
		"f82edb": 9402, // RTW
		"f82f5b": 9403, // eGauge
		"f82f65": 3334, // Huawei
		"f82f6a": 6281, // Itel
		"f8313e": 9404, // endeavour
		"f832e4": 1806, // ASUS
		"f83441": 421,  // Intel
		"f83553": 9405, // Magenta
		"f835dd": 1450, // Gemtek
		"f83869": 869,  // LG
		"f83880": 545,  // Apple
		"f83b1d": 6392, // Technicolor
		"f83b7e": 3334, // Huawei
		"f83cbf": 9406, // Botato
		"f83dff": 3334, // Huawei
		"f83e95": 3334, // Huawei
		"f83f51": 152,  // Samsung
		"f845ad": 3532, // Konka
		"f8461c": 101,  // Sony
		"f8462d": 9407, // SYNTEC
		"f84897": 89,   // Hitachi
		"f84a73": 9408, // Eumtech
		"f84a7f": 9409, // Innometriks
		"f84abf": 3334, // Huawei
		"f84cda": 3334, // Huawei
		"f84d33": 4921, // Fiberhome
		"f84d89": 545,  // Apple
		"f84e17": 101,  // Sony
		"f84e73": 545,  // Apple
		"f84f57": 2,    // Cisco
		"f85063": 9410, // Verathon
		"f85128": 9411, // SimpliSafe
		"f8516d": 9412, // Denwa
		"f852df": 9413, // VNL
		"f85329": 3334, // Huawei
		"f854b8": 5438, // Amazon
		"f855cd": 1399, // Visteon
		"f856c3": 3031, // ZTE
		"f85971": 421,  // Intel
		"f85a00": 9414, // Sanford
		"f85b3b": 2544, // Askey
		"f85b9c": 9415, // SB
		"f85bc9": 9416, // M-Cube
		"f85c4d": 457,  // Nokia
		"f85e42": 6392, // Technicolor
		"f85ea0": 421,  // Intel
		"f85f2a": 457,  // Nokia
		"f860f0": 1685, // Aruba
		"f86214": 545,  // Apple
		"f862aa": 9417, // xn
		"f8633f": 421,  // Intel
		"f863d9": 125,  // ARRIS
		"f864b8": 3031, // ZTE
		"f8665a": 545,  // Apple
		"f866f2": 2,    // Cisco
		"f86bd9": 2,    // Cisco
		"f86d73": 9418, // Zengge
		"f86ecf": 9419, // Arcx
		"f86eee": 3334, // Huawei
		"f86fc1": 545,  // Apple
		"f872ea": 2,    // Cisco
		"f87394": 1368, // Netgear
		"f873a2": 620,  // Avaya
		"f87588": 3334, // Huawei
		"f875a4": 4919, // LCFC
		"f8769b": 9420, // Neopis
		"f877b8": 152,  // Samsung
		"f8790a": 125,  // ARRIS
		"f87a41": 2,    // Cisco
		"f87aef": 9421, // Rosonix
		"f87b20": 2,    // Cisco
		"f87b7a": 125,  // ARRIS
		"f87fa5": 9422, // Greatek
		"f8811a": 9423, // Overkiz
		"f88200": 9424, // CaptionCall
		"f88479": 7722, // Yaojin
		"f884f2": 152,  // Samsung
		"f885f9": 374,  // Calix
		"f887f1": 545,  // Apple
		"f88b37": 125,  // ARRIS
		"f88c1c": 9425, // Kaishun
		"f88c21": 1595, // TP-Link
		"f88def": 9426, // Tenebraex
		"f88e85": 3869, // Comtrend
		"f88ea1": 6294, // Edgecore
		"f88f07": 152,  // Samsung
		"f88fca": 3521, // Google
		"f89066": 9427, // Nain
		"f893f3": 9428, // Volans
		"f894c2": 421,  // Intel
		"f89522": 3334, // Huawei
		"f895c7": 869,  // LG
		"f895ea": 545,  // Apple
		"f89753": 3334, // Huawei
		"f897a9": 306,  // Ericsson
		"f897cf": 9429, // Daeshin-Information
		"f8983a": 9430, // Leeman
		"f898b9": 3334, // Huawei
		"f898ef": 3334, // Huawei
		"f89955": 9431, // Fortress
		"f89a78": 3334, // Huawei
		"f89d0d": 1998, // Control
		"f89dbb": 9432, // Tintri
		"f89e94": 421,  // Intel
		"f8a03d": 9433, // Dinstar
		"f8a097": 125,  // ARRIS
		"f8a26d": 87,   // Canon
		"f8a2d6": 2873, // Liteon
		"f8a34f": 3031, // ZTE
		"f8a45f": 5697, // Xiaomi
		"f8a5c5": 2,    // Cisco
		"f8a73a": 2,    // Cisco
		"f8a91f": 9434, // ZVISION
		"f8a963": 358,  // Compal
		"f8a9d0": 869,  // LG
		"f8aa3f": 7446, // DWnet
		"f8aa8a": 9435, // Axview
		"f8ab05": 2048, // Sagemcom
		"f8ac65": 421,  // Intel
		"f8ac6d": 9436, // Deltenna
		"f8af05": 3334, // Huawei
		"f8afdb": 4921, // Fiberhome
		"f8b132": 3334, // Huawei
		"f8b156": 102,  // Dell
		"f8b1dd": 545,  // Apple
		"f8b46a": 302,  // HP
		"f8b54d": 421,  // Intel
		"f8b7e2": 2,    // Cisco
		"f8b95a": 869,  // LG
		"f8bae6": 457,  // Nokia
		"f8bbbf": 5820, // eero
		"f8bc0e": 5820, // eero
		"f8bc12": 102,  // Dell
		"f8bc41": 9437, // Rosslare
		"f8be0d": 9438, // A2UICT
		"f8bf09": 3334, // Huawei
		"f8c001": 826,  // Juniper
		"f8c091": 9439, // Highgates
		"f8c116": 826,  // Juniper
		"f8c249": 67,   // Private
		"f8c288": 2,    // Cisco
		"f8c397": 9440, // NZXT
		"f8c39e": 3334, // Huawei
		"f8c3cc": 545,  // Apple
		"f8c678": 9441, // Carefusion
		"f8c96c": 4921, // Fiberhome
		"f8ca85": 48,   // NEC
		"f8cab8": 102,  // Dell
		"f8cc6e": 9442, // DEPO
		"f8ce72": 1592, // Wistron
		"f8cfc5": 687,  // Motorola
		"f8d027": 46,   // Epson
		"f8d0ac": 101,  // Sony
		"f8d0bd": 152,  // Samsung
		"f8d111": 1595, // TP-Link
		"f8d3a9": 9443, // AXAN
		"f8d478": 833,  // Flextronics
		"f8dadf": 9444, // EcoTech
		"f8dae2": 9445, // NDC
		"f8db4c": 9446, // PNY
		"f8db7f": 1341, // HTC
		"f8db88": 102,  // Dell
		"f8dc7a": 9447, // Variscite
		"f8df15": 3939, // Sunitec
		"f8dfa8": 3031, // ZTE
		"f8dfe1": 9448, // MyLight
		"f8e079": 687,  // Motorola
		"f8e43b": 6722, // ASIX
		"f8e44e": 9449, // Mcot
		"f8e4e3": 421,  // Intel
		"f8e4fb": 2246, // Actiontec
		"f8e57e": 2,    // Cisco
		"f8e61a": 152,  // Samsung
		"f8e71e": 2737, // Ruckus
		"f8e7a0": 6369, // Vivo
		"f8e811": 3334, // Huawei
		"f8e903": 803,  // D-Link
		"f8e94e": 545,  // Apple
		"f8eda5": 125,  // ARRIS
		"f8f014": 9450, // RackWare
		"f8f082": 8147, // Nagtech
		"f8f1b6": 687,  // Motorola
		"f8f1e6": 152,  // Samsung
		"f8f21e": 421,  // Intel
		"f8f25a": 9451, // G-Lab
		"f8f532": 125,  // ARRIS
		"f8f7b9": 3334, // Huawei
		"f8f7d3": 639,  // International
		"f8f7ff": 9452, // SYN-Tech
		"f8fb2f": 9453, // Santur
		"f8fce1": 5438, // Amazon
		"f8fe5c": 9454, // Reciprocal
		"f8ff0b": 3690, // Electronic
		"f8ff5f": 8624, // Shenzhen
		"f8ffc2": 545,  // Apple
		"fc0012": 35,   // Toshiba
		"fc019e": 9455, // Vievu
		"fc0296": 5697, // Xiaomi
		"fc039f": 152,  // Samsung
		"fc0647": 9456, // Cortland
		"fc06ed": 5817, // M2Motive
		"fc084a": 4,    // Fujitsu
		"fc0a81": 185,  // Extreme
		"fc0fe6": 101,  // Sony
		"fc0fe7": 706,  // Microchip
		"fc1186": 9457, // Logic3
		"fc1193": 3334, // Huawei
		"fc122c": 3334, // Huawei
		"fc15b4": 302,  // HP
		"fc1607": 9458, // Taian
		"fc1794": 9459, // InterCreative
		"fc183c": 545,  // Apple
		"fc1910": 152,  // Samsung
		"fc1999": 5697, // Xiaomi
		"fc1a11": 6369, // Vivo
		"fc1bd1": 3334, // Huawei
		"fc1bff": 9460, // V-ZUG
		"fc1ca1": 457,  // Nokia
		"fc1d2a": 6369, // Vivo
		"fc1d43": 545,  // Apple
		"fc1d84": 9461, // Autobase
		"fc1e16": 9462, // IPEVO
		"fc1f19": 152,  // Samsung
		"fc1fc0": 9463, // Eurecam
		"fc2325": 9464, // EosTek
		"fc253f": 545,  // Apple
		"fc29e3": 6295, // Infinix
		"fc29f3": 9465, // McPay
		"fc2a9c": 545,  // Apple
		"fc2bb2": 2246, // Actiontec
		"fc2d5e": 3031, // ZTE
		"fc2e2d": 9466, // Lorom
		"fc2f40": 7844, // Calxeda
		"fc2f6b": 9467, // Everspin
		"fc2faa": 457,  // Nokia
		"fc2fef": 6515, // UTT
		"fc3342": 826,  // Juniper
		"fc335f": 9468, // Polyera
		"fc3497": 1806, // ASUS
		"fc3598": 9469, // Favite
		"fc35e6": 1399, // Visteon
		"fc3964": 6281, // Itel
		"fc3ce9": 9470, // Tsingtong
		"fc3da5": 2639, // Arcadyan
		"fc3f7c": 3334, // Huawei
		"fc3fdb": 302,  // HP
		"fc4009": 3031, // ZTE
		"fc4203": 152,  // Samsung
		"fc4482": 421,  // Intel
		"fc449f": 3031, // ZTE
		"fc4596": 358,  // Compal
		"fc48ef": 3334, // Huawei
		"fc492d": 5438, // Amazon
		"fc4ae9": 3797, // Castlenet
		"fc4bbc": 2426, // Sunplus
		"fc4da6": 3334, // Huawei
		"fc4ea4": 545,  // Apple
		"fc51a4": 125,  // ARRIS
		"fc528d": 6392, // Technicolor
		"fc589a": 2,    // Cisco
		"fc58df": 9471, // Interphone
		"fc5a1d": 870,  // Hitron
		"fc5b26": 9472, // MikroBits
		"fc5b39": 2,    // Cisco
		"fc5c45": 2737, // Ruckus
		"fc61e9": 4921, // Fiberhome
		"fc62b9": 430,  // Alpsalpine
		"fc643a": 152,  // Samsung
		"fc64ba": 5697, // Xiaomi
		"fc65b3": 3334, // Huawei
		"fc65de": 5438, // Amazon
		"fc66cf": 545,  // Apple
		"fc6c31": 9473, // LXinstruments
		"fc6dc0": 9474, // BME
		"fc6dd1": 5111, // APRESIA
		"fc6fb7": 125,  // ARRIS
		"fc71fa": 2656, // Trane
		"fc73fb": 3334, // Huawei
		"fc7516": 803,  // D-Link
		"fc75e6": 3621, // Handreamnet
		"fc7692": 4329, // Semptian
		"fc7774": 421,  // Intel
		"fc777b": 870,  // Hitron
		"fc7c02": 6848, // Phicomm
		"fc7ff1": 1685, // Aruba
		"fc8399": 620,  // Avaya
		"fc83c6": 9475, // N-Radio
		"fc8596": 9476, // Axonne
		"fc862a": 3334, // Huawei
		"fc8743": 3334, // Huawei
		"fc8a3d": 3031, // ZTE
		"fc8d3d": 9477, // Leapfive
		"fc8e6e": 9478, // StreamCCTV
		"fc8e7e": 125,  // ARRIS
		"fc8f90": 152,  // Samsung
		"fc8fc4": 1861, // Intelligent
		"fc90fa": 9479, // Independent
		"fc9114": 6392, // Technicolor
		"fc923b": 457,  // Nokia
		"fc9257": 5695, // Renesas
		"fc9435": 3334, // Huawei
		"fc946c": 9480, // Ubivelox
		"fc94ce": 3031, // ZTE
		"fc94e3": 6392, // Technicolor
		"fc956a": 4175, // Octagon
		"fc9643": 826,  // Juniper
		"fc9947": 2,    // Cisco
		"fc9bc6": 4654, // Sumavision
		"fc9bd4": 9481, // EdgeQ
		"fc9c98": 8405, // Arlo
		"fc9fae": 9482, // Fidus
		"fc9fe1": 9483, // CoNWIN.Tech
		"fca13e": 152,  // Samsung
		"fca183": 5438, // Amazon
		"fca621": 152,  // Samsung
		"fca667": 5438, // Amazon
		"fca6cd": 4921, // Fiberhome
		"fca841": 620,  // Avaya
		"fca84a": 9484, // Sentinum
		"fca89a": 3939, // Sunitec
		"fca9b0": 9485, // Miartech
		"fca9dc": 5695, // Renesas
		"fcaa14": 1929, // Giga-Byte
		"fcaa81": 545,  // Apple
		"fcaab6": 152,  // Samsung
		"fcab90": 3334, // Huawei
		"fcad0f": 9486, // QTS
		"fcae34": 125,  // ARRIS
		"fcaf6a": 9487, // Qulsar
		"fcafac": 9488, // Socionext
		"fcafbe": 9489, // TireCheck
		"fcb3bc": 421,  // Intel
		"fcb4e6": 2544, // Askey
		"fcb58a": 9490, // Wapice
		"fcb662": 9491, // IC
		"fcb698": 1640, // Cambridge
		"fcb6d8": 545,  // Apple
		"fcbc9c": 9492, // Vimar
		"fcbcd1": 3334, // Huawei
		"fcbd67": 3793, // Arista
		"fcbe7b": 6369, // Vivo
		"fcc233": 1806, // ASUS
		"fcc23d": 633,  // Atmel
		"fcc2de": 2056, // Murata
		"fcc734": 152,  // Samsung
		"fcc897": 3031, // ZTE
		"fccac4": 9493, // LifeHealth
		"fccce4": 9494, // Ascon
		"fccf62": 372,  // IBM
		"fcd436": 687,  // Motorola
		"fcd6bd": 1926, // Bosch
		"fcd733": 1595, // TP-Link
		"fcd848": 545,  // Apple
		"fcd908": 5697, // Xiaomi
		"fcdb21": 9495, // Samsara
		"fcdb96": 9496, // Enervalley
		"fcdbb3": 2056, // Murata
		"fcdc4a": 9497, // G-Wearables
		"fcde90": 152,  // Samsung
		"fce186": 9498, // A3M
		"fce1fb": 5955, // Array
		"fce26c": 545,  // Apple
		"fce33c": 3334, // Huawei
		"fce557": 457,  // Nokia
		"fce66a": 6570, // Industrial
		"fce806": 6445, // Edifier
		"fce998": 545,  // Apple
		"fcecda": 2978, // Ubiquiti
		"fcedb9": 9499, // Arrayent
		"fceee6": 9500, // Formike
		"fcf136": 152,  // Samsung
		"fcf152": 101,  // Sony
		"fcf1cd": 9501, // Optex-FA
		"fcf528": 2692, // ZyXEL
		"fcf5c4": 6376, // Espressif
		"fcf647": 4921, // Fiberhome
		"fcf77b": 3334, // Huawei
		"fcf8ae": 421,  // Intel
		"fcf8b7": 9502, // TRONTEQ
		"fcfbfb": 2,    // Cisco
		"fcfc48": 545,  // Apple
	},
	vendors: map[uint]string{
		7746: "01DB-Metravib",
		8820: "100fio",
		6309: "10NET/DCA",
		8465: "12Sided",
		3447: "16063",
		6627: "1MORE",
		4983: "1Net",
		2515: "2001",
		4785: "21168",
		2731: "27M",
		1094: "29530",
		7062: "2CRSI",
		3127: "2EI",
		8686: "2GIG",
		3688: "2M",
		1939: "2Wire",
		6564: "2bps",
		2499: "2wcom",
		2235: "2xWireless",
		3787: "30805",
		800:  "360",
		6844: "3ALogics",
		5335: "3CX",
		160:  "3Com",
		4374: "3DSP",
		5779: "3H",
		1421: "3J",
		3123: "3Leaf",
		6307: "3M",
		1294: "3UP",
		3579: "3Way",
		6681: "3bumen.com",
		8928: "3isysnetworks",
		4440: "3onedata",
		7482: "3pleplay",
		3243: "3soft",
		7797: "3view",
		766:  "3ware",
		2773: "4Access",
		6545: "4DReplay",
		4148: "4IPNET",
		3329: "4NSYS",
		4466: "4RF",
		9002: "5VT",
		1428: "6WIND",
		8550: "6harmonics",
		7166: "70mai",
		7034: "7HUGS",
		6260: "7INOVA",
		8677: "7signal",
		3729: "802automation",
		6597: "8D",
		8819: "8Devices",
		7916: "8Locations",
		6473: "8Mesh",
		5650: "8X8",
		8699: "9Solutions",
		899:  "@Track",
		1199: "@pos.com",
		1331: "A&D",
		8761: "A&R",
		7103: "A&T",
		2973: "A-First",
		3542: "A-Four",
		3584: "A-Link",
		537:  "A-One",
		4326: "A-Team",
		6223: "A-Trend",
		4150: "A-Trust",
		1127: "A-Z",
		6994: "A.D.C",
		8816: "A.N",
		6155: "A.T.N.R",
		4127: "A10",
		2319: "A2",
		4444: "A2B",
		9438: "A2UICT",
		9498: "A3M",
		4395: "A4SP",
		1009: "A5TEK",
		4674: "A7",
		2126: "AAC",
		5029: "AAE",
		1069: "AAEON",
		3269: "AAI",
		5619: "AANetcom",
		21:   "ABB",
		1961: "ABB./Tropos",
		2730: "ABB/Totalflow",
		7474: "ABC",
		8403: "ABIsystems",
		4441: "ABK",
		3552: "ACA-Digital",
		2476: "ACARD",
		1807: "ACC",
		7950: "ACES",
		8791: "ACK",
		1397: "ACKSYS",
		1799: "ACM",
		616:  "ACN",
		3431: "ACOGITO",
		443:  "ACR",
		8307: "ACT",
		491:  "ACT'L",
		8421: "ACTIVIO",
		8151: "ACTP",
		2213: "ACTi",
		7759: "ACX",
		5286: "AD",
		7777: "ADATA",
		6491: "ADB",
		9247: "ADC",
		4670: "ADC-Elektronik",
		6838: "ADD-Engineering",
		2203: "ADDI-DATA",
		2219: "ADDO-Japan",
		1254: "ADI",
		1653: "ADInstruments",
		5280: "ADS",
		5427: "ADVANCED",
		3355: "ADigit",
		2034: "AERAS",
		7087: "AES",
		8263: "AEV",
		1915: "AEWIN",
		1958: "AFAR",
		5107: "AFE",
		4437: "AFREEY",
		232:  "AG-E",
		8877: "AGAiT",
		1354: "AGFEO",
		9042: "AHN",
		1030: "AIDONIC",
		6388: "AIM",
		1905: "AINTech",
		3932: "AIOI",
		4087: "AIR802",
		2545: "AIRAYA",
		9203: "AIRSOUND",
		6446: "AITelecom",
		6224: "AK-Systems",
		7016: "AKELA",
		2653: "ALAXALA",
		3991: "ALCOMA",
		5526: "ALE",
		2746: "ALGOSYSTEM",
		8316: "ALT",
		9188: "ALTOGRAPHICS",
		5681: "ALi",
		469:  "AM",
		226:  "AMCC",
		13:   "AMD",
		8102: "AMERGINT",
		2222: "AMETEK",
		7394: "AMG",
		6622: "AMICCOM",
		5196: "AMIT",
		2790: "AMOD",
		4498: "AMPAK",
		5333: "AMTEC",
		5063: "ANI",
		3360: "ANSA",
		664:  "ANTARA.net",
		4653: "AOC",
		7534: "AOD",
		1872: "AOS",
		234:  "AOpen",
		5932: "APC",
		8945: "APCON",
		3219: "APD",
		4256: "API",
		9276: "API-K",
		1913: "APLUX",
		6524: "APR",
		5111: "APRESIA",
		8052: "APS",
		5473: "APT",
		8936: "AQ",
		4507: "ARBURG",
		6572: "ARCHEAN",
		8741: "AREC",
		1957: "ARIMA",
		2024: "ARION",
		276:  "ARK",
		2282: "ARKUS",
		462:  "ARM",
		2388: "ARMA",
		673:  "ARMITEL",
		6192: "ARN",
		125:  "ARRIS",
		2474: "ARTDIO",
		2561: "ARUZE",
		283:  "ARtem",
		1274: "ASA",
		854:  "ASB",
		9075: "ASD",
		726:  "ASE",
		2645: "ASEM",
		3531: "ASI",
		2011: "ASIP",
		6722: "ASIX",
		8379: "ASMedia",
		2871: "ASP",
		8745: "ASR",
		4719: "ASRock",
		7063: "ASSMANN",
		5072: "AST",
		4711: "ASTAK",
		3407: "ASTEL",
		1806: "ASUS",
		1057: "AT&T",
		8500: "ATAW",
		7993: "ATCOM",
		6991: "ATEK",
		1011: "ATI",
		5121: "ATM",
		3209: "ATMedia",
		6939: "ATN",
		1296: "ATO",
		6238: "ATOCS",
		7960: "ATOM",
		2830: "ATOMIC",
		6541: "ATP",
		3646: "ATRON",
		2356: "ATTO",
		8490: "ATW",
		2771: "ATX",
		4853: "AUTOVISION",
		2045: "AV",
		3118: "AVC",
		8358: "AVI",
		9252: "AVI-ON",
		6104: "AVIDIA",
		163:  "AVLAB",
		621:  "AVM",
		4588: "AVSystem",
		8398: "AVTrace",
		9249: "AWCER",
		2594: "AWIND",
		6970: "AWare",
		9443: "AXAN",
		1894: "AXELL",
		7617: "AXERRA",
		7409: "AXIM",
		8394: "AXPRO",
		5051: "AZS",
		4066: "AZTEQ",
		9180: "Aartesys",
		327:  "Aatr",
		7758: "Aava",
		1749: "Abatron",
		2655: "Abbey",
		6719: "Abeeway",
		907:  "Abeona",
		6401: "Abicom",
		6744: "Ability",
		5255: "Abit",
		4815: "Ablaze",
		6306: "Able",
		8032: "Ablelink",
		6215: "Abler",
		2215: "Ablerex",
		7555: "Abloomy",
		2556: "AboCom",
		7351: "Aboundi",
		1888: "AboveCable",
		6538: "Abrantix",
		825:  "AbsoluteValue",
		4903: "AcSiP",
		5767: "Acacia",
		1039: "AcceLight",
		3003: "Accedian",
		5209: "Accel",
		6932: "AccelStor",
		643:  "Accelent",
		2305: "Accelerated",
		7697: "Accelink",
		6347: "Accell",
		1014: "Accella",
		3210: "Accense",
		9074: "Accensus",
		5471: "Access",
		2310: "Accesslan",
		5407: "Acclaim",
		4082: "Accm",
		3509: "Acconet",
		5174: "Accord",
		8715: "Accordance",
		376:  "Accordion",
		2882: "Accsense",
		145:  "Accton",
		2016: "Accu-Sort",
		5375: "Accu-Time",
		6665: "AccuSpec",
		9273: "Accuenergy",
		8310: "Accupix",
		908:  "Accusys",
		8302: "Accutome",
		6517: "Ace",
		3249: "AceNet",
		857:  "Aceex",
		9005: "Acentic",
		143:  "Acer",
		6090: "Acetel",
		6166: "Ackfin",
		7710: "Aclima",
		7132: "Acoinfo",
		527:  "Acomz",
		107:  "Acorn",
		1231: "Acorp",
		820:  "Acqis",
		6025: "Acrison",
		277:  "Acromag",
		6289: "Acroname",
		759:  "Acronet",
		424:  "Acrosser",
		655:  "Acrowave",
		1735: "Actel",
		540:  "Actelis",
		1052: "Acterna",
		2293: "Actia",
		7025: "Actifio",
		9300: "Actility",
		7350: "Actineon",
		3632: "Action",
		8366: "Actioncable",
		2246: "Actiontec",
		4181: "Actis",
		2549: "ActivNetworks",
		294:  "Activetelco",
		7219: "Actlas",
		2497: "Actuality",
		7513: "Acubit",
		5604: "Acucomm",
		335:  "Aculab",
		7852: "Acumen",
		9343: "Acurix",
		9269: "Acuro",
		6089: "Acute",
		4914: "Adachi-Syokai",
		8287: "Adafruit",
		8776: "Adamson",
		6619: "Adanis",
		396:  "Adapcom",
		6516: "Adapt-IP",
		2777: "Adapt4",
		134:  "Adaptec",
		6244: "Adapteva",
		5467: "Adaptive",
		2595: "Adaptix",
		9339: "Adaptrum",
		4176: "Adastra",
		6495: "Adatis",
		1051: "Adax",
		8440: "Adaxys",
		2518: "Aday",
		2729: "Add-On",
		7215: "AddOn",
		411:  "AddPac",
		8072: "Addenergie",
		2192: "Adder",
		8751: "Additech",
		2616: "Addlogix",
		5281: "Addonics",
		5084: "Addtron",
		1665: "Addvalue",
		1179: "Adept",
		8363: "Adero",
		4360: "Adesys",
		2376: "Adhoc",
		3382: "Adhoco",
		4990: "Adid",
		2125: "Adimos",
		2847: "Aditec",
		4977: "Adlink",
		6165: "Admtek",
		7265: "Adnacom",
		5069: "Adobe",
		9369: "Adonit",
		72:   "Adra",
		5490: "Adsoft",
		342:  "Adtec",
		4986: "Adtech",
		202:  "Adtran",
		4909: "Adura",
		7466: "Advan",
		6121: "AdvanSys",
		587:  "Advanced",
		9214: "Advanced-Connectek",
		2342: "Advanet",
		6370: "Advansee",
		3402: "Advansus",
		7383: "Advantage",
		1703: "Advantech",
		7752: "Advas",
		987:  "Advent",
		4053: "Adventiq",
		7889: "Advidia",
		3086: "Aegate",
		4263: "Aehr",
		1622: "Aeluros",
		5070: "Aeon",
		8651: "Aeonsemi",
		1343: "AeroConcierge",
		3624: "AeroVIronment",
		5233: "Aerocomm",
		6850: "Aerodev",
		7707: "Aerodisk",
		5212: "Aeroflex",
		1721: "Aeronix",
		1863: "Aeroscout",
		2495: "Aerotech",
		2225: "Aerotelecom",
		539:  "Aeta",
		4728: "Aetas",
		7669: "Aetek",
		8356: "Aetheris",
		6645: "Aetheros",
		2827: "Aevoe",
		1944: "Afco",
		7423: "Affinegy",
		7863: "Affirmed",
		7837: "AgJunction",
		6435: "AgLogica",
		3546: "Agami",
		3864: "Agapha",
		8631: "Agatel",
		345:  "Agere",
		6339: "Agfa",
		1610: "Agile",
		7997: "AgileMesh",
		653:  "Agilent",
		5492: "Agilis",
		5659: "Agranat",
		3728: "AhnLab",
		7316: "AiNET",
		2958: "AiZen",
		8919: "Aiconn",
		7467: "Aidc",
		7201: "Aidon",
		8912: "Aifa",
		1844: "Ailocom",
		3534: "Aipermon",
		1702: "Aiphone",
		6243: "Aipon",
		1551: "Aiptek",
		4001: "Air2App",
		1611: "Air2U",
		8401: "AirCUVE",
		3089: "AirDefense",
		5290: "AirFiber",
		1718: "AirFlow",
		2828: "AirLink",
		1621: "AirLogic",
		1982: "AirMagnet",
		6583: "AirNav",
		978:  "AirRunner",
		7554: "AirScape",
		216:  "AirSwitch",
		1594: "AirVast",
		2920: "Airak",
		3427: "AireSpider",
		6454: "Airenetworks",
		3741: "Airgain",
		1599: "Airgo",
		774:  "Airocon",
		6465: "Airoha",
		8220: "Airpo",
		7073: "Airsonics",
		259:  "Airspan",
		4429: "Airtech",
		860:  "Airvana",
		3708: "Airvod",
		9190: "Airware",
		1994: "Airwave",
		6658: "Airwire",
		9150: "Aisai",
		7133: "Aisino",
		5505: "Aitech",
		6922: "Aitexin",
		5242: "Aiwa",
		2615: "Aiware",
		5836: "Ajile",
		4568: "Ajinextek",
		4862: "Ajoho",
		3819: "Akamai",
		4987: "Akamba",
		1726: "Akcp",
		8977: "Akenori",
		2794: "Akimbi",
		1691: "Akita",
		1342: "Akom",
		2868: "Akorri",
		2114: "Aksys",
		6413: "Akuvox",
		9109: "Akyllor",
		2318: "Alacritech",
		4426: "Alacron",
		4276: "Alamar",
		6050: "Alantro",
		5831: "Alaris",
		3856: "Alarm.com",
		7049: "Albahith",
		1504: "Albatron",
		3607: "Albis",
		2068: "Alcatel",
		6992: "Alcea",
		7202: "Alcomp",
		1460: "Alcon",
		3834: "Alektrona",
		3910: "Alereon",
		3756: "Alertme.com",
		5328: "Alerton",
		3915: "Alertus",
		838:  "Alesis",
		935:  "Alexon",
		5937: "Alfa",
		6123: "Alfatech",
		4535: "Alflex",
		6857: "Alge-Timing",
		4495: "Algo",
		3937: "Algolith",
		4565: "Algolware",
		5154: "Algorithmics",
		4604: "Algorithmix",
		4188: "Algorithms",
		5999: "Alidian",
		3670: "Alien",
		1680: "Align",
		6755: "Alinco",
		8166: "Alinket",
		4325: "AliphCom",
		1295: "Alistel",
		219:  "Alitec",
		2895: "Alive",
		7831: "Aliwei",
		365:  "All-Win",
		8675: "AllDSP",
		606:  "Allegro",
		5576: "Allgon",
		2652: "Alliant",
		8093: "Allied-telesisK.K",
		7498: "Allis",
		2263: "Allnet",
		586:  "Alloptic",
		1349: "Allot",
		93:   "Alloy",
		1111: "Allradio",
		3386: "Alltec",
		5537: "Allumer",
		4922: "Allwell",
		9083: "Allwinner",
		1587: "Allworx",
		5639: "Aloha",
		7080: "Along",
		4035: "Aloys",
		2236: "Alpha",
		5295: "Alpha-TOP",
		8852: "AlphaTheta",
		8223: "Alphatronics",
		4155: "Alphion",
		5751: "Alps",
		430:  "Alpsalpine",
		3610: "AlsterAero",
		2309: "Alta",
		3484: "Altai",
		6391: "Altasec",
		3393: "Altec",
		4215: "Altech",
		5408: "Alteon",
		1167: "Altera",
		1559: "AltiGen",
		2916: "Alticast",
		6824: "Altierre",
		5652: "Altiga",
		7031: "AltoBeam",
		128:  "Altos",
		8749: "Altronic",
		7385: "Alula",
		3765: "AmRoad",
		690:  "Amann",
		5253: "Amano",
		6029: "Amaquest",
		5404: "Amati",
		5438: "Amazon",
		7477: "Amazon.com",
		5001: "Amber",
		2349: "AmbiCom",
		3006: "Ambient",
		8218: "Ambiq",
		1043: "Ambrado",
		1601: "Ambri",
		8326: "Amcrest",
		5511: "Amdahl",
		3896: "Amec",
		2671: "Amedia",
		360:  "Amer.com",
		2358: "American",
		4488: "Amerigon",
		106:  "Ameristar",
		5733: "Ameritec",
		2418: "Amherst",
		3697: "Amic",
		5998: "Amigo",
		6458: "Amimon",
		320:  "Amino",
		1924: "Amity",
		4191: "Amkly",
		1972: "Ammasso",
		2579: "Amoi",
		6667: "Amon",
		8939: "Amosense",
		8114: "Amper",
		81:   "Ampere",
		1487: "Amperion",
		50:   "Ampex",
		1475: "Amphenol",
		6550: "Amphitech",
		754:  "Amphus",
		2062: "Ample",
		3325: "Amplex",
		6693: "Amplified",
		5101: "Ampro",
		5442: "Ampt",
		6361: "Amsc",
		2061: "Amtelco",
		3357: "AnNeal",
		6463: "AnaCom",
		2854: "Anagran",
		3586: "Analogic",
		8674: "Analytica",
		2283: "Anam",
		4484: "Anatek",
		1766: "Anator",
		3378: "Anaveo",
		5402: "Ancot",
		5297: "Anda",
		4570: "Andes",
		782:  "Andiamo",
		2573: "Andrew",
		7003: "Andtek",
		7036: "Anedo",
		6159: "Anerma",
		3514: "Angel",
		5734: "Angia",
		7655: "Angler",
		1786: "Animation",
		1468: "Annso",
		8972: "Anobit",
		1149: "Anoto",
		7122: "Anovo",
		98:   "Anritsu",
		5132: "Ansel",
		3593: "Anseri",
		6443: "Ansjer",
		4598: "Ansync",
		6503: "Antailiye",
		7955: "Antaira",
		2415: "Antec",
		7704: "Antex",
		1639: "Anthology",
		4562: "Antipode",
		5448: "Antlow",
		3489: "AnyDATA",
		8225: "Anywave",
		6942: "Anywire",
		1734: "Aoip",
		9088: "Apacer",
		4817: "Apacewave",
		1500: "Apani",
		4489: "Apass",
		7821: "Aperi",
		406:  "Apex",
		5699: "Apexx",
		7850: "Apiste",
		4295: "Aplicaciones",
		8860: "Aplicom",
		5577: "Aplio",
		591:  "Apogee",
		5423: "Apollo",
		4352: "AppTech",
		4962: "Appian",
		545:  "Apple",
		2728: "Application",
		5765: "Applicom",
		41:   "Applicon",
		4491: "Applition",
		2440: "Appointech",
		3637: "Apprion",
		1755: "Appro",
		8654: "Apption",
		3242: "Apricorn",
		47:   "Apricot",
		150:  "April",
		4546: "Aprius",
		4122: "Aprotech",
		5963: "Aptec",
		4459: "Aptiv",
		6111: "Aptix",
		8250: "Aptos",
		6888: "Aqavi.com",
		3253: "Aquantia",
		4668: "Aquila",
		4876: "Arada",
		3193: "Araneo",
		182:  "Aravox",
		7603: "Arbiter",
		3097: "Arbitron",
		859:  "Arbor",
		9018: "Arcadia",
		2639: "Arcadyan",
		5915: "Archipel",
		3150: "Archos",
		5862: "Arco",
		8136: "Arcom",
		1568: "Arcon",
		9221: "Arcontia",
		1788: "Arcor&Co",
		926:  "Arcturus",
		9419: "Arcx",
		6043: "Ardent",
		6920: "Ardomus",
		8466: "Arduino",
		5903: "Areanex",
		3663: "Areca",
		4477: "Ared",
		429:  "Arelnet",
		879:  "Arescom",
		8856: "Areson",
		3312: "Argard",
		5631: "Argon",
		2400: "Argosy",
		471:  "Argus",
		5542: "Aries",
		490:  "Arima",
		3793: "Arista",
		127:  "Arix",
		2684: "Arkados",
		8405: "Arlo",
		4029: "Armadeus",
		9095: "Armatura",
		1140: "Armillaire",
		3331: "Armorlink",
		6548: "Armtel",
		5955: "Array",
		2405: "ArrayComm",
		9499: "Arrayent",
		9148: "Arrcus",
		9359: "Arrikto",
		6941: "Arris",
		7815: "Arrive",
		3480: "Arrow7",
		2331: "ArrowPoint",
		2852: "ArrowSpan",
		2927: "Artech",
		5506: "Artel",
		3218: "Arti",
		2691: "Artila",
		2766: "Artimi",
		69:   "Artisoft",
		5372: "Artiza",
		3917: "Artjoy",
		1494: "Artnix",
		4883: "Artray",
		1685: "Aruba",
		8703: "Aryaka",
		5221: "Asaca",
		7324: "Asahi",
		856:  "Asahi-Engineering",
		100:  "Asante",
		4644: "Asantron",
		2600: "Ascalade",
		3297: "Ascend",
		1690: "Ascent",
		1692: "Ascom",
		9494: "Ascon",
		7964: "Asct",
		3289: "Asia",
		1511: "AsiaRF",
		2672: "Asiamajor",
		1666: "Asiarock",
		6976: "Asiatelco",
		4892: "Asiteq",
		2103: "Asix",
		1545: "Aska",
		2544: "Askey",
		2005: "Asmax",
		3077: "Asmobile",
		7470: "Asoni",
		343:  "Asound",
		5656: "Aspect",
		940:  "Aspen",
		7041: "Asperiq",
		6724: "Assa",
		8277: "Assembled",
		2461: "Assurance",
		2072: "Astarte",
		4078: "Astec",
		6350: "Astech",
		2638: "Astek",
		1567: "Astera",
		5761: "Aston",
		2266: "Astri",
		5377: "AstroNova",
		450:  "Astrodesign",
		6404: "Astrol",
		4447: "Astron",
		2218: "Astute",
		4685: "Asumo",
		3370: "AtRoad",
		7381: "Atamo",
		6068: "Atan",
		33:   "Atari",
		3401: "Atech",
		1154: "Atek",
		7624: "Ateme",
		2345: "Aten",
		3311: "Atera",
		138:  "Atex",
		8532: "AthenTek",
		4308: "Athena",
		5126: "Athenix",
		535:  "Atheros",
		9046: "Atil",
		5456: "Atlantix",
		7820: "Atlinks",
		8676: "Atlona",
		633:  "Atmel",
		7931: "Atmosic",
		2382: "Atmosphere",
		543:  "Atoga",
		9197: "Atol",
		9141: "Atomax",
		7512: "Atomos",
		5168: "Atomwide",
		5420: "Atop",
		9398: "Atopia",
		599:  "Atrica",
		2010: "Atrie",
		3842: "Attero",
		2831: "Atto",
		6594: "AuVerte",
		241:  "AudeSi",
		9126: "Audeze",
		1660: "Audio",
		1503: "Audio-Technica",
		7161: "AudioControl",
		2586: "AudioDev",
		378:  "AudioRamp.com",
		3854: "AudioScience",
		8531: "Audioengine",
		4315: "Audiovox",
		8461: "Audiowise",
		8464: "Audivo",
		6620: "Audoo",
		8756: "Audyssey",
		1364: "Auerswald",
		2162: "Augmentix",
		8138: "Augtek",
		9016: "Aurender",
		1119: "Aurora",
		8140: "Aus.Linx",
		37:   "Auspex",
		3472: "Austar",
		5746: "Austron",
		4236: "Autec",
		7601: "Auth-Servers",
		6387: "Auto",
		3122: "Auto-Maskin",
		1848: "AutoCell",
		8436: "AutoCrib",
		5812: "AutoGas",
		8213: "AutoHotBox",
		9461: "Autobase",
		3604: "Autocom",
		4210: "Autocomputer",
		7019: "Autoit",
		6189: "AutomatedLogic",
		8779: "Automatic",
		3186: "Automation",
		4689: "Autonet",
		7458: "Autonics",
		6925: "Autosales",
		1517: "Autostar",
		634:  "Autosys",
		8182: "Autotalks",
		3005: "Autotelenet",
		27:   "Autotote",
		2046: "Auvitran",
		4411: "Avaak",
		3214: "Avago",
		295:  "Avail",
		6904: "Availink",
		5206: "Aval",
		669:  "Avalue",
		2229: "Avamax",
		2228: "Avanex",
		2637: "Avantec",
		3645: "Avantis",
		8206: "Avanu",
		2071: "Avara",
		5570: "Avatar",
		620:  "Avaya",
		3283: "Avega",
		7930: "Avelon",
		8474: "Aventura",
		4456: "AverLogic",
		6551: "Avermetrics",
		5679: "Avex",
		5658: "Avici",
		5754: "Avid",
		3883: "Avidyne",
		1243: "Avilinks",
		3158: "Aviqtech",
		349:  "Avision",
		1275: "Avistar",
		6558: "Avistel",
		4522: "Avitech",
		8607: "Aviwest",
		2538: "Avix",
		8555: "Avizia",
		423:  "Avnet",
		2633: "Avolites",
		8272: "Avonic",
		8104: "Avotek",
		1528: "Avrio",
		1798: "Avtec",
		4803: "Avtex",
		7630: "Avvasi",
		1689: "Avvio",
		2370: "Aware",
		2919: "AwarePoint",
		3696: "Awox",
		9069: "Axacore",
		5722: "Axel",
		9128: "AxesNetwork",
		3504: "Axesstel",
		8304: "Axiim",
		8632: "Axilspot",
		5415: "Axiom",
		3918: "Axion",
		1945: "Axiowave",
		6429: "Axiros",
		5131: "Axis",
		6459: "Axium",
		5942: "Axon",
		4637: "Axona",
		9476: "Axonne",
		2523: "Axsun",
		9435: "Axview",
		6715: "Axxana",
		8929: "Ayecka",
		3529: "Ayecom",
		7564: "Ayla",
		3224: "Azalea",
		3238: "Azonic",
		5796: "Azonix",
		7500: "Azroad",
		4854: "Aztech",
		1357: "Aztek",
		2188: "Azul",
		8239: "Azuray",
		5571: "Azure",
		3667: "Azuretec",
		3004: "Azurewave",
		1478: "Azylex",
		2099: "B&B",
		7223: "B-Link",
		6073: "B2C2",
		3141: "BA",
		2351: "BACHMANN",
		1743: "BAE",
		9059: "BALMUDA",
		3533: "BARTEC",
		8730: "BASIC",
		6157: "BAY",
		3060: "BBH",
		8060: "BBMC",
		1993: "BBN",
		2581: "BBWM",
		2205: "BBox",
		382:  "BCM",
		8846: "BCTech",
		1871: "BECS",
		8152: "BEFEGA",
		7723: "BEKO",
		4075: "BEN-RI",
		2929: "BETA",
		9118: "BH",
		5744: "BHP",
		7416: "BHR",
		7212: "BHTC",
		4693: "BICOM",
		3894: "BIJ",
		8931: "BIOS",
		1716: "BIOTRONIK",
		4042: "BISA",
		7584: "BITwave",
		7135: "BK",
		9008: "BKAV",
		2084: "BLIP",
		8065: "BLT",
		7460: "BLU",
		2718: "BLwave",
		2763: "BM",
		4196: "BMC",
		9474: "BME",
		258:  "BMW",
		192:  "BNA",
		9360: "BND",
		6438: "BNS",
		1013: "BNTECHNOLOGY",
		2835: "BOAZ",
		2078: "BOE",
		3669: "BPL",
		3827: "BPT",
		7283: "BQ",
		7406: "BQT",
		446:  "BRECIS",
		5273: "BRODEL",
		8942: "BSE",
		6581: "BSMediasoft",
		6272: "BSP",
		1261: "BST",
		3512: "BSkyB",
		8683: "BT",
		3131: "BT-Links",
		8814: "BT9",
		5740: "BTG",
		2903: "BTI",
		1843: "BTU",
		7690: "BUJEON",
		5569: "BVM",
		2237: "BWA",
		8187: "BXB",
		4046: "BYD",
		6907: "BYK-Gardner",
		8807: "Bach-Simpson",
		4837: "Bachmann",
		7252: "Baicells",
		7288: "Baikal",
		4594: "Ball-It",
		3416: "Balluff",
		7472: "Baltech",
		6061: "Balthazar",
		7165: "Bamboo",
		4902: "BandRich",
		961:  "Banderacom",
		1008: "Bandspeed",
		1197: "Banksys",
		4595: "Banner",
		6326: "Banyan",
		9084: "Baraja",
		6623: "Barberry",
		721:  "Bardac",
		1297: "Barix",
		465:  "Barracuda",
		6258: "Barrot",
		1362: "Bartech",
		8848: "Barun",
		4967: "Basler",
		4347: "Bastec",
		5470: "Basys",
		3988: "Battistoni",
		2683: "BaudTec",
		8346: "Baumer",
		7434: "Baxter",
		84:   "Bay",
		8529: "Baycity",
		813:  "Baydel",
		5641: "Bayly",
		6187: "Bcom",
		1857: "BeWAN",
		7671: "Beacon",
		1582: "BeamReach",
		8242: "Beamex",
		6261: "Beats",
		8865: "Beautiful",
		8551: "Becker",
		8396: "Becker-Antriebe",
		3683: "Beckmann",
		6692: "Becon",
		8345: "Becton",
		9200: "Befs",
		5990: "Behavior",
		1086: "Beicom",
		5237: "Beijer",
		9182: "Beijing-Cloud",
		6910: "Beissbarth",
		7781: "Beken",
		7912: "Bekey",
		2661: "Belco",
		3426: "Belden",
		2468: "Belkin",
		15:   "Bell",
		3823: "Bellon",
		1062: "Bematech",
		7987: "BenQ",
		8681: "BenRui",
		2079: "Benchmark",
		4597: "Benein",
		4619: "Benign",
		9105: "Bentek",
		9358: "Benu",
		6154: "Berkeley",
		4084: "Berker",
		4324: "Berkshire",
		2040: "Bermai",
		213:  "Best",
		3926: "BestComm",
		7834: "Bestek",
		5264: "Beta",
		8637: "BetterBots",
		8327: "Beyless",
		3411: "Beyondwiz",
		3705: "Bharat",
		7647: "Bhuu",
		9297: "Bi2-Vision",
		1652: "BiTMICRO",
		972:  "Biacore",
		7872: "Biamp",
		505:  "BigBand",
		3397: "Bigfoot",
		2334: "Billionton",
		1951: "Bils",
		6368: "BinCube",
		1591: "Binatone",
		5814: "Bintec",
		8767: "Binxin",
		3763: "Bio-Rad",
		9238: "BioControl",
		7151: "Biodata",
		4610: "Biodevices",
		7344: "Biodit",
		425:  "Bionet",
		8444: "Bionics",
		8584: "Bionime",
		5607: "Biopac",
		4072: "Bioscrypt",
		8342: "Biosoundlab",
		2706: "Biospace",
		3710: "Bird",
		8762: "Biscotti",
		7299: "Bison",
		8778: "BitBox",
		7706: "Bita",
		2644: "Bitatek",
		6365: "Bitel",
		9143: "Bitmain",
		1048: "Bitrage",
		995:  "Bitran",
		4724: "Bitrode",
		6007: "Bitronics",
		3414: "BitsGen",
		6056: "Bitswitch",
		1638: "BittWare",
		1160: "Bitworks",
		920:  "Bivio",
		2997: "Bixolon",
		3611: "Biz-2-Me",
		8354: "BizLink",
		6426: "Bizlink",
		2220: "BlackBerry",
		8657: "Blaster",
		4100: "Blatand",
		9299: "Bloombase",
		4129: "Blue-White",
		903:  "Blue2space",
		5058: "BlueBit",
		9381: "BlueBite",
		2514: "BlueExpert",
		1311: "BlueKorea",
		8805: "BlueN",
		2623: "BluePacket",
		9285: "BlueRadios",
		1436: "BlueWINC",
		8438: "Bluebank",
		3981: "Bluecard",
		8682: "Bluecom",
		1108: "Bluegiga",
		4036: "Bluelight",
		3658: "Blueone",
		1307: "Bluetake",
		4833: "Bluetechnix",
		3711: "Blueway",
		1532: "Bluewire",
		3589: "Blusens",
		8058: "Bluwan",
		4318: "Blynke",
		7853: "BnCOM",
		6616: "BobjGear",
		3478: "Bobst",
		5900: "Boca",
		1684: "Bodet",
		321:  "Bodmann",
		8247: "BodyMedia",
		1041: "Boeing",
		3308: "Bogen",
		8667: "Bolymin",
		1236: "Bon",
		3519: "Bona",
		8204: "Bookeen",
		2858: "Bookham",
		6903: "Boosty",
		8614: "Borea",
		5041: "Borgardt",
		1092: "Boris",
		1926: "Bosch",
		1819: "Bose",
		5275: "Boser",
		8245: "Bostex",
		5816: "Boston",
		6690: "Bosung",
		9406: "Botato",
		4106: "Botik",
		6510: "Bowei",
		8317: "BoxCast",
		8378: "BoxLock",
		9036: "Boxin",
		4138: "Boyoung",
		8992: "Bragi",
		7773: "Brain",
		1146: "Brain21",
		7444: "BrainCo",
		1581: "Brainchild",
		2217: "Brainium",
		271:  "Brains",
		1361: "Braintree",
		5719: "Brand",
		4320: "Brandywine",
		795:  "Brans",
		971:  "Bravara",
		6763: "Braveridge",
		9119: "Bravo",
		8451: "Breathometer",
		700:  "Breezecom",
		165:  "Bri-Link",
		6245: "BriView",
		4651: "Bridge",
		517:  "BridgeWave",
		1173: "Bridgeco",
		629:  "Bridgeworks",
		1992: "Bright",
		8200: "BrightSign",
		1053: "Brightcom",
		7489: "Brightsource",
		7412: "Brilliantts",
		3961: "Briot",
		8750: "Brita",
		1130: "Britestream",
		3860: "Brivo",
		4994: "Brix",
		2569: "BroadEasy",
		3490: "Broadata",
		2422: "Broadband",
		1754: "Broadbus",
		858:  "Broadcom",
		405:  "Broadframe",
		8611: "Broadlink",
		6027: "Broadlogic",
		709:  "Broadmax",
		4585: "Broadway",
		90:   "Brocade",
		449:  "Bromax",
		5760: "Brooktrout",
		3230: "Brother",
		3053: "Browan",
		4024: "Brunata",
		4960: "Bryant",
		6252: "Bryston",
		7786: "Bryton",
		510:  "Bticino",
		1077: "Buffalo",
		9063: "Bulat",
		2098: "Burdick",
		6905: "Burg-Wachter",
		6662: "Burlywood",
		5885: "Burr-Brown",
		4552: "Burster",
		3100: "Bury",
		5085: "Bustek",
		8415: "Busware.DE",
		7523: "Buwon",
		3333: "Buyang",
		6183: "Byas",
		9163: "Bytedance",
		8547: "Bytemark",
		5146: "Bytex",
		4431: "Byzoro",
		8127: "C",
		1353: "C&I",
		1305: "C&S",
		3938: "C-BEL",
		305:  "C-CoM",
		338:  "C-CoR",
		404:  "C-CoR.net",
		4534: "C-Matic",
		2015: "C-guys",
		5173: "C.A.E.N",
		7571: "C.E",
		8209: "C.O.B.O",
		810:  "C.P",
		2809: "C4Line",
		454:  "CAB",
		2191: "CABLELOGIC",
		5232: "CAE",
		4404: "CAI",
		6603: "CAP-Tech",
		1326: "CAS",
		6372: "CASwell",
		1194: "CATS",
		3828: "CBC",
		5195: "CBL",
		4810: "CC",
		385:  "CC&C",
		4119: "CCS",
		816:  "CDS-Electronics",
		7503: "CDYNE",
		2555: "CE-Infosys",
		7465: "CELIZION",
		844:  "CELOX",
		9101: "CELOXICA",
		1234: "CEM",
		3301: "CENITS",
		4681: "CENTRAL",
		1084: "CENiX",
		4771: "CEStronics",
		2424: "CET",
		7658: "CETORY.TV",
		4733: "CEVA",
		8153: "CG",
		7347: "CHAHOO",
		2726: "CHIPS",
		2439: "CIMSYS",
		164:  "CIS",
		3072: "CITEL",
		7076: "CKD",
		8011: "CKS",
		834:  "CLCsoft",
		2012: "CMA/Microdialysis",
		5285: "CMC",
		4245: "CMD",
		3526: "CMOTECH",
		2363: "CMS",
		4004: "CNB",
		7496: "CNEX",
		1877: "CNMP",
		8955: "CNSLink",
		1066: "CNSystems",
		1253: "CNet",
		9259: "CPI",
		3892: "CQ",
		2923: "CREVIS",
		4009: "CRFS",
		7567: "CRU-Dataport",
		3076: "CReTE",
		8616: "CRemote",
		3362: "CS",
		6628: "CSE-Servelec",
		2390: "CSI-Control",
		5454: "CSK",
		8991: "CSM",
		34:   "CSS",
		4634: "CSSI",
		4722: "CTERA",
		2412: "CTI",
		2058: "CTS",
		8486: "CVC",
		9010: "CVT",
		3126: "CWT",
		2312: "CYQ've",
		654:  "CYZENTECH",
		9034: "CZ.NIC",
		6609: "Caavo",
		4173: "Cable",
		6511: "CableFree",
		8594: "CableWorld",
		1148: "Cabletime",
		16:   "Cabletron",
		420:  "Cablevision",
		2163: "Cableware",
		1492: "Cabot",
		5919: "Cache",
		5015: "CacheFlow",
		1328: "CacheVision",
		7839: "Cachengo",
		5449: "Cactus",
		8806: "Cadac",
		212:  "Cadant",
		2564: "Cadco",
		5550: "Cadre",
		2593: "Caen",
		1630: "Cal-Comp",
		7587: "CalDigit",
		7769: "Calantec",
		5949: "Calcomp",
		2839: "Calculex",
		8636: "Caldero",
		2401: "Calista",
		374:  "Calix",
		4760: "CallTechSolution",
		2250: "CallURL",
		3762: "Callpod",
		6752: "Calsys",
		7844: "Calxeda",
		5837: "Caly",
		3695: "Calyptech",
		5210: "Cambex",
		665:  "Cambium",
		1640: "Cambridge",
		8409: "Cambrionix",
		7196: "Camco",
		3383: "Cameo",
		8993: "Cameronet",
		8:    "Camex",
		6020: "Campio",
		3244: "Camrivox",
		158:  "Camtec",
		309:  "Camtel",
		5961: "Canary",
		112:  "Canberra",
		1136: "Candera",
		3025: "Canesta",
		3576: "Canhold",
		3270: "Canko",
		87:   "Canon",
		4179: "Canopus",
		5493: "Canstar",
		4055: "Cantronic",
		2677: "Cap",
		6810: "Capelec",
		3002: "Capelon",
		1403: "Capinfo",
		6853: "Capisco",
		1151: "Caporis",
		9424: "CaptionCall",
		5720: "Captor/SA",
		6480: "Cara",
		2759: "Carallon",
		2173: "CardioNet",
		7052: "Cardiopulmonary",
		5677: "Cardkey",
		6938: "CarePredict",
		9305: "CareView",
		4848: "Carecom",
		9441: "Carefusion",
		3007: "Caretech",
		6953: "Cargt",
		3049: "Carina",
		8981: "Carma",
		8901: "Carnegie",
		3840: "Carpoint",
		4288: "Carrera",
		370:  "Carrier",
		222:  "CarrierComm",
		8404: "Carry",
		3170: "Casa",
		7136: "Casacom",
		5183: "Cascade",
		8376: "Cashmaster",
		6345: "Casio",
		3144: "Caspian",
		3831: "CastGrabber",
		2090: "Castel",
		44:   "Castelle",
		1401: "Castle",
		3797: "Castlenet",
		3921: "Castles",
		2067: "Catalyst",
		552:  "Catapult",
		3307: "Catcher",
		2494: "Category",
		373:  "Catena",
		1531: "Caterpillar",
		8045: "Cathay",
		3493: "Cathexis",
		2796: "Cavera",
		2249: "Cavium",
		4504: "Cayee",
		91:   "Cayman",
		7393: "Cdoubles",
		4040: "Cdvi",
		7065: "Cedes",
		6590: "Cedint-UPM",
		4793: "Cedo",
		1032: "Ceemax",
		2448: "Ceicom",
		2853: "CelPlan",
		5959: "Celan",
		5348: "Celcore",
		3776: "Celeno",
		6212: "Celestica",
		218:  "Celestix",
		4038: "Celio",
		1093: "Celleritas",
		7597: "Cellient",
		2186: "Cellink",
		1518: "Cellinx",
		7255: "Cello",
		6136: "Cellport",
		483:  "Cellvision",
		8600: "Celona",
		993:  "Celsian",
		3640: "Celtro",
		3957: "Centec",
		3750: "Center",
		5587: "Centigram",
		5698: "Centillion",
		173:  "Centillium",
		266:  "Centos",
		3124: "CentraLite",
		4695: "Centrak",
		6862: "Centripetal",
		4712: "Centrofactor",
		5870: "Centrum",
		5499: "Century",
		8359: "CenturyLink",
		2324: "Ceologic",
		4321: "Cepheid",
		6591: "Cepton",
		8590: "Cera",
		8940: "CeraMicro",
		1490: "Ceragon",
		9019: "Cercacor",
		6625: "CerebrEX",
		7765: "Cerebras",
		8126: "Cerio",
		2867: "Cermate",
		6313: "Cern",
		4302: "Cernium",
		4400: "Certicom",
		2434: "Cesnet",
		2727: "Cetacea",
		1396: "Cetacean",
		140:  "Cetia",
		3506: "Cetis",
		4416: "Ceton",
		3039: "Ceyon",
		8034: "Chabrier",
		3432: "Chainleader",
		5241: "Chaintech",
		3899: "Chainzone",
		6442: "Chameleon",
		9174: "Championtech",
		7806: "ChangYang",
		6653: "Changhe",
		5391: "Channelmatic",
		2262: "Chantry",
		8428: "ChargeStorm",
		467:  "Charles",
		5549: "Chase",
		4940: "Checkout",
		3234: "Checkpoint",
		4896: "Cheerchip",
		7509: "Cheerstar",
		3029: "Cheertek",
		1079: "Chelsio",
		7726: "Chemoptics",
		7074: "Chemtronics",
		3709: "Cherry",
		4247: "Chess",
		509:  "Chiaro",
		1797: "Chic",
		7300: "Chicony",
		1742: "Chih-Kan",
		2554: "Chinasys",
		702:  "Chino",
		4313: "Chip-pro",
		605:  "Chip2Chip",
		6649: "ChipER",
		6359: "Chipcom",
		6439: "Chipsea",
		7826: "Chipsip",
		2488: "Chiron",
		7437: "Chitai",
		7140: "Chiyoda",
		2123: "Chiyu",
		4582: "Chroma",
		205:  "Chromatek",
		28:   "Chromatics",
		5194: "Chromatis",
		336:  "Chromisys",
		8639: "Chromlech",
		9286: "Chrontel",
		2987: "ChuanG",
		3779: "Chubb",
		6270: "Chun-il",
		5916: "Chuntex",
		5178: "Chuo",
		36:   "Chyron",
		717:  "Ciac",
		2745: "Ciara",
		4684: "Ciat",
		478:  "Cidco",
		1987: "Cidra",
		4564: "Ciena",
		7840: "Ciesse",
		7533: "Ciholas",
		8056: "Cilag",
		2361: "Cimetrics",
		114:  "Cimlinc",
		6208: "Cinco",
		6872: "Cincoze",
		4453: "Cinetal",
		781:  "Cinta",
		2795: "Cintech",
		5868: "Cipher",
		1455: "CipherOptics",
		4206: "Ciprico",
		3836: "Circleone",
		4013: "CiriTech",
		304:  "Cirilium",
		898:  "Cirkitech",
		3805: "Cirrascale",
		2:    "Cisco",
		5367: "Citadel",
		1033: "Citel",
		5711: "Citicorp/TTI",
		3485: "Citiway",
		9256: "Citrix",
		2585: "Citronix",
		5017: "City-Net",
		2824: "CityCom",
		6259: "Clack",
		7258: "ClarIDy",
		3897: "ClariPhy",
		5317: "Clariion",
		2369: "Clarinet",
		7129: "Clavister",
		823:  "ClearCube",
		5632: "ClearOne",
		4703: "ClearPath",
		4303: "Clearbox",
		5533: "Clearpoint",
		3907: "Clearwire",
		2070: "Clematic",
		8137: "Cleondris",
		5687: "Clevo",
		337:  "ClickTV",
		5653: "Clientec",
		3914: "Climax",
		6384: "CliniCare",
		1110: "Clipcomm",
		6230: "Cloos",
		9270: "CloudGenix",
		1701: "CloudShield",
		7544: "CloudSwitch",
		3189: "Cloudastructure",
		9161: "Cloudena",
		8635: "Cloudistics",
		6664: "Cloudium",
		9130: "Cloudleaf",
		6889: "Cloudtronics",
		8467: "Cloudview",
		5434: "Cloudwalk",
		4508: "Clover",
		1378: "Cloverleaf",
		580:  "Clovertech",
		1356: "Cluster",
		82:   "Clustrix",
		2326: "Cmicro",
		7327: "Cmitech",
		4980: "Cmos",
		5536: "Cnet",
		4452: "Cnrs",
		7928: "CoBI",
		8717: "CoDACO",
		2144: "CoM&C",
		2956: "CoM.s.a.t",
		5749: "CoM21",
		6840: "CoMESTA",
		4776: "CoMFILE",
		8112: "CoMMANDO",
		4855: "CoMMidt",
		7972: "CoMTEC",
		4577: "CoMXION",
		6578: "CoNCH",
		7364: "CoNELCOM",
		6720: "CoNSTELL8",
		7329: "CoNTROLtronic",
		9483: "CoNWIN.Tech",
		3737: "CoNWISE",
		285:  "CoNet",
		3040: "CoOLKSKY",
		2113: "CoPAN",
		5392: "CoRELIS",
		4372: "CoRNELL",
		570:  "CoT",
		3786: "CoVAX",
		2781: "CoVi",
		3990: "CoWON",
		4747: "CoachComm",
		9382: "Coagent",
		5747: "Coastcom",
		2487: "Coaxial",
		9329: "CobaltRay",
		8429: "Cobham",
		6847: "Cobs",
		8772: "Coby",
		8036: "Cochlear",
		2423: "Cocom",
		2651: "Codan",
		1740: "Code",
		2841: "Coder",
		8603: "Codetek",
		6320: "Codex",
		1943: "Codian",
		5147: "Codonics",
		7436: "Coflec",
		9194: "Cofractal",
		1675: "Cogent",
		5982: "Cognex",
		2245: "Cognio",
		9032: "Cognitas",
		7059: "Cognitec",
		7452: "Cognitive",
		8678: "Cohere",
		3259: "Coherent",
		7221: "Cohesity",
		1458: "Cohu",
		1914: "Collex",
		6708: "Collinear",
		9167: "Color-Chip",
		7001: "ColorTokens",
		5075: "Colorgraph",
		512:  "Colubris",
		6340: "ComDesign",
		4481: "ComWorth",
		7184: "Comat",
		7615: "CombiQ",
		5181: "Combinet",
		6030: "Comcam",
		5967: "Comda",
		439:  "Comdial",
		5939: "Comelta",
		5108: "Comendec",
		1733: "Comflux",
		8777: "Comigo",
		8157: "Comlab",
		4381: "CommAgility",
		9056: "CommFront",
		5076: "CommScope",
		6981: "CommSky",
		1224: "Command-e",
		1641: "Commax",
		1842: "Commend",
		3418: "CommerceGuard",
		3438: "Commerciant",
		1150: "Commil",
		5457: "Commodore",
		6804: "Common",
		6052: "Commscope",
		5524: "Commscraft",
		7816: "Commsen",
		4831: "Commtact",
		757:  "Commtech",
		4423: "Communication",
		2392: "Communications",
		5771: "Commvision",
		7281: "Comnect",
		5020: "Comone",
		2002: "CompXs",
		5079: "Compac",
		9026: "Compacta",
		358:  "Compal",
		4101: "Compass",
		7108: "Compass-EOS",
		1906: "Compellent",
		5081: "Compex",
		767:  "Complementary",
		3392: "Compro",
		130:  "Compu-Shack",
		275:  "CompuLab",
		5450: "Compuadd",
		1775: "Compucase",
		55:   "Compucorp",
		1887: "Compulogic",
		4413: "Compumedics",
		1566: "Compunetix",
		7856: "Compupal",
		5417: "Compuserve",
		6064: "Computational",
		3651: "Computec",
		2328: "Computer",
		5151: "Computerm",
		6299: "Computervision",
		5701: "Computex",
		3993: "Computime",
		5497: "Computone",
		1325: "Computrols",
		6149: "Compuware",
		5766: "Comsat",
		8497: "Comsis",
		5068: "Comsoft",
		594:  "Comspace",
		3982: "Comsys",
		6100: "Comtec",
		4260: "Comtech",
		5940: "Comtree",
		3869: "Comtrend",
		308:  "Comtrol",
		4208: "Comtron",
		7397: "ConMet",
		2574: "ConSentry",
		8641: "Conception",
		4500: "Conceptronic",
		66:   "Concord",
		5140: "Concurrent",
		3707: "Condalo",
		485:  "Condev",
		5491: "Condor",
		8344: "Conductix-Wampfler",
		7297: "Conet",
		5024: "Conexant",
		6882: "Conextop",
		5368: "ConferTech",
		3898: "Confidant",
		2740: "Congatec",
		4931: "Congruency",
		223:  "Conklin",
		7112: "Conlog",
		22:   "Connect",
		2660: "ConnectBlue",
		7060: "ConnectQuest",
		6008: "Connected",
		912:  "Connection",
		8446: "Connex",
		1713: "Connexionz",
		281:  "Conrad",
		1942: "Consensys",
		8608: "Consert",
		4373: "Consilium",
		2832: "Consultronics",
		4662: "ConteXtream",
		5482: "Contec",
		3167: "Contela",
		2897: "Contemporary",
		1276: "Contex",
		6610: "Continental",
		5381: "Continuum",
		1998: "Control",
		2292: "Control4",
		1856: "ControlNet",
		8638: "ControlThings",
		4959: "Controlled",
		5347: "Controlware",
		970:  "Convedia",
		8668: "Convergence",
		5009: "Convergenet",
		3507: "Convergens",
		5972: "Convergent",
		8284: "ConversDigital",
		5793: "Convex",
		4699: "Convey",
		6086: "Convision",
		110:  "Conware",
		4852: "Coordiwise",
		5116: "Copernique",
		5692: "CopperCom",
		2725: "CorEdge",
		5780: "Cordant",
		5239: "Corder",
		475:  "Core",
		1580: "CoreBell",
		6373: "CoreEdge",
		8619: "CoreOS",
		2888: "CoreStar",
		6624: "CoreTrust",
		5651: "Corecess",
		3587: "Corelatus",
		8392: "Coresys",
		4163: "Coretree",
		5220: "Coretronic",
		6680: "Coriant",
		1714: "Corinex",
		8291: "Corintech",
		386:  "Coriolis",
		8175: "Cornami",
		489:  "Cornet",
		4949: "Corning",
		5118: "Corollary",
		2849: "Corona",
		1609: "Corrent",
		1291: "Corrigent",
		6799: "Corsa",
		2678: "Cortina",
		9456: "Cortland",
		1133: "Corvalent",
		4420: "Corventis",
		5243: "Corvis",
		6329: "Corvus",
		7951: "Cosco",
		6032: "Cosine",
		1981: "Cosmic",
		354:  "Cosmo",
		4779: "Costar",
		9312: "Costel",
		6182: "Cosworth",
		5204: "Cosystems",
		2484: "Cotron",
		9206: "Cots",
		4677: "Coulomb",
		5715: "Coulter",
		9375: "CountMax",
		4265: "Counter",
		6334: "Counterpoint",
		7644: "Countwise",
		1163: "Coup",
		6718: "Coval",
		2910: "Covergence",
		4655: "Covia",
		3886: "Covote",
		3841: "Cowbell",
		8694: "Cox",
		2384: "Coyote",
		7269: "Cozybit",
		4989: "Cqos",
		1155: "Cradle",
		4958: "CradlePoint",
		6102: "Cradlepoint",
		3802: "Cranite",
		618:  "Cratos",
		68:   "Cray",
		8339: "Creatcomm",
		7045: "Creation",
		357:  "Creative",
		6880: "Creatz",
		9116: "Cree",
		6775: "Crenus",
		8257: "Creowave",
		2030: "Crere",
		3925: "Crescendo",
		6058: "Crescent",
		2352: "Crestron",
		3282: "Cresyn",
		1541: "Creval",
		384:  "Crewave",
		1519: "Crinis",
		5786: "Cristie",
		7180: "Criticare",
		6300: "Cromemco",
		1400: "Cronyx",
		585:  "Crossbeam",
		3549: "Crossbow",
		103:  "Crosscomm",
		902:  "Crossport",
		6092: "Crossroads",
		2747: "Crow",
		94:   "Cryptek",
		5597: "Crypto",
		3979: "CryptoMetrics",
		2851: "Cryptosoft",
		4032: "Cryptsoft",
		5705: "Csir",
		6126: "Cspi",
		6256: "Csst",
		6771: "Ctek",
		3466: "Ctring",
		2959: "Cube",
		4249: "Cubix",
		9353: "Cuciniale",
		8108: "Cudo",
		8758: "Cuff",
		8569: "Cummings",
		3846: "Cummins",
		4590: "Cummins-Allison",
		7171: "Cumulus",
		8897: "Curiouser",
		6453: "Currant",
		8149: "Current",
		3656: "Curtis",
		7137: "Curvalux",
		3043: "Curves",
		2122: "Custom",
		4809: "Cutera",
		6055: "Cutler-Hammer",
		7362: "CviLux",
		3742: "Cwlinux",
		3972: "CyVerse",
		8801: "Cybelec",
		677:  "Cyber",
		8737: "Cyber-Rain",
		2592: "CyberNet",
		227:  "CyberOptics",
		1761: "CyberPower",
		189:  "CyberTAN",
		8740: "Cybera",
		642:  "Cyberboard",
		4287: "Cyberdata",
		4399: "Cyberdyne",
		5496: "Cybergraphic",
		1654: "Cybernetics",
		4275: "Cybertec",
		4134: "Cybertech",
		7491: "Cybertelbridge",
		3950: "Cybiotronics",
		6854: "Cybo",
		5329: "Cyclades",
		5344: "Cycle",
		8258: "Cydle",
		215:  "Cygnet",
		5723: "Cylink",
		3213: "Cymphonix",
		4301: "Cymtec",
		7902: "Cynosure",
		8687: "Cynove",
		4482: "Cypak",
		4780: "Cypress",
		4997: "Cyra",
		481:  "Cyras",
		2967: "Cytyc",
		873:  "D&M",
		2200: "D&R",
		4758: "D-BOX",
		803:  "D-Link",
		2991: "D-MAX",
		1748: "D-NET",
		4333: "D-TACQ",
		7944: "D.SignT",
		4427: "DAC",
		3271: "DACOS",
		3201: "DAGS",
		1070: "DANCONTROL",
		876:  "DAP",
		2676: "DAQ",
		7408: "DASAN",
		8734: "DASCOM",
		6196: "DATENTECHNIK",
		8937: "DATTUS",
		7521: "DAVOLINK",
		6683: "DB",
		9186: "DBG",
		7007: "DBL",
		49:   "DCI",
		8926: "DCONWORKS",
		5403: "DCS",
		4622: "DCT-Delta",
		5964: "DDL",
		4526: "DDRdrive",
		9266: "DDoS-Guard",
		2255: "DELCOMp",
		4615: "DELEC",
		7174: "DEOTRON",
		7625: "DEP",
		9442: "DEPO",
		8340: "DERA",
		8596: "DEXON",
		2483: "DEXTER",
		180:  "DFI",
		2968: "DFM",
		2976: "DG2L",
		2673: "DGSTATION",
		6150: "DH",
		7896: "DHC",
		1520: "DHD",
		9014: "DIAODIAOTechnology",
		7241: "DIGALOG",
		1999: "DIGINICS",
		3942: "DINEC",
		8372: "DIRECTV",
		8159: "DK",
		3469: "DKT",
		8845: "DLOGIC",
		9263: "DLX",
		8503: "DMATEK",
		3733: "DMP",
		397:  "DNE",
		4666: "DO-Monix",
		5670: "DPAC",
		2354: "DPS",
		9366: "DQ",
		5030: "DResearch",
		4698: "DS",
		5053: "DSA",
		5323: "DSC",
		5042: "DSG",
		4167: "DSP",
		7099: "DSPECIALISTS",
		6236: "DSPWorks",
		7541: "DSSD",
		7199: "DT",
		4893: "DTI",
		2168: "DUALi",
		6110: "DUX",
		799:  "DVC",
		161:  "DVICO",
		6736: "DVL",
		9302: "DVN",
		5359: "DVS",
		5625: "DVT",
		7446: "DWnet",
		5509: "DY-4",
		6203: "DZS",
		1814: "DaTARIUS",
		5514: "Dacoll",
		6777: "Dadoutek",
		1099: "Daehanet",
		742:  "Daeryung",
		9429: "Daeshin-Information",
		5617: "Daewoo",
		6592: "Daewoois",
		6452: "Dags",
		504:  "Daiden",
		1085: "Daihen",
		1604: "Daiichi",
		4123: "Daiichi-Dentsu",
		5483: "Daikin",
		4445: "Daintree",
		1465: "Daisy",
		1454: "Daktronics",
		1605: "Dallmeier",
		7954: "Damalisk",
		3056: "Dametric",
		80:   "Dana",
		7817: "Danal",
		784:  "Danam",
		2533: "Danelec",
		1096: "Danfoss",
		2872: "Daniels",
		6395: "Danlaw",
		3441: "Danpex",
		3129: "Dansensor",
		4279: "Dantel",
		9068: "Daontec",
		422:  "Daphne",
		8085: "DarbeeVision",
		3232: "Darts",
		6067: "Dasan",
		111:  "Dassault",
		2568: "Data",
		7203: "DataCore",
		4794: "DataFab",
		8813: "DataGravity",
		1139: "DataLogic",
		1506: "DataPower",
		8922: "DataRemote",
		2265: "DataWind",
		116:  "Datability",
		2212: "Datacap",
		5065: "Datacom",
		5032: "Datacore",
		5888: "Dataexpert",
		9199: "Datafox",
		5429: "Datafusion",
		4310: "Dataline",
		5879: "Datalux",
		1937: "Datamax",
		5510: "Datamedia",
		4202: "Datametrics",
		5756: "Dataplex",
		10:   "Datapoint",
		1969: "Dataprobe",
		5866: "Dataproducts",
		4814: "Dataram",
		953:  "Datasound",
		2908: "Datastore",
		4626: "Datastrip",
		5109: "Datatech",
		4278: "Datatrek",
		3720: "Dataupia",
		2190: "Datawire",
		6202: "Datax",
		7673: "Datecs",
		1732: "Datel",
		1022: "Dateno",
		5218: "Datong",
		8430: "Datrium",
		3449: "Datum",
		5525: "Datus",
		5956: "Dauphin",
		2985: "Dave",
		5840: "David",
		3719: "Daviscomms",
		1211: "Davolink",
		3304: "Dawevision",
		5752: "Dawn",
		4842: "DaySequerra",
		5462: "Dayna",
		8118: "Dayouplus",
		4364: "Dbii",
		2413: "Dbtel",
		5257: "Dctri",
		8433: "DeLaval",
		7694: "Decatur",
		6220: "Decision",
		2060: "Decru",
		9392: "Dediprog",
		7432: "Deeplet",
		3687: "Deepsound",
		3192: "Defidev",
		8273: "Definium",
		2811: "Deicy",
		4856: "Deif",
		7638: "Dejai",
		5614: "Delem",
		102:  "Dell",
		6825: "DellKing",
		3585: "Delorme",
		8969: "Delphian",
		7508: "Delphin",
		957:  "Delta",
		4519: "Deltacom",
		7457: "Deltanet",
		3527: "Deltanode",
		9436: "Deltenna",
		689:  "Deltron",
		3460: "Demant",
		6791: "Dematic",
		4603: "Demco",
		131:  "Densan",
		5889: "Denso",
		2211: "Dentum",
		9412: "Denwa",
		815:  "Deonet",
		8685: "Derek",
		7102: "DerekLimited",
		592:  "Desana",
		7224: "Desay",
		3861: "Design",
		4647: "Design-Com",
		4663: "DesignArt",
		6125: "DeskStation",
		5763: "Desknet",
		9055: "Detracom",
		7549: "Deutron",
		3010: "Develco",
		133:  "Develcon",
		7953: "Devialet",
		6891: "DeviceDesign",
		2065: "Devicescape",
		6747: "Devlin",
		6845: "Devline",
		7057: "Dewar",
		7749: "Dewav",
		7069: "Dexatek",
		5610: "Dexdyne",
		7582: "Dexin",
		3974: "DiMoto",
		2332: "Diablo",
		1005: "Diagraph",
		3806: "Dial",
		2244: "Dialog",
		5353: "Dialogic",
		1369: "Dialogue",
		9216: "Diamanti",
		1025: "Diamond",
		6094: "Diba",
		1303: "Dibal",
		3104: "Diboss",
		235:  "Dica",
		6084: "Dictaphone",
		511:  "Diebold",
		9368: "Diffon",
		5167: "Digalog",
		968:  "Digeo",
		5649: "Digi-Data",
		5139: "DigiBoard",
		911:  "DigiPower",
		1980: "DigiRose",
		5283: "Digianswer",
		6211: "Digicom",
		2143: "Digicube",
		6448: "Digience",
		3968: "Digifriends",
		6017: "Digigram",
		3303: "Digilent",
		8167: "Digimore",
		2501: "Diginfo",
		7688: "Digiquest",
		3176: "Digit",
		6797: "Digita",
		1565: "Digital",
		5645: "DigitalScape",
		515:  "DigitalSis",
		2900: "DigitalZone",
		4697: "Digitalbox",
		5635: "Digitalcast",
		2142: "Digitalks",
		9159: "Digitalwatt",
		4835: "Digitec",
		5731: "Digitech",
		496:  "Digitel",
		3912: "Digitize",
		4972: "Digitra",
		3738: "Digitrax",
		4081: "Digitview",
		2268: "Digium",
		2850: "Digiwell",
		3739: "Dignsys",
		2975: "Dilithium",
		9433: "Dinstar",
		6214: "Dionex",
		3567: "Diostech",
		7809: "DirectPacket",
		794:  "Disco",
		6890: "Discovergy",
		8312: "Discovery",
		1238: "Dish",
		1192: "Diskbank",
		2775: "Diskware",
		1907: "DispenseSource",
		8469: "Displaire",
		8858: "Display",
		2232: "DisplayLink",
		8289: "Disruptive",
		5966: "Ditech",
		1922: "DivR",
		7733: "DivUS",
		1209: "DivergeNet",
		8221: "Diversey",
		4180: "Diversified",
		4934: "Divio",
		4531: "Dizipia",
		8203: "Dlogixs",
		7963: "Dmet",
		2336: "Dnpg",
		3437: "DoCoMo",
		316:  "DoTop",
		4044: "Doble",
		5544: "Document",
		5898: "Docupoint",
		6001: "Dolby",
		5579: "Domex",
		2480: "Domo",
		5227: "Doms",
		3538: "Dongahelecomm",
		7959: "Dongnian",
		3106: "Donjin",
		6076: "Dooin",
		3439: "Doorking",
		8546: "Doppler",
		2438: "Doremi",
		3873: "Doro",
		4609: "Double-Take",
		4254: "Dovatron",
		118:  "Dove",
		8456: "Dragino",
		1091: "DragonWave",
		8923: "Dragonchip",
		7438: "Dragos",
		3922: "DrayTek",
		3028: "Dream",
		1346: "Dream-Multimedia-Tv",
		3369: "Dreamtech",
		5135: "Dressler",
		733:  "Drew",
		7311: "Drimsys",
		3134: "DriveCam",
		7643: "DriveScale",
		8028: "Drivenets",
		8515: "Drivven",
		8722: "Drogoo",
		6919: "Dropcam",
		8840: "Drtech",
		7430: "Drust",
		7578: "Dryad",
		331:  "Dtvro",
		5921: "Dual",
		7191: "DualShine",
		3399: "Duaxes",
		8947: "Ducere",
		4877: "Duelco",
		9166: "Dunkermotoren",
		6493: "DuroByte",
		6508: "Duskrise",
		3169: "Dust",
		8726: "Dusun",
		9306: "Dycon",
		4673: "Dycor",
		4099: "Dyna",
		6041: "Dynacolor",
		2724: "Dynalab",
		7354: "Dynalec",
		1941: "Dynamic",
		6554: "Dynapower",
		5048: "Dynapro",
		5675: "Dynarc",
		5451: "Dynatech",
		5877: "Dynatem",
		1986: "Dynavac",
		7189: "Dyson",
		7921: "E-Band",
		248:  "E-Control",
		7755: "E-Domus",
		8577: "E-Fuel",
		476:  "E-Globaledge",
		8207: "E-Lead",
		5171: "E-M",
		4573: "E-Mon",
		8605: "E-Prime",
		4027: "E-Senza",
		5110: "E-Systems./Garland",
		2713: "E-TEC",
		7915: "E-TRON",
		5634: "E-Tech",
		455:  "E.D.&A",
		6116: "E.E.P.D",
		2243: "E2O",
		2619: "E2S",
		7168: "E3",
		2705: "EAB/RWI/K",
		4390: "EASY3CALL",
		7376: "EBN",
		4178: "EBRAINS",
		2999: "ECA-Sinters",
		5609: "ECCS",
		3767: "ECKey",
		2679: "ECM",
		8442: "ECOtality",
		6842: "EControls",
		3631: "ED",
		6456: "EDIC",
		3246: "EDM",
		6661: "EDMI",
		2340: "EDNT",
		3465: "EDO-EVI",
		229:  "EDSL",
		3316: "EDSLAB",
		4716: "EDT",
		1831: "EE",
		6177: "EES",
		1252: "EFM",
		1590: "EG",
		3032: "EGO",
		6679: "EID",
		2307: "EION",
		5201: "EIS",
		5646: "EIZO",
		8232: "EKE",
		4362: "EKE-Electronics",
		3336: "EL-Tech",
		645:  "ELANsat",
		6077: "ELE-Chem",
		715:  "ELESIGN",
		8005: "ELFTECH",
		1771: "ELIAS",
		7459: "ELL-IoT",
		1738: "ELM",
		214:  "ELMEX",
		2632: "ELPRO",
		7660: "ELS-GmbH",
		1245: "EM",
		8739: "EM-Tech",
		6823: "EMAC",
		4351: "EMC",
		445:  "EMPEG",
		9390: "EMTAKE",
		6478: "EMU",
		8286: "EMW",
		4327: "EN",
		611:  "ENEGATE",
		3498: "ENENSYS",
		8688: "ENERES",
		7938: "ENMAS",
		8370: "ENVINET",
		8414: "EOC",
		126:  "EON",
		5946: "EOS",
		3365: "EOps",
		7738: "EPCOMM",
		2433: "EPIN",
		3374: "EPL",
		671:  "EPOX",
		4583: "EQ-Sys",
		9154: "EQUES",
		6947: "ER-Telecom",
		5498: "ERI",
		2578: "ERUNE",
		4679: "ESCATRONIC",
		340:  "ESD",
		8155: "ESG",
		4963: "ESI",
		4385: "ESP",
		9388: "ESPOD",
		249:  "ESS",
		9133: "ESSENCORE",
		5709: "EST",
		3456: "ESTIC",
		2875: "ESU",
		7361: "ESYLUX",
		1674: "ET&T",
		451:  "ETAS",
		581:  "ETEN",
		7830: "ETH",
		4054: "ETL",
		6451: "EUCAST",
		2347: "EUREM",
		7943: "EVBox",
		4140: "EVGA",
		282:  "EVR",
		7456: "EVRsafe",
		8355: "EVUlution",
		4389: "EW3",
		2267: "EWA",
		5425: "EXP",
		8924: "EXTEN",
		6099: "EXXACT",
		5835: "Eacem",
		4267: "Eagle",
		6498: "Earda",
		3211: "Earforce",
		8775: "EarthCam",
		7747: "Eastcompeace",
		4753: "Eastern",
		1413: "Eastmode",
		6455: "Eastriver",
		7870: "EasySYNC",
		1855: "Eaton",
		1818: "Eazix",
		7847: "Ebay",
		1841: "Ebtron",
		2092: "Ecastle",
		5878: "Ecci",
		5789: "Ecessa",
		6026: "Echelon",
		3743: "Echo360",
		6950: "Echodyne",
		2636: "Echolab",
		9331: "Echosens",
		9370: "Echowell",
		8508: "Eclipse",
		9444: "EcoTech",
		7366: "Ecocom",
		3963: "Ecolab",
		9073: "Ecoliv",
		2760: "Ecom",
		6957: "Ecosense",
		5014: "Ectel",
		6003: "Ecton",
		2816: "Edata",
		5315: "Edec",
		8296: "Edeltech",
		5235: "Edeva",
		2842: "Edge",
		7774: "EdgeCore",
		6143: "EdgePoint",
		9481: "EdgeQ",
		4153: "EdgeVelocity",
		1165: "EdgeWave",
		6294: "Edgecore",
		7389: "Edgewater",
		6445: "Edifier",
		115:  "Edimax",
		7825: "Edison",
		5620: "Edmi",
		9212: "Edup",
		1805: "Edwards",
		4280: "Efficient",
		412:  "Effinet",
		3691: "Efkon",
		7331: "Egardia",
		883:  "Egenera",
		2157: "Ego",
		8985: "Eidetic",
		3156: "Eidicom",
		6978: "Eidolon",
		2470: "Eidsvoll",
		1629: "Eiki",
		6871: "Eintechno",
		3094: "Eishin",
		6385: "Eito",
		2805: "Eka",
		3340: "Ekahau",
		4966: "Elastic",
		8718: "Elastifile",
		6704: "Elatec",
		627:  "Elau",
		4465: "Elbit",
		5410: "Eldat",
		5693: "Elecom",
		1250: "Elecs",
		6602: "Elecsys",
		210:  "Electro",
		5057: "Electro-Metrics",
		7173: "Electrolux",
		3690: "Electronic",
		6849: "Electronics",
		6175: "Electrosonic",
		3140: "Eleksen",
		3962: "Elelux",
		3764: "Element",
		4652: "Elentec",
		3221: "Elesta",
		8552: "Elesys",
		1172: "Eletex",
		2517: "Elexol",
		2184: "Elextech",
		8434: "Elgama-Elektronika",
		1023: "Elgar",
		8558: "Elim",
		8606: "Eline",
		6672: "EliteGroup",
		1115: "Elitegroup",
		8266: "Elma",
		5411: "Elmic",
		1404: "Elmo",
		7895: "Elno",
		5891: "Elonex",
		2053: "Elphel",
		8708: "Elprotronic",
		1106: "Elrest",
		3562: "Elster",
		8495: "Eltex",
		7519: "EltexAlatau",
		8217: "Elvaco",
		6353: "Elxsi",
		8612: "EmBestor",
		8119: "EmFirst",
		5724: "Email",
		8010: "Embed",
		7494: "EmbedWay",
		6468: "Embedian",
		522:  "Embedone",
		1893: "Embedtronics",
		1936: "Ember",
		8043: "Embertec",
		8875: "Embrane",
		5881: "Emcom",
		1979: "Emcore",
		7101: "Emcraft",
		1195: "EmergeCore",
		3792: "Emergent",
		1965: "Emerging",
		4059: "Emfinity",
		3790: "Emfit",
		6269: "Emicon",
		3905: "Emitech",
		3020: "Emitor",
		8288: "Emizon",
		3807: "Empacket",
		880:  "Empirix",
		7266: "Emplus",
		2822: "Empower",
		9362: "Emstone",
		5405: "Emtrakorporated",
		8948: "Emtronix",
		129:  "Emulex",
		6060: "Emutec",
		1646: "Emuzed",
		8103: "EnGenius",
		7071: "EnTek",
		4812: "Enabling",
		4773: "Enalasys",
		9007: "Enatel",
		2303: "Encanto",
		8142: "Encell",
		6732: "Enclustra",
		3302: "Encore",
		433:  "EndPoints",
		2139: "EndRun",
		2085: "Endace",
		2172: "Endeleo",
		2611: "Endevco",
		6977: "Enecsys",
		8261: "Enelps",
		7404: "EnerAccess",
		1260: "EnerLinx.com",
		2597: "Enerdyne",
		9108: "Energica",
		9231: "Energotest",
		4727: "Energy",
		4882: "EnergyHub",
		4625: "EnergyICT",
		3251: "Enermet",
		6758: "Enernet",
		1502: "Enerpoint",
		9496: "Enervalley",
		5960: "Engage",
		2693: "Engim",
		5395: "Enginuity",
		7480: "Enguity",
		179:  "EnjoyWeb",
		7958: "Enmotus",
		200:  "Ennovate",
		9035: "Eno",
		204:  "Ensemble",
		4393: "Enseo",
		8627: "Ensequence",
		5038: "Ensim",
		8517: "Enspert",
		986:  "Ensure",
		2890: "Enswer",
		312:  "Enterasys",
		7922: "Entertainment",
		6560: "Entis",
		1422: "Entise",
		596:  "Entone",
		4433: "Entorian",
		4725: "Entourage",
		565:  "Entra",
		6181: "Entrada",
		2383: "Entrata",
		5301: "Entrega",
		2189: "Entrelogic",
		5612: "Entridia",
		1221: "Entrisphere",
		1394: "Entropic",
		6750: "Entropix",
		1298: "Envenergy",
		8724: "Environics",
		3849: "Environnement",
		3755: "Envisacor",
		4910: "Envisionnovation",
		8313: "EnvyLogic",
		1772: "Enwiser",
		3813: "Enzytek",
		6499: "EoCell",
		6874: "Eolo",
		1129: "Eolring",
		7799: "Eoptolink",
		9464: "EosTek",
		8746: "Eoslink",
		7754: "EpSpot",
		8764: "Epec",
		4049: "Epic",
		2112: "Epicenter",
		2669: "Epicom",
		5613: "Epigram",
		6070: "Epilog",
		1332: "EpoX",
		46:   "Epson",
		1425: "Epygi",
		1393: "EqualLogic",
		3674: "Equaline",
		1131: "Equator",
		4985: "Equiinet",
		5508: "Equinox",
		159:  "Equip'Trans",
		461:  "Equipe",
		7153: "Equitech",
		718:  "Equitrac",
		3486: "Equustek",
		8389: "EraThink",
		1521: "Eracom",
		3368: "Erae",
		7646: "Ereca",
		7848: "Ergophone",
		7691: "Erhardt+Leimer",
		306:  "Ericsson",
		8567: "Ericsson-LG",
		2014: "Ermme",
		4252: "Ersat",
		1270: "Erskine",
		7322: "Es-tech",
		4676: "Esab",
		9292: "Esan",
		5688: "Escalate",
		3388: "Escape",
		4322: "Escherlogic",
		9038: "Escort",
		7032: "Espec",
		1188: "Espera-Werke",
		6828: "Espotel",
		6376: "Espressif",
		5459: "Esprit",
		9335: "Essec",
		6629: "Essel-T",
		4849: "Essensium",
		5758: "Essential",
		3522: "Essilor",
		7971: "Esson",
		6830: "Essys",
		2539: "Estari",
		7559: "Estech",
		1318: "Esteem",
		6816: "Etek",
		7261: "Eterna",
		6180: "EtherWAN",
		5804: "Ethercom",
		4091: "Etherstack",
		1491: "Etherstuff",
		7687: "Ethertronics",
		7240: "EthoSwitch",
		3592: "Ethos",
		1712: "Etin",
		5642: "Etrend",
		7064: "Etronic",
		3442: "Etrovision",
		8364: "Ettus",
		3268: "Etymotic",
		2862: "Eubus",
		3556: "Euchner+Co",
		7387: "Euclid",
		8927: "Eukrea",
		9408: "Eumtech",
		3801: "Eunicorn",
		4857: "Euphonic",
		5672: "Euphonix",
		982:  "Euracom",
		9463: "Eurecam",
		2531: "Eurilogic",
		3101: "EuroCB",
		3655: "EuroTel",
		2261: "Eurocom",
		4769: "Eurodesign",
		502:  "Eurologic",
		7015: "Euronda",
		2248: "Europlex",
		5313: "Eurotech",
		1543: "Eurotherm",
		4199: "Eurotime",
		6544: "Eutronix",
		8240: "Evantage",
		1804: "Eve",
		3035: "Eventide",
		2432: "EverFocus",
		2097: "Everbee",
		6975: "Everest",
		929:  "Everex",
		4225: "Evergreen",
		9467: "Everspin",
		4796: "Everspring",
		6969: "Evervictory",
		7222: "Everysight",
		3627: "Evolis",
		8124: "Evoluzn",
		7846: "Evrisko",
		1477: "ExPet",
		3810: "ExaDigm",
		7585: "Exablaze",
		8618: "Exablox",
		6035: "Exabyte",
		3166: "Exalt",
		5848: "Exar",
		3018: "Exartech",
		2744: "Exavera",
		4184: "Excel",
		6305: "Excelan",
		6197: "Excellent",
		3799: "Excelpoint",
		3839: "Exegin",
		2399: "Exelis",
		4690: "Exeltech",
		3082: "Exeo",
		466:  "Exfo",
		447:  "ExiO",
		4924: "Expand",
		6330: "ExperData",
		3323: "Expertise",
		9104: "Explora",
		9138: "Exponent",
		2925: "Extandon",
		5113: "Extended",
		5125: "Extension",
		731:  "Extenway",
		3288: "Exterity",
		169:  "Extratech",
		185:  "Extreme",
		2121: "ExtremeSpeed",
		2751: "Extricom",
		847:  "Extron",
		6304: "Exxon",
		1725: "EyeCross",
		3313: "EyeFi",
		3901: "Eyeheight",
		4330: "Eyeview",
		2131: "Ezquest",
		3125: "Ezurio",
		9295: "F-Secure",
		3396: "F1MEDIA",
		492:  "F3",
		289:  "F5",
		3275: "FA",
		7529: "FACTS",
		9255: "FADU",
		9357: "FAG",
		790:  "FAST",
		9282: "FCA",
		5271: "FDK",
		9372: "FDT",
		1557: "FEI",
		3499: "FEI-Zyfer",
		3816: "FEIG",
		3591: "FFEI",
		9218: "FIRS",
		6481: "FLECTRON",
		3723: "FLIR",
		2242: "FMN",
		6398: "FMTech",
		6440: "FN-Link",
		7183: "FONsystem",
		2381: "FOR-A",
		7527: "FORICS",
		3335: "FORMOSA21",
		7735: "FRAMOS",
		4822: "FRC",
		2635: "FSI",
		2206: "FTA",
		3085: "FUHO",
		7871: "FUJITU",
		9272: "FUSION",
		3184: "FXC",
		9021: "FXI",
		708:  "FabiaTech",
		1966: "Fabric7",
		1065: "Fabricom",
		7234: "Facebook",
		1956: "Factum",
		6855: "Fagor",
		8773: "Fairfield",
		8054: "Fairphone",
		1940: "FalconStor",
		1444: "Falt",
		4443: "FamilyPhone",
		7020: "Fanhattan",
		3280: "Fanstel",
		5552: "Fantum",
		6427: "Fanvil",
		8164: "Far-sighted",
		2690: "Fargo",
		8796: "Farlink",
		7164: "Farmage",
		298:  "Fast",
		6608: "Fastback",
		3540: "Faster",
		541:  "Fastfame",
		2553: "Fastrax",
		1265: "Fastwel",
		9469: "Favite",
		7586: "Feitian",
		328:  "Fenecom",
		8593: "Fengfan",
		6790: "Fermax",
		5488: "Fermilab",
		6321: "Ferranti",
		2130: "Festo",
		4756: "FiberPlex",
		4367: "Fiberblaze",
		5507: "Fibercom",
		628:  "Fibercycle",
		5952: "Fiberdata",
		8275: "Fibergate",
		4921: "Fiberhome",
		2395: "Fiberlane",
		5485: "Fibermux",
		5157: "Fibernet",
		7070: "Fiberpro",
		5662: "Fibex",
		7413: "Fibrain",
		8411: "Fibrlink",
		821:  "FibroLAN",
		3:    "Fibronics",
		6655: "Ficer",
		7640: "Ficosa",
		8185: "Fida",
		8998: "Fidelix",
		9482: "Fidus",
		4020: "Fidustron",
		6298: "Fihonest",
		7029: "Fijowave",
		8979: "Fike",
		6045: "Filanet",
		6315: "Filenet",
		1360: "FillFactory",
		8907: "Filmetrics",
		2111: "Filtronic",
		5894: "Fine-PAL",
		4237: "Finecom",
		1056: "Finedigital",
		8769: "Finis",
		5622: "Finisar",
		1892: "Finlux",
		3513: "Finnzymes",
		8222: "Finsecur",
		3685: "Finsoft",
		546:  "FireBrick",
		8803: "Fireflies",
		5714: "Firepower",
		3361: "Firetide",
		6721: "Firewalla",
		208:  "Firewiredirect.com",
		8459: "Firich",
		8101: "FirmTek",
		1978: "First",
		1330: "Firstech",
		6973: "Fischer",
		4490: "Fisher-Rosemount",
		6520: "Fisys",
		6587: "Fitbit",
		7162: "Fitview",
		7066: "FiveCo",
		5991: "Fivemere",
		2668: "Flaircomm",
		1072: "Flarion",
		1029: "Flash",
		7243: "Flashbay",
		8408: "FlatFrog",
		409:  "Flatstack",
		3084: "Flex-P",
		3758: "FlexRadio",
		3014: "FlexiPanel",
		1100: "Flexlight",
		7017: "Flexoptix",
		833:  "Flextronics",
		984:  "Flexus",
		8178: "Flonidan",
		3013: "Flovel",
		4227: "Flowpoint",
		5098: "Fluent",
		8290: "Flyaudio",
		5424: "Flytech",
		1042: "Fnet",
		6122: "Focon",
		6632: "Fon",
		8733: "FonSee",
		5907: "Fonsys",
		3595: "Fontal",
		2271: "Force",
		303:  "Force10",
		4880: "Ford",
		5360: "Fore",
		5250: "Foresson",
		7974: "ForgetBox",
		5091: "Forks",
		6487: "FormericaOE",
		9500: "Formike",
		6998: "Formlabs",
		8752: "Fort-Telecom",
		2513: "Fortelink",
		4483: "Fortex",
		6877: "Fortify",
		1323: "Fortinet",
		9431: "Fortress",
		2505: "Fortuna",
		3495: "Fortunetek",
		8002: "Forworld",
		950:  "Fostex",
		5605: "Fountain",
		7836: "Four",
		4518: "Fourier",
		5422: "Fourthtrack",
		6039: "Foveon",
		437:  "FoxJet",
		221:  "Foxconn",
		7313: "Foxda",
		6703: "Foxtech",
		2467: "Francotyp-Postalia",
		2428: "Franklin",
		1720: "Free2move",
		9386: "FreeBit",
		1870: "FreeHand",
		729:  "FreeMs",
		8473: "FreeTek",
		1164: "FreeWave",
		293:  "Freecom",
		3195: "Freedom9",
		6106: "Freegate",
		4408: "Freegene",
		7957: "Freestyle",
		270:  "Frequentis",
		2685: "Fresenius-Vial",
		7285: "Fresh",
		6774: "Friem",
		4058: "From2",
		5487: "Frontier",
		6801: "Frontiir",
		4621: "FuGang",
		1050: "Fujant",
		6986: "Fujikon",
		1698: "Fujikura",
		2138: "Fujinon",
		4:    "Fujitsu",
		8574: "Fulan",
		1773: "FullWave",
		7418: "Fullpower",
		262:  "Fulltek",
		5104: "Funasset",
		5933: "Funk",
		8480: "Furrion",
		3989: "Further",
		4230: "Fusion",
		8642: "Fusion-io",
		3580: "FusionDynamic",
		3087: "Fusiontech",
		9319: "Futaba",
		8556: "Futaba-Kikaku",
		3659: "Futarque",
		8604: "Futecho",
		3149: "Futronic",
		7674: "Futura",
		5105: "Future",
		4418: "FutureLogic",
		878:  "FutureSmart",
		1571: "Futuretel",
		7310: "Fuze",
		5240: "G-Connect",
		9451: "G-Lab",
		2291: "G-PRO",
		8173: "G-Printec",
		1620: "G-Star",
		3746: "G-Technology",
		2039: "G-Tek",
		9497: "G-Wearables",
		6205: "G2",
		6626: "G2C",
		206:  "G3M",
		3250: "GAI-Tronics",
		7903: "GBO",
		4450: "GBS",
		5928: "GCC",
		4290: "GDE",
		4152: "GDI",
		3290: "GDX",
		5810: "GE",
		6363: "GECO",
		1564: "GENETEC",
		1105: "GERSTEL",
		2577: "GES",
		6411: "GET",
		5382: "GEW",
		4561: "GGH",
		8193: "GH",
		4297: "GHI",
		8626: "GHIA",
		6786: "GHT",
		9262: "GIRD",
		9268: "GIT",
		7561: "GITSN",
		5083: "GK",
		4371: "GL",
		4696: "GLOBAL",
		1191: "GM-2",
		4807: "GMK",
		5555: "GMX/Gimix",
		2288: "GN&S",
		2311: "GNP",
		9219: "GNTEK",
		8585: "GOEFER",
		6507: "GOLD3LINK",
		8825: "GOPEACE",
		8017: "GP",
		8086: "GSI",
		5755: "GSM-Syntel",
		3477: "GSP",
		9318: "GST",
		2825: "GSTeletech",
		3525: "GT&T",
		5924: "GVC",
		2343: "GVN",
		2281: "GW",
		8609: "GX",
		3286: "GZ",
		6108: "Gadzoox",
		7525: "Gafachi",
		3935: "GainSpan",
		6860: "Gainspeed",
		5630: "Gaio",
		1447: "Galaxis",
		3847: "Galaxy",
		1921: "Galazar",
		5995: "Galileo",
		8988: "Galleon",
		8298: "Galore",
		1175: "Galtronics",
		4368: "Galvanic",
		1212: "Gamatronic",
		5925: "Gambit",
		8081: "Game",
		4220: "Gammadata",
		9315: "Garland",
		797:  "Garmin",
		5036: "Garnet",
		7275: "Garo",
		4171: "Garrett",
		4221: "GarrettCom",
		1840: "Garuda",
		7741: "Gastron",
		5277: "Gatan",
		2667: "GateConnect",
		2202: "GateWare",
		8172: "Gatekeeper",
		7933: "GatesAir",
		64:   "Gateway",
		5973: "Gateworks",
		3295: "Gatsometer",
		2089: "Gawell",
		1144: "Gcom",
		9004: "Gearlinx",
		9185: "Geberit",
		4915: "Geenovo",
		3808: "Gefen",
		323:  "Gefran",
		7476: "Gehirn",
		2063: "GemWon",
		8096: "Gembird",
		5278: "Gemflex",
		8089: "Gemicom",
		8267: "Geminico",
		9152: "Gemintek",
		1747: "Gemstone",
		1450: "Gemtek",
		6802: "Gen2wave",
		8510: "Genelec",
		3773: "Gener8",
		1854: "Genera",
		2750: "General",
		7686: "Generalplus",
		8095: "Generiton",
		5849: "Genetec",
		931:  "Genetel",
		3800: "Genew",
		2195: "Genexis",
		5205: "Genicom",
		4764: "Genie",
		4204: "Genitech",
		1589: "Gennum",
		6042: "Genoa",
		6779: "Genoray",
		839:  "Genotech",
		5207: "Genrad",
		6186: "Genroco",
		1522: "GentechMedia",
		3524: "Gentex/Electro-Acoustic",
		916:  "Gentner",
		7813: "Gentrice",
		4723: "Genuine",
		2780: "GeoVision",
		7672: "GeoWAN",
		9210: "Geodesic",
		1778: "Geomation",
		891:  "Geospace",
		7976: "Germaneers",
		5951: "Gespac",
		8445: "Gessler",
		9106: "Getck",
		3067: "Geutebruck",
		368:  "Geyser",
		8196: "Gheo",
		1380: "Giant",
		722:  "Giantec",
		7125: "Gifa",
		1929: "Giga-Byte",
		2792: "Giga-byte",
		3151: "Gigabeam",
		2234: "Gigabit",
		4056: "Gigafin",
		7360: "Gigafirm",
		5363: "Gigalabs",
		273:  "Gigalink",
		3420: "Gigamips",
		3923: "Gigamon",
		1310: "Gigaphoton",
		4298: "Gigaset",
		8760: "Gigastone",
		7377: "Gigawave",
		2224: "Gigawavetech",
		203:  "Gilat",
		5251: "Gilbarco",
		7892: "Gimasi",
		8443: "Gimbal",
		2029: "Gincom",
		1134: "Ginganet",
		7845: "Gint",
		4667: "Ginzinger",
		9271: "Giroptic",
		2610: "Gizmondo",
		4396: "Glensound",
		1765: "Glimmerglass",
		3183: "Global",
		6800: "GlobalBeiMing",
		641:  "GlobalStreams",
		1389: "GlobalTop",
		5851: "Globalnet",
		5261: "Globaloop",
		1974: "Globalsat",
		9321: "Globalscale",
		6254: "Globalstar",
		400:  "Globetek",
		2869: "Globo",
		4606: "Glocom",
		743:  "Glonet",
		5738: "Glory",
		4512: "Gloscom",
		7431: "Glovast",
		7700: "Glowforge",
		1782: "Glsystech",
		1182: "Gluon",
		1358: "Glyph",
		8051: "Gnodal",
		2804: "Go",
		6241: "GoPro",
		8134: "Goden",
		3916: "Godex",
		969:  "Goepel",
		6366: "Goertek",
		1120: "Golden",
		5141: "Goldstar",
		4356: "Goliath",
		7602: "GooWi",
		2459: "GoodMan",
		4131: "Goodmill",
		3521: "Google",
		8684: "Googol",
		4509: "Gorba",
		2646: "Gossen-Metrawatt-GmbH",
		6717: "Gotech",
		166:  "Gotham",
		6333: "Gould",
		4439: "Gowell",
		8873: "Goyoo",
		8387: "Gpms",
		7013: "Grabango",
		5325: "Gradient",
		4678: "Graf-Syteco",
		4913: "Grainmustards",
		4064: "Granch",
		1883: "Grand",
		6937: "Grandbeing",
		7119: "Grandex",
		2452: "Grandeye",
		1683: "Grandstream",
		2799: "Grandtec",
		9222: "Graphiant",
		8537: "Graphite",
		5553: "Graphon",
		530:  "Graphtec",
		2304: "Grayhill",
		8170: "Great",
		9422: "Greatek",
		6405: "Green",
		9136: "GreenBytes",
		4017: "GreenLine",
		1514: "GreenNET",
		2971: "GreenPeak",
		9100: "GreenPriz",
		2164: "Greenbell",
		7338: "Greenlee",
		8528: "Greenliant",
		8041: "Greenvity",
		7214: "Greenwald",
		142:  "Grid",
		6805: "GridCentric",
		4553: "GridIron",
		4680: "GridPoint",
		6530: "Gridco",
		9187: "Gridlink",
		3453: "Gridpoint",
		6379: "Gridstore",
		7502: "Gridwiz",
		2614: "Griffin",
		2397: "Grips",
		7008: "Grotech",
		2833: "Grundfos",
		6072: "Grundig",
		4282: "Gtech",
		4128: "Gtran",
		3966: "Gtri",
		4520: "Guangda",
		1542: "Guardware",
		3417: "Gude",
		4545: "Gulfstream",
		3015: "Gumstix",
		7121: "GuoTengShengHua",
		1510: "GyroSignal",
		5512: "H-Three",
		6382: "H2A",
		9077: "HAKKO",
		5137: "HAL",
		6987: "HAN",
		1412: "HARTEC",
		1596: "HARTING",
		2096: "HASHIMOTO",
		1287: "HASSNET",
		9320: "HBC-radiomatic",
		1438: "HBrain",
		71:   "HCL",
		9080: "HCT",
		1495: "HCV",
		4097: "HD",
		8192: "HDR10+",
		5370: "HE",
		2036: "HEINESYS",
		9364: "HENGBAO",
		9189: "HES-SO",
		2880: "HF",
		6483: "HFC",
		4575: "HFR",
		1862: "HGST",
		8496: "HI",
		2911: "HI-P",
		980:  "HID",
		1731: "HIMS",
		1268: "HIT",
		3055: "HIVION",
		9012: "HIWIFI",
		3862: "HM",
		5846: "HME",
		837:  "HMS",
		7560: "HNS",
		53:   "HOB",
		3713: "HOLUX",
		779:  "HORIBA",
		1288: "HOW",
		1627: "HOWTEL",
		416:  "HOYA",
		302:  "HP",
		7295: "HRD",
		5330: "HRK",
		5376: "HT",
		1341: "HTC",
		8205: "HUG-Witschi",
		531:  "HUMAX",
		1753: "HUMEX",
		5931: "HVE",
		9175: "HW",
		6788: "HWH",
		848:  "HYPERCHIP",
		3384: "Hacetron",
		2023: "Hach",
		5901: "Haft",
		4686: "Hager",
		5334: "Hagiwara",
		4767: "Haier",
		7679: "Hailo",
		4683: "Hainzl",
		7154: "Hajen",
		5306: "Hakko",
		5292: "Hakusan",
		4074: "Hakusan.Mfg",
		2785: "Halcro",
		8849: "Halfa",
		1405: "Haliplex",
		8314: "Hame",
		6943: "Hamee",
		7279: "Hamilton",
		1912: "Hammerhead",
		5580: "HanA",
		3178: "Hanazeder",
		1700: "Hanback",
		1216: "Hanbit",
		8471: "HanbitEDS",
		2526: "HandEra",
		3520: "HandHeld",
		8665: "Handaer",
		5583: "Handlink",
		3621: "Handreamnet",
		6902: "Hanilstm",
		4141: "Hanlong",
		2028: "Hannae",
		9289: "Hannto",
		9243: "Hansgrohe",
		4797: "Hanshinit",
		3954: "Hanson",
		3345: "Hansun",
		7910: "Hansung",
		1796: "Hanwang",
		1588: "Happy",
		7785: "Happyelectronics",
		9011: "Harley-Davidson",
		1526: "Harmonic",
		4945: "Harmonix",
		124:  "Harris",
		4973: "Hasselblad",
		3400: "Hasware",
		5060: "Hatteland",
		9110: "Haverford",
		9378: "Hawa",
		7123: "Hawkeye",
		2037: "Hawking",
		8549: "Hay",
		5721: "Hazeltine",
		7479: "Hazemeyer",
		7461: "Hccp",
		8995: "Hdpro",
		6709: "Hdsn",
		9301: "Heat",
		2437: "Hectrix",
		2994: "Hectronic",
		246:  "HeiSei",
		2547: "Heidelberg",
		2838: "Heim",
		8462: "Heimgard",
		5844: "Heinzinger",
		6968: "Heinzmann",
		6098: "Heiwa",
		2878: "Helicomm",
		7577: "Heliospectra",
		2338: "Helioss",
		7539: "Helium",
		2158: "Helius",
		1322: "Helix",
		8369: "Helixtech",
		59:   "Hellige",
		7488: "Hello",
		6785: "Helmholz",
		8559: "Helvetia",
		1546: "HemoCue",
		6423: "Hengstler",
		1963: "Heraeus",
		8598: "Heran",
		5590: "Hermes",
		595:  "Hermstedt",
		6191: "Hero",
		5567: "Heurikon",
		8523: "Hexatronic",
		9081: "Heyrex",
		9127: "Hi-P",
		1373: "Hi-Techniques",
		8543: "Hi-flying",
		1600: "HiConnect",
		5827: "HiQ",
		6948: "HiTEM",
		9057: "Hichan",
		4336: "Hidea",
		2836: "Hifn",
		8407: "Hifocus",
		375:  "High",
		3423: "HighPoint",
		9439: "Highgates",
		7273: "Hightech",
		1277: "Hikari",
		5782: "Hilan",
		2891: "Hill-Rom",
		4765: "Hills",
		3777: "Hillstone",
		410:  "Hilscher",
		8216: "Himax",
		4602: "Hinke",
		915:  "Hinox",
		7881: "Hioso",
		1619: "Hirata",
		6596: "Hirose",
		4510: "Hirsch",
		2171: "Hisharp",
		89:   "Hitachi",
		1830: "Hitech",
		4981: "Hitex",
		1340: "Hitpoint",
		870:  "Hitron",
		6866: "HiveMotion",
		6267: "Hivesystem",
		4648: "Hivision",
		3083: "Hoatech",
		1556: "Hochiki",
		4642: "Hokkaido",
		3279: "Hokkei",
		419:  "Hokubu",
		8565: "Holl",
		6138: "Holontech",
		8843: "Holoplot",
		1207: "Homag",
		2479: "HomeLogic",
		1359: "Homenet",
		8208: "Homerider",
		7248: "Homewins",
		2712: "Honda",
		8447: "Honest",
		5828: "Honewell",
		934:  "Honeywell",
		5599: "Honeywell-Dating",
		4359: "Honeywld",
		5166: "Hong",
		7796: "Hongyu",
		2296: "Horanet",
		5265: "Horizon",
		2941: "Horoquartz",
		315:  "Horoscas",
		3694: "Hose-McCann",
		1730: "Hosiden",
		562:  "Hospira",
		6144: "Host",
		2335: "Hostlink",
		1628: "Hostnet",
		2640: "HotLava",
		2829: "Hotway",
		7195: "Hoya",
		4277: "Hsing",
		8650: "Htel",
		6576: "Hu",
		1453: "HuMANDATA",
		4682: "HuRob",
		3334: "Huawei",
		7053: "Huayuan",
		4480: "Hub-Tech",
		7731: "Hubbell",
		9298: "Huike",
		7312: "Huiwan",
		7089: "Huiyang",
		949:  "Human",
		8505: "Humannix",
		7495: "Humanorporated",
		4801: "Humanware",
		4011: "Hunkeler",
		2150: "Hunt",
		472:  "Hunter",
		8023: "Huria",
		8251: "Husqvarna",
		8116: "Hutec",
		6749: "Hutek",
		2027: "Huwell",
		4272: "Hybrid",
		4665: "Hybus",
		6328: "Hydra",
		8711: "Hydro",
		1227: "Hyglo",
		4305: "Hylab",
		2767: "Hymatom",
		4228: "Hynet",
		1657: "HyperEdge",
		5176: "Hypercom",
		3680: "Hypermedia",
		3749: "Hyperstone",
		5097: "Hypertec",
		3971: "Hypertherm",
		5447: "Hytec",
		7600: "Hytera",
		1374: "HyunJu",
		1351: "Hyundai",
		7591: "Hyunjin.com",
		7842: "Hyunteck",
		8035: "I&C",
		6727: "I'M",
		5644: "I-BUS",
		5331: "I-Cube",
		5874: "I-Cubed",
		6684: "I-LAX",
		2414: "I-O",
		6636: "I-Storm",
		8571: "I-Sys",
		4804: "I-TEC",
		4775: "I-Tech",
		4383: "I-WIN",
		867:  "I/F-CoM",
		236:  "I2SE",
		7520: "I4VINE",
		8696: "IAI",
		2022: "IAV",
		6886: "IAdea",
		494:  "IBASE",
		7932: "IBC",
		372:  "IBM",
		4229: "IBR",
		9491: "IC",
		269:  "IC-Net",
		1184: "ICA",
		62:   "ICANN",
		5074: "ICC",
		42:   "ICE",
		600:  "ICG",
		1027: "ICHIPS",
		5062: "ICM",
		1249: "ICP",
		1997: "ICPDAS",
		8601: "ICR",
		763:  "ICS",
		575:  "ICUE",
		3494: "ICX",
		7862: "ID",
		9070: "IDFone",
		488:  "IDIS",
		1299: "IDK",
		558:  "IDOT",
		9058: "IDS",
		3225: "IDT",
		4073: "IDX",
		6656: "IDY",
		2963: "IEC",
		2628: "IEE",
		7231: "IEP",
		745:  "IER",
		9342: "IES",
		319:  "IFM",
		5248: "IFP",
		7664: "IGI",
		8980: "IGRS",
		7323: "IGShare",
		6085: "IGT",
		905:  "IGYS",
		4335: "IHSE",
		8835: "II-VI",
		8265: "IJ",
		8449: "IMAGO",
		6851: "IMAX",
		26:   "IMC",
		5279: "IMD",
		5145: "IMF",
		1387: "IMI",
		8854: "IMK",
		9283: "IMS",
		3177: "IMV",
		5480: "IN-NET",
		5818: "INAT",
		353:  "INIT",
		1579: "INITECH",
		1651: "INITIUM",
		4781: "INNERINT",
		3194: "INNOTZ",
		1017: "INNOVI",
		996:  "INNOWELL",
		3543: "INOCOVA",
		7667: "INQ",
		4392: "INRange",
		6218: "INSIDE",
		6489: "INTARSO",
		5794: "INTEGRATED",
		7618: "INVENTEC",
		2902: "INVISIO",
		5267: "IO",
		6951: "IO-Power",
		1170: "IOA",
		4345: "IOGEAR",
		390:  "IOI",
		4524: "IOLAN",
		4205: "ION",
		4475: "IONODES",
		2770: "IONOS",
		6767: "IOTTECH",
		7185: "IP-Line",
		6731: "IP-NET",
		402:  "IP.Access",
		778:  "IPAS",
		5945: "IPC",
		1583: "IPCserv",
		9462: "IPEVO",
		1463: "IPFLEX",
		1087: "IPFRONT",
		1283: "IPMobileNet",
		893:  "IPOptical",
		9013: "IPROAD",
		7082: "IPS",
		737:  "IPWireless",
		4601: "IPaXiom",
		8148: "IPitomy",
		831:  "IPmental",
		8122: "IPmotion",
		4600: "IPnect",
		679:  "IPrad",
		5561: "IQ",
		5197: "IQinVision",
		4890: "IRCA",
		4527: "IRD",
		2419: "IRIICHI",
		2618: "IRT",
		4425: "IRTrans",
		7595: "IRay",
		5743: "ISA",
		1717: "ISAC",
		6274: "ISB",
		3171: "ISCO",
		297:  "ISDN",
		4824: "ISGUS",
		3601: "ISL",
		3364: "ISONAS",
		2626: "ISR",
		8156: "ISSC",
		2666: "IT-Factory",
		6765: "IT-IS",
		7018: "ITC",
		1176: "ITDevices",
		8494: "ITE",
		3995: "ITEC",
		4517: "ITF",
		2009: "ITFOR",
		6471: "ITG",
		1991: "ITI",
		6792: "ITL",
		2297: "ITO",
		8706: "ITON",
		1026: "ITRAN",
		1890: "ITSupported",
		1178: "ITT",
		692:  "ITTC",
		8563: "ITTIM",
		3474: "ITU-T",
		4293: "ITV",
		4954: "IVA",
		8003: "IVS",
		2990: "IVT",
		2537: "IWICS",
		7531: "IXECLOUD",
		2446: "IZT",
		6079: "Ibond",
		2776: "Ibtek",
		7824: "Icarvisions",
		5668: "Icom",
		7114: "Icomera",
		760:  "Iconag",
		2156: "Icotera",
		3638: "Icron",
		5770: "Ictv",
		7517: "IdaTech",
		4799: "Idaho",
		3042: "Ideal",
		4494: "Idealbt",
		4675: "Idemia",
		3721: "Identec",
		1497: "Identicard",
		1137: "Identix",
		7090: "Iders",
		4957: "Idream",
		4732: "IfTA",
		7190: "Iffu",
		6745: "Iflytek",
		2659: "Ifotec",
		1752: "Iftest",
		7891: "Igneous",
		6812: "IgniteNet",
		2095: "Iiga",
		1414: "Ikanos",
		1272: "Ilevo",
		1122: "Ilinx",
		3415: "ImCoSys",
		7356: "ImTech",
		4350: "Imacs",
		824:  "ImageCom",
		495:  "Imagenics",
		5602: "Imagic",
		3508: "Imagination",
		5690: "Imagine",
		8888: "Imaqliq",
		554:  "Imation",
		2294: "Imatron",
		2378: "Imci",
		1482: "Imerge",
		4031: "ImesD",
		5535: "Imlogix",
		8524: "ImmediaTV",
		4979: "Impacct",
		5089: "Impact",
		3596: "Impatica",
		351:  "Imperial",
		8195: "Impex-Sat&Co",
		3062: "Impinj",
		6004: "Impresstek",
		3566: "Impro",
		5896: "Impulse",
		1711: "Imsys",
		4154: "In-Circuit",
		1813: "In-Tech",
		3165: "InGrid",
		4560: "InMage",
		3895: "InPhase",
		1202: "InPro",
		7766: "InShow",
		9001: "InView",
		8891: "InVue",
		6460: "Inala",
		4943: "Inalp",
		1035: "Inara",
		1739: "Inc.jetorporated",
		6209: "IncAA",
		2892: "IncNETWORKS",
		3617: "IncOTEC",
		4905: "IncOstartec",
		2026: "Incipient",
		6549: "Incipio",
		8044: "Incoax",
		8053: "Incognito",
		5980: "Incredible",
		9346: "IndUSNET",
		2147: "Indagon",
		6348: "Indata",
		649:  "Indel",
		9479: "Independent",
		7698: "Indieon",
		5035: "Indigita",
		4165: "Indigo",
		6829: "Individual",
		7337: "Indu-Sol",
		2582: "Inducon",
		6570: "Industrial",
		1795: "Indyme",
		1286: "IneoQuest",
		2696: "Inepro",
		7675: "Inesa",
		4266: "Inet",
		762:  "Inetcam",
		5935: "Inex",
		8970: "InfORM",
		4472: "InfORSON",
		3138: "InfRANET",
		644:  "InfiNet",
		8441: "Inficomm",
		5728: "Inficon",
		482:  "Infineon",
		1746: "Infinera",
		4832: "Infineta",
		954:  "InfiniCon",
		6571: "InfiniWing",
		2915: "Infinias",
		2132: "Infinico",
		7807: "Infinidat",
		251:  "Infinilink",
		746:  "Infiniswitch",
		752:  "Infinite",
		2066: "Infinity",
		6295: "Infinix",
		3339: "Infinova",
		1880: "InfoExpress",
		5822: "InfoGear",
		7637: "InfoSight",
		9373: "Infoblox",
		3488: "Infocrypt",
		6131: "Infocus",
		2739: "Infohand",
		5224: "Infolibria",
		3976: "Infomark",
		8974: "Infopia",
		6417: "Inform",
		4384: "Informatics",
		7:    "Information",
		8984: "Informtekhnika",
		5981: "Infortrend",
		1962: "Infotec",
		5808: "Infotek",
		92:   "Infotron",
		1964: "Infrant",
		8729: "Ingate",
		536:  "Ingenico",
		6137: "Ingenieria",
		348:  "Ingersoll-Rand",
		4043: "Ingespace",
		4592: "Inhand",
		7965: "Inhon",
		5441: "Inid",
		8491: "Inim",
		2298: "Initio",
		224:  "Inkel",
		623:  "Inkra",
		6944: "Inlab",
		6642: "Inmarsat",
		900:  "Inncom",
		4789: "InnerSpace",
		2864: "InnerWireless",
		3845: "Innes",
		8892: "InnoDigital",
		2320: "InnoLabs",
		2366: "InnoMedia",
		6112: "InnoMediaLogic",
		5309: "InnoSys",
		7342: "Innocent",
		675:  "Innocom",
		8457: "Innogrit",
		7192: "Innolight",
		8348: "Innolux",
		812:  "Innomedia",
		9409: "Innometriks",
		8306: "Innonet",
		9132: "Innophase",
		1384: "Innopia",
		7341: "InnosiliconTechnology",
		3435: "Innotech",
		8437: "Innotube",
		6206: "Innova",
		1064: "Innovance",
		5601: "Innovaphone",
		4694: "Innovar",
		6210: "Innovat",
		428:  "Innovative",
		2372: "Innovex",
		9129: "Innovid",
		6689: "Innovolt",
		5028: "Inova",
		4211: "Inovis",
		9115: "Inovonics",
		6916: "Inpeco",
		4790: "InpegVision",
		4369: "Inphi",
		5387: "Input/Output",
		1891: "Inqnet",
		7289: "Insensi",
		4885: "Insightek",
		7260: "Insigma",
		8365: "Inspire",
		8869: "Inspiremobile",
		7721: "Inspur",
		5791: "Instem",
		4872: "Instrumentation",
		6331: "Insystec",
		7109: "Intai",
		6865: "IntalTech",
		3946: "Intego",
		8880: "IntegraOptics",
		5873: "Integrated",
		8511: "Integri-Sys.com",
		5686: "Integrix",
		421:  "Intel",
		3541: "Intelbras",
		4825: "InteliCloud",
		3413: "Intelicis",
		7593: "Intelight",
		2486: "Intellambda",
		3706: "Intellect",
		9240: "IntelliVoice",
		6378: "Intellian",
		4984: "Intellibyte",
		5486: "Intellicom",
		1861: "Intelligent",
		6543: "Intellimedia",
		3483: "Intellio",
		7238: "Intellion",
		7999: "Intellisis",
		7486: "Intellithings",
		7576: "Intellivision",
		6013: "Intelliworxx",
		7336: "Intent",
		3868: "Inter-M",
		9459: "InterCreative",
		2493: "InterEnergy",
		1398: "InterEpoch",
		4080: "Interacoustics",
		2682: "Interactek",
		3613: "Interactivetv",
		1047: "Interalia",
		3515: "Interay",
		9211: "Intercept",
		9096: "Intercom",
		8899: "Intercon",
		3330: "Intercross",
		1410: "Interface",
		5272: "Intergon",
		6317: "Intergraph",
		3023: "Interlink",
		2317: "Intermec",
		5920: "Intermedium",
		639:  "International",
		8602: "Internet",
		5345: "Internix",
		77:   "Interphase",
		9471: "Interphone",
		5011: "Intersil",
		347:  "Intersoft",
		4861: "Interspiro",
		3967: "Intertain",
		8280: "Intertech",
		514:  "Intervoice-Brite",
		5120: "Interware",
		1215: "IntiGate",
		1034: "Intime",
		3928: "Intoto",
		5349: "IntraServer",
		4348: "Intraco",
		7415: "Intrakey",
		1067: "Intransa",
		919:  "Intraserver",
		9120: "Intrigue",
		927:  "Intrinsyc",
		9229: "Intrising",
		1512: "Intronicsorporated",
		54:   "Intrusion.com",
		983:  "Intruvert",
		3941: "Intuicom",
		6806: "Intune",
		1544: "Invacom",
		6873: "Invenit",
		4251: "Invensys",
		3978: "Inventec",
		8817: "Inventek",
		1075: "Inventel",
		2648: "Invento",
		7348: "Inventum",
		6037: "Invertex",
		936:  "Invicta",
		5073: "Invisible",
		1476: "Invivo",
		6637: "Invoxia",
		4541: "Inyuan",
		8548: "IoT",
		6059: "Iomega",
		2004: "Ionix",
		1673: "Iosoft",
		6743: "Iota",
		6036: "Iowave",
		704:  "Ipanema",
		1348: "Ipetronik",
		4800: "Iphion",
		2779: "Iprobe",
		1200: "Ipsilorporated",
		3980: "Ipte",
		6497: "Iqsim",
		3011: "Iqua",
		3554: "Ircona",
		5192: "Iris",
		6908: "Irlab",
		4992: "Irlan",
		3408: "Irlink",
		5099: "Ironicsorporated",
		6753: "Irts",
		4846: "Irumtek",
		3672: "IsaacLandKorea",
		1853: "Iscutum",
		4223: "Isdn",
		4222: "Isdyne",
		4935: "Ishida",
		146:  "Isicad",
		2937: "Isilon",
		1852: "Iskraemeco",
		5882: "Isolation",
		7150: "Ison",
		4554: "Isotek",
		9060: "Isung",
		1869: "Itcare",
		930:  "Itcn",
		6281: "Itel",
		2863: "Iteris",
		9051: "Itibia",
		5640: "Itis",
		6868: "Iton",
		1109: "Itron",
		8278: "Itronic",
		2161: "Itronix",
		8074: "Ivenix",
		5594: "Ivex",
		4993: "Ivron",
		6021: "Iwill",
		1107: "J",
		7913: "J-MEX",
		3113: "J-TEK",
		1501: "J-THREE",
		805:  "J-Works",
		5943: "J1",
		4246: "J125",
		1873: "JAI",
		468:  "JAMA",
		9092: "JANUS",
		9391: "JBL",
		6901: "JDA",
		7689: "JDC",
		1143: "JEAN",
		1578: "JEPICO",
		3346: "JESS-LINK",
		7589: "JETTER",
		9393: "JJ",
		2972: "JJPlus",
		2457: "JL",
		8432: "JLG",
		6449: "JM-DATA",
		1021: "JMI",
		965:  "JMP",
		8390: "JMR",
		3689: "JMicron",
		7929: "JNC",
		3389: "JOYTOTO",
		2021: "JPS",
		6612: "JRC",
		4826: "JRD",
		8851: "JRI",
		5741: "JRL",
		2504: "JStream",
		8905: "JTD",
		9304: "JTECH",
		6895: "JTI",
		8957: "JUNGJIN",
		7061: "JW",
		8199: "JWEntertainment",
		3545: "JWTrading",
		6599: "JY",
		7861: "Jabil",
		2758: "Jablotron",
		2522: "Jacobsons",
		5762: "Jacomo",
		7055: "Jadak",
		7492: "Jaewoncnc",
		1851: "Jamex",
		4506: "Janam",
		2055: "Janitza",
		8918: "Janteq",
		5683: "Janz",
		5453: "Japan",
		6336: "Jarogate",
		5869: "Jasco",
		1222: "Jascom",
		521:  "Jasmine",
		4068: "Jastec",
		6160: "Jato",
		5198: "Jatom",
		6124: "Jaton",
		3160: "Jaty",
		5666: "Javelin",
		8999: "JayBird",
		6147: "Jaycor",
		2814: "JazzMutant",
		6410: "Jebsee",
		7256: "Jeda",
		2992: "Jennic",
		660:  "Jenoptik",
		2907: "Jeongmin",
		4738: "Jeorich",
		5985: "Jetcell",
		8878: "Jetlun",
		8110: "Jetmobile",
		711:  "Jetstream",
		5282: "Jetter",
		8818: "Jfcontrol",
		8794: "JiQiDao",
		7975: "JianLing",
		7368: "Jibo",
		7138: "Jide",
		8141: "Jigowatts",
		6945: "JinQianMao",
		7483: "Jingsheng",
		7941: "Jinmuyu",
		8468: "Jiwumedia",
		3625: "Joby",
		3782: "JoeScan",
		7042: "Jolata",
		7333: "Jolla",
		3462: "Jorjin",
		1020: "Jotron",
		8248: "Joview",
		8527: "Jovision",
		3597: "Joybien",
		8202: "Joyent",
		1650: "Joymax",
		1996: "Joyteck",
		7095: "Jubix",
		7883: "Juice",
		7187: "Juin",
		4201: "Juki",
		5123: "Juko",
		8352: "Julong",
		4424: "JumpGen",
		9184: "Jumptronic",
		6375: "Junilab",
		826:  "Juniper",
		5452: "Jupiter",
		772:  "Jupiters",
		7829: "Jurumani",
		1719: "Jusan",
		863:  "JustEzy",
		6472: "Justec",
		9258: "Justone",
		5522: "Justsystem",
		8712: "Jwcnetworks",
		2906: "K",
		7990: "K's",
		2355: "K-BOT",
		5768: "K-NET",
		1935: "K-Patents",
		8960: "K2NET",
		2859: "K40",
		2466: "KAYABA",
		1001: "KB",
		7048: "KBC",
		3133: "KBS",
		3091: "KBT",
		999:  "KC",
		7854: "KCE",
		8243: "KCF",
		2534: "KCodes",
		188:  "KDC",
		3161: "KDE",
		769:  "KDT",
		5389: "KEBA",
		9341: "KEEBOX",
		6396: "KEMAS",
		8019: "KERNEL-I",
		7443: "KETEK",
		4012: "KIMIN",
		1433: "KINGENE",
		8425: "KLINFO",
		7426: "KMB",
		3983: "KMW",
		5904: "KNX",
		583:  "KOANKEISO",
		3111: "KORWIN",
		6769: "KSH",
		6364: "KSP",
		7290: "KST",
		5853: "KT",
		3026: "KT&C",
		4542: "KTC",
		3261: "KTF",
		151:  "KTI",
		7197: "KTIS",
		8853: "KUNBUS",
		4157: "KURUSUGAWA",
		5742: "KVB/Analect",
		8902: "KVH",
		6228: "KWB",
		5948: "KYE",
		4041: "KYLAND",
		3033: "KYLINK",
		186:  "KYOWA",
		1860: "Ka-Ro",
		4664: "Kaga",
		7869: "Kai-EE",
		6582: "Kaiam",
		2723: "Kaicom",
		2591: "Kaimei",
		3833: "Kaise",
		9425: "Kaishun",
		6419: "Kakao",
		1662: "Kaleidescape",
		4770: "Kalki",
		5466: "Kalpana",
		8786: "Kamama",
		3390: "Kameleon",
		4706: "Kaminario",
		6436: "Kamo",
		2787: "Kamstrup",
		1418: "Kanematsu",
		6532: "KangSheng",
		1269: "Kaonmedia",
		6047: "Kapadia",
		1682: "Kaparel",
		7078: "Kapelse",
		6148: "Kapsch",
		4388: "Kapsys",
		1259: "Karam",
		4233: "Kardios",
		1284: "Karel",
		2134: "Kasda",
		7926: "Kastle",
		4185: "Katana",
		6009: "Kathrein-Werke",
		9153: "Katoudenkikougyousyo",
		5180: "Katron",
		6057: "Katsujima",
		5414: "Kayser-Threde",
		4639: "Kaztek",
		4530: "Kedah",
		7374: "Keenetic",
		3128: "Kei",
		5799: "Keisokugiken",
		3240: "Kelman",
		9385: "Kelvin",
		9267: "Kemppi",
		7629: "Kenade",
		243:  "Kenetec",
		7642: "Kenstel",
		1602: "Kentec",
		2148: "Kentima",
		4218: "Kentrox",
		3736: "Kenwin",
		1920: "Kenwood",
		9065: "Keopsys",
		1688: "Kerajet",
		4941: "Kerbango",
		2823: "Keri",
		7776: "Kerlink",
		5244: "Kestrel",
		8435: "Ketra",
		3616: "Keumbee",
		5394: "KeunYoung",
		1189: "Key",
		2749: "KeyEye",
		2260: "KeyMed",
		6033: "Keycorp",
		317:  "Keyence",
		3599: "Keysight",
		3964: "Keytronix",
		9395: "Khomp",
		8790: "Khwahish",
		1710: "Kihoku",
		6961: "Kiio",
		2264: "Kikusui",
		3293: "Kimaldi",
		6657: "Kimball",
		5163: "Kimpsion",
		5546: "Kimtron",
		8910: "Kinestral",
		6354: "Kinetics",
		8049: "Kinexon",
		9045: "KingTing",
		2734: "Kinghold",
		3321: "Kingjim",
		5702: "Kingmax",
		7623: "Kingnetic",
		7814: "Kingsignal",
		5142: "Kingstar",
		3347: "Kingstate",
		4881: "Kingston",
		3038: "Kingtronics",
		1672: "Kingwave",
		7628: "Kinion",
		8647: "Kinova",
		299:  "Kinpo",
		8158: "Kioxia",
		1153: "Kirana",
		8211: "Kirisun",
		4726: "Kiryung",
		4514: "Kisan",
		3559: "KitWorks.fi",
		7422: "Kivic",
		9227: "Kivo",
		3215: "Kiyon",
		2791: "Klas",
		840:  "Kleinknecht",
		5077: "Klever",
		8282: "Kmdata",
		670:  "Knilink",
		2094: "Knovative",
		1635: "Knurr",
		7917: "KoamTac",
		2948: "Kocom",
		1955: "Koden",
		1671: "Kodeos",
		5463: "Kodiak",
		1820: "Kodicom",
		1204: "Koei",
		5216: "Koga",
		2856: "Kohler",
		2402: "Kollmorgen",
		334:  "Kollmorgen-Servotronix",
		5902: "Komatsu",
		2008: "Komax",
		183:  "Komodo",
		964:  "Konami",
		7435: "Koncar",
		2817: "Koncept",
		6566: "KongTop",
		3532: "Konka",
		8246: "KonnectONE",
		7633: "Konten",
		63:   "Kontron",
		4397: "Koratek",
		853:  "Korea",
		2605: "Korenix",
		2548: "Korg",
		6470: "Korins",
		8754: "Koss",
		846:  "Kott",
		8113: "Koubachi",
		5231: "Kouwell",
		7282: "Kove",
		802:  "Kowa",
		6031: "Koyo",
		4037: "Kozio",
		2954: "Kprotech",
		6670: "Kraftway",
		3891: "Kramer",
		407:  "Kreatel",
		4792: "Kroger",
		3885: "Krohne",
		3952: "Ksic",
		4819: "Ktnf",
		5908: "Kubota",
		5726: "Kubotek",
		9039: "Kuhn",
		6711: "Kuipers",
		720:  "Kumahira",
		8838: "Kumalift",
		6595: "Kummler+Matter",
		1147: "Kumoh",
		2733: "Kumyoung",
		7546: "KunTeng",
		992:  "Kuokoa",
		6643: "Kurth",
		6075: "Kuzumi",
		6000: "Kvaser",
		6525: "Kwangsung",
		7764: "Kwangwon",
		3994: "KwikByte",
		3735: "Kworld",
		7906: "Kyland-USA",
		2848: "Kyocera",
		3788: "Kyohritsu",
		7810: "Kyolis",
		1898: "Kyoto",
		7713: "Kyowa",
		2241: "Kyushu",
		2183: "Kyushu-kyohan",
		8014: "Kyynel",
		1419: "L&F",
		5892: "L&N",
		3157: "L-3",
		6638: "L-Tech",
		7818: "LAM",
		6311: "LAN-TEC",
		6198: "LANBit",
		5737: "LANCOM",
		1180: "LANergy",
		7619: "LARK",
		4918: "LARsys-Automation",
		3412: "LBNL",
		4919: "LCFC",
		1850: "LEA",
		776:  "LEA*D",
		7590: "LENUS",
		2912: "LET'S",
		1829: "LETEK",
		247:  "LEUNIG",
		869:  "LG",
		8599: "LG-Ericsson",
		7468: "LGE",
		6188: "LINK2IT",
		4839: "LIOS",
		452:  "LITE-ON",
		4309: "LInTech",
		6979: "LM",
		8069: "LMI",
		4511: "LNC",
		6486: "LNT-Automation",
		8885: "LOCOSYS",
		5312: "LOGWARE",
		4376: "LOHUIS",
		3684: "LORD",
		1563: "LOYTEC",
		1626: "LS",
		1116: "LSI",
		6161: "LTX-Credence",
		8911: "LVS",
		237:  "LXCO",
		5184: "LXE",
		9473: "LXinstruments",
		7768: "LaVision",
		8171: "LabJack",
		7421: "Labris",
		78:   "Labtam",
		1089: "Lafon",
		4006: "Lagotek",
		4536: "Laird",
		5636: "Lake",
		3220: "Laketune",
		7709: "Lampex",
		9277: "Lamprey",
		364:  "Lampus",
		4911: "LanPro",
		2503: "LanReady",
		5858: "Lanart",
		6051: "Lanbird",
		8576: "Lanbowan",
		5117: "Lanco",
		5390: "Land",
		2530: "Land-Cellular",
		7732: "Landauer",
		4248: "Landings",
		2227: "Landis+Gyr",
		5543: "Lanex",
		5582: "Lanner",
		5850: "Lanoptics",
		5867: "Lans",
		8368: "Lansen",
		7083: "Lansentechnology",
		6794: "Lantis",
		5528: "Lantronix",
		1416: "Lanvoice",
		5160: "Lanwan",
		6601: "Laon",
		6466: "LaonPeople",
		4244: "Lapis",
		5019: "Lara",
		8189: "Lars",
		4269: "Larscom",
		2306: "Lasat",
		4121: "Lascar",
		4515: "Lasercraft",
		5426: "Lasergraphics",
		5864: "Lasermaster",
		1479: "Lassen",
		4455: "Lastar",
		5843: "Latch",
		8958: "Latecoere",
		7973: "Latticework",
		5911: "Lauterbach",
		647:  "Lava",
		3137: "Law-Chain",
		1670: "Lawo",
		7115: "Layer3TV",
		6808: "Layon",
		2572: "LeWiz",
		9043: "Leadcore",
		804:  "Leader",
		1411: "Leadfly",
		4926: "Leadtek",
		3433: "Leaf",
		5859: "Leap",
		2589: "LeapComm",
		7980: "LeapFrog",
		9477: "Leapfive",
		3535: "Lear",
		403:  "Lectron",
		4631: "Lectrosonics",
		4126: "Ledco",
		9326: "LeddarTech",
		9328: "Ledvance",
		6729: "Lee-Dickens",
		9430: "Leeman",
		7369: "Leeo",
		6837: "Leetek",
		2141: "Legra",
		682:  "Legrand",
		5671: "Leichu",
		9195: "Leifheit",
		6222: "Leightronix",
		8534: "Leimac",
		2935: "Leipold+Co.GmbH",
		162:  "Leiser",
		9029: "Leitner",
		2665: "Lely",
		5114: "Lemcom",
		8181: "Lenbrook",
		2938: "Leneco",
		4559: "Lengda",
		3351: "Lenntek",
		4474: "Lenord",
		2664: "Lenovo",
		1383: "Lenten",
		1540: "Lenze",
		8482: "Leoni",
		8325: "Leonton",
		8349: "Lesira",
		3191: "LeucotronEquipamentos",
		2984: "Leuze",
		7775: "LevelOne",
		7230: "Levven",
		2928: "LexBox",
		7989: "Lexar",
		8520: "Lexking",
		613:  "Lexmark",
		3953: "LiComm",
		3309: "Lianhe",
		5529: "Liberty",
		8811: "Libratone",
		3051: "LibreStream",
		4818: "Licera",
		1233: "Liebert-Hiross",
		3817: "Liecthi",
		7126: "Life",
		6660: "LifeBEAM",
		9493: "LifeHealth",
		2797: "LifeSize",
		2898: "LifeSync",
		2198: "Lifetron",
		7009: "Light",
		220:  "LightChip",
		881:  "LightSand",
		4721: "Lightcomm",
		6081: "Lightera",
		5003: "Lightner",
		744:  "Lightpointe",
		8557: "Lightspeed",
		4946: "Lightwave",
		7802: "LightwaveRF",
		3422: "LigoWave",
		9168: "Lilee",
		3835: "Limetek",
		3235: "LinTech",
		6927: "Linctronix",
		6701: "Linea",
		9327: "Linear",
		8111: "Lineeye",
		4438: "LinkSprite",
		8949: "Linkcom",
		4532: "Linkflex",
		7610: "Linkflow",
		7355: "Linkintec",
		1783: "Linksys",
		7682: "Linktel",
		8951: "Linktop",
		6023: "Linkup",
		3812: "Linkwise",
		4816: "Linn",
		4454: "LinoWave",
		83:   "Linotype-Hell",
		4259: "Linq",
		9158: "Lintes",
		1004: "Linxtek",
		1278: "Liontech",
		7592: "Liquidtool",
		8460: "Lisantech",
		1669: "Litchfield",
		5788: "Lite-ON",
		6402: "Lite-On",
		7708: "LiteON",
		2674: "LiteTouch",
		8640: "Litemax",
		2873: "Liteon",
		1300: "Littlefeet",
		5494: "Litton",
		5339: "Litton/Poly-Scientific",
		4377: "LiveTV",
		7875: "LiveU",
		6519: "Liverock",
		9090: "Livescribe",
		7751: "Livesecu",
		5852: "Livingston",
		3090: "Liyuh",
		3227: "LoBenn",
		4737: "LoJack",
		1668: "Load",
		2280: "Lobos",
		5829: "LocSoft",
		828:  "Locusorporated",
		4109: "Lodam",
		1219: "LodgeNet",
		9147: "Loenk",
		4587: "LogMeIn",
		6562: "Logi-D",
		4200: "LogiCan",
		1168: "LogiSync",
		6006: "Logibag",
		5078: "Logic",
		9457: "Logic3",
		2020: "LogicaCMG",
		253:  "Logical",
		9091: "Logicom",
		51:   "Logicraft",
		8588: "Logipix",
		4422: "Logiplus",
		6674: "Logistic",
		7141: "Logital",
		242:  "Logitec",
		4079: "Logitech",
		4718: "Logitek",
		4845: "Logiways",
		6924: "Logosol",
		597:  "Logostek",
		6535: "LongSung",
		4578: "Longcheer",
		6418: "Longconn",
		7220: "Longicorn",
		3877: "Longkay",
		4547: "Lookit",
		3619: "Loopcomm",
		6479: "Looxcie",
		5430: "Loran",
		9114: "Lorent",
		4098: "Lorex",
		3700: "Lorica",
		9466: "Lorom",
		2287: "Loud",
		7760: "Low",
		2102: "Lowrance",
		7330: "Loxone",
		6129: "Lsics",
		4114: "Lucas",
		827:  "Lucent",
		5807: "Lucidata",
		3814: "Lucky",
		4633: "Ludl",
		7454: "Lufkin",
		4432: "Lumasense",
		1729: "Lumenera",
		8866: "Lumenpulse",
		8930: "Lumewave",
		3970: "Lumexis",
		8617: "Lumigon",
		8063: "Luminar",
		6807: "Luminator",
		3152: "Lundinova",
		2279: "Lutron",
		5022: "LuxN",
		5523: "Luxcom",
		4751: "Luxtera",
		8406: "Luxul",
		8026: "Lvswitches",
		1244: "Lyan",
		1681: "Lycium",
		6151: "Lynk",
		8305: "Lynxspringl",
		8135: "Lytro",
		6858: "Lytx",
		4579: "Lyyn",
		8012: "LzLabs",
		1794: "M&S",
		1934: "M-Audio",
		9416: "M-Cube",
		2368: "M-System",
		435:  "M.C.C.I",
		5772: "M.I",
		7767: "M2Communication",
		2939: "M2I",
		4869: "M2Mnet",
		5817: "M2Motive",
		1793: "MACKIE",
		6322: "MACNICA",
		2989: "MAKUS",
		3263: "MARA",
		230:  "MARGI",
		1513: "MARKEM",
		1031: "MAT",
		553:  "MAVIX",
		7410: "MAX-Tech",
		3450: "MAZeT",
		948:  "MBM",
		3844: "MBS",
		2278: "MBTech",
		8968: "MC",
		6940: "MCD",
		1420: "MCM",
		9246: "MCNEX",
		8418: "MCOT",
		6682: "MCT",
		2662: "MDK",
		2377: "MEDIA4",
		8824: "MEDIAEDGE",
		7919: "MEG",
		7030: "MEIZU",
		4007: "MEL",
		8373: "MELPER",
		8162: "MESADA",
		7945: "MHL",
		4838: "MIA",
		4996: "MICRILOR",
		8561: "MICRODIA",
		7828: "MICRODIGTAL",
		5383: "MICROSENS",
		6133: "MICROWI",
		1946: "MIDAS",
		5676: "MIHARU",
		6169: "MII",
		9361: "MIMO",
		8362: "MINIX",
		4917: "MIPRO",
		1687: "MITEQ",
		4904: "MKD",
		601:  "MKNet",
		4638: "MMB",
		4932: "MMC",
		9354: "MMPC",
		9253: "MOCACARE",
		8625: "MODI",
		2508: "MOIMSTONE",
		7874: "MOKO",
		7789: "MONAD",
		1316: "MONDIAL",
		7745: "MPB",
		8513: "MPI",
		5400: "MPL",
		9196: "MR&D",
		5199: "MRG",
		9121: "MRS",
		2253: "MRV",
		7956: "MS-Magnet",
		3054: "MSI",
		6569: "MST",
		4864: "MTA",
		7811: "MTG",
		2772: "MTI",
		8066: "MTMCommunications",
		6746: "MTN",
		577:  "MTS",
		2521: "MTT",
		104:  "MTX",
		8962: "MVTECH",
		917:  "MWE",
		1577: "MWS",
		8868: "MXCHIP",
		7169: "MXT",
		6780: "MYK",
		2101: "MYNAH",
		1338: "MYTECS",
		2076: "Maas",
		4951: "Maatel",
		3775: "Macab",
		9099: "Macandc",
		2137: "Macey",
		3326: "Mackware",
		5153: "Macq",
		5397: "Macraigor",
		5839: "MacroSAN",
		1012: "Macrolink",
		5021: "Macromate",
		6930: "Macrotech",
		564:  "Macrotek",
		5326: "Macrovision",
		5969: "Mactell",
		70:   "Madge",
		8545: "MadgeTech",
		7263: "Maestronic",
		4660: "Mag",
		252:  "Mag-Tek",
		3948: "Magellan",
		9405: "Magenta",
		4947: "MagicRam",
		3063: "Magicard",
		7699: "Magicjack",
		153:  "Magna",
		3999: "Magna-Power",
		6793: "MagnaCom",
		8768: "MagneMotion",
		725:  "Magnipix",
		701:  "Mahi",
		6917: "Maike",
		523:  "Mainnet",
		4584: "Mainpine",
		8920: "Maipu",
		7072: "MakerBot",
		6899: "Mako",
		4762: "Maksat",
		1949: "Makus",
		945:  "Malachite",
		6464: "Malgn",
		713:  "Malibu",
		1426: "Mamiya-OP",
		7239: "ManTechnology",
		2146: "Mangrove",
		1530: "Manticom",
		350:  "Mantra",
		6279: "Manycolors",
		8890: "Manzanita",
		1493: "Maple",
		657:  "Mapletree",
		4571: "Mapower",
		9265: "Mapper.ai",
		6967: "Maquet",
		1699: "Maranti",
		5926: "Marben",
		2387: "March",
		31:   "Marconi",
		5217: "Mariner",
		5648: "Markem-Imaje",
		7246: "Marketaxess",
		6447: "Marketech",
		9244: "Markov",
		5169: "Marner",
		8774: "MarqMetrix",
		5739: "Marquip",
		6696: "Marshal",
		4574: "Marshall",
		2386: "Martinho-Davis",
		3943: "Marusys",
		6362: "Masimo",
		1166: "Massana",
		6342: "Masscomp",
		4319: "Masterclock",
		4607: "Masternaut",
		1423: "MathStar",
		3070: "Mathtech",
		8400: "Matis",
		2473: "Matisse",
		6343: "Matra",
		1885: "Matrics",
		1:    "Matrix",
		4292: "Matrox",
		9365: "Matsufu",
		148:  "Matsushita",
		4464: "Mattel",
		712:  "Mavenir",
		3878: "Maverick",
		1448: "Mavin",
		6559: "Max",
		9324: "MaxID",
		8787: "MaxMedia",
		2748: "MaxStream",
		9028: "MaxTronic",
		4912: "MaxVision",
		914:  "Maxan",
		2431: "Maxanna",
		2221: "Maxcess",
		8693: "Maxeler",
		3496: "Maxfor",
		4566: "Maxian",
		7157: "Maxio",
		5298: "Maxlinear",
		1555: "Maxlink",
		5534: "Maxpeed",
		7293: "Maxphotonics",
		5122: "Maxton",
		2385: "Maxtor",
		6522: "Maxway",
		1036: "Maxxan",
		4778: "Maya-Creation",
		5214: "Mayan",
		1321: "Mayekawa",
		8521: "Maytronics",
		1803: "Mbari",
		2082: "McAfee",
		6469: "McCain",
		9465: "McPay",
		2120: "Mcharge",
		8823: "Mciao",
		5504: "Mcnc",
		6114: "Mcns",
		9449: "Mcot",
		6074: "Mcquay",
		7359: "Measy",
		1083: "Mecalc",
		8833: "Mechatro",
		3647: "Med-Eng",
		9078: "MedHab",
		555:  "Medea",
		3852: "Media",
		2451: "MediaCell",
		1781: "MediaChorus",
		1760: "MediaQ",
		4827: "MediaSputnik",
		1876: "MediaTek",
		6552: "Mediabridge",
		5219: "Mediafire",
		2512: "Medialink-i",
		688:  "Medialogic",
		5200: "Mediastar",
		1466: "Mediatek",
		8763: "Medicaroid",
		3614: "Medicis",
		1432: "Medicore",
		1849: "Medion",
		6936: "Mediplan",
		855:  "Medison",
		661:  "Medrad",
		1897: "Mega-Trend",
		5386: "MegaChips",
		6505: "Megabyte",
		132:  "Megadata",
		6710: "Megafone",
		88:   "Megahertz",
		2166: "Megapower",
		2140: "Megasolution",
		6310: "Megatek",
		1792: "Megatel",
		4187: "Megatron",
		8328: "Megatronix",
		4871: "Megger",
		5129: "Meidensha",
		157:  "Meiko",
		4355: "Mekics",
		7088: "Melange",
		904:  "Melco",
		8996: "Meld",
		9215: "Melec",
		6115: "Melita",
		432:  "Mellanox",
		8046: "Memjet",
		748:  "Memobox",
		1833: "MemoryLink",
		4242: "Memotec",
		3634: "Mendocino",
		5223: "Menicx",
		1821: "Mentor",
		7545: "Merchandising",
		9401: "Mercku",
		4253: "Mercury",
		6227: "Mercusys",
		6926: "Merging",
		2380: "Meridian",
		786:  "Merilus",
		6450: "Meritec",
		7023: "Meritech",
		408:  "Merix",
		480:  "Merlin",
		5246: "Merlot",
		6146: "Merrimac",
		3537: "Merten&CoKG",
		5319: "Mesa",
		923:  "Mesco",
		3226: "Meshcom",
		7947: "Mesmo",
		6611: "Messoa",
		8620: "Meta-Networks",
		2688: "MetaSwitch",
		3443: "Metacom",
		5958: "Metacomp",
		1390: "Metalink",
		1709: "Metalligence",
		7925: "Metamako",
		2088: "Metanoia",
		6312: "Metaphor",
		3337: "Metasystem",
		875:  "Metavector",
		2353: "Metawave",
		4661: "Meteocontrol",
		2942: "Meteor",
		8038: "Meter",
		538:  "Metera",
		3163: "Methode",
		6988: "Metis",
		8061: "Metrascale",
		4261: "Metricom",
		1839: "Metro",
		325:  "Metro-Optix",
		5909: "Metrodata",
		2559: "Metrohm",
		5102: "Metronix",
		7794: "Metrum",
		7128: "Mettle",
		2327: "Mettler-Toledo",
		4858: "Metz-Werke",
		7037: "Mevo",
		1549: "Mewtel",
		6781: "Mexus",
		4538: "MiXTelematics",
		9485: "Miartech",
		1649: "Micetek",
		7833: "Micobiomed",
		5156: "Micom",
		3960: "Micran",
		43:   "Micro",
		117:  "Micro-Matic",
		1759: "Micro-Optronic-Messtechnik",
		1812: "Micro-Star",
		5373: "Micro/Sys",
		5177: "MicroBrain",
		7226: "MicroPower",
		7542: "MicroSys",
		1764: "MicroWeb",
		5182: "Microboards",
		706:  "Microchip",
		97:   "Microcom",
		1911: "Microelectronics",
		6040: "Microfirst",
		6316: "Microfive",
		137:  "Micrognosis",
		2226: "Microhard",
		254:  "Microlink",
		8274: "Micromedia",
		4363: "Micromint",
		4627: "Micron",
		2456: "Micronas",
		2455: "Micronet",
		5545: "Micronics",
		326:  "Micronpc.com",
		5502: "Microplex",
		5706: "Micropolis",
		149:  "Microprocess",
		5736: "Micropross",
		4873: "Microrobot",
		156:  "Microsage",
		1642: "Microscan",
		5757: "Microsemi",
		6501: "Microseven",
		6172: "Microslate",
		612:  "Microsoft",
		4241: "Microtech",
		1408: "Microtechno",
		5517: "Microtek",
		5527: "Microtest",
		6836: "Microtime",
		2052: "Microtrol",
		4599: "Microtronic",
		994:  "Microtune",
		5143: "Microunity",
		2204: "Midas",
		6546: "Midicom",
		2492: "Midmark",
		8542: "Midokura",
		8826: "Midori",
		6176: "Midsco",
		765:  "Midstream",
		3676: "Midtronics",
		9472: "MikroBits",
		3440: "MikroM",
		6156: "Mikrodidakt",
		4234: "Mikron",
		9151: "MilDef",
		5158: "Milan",
		6695: "MileSight",
		8889: "Miljovakt",
		4331: "Millinet",
		3197: "Millipore",
		6869: "Milper",
		5975: "Mimaki",
		7783: "Mimodisplaykorea",
		6733: "Mimosa",
		2155: "Mindray",
		5033: "Mindready",
		1928: "Minds",
		1239: "Minds@Work",
		5284: "Minerva",
		636:  "Minet",
		8226: "Mini-Cam",
		8944: "Mini-Circuits",
		7407: "Minibar",
		8822: "Minieum",
		5087: "Minim",
		4730: "Minimax",
		73:   "Miniware",
		9181: "Minrray",
		8629: "Minsung",
		1838: "Mintera",
		2087: "Mintron",
		8961: "Mios",
		8062: "Miovision",
		1226: "Mipsys",
		6784: "MirWifi",
		3679: "MiraLink",
		498:  "Mirae",
		9274: "MiraeRecognition",
		7695: "MiraeSignal",
		520:  "Miraesys",
		8695: "Mirka",
		4217: "Miro",
		6521: "Miromico",
		7484: "Mist",
		1933: "Mita-Teknik",
		513:  "Mitac",
		1289: "Mitadenshi",
		5884: "Mitec",
		1217: "Mitel",
		6428: "MitraStar",
		3884: "Mitron",
		7717: "Mitsuba",
		2393: "Mitsubishi",
		6673: "Mitsunami",
		8780: "Mitsuya",
		3285: "Mitutoyo",
		3717: "Miura",
		4707: "Miyoshi",
		5566: "Mizar",
		4402: "MoCA",
		1741: "MoTEX",
		1055: "Mobiis",
		6983: "MobilMAX",
		4687: "Mobilarm",
		6380: "Mobile",
		4850: "MobileAccess",
		2608: "MobileAria",
		3863: "MobileCompia",
		7110: "Mobileeco",
		3848: "Mobilesoft",
		7952: "Mobilicom",
		985:  "Mobillian",
		3557: "Mobinnova",
		3714: "Mobisolution",
		9338: "Mobitec",
		3500: "Mobitek",
		5128: "Mobius",
		668:  "Mobiwave",
		574:  "Mobotix",
		4975: "MobyTEL",
		7098: "Moda-InnoChips",
		3913: "Modacom",
		7739: "Modcam",
		8024: "Modelleisenbahn",
		3673: "Modnnet",
		9117: "Modoosis",
		6934: "Moduel",
		6648: "Moduletek",
		3626: "ModusLink",
		8952: "Moen",
		9172: "MofiNetwork",
		7035: "Mohlenhoff",
		8489: "Moimstone",
		4092: "Mojix",
		2485: "Mojo",
		6504: "Molekule",
		579:  "Momentum",
		6557: "Monaco",
		4736: "Moneytech",
		2498: "Monitoring",
		6088: "Monterey",
		3341: "Montgomery",
		2613: "Montilio",
		1439: "Moog",
		3881: "Moohadigital",
		851:  "Moore",
		7743: "Mooredoll",
		9191: "Mophie",
		2844: "Moram",
		4407: "Morega",
		3377: "Morey",
		6476: "Morion",
		4748: "Morningstar",
		5832: "Morrow",
		886:  "Mosaic",
		3114: "Moser-Baer",
		1104: "Motah",
		3274: "Motech",
		687:  "Motorola",
		7823: "Movek",
		1435: "Movistec",
		399:  "Movita",
		2722: "Movon",
		5682: "Moxa",
		4744: "Mpedia",
		8131: "Mpgio",
		7522: "Mplus",
		6579: "Mpmkvvcl",
		8628: "MtM",
		4645: "MuLogic",
		8879: "Muehlbauer",
		6430: "Mueller",
		4061: "Mueller-Elektronik",
		8554: "Muller",
		2861: "MultiCom",
		5018: "Multidata",
		2865: "Multilink",
		4435: "Multimedia",
		1715: "Multiplex",
		4255: "Multipoint",
		1183: "Multitech",
		6018: "Multitel",
		1208: "Multitone",
		2056: "Murata",
		2233: "Murrelektronik",
		1121: "Musashi",
		3882: "Muscle",
		3830: "MusicianLink",
		7887: "Musilab",
		5703: "Mutoh",
		8492: "Muuselabs/SA",
		582:  "Muxcom",
		3615: "Mvox",
		2019: "MyA",
		7010: "MyCloud",
		9448: "MyLight",
		5887: "Myco",
		2469: "Mykotronx",
		5413: "Myricom",
		5930: "Myson",
		861:  "Myspace",
		5600: "Mysticom",
		4050: "Mytek",
		3698: "MyungMin",
		9475: "N-Radio",
		7146: "N-iTUS",
		5094: "N.A.T",
		2945: "N3",
		3872: "NAC-Intercom",
		947:  "NADEX",
		6082: "NAKAYO",
		5270: "NALTEC",
		7294: "NANOWAVE",
		6654: "NARI",
		7860: "NASCENT",
		8815: "NAVIS",
		8613: "NB",
		6308: "NBI",
		756:  "NBS",
		6184: "NBX",
		5572: "NC&C",
		6302: "NCR",
		8351: "NCTech",
		2276: "NComputing",
		9445: "NDC",
		1919: "NDR",
		2007: "NDS",
		48:   "NEC",
		8913: "NECMagnusCommunications",
		5628: "NEO",
		7961: "NEON",
		1985: "NEOSMART",
		4669: "NES",
		5161: "NET-Source",
		8057: "NET2GRID",
		5316: "NET2NET",
		3951: "NETCLEUS",
		3430: "NETCONS",
		9207: "NETINT",
		7742: "NEXCONTROL",
		4768: "NEXTEK",
		1097: "NEXTEYE",
		2901: "NF",
		9308: "NGD",
		5802: "NHC",
		8732: "NHN",
		444:  "NICE",
		6754: "NIIC",
		4144: "NIW",
		1834: "NKE",
		2813: "NL",
		8318: "NMR",
		4189: "NMS",
		4867: "NOOLIX",
		1229: "NORTHDATA",
		5638: "NOT",
		8254: "NOVA",
		3344: "NPCore",
		8591: "NRG",
		5291: "NS",
		9220: "NSD",
		7096: "NSE",
		3902: "NSGate",
		2406: "NSI",
		817:  "NSM",
		4177: "NSSLGlobal",
		2259: "NST",
		4378: "NTC-Metrotek",
		7936: "NTI",
		6989: "NTT",
		2704: "NTTPC",
		7134: "NTmore",
		4617: "NUETEQ",
		4063: "NUM",
		2757: "NUMA",
		7967: "NUUO",
		7245: "NVIDIA",
		7353: "NX",
		526:  "NXTV",
		9440: "NZXT",
		9072: "Nabtesco",
		8736: "Nabto",
		652:  "Nact",
		2949: "Nadam",
		8321: "Nadasnv",
		381:  "Nadatel",
		9264: "Nagravision",
		8147: "Nagtech",
		9427: "Nain",
		1868: "Nallatech",
		8423: "NamJunSa",
		8123: "NambooSolution",
		5598: "Namco",
		7996: "Nanoleaf",
		4219: "Nanomatic",
		8168: "Nanomegas",
		2458: "Nanometrics",
		7744: "Nanoptix",
		7208: "Nanotec",
		4558: "Nanoteq",
		6588: "Nanotron",
		9053: "Nantworks",
		3859: "Napera",
		906:  "Narad",
		2710: "Narayon",
		2001: "Nasaco",
		5708: "Nascent",
		5806: "Nashoba",
		7778: "Nastec",
		9137: "Nata-Info",
		9071: "National",
		6915: "Nations",
		7886: "Nationz",
		324:  "Native",
		4709: "NaturalPoint",
		7692: "Nautronix",
		3349: "Navcast",
		1161: "Navcom",
		7417: "Navdy",
		8198: "Naver",
		5203: "Navic",
		3671: "Navigon",
		7147: "Naviit",
		676:  "Navini",
		3934: "Navionics",
		6972: "Navtech",
		5775: "Navtel",
		695:  "Nayna",
		2698: "Naztec",
		5689: "Nbase",
		6760: "Ncse",
		5159: "Ncube",
		4406: "Neat",
		8810: "Neatframe",
		6631: "Nebra",
		8871: "Nebula",
		9376: "Nebulon",
		8070: "Nebusens",
		341:  "Necsom",
		6598: "Nectarsoft",
		4268: "Nectec",
		8007: "Neets",
		4757: "Neli",
		4754: "Nemo-Q",
		6547: "Nemoa",
		1598: "NeoAxiom",
		2542: "NeoMedia",
		278:  "NeoWave",
		3190: "Neology",
		9281: "Neomontana",
		5043: "Neon",
		3322: "Neonode",
		6162: "Neoparadigm",
		9420: "Neopis",
		3630: "Neopost",
		4216: "Neoproducts",
		1445: "Neoscale",
		7077: "Neosfar",
		4630: "Neostar",
		7992: "Neosystem",
		6677: "Neotech",
		4557: "Neotion",
		7898: "Neousys",
		4823: "Neovia",
		7026: "Nephos",
		5398: "Nera",
		7787: "Nero",
		8539: "Ness",
		4899: "Nesslab",
		6635: "Nest",
		1480: "Nestar",
		8653: "Nesys",
		318:  "Net",
		1904: "Net2Edge",
		5857: "NetAlly",
		5558: "NetApp",
		5673: "NetBoost",
		442:  "NetBotz",
		604:  "NetBurner",
		3555: "NetCare",
		549:  "NetChip",
		2588: "NetEffect",
		707:  "NetEnabled",
		1113: "NetEngines",
		2525: "NetEnrich",
		1271: "NetExcell",
		5432: "NetICs",
		568:  "NetKit",
		2081: "NetKlass",
		240:  "NetLinks",
		7159: "NetMan",
		529:  "NetMedia",
		2445: "NetModule",
		728:  "NetMount",
		973:  "NetNearU",
		2475: "NetStreams",
		4623: "NetUP",
		8215: "NetView",
		5838: "NetWorth",
		1141: "NetZerver",
		2013: "Netac",
		5258: "Netaccess",
		1431: "Netas",
		7798: "Netatmo",
		1248: "Netbind",
		5215: "Netcam",
		2033: "Netcodec",
		830:  "Netcom",
		5080: "Netcomm",
		1415: "Netcontrol",
		5179: "Netcor",
		2362: "Netcore",
		5700: "Netcorp",
		7170: "Neterix",
		5863: "Netexpress",
		960:  "Netezza",
		2563: "Netfabric",
		3947: "Netflix",
		2482: "Netforyou",
		1368: "Netgear",
		640:  "Netgem",
		9054: "Netgenetech",
		572:  "Netility",
		3734: "Netio",
		5007: "Netiverse",
		2367: "Netline",
		8450: "Netlist",
		6634: "Netlogic",
		8071: "Netmoon",
		9237: "Netonix",
		638:  "Netous",
		1117: "Netpower",
		4186: "Netquest",
		5856: "Netrix",
		5336: "Netro",
		7995: "Netronics",
		9:    "Netronix",
		2962: "Netronome",
		785:  "Nets",
		2315: "Netschools",
		5515: "Netscout",
		448:  "Netsec",
		267:  "Netsensity",
		5855: "Netspan",
		170:  "Netspect",
		6193: "Netspeed",
		2170: "Netstar",
		6856: "Netstor",
		2709: "Neturity",
		5934: "Netvantage",
		5401: "Netvision",
		2721: "Netvox",
		5871: "Netwiz",
		109:  "Network",
		9052: "NetworkAccountant",
		5478: "Networld",
		79:   "Networth",
		6226: "Netzin",
		3185: "NeuLion",
		8870: "Neul",
		1875: "NeuroCom",
		4203: "Neuron",
		3284: "Neuros",
		8000: "Neurotek",
		6232: "Nevatec",
		2983: "Nevis",
		7027: "Nevro",
		8886: "NewSharp",
		2982: "NewSoft",
		7991: "Newag",
		3789: "Newbury",
		2091: "Newcotech",
		5249: "Newer",
		5136: "Newgen",
		7756: "Newings",
		1352: "Newisys",
		1388: "Newport",
		8039: "Newrun",
		925:  "Newtec",
		1659: "Newtech",
		7274: "Newtek",
		8260: "NexFi",
		4071: "NexG",
		2567: "NexQL",
		8836: "Nexar",
		2416: "Nexcom",
		6038: "Nexcomm",
		8422: "Nexell",
		5188: "Nexo",
		6980: "Nexpring",
		614:  "Nexsan",
		615:  "Nexsi",
		5:    "Next",
		3866: "NextGTV",
		1301: "NextGig",
		4132: "NextIO",
		6820: "NextNav",
		5834: "Nextcell",
		255:  "Nextcomm",
		9337: "Nextek",
		3318: "Nexterm",
		5428: "Nextest",
		1859: "Nextlink",
		6080: "Nextone",
		6959: "Nextwill",
		2240: "Nexus",
		6201: "Nexware",
		8755: "Nexxt",
		6676: "Nhnetworks",
		5944: "Nice",
		1693: "NiceTechVision",
		2197: "Nicety",
		3021: "Nicevt",
		4521: "Nicira",
		1745: "Nidek",
		1060: "Nielsen",
		7167: "NietZsche",
		3468: "Nifty",
		1242: "Nihon",
		8333: "Nike",
		8909: "Nikkiso",
		2444: "Niko",
		6513: "Niko-Servodan",
		5661: "Nikon",
		8507: "Nilan",
		8191: "Nilfisk",
		7986: "Nimbus",
		5654: "NineTiles",
		1918: "Ninelanes",
		1427: "Nintendo",
		9022: "Nippon",
		2154: "Nisca",
		5575: "Nishimu",
		7790: "Nishiyama",
		4851: "Nissho-denki",
		1910: "Nissin",
		1744: "Nitgen",
		2964: "Nits",
		3052: "Nittan",
		4901: "Nivetti",
		4503: "Nivis",
		3064: "Nivus",
		6225: "Nixdorf",
		1751: "Nixvue",
		479:  "Nobell",
		8830: "Noblex",
		6555: "Noccela",
		7227: "Nocsys",
		457:  "Nokia",
		2879: "Nokota",
		2980: "Nolan",
		5293: "Nomadix",
		6249: "Nome",
		4446: "Nomus",
		7734: "Noon",
		5929: "Norand",
		4468: "Norcott",
		7267: "Noregon",
		1507: "Noritz",
		7901: "Norma",
		6485: "Norphonic",
		7942: "Nortec",
		2918: "Nortech",
		8132: "Nortek-AS",
		76:   "Nortel",
		2182: "Northover",
		783:  "Northstar",
		6884: "Nothing",
		4339: "NovAtel",
		8841: "Nova",
		7497: "NovaComm",
		1603: "NovaPal",
		9260: "NovaSparks",
		3791: "Novacomm",
		7268: "Novakon",
		5365: "Novalink",
		7715: "Novar",
		1128: "Novasonics",
		1263: "Novatec",
		1214: "Novatechnology",
		3255: "Novatron",
		9245: "NovelSat",
		14:   "Novell",
		7985: "Noviga",
		4467: "Novita",
		2650: "Novomatic",
		962:  "Novra",
		6049: "Novtek",
		5521: "Novus",
		8292: "Noxus",
		8271: "Nsolution",
		8564: "NuLEDs",
		4802: "NuVo",
		770:  "Nuark",
		9111: "Nubia",
		5776: "Nucom",
		4415: "Nucomm",
		3752: "Nucsafe",
		1391: "Nudian",
		2348: "Nuera",
		9023: "Nuheara",
		3385: "Numata",
		2943: "Numatics",
		8663: "Numera",
		2041: "Numesa",
		9156: "Nureva",
		7843: "Nusoft",
		7340: "Nutanix",
		7487: "Nuvico",
		7172: "Nuvolt",
		7599: "Nuvon",
		5551: "Nuvotech",
		7345: "Nuvyyo",
		1724: "Nvergence",
		659:  "Nvidia",
		9314: "NxtConect",
		1990: "O'Rite",
		6958: "O-Net",
		5795: "O.N",
		2617: "O2Micro",
		5954: "OAK",
		8982: "OCOSMOS",
		4720: "OCP",
		7832: "OCT",
		1770: "OCTTEL",
		8541: "ODA",
		4533: "OJ-Electronics",
		5006: "OL'E",
		6735: "OMICRON",
		4798: "OMNI-WiFi",
		3350: "OMNIKEY",
		7100: "OMRON",
		3425: "ON",
		1071: "ONStor",
		7004: "ONvocal",
		7014: "OOSIC",
		2299: "OOmon",
		178:  "OPEN",
		9234: "OPMEX",
		1570: "OPNET",
		2599: "OPTOELECTRONICS",
		7680: "OPWILL",
		1828: "OQO",
		7970: "ORA",
		1763: "ORACOM",
		9291: "ORICO",
		5825: "ORSYS",
		9140: "ORTHOsoft",
		7948: "OSRAM",
		7979: "OT",
		2950: "OTE",
		8308: "OTECTechnology",
		8357: "OTSL",
		6846: "OVH",
		7044: "OWIN",
		8827: "OWLink",
		501:  "Oak",
		8424: "Oakley",
		6386: "Obelux",
		8336: "Obihai",
		2308: "Objective",
		7946: "Objenious",
		7365: "Observint",
		3222: "Obsidian",
		4039: "Obvius",
		2199: "Obzerv",
		395:  "Occam",
		6283: "Occuspace",
		8679: "OceanServer",
		8903: "Ocom",
		4175: "Octagon",
		1822: "Octasic",
		1489: "Octave",
		4291: "Octel",
		8320: "Octonion",
		7081: "Octopod",
		7678: "Octopus",
		5845: "Octothorpe",
		631:  "Ocular",
		4342: "Odva",
		3421: "Oesolutions",
		4183: "Ohler",
		697:  "Ohm",
		6577: "Ohsung",
		7084: "Oilfind",
		5872: "Okuma",
		5225: "Olencom",
		7780: "OleumTech",
		9363: "Olibra",
		5296: "Olicom",
		1533: "Olitec",
		45:   "Olivetti",
		1971: "Olym-tech",
		168:  "Olympus",
		497:  "Omega",
		7433: "Omneality",
		6421: "Omni-ID",
		571:  "OmniCluster",
		8844: "OmniLync",
		3155: "OmniSense",
		5665: "Omnia",
		5841: "Omnibit",
		5112: "Omnibyte",
		7094: "Omnifi",
		1344: "Omnilux",
		7864: "Omnima",
		8804: "Omniprint",
		5190: "Omnisec",
		6795: "Omnisense",
		975:  "Omnitron",
		7218: "Omnitronics",
		2374: "Omnitronix",
		9279: "Omntec",
		4401: "OnLive",
		2703: "OnSite",
		7748: "OnTarget",
		8904: "OnTime",
		4836: "Onbnetech",
		4657: "Onda",
		6827: "One",
		2658: "OneAccess",
		4616: "OnePath",
		6923: "OnePlus",
		4239: "Oneac",
		5500: "Onelan",
		602:  "Oneline",
		7855: "Onface",
		1523: "OngCorp",
		1430: "Onity",
		2252: "Onkey",
		3200: "Online",
		4953: "Onnto",
		4988: "Onprem",
		3510: "Onset",
		2630: "Ontimetek",
		6870: "Onzo",
		3319: "Ooma",
		4611: "Op-Tection",
		3517: "OpVista",
		3998: "Opaque",
		5086: "Opcom",
		1817: "Opelcomm",
		1879: "Open",
		7424: "Open-m",
		5269: "OpenCon",
		2769: "OpenGear",
		2803: "OpenIB",
		7621: "OpenPattern",
		8385: "OpenVox",
		6412: "OpenXS",
		2932: "Openbrain",
		4070: "Openmoko",
		8009: "Openpeak",
		832:  "Opentech",
		1068: "Ophir-Spiricon",
		1345: "Ophit",
		4942: "Opicom",
		723:  "OptXCon",
		8859: "Optcom",
		1124: "Opteon",
		4146: "Optex",
		9501: "Optex-FA",
		928:  "Opthos",
		4928: "Opti",
		1959: "Opti-cell",
		8915: "OptiLogix",
		550:  "OptiMight",
		3516: "Optibase",
		3048: "Optica",
		1484: "Optical",
		3266: "Opticom",
		3811: "Opticomm",
		6297: "Opticore",
		8214: "Optictimes",
		1037: "Optillion",
		5667: "Optim",
		3394: "Optimal",
		5256: "Optimation",
		3843: "Optimedical",
		5175: "Optimem",
		2290: "Optinel",
		5801: "Optiquest",
		2044: "Optium",
		5655: "Optivision",
		5727: "Opto-22",
		9169: "OptoMET",
		7876: "Optonic",
		4314: "Optos",
		2136: "Optoway",
		7803: "Optowiz",
		484:  "Optranet",
		5016: "Optronic",
		5268: "Optronics",
		3747: "Optsys",
		8077: "Opulinks",
		5563: "Opus",
		6776: "Opzoon",
		11:   "Oracle",
		6878: "Oraimo",
		7859: "Orantek",
		9146: "Orb",
		4961: "Orbacom",
		2129: "Orban",
		6229: "Orbis",
		4182: "Orbotech",
		7666: "Orbotix",
		6318: "Orcatech",
		5416: "Orckit",
		3470: "Ordyn",
		6054: "Oresis",
		7024: "Orga",
		8652: "Oriental",
		1078: "Ormazabal",
		5989: "Ormec",
		4580: "Ortana",
		4224: "Osaka",
		5950: "Ositech",
		1125: "Ositis",
		8334: "Osorno",
		5647: "Osprey",
		8018: "Osram",
		897:  "Otanikeiki",
		791:  "Otari",
		3050: "Otsuka",
		9173: "Ottec",
		7175: "OttoQ",
		7719: "Otus",
		2736: "Ouen",
		4462: "Ouman",
		8705: "Ouster",
		7665: "OuterLink",
		1917: "Outline",
		8664: "Ouya",
		5748: "Ovation",
		9423: "Overkiz",
		818:  "Overture",
		4002: "Owitek",
		8744: "Owl",
		1780: "Oxance",
		8455: "Oxide",
		5303: "Oxtel",
		1648: "Oxygnet",
		7651: "P&S",
		2350: "P-CoM",
		346:  "P-Cube",
		8381: "P.T.I",
		7607: "P2",
		177:  "PAC",
		5300: "PAN-International",
		3936: "PAV",
		3217: "PAX",
		287:  "PAXCOMM",
		6798: "PBR",
		1704: "PC-PoS",
		4405: "PCI",
		1218: "PCO",
		4000: "PCS",
		460:  "PCTEL",
		2657: "PDH",
		2821: "PDL",
		5593: "PEC",
		5346: "PERIPHERALS",
		5461: "PFU",
		7050: "PHAZR",
		6139: "PHC",
		7552: "PHYTRONIX",
		4501: "PIMA",
		997:  "PINON",
		3024: "PKC",
		4964: "PLANET",
		1539: "PLAT'C2",
		6906: "PLATH",
		3277: "PLAYLINE",
		8347: "PLC",
		6540: "PLNetworks",
		7086: "PLUMgrid",
		2558: "PLUS",
		2119: "PLX",
		415:  "PLcom",
		3594: "PMC",
		6093: "PMC-Sierra",
		7178: "PNC",
		1758: "PNI",
		9446: "PNY",
		534:  "PORTech",
		7046: "PRAVIS",
		9230: "PRF",
		5481: "PRO-LOG",
		938:  "PRO-NETS",
		8323: "PRO-VISION",
		3574: "PROBA",
		6553: "PROCOM",
		4060: "PROTEI",
		8690: "PSTec",
		8881: "PTCOM",
		9303: "PYRAMID",
		3561: "PacStar",
		7900: "Pace-O-Matic",
		8782: "Pacidal",
		811:  "Pacific",
		8297: "PacketAccess",
		924:  "PacketAir",
		4105: "PacketFlux",
		2947: "PacketHop",
		895:  "PacketLight",
		2717: "PacketMotion",
		7681: "PacketStorm",
		5431: "Packeteer",
		6127: "Pacom",
		3778: "Pado",
		5162: "Pagine",
		4257: "Pairgain",
		8867: "Pakton",
		1158: "Palm",
		2258: "PalmPalm",
		1355: "Palmmicro",
		4137: "Paltronics",
		3300: "PanAccess",
		7012: "Panaccess",
		6496: "Panamax",
		1315: "Panasas",
		2153: "Panasonic",
		8713: "Pandachip",
		4286: "Pandatel",
		7662: "Pandora",
		2231: "Panduit",
		3078: "Pangolin",
		1553: "Pannaway",
		8067: "Panodic",
		9367: "Panoptic",
		7291: "Panoptics",
		2738: "Panta",
		2277: "Pantech",
		5149: "ParTech",
		3855: "Parade",
		2708: "Paradigm",
		6911: "Paradom",
		8040: "Paradox",
		6119: "Paradyne",
		560:  "Paragea",
		352:  "Paragon",
		3110: "Paralan",
		4289: "Paralink",
		3769: "Parallels",
		332:  "Paralon",
		1488: "Parama",
		3198: "Parama-tech",
		7594: "ParanTek",
		2670: "Paravirtual",
		2562: "Parrot",
		6982: "Parsec",
		5503: "Parsytec",
		8509: "Parta",
		8244: "Particle",
		1867: "Partner",
		6239: "Partron",
		2322: "Parvus",
		8235: "Pason",
		1866: "Passave",
		7139: "PassivSystems",
		7920: "Patech",
		9087: "Paterson",
		5340: "Pathlight",
		6065: "Pathway",
		7981: "Patlite",
		5781: "Patton",
		2371: "Pavo",
		7428: "Pax",
		5133: "Paxdata",
		747:  "Paxonet",
		8669: "PayPal",
		7939: "PayRange",
		3405: "PayTec",
		8698: "Payter",
		4889: "Pcube",
		3612: "PePWave",
		1791: "Pedestal",
		5559: "Peer",
		3252: "Peerless",
		3575: "Pegasus",
		5439: "Pegatron",
		758:  "Pelago",
		3034: "Peltor",
		3718: "Pempek",
		8692: "Penetek",
		1255: "Peninsula",
		5821: "Pensando",
		5968: "Pentacom",
		3622: "Pentaone",
		5357: "Pentek",
		6756: "PentronicAB",
		8283: "PeopleNet",
		2481: "Peplink",
		1947: "Pepperl+Fuchs",
		8483: "Pepwave",
		9067: "Pepxim",
		5259: "Peracom",
		6739: "Peraso",
		5750: "Perception",
		4641: "Perceptron",
		6355: "PerfTech",
		2643: "Perfect",
		2977: "Perfisans",
		5912: "Performance",
		2194: "Peribit",
		7808: "Perinet",
		5445: "Periphonics",
		5378: "Perkin-Elmer",
		5061: "Perle",
		6492: "Perma-Pipe",
		6663: "PernixData",
		8048: "Perples",
		3798: "Perq",
		3356: "Persistent",
		5685: "Personal",
		6863: "Pertronic",
		1636: "Pesa",
		990:  "Petards",
		8646: "Petatel",
		1984: "Petcomkorea",
		272:  "Peterson",
		6707: "Petite-En",
		7573: "Petra",
		3636: "Petratec",
		9142: "Pevco",
		3759: "Pfister",
		4113: "Phabrix",
		2732: "Phantom",
		4618: "PharmaSmart",
		7253: "Pharos",
		6528: "PhaseSpace",
		2314: "Phasecom",
		5379: "Phast",
		4925: "Phasys",
		3550: "PheeNet",
		6848: "Phicomm",
		3635: "Phidgets",
		6640: "Philio",
		796:  "Philips",
		7604: "Phison",
		6276: "Phistek",
		3641: "Phoenix",
		2193: "Phonak",
		2876: "Phonic",
		9340: "Phorm",
		7548: "Phorus",
		2239: "Photometrics",
		4449: "Photon",
		681:  "Photonex",
		3135: "Photonicbridges",
		5918: "Photonics",
		516:  "Photron",
		946:  "Photuris",
		1948: "Phsnet",
		4650: "Phybridge",
		4813: "Phyllis",
		7144: "Physio-Control",
		2889: "Physiometrix",
		8270: "Phytium",
		7940: "Phytrex",
		6834: "Pi-Coral",
		7236: "Pica8",
		5213: "Picazo",
		6896: "Piciorgros",
		5554: "Picker",
		456:  "Pico",
		6237: "PicoCELA",
		691:  "PicoLight",
		3027: "Picochip",
		3451: "Picotest",
		2301: "PictureTel",
		9224: "PiiGAB",
		4005: "Pika",
		3108: "Pilkor",
		4808: "Pilot",
		808:  "Piltofish",
		366:  "Pilz",
		458:  "Pinetron",
		4357: "Ping",
		4104: "Pingood",
		9177: "Pingtek",
		5977: "Pingtel",
		4995: "Pinnacle",
		1016: "Piolink",
		6117: "Pioneer",
		9134: "Pioneercorporation",
		1562: "Pipal",
		5596: "Pipelinks",
		8832: "Piper",
		8633: "Piranti",
		8883: "Pishion",
		4543: "Pitronot",
		4735: "Pittasoft",
		3216: "Pivot3",
		5064: "Pivotal",
		922:  "Pivotech",
		5234: "PixStream",
		8753: "Pixavi",
		7319: "PixeLINK",
		6323: "Pixel",
		4743: "Pixel8",
		4062: "Pixelmetrix",
		1196: "Pixelworks",
		2566: "Pixim",
		635:  "Pixord",
		6291: "Pjrc.com",
		4103: "Planar",
		7339: "Planet",
		4478: "Planex",
		1333: "Planmeca",
		6327: "Planning",
		542:  "Plantronics",
		607:  "Plast-Control",
		4702: "Plaster",
		6512: "Plastoform",
		7317: "Platina",
		3826: "Plato",
		1615: "Platypus",
		1038: "Platys",
		9290: "PlayFusion",
		6974: "Plds",
		609:  "Pleiades",
		2436: "Pleora",
		6337: "Plessey",
		5953: "Plexcom",
		6931: "Plexonics",
		1970: "Plexus",
		9123: "Plexxi",
		8145: "Plugable",
		7575: "Pluribus",
		5971: "Pluris",
		2930: "Plus",
		2993: "Plustek.Inc",
		6097: "Pluto",
		8080: "Pocketbook",
		8953: "PoeWit",
		7750: "Poindus",
		4612: "Pointmobile",
		4479: "Polar",
		8397: "PolarLink",
		1213: "Polaris",
		656:  "Polaroid",
		4766: "Pole/Zero",
		749:  "Polestar",
		4502: "Pollin",
		6633: "Polostar",
		7228: "Poly",
		751:  "Polycom",
		9468: "Polyera",
		3644: "Polygon",
		2181: "Polypix",
		7611: "Polytec",
		2075: "Ponico",
		7923: "Portsmith",
		5691: "Portwell",
		7514: "Posbank",
		4125: "Posbro",
		694:  "Poscon",
		2986: "Posdata",
		6742: "Posh",
		3406: "Posiflex",
		4088: "Positron",
		9323: "Poslab",
		3731: "Postek",
		2576: "Posystech",
		6921: "Power",
		1515: "Power-One",
		7648: "PowerCloud",
		4952: "PowerCom",
		7612: "PowerComm",
		3116: "PowerLink",
		3395: "PowerQuattro",
		7653: "PowerRay",
		8783: "Powercode",
		2520: "Powercom",
		4974: "Powerfile",
		7572: "Powerlinq",
		1379: "Powernet",
		2951: "Powertech",
		7400: "Powervision",
		8099: "Poynt",
		5627: "Praxon",
		6803: "Preceno",
		8161: "Precepscion",
		6518: "Precia",
		7890: "Precidata",
		174:  "Precidia",
		263:  "Precision",
		6887: "Precor",
		1908: "Prediwave",
		9125: "Premietech",
		5707: "Premisys",
		7988: "Prescope",
		6005: "Presence",
		8789: "Presition",
		1547: "Presonus",
		4338: "Presstek",
		6163: "Pressure",
		4240: "Presticom",
		4271: "Pretec",
		6066: "Prevas",
		1902: "Primagraphics",
		998:  "Primarion",
		389:  "Primax",
		1040: "Prime",
		1827: "PrimeNet",
		6431: "PrimeVOLT",
		8525: "Primera",
		8691: "Primmcon",
		850:  "Princeton",
		8343: "PrintCounts",
		5071: "Printer",
		4158: "Printrex",
		1246: "Printronix",
		5358: "Prisa",
		909:  "Prism",
		2110: "Privaris",
		67:   "Private",
		6898: "ProLon",
		6145: "ProMax",
		1264: "ProQuent",
		3832: "ProStor",
		2702: "ProTelevision",
		4492: "ProVision",
		7395: "Probedigital",
		4795: "Probits",
		3568: "Procare",
		8337: "Procentec",
		1499: "Procera",
		1429: "Proces-Data",
		740:  "Procket",
		3092: "Prod-El",
		25:   "Prodigy",
		1634: "Productivity",
		3796: "Prodys",
		6726: "Profalux",
		6406: "Proformatique",
		8560: "Progeny",
		2857: "Prokom",
		5245: "Prolific",
		2465: "Proliphix",
		4341: "Prologix",
		6861: "Promate",
		7204: "Promax",
		7737: "Promethean",
		6120: "Prominet",
		4076: "Prominvest",
		207:  "Promise",
		6999: "Promzakaz",
		4209: "Pronet",
		1801: "Pronto",
		7657: "Prophet",
		8047: "Propulsion",
		532:  "Proscend",
		1954: "Prosoft",
		3375: "Prostar",
		5962: "Prosum",
		8634: "Prosyst",
		1329: "Proteam",
		5578: "Protech",
		4486: "Protecta",
		4847: "ProtectedLogic",
		7386: "Protectron",
		99:   "Proteon",
		4567: "Proteus",
		2391: "Protocol",
		5581: "Proton",
		2870: "Protronic",
		4886: "Proventix",
		735:  "Proview",
		3182: "Proware",
		6770: "Prowave",
		4829: "Proxense",
		3652: "Proxicast",
		6118: "Proxima",
		7411: "Proximus",
		9033: "Proxis",
		8671: "Prudential",
		8792: "Prysm",
		8723: "PsiKick",
		2789: "Psia",
		5854: "Psitech",
		2793: "Psitek",
		3305: "Pulsar-Telecom",
		4264: "Pulse",
		672:  "Pulse-Link",
		7249: "PulseOn",
		836:  "Pultek",
		8078: "Pulzze",
		3784: "Pumpkin",
		1223: "PurOptix",
		3260: "PureTech",
		3911: "PureWave",
		4537: "Purechoice",
		777:  "Puretek",
		7772: "Purple",
		4887: "PurpleComm",
		3875: "Pylone",
		6356: "Pyramid",
		5684: "Pyrescom",
		3502: "Pyronix",
		4458: "Pyung-HWA",
		8431: "Q",
		6514: "Q-Lab",
		4581: "Q-Light",
		7885: "Q9",
		3119: "QDI",
		845:  "QEI",
		2025: "QLogic",
		6761: "QNAP",
		311:  "QPS",
		5356: "QSC",
		566:  "QSI",
		9486: "QTS",
		4968: "QTelNet",
		4691: "QUALICA",
		4865: "QVidium",
		2210: "Qamcom",
		7507: "Qardio",
		4089: "Qbit",
		2003: "Qcom",
		8874: "Qdis",
		8098: "Qfiednet",
		6669: "QianGua",
		7914: "QianTang",
		7352: "Qibixx",
		9251: "Qihan",
		6425: "QiiQ",
		2541: "QinetiQ",
		7429: "Qingping",
		551:  "Qisda",
		3824: "Qniq",
		5095: "Qnix",
		3173: "Qno",
		1787: "QoStek",
		7058: "Qolsys",
		4513: "Qool",
		6652: "Qorvo",
		6723: "Qotom",
		2609: "Qovia",
		8094: "Qpcom",
		30:   "Qpsx",
		1174: "Qqest",
		2720: "Qsan",
		7304: "Qsono",
		3172: "Qstik",
		4145: "Qtech",
		3770: "Qtum",
		5433: "QuVis",
		5936: "Quad/Graphics",
		141:  "Quadram",
		3120: "Quadrics",
		1247: "Quake",
		5787: "Qualcomm",
		8793: "Qualisys",
		3487: "Qualitrol",
		1210: "Qualstar",
		9006: "Quanergy",
		8117: "QuantHouse",
		3071: "Quanta",
		5565: "Quantel",
		3093: "Quantier",
		8861: "Quantify",
		2083: "Quantum",
		3876: "QuantumVision",
		7858: "Quantumsolution",
		6675: "Quarion",
		280:  "Quarry",
		3858: "Quartics",
		4734: "Quasar",
		1625: "Quatech",
		7558: "Quatius",
		8293: "Qudelix",
		2507: "Quest",
		5730: "Questech",
		7801: "Quext",
		3223: "QuickTel",
		5170: "Quin",
		5875: "Quintar",
		6390: "Quintic",
		2687: "Quintron",
		7962: "Quirky",
		2896: "Quixant",
		9487: "Qulsar",
		5124: "Qume",
		3548: "Qumranet",
		4708: "QuoPin",
		155:  "Quotron",
		6644: "Qvis",
		6952: "R&M",
		5784: "R.A",
		414:  "RACOM",
		588:  "RADVision",
		2974: "RADWIN",
		7535: "RADiflow",
		2622: "RAE",
		8506: "RAID",
		6502: "RAONIX",
		4811: "RAUMFELD",
		9098: "RBcloudtech",
		105:  "RC",
		3491: "RCG",
		8121: "RDA",
		1402: "RDI",
		2152: "RDM",
		3299: "REJ",
		2118: "REMEC",
		2855: "RF",
		7111: "RFI",
		2953: "RFID",
		4978: "RFM",
		1019: "RFNET",
		573:  "RFTNC",
		3712: "RFTech",
		2427: "RGB",
		5970: "RHK",
		508:  "RIAS",
		4555: "RIM",
		4147: "RIVA",
		3203: "RIX",
		9333: "RLH",
		2995: "RLW",
		5187: "RLX",
		9261: "RM",
		6821: "RNET",
		6809: "RNware",
		8666: "ROBOTIS",
		4026: "ROBOTOUS",
		239:  "ROI",
		871:  "ROIS",
		7247: "ROLI",
		1302: "ROMWin",
		1847: "ROUND",
		4151: "RSD",
		6778: "RSF",
		5254: "RSI",
		8064: "RTC",
		171:  "RTUnet",
		9402: "RTW",
		5568: "RWT",
		3363: "Raba",
		1045: "Racal-Datacom",
		1367: "Racewood",
		8802: "RackTop",
		9450: "RackWare",
		3424: "Rackable",
		8088: "Racktivity",
		58:   "Racore",
		4643: "RadChips",
		3473: "RadarFind",
		4955: "Radcom",
		8332: "Raden",
		5811: "Radguard",
		7761: "RadiAnt",
		2756: "Radiance",
		228:  "Radiant",
		1576: "Radiantech",
		5056: "Radicom",
		3930: "Radiient",
		8150: "Radinet",
		3342: "RadioCOM",
		2965: "RadioPulse",
		4086: "Radiocomp",
		2940: "Radiocrafts",
		5792: "Radiolan",
		1138: "Radionet",
		3940: "Radionor",
		7302: "Radios",
		3900: "Radioscape",
		52:   "Radisys",
		684:  "Radius",
		1081: "Radlan",
		4328: "Radlive",
		6702: "Radmax",
		7851: "Radspin",
		5516: "Radstone",
		563:  "Radware",
		9192: "Radwin",
		2337: "Radyne",
		7911: "Rafael",
		6839: "Ragentek",
		7835: "Ragile",
		8439: "Ragsdale",
		4213: "Ragula",
		4646: "Raidon",
		2408: "Raidtec",
		7306: "RailComm",
		7105: "Railtec",
		7905: "RainUs",
		1538: "Rainsun",
		2050: "Raisecom",
		8042: "Rajant",
		1785: "Ralink",
		6242: "Ramaxel",
		9275: "Rami",
		5047: "Ramix",
		23:   "Ramtek",
		1837: "Rancho",
		8377: "Rancore",
		5419: "Randata",
		2546: "Rane",
		5824: "Rantic",
		9317: "RaonThink",
		1186: "Raonet",
		901:  "RapidWAN",
		8784: "Rapidmax",
		2238: "Raptor",
		1927: "Raritan",
		626:  "Rasteme",
		1228: "Rasvia",
		5562: "Rational",
		8470: "Ratp",
		5621: "Rauland-Borg",
		8015: "RayTight",
		9131: "Raybased",
		3476: "Raycom",
		5130: "Raylan",
		4428: "Raylase",
		8255: "Raylios",
		4284: "Raynet",
		362:  "Raysis",
		2601: "Rayson",
		5876: "Raytech",
		96:   "Raytheon",
		4649: "Rayzone",
		7181: "Razer",
		3818: "Razorstream",
		8704: "Rcntec",
		6912: "Reach",
		6531: "Reacheng",
		2764: "ReadyLinks",
		250:  "ReadyNet",
		4008: "Realease",
		7188: "Reallin",
		2701: "Realm",
		2330: "Rebel.com",
		4021: "Recall",
		8228: "Rechnerbetriebsgruppe",
		9454: "Reciprocal",
		487:  "Reco",
		6383: "Recovision",
		2981: "Red-Lemon",
		1485: "Red-M",
		5712: "Redcom",
		5355: "Redcreek",
		761:  "Reddo",
		4991: "Redflex",
		7685: "Redflow",
		1314: "Redline",
		885:  "Rednix",
		910:  "Redswitch",
		2602: "Redux",
		9344: "Reduxio",
		9250: "Redwire",
		1932: "Redwood",
		7556: "Refined",
		238:  "Refraction",
		6864: "Regenersis",
		5979: "Regent",
		3372: "Reigncom",
		739:  "Relax",
		8622: "Relay2",
		2375: "Reliance",
		1722: "Remopro",
		2944: "Remote",
		1281: "Remotec",
		1509: "Remotek",
		2100: "Remsdaq",
		2802: "Renasis",
		5695: "Renesas",
		5774: "Renex",
		882:  "Renishaw",
		3725: "Renkus-Heinz",
		3273: "Renu",
		1201: "Repeatit",
		9139: "Repotec",
		5531: "Republic",
		4551: "ResMed",
		8621: "Research",
		5986: "Resideo",
		5385: "Resilience",
		6190: "Reson",
		4166: "Respironics",
		6737: "RetailNext",
		3022: "Reti",
		4232: "Reudo",
		5803: "Reuters",
		8570: "Revolv",
		3001: "Rextechnik",
		7822: "Rezolt",
		5573: "Rftech",
		7198: "Rhino",
		8900: "Rhombus",
		1707: "RiT",
		6607: "Riava",
		4164: "Riccius+Sohn",
		75:   "Ricoh",
		6341: "Ridge",
		3452: "Riedel",
		8989: "Riedo",
		1769: "Rifatron",
		8230: "Rigado",
		1181: "Rigaku",
		3583: "RightHand",
		5034: "Rightech",
		3479: "Rigol",
		1658: "Rincon",
		6946: "Ring",
		4784: "RingBell",
		4739: "RingCube",
		7712: "Rinicom",
		7762: "Rinstrum",
		2325: "Rion",
		5222: "Rioworks",
		3681: "Ripcode",
		194:  "RiscStation",
		4067: "Risco",
		2403: "Rise",
		6716: "Risk",
		3237: "Rittmeyer",
		4250: "Riva",
		5013: "RiverDelta",
		2093: "Riverbed",
		1823: "Riverhead",
		394:  "Riverstone",
		1811: "Rivertec",
		2800: "Rivertree",
		8338: "Rivet",
		862:  "Roax",
		3000: "Robatech",
		956:  "Robinson",
		4791: "Robonica",
		2180: "Robox",
		9027: "Rockeetech",
		662:  "RocketLogix",
		5996: "Rocketchips",
		6286: "Rockport",
		3871: "Rockridgesound",
		4403: "Rohati",
		3865: "Rohm",
		1916: "Roku",
		1790: "Roll",
		1874: "Rolls-Royce",
		6095: "Rooftop",
		7631: "RoomReady/Zdi",
		5974: "Root",
		6965: "Rootech",
		8485: "Roqos",
		8177: "Rorze",
		646:  "Rosco",
		5144: "Rose",
		4093: "Roseman",
		4821: "Rosemount",
		7635: "Rosewill",
		9421: "Rosonix",
		5725: "Ross",
		9437: "Rosslare",
		8454: "Rossma",
		4717: "Rotel",
		1267: "RouteFree",
		1784: "Routerboard.com",
		507:  "Routrek",
		951:  "Roving",
		7878: "Roxton",
		4077: "RoyalTek",
		3154: "Royaldigital",
		2410: "Rpcg",
		7518: "Rramac",
		9351: "RtBrick",
		7493: "Rtnet",
		7155: "Rubezh",
		4628: "Ruby",
		2737: "Ruckus",
		1586: "RuggedCom",
		5440: "Ruijie",
		4731: "Runcom",
		524:  "Runtop",
		8581: "Ruroc",
		4375: "Russound",
		4493: "Rustelcom",
		8797: "Rusteletech",
		2762: "RyCo",
		6393: "Ryowa",
		2778: "Ryvor",
		3889: "S&O",
		4788: "S-Access",
		7551: "S-Bluetech",
		2807: "S-TEC",
		1777: "S-Takaya",
		4635: "S.A.A.A",
		1968: "S.A.Tehnology",
		330:  "S.D.E.L",
		5861: "S.E.R.C.E.L",
		8884: "S.E.Technologies",
		7163: "S.FAC",
		2404: "S.I",
		211:  "S1",
		7442: "S2M",
		1882: "S2io",
		3359: "S3C",
		3744: "SAE",
		5501: "SAI",
		7296: "SAM",
		6811: "SAMJIN",
		4586: "SAMSUNG",
		8367: "SAMWONFA",
		2429: "SANBlaze",
		3661: "SANION",
		1923: "SANYCOM",
		6174: "SAT",
		3530: "SATEC",
		344:  "SAXA",
		9415: "SB",
		135:  "SBE",
		8728: "SBM",
		3969: "SBN",
		5976: "SBS",
		6699: "SCAPS",
		4884: "SCDI",
		8921: "SCR",
		8169: "SCS",
		8834: "SCSpro",
		7614: "SDJ",
		6132: "SDL",
		1232: "SDSystem",
		6415: "SDTEC",
		4741: "SE-Elektronic",
		5088: "SEA-Ilan",
		463:  "SEAKR",
		7462: "SECUDOS",
		1461: "SED",
		2957: "SEECODE",
		4430: "SEEnergy",
		8649: "SEMA",
		9387: "SEMOCON",
		5444: "SENTRY",
		979:  "SEPATON",
		4124: "SERONICS",
		8735: "SES-imagotag",
		4948: "SET",
		307:  "SETA",
		6772: "SEnergy",
		6394: "SF",
		4172: "SFA",
		1865: "SFOM",
		7056: "SFORZATO",
		3187: "SFR",
		24:   "SHA-KEN",
		7877: "SHIFT",
		6204: "SI",
		1618: "SIBCO",
		764:  "SICOM",
		7645: "SIEMENS",
		7213: "SIFROM",
		1836: "SIGMACOM",
		4860: "SIL3",
		3653: "SIMS",
		8938: "SINTRONES",
		7152: "SK",
		6207: "SK-Elektronik",
		268:  "SKF",
		3098: "SKNET",
		8917: "SM-Electronic",
		4971: "SMAR",
		388:  "SMART",
		9000: "SMAX",
		741:  "SMC",
		6641: "SMG",
		244:  "SMK-M",
		4923: "SMP",
		4419: "SMT&C",
		9155: "SNB",
		6537: "SNK",
		3933: "SNR",
		896:  "SNS",
		1467: "SOHOware",
		4868: "SOKRAT",
		8279: "SOL",
		7278: "SOLARWATT",
		7451: "SOLEM",
		2275: "SOLOMON",
		9102: "SOLiD",
		1550: "SONICblue",
		932:  "SONO",
		648:  "SONOS",
		7740: "SOUND4",
		8925: "SOYEA",
		4895: "SP",
		1237: "SPAUN",
		6113: "SPC",
		5489: "SPHINX",
		2774: "SPIDCOM",
		9348: "SPON",
		6457: "SPRINGWAVE",
		8701: "SPS",
		5236: "SPX-Ateg",
		7650: "SPnS",
		8076: "SRC",
		3080: "SSD",
		4624: "SSI",
		4755: "ST",
		8084: "STABILO",
		2742: "STAC",
		9280: "STI",
		4759: "STJ",
		6527: "STMicrolectronics",
		3965: "STN",
		3066: "STNet",
		7616: "STORDIS",
		5797: "STS",
		7894: "STULZ",
		5564: "SUN",
		7142: "SUNGSAM",
		1472: "SUNIX",
		3059: "SUNWAVETEC",
		3850: "SUPoX",
		363:  "SURECOM",
		7528: "SUYIN",
		1486: "SVA",
		1049: "SVA-Intrusion",
		8512: "SVS-VISTEK",
		3046: "SWEEX",
		3310: "SWsoft",
		9452: "SYN-Tech",
		9407: "SYNTEC",
		3879: "SYRIS",
		2711: "Saab",
		6335: "Saber",
		3919: "SabiOso",
		9171: "Sabre",
		6574: "Sabtech",
		2032: "Safari",
		5343: "Safe-Com",
		7934: "SafeTone",
		1977: "SafeWeb",
		5988: "Safetran",
		7349: "Safetrust",
		8022: "Saffron",
		3503: "Sagamore",
		619:  "Sage",
		4572: "SageTV",
		2048: "Sagemcom",
		6586: "Sagittar",
		4448: "Sagrad",
		5393: "Sahara",
		6614: "Sailer",
		7158: "Salcomp",
		3492: "Salicru",
		5260: "Salix",
		1845: "Salland",
		7968: "Salutron",
		7308: "SamJi",
		8839: "SambaNova",
		4569: "Samjeon",
		1768: "Sammy",
		2970: "Sampo",
		5586: "Samsan",
		9495: "Samsara",
		6168: "Samson",
		152:  "Samsung",
		6275: "Samtec",
		3620: "Samyoung",
		4019: "San-Eisha",
		608:  "SanCastle",
		5332: "SanCom",
		3657: "SanDisk",
		8661: "SanJet",
		7720: "SanLogic",
		4529: "SandForce",
		4349: "SandLinks",
		710:  "SandStream",
		2209: "Sandburst",
		1901: "Sanden",
		2128: "SandmartinElectronics",
		7530: "Sandstone",
		1347: "Sandvine",
		9414: "Sanford",
		1258: "SangSang",
		3867: "Sangean",
		8476: "Sangfor",
		5228: "Sangoma",
		9040: "Sanix",
		617:  "Sanko",
		6441: "Sankosha",
		7879: "SankyuElectronics",
		2208: "Sanmei",
		3536: "Sanmina-SCI",
		1046: "Sanritz",
		7277: "Sanshin",
		4843: "Sansonic",
		3986: "Santec",
		4976: "Santera",
		9453: "Santur",
		1337: "Sanyo",
		2606: "Sanyu",
		7524: "Sapling",
		9097: "Sapphire",
		637:  "Sarian",
		6170: "Sarnoff",
		2270: "Sarotech",
		4258: "Sast",
		888:  "Satec",
		4361: "Satel",
		7949: "Satelco",
		5446: "Satelcom",
		8008: "Satmap",
		3464: "Sato",
		2961: "Saunders",
		3588: "Savant",
		6561: "Savitech",
		2819: "Savvius",
		7924: "Sawwave",
		683:  "Scalant",
		7792: "Scalar",
		2086: "Scalent",
		7724: "Scalys",
		3686: "Scan",
		5465: "Scan-Optics",
		5350: "Scanivalve",
		1903: "Scanmatic",
		417:  "Scannex",
		2675: "Scanvaegt",
		1975: "Schiller",
		685:  "Schlumberger",
		4777: "Schmartz",
		56:   "Schneider",
		6584: "Schreder",
		5005: "Schweitzer",
		4476: "SciLog",
		3036: "Science",
		8959: "Scientech",
		1474: "Scientific",
		3115: "Scientific-Atlanta",
		4316: "Scimolex",
		4197: "Scinets",
		1624: "Scion",
		7068: "Sciovid",
		57:   "Scitex",
		5897: "Scope",
		6240: "Scopus-Belgium",
		2808: "Scosche",
		1593: "ScottCare",
		3264: "Screen",
		8983: "ScreenBeam",
		1656: "ScriptPro",
		6420: "SeAH",
		4332: "SeaMicro",
		1569: "Seabridge",
		732:  "Seagate",
		7537: "Seal",
		1473: "Sealevel",
		6271: "Seamap",
		4705: "Seanodes",
		1884: "Seaway",
		1386: "SecWell",
		590:  "Secheron",
		2464: "Secom-Industry",
		7490: "Secret",
		4907: "Sectronic",
		8895: "SecuGen",
		7028: "SecuWorks",
		807:  "Secui.com",
		2969: "Securaplane",
		1382: "Securebase",
		5046: "Securelogix",
		7357: "Securetech",
		9179: "Securifi",
		1225: "Securiton",
		8771: "Securosys",
		3256: "Sedo",
		5037: "Sedona",
		359:  "Seedek",
		933:  "Seedsware",
		6893: "Seeed",
		8788: "SeekTech",
		2845: "SeekerNet",
		5923: "Seel",
		5753: "Seeq",
		4701: "Seer",
		6408: "Seers",
		8908: "Seetech",
		6087: "Sega",
		2346: "Segate",
		3795: "Segnet",
		4094: "Segnetics",
		5460: "Seiko",
		119:  "Seikosha",
		814:  "Seiwa",
		4346: "Sejin",
		4380: "Sekonic",
		4069: "Select",
		187:  "Selectron",
		3180: "Selex",
		2409: "Selsius",
		3563: "Seluxit",
		6666: "Selve",
		5890: "Semaphore",
		4888: "Semihalf",
		3666: "Semindia",
		4329: "Semptian",
		3139: "Semtech",
		9284: "SenRa",
		4749: "SenTec",
		245:  "Sena",
		383:  "Senao",
		937:  "Sencore",
		2755: "Sendo",
		195:  "Sendtek",
		3148: "Senea",
		603:  "Seneca",
		4713: "Senet",
		5978: "Senetas",
		2885: "Senhai",
		8562: "Senient",
		3117: "Sennheiser",
		9356: "Senor",
		1694: "SensAble",
		6688: "Sensaio",
		1177: "Sensaphone",
		4866: "Senscient",
		9209: "Sense",
		4334: "SenseAnywhere",
		8586: "Senseit",
		8499: "Senselogix",
		3410: "Sensicast",
		8719: "SensingTek",
		3444: "SensoPart",
		9330: "Sensometrix",
		6738: "SensorPush",
		7038: "SensorTec-Canada",
		829:  "Sensoria",
		5302: "Sensormatic",
		2612: "Sensory",
		2826: "Sensovation",
		3870: "Sensus",
		9076: "Sentec",
		4528: "Senticare",
		5362: "Sentient",
		2604: "Sentilla",
		9217: "Sentinhealth",
		9484: "Sentinum",
		3247: "Sentivision",
		7148: "Seongji",
		1774: "SeorimTechnology",
		8241: "Seoul",
		4299: "Seowonintech",
		8967: "Seowoo",
		2905: "Sepsa",
		4016: "Sepura",
		3044: "Sequans",
		4311: "Sequel",
		6324: "Sequent",
		1442: "Seranoa",
		1728: "SercoNet",
		2074: "Sercomm",
		8766: "Serelec",
		4382: "SerialTek",
		7626: "Serialway",
		4168: "Seritech",
		6255: "Sernet",
		355:  "Serome",
		6246: "Sertel",
		1552: "Server",
		3107: "ServerEngines",
		8029: "ServerU",
		8144: "Servercom",
		7047: "Servergy",
		4057: "Servimat",
		3081: "SetOne",
		2575: "SetaBox",
		4878: "Setrix",
		7581: "Seungil",
		2472: "Sevis",
		7206: "Sewoo",
		4982: "Seyeon",
		5008: "Shanghai",
		6650: "Shannon",
		6533: "Shany",
		894:  "ShareGate",
		6024: "Sharewave",
		6462: "SharkNinja",
		3205: "Sharp",
		5606: "Shasta",
		6432: "Shaw",
		1044: "Sheba",
		191:  "Shelcad",
		771:  "Shellcomm",
		7292: "ShengHai",
		8624: "Shenzhen",
		3618: "Shenztech",
		6142: "Sherwood",
		2031: "Shester",
		1350: "ShibaSoku",
		6167: "Shimadzu",
		1457: "Shimizu",
		8419: "Shin-IL",
		1597: "Shin-OH",
		3008: "ShinMaywa",
		8477: "Shinbo",
		2454: "Shinboram",
		3162: "Shinco",
		7445: "Shineway",
		8896: "Shiningtek",
		5547: "Shinnihondenko",
		4161: "Shinsei",
		3906: "Shinwa",
		7893: "Shinybow",
		9085: "Shinyei",
		3975: "Shireen",
		5548: "Shiva",
		5134: "Shographics",
		5418: "Shomiti",
		8463: "Shoogee",
		2321: "ShoreTel",
		8971: "ShotSpotter",
		7693: "ShotTracker",
		4525: "Shouyo",
		2035: "Shuko",
		3815: "Shunra",
		2117: "Shure",
		4937: "Shuttle",
		4391: "Si14",
		367:  "SiByte",
		548:  "SiConnect",
		2798: "SiCortex",
		2571: "SiNett",
		3142: "SiRF",
		7516: "Siama",
		963:  "Sick",
		8800: "Siconix",
		2810: "Sidsa",
		74:   "Siecor",
		4516: "Sielox",
		300:  "Siemens",
		3996: "Siemon",
		329:  "SierraCom",
		7373: "Sify",
		3281: "SightLogix",
		291:  "Sigma",
		1241: "Sigma-Links",
		1802: "SigmaTel",
		4015: "Sigmalink",
		6575: "Sigmastar",
		5299: "Sigmatek",
		144:  "Sigmex",
		569:  "Signal",
		4118: "Signalion",
		3903: "Signamax",
		3662: "Signtech",
		5289: "Signum",
		3041: "Sigpro",
		7453: "Sigrand",
		9294: "Sigrist-Photometer",
		120:  "Siig",
		4672: "Siklu",
		6996: "Silca",
		5518: "Silex",
		6213: "Silicom",
		1655: "Silicon",
		3153: "SiliconStor",
		3376: "Silicondust",
		2584: "Silink",
		7284: "Silkan",
		7315: "SilverNet",
		7515: "SilverPlus",
		7908: "Silverbrook",
		7176: "Silverflare",
		4156: "Simet",
		3518: "Simoco",
		9350: "Simon-Kaloi",
		8229: "SimonsVoss",
		39:   "Simpact",
		4170: "Simple",
		3643: "SimpleComTools",
		279:  "Simpler",
		9411: "SimpliSafe",
		5629: "Simrad",
		2511: "Simtec",
		8083: "Simton",
		5584: "Simulation",
		2179: "Sinar",
		1451: "Sinbon",
		3292: "Sindoricoh",
		1101: "Sinetica",
		7897: "Sinewy",
		4740: "Singapore",
		2490: "Singim",
		2489: "Singular",
		3648: "Sinkyo",
		7562: "Sino-Telecom",
		3924: "Sinotech",
		8978: "Sinwatec",
		9325: "Sipod",
		3241: "Sirit",
		8183: "Sirius",
		8458: "Siselectron",
		8197: "Sisnet",
		260:  "Sitara",
		6338: "Sitasys",
		8262: "Sitcorp",
		1881: "Sitecom",
		1366: "Sitecsoft",
		5031: "Sitek",
		301:  "Sitera",
		9113: "Sitronik",
		7210: "Sius",
		6712: "Sivantos",
		2565: "Skardin",
		686:  "Skidata",
		7358: "Skipper",
		8120: "Skiva",
		6605: "Skorpios",
		2109: "Skov",
		6262: "Skspruce",
		7039: "Skullcandy",
		6997: "Sky-City",
		7684: "SkyBell",
		8341: "SkyDisk",
		8025: "SkyHawke",
		8956: "Skybell",
		4117: "Skydigital",
		7000: "Skydio",
		7303: "Skyera",
		6913: "Skymotion",
		4938: "Skystream",
		8391: "Skytap",
		9017: "Skyviia",
		1706: "Smallbig",
		3455: "SmarDTV",
		3239: "SmarTire",
		7321: "Smart",
		6741: "SmartCap",
		8329: "SmartDoor",
		6782: "SmartDrive",
		8264: "SmartOptics",
		4549: "SmartRG",
		4135: "SmartShare",
		4142: "SmartSynch",
		6789: "SmartThings",
		4936: "Smartbridges",
		9287: "SmarteBuilding",
		8597: "Smartisan",
		2133: "Smartlabs",
		1015: "Smartmatic",
		7380: "Smarto",
		775:  "Smartronix",
		8526: "Smartrove",
		6044: "Smartsan",
		2765: "Smartvue",
		4999: "Smartware",
		6748: "Smnd",
		6268: "Smobile",
		8770: "SnD",
		8294: "Snap",
		6556: "SnapAV",
		7145: "SnapRoute",
		1320: "SnedFar",
		822:  "Snell",
		5409: "Snmp",
		753:  "SnowShore",
		7217: "Snuza",
		6613: "Soarnex",
		9488: "Socionext",
		5860: "Socket",
		3196: "Socomec",
		7909: "Socus",
		2329: "Sodick",
		1290: "SofaWare",
		3628: "Sofacreal",
		1220: "SoftEnergy",
		1162: "SoftRadio",
		5538: "Softcom",
		3296: "Softier",
		959:  "Softing",
		9162: "Softiron",
		5321: "Softlab",
		4608: "Softwell",
		8997: "Sogecam",
		3722: "Sogestmatic",
		1441: "Solacom",
		7540: "Soladigm",
		3208: "Solar",
		4906: "SolarEdge",
		7977: "Solarbridge",
		1524: "Solarflare",
		95:   "Solbourne",
		5895: "Solectek",
		9371: "SolidFire",
		8941: "SolidRun",
		7620: "Solidica",
		6078: "Solidum",
		9399: "Solidwintech",
		874:  "Solinet",
		3649: "Solitech",
		6883: "Soliton",
		5049: "Sollae",
		4343: "Soltech",
		2435: "Solteras",
		4386: "Solutronic",
		369:  "Soma",
		6010: "Somat",
		5886: "Somelec",
		7301: "Somfy",
		7107: "Sonance",
		7441: "Sonar",
		8160: "Sonardyne",
		9050: "Sonavation",
		1279: "Soneticom",
		5066: "Sonic",
		8382: "SonicSensory",
		3262: "SonicWALL",
		6639: "SonicWall",
		1003: "Sonicwall",
		8950: "Sonifex",
		4671: "Sonim",
		1372: "Sonitor",
		3458: "Sonitrol",
		4243: "Sonix",
		4998: "Sonnet",
		1309: "SonoSite",
		2931: "Sonoa",
		6360: "Sonoma",
		2047: "Sonos",
		6859: "Sonova",
		9397: "Sontex",
		2339: "Sonus",
		101:  "Sony",
		5539: "Sophia",
		3578: "Sophos",
		7307: "Soraa",
		5311: "Sord",
		2274: "Sordin",
		7233: "Soreel",
		1230: "Sorenson",
		1073: "Soriya",
		864:  "Soronti",
		5790: "Sotas",
		5310: "Sotec",
		8386: "SoundBridge",
		4414: "SoundEar",
		7701: "SoundHawk",
		2269: "Soundcraft",
		6867: "Soundmatters",
		6484: "Soundmax",
		6526: "Source",
		4174: "Source-Comm",
		9310: "Sourcefire",
		6822: "SourcingOverseas",
		8100: "Sovico",
		2700: "Soyal",
		5208: "Soyo",
		6482: "Spacelink",
		1976: "Spagat",
		2073: "SparkLAN",
		3228: "Sparr",
		7179: "Spawn",
		8680: "Spbec-Mining",
		6296: "Speaker",
		4090: "Speakercraft",
		5541: "Specialix",
		7510: "Speco",
		2654: "Spectec",
		6046: "Spectel",
		2587: "Spectra",
		32:   "Spectragraphics",
		5633: "Spectralink",
		4285: "Spectrix",
		9049: "Spectronix",
		5530: "Speed",
		8821: "Speedytel",
		3754: "Sphairon",
		6101: "Sphere",
		8837: "Spica",
		154:  "Spider",
		3888: "Spinetix",
		622:  "Spinnaker",
		2333: "Spirent",
		1152: "Splicecom",
		8697: "SpotCam",
		7116: "Spreadtrum",
		6819: "Sprocomm",
		5729: "Spur",
		3409: "Sputnik",
		6751: "Squarehead",
		7471: "Squirrels",
		2837: "Srisa",
		5093: "Ssangyong",
		3434: "Staccato",
		5322: "Stallion",
		8219: "Stalmart",
		3482: "Stanford",
		5735: "Stanilite",
		5100: "Star",
		7194: "Star-Net",
		5106: "Star-TEK",
		4863: "StarLeaf",
		9233: "StarTech.com",
		3716: "StarVedia",
		4083: "Starbridge",
		5045: "Stardot",
		793:  "Starent",
		1561: "Stargames",
		8331: "Starkey",
		5155: "Starlight",
		2365: "Starnet",
		2536: "Starnex",
		8143: "Starry",
		9257: "Startel",
		8672: "Stateless",
		5809: "Staubli",
		5152: "Stealth",
		1529: "Stec",
		7757: "Steelcase",
		7363: "Steffes",
		4207: "Steinbrecher",
		8914: "Steinel",
		9037: "Steinsvik",
		7124: "Stella-Green",
		5540: "Stellar",
		533:  "Stellcom",
		7866: "Stephen",
		7714: "Stereotaxis",
		7884: "Sterlite",
		3973: "Stim",
		2603: "Stoke",
		2917: "Stolinx",
		4051: "Stoneridge",
		7440: "Stonesoft",
		1835: "StorCase",
		7237: "StorSimple",
		2398: "Storage",
		8847: "Storagedata",
		7244: "Store",
		1973: "Stormshield",
		4636: "Storwize",
		1098: "Stralfors",
		2955: "StrataLight",
		6014: "Stratabeam",
		6892: "Stratacache",
		5880: "Stratacom",
		6034: "Strategy",
		719:  "Stratex",
		108:  "Stratus",
		1483: "Stream",
		9478: "StreamCCTV",
		5421: "StreamLogic",
		2187: "StreamScale",
		7318: "StreamUnlimited",
		2443: "Streamit",
		2178: "Stretch",
		944:  "Strix",
		4875: "Strong",
		9377: "Structab",
		4421: "Strukton",
		1988: "Stryker",
		9176: "Stuart",
		632:  "Studio",
		8276: "Sub10",
		7702: "Subscriber",
		4805: "Suga",
		714:  "Sullair",
		4654: "Sumavision",
		872:  "Sumtel",
		819:  "Sun",
		2649: "SunCorp",
		2812: "SunKwang",
		4497: "SunPower",
		7711: "SunReports",
		2411: "Sundance",
		3753: "Sunell",
		8857: "Sunflex",
		4231: "Sungwoon",
		3079: "Sunhillo",
		3939: "Sunitec",
		2754: "Sunmyung",
		5435: "Sunnovo",
		2426: "Sunplus",
		2285: "Sunray",
		8932: "Sunrex",
		1575: "Sunrich",
		8864: "Sunrise",
		3949: "Sunshine",
		9183: "Sunstar",
		9145: "Suntaili",
		8322: "Suntec",
		6284: "Suntech",
		7636: "Sunwave",
		3147: "Sunways",
		7770: "Sunwoda",
		2834: "SuperVision",
		1365: "Supercaller",
		4014: "Supercom",
		5560: "Supercomputing",
		86:   "Supernet",
		4927: "Superpower",
		3276: "Suprema",
		19:   "Sureman",
		2621: "Surf",
		703:  "Surgient",
		5475: "Surigiken",
		2323: "Surtec",
		1114: "Sutron",
		4473: "Sutus",
		6437: "Suunto",
		3602: "Suzuken",
		5899: "Svec",
		6389: "Svyazkomplektservice",
		6403: "Swaive",
		8721: "Swann",
		2607: "Swegon",
		5307: "Swelaser",
		7272: "Sweroam",
		8315: "SwiftTest",
		3373: "Swirlnet",
		8130: "Swissbit",
		4096: "Swissdis",
		835:  "Swissvoice",
		5263: "Switchcore",
		3675: "Swyx",
		1028: "Syabas",
		3959: "Syba",
		5172: "Sybus",
		6928: "Sycada",
		4214: "Sycamore",
		1667: "Sychip",
		4944: "Sylantro",
		4344: "SymCom",
		7251: "Symanitron",
		1989: "Symantec",
		6012: "Symbionics",
		6301: "Symbolics",
		7388: "Symeo",
		1864: "Symetrix",
		5327: "Symicron",
		122:  "Symmetric",
		4294: "Symmetrical",
		5745: "Symmetricom",
		6185: "Symon",
		5185: "Symplex",
		2840: "Symwave",
		4025: "Symx",
		6773: "SynTrust",
		5055: "SynapSense",
		3757: "Synapse",
		7569: "Synaptec",
		7117: "Synapticon",
		7969: "Synaptics",
		5119: "Sync",
		9202: "Syncbak",
		3145: "Synccom",
		2149: "Synchronic",
		801:  "Synchronous",
		2230: "Synchrony",
		5354: "Synclayer",
		7326: "Syncmold",
		1145: "Synectics",
		7067: "Synerchip",
		7547: "Synergics",
		4238: "Synergy",
		5090: "Synerjet",
		5476: "Synernetics",
		9149: "Synertau",
		2450: "Synology",
		1102: "Synoptics",
		7511: "Syntech",
		5941: "Syntellect",
		40:   "Syntrex",
		6783: "Syntronic",
		4745: "Syphan",
		2506: "Sypixx",
		4499: "Syracuse",
		3257: "Syrinx",
		7022: "Syrotech",
		6251: "SysDINE",
		8281: "SysGRATION",
		60:   "SysKonnect",
		7254: "SysTec",
		5115: "Sysgen",
		9048: "Syslane",
		2788: "Sysmaster",
		5997: "Sysmate",
		955:  "Sysmex",
		1266: "Syspol",
		1257: "Systec",
		5479: "Systech",
		1737: "Systegra",
		1554: "Systek",
		4544: "Systel",
		578:  "SystemGear",
		1324: "SystemK",
		9103: "Systembase",
		5148: "Systemforschung",
		7464: "Systemhouse",
		1187: "Systemonic",
		7634: "Systems",
		2874: "Systimax",
		5678: "Systran",
		6409: "Systrome",
		8872: "Systronics",
		1617: "Systronix",
		3794: "Syswan",
		209:  "Syswave",
		7622: "Syszone",
		6:    "Sytek",
		1953: "T&D",
		4048: "T&W",
		7104: "T-Mobile",
		4844: "T-Platforms",
		2843: "T-Vips",
		4190: "T.C",
		5557: "T.D.I",
		4834: "T.M",
		2580: "T.O.M",
		2300: "T.Sqware",
		7579: "T3",
		2379: "TAC",
		5262: "TAG",
		7332: "TAKT",
		503:  "TAMI",
		3745: "TANITA",
		4658: "TAP.tv",
		3774: "TASA",
		8311: "TASCAN",
		2442: "TASI",
		8518: "TATUNG",
		4034: "TBTech",
		8595: "TC",
		139:  "TCL",
		3887: "TCM",
		7455: "TCPlink",
		6433: "TCT",
		698:  "TD",
		3146: "TDA",
		6691: "TDC",
		5520: "TDK",
		3511: "TDK-Lambda",
		8842: "TDSi",
		199:  "TDT",
		2160: "TEAL",
		7320: "TECHNART",
		576:  "TECOM",
		7606: "TEKTELIC",
		7401: "TELCO",
		1727: "TELDIX",
		2049: "TELEFIELD",
		1007: "TELEM",
		181:  "TELENET",
		3726: "TELENOT",
		4539: "TEM",
		2519: "TEN",
		8180: "TES",
		1967: "TElectronics",
		6288: "THUB",
		2663: "THX",
		1193: "TIL",
		4920: "TITECH",
		7131: "TITENG",
		193:  "TIW",
		4897: "TKM",
		5337: "TKS",
		6053: "TL",
		7439: "TLS",
		8484: "TM-Research",
		8371: "TMCT",
		7763: "TMRG",
		1313: "TMT",
		1931: "TMT&D",
		9205: "TMY",
		6529: "TNK",
		2373: "TNS",
		892:  "TOA",
		8973: "TOHO",
		6933: "TOP-Access",
		1595: "TP-Link",
		3095: "TPS",
		3642: "TPine",
		2904: "TRENDnet",
		4591: "TRG",
		121:  "TRI-Data",
		8412: "TRIZ",
		6221: "TRL",
		9502: "TRONTEQ",
		9107: "TRP",
		477:  "TRsystems",
		4939: "TSI",
		6346: "TSL",
		9213: "TSMART",
		6964: "TTE",
		1082: "TURCK",
		4451: "TV-Numeric",
		5364: "TV/CoM",
		4605: "TVLogic",
		5474: "TVS",
		3358: "TVT",
		3047: "TVWorks",
		5026: "TWO",
		8504: "TXTR",
		2921: "TZero",
		2108: "Tableau",
		2719: "Tabor",
		4929: "Tachion",
		1376: "Tachyon",
		8224: "Taco",
		943:  "Tactel",
		8481: "Tactical",
		2018: "Tadlys",
		85:   "Tadpole",
		1606: "Taekwang",
		2478: "Taelim",
		3314: "TagMaster",
		8799: "Tagatec",
		1024: "Tahoe",
		4891: "TaiDoc",
		6253: "TaiYear",
		9458: "Taian",
		1676: "Taifatech",
		3317: "Taiguen",
		806:  "Tailyn",
		6826: "Taimag",
		2453: "Taishin",
		1983: "Tait",
		4003: "Taiwick",
		4194: "Taiyo",
		3547: "Takacom",
		736:  "Takagi",
		3920: "Takahata",
		966:  "Takasago",
		584:  "Takaya",
		7216: "Talari",
		65:   "Talaris",
		4102: "Talent",
		8575: "Taleo",
		4052: "Talk-A-Phone",
		2185: "TalkSwitch",
		8252: "Tallac",
		7469: "Talon",
		6135: "Talx",
		8107: "Tamaggo",
		9316: "Tamio",
		8068: "Tamron",
		6593: "Tamtron",
		5914: "Tamura",
		6358: "Tandem",
		8986: "TangoWiFi.com",
		8882: "Tangtop",
		2752: "Tanisys",
		4746: "Tantalus",
		1950: "Tapwave",
		6028: "Taqua",
		1505: "Targa",
		8021: "Tariox",
		4130: "Taseon",
		7040: "Tatsuno",
		5477: "Tatung",
		6199: "Tazmo",
		7574: "Taztag",
		730:  "Tdsoft",
		3523: "Teak",
		7305: "Team",
		4589: "Team-R",
		3371: "Teamcast",
		4159: "Tecan",
		9241: "Tecc",
		8013: "Tech4home",
		4111: "TechNexion",
		9082: "TechSAT",
		3909: "TechTrex",
		6668: "Techaya",
		3544: "Techcity",
		3164: "Techfaithwireless",
		6407: "Techman",
		8259: "Techmation",
		2716: "Techmetro",
		7005: "Technica",
		2924: "Technical",
		6392: "Technicolor",
		6371: "Technity",
		1185: "Techno-Holon",
		7382: "Techno-Innov",
		1757: "Techno-One",
		4772: "TechnoDigital",
		225:  "TechnoLand",
		2344: "Technobox",
		431:  "Technocom",
		6022: "Technologic",
		2784: "Technolution",
		7242: "Technomate",
		7718: "Technonia",
		1377: "Technoventure",
		699:  "Technovision",
		3539: "Technowave",
		667:  "Techsan",
		2934: "Techsphere",
		4281: "Techware",
		4442: "Techway",
		6344: "Tecmar",
		8309: "Tecmobile",
		5513: "Tecnetics",
		6265: "Tecno",
		8253: "Tecnobit",
		6105: "Tecnomen",
		1750: "Tecnova",
		3367: "Tecobest",
		6734: "TeconGroup",
		7156: "Tecore",
		7795: "Tecsen",
		1534: "Tecton",
		6671: "Tegile",
		3665: "Tehuti",
		7229: "Tek-Air",
		5927: "Teklogix",
		5680: "Teknema",
		8748: "Teko",
		4045: "Tekon-Automatics",
		7343: "Tekpea",
		5371: "Tekram",
		3904: "Tekron",
		4782: "Tektrap",
		6303: "Tektronix",
		7641: "Teladin",
		3703: "Telchemyorporated",
		18:   "Telco",
		1779: "TelcoBridges",
		5713: "Teldat",
		8006: "TeleAdapt",
		418:  "TeleCruz",
		436:  "TeleDream",
		4028: "TeleWell",
		1334: "Telebau",
		5186: "Telebit",
		1370: "Telebyte",
		2463: "Telecard-Pribor",
		7075: "Telechips",
		3571: "Teleco",
		1132: "Telecom",
		3570: "Telecomunication",
		6831: "Telecor",
		1336: "Telecore",
		6416: "Telecsys",
		865:  "Telect",
		7271: "Teledata",
		1470: "Teledex",
		3233: "Teledyne",
		8716: "Teleepoch",
		7728: "Telefield",
		274:  "Teleforce",
		4387: "Telegesis",
		680:  "Telelynx",
		5664: "Telemann",
		61:   "Telematics",
		1061: "Telemax",
		981:  "Telemonitor",
		2753: "Telemotive",
		1293: "Telena",
		650:  "Telencomm",
		4366: "Telephonics",
		4115: "Teleplan",
		7091: "Telepower",
		4192: "Teleprocessing",
		4149: "Telerad",
		6219: "Teles",
		5938: "Telesciences",
		4463: "Telesis",
		473:  "Telesoft",
		9024: "Telesquare",
		5615: "Teleste",
		5618: "Telestream",
		5352: "Telesync",
		2500: "Telesynergy",
		7250: "Teletics",
		3677: "Teletrak",
		5012: "Teletrol",
		113:  "Televideo",
		6705: "Teleview",
		8448: "Telewave",
		1409: "Telewise",
		1679: "Telex",
		2425: "Telexy",
		4969: "Telgen",
		3760: "Telian",
		5050: "Telica",
		1762: "Telio",
		5252: "Telkom",
		1537: "Telkonet",
		5027: "Tellabs",
		8540: "Telldus",
		2116: "Tellion",
		1815: "Tellium",
		4307: "Tellord",
		3030: "Tellumat",
		5694: "Tellus",
		6179: "Telmax",
		2407: "Telocityorporated",
		6173: "Telogy",
		6107: "Telrad",
		525:  "Telsey",
		6194: "Telsis",
		17:   "Telsist",
		1810: "Telson",
		5164: "Telspec",
		6787: "Telstra",
		5287: "Telstrat",
		3985: "Teltonika",
		5704: "Teltrend",
		7375: "Teltronic",
		290:  "Teltronics",
		3678: "Telular",
		7257: "Telvent",
		976:  "Telways",
		5320: "Telxon",
		7384: "Tely",
		1809: "Tempearl",
		7225: "Tempered",
		8863: "Tempo",
		2946: "TenX",
		6266: "Tenda",
		3927: "Tendril",
		6494: "Tendyron",
		9426: "Tenebraex",
		3061: "Teneros",
		3851: "Tenlon",
		5798: "Tennyson",
		5023: "Tenor",
		2783: "Tenosys",
		1076: "Tenovis",
		9374: "Tensorcom",
		6725: "Tenstorrent",
		2715: "Tentaculus",
		7935: "Tenyu",
		696:  "Teo",
		1273: "Tepg-US",
		1169: "TeraBurst",
		5039: "TeraForce",
		314:  "TeraGlobal",
		2360: "TeraLogic",
		3501: "TeraMage",
		3016: "TeraRecon",
		4820: "Teracom",
		5092: "Teradata",
		4434: "Teradici",
		1536: "Teradon",
		1663: "Teralink",
		2884: "Teranetics",
		8092: "Teraon",
		2998: "Terascala",
		6678: "Terasic",
		8615: "Teraspek",
		6063: "Teratech",
		4108: "Teraview",
		5052: "Terawave",
		6956: "Teraworks",
		8828: "Tercel",
		6971: "Terewave",
		2735: "Termtek",
		3820: "Terra",
		9380: "TerraSem",
		1560: "TerraTec",
		2551: "Terrasat",
		7568: "Terumo",
		3809: "Tervela",
		7051: "Tescom",
		8256: "Teseq",
		7309: "Tesla",
		1678: "Test-Um",
		1406: "Testech",
		3428: "Testo",
		6250: "Testop",
		2251: "Tevebox",
		1203: "Texa",
		470:  "Texcel",
		313:  "Texio",
		4879: "Thales",
		8655: "Thalmic",
		6091: "Thamway",
		6216: "That",
		7501: "Theben",
		2926: "Thecus",
		8269: "Theobroma",
		7475: "There",
		380:  "Thermalogic",
		2590: "Thermo",
		233:  "ThermoQuest",
		8710: "ThinGlobal",
		6263: "ThinPAD",
		3558: "Thincom",
		2625: "ThingMagic",
		9254: "Things",
		7328: "ThingsMatrix",
		7984: "ThinkEco",
		4563: "ThinkFlood",
		1171: "Thinkengine",
		7205: "Thinking",
		4095: "Thinkware",
		3267: "Thinlinx",
		5458: "Thomas-Conrad",
		9086: "Thomas-Krenn.AG",
		1142: "Thomson",
		6929: "Thread",
		9201: "Throughtek",
		3204: "ThruVision",
		8630: "Thuh",
		8350: "Thum+Mahr",
		9309: "Thundercomm",
		610:  "TiMetra",
		2528: "TiVo",
		7093: "Tiandy",
		5226: "Tiara",
		4659: "Tibbo",
		3258: "Tibetsystem",
		939:  "Tiburon",
		3605: "Tidel",
		5384: "Tidomat",
		6152: "Tiernan",
		1925: "Tiesse",
		506:  "Tietech",
		6314: "Tigan",
		3236: "Tigi",
		6539: "Tiinlab",
		3603: "Tilera",
		377:  "Tilgin",
		9198: "Time-O-Matic",
		4133: "TimeIPS",
		4410: "TimeKeeping",
		5764: "TimeStep",
		29:   "Timeplex",
		2247: "Timespace",
		453:  "Timeware",
		3682: "Tintometer",
		9432: "Tintri",
		1858: "Tiptel",
		3352: "Tiqit",
		9489: "TireCheck",
		738:  "Titan",
		2177: "Tivella",
		2535: "Tixi.com",
		2694: "ToGoldenNet",
		4688: "Tobii",
		7838: "Todaair",
		3159: "Tohken",
		3130: "Toho",
		6835: "Tokheim",
		5318: "Tokimec",
		4970: "Toko",
		1205: "Tokyo",
		5211: "TollBridge",
		2316: "Tollgrade",
		3893: "Tom",
		2714: "TomTom",
		6352: "Tomen",
		387:  "Tommy",
		6490: "Tonal",
		8285: "Tongfang",
		6894: "Tonly",
		3254: "Tonze",
		6900: "Top-Unum",
		4354: "TopControl",
		6374: "Topaz",
		1090: "Topcall",
		8487: "Topcon",
		2080: "Topfield",
		852:  "Topspin",
		1664: "Topview",
		8943: "Topwell",
		2818: "Toradex",
		5815: "Toray",
		4613: "Toro",
		6319: "Torus",
		35:   "Toshiba",
		6506: "Tosibox",
		426:  "Totsu",
		6399: "Totus",
		6813: "Touch",
		921:  "TouchStar",
		6130: "Touchwave",
		7399: "Toyo",
		1262: "Toyo-Linx",
		734:  "Toyokeiki",
		7463: "Tozo",
		2059: "Tpack",
		2364: "Tracewell",
		7447: "TrackNet",
		3088: "Trackflow",
		5468: "Tradpost",
		7314: "TrafficCast",
		2741: "TrafficSim",
		2043: "Trajet",
		6277: "Traka",
		8656: "Trakm8",
		5716: "Trancell",
		2656: "Trane",
		296:  "Trango",
		2054: "TransCore",
		5266: "TransMedia",
		7729: "TransPacket",
		2038: "Transact",
		4306: "Transcon",
		8075: "Transics",
		5957: "Transition",
		5717: "Transitions",
		4485: "Translogic",
		5830: "Transmeta",
		5805: "Transmitton",
		989:  "Transmode",
		1443: "Transtech",
		5138: "Transware",
		5412: "Transys",
		3804: "Transystem",
		7565: "Tranwo",
		3565: "Tranzas",
		1612: "Trapeze",
		9389: "Traqueur",
		4085: "Travelping",
		556:  "Traxit",
		7372: "Treehouse",
		6935: "Treeview",
		7193: "Trek",
		8530: "TrekStor",
		3984: "TrellisWare",
		7106: "Tremol",
		6016: "Tremon",
		176:  "Trend",
		8689: "TrendPoint",
		2341: "Trenton",
		261:  "Trex",
		1058: "Tri-M",
		7804: "Tri-Sen",
		8916: "Tri-Systems",
		8027: "Tri-Tech",
		1695: "TriBeam",
		434:  "TriState",
		6984: "Tributary",
		9030: "Tricascade",
		5710: "Tricord",
		310:  "Tridium",
		8589: "Tridonic",
		5082: "Trigem",
		391:  "Trilithic",
		1585: "Trilliant",
		8233: "Trilobit",
		1312: "Trilogy",
		1375: "Trimble",
		3944: "Trimm",
		9236: "Trinus",
		8578: "Tripwire",
		8194: "Trison",
		5165: "Tritec",
		8361: "Triteka",
		8393: "Triton",
		8493: "Tritonwave",
		6287: "Triumph-Adler",
		4874: "Trixell",
		866:  "Triz",
		5623: "Troika",
		1240: "Tropic",
		6879: "True",
		6400: "Truebeyond",
		3781: "Truen",
		4828: "Truesell",
		5010: "TrunkNet",
		8887: "Trusonus",
		2510: "Trustable",
		1059: "Trutzschler",
		7553: "Trymus",
		7392: "Tsat",
		9470: "Tsingtong",
		3650: "Tsubata",
		5230: "Tundo",
		7367: "TurControlSystme",
		1896: "Turbo",
		3463: "TurboChef",
		4956: "TurboComm",
		884:  "TurboWave",
		5040: "Turbonet",
		9235: "Turbostor",
		265:  "Turin",
		5626: "Turnstone",
		6700: "Turtle",
		5865: "Tutankhamon",
		7425: "Tuttnaer",
		6475: "Tvip",
		3168: "Twig",
		1280: "TwinHan",
		1696: "TwinMOS",
		184:  "Twinhead",
		9025: "Twinlinx",
		7427: "Twpi",
		6158: "Tyan",
		4461: "Tyco",
		7557: "Tzukuri",
		6500: "U&U",
		2532: "U-MEDIA",
		8809: "U-Raku",
		2988: "U-WAY",
		8831: "U2S",
		2996: "U4EA",
		3701: "UBI&MOBI",
		2216: "UBSTORAGE",
		8129: "UCI",
		7505: "UCZOON",
		6573: "UCloud",
		5096: "UDC",
		7649: "UGame",
		1251: "UHD-Elektronik",
		1895: "UHS",
		5103: "UKG",
		8427: "ULTIMEDIA",
		6993: "ULVAC",
		3419: "UMB",
		7812: "UNINET",
		1846: "UNION",
		991:  "UNIQA",
		8301: "UNITEC",
		1508: "UNITEK",
		8573: "URadio",
		4283: "USC",
		1135: "UTStarcom",
		6515: "UTT",
		7002: "UTran",
		2881: "UbONE",
		1832: "UbeeAirWalk",
		3181: "Ubicod",
		8238: "Ubicquia",
		9307: "Ubiik",
		1464: "Ubinetics",
		9355: "Ubiqam",
		3073: "Ubiquam",
		2978: "Ubiquiti",
		6897: "Ubiquitous",
		1103: "Ubiquoss",
		2524: "Ubisense",
		3528: "Ubistar",
		5127: "Ubitrex",
		9480: "Ubivelox",
		3287: "Ubixon",
		7262: "Ubizcore",
		6536: "Ubyon",
		3783: "Ucamp",
		8535: "Ufine",
		9226: "Ufispace",
		528:  "Ulan",
		7414: "Ulterius",
		5616: "Ultimate",
		3429: "Ultra",
		7867: "UltraClenz",
		2077: "Ultracker",
		6011: "Ultrak",
		3715: "Ultratec",
		3732: "Ultratronik",
		2449: "Unatech",
		7907: "Unetconvergence",
		256:  "Unex",
		5847: "Ungermann-Bass",
		5696: "Uni-Link",
		4629: "Uni-v",
		8472: "UniComm",
		493:  "UniData",
		8933: "UniPrint",
		8533: "Uniband",
		6960: "Unicard",
		2509: "Uniclass",
		1080: "Unico",
		6604: "Unicoi",
		5785: "Unicomputer",
		9170: "Unicore",
		5288: "Unicorn",
		6096: "Uniden",
		8016: "Unidis",
		4033: "Unifat",
		5044: "Uniform",
		322:  "Unify",
		4548: "Unigen",
		5922: "Unigraf",
		3278: "Unigrand",
		7937: "Unikey",
		5883: "Unimicro",
		1462: "Unimo",
		3505: "Unionman",
		8055: "Unipattern",
		4304: "Uniphone",
		3600: "Unipoint",
		5369: "Unipulse",
		6414: "Uniqoteq",
		2420: "Unique",
		7563: "Unisem",
		5591: "Unisphere",
		38:   "Unisys",
		9044: "Unit-EM",
		6140: "Unitec",
		3045: "Unitech",
		1826: "United",
		6444: "Unitend",
		3992: "Unitron",
		1900: "Unitronics",
		6325: "Univation",
		3857: "Universal",
		2257: "Uniwell",
		918:  "Uniwide",
		474:  "Uniwill",
		557:  "Unixtar",
		172:  "Unizone",
		8125: "Unmonday",
		6841: "Unowhy",
		4143: "Up-Today",
		4296: "UpdateLogic",
		2165: "Uplogix",
		147:  "Upnod",
		958:  "Upponetti",
		7346: "Uptmate",
		9003: "Upvel",
		8946: "Upwis",
		3017: "Uquest",
		2936: "Uriel",
		3803: "Uriver",
		4047: "Urmet",
		7232: "Uros",
		8001: "Usag",
		6963: "UserGate",
		6567: "Usys",
		8190: "Utek",
		5992: "Utilicom",
		1424: "Utility",
		8514: "Utillink",
		6164: "Utstarcom",
		5910: "Uunet",
		4436: "Uwin",
		7608: "V",
		6071: "V-Bits",
		2642: "V-Show",
		9460: "V-ZUG",
		5813: "V.I",
		9165: "V2",
		1995: "VAC",
		8894: "VAIO",
		3768: "VDG-Security",
		398:  "VDSL",
		6714: "VECOW",
		6381: "VEO-Labs",
		4162: "VIA",
		7200: "VIG",
		8731: "VILLBAU",
		9336: "VIPAR",
		3132: "VK",
		5913: "VMX",
		809:  "VMware",
		9413: "VNL",
		7276: "VOGTEC",
		5842: "VOIM",
		1190: "VOIX",
		6966: "VPI",
		5380: "VPNet",
		3730: "VR",
		4312: "VRmagic",
		5304: "VSK",
		2680: "VSST",
		5965: "VST",
		6332: "VTA",
		8319: "VTC",
		7819: "VTS",
		1319: "VTech",
		8295: "VVDN",
		3391: "VVOND",
		4742: "VXi",
		7880: "Vachen",
		4358: "Vacon",
		2686: "VadaTech",
		8176: "Vadaro",
		9160: "Vaddio",
		5657: "Vadem",
		8050: "Vaillant",
		6015: "Valcom",
		2289: "Valcretec",
		3436: "Valemount",
		9345: "ValenceTech",
		7725: "Valentine",
		4710: "Valiant",
		5663: "Valid",
		6914: "Validus",
		8862: "Valink",
		792:  "Valley",
		8566: "Vallox",
		1535: "Valo",
		2697: "Valox",
		6424: "Valuable",
		2598: "Value",
		9009: "ValueHD",
		2462: "ValuePoint",
		8519: "Valueplus",
		9122: "Valve",
		3608: "Vamp",
		887:  "Vanderbilt",
		3174: "Vansco",
		4018: "Vantanol",
		6630: "Vanu",
		500:  "Vaone",
		5893: "Vari-Lite",
		5374: "Varian",
		8580: "Varikorea",
		9447: "Variscite",
		5472: "Varityper",
		3068: "Vativ",
		3908: "VaultStor",
		1157: "Vbrick",
		3306: "Vecima",
		3664: "Vector",
		2313: "Vectron",
		5443: "Veea",
		8384: "Veedims",
		6685: "Veethree",
		8082: "Velankani",
		8587: "Vello",
		9313: "VeloCloud",
		6285: "Velocytech",
		7538: "Velodyne",
		7596: "Velux",
		7966: "Vence",
		3724: "Vencer",
		3577: "Venergy",
		7730: "Venetex",
		4714: "Venntis",
		2441: "Venstar",
		2952: "Ventus",
		2196: "Veo",
		3829: "VerScient",
		9094: "Verana",
		2707: "Verascape",
		9410: "Verathon",
		2786: "VeriWave",
		8893: "Verifi",
		1647: "Verifone",
		4212: "Verilink",
		1118: "Verint",
		6048: "Veris",
		4457: "Verismo",
		666:  "Veristar",
		6351: "Veritas",
		3320: "Veritech",
		6621: "Verizon",
		9144: "Verkada",
		3379: "Verkerk",
		1616: "Vernier",
		6875: "Veroguard",
		6618: "Veros",
		8413: "Vers",
		4965: "Versa",
		724:  "VersaLogic",
		4273: "Versalynx",
		3929: "Versamed",
		6195: "Versanet",
		2417: "Vertical",
		6955: "Vertu",
		868:  "Verytech",
		4540: "Vestac",
		1449: "Vestel",
		5637: "Vetronix",
		7079: "Vexata",
		2820: "ViPowER",
		7613: "ViVOtech",
		2922: "ViXS",
		393:  "ViaClix",
		5826: "ViaGate",
		3727: "ViaLogy",
		6200: "ViaVideo",
		6488: "Viaanix",
		4715: "Viaas",
		8610: "Viableware",
		8380: "Vialis",
		2634: "Vialta",
		5004: "Vianet",
		5783: "Viasatorporated",
		8426: "Vibicom",
		1381: "Vibration",
		3997: "Vibro-Meter",
		1481: "Vichel",
		5643: "Vickers",
		5324: "Vicom",
		952:  "Vicon",
		3880: "Viconics",
		7899: "Vicos",
		4270: "Victron",
		9239: "VidaBox",
		5455: "Videcom",
		5150: "Video",
		9228: "VideoHome",
		5406: "VideoServer",
		5247: "Videocon",
		361:  "Videoframe",
		5238: "Videojet",
		8658: "Videoswitch",
		4365: "Videotec",
		716:  "Videotek",
		2069: "Videotron",
		1407: "Videx",
		3945: "Vidient",
		1686: "Vidisco",
		464:  "Viditec",
		2295: "Vienna",
		5994: "Vieo",
		9455: "Vievu",
		1614: "ViewSonic",
		3553: "ViewTel",
		8087: "Viewcooper",
		8059: "Vieworks",
		2496: "Viewtran",
		6647: "Viking",
		9492: "Vimar",
		3112: "Vimicro",
		3121: "Vimtron",
		5342: "Vina",
		1930: "Vinci",
		1584: "Vindicator",
		2107: "VineSys",
		8133: "Viogem",
		3693: "Violin",
		4274: "Vipa",
		3366: "Viprinet",
		8004: "Viptela",
		2570: "Virbiage",
		379:  "Virditech",
		6918: "Virgilant",
		8299: "Virident",
		4614: "Virtual",
		7097: "Virtualtek",
		8798: "Virtuozzo",
		8154: "VisSim",
		8906: "Viscount",
		2641: "Viseon",
		356:  "Visicom",
		1644: "Visimetrics",
		2689: "Vision",
		789:  "VisionTek",
		7566: "VisionVera",
		2017: "Visionary",
		5305: "Visioncomm",
		4022: "Visioneering",
		4933: "Visionetics",
		4894: "Visionhitech",
		1434: "Visionics",
		3590: "Visionite",
		6706: "Visionscape",
		1645: "Visiowave",
		1399: "Visteon",
		20:   "Visual",
		8536: "Visualedge",
		3380: "Visualgate",
		2624: "Vita",
		5532: "Vitacom",
		6128: "VitalCom",
		284:  "VitalPoint",
		7054: "VitalThings",
		6349: "Vitalink",
		4763: "Vitality",
		217:  "Vitana",
		7287: "Vitec",
		2894: "Vitelcom",
		2883: "Vitelec",
		8538: "Vitsmo",
		7259: "Vity",
		2529: "Vivaas",
		413:  "Vivace",
		7378: "Vivago",
		8643: "Vivalnk",
		8402: "Vivatel",
		1633: "Vivato",
		2631: "Vivavis",
		5833: "Viveris",
		6062: "Vivid",
		1392: "VividLogic",
		674:  "Vivity",
		6369: "Vivo",
		7543: "Vivonic",
		440:  "Vivotek",
		3454: "Vivox",
		6234: "Viwone",
		1776: "Vixen",
		6565: "Vixtel",
		6759: "Vizeo",
		4752: "Vizimax",
		3467: "Vizio",
		7849: "Vizmonet",
		8660: "Vlatacom",
		5987: "Vlsi",
		5341: "Vmetro",
		8374: "Vnpt",
		1456: "Vocera",
		2815: "Vocollect",
		9062: "Vodia",
		519:  "Vodtel",
		5732: "Voelker",
		5274: "Voiceboard",
		8330: "Voippartners",
		8742: "Voismart",
		6728: "Volacomm",
		2394: "Volamp",
		9428: "Volans",
		6542: "Volex",
		1607: "Volktek",
		3668: "Vololink",
		8759: "Volta",
		1304: "Voltaire",
		9349: "Volterra",
		5773: "Vorax",
		1632: "Vormetric",
		1006: "Vorne",
		6606: "VostroNet",
		3581: "Votronic",
		5984: "Voxent",
		1608: "Voxpath",
		3639: "Voxtel",
		4640: "Voyant",
		889:  "Vrcom",
		8163: "Vsoontech",
		2502: "Vtech",
		1960: "Vtera",
		7639: "Vubiq",
		5436: "Vutrix",
		8300: "Vuzix",
		3660: "W&W",
		1000: "W-Link",
		2167: "W-Linx",
		1548: "W2",
		8898: "WAAV",
		2768: "WACOM",
		8516: "WAMA",
		4593: "WAREMA",
		8747: "WATERWORLD",
		2256: "WAVE",
		5585: "WAVTrace",
		3398: "WB",
		8395: "WBS",
		123:  "WD",
		5002: "WEBGATE",
		3105: "WEBIO",
		7006: "WEG",
		1459: "WELL",
		1327: "WEM",
		2806: "WEPIO",
		9384: "WEY",
		4840: "WFE",
		201:  "WIN",
		292:  "WINCOMM",
		4394: "WIRECOM",
		1159: "WIS",
		1899: "WISCORE",
		7420: "WISEWARE",
		2846: "WJ",
		7160: "WKK",
		5777: "WMS",
		8236: "WOM",
		5059: "WONWOO",
		1705: "WOOJU",
		7627: "WOORI",
		7570: "WOORISYSTEMS",
		8030: "WOXTER",
		6377: "WSH",
		2106: "WTSS",
		7186: "Wacom",
		9293: "Wafa",
		1525: "Walchem",
		9296: "Waldo",
		5624: "WalkAbout",
		8188: "Wally",
		136:  "Wang",
		2960: "Wanshih",
		1306: "Wany",
		9490: "Wapice",
		8324: "Warehouse",
		2176: "Wasabi",
		8234: "Wata",
		175:  "WatchGuard",
		7526: "Water-i.d",
		7676: "WaterFurnace",
		2006: "Watertek",
		559:  "Watlow",
		7696: "Wattty",
		3560: "Wave",
		5361: "WaveAccess",
		6235: "WaveIP",
		7473: "WaveLynx",
		624:  "WaveSmith",
		5399: "WaveSpan",
		2151: "WaveSplitter",
		3766: "WaveStorm",
		9347: "Wavelink",
		5314: "Wavenet",
		2042: "Waveplus",
		5669: "Waverider",
		2761: "Wavesat",
		598:  "Wavesight",
		8582: "Wavetel",
		3874: "Wavetrend",
		6694: "Wavetronix",
		8702: "WayTools",
		264:  "Wayport",
		6290: "We",
		4398: "WeLink",
		8623: "WeTelecom",
		8965: "Wearable",
		6698: "Wearhaus",
		7120: "Wearsafe",
		3136: "Weathernews",
		6231: "WebSilicon",
		2359: "WebSonic",
		6217: "WebSprint",
		6103: "WebTV",
		1471: "WebWayOne",
		7784: "Weber-Stephen",
		6153: "Webgear",
		2471: "Webpro",
		12:   "Webster",
		5396: "Webtronics",
		890:  "Webyn",
		1112: "Wegener",
		7583: "Weidu",
		544:  "Weinschel",
		1767: "Weintek",
		4656: "Weinzierl",
		3785: "Weiss",
		3403: "Welcat",
		3348: "Weldex",
		3387: "Welding",
		7270: "Welgate",
		8375: "Wellav",
		7793: "Wellcore",
		257:  "Welltech",
		4193: "Welltronix",
		9112: "Wellysis",
		7882: "Welotec",
		3564: "Weltec",
		4353: "Wescon",
		9332: "Wesine",
		6768: "Westcontrol",
		2273: "Westell",
		5823: "Westport",
		7659: "Westunitis",
		1010: "Westwave",
		1623: "Wetek",
		8335: "Whaley",
		4556: "Whdi",
		750:  "WhereNet",
		8268: "WhereWhen",
		8105: "Whirlpool",
		5351: "Whistle",
		5674: "Whitecross",
		8795: "WhizNets",
		2966: "Wi-Gear",
		3987: "Wi-Links",
		8568: "Wi-NEXT",
		3457: "Wi2Wi",
		7656: "Wi3",
		6985: "WiBotic",
		3825: "WiChorus",
		3771: "WiDeFi",
		7716: "WiFiSONG",
		7033: "WiFiSong",
		2629: "WiLife",
		4136: "WiMate",
		3475: "WiQuest",
		4235: "WiSE",
		8227: "WiSilica",
		4110: "WiWide",
		7177: "WiZ",
		7478: "Wiatec",
		3955: "Wibrain",
		1469: "Widax",
		518:  "Widcomm",
		4632: "Wide",
		2557: "WideRay",
		2223: "WideView",
		9031: "Wideband",
		2540: "Wideful",
		7118: "Widex",
		8360: "Wieson",
		7771: "Wifi-soft",
		5819: "WigWag",
		8165: "Wiio",
		2860: "Wiline",
		7873: "Will-Burt",
		2389: "Willnet",
		5660: "Willowbrook",
		787:  "Willowglen",
		6278: "Wilocity",
		6171: "Wiltron",
		3822: "Win4NET",
		967:  "WinCom",
		3338: "WinNet",
		1789: "Winbest",
		2302: "Winbond",
		6740: "Wincal",
		4323: "Winchester",
		5067: "Windata",
		1198: "Windigo",
		9379: "WindowMaster",
		8353: "Winduskorea",
		3175: "Winegard",
		8237: "Winfirm",
		4030: "Wingtech",
		3404: "Winix",
		2893: "Wink",
		3780: "Winland",
		593:  "Winmate",
		1952: "Winners",
		1574: "Winnow",
		941:  "Winpresa",
		7983: "Winstars",
		196:  "Winsystems",
		2207: "Wintec",
		3037: "Wintecronics",
		8963: "Wintop",
		6019: "Wintriss",
		4692: "Winward",
		3956: "Winy",
		1909: "Wiplug",
		1235: "Wipotec",
		7130: "WiredIQ",
		231:  "Wireless",
		7092: "WirelessWERX",
		705:  "Wirelink",
		3328: "Wiremold",
		7791: "Wirepas",
		8388: "Wiscloud",
		4198: "Wisdm",
		3853: "Wiseblue",
		547:  "Wisi",
		7235: "Wisol",
		1592: "Wistron",
		2396: "Witcom",
		4112: "Witelcom",
		1282: "With-Net",
		4700: "Withings",
		8186: "Withrobot",
		4505: "Witron",
		4729: "Wittenstein",
		6990: "WizLAN",
		7371: "Wizardlab",
		8583: "Wizitdongdo",
		2886: "Wizlogics",
		1292: "Wiznet",
		3629: "Wizyoung",
		8592: "Woan",
		8645: "Wohler",
		1824: "WolfVision",
		6367: "WondaLink",
		8765: "WonderSound",
		7609: "Wonderlan",
		4830: "Wongs",
		3272: "WooJooIT",
		7264: "Woodstream",
		6600: "Woojin",
		3058: "Woojinnet",
		1206: "Wooksung",
		401:  "Woorigisool",
		5000: "Workbit",
		441:  "Workstation",
		2159: "WorldAccxx",
		5608: "WorldCast",
		371:  "WorldGate",
		8659: "Worldplay",
		3623: "Woven",
		3324: "WowWee",
		9064: "Wush",
		7334: "WyTec",
		1808: "Wybron",
		8488: "Wyconn",
		849:  "Wyle",
		8855: "Wyler",
		4704: "Wynmax",
		9193: "WyreStorm",
		5495: "Wyse",
		6954: "Wytek",
		6881: "Wyze",
		1677: "X-CoM",
		198:  "X-traWeb",
		4379: "X2E",
		7788: "X6D",
		8644: "XADA",
		8416: "XAG",
		190:  "XAVi",
		4950: "XCP",
		5294: "XEL",
		3207: "XENOLINK",
		339:  "XESystems",
		6280: "XIAOMI",
		3353: "XIP",
		5800: "XKL",
		4023: "XLN-t",
		6002: "XN",
		6843: "XOS",
		4900: "XRPLUS",
		3298: "XStreamHD",
		8079: "XTA",
		5556: "XTP",
		8452: "XTS",
		2284: "XTecorporated",
		8727: "XTrillion",
		3481: "XYnetsoft",
		4226: "Xact",
		4916: "Xagyl",
		1002: "Xalted",
		2175: "Xalyo",
		2169: "Xambala",
		3065: "Xanboo",
		5906: "Xante",
		3019: "Xantech",
		7998: "Xapt",
		9208: "Xaptec",
		8579: "Xaptum",
		6178: "Xaqti",
		5366: "Xata",
		9311: "Xcellen",
		7391: "Xcheng",
		3143: "Xcute",
		5905: "Xedia",
		1825: "Xeline",
		4550: "Xembedded",
		3702: "XenICs",
		6293: "Xena",
		4139: "Xenatech",
		7021: "Xenox",
		3075: "Xensource",
		0:    "Xerox",
		7800: "Xetawave",
		9322: "Xi3",
		2115: "XiNCOM",
		5697: "Xiaomi",
		1498: "Xilinx",
		6477: "Ximea",
		5464: "Xinetron",
		7209: "Xingfei",
		8954: "Xingluo",
		4120: "Xiotech",
		6083: "Xiox",
		3448: "Xipher",
		2979: "Xiranet",
		788:  "Xircom",
		2214: "Xirrus",
		5276: "Xitron",
		7865: "Xlab",
		4576: "Xmark",
		6233: "Xmit",
		773:  "Xnet",
		8990: "XonTel",
		7580: "Xorcom",
		3931: "Xortec",
		5437: "Xovis",
		486:  "Xpeed",
		5759: "Xpoint",
		5191: "Xrite",
		5983: "Xrosstech",
		561:  "Xsense",
		1496: "Xsido",
		2357: "Xstreamis",
		3069: "Xteam",
		755:  "Xtera",
		4460: "Xtramus",
		8383: "Xtreme",
		1756: "XtremeSpectrum",
		6585: "Xtremio",
		9015: "Xunison",
		8757: "XySystem",
		5308: "Xycom",
		798:  "Xycotec",
		5469: "Xylogics",
		5769: "Xyplex",
		5519: "Xyron",
		6357: "Xyvision",
		8146: "Y-cam",
		1126: "Y.D.K",
		693:  "YAFO",
		9352: "YAMABISHI",
		2695: "YDT",
		3598: "YEC",
		1452: "YEM",
		1308: "YESTECHNOLOGY",
		8707: "YF",
		7485: "YIK",
		3633: "YMC",
		7683: "YODO",
		2105: "YOKO",
		1371: "YOZAN",
		7927: "YSI",
		8231: "YST",
		7450: "YSTen",
		6766: "YUKAI",
		8700: "YWire",
		727:  "Yamaha",
		5189: "Yamashita",
		4169: "Yamatake-Honeywell",
		8670: "Yandex",
		2782: "Yangjae",
		7722: "Yaojin",
		9248: "Yaptv",
		1697: "Yasing",
		2272: "Yazaki",
		9400: "Yeelink",
		3654: "Yi-Qing",
		7325: "Yichen",
		3381: "Yiguang",
		942:  "Yipee",
		6686: "Ynomia",
		8934: "Yo",
		1335: "Yoda",
		8964: "Yoga",
		4908: "Yoisys",
		7605: "Yokota",
		5917: "Yoshiki",
		2516: "Yoshimiya",
		3245: "Yosin",
		842:  "Yotta",
		663:  "YottaYotta",
		9334: "Yottabyte",
		1074: "Young",
		2000: "Youngbo",
		7805: "Youngkook",
		4337: "Yournet",
		9242: "Ypsomed",
		6832: "Yuduan",
		1889: "Yuehua",
		2620: "Yulinet",
		3096: "Yulong",
		3606: "Yupiteru",
		1095: "Yuxing",
		3074: "Yves",
		7504: "Yytek",
		5388: "Z-CoM",
		3445: "Z-Com",
		8714: "Z-TEC",
		7532: "Z-meta",
		7149: "Z3",
		5588: "ZAC",
		567:  "ZACCESS",
		3890: "ZANTAZ",
		974:  "ZARDCOM",
		4262: "ZAX",
		3958: "ZEFATEK",
		9223: "ZETLAB",
		2877: "ZHIYUAN",
		7113: "ZIV",
		6282: "ZNV",
		2477: "ZNYX",
		1285: "ZOOM",
		4487: "ZORT",
		9394: "ZOYI",
		3837: "ZP",
		9164: "ZPE",
		1572: "ZPSYS",
		3031: "ZTE",
		9434: "ZVISION",
		6134: "ZX",
		5025: "Zaffire",
		4370: "Zala",
		6796: "Zalliant",
		1573: "Zambeel",
		7396: "Zaplox",
		7703: "Zappware",
		3740: "Zavio",
		2699: "Zcomax",
		5947: "Zcomm",
		5611: "Zeal",
		768:  "Zebra",
		3206: "Zed-3",
		3838: "ZeeVee",
		9066: "Zeebo",
		4470: "Zeeport",
		7841: "Zegna-Daidong",
		7506: "Zektor",
		3573: "Zelax",
		7403: "Zelfy",
		8673: "Zencheer",
		9418: "Zengge",
		1800: "Zenith",
		3231: "Zenitron",
		6273: "Zenner",
		843:  "Zenocom",
		6615: "Zenotech",
		6434: "Zenovia",
		3977: "Zensys",
		8303: "Zentan",
		4859: "Zenterio",
		7280: "Zentri",
		3099: "Zenway",
		7286: "Zeo",
		6995: "Zeppelin",
		9079: "Zera",
		8876: "Zero1.tv",
		6474: "ZeroDesktop",
		4195: "Zeta",
		913:  "Zetari",
		2899: "Zetec",
		499:  "Zetes",
		1613: "Zetron",
		2647: "Zetta",
		9288: "ZettaHash",
		3446: "Zeugma",
		3497: "Zeutschel",
		9020: "Zexelon",
		8966: "Zhehua",
		197:  "Zhone",
		8720: "ZhongMiao",
		7670: "Zhuolian",
		1063: "Zi",
		5484: "Ziatech",
		8174: "Zicon",
		2421: "Zida",
		4783: "ZillionTV",
		5595: "Zilog",
		7736: "Zimi",
		8976: "Zimory",
		3212: "Zinwave",
		841:  "Zinwell",
		2127: "Zioncom",
		2174: "Zipher",
		3354: "Zippy",
		8090: "Zitte",
		6247: "Zmodo",
		3332: "Zodianet",
		3772: "Zoltan",
		780:  "Zoltrix",
		7632: "Zonar",
		5202: "Zonet",
		5993: "Zoneworx",
		8475: "Zonoff",
		6687: "Zoovel",
		2104: "Zoran",
		7011: "Ztron",
		1736: "Zultys",
		333:  "Zuma",
		7449: "ZuniData",
		7868: "Zurn",
		9396: "ZyCast",
		1527: "ZyFLEX",
		438:  "ZyGate",
		2692: "ZyXEL",
		7654: "Zycoo",
		3821: "Zygo",
		5778: "Zykronix",
		3551: "Zylaya",
		4523: "Zylin",
		5589: "Zypcom",
		8572: "Zyptonite",
		7127: "aFUN",
		6876: "acromate",
		2051: "activ-net",
		3343: "ads-tec",
		8478: "aizo",
		1516: "almedio",
		4787: "altek",
		6949: "amazipoint",
		9204: "amnimo",
		8808: "ants",
		8785: "arkona",
		7779: "as",
		392:  "ask-technologies.com",
		3471: "asteel",
		6523: "automationNEXT",
		3461: "avinfo",
		7978: "azeti",
		7677: "b-plus",
		8498: "ball-b",
		6509: "base",
		6852: "bct",
		6697: "bebro",
		9061: "beroNet",
		4469: "beyerdynamic",
		2909: "bio-logic",
		6833: "bioMerieux",
		2124: "bitWallet",
		9383: "blackned",
		8410: "bluesky",
		1938: "boca",
		1631: "bplan",
		4300: "byd:sign",
		7904: "c-scape",
		4898: "cTrixs",
		6248: "camtron",
		1723: "cd3o",
		8109: "cellon",
		7918: "chaowifi.com",
		2145: "cim-usa",
		7782: "citygrow",
		6568: "currentoptronics",
		3572: "cyber-blue",
		651:  "cyberPIXIE",
		1886: "cybernet",
		8128: "d-broad",
		8743: "d2d",
		7588: "dSPACE",
		3327: "dSys",
		9124: "data-Complex",
		2913: "datacom",
		2596: "daum",
		7379: "deister",
		1637: "devolo",
		8115: "di-soric",
		4116: "digEcor",
		8553: "digitron",
		6109: "dit",
		427:  "dotRocket",
		4317: "dresden-elektronik",
		7390: "duagon",
		6141: "e-Net",
		2460: "e-SMARTCOM",
		8709: "e-Smart",
		5193: "e-Tek",
		1123: "e-Watch",
		1417: "e-generis",
		2254: "e-w/you",
		2491: "e-zy.net",
		3294: "e2v",
		3057: "e:cue",
		1156: "eCopilt",
		4010: "eCopy",
		459:  "eDevice",
		9403: "eGauge",
		9047: "eMegatech",
		3102: "eOn",
		1558: "ePipe",
		877:  "eProduction",
		7663: "eSSys",
		4409: "eSang",
		1446: "eSpace",
		7370: "eWBM",
		2527: "eWerks",
		2866: "eXS",
		8850: "eZEX",
		8097: "easynetworks",
		7182: "ecobee",
		5820: "eero",
		988:  "egnite",
		3248: "elab-experience",
		1363: "elmegt",
		3751: "emtrion",
		7652: "enLighted",
		6815: "enblink",
		9404: "endeavour",
		8399: "enimai",
		5603: "ens",
		4471: "epro",
		7727: "essensys",
		7598: "evon",
		8091: "exands",
		6817: "exlar",
		6659: "eyevis",
		1385: "fSONA",
		8417: "fenglian",
		6713: "feno",
		8033: "ffly4u",
		3582: "fitivision",
		7753: "fos4X",
		8453: "frogblue",
		7982: "ghe-ces",
		7661: "grandcentrix",
		7043: "home2net",
		6651: "i3",
		1440: "iAd",
		5229: "iBAHN",
		1054: "iCanTek",
		3199: "iCatch",
		8522: "iConnectivity",
		3704: "iControl",
		7448: "iD",
		8994: "iDevices",
		167:  "iDigm",
		5338: "iDirect",
		6461: "iDruide",
		6962: "iFORCOM",
		7085: "iKey",
		6397: "iKuai",
		1256: "iLogic",
		3265: "iMCA-GmbH",
		589:  "iMPath",
		1437: "iMaxNetworksLimited",
		3761: "iNEWiT",
		2135: "iPAC",
		3609: "iPOX",
		3748: "iPhotonix",
		658:  "iPolicy",
		9093: "iPort",
		2447: "iPulse",
		1339: "iQstor",
		2552: "iQuest",
		9135: "iRay",
		3103: "iRex",
		7298: "iRobot",
		6757: "iRule",
		9232: "iS5",
		6589: "iSonea",
		2560: "iStor",
		6069: "iSystem",
		4065: "iTAS",
		1317: "iTEC",
		4340: "iVeia",
		4496: "iWDL",
		4620: "iWOW",
		2681: "iWise",
		3459: "iXSea",
		8829: "iZotope",
		1661: "ib-mohnen",
		3009: "iba",
		2201: "icube",
		3109: "id-Confirm",
		7827: "ifm",
		7888: "iiNet",
		3188: "iiTron",
		2057: "in2",
		8479: "inMotion",
		286:  "inXtron",
		6764: "innodisk",
		8106: "innomdlelab",
		2627: "intec",
		5718: "interWAVE",
		4750: "interbro",
		3699: "intotech",
		8975: "ioBridge",
		2743: "ioIMAGE",
		7402: "iodyne",
		2064: "ionSign",
		630:  "ipDialog",
		678:  "ipUnplugged",
		1395: "ipcas",
		7481: "iryx",
		7536: "isepos",
		7143: "itelio",
		4774: "itron",
		7499: "ittim",
		6264: "iway",
		8212: "jinyoung",
		1878: "jp-embedded",
		4786: "kasercorp",
		2801: "kk-electronic",
		3692: "l-acoustics",
		6617: "lemonbeat",
		1018: "lesswire",
		8662: "lignex1",
		8935: "littleBits",
		1088: "m-u-t",
		3179: "m2c",
		2887: "mCubelogics",
		6580: "mLogic",
		7857: "mMax",
		4930: "mPHASE",
		6885: "machineQ",
		4417: "maintech",
		288:  "manroland",
		9178: "messMa",
		6730: "miControl",
		2933: "mingjong",
		6534: "mirusystems",
		5574: "mixi",
		2914: "mm-lab",
		4870: "moblic",
		8184: "modas",
		6257: "moobox",
		5592: "mps",
		8812: "myIDkey",
		6814: "myenergi",
		3202: "nFore",
		8987: "nSTREAMS",
		1708: "nStor",
		2583: "nVent",
		8201: "netKTI",
		4806: "netTALK.com",
		2430: "netplat",
		2286: "nex-G",
		7550: "nextLAP",
		6422: "noax",
		7668: "nuoxc",
		6563: "nyantec",
		7705: "o2ones",
		8648: "optilink",
		7994: "oshkosh",
		6467: "pomdevices",
		8037: "profichip",
		4841: "r2p",
		6762: "razberi",
		9225: "sTraffic",
		4596: "saxnet",
		6818: "seca",
		8502: "silergy",
		8020: "silex",
		8210: "silicom",
		8420: "smart-electronic",
		8738: "smartAC.com",
		625:  "snom",
		5054: "so-logic",
		1643: "sofrel",
		8073: "storONE",
		6909: "streamnow",
		1816: "synertronixx",
		7405: "t-mac",
		9278: "tado",
		8139: "taskit",
		8544: "tci",
		8725: "tdvine",
		3012: "technicob",
		7211: "tide",
		8249: "trivum",
		7419: "u-blox",
		2550: "u10",
		7398: "uAvionix",
		3315: "uControl",
		4160: "ubisys",
		3569: "ubtos",
		7207: "udworks",
		8179: "unGlue",
		4412: "uv-electronic",
		8031: "vArmour",
		9157: "vastriver",
		3291: "w5networks",
		6646: "wanze",
		7335: "waytotec",
		8501: "wi-daq",
		3229: "wisembed",
		2543: "woori-net",
		9089: "worldcns",
		6292: "x-fabric",
		4761: "x-star",
		9041: "x.o.ware",
		4107: "xG",
		9417: "xn",
		8781: "xvtec",
		977:  "yLez",
	},
}
