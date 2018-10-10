// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

@file:Suppress("UnusedImport", "unused")

package proto.fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import proto.*

// Fast Binary Encoding proto final sender class
@Suppress("MemberVisibilityCanBePrivate", "PropertyName")
open class FinalSender : fbe.Sender
{
    // Sender models accessors
    val OrderModel: OrderFinalModel
    val BalanceModel: BalanceFinalModel
    val AccountModel: AccountFinalModel

    constructor()
    {
        final = true
        OrderModel = OrderFinalModel(buffer)
        BalanceModel = BalanceFinalModel(buffer)
        AccountModel = AccountFinalModel(buffer)
    }

    constructor(buffer: Buffer) : super(buffer)
    {
        final = true
        OrderModel = OrderFinalModel(buffer)
        BalanceModel = BalanceFinalModel(buffer)
        AccountModel = AccountFinalModel(buffer)
    }

    fun send(value: proto.Order): Long
    {
        // Serialize the value into the FBE stream
        val serialized = OrderModel.serialize(value)
        assert(serialized > 0) { "proto.Order serialization failed!" }
        assert(OrderModel.verify()) { "proto.Order validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: proto.Balance): Long
    {
        // Serialize the value into the FBE stream
        val serialized = BalanceModel.serialize(value)
        assert(serialized > 0) { "proto.Balance serialization failed!" }
        assert(BalanceModel.verify()) { "proto.Balance validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: proto.Account): Long
    {
        // Serialize the value into the FBE stream
        val serialized = AccountModel.serialize(value)
        assert(serialized > 0) { "proto.Account serialization failed!" }
        assert(AccountModel.verify()) { "proto.Account validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }

    // Send message handler
    override fun onSend(buffer: ByteArray, offset: Long, size: Long): Long { throw UnsupportedOperationException("proto.fbe.Sender.onSend() not implemented!") }
}
