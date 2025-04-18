/*
 * Copyright (c) 2021 IBM Corp and others.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * and Eclipse Distribution License v1.0 which accompany this distribution.
 *
 * The Eclipse Public License is available at
 *    https://www.eclipse.org/legal/epl-2.0/
 * and the Eclipse Distribution License is available at
 *   http://www.eclipse.org/org/documents/edl-v10.php.
 *
 * Contributors:
 *    Allan Stockdill-Mander
 */

package packets

import (
	"io"
)

// DisconnectPacket is an internal representation of the fields of the
// Disconnect MQTT packet
type DisconnectPacket struct {
	FixedHeader
}

func (d *DisconnectPacket) String() string {
	return d.FixedHeader.String()
}

func (d *DisconnectPacket) Write(w io.Writer) error {
	packet := d.FixedHeader.pack()
	_, err := packet.WriteTo(w)

	return err
}

// Unpack decodes the details of a ControlPacket after the fixed
// header has been read
func (d *DisconnectPacket) Unpack(b io.Reader) error {
	return nil
}

// Details returns a Details struct containing the Qos and
// MessageID of this ControlPacket
func (d *DisconnectPacket) Details() Details {
	return Details{Qos: 0, MessageID: 0}
}

// Copy creates a deep copy of the DisconnectPacket
func (d *DisconnectPacket) Copy() ControlPacket {
	cp := NewControlPacket(Disconnect).(*DisconnectPacket)

	*cp = *d

	return cp
}
