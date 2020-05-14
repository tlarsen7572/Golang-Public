package recordinfo_test

import (
	"github.com/tlarsen7572/Golang-Public/goalteryx/recordinfo"
	"testing"
)

func TestInstantiateRecordInfoFromXml(t *testing.T) {
	recordInfoXml := `<MetaInfo connection="Output">
		<RecordInfo>
		<Field name="Field1" source="TextInput:" type="Int64"/>
		<Field name="Field3" size="30" source="TextInput:" type="V_WString"/>
		<Field name="Field2" size="10" source="TextInput:" type="String"/>
		<Field name="Field4" size="1073741823" source="Formula: &apos;la la la&apos;" type="V_WString"/>
		<Field name="Field5" size="1073741823" source="Formula: PadRight(&apos;&apos;, 500, &apos;1&apos;)" type="V_WString"/>
		</RecordInfo>
	</MetaInfo>
	`

	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if count := recordInfo.NumFields(); count != 5 {
		t.Fatalf(`expecpted 5 fields but got %v`, count)
	}
	if field, _ := recordInfo.GetFieldByIndex(0); field.Name != `Field1` {
		t.Fatalf(`expected 'Field1' but got '%v'`, field.Name)
	}
}
