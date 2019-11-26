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

const OnuPowerSheddingClassId ClassID = ClassID(133)

var onupowersheddingBME *ManagedEntityDefinition

// OnuPowerShedding (class ID #133)
//	This ME models the ONU's ability to shed services when the ONU goes into battery operation mode
//	after AC power failure. Shedding classes are defined in the following table, which may span
//	multiple circuit pack types. This feature works in conjunction with the power shed override
//	attribute of the circuit pack ME, which can selectively prevent power shedding of priority
//	ports.
//
//	An ONU that supports power shedding automatically creates an instance of this ME.
//
//	The following table defines the binding of shedding class and PPTP type. The coding is taken
//	from Table 9.1.5-1. In the case of hybrid circuit pack types, multiple shedding classes may
//	affect a circuit pack if the hardware is capable of partial power shedding.
//
//	An ONU may choose to model its ports with the port-mapping package of clause 9.1.8, rather than
//	with real or virtual circuit packs. In this case, power shedding pertains to individual PPTPs
//	(listed in column 2 of the table).
//
//	Relationships
//		One instance of this ME is associated with the ONU ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only
//			one instance, number 0. (R) (mandatory) (2 bytes)
//
//		Restore Power Timer Reset Interval
//			Restore power timer reset interval: The time delay, in seconds, before resetting the power-
//			shedding timers after full power restoration. Upon ME instantiation, the ONU sets this attribute
//			to 0. (R, W) (mandatory) (2 bytes)
//
//		Data Class Shedding Interval
//			Data class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Voice Class Shedding Interval
//			Voice class shedding interval: This attribute only pertains to voice services that terminate on
//			the ONU and are under the management control of the OMCI. (R, W) (mandatory) (2 bytes)
//
//		Video Overlay Class Shedding Interval
//			Video overlay class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Video Return Class Shedding Interval
//			Video return class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Digital Subscriber Line Class Shedding Interval
//			Digital subscriber line (DSL) class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Atm Class Shedding Interval
//			ATM class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Ces Class Shedding Interval
//			CES class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Frame Class Shedding Interval
//			Frame class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Sdh_Sonet Class Shedding Interval
//			Sdh-sonet class shedding interval:	(R, W) (mandatory) (2 bytes)
//
//		Shedding Status
//			The ONU sets each bit to 1 when power shedding is active, and clears it to 0 when the service is
//			restored. (R) (optional) (2 bytes)
//
type OnuPowerShedding struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	onupowersheddingBME = &ManagedEntityDefinition{
		Name:    "OnuPowerShedding",
		ClassID: 133,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFE0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  Uint16Field("RestorePowerTimerResetInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  Uint16Field("DataClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
			3:  Uint16Field("VoiceClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  Uint16Field("VideoOverlayClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  Uint16Field("VideoReturnClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  Uint16Field("DigitalSubscriberLineClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  Uint16Field("AtmClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  Uint16Field("CesClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  Uint16Field("FrameClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 9),
			10: Uint16Field("SdhSonetClassSheddingInterval", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 10),
			11: Uint16Field("SheddingStatus", 0, mapset.NewSetWith(Read), true, false, true, false, 11),
		},
	}
}

// NewOnuPowerShedding (class ID 133 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewOnuPowerShedding(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(onupowersheddingBME, params...)
}
