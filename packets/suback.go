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
	"bytes"
	"fmt"
	"io"
)

// SubackPacket is an internal representation of the fields of the
// Suback MQTT packet
type SubackPacket struct {
	FixedHeader
	MessageID   uint16
	ReturnCodes []byte
}

func (sa *SubackPacket) String() string {
	return fmt.Sprintf("%s MessageID: %d", sa.FixedHeader, sa.MessageID)
}

func (sa *SubackPacket) Write(w io.Writer) error {
	var body bytes.Buffer
	var err error
	body.Write(encodeUint16(sa.MessageID))
	body.Write(sa.ReturnCodes)
	sa.FixedHeader.RemainingLength = body.Len()
	packet := sa.FixedHeader.pack()
	packet.Write(body.Bytes())
	_, err = packet.WriteTo(w)

	return err
}

// Unpack decodes the details of a ControlPacket after the fixed
// header has been read
func (sa *SubackPacket) Unpack(b io.Reader) error {
	var qosBuffer bytes.Buffer
	var err error
	sa.MessageID, err = decodeUint16(b)
	if err != nil {
		return err
	}

	_, err = qosBuffer.ReadFrom(b)
	if err != nil {
		return err
	}
	sa.ReturnCodes = qosBuffer.Bytes()

	return nil
}

// Details returns a Details struct containing the Qos and
// MessageID of this ControlPacket
func (sa *SubackPacket) Details() Details {
	return Details{Qos: 0, MessageID: sa.MessageID}
}

// Copy creates a deep copy of the SubackPacket
func (sa *SubackPacket) Copy() ControlPacket {
	cp := NewControlPacket(Suback).(*SubackPacket)

	*cp = *sa

	if len(sa.ReturnCodes) > 0 {
		cp.ReturnCodes = make([]byte, len(sa.ReturnCodes))
		copy(cp.ReturnCodes, sa.ReturnCodes)
	}

	return cp
}
