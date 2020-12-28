package smart

import (
	// "bytes"
	"context"
	"encoding/base64"
	// "encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	// "unsafe"

	"go-test/initial"
)

// var (
// 	NativeEndian binary.ByteOrder
// )

// // Determine native endianness of system
// func init() {
// 	i := uint32(1)
// 	b := (*[4]byte)(unsafe.Pointer(&i))
// 	if b[0] == 1 {
// 		NativeEndian = binary.LittleEndian
// 	} else {
// 		NativeEndian = binary.BigEndian
// 	}
// }

// Individual SMART attribute (12 bytes)
type smartAttr struct {
	Id          uint8
	Flags       uint16
	Value       uint8   // normalised value
	Worst       uint8   // worst value
	VendorBytes [6]byte // vendor-specific (and sometimes device-specific) data
	Reserved    uint8
}

type finalSmartAttr struct {
	Id       uint8  `json: "id" bson:"id"`
	Name     string `json: "name" bson:"name"`
	Flags    uint16 `json: "flag" bson:"flag"`
	Value    uint8  `json: "value" bson:"value"` // normalised value
	Worst    uint8  `json: "worst" bson:"worst"` // worst value
	Reserved uint8  `json: "reserved" bson:"reserved"`
	RawValue string `json: "rawValue" bson:"rawValue"` // raw data
	Type     string `json: "type" bson:"type"`
	Updated  string `json: "updated" bson:"updated"`
}

type FinalSmartInfo struct {
	DiskMD  string           `json: "disk_md" bson:"disk_md"`
	DiskSN  string           `json: "disk_sn" bson:"disk_sn"`
	Version uint16           `json: "version" bson:"version"`
	Attrs   []finalSmartAttr `json: "attrs" bson:"attrs"`
}

// Page of 30 SMART attributes as per ATA spec
type SmartPage struct {
	Version uint16
	Attrs   []smartAttr
}

// decodeVendorBytes decodes the six-byte vendor byte array based on the conversion rule passed as
// conv. The conversion may also include the reserved byte, normalised value or worst value byte.
func (sa *smartAttr) decodeVendorBytes(conv string) uint64 {
	var (
		byteOrder string
		r         uint64
	)

	// Default byte orders if not otherwise specified in drivedb
	switch conv {
	case "raw64", "hex64":
		byteOrder = "543210wv"
	case "raw56", "hex56", "raw24/raw32", "msec24hour32":
		byteOrder = "r543210"
	default:
		byteOrder = "543210"
	}

	// Pick bytes from smartAttr in order specified by byteOrder
	for _, i := range byteOrder {
		var b byte

		switch i {
		case '0', '1', '2', '3', '4', '5':
			b = sa.VendorBytes[i-48]
		case 'r':
			b = sa.Reserved
		case 'v':
			b = sa.Value
		case 'w':
			b = sa.Worst
		default:
			b = 0
		}

		r <<= 8
		r |= uint64(b)
	}

	fmt.Printf("value==========:%v\n", r)

	return r
}

func (s *SmartPage) GetSmartPage(msmart map[byte][]byte) {
	for k, v := range msmart {
		sa := smartAttr{}
		sa.Id = k
		sa.Flags = uint16(v[1])
		sa.Value = v[3]
		sa.Worst = v[4]
		sa.VendorBytes = [6]byte{v[5], v[6], v[7], v[8], v[9], v[10]}
		sa.Reserved = v[11]
		s.Attrs = append(s.Attrs, sa)
	}
	sort.Slice(s.Attrs, func(i, j int) bool {
		return s.Attrs[i].Id < s.Attrs[j].Id
	})
}

func checkTempRange(t int8, ut1, ut2 uint8, lo, hi *int) bool {
	t1, t2 := int8(ut1), int8(ut2)

	if t1 > t2 {
		t1, t2 = t2, t1
	}

	if (-60 <= t1) && (t1 <= t) && (t <= t2) && (t2 <= 120) && !(t1 == -1 && t2 <= 0) {
		*lo, *hi = int(t1), int(t2)
		return true
	}

	return false
}

func checkTempWord(word uint16) int {
	switch {
	case word <= 0x7f:
		return 0x11 // >= 0, signed byte or word
	case word <= 0xff:
		return 0x01 // < 0, signed byte
	case word > 0xff80:
		return 0x10 // < 0, signed word
	default:
		return 0x00
	}
}

func formatRawValue(v uint64, conv string) (s string) {
	var (
		raw  [6]uint8
		word [3]uint16
	)

	// Split into bytes
	for i := 0; i < 6; i++ {
		raw[i] = uint8(v >> uint(i*8))
	}

	// Split into words
	for i := 0; i < 3; i++ {
		word[i] = uint16(v >> uint(i*16))
	}

	switch conv {
	case "raw8":
		s = fmt.Sprintf("%d %d %d %d %d %d",
			raw[5], raw[4], raw[3], raw[2], raw[1], raw[0])
	case "raw16":
		s = fmt.Sprintf("%d %d %d", word[2], word[1], word[0])
	case "raw48", "raw56", "raw64":
		s = fmt.Sprintf("%d", v)
	case "hex48":
		s = fmt.Sprintf("%#012x", v)
	case "hex56":
		s = fmt.Sprintf("%#014x", v)
	case "hex64":
		s = fmt.Sprintf("%#016x", v)
	case "raw16(raw16)":
		s = fmt.Sprintf("%d", word[0])
		if (word[1] != 0) || (word[2] != 0) {
			s += fmt.Sprintf(" (%d %d)", word[2], word[1])
		}
	case "raw16(avg16)":
		s = fmt.Sprintf("%d", word[0])
		if word[1] != 0 {
			s += fmt.Sprintf(" (Average %d)", word[1])
		}
	case "raw24(raw8)":
		s = fmt.Sprintf("%d", v&0x00ffffff)
		if (raw[3] != 0) || (raw[4] != 0) || (raw[5] != 0) {
			s += fmt.Sprintf(" (%d %d %d)", raw[5], raw[4], raw[3])
		}
	case "raw24/raw24":
		s = fmt.Sprintf("%d/%d", v>>24, v&0x00ffffff)
	case "raw24/raw32":
		s = fmt.Sprintf("%d/%d", v>>32, v&0xffffffff)
	case "min2hour":
		// minutes
		minutes := uint64(word[0]) + uint64(word[1])<<16
		s = fmt.Sprintf("%dh+%02dm", minutes/60, minutes%60)
		if word[2] != 0 {
			s += fmt.Sprintf(" (%d)", word[2])
		}
	case "sec2hour":
		// seconds
		hours := v / 3600
		minutes := (v % 3600) / 60
		seconds := v % 60
		s = fmt.Sprintf("%dh+%02dm+%02ds", hours, minutes, seconds)
	case "halfmin2hour":
		// 30-second counter
		hours := v / 120
		minutes := (v % 120) / 2
		s = fmt.Sprintf("%dh+%02dm", hours, minutes)
	case "msec24hour32":
		// hours + milliseconds
		hours := v & 0xffffffff
		milliseconds := v >> 32
		seconds := milliseconds / 1000
		s = fmt.Sprintf("%dh+%02dm+%02d.%03ds",
			hours, seconds/60, seconds%60, milliseconds)
	case "tempminmax":
		var tFormat, lo, hi int

		t := int8(raw[0])
		ctw0 := checkTempWord(word[0])

		if word[2] == 0 {
			if (word[1] == 0) && (ctw0 != 0) {
				// 00 00 00 00 xx TT
				tFormat = 0
			} else if (ctw0 != 0) && checkTempRange(t, raw[2], raw[3], &lo, &hi) {
				// 00 00 HL LH xx TT
				tFormat = 1
			} else if (raw[3] == 0) && checkTempRange(t, raw[1], raw[2], &lo, &hi) {
				// 00 00 00 HL LH TT
				tFormat = 2
			} else {
				tFormat = -1
			}
		} else if ctw0 != 0 {
			if (ctw0&checkTempWord(word[1])&checkTempWord(word[2]) != 0x00) && checkTempRange(t, raw[2], raw[4], &lo, &hi) {
				// xx HL xx LH xx TT
				tFormat = 3
			} else if (word[2] < 0x7fff) && checkTempRange(t, raw[2], raw[3], &lo, &hi) && (hi >= 40) {
				// CC CC HL LH xx TT
				tFormat = 4
			} else {
				tFormat = -2
			}
		} else {
			tFormat = -3
		}

		switch tFormat {
		case 0:
			s = fmt.Sprintf("%d", t)
		case 1, 2, 3:
			s = fmt.Sprintf("%d (Min/Max %d/%d)", t, lo, hi)
		case 4:
			s = fmt.Sprintf("%d (Min/Max %d/%d #%d)", t, lo, hi, word[2])
		default:
			s = fmt.Sprintf("%d (%d %d %d %d %d)",
				raw[0], raw[5], raw[4], raw[3], raw[2], raw[1])
		}
	case "temp10x":
		// ten times temperature in Celsius
		s = fmt.Sprintf("%d.%d", word[0]/10, word[0]%10)
	default:
		s = "?"
	}

	fmt.Printf("format value============:%s\n", s)
	return s
}

// GetSmartMapByBytes get smart info
func GetSmartMapByBytes(smartBytes []byte) map[byte][]byte {
	var attrBytes []byte
	attrMap := make(map[byte][]byte)
	for i, v := range smartBytes {
		attrBytes = append(attrBytes, v)
		if (i+1)%12 == 0 {
			// fmt.Println(attrBytes)
			attrMap[attrBytes[0]] = attrBytes
			attrBytes = []byte{}
		}
	}
	delete(attrMap, 0)
	fmt.Println(attrMap)
	return attrMap
}

func getAttrBytesByStr() []byte {
	vendorSpecific := "1,0,5,50,0,100,100,0,0,0,0,0,0,0,9,50,0,100,100,188,18,0,0,0,0,0,12,50,0,100,100,66,3,0,0,0,0,0,165,50,0,100,100,131,5,0,0,0,0,0,166,50,0,100,100,7,0,0,0,0,0,0,167,50,0,100,100,0,0,0,0,0,0,0,168,50,0,100,100,21,0,0,0,0,0,0,169,50,0,100,100,172,0,0,0,0,0,0,170,50,0,100,100,0,0,0,0,0,0,0,171,50,0,100,100,0,0,0,0,0,0,0,172,50,0,100,100,0,0,0,0,0,0,0,173,50,0,100,100,7,0,0,0,0,0,0,174,50,0,100,100,6,0,0,0,0,0,0,184,50,0,100,100,0,0,0,0,0,0,0,187,50,0,100,100,9,0,0,0,0,0,0,188,50,0,100,100,0,0,0,0,0,0,0,194,34,0,55,56,45,0,18,0,56,0,0,199,50,0,100,100,0,0,0,0,0,0,0,230,50,0,100,100,3,4,40,1,3,4,0,232,51,0,100,100,100,0,0,0,0,0,4,233,50,0,100,100,212,6,0,0,0,0,0,234,50,0,100,100,4,42,0,0,0,0,0,241,48,0,100,100,233,14,0,0,0,0,0,242,48,0,100,100,107,8,0,0,0,0,0,244,50,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0"
	vendorBytes := strings.Split(vendorSpecific, ",")
	smartBytes := vendorBytes[:362]
	var attrBytes []byte
	for _, a := range smartBytes {
		val, _ := strconv.Atoi(a)
		attrBytes = append(attrBytes, byte(val))
	}
	return attrBytes
}

func getAttrBytesByBase64() []byte {
	str := "EQKPyQEAAQEBAQABAQEAAAAAAAAAAGoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPbuuwcHAQAAAAAAAAAAAAAAAAAAAAAAAQEBAQABAQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAGU9EBAAAAAAAAAAAAAAAAAAAAAAAAAQEBAQABAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALzebnkOAAAAAAAAAAAAAAAAAAAAAAAAHwAyMDE2MjkAAFDDAABJDAAAwCcJAEkMAAADAIEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(decodeBytes)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(decodeBytes))
	attrBytes := decodeBytes[:362]
	return attrBytes
}

func getAttrBytes() []byte {
	attrBytes := []byte{10, 0, 1, 15, 0, 83, 64, 254, 190, 199, 10, 0, 0, 0, 3, 3, 0, 92, 92, 0, 0, 0, 0, 0, 0, 0, 4, 50, 0, 98, 98, 171, 9, 0, 0, 0, 0, 0, 5, 51, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 7, 15, 0, 78, 60, 98, 150, 157, 13, 4, 0, 0, 9, 50, 0, 71, 71, 210, 99, 0, 0, 0, 0, 0, 10, 19, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 12, 50, 0, 98, 98, 170, 9, 0, 0, 0, 0, 0, 184, 50, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 187, 50, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 188, 50, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 189, 58, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 190, 34, 0, 81, 54, 19, 0, 16, 27, 0, 0, 0, 191, 50, 0, 100, 100, 138, 0, 0, 0, 0, 0, 0, 192, 50, 0, 1, 1, 75, 157, 7, 0, 0, 0, 0, 193, 50, 0, 1, 1, 199, 158, 7, 0, 0, 0, 0, 194, 34, 0, 19, 46, 19, 0, 0, 0, 16, 0, 0, 195, 26, 0, 1, 1, 254, 190, 199, 10, 0, 0, 0, 197, 18, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 198, 16, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 199, 62, 0, 200, 200, 0, 0, 0, 0, 0, 0, 0, 240, 0, 0, 100, 253, 204, 97, 0, 0, 204, 102, 54, 241, 0, 0, 100, 253, 246, 170, 214, 82, 6, 0, 0, 242, 0, 0, 100, 253, 73, 177, 141, 139, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 130, 0, 105, 2, 0, 123, 3, 0, 1, 0, 1, 255, 2, 78, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 9, 8, 9, 8, 8, 8, 8, 8, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 138, 0, 0, 0, 182, 83, 81, 111, 171, 83, 0, 0, 0, 0, 0, 0, 1, 0, 239, 58, 246, 170, 214, 82, 6, 0, 0, 0, 73, 177, 141, 139, 1, 0, 0, 0, 0, 0, 0, 0, 224, 26, 30, 13, 0, 0, 0, 0, 0, 0, 0, 0, 154, 37, 0, 0, 9, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 24, 0, 0, 0, 250}
	return attrBytes[:362]
}

//GetSmartFromJSON receive a json byte,parse to smart struct
func GetSmartFromJSON(byteJson []byte) error {
	jsonMap := make(map[string]string)
	if err := json.Unmarshal(byteJson, &jsonMap); err != nil {
		fmt.Println(err)
		return err
	}
	diskSN := jsonMap["disk_sn"]
	diskMD := jsonMap["disk_md"]
	smartStr := jsonMap["details"]
	decodeBytes, err := base64.StdEncoding.DecodeString(smartStr)
	fmt.Println(decodeBytes)
	if err != nil {
		log.Fatalln(err)
	}
	version := uint16(decodeBytes[0])

	finalSmartInfo := parseSmartAttr(decodeBytes[2:362])
	finalSmartInfo.Version = version
	finalSmartInfo.DiskMD = diskMD
	finalSmartInfo.DiskSN = diskSN

	jsonResult, err := json.Marshal(finalSmartInfo)
	fmt.Println(string(jsonResult))

	err = insertSmart2Mongo(finalSmartInfo)

	return err
}

func insertSmart2Mongo(finalSmartInfo FinalSmartInfo) error {
	collection := initial.IsddcMongoDb.Collection("smart")
	_, err := collection.InsertOne(context.Background(), finalSmartInfo)
	if err != nil {
		fmt.Printf("%s:%s insert failed", finalSmartInfo.DiskMD, finalSmartInfo.DiskSN)
		return err
	}
	fmt.Printf("%s:%s insert successful", finalSmartInfo.DiskMD, finalSmartInfo.DiskSN)

	return nil
}

func parseSmartAttr(attrBytes []byte) FinalSmartInfo {
	smartMap := GetSmartMapByBytes(attrBytes)

	db, err := OpenDriveDb("./smart/drivedb.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// smartMap := GetSmartMapByBytes(attrBytes)
	// attrMap, _ := GetAttrMap("./smart/attribute.yaml")
	smart := SmartPage{}
	// binary.Read(bytes.NewBuffer(attrBytes), binary.LittleEndian, &smart)
	smart.GetSmartPage(smartMap)

	// fmt.Printf("ID# ATTRIBUTE_NAME           FLAG     VALUE WORST THRESHOLD TYPE     UPDATED RAW_VALUE\n")

	finalSmartInfo := FinalSmartInfo{}
	// finalSmartInfo.Version = uint16(attrBytes[0])
	for _, attr := range smart.Attrs {
		var (
			rawValue              uint64
			conv                  AttrConv
			attrType, attrUpdated string
		)

		if attr.Id == 0 {
			break
		}

		conv, ok := db.Drive.Presets[strconv.Itoa(int(attr.Id))]
		// fmt.Printf("drive=====:%v\n", drive)
		if ok {
			// fmt.Println("ok!!!!!!!!!!!!!!!!!!!!!!!")
			rawValue = attr.decodeVendorBytes(conv.Conv)
		}

		// fmt.Printf("raw value==========:%v\n", rawValue)

		// Pre-fail / advisory bit
		if attr.Flags&0x0001 != 0 {
			attrType = "Pre-fail"
		} else {
			attrType = "Old_age"
		}

		// Online data collection bit
		if attr.Flags&0x0002 != 0 {
			attrUpdated = "Always"
		} else {
			attrUpdated = "Offline"
		}

		rawValueStr := formatRawValue(rawValue, conv.Conv)

		finalSmart := finalSmartAttr{}
		finalSmart.Id = attr.Id
		finalSmart.Name = conv.Name
		finalSmart.Flags = attr.Flags
		finalSmart.Value = attr.Value
		finalSmart.Worst = attr.Worst
		finalSmart.Reserved = attr.Reserved
		finalSmart.Type = attrType
		finalSmart.Updated = attrUpdated
		finalSmart.RawValue = rawValueStr

		finalSmartInfo.Attrs = append(finalSmartInfo.Attrs, finalSmart)

		// fmt.Fprintf(w, "%3d %-24s %#04x   %03d   %03d   %03d      %-8s %-7s %s\n",
		// 	attr.Id, conv.Name, attr.Flags, attr.Value, attr.Worst, attr.Reserved, attrType,
		// 	attrUpdated, formatRawValue(rawValue, conv.Conv))
	}

	return finalSmartInfo
}
