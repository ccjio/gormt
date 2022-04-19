package cnf

// EImportsHead imports head options. import包含选项
var EImportsHead = map[string]string{
	"stirng":          `"string"`,
	"time.Time":       `"time"`,
	"gorm.Model":      `"gorm.io/gorm"`,
	"fmt":             `"fmt"`,
	"datatypes.JSON":  `"gorm.io/datatypes"`,
	"datatypes.Date":  `"gorm.io/datatypes"`,
	"decimal.Decimal": `"github.com/shopspring/decimal"`,
}

// TypeMysqlDicMp Accurate matching type.精确匹配类型
var TypeMysqlDicMp = map[string]string{
	"smallint":            "int",
	"smallint unsigned":   "int16",
	"int":                 "int",
	"int unsigned":        "int",
	"bigint":              "int",
	"bigint unsigned":     "int",
	"varchar":             "string",
	"char":                "string",
	"date":                "datatypes.Date",
	"datetime":            "time.Time",
	"bit(1)":              "[]uint8",
	"tinyint":             "int",
	"tinyint unsigned":    "int",
	"tinyint(1)":          "int", // tinyint(1) 默认设置成bool
	"tinyint(1) unsigned": "int", // tinyint(1) 默认设置成bool
	"json":                "datatypes.JSON",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"double unsigned":     "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"float unsigned":      "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
	"integer":             "int64",
}

// TypeMysqlMatchList Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchList = []struct {
	Key   string
	Value string
}{
	{`^(tinyint)[(]\d+[)] unsigned`, "int"},
	{`^(smallint)[(]\d+[)] unsigned`, "int"},
	{`^(int)[(]\d+[)] unsigned`, "int"},
	{`^(bigint)[(]\d+[)] unsigned`, "uint"},
	{`^(float)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(double)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(tinyint)[(]\d+[)]`, "int"},
	{`^(smallint)[(]\d+[)]`, "int"},
	{`^(int)[(]\d+[)]`, "int"},
	{`^(bigint)[(]\d+[)]`, "int"},
	{`^(char)[(]\d+[)]`, "string"},
	{`^(enum)[(](.)+[)]`, "string"},
	{`^(varchar)[(]\d+[)]`, "string"},
	{`^(varbinary)[(]\d+[)]`, "[]byte"},
	{`^(blob)[(]\d+[)]`, "[]byte"},
	{`^(binary)[(]\d+[)]`, "[]byte"},
	{`^(decimal)[(]\d+,\d+[)]`, "decimal.Decimal"},
	{`^(mediumint)[(]\d+[)]`, "string"},
	{`^(double)[(]\d+,\d+[)]`, "float64"},
	{`^(float)[(]\d+,\d+[)]`, "float64"},
	{`^(datetime)[(]\d+[)]`, "time.Time"},
	{`^(bit)[(]\d+[)]`, "[]uint8"},
	{`^(text)[(]\d+[)]`, "string"},
	{`^(integer)[(]\d+[)]`, "int"},
}
