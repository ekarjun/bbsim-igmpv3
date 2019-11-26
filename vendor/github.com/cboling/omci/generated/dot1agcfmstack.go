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

const Dot1AgCfmStackClassId ClassID = ClassID(305)

var dot1agcfmstackBME *ManagedEntityDefinition

// Dot1AgCfmStack (class ID #305)
//	This ME reports the maintenance status of a bridge port at any given time. An ONU that supports
//	[IEEE 802.1ag] functionality automatically creates an instance of the dot1ag CFM stack ME for
//	each MAC bridge or IEEE 802.1p mapper, depending on its provisioning model.
//
//	The dot1ag CFM stack also lists any VLANs and bridge ports against which configuration errors
//	are currently identified. The ONU should reject operations that create configuration errors.
//	However, these errors can arise because of operations on other MEs that are not necessarily
//	possible to detect during CFM configuration.
//
//	Relationships
//		An ONU that supports [IEEE 802.1ag] creates one instance of this ME for each MAC bridge or IEEE
//		802.1p mapper, depending on its provisioning model. It should not create an instance for an
//		IEEE 802.1p mapper that is associated with a MAC bridge.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies an instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the MAC bridge service profile ME
//			or an IEEE 802.1p mapper ME. It is expected that an ONU will implement CFM on bridges or on
//			IEEE 802.1p mappers, but not both. For precision, the reference is disambiguated by the value of
//			the layer 2 type pointer attribute. (R) (mandatory) (2 bytes)
//
//		Layer 2 Type
//			Layer 2 type:	This attribute specifies whether the dot1ag CFM stack is associated with a MAC
//			bridge service profile (value 0) or an IEEE 802.1p mapper (value 1). (R) (mandatory) (1 byte)
//
//		Mp Status Table
//			(R) (mandatory) (18N bytes)
//
//		Configuration Error List Table
//			(R) (mandatory) (5N bytes)
//
type Dot1AgCfmStack struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1agcfmstackBME = &ManagedEntityDefinition{
		Name:    "Dot1AgCfmStack",
		ClassID: 305,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
		),
		AllowedAttributeMask: 0XE000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1: ByteField("Layer2Type", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: TableField("MpStatusTable", TableInfo{nil, 18}, mapset.NewSetWith(Read), false, false, false, 2),
			3: TableField("ConfigurationErrorListTable", TableInfo{nil, 5}, mapset.NewSetWith(Read), true, false, false, 3),
		},
	}
}

// NewDot1AgCfmStack (class ID 305 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1AgCfmStack(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(dot1agcfmstackBME, params...)
}
