/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import "github.com/deckarep/golang-set"

const VpPerformanceMonitoringHistoryDataClassId ClassID = ClassID(62)

var vpperformancemonitoringhistorydataBME *ManagedEntityDefinition

// VpPerformanceMonitoringHistoryData (class ID #62)
//	This ME collects PM data associated with a VP network CTP. Instances of this ME are created and
//	deleted by the OLT.
//
//	Relationships
//		An instance of this ME is associated with an instance of the VP network CTP ME. The performance
//		of upstream ATM flows is reported.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the VP network CTP. (R,
//			setbycreate) (mandatory) (2 bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15 min interval. (R)
//			(mandatory) (1 byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R, W, setbycreate) (mandatory) (2 bytes)
//
//		Lost C 0 1 Cells
//			Lost C = 0 + 1 cells: This attribute counts all cell loss. It cannot distinguish between cells
//			lost because of header bit errors, ATM-level header errors, cell policing, or buffer overflows.
//			It records only loss of information, independent of the priority of the cell. (R) (mandatory)
//			(2 bytes)
//
//		Lost C_= 0 Cells
//			Lost C = 0 cells: This attribute counts loss of high priority cells. It cannot distinguish
//			between cells lost because of header bit errors, ATM-level header errors, cell policing, or
//			buffer overflows. It records only loss of high priority cells. (R) (mandatory) (2 bytes)
//
//		Misinserted Cells
//			Misinserted cells: This attribute counts cells that are misrouted to a monitored VP. (R)
//			(mandatory) (2 bytes)
//
//		Transmitted C_= 0 _ 1 Cells
//			Transmitted C = 0 + 1 cells: This attribute counts cells originated by the transmitting end
//			point (i.e., backward reporting is assumed). (R) (mandatory) (5 bytes)
//
//		Transmitted C_= 0 Cells
//			Transmitted C = 0 cells: This attribute counts high priority cells originated by the
//			transmitting end point (i.e., backward reporting is assumed). (R) (mandatory) (5 bytes)
//
//		Impaired Block
//			Impaired blocks: This severely errored cell block counter is incremented whenever one of the
//			following events takes place: the number of misinserted cells reaches its threshold; the number
//			of bipolar violations reaches its threshold; or the number of lost cells reaches its threshold.
//			Threshold values are based on vendor-operator negotiation. (R) (mandatory) (2 bytes)
//
type VpPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vpperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "VpPerformanceMonitoringHistoryData",
		ClassID: 62,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFF00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint16Field("LostC01Cells", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint16Field("LostC=0Cells", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint16Field("MisinsertedCells", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6: MultiByteField("TransmittedC=01Cells", 5, nil, mapset.NewSetWith(Read), false, false, false, false, 6),
			7: MultiByteField("TransmittedC=0Cells", 5, nil, mapset.NewSetWith(Read), false, false, false, false, 7),
			8: Uint16Field("ImpairedBlock", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
		},
	}
}

// NewVpPerformanceMonitoringHistoryData (class ID 62 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVpPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(vpperformancemonitoringhistorydataBME, params...)
}
