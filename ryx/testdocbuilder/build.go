package testdocbuilder

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func RebuildTestdocs(baseFolder string) {
	checkErr(os.RemoveAll(baseFolder))
	checkErr(os.Mkdir(baseFolder, 0777))
	checkErr(os.Mkdir(filepath.Join(baseFolder, `macros`), 0777))
	checkErr(ioutil.WriteFile(filepath.Join(baseFolder, `MultiInOut.yxmc`), []byte(multiInOutMacro), 0644))
	checkErr(ioutil.WriteFile(filepath.Join(baseFolder, `MultiInOut.yxmd`), []byte(multiInOutWorkflow), 0644))
	checkErr(ioutil.WriteFile(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`), []byte(setleafYxmd), 0644))
	checkErr(ioutil.WriteFile(filepath.Join(baseFolder, `Calculate Filter Expression.yxmc`), []byte(calcFilterExpression), 0644))
	checkErr(ioutil.WriteFile(filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`), []byte(tagWithSets), 0644))
	checkErr(ioutil.WriteFile(filepath.Join(baseFolder, `New Workflow 1.yxwz`), []byte(emptyYxwz), 0644))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var setleafYxmd = `<?xml version="1.0"?>
<AlteryxDocument yxmdVer="2019.3">
  <Nodes>
    <Node ToolID="1">
      <GuiSettings Plugin="AlteryxBasePluginsGui.TextInput.TextInput">
        <Position x="54" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <NumRows value="13" />
          <Fields>
            <Field name="MANDT" />
            <Field name="SETCLASS" />
            <Field name="SUBCLASS" />
            <Field name="SETNAME" />
            <Field name="LINEID" />
            <Field name="VALSIGN" />
            <Field name="VALOPTION" />
            <Field name="VALFROM" />
            <Field name="VALTO" />
            <Field name="SEQNR" />
          </Fields>
          <Data>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST01</c>
              <c>00001</c>
              <c>I</c>
              <c>EQ</c>
              <c>A1</c>
              <c />
              <c>1</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST01</c>
              <c>00002</c>
              <c>E</c>
              <c>EQ</c>
              <c>E1</c>
              <c />
              <c>2</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST01</c>
              <c>00003</c>
              <c>I</c>
              <c>BT</c>
              <c>C0</c>
              <c>L9</c>
              <c>3</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST01</c>
              <c>00004</c>
              <c>E</c>
              <c>BT</c>
              <c>H0</c>
              <c>J9</c>
              <c>4</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST02</c>
              <c>00001</c>
              <c>I</c>
              <c>LE</c>
              <c>D1</c>
              <c />
              <c>1</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST02</c>
              <c>00002</c>
              <c>I</c>
              <c>GT</c>
              <c>W1</c>
              <c />
              <c>2</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST03</c>
              <c>00001</c>
              <c>I</c>
              <c>GE</c>
              <c>X1</c>
              <c />
              <c>1</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST03</c>
              <c>00002</c>
              <c>I</c>
              <c>LT</c>
              <c>C1</c>
              <c />
              <c>2</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST04</c>
              <c>00001</c>
              <c>I</c>
              <c>CP</c>
              <c>A*</c>
              <c />
              <c>1</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST04</c>
              <c>00002</c>
              <c>E</c>
              <c>CP</c>
              <c>+3</c>
              <c />
              <c>2</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST05</c>
              <c>00001</c>
              <c>I</c>
              <c>NE</c>
              <c>A1</c>
              <c />
              <c>1</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST06</c>
              <c>00001</c>
              <c>I</c>
              <c>NB</c>
              <c>A4</c>
              <c>Z3</c>
              <c>1</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST07</c>
              <c>00001</c>
              <c>I</c>
              <c>NP</c>
              <c>+1</c>
              <c />
              <c>1</c>
            </r>
          </Data>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>SETLEAF</AnnotationText>
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxTextInput" />
    </Node>
    <Node ToolID="4">
      <GuiSettings Plugin="AlteryxBasePluginsGui.TextInput.TextInput">
        <Position x="54" y="138" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <NumRows value="7" />
          <Fields>
            <Field name="MANDT" />
            <Field name="SETCLASS" />
            <Field name="SUBCLASS" />
            <Field name="SETNAME" />
            <Field name="FIELDNAME" />
          </Fields>
          <Data>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST01</c>
              <c>RACCT</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST02</c>
              <c>RACCT</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST03</c>
              <c>RACCT</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST04</c>
              <c>RACCT</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST05</c>
              <c>RACCT</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST06</c>
              <c>RACCT</c>
            </r>
            <r>
              <c>000</c>
              <c>0000</c>
              <c><![CDATA[ ]]></c>
              <c>TEST07</c>
              <c>RACCT</c>
            </r>
          </Data>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>SETHEADER</AnnotationText>
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxTextInput" />
    </Node>
    <Node ToolID="5">
      <GuiSettings Plugin="AlteryxBasePluginsGui.TextInput.TextInput">
        <Position x="54" y="222" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <NumRows value="11" />
          <Fields>
            <Field name="RCLNT" />
            <Field name="RYEAR" />
            <Field name="DOCNR" />
            <Field name="RLDNR" />
            <Field name="RBUKRS" />
            <Field name="DOCLN" />
            <Field name="RACCT" />
            <Field name="HSL" />
          </Fields>
          <Data>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000001</c>
              <c>0L</c>
              <c>0001</c>
              <c>000001</c>
              <c>A1</c>
              <c>12.00</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000001</c>
              <c>0L</c>
              <c>0001</c>
              <c>000002</c>
              <c>G7</c>
              <c>-12.00</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000002</c>
              <c>0L</c>
              <c>0001</c>
              <c>000001</c>
              <c>A3</c>
              <c>22.10</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000002</c>
              <c>0L</c>
              <c>0001</c>
              <c>000002</c>
              <c>A8</c>
              <c>1.30</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000002</c>
              <c>0L</c>
              <c>0001</c>
              <c>000003</c>
              <c>B2</c>
              <c>-23.40</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000003</c>
              <c>0L</c>
              <c>0001</c>
              <c>000001</c>
              <c>C2</c>
              <c>14.20</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000003</c>
              <c>0L</c>
              <c>0001</c>
              <c>000002</c>
              <c>I4</c>
              <c>-14.20</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000004</c>
              <c>0L</c>
              <c>0001</c>
              <c>000001</c>
              <c>J1</c>
              <c>10.20</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000004</c>
              <c>0L</c>
              <c>0001</c>
              <c>000002</c>
              <c>Z8</c>
              <c>3.50</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000004</c>
              <c>0L</c>
              <c>0001</c>
              <c>000003</c>
              <c>K9</c>
              <c>-5.00</c>
            </r>
            <r>
              <c>010</c>
              <c>2019</c>
              <c>0000000004</c>
              <c>0L</c>
              <c>0001</c>
              <c>000004</c>
              <c>M3</c>
              <c>-8.70</c>
            </r>
          </Data>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>FAGLFLEXA</AnnotationText>
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxTextInput" />
    </Node>
    <Node ToolID="6">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Join.Join">
        <Position x="150" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration joinByRecordPos="False">
          <JoinInfo connection="Left">
            <Field field="MANDT" />
            <Field field="SETCLASS" />
            <Field field="SUBCLASS" />
            <Field field="SETNAME" />
          </JoinInfo>
          <JoinInfo connection="Right">
            <Field field="MANDT" />
            <Field field="SETCLASS" />
            <Field field="SUBCLASS" />
            <Field field="SETNAME" />
          </JoinInfo>
          <SelectConfiguration>
            <Configuration outputConnection="Join">
              <OrderChanged value="False" />
              <CommaDecimal value="False" />
              <SelectFields>
                <SelectField field="Right_MANDT" selected="False" rename="Right_MANDT" input="Right_" />
                <SelectField field="Right_SETCLASS" selected="False" rename="Right_SETCLASS" input="Right_" />
                <SelectField field="Right_SUBCLASS" selected="False" rename="Right_SUBCLASS" input="Right_" />
                <SelectField field="Right_SETNAME" selected="False" rename="Right_SETNAME" input="Right_" />
                <SelectField field="*Unknown" selected="True" />
              </SelectFields>
            </Configuration>
          </SelectConfiguration>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxJoin" />
    </Node>
    <Node ToolID="12">
      <GuiSettings>
        <Position x="246" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration />
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings Macro="Calculate Filter Expression.yxmc" />
    </Node>
    <Node ToolID="13">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Filter.Filter">
        <Position x="342" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Expression>[VALSIGN] = "I"</Expression>
          <Mode>Simple</Mode>
          <Simple>
            <Operator>=</Operator>
            <Field>VALSIGN</Field>
            <Operands>
              <IgnoreTimeInDateTime>True</IgnoreTimeInDateTime>
              <DateType>fixed</DateType>
              <PeriodDate>2019-10-18 05:18:54</PeriodDate>
              <PeriodType>
              </PeriodType>
              <PeriodCount>0</PeriodCount>
              <StartDate>2019-10-18 05:18:54</StartDate>
              <EndDate>2019-10-18 05:18:54</EndDate>
              <Operand>I</Operand>
            </Operands>
          </Simple>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>Split I / E</AnnotationText>
          <DefaultAnnotationText>[VALSIGN] = "I"</DefaultAnnotationText>
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFilter" />
    </Node>
    <Node ToolID="14">
      <GuiSettings Plugin="AlteryxSpatialPluginsGui.Summarize.Summarize">
        <Position x="438" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <SummarizeFields>
            <SummarizeField field="MANDT" action="GroupBy" rename="MANDT" />
            <SummarizeField field="SETCLASS" action="GroupBy" rename="SETCLASS" />
            <SummarizeField field="SUBCLASS" action="GroupBy" rename="SUBCLASS" />
            <SummarizeField field="SETNAME" action="GroupBy" rename="SETNAME" />
            <SummarizeField field="Filter" action="Concat" rename="Filter">
              <Concat_Start>(</Concat_Start>
              <Separator><![CDATA[ || ]]></Separator>
              <Concat_End>)</Concat_End>
            </SummarizeField>
          </SummarizeFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxSpatialPluginsEngine.dll" EngineDllEntryPoint="AlteryxSummarize" />
    </Node>
    <Node ToolID="15">
      <GuiSettings Plugin="AlteryxSpatialPluginsGui.Summarize.Summarize">
        <Position x="438" y="138" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <SummarizeFields>
            <SummarizeField field="MANDT" action="GroupBy" rename="MANDT" />
            <SummarizeField field="SETCLASS" action="GroupBy" rename="SETCLASS" />
            <SummarizeField field="SUBCLASS" action="GroupBy" rename="SUBCLASS" />
            <SummarizeField field="SETNAME" action="GroupBy" rename="SETNAME" />
            <SummarizeField field="Filter" action="Concat" rename="Filter">
              <Concat_Start>(</Concat_Start>
              <Separator><![CDATA[ && ]]></Separator>
              <Concat_End>)</Concat_End>
            </SummarizeField>
          </SummarizeFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxSpatialPluginsEngine.dll" EngineDllEntryPoint="AlteryxSummarize" />
    </Node>
    <Node ToolID="16">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Union.Union">
        <Position x="534" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <ByName_ErrorMode>Error</ByName_ErrorMode>
          <ByName_OutputMode>All</ByName_OutputMode>
          <Mode>ByName</Mode>
          <SetOutputOrder value="False" />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxUnion" />
    </Node>
    <Node ToolID="17">
      <GuiSettings Plugin="AlteryxSpatialPluginsGui.Summarize.Summarize">
        <Position x="630" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <SummarizeFields>
            <SummarizeField field="MANDT" action="GroupBy" rename="MANDT" />
            <SummarizeField field="SETCLASS" action="GroupBy" rename="SETCLASS" />
            <SummarizeField field="SUBCLASS" action="GroupBy" rename="SUBCLASS" />
            <SummarizeField field="SETNAME" action="GroupBy" rename="SETNAME" />
            <SummarizeField field="Filter" action="Concat" rename="Filter">
              <Concat_Start>(</Concat_Start>
              <Separator><![CDATA[ && ]]></Separator>
              <Concat_End>)</Concat_End>
            </SummarizeField>
          </SummarizeFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxSpatialPluginsEngine.dll" EngineDllEntryPoint="AlteryxSummarize" />
    </Node>
    <Node ToolID="18">
      <GuiSettings>
        <Position x="726" y="222" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Value name="BatchMacroGroupBy" />
          <Value name="ControlParams"><![CDATA[Filter=Filter
SetName=SETNAME
]]></Value>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="Output7">
          <RecordInfo>
            <Field name="RCLNT" size="3" source="TextInput: (BatchMacro) (BatchMacro)" type="String" />
            <Field name="RYEAR" source="TextInput: (BatchMacro) (BatchMacro)" type="Int16" />
            <Field name="DOCNR" size="10" source="TextInput: (BatchMacro) (BatchMacro)" type="String" />
            <Field name="RLDNR" size="2" source="TextInput: (BatchMacro) (BatchMacro)" type="String" />
            <Field name="RBUKRS" size="4" source="TextInput: (BatchMacro) (BatchMacro)" type="String" />
            <Field name="DOCLN" size="6" source="TextInput: (BatchMacro) (BatchMacro)" type="String" />
            <Field name="RACCT" size="2" source="TextInput: (BatchMacro) (BatchMacro)" type="String" />
            <Field name="HSL" source="TextInput: (BatchMacro) (BatchMacro)" type="Double" />
            <Field name="Set Name" size="1073741823" source="Formula: &quot;TEST01&quot; (BatchMacro) (BatchMacro)" type="V_WString" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings Macro="macros\Tag with Sets.yxmc" />
    </Node>
    <Node ToolID="19">
      <GuiSettings Plugin="AlteryxBasePluginsGui.BrowseV2.BrowseV2">
        <Position x="822" y="222" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <TempFile>C:\Users\tlarsen\AppData\Local\Temp\Engine_3296_30e9198a94364271a3393dc5a4d3ed7f_\Engine_9220_c2e1dcdb71aa4c65a5a2fb3ae9ee9921_.yxdb</TempFile>
          <TempFileDataProfiling />
          <Layout>
            <ViewMode>Single</ViewMode>
            <ViewSize value="100" />
            <View1>
              <DefaultTab>Profile</DefaultTab>
              <Hints>
                <Table />
              </Hints>
            </View1>
            <View2 />
          </Layout>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxBrowseV2" />
    </Node>
    <Node ToolID="20">
      <GuiSettings Plugin="AlteryxGuiToolkit.ToolContainer.ToolContainer">
        <Position x="269" y="292" width="405.0632" height="188" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Caption>Container 20</Caption>
          <Style TextColor="#314c4a" FillColor="#ecf2f2" BorderColor="#314c4a" Transparency="25" Margin="25" />
          <Disabled value="False" />
          <Folded value="False" />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <ChildNodes>
        <Node ToolID="21">
          <GuiSettings Plugin="AlteryxGuiToolkit.TextBox.TextBox">
            <Position x="294" y="378" width="100" height="40" />
          </GuiSettings>
          <Properties>
            <Configuration>
              <Text>One level in</Text>
              <Font name="Arial" size="8.25" style="0" />
              <TextColor name="Black" />
              <FillColor name="White" />
              <Shape shape="0" />
              <Justification Justification="4" />
            </Configuration>
            <Annotation DisplayMode="0">
              <Name />
              <DefaultAnnotationText />
              <Left value="False" />
            </Annotation>
          </Properties>
        </Node>
        <Node ToolID="22">
          <GuiSettings Plugin="AlteryxGuiToolkit.ToolContainer.ToolContainer">
            <Position x="497" y="341" width="152.0632" height="114" />
          </GuiSettings>
          <Properties>
            <Configuration>
              <Caption>Container 22</Caption>
              <Style TextColor="#314c4a" FillColor="#ecf2f2" BorderColor="#314c4a" Transparency="25" Margin="25" />
              <Disabled value="False" />
              <Folded value="False" />
            </Configuration>
            <Annotation DisplayMode="0">
              <Name />
              <DefaultAnnotationText />
              <Left value="False" />
            </Annotation>
          </Properties>
          <ChildNodes>
            <Node ToolID="23">
              <GuiSettings Plugin="AlteryxGuiToolkit.TextBox.TextBox">
                <Position x="522" y="390" width="100" height="40" />
              </GuiSettings>
              <Properties>
                <Configuration>
                  <Text>Two levels in</Text>
                  <Font name="Arial" size="8.25" style="0" />
                  <TextColor name="Black" />
                  <FillColor name="White" />
                  <Shape shape="0" />
                  <Justification Justification="4" />
                </Configuration>
                <Annotation DisplayMode="0">
                  <Name />
                  <DefaultAnnotationText />
                  <Left value="False" />
                </Annotation>
              </Properties>
            </Node>
          </ChildNodes>
        </Node>
      </ChildNodes>
    </Node>
  </Nodes>
  <Connections>
    <Connection>
      <Origin ToolID="1" Connection="Output" />
      <Destination ToolID="6" Connection="Left" />
    </Connection>
    <Connection>
      <Origin ToolID="4" Connection="Output" />
      <Destination ToolID="6" Connection="Right" />
    </Connection>
    <Connection>
      <Origin ToolID="5" Connection="Output" />
      <Destination ToolID="18" Connection="Input1" />
    </Connection>
    <Connection>
      <Origin ToolID="17" Connection="Output" />
      <Destination ToolID="18" Connection="Control" />
    </Connection>
    <Connection>
      <Origin ToolID="6" Connection="Join" />
      <Destination ToolID="12" Connection="Input1" />
    </Connection>
    <Connection>
      <Origin ToolID="12" Connection="Output" />
      <Destination ToolID="13" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="13" Connection="True" />
      <Destination ToolID="14" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="13" Connection="False" />
      <Destination ToolID="15" Connection="Input" />
    </Connection>
    <Connection name="#1">
      <Origin ToolID="14" Connection="Output" />
      <Destination ToolID="16" Connection="Input" />
    </Connection>
    <Connection name="#2">
      <Origin ToolID="15" Connection="Output" />
      <Destination ToolID="16" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="16" Connection="Output" />
      <Destination ToolID="17" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="18" Connection="Output7" />
      <Destination ToolID="19" Connection="Input" />
    </Connection>
  </Connections>
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name>01 SETLEAF Equations Completed</Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="True" />
    </Events>
  </Properties>
</AlteryxDocument>`

var calcFilterExpression = `<?xml version="1.0"?>
<AlteryxDocument yxmdVer="2019.3">
  <Nodes>
    <Node ToolID="1">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroInput.MacroInput">
        <Position x="102" y="126" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <UseFileInput value="False" />
          <Name>Input1</Name>
          <Abbrev />
          <ShowFieldMap value="False" />
          <Optional value="False" />
          <TextInput>
            <Configuration>
              <NumRows value="13" />
              <Fields>
                <Field name="MANDT" />
                <Field name="SETCLASS" />
                <Field name="SUBCLASS" />
                <Field name="SETNAME" />
                <Field name="LINEID" />
                <Field name="VALSIGN" />
                <Field name="VALOPTION" />
                <Field name="VALFROM" />
                <Field name="VALTO" />
                <Field name="SEQNR" />
                <Field name="FIELDNAME" />
              </Fields>
              <Data>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST01</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>EQ</c>
                  <c>A1</c>
                  <c />
                  <c>1</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST01</c>
                  <c>00002</c>
                  <c>E</c>
                  <c>EQ</c>
                  <c>E1</c>
                  <c />
                  <c>2</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST01</c>
                  <c>00003</c>
                  <c>I</c>
                  <c>BT</c>
                  <c>C0</c>
                  <c>L9</c>
                  <c>3</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST01</c>
                  <c>00004</c>
                  <c>E</c>
                  <c>BT</c>
                  <c>H0</c>
                  <c>J9</c>
                  <c>4</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST02</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>LE</c>
                  <c>D1</c>
                  <c />
                  <c>1</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST02</c>
                  <c>00002</c>
                  <c>I</c>
                  <c>GT</c>
                  <c>W1</c>
                  <c />
                  <c>2</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST03</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>GE</c>
                  <c>X1</c>
                  <c />
                  <c>1</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST03</c>
                  <c>00002</c>
                  <c>I</c>
                  <c>LT</c>
                  <c>C1</c>
                  <c />
                  <c>2</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST04</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>CP</c>
                  <c>A*</c>
                  <c />
                  <c>1</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST04</c>
                  <c>00002</c>
                  <c>E</c>
                  <c>CP</c>
                  <c>+3</c>
                  <c />
                  <c>2</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST05</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>NE</c>
                  <c>A1</c>
                  <c />
                  <c>1</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST06</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>NB</c>
                  <c>A4</c>
                  <c>Z3</c>
                  <c>1</c>
                  <c>RACCT</c>
                </r>
                <r>
                  <c>000</c>
                  <c>0000</c>
                  <c><![CDATA[ ]]></c>
                  <c>TEST07</c>
                  <c>00001</c>
                  <c>I</c>
                  <c>NP</c>
                  <c>+1</c>
                  <c />
                  <c>1</c>
                  <c>RACCT</c>
                </r>
              </Data>
            </Configuration>
          </TextInput>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="Output">
          <RecordInfo>
            <Field name="MANDT" size="3" source="TextInput:" type="String" />
            <Field name="SETCLASS" size="4" source="TextInput:" type="String" />
            <Field name="SUBCLASS" size="1" source="TextInput:" type="String" />
            <Field name="SETNAME" size="6" source="TextInput:" type="String" />
            <Field name="LINEID" size="5" source="TextInput:" type="String" />
            <Field name="VALSIGN" size="1" source="TextInput:" type="String" />
            <Field name="VALOPTION" size="2" source="TextInput:" type="String" />
            <Field name="VALFROM" size="2" source="TextInput:" type="String" />
            <Field name="VALTO" size="2" source="TextInput:" type="String" />
            <Field name="SEQNR" source="TextInput:" type="Byte" />
            <Field name="FIELDNAME" size="5" source="TextInput:" type="String" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroInput" />
    </Node>
    <Node ToolID="2">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Formula.Formula">
        <Position x="198" y="126" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <FormulaFields>
            <FormulaField expression="IF   [VALOPTION] = 'EQ'&#xA;THEN &quot;[[FIELDNAME]] = '[VALFROM]'&quot;&#xA;ELSE &quot;&quot;&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'NE'&#xA;THEN &quot;[[FIELDNAME]] != '[VALFROM]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'GT'&#xA;THEN &quot;[[FIELDNAME]] &gt; '[VALFROM]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'GE'&#xA;THEN &quot;[[FIELDNAME]] &gt;= '[VALFROM]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'LT'&#xA;THEN &quot;[[FIELDNAME]] &lt; '[VALFROM]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'LE'&#xA;THEN &quot;[[FIELDNAME]] &lt;= '[VALFROM]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'BT'&#xA;THEN &quot;[[FIELDNAME]] &gt;= '[VALFROM]' &amp;&amp; [[FIELDNAME]] &lt;= '[VALTO]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'NB'&#xA;THEN &quot;[[FIELDNAME]] &lt; '[VALFROM]' || [[FIELDNAME]] &gt; '[VALTO]'&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'CP'&#xA;THEN &quot;REGEX_Match([[FIELDNAME]], '^[VALFROM]$')&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] = 'NP'&#xA;THEN &quot;!REGEX_Match([[FIELDNAME]], '^[VALFROM]$')&quot;&#xA;ELSE [Filter]&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
          </FormulaFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>Filter</AnnotationText>
          <DefaultAnnotationText>Filter = IF   [VALOPTION] = 'EQ'
THEN "[[FIELDNAME]] = '[VALFROM]'"
ELSE ""
ENDI...</DefaultAnnotationText>
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="Output">
          <RecordInfo>
            <Field name="MANDT" size="3" source="TextInput:" type="String" />
            <Field name="SETCLASS" size="4" source="TextInput:" type="String" />
            <Field name="SUBCLASS" size="1" source="TextInput:" type="String" />
            <Field name="SETNAME" size="6" source="TextInput:" type="String" />
            <Field name="LINEID" size="5" source="TextInput:" type="String" />
            <Field name="VALSIGN" size="1" source="TextInput:" type="String" />
            <Field name="VALOPTION" size="2" source="TextInput:" type="String" />
            <Field name="VALFROM" size="2" source="TextInput:" type="String" />
            <Field name="VALTO" size="2" source="TextInput:" type="String" />
            <Field name="SEQNR" source="TextInput:" type="Byte" />
            <Field name="FIELDNAME" size="5" source="TextInput:" type="String" />
            <Field name="Filter" size="1073741823" source="Formula: IF   [VALOPTION] = 'NP'&#xA;THEN &quot;!REGEX_Match([[FIELDNAME]], '^[VALFROM]$')&quot;&#xA;ELSE [Filter]&#xA;ENDIF" type="V_WString" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFormula" />
    </Node>
    <Node ToolID="3">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Formula.Formula">
        <Position x="294" y="126" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <FormulaFields>
            <FormulaField expression="Replace([Filter], &quot;[FIELDNAME]&quot;, [FIELDNAME])" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="Replace([Filter], &quot;[VALTO]&quot;, [VALTO])" field="Filter" size="1073741823" type="V_WString" />
            <FormulaField expression="IF   [VALOPTION] in ('CP','NP')&#xA;THEN Replace([Filter], '[VALFROM]', Replace(Replace([VALFROM], &quot;+&quot;, &quot;.&quot;), &quot;*&quot;, &quot;.*&quot;))&#xA;ELSE REPLACE([Filter], '[VALFROM]', [VALFROM])&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
          </FormulaFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>Filter</AnnotationText>
          <DefaultAnnotationText>Filter = Replace([Filter], "[FIELDNAME]", [FIELDNAME])
Filter = Replace([Filter]...</DefaultAnnotationText>
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="Output">
          <RecordInfo>
            <Field name="MANDT" size="3" source="TextInput:" type="String" />
            <Field name="SETCLASS" size="4" source="TextInput:" type="String" />
            <Field name="SUBCLASS" size="1" source="TextInput:" type="String" />
            <Field name="SETNAME" size="6" source="TextInput:" type="String" />
            <Field name="LINEID" size="5" source="TextInput:" type="String" />
            <Field name="VALSIGN" size="1" source="TextInput:" type="String" />
            <Field name="VALOPTION" size="2" source="TextInput:" type="String" />
            <Field name="VALFROM" size="2" source="TextInput:" type="String" />
            <Field name="VALTO" size="2" source="TextInput:" type="String" />
            <Field name="SEQNR" source="TextInput:" type="Byte" />
            <Field name="FIELDNAME" size="5" source="TextInput:" type="String" />
            <Field name="Filter" size="1073741823" source="Formula: IF   [VALOPTION] in ('CP','NP')&#xA;THEN Replace([Filter], '[VALFROM]', Replace(Replace([VALFROM], &quot;+&quot;, &quot;.&quot;), &quot;*&quot;, &quot;.*&quot;))&#xA;ELSE REPLACE([Filter], '[VALFROM]', [VALFROM])&#xA;ENDIF" type="V_WString" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFormula" />
    </Node>
    <Node ToolID="4">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Formula.Formula">
        <Position x="390" y="126" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <FormulaFields>
            <FormulaField expression="IF   [VALSIGN] = 'E'&#xA;THEN &quot;!(&quot; + [Filter] + &quot;)&quot;&#xA;ELSE &quot;(&quot; + [Filter] + &quot;)&quot;&#xA;ENDIF" field="Filter" size="1073741823" type="V_WString" />
          </FormulaFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>Filter</AnnotationText>
          <DefaultAnnotationText>Filter = IF   [VALSIGN] = 'E'
THEN "!(" + [Filter] + ")"
ELSE "(" + [Filter] + "...</DefaultAnnotationText>
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="Output">
          <RecordInfo>
            <Field name="MANDT" size="3" source="TextInput:" type="String" />
            <Field name="SETCLASS" size="4" source="TextInput:" type="String" />
            <Field name="SUBCLASS" size="1" source="TextInput:" type="String" />
            <Field name="SETNAME" size="6" source="TextInput:" type="String" />
            <Field name="LINEID" size="5" source="TextInput:" type="String" />
            <Field name="VALSIGN" size="1" source="TextInput:" type="String" />
            <Field name="VALOPTION" size="2" source="TextInput:" type="String" />
            <Field name="VALFROM" size="2" source="TextInput:" type="String" />
            <Field name="VALTO" size="2" source="TextInput:" type="String" />
            <Field name="SEQNR" source="TextInput:" type="Byte" />
            <Field name="FIELDNAME" size="5" source="TextInput:" type="String" />
            <Field name="Filter" size="1073741823" source="Formula: IF   [VALSIGN] = 'E'&#xA;THEN &quot;!(&quot; + [Filter] + &quot;)&quot;&#xA;ELSE &quot;(&quot; + [Filter] + &quot;)&quot;&#xA;ENDIF" type="V_WString" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFormula" />
    </Node>
    <Node ToolID="5">
      <GuiSettings Plugin="AlteryxGuiToolkit.Questions.Tab.Tab">
        <Position x="0" y="0" width="59" height="59" />
      </GuiSettings>
      <Properties>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </Node>
    <Node ToolID="6">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroOutput.MacroOutput">
        <Position x="486" y="126" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Name>Output</Name>
          <Abbrev />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroOutput" />
    </Node>
  </Nodes>
  <Connections>
    <Connection>
      <Origin ToolID="1" Connection="Output" />
      <Destination ToolID="2" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="2" Connection="Output" />
      <Destination ToolID="3" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="3" Connection="Output" />
      <Destination ToolID="4" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="4" Connection="Output" />
      <Destination ToolID="6" Connection="Input" />
    </Connection>
  </Connections>
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <Constants>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>Macro Output (6)</Name>
        <Value />
        <IsNumeric value="False" />
      </Constant>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>Macro Input (1)</Name>
        <Value />
        <IsNumeric value="False" />
      </Constant>
    </Constants>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name>Calculate Filter Expression</Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="True" />
    </Events>
    <RuntimeProperties>
      <Actions />
      <Questions>
        <Question>
          <Type>Tab</Type>
          <Description>Questions</Description>
          <Name>Tab (5)</Name>
          <ToolId value="5" />
          <Questions>
            <Question>
              <Type>MacroOutput</Type>
              <Description>Macro Output (6)</Description>
              <Name>Macro Output (6)</Name>
              <ToolId value="6" />
            </Question>
            <Question>
              <Type>MacroInput</Type>
              <Description />
              <Name>Macro Input (1)</Name>
              <ToolId value="1" />
            </Question>
          </Questions>
        </Question>
      </Questions>
      <ModuleType>Macro</ModuleType>
      <MacroCustomHelp value="False" />
      <MacroDynamicOutputFields value="False" />
      <MacroImage>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
</MacroImage>
      <MacroInputs />
      <MacroOutputs />
      <Wiz_CustomHelp value="False" />
      <Wiz_CustomGraphic value="False" />
      <Wiz_ShowOutput value="True" />
      <Wiz_OpenOutputTools>
        <Tool ToolId="6" Selected="True" />
      </Wiz_OpenOutputTools>
      <Wiz_OutputMessage />
      <Wiz_NoOutputFilesMessage />
      <Wiz_ChainRunWizard />
    </RuntimeProperties>
  </Properties>
</AlteryxDocument>`

var tagWithSets = `<?xml version="1.0"?>
<AlteryxDocument yxmdVer="2019.3">
  <Nodes>
    <Node ToolID="1">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroInput.MacroInput">
        <Position x="222" y="270" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <UseFileInput value="False" />
          <Name>Input1</Name>
          <Abbrev />
          <ShowFieldMap value="False" />
          <Optional value="False" />
          <TextInput>
            <Configuration>
              <NumRows value="11" />
              <Fields>
                <Field name="RCLNT" />
                <Field name="RYEAR" />
                <Field name="DOCNR" />
                <Field name="RLDNR" />
                <Field name="RBUKRS" />
                <Field name="DOCLN" />
                <Field name="RACCT" />
                <Field name="HSL" />
              </Fields>
              <Data>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000001</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000001</c>
                  <c>A1</c>
                  <c>12.00</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000001</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000002</c>
                  <c>G7</c>
                  <c>-12.00</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000002</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000001</c>
                  <c>A3</c>
                  <c>22.10</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000002</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000002</c>
                  <c>A8</c>
                  <c>1.30</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000002</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000003</c>
                  <c>B2</c>
                  <c>-23.40</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000003</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000001</c>
                  <c>C2</c>
                  <c>14.20</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000003</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000002</c>
                  <c>I4</c>
                  <c>-14.20</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000004</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000001</c>
                  <c>J1</c>
                  <c>10.20</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000004</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000002</c>
                  <c>Z8</c>
                  <c>3.50</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000004</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000003</c>
                  <c>K9</c>
                  <c>-5.00</c>
                </r>
                <r>
                  <c>010</c>
                  <c>2019</c>
                  <c>0000000004</c>
                  <c>0L</c>
                  <c>0001</c>
                  <c>000004</c>
                  <c>M3</c>
                  <c>-8.70</c>
                </r>
              </Data>
            </Configuration>
          </TextInput>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>FAGLFLEXA</AnnotationText>
          <DefaultAnnotationText />
          <Left value="True" />
        </Annotation>
        <MetaInfo connection="Output">
          <RecordInfo>
            <Field name="RCLNT" size="3" source="TextInput:" type="String" />
            <Field name="RYEAR" source="TextInput:" type="Int16" />
            <Field name="DOCNR" size="10" source="TextInput:" type="String" />
            <Field name="RLDNR" size="2" source="TextInput:" type="String" />
            <Field name="RBUKRS" size="4" source="TextInput:" type="String" />
            <Field name="DOCLN" size="6" source="TextInput:" type="String" />
            <Field name="RACCT" size="2" source="TextInput:" type="String" />
            <Field name="HSL" source="TextInput:" type="Double" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroInput" />
    </Node>
    <Node ToolID="2">
      <GuiSettings Plugin="AlteryxGuiToolkit.Questions.Tab.Tab">
        <Position x="0" y="0" width="59" height="59" />
      </GuiSettings>
      <Properties>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </Node>
    <Node ToolID="3">
      <GuiSettings Plugin="AlteryxGuiToolkit.Questions.ControlParam.ControlParam">
        <Position x="162" y="66" width="59" height="59" />
      </GuiSettings>
      <Properties>
        <Configuration />
        <Annotation DisplayMode="0">
          <Name>Filter</Name>
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </Node>
    <Node ToolID="4">
      <GuiSettings Plugin="AlteryxGuiToolkit.Questions.ControlParam.ControlParam">
        <Position x="162" y="150" width="59" height="59" />
      </GuiSettings>
      <Properties>
        <Configuration />
        <Annotation DisplayMode="0">
          <Name>SetName</Name>
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </Node>
    <Node ToolID="5">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Filter.Filter">
        <Position x="318" y="270" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Expression>%Question.ControlParam.Filter%</Expression>
          <Mode>Custom</Mode>
          <Simple>
            <Operator>=</Operator>
            <Operands>
              <IgnoreTimeInDateTime>True</IgnoreTimeInDateTime>
              <DateType>fixed</DateType>
              <PeriodDate>2019-10-18 05:21:35</PeriodDate>
              <PeriodType>
              </PeriodType>
              <PeriodCount>0</PeriodCount>
              <StartDate>2019-10-18 05:21:35</StartDate>
              <EndDate>2019-10-18 05:21:35</EndDate>
              <Operand>
              </Operand>
            </Operands>
          </Simple>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>Filter</AnnotationText>
          <DefaultAnnotationText>%Question.ControlParam.Filter%</DefaultAnnotationText>
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="True">
          <RecordInfo>
            <Field name="RCLNT" size="3" source="TextInput:" type="String" />
            <Field name="RYEAR" source="TextInput:" type="Int16" />
            <Field name="DOCNR" size="10" source="TextInput:" type="String" />
            <Field name="RLDNR" size="2" source="TextInput:" type="String" />
            <Field name="RBUKRS" size="4" source="TextInput:" type="String" />
            <Field name="DOCLN" size="6" source="TextInput:" type="String" />
            <Field name="RACCT" size="2" source="TextInput:" type="String" />
            <Field name="HSL" source="TextInput:" type="Double" />
          </RecordInfo>
        </MetaInfo>
        <MetaInfo connection="False">
          <RecordInfo>
            <Field name="RCLNT" size="3" source="TextInput:" type="String" />
            <Field name="RYEAR" source="TextInput:" type="Int16" />
            <Field name="DOCNR" size="10" source="TextInput:" type="String" />
            <Field name="RLDNR" size="2" source="TextInput:" type="String" />
            <Field name="RBUKRS" size="4" source="TextInput:" type="String" />
            <Field name="DOCLN" size="6" source="TextInput:" type="String" />
            <Field name="RACCT" size="2" source="TextInput:" type="String" />
            <Field name="HSL" source="TextInput:" type="Double" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFilter" />
    </Node>
    <Node ToolID="6">
      <GuiSettings Plugin="AlteryxBasePluginsGui.Formula.Formula">
        <Position x="414" y="270" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <FormulaFields>
            <FormulaField expression="&quot;%Question.ControlParam.SetName%&quot;" field="Set Name" size="1073741823" type="V_WString" />
          </FormulaFields>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <AnnotationText>Set Name</AnnotationText>
          <DefaultAnnotationText><![CDATA[Set Name = "%Question.ControlParam.SetName%"
]]></DefaultAnnotationText>
          <Left value="False" />
        </Annotation>
        <MetaInfo connection="Output">
          <RecordInfo>
            <Field name="RCLNT" size="3" source="TextInput:" type="String" />
            <Field name="RYEAR" source="TextInput:" type="Int16" />
            <Field name="DOCNR" size="10" source="TextInput:" type="String" />
            <Field name="RLDNR" size="2" source="TextInput:" type="String" />
            <Field name="RBUKRS" size="4" source="TextInput:" type="String" />
            <Field name="DOCLN" size="6" source="TextInput:" type="String" />
            <Field name="RACCT" size="2" source="TextInput:" type="String" />
            <Field name="HSL" source="TextInput:" type="Double" />
            <Field name="Set Name" size="1073741823" source="Formula: &quot;TEST1&quot;" type="V_WString" />
          </RecordInfo>
        </MetaInfo>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFormula" />
    </Node>
    <Node ToolID="7">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroOutput.MacroOutput">
        <Position x="510" y="270" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Name>Output7</Name>
          <Abbrev />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroOutput" />
    </Node>
  </Nodes>
  <Connections>
    <Connection>
      <Origin ToolID="1" Connection="Output" />
      <Destination ToolID="5" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="5" Connection="True" />
      <Destination ToolID="6" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="6" Connection="Output" />
      <Destination ToolID="7" Connection="Input" />
    </Connection>
  </Connections>
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <Constants>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>ControlParam.Filter</Name>
        <Value>[RACCT] = 'A1'</Value>
        <IsNumeric value="False" />
      </Constant>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>ControlParam.SetName</Name>
        <Value>TEST1</Value>
        <IsNumeric value="False" />
      </Constant>
    </Constants>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name>Tag with Sets</Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="False" />
    </Events>
    <RuntimeProperties>
      <Actions />
      <Questions>
        <Question>
          <Type>Tab</Type>
          <Description>Questions</Description>
          <Name>Tab (2)</Name>
          <ToolId value="2" />
          <Questions>
            <Question>
              <Type>MacroInput</Type>
              <Description />
              <Name>Macro Input (1)</Name>
              <ToolId value="1" />
            </Question>
            <Question>
              <Type>ControlParam</Type>
              <Description>Filter</Description>
              <Name>Filter</Name>
              <ToolId value="3" />
            </Question>
            <Question>
              <Type>ControlParam</Type>
              <Description>Set Name</Description>
              <Name>SetName</Name>
              <ToolId value="4" />
            </Question>
          </Questions>
        </Question>
      </Questions>
      <ModuleType>Macro</ModuleType>
      <MacroCustomHelp value="False" />
      <MacroDynamicOutputFields value="False" />
      <MacroImageStd value="39" />
      <MacroInputs />
      <MacroOutputs />
      <BatchMacro>
        <OutputMode>AllSame</OutputMode>
        <ControlParams>
          <ControlParam>
            <Name>Filter</Name>
            <Description>Filter</Description>
          </ControlParam>
          <ControlParam>
            <Name>SetName</Name>
            <Description>Set Name</Description>
          </ControlParam>
        </ControlParams>
      </BatchMacro>
      <Wiz_CustomHelp value="False" />
      <Wiz_CustomGraphic value="False" />
      <Wiz_ShowOutput value="True" />
      <Wiz_OpenOutputTools>
        <Tool ToolId="7" Selected="False" />
      </Wiz_OpenOutputTools>
      <Wiz_OutputMessage />
      <Wiz_NoOutputFilesMessage />
      <Wiz_ChainRunWizard />
    </RuntimeProperties>
  </Properties>
</AlteryxDocument>`

var emptyYxwz = `<?xml version="1.0"?>
<AlteryxDocument yxmdVer="2019.4">
  <Nodes>
    <Node ToolID="1">
      <GuiSettings Plugin="AlteryxGuiToolkit.Questions.Tab.Tab">
        <Position x="0" y="0" width="59" height="59" />
      </GuiSettings>
      <Properties>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </Node>
  </Nodes>
  <Connections />
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name>New Workflow1</Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="True" />
    </Events>
    <RuntimeProperties>
      <Actions />
      <Questions>
        <Question>
          <Type>Tab</Type>
          <Description>Questions</Description>
          <Name>Tab (1)</Name>
          <ToolId value="1" />
          <Questions />
        </Question>
      </Questions>
      <ModuleType>Wizard</ModuleType>
      <MacroCustomHelp value="False" />
      <MacroDynamicOutputFields value="False" />
      <MacroImageStd value="39" />
      <MacroInputs />
      <MacroOutputs />
      <Wiz_CustomHelp value="False" />
      <Wiz_CustomGraphic value="False" />
      <Wiz_ShowOutput value="True" />
      <Wiz_OpenOutputTools />
      <Wiz_OutputMessage />
      <Wiz_NoOutputFilesMessage />
      <Wiz_ChainRunWizard />
    </RuntimeProperties>
  </Properties>
</AlteryxDocument>`

var multiInOutMacro = `<?xml version="1.0"?>
<AlteryxDocument yxmdVer="2019.4">
  <Nodes>
    <Node ToolID="1">
      <GuiSettings Plugin="AlteryxGuiToolkit.Questions.Tab.Tab">
        <Position x="0" y="0" width="59" height="59" />
      </GuiSettings>
      <Properties>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </Node>
    <Node ToolID="2">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroInput.MacroInput">
        <Position x="114" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <UseFileInput value="False" />
          <Name>Input2</Name>
          <Abbrev />
          <ShowFieldMap value="False" />
          <Optional value="False" />
          <TextInput>
            <Configuration>
              <NumRows value="0" />
              <Fields />
              <Data />
            </Configuration>
          </TextInput>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="True" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroInput" />
    </Node>
    <Node ToolID="3">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroInput.MacroInput">
        <Position x="114" y="150" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <UseFileInput value="False" />
          <Name>Input3</Name>
          <Abbrev />
          <ShowFieldMap value="False" />
          <Optional value="False" />
          <TextInput>
            <Configuration>
              <NumRows value="0" />
              <Fields />
              <Data />
            </Configuration>
          </TextInput>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="True" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroInput" />
    </Node>
    <Node ToolID="4">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroOutput.MacroOutput">
        <Position x="222" y="54" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Name>Output4</Name>
          <Abbrev />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroOutput" />
    </Node>
    <Node ToolID="5">
      <GuiSettings Plugin="AlteryxBasePluginsGui.MacroOutput.MacroOutput">
        <Position x="222" y="150" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Name>Output5</Name>
          <Abbrev />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxMacroOutput" />
    </Node>
  </Nodes>
  <Connections>
    <Connection>
      <Origin ToolID="2" Connection="Output" />
      <Destination ToolID="4" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="3" Connection="Output" />
      <Destination ToolID="5" Connection="Input" />
    </Connection>
  </Connections>
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <Constants>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>Macro Input (2)</Name>
        <Value />
        <IsNumeric value="False" />
      </Constant>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>Macro Input (3)</Name>
        <Value />
        <IsNumeric value="False" />
      </Constant>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>Macro Output (4)</Name>
        <Value />
        <IsNumeric value="False" />
      </Constant>
      <Constant>
        <Namespace>Question</Namespace>
        <Name>Macro Output (5)</Name>
        <Value />
        <IsNumeric value="False" />
      </Constant>
    </Constants>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name>MultiInOut</Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="True" />
    </Events>
    <RuntimeProperties>
      <Actions />
      <Questions>
        <Question>
          <Type>Tab</Type>
          <Description>Questions</Description>
          <Name>Tab (1)</Name>
          <ToolId value="1" />
          <Questions>
            <Question>
              <Type>MacroInput</Type>
              <Description>Macro Input (2)</Description>
              <Name>Macro Input (2)</Name>
              <ToolId value="2" />
            </Question>
            <Question>
              <Type>MacroInput</Type>
              <Description>Macro Input (3)</Description>
              <Name>Macro Input (3)</Name>
              <ToolId value="3" />
            </Question>
            <Question>
              <Type>MacroOutput</Type>
              <Description>Macro Output (4)</Description>
              <Name>Macro Output (4)</Name>
              <ToolId value="4" />
            </Question>
            <Question>
              <Type>MacroOutput</Type>
              <Description>Macro Output (5)</Description>
              <Name>Macro Output (5)</Name>
              <ToolId value="5" />
            </Question>
          </Questions>
        </Question>
      </Questions>
      <ModuleType>Macro</ModuleType>
      <MacroCustomHelp value="False" />
      <MacroDynamicOutputFields value="False" />
      <MacroImageStd value="39" />
      <MacroInputs />
      <MacroOutputs />
      <Wiz_CustomHelp value="False" />
      <Wiz_CustomGraphic value="False" />
      <Wiz_ShowOutput value="True" />
      <Wiz_OpenOutputTools>
        <Tool ToolId="2" Selected="True" />
        <Tool ToolId="3" Selected="True" />
        <Tool ToolId="4" Selected="True" />
        <Tool ToolId="5" Selected="True" />
      </Wiz_OpenOutputTools>
      <Wiz_OutputMessage />
      <Wiz_NoOutputFilesMessage />
      <Wiz_ChainRunWizard />
    </RuntimeProperties>
  </Properties>
</AlteryxDocument>`

var multiInOutWorkflow = `<?xml version="1.0"?>
<AlteryxDocument yxmdVer="2019.4">
  <Nodes>
    <Node ToolID="1">
      <GuiSettings Plugin="AlteryxBasePluginsGui.TextInput.TextInput">
        <Position x="162" y="114" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <NumRows value="1" />
          <Fields>
            <Field name="Field1" />
          </Fields>
          <Data>
            <r>
              <c>A</c>
            </r>
          </Data>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxTextInput" />
    </Node>
    <Node ToolID="2">
      <GuiSettings Plugin="AlteryxBasePluginsGui.TextInput.TextInput">
        <Position x="162" y="186" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <NumRows value="1" />
          <Fields>
            <Field name="Field1" />
          </Fields>
          <Data>
            <r>
              <c>B</c>
            </r>
          </Data>
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxTextInput" />
    </Node>
    <Node ToolID="6">
      <GuiSettings>
        <Position x="258" y="150" />
      </GuiSettings>
      <Properties>
        <Configuration />
        <Annotation DisplayMode="0">
          <Name>MultiInOut (3)</Name>
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings Macro="MultiInOut.yxmc" />
    </Node>
    <Node ToolID="4">
      <GuiSettings Plugin="AlteryxBasePluginsGui.BrowseV2.BrowseV2">
        <Position x="354" y="114" />
      </GuiSettings>
      <Properties>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxBrowseV2" />
    </Node>
    <Node ToolID="5">
      <GuiSettings Plugin="AlteryxBasePluginsGui.BrowseV2.BrowseV2">
        <Position x="354" y="186" />
      </GuiSettings>
      <Properties>
        <Configuration />
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxBrowseV2" />
    </Node>
  </Nodes>
  <Connections>
    <Connection>
      <Origin ToolID="1" Connection="Output" />
      <Destination ToolID="6" Connection="Input2" />
    </Connection>
    <Connection>
      <Origin ToolID="2" Connection="Output" />
      <Destination ToolID="6" Connection="Input3" />
    </Connection>
    <Connection>
      <Origin ToolID="6" Connection="Output4" />
      <Destination ToolID="4" Connection="Input" />
    </Connection>
    <Connection>
      <Origin ToolID="6" Connection="Output5" />
      <Destination ToolID="5" Connection="Input" />
    </Connection>
  </Connections>
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name>MultiInOut</Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="True" />
    </Events>
  </Properties>
</AlteryxDocument>`
