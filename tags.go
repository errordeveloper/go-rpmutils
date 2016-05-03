/*
 * Copyright (c) SAS Institute, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpmutils

const (
	_GENERAL_TAG_BASE = 1000
	NAME              = 1000
	VERSION           = 1001
	RELEASE           = 1002
	EPOCH             = 1003
	SUMMARY           = 1004
	DESCRIPTION       = 1005
	VENDOR            = 1011
	LICENSE           = 1014
	SOURCE            = 1018
	ARCH              = 1022
	PREIN             = 1023
	POSTIN            = 1024
	PREUN             = 1025
	POSTUN            = 1026
	OLDFILENAMES      = 1027
	FILESIZES         = 1028
	FILEMODES         = 1030
	FILERDEVS         = 1033
	FILEMTIMES        = 1034
	FILEDIGESTS       = 1035 // AKA FILEMD5S
	FILELINKTOS       = 1036
	FILEFLAGS         = 1037 // bitmask: RPMFILE_* are bitmasks to interpret
	FILEUSERNAME      = 1039
	FILEGROUPNAME     = 1040
	SOURCERPM         = 1044
	FILEVERIFYFLAGS   = 1045 // bitmask: RPMVERIFY_* are bitmasks to interpret
	PROVIDENAME       = 1047
	REQUIRENAME       = 1049
	REQUIREVERSION    = 1050
	RPMVERSION        = 1064
	TRIGGERSCRIPTS    = 1065
	TRIGGERNAME       = 1066
	TRIGGERVERSION    = 1067
	TRIGGERFLAGS      = 1068 // bitmask: RPMSENSE_* are bitmasks to interpret
	TRIGGERINDEX      = 1069
	VERIFYSCRIPT      = 1079
	PREINPROG         = 1085
	POSTINPROG        = 1086
	PREUNPROG         = 1087
	POSTUNPROG        = 1088
	OBSOLETENAME      = 1090
	PROVIDEVERSION    = 1113
	OBSOLETEFLAGS     = 1114
	OBSOLETEVERSION   = 1115
	VERIFYSCRIPTPROG  = 1091
	TRIGGERSCRIPTPROG = 1092
	DIRINDEXES        = 1116
	BASENAMES         = 1117
	DIRNAMES          = 1118
	PAYLOADFORMAT     = 1124
	PAYLOADCOMPRESSOR = 1125
	FILECOLORS        = 1140
	// BLINK*, FLINK*, and TRIGGERPREIN included from SUSE fork of RPM
	BLINKPKGID   = 1164
	BLINKHDRID   = 1165
	BLINKNEVRA   = 1166
	FLINKPKGID   = 1167
	FLINKHDRID   = 1168
	FLINKNEVRA   = 1169
	TRIGGERPREIN = 1170

	SIG_BASE = 256
	SIG_SHA1 = 269

	// Given that there is overlap between signature tag headers and general tag
	// headers, we offset the signature ones by some amount
	_SIGHEADER_TAG_BASE = 16384
	SIG_SIZE            = _SIGHEADER_TAG_BASE + 1000 // Header + Payload size
	SIG_MD5             = _SIGHEADER_TAG_BASE + 1004 // MD5SUM of header + payload
	SIG_GPG             = _SIGHEADER_TAG_BASE + 1005

	// FILEFLAGS bitmask elements:
	RPMFILE_NONE      = 0
	RPMFILE_CONFIG    = 1 << 0
	RPMFILE_DOC       = 1 << 1
	RPMFILE_ICON      = 1 << 2
	RPMFILE_MISSINGOK = 1 << 3
	RPMFILE_NOREPLACE = 1 << 4
	RPMFILE_SPECFILE  = 1 << 5
	RPMFILE_GHOST     = 1 << 6
	RPMFILE_LICENSE   = 1 << 7
	RPMFILE_README    = 1 << 8
	RPMFILE_EXCLUDE   = 1 << 9
	RPMFILE_UNPATCHED = 1 << 10
	RPMFILE_PUBKEY    = 1 << 11
	RPMFILE_POLICY    = 1 << 12

	// FILEVERIFYFLAGS bitmask elements:
	RPMVERIFY_NONE       = 0
	RPMVERIFY_MD5        = 1 << 0
	RPMVERIFY_FILEDIGEST = 1 << 0
	RPMVERIFY_FILESIZE   = 1 << 1
	RPMVERIFY_LINKTO     = 1 << 2
	RPMVERIFY_USER       = 1 << 3
	RPMVERIFY_GROUP      = 1 << 4
	RPMVERIFY_MTIME      = 1 << 5
	RPMVERIFY_MODE       = 1 << 6
	RPMVERIFY_RDEV       = 1 << 7
	RPMVERIFY_CAPS       = 1 << 8
	RPMVERIFY_CONTEXTS   = 1 << 15

	// TRIGGERFLAGS bitmask elements -- not all rpmsenseFlags make sense
	// in TRIGGERFLAGS
	RPMSENSE_ANY           = 0
	RPMSENSE_LESS          = 1 << 1
	RPMSENSE_GREATER       = 1 << 2
	RPMSENSE_EQUAL         = 1 << 3
	RPMSENSE_TRIGGERIN     = 1 << 16
	RPMSENSE_TRIGGERUN     = 1 << 17
	RPMSENSE_TRIGGERPOSTUN = 1 << 18
	RPMSENSE_TRIGGERPREIN  = 1 << 25
)
