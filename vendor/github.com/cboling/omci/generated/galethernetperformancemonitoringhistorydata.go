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

const GalEthernetPerformanceMonitoringHistoryDataClassId ClassID = ClassID(276)

var galethernetperformancemonitoringhistorydataBME *ManagedEntityDefinition

// GalEthernetPerformanceMonitoringHistoryData (class ID #276)
//	This ME collects PM data associated with a GEM IW TP when the GEM layer supports an Ethernet
//	service. Instances of this ME are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of the GEM IW TP ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the GEM IW TP. (R, setbycreate)
//			(mandatory) (2 bytes)
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
//		Discarded Frames
//			Discarded frames: This attribute counts the number of downstream GEM frames discarded for any
//			reason [erroneous frame check sequence (FCS), too long length, buffer overflow, etc.]. (R)
//			(mandatory) (4 bytes)
//
type GalEthernetPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	galethernetperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "GalEthernetPerformanceMonitoringHistoryData",
		ClassID: 276,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XE000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("DiscardedFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
		},
	}
}

// NewGalEthernetPerformanceMonitoringHistoryData (class ID 276 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewGalEthernetPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(galethernetperformancemonitoringhistorydataBME, params...)
}